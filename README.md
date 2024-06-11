# go-hackers-delight

[![codecov](https://codecov.io/gh/nikolaydubina/go-hackers-delight/graph/badge.svg?token=660JQtUmiO)](https://codecov.io/gh/nikolaydubina/go-hackers-delight)
![fuzzing](https://img.shields.io/badge/fuzzing-active-brightgreen)

An interactive Go showcase of ["Hacker's Delight"](https://en.wikipedia.org/wiki/Hacker%27s_Delight) 2nd Edition by Henrey S.Warren Jr, 2013.[^1][^2]

> If you write optimizing compilers or high-performance code, you must read this book.
> — [Guy L. Steele](https://en.wikipedia.org/wiki/Guy_L._Steele_Jr.) [PhD, ACM Award, AAAI Fellow, Sun Fellow, core Java design]

> This is the first book that promises to tell the deep, dark secrets of computer arithmetic, and it delivers in spades. It contains every trick I knew plus many, many more.
> A godsend for library developers, compiler writers, and lovers of elegant hacks.
> It deserves a spot on your self right next to Knuth.
> In the ten years since the first edition came out, it's been absolutely invaluable to my work at Sun and Google.
> — [Joshua Bloch](https://en.wikipedia.org/wiki/Joshua_Bloch) [PhD, Distinguished Engineer at Sun, Chief Java Architect at Google 2004~2012]

### Methodology

* extensively fuzzing
* not using any Go packages, not even standard library
* using generics whenever possible
* verifying compiled code via https://godbolt.org

### Observations

* native `Abs` performance is the same
* native `math/bits.Mul32` and `math/bits.Mul64`[^3] performance is the same
* native `Div` and `Mod` by small constants performance is better
* native `math.Sqrt` performance is better

<details><summary>Appendix: Benchmarks</summary>

```bash
$ go test -bench .
goos: darwin
goarch: arm64
pkg: github.com/nikolaydubina/go-hackers-delight
BenchmarkNoop/---------------------------------16         	1000000000	         0.0000001 ns/op
BenchmarkAbs/basic-16                                     	1000000000	         0.9349 ns/op
BenchmarkAbs/Abs-16                                       	1000000000	         0.9343 ns/op
BenchmarkAbs/Abs2-16                                      	1000000000	         0.9370 ns/op
BenchmarkAbs/Abs3-16                                      	1000000000	         0.9317 ns/op
BenchmarkAbs/Abs4-16                                      	1000000000	         0.9537 ns/op
BenchmarkAbs/AbsFastMul-16                                	1000000000	         0.9445 ns/op
BenchmarkCompress/Compress-16                             	100000000	        10.64 ns/op
BenchmarkCompress/Compress2-16                            	57345729	        21.44 ns/op
BenchmarkMul/uint32/basic-16                              	602453996	         1.998 ns/op
BenchmarkMul/uint32/MultiplyHighOrder32-16                	592384160	         2.023 ns/op
BenchmarkMul/uint64/basic-16                              	996021441	         1.202 ns/op
BenchmarkMul/uint64/MultiplyHighOrder64-16                	594833374	         2.025 ns/op
BenchmarkDivMod/DivMod/3/basic-16                         	1000000000	         0.8379 ns/op
BenchmarkDivMod/DivMod/3/DivMod3Signed-16                 	617622705	         1.933 ns/op
BenchmarkDivMod/DivMod/3/DivMod3Signed2-16                	1000000000	         1.072 ns/op
BenchmarkDivMod/DivMod/7/basic-16                         	1000000000	         0.8327 ns/op
BenchmarkDivMod/DivMod/7/DivMod7Signed-16                 	562568671	         2.126 ns/op
BenchmarkDivMod/Div/3/basic-16                            	1000000000	         0.8312 ns/op
BenchmarkDivMod/Div/3/Div3Signed-16                       	791429696	         1.518 ns/op
BenchmarkDivMod/Div/3/Div3ShiftSigned-16                  	889857019	         1.329 ns/op
BenchmarkDivMod/Div/7/basic-16                            	1000000000	         0.8317 ns/op
BenchmarkDivMod/Div/7/Div7Signed-16                       	745011523	         1.608 ns/op
BenchmarkDivMod/Div/7/Div7ShiftSigned-16                  	841457688	         1.426 ns/op
BenchmarkDivMod/Mod/3/basic-16                            	1000000000	         0.8315 ns/op
BenchmarkDivMod/Mod/3/Mod3Signed-16                       	811895162	         1.499 ns/op
BenchmarkDivMod/Mod/3/Mod3Signed2-16                      	1000000000	         0.8305 ns/op
BenchmarkDivMod/Mod/7/basic-16                            	1000000000	         0.8369 ns/op
BenchmarkDivMod/Mod/7/Mod7Signed-16                       	766542748	         1.562 ns/op
BenchmarkDivMod/Mod/7/Mod7Signed2-16                      	1000000000	         1.094 ns/op
BenchmarkDivMod/Mod/10/basic-16                           	1000000000	         0.8311 ns/op
BenchmarkDivMod/Mod/10/Mod10Signed-16                     	835493868	         1.379 ns/op
BenchmarkDivMod/DivExact/7/basic-16                       	1000000000	         0.9218 ns/op
BenchmarkDivMod/DivExact/7/DivExact7-16                   	1000000000	         0.9235 ns/op
BenchmarkDivMod/DivExact/7/Div7Signed-16                  	718700412	         1.665 ns/op
BenchmarkDivMod/DivExact/7/Div7ShiftSigned-16             	807302046	         1.483 ns/op
BenchmarkSqrt/basic-16                                    	1000000000	         1.020 ns/op
BenchmarkSqrt/SqrtNewton-16                               	170720365	         5.981 ns/op
BenchmarkSqrt/SqrtBinarySearch-16                         	77295098	        15.77 ns/op
BenchmarkSqrt/SqrtShiftAndSubtract-16                     	136669647	         8.790 ns/op
PASS
ok  	github.com/nikolaydubina/go-hackers-delight	49.760s
```
</details>

[^1]: showcase in `C` — https://github.com/hcs0/Hackers-Delight
[^2]: showcase in `Rust` — https://github.com/victoryang00/Delight
[^3]: given manual inlining of generic type, which produces equivalent Go code
