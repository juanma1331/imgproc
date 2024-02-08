# imgproc CLI

## Overview

imgproc is a command-line interface (CLI) tool designed for image processing tasks such as resizing, cropping, and converting images to different formats. It supports processing individual images or batches of images within a directory.

## Features

- Resize images to specified dimensions.
- Crop images to specified dimensions and positions.
- Convert images to different formats (JPG, PNG, WEBP).
- Process a single image or an entire directory of images.

## Installation

Clone this repository and navigate into the project directory. Ensure you have Go installed on your system and then run:

```bash
make build
```

This will compile the CLI tool and place the executable in the `bin` directory.

## Usage

To use imgproc, you can execute the binary directly from the command line. Here are some examples of how to use the tool:

```bash
./bin/imgproc --directory="/path/to/images" --resize="100x200" --crop="10,20,50,50" --format="webp"
```

### Available Flags

- `--directory`: Specify the directory to process images. Defaults to the current directory.
- `--resize`: Resize images to specified dimensions (width x height).
- `--crop`: Crop images to specified dimensions and position (width,height,x,y).
- `--format`: Convert images to a specified format (jpg, png, webp).

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
