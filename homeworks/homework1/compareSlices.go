package hw1

func sort(slice1 []int) []int {
	for i := 0; i < len(slice1)-1; i++ {
		swapped := false
		for j := 0; j < len(slice1)-1-i; j++ {
			if slice1[j] > slice1[j+1] {
				slice1[j], slice1[j+1] = slice1[j+1], slice1[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return slice1
}

func Compare(slice1, slice2 []int) bool {
	slice1 = sort(slice1)
	slice2 = sort(slice2)
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
