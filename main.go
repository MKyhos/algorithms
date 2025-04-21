package main

import (
	"fmt"

	"github.com/mkyhos/algorithms/cluster/dbscan"
	"github.com/mkyhos/algorithms/index/rtree"
)

func main() {
	data := []rtree.Point{
		{1.0, 2.0},
		{1.5, 1.8},
		{1.3, 2.2},
		{10.0, 10.0},
		{10.3, 9.9},
		{10.5, 10.2},
		{99.0, 99.0}, // Outlier
	}

	result := dbscan.DBSCAN(data, 1.0, 2)

	fmt.Printf("Found %d clusters\n", result.ClusterCount)
	fmt.Println("Point clusters:")

	for i, label := range result.Labels {
		if label == -1 {
			fmt.Printf("Point %d: Noise\n", i)
		} else {
			fmt.Printf("Point %d: Cluster %d\n", i, label)
		}
	}
}
