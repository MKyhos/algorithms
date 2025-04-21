package rtree

import (
	"math"
	"sort"
)

type Entry struct {
	MBR      Rectangle
	ChildPtr *Node
	DataID   int
}

type Node struct {
	Entries []*Entry
	IsLeaf  bool
	Parent  *Node
	Level   int
}
type RPlusTree struct {
	Root       *Node
	MaxEntries int
	MinEntries int
	Dimensions int
	Height     int
	Size       int
}

func NewRPlusTree(maxEntries, dimensions int) *RPlusTree {
	minEntries := max(maxEntries/2, 2)

	root := &Node{
		Entries: make([]*Entry, 0),
		IsLeaf:  true,
		Level:   0,
	}
	return &RPlusTree{
		Root:       root,
		MaxEntries: maxEntries,
		MinEntries: minEntries,
		Dimensions: dimensions,
		Height:     1,
		Size:       0,
	}
}

func (tree *RPlusTree) Insert(point Point, id int) {
	rect := NewRectangle(point, 0)
	leaf := tree.findLeaf(tree.Root, &rect)

	entry := &Entry{
		MBR:      rect,
		ChildPtr: nil,
		DataID:   id,
	}

	leaf.Entries = append(leaf.Entries, entry)
	tree.Size++

	if len(leaf.Entries) > tree.MaxEntries {
		tree.splitNode(leaf)
	}
}

func (tree *RPlusTree) Search(query Point, eps float64) []int {
	searchRect := NewRectangle(query, eps)
	results := make([]int, 0)
	tree.searchNode(tree.Root, &searchRect, &results)
	return results
}

func (tree *RPlusTree) splitNode(node *Node) {
	axis, index := tree.chooseSplitAxis(node)
	sort.Slice(node.Entries, func(i, j int) bool {
		return node.Entries[i].MBR.GetCentroid(axis) < node.Entries[j].MBR.GetCentroid(axis)
	})

	newNode := &Node{
		Entries: make([]*Entry, 0, tree.MaxEntries),
		IsLeaf:  node.IsLeaf,
		Level:   node.Level,
		Parent:  node.Parent,
	}
	newNode.Entries = append(newNode.Entries, node.Entries[index:]...)
	node.Entries = node.Entries[:index]

	if !newNode.IsLeaf {
		for _, entry := range newNode.Entries {
			entry.ChildPtr.Parent = newNode
		}
	}

	if node.Parent == nil {
		newRoot := &Node{
			Entries: make([]*Entry, 0, tree.MaxEntries),
			IsLeaf:  false,
			Level:   node.Level + 1,
		}

		nodeMBR := tree.calculateMBR(node)
		newNodeMBR := tree.calculateMBR(newNode)
		newRoot.Entries = append(
			newRoot.Entries,
			&Entry{MBR: nodeMBR, ChildPtr: node, DataID: -1},
			&Entry{MBR: newNodeMBR, ChildPtr: newNode, DataID: -1},
		)
		node.Parent = newRoot
		newNode.Parent = newRoot

		tree.Root = newRoot
		tree.Height++
	} else {
		parentNode := node.Parent
		newNodeMBR := tree.calculateMBR(newNode)
		parentNode.Entries = append(
			parentNode.Entries,
			&Entry{MBR: newNodeMBR, ChildPtr: newNode, DataID: -1},
		)

		for i, entry := range parentNode.Entries {
			if entry.ChildPtr == node {
				parentNode.Entries[i].MBR = tree.calculateMBR(node)
				break
			}
		}
		if len(parentNode.Entries) > tree.MaxEntries {
			tree.splitNode(parentNode)
		}
	}
}

func (tree *RPlusTree) chooseSplitAxis(node *Node) (int, int) {
	bestAxis := 0
	bestIndex := tree.MinEntries
	minOverlap := math.MaxFloat64
	bestArea := math.MaxFloat64

	for axis := range tree.Dimensions {
		sort.Slice(node.Entries, func(i, j int) bool {
			return node.Entries[i].MBR.GetCentroid(bestAxis) < node.Entries[j].MBR.GetCentroid(bestAxis)
		})

		for index := tree.MinEntries; index <= len(node.Entries)-tree.MinEntries; index++ {
			leftMBR := tree.calculateMBRForEntries(node.Entries[:index])
			rightMBR := tree.calculateMBRForEntries(node.Entries[index:])
			overlap := leftMBR.Overlap(&rightMBR)
			if overlap < minOverlap || (overlap == minOverlap && leftMBR.Area()+rightMBR.Area() < bestArea) {
				minOverlap = overlap
				bestAxis = axis
				bestIndex = index
				bestArea = leftMBR.Area() + rightMBR.Area()
			}
		}
	}
	sort.Slice(node.Entries, func(i, j int) bool {
		return node.Entries[i].MBR.GetCentroid(bestAxis) < node.Entries[j].MBR.GetCentroid(bestAxis)
	})
	return bestAxis, bestIndex
}

func (tree *RPlusTree) calculateMBR(node *Node) Rectangle {
	return tree.calculateMBRForEntries(node.Entries)
}

func (tree *RPlusTree) calculateMBRForEntries(entries []*Entry) Rectangle {
	mbr := Rectangle{
		Min: make(Point, tree.Dimensions),
		Max: make(Point, tree.Dimensions),
	}
	if len(entries) == 0 {
		return mbr
	}
	for i := range mbr.Min {
		mbr.Min[i] = entries[0].MBR.Min[i]
		mbr.Max[i] = entries[0].MBR.Max[i]
	}
	for _, entry := range entries[1:] {
		for i := range mbr.Min {
			mbr.Min[i] = math.Min(mbr.Min[i], entry.MBR.Min[i])
			mbr.Max[i] = math.Max(mbr.Max[i], entry.MBR.Max[i])
		}
	}
	return mbr
}

func (tree *RPlusTree) searchNode(node *Node, searchRect *Rectangle, results *[]int) {
	for _, entry := range node.Entries {
		if entry.MBR.Overlaps(searchRect) {
			if node.IsLeaf {
				*results = append(*results, entry.DataID)
			} else {
				tree.searchNode(entry.ChildPtr, searchRect, results)
			}
		}
	}
}

func (tree *RPlusTree) findLeaf(node *Node, rect *Rectangle) *Node {
	if node.IsLeaf {
		return node
	}

	minEnlargement := math.MaxFloat64
	var selectedChild *Node
	for _, entry := range node.Entries {
		originalArea := entry.MBR.Area()
		enlarged := entry.MBR
		enlarged.Enlarge(rect)
		enlargement := enlarged.Area() - originalArea

		if enlargement < minEnlargement ||
			(enlargement == minEnlargement &&
				entry.MBR.Area() < selectedChild.Entries[0].MBR.Area()) {

			minEnlargement = enlargement
			selectedChild = entry.ChildPtr
		}
	}
	return tree.findLeaf(selectedChild, rect)
}
