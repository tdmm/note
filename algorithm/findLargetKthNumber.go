package main

import "fmt"

func main() {
	array := []int{
		3,2,3,1,2,4,5,5,6,
	}
	fmt.Println(findKthLargest(array,4))
	//fmt.Println(array)
}

func findKthLargest(nums []int, k int) int {
	//
	for i:=k/2-1;i>=0;i--{
		MinHeapify(nums,i,k-1)
		//fmt.Println("build", nums)
	}
	for i:=k;i< len(nums);i++{
		if nums[i]>nums[0]{
			nums[i],nums[0]=nums[0],nums[i]
			MinHeapify(nums,0,k-1)
			//fmt.Println("find", nums)
		}
	}
	return nums[0]
}

func MinHeapify(array []int,index,end int){
	left:=index*2+1
	right:=index*2+2
	min:=index

	if left<=end&&array[left]<array[index]{
		min=left
	}
	if right<=end&&array[right]<array[min]{
		min=right
	}

	if min!=index{
		array[min],array[index]=array[index],array[min]
		MinHeapify(array,min,end)
	}

	//fmt.Println(array)


}
