// Code generated by go2go; DO NOT EDIT.


//line search.go2:7
package sort2

//line search.go2:7
import (
//line search.go2:7
 "fmt"
//line search.go2:7
 "math"
//line search.go2:7
 "math/rand"
//line search.go2:7
 "runtime"
//line search.go2:7
 "strconv"
//line search.go2:7
 "testing"
//line search.go2:7
)

//line search.go2:59
func Search(n int, f func(int) bool) int {

//line search.go2:62
 i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1)

		if !f(h) {
			i = h + 1
		} else {
			j = h
		}
	}

	return i
}

//line search.go2:83
func SearchInts(a []int, x int) int {
	return Search(len(a), func(i int) bool { return a[i] >= x })
}

//line search.go2:92
func SearchFloat64s(a []float64, x float64) int {
	return Search(len(a), func(i int) bool { return a[i] >= x })
}

//line search.go2:101
func SearchStrings(a []string, x string) int {
	return Search(len(a), func(i int) bool { return a[i] >= x })
}

//line search.go2:106
func (p IntSlice) Search(x int) int { return SearchInts(p, x) }

//line search.go2:109
func (p Float64Slice) Search(x float64) int { return SearchFloat64s(p, x) }

//line search.go2:112
func (p StringSlice) Search(x string) int { return SearchStrings(p, x) }

//line search.go2:112
var _ = fmt.Errorf
//line search.go2:112
var _ = math.Abs
//line search.go2:112
var _ = rand.ExpFloat64
//line search.go2:112
var _ = runtime.BlockProfile
//line search.go2:112
var _ = strconv.AppendBool
//line search.go2:112
var _ = testing.AllocsPerRun
