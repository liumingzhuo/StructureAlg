package main

import (
	"fmt"

	"github.com/liumingzhuo/StructureAlg/Array/mergesorted88"
)

func mergesortedRun() {
	//nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
	//[1,2,2,3,5,6]

	//nums1 = [1], m = 1, nums2 = [], n = 0
	//[1]
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	mergesorted88.Merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

func main() {
	mergesortedRun()
}
