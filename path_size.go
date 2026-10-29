// Package code provides functions for getting size of files and dirs
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

// GetPathSize returns the size of the file or directory at systemPath in bytes.
func GetPathSize(systemPath string, includeHidden, recursive bool) (int64, error) {
	var res int64

	fiRoot, err := os.Lstat(systemPath)
	if err != nil {
		return 0, err
	}

	if strings.HasPrefix(fiRoot.Name(), ".") && !includeHidden {
		return 0, errors.New("ошибка: если вы указываете путь до файла/директории, который начинается с '.' нужно передать флаг --all")
	}

	if fiRoot.IsDir() {
		files, err := os.ReadDir(systemPath)
		if err != nil {
			return 0, err
		}

		for _, file := range files {
			if !includeHidden && strings.HasPrefix(file.Name(), ".") {
				continue
			}

			fi, err := file.Info()
			if err != nil {
				return 0, err
			}

			if fi.IsDir() && recursive {
				dirSum, err := GetPathSize(filepath.Join(systemPath, fi.Name()), includeHidden, recursive)
				if err != nil {
					return 0, err
				}

				res += dirSum
			} else {
				res += fi.Size()
			}
		}
	} else {
		res = fiRoot.Size()
	}

	return res, nil
}

// FormatUnitString returns formatted string <size><unit>.
func FormatUnitString(sizeBytes int64, chooseUnits bool) string {
	size := float64(sizeBytes)

	res := fmt.Sprintf("%dB", sizeBytes)
	if size < unitDiv || !chooseUnits {
		return res
	}

	for _, unit := range units {
		if size >= unitDiv {
			size /= unitDiv
			continue
		}

		res = fmt.Sprintf("%.1f%s", size, unit)

		break
	}

	return res
}
