package main

import "fmt"

func main() {
	array := []int{
		55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9,
	}
	HeapSort(array)
	fmt.Println(array)
}

func HeapSort(array []int) {
	for i := len(array)/2 - 1; i >= 0; i-- {
		MaxHeapify(array, i, len(array)-1)
	}

	for i:= len(array)-1;i>=0;i--{
		array[i],array[0]=array[0],array[i]
		MaxHeapify(array,0,i-1)
	}

}

func MaxHeapify(array []int, index, end int) {
	left := index*2 + 1
	right := index*2 + 2
	largest := index

	if left <= end && array[left] > array[largest] {
		largest = left
	}
	if right <= end && array[right] > array[largest] {
		largest = right
	}

	if largest != index {
		array[largest], array[index] = array[index], array[largest]
		MaxHeapify(array, largest, end)
	}
}


//伪代码
//
//BuildHeap(A)
//n = length(A)
//for  i = n/2 downto 1  do   //从非叶子节点开始，自底往上，使A变成最大堆
//Max_Heapify(A, i, n)
//end

//Max_Heapify(A,idx,max) //idx：数组开始的下标，max：最大的数组下标
//    left = 2*idx
//    right = 2*idx
//    if(left<max and A[left]>A[idx]) then
//        largest = left
//    else
//        largest = idx
//    if(right < max and A[right]>A[largest]) then
//        largest = ritht
//
//    if(largest != idx) then
//        exchange A[largest] with A[idx]
//        Max_Heapify(A,largest,max) //交换位置后，还需要调整它的子树
//end

//HeapSort(A)
//
//      BuildHeap(A)
//
//      for i = length(A) downto 2   do
//
//             exchange  A[1] with A[i] //把最大堆根节点与最后一个互换
//
//             Max_Heapify(A,1, i-1) //把交互后的排除在堆之外，重新从1到i-1,调整堆
//
//end
