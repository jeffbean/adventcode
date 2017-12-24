package main

import "testing"

var testString = `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`

func TestPart1(t *testing.T) {
	// f := readInGraph([]byte(testString))
	// t.Log(f)
	// for _, node := range f {
	// 	fmt.Println(node)
	// }
}

func TestPart2(t *testing.T) {
	f := readInGraph([]byte(testString))
	findWeights(f)
}
