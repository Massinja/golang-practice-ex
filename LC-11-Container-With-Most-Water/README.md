### Problem
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
Find two lines that together with the x-axis form a container, such that the container contains the most water.  
Return the maximum amount of water a container can store. Notice that you may not slant the container.

#### Constraints:

    n == height.length
    2 <= n <= 105
    0 <= height[i] <= 104

![question_11](https://github.com/Massinja/golang-practice-ex/assets/84927906/6abc4f91-8139-470d-b0ab-09c7109cc0c3)

### Solution
Code in __container.go__ file is trying naive approach using nested for loops, which was okay on smaller arrays, however, times out on larger ones. 
That lead me to look for a different approach called Two Point Technique. It's implemented in __container-two-pointer.go__. It is certainly a winner with O(n) time and O(1) space complexities   
This approach works because it starts with the widest container (i.e., right - left is maximum) and then narrows down the width while attempting to increase the height. By moving the pointers towards each other, it guarantees that the area calculation will always consider the possibility of a higher height in the future, maximizing the area.

