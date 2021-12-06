package util

func RemoveDuplicateStrings(strSlice []string) []string {
	allKeys := make(map[string]struct{})
	list := []string{}
	for _, item := range strSlice {
		if _, ok := allKeys[item]; !ok {
			allKeys[item] = struct{}{}
			list = append(list, item)
		}
	}
	return list
}
