/*
# Chapter 10: Integer Division by Constants

On many computers, division is very time consuming and is to be avoided.
A value of 20 or more elementary Add instructions is usually needed and execution time is usually large.
When the divisor is a constant, the division can be replaced by faster operations.

The basic trick for signed division of non-powers of two, is to multiplied compute magic number 2**32/d and then extract leftmost 32 bits of the product.

Exact division by constants can be computed quickly by using multiplicative inverse.
It relies on theorem "if a and m are relatively prime integers, then there exists an integer ā such that a*ā = 1 (mod m)".
*/
package hd
