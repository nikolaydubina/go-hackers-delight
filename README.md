# go-hackers-delight

[![codecov](https://codecov.io/gh/nikolaydubina/go-hackers-delight/graph/badge.svg?token=660JQtUmiO)](https://codecov.io/gh/nikolaydubina/go-hackers-delight)

This is an interactive Go showcase for "[Hacker's Delight](https://www.amazon.com/Hackers-Delight-2nd-Henry-Warren/dp/0321842685)" 2nd Edition by Henrey S.Warren Jr, 2013.

## Ch2

> **Right-to-Left Computability.** A function mapping words to words can be implemented with word-parallel `add`, `substract`, `and`, `or` and `not` instructions if and only if
> each bit of of the result depends only on the bits at and to the right of each input operand.

— all these operands themselves depend only on right bits, so any of their composition also depends on the right bits.

----

It can be shown that in the ordinary addition of binary numbers with each bit independently equally likely to be 0 or 1, a carry occurs at each position with probability about 0.5.

----

## References
  
- [wiki: Hackers Delight](https://en.wikipedia.org/wiki/Hacker%27s_Delight)
- [showcase in `C`](https://github.com/hcs0/Hackers-Delight)
