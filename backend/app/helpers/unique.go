package helpers

// UniqStrings returns a list of strings with no duplicates
func UniqStrings(slice []string) []string {
	var keys map[string]struct{} = map[string]struct{}{}
	for i := 0; i < len(slice); i++ {
		if _, ok := keys[slice[i]]; !ok {
			keys[slice[i]] = struct{}{}
		}
	}
	var unique []string = make([]string, 0, len(keys))
	for value := range keys {
		unique = append(unique, value)
	}
	return unique
}
