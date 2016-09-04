package ai

import (
	"encoding/json"
	"fmt"
)

/*
Node Transitable State
*/
type Node struct {
	NodeID   int `json:"nodeid",desc:"노드 번호"`
	ParentID int `json:"parentid",desc:"상위 노드 번호"`
	Value int  `json:"value",desc:"노드의 값"`
}

// RootNode No Parent, No Value.
func RootNode() Node {
	return Node{1,0,0}
}

func (n Node) String() string {

	rslt, _ := json.Marshal(n)
	return fmt.Sprint(string(rslt))
}


// OpenList Width-First Search
type OpenList []Node

// Pop FIFO Pop
func (ol OpenList) Pop() (Node, OpenList) {
	node, ol := ol[0], ol[1:]
	return node, ol
}

// Load tree data
func loadTree(openList OpenList, data map[int]int) OpenList {

	for nodeID, parentID := range data {
		openList=append(openList,Node{nodeID,parentID,0}) 
	}
	return openList
}

func initOpenList() OpenList{
	var openList OpenList
	openList = append(openList, RootNode())
	return openList
}
// ClosedList Width-First Search
type ClosedList []Node
