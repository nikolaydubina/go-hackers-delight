/*
# Chapter 2: Basics

For better performance, it is advisable to avoid branches in simple functions.

Right-to-Left Computability — a function mapping words to words can be implemented with word-parallel
`add`, `substract`, `and`, `or` and `not` instructions if and only if each bit of of the result depends only on the bits at and to the right of each input operand.
All these operands themselves depend only on right bits, so any of their composition also depends on the right bits.

It can be shown that in the ordinary addition of binary numbers with each bit independently equally likely to be 0 or 1,
a carry occurs at each position with probability about 0.5.

If instruction set includes an instruction for each of 16 Boolean functions of two variables,
then any Boolean function of three variables can be implemented with four or fewer instructions (and four variables with seven instructions).
*/
package hd
