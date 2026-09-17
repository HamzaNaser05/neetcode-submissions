import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	var result [][]string
	if len(strs) == 0 {
		return result
	}

	groups := map[string][]string{}
	for _, str := range strs {
		runes := []rune(str)
		slices.Sort(runes)
		sortedStr := string(runes)
		groups[sortedStr] = append(groups[sortedStr], str)
	}
	for key:= range groups {
		result = append(result, groups[key])
	}
	return result
	//loop over the str[]
	//sort each entry
	//add it to the map key
	//loop to add the entries the the [][]str


}
