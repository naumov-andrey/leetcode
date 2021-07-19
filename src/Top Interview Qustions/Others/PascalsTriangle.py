# https://leetcode.com/problems/pascals-triangle/

class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        result = [[1]]
        for _ in range(numRows - 1):
            result.append(self.generateRow(result[-1]))
        return result

    def generateRow(self, previousRow: List[int]) -> List[int]:
        i = 0
        row = [1]
        while i < len(previousRow) // 2:
            row.append(previousRow[i] + previousRow[i + 1])
            i += 1
        while i < len(previousRow):
            row.append(row[len(previousRow) - i - 1])
            i += 1
        return row
