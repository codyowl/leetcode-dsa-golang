func removeDuplicates(nums []int) int {
    /* iterate the array
    take the first element 
    check that with next element
    elements can appear two times
    so current_index + 2 = different should be replaced with next element 
    */
    indexSpace := 0
    for i := 0;  i<len(nums); i ++{
        if indexSpace < 2 || nums[i] != nums[indexSpace-2]{
            nums[indexSpace] = nums[i]
            indexSpace ++
        }
    }
    return indexSpace
}