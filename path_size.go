// Package code provides functions for getting size of files and dirs.
package code

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var units = [7]string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

const unitDiv float64 = 1000.0

// GetPathSize returns the size of the file or directory at path
// as a formatted string, e.g. "123B" or "1.5KB" when human is true.
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if isHidden(info.Name()) && !all {
		return "", errors.New("ошибка: если вы указываете путь до файла/директории, который начинается с '.' нужно передать флаг --all")
	}

	size, err := getSize(path, info, recursive, all)
	if err != nil {
		return "", err
	}

	return formatSize(size, human), nil
}

// getSize returns the size in bytes of the file or directory at path.
func getSize(path string, info os.FileInfo, recursive, all bool) (int64, error) {
	if !info.IsDir() {
		return info.Size(), nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var total int64

	for _, entry := range entries {
		if !all && isHidden(entry.Name()) {
			continue
		}

		entryInfo, err := entry.Info()
		if err != nil {
			return 0, err
		}

		if entryInfo.IsDir() {
			if !recursive {
				continue
			}

			dirSize, err := getSize(filepath.Join(path, entry.Name()), entryInfo, recursive, all)
			if err != nil {
				return 0, err
			}

			total += dirSize

			continue
		}

		total += entryInfo.Size()
	}

	return total, nil
}

// formatSize returns size as "<size><unit>"; units are auto-selected when human is true.
func formatSize(sizeBytes int64, human bool) string {
	size := float64(sizeBytes)

	if !human || size < unitDiv {
		return fmt.Sprintf("%dB", sizeBytes)
	}

	for _, unit := range units {
		if size < unitDiv {
			return fmt.Sprintf("%.1f%s", size, unit)
		}

		size /= unitDiv
	}

	return fmt.Sprintf("%.1f%s", size, units[len(units)-1])
}

func isHidden(name string) bool {
	return strings.HasPrefix(name, ".")
}
