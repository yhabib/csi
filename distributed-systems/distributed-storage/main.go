package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	memory := make(map[string]string)
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Distributed Storage")
	fmt.Println("---------------------")

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	go func() {
		for {
			fmt.Print("-> ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimRight(text, "\n")
			strs := strings.SplitN(text, " ", 2)

			if len(strs) < 2 {
				panic("Missing arguments")
			}

			switch action := strs[0]; action {
			case "set":
				keys := strings.SplitN(strs[1], "=", 2)
				fmt.Println("setting " + keys[0] + " to " + keys[1])
				memory[keys[0]] = keys[1]
			case "get":
				value := memory[strs[1]]
				fmt.Println(strs[1] + ": " + value)
			}
		}
	}()

	<-done
	fmt.Println("exiting")
}
