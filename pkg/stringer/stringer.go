package stringer

import "strconv"

func Reverse(input string) (result string) {
	for _, c := range input {
		result = string(c) + result
	}
	return result
}

func Inspect(input string, isDigits bool) (count int, kind string) {
	if isDigits {
		return InspectNumbers(input), "digits"
	}
	return len(input), "char"
}

func InspectNumbers(input string) (count int) {
	for _, c := range input {
		_, err := strconv.Atoi(string(c))
		if err == nil {
			count++
		}
	}
	return count
}
