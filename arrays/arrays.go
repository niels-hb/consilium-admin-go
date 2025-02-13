package arrays

func Contains(haystack []string, needle string) bool {
	for _, element := range haystack {
		if element == needle {
			return true
		}
	}

	return false
}
