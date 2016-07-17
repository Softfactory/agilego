package ai

import (
	"container/heap"
	"fmt"
)

// Vertex 꼭짓점은 어떤 자료형도 받을 수 있어야 한다.
type Vertex interface{}

// Edge 변은 두개의 꼭짓점을 연결하고 속성값을 갖는다.
type Edge struct {
	start Vertex
	end   Vertex
	value Property
}

/*
Property 변의 속성은 어떤 자료형이 될지 알 수 없다.
 TODO(me): 나중에 자료형을 지정해야 한다.
*/
type Property interface{}

// Vertices Heap인터페이스를 구현합니다.
type Vertices []Vertex

func (v Vertices) Len() int {
	return len(v)
}

func (v Vertices) Less(i, j int) bool {
	return v[i].(int) < v[j].(int)
}

func (v Vertices) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

// Push Vertex를 추가한다.
func (v *Vertices) Push(element interface{}) {
	if !v.exist(element) {
		*v = append(*v, element.(Vertex))
	}
}

// Pop 제일 상단의 Vertex를 가져온다.
func (v *Vertices) Pop() interface{} {
	old := *v
	n := len(old)
	element := old[n-1]
	*v = old[0 : n-1]
	return element
}

func (v Vertices) exist(i Vertex) bool {
	for vertex := range v {
		if vertex == i {
			return true
		}
	}
	return false
}

/*
Graph  두개 이상의 꼭짓점과 그 꼭짓점들을 연결한 변을 갖는다.
*/
type Graph struct {
	// Vertices map[Vertex]*list.List
	vertices Vertices
	oriented bool
}

// InitGraph 그래프를 생성합니다.
func InitGraph() Graph {
	return Graph{
		make(Vertices, 0),
		false,
	}
}

func (graph *Graph) add(edge Edge) {
	v := &(graph.vertices)
	heap.Init(v)
	// if !v.exist(edge.start) {
	heap.Push(v, edge.start)
	// }
	heap.Push(v, edge.end)
	fmt.Print(v)
}
