package transpiler

func expand(s string, f func(string) string) string {
	ret := ""
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if c != "$" {
			ret += c
		} else {
			if i+len("$animal") <= len(s) {
				ret += f("mew")
				i += len("$animal") - 1
			}
		}
	}
	return ret
}

func Transpiler(s string) string {
	wordSet := map[string]string{
		"mew":    "cat",
		"bow":    "dog",
		"ribbit": "frog",
		"squeak": "mouse",
	}
	return wordSet[s]
}
