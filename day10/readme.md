## Part2

Combinations calculations is https://en.wikipedia.org/wiki/Combination#Number_of_k-combinations_for_all_k

> Given 3 cards numbered 1 to 3, there are 8 distinct combinations (subsets), including the empty set:  2^3 =8

In our case when length is bigger than the adapter rating difference we cannot have empty set (or adapters would not plug) so it is (2^3) - 1 = 7.

### First Case:
0, 1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19,

4,**5,6,**7 -  2^2 = 4

10,**11**,12 - 2 ^ 1 = 2

2*4 = 8

### Second case:

0, 1, 2, 3, 4, 7, 8, 9, 10, 11, 14, 17, 18, 19, 20, 23, 24, 25, 28, 31, 32, 33, 34, 35, 38, 39, 42, 45, 46, 47, 48, 49, 

0,**1,2,3**,4.  		2^3 -1  = 7 (can’t have none match in middle)

~1,**2,3**,4. 			2^2     = 4 (forgot the 0 case initially!)~

7,**8,9,10**,11 		2^3 -1  = 7 (can’t have none match in middle)

17,**18,19**,20.  		2^2     = 4

23,**24**,25,		    2^1     = 2

31,**32,33,34**,35      2^3 -1  = 7 (can’t have none match in middle)

45,**46,47,48**,49. 	2^3 -1  = 7 (can’t have none match in middle)

7 * 7 * 4 * 2 * 7 * 7 = 19208

Good Explanation:
https://github.com/tudorpavel/advent-of-code-2020/tree/master/day10#part-2

Cool way to visualize: 
https://www.reddit.com/r/adventofcode/comments/ka8z8x/2020_day_10_solutions/gf9osmx/?utm_source=reddit&utm_medium=web2x&context=3
https://i.imgur.com/WfnLpA7.png