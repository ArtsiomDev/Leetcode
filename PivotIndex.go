func pivotIndex(nums []int) int {
    sum:=0
    for i, _ := range nums {
        sum+=nums[i]
    }
    sumLeft :=0
    for i, _ := range nums {
        sumLeft+=nums[i]
        if sumLeft - nums[i] == sum - sumLeft {
            return i
        }
    }
    return -1
}
