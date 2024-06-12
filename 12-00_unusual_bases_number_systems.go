/*
# Chapter 12: Unusual Bases for Number Systems

Standard binary representation of an non-negative integer number is base 2.

  - (aₙ, ... a₃ a₂, a₁, a₀) = aₙ ⋅ 2ⁿ + ... + a₃ ⋅ 2³ + a₂ ⋅ 2² + a₁ ⋅ 2¹ + a₀ ⋅ 2⁰

However, there are alternative bases, including complex numbers.

  - (aₙ, ... a₃ a₂, a₁, a₀) = aₙ ⋅ -2ⁿ + ... + a₃ ⋅ -2³ + a₂ ⋅ -2² + a₁ ⋅ -2¹ + a₀ ⋅ -2⁰
  - (aₙ, ... a₃ a₂, a₁, a₀) = aₙ ⋅ (-1 + 𝑖)ⁿ + ... + a₃ ⋅ (-1 + 𝑖)³ + a₂ ⋅ (-1 + 𝑖)² + a₁ ⋅ (-1 + 𝑖)¹ + a₀ ⋅ (-1 + 𝑖)⁰
  - (aₙ, ... a₃ a₂, a₁, a₀) = aₙ ⋅ (-1 - 𝑖)ⁿ + ... + a₃ ⋅ (-1 - 𝑖)³ + a₂ ⋅ (-1 - 𝑖)² + a₁ ⋅ (-1 - 𝑖)¹ + a₀ ⋅ (-1 - 𝑖)⁰

Which base is the most efficient?
It can be argued that under standard assumptions of hardware design (e.g. "cost of b-state circuit" is proportional to b) the most efficient base is 𝑒.
However, it is not practical.
Meanwhile, base 2 is only 5% more costly than base 3 and only 6% more costly than base 𝑒.
Which is an interesting result, given base 2 is both convenient to use and is theoretically optimal.
*/
package hd
