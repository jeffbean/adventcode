package day02

func part1(in []string) int {
	var (
		totalTwice  int
		totalThrice int
	)
	for _, s := range in {
		tw, th := processID(s)
		totalTwice += tw
		totalThrice += th
	}

	return totalTwice * totalThrice
}

func processID(id string) (twice, thrice int) {
	counters := make(map[rune]int)

	for _, c := range id {
		counters[c]++
	}

	for _, num := range counters {
		switch num {
		case 2:
			twice = 1
		case 3:
			thrice = 1
		}
	}

	return twice, thrice
}
