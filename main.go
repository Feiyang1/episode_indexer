package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	inputDir := flag.String("i", "", "The directory where the episodes are stored.")
	flag.Parse()

	if *inputDir == "" {
		log.Fatal("-i is required, but nothing is provided.")
	}

	// Obtain the absolute dir path in case it's a relative path
	absoluteDirPath, err := filepath.Abs(*inputDir)

	if err != nil {
		log.Fatalf("Failed to obtain the absolute directory path for %s.", *inputDir)
	}

	// Read files in the directory and prepend episode number to file name.
	entries, err := os.ReadDir(*inputDir)

	if err != nil {
		log.Fatalf("Failed to read directory %s", *inputDir)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		fileName := entry.Name()
		// if the filename already starts with a number, assume it's the episode number and doesn't need change.
		already_good, _ := regexp.MatchString(`^\d`, fileName)
		if already_good {
			continue
		}

		r, _ := regexp.Compile(`第(\d+)集`)
		match := r.FindStringSubmatch(fileName)
		if len(match) == 0 {
			continue
		}
		episode_number := match[1]

		fullOriginalFileName := filepath.Join(absoluteDirPath, fileName)
		newFileName := filepath.Join(absoluteDirPath, episode_number+". "+fileName)

		fmt.Printf("%s -> %s\n", fullOriginalFileName, newFileName)
		os.Rename(fullOriginalFileName, newFileName)
	}

	fmt.Println("Hello World!")
}
