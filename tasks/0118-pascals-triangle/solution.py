class Solution:
    def generate(self, n: int) -> List[List[int]]:
        if n == 0:
            return []

        result = [[1]]

        for i in range(1, n+1):
            prev = result[len(result)-1]
            r = []
            for j in range(0, i):
                if j == 0 or j == i-1:
                    r.append(1)
                else:
                    r.append(prev[j]+prev[j-1])
            
            result.append(r)

        return result[1:]
            
       
# [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1],[1,5,10,10,5,1]]

