// Code generated by go2go; DO NOT EDIT.


//line sort.go2:7
package sort2

//line sort.go2:7
import (
//line sort.go2:7
 "fmt"
//line sort.go2:7
 "math"
//line sort.go2:7
 "math/rand"
//line sort.go2:7
 "runtime"
//line sort.go2:7
 "strconv"
//line sort.go2:7
 "testing"
//line sort.go2:7
)

//line sort.go2:12
type Interface interface {
//line sort.go2:14
 Len() int

//line sort.go2:17
 Less(i, j int) bool

	Swap(i, j int)
}

type lessSwap interface {
	Less(i, j int) bool
	Swap(i, j int)
}

//line sort.go2:264
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

//line sort.go2:271
func (p IntSlice) Sort() { instantiate୦୦Sort୦sort2୮aIntSlice(p) }

//line sort.go2:275
type Float64Slice []float64

func (p Float64Slice) Len() int           { return len(p) }
func (p Float64Slice) Less(i, j int) bool { return p[i] < p[j] || isNaN(p[i]) && !isNaN(p[j]) }
func (p Float64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

//line sort.go2:282
func isNaN(f float64) bool {
	return f != f
}

//line sort.go2:287
func (p Float64Slice) Sort() { instantiate୦୦Sort୦sort2୮aFloat64Slice(p) }

//line sort.go2:290
type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

//line sort.go2:297
func (p StringSlice) Sort() { instantiate୦୦Sort୦sort2୮aStringSlice(p) }

//line sort.go2:302
func Ints(a []int) { instantiate୦୦Sort୦sort2୮aIntSlice(IntSlice(a)) }

//line sort.go2:306
func Float64s(a []float64) { instantiate୦୦Sort୦sort2୮aFloat64Slice(Float64Slice(a)) }

//line sort.go2:309
func Strings(a []string) { instantiate୦୦Sort୦sort2୮aStringSlice(StringSlice(a)) }

//line sort.go2:312
func IntsAreSorted(a []int) bool { return instantiate୦୦IsSorted୦sort2୮aIntSlice(IntSlice(a)) }

//line sort.go2:316
func Float64sAreSorted(a []float64) bool {
//line sort.go2:316
 return instantiate୦୦IsSorted୦sort2୮aFloat64Slice(Float64Slice(a))
//line sort.go2:316
}

//line sort.go2:319
func StringsAreSorted(a []string) bool {
//line sort.go2:319
 return instantiate୦୦IsSorted୦sort2୮aStringSlice(StringSlice(a))
//line sort.go2:319
}
//line sort.go2:219
func instantiate୦୦Sort୦sort2୮aIntSlice(data (IntSlice),) {
			n := data.Len()
//line sort.go2:220
 instantiate୦୦quickSort୦sort2୮aIntSlice(data, 0, n, instantiate୦୦maxDepth୦sort2୮aIntSlice(n))
//line sort.go2:222
}
//line sort.go2:219
func instantiate୦୦Sort୦sort2୮aFloat64Slice(data (Float64Slice),) {
			n := data.Len()
//line sort.go2:220
 instantiate୦୦quickSort୦sort2୮aFloat64Slice(data, 0, n, instantiate୦୦maxDepth୦sort2୮aFloat64Slice(n))
//line sort.go2:222
}
//line sort.go2:219
func instantiate୦୦Sort୦sort2୮aStringSlice(data (StringSlice),) {
			n := data.Len()
//line sort.go2:220
 instantiate୦୦quickSort୦sort2୮aStringSlice(data, 0, n, instantiate୦୦maxDepth୦sort2୮aStringSlice(n))
//line sort.go2:222
}

//line sort.go2:251
func instantiate୦୦IsSorted୦sort2୮aIntSlice(data (IntSlice),) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}
//line sort.go2:251
func instantiate୦୦IsSorted୦sort2୮aFloat64Slice(data (Float64Slice),) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}
//line sort.go2:251
func instantiate୦୦IsSorted୦sort2୮aStringSlice(data (StringSlice),) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}
//line sort.go2:186
func instantiate୦୦quickSort୦sort2୮aIntSlice(data (IntSlice), a, b, maxDepth int) {
	for b-a > 12 {
		if maxDepth == 0 {
//line sort.go2:188
   instantiate୦୦heapSort୦sort2୮aIntSlice(data, a, b)
//line sort.go2:190
   return
		}
				maxDepth--
				mlo, mhi := instantiate୦୦doPivot୦sort2୮aIntSlice(data, a, b)

//line sort.go2:196
  if mlo-a < b-mhi {
//line sort.go2:196
   instantiate୦୦quickSort୦sort2୮aIntSlice(data, a, mlo, maxDepth)
//line sort.go2:198
   a = mhi
		} else {
//line sort.go2:199
   instantiate୦୦quickSort୦sort2୮aIntSlice(data, mhi, b, maxDepth)
//line sort.go2:201
   b = mlo
		}
	}
	if b-a > 1 {

//line sort.go2:207
  for i := a + 6; i < b; i++ {
			if data.Less(i, i-6) {
				data.Swap(i, i-6)
			}
		}
//line sort.go2:211
  instantiate୦୦insertionSort୦sort2୮aIntSlice(data, a, b)
//line sort.go2:213
 }
}

//line sort.go2:226
func instantiate୦୦maxDepth୦sort2୮aIntSlice(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}
//line sort.go2:186
func instantiate୦୦quickSort୦sort2୮aFloat64Slice(data (Float64Slice), a, b, maxDepth int) {
	for b-a > 12 {
		if maxDepth == 0 {
//line sort.go2:188
   instantiate୦୦heapSort୦sort2୮aFloat64Slice(data, a, b)
//line sort.go2:190
   return
		}
				maxDepth--
				mlo, mhi := instantiate୦୦doPivot୦sort2୮aFloat64Slice(data, a, b)

//line sort.go2:196
  if mlo-a < b-mhi {
//line sort.go2:196
   instantiate୦୦quickSort୦sort2୮aFloat64Slice(data, a, mlo, maxDepth)
//line sort.go2:198
   a = mhi
		} else {
//line sort.go2:199
   instantiate୦୦quickSort୦sort2୮aFloat64Slice(data, mhi, b, maxDepth)
//line sort.go2:201
   b = mlo
		}
	}
	if b-a > 1 {

//line sort.go2:207
  for i := a + 6; i < b; i++ {
			if data.Less(i, i-6) {
				data.Swap(i, i-6)
			}
		}
//line sort.go2:211
  instantiate୦୦insertionSort୦sort2୮aFloat64Slice(data, a, b)
//line sort.go2:213
 }
}

//line sort.go2:226
func instantiate୦୦maxDepth୦sort2୮aFloat64Slice(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}
//line sort.go2:186
func instantiate୦୦quickSort୦sort2୮aStringSlice(data (StringSlice), a, b, maxDepth int) {
	for b-a > 12 {
		if maxDepth == 0 {
//line sort.go2:188
   instantiate୦୦heapSort୦sort2୮aStringSlice(data, a, b)
//line sort.go2:190
   return
		}
				maxDepth--
				mlo, mhi := instantiate୦୦doPivot୦sort2୮aStringSlice(data, a, b)

//line sort.go2:196
  if mlo-a < b-mhi {
//line sort.go2:196
   instantiate୦୦quickSort୦sort2୮aStringSlice(data, a, mlo, maxDepth)
//line sort.go2:198
   a = mhi
		} else {
//line sort.go2:199
   instantiate୦୦quickSort୦sort2୮aStringSlice(data, mhi, b, maxDepth)
//line sort.go2:201
   b = mlo
		}
	}
	if b-a > 1 {

//line sort.go2:207
  for i := a + 6; i < b; i++ {
			if data.Less(i, i-6) {
				data.Swap(i, i-6)
			}
		}
//line sort.go2:211
  instantiate୦୦insertionSort୦sort2୮aStringSlice(data, a, b)
//line sort.go2:213
 }
}

//line sort.go2:226
func instantiate୦୦maxDepth୦sort2୮aStringSlice(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}
//line sort.go2:56
func instantiate୦୦heapSort୦sort2୮aIntSlice(data (IntSlice), a, b int) {
			first := a
			lo := 0
			hi := b - a

//line sort.go2:62
 for i := (hi - 1) / 2; i >= 0; i-- {
//line sort.go2:62
  instantiate୦୦siftDown୦sort2୮aIntSlice(data, i, hi, first)
//line sort.go2:64
 }

//line sort.go2:67
 for i := hi - 1; i >= 0; i-- {
				data.Swap(first, first+i)
//line sort.go2:68
  instantiate୦୦siftDown୦sort2୮aIntSlice(data, lo, i, first)
//line sort.go2:70
 }
}

//line sort.go2:99
func instantiate୦୦doPivot୦sort2୮aIntSlice(data (IntSlice), lo, hi int) (midlo, midhi int) {
	m := int(uint(lo+hi) >> 1)
	if hi-lo > 40 {

				s := (hi - lo) / 8
//line sort.go2:103
  instantiate୦୦medianOfThree୦sort2୮aIntSlice(data, lo, lo+s, lo+2*s)
				instantiate୦୦medianOfThree୦sort2୮aIntSlice(data, m, m-s, m+s)
				instantiate୦୦medianOfThree୦sort2୮aIntSlice(data, hi-1, hi-1-s, hi-1-2*s)
//line sort.go2:107
 }
//line sort.go2:107
 instantiate୦୦medianOfThree୦sort2୮aIntSlice(data, lo, m, hi-1)

//line sort.go2:117
 pivot := lo
	a, c := lo+1, hi-1

	for ; a < c && data.Less(a, pivot); a++ {
	}
	b := a
	for {
		for ; b < c && !data.Less(pivot, b); b++ {
		}
		for ; b < c && data.Less(pivot, c-1); c-- {
		}
		if b >= c {
			break
		}

		data.Swap(b, c-1)
		b++
		c--
	}

//line sort.go2:138
 protect := hi-c < 5
	if !protect && hi-c < (hi-lo)/4 {

		dups := 0
		if !data.Less(pivot, hi-1) {
			data.Swap(c, hi-1)
			c++
			dups++
		}
		if !data.Less(b-1, pivot) {
			b--
			dups++
		}

//line sort.go2:154
  if !data.Less(m, pivot) {
			data.Swap(m, b-1)
			b--
			dups++
		}

		protect = dups > 1
	}
	if protect {

//line sort.go2:167
  for {
			for ; a < b && !data.Less(b-1, pivot); b-- {
			}
			for ; a < b && data.Less(a, pivot); a++ {
			}
			if a >= b {
				break
			}

			data.Swap(a, b-1)
			a++
			b--
		}
	}

	data.Swap(pivot, b-1)
	return b - 1, c
}
//line sort.go2:28
func instantiate୦୦insertionSort୦sort2୮aIntSlice(data (IntSlice), a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

//line sort.go2:56
func instantiate୦୦heapSort୦sort2୮aFloat64Slice(data (Float64Slice), a, b int) {
			first := a
			lo := 0
			hi := b - a

//line sort.go2:62
 for i := (hi - 1) / 2; i >= 0; i-- {
//line sort.go2:62
  instantiate୦୦siftDown୦sort2୮aFloat64Slice(data, i, hi, first)
//line sort.go2:64
 }

//line sort.go2:67
 for i := hi - 1; i >= 0; i-- {
				data.Swap(first, first+i)
//line sort.go2:68
  instantiate୦୦siftDown୦sort2୮aFloat64Slice(data, lo, i, first)
//line sort.go2:70
 }
}

//line sort.go2:99
func instantiate୦୦doPivot୦sort2୮aFloat64Slice(data (Float64Slice), lo, hi int) (midlo, midhi int) {
	m := int(uint(lo+hi) >> 1)
	if hi-lo > 40 {

				s := (hi - lo) / 8
//line sort.go2:103
  instantiate୦୦medianOfThree୦sort2୮aFloat64Slice(data, lo, lo+s, lo+2*s)
				instantiate୦୦medianOfThree୦sort2୮aFloat64Slice(data, m, m-s, m+s)
				instantiate୦୦medianOfThree୦sort2୮aFloat64Slice(data, hi-1, hi-1-s, hi-1-2*s)
//line sort.go2:107
 }
//line sort.go2:107
 instantiate୦୦medianOfThree୦sort2୮aFloat64Slice(data, lo, m, hi-1)

//line sort.go2:117
 pivot := lo
	a, c := lo+1, hi-1

	for ; a < c && data.Less(a, pivot); a++ {
	}
	b := a
	for {
		for ; b < c && !data.Less(pivot, b); b++ {
		}
		for ; b < c && data.Less(pivot, c-1); c-- {
		}
		if b >= c {
			break
		}

		data.Swap(b, c-1)
		b++
		c--
	}

//line sort.go2:138
 protect := hi-c < 5
	if !protect && hi-c < (hi-lo)/4 {

		dups := 0
		if !data.Less(pivot, hi-1) {
			data.Swap(c, hi-1)
			c++
			dups++
		}
		if !data.Less(b-1, pivot) {
			b--
			dups++
		}

//line sort.go2:154
  if !data.Less(m, pivot) {
			data.Swap(m, b-1)
			b--
			dups++
		}

		protect = dups > 1
	}
	if protect {

//line sort.go2:167
  for {
			for ; a < b && !data.Less(b-1, pivot); b-- {
			}
			for ; a < b && data.Less(a, pivot); a++ {
			}
			if a >= b {
				break
			}

			data.Swap(a, b-1)
			a++
			b--
		}
	}

	data.Swap(pivot, b-1)
	return b - 1, c
}
//line sort.go2:28
func instantiate୦୦insertionSort୦sort2୮aFloat64Slice(data (Float64Slice), a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

//line sort.go2:56
func instantiate୦୦heapSort୦sort2୮aStringSlice(data (StringSlice), a, b int) {
			first := a
			lo := 0
			hi := b - a

//line sort.go2:62
 for i := (hi - 1) / 2; i >= 0; i-- {
//line sort.go2:62
  instantiate୦୦siftDown୦sort2୮aStringSlice(data, i, hi, first)
//line sort.go2:64
 }

//line sort.go2:67
 for i := hi - 1; i >= 0; i-- {
				data.Swap(first, first+i)
//line sort.go2:68
  instantiate୦୦siftDown୦sort2୮aStringSlice(data, lo, i, first)
//line sort.go2:70
 }
}

//line sort.go2:99
func instantiate୦୦doPivot୦sort2୮aStringSlice(data (StringSlice), lo, hi int) (midlo, midhi int) {
	m := int(uint(lo+hi) >> 1)
	if hi-lo > 40 {

				s := (hi - lo) / 8
//line sort.go2:103
  instantiate୦୦medianOfThree୦sort2୮aStringSlice(data, lo, lo+s, lo+2*s)
				instantiate୦୦medianOfThree୦sort2୮aStringSlice(data, m, m-s, m+s)
				instantiate୦୦medianOfThree୦sort2୮aStringSlice(data, hi-1, hi-1-s, hi-1-2*s)
//line sort.go2:107
 }
//line sort.go2:107
 instantiate୦୦medianOfThree୦sort2୮aStringSlice(data, lo, m, hi-1)

//line sort.go2:117
 pivot := lo
	a, c := lo+1, hi-1

	for ; a < c && data.Less(a, pivot); a++ {
	}
	b := a
	for {
		for ; b < c && !data.Less(pivot, b); b++ {
		}
		for ; b < c && data.Less(pivot, c-1); c-- {
		}
		if b >= c {
			break
		}

		data.Swap(b, c-1)
		b++
		c--
	}

//line sort.go2:138
 protect := hi-c < 5
	if !protect && hi-c < (hi-lo)/4 {

		dups := 0
		if !data.Less(pivot, hi-1) {
			data.Swap(c, hi-1)
			c++
			dups++
		}
		if !data.Less(b-1, pivot) {
			b--
			dups++
		}

//line sort.go2:154
  if !data.Less(m, pivot) {
			data.Swap(m, b-1)
			b--
			dups++
		}

		protect = dups > 1
	}
	if protect {

//line sort.go2:167
  for {
			for ; a < b && !data.Less(b-1, pivot); b-- {
			}
			for ; a < b && data.Less(a, pivot); a++ {
			}
			if a >= b {
				break
			}

			data.Swap(a, b-1)
			a++
			b--
		}
	}

	data.Swap(pivot, b-1)
	return b - 1, c
}
//line sort.go2:28
func instantiate୦୦insertionSort୦sort2୮aStringSlice(data (StringSlice), a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

//line sort.go2:38
func instantiate୦୦siftDown୦sort2୮aIntSlice(data (IntSlice), lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && data.Less(first+child, first+child+1) {
			child++
		}
		if !data.Less(first+root, first+child) {
			return
		}
		data.Swap(first+root, first+child)
		root = child
	}
}

//line sort.go2:77
func instantiate୦୦medianOfThree୦sort2୮aIntSlice(data (IntSlice), m1, m0, m2 int) {

	if data.Less(m1, m0) {
		data.Swap(m1, m0)
	}

	if data.Less(m2, m1) {
		data.Swap(m2, m1)

		if data.Less(m1, m0) {
			data.Swap(m1, m0)
		}
	}

}
//line sort.go2:38
func instantiate୦୦siftDown୦sort2୮aFloat64Slice(data (Float64Slice), lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && data.Less(first+child, first+child+1) {
			child++
		}
		if !data.Less(first+root, first+child) {
			return
		}
		data.Swap(first+root, first+child)
		root = child
	}
}

//line sort.go2:77
func instantiate୦୦medianOfThree୦sort2୮aFloat64Slice(data (Float64Slice), m1, m0, m2 int) {

	if data.Less(m1, m0) {
		data.Swap(m1, m0)
	}

	if data.Less(m2, m1) {
		data.Swap(m2, m1)

		if data.Less(m1, m0) {
			data.Swap(m1, m0)
		}
	}

}
//line sort.go2:38
func instantiate୦୦siftDown୦sort2୮aStringSlice(data (StringSlice), lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && data.Less(first+child, first+child+1) {
			child++
		}
		if !data.Less(first+root, first+child) {
			return
		}
		data.Swap(first+root, first+child)
		root = child
	}
}

//line sort.go2:77
func instantiate୦୦medianOfThree୦sort2୮aStringSlice(data (StringSlice), m1, m0, m2 int) {

	if data.Less(m1, m0) {
		data.Swap(m1, m0)
	}

	if data.Less(m2, m1) {
		data.Swap(m2, m1)

		if data.Less(m1, m0) {
			data.Swap(m1, m0)
		}
	}

}

//line sort.go2:91
var _ = fmt.Errorf
//line sort.go2:91
var _ = math.Abs
//line sort.go2:91
var _ = rand.ExpFloat64
//line sort.go2:91
var _ = runtime.BlockProfile
//line sort.go2:91
var _ = strconv.AppendBool
//line sort.go2:91
var _ = testing.AllocsPerRun
