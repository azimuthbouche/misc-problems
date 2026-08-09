// variation of fizzbuzz, convert number into corresponding raindrop sounds
// divis by 3, 5, or 7, goes Pling, Plang, or Plong
// if not divis by above, returns num as s tr

package raindrops

import "strconv"

var sounds = []struct {
	factor int
	sound  string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(n int) string {
	result := ""
	for _, s := range sounds {
		if n%s.factor == 0 {
			result += s.sound
		}
	}
	if result == "" {
		return strconv.Itoa(n)
	}
	return result
}
