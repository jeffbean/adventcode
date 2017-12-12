package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	g := Graph{}
	z := g.newNode(0)
	one := g.newNode(1)
	two := g.newNode(2)
	three := g.newNode(3)
	four := g.newNode(4)
	five := g.newNode(5)
	six := g.newNode(6)

	g.newEdge(z, two)
	g.newEdge(one, one)

	g.newEdge(two, z)
	g.newEdge(two, three)
	g.newEdge(two, four)

	g.newEdge(three, two)
	g.newEdge(three, four)

	g.newEdge(four, two)
	g.newEdge(four, three)
	g.newEdge(four, six)

	g.newEdge(five, six)

	g.newEdge(six, four)
	g.newEdge(six, five)

	for _, n := range g.nodes {
		t.Logf("%#v", n)
	}
	nodes := make(map[int]struct{})
	n := g.findInGroup(0, nodes)
	t.Log(n)
	assert.Equal(t, 6, len(n))

	// gr := &Graph{
	// 	nodes: []*Node{},
	// }
	// testData := []byte(`0 <-> 1
	// 	1 <-> 1
	// 	2 <-> 0`)
	// assert.Equal(t, gr, readInGraph(testData))
}
func TestPart2(t *testing.T) {

}
