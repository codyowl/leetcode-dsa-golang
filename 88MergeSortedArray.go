func merge(nums1 []int, m int, nums2 []int, n int)  {
    writeIndex := m + n - 1
    // setting the index for nums1 and nums2
    i := m - 1 // basically lenth -1 
    j := n - 1
    // making sure to iterate across num1 and num2 elements, compare it, put the 
    // greater one at the end of the index of num1
    for j >= 0 {
        if i >=0 && nums1[i] > nums2[j] {
            nums1[writeIndex] = nums1[i]
            i --
        }else {
            nums1[writeIndex] = nums2[j]
            j --
    }
    // finally reducing the write index else the write index will keep on same state
    writeIndex--
}
}