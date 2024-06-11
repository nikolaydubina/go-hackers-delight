/*
# Chapter 12: Unusual Bases for Number Systems

Standard binary representation of an integer number is in base 2.
(a_n, a_n-1, ... a3, a2, a1, a0) = a_n * 2^n + a_n-1 * 2^n-1 + ... + a3 * 2^3 + a2 * 2^2 + a1 * 2^1 + a0 *2^0

However, there can be alternative definitions, even including complex numbers as bases.
* base -2: (a_n, a_n-1, ... a3, a2, a1, a0) = a_n * (-2)^n + a_n-1 * (-2)^n-1 + ... + a3 * (-2)^3 + a2 * (-2)^2 + a1 * (-2)^1 + a0 *(-2)^0
* base (-1 + i): (a_n, a_n-1, ... a3, a2, a1, a0) = a_n * (-1 + i)^n + a_n-1 * (-1 + i)^n-1 + ... + a3 * (-1 + i)^3 + a2 * (-1 + i)^2 + a1 * (-1 + i)^1 + a0 *(-1 + i)^0
* base (-1 - i): (a_n, a_n-1, ... a3, a2, a1, a0) = a_n * (-1 - i)^n + a_n-1 * (-1 - i)^n-1 + ... + a3 * (-1 - i)^3 + a2 * (-1 - i)^2 + a1 * (-1 - i)^1 + a0 *(-1 - i)^0

Which base is most efficient? It can be argued that
under standard assumptions of hardware design (e.g. "cost of b-state circuit" is proportional to b)
the most efficient base is e.
However, it maybe be not practical.
Meanwhile, base 2 is 5% more costly than base 3 and 6% more costly than base e.
Which is interesting result, given base 2 is convenient to use and theoretically optimal.
*/
package hd
