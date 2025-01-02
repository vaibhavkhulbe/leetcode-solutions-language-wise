// O(n)
func maxScoreSightseeingPair(values []int) int {
    var max = math.MinInt64
    var iter = len(values) - 2
    var maxJ = math.MinInt64
    for iter >= 0 {
        if values[iter+1] - (iter+1) > maxJ {
            maxJ = values[iter+1] - (iter+1)
        }
        if values[iter] + iter + maxJ > max {
            max = values[iter] + iter + maxJ
        }  
        iter --      
    }
    return max
}

// O(n^2)
func maxScoreSightseeingPair(values []int) int {
    var max = math.MinInt64
    for firstIter := 0; firstIter < len(values); firstIter++ {
        for secondIter := firstIter + 1; secondIter < len(values); secondIter++ {
            if values[firstIter] + firstIter + values[secondIter] - secondIter > max         {
                max = values[firstIter] + firstIter + values[secondIter] - secondIter
            }
        }
    }
    return max
}