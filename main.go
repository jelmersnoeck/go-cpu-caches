package main

import (
	"fmt"
	"time"
)

var (
	matrixData [][]int32
	linkedList *ListItem

	matrixSize = 100 // 40KB for whole matrix
	//matrixSize = 10000 // 40KB/row -> 1 row fits in L1 core cache (64KB)
	//matrixSize = 30000 // 200KB/row -> 1 row fits in L2 core cache (256KB)
	//matrixSize = 50000 // 200KB/row -> 1 row fits in L2 core cache (3MB)
)

type ListItem struct {
	n    int32
	next *ListItem
}

func init() {
	matrixData = make([][]int32, matrixSize)
	linkedList = &ListItem{}

	for c := 0; c < matrixSize; c++ {
		rowData := make([]int32, matrixSize)
		for r := 0; r < matrixSize; r++ {
			if c != 0 && r != 0 {
				linkedList = &ListItem{
					n:    int32(r),
					next: linkedList,
				}
			}

			rowData[r] = int32(r)
		}
		matrixData[c] = rowData
	}
}

func loopRows() {
	sum := int32(0)
	for r := 0; r < matrixSize; r++ {
		for c := 0; c < matrixSize; c++ {
			sum += matrixData[r][c]
		}
	}
}

func loopColumns() {
	sum := int32(0)
	for r := 0; r < matrixSize; r++ {
		for c := 0; c < matrixSize; c++ {
			sum += matrixData[c][r]
		}
	}
}

func loopLinkedList() {
	sum := int32(0)

	listItem := linkedList
	for listItem.next != nil {
		sum += listItem.n
		listItem = listItem.next
	}
}

func timer(name string, f func()) {
	start := time.Now()
	f()
	fmt.Printf("%s took %s\n", name, time.Now().Sub(start))
}
