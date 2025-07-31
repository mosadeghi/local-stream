package util

import (
	"strconv"
	"strings"
)

func ParseRange(header string, fileSize int64) (int64, int64) {
	const prefix = "bytes="
	if !strings.HasPrefix(header, prefix) {
		return -1, -1
	}
	rangeSpec := header[len(prefix):]
	var start, end int64
	var err error

	if dash := strings.Index(rangeSpec, "-"); dash != -1 {
		start, err = strconv.ParseInt(rangeSpec[:dash], 10, 64)
		if err != nil || start >= fileSize {
			return -1, -1
		}
		if dash+1 < len(rangeSpec) {
			end, err = strconv.ParseInt(rangeSpec[dash+1:], 10, 64)
			if err != nil || end < start {
				return -1, -1
			}
		} else {
			end = fileSize - 1
		}
		return start, min(end, fileSize-1)
	}

	return -1, -1
}
