# go-hackers-delight

[![codecov](https://codecov.io/gh/nikolaydubina/go-hackers-delight/graph/badge.svg?token=660JQtUmiO)](https://codecov.io/gh/nikolaydubina/go-hackers-delight)
![fuzzing](https://img.shields.io/badge/fuzzing-active-brightgreen)

An interactive Go showcase of ["Hacker's Delight"](https://www.amazon.com/Hackers-Delight-2nd-Henry-Warren/dp/0321842685)[^1][^2][^3] 2nd Edition by Henrey S.Warren Jr, 2013.

> In Hacker’s Delight, Second Edition, Hank Warren once again compiles an irresistible collection of programming hacks: timesaving techniques, algorithms, and tricks that help programmers build more elegant and efficient software, while also gaining deeper insights into their craft.
> Warren’s hacks are eminently practical, but they’re also intrinsically interesting, and sometimes unexpected, much like the solution to a great puzzle. They are, in a word, a delight to any programmer who is excited by the opportunity to improve.
> — Amazon

> The author, an IBM researcher working on systems ranging from the IBM 704 to the PowerPC, collected what he called "programming tricks" over the course of his career.
> These tricks concern efficient low-level manipulation of bit strings and numbers. According to the book's foreword by [Guy L. Steele](https://en.wikipedia.org/wiki/Guy_L._Steele_Jr.) [PhD, ACM Award, AAAI Fellow, Sun Fellow, core Java design], the target audience includes compiler writers and people writing high-performance code.
> — Wikipedia

> This is the first book that promises to tell the dep, dark secrets of computer arithmetic, and it delivers in spades. It contains every trick I knew plus many, many more.
> A godsend for library developers, compiler writers, and lovers of elegant hacks.
> It deserves a spot on your self right next to Knuth.
> In the tehn years since the first edition came out, it's been absolutely invaluable to my work at Sun and Google.
> — [Joshua Bloch](https://en.wikipedia.org/wiki/Joshua_Bloch) [PhD, Distinguished Engineer at Sun Microsystems, Chief Java Architect at Google 2004~2012]

## Ch2

It is advisable to avoid branches in simple functions.
This book shows many examples on how to compute common functions with few branch-free instructions.

----

> **Right-to-Left Computability.** A function mapping words to words can be implemented with word-parallel `add`, `substract`, `and`, `or` and `not` instructions if and only if
> each bit of of the result depends only on the bits at and to the right of each input operand.

— all these operands themselves depend only on right bits, so any of their composition also depends on the right bits.

----

It can be shown that in the ordinary addition of binary numbers with each bit independently equally likely to be 0 or 1, a carry occurs at each position with probability about 0.5.

----

> If a computer's instruction set includes an instruction for each of 16 Boolean functions of two variables,
> then any Boolean function of three variables can be implemented with four (or fewer) instructions.
> (and four variables with seven instructions).

----

[^1]: https://en.wikipedia.org/wiki/Hacker%27s_Delight
[^2]: showcase in `C` — https://github.com/hcs0/Hackers-Delight
[^3]: showcase in `Rust` — https://github.com/victoryang00/Delight
