# go-hackers-delight

[![codecov](https://codecov.io/gh/nikolaydubina/go-hackers-delight/graph/badge.svg?token=660JQtUmiO)](https://codecov.io/gh/nikolaydubina/go-hackers-delight)
![fuzzing](https://img.shields.io/badge/fuzzing-active-brightgreen)
[![Hits](https://hits.sh/github.com/nikolaydubina/go-hackers-delight.svg?view=today-total)](https://hits.sh/github.com/nikolaydubina/go-hackers-delight/)
[![Go Reference](https://pkg.go.dev/badge/github.com/nikolaydubina/go-hackers-delight.svg)](https://pkg.go.dev/github.com/nikolaydubina/go-hackers-delight)


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
* native `Div` and `Mod` by small constants performance is better
* native `math/bits.Mul32` and `math/bits.Mul64`[^3] performance is the same
* native `math/bits.LeadingZeros` performance is better than `LeadingZeroes` algorithm
* native `math.Sqrt` performance is better
* native `math.Pow(x, 1./3)`[^4] performance is worse than `Cbrt` algorithm 💥
* native `math.Pow(x, 1./3)`[^4][^5] underflows but `Cbrt` algorithm is correct 💥
* native `math.Pow(x, n)`[^4] performance is worse than `Pow` algorithm 💥
* native `math.Log2(x)`[^4] performance is worse than `Log2` algorithm 💥
* native `math.Log10(x)`[^4] performance is worse than `Log10` algorithm 💥
* native `math.Log10(x)` [^4][^5] overflows for `math.MaxUint64` but `Log10` algorithm is correct 💥
* native `1/math.Sqrt(x)` performance is `1.5x` better than `RSqrtFloat32` algorithm
* native `switch` performance is better than `CycleThreeValues` algorithm
* native `crc32.Checksum` performance is `30x`~`500x` better than `CRC32` algorithms
* simplistic `LRUCache` performance is `3x` worse than `LRUCache` algorithm 💥 

<details><summary>Appendix: Benchmarks</summary>

```bash
$ go test -bench .
goos: darwin
goarch: arm64
pkg: github.com/nikolaydubina/go-hackers-delight
BenchmarkNoop/---------------------------------16         	1000000000	         0.0000001 ns/op
BenchmarkAbs/basic-16                                     	1000000000	         0.9826 ns/op
BenchmarkAbs/Abs-16                                       	1000000000	         0.9647 ns/op
BenchmarkAbs/Abs2-16                                      	1000000000	         0.9943 ns/op
BenchmarkAbs/Abs3-16                                      	1000000000	         0.9819 ns/op
BenchmarkAbs/Abs4-16                                      	1000000000	         1.003 ns/op
BenchmarkAbs/AbsFastMul-16                                	1000000000	         0.9598 ns/op
BenchmarkAvg/basic-16                                     	973716225	         2.045 ns/op
BenchmarkAvg/AvgFloor-16                                  	602586224	         2.050 ns/op
BenchmarkAvg/AvgCeil-16                                   	582029594	         2.054 ns/op
BenchmarkCycleThree/basic-16                              	767160418	         1.560 ns/op
BenchmarkCycleThree/CycleThreeValues-16                   	438818894	         2.729 ns/op
BenchmarkLeadingZeros/uint32/basic-16                     	1000000000	         0.9419 ns/op
BenchmarkLeadingZeros/uint32/LeadingZerosUint32-16        	1000000000	         1.124 ns/op
BenchmarkLeadingZeros/uint64/basic-16                     	1000000000	         0.9230 ns/op
BenchmarkLeadingZeros/uint64/LeadingZerosUint64-16        	898095195	         1.336 ns/op
BenchmarkCompress/Compress-16                             	100000000	        10.60 ns/op
BenchmarkCompress/Compress2-16                            	55584826	        21.52 ns/op
BenchmarkLRU/basic-16                                     	246358870	         4.870 ns/op
BenchmarkLRU/LRUCache-16                                  	960896830	         1.239 ns/op
BenchmarkMul/uint32/basic-16                              	593555838	         1.892 ns/op
BenchmarkMul/uint32/MultiplyHighOrder32-16                	951445552	         2.046 ns/op
BenchmarkMul/uint64/basic-16                              	977065424	         1.220 ns/op
BenchmarkMul/uint64/MultiplyHighOrder64-16                	675693746	         2.042 ns/op
BenchmarkDivMod/DivMod/3/basic-16                         	1000000000	         0.8500 ns/op
BenchmarkDivMod/DivMod/3/DivMod3Signed-16                 	605588445	         1.970 ns/op
BenchmarkDivMod/DivMod/3/DivMod3Signed2-16                	1000000000	         1.078 ns/op
BenchmarkDivMod/DivMod/7/basic-16                         	1000000000	         0.8311 ns/op
BenchmarkDivMod/DivMod/7/DivMod7Signed-16                 	582087586	         2.105 ns/op
BenchmarkDivMod/Div/3/basic-16                            	1000000000	         0.8325 ns/op
BenchmarkDivMod/Div/3/Div3Signed-16                       	793883130	         1.509 ns/op
BenchmarkDivMod/Div/3/Div3ShiftSigned-16                  	907116610	         1.320 ns/op
BenchmarkDivMod/Div/7/basic-16                            	1000000000	         0.8344 ns/op
BenchmarkDivMod/Div/7/Div7Signed-16                       	755509315	         1.590 ns/op
BenchmarkDivMod/Div/7/Div7ShiftSigned-16                  	841563656	         1.424 ns/op
BenchmarkDivMod/Mod/3/basic-16                            	1000000000	         0.8309 ns/op
BenchmarkDivMod/Mod/3/Mod3Signed-16                       	812136249	         1.466 ns/op
BenchmarkDivMod/Mod/3/Mod3Signed2-16                      	1000000000	         0.8410 ns/op
BenchmarkDivMod/Mod/7/basic-16                            	1000000000	         0.8332 ns/op
BenchmarkDivMod/Mod/7/Mod7Signed-16                       	766677633	         1.564 ns/op
BenchmarkDivMod/Mod/7/Mod7Signed2-16                      	1000000000	         1.095 ns/op
BenchmarkDivMod/Mod/10/basic-16                           	1000000000	         0.8318 ns/op
BenchmarkDivMod/Mod/10/Mod10Signed-16                     	868932930	         1.441 ns/op
BenchmarkDivMod/DivExact/7/basic-16                       	1000000000	         0.9247 ns/op
BenchmarkDivMod/DivExact/7/DivExact7-16                   	1000000000	         0.9238 ns/op
BenchmarkDivMod/DivExact/7/Div7Signed-16                  	718667949	         1.668 ns/op
BenchmarkDivMod/DivExact/7/Div7ShiftSigned-16             	802988229	         1.490 ns/op
BenchmarkCbrt/basic-16                                    	47340079	        26.01 ns/op
BenchmarkCbrt/Cbrt-16                                     	85196262	        14.55 ns/op
BenchmarkPow/basic-16                                     	24005180	        48.25 ns/op
BenchmarkPow/Pow-16                                       	65121390	        19.08 ns/op
BenchmarkLog/uint32/2/basic-16                            	99810775	        12.06 ns/op
BenchmarkLog/uint32/2/Log2-16                             	984283590	         1.223 ns/op
BenchmarkLog/uint32/10/basic-16                           	140540709	         8.516 ns/op
BenchmarkLog/uint32/10/Log10-16                           	539441811	         2.220 ns/op
BenchmarkLog/uint64/2/basic-16                            	100000000	        11.73 ns/op
BenchmarkLog/uint64/2/Log2-16                             	839779903	         1.419 ns/op
BenchmarkLog/uint64/10/basic-16                           	142679388	         8.419 ns/op
BenchmarkLog/uint64/10/Log10-16                           	538269764	         2.228 ns/op
BenchmarkSqrt/basic-16                                    	1000000000	         1.019 ns/op
BenchmarkSqrt/SqrtNewton-16                               	188142513	         6.538 ns/op
BenchmarkSqrt/SqrtBinarySearch-16                         	74752382	        15.71 ns/op
BenchmarkSqrt/SqrtShiftAndSubtract-16                     	136426688	         8.834 ns/op
BenchmarkCRC32/basic-16                                   	220094371	         5.395 ns/op
BenchmarkCRC32/CRC32Basic-16                              	  448540	      2510 ns/op
BenchmarkCRC32/CRC32TableLookup-16                        	 8108901	       147.5 ns/op
BenchmarkRSqrtFloat32/basic-16                            	1000000000	         0.9354 ns/op
BenchmarkRSqrtFloat32/RSqrtFloat32-16                     	828149971	         1.448 ns/op
PASS
ok  	github.com/nikolaydubina/go-hackers-delight	91.183s
```
</details>

[^1]: showcase in `C` — https://github.com/hcs0/Hackers-Delight
[^2]: showcase in `Rust` — https://github.com/victoryang00/Delight
[^3]: given manual inlining of generic type, which produces equivalent Go code
[^4]: we are comparing native float64 result converted to uint32, as there is no better standard function
[^5]: which is due to `float64` not having enough precision
