package main

import (
	"fmt"
	"imgproc-cli/imgproc/file"
	"io"
	"os"
	"path/filepath"
)

const FILES_TO_CLONE = 100

func main() {
	filesProvider := file.OSFilesProvider{Directory: "images"}

	// clone the image on the same directory
	files, err := filesProvider.Files()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(files) == 0 {
		fmt.Println("no images found")
		return
	}

	fileDirectory := filepath.Dir(files[0])
	fileBaseName := filepath.Base(files[0])

	for i := 0; i < FILES_TO_CLONE; i++ {
		newFileName := fmt.Sprintf("%d-%s", i, fileBaseName)
		newFilePath := filepath.Join(fileDirectory, newFileName)

		if filesProvider.AlreadyExists(newFilePath) {
			fmt.Printf("file %s already exists\n", newFilePath)
			continue
		}

		err := copyFile(files[0], newFilePath)
		if err != nil {
			fmt.Println(err)
			return
		}

	}

}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}
