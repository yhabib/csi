func canVisitAllRooms(rooms [][]int) bool {
	visitedRooms := make(map[int]bool)
	queue := []int{0}
	numberOfRoomsVisited := 0

	for len(queue) > 0 {
		key := queue[0]
		queue = queue[1:]

		if visitedRooms[key] {
			continue
		}
		numberOfRoomsVisited++
		if numberOfRoomsVisited == len(rooms) {
			return true
		}
		visitedRooms[key] = true
		queue = append(queue, rooms[key]...)
	}
	return numberOfRoomsVisited == len(rooms)

}