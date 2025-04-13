package gtfsdata

import (
	"bytes"
	"log/slog"
	"slices"

	"github.com/seannyphoenix/bogie/pkg/collections"
)

type keyedData interface {
	partitionKey() []byte
	sortKey() []byte
}

func sortedKeyedData[T keyedData, K comparable](dataMap map[K]T) []T {
	slog.Info("Sorting keyed data")
	dataSlice := collections.MapValues(dataMap)
	slices.SortStableFunc(dataSlice, sortFunc)
	return dataSlice
}

func sortFunc[T keyedData](a, b T) int {
	partition := bytes.Compare(a.partitionKey(), b.partitionKey())
	if partition != 0 {
		return partition
	}

	return bytes.Compare(a.sortKey(), b.sortKey())
}
