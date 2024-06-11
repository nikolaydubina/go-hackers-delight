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

<details><summary>Methodology</summary>

* implementing all code close to original, leaving same comments as originals
* extensively fuzzing
* not using any Go packages, not even standard library
* using generics whenever possible
* verifying compiled code via https://godbolt.org

</details>

<details><summary>Appendix: Benchmarks</summary>

```bash
$ go test -bench .        
goos: darwin
goarch: arm64
pkg: github.com/nikolaydubina/go-hackers-delight
BenchmarkNoop/---------------------------------16         	1000000000	         0.0000001 ns/op
BenchmarkAbs/basic-16                                     	1000000000	         0.9240 ns/op
BenchmarkAbs/Abs-16                                       	1000000000	         0.9593 ns/op
BenchmarkAbs/Abs2-16                                      	1000000000	         0.9418 ns/op
BenchmarkAbs/Abs3-16                                      	1000000000	         0.9391 ns/op
BenchmarkAbs/Abs4-16                                      	1000000000	         0.9485 ns/op
BenchmarkAbs/AbsFastMul-16                                	1000000000	         0.9408 ns/op
BenchmarkCompress/Compress-16                             	100000000	        10.86 ns/op
BenchmarkCompress/Compress2-16                            	56361072	        21.37 ns/op
BenchmarkMul/uint32/basic-16                              	897062446	         1.982 ns/op
BenchmarkMul/uint32/MultiplyHighOrder32-16                	596481382	         2.012 ns/op
BenchmarkMul/uint64/basic-16                              	999763938	         1.205 ns/op
BenchmarkMul/uint64/MultiplyHighOrder64-16                	960107532	         2.027 ns/op
BenchmarkDivMod/DivMod/3/basic-16                         	1000000000	         0.8305 ns/op
BenchmarkDivMod/DivMod/3/DivMod3Signed-16                 	612924534	         1.948 ns/op
BenchmarkDivMod/DivMod/3/DivMod3Signed2-16                	1000000000	         1.068 ns/op
BenchmarkDivMod/DivMod/7/basic-16                         	1000000000	         0.8445 ns/op
BenchmarkDivMod/DivMod/7/DivMod7Signed-16                 	571082296	         2.132 ns/op
BenchmarkDivMod/Div/3/basic-16                            	1000000000	         0.8319 ns/op
BenchmarkDivMod/Div/3/Div3Signed-16                       	780752707	         1.608 ns/op
BenchmarkDivMod/Div/3/Div3ShiftSigned-16                  	887261965	         1.345 ns/op
BenchmarkDivMod/Div/7/basic-16                            	1000000000	         0.8461 ns/op
BenchmarkDivMod/Div/7/Div7Signed-16                       	740364465	         1.621 ns/op
BenchmarkDivMod/Div/7/Div7ShiftSigned-16                  	823159357	         1.432 ns/op
BenchmarkDivMod/Mod/3/basic-16                            	1000000000	         0.8322 ns/op
BenchmarkDivMod/Mod/3/Mod3Signed-16                       	809761676	         1.490 ns/op
BenchmarkDivMod/Mod/3/Mod3Signed2-16                      	1000000000	         0.8319 ns/op
BenchmarkDivMod/Mod/7/basic-16                            	1000000000	         0.8503 ns/op
BenchmarkDivMod/Mod/7/Mod7Signed-16                       	755515261	         1.566 ns/op
BenchmarkDivMod/Mod/7/Mod7Signed2-16                      	1000000000	         1.109 ns/op
BenchmarkDivMod/Mod/10/basic-16                           	1000000000	         0.8413 ns/op
BenchmarkDivMod/Mod/10/Mod10Signed-16                     	831959974	         1.455 ns/op
BenchmarkDivMod/DivExact/7/basic-16                       	1000000000	         0.9366 ns/op
BenchmarkDivMod/DivExact/7/DivExact7-16                   	1000000000	         0.9434 ns/op
BenchmarkDivMod/DivExact/7/Div7Signed-16                  	685721142	         1.706 ns/op
BenchmarkDivMod/DivExact/7/Div7ShiftSigned-16             	801524006	         1.488 ns/op
PASS
ok  	github.com/nikolaydubina/go-hackers-delight	45.723s
```
</details>

[^1]: showcase in `C` — https://github.com/hcs0/Hackers-Delight
[^2]: showcase in `Rust` — https://github.com/victoryang00/Delight
