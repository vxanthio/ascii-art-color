package main

import (
	"os/exec"
	"strings"
	"testing"
)

// ============================================
// INTEGRATION TESTS - FULL PROGRAM EXECUTION
// ============================================

// Integration test: Run actual program and check output
func TestMainProgram_Integration(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectError bool
		checkOutput func(string) bool
	}{
		{
			name:        "Hello with standard banner",
			args:        []string{"Hello"},
			expectError: false,
			checkOutput: func(output string) bool {
				// Should have 8 lines
				lines := strings.Count(output, "\n")
				return lines == 8
			},
		},
		{
			name:        "Empty string",
			args:        []string{""},
			expectError: false,
			checkOutput: func(output string) bool {
				// Should have 8 empty lines
				return output == "\n\n\n\n\n\n\n\n"
			},
		},
		{
			name:        "With shadow banner",
			args:        []string{"Hi", "shadow"},
			expectError: false,
			checkOutput: func(output string) bool {
				// Should have 8 lines
				lines := strings.Count(output, "\n")
				return lines == 8
			},
		},
		{
			name:        "With thinkertoy banner",
			args:        []string{"Go", "thinkertoy"},
			expectError: false,
			checkOutput: func(output string) bool {
				// Should have 8 lines
				lines := strings.Count(output, "\n")
				return lines == 8
			},
		},
		{
			name:        "Multiple words with spaces",
			args:        []string{"Hello World"},
			expectError: false,
			checkOutput: func(output string) bool {
				// Should have 8 lines and contain space
				lines := strings.Count(output, "\n")
				return lines == 8 && len(output) > 0
			},
		},
		{
			name:        "Text with newline",
			args:        []string{"Hello\nWorld"},
			expectError: false,
			checkOutput: func(output string) bool {
				// Should have 16 lines (8 per text line)
				lines := strings.Count(output, "\n")
				return lines == 16
			},
		},
		{
			name:        "No arguments - usage error",
			args:        []string{},
			expectError: true,
			checkOutput: nil,
		},
		{
			name:        "Too many arguments",
			args:        []string{"Hello", "standard", "extra"},
			expectError: true,
			checkOutput: nil,
		},
		{
			name:        "Invalid banner",
			args:        []string{"Hello", "invalid"},
			expectError: true,
			checkOutput: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Build command
			args := append([]string{"run", "main.go"}, tt.args...)
			cmd := exec.Command("go", args...)
			
			// Run command
			output, err := cmd.CombinedOutput()
			
			// Check error expectation
			if tt.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v\nOutput: %s", err, output)
			}
			
			// Check output if provided
			if !tt.expectError && tt.checkOutput != nil {
				if !tt.checkOutput(string(output)) {
					t.Errorf("Output check failed.\nOutput:\n%s", output)
				}
			}
		})
	}
}

// Integration test: Test with actual banner files
func TestMainProgram_RealBannerFiles(t *testing.T) {
	// This test requires banner files to exist
	banners := []string{"standard", "shadow", "thinkertoy"}
	
	for _, banner := range banners {
		t.Run("Banner_"+banner, func(t *testing.T) {
			cmd := exec.Command("go", "run", "main.go", "ABC", banner)
			output, err := cmd.CombinedOutput()
			
			if err != nil {
				t.Errorf("Failed to run with %s banner: %v\nOutput: %s", 
					banner, err, output)
			}
			
			// Verify output has content
			if len(output) == 0 {
				t.Errorf("Expected output for banner %s, got empty", banner)
			}
			
			// Verify correct number of lines
			lines := strings.Count(string(output), "\n")
			if lines != 8 {
				t.Errorf("Expected 8 lines for banner %s, got %d", banner, lines)
			}
		})
	}
}

// Integration test: Error handling
func TestMainProgram_ErrorHandling(t *testing.T) {
	errorTests := []struct {
		name     string
		args     []string
		errorMsg string
	}{
		{
			name:     "No arguments",
			args:     []string{},
			errorMsg: "Usage:",
		},
		{
			name:     "Invalid banner",
			args:     []string{"Hello", "notexist"},
			errorMsg: "invalid banner",
		},
	}
	
	for _, tt := range errorTests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", append([]string{"run", "main.go"}, tt.args...)...)
			output, err := cmd.CombinedOutput()
			
			if err == nil {
				t.Errorf("Expected error for %s, got none", tt.name)
			}
			
			if !strings.Contains(string(output), tt.errorMsg) {
				t.Errorf("Expected error message containing %q, got: %s", 
					tt.errorMsg, output)
			}
		})
	}
}
