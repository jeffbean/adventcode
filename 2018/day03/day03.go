package day03

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func part1(in []string) int {
	pices := []*fabric{}
	for _, l := range in {
		f, err := parseLine(l)
		if err != nil {
			log.Fatal(err)
		}
		pices = append(pices, f)
	}
	grid := calcOverlap(pices)
	return gridOverlab(grid)
}

func part2(in []string) string {
	pices := []*fabric{}
	for _, l := range in {
		f, err := parseLine(l)
		if err != nil {
			log.Fatal(err)
		}
		pices = append(pices, f)
	}
	grid := calcOverlap(pices)
	return theOne(grid).id
}

type fabric struct {
	id         string
	posX, posY int
	l, w       int
}

func parseLine(line string) (*fabric, error) {
	parts := strings.Split(line, " ")
	// 0  1 2        3
	// #8 @ 192,643: 16x15
	f := &fabric{}
	posParts := strings.Split(strings.TrimRight(parts[2], ":"), ",")
	f.id = parts[0]
	f.posX = strToInt(posParts[0])
	f.posY = strToInt(posParts[1])

	sizeParts := strings.Split(parts[3], "x")
	f.w = strToInt(sizeParts[0])
	f.l = strToInt(sizeParts[1])
	return f, nil
}

type track struct {
	count int
	fabs  []*fabric
}

type grid [][]*track

func calcOverlap(pcs []*fabric) grid {
	grid := make([][]*track, 999)
	for i := range grid {
		grid[i] = make([]*track, 999)
	}

	for _, fab := range pcs {
		for i := 0; i < fab.w; i++ {
			for j := 0; j < fab.l; j++ {
				tracker := grid[fab.posX+i][fab.posY+j]
				if tracker == nil {
					grid[fab.posX+i][fab.posY+j] = &track{
						fabs:  []*fabric{fab},
						count: 1,
					}
				} else {
					tracker.count++
					tracker.fabs = append(tracker.fabs, fab)
				}
			}
		}
	}

	return grid
}
func gridOverlab(g grid) int {
	var totalOverlap int
	for i := range g {
		for j := range g[i] {
			t := g[i][j]

			if t != nil && t.count > 1 {
				totalOverlap++
			}
		}
	}
	return totalOverlap
}

func theOne(g grid) *fabric {
	hmm := map[*fabric]bool{}
	for i := range g {
		for j := range g[i] {
			t := g[i][j]

			if t != nil && len(t.fabs) != 1 {
				for _, thing := range t.fabs {
					hmm[thing] = false
				}
			}
			if t != nil && len(t.fabs) == 1 {
				_, ok := hmm[t.fabs[0]]
				if !ok {
					hmm[t.fabs[0]] = true
				}
			}

		}
	}
	found := []*fabric{}
	for f, value := range hmm {
		if value {
			found = append(found, f)
		}
	}
	fmt.Printf("%#v\n", found)
	return found[0]
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
