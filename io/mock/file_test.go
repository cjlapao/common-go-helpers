package mock

import (
	"errors"
	"io/fs"
	"os"
	"testing"

	helpers_io "github.com/cjlapao/common-go-helpers/io"
	"github.com/stretchr/testify/assert"
)

func TestGetMockFuncArgumentValue(t *testing.T) {
	args := []MockFuncArgument{
		{Name: "arg1", Value: 10},
		{Name: "arg2", Value: "hello"},
		{Name: "arg3", Value: true},
	}

	t.Run("Existing Argument", func(t *testing.T) {
		value, ok := GetMockFuncArgumentValue[string](args, "arg2")
		assert.True(t, ok)
		assert.Equal(t, "hello", value)
	})

	t.Run("Non-Existing Argument", func(t *testing.T) {
		value, ok := GetMockFuncArgumentValue[string](args, "arg4")
		assert.False(t, ok)
		assert.Equal(t, "", value)
	})

	t.Run("Invalid Type Assertion", func(t *testing.T) {
		value, ok := GetMockFuncArgumentValue[string](args, "arg1")
		assert.False(t, ok)
		assert.Equal(t, "", value)
	})
}

func TestGetOperatingSystem(t *testing.T) {
	t.Run("Mock Function Returns Operating System", func(t *testing.T) {
		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method: "GetOperatingSystem",
					Func: func(args ...MockFuncArgument) interface{} {
						return helpers_io.WindowsOs
					},
				},
			},
		}

		expectedResult := helpers_io.WindowsOs
		actualResult := mockFileIo.GetOperatingSystem()

		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("Mock Function Returns Invalid Result Type", func(t *testing.T) {
		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method: "GetOperatingSystem",
					Func: func(args ...MockFuncArgument) interface{} {
						return "invalid"
					},
				},
			},
		}

		expectedResult := helpers_io.UnknownOs
		actualResult := mockFileIo.GetOperatingSystem()

		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("Mock Function Returns Error", func(t *testing.T) {
		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method: "GetOperatingSystem",
					Func: func(args ...MockFuncArgument) interface{} {
						return errors.New("error")
					},
				},
			},
		}

		expectedResult := helpers_io.UnknownOs
		actualResult := mockFileIo.GetOperatingSystem()

		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("No Mock Function Found", func(t *testing.T) {
		mockFileIo := MockFileIo{
			mocks: []*MockOperation{},
		}

		expectedResult := helpers_io.UnknownOs
		actualResult := mockFileIo.GetOperatingSystem()

		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("Mock Result", func(t *testing.T) {
		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method:      "GetOperatingSystem",
					ReturnValue: helpers_io.LinuxOs,
				},
			},
		}

		expectedResult := helpers_io.LinuxOs
		actualResult := mockFileIo.GetOperatingSystem()

		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("Mock Result", func(t *testing.T) {
		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method:      "GetOperatingSystem",
					ReturnValue: "aa",
				},
			},
		}

		expectedResult := helpers_io.UnknownOs
		actualResult := mockFileIo.GetOperatingSystem()

		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("Mock Nil Result", func(t *testing.T) {
		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method: "GetOperatingSystem",
				},
			},
		}

		expectedResult := helpers_io.UnknownOs
		actualResult := mockFileIo.GetOperatingSystem()

		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestMockFileIo_FileExists(t *testing.T) {
	mockFileIo := MockFileIo{
		mocks: []*MockOperation{
			{
				Method: "FileExists",
				Func: func(arg ...MockFuncArgument) interface{} {
					path, ok := GetMockFuncArgumentValue[string](arg, "path")
					if !ok {
						return false
					}

					if path == "existing_file.txt" {
						return true
					} else if path == "non_existing_file.txt" {
						return false
					}
					return nil
				},
				ReturnValue: true,
			},
		},
	}

	t.Run("Mock no Function", func(t *testing.T) {
		mockFileIo := MockFileIo{}
		result := mockFileIo.FileExists("test_file")
		assert.False(t, result)
	})

	t.Run("Existing File", func(t *testing.T) {
		path := "existing_file.txt"
		expectedResult := true

		result := mockFileIo.FileExists(path)

		assert.Equal(t, expectedResult, result)
	})

	t.Run("Non-existing File", func(t *testing.T) {
		path := "non_existing_file.txt"
		expectedResult := false

		result := mockFileIo.FileExists(path)

		assert.Equal(t, expectedResult, result)
	})

	t.Run("Unknown File", func(t *testing.T) {
		path := "unknown_file.txt"
		expectedResult := false

		result := mockFileIo.FileExists(path)

		assert.Equal(t, expectedResult, result)
	})

	t.Run("Mock Result", func(t *testing.T) {
		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method:      "FileExists",
					ReturnValue: false,
				},
			},
		}

		path := "unknown_file.txt"
		expectedResult := false

		result := mockFileIo.FileExists(path)

		assert.Equal(t, expectedResult, result)
	})

	t.Run("Mock Result true", func(t *testing.T) {
		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method:      "FileExists",
					ReturnValue: true,
				},
			},
		}

		path := "unknown_file.txt"
		expectedResult := true

		result := mockFileIo.FileExists(path)

		assert.Equal(t, expectedResult, result)
	})
}

func TestMockFileIo_DirExists(t *testing.T) {
	mockFileIo := MockFileIo{
		mocks: []*MockOperation{
			{
				Method: "DirExists",
				Func: func(arg ...MockFuncArgument) interface{} {
					path, ok := GetMockFuncArgumentValue[string](arg, "folderPath")
					if !ok {
						return false
					}

					return path == "/path/to/folder"
				},
			},
		},
	}

	t.Run("Mock no Function", func(t *testing.T) {
		mockFileIo := MockFileIo{}
		result := mockFileIo.DirExists("test_file")
		assert.False(t, result)
	})

	t.Run("Mocked Result: true", func(t *testing.T) {
		folderPath := "/path/to/folder"
		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method:      "DirExists",
					ReturnValue: true,
				},
			},
		}
		result := mockFileIo.DirExists(folderPath)
		assert.True(t, result)
	})

	t.Run("Mocked Result: false", func(t *testing.T) {
		folderPath := "/path/to/other/folder"
		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method:      "DirExists",
					ReturnValue: false,
				},
			},
		}
		result := mockFileIo.DirExists(folderPath)
		assert.False(t, result)
	})

	t.Run("Mocked Function", func(t *testing.T) {
		folderPath := "/path/to/folder"
		result := mockFileIo.DirExists(folderPath)
		assert.True(t, result)
	})

	t.Run("No Mock Match", func(t *testing.T) {
		folderPath := "/path/to/invalid/folder"
		result := mockFileIo.DirExists(folderPath)
		assert.False(t, result)
	})
}

func TestMockFileIo_CreateDir(t *testing.T) {
	mockFileIo := MockFileIo{
		mocks: []*MockOperation{
			{
				Method: "CreateDir",
				Func: func(args ...MockFuncArgument) interface{} {
					return nil
				},
			},
		},
	}

	t.Run("Mock no Function", func(t *testing.T) {
		mockFileIo := MockFileIo{}
		err := mockFileIo.CreateDir("test_folder", os.ModePerm)
		assert.Nil(t, err)
	})

	t.Run("Mock Function", func(t *testing.T) {
		folderPath := "test_folder"
		mode := os.ModePerm

		err := mockFileIo.CreateDir(folderPath, mode)
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
	})

	t.Run("Mock Function Error", func(t *testing.T) {
		folderPath := "test_folder"
		mode := os.ModePerm
		expectedErr := errors.New("mock error")

		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method: "CreateDir",
					Func: func(args ...MockFuncArgument) interface{} {
						return expectedErr
					},
				},
			},
		}

		err := mockFileIo.CreateDir(folderPath, mode)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("Mock Result", func(t *testing.T) {
		expectedErr := errors.New("mock error")

		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method:      "CreateDir",
					ReturnValue: expectedErr,
				},
			},
		}

		folderPath := "test_folder"
		mode := os.ModePerm

		err := mockFileIo.CreateDir(folderPath, mode)
		if err != expectedErr {
			t.Errorf("Expected error: %v, but got: %v", expectedErr, err)
		}
	})

	t.Run("Mock Nil Result", func(t *testing.T) {
		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method:      "CreateDir",
					ReturnValue: nil,
				},
			},
		}

		folderPath := "test_folder"
		mode := os.ModePerm

		err := mockFileIo.CreateDir(folderPath, mode)
		assert.Nil(t, err)
	})

	t.Run("Mock Wrong Result", func(t *testing.T) {
		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method:      "CreateDir",
					ReturnValue: false,
				},
			},
		}

		folderPath := "test_folder"
		mode := os.ModePerm

		err := mockFileIo.CreateDir(folderPath, mode)
		assert.Nil(t, err)
	})
}

func TestMockFileIo_GetExecutionPath(t *testing.T) {
	mockFileIo := MockFileIo{
		mocks: []*MockOperation{
			{
				Method: "GetExecutionPath",
				Func: func(args ...MockFuncArgument) interface{} {
					return "/path/to/executable"
				},
			},
		},
	}

	t.Run("Mock No Op", func(t *testing.T) {
		mockFileIo := MockFileIo{}
		result := mockFileIo.GetExecutionPath()
		assert.Empty(t, result)
	})

	t.Run("Mock Function", func(t *testing.T) {
		result := mockFileIo.GetExecutionPath()
		assert.Equal(t, "/path/to/executable", result)
	})

	t.Run("Mock Result", func(t *testing.T) {
		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method:      "GetExecutionPath",
					ReturnValue: "/path/to/executable",
				},
			},
		}

		folderPath := "/path/to/executable"
		err := mockFileIo.GetExecutionPath()
		assert.Equal(t, folderPath, err)
	})
}

func TestMockFileIo_GetOsPathSeparator(t *testing.T) {
	mockFileIo := MockFileIo{
		mocks: []*MockOperation{
			{
				Method: "GetOsPathSeparator",
				Func: func(args ...MockFuncArgument) interface{} {
					return "\\"
				},
			},
		},
	}

	t.Run("Mock No Op", func(t *testing.T) {
		mockFileIo := MockFileIo{}
		result := mockFileIo.GetOsPathSeparator()
		assert.Equal(t, "/", result)
	})

	t.Run("Mock Function", func(t *testing.T) {
		result := mockFileIo.GetOsPathSeparator()
		assert.Equal(t, "\\", result)
	})

	t.Run("Mock Result", func(t *testing.T) {
		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method:      "GetOsPathSeparator",
					ReturnValue: "/",
				},
			},
		}

		folderPath := "/"
		err := mockFileIo.GetOsPathSeparator()
		assert.Equal(t, folderPath, err)
	})
}

func TestMockFileIo_ToOsPath(t *testing.T) {
	mockFileIo := MockFileIo{
		mocks: []*MockOperation{
			{
				Method: "ToOsPath",
				Func: func(args ...MockFuncArgument) interface{} {
					return "/to/path"
				},
			},
		},
	}

	t.Run("Mock No Op", func(t *testing.T) {
		mockFileIo := MockFileIo{}
		result := mockFileIo.ToOsPath("non_existing_file.txt")
		assert.Equal(t, "non_existing_file.txt", result)
	})

	t.Run("Mock Function", func(t *testing.T) {
		result := mockFileIo.ToOsPath("\\to\\path")
		assert.Equal(t, "/to/path", result)
	})

	t.Run("Mock Result", func(t *testing.T) {
		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method:      "ToOsPath",
					ReturnValue: "/to/path",
				},
			},
		}

		folderPath := "/to/path"
		err := mockFileIo.ToOsPath("\\to\\path")
		assert.Equal(t, folderPath, err)
	})
}

func TestMockFileIo_ReadFile(t *testing.T) {
	filepath := "test_file.txt"
	mockFileIo := MockFileIo{
		mocks: []*MockOperation{
			{
				Method: "ReadFile",
				FuncWithErr: func(args ...MockFuncArgument) (interface{}, error) {
					return []byte("file content"), nil
				},
			},
		},
	}

	t.Run("Mock no Function", func(t *testing.T) {
		mockFileIo := MockFileIo{}
		content, err := mockFileIo.ReadFile(filepath)
		assert.Error(t, err)
		assert.Nil(t, content)
	})

	t.Run("Mock Function", func(t *testing.T) {
		content, err := mockFileIo.ReadFile(filepath)
		assert.NoError(t, err)
		assert.Equal(t, []byte("file content"), content)
	})

	t.Run("Mock Function Error", func(t *testing.T) {
		expectedErr := errors.New("mock error")

		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method: "ReadFile",
					FuncWithErr: func(args ...MockFuncArgument) (interface{}, error) {
						return "", expectedErr
					},
				},
			},
		}

		result, err := mockFileIo.ReadFile(filepath)
		assert.Equal(t, expectedErr, err)
		assert.Empty(t, result)
	})

	t.Run("Mock Result", func(t *testing.T) {
		expectedErr := errors.New("mock error")

		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method:      "ReadFile",
					ReturnError: expectedErr,
					ReturnValue: "",
				},
			},
		}

		content, err := mockFileIo.ReadFile(filepath)
		assert.Equal(t, expectedErr, err)
		assert.Empty(t, content)
	})

	t.Run("Mock Nil Result", func(t *testing.T) {
		expectedErr := errors.New("mock error")
		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method:      "ReadFile",
					ReturnValue: nil,
					ReturnError: expectedErr,
				},
			},
		}

		result, err := mockFileIo.ReadFile(filepath)
		assert.Equal(t, err, expectedErr)
		assert.Nil(t, result)
	})

	t.Run("Mock Wrong Result", func(t *testing.T) {
		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method:      "ReadFile",
					ReturnValue: false,
				},
			},
		}

		content, err := mockFileIo.ReadFile(filepath)
		assert.Nil(t, err)
		assert.Empty(t, content)
	})
}

func TestMockFileIo_ReadBufferedFile(t *testing.T) {
	filepath := "test_file.txt"
	from := 0
	to := 10
	mockFileIo := MockFileIo{
		mocks: []*MockOperation{
			{
				Method: "ReadBufferedFile",
				FuncWithErr: func(args ...MockFuncArgument) (interface{}, error) {
					return []byte("file content"), nil
				},
			},
		},
	}

	t.Run("Mock no Function", func(t *testing.T) {
		mockFileIo := MockFileIo{}
		content, err := mockFileIo.ReadBufferedFile(filepath, from, to)
		assert.Error(t, err)
		assert.Nil(t, content)
	})

	t.Run("Mock Function", func(t *testing.T) {
		content, err := mockFileIo.ReadBufferedFile(filepath, from, to)
		assert.NoError(t, err)
		assert.Equal(t, []byte("file content"), content)
	})

	t.Run("Mock Function Error", func(t *testing.T) {
		expectedErr := errors.New("mock error")

		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method: "ReadBufferedFile",
					FuncWithErr: func(args ...MockFuncArgument) (interface{}, error) {
						return "", expectedErr
					},
				},
			},
		}

		result, err := mockFileIo.ReadBufferedFile(filepath, from, to)
		assert.Equal(t, expectedErr, err)
		assert.Empty(t, result)
	})

	t.Run("Mock Result", func(t *testing.T) {
		expectedErr := errors.New("mock error")

		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method:      "ReadBufferedFile",
					ReturnError: expectedErr,
					ReturnValue: "",
				},
			},
		}

		content, err := mockFileIo.ReadBufferedFile(filepath, from, to)
		assert.Equal(t, expectedErr, err)
		assert.Empty(t, content)
	})

	t.Run("Mock Nil Result", func(t *testing.T) {
		expectedErr := errors.New("mock error")
		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method:      "ReadBufferedFile",
					ReturnValue: nil,
					ReturnError: expectedErr,
				},
			},
		}

		result, err := mockFileIo.ReadBufferedFile(filepath, from, to)
		assert.Equal(t, err, expectedErr)
		assert.Nil(t, result)
	})

	t.Run("Mock Wrong Result", func(t *testing.T) {
		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method:      "ReadBufferedFile",
					ReturnValue: false,
				},
			},
		}

		content, err := mockFileIo.ReadBufferedFile(filepath, from, to)
		assert.Nil(t, err)
		assert.Empty(t, content)
	})
}

func TestMockFileIo_WriteFile(t *testing.T) {
	filepath := "test_file.txt"
	data := []byte("file content")
	mode := os.ModePerm
	expectedError := errors.New("error")

	mockFileIo := MockFileIo{
		mocks: []*MockOperation{
			{
				Method: "WriteFile",
				Func: func(args ...MockFuncArgument) interface{} {
					path, ok := GetMockFuncArgumentValue[string](args, "path")
					if !ok {
						return nil
					}

					if path == "test_file.txt" {
						return nil
					} else {
						return expectedError
					}
				},
			},
		},
	}

	t.Run("Mock No Op", func(t *testing.T) {
		mockFileIo := MockFileIo{}
		result := mockFileIo.WriteFile(filepath, data, mode)
		assert.Empty(t, result)
	})

	t.Run("Mock Function with no error", func(t *testing.T) {
		result := mockFileIo.WriteFile(filepath, data, mode)
		assert.Nil(t, result)
	})

	t.Run("Mock Function with error", func(t *testing.T) {
		result := mockFileIo.WriteFile("unknown_file.txt", data, mode)
		assert.Equal(t, result, expectedError)
	})

	t.Run("Mock Result", func(t *testing.T) {
		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method:      "WriteFile",
					ReturnValue: expectedError,
				},
			},
		}

		result := mockFileIo.WriteFile(filepath, data, mode)
		assert.Equal(t, result, expectedError)
	})
}

func TestMockFileIo_WriteBufferedFile(t *testing.T) {
	filepath := "test_file.txt"
	data := []byte("file content")
	mode := os.ModePerm
	expectedError := errors.New("error")

	mockFileIo := MockFileIo{
		mocks: []*MockOperation{
			{
				Method: "WriteBufferedFile",
				Func: func(args ...MockFuncArgument) interface{} {
					path, ok := GetMockFuncArgumentValue[string](args, "path")
					if !ok {
						return nil
					}

					if path == "test_file.txt" {
						return nil
					} else {
						return expectedError
					}
				},
			},
		},
	}

	t.Run("Mock No Op", func(t *testing.T) {
		mockFileIo := MockFileIo{}
		result := mockFileIo.WriteBufferedFile(filepath, data, 10, mode)
		assert.Empty(t, result)
	})

	t.Run("Mock Function with no error", func(t *testing.T) {
		result := mockFileIo.WriteBufferedFile(filepath, data, 10, mode)
		assert.Nil(t, result)
	})

	t.Run("Mock Function with error", func(t *testing.T) {
		result := mockFileIo.WriteBufferedFile("unknown_file.txt", data, 10, mode)
		assert.Equal(t, result, expectedError)
	})

	t.Run("Mock Result", func(t *testing.T) {
		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method:      "WriteBufferedFile",
					ReturnValue: expectedError,
				},
			},
		}

		result := mockFileIo.WriteBufferedFile(filepath, data, 10, mode)
		assert.Equal(t, mockFileIo.mocks[0].CalledWith[0].Value, filepath)
		assert.Equal(t, result, expectedError)
	})
}

func TestMockFileIo_ReadDir(t *testing.T) {
	filepath := "test_file.txt"
	from := 0
	to := 10
	mockFileIo := MockFileIo{
		mocks: []*MockOperation{
			{
				Method: "ReadDir",
				FuncWithErr: func(args ...MockFuncArgument) (interface{}, error) {
					return []fs.DirEntry{}, nil
				},
			},
		},
	}

	t.Run("Mock no Function", func(t *testing.T) {
		mockFileIo := MockFileIo{}
		content, err := mockFileIo.ReadDir(filepath)
		assert.Error(t, err)
		assert.Nil(t, content)
	})

	t.Run("Mock Function", func(t *testing.T) {
		content, err := mockFileIo.ReadDir(filepath)
		assert.NoError(t, err)
		assert.Equal(t, []fs.DirEntry{}, content)
	})

	t.Run("Mock Function Error", func(t *testing.T) {
		expectedErr := errors.New("mock error")

		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method: "ReadDir",
					FuncWithErr: func(args ...MockFuncArgument) (interface{}, error) {
						return nil, expectedErr
					},
				},
			},
		}

		result, err := mockFileIo.ReadDir(filepath)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
	})

	t.Run("Mock Result", func(t *testing.T) {
		expectedErr := errors.New("mock error")

		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method:      "ReadBufferedFile",
					ReturnError: expectedErr,
					ReturnValue: nil,
				},
			},
		}

		content, err := mockFileIo.ReadDir(filepath)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, content)
	})

	t.Run("Mock Nil Result", func(t *testing.T) {
		expectedErr := errors.New("mock error")
		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method:      "ReadDir",
					ReturnValue: nil,
					ReturnError: expectedErr,
				},
			},
		}

		result, err := mockFileIo.ReadDir(filepath)
		assert.Equal(t, err, expectedErr)
		assert.Nil(t, result)
	})

	t.Run("Mock Wrong Result", func(t *testing.T) {
		mockFileIo := MockFileIo{
			mocks: []*MockOperation{
				{
					Method:      "ReadBufferedFile",
					ReturnValue: false,
				},
			},
		}

		content, err := mockFileIo.ReadBufferedFile(filepath, from, to)
		assert.Nil(t, err)
		assert.Nil(t, content)
	})
}
