package utils

func IndexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func Mod(a, b int) int {
	return (a%b + b) % b
}
