package graff

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectedGraphDOTGraph(t *testing.T) {
	graph := NewDirectedGraph()
	graph.AddNodes("A", "B", "C", "D")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "C")
	graph.AddEdge("D", "B")

	output := graph.DOTGraph()
	output = strings.Replace(output, "\t", "", -1)
	output = strings.TrimSpace(output)

	assert.Equal(t, `digraph G {
A;
B;
C;
D;
A -> C
B -> C
D -> B
}`, output, "graph.DOTGraph() output is incorrect")
}