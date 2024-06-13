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
* native `math.Pow(x, 1./3)`[^4] performance is worse than `Cbrt` algorithm 💥
* native `math.Pow(x, 1./3)`[^4][^5] underflows but `Cbrt` algorithm is correct 💥
* native `math.Pow(x, n)`[^4] performance is worse than `Pow` algorithm 💥
* native `math.Log2(x)`[^4] performance is worse than `Log2` algorithm 💥
* native `math.Log10(x)`[^4] performance is worse than `Log10` algorithm 💥
* native `math.Log10(x)` [^4][^5] overflows for `math.MaxUint64` but `Log10` algorithm is correct 💥
* native `crc32.Checksum` performance is `30x`~`500x` better

<details><summary>Appendix: Benchmarks</summary>

```bash
$ go test -bench .
goos: darwin
goarch: arm64
pkg: github.com/nikolaydubina/go-hackers-delight
BenchmarkNoop/---------------------------------16         	1000000000	         0.0000000 ns/op
BenchmarkAbs/basic-16                                     	1000000000	         0.9244 ns/op
BenchmarkAbs/Abs-16                                       	1000000000	         0.9228 ns/op
BenchmarkAbs/Abs2-16                                      	1000000000	         0.9222 ns/op
BenchmarkAbs/Abs3-16                                      	1000000000	         0.9266 ns/op
BenchmarkAbs/Abs4-16                                      	1000000000	         0.9312 ns/op
BenchmarkAbs/AbsFastMul-16                                	1000000000	         0.9387 ns/op
BenchmarkCompress/Compress-16                             	100000000	        10.86 ns/op
BenchmarkCompress/Compress2-16                            	57372118	        21.39 ns/op
BenchmarkMul/uint32/basic-16                              	595723204	         2.015 ns/op
BenchmarkMul/uint32/MultiplyHighOrder32-16                	589839883	         2.041 ns/op
BenchmarkMul/uint64/basic-16                              	984232455	         1.215 ns/op
BenchmarkMul/uint64/MultiplyHighOrder64-16                	591391069	         2.030 ns/op
BenchmarkDivMod/DivMod/3/basic-16                         	1000000000	         0.8394 ns/op
BenchmarkDivMod/DivMod/3/DivMod3Signed-16                 	617556885	         1.973 ns/op
BenchmarkDivMod/DivMod/3/DivMod3Signed2-16                	1000000000	         1.080 ns/op
BenchmarkDivMod/DivMod/7/basic-16                         	1000000000	         0.8374 ns/op
BenchmarkDivMod/DivMod/7/DivMod7Signed-16                 	564408957	         2.145 ns/op
BenchmarkDivMod/Div/3/basic-16                            	1000000000	         0.8357 ns/op
BenchmarkDivMod/Div/3/Div3Signed-16                       	776687295	         1.542 ns/op
BenchmarkDivMod/Div/3/Div3ShiftSigned-16                  	888967630	         1.344 ns/op
BenchmarkDivMod/Div/7/basic-16                            	1000000000	         0.8487 ns/op
BenchmarkDivMod/Div/7/Div7Signed-16                       	738237795	         1.624 ns/op
BenchmarkDivMod/Div/7/Div7ShiftSigned-16                  	826423514	         1.451 ns/op
BenchmarkDivMod/Mod/3/basic-16                            	1000000000	         0.8469 ns/op
BenchmarkDivMod/Mod/3/Mod3Signed-16                       	798928768	         1.515 ns/op
BenchmarkDivMod/Mod/3/Mod3Signed2-16                      	1000000000	         0.8473 ns/op
BenchmarkDivMod/Mod/7/basic-16                            	1000000000	         0.8482 ns/op
BenchmarkDivMod/Mod/7/Mod7Signed-16                       	734880409	         1.632 ns/op
BenchmarkDivMod/Mod/7/Mod7Signed2-16                      	1000000000	         1.115 ns/op
BenchmarkDivMod/Mod/10/basic-16                           	1000000000	         0.8499 ns/op
BenchmarkDivMod/Mod/10/Mod10Signed-16                     	824417413	         1.463 ns/op
BenchmarkDivMod/DivExact/7/basic-16                       	1000000000	         0.9435 ns/op
BenchmarkDivMod/DivExact/7/DivExact7-16                   	1000000000	         0.9407 ns/op
BenchmarkDivMod/DivExact/7/Div7Signed-16                  	709120321	         1.686 ns/op
BenchmarkDivMod/DivExact/7/Div7ShiftSigned-16             	792262238	         1.511 ns/op
BenchmarkCbrt/basic-16                                    	45397081	        26.45 ns/op
BenchmarkCbrt/Cbrt-16                                     	68115703	        17.60 ns/op
BenchmarkPow/basic-16                                     	20584660	        56.07 ns/op
BenchmarkPow/Pow-16                                       	62423868	        19.33 ns/op
BenchmarkLog/uint32/2/basic-16                            	97335441	        12.09 ns/op
BenchmarkLog/uint32/2/Log2-16                             	981750010	         1.221 ns/op
BenchmarkLog/uint32/10/basic-16                           	141246283	         8.538 ns/op
BenchmarkLog/uint32/10/Log10-16                           	540736009	         2.220 ns/op
BenchmarkLog/uint64/2/basic-16                            	100000000	        11.81 ns/op
BenchmarkLog/uint64/2/Log2-16                             	848453780	         1.412 ns/op
BenchmarkLog/uint64/10/basic-16                           	142674736	         8.369 ns/op
BenchmarkLog/uint64/10/Log10-16                           	538090652	         2.263 ns/op
BenchmarkSqrt/basic-16                                    	1000000000	         1.047 ns/op
BenchmarkSqrt/SqrtNewton-16                               	190331125	         5.673 ns/op
BenchmarkSqrt/SqrtBinarySearch-16                         	77246794	        15.66 ns/op
BenchmarkSqrt/SqrtShiftAndSubtract-16                     	138441583	         8.693 ns/op
PASS
ok  	github.com/nikolaydubina/go-hackers-delight	66.962s
```
</details>

[^1]: showcase in `C` — https://github.com/hcs0/Hackers-Delight
[^2]: showcase in `Rust` — https://github.com/victoryang00/Delight
[^3]: given manual inlining of generic type, which produces equivalent Go code
[^4]: we are comparing native float64 result converted to uint32, as there is no better standard function
[^5]: which is due to `float64` not having enough precision
