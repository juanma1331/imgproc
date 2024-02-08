/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"imgproc-cli/imgproc"
	"imgproc-cli/imgproc/file"
	"imgproc-cli/imgproc/operations"
	"imgproc-cli/imgproc/parser"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	directory string
	resize    string
	crop      string
	format    string

	rootCmd = &cobra.Command{
		Use:   "imgproc",
		Short: "A cli tool to process images.",
		Long: ` 
			imgproc is a cli tool to process images. It can resize, crop and convert images to different formats.
			You can use it to process a single image or a whole directory of images.
			Example: imgproc --directory="/path/to/images" --resize="100x200" --crop="10,20,50,50" --format="webp"
			The order of the operations matters. The operations are applied in the order they are provided.`,

		Run: func(cmd *cobra.Command, args []string) {
			ops := make([]imgproc.Operation, 0)

			if cmd.Flags().Changed("resize") {
				w, h, err := parser.ParseResizeFlag(resize)
				if err != nil {
					log.Fatal(err)
					return
				}

				ops = append(ops, operations.ResizeOperation{Width: w, Height: h})
			}

			if cmd.Flags().Changed("crop") {
				w, h, x, y, err := parser.ParseCropFlag(crop)
				if err != nil {
					log.Fatal(err)
					return
				}

				ops = append(ops, operations.CropOperation{Width: w, Height: h, X: x, Y: y})
			}

			filesProvider := file.OSFilesProvider{Directory: directory}

			filesWriter := file.OSFilesWriter{}

			imageProcessor := imgproc.ImageProcessor{
				Provider:   filesProvider,
				Writer:     filesWriter,
				Operations: ops,
			}

			err := imageProcessor.Process(format)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().StringVar(&directory, "directory", ".", "Directory to process images")
	rootCmd.Flags().StringVar(&resize, "resize", "", "Resize images")
	rootCmd.Flags().StringVar(&crop, "crop", "", "Crop images")
	rootCmd.Flags().StringVar(&format, "format", "", "Format to convert images")
}
