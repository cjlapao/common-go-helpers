package mock

import (
	"errors"
	"os"
	"testing"

	helpers_io "github.com/cjlapao/common-go-helpers/io"
	"github.com/stretchr/testify/assert"
)

func TestGetOperatingSystem(t *testing.T) {
	t.Run("Mock Function Returns Operating System", func(t *testing.T) {
		mockFileIo := MockFileIo{
			mocks: []MockOperation{
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
			mocks: []MockOperation{
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
			mocks: []MockOperation{
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
			mocks: []MockOperation{},
		}

		expectedResult := helpers_io.UnknownOs
		actualResult := mockFileIo.GetOperatingSystem()

		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("Mock Result", func(t *testing.T) {
		mockFileIo := MockFileIo{
			mocks: []MockOperation{
				{
					Method: "GetOperatingSystem",
					Result: helpers_io.LinuxOs,
				},
			},
		}

		expectedResult := helpers_io.LinuxOs
		actualResult := mockFileIo.GetOperatingSystem()

		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("Mock Result", func(t *testing.T) {
		mockFileIo := MockFileIo{
			mocks: []MockOperation{
				{
					Method: "GetOperatingSystem",
					Result: "aa",
				},
			},
		}

		expectedResult := helpers_io.UnknownOs
		actualResult := mockFileIo.GetOperatingSystem()

		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("Mock Nil Result", func(t *testing.T) {
		mockFileIo := MockFileIo{
			mocks: []MockOperation{
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
		mocks: []MockOperation{
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
				Result: true,
			},
		},
	}

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
}

func TestMockFileIo_DirExists(t *testing.T) {
	mockFileIo := MockFileIo{
		mocks: []MockOperation{
			{
				Method: "DirExists",
				Func: func(args ...MockFuncArgument) interface{} {
					folderPath, ok := GetMockFuncArgumentValue[string](args, "folderPath")
					if !ok {
						return ""
					}

					return folderPath == "/existing/folder"
				},
			},
		},
	}

	t.Run("Directory Exists", func(t *testing.T) {
		folderPath := "/existing/folder"
		result := mockFileIo.DirExists(folderPath)
		assert.True(t, result, "Expected directory to exist")
	})

	t.Run("Directory Does Not Exist", func(t *testing.T) {
		folderPath := "/non_existing/folder"
		result := mockFileIo.DirExists(folderPath)
		assert.False(t, result, "Expected directory to not exist")
	})

	t.Run("Custom Directory Check", func(t *testing.T) {
		mockFileIo := NewMockFileIo()
		mockFileIo.On(MockOperation{
			Method: "DirExists",
			Result: interface{}(true),
		})

		folderPath := "/path/to/folder"
		result := mockFileIo.DirExists(folderPath)
		assert.True(t, result, "Expected custom directory check to return true")
	})
}

func TestMockFileIo_CreateDir(t *testing.T) {
	mockFileIo := MockFileIo{
		mocks: []MockOperation{
			{
				Method: "CreateDir",
				Func: func(args ...MockFuncArgument) interface{} {
					return nil
				},
			},
		},
	}

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
			mocks: []MockOperation{
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
			mocks: []MockOperation{
				{
					Method: "CreateDir",
					Result: expectedErr,
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
			mocks: []MockOperation{
				{
					Method: "CreateDir",
					Result: nil,
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
			mocks: []MockOperation{
				{
					Method: "CreateDir",
					Result: false,
				},
			},
		}

		folderPath := "test_folder"
		mode := os.ModePerm

		err := mockFileIo.CreateDir(folderPath, mode)
		assert.Nil(t, err)
	})
}
