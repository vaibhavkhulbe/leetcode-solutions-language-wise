// O(n)
class Solution {
public:
    int maxScoreSightseeingPair(vector<int>& values) {
        int maxAns = INT_MIN;
        int firstIter = values.size() - 2;
        int maxJ = INT_MIN;
        while(firstIter >= 0) {
            maxJ = max(maxJ, values.at(firstIter + 1) - firstIter - 1);
            maxAns = max(maxAns, values.at(firstIter) + firstIter + maxJ);
            firstIter --;
        }
        return maxAns;
    }
};

// O(n^2)
class Solution2 {
public:
    int maxScoreSightseeingPair(vector<int>& values) {
        int maxAns = INT_MIN;
        for (int firstIter = 0; firstIter < values.size(); firstIter++) {
            for (int secondIter = firstIter + 1; secondIter < values.size(); secondIter++) {
                if (values.at(firstIter) + values.at(secondIter) + firstIter - secondIter > maxAns)
                    maxAns = values.at(firstIter) + values.at(secondIter) + firstIter - secondIter;
            }
        }
        return maxAns;
    }
};