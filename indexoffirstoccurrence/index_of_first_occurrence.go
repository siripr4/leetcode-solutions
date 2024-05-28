package indexoffirstoccurrence

func main() {

}

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	i := 0
	j := len(needle) - 1

	for j < len(haystack) {
		for j < len(haystack)-1 && haystack[i] != needle[0] {
			i++
			j++
		}

		if haystack[i:j-1] == needle {
			return i
		} else {
			i++
			j++
		}
	}
	return -1
}
