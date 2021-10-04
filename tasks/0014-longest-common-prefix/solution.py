class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        common = ""
        for index, char in enumerate(strs[0]):
            for word in strs:
                print(word)
                try:
                    if word[index] != char:
                        return common
                except IndexError:
                    return common
            common += char
            
        return common
