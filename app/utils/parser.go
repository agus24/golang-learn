package utils

func ParseID(id any) int64 {
	return int64(id.(float64))
}
