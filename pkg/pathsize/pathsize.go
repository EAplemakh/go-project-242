// Package pathsize provides functions for calculating file and directory sizes.
package pathsize

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

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
