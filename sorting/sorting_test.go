package sorting

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	BubbleSort(arr)
	expected := []int{11, 12, 22, 25, 34, 64, 90}
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("BubbleSort failed, expected %v, got %v", expected, arr)
	}
}

func TestSelectionSort(t *testing.T) {
	arr := []int{64, 25, 12, 22, 11}
	SelectionSort(arr)
	expected := []int{11, 12, 22, 25, 64}
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("SelectionSort failed, expected %v, got %v", expected, arr)
	}
}

func TestInsertionSort(t *testing.T) {
	arr := []int{12, 11, 13, 5, 6}
	InsertionSort(arr)
	expected := []int{5, 6, 11, 12, 13}
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("InsertionSort failed, expected %v, got %v", expected, arr)
	}
}

func TestMergeSort(t *testing.T) {
	arr := []int{12, 11, 13, 5, 6, 7}
	sortedArr := MergeSort(arr)
	expected := []int{5, 6, 7, 11, 12, 13}
	if !reflect.DeepEqual(sortedArr, expected) {
		t.Errorf("MergeSort failed, expected %v, got %v", expected, sortedArr)
	}
}

func TestQuickSort(t *testing.T) {
	arr := []int{10, 80, 30, 90, 40, 50, 70}
	sortedArr := QuickSort(arr)
	expected := []int{10, 30, 40, 50, 70, 80, 90}
	if !reflect.DeepEqual(sortedArr, expected) {
		t.Errorf("QuickSort failed, expected %v, got %v", expected, sortedArr)
	}
}
