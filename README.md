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

* native `math.Pow(x, 1./3)`[^4] performance is worse than `Cbrt` algorithm 💥
* native `math.Pow(x, 1./3)`[^4][^5] underflows but `Cbrt` algorithm is correct 💥
* native `math.Pow(x, n)`[^4] performance is worse than `Pow` algorithm 💥
* native `math.Log2(x)`[^4] performance is worse than `Log2` algorithm 💥
* native `math.Log10(x)`[^4] performance is worse than `Log10` algorithm 💥
* native `math.Log10(x)` [^4][^5] overflows for `math.MaxUint64` but `Log10` algorithm is correct 💥
* simplistic `LRUCache` performance is `2x` worse than `LRUCache` algorithm 💥 

<details><summary>Appendix: Benchmarks</summary>

```bash
$ go test -bench .
goos: darwin
goarch: arm64
pkg: github.com/nikolaydubina/go-hackers-delight
cpu: Apple M3 Max
BenchmarkNoop/---------------------------------16         	1000000000	         0.0000001 ns/op
BenchmarkAbs/basic-16                                     	 1489353	       811.6 ns/op
BenchmarkAbs/Abs-16                                       	 1460946	       818.8 ns/op
BenchmarkAbs/Abs2-16                                      	 1464883	       816.4 ns/op
BenchmarkAbs/Abs3-16                                      	 1443763	       830.9 ns/op
BenchmarkAbs/Abs4-16                                      	 1431141	       882.4 ns/op
BenchmarkAbs/AbsFastMul-16                                	 1494900	       803.1 ns/op
BenchmarkAvg/basic-16                                     	  611107	      1970 ns/op
BenchmarkAvg/AvgFloor-16                                  	  811885	      1603 ns/op
BenchmarkAvg/AvgCeil-16                                   	  609648	      1916 ns/op
BenchmarkCycleThree/basic-16                              	1000000000	         1.105 ns/op
BenchmarkCycleThree/CycleThreeValues-16                   	1000000000	         1.000 ns/op
BenchmarkLeadingZeros/uint32/basic-16                     	 1474354	       816.2 ns/op
BenchmarkLeadingZeros/uint32/LeadingZerosUint32-16        	 1000000	      1174 ns/op
BenchmarkLeadingZeros/uint64/basic-16                     	 1483987	       808.8 ns/op
BenchmarkLeadingZeros/uint64/LeadingZerosUint64-16        	  892615	      1349 ns/op
BenchmarkCompress/Compress-16                             	150361084	         8.064 ns/op
BenchmarkCompress/Compress2-16                            	54269787	        21.47 ns/op
BenchmarkLRU/basic-16                                     	447112422	         2.689 ns/op
BenchmarkLRU/LRUCache-16                                  	625848724	         1.878 ns/op
BenchmarkMul/uint32/basic-16                              	  610756	      1963 ns/op
BenchmarkMul/uint32/MultiplyHighOrder32-16                	  590781	      1851 ns/op
BenchmarkMul/uint64/basic-16                              	  109549	     11017 ns/op
BenchmarkMul/uint64/MultiplyHighOrder64-16                	   71486	     19700 ns/op
BenchmarkDivMod/DivMod/3/basic-16                         	 1483515	       809.9 ns/op
BenchmarkDivMod/DivMod/3/DivMod3Signed-16                 	  630376	      1903 ns/op
BenchmarkDivMod/DivMod/3/DivMod3Signed2-16                	 1000000	      1078 ns/op
BenchmarkDivMod/DivMod/7/basic-16                         	 1470925	       813.2 ns/op
BenchmarkDivMod/DivMod/7/DivMod7Signed-16                 	  617737	      1977 ns/op
BenchmarkDivMod/Div/3/basic-16                            	 1474382	       816.0 ns/op
BenchmarkDivMod/Div/3/Div3Signed-16                       	  855132	      1387 ns/op
BenchmarkDivMod/Div/3/Div3ShiftSigned-16                  	  958530	      1245 ns/op
BenchmarkDivMod/Div/7/basic-16                            	 1480100	       808.9 ns/op
BenchmarkDivMod/Div/7/Div7Signed-16                       	  831228	      1438 ns/op
BenchmarkDivMod/Div/7/Div7ShiftSigned-16                  	  890666	      1347 ns/op
BenchmarkDivMod/Mod/3/basic-16                            	 1494051	       807.6 ns/op
BenchmarkDivMod/Mod/3/Mod3Signed-16                       	  896326	      1306 ns/op
BenchmarkDivMod/Mod/3/Mod3Signed2-16                      	 1491132	       804.5 ns/op
BenchmarkDivMod/Mod/7/basic-16                            	 1489940	       804.5 ns/op
BenchmarkDivMod/Mod/7/Mod7Signed-16                       	  865971	      1394 ns/op
BenchmarkDivMod/Mod/7/Mod7Signed2-16                      	 1000000	      1108 ns/op
BenchmarkDivMod/Mod/10/basic-16                           	 1470753	       816.9 ns/op
BenchmarkDivMod/Mod/10/Mod10Signed-16                     	 1000000	      1113 ns/op
BenchmarkDivMod/DivExact/7/basic-16                       	 1492069	       803.8 ns/op
BenchmarkDivMod/DivExact/7/DivExact7-16                   	 1452621	       820.2 ns/op
BenchmarkDivMod/DivExact/7/Div7Signed-16                  	  763095	      1558 ns/op
BenchmarkDivMod/DivExact/7/Div7ShiftSigned-16             	  869890	      1382 ns/op
BenchmarkCbrt/basic-16                                    	    4669	    260382 ns/op
BenchmarkCbrt/Cbrt-16                                     	    7956	    150951 ns/op
BenchmarkPow/basic-16                                     	    2311	    520212 ns/op
BenchmarkPow/Pow-16                                       	    6334	    188887 ns/op
BenchmarkLog/uint32/2/basic-16                            	   10000	    120539 ns/op
BenchmarkLog/uint32/2/Log2-16                             	   96388	     12463 ns/op
BenchmarkLog/uint32/10/basic-16                           	   13887	     86154 ns/op
BenchmarkLog/uint32/10/Log10-16                           	   54652	     21962 ns/op
BenchmarkLog/uint64/2/basic-16                            	   10000	    120627 ns/op
BenchmarkLog/uint64/2/Log2-16                             	   81048	     15237 ns/op
BenchmarkLog/uint64/10/basic-16                           	   14146	     84683 ns/op
BenchmarkLog/uint64/10/Log10-16                           	   51010	     23544 ns/op
BenchmarkSqrt/basic-16                                    	  134250	      8954 ns/op
BenchmarkSqrt/SqrtNewton-16                               	   18208	     63912 ns/op
BenchmarkSqrt/SqrtBinarySearch-16                         	    7302	    168653 ns/op
BenchmarkSqrt/SqrtShiftAndSubtract-16                     	   10000	    114464 ns/op
BenchmarkCRC32/basic-16                                   	  181860	      6610 ns/op
BenchmarkCRC32/CRC32Basic-16                              	     434	   2778007 ns/op
BenchmarkCRC32/CRC32TableLookup-16                        	    7068	    168972 ns/op
BenchmarkRSqrtFloat32/basic-16                            	  139954	      8663 ns/op
BenchmarkRSqrtFloat32/RSqrtFloat32-16                     	   99153	     12287 ns/op
PASS
ok  	github.com/nikolaydubina/go-hackers-delight	83.522s

```
</details>

[^1]: showcase in `C` — https://github.com/hcs0/Hackers-Delight
[^2]: showcase in `Rust` — https://github.com/victoryang00/Delight
[^4]: we are comparing native float64 result converted to uint32, as there is no better standard function
[^5]: which is due to `float64` not having enough precision
