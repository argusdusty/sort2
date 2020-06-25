# sort2
A performant go2go drop-in replacement for "sort", via the [cmd/go2go](https://go.googlesource.com/go/+/refs/heads/dev.go2go/README.go2go.md) generics experiment.

This is intended as an experimental demonstration of the more performant and memory-efficient sorting that will come with Go 2 generics. Go2go is not stable and should not be used in production. However, this is good for getting a feel for how generic sorting may work in Go 2.

This follows [the June 16, 2020 generics design draft](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md).

### How?

When using go2go, change `import "sort"` to `import "github.com/argusdusty/sort2"` and `sort.` calls to `sort2.` calls. Compile with cmd/go2go as usual. That's it.

The one exception is when you're using Sort/etc. as a variable of type func(Interface), you'll need to instantiate it specifically to the desired type that implements Interface.

### Benchmarks

```
C:\Users\argusdusty\Code\src\github.com\argusdusty\sort2>go test -run=None -bench=. -benchmem -cpu=1 sort > old.txt
C:\Users\argusdusty\Code\src\github.com\argusdusty\sort2>go test -run=None -bench=. -benchmem -cpu=1 > new.txt
C:\Users\argusdusty\Code\src\github.com\argusdusty\sort2>benchcmp old.txt new.txt
benchmark                       old ns/op      new ns/op      delta
BenchmarkSearchWrappers         113            95.6           -15.40%
BenchmarkSortString1K           129093         77974          -39.60%
BenchmarkSortString1K_Slice     118827         100123         -15.74%
BenchmarkStableString1K         159700         86238          -46.00%
BenchmarkSortInt1K              64380          11499          -82.14%
BenchmarkStableInt1K            77514          19126          -75.33%
BenchmarkStableInt1K_Slice      66054          37648          -43.00%
BenchmarkSortInt64K             6887104        1985236        -71.17%
BenchmarkSortInt64K_Slice       6021752        5013531        -16.74%
BenchmarkStableInt64K           7655556        1537283        -79.92%
BenchmarkSort1e2                39004          17308          -55.63%
BenchmarkStable1e2              78648          37286          -52.59%
BenchmarkSort1e4                8603409        3472625        -59.64%
BenchmarkStable1e4              26374358       12730205       -51.73%
BenchmarkSort1e6                1309289400     515116800      -60.66%
BenchmarkStable1e6              5547247900     2591584200     -53.28%

benchmark                       old allocs     new allocs     delta
BenchmarkSearchWrappers         0              0              +0.00%
BenchmarkSortString1K           1              0              -100.00%
BenchmarkSortString1K_Slice     2              0              -100.00%
BenchmarkStableString1K         1              0              -100.00%
BenchmarkSortInt1K              1              0              -100.00%
BenchmarkStableInt1K            1              0              -100.00%
BenchmarkStableInt1K_Slice      2              0              -100.00%
BenchmarkSortInt64K             1              0              -100.00%
BenchmarkSortInt64K_Slice       2              0              -100.00%
BenchmarkStableInt64K           1              0              -100.00%
BenchmarkSort1e2                7              0              -100.00%
BenchmarkStable1e2              7              0              -100.00%
BenchmarkSort1e4                7              0              -100.00%
BenchmarkStable1e4              7              0              -100.00%
BenchmarkSort1e6                10             0              -100.00%
BenchmarkStable1e6              10             0              -100.00%

benchmark                       old bytes     new bytes     delta
BenchmarkSearchWrappers         0             0             +0.00%
BenchmarkSortString1K           32            0             -100.00%
BenchmarkSortString1K_Slice     64            0             -100.00%
BenchmarkStableString1K         32            0             -100.00%
BenchmarkSortInt1K              32            0             -100.00%
BenchmarkStableInt1K            32            0             -100.00%
BenchmarkStableInt1K_Slice      64            0             -100.00%
BenchmarkSortInt64K             32            0             -100.00%
BenchmarkSortInt64K_Slice       64            0             -100.00%
BenchmarkStableInt64K           32            0             -100.00%
BenchmarkSort1e2                224           0             -100.00%
BenchmarkStable1e2              224           0             -100.00%
BenchmarkSort1e4                225           0             -100.00%
BenchmarkStable1e4              230           0             -100.00%
BenchmarkSort1e6                480           0             -100.00%
BenchmarkStable1e6              480           0             -100.00%
```

sort2 makes 0 additional allocations (whereas the old sort needs to convert between interface{} and the raw type, each time requiring 1 word allocation), and usually runs **2-5x** faster.

There is the exception of Slice (only 15% faster), which doesn't lend itself well to generics, however I've added a new Sortable interface that applies to the elements of the slice, and corresponding SortableSlice functions, which result in similar speed-ups to the rest of the package:
```
BenchmarkSortString1K_SortableSlice        14902             82155 ns/op               0 B/op          0 allocs/op
BenchmarkStableInt1K_SortableSlice         61521             20437 ns/op               0 B/op          0 allocs/op
BenchmarkSortInt64K_SortableSlice            592           1980180 ns/op               0 B/op          0 allocs/op
```