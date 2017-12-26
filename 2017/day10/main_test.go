package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	k := NewKnot(5)

	k.twitshHashLengths([]int{3, 4, 1, 5})

	fmt.Println(k)

	part1([]byte("3,4,1,5"))
}

func TestPart2(t *testing.T) {

}
