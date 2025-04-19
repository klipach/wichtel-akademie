package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Photo struct {
	PhotoId        int    `json:"PhotoId"`
	FileName       string `json:"FileName"`
	Description    string `json:"Description"`
	Date           string `json:"Date"`
	GpsCoordinates struct {
		Longitude string `json:"Longitude"`
		Latitude  string `json:"Latitude"`
	} `json:"GpsCoordinates"`
}

func main() {
	// Read the JSON file
	jsonData, err := os.ReadFile("photos.json")
	if err != nil {
		fmt.Printf("Error reading JSON file: %v\n", err)
		return
	}

	var photos []Photo
	if err := json.Unmarshal(jsonData, &photos); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	// Desired GPS coordinates
	latitude := "48.1313726"
	longitude := "11.6585446"

	// Process each photo
	for _, photo := range photos {
		// Skip if the file doesn't exist
		fileName := fmt.Sprintf("photos/%d.jpg", photo.PhotoId)
		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			fmt.Printf("Skipping %s: file not found\n", fileName)
			continue
		}

		// Parse the date string
		photoDate, err := time.Parse("2006-01-02T15:04:05", photo.Date)
		if err != nil {
			fmt.Printf("Error parsing date for %s: %v\n", photo.FileName, err)
			continue
		}

		// Format date for exiftool
		dateStr := photoDate.Format("2006:01:02 15:04:05")

		args := []string{
			"exiftool",
			"-overwrite_original",
			fmt.Sprintf("-Description='%s'", escapeQuotes(photo.Description)),
			fmt.Sprintf("-DateTimeOriginal='%s'", dateStr),
			fmt.Sprintf("-GPSLatitude=%s", latitude),
			fmt.Sprintf("-GPSLongitude=%s", longitude),
			fmt.Sprintf("-GPSLatitudeRef=%s", "N"),
			fmt.Sprintf("-GPSLongitudeRef=%s", "E"),
			fileName,
		}

		// Execute the command
		cmd := exec.Command(args[0], args[1:]...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error executing command for %s: %v\nOutput: %s\n", fileName, err, string(output))
			continue
		}

		// Print success message
		fmt.Printf("Successfully updated EXIF data for %s\n", fileName)
	}
}

// escapeQuotes escapes single quotes in the description for shell safety
func escapeQuotes(s string) string {
	return strings.ReplaceAll(s, "'", "'\\''")
}
