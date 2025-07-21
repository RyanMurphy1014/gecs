package main

import "fmt"

func main() {
	ecs, e := NewECS(WithAttributes(true))
	fmt.Println(ecs.EntToArche[e])
}
