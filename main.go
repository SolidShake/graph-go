package main

import "fmt"

type Vertex struct {
	value string
}

func (n *Vertex) toString() string {
	return fmt.Sprintf("%s ", n.value)
}

type Graph struct {
	vertices []Vertex
	edges    map[Vertex][]Vertex
}

func (g *Graph) addVertex(v Vertex) {
	g.vertices = append(g.vertices, v)
}

func (g *Graph) addEdge(src, dest Vertex) {
	if g.edges == nil {
		g.edges = make(map[Vertex][]Vertex)
	}
	g.edges[src] = append(g.edges[src], dest)
}

func (g *Graph) toString() (graphToString string) {
	for i := 0; i < len(g.vertices); i++ {
		graphToString += g.vertices[i].toString()

		var edgesToString string
		for j := 0; j < len(g.edges[g.vertices[i]]); j++ {
			edgesToString += g.edges[g.vertices[i]][j].toString()
		}

		if len(g.edges[g.vertices[i]]) > 0 {
			graphToString += fmt.Sprintf("-> %s \n", edgesToString)
		}
	}

	return graphToString
}

func main() {
	vA := Vertex{"A"}
	vB := Vertex{"B"}
	vC := Vertex{"C"}
	vD := Vertex{"D"}

	var testGraph Graph

	testGraph.addVertex(vA)
	testGraph.addVertex(vB)
	testGraph.addVertex(vC)
	testGraph.addVertex(vD)

	testGraph.addEdge(vA, vB)
	testGraph.addEdge(vA, vC)
	testGraph.addEdge(vA, vD)
	testGraph.addEdge(vB, vA)
	testGraph.addEdge(vC, vA)
	testGraph.addEdge(vD, vB)

	fmt.Println(testGraph.toString())
}
