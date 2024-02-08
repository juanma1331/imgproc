package file

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"regexp"

	"github.com/chai2010/webp"
)

const ACCEPTED_FORMATS = "jpg|jpeg|png|webp"

type OSFilesProvider struct {
	Directory string
}

func (o OSFilesProvider) Files() ([]string, error) {
	files := make([]string, 0)
	regexPattern := regexp.MustCompile(fmt.Sprintf(`(?i)\.(%s)$`, ACCEPTED_FORMATS))

	err := filepath.Walk(o.Directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && regexPattern.MatchString(path) {
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error walking the path %s: %w", o.Directory, err)
	}

	return files, nil
}

func (o OSFilesProvider) AlreadyExists(file string) bool {
	_, err := os.Stat(file)

	return err == nil
}

type OSFilesWriter struct {
}

func (o OSFilesWriter) Write(file string, img image.Image) error {
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("error creating the file %s: %w", file, err)
	}
	defer f.Close()

	switch filepath.Ext(file) {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(f, img, nil)
	case ".png":
		err = png.Encode(f, img)
	case ".webp":
		err = webp.Encode(f, img, nil)
	}

	if err != nil {
		return fmt.Errorf("error encoding the image to %s: %w", file, err)
	}

	return nil
}
