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
* verifying compiled code via https://godbolt.org/ (e.g. inlining, generics)
</details>

<details><summary>Appendix: Benchmarks</summary>

```bash
$ go test -bench .        
goos: darwin
goarch: arm64
pkg: github.com/nikolaydubina/go-hackers-delight
BenchmarkCompress/Compress-16         	            100000000	        10.37 ns/op
BenchmarkCompress/Compress2-16        	            58364338	        20.67 ns/op
BenchmarkDivMod/DivMod/3/basic-16     	            1000000000	         0.8335 ns/op
BenchmarkDivMod/DivMod/3/DivMod3Signed-16         	622294209	         1.919 ns/op
BenchmarkDivMod/DivMod/3/DivMod3Signed2-16        	1000000000	         1.074 ns/op
BenchmarkDivMod/DivMod/7/basic-16                 	1000000000	         0.8324 ns/op
BenchmarkDivMod/DivMod/7/DivMod7Signed-16         	578056261	         2.071 ns/op
BenchmarkDivMod/Div/3/basic-16                    	1000000000	         0.8361 ns/op
BenchmarkDivMod/Div/3/Div3Signed-16               	793247480	         1.542 ns/op
BenchmarkDivMod/Div/3/Div3ShiftSigned-16          	908296149	         1.322 ns/op
BenchmarkDivMod/Div/7/basic-16                    	1000000000	         0.8328 ns/op
BenchmarkDivMod/Div/7/Div7Signed-16               	755478798	         1.591 ns/op
BenchmarkDivMod/Div/7/Div7ShiftSigned-16          	842272730	         1.425 ns/op
BenchmarkDivMod/Mod/3/basic-16                    	1000000000	         0.8355 ns/op
BenchmarkDivMod/Mod/3/Mod3Signed-16               	809193108	         1.482 ns/op
BenchmarkDivMod/Mod/3/Mod3Signed2-16              	1000000000	         0.8327 ns/op
BenchmarkDivMod/Mod/7/basic-16                    	1000000000	         0.8500 ns/op
BenchmarkDivMod/Mod/7/Mod7Signed-16               	768375580	         1.564 ns/op
BenchmarkDivMod/Mod/7/Mod7Signed2-16              	1000000000	         1.105 ns/op
BenchmarkDivMod/Mod/10/basic-16                   	1000000000	         0.8494 ns/op
BenchmarkDivMod/Mod/10/Mod10Signed-16             	869101016	         1.454 ns/op
BenchmarkDivMod/DivExact/7/basic-16               	1000000000	         0.9297 ns/op
BenchmarkDivMod/DivExact/7/DivExact7-16           	1000000000	         0.9259 ns/op
BenchmarkDivMod/DivExact/7/Div7Signed-16          	719832459	         1.670 ns/op
BenchmarkDivMod/DivExact/7/Div7ShiftSigned-16     	801826603	         1.493 ns/op
PASS
ok  	github.com/nikolaydubina/go-hackers-delight	32.676s
```
</details>

[^1]: showcase in `C` — https://github.com/hcs0/Hackers-Delight
[^2]: showcase in `Rust` — https://github.com/victoryang00/Delight
