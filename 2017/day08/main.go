package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/jeffbean/adventcode/2017/common"
)

func handleIncDec(op string, in, by int) (out int) {
	switch op {
	case "inc":
		out = in + by
	case "dec":
		out = in - by
	}
	// fmt.Printf("changed %v to %v\n", in, out)
	return out
}

func part1(data []byte) int {
	re := regexp.MustCompile(`(?P<reg>\w+)\s(?P<ind_dec>\w+)\s(?P<count_value>[-\d]+)\s\w+\s(?P<conditional_register>\w+)\s(?P<operation>[>!=<]{1,2})\s(?P<op_value>[-\d]+)`)
	n1 := re.SubexpNames()
	max := -99999
	min := 99999

	lines := strings.Split(string(data), "\n")
	g := make(map[string]int, len(lines))
	for _, line := range lines {
		r2 := re.FindAllStringSubmatch(line, -1)[0]

		md := map[string]string{}
		for i, n := range r2 {
			// fmt.Printf("%d. match='%s'\tname='%s'\n", i, n, n1[i])
			md[n1[i]] = n
		}
		val, ok := g[md["conditional_register"]]
		if !ok {
			g[md["conditional_register"]] = 0
			val = 0
		}
		// fmt.Printf("process reg %q op %q", md["reg"], )
		opVal, err := strconv.Atoi(md["op_value"])
		if err != nil {
			log.Fatal(err)
		}
		countVal, err := strconv.Atoi(md["count_value"])
		if err != nil {
			log.Fatal(err)
		}
		switch md["operation"] {
		case ">":
			if val > opVal {
				g[md["reg"]] = handleIncDec(md["ind_dec"], g[md["reg"]], countVal)
			}
		case "<":
			if val < opVal {
				g[md["reg"]] = handleIncDec(md["ind_dec"], g[md["reg"]], countVal)
			}
		case "==":
			if val == opVal {
				g[md["reg"]] = handleIncDec(md["ind_dec"], g[md["reg"]], countVal)
			}
		case "!=":
			if val != opVal {
				g[md["reg"]] = handleIncDec(md["ind_dec"], g[md["reg"]], countVal)
			}
		case ">=":
			if val >= opVal {
				g[md["reg"]] = handleIncDec(md["ind_dec"], g[md["reg"]], countVal)
			}
		case "<=":
			if val <= opVal {
				g[md["reg"]] = handleIncDec(md["ind_dec"], g[md["reg"]], countVal)
			}
		default:
			log.Fatal("unknown operation")
		}
		if g[md["reg"]] < min {
			min = g[md["reg"]]
		}
		if g[md["reg"]] > max {
			max = g[md["reg"]]
		}

	}

	fmt.Printf("max: %v\n", max)
	fmt.Printf("min: %v\n", min)
	return 0
}

func part2() int {
	return 0
}

func main() {
	data := common.GetInputFile()
	// log.Println(data)
	fmt.Printf("%v\n", part1(data))
	// fmt.Printf("%v\n", part2())
}
