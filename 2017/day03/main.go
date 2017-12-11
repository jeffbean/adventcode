package main

import (
	"fmt"
	"log"
	"math"
)

func findCorrdinates(desiredSprialNumber float64) (float64, float64) {
	// This is calculation based on x,y corrdinates thanks to the hint in the problem
	k := math.Ceil((math.Sqrt(desiredSprialNumber) - 1) / 2)
	t := 2*k + 1
	m := math.Pow(t, 2)
	t = t - 1
	// f(n) = 4 n^2 + n + 1
	// log.Printf("k: %v, t: %v, m: %v, t: %v", k, t, m, t)
	if desiredSprialNumber >= m-t {
		return k - (m - desiredSprialNumber), -k
	}
	m = m - t
	// log.Printf("k: %v, t: %v, m: %v, t: %v", k, t, m, t)
	if desiredSprialNumber >= m-t {
		// log.Printf("k: %v, t: %v, m: %v, t: %v", k, t, m, t)
		return -k, -k + (m - desiredSprialNumber)
	}
	m = m - t
	if desiredSprialNumber >= m-t {
		return -k + (m - desiredSprialNumber), k
	}
	return k, k - (m - desiredSprialNumber - t)
}

func part1(desiredSprialNumber float64) float64 {
	x, y := findCorrdinates(desiredSprialNumber)
	return math.Abs(x) + math.Abs(y)
}

type grid struct {
	m [][]int
}

func (g *grid) getValue(posX, posY int) int {
	var sum = 0
	for x := posX - 1; x <= posX+1; x++ {
		for y := posY - 1; y <= posY+1; y++ {
			if x == posX && y == posY {
				continue
			}
			sum += g.m[x][y]
			// log.Printf("summing (%v,%v): x, y, sum: (%v, %v) -> %v", posX, posY, x, y, sum)
		}
	}
	return sum
}

func buildGrid(x, y int) *grid {
	grid := &grid{}
	a := make([][]int, x*2)
	for i := range a {
		a[i] = make([]int, y*2)
	}
	grid.m = a
	grid.m[x][y] = 1
	return grid
}

func (g *grid) takeStep(x, y int, d string) int {
	// log.Printf("taking step to the %v", d)
	val := g.getValue(x, y)

	fmt.Printf("Setting value: [%-3v] for cor (%-2v, %-2v) headed %v\n", val, x, y, d)
	g.m[x][y] = val
	return val
}

func part2BuildSpiral(gridS int, theNumber int) (int, int) {
	grid := buildGrid(gridS, gridS)
	x := gridS
	y := gridS

	minX := x
	maxX := x
	minY := y
	maxY := y
	dir := "right"
	val := 0

	for val < theNumber {
		// change directionssss.....
		switch dir {
		case "right":
			x++
			// we are movig right one on the x axis
			val = grid.takeStep(x, y, dir)
			log.Printf("(%v, %v) minX, MaxX (%v, %v) - minY MaxY (%v, %v)", x, y, minX, maxX, minY, maxY)

			if x < maxX+1 {
				break
			}
			maxX = x

			dir = "up"
			break
		case "up":
			y++
			val = grid.takeStep(x, y, dir)
			log.Printf("(%v, %v) minX, MaxX (%v, %v) - minY MaxY (%v, %v)", x, y, minX, maxX, minY, maxY)

			if y < maxY+1 {
				break
			}
			maxY = y

			dir = "left"
			break
		case "left":
			x--
			val = grid.takeStep(x, y, dir)
			// log.Printf("(%v, %v) minX, MaxX (%v, %v) - minY MaxY (%v, %v)", x, y, minX, maxX, minY, maxY)

			if x > minX-1 {
				break
			}
			minX = x

			dir = "down"
			break
		case "down":
			y--
			//
			val = grid.takeStep(x, y, dir)
			// log.Printf("(%v, %v) minX, MaxX (%v, %v) - minY MaxY (%v, %v)", x, y, minX, maxX, minY, maxY)

			if y > minY-1 {
				break
			}
			minY = y

			dir = "right"
			break
		}
		if grid.getValue(x, y) > theNumber {
			return x, y
		}
	}

	return x, y
}

func part2(input int) float64 {
	x, y := part2BuildSpiral(512, input)
	return math.Abs(float64(x)-512.0) + math.Abs(float64(y)-512.0)
}

func main() {
	fmt.Println(part2(265149 + 10))
}
