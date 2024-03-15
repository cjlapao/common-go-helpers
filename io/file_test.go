package io

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestPath() string {
	environmentDir := os.Getenv("TEST_DIR")
	if environmentDir != "" {
		return environmentDir
	}

	return "./tests"
}

func TestFileExists(t *testing.T) {
	t.Run("File Exists", func(t *testing.T) {
		defaultClient := Default()
		existingFilePath := filepath.Join(getTestPath(), "test_file_1.txt")
		if !defaultClient.FileExists(existingFilePath) {
			t.Errorf("Expected file to exist, but it doesn't")
		}
	})

	t.Run("File Does Not Exist", func(t *testing.T) {
		defaultClient := Default()
		nonExistingFilePath := filepath.Join(getTestPath(), "non_existing_file.txt")
		if defaultClient.FileExists(nonExistingFilePath) {
			t.Errorf("Expected file to not exist, but it does")
		}
	})
}

func TestDirExists(t *testing.T) {
	t.Run("Directory Exists", func(t *testing.T) {
		defaultClient := Default()
		existingDirPath := filepath.Join(getTestPath(), "test_dir_1")
		if !defaultClient.DirExists(existingDirPath) {
			t.Errorf("Expected directory to exist, but it doesn't")
		}
	})

	t.Run("Directory Does Not Exist", func(t *testing.T) {
		defaultClient := Default()
		nonExistingDirPath := filepath.Join(getTestPath(), "non_existing_dir")
		if defaultClient.DirExists(nonExistingDirPath) {
			t.Errorf("Expected directory to not exist, but it does")
		}
	})
}

func TestCreateDir(t *testing.T) {
	t.Run("Create Directory", func(t *testing.T) {
		defaultClient := Default()
		testDirPath := filepath.Join(getTestPath(), "test_dir")
		testDirMode := os.ModePerm

		err := defaultClient.CreateDir(testDirPath, testDirMode)
		if err != nil {
			t.Errorf("Failed to create directory: %v", err)
		}

		// Verify if the directory exists
		if !defaultClient.DirExists(testDirPath) {
			t.Errorf("Expected directory to exist, but it doesn't")
		}

		// Clean up: remove the created directory
		err = os.Remove(testDirPath)
		if err != nil {
			t.Errorf("Failed to remove directory: %v", err)
		}
	})
}

func TestGetExecutionPath(t *testing.T) {
	t.Run("Get Execution Path", func(t *testing.T) {
		defaultClient := Default()
		expectedPath := os.Args[0]
		actualPath := defaultClient.GetExecutionPath()

		if actualPath != expectedPath {
			t.Errorf("Expected execution path to be %s, but got %s", expectedPath, actualPath)
		}
	})
}

func TestGetOsPathSeparator(t *testing.T) {
	t.Run("Windows OS", func(t *testing.T) {
		os.Setenv("TEST_OS_OVERRIDE", "windows")
		defaultClient := Default()
		expectedSeparator := "\\"
		actualSeparator := defaultClient.GetOsPathSeparator()

		if actualSeparator != expectedSeparator {
			t.Errorf("Expected path separator to be %s, but got %s", expectedSeparator, actualSeparator)
		}
	})

	t.Run("Linux OS", func(t *testing.T) {
		os.Setenv("TEST_OS_OVERRIDE", "linux")
		defaultClient := Default()
		expectedSeparator := "/"
		actualSeparator := defaultClient.GetOsPathSeparator()

		if actualSeparator != expectedSeparator {
			t.Errorf("Expected path separator to be %s, but got %s", expectedSeparator, actualSeparator)
		}
	})

	t.Run("Mac OS", func(t *testing.T) {
		os.Setenv("TEST_OS_OVERRIDE", "darwin")
		defaultClient := Default()
		expectedSeparator := "/"
		actualSeparator := defaultClient.GetOsPathSeparator()

		if actualSeparator != expectedSeparator {
			t.Errorf("Expected path separator to be %s, but got %s", expectedSeparator, actualSeparator)
		}
	})

	t.Run("Unknown OS", func(t *testing.T) {
		os.Setenv("TEST_OS_OVERRIDE", "unknown")
		defaultClient := Default()
		expectedSeparator := "/"
		actualSeparator := defaultClient.GetOsPathSeparator()

		if actualSeparator != expectedSeparator {
			t.Errorf("Expected path separator to be %s, but got %s", expectedSeparator, actualSeparator)
		}
	})
}

func TestToOsPath(t *testing.T) {
	t.Run("Windows OS", func(t *testing.T) {
		defaultClient := Default()
		os.Setenv("TEST_OS_OVERRIDE", "windows")
		expectedPath := "C:\\path\\to\\file"
		actualPath := defaultClient.ToOsPath("C:/path/to/file")

		if actualPath != expectedPath {
			t.Errorf("Expected OS path to be %s, but got %s", expectedPath, actualPath)
		}
	})

	t.Run("Linux OS", func(t *testing.T) {
		defaultClient := Default()
		os.Setenv("TEST_OS_OVERRIDE", "linux")
		expectedPath := "/path/to/file"
		actualPath := defaultClient.ToOsPath("C:\\path\\to\\file")

		if actualPath != expectedPath {
			t.Errorf("Expected OS path to be %s, but got %s", expectedPath, actualPath)
		}
	})

	t.Run("Mac OS", func(t *testing.T) {
		defaultClient := Default()
		os.Setenv("TEST_OS_OVERRIDE", "darwin")
		expectedPath := "/path/to/file"
		actualPath := defaultClient.ToOsPath("C:\\path\\to\\file")

		if actualPath != expectedPath {
			t.Errorf("Expected OS path to be %s, but got %s", expectedPath, actualPath)
		}
	})

	t.Run("Unknown OS", func(t *testing.T) {
		defaultClient := Default()
		os.Setenv("TEST_OS_OVERRIDE", "unknown")
		expectedPath := "C:\\path\\to\\file"
		actualPath := defaultClient.ToOsPath("C:\\path\\to\\file")

		if actualPath != expectedPath {
			t.Errorf("Expected OS path to be %s, but got %s", expectedPath, actualPath)
		}
	})
}

func TestReadFile(t *testing.T) {
	t.Run("File Exists", func(t *testing.T) {
		defaultClient := Default()
		existingFilePath := filepath.Join(getTestPath(), "test_file_1.txt")
		expectedFileContent := `Initial bytes
This is Second Line
More Text`

		data, err := defaultClient.ReadFile(existingFilePath)
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}

		// Verify if the data is not empty
		if len(data) == 0 {
			t.Errorf("Expected non-empty data, but got empty")
		}

		assert.Equal(t, string(data), expectedFileContent)
	})

	t.Run("File Does Not Exist", func(t *testing.T) {
		defaultClient := Default()
		nonExistingFilePath := filepath.Join(getTestPath(), "non_existing_file.txt")

		_, err := defaultClient.ReadFile(nonExistingFilePath)
		if !errors.Is(err, os.ErrNotExist) {
			t.Errorf("Expected error os.ErrNotExist, but got: %v", err)
		}
	})
}

func TestReadBufferedFile(t *testing.T) {
	t.Run("File Exists", func(t *testing.T) {
		defaultClient := Default()
		existingFilePath := filepath.Join(getTestPath(), "test_file_1.txt")
		expectedBuffer := []byte("Initial")

		buffer, err := defaultClient.ReadBufferedFile(existingFilePath, 0, 7)
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}

		assert.Equal(t, buffer, expectedBuffer)
	})

	t.Run("File Does Not Exist", func(t *testing.T) {
		defaultClient := Default()
		nonExistingFilePath := filepath.Join(getTestPath(), "non_existing_file.txt")

		_, err := defaultClient.ReadBufferedFile(nonExistingFilePath, 0, 50)
		if !errors.Is(err, os.ErrNotExist) {
			t.Errorf("Expected error os.ErrNotExist, but got: %v", err)
		}
	})

	t.Run("Read Full File", func(t *testing.T) {
		defaultClient := Default()
		existingFilePath := filepath.Join(getTestPath(), "test_file_1.txt")
		expectedBuffer := []byte("Initial bytes\nThis is Second Line\nMore Text")

		buffer, err := defaultClient.ReadBufferedFile(existingFilePath, 0, 0)
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}

		assert.Equal(t, buffer, expectedBuffer)
	})

	t.Run("Read Partial File", func(t *testing.T) {
		defaultClient := Default()
		existingFilePath := filepath.Join(getTestPath(), "test_file_1.txt")
		expectedBuffer := []byte("This is Second Line\nMore Text")

		buffer, err := defaultClient.ReadBufferedFile(existingFilePath, 14, 50)
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}

		assert.Equal(t, buffer, expectedBuffer)
	})

	t.Run("Read Beyond File Size", func(t *testing.T) {
		defaultClient := Default()
		existingFilePath := filepath.Join(getTestPath(), "test_file_1.txt")
		expectedBuffer := []byte("This is Second Line\nMore Text")

		buffer, err := defaultClient.ReadBufferedFile(existingFilePath, 14, 1000)
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}

		assert.Equal(t, buffer, expectedBuffer)
	})
}

func TestWriteFile(t *testing.T) {
	t.Run("Write File", func(t *testing.T) {
		defaultClient := Default()
		testFilePath := filepath.Join(getTestPath(), "test_file.txt")
		testFileMode := os.ModePerm
		testData := []byte("Test data")

		err := defaultClient.WriteFile(testFilePath, testData, testFileMode)
		if err != nil {
			t.Errorf("Failed to write file: %v", err)
		}

		// Verify if the file exists
		if !defaultClient.FileExists(testFilePath) {
			t.Errorf("Expected file to exist, but it doesn't")
		}

		// Verify if the file content matches the test data
		fileContent, err := defaultClient.ReadFile(testFilePath)
		if err != nil {
			t.Errorf("Failed to read file: %v", err)
		}

		if string(fileContent) != string(testData) {
			t.Errorf("Expected file content to be %s, but got %s", string(testData), string(fileContent))
		}

		// Clean up: remove the created file
		err = os.Remove(testFilePath)
		if err != nil {
			t.Errorf("Failed to remove file: %v", err)
		}
	})
}

func TestWriteBufferedFile(t *testing.T) {
	t.Run("Write Buffered File", func(t *testing.T) {
		defaultClient := Default()
		testFilePath := filepath.Join(getTestPath(), "test_file.txt")
		testData := []byte("Test data")
		bufferSize := 5

		err := defaultClient.WriteBufferedFile(testFilePath, testData, bufferSize, os.ModePerm)
		if err != nil {
			t.Errorf("Failed to write buffered file: %v", err)
		}

		// Verify if the file exists
		if !defaultClient.FileExists(testFilePath) {
			t.Errorf("Expected file to exist, but it doesn't")
		}

		// Verify if the file content matches the test data
		fileContent, err := defaultClient.ReadFile(testFilePath)
		if err != nil {
			t.Errorf("Failed to read file: %v", err)
		}

		if string(fileContent) != string(testData) {
			t.Errorf("Expected file content to be %s, but got %s", string(testData), string(fileContent))
		}

		// Clean up: remove the created file
		err = os.Remove(testFilePath)
		if err != nil {
			t.Errorf("Failed to remove file: %v", err)
		}
	})

	t.Run("Write Empty Buffered File", func(t *testing.T) {
		defaultClient := Default()
		testFilePath := filepath.Join(getTestPath(), "test_file.txt")
		testData := []byte{}
		bufferSize := 5

		err := defaultClient.WriteBufferedFile(testFilePath, testData, bufferSize, os.ModePerm)
		if err != nil {
			t.Errorf("Failed to write buffered file: %v", err)
		}

		// Verify if the file exists
		if !defaultClient.FileExists(testFilePath) {
			t.Errorf("Expected file to exist, but it doesn't")
		}

		// Verify if the file content matches the test data
		fileContent, err := defaultClient.ReadFile(testFilePath)
		if err != nil {
			t.Errorf("Failed to read file: %v", err)
		}

		if string(fileContent) != string(testData) {
			t.Errorf("Expected file content to be %s, but got %s", string(testData), string(fileContent))
		}

		// Clean up: remove the created file
		err = os.Remove(testFilePath)
		if err != nil {
			t.Errorf("Failed to remove file: %v", err)
		}
	})

	t.Run("Write Buffered File with Large Data", func(t *testing.T) {
		defaultClient := Default()
		testFilePath := filepath.Join(getTestPath(), "test_file.txt")
		testData := []byte("This is a large amount of data")
		bufferSize := 10

		err := defaultClient.WriteBufferedFile(testFilePath, testData, bufferSize, os.ModePerm)
		if err != nil {
			t.Errorf("Failed to write buffered file: %v", err)
		}

		// Verify if the file exists
		if !defaultClient.FileExists(testFilePath) {
			t.Errorf("Expected file to exist, but it doesn't")
		}

		// Verify if the file content matches the test data
		fileContent, err := defaultClient.ReadFile(testFilePath)
		if err != nil {
			t.Errorf("Failed to read file: %v", err)
		}

		if string(fileContent) != string(testData) {
			t.Errorf("Expected file content to be %s, but got %s", string(testData), string(fileContent))
		}

		// Clean up: remove the created file
		err = os.Remove(testFilePath)
		if err != nil {
			t.Errorf("Failed to remove file: %v", err)
		}
	})
}

func TestReadDir(t *testing.T) {
	t.Run("Read Directory", func(t *testing.T) {
		defaultClient := Default()
		testDirPath := filepath.Join(getTestPath(), "test_dir_1")

		entries, err := defaultClient.ReadDir(testDirPath)
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}

		// Verify if the entries are not empty
		if len(entries) == 0 {
			t.Errorf("Expected non-empty entries, but got empty")
		}

		// TODO: Add assertions for the expected directory entries
	})
}

func TestJoinPath(t *testing.T) {
	defaultClient := Default()

	t.Run("Join Path with Windows Separator", func(t *testing.T) {
		os.Setenv("TEST_OS_OVERRIDE", "windows")
		expectedPath := "path\\to\\file"
		actualPath := defaultClient.JoinPath("path\\", "to\\", "file")

		if actualPath != expectedPath {
			t.Errorf("Expected joined path to be %s, but got %s", expectedPath, actualPath)
		}
	})

	t.Run("Join Path with Linux Separator", func(t *testing.T) {
		os.Setenv("TEST_OS_OVERRIDE", "linux")
		expectedPath := "path/to/file"
		actualPath := defaultClient.JoinPath("path/", "to/", "file")

		if actualPath != expectedPath {
			t.Errorf("Expected joined path to be %s, but got %s", expectedPath, actualPath)
		}
	})

	t.Run("Join Path with Mixed Separators", func(t *testing.T) {
		os.Setenv("TEST_OS_OVERRIDE", "windows")
		expectedPath := "path\\to\\file"
		actualPath := defaultClient.JoinPath("path\\", "to/", "file")

		if actualPath != expectedPath {
			t.Errorf("Expected joined path to be %s, but got %s", expectedPath, actualPath)
		}
	})
}

func TestCopyFile(t *testing.T) {
	t.Run("Copy File", func(t *testing.T) {
		defaultClient := Default()
		sourceFilePath := filepath.Join(getTestPath(), "source_file.txt")
		destinationFilePath := filepath.Join(getTestPath(), "destination_file.txt")

		// Create a source file with some content
		sourceContent := []byte("This is the source file")
		err := defaultClient.WriteFile(sourceFilePath, sourceContent, os.ModePerm)
		if err != nil {
			t.Errorf("Failed to create source file: %v", err)
		}

		// Copy the source file to the destination file
		err = defaultClient.CopyFile(sourceFilePath, destinationFilePath)
		if err != nil {
			t.Errorf("Failed to copy file: %v", err)
		}

		// Verify if the destination file exists
		if !defaultClient.FileExists(destinationFilePath) {
			t.Errorf("Expected destination file to exist, but it doesn't")
		}

		// Verify if the destination file content matches the source file content
		sourceFileContent, err := defaultClient.ReadFile(sourceFilePath)
		if err != nil {
			t.Errorf("Failed to read source file: %v", err)
		}

		destinationFileContent, err := defaultClient.ReadFile(destinationFilePath)
		if err != nil {
			t.Errorf("Failed to read destination file: %v", err)
		}

		if string(destinationFileContent) != string(sourceFileContent) {
			t.Errorf("Expected destination file content to be %s, but got %s", string(sourceFileContent), string(destinationFileContent))
		}

		// Clean up: remove the source and destination files
		err = os.Remove(sourceFilePath)
		if err != nil {
			t.Errorf("Failed to remove source file: %v", err)
		}

		err = os.Remove(destinationFilePath)
		if err != nil {
			t.Errorf("Failed to remove destination file: %v", err)
		}
	})

	t.Run("Copy Non-Existent File", func(t *testing.T) {
		defaultClient := Default()
		nonExistentFilePath := filepath.Join(getTestPath(), "non_existent_file.txt")
		destinationFilePath := filepath.Join(getTestPath(), "destination_file.txt")

		// Copy the non-existent file to the destination file
		err := defaultClient.CopyFile(nonExistentFilePath, destinationFilePath)
		if !errors.Is(err, os.ErrNotExist) {
			t.Errorf("Expected error os.ErrNotExist, but got: %v", err)
		}
	})

	t.Run("Copy to Existing File", func(t *testing.T) {
		defaultClient := Default()
		sourceFilePath := filepath.Join(getTestPath(), "source_file.txt")
		destinationFilePath := filepath.Join(getTestPath(), "destination_file.txt")

		// Create a source file with some content
		sourceContent := []byte("This is the source file")
		err := defaultClient.WriteFile(sourceFilePath, sourceContent, os.ModePerm)
		if err != nil {
			t.Errorf("Failed to create source file: %v", err)
		}

		// Create a destination file
		destinationContent := []byte("This is the destination file")
		err = defaultClient.WriteFile(destinationFilePath, destinationContent, os.ModePerm)
		if err != nil {
			t.Errorf("Failed to create destination file: %v", err)
		}

		// Copy the source file to the existing destination file
		err = defaultClient.CopyFile(sourceFilePath, destinationFilePath)
		if err != nil {
			t.Errorf("Failed to copy file: %v", err)
		}

		// Verify if the destination file content is updated
		destinationFileContent, err := defaultClient.ReadFile(destinationFilePath)
		if err != nil {
			t.Errorf("Failed to read destination file: %v", err)
		}

		if string(destinationFileContent) != string(sourceContent) {
			t.Errorf("Expected destination file content to be %s, but got %s", string(sourceContent), string(destinationFileContent))
		}

		// Clean up: remove the source and destination files
		err = os.Remove(sourceFilePath)
		if err != nil {
			t.Errorf("Failed to remove source file: %v", err)
		}

		err = os.Remove(destinationFilePath)
		if err != nil {
			t.Errorf("Failed to remove destination file: %v", err)
		}
	})
}

func TestDeleteFile(t *testing.T) {
	t.Run("Delete Existing File", func(t *testing.T) {
		defaultClient := Default()
		testFilePath := filepath.Join(getTestPath(), "test_file.txt")
		testFileMode := os.ModePerm
		testData := []byte("Test data")

		// Create a test file
		err := defaultClient.WriteFile(testFilePath, testData, testFileMode)
		if err != nil {
			t.Errorf("Failed to create test file: %v", err)
		}

		// Verify if the file exists before deletion
		if !defaultClient.FileExists(testFilePath) {
			t.Errorf("Expected file to exist, but it doesn't")
		}

		// Delete the file
		err = defaultClient.DeleteFile(testFilePath)
		if err != nil {
			t.Errorf("Failed to delete file: %v", err)
		}

		// Verify if the file is deleted
		if defaultClient.FileExists(testFilePath) {
			t.Errorf("Expected file to be deleted, but it still exists")
		}
	})

	t.Run("Delete Non-existing File", func(t *testing.T) {
		defaultClient := Default()
		nonExistingFilePath := filepath.Join(getTestPath(), "non_existing_file.txt")

		// Delete the non-existing file
		err := defaultClient.DeleteFile(nonExistingFilePath)
		if !errors.Is(err, os.ErrNotExist) {
			t.Errorf("Expected error os.ErrNotExist, but got: %v", err)
		}
	})
}

func TestCopyDir(t *testing.T) {
	t.Run("Copy Directory", func(t *testing.T) {
		defaultClient := Default()
		sourceDir := filepath.Join(getTestPath(), "source_dir")
		sourceSubDir := filepath.Join(sourceDir, "sub_dir")
		destinationDir := filepath.Join(getTestPath(), "destination_dir")

		// Create a source directory with some files
		err := defaultClient.CreateDir(sourceDir, os.ModePerm)
		if err != nil {
			t.Errorf("Failed to create source directory: %v", err)
		}

		err = defaultClient.CreateDir(sourceSubDir, os.ModePerm)
		if err != nil {
			t.Errorf("Failed to create source sub directory: %v", err)
		}

		err = defaultClient.WriteFile(filepath.Join(sourceDir, "file1.txt"), []byte("File 1"), os.ModePerm)
		if err != nil {
			t.Errorf("Failed to create file1.txt: %v", err)
		}

		err = defaultClient.WriteFile(filepath.Join(sourceDir, "file2.txt"), []byte("File 2"), os.ModePerm)
		if err != nil {
			t.Errorf("Failed to create file2.txt: %v", err)
		}

		err = defaultClient.WriteFile(filepath.Join(sourceSubDir, "file3.txt"), []byte("File 3"), os.ModePerm)
		if err != nil {
			t.Errorf("Failed to create file3.txt: %v", err)
		}

		// Copy the source directory to the destination directory
		err = defaultClient.CopyDir(sourceDir, destinationDir)
		if err != nil {
			t.Errorf("Failed to copy directory: %v", err)
		}

		// Verify if the destination directory exists
		if !defaultClient.DirExists(destinationDir) {
			t.Errorf("Expected destination directory to exist, but it doesn't")
		}

		// Verify if the files in the destination directory match the source directory
		sourceFiles, err := defaultClient.ReadDir(sourceDir)
		if err != nil {
			t.Errorf("Failed to read source directory: %v", err)
		}

		destinationFiles, err := defaultClient.ReadDir(destinationDir)
		if err != nil {
			t.Errorf("Failed to read destination directory: %v", err)
		}

		if len(sourceFiles) != len(destinationFiles) {
			t.Errorf("Expected number of files in source and destination directories to be the same")
		}

		for _, sourceFile := range sourceFiles {
			sourceFilePath := filepath.Join(sourceDir, sourceFile.Name())
			destinationFilePath := filepath.Join(destinationDir, sourceFile.Name())

			if sourceFile.IsDir() {
				continue
			}

			sourceFileContent, err := defaultClient.ReadFile(sourceFilePath)
			if err != nil {
				t.Errorf("Failed to read source file: %v", err)
			}

			destinationFileContent, err := defaultClient.ReadFile(destinationFilePath)
			if err != nil {
				t.Errorf("Failed to read destination file: %v", err)
			}

			if string(sourceFileContent) != string(destinationFileContent) {
				t.Errorf("Expected file content in source and destination directories to be the same")
			}
		}

		// Clean up: remove the source and destination directories
		err = os.RemoveAll(sourceDir)
		if err != nil {
			t.Errorf("Failed to remove source directory: %v", err)
		}

		err = os.RemoveAll(destinationDir)
		if err != nil {
			t.Errorf("Failed to remove destination directory: %v", err)
		}
	})
}

func TestDeleteDir(t *testing.T) {
	t.Run("Delete Directory", func(t *testing.T) {
		defaultClient := Default()
		testDirPath := filepath.Join(getTestPath(), "test_delete_dir")

		// Create a test directory
		err := os.MkdirAll(testDirPath, os.ModePerm)
		if err != nil {
			t.Fatalf("Failed to create test directory: %v", err)
		}

		// create a file in the folder
		err = defaultClient.WriteFile(filepath.Join(testDirPath, "file1.txt"), []byte("File 1"), os.ModePerm)
		if err != nil {
			t.Errorf("Failed to create file1.txt: %v", err)
		}

		// Verify if the directory exists before deletion
		if !defaultClient.DirExists(testDirPath) {
			t.Errorf("Expected directory to exist, but it doesn't")
		}

		// Delete the directory
		err = defaultClient.DeleteDir(testDirPath)
		if err != nil {
			t.Errorf("Failed to delete directory: %v", err)
		}

		// Verify if the directory is deleted
		if defaultClient.DirExists(testDirPath) {
			t.Errorf("Expected directory to be deleted, but it still exists")
		}
	})

	t.Run("Delete Non-existing Directory", func(t *testing.T) {
		defaultClient := Default()
		nonExistingDirPath := filepath.Join(getTestPath(), "non_existing_dir")

		// Verify if the directory does not exist before deletion
		if defaultClient.DirExists(nonExistingDirPath) {
			t.Errorf("Expected directory to not exist, but it does")
		}

		// Delete the directory
		err := defaultClient.DeleteDir(nonExistingDirPath)
		if err != nil {
			t.Errorf("Failed to delete directory: %v", err)
		}
	})
}

func TestChecksum(t *testing.T) {
	t.Run("File does not exist", func(t *testing.T) {
		defaultClient := Default()
		testFilePath := filepath.Join(getTestPath(), "test_file.txt")
		_, err := defaultClient.Checksum(testFilePath, ChecksumMD5)

		assert.Error(t, err)
	})

	t.Run("MD5 Checksum", func(t *testing.T) {
		defaultClient := Default()
		testFilePath := filepath.Join(getTestPath(), "test_file_1.txt")
		expectedChecksum := "bad71408e80acc34a474d42ce219d154"

		checksum, err := defaultClient.Checksum(testFilePath, ChecksumMD5)
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}

		assert.Equal(t, checksum, expectedChecksum)
	})

	t.Run("SHA1 Checksum", func(t *testing.T) {
		defaultClient := Default()
		testFilePath := filepath.Join(getTestPath(), "test_file_1.txt")
		expectedChecksum := "346722065c7c68422dcfbfa6bb6280300aa168a6"

		checksum, err := defaultClient.Checksum(testFilePath, ChecksumSHA1)
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}

		assert.Equal(t, checksum, expectedChecksum)
	})

	t.Run("SHA256 Checksum", func(t *testing.T) {
		defaultClient := Default()
		testFilePath := filepath.Join(getTestPath(), "test_file_1.txt")
		expectedChecksum := "030685cfa852639dee5e327f54153df00af48f75e146331b44ee72fe3b0cee6a"

		checksum, err := defaultClient.Checksum(testFilePath, ChecksumSHA256)
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}

		assert.Equal(t, checksum, expectedChecksum)
	})

	t.Run("Invalid Checksum Method", func(t *testing.T) {
		defaultClient := Default()
		testFilePath := filepath.Join(getTestPath(), "test_file_1.txt")

		_, err := defaultClient.Checksum(testFilePath, 10)
		if err == nil {
			t.Errorf("Expected an error, but got nil")
		}
	})
}

func TestFileInfo(t *testing.T) {
	t.Run("File Exists", func(t *testing.T) {
		defaultClient := Default()
		existingFilePath := filepath.Join(getTestPath(), "test_file_1.txt")

		fileInfo, err := defaultClient.FileInfo(existingFilePath)
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}

		// Verify if the file info is not nil
		if fileInfo == nil {
			t.Errorf("Expected non-nil file info, but got nil")
		}

		// Verify if the file exists
		if !defaultClient.FileExists(existingFilePath) {
			t.Errorf("Expected file to exist, but it doesn't")
		}
	})

	t.Run("File Does Not Exist", func(t *testing.T) {
		defaultClient := Default()
		nonExistingFilePath := filepath.Join(getTestPath(), "non_existing_file.txt")

		_, err := defaultClient.FileInfo(nonExistingFilePath)
		if !errors.Is(err, os.ErrNotExist) {
			t.Errorf("Expected error os.ErrNotExist, but got: %v", err)
		}
	})
}
