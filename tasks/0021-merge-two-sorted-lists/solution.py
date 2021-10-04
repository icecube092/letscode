# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    numbers = []
    index = 0
    def mergeTwoLists(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        result = []
        firsts = []
        seconds = []
        
        node1 = l1
        node2 = l2
        self.gather(firsts, node1)
        self.gather(seconds, node2)
        
        firsts.extend(seconds)
        self.numbers = sorted(firsts)
        
        main_node = self.node()
        return main_node
        
    def node(self) -> ListNode:
        if self.index == len(self.numbers):
            return None
        l = ListNode()
        l.val = self.numbers[self.index]
        self.index += 1
        l.next = self.node()
        return l
    
    def gather(self, arr, node):
        while True:
            if node == None:
                break
            arr.append(node.val)
            node = node.next
            
        
