package godiffer

func contains(arr []byte, b byte) bool {
	for _, v := range arr {
		if v == b {
			return true
		}
	}
	return false
}

// func Complement(str1, str2 string) string {

// }

func Union(str1, str2 string) string {
	strArr1 := []byte(str1)
	strArr2 := []byte(str2)
	result := []byte{}
	for _, v := range strArr1 {
		if !contains(result, v) {
			result = append(result, v)
		}
	}
	for _, v := range strArr2 {
		if !contains(result, v) {
			result = append(result, v)
		}
	}
	return string(result)
}

func Intersection(str1, str2 string) string {
	strArr1 := []byte(str1)
	strArr2 := []byte(str2)
	result := []byte{}
	for _, v := range strArr1 {
		if contains(strArr2, v) {
			result = append(result, v)
		}
	}
	return string(result)
}

func Difference(str1, str2 string) string {
	strArr1 := []byte(str1)
	strArr2 := []byte(str2)
	result := []byte{}
	for _, v := range strArr1 {
		if !contains(strArr2, v) {
			result = append(result, v)
		}
	}
	return string(result)
}

func SymetricDifference(str1, str2 string) string {
	strArr1 := []byte(str1)
	strArr2 := []byte(str2)
	result := []byte{}
	for _, v := range strArr1 {
		if !contains(strArr2, v) {
			result = append(result, v)
		}
	}
	for _, v := range strArr2 {
		if !contains(strArr1, v) {
			result = append(result, v)
		}
	}
	return string(result)
}
