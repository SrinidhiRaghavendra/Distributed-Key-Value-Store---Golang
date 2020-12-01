package main

import "fmt"
import "server"

func main() {
	fmt.Println("hello server xD\n");
	a := server.GetReplicasForKey(4)
	fmt.Println(a)
}
