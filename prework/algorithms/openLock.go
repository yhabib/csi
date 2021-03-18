
import "strings"

func openLock(deadends []string, target string) int {
	startingPoint := "0000"
	visited := make(map[string]bool)
	queue := []string{startingPoint}
	minDistance := 0

	// To make it constant lookup
	dead_ends := make(map[string]bool)
	for _, d := range deadends {
		dead_ends[d] = true
	}

	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			vertex := queue[0]
			queue = queue[1:]

			if vertex == target {
				return minDistance
			}

			if dead_ends[vertex] {
				size--
				continue
			}

			for i := 0; i < len(vertex); i++ {
				n := (int(vertex[i]-'0') + 1) % 10
				p := (int(vertex[i]-'0') - 1 + 10) % 10
				next := vertex[:i] + strconv.Itoa(n) + vertex[i+1:]
				prev := vertex[:i] + strconv.Itoa(p) + vertex[i+1:]
				if !dead_ends[next] && !visited[next] {
					queue = append(queue, next)
					visited[next] = true
				}
				if !dead_ends[prev] && !visited[prev] {
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
