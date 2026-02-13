package main

import "fmt"

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
func isValidBanner(name string) bool {
	_, exists := bannerPaths[name]
	return exists
}
