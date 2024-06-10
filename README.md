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

* not using any Go packages, not even standard library

<details><summary>Appendix: Benchmarks</summary>

```bash
$ go test -bench .        
goos: darwin
goarch: arm64
pkg: github.com/nikolaydubina/go-hackers-delight
BenchmarkCompress/Compress-16         	            100000000	        10.05 ns/op
BenchmarkCompress/Compress2-16        	            59208940	        20.69 ns/op
BenchmarkDivMod/DivMod/3/basic-16     	            1000000000	         0.843 ns/op
BenchmarkDivMod/DivMod/3/DivMod3Signed-16         	615786322	         1.931 ns/op
BenchmarkDivMod/DivMod/3/DivMod3Signed2-16        	1000000000	         1.091 ns/op
BenchmarkDivMod/DivMod/7/basic-16                 	1000000000	         0.836 ns/op
BenchmarkDivMod/DivMod/7/DivMod7Signed-16         	578422084	         2.074 ns/op
BenchmarkDivMod/Div/3/basic-16                    	1000000000	         0.834 ns/op
BenchmarkDivMod/Div/3/Div3Signed-16               	791803947	         1.522 ns/op
BenchmarkDivMod/Div/3/Div3ShiftSigned-16          	903334430	         1.328 ns/op
BenchmarkDivMod/Div/7/basic-16                    	1000000000	         0.836 ns/op
BenchmarkDivMod/Div/7/Div7Signed-16               	753299805	         1.607 ns/op
BenchmarkDivMod/Div/7/Div7ShiftSigned-16          	824427794	         1.455 ns/op
BenchmarkDivMod/DivExact/7/basic-16               	1000000000	         1.095 ns/op
BenchmarkDivMod/DivExact/7/DivExact7-16           	1000000000	         1.105 ns/op
BenchmarkDivMod/DivExact/7/Div7Signed-16          	709693124	         1.693 ns/op
BenchmarkDivMod/DivExact/7/Div7ShiftSigned-16     	793241796	         1.509 ns/op
PASS
ok  	github.com/nikolaydubina/go-hackers-delight	23.260s
```
</details>

[^1]: showcase in `C` — https://github.com/hcs0/Hackers-Delight
[^2]: showcase in `Rust` — https://github.com/victoryang00/Delight
