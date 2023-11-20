/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"example/qufi/cmd"
	"math/rand"
	"time"
)

func main() {
	//seeding for random number generator
	rand.NewSource(time.Now().UnixNano())
	cmd.Execute()
}
