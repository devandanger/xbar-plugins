//usr/bin/env go run $0 $@; exit

// # <xbar.title>My Azure Devops Pull Requests</xbar.title>
// # <xbar.version>v1.0</xbar.version>
// # <xbar.author>Evan Anger</xbar.author>
// # <xbar.author.github>devandanger</xbar.author.github>
// # <xbar.desc>Access Pull Requests of Interests</xbar.desc>
// # <xbar.dependencies>swift</xbar.dependencies>

package main

import (
	"fmt"
	"time"
)

func main() {
	printOutput()
}

func printOutput() {
	fmt.Println("")
	fmt.Println("My Pull Requests")
	fmt.Println("---")
	fmt.Println("Assigned Pull Requests")
	fmt.Println("---")
	val := DoneAsync()
	fmt.Println("Done is running ...")
	fmt.Println(<-val)
}

func DoneAsync() chan int {
	r := make(chan int)
	fmt.Println("Warming up ...")
	go func() {
		time.Sleep(3 * time.Second)
		r <- 1
		fmt.Println("Done ...")
	}()
	return r
}
