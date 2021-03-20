func areWordsDiffByOne(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	var diff int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff == 1
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	queue := []string{beginWord}
	visited := make(map[string]bool)
	minDistance := 1
	visited[beginWord] = true

	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			word := queue[0]
			queue = queue[1:]
			if word == endWord {
				return minDistance
			}
			for _, w := range wordList {
				if areWordsDiffByOne(word, w) && !visited[w] {
					visited[w] = true
					queue = append(queue, w)
				}
			}
			size--
		}
		minDistance++
	}
	return 0
}