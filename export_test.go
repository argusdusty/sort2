// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sort2

func Heapsort[T Interface](data T) {
	heapSort(data, 0, data.Len())
}
