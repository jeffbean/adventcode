package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/jeffbean/adventcode/2017/common"
)

type Tree map[string]*node

type node struct {
	parent   *node
	name     string
	weight   int
	children []*node
}

// func (n node) String() string {
// 	var r string
// 	r += fmt.Sprintf("\t%v (%v)", n.name, n.weight)
// 	for _, child := range n.children {
// 		r += fmt.Sprintf("%v\n", child.String())
// 	}
// 	return r
// }

func (n node) String() string {
	var r string
	r += fmt.Sprintf("%v(%v)", n.name, n.weight)
	return r
}

func readInGraph(data []byte) Tree {
	re := regexp.MustCompile(`(?P<node>\w+)\s\((?P<value>\d+)\)(?:\s->\s(?P<children>.*))?`)
	n1 := re.SubexpNames()

	lines := strings.Split(string(data), "\n")
	g := make(map[string]*node, len(lines))
	for _, line := range lines {
		r2 := re.FindAllStringSubmatch(line, -1)[0]

		md := map[string]string{}
		for i, n := range r2 {
			// fmt.Printf("%d. match='%s'\tname='%s'\n", i, n, n1[i])
			md[n1[i]] = n
		}
		val, ok := md["value"]
		if !ok {
			log.Fatalf("number not found on line %s", line)
		}
		num, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("number found: %v\n", num)
		// fmt.Printf("node name: %v\n", md["node"])
		n, ok := g[md["node"]]
		if ok {
			n.weight = num
		} else {
			n = &node{name: md["node"], weight: num}
			g[md["node"]] = n
		}

		children := strings.Split(md["children"], ", ")
		if len(children) > 1 {
			for _, child := range children {
				c, ok := g[child]
				if ok {
					c.parent = n
				} else {
					c = &node{parent: n, name: child}
					g[child] = c
				}
				n.children = append(n.children, c)
			}
		}
	}
	return g
}

func findRoot(t Tree) string {
	for _, node := range t {
		if node.parent == nil {
			return node.name
		}
	}
	return ""
}

func findWeights(t Tree) {
	rootNode := findRoot(t)
	// found := false
	theNode, _ := t[rootNode]
	// ws := make(map[string]int, len(theNode.children))

	fmt.Printf("found rootnode children weights: %v\n", rootNode)
	g := findUnbalancedTower(t, theNode)
	fmt.Printf("\n\tnode: %v, (%v)\n\n", g.name, g.weight)
	// for n, w := range ws {
	// 	fmt.Printf("\tnode %-10v: %-10v\n", n, w)
	// }

}

func findUnbalancedTower(t Tree, nod *node) *node {
	ws := make(map[string]int, len(nod.children))
	wsCount := make(map[int]int, len(nod.children))
	for _, child := range nod.children {
		ws[child.name] = findTowerWeight(child)
		wsCount[ws[child.name]]++
	}
	fmt.Printf("found child tower weights %v: %v\n", nod.name, wsCount)
	for n, w := range wsCount {
		if w != 1 {
			continue
		}
		fmt.Printf("Found the odd one: %v\n", n)
		for nname, val := range ws {
			// fmt.Printf("here we compare %v==%v\n", val, n)
			if val == n {
				fmt.Printf("calling recurse function on unbalenced %q\n", t[nname].name)
				return findUnbalancedTower(t, t[nname])
			}
		}
	}
	return nod
}

func findTowerWeight(node *node) int {
	towerW := 0
	towerW += node.weight
	// fmt.Printf("node: %10s     (%8v)\n", node.name, node.weight)
	for _, child := range node.children {
		// fmt.Printf("%20s (%8v)\n", child.name, child.weight)
		towerW += findTowerWeight(child)
	}
	// fmt.Printf("--------------------%10v\n", towerW)
	return towerW
}

func part1(data []byte) string {
	tree := readInGraph(data)
	return findRoot(tree)
}

func part2(data []byte) {
	tree := readInGraph(data)
	findWeights(tree)
}

func main() {
	data := common.GetInputFile()
	// fmt.Printf("%v\n", part1(data))
	part2(data)
}
