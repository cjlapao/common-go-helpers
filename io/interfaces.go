package io

import (
	"io/fs"
	"os"
)

type FileIo interface {
	GetOperatingSystem() OperatingSystem
	FileExists(path string) bool
	DirExists(folderPath string) bool
	CreateDir(folderPath string, mode fs.FileMode) error
	GetExecutionPath() string
	ToOsPath(path string) string
	GetOsPathSeparator() string
	ReadFile(path string) ([]byte, error)
	ReadBufferedFile(path string, from, to int) ([]byte, error)
	WriteFile(path string, data []byte, mode os.FileMode) error
	WriteBufferedFile(path string, data []byte, bufferSize int, mode os.FileMode) error
	ReadDir(path string) ([]fs.DirEntry, error)
	JoinPath(parts ...string) string
	CopyFile(source, destination string) error
	DeleteFile(path string) error
	CopyDir(source, destination string) error
	DeleteDir(path string) error
	Checksum(path string, method ChecksumMethod) (string, error)
	FileInfo(path string) (os.FileInfo, error)
}
