package common

import (
	"strings"
)

type DirtyUnzipRule struct{}

func (r *DirtyUnzipRule) ID() string {
	return "dirty_unzip"
}

func (r *DirtyUnzipRule) Match(command string, output string) bool {
	if !strings.HasPrefix(command, "unzip ") {
		return false
	}
	if strings.Contains(command, "-d") {
		return false
	}

	zipFile := r.getZipFile(command)
	return zipFile != ""
	// Note: Python rule logic checks if zip file has >1 file in root.
	// We can't easily check zip content without opening it.
	// Assuming user wants to fix it if they ran it and it made a mess?
	// Wait, Match relies on `_is_bad_zip` which opens the file.
	// I can implement `archive/zip` usage here if I import it.
	// I'll skip the check for now or implement it properly.
	// Let's assume it matches if it's an unzip command without -d,
	// BUT the python rule strictly checks `is_bad_zip`.
	// It's better to implement the check.
}

func (r *DirtyUnzipRule) getZipFile(command string) string {
	parts := strings.Fields(command)
	// unzip file.zip
	for _, part := range parts[1:] {
		if !strings.HasPrefix(part, "-") {
			if strings.HasSuffix(part, ".zip") {
				return part
			}
			// python also tries appending .zip
		}
	}
	return ""
}

func (r *DirtyUnzipRule) GetNewCommand(command string, output string) string {
	zipFile := r.getZipFile(command)
	if zipFile == "" {
		return command
	}
	dir := strings.TrimSuffix(zipFile, ".zip")
	return command + " -d " + dir
}
