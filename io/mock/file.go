package mock

import (
	"io/fs"
	"os"

	helpers_io "github.com/cjlapao/common-go-helpers/io"
)

type MockOperation struct {
	Method      string
	Func        func(args ...MockFuncArgument) interface{}
	FuncWithErr func(args ...MockFuncArgument) (interface{}, error)
	ReturnError error
	CalledWith  []MockFuncArgument
	ReturnValue interface{}
}

type MockFuncArgument struct {
	Name  string
	Value interface{}
}

type MockFileIo struct {
	mocks []*MockOperation
}

func NewMockFileIo() *MockFileIo {
	return &MockFileIo{
		mocks: []*MockOperation{},
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

	var def T
	return def, false
}

func (f *MockFileIo) On(op MockOperation) *MockOperation {
	if f.mocks == nil {
		f.mocks = []*MockOperation{}
	}

	f.mocks = append(f.mocks, &op)
	return f.mocks[len(f.mocks)-1]
}

func (f MockFileIo) GetOperatingSystem() helpers_io.OperatingSystem {
	for _, op := range f.mocks {
		if op.Method == "GetOperatingSystem" {
			if op.Func != nil {
				return processFunction[helpers_io.OperatingSystem](op.Func)
			} else {
				return processResult[helpers_io.OperatingSystem](op.ReturnValue)
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
				op.CalledWith = append(op.CalledWith, argument)
				return processFunction[bool](op.Func, argument)
			} else {
				return processResult[bool](op.ReturnValue)
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
				op.CalledWith = append(op.CalledWith, argument)
				return processFunction[bool](op.Func, argument)
			} else {
				return processResult[bool](op.ReturnValue)
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
				op.CalledWith = append(op.CalledWith, argument)
				return processFunction[error](op.Func, argument)
			} else {
				return processResult[error](op.ReturnValue)
			}
		}
	}

	return nil
}

func (f MockFileIo) GetExecutionPath() string {
	for _, op := range f.mocks {
		if op.Method == "GetExecutionPath" {
			if op.Func != nil {
				return processFunction[string](op.Func)
			} else {
				return processResult[string](op.ReturnValue)
			}
		}
	}

	return ""
}

func (f MockFileIo) GetOsPathSeparator() string {
	for _, op := range f.mocks {
		if op.Method == "GetOsPathSeparator" {
			if op.Func != nil {
				return processFunction[string](op.Func)
			} else {
				return processResult[string](op.ReturnValue)
			}
		}
	}

	return "/"
}

func (f MockFileIo) ToOsPath(path string) string {
	for _, op := range f.mocks {
		if op.Method == "ToOsPath" {
			if op.Func != nil {
				argument := MockFuncArgument{
					Name:  "path",
					Value: path,
				}
				op.CalledWith = append(op.CalledWith, argument)
				return processFunction[string](op.Func, argument)
			} else {
				return processResult[string](op.ReturnValue)
			}
		}
	}

	return path
}

func (f MockFileIo) ReadFile(path string) ([]byte, error) {
	for _, op := range f.mocks {
		if op.Method == "ReadFile" {
			if op.FuncWithErr != nil {
				argument := MockFuncArgument{
					Name:  "path",
					Value: path,
				}
				op.CalledWith = append(op.CalledWith, argument)
				return processFunctionWithErr[[]byte](op.FuncWithErr, op.ReturnError, argument)
			} else {
				return processResult[[]byte](op.ReturnValue), op.ReturnError
			}
		}
	}

	return nil, os.ErrNotExist
}

func (f MockFileIo) ReadBufferedFile(path string, from, to int) ([]byte, error) {
	for _, op := range f.mocks {
		if op.Method == "ReadBufferedFile" {
			if op.FuncWithErr != nil {
				argument1 := MockFuncArgument{
					Name:  "path",
					Value: path,
				}
				argument2 := MockFuncArgument{
					Name:  "from",
					Value: from,
				}
				argument3 := MockFuncArgument{
					Name:  "to",
					Value: to,
				}
				op.CalledWith = append(op.CalledWith, argument1, argument2, argument3)
				return processFunctionWithErr[[]byte](op.FuncWithErr, op.ReturnError, argument1, argument2, argument3)
			} else {
				return processResult[[]byte](op.ReturnValue), op.ReturnError
			}
		}
	}

	return nil, os.ErrNotExist
}

func (f MockFileIo) WriteFile(path string, data []byte, mode os.FileMode) error {
	for _, op := range f.mocks {
		if op.Method == "WriteFile" {
			if op.Func != nil {
				argument1 := MockFuncArgument{
					Name:  "path",
					Value: path,
				}
				argument2 := MockFuncArgument{
					Name:  "data",
					Value: data,
				}
				argument3 := MockFuncArgument{
					Name:  "mode",
					Value: mode,
				}
				op.CalledWith = append(op.CalledWith, argument1, argument2, argument3)
				return processFunction[error](op.Func, argument1, argument2, argument3)
			} else {
				return processResult[error](op.ReturnValue)
			}
		}
	}

	return nil
}

func (f MockFileIo) WriteBufferedFile(path string, data []byte, bufferSize int, mode os.FileMode) error {
	for _, op := range f.mocks {
		if op.Method == "WriteBufferedFile" {
			if op.Func != nil {
				argument1 := MockFuncArgument{
					Name:  "path",
					Value: path,
				}
				argument2 := MockFuncArgument{
					Name:  "data",
					Value: data,
				}
				argument3 := MockFuncArgument{
					Name:  "bufferSize",
					Value: bufferSize,
				}
				op.CalledWith = append(op.CalledWith, argument1, argument2, argument3)
				return processFunction[error](op.Func, argument1, argument2, argument3)
			} else {
				return processResult[error](op.ReturnValue)
			}
		}
	}

	return nil
}

func (f MockFileIo) ReadDir(path string) ([]fs.DirEntry, error) {
	for _, op := range f.mocks {
		if op.Method == "ReadDir" {
			if op.FuncWithErr != nil {
				argument := MockFuncArgument{
					Name:  "path",
					Value: path,
				}
				op.CalledWith = append(op.CalledWith, argument)
				return processFunctionWithErr[[]fs.DirEntry](op.FuncWithErr, op.ReturnError, argument)
			} else {
				return processResult[[]fs.DirEntry](op.ReturnValue), op.ReturnError
			}
		}
	}

	return nil, os.ErrNotExist
}

// func (f MockFileIo) JoinPath(parts ...string) string {
// }

// func (f MockFileIo) CopyFile(source, destination string) error {
// }

// func (f MockFileIo) DeleteFile(path string) error {
// }

// func (f MockFileIo) CopyDir(source, destination string) error {
// }

// func (f MockFileIo) DeleteDir(path string) error {
// }

// func (f MockFileIo) Checksum(path string, method ChecksumMethod) (string, error) {
// }

func processFunction[T any](fn func(args ...MockFuncArgument) interface{}, args ...MockFuncArgument) T {
	var def T
	if fn != nil {
		result, ok := fn(args...).(T)
		if !ok {
			return def
		}
		return result
	}

	return def
}

func processFunctionWithErr[T any](fn func(args ...MockFuncArgument) (interface{}, error), err error, args ...MockFuncArgument) (T, error) {
	var def T
	if fn != nil {
		result, funcErr := fn(args...)
		if result, ok := result.(T); !ok {
			return def, funcErr
		} else {
			return result, funcErr
		}
	}

	return def, err
}

func processResult[T any](result interface{}) T {
	var def T
	if result != nil {
		if result, ok := result.(T); ok {
			return result
		} else {
			return def
		}
	}

	return def
}
