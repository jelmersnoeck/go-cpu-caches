package main

import "testing"

func BenchmarkRows(b *testing.B) {
	for n := 0; n < b.N; n++ {
		loopRows()
	}
}

func BenchmarkColumns(b *testing.B) {
	for n := 0; n < b.N; n++ {
		loopColumns()
	}
}

func BenchmarkLinkedList(b *testing.B) {
	for n := 0; n < b.N; n++ {
		loopLinkedList()
	}
}
