package ai

import (
	"testing"
)

func TestGraph(t *testing.T) {
	//새 그래프를 생성한다.
	graph := InitGraph()

	//새로 생성된 그래프의 꼭짓점수는 0이다.
	if len(graph.vertices) != 0 {
		t.Errorf("그래프는 초기화시 꼭짓점이 0개여야 합니다. %d", len(graph.vertices))
	}

	//두개의 꼭짓점과 속성을 갖는 변을 추가한다.
	graph.add(Edge{1, 2, "blue"})
	if len(graph.vertices) != 2 {
		t.Errorf("엣지가 등록되면 꼭짓점은 2개여야 합니다. %d", len(graph.vertices))
	}

	//같은 꼭짓점을 추가하면 제외한다.
	graph.add(Edge{1, 5, "red"})
	if len(graph.vertices) != 3 {
		t.Errorf("엣지가 추가되면 꼭짓점은 3개여야 합니다. %d", len(graph.vertices))
	}

	t.Log(graph.vertices)

	//같은 꼭짓점을 추가하면 제외한다.
	graph.add(Edge{4, 7, "red"})
	if len(graph.vertices) != 5 {
		t.Errorf("엣지가 추가되면 꼭짓점은 5개여야 합니다. %d", len(graph.vertices))
	}

}

//
// graph := NewGraph()
//
// if graph.NVertices() != 0 {
//     t.Error("empty graph should not have any vertices")
// }
//
// graph.AddEdge("a", "b", 0)
// if graph.NVertices() != 2 {
//     t.Error("graph should have two vertices")
// }
