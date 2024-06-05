/*
# Chapter 10: Integer Division by Constants

On many computers, division is very time consuming and is to be avoided.
A value of 20 or more elementary Add instructions is usually needed and execution time is usually large.
When the divisor is a constant, the division can be replaced by faster operations.

The basic trick for signed division of non-powers of two, is to multiplied compute magic number 2**32/d and then extract leftmost 32 bits of the product.
*/
package hd
