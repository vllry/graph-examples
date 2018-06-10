package main

type vertex struct {
	id string
	adjacentTo map[string]*vertex
	adjacentFrom map[string]*vertex
}

func (v1 *vertex) adjacentDirected(v2 *vertex) {
	v1.adjacentTo[v2.id] = v2
	v2.adjacentFrom[v1.id] = v1
}

func (v1 *vertex) adjacentUndirected(v2 *vertex) {
	v1.adjacentDirected(v2)
	v2.adjacentDirected(v1)
}

type graph struct {
	vertices map[string]*vertex
}

// TODO: wrap creation, and handle adjacency setting.s
func (g *graph) add(v *vertex) {
	g.vertices[v.id] = v
}