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
BenchmarkDivMagic_objIndex-8      	  474774	      7129 ns/op
BenchmarkDivMagic_objIndex2-8     	  569115	      6926 ns/op
BenchmarkDivMagic_objIndex3-8     	  596546	      6171 ns/op
BenchmarkDivMagic_objIndex4-8     	  643726	      5574 ns/op
BenchmarkDivMagic_findObject-8    	  523143	      7707 ns/op
BenchmarkDivMagic_findObject2-8   	  562238	      6980 ns/op
BenchmarkDivMagic_findObject3-8   	  515562	      7435 ns/op
BenchmarkDivMagic_findObject4-8   	  513784	      7625 ns/op
BenchmarkLemire_objIndex-8        	 1000000	      3347 ns/op
BenchmarkLemire_findObject-8      	  799466	      4467 ns/op

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
BenchmarkDivMagic_objIndex-36       	  420691	      8546 ns/op
BenchmarkDivMagic_objIndex2-36      	  480241	      7490 ns/op
BenchmarkDivMagic_objIndex3-36      	  498822	      7255 ns/op
BenchmarkDivMagic_objIndex4-36      	  545398	      6673 ns/op
BenchmarkDivMagic_findObject-36     	  417222	      8620 ns/op
BenchmarkDivMagic_findObject2-36    	  459343	      7833 ns/op
BenchmarkDivMagic_findObject3-36    	  427113	      8410 ns/op
BenchmarkDivMagic_findObject4-36    	  427814	      8412 ns/op
BenchmarkLemire_objIndex-36         	  923226	      3892 ns/op
BenchmarkLemire_findObject-36       	  673830	      5332 ns/op

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
