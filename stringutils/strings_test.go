// Copyright 2020 Michael Treanor. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package strings implements additional functions to support the go library:
// .../go/1.15.3/libexec/src/strings/strings.go
//
// For information about UTF-8 strings in Go, see https://blog.golang.org/strings.
package stringutils

import (
	"math/rand"
	"time"
	"unicode"
	"unicode/utf8"
)

const (
    TAB = 0x09  // '\t'
    LF = 0x0A   // '\n'
    VT = 0x0B   // '\v'
    FF = 0x0C   // '\f'
    CR = 0x0D   // '\r'
    SPACE = ' '
    RuneSelf = utf8.RuneSelf
    NBSP = 0x00A0
    NEL = 0x0085

    defaultSamples = 1<<8 - 1
    maxSamples     = 1<<32 - 1

    numSamples = 1<<4
)

// Benchmark results
/*
IsWhiteSpace switch/case is significantly slower than multiple || with 6 checks

BenchmarkIsAlphaNum-8      	770840788	         1.51 ns/op	       0 B/op	       0 allocs/op
// removing the 'if' and simply returning the boolean result is 50% faster
BenchmarkIsAlphaNum2-8     	1000000000	         0.923 ns/op	       0 B/op	       0 allocs/op


BenchmarkIsAlpha-8         	967973194	         1.31 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsDigit-8         	971392962	         1.21 ns/op	       0 B/op	       0 allocs/op

Small sample size
================================================================
BenchmarkIsWhiteSpace-8    	227064271	         5.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsWhiteSpace2-8   	235199743	         5.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaSwitch-8   	585296522	         1.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaNum-8      	124885717	         9.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaNum2-8     	195595128	         6.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlpha-8         	990591926	         1.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsDigit-8         	937417184	         1.30 ns/op	       0 B/op	       0 allocs/op

================================================================
With ByteSamples() and RuneSamples()  ... consistent samples

BenchmarkIsAlpha-8         	47187997	        25.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsDigit-8         	51663846	        21.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaSwitch-8   	36453657	        32.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsWhiteSpace-8    	162498454	         7.37 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsWhiteSpace2-8   	48987033	        24.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaNum-8      	43917738	        26.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaNum2-8     	58560900	        19.4 ns/op	       0 B/op	       0 allocs/op

================================================================
sample size x8

BenchmarkIsAlpha-8         	 4349080	       273 ns/op	     144 B/op	       1 allocs/op
BenchmarkIsDigit-8         	 5197915	       231 ns/op	     144 B/op	       1 allocs/op
BenchmarkIsAlphaSwitch-8   	 3905835	       309 ns/op	     144 B/op	       1 allocs/op
BenchmarkIsWhiteSpace-8    	 6110427	       196 ns/op	     576 B/op	       1 allocs/op
BenchmarkIsWhiteSpace2-8   	 3619970	       332 ns/op	     576 B/op	       1 allocs/op
BenchmarkIsAlphaNum-8      	 4399827	       282 ns/op	     144 B/op	       1 allocs/op
BenchmarkIsAlphaNum2-8     	 2731586	       444 ns/op	     144 B/op	       1 allocs/op

================================================================
using n = 1024
BenchmarkIsAlpha-8         	   39996	     30380 ns/op	    1024 B/op	       1 allocs/op
BenchmarkIsDigit-8         	   42748	     28097 ns/op	    1024 B/op	       1 allocs/op
BenchmarkIsAlphaSwitch-8   	   39072	     31276 ns/op	    1024 B/op	       1 allocs/op
BenchmarkIsWhiteSpace-8    	   44443	     27086 ns/op	    5120 B/op	       2 allocs/op
BenchmarkIsWhiteSpace2-8   	   43936	     27800 ns/op	    5120 B/op	       2 allocs/op
BenchmarkIsAlphaNum-8      	   40113	     29856 ns/op	    1024 B/op	       1 allocs/op
BenchmarkIsAlphaNum2-8     	   43183	     28851 ns/op	    1024 B/op	       1 allocs/op

(without preallocation of rune make buffer; e.g. the 'n' in retval := make([]byte, 0, n)
BenchmarkIsWhiteSpace-8    	   43213	     28103 ns/op	    9208 B/op	      11 allocs/op
BenchmarkIsWhiteSpace2-8   	   41742	     29355 ns/op	    9208 B/op	      11 allocs/op

(without preallocation of either buffer)
BenchmarkIsAlpha-8         	   39459	     30819 ns/op	    2040 B/op	       8 allocs/op
BenchmarkIsDigit-8         	   44440	     26921 ns/op	    2040 B/op	       8 allocs/op
BenchmarkIsAlphaSwitch-8   	   39086	     31031 ns/op	    2040 B/op	       8 allocs/op
BenchmarkIsWhiteSpace-8    	   42594	     28711 ns/op	   10224 B/op	      18 allocs/op
BenchmarkIsWhiteSpace2-8   	   40940	     30592 ns/op	   10224 B/op	      18 allocs/op
BenchmarkIsAlphaNum-8      	   39778	     30075 ns/op	    2040 B/op	       8 allocs/op
BenchmarkIsAlphaNum2-8     	   43225	     28086 ns/op	    2040 B/op	       8 allocs/op

================================================================
n = 65535 (with preallocation) ( ... preallocation is generally good)

BenchmarkIsAlpha-8         	     619	   1899081 ns/op	   65539 B/op	       1 allocs/op
BenchmarkIsDigit-8         	     721	   1660150 ns/op	   65536 B/op	       1 allocs/op
BenchmarkIsAlphaSwitch-8   	     636	   1888789 ns/op	   65536 B/op	       1 allocs/op
BenchmarkIsWhiteSpace-8    	     708	   1717852 ns/op	  327681 B/op	       2 allocs/op
BenchmarkIsWhiteSpace2-8   	     688	   1718280 ns/op	  327681 B/op	       2 allocs/op
BenchmarkIsAlphaNum-8      	     643	   1875367 ns/op	   65536 B/op	       1 allocs/op
BenchmarkIsAlphaNum2-8     	     690	   1746706 ns/op	   65536 B/op	       1 allocs/op

(with special 'common' case of space)
BenchmarkIsWhiteSpace-8    	     657	   1740744 ns/op	  327681 B/op	       2 allocs/op
BenchmarkIsWhiteSpace2-8   	     676	   1800894 ns/op	  327681 B/op	       2 allocs/op

(no preallocation)
BenchmarkIsAlpha-8         	     613	   1914118 ns/op	  284666 B/op	      23 allocs/op
BenchmarkIsDigit-8         	     702	   1726401 ns/op	  284666 B/op	      23 allocs/op
BenchmarkIsAlphaSwitch-8   	     594	   1990682 ns/op	  284669 B/op	      23 allocs/op
BenchmarkIsWhiteSpace-8    	     608	   1941723 ns/op	 1693433 B/op	      49 allocs/op
BenchmarkIsWhiteSpace2-8   	     594	   2052508 ns/op	 1693435 B/op	      49 allocs/op
BenchmarkIsAlphaNum-8      	     614	   1972096 ns/op	  284665 B/op	      23 allocs/op
BenchmarkIsAlphaNum2-8     	     666	   1828022 ns/op	  284666 B/op	      23 allocs/op

*/

// Alternate methods that generated worse results...
/*
================================================================
using declared variables for ByteSamples() and RuneSamples()  byteSamples / runeSamples (this is slightly slower??  )

BenchmarkIsAlpha-8         	44298718	        29.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsDigit-8         	49668928	        25.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaSwitch-8   	35638029	        33.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsWhiteSpace-8    	124931196	         9.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsWhiteSpace2-8   	51955320	        23.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaNum-8      	44332395	        26.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaNum2-8     	74700576	        15.9 ns/op	       0 B/op	       0 allocs/op

================================================================
using declared variables for func names ... this is MUCH slower .. it must be preventing compiler optimization since the
function is a variable and might ... vary. Thus cannot be replaced at compile time (??)
// byteSamples  = ByteSamples
// for _, c := range byteSamples()

BenchmarkIsAlpha-8         	20974160	        53.0 ns/op	      32 B/op	       1 allocs/op
BenchmarkIsDigit-8         	25314271	        49.0 ns/op	      32 B/op	       1 allocs/op
BenchmarkIsAlphaSwitch-8   	21365545	        56.5 ns/op	      32 B/op	       1 allocs/op
BenchmarkIsWhiteSpace-8    	26824587	        45.6 ns/op	      80 B/op	       1 allocs/op
BenchmarkIsWhiteSpace2-8   	19660117	        59.6 ns/op	      80 B/op	       1 allocs/op
BenchmarkIsAlphaNum-8      	22809402	        53.9 ns/op	      32 B/op	       1 allocs/op
BenchmarkIsAlphaNum2-8     	25673344	        46.1 ns/op	      32 B/op	       1 allocs/op

================================================================
using maps for checks (3 [string]string/ 4 [rune]string ) (this is much worse than I expected!!)
BenchmarkIsWhiteSpace-8    	     708	   1681603 ns/op	  327680 B/op	       2 allocs/op
BenchmarkIsWhiteSpace2-8   	     666	   1775882 ns/op	  327681 B/op	       2 allocs/op

 using map [string]string
BenchmarkIsWhiteSpace3-8   	      18	  63820745 ns/op	76760858 B/op	   68792 allocs/op
 using [rune]string
BenchmarkIsWhiteSpace4-8   	      22	  51020315 ns/op	47036008 B/op	   68787 allocs/op
 using [rune]bool (returned on the fly from a function call)
BenchmarkIsWhiteSpace5-8   	      30	  38028585 ns/op	15449897 B/op	   78547 allocs/op
 using [rune]bool from var(map literal)
BenchmarkIsWhiteSpace5-8   	     417	   2779469 ns/op	  327681 B/op	       2 allocs/op


using pre-made maps (var = map literal) for 4 and 5 .. 3 is still generated within function
BenchmarkIsWhiteSpace-8    	10117532	       119 ns/op	      16 B/op	       2 allocs/op
BenchmarkIsWhiteSpace2-8   	 9799669	       117 ns/op	      16 B/op	       2 allocs/op
BenchmarkIsWhiteSpace3-8   	  402494	      2942 ns/op	    3515 B/op	       5 allocs/op
BenchmarkIsWhiteSpace4-8   	 7176753	       164 ns/op	      16 B/op	       2 allocs/op
BenchmarkIsWhiteSpace5-8   	 7355563	       164 ns/op	      16 B/op	       2 allocs/op


3 is replaced with unicode.IsSpace(c)
BenchmarkIsWhiteSpace-8    	10207522	       118 ns/op	      16 B/op	       2 allocs/op
BenchmarkIsWhiteSpace2-8   	10175007	       123 ns/op	      16 B/op	       2 allocs/op
BenchmarkIsWhiteSpace3-8   	 8766572	       129 ns/op	      16 B/op	       2 allocs/op
BenchmarkIsWhiteSpace4-8   	 7361227	       166 ns/op	      16 B/op	       2 allocs/op
BenchmarkIsWhiteSpace5-8   	 6595645	       173 ns/op	      16 B/op	       2 allocs/op

BenchmarkIsWhiteSpace-8    	 9804511	       122 ns/op	      16 B/op	       2 allocs/op
BenchmarkIsWhiteSpace2-8   	 9558478	       128 ns/op	      16 B/op	       2 allocs/op
BenchmarkIsWhiteSpace3-8   	 8892658	       129 ns/op	      16 B/op	       2 allocs/op
BenchmarkIsWhiteSpace4-8   	 7139104	       165 ns/op	      16 B/op	       2 allocs/op
BenchmarkIsWhiteSpace5-8   	 7252981	       166 ns/op	      16 B/op	       2 allocs/op

================================================================
using bytes instead of runes for (1) and (2)
BenchmarkIsWhiteSpace-8    	13707632	        90.0 ns/op	       3 B/op	       1 allocs/op
BenchmarkIsWhiteSpace2-8   	12946368	        94.3 ns/op	       3 B/op	       1 allocs/op
BenchmarkIsWhiteSpace3-8   	 8932070	       130 ns/op	      16 B/op	       2 allocs/op
BenchmarkIsWhiteSpace4-8   	 7197372	       167 ns/op	      16 B/op	       2 allocs/op
BenchmarkIsWhiteSpace5-8   	 7093136	       166 ns/op	      16 B/op	       2 allocs/op

BenchmarkIsSpace-8               	 9312074	       126 ns/op	      16 B/op	       2 allocs/op
BenchmarkIsWhiteSpace-8          	13613052	        92.1 ns/op	       3 B/op	       1 allocs/op
BenchmarkIsWhiteSpace2-8         	12440322	        93.1 ns/op	       3 B/op	       1 allocs/op
BenchmarkIsWhiteSpace3-8         	 9552418	       127 ns/op	      16 B/op	       2 allocs/op
BenchmarkIsWhiteSpace4-8         	 7420647	       163 ns/op	      16 B/op	       2 allocs/op
BenchmarkIsWhiteSpace5-8         	 7419589	       160 ns/op	      16 B/op	       2 allocs/op
BenchmarkIsUnicodeWhiteSpace-8   	10632638	       118 ns/op	      16 B/op	       2 allocs/op
*/

// Final Benchmark Results
/*
================================================================
n = 3
================================================================
BenchmarkIsASCIISpace-8          	12261010	        92.3 ns/op	       3 B/op	       1 allocs/op
BenchmarkIsWhiteSpace-8          	12955863	        89.8 ns/op	       3 B/op	       1 allocs/op
BenchmarkIsWhiteSpace2-8         	11787762	        94.6 ns/op	       3 B/op	       1 allocs/op
BenchmarkIsWhiteSpace6-8         	10821630	       117 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsUnicodeWhiteSpace-8   	 8072452	       144 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsAnySpace-8            	 8138602	       145 ns/op	      16 B/op	       1 allocs/op
BenchmarkUnicode_IsSpace-8       	 7953177	       152 ns/op	      16 B/op	       1 allocs/op

================================================================
n = 1<<2
================================================================
BenchmarkIsASCIISpace-8          	10180090	       119 ns/op	       4 B/op	       1 allocs/op
BenchmarkIsWhiteSpace-8          	10503676	       120 ns/op	       4 B/op	       1 allocs/op
BenchmarkIsWhiteSpace2-8         	 9729558	       126 ns/op	       4 B/op	       1 allocs/op
BenchmarkIsWhiteSpace6-8         	 8457750	       143 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsUnicodeWhiteSpace-8   	 6555621	       191 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsAnySpace-8            	 6257503	       186 ns/op	      16 B/op	       1 allocs/op
BenchmarkUnicode_IsSpace-8       	 6285614	       195 ns/op	      16 B/op	       1 allocs/op

================================================================
n = 1<<3
================================================================
BenchmarkIsASCIISpace-8          	 5473182	       229 ns/op	       8 B/op	       1 allocs/op
BenchmarkIsWhiteSpace-8          	 5677147	       235 ns/op	       8 B/op	       1 allocs/op
BenchmarkIsWhiteSpace2-8         	 4445811	       257 ns/op	       8 B/op	       1 allocs/op
BenchmarkIsWhiteSpace6-8         	 4806199	       262 ns/op	      32 B/op	       1 allocs/op
BenchmarkIsUnicodeWhiteSpace-8   	 3098326	       374 ns/op	      32 B/op	       1 allocs/op
BenchmarkIsAnySpace-8            	 3237192	       365 ns/op	      32 B/op	       1 allocs/op
BenchmarkUnicode_IsSpace-8       	 3176928	       420 ns/op	      32 B/op	       1 allocs/op

================================================================
n = 1<<4
================================================================
BenchmarkIsASCIISpace-8          	 2850421	       426 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsWhiteSpace-8          	 2863578	       423 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsWhiteSpace2-8         	 2745991	       438 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsWhiteSpace6-8         	 2486973	       502 ns/op	      64 B/op	       1 allocs/op
BenchmarkIsUnicodeWhiteSpace-8   	 1717446	       707 ns/op	      64 B/op	       1 allocs/op
BenchmarkIsAnySpace-8            	 1688938	       700 ns/op	      64 B/op	       1 allocs/op
BenchmarkUnicode_IsSpace-8       	 1641483	       738 ns/op	      64 B/op	       1 allocs/op

================================================================
n = 1<<8
================================================================
BenchmarkIsASCIISpace-8          	  171495	      6260 ns/op	     256 B/op	       1 allocs/op
BenchmarkIsWhiteSpace-8          	  193882	      6198 ns/op	     256 B/op	       1 allocs/op
BenchmarkIsWhiteSpace2-8         	  187462	      6511 ns/op	     256 B/op	       1 allocs/op
BenchmarkIsWhiteSpace6-8         	  160911	      7757 ns/op	    1024 B/op	       1 allocs/op
BenchmarkIsUnicodeWhiteSpace-8   	  110793	     10816 ns/op	    1024 B/op	       1 allocs/op
BenchmarkIsAnySpace-8            	  111583	     10591 ns/op	    1024 B/op	       1 allocs/op
BenchmarkUnicode_IsSpace-8       	  109530	     11113 ns/op	    1024 B/op	       1 allocs/op

================================================================
n = 1<<12
================================================================
BenchmarkIsASCIISpace-8          	   12034	    100013 ns/op	    4096 B/op	       1 allocs/op
BenchmarkIsWhiteSpace-8          	   12196	    100702 ns/op	    4096 B/op	       1 allocs/op
BenchmarkIsWhiteSpace2-8         	   10000	    103272 ns/op	    4096 B/op	       1 allocs/op
BenchmarkIsWhiteSpace6-8         	   10000	    114599 ns/op	   16384 B/op	       1 allocs/op
BenchmarkIsUnicodeWhiteSpace-8   	    6301	    193826 ns/op	   16384 B/op	       1 allocs/op
BenchmarkIsAnySpace-8            	    6343	    192879 ns/op	   16384 B/op	       1 allocs/op
BenchmarkUnicode_IsSpace-8       	    6753	    190504 ns/op	   16384 B/op	       1 allocs/op

================================================================
n = 1<<16 - 1
================================================================
BenchmarkIsASCIISpace-8          	     614	   1665135 ns/op	   65539 B/op	       1 allocs/op
BenchmarkIsWhiteSpace-8          	     754	   1581501 ns/op	   65536 B/op	       1 allocs/op
BenchmarkIsWhiteSpace2-8         	     726	   1659465 ns/op	   65536 B/op	       1 allocs/op
BenchmarkIsWhiteSpace6-8         	     654	   1885503 ns/op	  262145 B/op	       1 allocs/op
BenchmarkIsUnicodeWhiteSpace-8   	     386	   3095246 ns/op	  262145 B/op	       1 allocs/op
BenchmarkIsAnySpace-8            	     382	   3108825 ns/op	  262144 B/op	       1 allocs/op
BenchmarkUnicode_IsSpace-8       	     416	   2908501 ns/op	  262145 B/op	       1 allocs/op

================================================================
n = 1<<24 - 1
================================================================
BenchmarkIsASCIISpace-8          	       3	 394559021 ns/op	16777312 B/op	       2 allocs/op
BenchmarkIsWhiteSpace-8          	       3	 394949234 ns/op	16777248 B/op	       1 allocs/op
BenchmarkIsWhiteSpace2-8         	       3	 419641493 ns/op	16777218 B/op	       1 allocs/op
BenchmarkIsWhiteSpace6-8         	       3	 454344906 ns/op	67108864 B/op	       1 allocs/op
BenchmarkIsUnicodeWhiteSpace-8   	       2	 675746892 ns/op	67108912 B/op	       1 allocs/op
BenchmarkIsAnySpace-8            	       2	 674386516 ns/op	67108864 B/op	       1 allocs/op
BenchmarkUnicode_IsSpace-8       	       2	 714679118 ns/op	67108864 B/op	       1 allocs/op
*/


var Want = unicode.IsSpace


func init() {
    rand.Seed(time.Now().UnixNano())
}


func SmallRuneSamples() []rune {
	return []rune{
		'A', '0', 65, 't', 'n', 'f', 'r', 'v', '\t', '\n', '\f', '\r', '\v', 48, 12, ' ', 0x20, 8, 0x1680, 0x2028, 0x3000, 0x1680, 0x200C, 0x2123, 0x3333, 0xFFDF, 0xFFEE,
	}
}

func SmallByteSamples() []byte {
	return []byte{
		'A', '0', 65, 't', 'n', 'f', 'r', 'v', '\t', '\n', '\f', '\r', '\v', 48, 12, ' ', 0x20, 8, 0xFF,
	}
}

func SmallByteStringSamples() (list []string) {
    for _, c := range SmallByteSamples() {
        list = append(list, string(c))
    }
    return
}

func SmallRuneStringSamples() (list []string) {
    for _, r := range SmallRuneSamples() {
        list = append(list, string(r))
    }
    return
}

func ByteSamples() []byte {
    n := numSamples
	if n < 2 || n > maxSamples {
		n = defaultSamples
	}
	retval := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		retval = append(retval, byte(rand.Intn(126)))
    }
    retval = append(retval, 0xFF)
    return retval
}

func RuneSamples() []rune {
    n := numSamples
	if n < 2 || n > maxSamples {
		n = defaultSamples
	}
	retval := make([]rune, 0, n)
	for i := 0; i < n; i++ {
        retval = append(retval, rune(rand.Intn(0x3000)))
	}
	return retval
}

func byteStringSamples() (list []string) {
    for _, r := range RuneSamples() {
        list = append(list, string(r))
    }
    return
}

func runeStringSamples() (list []string) {
    for _, r := range RuneSamples() {
        list = append(list, string(r))
    }
    return
}

// This is a horrible idea ... much slower
// var (
// 	byteSamples  = ByteSamples
// 	runeSamples  = RuneSamples
// )

// func BenchmarkIsAlpha(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		for _, c := range ByteSamples() {
// 			IsAlpha(c)
// 		}
// 	}
// }

// func BenchmarkIsDigit(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		for _, c := range ByteSamples() {
// 			IsDigit(c)
// 		}
// 	}
// }

// func BenchmarkIsAlphaSwitch(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		for _, c := range ByteSamples() {
// 			IsAlphaSwitch(c)
// 		}
// 	}
// }

// func BenchmarkIsAlphaNum(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		for _, c := range ByteSamples() {
// 			IsAlphaNum(c)
// 		}
// 	}
// }

// func BenchmarkIsAlphaNum2(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		for _, c := range ByteSamples() {
// 			IsAlphaNum2(c)
// 		}
// 	}
// }


// func BenchmarkIsWhiteSpace5(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		for _, r := range RuneSamples() {
// 			isWhiteSpace5(r)
// 		}
// 	}
// }

// func BenchmarkIsWhiteSpace4(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		for _, r := range RuneSamples() {
// 			isWhiteSpace4(r)
// 		}
// 	}
// }
