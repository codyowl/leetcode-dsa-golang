func removeDuplicates(nums []int) int {
    /*
    Take the first index and the value 
    if the next value is same , just increase the index and keep on iterating
    if the nex value is different update the value and increase the index
    */
    uniqueIndex := 0
    uniqueValue := nums[0]
    for i := 0; i < len(nums); i ++ {
        if nums[i] != uniqueValue{
            uniqueIndex ++
            nums[uniqueIndex] = nums[i]
            uniqueValue = nums[i]
        }
    }
    // uniqueIndex will count the index so if its 1 that mean o and 1 by count 2 
    // so we are returning uniqueIndex + 1
    return uniqueIndex + 1
}