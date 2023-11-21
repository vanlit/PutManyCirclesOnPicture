package main

import (
	"encoding/csv"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"

	"github.com/vanlit/PutCircleOnPicture"
)

type CircleDescription struct {
	X, Y  int
	Size  int
	Color color.RGBA
}

const usage = `Usage: PutManyCirclesOnPicture <input_image_path> <output_image_path> <csvFileWithCirclesDescriptions>
The structure of the csv:
	{
		X: integer,
		Y: integer,
		Color: HTML-notation of the color
	}
`

func main() {
	if len(os.Args) != 4 {
		fmt.Print(usage)
		os.Exit(1)
	}

	inputImagePath := os.Args[1]
	outputImagePath := os.Args[2]
	csvFilePath := os.Args[3]

	circleDescriptions, cdErr := readCSVOfCircleDescriptions(csvFilePath)
	if cdErr != nil {
		fmt.Printf("Error while reading the csv: %v\n", cdErr)
		os.Exit(1)
	}

	err := processImage(inputImagePath, outputImagePath, circleDescriptions)
	if err != nil {
		fmt.Printf("Error processing image: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Image with %d colored square(s) saved to %s\n", len(circleDescriptions), outputImagePath)
}

func readCSVOfCircleDescriptions(filePath string) ([]CircleDescription, error) {
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all the records from CSV
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Convert records to RowType
	var circleDescriptions []CircleDescription
	for rowIndex, record := range records[1:] {
		x := parseInt(record[0])
		y := parseInt(record[1])
		size := parseInt(record[2])

		row := CircleDescription{
			X:     x,
			Y:     y,
			Size:  size,
			Color: parseColor(record[3]),
		}

		validationErr := validateInputs(row)
		if validationErr != nil {
			return nil, fmt.Errorf("invalid value in row #%d: err: %v", rowIndex, validationErr)
		}

		circleDescriptions = append(circleDescriptions, row)
	}

	return circleDescriptions, nil
}

func validateInputs(circleDesc CircleDescription) error {
	// Validate x and y coordinates
	err := validatePositiveInt(circleDesc.X)
	if err != nil {
		return fmt.Errorf("invalid x coordinate: %v", err)
	}

	err = validatePositiveInt(circleDesc.Y)
	if err != nil {
		return fmt.Errorf("invalid y coordinate: %v", err)
	}

	// Validate square size
	err = validatePositiveInt(circleDesc.Size)
	if err != nil {
		return fmt.Errorf("invalid size: %v", err)
	}

	return nil
}

func processImage(inputPath, outputPath string, circleDescriptions []CircleDescription) error {
	// Read the input image
	img, err := readImage(inputPath)
	if err != nil {
		return fmt.Errorf("error reading input image: %v", err)
	}

	// Create a new draw.Image from the existing image
	drawImg := image.NewRGBA(img.Bounds())
	draw.Draw(drawImg, drawImg.Bounds(), img, image.Point{}, draw.Over)

	for _, cd := range circleDescriptions {
		PutCircleOnPicture.DrawFilledCircle(drawImg, cd.X, cd.Y, cd.Size, cd.Color)
	}

	// Save the modified image to the output path
	err = saveImage(outputPath, drawImg)
	if err != nil {
		return fmt.Errorf("error saving output image: %v", err)
	}

	return nil
}

func saveImage(path string, img image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	err = jpeg.Encode(file, img, nil)
	if err != nil {
		return err
	}

	return nil
}

func readImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func validatePositiveInt(val int) error {
	if val <= 0 {
		return fmt.Errorf("value must be a positive number")
	}

	return nil
}

func parseColor(colorStr string) color.RGBA {
	var r, g, b, a uint8
	fmt.Sscanf(colorStr, "%02x%02x%02x%02x", &r, &g, &b, &a)
	return color.RGBA{r, g, b, a}
}

func parseInt(valStr string) int {
	var val int
	fmt.Sscanf(valStr, "%d", &val)
	return val
}
