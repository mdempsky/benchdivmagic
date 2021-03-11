This repo has benchmarks to compare the current "divmagic" code used by the Go
runtime for computing object index / base addresses from span offsets with the
technique described in ["Faster remainders when the divisor is a constant:
beating compilers and
libdivide"](https://lemire.me/blog/2019/02/08/faster-remainders-when-the-divisor-is-a-constant-beating-compilers-and-libdivide/).

```
goos: linux
goarch: amd64
pkg: github.com/mdempsky/benchdivmagic
cpu: Intel(R) Core(TM) i7-8650U CPU @ 1.90GHz
BenchmarkDivMagic_objIndex-8      	  494898	      6729 ns/op
BenchmarkDivMagic_objIndex2-8     	  576932	      6024 ns/op
BenchmarkDivMagic_objIndex3-8     	  616864	      5944 ns/op
BenchmarkDivMagic_objIndex4-8     	  656739	      5165 ns/op
BenchmarkDivMagic_findObject-8    	  515413	      6963 ns/op
BenchmarkDivMagic_findObject2-8   	  577012	      6179 ns/op
BenchmarkDivMagic_findObject3-8   	  541370	      6763 ns/op
BenchmarkDivMagic_findObject4-8   	  542636	      6703 ns/op
BenchmarkLemire_objIndex-8        	 1000000	      3013 ns/op
BenchmarkLemire_findObject-8      	  719557	      4876 ns/op

goos: linux
goarch: amd64
pkg: github.com/mdempsky/benchdivmagic
cpu: Intel(R) Xeon(R) Gold 6154 CPU @ 3.00GHz
BenchmarkDivMagic_objIndex-36       	  421186	      8534 ns/op
BenchmarkDivMagic_objIndex2-36      	  479966	      7496 ns/op
BenchmarkDivMagic_objIndex3-36      	  498529	      7227 ns/op
BenchmarkDivMagic_objIndex4-36      	  546093	      6610 ns/op
BenchmarkDivMagic_findObject-36     	  412866	      8642 ns/op
BenchmarkDivMagic_findObject2-36    	  460458	      7817 ns/op
BenchmarkDivMagic_findObject3-36    	  428541	      8398 ns/op
BenchmarkDivMagic_findObject4-36    	  428371	      8399 ns/op
BenchmarkLemire_objIndex-36         	  923715	      3893 ns/op
BenchmarkLemire_findObject-36       	  571544	      6307 ns/op
```
