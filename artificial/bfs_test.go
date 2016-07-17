package ai

import (
	"testing"
)

//RootNode는 ID가 1이고, 부모는 0이다.
func TestNode(t *testing.T) {
	var node = RootNode()

	if node.NodeID!= 1 {
		t.Errorf("RootNode is invalid, %s", node)
	}
	if node.ParentID!=0 {
		t.Errorf("RootNode is invalid, %s", node)
	}
}

//OpenList는 초기화될 때 RootNode를 갖는다.
func TestOpenList(t *testing.T) {
	var openList = initOpenList()

	var rootNode = RootNode()
	node, _ := openList.Pop()

	if rootNode != node {
		t.Errorf("OpenList have to be initialized with RootNode %s", openList[0])
	}
}

// 트리 데이터를 적재한다.
func TestLoad(t *testing.T) {
	// treeData := make(map[int]int)
	treeData := map[int]int {
	2:1,
	3:1,
	5:1,
}
	var openList = initOpenList()
	openList=loadTree(openList, treeData)

	var rootNode = RootNode()

	node, _ :=openList.Pop()

	if rootNode != node {
		t.Errorf("OpenList have to be initialized with RootNode %s", openList[0])
	}
}

// //openList와 closedList를 리스트로 정의한다.
// func TestOpenAndClosedList(t * testing.T) {
// 	var openList OpenList
// 	var closedList ClosedList
//
// 	var node=Node{1,0,nil}
//
// 	openList =append(openList,node)
// 	closedList =append(closedList,node)
//
// 	var j =[]int{1,2,3,4}
//
// 	for _, i := range j {
// 		openList =append(openList,Node{i,1,nil})
// 		closedList =append(closedList,Node{i,1,nil})
// 	}
//
// 	//Node 구조체는 문자열로 동치성이 구현되어 있다.
// 	if openList[0]!=node {
// 		t.Errorf(" %s is expected, but %s", node, openList[0]  )
// 	}
// 	// t.Log(openList)
// }

//
//openList에 있는 모든 값을 탐색한다.
// func TestSearchTree(t *testing.T) {
// 	var openList OpenList
//
// 	var nodeCount  [10]int
// 	for  i := range nodeCount {
// 		openList = append(openList, Node{i,0,i/3})
// 	}
//
// 	for len(openList)>0 {
//
// 		node, openList := openList.Pop()
// 		//
// 		// node, openList := openList[0], openList[1:]
// 		if node.Value > 1 {
// 			t.Logf("%s is founded", node)
// 		}
// 	}
// 	// for j :=0; j< len(openList); j++ {
// 	// 	node :=openList[j]
// 	// 	if node.Value > 1{
// 	//
// 	// 	}
// 	// }
// 	// t.Log(openList)
//
// }


	// openList := OpenList{}

//
// 	var node1 = Node{1,0}
// 	var node2 = Node{2,1}
// 	openList.push(node1)
// 	var result= openList.pop()
// strNode, _ := json.Marshal(result)
// 	if (node1.NodeID!=result.NodeID) {
// 		t.Errorf("Must implement container/Heap Interface %s", strNode)
// 	}


// // node := new(Node)
//
// var id = reflect.ValueOf(node).FieldByName("NodeID")
//
// if id.CanSet() {
// 	t.Error("NodeId is not declared", id)
// }
//
//
// node.NodeID = 1
// node.ParentID = 0
//
// nodes := make([]Node, 1, 100)
//
// nodes[0] = Node{1, 0}
//
// doc, _ := json.MarshalIndent(nodes,"\n"," ")
//
//
// t.Errorf("Node is not valid, %s", string(doc))
//


// openList는 Queue로 Wrapping 한다.
//최적 또는 목표를 찾기 위해 비교 연산을 구현한다.
//모든 탐색이 종료되면 목표노드를 출력하고 프로그램은 종료된다.
