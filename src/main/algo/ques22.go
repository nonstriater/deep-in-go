package algo

func Solution(s string) int {
	li := 0
	ret := 0
	m := make(map[rune]int, 26)
	for i := 0; i < len(s); i++ {
		if _, ok := m[rune(s[i])]; ok {
			li = m[rune(s[i])] + 1
		}
		if (i - li + 1) > ret {
			ret = i - li + 1
		}
		m[rune(s[i])] = i
	}
	return ret
}



