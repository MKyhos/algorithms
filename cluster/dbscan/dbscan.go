package dbscan

import (
	"github.com/mkyhos/algorithms/index/rtree"
)

type ClusterResult struct {
	Labels       []int
	ClusterCount int
}

func DBSCAN(data []rtree.Point, eps float64, minPts int) ClusterResult {
	n := len(data)

	labels := make([]int, n)
	for i := range labels {
		labels[i] = -2
	}

	// Setup R+ Tree
	dimensions := len(data[0])
	tree := rtree.NewRPlusTree(50, dimensions)
	for i, point := range data {
		tree.Insert(point, i)
	}

	clusterID := 0
	for pointIdx := range n {
		// skip if already processed
		if labels[pointIdx] != -2 {
			continue
		}
		// Mark as visited - default to 'noise'
		labels[pointIdx] = -1

		neighborIndices := tree.Search(data[pointIdx], eps)

		if len(neighborIndices) < minPts {
			continue // it is noise
		}
		clusterID++
		expandCluster(data, tree, labels, pointIdx, neighborIndices, clusterID, eps, minPts)
	}

	return ClusterResult{
		Labels:       labels,
		ClusterCount: clusterID,
	}
}

func expandCluster(data []rtree.Point, tree *rtree.RPlusTree, labels []int, pointIdx int, neighborIndices []int, clusterID int, eps float64, minPts int) {
	labels[pointIdx] = clusterID

	i := 0
	for i < len(neighborIndices) {
		neighborIdx := neighborIndices[i]
		i++

		if labels[neighborIdx] == -1 {
			labels[neighborIdx] = clusterID
		}

		if labels[neighborIdx] != -2 {
			continue
		}
		labels[neighborIdx] = clusterID
		newNeighbors := tree.Search(data[neighborIdx], eps)
		if len(newNeighbors) >= minPts {
			neighborIndices = append(neighborIndices, newNeighbors...)
		}
	}
}
