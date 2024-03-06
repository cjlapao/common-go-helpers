package mock

import (
	"os"

	helpers_io "github.com/cjlapao/common-go-helpers/io"
)

type MockOperation struct {
	Method string
	Func   func(args ...MockFuncArgument) interface{}
	Args   []interface{}
	Result interface{}
}

type MockFuncArgument struct {
	Name  string
	Value interface{}
}

type MockFileIo struct {
	mocks []MockOperation
}

func NewMockFileIo() *MockFileIo {
	return &MockFileIo{
		mocks: []MockOperation{},
	}
}

func GetMockFuncArgumentValue[T any](args []MockFuncArgument, name string) (T, bool) {
	for _, arg := range args {
		if arg.Name == name {
			if value, ok := arg.Value.(T); ok {
				return value, true
			}
		}
	}

	t := interface{}(nil).(T)
	return t, false
}

func (f *MockFileIo) On(op MockOperation) *MockOperation {
	if f.mocks == nil {
		f.mocks = []MockOperation{}
	}

	f.mocks = append(f.mocks, op)
	return &f.mocks[len(f.mocks)-1]
}

func (f MockFileIo) GetOperatingSystem() helpers_io.OperatingSystem {
	for _, op := range f.mocks {
		if op.Method == "GetOperatingSystem" {
			if op.Func != nil {
				result, ok := op.Func().(helpers_io.OperatingSystem)
				if !ok {
					return helpers_io.UnknownOs
				}
				return result
			} else {
				if result, ok := op.Result.(helpers_io.OperatingSystem); ok {
					return result
				} else {
					return helpers_io.UnknownOs
				}
			}
		}
	}

	return helpers_io.UnknownOs
}

func (f MockFileIo) FileExists(path string) bool {
	for _, op := range f.mocks {
		if op.Method == "FileExists" {
			if op.Func != nil {
				argument := MockFuncArgument{
					Name:  "path",
					Value: path,
				}
				result, ok := op.Func(argument).(bool)
				if !ok {
					return false
				}
				return result
			} else {
				return op.Result.(bool)
			}
		}
	}

	return false
}

func (f MockFileIo) DirExists(folderPath string) bool {
	for _, op := range f.mocks {
		if op.Method == "DirExists" {
			if op.Func != nil {
				argument := MockFuncArgument{
					Name:  "folderPath",
					Value: folderPath,
				}
				result, ok := op.Func(argument).(bool)
				if !ok {
					return false
				}
				return result
			} else {
				return op.Result.(bool)
			}
		}
	}

	return false
}

func (f MockFileIo) CreateDir(folderPath string, mode os.FileMode) error {
	for _, op := range f.mocks {
		if op.Method == "CreateDir" {
			if op.Func != nil {
				argument := MockFuncArgument{
					Name:  "folderPath",
					Value: folderPath,
				}
				result, ok := op.Func(argument).(error)
				if !ok {
					return nil
				}
				return result
			} else {
				if op.Result == nil {
					return nil
				}
				if result, ok := op.Result.(error); ok {
					return result
				} else {
					return nil
				}
			}
		}
	}

	return nil
}

// func (f MockFileIo) GetExecutionPath() string {
// 	return os.Args[0]
// }

// func (f MockFileIo) GetOsPathSeparator() string {
// 	switch getOperatingSystem() {
// 	case WindowsOs:
// 		return "\\"
// 	case LinuxOs:
// 		return "/"
// 	case MacOs:
// 		return "/"
// 	default:
// 		return "/"
// 	}
// }

// func (f MockFileIo) ToOsPath(path string) string {
// 	switch getOperatingSystem() {
// 	case WindowsOs:
// 		return strings.ReplaceAll(path, "/", "\\")
// 	case LinuxOs:
// 		if strings.ContainsAny(path, ":") {
// 			pathParts := strings.Split(path, ":")
// 			path = pathParts[1]
// 		}
// 		return strings.ReplaceAll(path, "\\", "/")
// 	case MacOs:
// 		if strings.ContainsAny(path, ":") {
// 			pathParts := strings.Split(path, ":")
// 			path = pathParts[1]
// 		}
// 		return strings.ReplaceAll(path, "\\", "/")
// 	default:
// 		return path
// 	}
// }

// func (f MockFileIo) ReadFile(path string) ([]byte, error) {
// 	if !f.FileExists(path) {
// 		return nil, os.ErrNotExist
// 	}

// 	data, err := os.ReadFile(path)
// 	return data, err
// }

// func (f MockFileIo) ReadBufferedFile(path string, from, to int) ([]byte, error) {
// 	if !f.FileExists(path) {
// 		return nil, os.ErrNotExist
// 	}

// 	file, err := os.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	stat, err := file.Stat()
// 	if err != nil {
// 		return nil, err
// 	}

// 	if int(stat.Size()) < to || to == 0 {
// 		to = int(stat.Size())
// 	}

// 	buffer := make([]byte, to-from)
// 	_, err = file.ReadAt(buffer, int64(from))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return buffer, nil
// }

// func (f MockFileIo) WriteFile(path string, data []byte, mode os.FileMode) error {
// 	err := os.WriteFile(path, data, mode)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (f MockFileIo) WriteBufferedFile(path string, data []byte, bufferSize int) error {
// 	file, err := os.Create(path)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	for i := 0; i < len(data); i += bufferSize {
// 		end := i + bufferSize
// 		if end > len(data) {
// 			end = len(data)
// 		}
// 		_, err = file.Write(data[i:end])
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (f MockFileIo) ReadDir(path string) ([]fs.DirEntry, error) {
// 	dir, err := os.ReadDir(path)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return dir, nil
// }

// func (f MockFileIo) JoinPath(parts ...string) string {
// 	for i := range parts {
// 		parts[i] = strings.ReplaceAll(parts[i], "\\", "")
// 		parts[i] = strings.ReplaceAll(parts[i], "/", "")
// 	}

// 	return strings.Join(parts, f.GetOsPathSeparator())
// }

// func (f MockFileIo) CopyFile(source, destination string) error {
// 	sourceFile, err := os.Open(source)
// 	if err != nil {
// 		return err
// 	}
// 	defer sourceFile.Close()

// 	destinationFile, err := os.Create(destination)
// 	if err != nil {
// 		return err
// 	}
// 	defer destinationFile.Close()

// 	_, err = io.Copy(destinationFile, sourceFile)
// 	if err != nil {
// 		return err
// 	}

// 	err = destinationFile.Sync()
// 	if err != nil {
// 		return err
// 	}

// 	si, err := sourceFile.Stat()
// 	if err != nil {
// 		return err
// 	}

// 	err = os.Chmod(destination, si.Mode())
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (f MockFileIo) DeleteFile(path string) error {
// 	err := os.Remove(path)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (f MockFileIo) CopyDir(source, destination string) error {
// 	sourceInfo, err := os.Stat(source)
// 	if err != nil {
// 		return err
// 	}

// 	err = os.MkdirAll(destination, sourceInfo.Mode())
// 	if err != nil {
// 		return err
// 	}

// 	directory, err := os.ReadDir(source)
// 	if err != nil {
// 		return err
// 	}

// 	for _, file := range directory {
// 		sourcePath := filepath.Join(source, file.Name())
// 		destinationPath := filepath.Join(destination, file.Name())

// 		if file.IsDir() {
// 			err = f.CopyDir(sourcePath, destinationPath)
// 			if err != nil {
// 				return err
// 			}
// 		} else {
// 			err = f.CopyFile(sourcePath, destinationPath)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	return nil
// }

// func (f MockFileIo) DeleteDir(path string) error {
// 	err := os.RemoveAll(path)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (f MockFileIo) Checksum(path string, method ChecksumMethod) (string, error) {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return "", err
// 	}
// 	var hash hash.Hash
// 	switch method {
// 	case MD5:
// 		hash = md5.New()
// 		_, err := io.Copy(hash, file)
// 		if err != nil {
// 			return "", err
// 		}
// 	case SHA1:
// 		hash = sha1.New()
// 		_, err := io.Copy(hash, file)
// 		if err != nil {
// 			return "", err
// 		}
// 	case SHA256:
// 		hash = sha256.New()
// 		_, err := io.Copy(hash, file)
// 		if err != nil {
// 			return "", err
// 		}
// 	default:
// 		return "", errors.New("invalid checksum method")
// 	}

// 	return fmt.Sprintf("%x", hash.Sum(nil)), nil
// }
