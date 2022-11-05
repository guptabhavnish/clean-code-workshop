package utils

import "strconv"

const (
	TB = 1000 * GB
	GB = 1000 * MB
	MB = 1000 * KB
	KB = 1000
)

func ToReadableSize(nbytes int64) string {
	if nbytes > TB {
		return ConvertINT64ToString(nbytes, TB) + " TB"
	}
	if nbytes > GB {
		return ConvertINT64ToString(nbytes, GB) + " GB"
	}
	if nbytes > MB {
		return ConvertINT64ToString(nbytes, MB) + " MB"
	}
	if nbytes > KB {
		return ConvertINT64ToString(nbytes, KB) + " KB"
	}
	return strconv.FormatInt(nbytes, 10) + " B"
}

func ConvertINT64ToString(nbytes int64, size int64) string {
	return strconv.FormatInt(nbytes/size, 10)
}
