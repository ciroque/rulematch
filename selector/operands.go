package selector

// and returns true if all expected values are found in actual.
func and(expected, actual []string) bool {
	for _, exp := range expected {
		found := false
		for _, act := range actual {
			if exp == act {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// not returns true if no expected values are found in actual.
func not(expected, actual []string) bool {
	for _, exp := range expected {
		for _, act := range actual {
			if exp == act {
				return false
			}
		}
	}
	return true
}

// or returns true if any expected value is found in actual.
func or(expected, actual []string) bool {
	for _, exp := range expected {
		for _, act := range actual {
			if exp == act {
				return true
			}
		}
	}
	return false
}
