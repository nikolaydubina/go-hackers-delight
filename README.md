# go-hackers-delight

[![codecov](https://codecov.io/gh/nikolaydubina/go-hackers-delight/graph/badge.svg?token=660JQtUmiO)](https://codecov.io/gh/nikolaydubina/go-hackers-delight)
![fuzzing](https://img.shields.io/badge/fuzzing-active-brightgreen)
[![Go Reference](https://pkg.go.dev/github.com/nikolaydubina/go-hackers-delight)

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
BenchmarkAbs/basic-16                                     	1000000000	         0.9330 ns/op
BenchmarkAbs/Abs-16                                       	1000000000	         0.9326 ns/op
BenchmarkAbs/Abs2-16                                      	1000000000	         0.9362 ns/op
BenchmarkAbs/Abs3-16                                      	1000000000	         0.9324 ns/op
BenchmarkAbs/Abs4-16                                      	1000000000	         0.9414 ns/op
BenchmarkAbs/AbsFastMul-16                                	1000000000	         0.9493 ns/op
BenchmarkAvg/basic-16                                     	590327602	         2.045 ns/op
BenchmarkAvg/AvgFloor-16                                  	594028404	         2.004 ns/op
BenchmarkAvg/AvgCeil-16                                   	592978882	         2.020 ns/op
BenchmarkCycleThree/basic-16                              	787989288	         1.519 ns/op
BenchmarkCycleThree/CycleThreeValues-16                   	474497178	         2.510 ns/op
BenchmarkLeadingZeros/uint32/basic-16                     	1000000000	         0.9270 ns/op
BenchmarkLeadingZeros/uint32/LeadingZerosUint32-16        	1000000000	         1.136 ns/op
BenchmarkLeadingZeros/uint64/basic-16                     	1000000000	         1.099 ns/op
BenchmarkLeadingZeros/uint64/LeadingZerosUint32-16        	877795342	         1.512 ns/op
BenchmarkCompress/Compress-16                             	100000000	        10.77 ns/op
BenchmarkCompress/Compress2-16                            	58750090	        21.02 ns/op
BenchmarkLRU/basic-16                                     	239493868	         4.978 ns/op
BenchmarkLRU/LeadingZerosUint32-16                        	981022663	         1.223 ns/op
BenchmarkMul/uint32/basic-16                              	656605255	         2.006 ns/op
BenchmarkMul/uint32/MultiplyHighOrder32-16                	591757071	         1.743 ns/op
BenchmarkMul/uint64/basic-16                              	993127761	         1.207 ns/op
BenchmarkMul/uint64/MultiplyHighOrder64-16                	975349400	         2.025 ns/op
BenchmarkDivMod/DivMod/3/basic-16                         	1000000000	         0.8625 ns/op
BenchmarkDivMod/DivMod/3/DivMod3Signed-16                 	607501119	         1.970 ns/op
BenchmarkDivMod/DivMod/3/DivMod3Signed2-16                	1000000000	         1.117 ns/op
BenchmarkDivMod/DivMod/7/basic-16                         	1000000000	         0.8537 ns/op
BenchmarkDivMod/DivMod/7/DivMod7Signed-16                 	562010763	         2.144 ns/op
BenchmarkDivMod/Div/3/basic-16                            	1000000000	         0.8432 ns/op
BenchmarkDivMod/Div/3/Div3Signed-16                       	794080347	         1.528 ns/op
BenchmarkDivMod/Div/3/Div3ShiftSigned-16                  	892143838	         1.332 ns/op
BenchmarkDivMod/Div/7/basic-16                            	1000000000	         0.8831 ns/op
BenchmarkDivMod/Div/7/Div7Signed-16                       	725504583	         1.657 ns/op
BenchmarkDivMod/Div/7/Div7ShiftSigned-16                  	828426534	         1.443 ns/op
BenchmarkDivMod/Mod/3/basic-16                            	1000000000	         0.8523 ns/op
BenchmarkDivMod/Mod/3/Mod3Signed-16                       	780382692	         1.518 ns/op
BenchmarkDivMod/Mod/3/Mod3Signed2-16                      	1000000000	         0.8461 ns/op
BenchmarkDivMod/Mod/7/basic-16                            	1000000000	         0.8449 ns/op
BenchmarkDivMod/Mod/7/Mod7Signed-16                       	747462322	         1.631 ns/op
BenchmarkDivMod/Mod/7/Mod7Signed2-16                      	1000000000	         1.100 ns/op
BenchmarkDivMod/Mod/10/basic-16                           	1000000000	         0.8332 ns/op
BenchmarkDivMod/Mod/10/Mod10Signed-16                     	815590462	         1.436 ns/op
BenchmarkDivMod/DivExact/7/basic-16                       	1000000000	         0.9324 ns/op
BenchmarkDivMod/DivExact/7/DivExact7-16                   	1000000000	         0.9513 ns/op
BenchmarkDivMod/DivExact/7/Div7Signed-16                  	698338923	         1.695 ns/op
BenchmarkDivMod/DivExact/7/Div7ShiftSigned-16             	802803792	         1.504 ns/op
BenchmarkCbrt/basic-16                                    	46489705	        26.13 ns/op
BenchmarkCbrt/Cbrt-16                                     	82498086	        15.14 ns/op
BenchmarkPow/basic-16                                     	22573610	        51.55 ns/op
BenchmarkPow/Pow-16                                       	62801604	        19.12 ns/op
BenchmarkLog/uint32/2/basic-16                            	100000000	        11.93 ns/op
BenchmarkLog/uint32/2/Log2-16                             	977360526	         1.228 ns/op
BenchmarkLog/uint32/10/basic-16                           	139321425	         8.559 ns/op
BenchmarkLog/uint32/10/Log10-16                           	539884531	         2.237 ns/op
BenchmarkLog/uint64/2/basic-16                            	100000000	        11.75 ns/op
BenchmarkLog/uint64/2/Log2-16                             	845898448	         1.419 ns/op
BenchmarkLog/uint64/10/basic-16                           	144678940	         8.284 ns/op
BenchmarkLog/uint64/10/Log10-16                           	539179033	         2.222 ns/op
BenchmarkSqrt/basic-16                                    	1000000000	         1.019 ns/op
BenchmarkSqrt/SqrtNewton-16                               	170914454	         6.656 ns/op
BenchmarkSqrt/SqrtBinarySearch-16                         	73125955	        16.40 ns/op
BenchmarkSqrt/SqrtShiftAndSubtract-16                     	135757068	         8.813 ns/op
BenchmarkCRC32/basic-16                                   	222373198	         5.397 ns/op
BenchmarkCRC32/CRC32Basic-16                              	  447698	      2522 ns/op
BenchmarkCRC32/CRC32TableLookup-16                        	 8125387	       147.7 ns/op
BenchmarkRSqrtFloat32/basic-16                            	1000000000	         0.9349 ns/op
BenchmarkRSqrtFloat32/CRC32Basic-16                       	830151805	         1.447 ns/op
PASS
ok  	github.com/nikolaydubina/go-hackers-delight	89.462s
```
</details>

[^1]: showcase in `C` — https://github.com/hcs0/Hackers-Delight
[^2]: showcase in `Rust` — https://github.com/victoryang00/Delight
[^3]: given manual inlining of generic type, which produces equivalent Go code
[^4]: we are comparing native float64 result converted to uint32, as there is no better standard function
[^5]: which is due to `float64` not having enough precision
