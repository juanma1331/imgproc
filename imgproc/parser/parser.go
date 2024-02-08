package parser

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseResizeFlag(resize string) (int, int, error) {
	// Example: 100x200
	split := strings.Split(resize, "x")

	width, err := strconv.Atoi(split[0])
	if err != nil {
		return 0, 0, fmt.Errorf("error converting the width to an integer: %w", err)
	}

	height, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, 0, fmt.Errorf("error converting the height to an integer: %w", err)
	}

	return width, height, nil
}

func ParseCropFlag(crop string) (int, int, int, int, error) {
	// Example: w,h,x,y
	split := strings.Split(crop, ",")

	width, err := strconv.Atoi(split[0])
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("error converting the width to an integer: %w", err)
	}

	height, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("error converting the height to an integer: %w", err)
	}

	x, err := strconv.Atoi(split[2])
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("error converting the x to an integer: %w", err)
	}

	y, err := strconv.Atoi(split[3])
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("error converting the y to an integer: %w", err)
	}

	return width, height, x, y, nil
}
