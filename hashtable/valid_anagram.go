package hashtable

import "fmt"

func isAnagram(s string, t string) bool {
	m := make(map[int32]int, 26)
	for _, letter := range s {
		m[letter-'a']++
	}

	for _, letter := range t {
		m[letter-'a']--
	}

	for _, n := range m {
		if n != 0 {
			return false
		}
	}
	return true
}

func TestIsAnagram() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}
