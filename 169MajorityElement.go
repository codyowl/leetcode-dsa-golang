func majorityElement(nums []int) int {
    /* The optimal approach is to use booyer moore algorithm, this
        solution is moreove like a trick, not much interesting*/
    count := 0
    value := 0
    for _, num := range nums{
        if count == 0 {
            value = num
            count += 1
        }else if value == num {
            count ++
        }else{
            count -- 
        }
    }   
    return value 
}