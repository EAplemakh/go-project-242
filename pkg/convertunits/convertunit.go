// Package convertunits provides function for converting units.
package convertunits

import "fmt"

var units = [7]string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

const unitDiv float64 = 1000.0

// FormatUnitString returns formatted string <size><unit>.
func FormatUnitString(sizeBytes int64, useBytes bool) string {
	size := float64(sizeBytes)

	res := fmt.Sprintf("%dB", sizeBytes)
	if size < unitDiv || useBytes {
		return res
	}

	for _, unit := range units {
		if size >= unitDiv {
			size /= unitDiv
			continue
		}

		res = fmt.Sprintf("%.1f%s", size, unit)
	}

	return res
}
