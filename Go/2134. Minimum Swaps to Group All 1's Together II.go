// solution using modulo for circular and num of 1s (0ms)
func minSwaps(nums []int) int {
    numOfOnes := 0
    lenOfNums := len(nums)
    for _, num := range nums {
        numOfOnes += num
    }
    if numOfOnes == 0 || numOfOnes == len(nums) {
        return 0
    }
    numOnesInSubArray := 0
    minZeros := math.MaxInt
    for index := 0; index < numOfOnes; index++ {
        numOnesInSubArray+=nums[index]
    }
    for index := numOfOnes; index < lenOfNums + numOfOnes; index++ {
        numOnesInSubArray -= nums[index-numOfOnes]
        numOnesInSubArray += nums[(index)%lenOfNums]
        if numOfOnes-numOnesInSubArray < minZeros {
            minZeros = numOfOnes-numOnesInSubArray
        }
        if minZeros == 0 {
            break
        }
    }
    return minZeros
}

// solution using array repetition and num of 0s (5ms)
func minSwaps(nums []int) int {
    numOfOnes := 0
    lenOfNums := len(nums)
    for _, num := range nums {
        numOfOnes += num
    }
    if numOfOnes == 0 {
        return 0
    }
    nums = append(nums, nums...)
    numZeroInSubArray := 0
    minZeros := math.MaxInt
    for index := 0; index < numOfOnes; index++ {
        if nums[index] == 0 {
            numZeroInSubArray++
        }
    }
    for index := numOfOnes; index < lenOfNums + numOfOnes; index++ {
        if nums[index-numOfOnes] == 0 {
            numZeroInSubArray --
        }
        if nums[index] == 0 {
            numZeroInSubArray++
        }
        if numZeroInSubArray < minZeros {
            minZeros = numZeroInSubArray
        }
        if minZeros == 0 {
            break
        }
    }
    return minZeros
}