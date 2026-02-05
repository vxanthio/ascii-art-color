// Package flagparser_test contains unit tests for the flagparser package.
//
// These tests verify that ParseArgs correctly validates command-line arguments,
// including argument count, flag syntax, flag position, banner handling,
// and supported color formats.
package flagparser_test

import (
	"ascii-art-color/internal/flagparser"
	"testing"
)

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "no arguments",
			args:    []string{"program"},
			wantErr: true,
		},
		{
			name:    "too many arguments",
			args:    []string{"program", "banner", "--color=red", "substring", "some text", "EXTRA"},
			wantErr: true,
		},
		{
			name:    "invalid color flag prefix",
			args:    []string{"program", "-color:black", "some text"},
			wantErr: true,
		},
		{
			name:    "invalid color flag format without equals",
			args:    []string{"program", "--color:red", "some text"},
			wantErr: true,
		},
		{
			name:    "valid single string",
			args:    []string{"program", "text"},
			wantErr: false,
		},
		{
			name:    "color flag with string and no substring",
			args:    []string{"program", "--color=red", "text"},
			wantErr: false,
		},
		{
			name:    "color flag with string and substring",
			args:    []string{"program", "--color=red", "text", "substring"},
			wantErr: false,
		},
		{
			name:    "missing string after color flag",
			args:    []string{"program", "--color=red"},
			wantErr: true,
		},
		{
			name:    "multiple color flags",
			args:    []string{"program", "--color=red", "--color=blue", "text"},
			wantErr: true,
		},
		{
			name:    "invalid color flag position",
			args:    []string{"program", "text", "--color=red"},
			wantErr: true,
		},
		{
			name:    "valid RGB color",
			args:    []string{"program", "--color=rgb(255,0,0)", "text"},
			wantErr: false,
		},
		{
			name:    "invalid RGB out of range",
			args:    []string{"program", "--color=rgb(300,0,0)", "text"},
			wantErr: true,
		},
		{
			name:    "valid HEX color",
			args:    []string{"program", "--color=#ff0000", "text"},
			wantErr: false,
		},
		{
			name:    "invalid HEX length",
			args:    []string{"program", "--color=#123", "text"},
			wantErr: true,
		},
		{
			name:    "invalid color name",
			args:    []string{"program", "--color=tirquaz", "text"},
			wantErr: true,
		},
		{
			name:    "valid banner with color",
			args:    []string{"program", "--color=red", "text", "standard"},
			wantErr: false,
		},
		{
			name:    "valid banner without color",
			args:    []string{"program", "text", "standard"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := flagparser.ParseArgs(tt.args)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
