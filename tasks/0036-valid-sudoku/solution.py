class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        values = {"1": 0, "2": 0, "3": 0, "4": 0, "5": 0, "6": 0, "7": 0,
                  "8": 0, "9": 0, ".": 0}
        rows = []
        columns = []
        boxes = []

        for i in range(0, 9):
            rows.append(values.copy())
            columns.append(values.copy())
            boxes.append(values.copy())

        for ri in range(0, len(board)):
            row = board[ri]
            for i in range(0, len(board[ri])):
                rows[ri][row[i]] += 1
                columns[i][row[i]] += 1

                boxes[ri // 3 * 3 + i // 3][row[i]] += 1
        values.pop(".")
        for r, c, b in zip(rows, columns, boxes):
            for k in values.keys():
                if r[k] > 1 or c[k] > 1 or b[k] > 1:
                    return False
        return True
