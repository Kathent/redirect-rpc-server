package utils

func Merge(slice1, slice2 []string) []string{
	if len(slice1) <= 0{
		return slice2
	}else if len(slice2) <= 0 {
		return slice1
	}

	tmp := make([]string, len(slice1) + len(slice2))
	copy(tmp, slice1)
	copy(tmp[(len(slice1)-1):], slice2)
	return tmp
}