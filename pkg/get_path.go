package pkg

import (
	"os"
)

func GetPathSize(systemPath string) (int64, error) {
	var res int64
	fiRoot, err := os.Lstat(systemPath)
	if err != nil {
		return 0, err
	}

	if fiRoot.IsDir() {
		files, err := os.ReadDir(".")
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
