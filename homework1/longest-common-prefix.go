package hw1


// эта функция имеет  time complexity O(n^2), так как мы имее два цикла один из которых вложен в другой. 
func LongestCommonPrefix(strs []string) string {
	longestComPrefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for !isPrefix(strs[i], longestComPrefix) {
			longestComPrefix = longestComPrefix[:len(longestComPrefix)-1]
		}
	}
	return longestComPrefix
}

func isPrefix(str string, substr string) bool {
	if len(str) >= len(substr) && str[:len(substr)] == substr {
		return true
	}
	return false
}
