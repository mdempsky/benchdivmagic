This repo has benchmarks to compare the current "divmagic" code used by the Go
runtime for computing object index / base addresses from span offsets with the
technique described in ["Faster remainders when the divisor is a constant:
beating compilers and
libdivide"](https://lemire.me/blog/2019/02/08/faster-remainders-when-the-divisor-is-a-constant-beating-compilers-and-libdivide/).

Benchmarked at golang/go#e8b82789cda6c9d9e3dfc9a652b4d7a823b834f2.

`BenchmarkDivMagic_objIndex` and `BenchmarkDivMagic_findObject` are the code
currently used by the Go runtime. The 2, 3, and 4 variants are attempts at
simplifying them to improve performance (e.g., masking shift counts so the
compiler doesn't need to insert bounds checks, or eliminating additional
comparisons to reduce branchiness).

`BenchmarkLemire_objIndex` and `BenchmarkLemire_findObject` are the rewrites
included in golang.org/cl/300994.

Lenovo ThinkPad X1 Gen 6 (Kaby Lake R):

```
goos: linux
goarch: amd64
pkg: github.com/mdempsky/benchdivmagic
cpu: Intel(R) Core(TM) i7-8650U CPU @ 1.90GHz
BenchmarkDivMagic_objIndex-8      	  808069	      7246 ns/op
BenchmarkDivMagic_objIndex2-8     	  887394	      6999 ns/op
BenchmarkDivMagic_objIndex3-8     	  933570	      6569 ns/op
BenchmarkDivMagic_objIndex4-8     	  993865	      5510 ns/op
BenchmarkDivMagic_findObject-8    	  827413	      7357 ns/op
BenchmarkDivMagic_findObject2-8   	  922064	      6500 ns/op
BenchmarkDivMagic_findObject3-8   	  806523	      7090 ns/op
BenchmarkDivMagic_findObject4-8   	  864450	      7208 ns/op
BenchmarkLemire_objIndex-8        	 2177156	      2709 ns/op
BenchmarkLemire_findObject-8      	 1697712	      3568 ns/op

goos: linux
goarch: 386
pkg: github.com/mdempsky/benchdivmagic
cpu: Intel(R) Core(TM) i7-8650U CPU @ 1.90GHz
BenchmarkDivMagic_objIndex-8      	  326536	     18334 ns/op
BenchmarkDivMagic_objIndex2-8     	  562065	     11277 ns/op
BenchmarkDivMagic_objIndex3-8     	  525712	     11556 ns/op
BenchmarkDivMagic_objIndex4-8     	  300144	     19092 ns/op
BenchmarkDivMagic_findObject-8    	  611988	      9641 ns/op
BenchmarkDivMagic_findObject2-8   	  731283	      7716 ns/op
BenchmarkDivMagic_findObject3-8   	  790094	      7658 ns/op
BenchmarkDivMagic_findObject4-8   	  829664	      7214 ns/op
BenchmarkLemire_objIndex-8        	 1657843	      3587 ns/op
BenchmarkLemire_findObject-8      	 1000000	      5175 ns/op
```

Lenovo ThinkStation P920 (Skylake):

```
goos: linux
goarch: amd64
pkg: github.com/mdempsky/benchdivmagic
cpu: Intel(R) Xeon(R) Gold 6154 CPU @ 3.00GHz
BenchmarkDivMagic_objIndex-36       	  701734	      8636 ns/op
BenchmarkDivMagic_objIndex2-36      	  795364	      7485 ns/op
BenchmarkDivMagic_objIndex3-36      	  823696	      7215 ns/op
BenchmarkDivMagic_objIndex4-36      	  908626	      6600 ns/op
BenchmarkDivMagic_findObject-36     	  691534	      8673 ns/op
BenchmarkDivMagic_findObject2-36    	  766933	      7822 ns/op
BenchmarkDivMagic_findObject3-36    	  713942	      8406 ns/op
BenchmarkDivMagic_findObject4-36    	  714055	      8405 ns/op
BenchmarkLemire_objIndex-36         	 1896676	      3162 ns/op
BenchmarkLemire_findObject-36       	 1371702	      4388 ns/op

goos: linux
goarch: 386
pkg: github.com/mdempsky/benchdivmagic
cpu: Intel(R) Xeon(R) Gold 6154 CPU @ 3.00GHz
BenchmarkDivMagic_objIndex-36       	  300124	     20509 ns/op
BenchmarkDivMagic_objIndex2-36      	  472482	     12865 ns/op
BenchmarkDivMagic_objIndex3-36      	  458918	     13066 ns/op
BenchmarkDivMagic_objIndex4-36      	  247666	     24123 ns/op
BenchmarkDivMagic_findObject-36     	  525354	     11339 ns/op
BenchmarkDivMagic_findObject2-36    	  675114	      8815 ns/op
BenchmarkDivMagic_findObject3-36    	  643443	      9322 ns/op
BenchmarkDivMagic_findObject4-36    	  714792	      8391 ns/op
BenchmarkLemire_objIndex-36         	 1372545	      4386 ns/op
BenchmarkLemire_findObject-36       	 1000000	      5825 ns/op
```
