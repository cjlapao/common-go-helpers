package io

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"errors"
	"fmt"
	"hash"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type DefaultFileIo struct{}

func Default() DefaultFileIo {
	return DefaultFileIo{}
}

func (f DefaultFileIo) GetOperatingSystem() OperatingSystem {
	return getOperatingSystem()
}

func (f DefaultFileIo) FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (f DefaultFileIo) DirExists(folderPath string) bool {
	_, err := os.Stat(folderPath)
	return !os.IsNotExist(err)
}

func (f DefaultFileIo) CreateDir(folderPath string, mode os.FileMode) error {
	err := os.Mkdir(folderPath, mode)
	return err
}

func (f DefaultFileIo) GetExecutionPath() string {
	return os.Args[0]
}

func (f DefaultFileIo) GetOsPathSeparator() string {
	switch getOperatingSystem() {
	case WindowsOs:
		return "\\"
	case LinuxOs:
		return "/"
	case MacOs:
		return "/"
	default:
		return "/"
	}
}

func (f DefaultFileIo) ToOsPath(path string) string {
	switch getOperatingSystem() {
	case WindowsOs:
		return strings.ReplaceAll(path, "/", "\\")
	case LinuxOs:
		if strings.ContainsAny(path, ":") {
			pathParts := strings.Split(path, ":")
			path = pathParts[1]
		}
		return strings.ReplaceAll(path, "\\", "/")
	case MacOs:
		if strings.ContainsAny(path, ":") {
			pathParts := strings.Split(path, ":")
			path = pathParts[1]
		}
		return strings.ReplaceAll(path, "\\", "/")
	default:
		return path
	}
}

func (f DefaultFileIo) ReadFile(path string) ([]byte, error) {
	if !f.FileExists(path) {
		return nil, os.ErrNotExist
	}

	data, err := os.ReadFile(path)
	return data, err
}

func (f DefaultFileIo) ReadBufferedFile(path string, from, to int) ([]byte, error) {
	if !f.FileExists(path) {
		return nil, os.ErrNotExist
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if int(stat.Size()) < to || to == 0 {
		to = int(stat.Size())
	}

	buffer := make([]byte, to-from)
	_, err = file.ReadAt(buffer, int64(from))
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func (f DefaultFileIo) WriteFile(path string, data []byte, mode os.FileMode) error {
	err := os.WriteFile(path, data, mode)
	if err != nil {
		return err
	}

	return nil
}

func (f DefaultFileIo) WriteBufferedFile(path string, data []byte, bufferSize int, mode os.FileMode) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := file.Chmod(mode); err != nil {
		return err
	}

	for i := 0; i < len(data); i += bufferSize {
		end := i + bufferSize
		if end > len(data) {
			end = len(data)
		}

		_, err = file.Write(data[i:end])
		if err != nil {
			return err
		}
	}

	return nil
}

func (f DefaultFileIo) ReadDir(path string) ([]fs.DirEntry, error) {
	dir, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	return dir, nil
}

func (f DefaultFileIo) JoinPath(parts ...string) string {
	for i := range parts {
		parts[i] = strings.ReplaceAll(parts[i], "\\", "")
		parts[i] = strings.ReplaceAll(parts[i], "/", "")
	}

	return strings.Join(parts, f.GetOsPathSeparator())
}

func (f DefaultFileIo) CopyFile(source, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	err = destinationFile.Sync()
	if err != nil {
		return err
	}

	si, err := sourceFile.Stat()
	if err != nil {
		return err
	}

	err = os.Chmod(destination, si.Mode())
	if err != nil {
		return err
	}

	return nil
}

func (f DefaultFileIo) DeleteFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}

func (f DefaultFileIo) CopyDir(source, destination string) error {
	sourceInfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	err = os.MkdirAll(destination, sourceInfo.Mode())
	if err != nil {
		return err
	}

	directory, err := os.ReadDir(source)
	if err != nil {
		return err
	}

	for _, file := range directory {
		sourcePath := filepath.Join(source, file.Name())
		destinationPath := filepath.Join(destination, file.Name())

		if file.IsDir() {
			err = f.CopyDir(sourcePath, destinationPath)
			if err != nil {
				return err
			}
		} else {
			err = f.CopyFile(sourcePath, destinationPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (f DefaultFileIo) DeleteDir(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}

	return nil
}

func (f DefaultFileIo) Checksum(path string, method ChecksumMethod) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	var hash hash.Hash
	switch method {
	case ChecksumMD5:
		hash = md5.New()
		_, err := io.Copy(hash, file)
		if err != nil {
			return "", err
		}
	case ChecksumSHA1:
		hash = sha1.New()
		_, err := io.Copy(hash, file)
		if err != nil {
			return "", err
		}
	case ChecksumSHA256:
		hash = sha256.New()
		_, err := io.Copy(hash, file)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("invalid checksum method")
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func (f DefaultFileIo) FileInfo(path string) (os.FileInfo, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	return fileInfo, nil
}
