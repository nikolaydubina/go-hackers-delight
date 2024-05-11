# go-hackers-delight

[![codecov](https://codecov.io/gh/nikolaydubina/go-hackers-delight/graph/badge.svg?token=660JQtUmiO)](https://codecov.io/gh/nikolaydubina/go-hackers-delight)

## Ch2

> **Right-to-Left Computability.** A function mapping words to words can be implemented with word-parallel `add`, `substract`, `and`, `or` and `not` instructions if and only if
> each bit of of the result depends only on the bits at and to the right of each input operand.

— all these operands themselves depend only on right bits, so any of their composition also depends on the right bits.
