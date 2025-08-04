package graph

type Graph struct {
	adjList map[string][]string
}

func (g *Graph) AddVertex(vertex string) bool {
	if !g.HasVertex(vertex) {
		g.adjList[vertex] = []string{}
		return true
	}
	return false
}

func (g *Graph) AddEdge(v1, v2 string) bool {
	if g.HasVertex(v1) && g.HasVertex(v2) {
		g.adjList[v1] = append(g.adjList[v1], v2)
		g.adjList[v2] = append(g.adjList[v2], v1)
		return true
	}
	return false
}

func (g *Graph) RemoveEdge(v1, v2 string) bool {
	if g.HasVertex(v1) && g.HasVertex(v2) {
		g.adjList[v1] = remove(g.adjList[v1], v2)
		g.adjList[v2] = remove(g.adjList[v2], v1)
		return true
	}
	return false
}

func (g *Graph) RemvoeVertex(v string) bool {
	if !g.HasVertex(v) {
		return false
	}

	for k := range g.adjList {
		g.adjList[k] = remove(g.adjList[k], v)
		if k == v {
			delete(g.adjList, k)
		}
	}
	return true

}

func (g *Graph) HasVertex(v string) bool {
	return g.adjList[v] != nil
}

func remove(slice []string, value string) []string {
	for i, s := range slice {
		if s == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return []string{}
}
