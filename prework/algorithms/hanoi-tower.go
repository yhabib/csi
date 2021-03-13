package main

import "fmt"

func moveTower(height int, fromPole string, toPole string, withPole string) {
	if height >= 1 {
		moveTower(height-1, fromPole, withPole, toPole)
		moveDisk(fromPole, toPole)
		moveTower(height-1, withPole, toPole, fromPole)
	}
}

func moveDisk(fromPole string, toPole string) {
	fmt.Println("Movig disc from pole", fromPole, "to pole", toPole)
}

func main() {
	moveTower(2, "A", "B", "C")
}
