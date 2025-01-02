func rangeSum(nums []int, n int, left int, right int) int {
    newNums := make([]int , 0, (n*(n+1))/2)
    for index := 0; index < n; index++ {
        sum := 0
        for window := 0; window < n-index; window++ {
            sum += nums[index+window]
            newNums = append(newNums, sum)
        }
    }
    sort.Ints(newNums)
    ansSum := 0
    for index := left-1; index < right; index++ {
        ansSum = (ansSum + newNums[index]) % 1000000007
    }
    return ansSum
}