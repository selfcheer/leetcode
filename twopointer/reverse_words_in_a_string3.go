package twopointer

import "fmt"

func reverseWords(s string) string {
	n := len(s)
	a := make([]rune, 0, n)
	for slow, fast,temp := 0, 0, 0; fast < n; fast ++ {
		if s[fast] == ' ' || fast == n - 1 {
			if fast == n - 1 {
				temp = fast
			}else {
				temp = fast - 1
			}

			for temp >= slow {
				a = append(a, rune(s[temp]))
				temp --
			}
			if  s[fast] == ' ' {
				a = append(a, rune(s[fast]))
			}
			slow = fast + 1
		}

	}
	return string(a)
}


func TestReverseWords() {
	s := "Let's take LeetCode contest"
	fmt.Println(s)
	fmt.Println(reverseWords(s))
}
