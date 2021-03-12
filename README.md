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
BenchmarkDivMagic_objIndex-8      	  431764	      7190 ns/op
BenchmarkDivMagic_objIndex2-8     	  500112	      6887 ns/op
BenchmarkDivMagic_objIndex3-8     	  580312	      6245 ns/op
BenchmarkDivMagic_objIndex4-8     	  625574	      5821 ns/op
BenchmarkDivMagic_findObject-8    	  438514	      7527 ns/op
BenchmarkDivMagic_findObject2-8   	  301363	     11383 ns/op
BenchmarkDivMagic_findObject3-8   	  436406	      7122 ns/op
BenchmarkDivMagic_findObject4-8   	  508644	      7297 ns/op
BenchmarkLemire_objIndex-8        	 1000000	      3081 ns/op
BenchmarkLemire_findObject-8      	  845827	      4081 ns/op

goos: linux
goarch: 386
pkg: github.com/mdempsky/benchdivmagic
cpu: Intel(R) Core(TM) i7-8650U CPU @ 1.90GHz
BenchmarkDivMagic_objIndex-8      	  204361	     18496 ns/op
BenchmarkDivMagic_objIndex2-8     	  324112	     11293 ns/op
BenchmarkDivMagic_objIndex3-8     	  317346	     11666 ns/op
BenchmarkDivMagic_objIndex4-8     	  141831	     22216 ns/op
BenchmarkDivMagic_findObject-8    	  287307	     12174 ns/op
BenchmarkDivMagic_findObject2-8   	  405092	      8551 ns/op
BenchmarkDivMagic_findObject3-8   	  508814	      7591 ns/op
BenchmarkDivMagic_findObject4-8   	  435669	      9405 ns/op
BenchmarkLemire_objIndex-8        	  965276	      3801 ns/op
BenchmarkLemire_findObject-8      	  676478	      5309 ns/op

goos: linux
goarch: amd64
pkg: github.com/mdempsky/benchdivmagic
cpu: Intel(R) Xeon(R) Gold 6154 CPU @ 3.00GHz
BenchmarkDivMagic_objIndex-36       	  421322	      8587 ns/op
BenchmarkDivMagic_objIndex2-36      	  475767	      7491 ns/op
BenchmarkDivMagic_objIndex3-36      	  497272	      7271 ns/op
BenchmarkDivMagic_objIndex4-36      	  541953	      6599 ns/op
BenchmarkDivMagic_findObject-36     	  417238	      8630 ns/op
BenchmarkDivMagic_findObject2-36    	  277136	     13007 ns/op
BenchmarkDivMagic_findObject3-36    	  428490	      8399 ns/op
BenchmarkDivMagic_findObject4-36    	  428677	      8401 ns/op
BenchmarkLemire_objIndex-36         	 1000000	      3408 ns/op
BenchmarkLemire_findObject-36       	  740534	      4862 ns/op

goos: linux
goarch: 386
pkg: github.com/mdempsky/benchdivmagic
cpu: Intel(R) Xeon(R) Gold 6154 CPU @ 3.00GHz
BenchmarkDivMagic_objIndex-36       	  174018	     20755 ns/op
BenchmarkDivMagic_objIndex2-36      	  280549	     12818 ns/op
BenchmarkDivMagic_objIndex3-36      	  275816	     13047 ns/op
BenchmarkDivMagic_objIndex4-36      	  152755	     24072 ns/op
BenchmarkDivMagic_findObject-36     	  268599	     13400 ns/op
BenchmarkDivMagic_findObject2-36    	  374582	      9629 ns/op
BenchmarkDivMagic_findObject3-36    	  428865	      8388 ns/op
BenchmarkDivMagic_findObject4-36    	  385939	      9319 ns/op
BenchmarkLemire_objIndex-36         	  822868	      4371 ns/op
BenchmarkLemire_findObject-36       	  617586	      5822 ns/op
```
