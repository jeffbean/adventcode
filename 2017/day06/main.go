package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jeffbean/adventcode/2017/common"
)

const _maxLoopCount = 100

type blocks struct {
	banks []int
}

func (b *blocks) String() string {
	return fmt.Sprintf("%v", b.banks)
}

func readInBlocks(data []byte) (a []int) {
	for _, line := range strings.Split(string(data), "\n") {
		num, err := common.StringToInts(line, "\t")
		if err != nil {
			log.Fatal(err)
		}
		return num
	}
	return nil
}

func findLargest(a blocks) int {
	max := 0
	maxI := 0
	for i, n := range a.banks {
		if n > max {
			max = n
			maxI = i
		}
	}
	return maxI
}

func redistibuteBlocks(b blocks) int {
	loopCount := 1
	foundBlocks := make(map[string]struct{})
	for i := 2; i < _maxLoopCount; i++ {
		if i >= len(b.banks) {
			i = 0
		}
		largestBlockIndex := findLargest(b)        // found the largest
		largestValue := b.banks[largestBlockIndex] // get that value to spread
		b.banks[largestBlockIndex] = 0
		// go to the next index and add one subtracting from the value
		for j := largestBlockIndex + 1; largestValue > 0; j++ {
			if j >= len(b.banks) {
				j = 0
			}
			// log.Printf("redistribute banks: j: %v , value: %v ", j, largestValue)
			b.banks[j]++
			largestValue--
			// log.Printf("after banks: j: %v , value: %v ", j, largestValue)

		}
		log.Printf("loop [%v] : %v ", loopCount, b.String())

		if _, ok := foundBlocks[b.String()]; ok {
			return loopCount
		}
		foundBlocks[b.String()] = struct{}{}
		loopCount++
	}

	return loopCount
}

func part1() int {
	part1Blocks := blocks{
		banks: readInBlocks(common.GetInputFile()),
	}
	return redistibuteBlocks(part1Blocks)
}

func part2() {
}

func main() {
	fmt.Printf("%v\n", part1())
}
