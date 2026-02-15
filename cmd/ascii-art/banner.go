package main

import (
	"embed"
	"fmt"
	"io/fs"
)

// bannerFS embeds the testdata directory into the compiled binary.
// This makes the binary fully self-contained and relocatable - it can be run
// from any directory without requiring testdata files to exist on disk.
// The embedded files are read-only and frozen at compile time.
//
//go:embed testdata/*.txt
var bannerFS embed.FS

var bannerPaths = map[string]string{
	"standard":   "testdata/standard.txt",
	"shadow":     "testdata/shadow.txt",
	"thinkertoy": "testdata/thinkertoy.txt",
}

// GetBannerPath converts a banner name to its corresponding file path.
//
// The function validates the banner name against a predefined map of valid banners
// (standard, shadow, thinkertoy) and returns the appropriate file path in the testdata
// directory.
//
// Parameters:
//   - banner: The banner name to resolve.
//
// Returns:
//   - The file path to the banner file.
//   - An error if the banner name is invalid.
func GetBannerPath(banner string) (string, error) {
	path, exists := bannerPaths[banner]
	if !exists {
		return "", fmt.Errorf("invalid banner name: %q\nValid options: standard, shadow, thinkertoy", banner)
	}

	return path, nil
}

// isValidBanner checks whether a string is a recognized banner name.
//
// Parameters:
//   - name: The banner name to validate.
//
// Returns:
//   - true if the banner name is valid (standard, shadow, or thinkertoy), false otherwise.
func isValidBanner(name string) bool {
	_, exists := bannerPaths[name]
	return exists
}

// GetBannerFS returns the embedded filesystem containing banner files.
//
// The filesystem is embedded at compile time and contains all banner files
// from the testdata directory. This allows the binary to be fully relocatable
// and run from any directory without requiring external data files.
//
// Returns:
//   - An fs.FS interface to the embedded testdata directory.
func GetBannerFS() fs.FS {
	return bannerFS
}
