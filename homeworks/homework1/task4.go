package hw1

type MyStruct struct {
	slice1 []int
}

// У этой функции  time complexity О(n^2), так как мы имее два цикла один из которых вложен в другой.
func (m *MyStruct) sort() []int {
	for i := 0; i < len(m.slice1)-1; i++ {
		swapped := false
		for j := 0; j < len(m.slice1)-1-i; j++ {
			if m.slice1[j] > m.slice1[j+1] {
				m.slice1[j], m.slice1[j+1] = m.slice1[j+1], m.slice1[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return m.slice1
}
