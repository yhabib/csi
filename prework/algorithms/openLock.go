
import "strings"

// Notes:
//  - Initially I had a map for visited and another one for dead_ends but they can be merged into one

func openLock(deadends []string, target string) int {
	startingPoint := "0000"
	visited := make(map[string]bool)
	queue := []string{startingPoint}
	minDistance := 0

	for _, d := range deadends {
		visited[d] = true
	}

	if visited[startingPoint] {
		return -1
	}

	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			vertex := queue[0]
			queue = queue[1:]

			if vertex == target {
				return minDistance
			}

			for i := 0; i < len(vertex); i++ {
				n := (int(vertex[i]-'0') + 1) % 10
				p := (int(vertex[i]-'0') - 1 + 10) % 10
				next := vertex[:i] + strconv.Itoa(n) + vertex[i+1:]
				prev := vertex[:i] + strconv.Itoa(p) + vertex[i+1:]
				if !visited[next] {
					queue = append(queue, next)
					visited[next] = true
				}
				if !visited[prev] {
					queue = append(queue, prev)
					visited[prev] = true
				}
			}
			size--
		}
		minDistance++
	}
	return -1
}
