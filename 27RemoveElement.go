func removeElement(nums []int, val int) int {
    /*
     Iterate through the nums array 
     The iteration should happen till the len of nums array
     if non val comes then update the writeIndex, place it in writeIndex
    */
    // for holding the index need to update it whenever we come across non val
    nonValIndex := 0
    for i := 0; i <len(nums); i++ {
        if nums[i] != val{
            nums[nonValIndex] = nums[i]
            nonValIndex ++
        }
    }
    return nonValIndex
}