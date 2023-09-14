package hw1
// У этой функции time complexity O(2n)=O(n), так как у нас простые два цикла идущие друг за другом.
func Compare(slice1, slice2 []int) bool {
	map1 := make(map[int]int,len(slice1))
	map2 := make(map[int]int,len(slice2))
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		map1[slice1[i]]++
		map2[slice2[i]]++
	}
	for key,val := range map1 {
		if val!= map2[key] {
            return false
        }
    }
	return true
}
