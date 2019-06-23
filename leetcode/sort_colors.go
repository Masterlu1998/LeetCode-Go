package leetcode

// 颜色分类

// 给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

// 此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

// 注意:
// 不能使用代码库中的排序函数来解决这道题。

// 示例:

// 输入: [2,0,2,1,1,0]
// 输出: [0,0,1,1,2,2]
// 进阶：

// 一个直观的解决方案是使用计数排序的两趟扫描算法。
// 首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
// 你能想出一个仅使用常数空间的一趟扫描算法吗？


// 思路1：用各种排序方法直接完成题目，这里使用了堆排序进行实现

func sortColorsFunc1(nums []int)  {
    heapSort(nums)
}

func heapSort(sli []int) {
    length := len(sli)
    buildMaxHeap(sli)
    for i := length - 1; i >= 0; i-- {
        sli[i], sli[0] = sli[0], sli[i]
        length--
        heapify(sli, 0, length)
    }
}

func buildMaxHeap(sli []int) {
    for i := len(sli) / 2 - 1; i >= 0; i-- {
        heapify(sli, i, len(sli))
    }
}

func heapify(sli []int, node, length int) {
    largest := node
    left := node * 2 + 1
    right := node * 2 + 2
    
    if left < length && sli[left] > sli[largest] {
        largest = left
    } 
    
    if right < length && sli[right] > sli[largest] {
        largest = right
    }
    
    if largest != node {
        sli[largest], sli[node] = sli[node], sli[largest]
        heapify(sli,  largest, length)
    }
}

// 思路2：三指针法，p0始终指向值为0的边界，p2始终指向值为2的边界，p1代表当前指针。当当前指针的值为2的时候与p2
// 交换，p2--；当当前指针的值为0的时候与p0交换，p0++，p1++；其他情况p1++

func sortColorsFunc2(nums []int)  {
    p0, p1, p2 := 0, 0, len(nums)-1
    
    for p1 <= p2 {
        if nums[p1] == 0 {
            nums[p1], nums[p0] = nums[p0], nums[p1]
            p1++
            p0++
        } else if nums[p1] == 2 {
            nums[p1], nums[p2] = nums[p2], nums[p1]
            p2--
        } else {
            p1++
        }
    }
}