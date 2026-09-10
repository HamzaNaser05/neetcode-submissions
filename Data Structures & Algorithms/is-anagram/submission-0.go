func isAnagram(s string, t string) bool {
	s_map := map[rune]int{}
	t_map := map[rune]int{}

	for _, value := range s{
		s_map[value]++
	}
	for _, value := range t{
		t_map[value]++
	}

	if len(s_map) != len(t_map){
		return false
	}

	for key,_ := range s_map{
		if s_map[key] != t_map[key]{
			return false
		}
	}
	return true
}
