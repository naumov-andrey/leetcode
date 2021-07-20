# https://leetcode.com/problems/min-stack/

class MinStack:

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.min_ = []
        self.data_ = []

    def push(self, val: int) -> None:
        if not self.min_:
            self.min_.append(val)
        else:
            self.min_.append(min(val, self.min_[-1]))
        self.data_.append(val)

    def pop(self) -> None:
        self.min_.pop()
        self.data_.pop()

    def top(self) -> int:
        return self.data_[-1]

    def getMin(self) -> int:
        return self.min_[-1]


# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(val)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()
