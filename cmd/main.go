package main

import (
	"algosort/sorting"
	"fmt"
)

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Original array:", arr)

	// Bubble Sort
	bubbleSorted := make([]int, len(arr))
	copy(bubbleSorted, arr)
	sorting.BubbleSort(bubbleSorted)
	fmt.Println("Bubble sorted:", bubbleSorted)

	// Selection Sort
	selectionSorted := make([]int, len(arr))
	copy(selectionSorted, arr)
	sorting.SelectionSort(selectionSorted)
	fmt.Println("Selection sorted:", selectionSorted)

	// Insertion Sort
	insertionSorted := make([]int, len(arr))
	copy(insertionSorted, arr)
	sorting.InsertionSort(insertionSorted)
	fmt.Println("Insertion sorted:", insertionSorted)

	// Merge Sort
	mergeSorted := sorting.MergeSort(arr)
	fmt.Println("Merge sorted:", mergeSorted)

	// Quick Sort
	quickSorted := sorting.QuickSort(arr)
	fmt.Println("Quick sorted:", quickSorted)
}
