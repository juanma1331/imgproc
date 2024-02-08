package imgproc

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Operation interface {
	Apply(image.Image) (image.Image, error)
	Name() string
}

type FilesProvider interface {
	Files() ([]string, error)
	AlreadyExists(string) bool
}

type FilesWriter interface {
	Write(string, image.Image) error
}

type ImageProcessor struct {
	Operations []Operation
	Provider   FilesProvider
	Writer     FilesWriter
}

func (i ImageProcessor) Process(outputExt string) error {
	start := time.Now()
	files, err := i.Provider.Files()
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return fmt.Errorf("no images found")
	}

	jobs := make(chan string, len(files))
	results := make(chan error, len(files))
	var wg sync.WaitGroup

	numWorkers := runtime.NumCPU()
	fmt.Printf("Using %d workers\n", numWorkers)
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, i.Operations, jobs, results, i.Provider, i.Writer, outputExt, &wg)
	}

	for _, file := range files {
		jobs <- file
	}
	close(jobs)

	wg.Wait()
	close(results)

	for err := range results {
		if err != nil {
			return err
		}
	}

	fmt.Printf("Processed %d images in %v\n", len(files), time.Since(start))

	return nil
}

func worker(id int, operations []Operation, jobs <-chan string, results chan<- error, provider FilesProvider, writer FilesWriter, outputExt string, wg *sync.WaitGroup) {
	defer wg.Done()

	for file := range jobs {
		fullName := buildFullName(file, outputExt, operations)

		if provider.AlreadyExists(fullName) {
			fmt.Printf("File %s already exists, skipping\n", fullName)
			results <- nil
			continue
		}

		img, err := loadImage(file)
		if err != nil {
			results <- err
			continue
		}

		for _, operation := range operations {
			img, err = operation.Apply(img)
			if err != nil {
				results <- err
				break
			}
		}

		if err == nil {
			err = writer.Write(fullName, img)
			if err != nil {
				results <- err
			} else {
				results <- nil
			}
		}
	}
}

func loadImage(file string) (image.Image, error) {
	imageFile, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("error opening the file %s: %w", file, err)
	}
	defer imageFile.Close()

	img, _, err := image.Decode(imageFile)
	if err != nil {
		return nil, fmt.Errorf("error decoding the image: %w", err)
	}

	return img, nil
}

func buildFullName(file, fileExt string, operations []Operation) string {
	baseName := file[:len(file)-len(filepath.Ext(file))] // Extract the base name of the file without its extension
	operationNames := getOperationNames(file, operations)

	if fileExt != "" {
		return fmt.Sprintf("%s%s.%s", baseName, operationNames, fileExt)
	}
	return fmt.Sprintf("%s%s%s", baseName, operationNames, filepath.Ext(file))
}

func getOperationNames(file string, operations []Operation) string {
	var operationNames strings.Builder
	currentName := filepath.Base(file)

	for _, operation := range operations {
		opName := fmt.Sprintf("_%s", operation.Name())
		if !strings.Contains(currentName, opName) {
			operationNames.WriteString(opName)
		}
	}

	return operationNames.String()
}
