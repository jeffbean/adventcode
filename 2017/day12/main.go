package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jeffbean/adventcode/2017/common"
)

type Graph struct {
	nodes []*Node
}

type Node struct {
	index int // just tracking what node we are in the loops
	edges []Edge
	value int
}

type Edge struct {
	end *Node
}

func (n Node) String() string {
	stringBuilder := ""
	stringBuilder += fmt.Sprintf("%v -> [", n.value)
	for _, e := range n.edges {
		stringBuilder += fmt.Sprintf("%v, ", e.String())
	}
	stringBuilder += "]"

	return stringBuilder
}

func (e Edge) String() string {
	return fmt.Sprintf("%v", e.end.value)
}

func (g *Graph) newNode(value int) *Node {
	// make the node and add to graph
	for _, n := range g.nodes {
		if n.value == value {
			return n
		}
	}
	node := &Node{index: len(g.nodes), value: value}
	g.nodes = append(g.nodes, node)
	return node
}

func (g *Graph) newEdge(from *Node, to *Node) error {
	// does the node esixt in the graph
	if from != g.nodes[from.index] {
		return fmt.Errorf("the from node [%v] is not in the graph", from)
	}
	if to != g.nodes[to.index] {
		return fmt.Errorf("the to node [%v] is not in the graph", to)
	}
	edge := Edge{end: to}
	for _, e := range from.edges {
		if e == edge {
			// log.Printf("found the edge already from [%v] -> [%v]", e, to.value)
			return nil
		}
	}
	from.edges = append(from.edges, edge)

	edgeFrom := Edge{end: from}
	for _, e := range to.edges {
		if e == edgeFrom {
			// log.Printf("found the edge already from [%v] -> [%v]", e, to.value)
			return nil
		}
	}

	to.edges = append(to.edges, edgeFrom)
	return nil
}

func (g *Graph) findInGroup(groupIndex int, nodeIndex int, nodes map[int]int) map[int]int {
	node := g.nodes[nodeIndex]

	// log.Printf("node list %v\n\n", nodes)
	// log.Printf("Found indexed  node [%v]\n", node)

	if _, ok := nodes[node.index]; ok {
		return nodes
	} // we need to basically check to see if this node already has a thing
	nodes[node.index] = groupIndex

	for _, edge := range node.edges {
		if edge.end.index == node.index {
			continue
		}
		// log.Printf("looking at edge %v", edge.end.index)
		nodes = g.findInGroup(groupIndex, edge.end.index, nodes)
		// g.findInGroup(edge.end.index, nodes)
	}
	return nodes
}

func readInGraph(data []byte) *Graph {
	var g = &Graph{}
	for _, line := range strings.Split(string(data), "\n") {
		subLine := strings.Split(line, "<->")
		if len(subLine) != 2 {
			log.Fatalf("Failed to parse line %v", line)
		}

		num, err := strconv.Atoi(strings.TrimSpace(subLine[0]))
		if err != nil {
			log.Fatal(err)
		}
		theOneNode := g.newNode(num)

		nodes, err := common.StringToInts(subLine[1], ",")
		if err != nil {
			log.Fatal(err)
		}
		// log.Printf("file input line nodes to [%v] <- %v", num, nodes)
		for _, n := range nodes {
			o := g.newNode(n)
			err := g.newEdge(theOneNode, o)
			if err != nil {
				log.Fatal(err)
			}
		}

	}
	return g
}

func part1() int {
	g := readInGraph(common.GetInputFile())
	// fmt.Print(g)
	thing := make(map[int]int)
	nums := g.findInGroup(0, 0, thing)
	// fmt.Println(nums)
	return len(nums)
}

func part2() int {
	g := readInGraph(common.GetInputFile())
	// fmt.Print(g)
	nums := make(map[int]int)
	for i := 0; i < len(g.nodes); i++ {
		nums = g.findInGroup(i, i, nums)

	}
	fmt.Println(nums)
	mapThing := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, ok := mapThing[num]; ok {
			continue
		}
		mapThing[num] = struct{}{}
	}
	return len(mapThing)
}

func main() {
	// data := common.GetInputFile()
	// log.Println(data)
	fmt.Printf("%v\n", part1())
	fmt.Printf("%v\n", part2())
}
