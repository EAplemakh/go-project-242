// Package pathsize provides functions for calculating file and directory sizes.
package pathsize

import (
	"os"
)

// GetPathSize returns the size of the file or directory at systemPath in bytes.
func GetPathSize(systemPath string) (int64, error) {
	var res int64

	fiRoot, err := os.Lstat(systemPath)
	if err != nil {
		return 0, err
	}

	if fiRoot.IsDir() {
		files, err := os.ReadDir(systemPath)
		if err != nil {
			return 0, err
		}

		for _, file := range files {
			fi, err := file.Info()
			if err != nil {
				return 0, err
			}

			res += fi.Size()
		}
	} else {
		res = fiRoot.Size()
	}

	return res, nil
}
