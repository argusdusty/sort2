# sort2
A mostly drop-in replacement for "sort" that uses the new [generics](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md), via the Go 1.18 beta version 1.

This was intended as an experimental demonstration of the more performant and memory-efficient sorting that will come with generics. Unfortunately, these performance benefits are not achieved with Go 1.18 generics beyond union types. Instead, in Go 1.18, generic functions over a generic type that has a function constraint will convert to an interface object before every such function call, leading to significant memory allocations and massive performance losses.

### How?

When using Go 1.18, change `import "sort"` to `import "github.com/argusdusty/sort2"` and `sort.` calls to `sort2.` calls. That's basically it. You may need to instantiate some sort functions to your generic type to use them.

### Benchmarks

```
C:\Users\argusdusty\Code\src\github.com\argusdusty\sort2>go version
go version go1.18beta1 windows/amd64
C:\Users\argusdusty\Code\src\github.com\argusdusty\sort2>go test -run=None -bench=. -benchmem -cpu=1 -count=10 sort > old.txt
C:\Users\argusdusty\Code\src\github.com\argusdusty\sort2>go test -run=None -bench=. -benchmem -cpu=1 -count=5 > new.txt
... Remove the headers from the new.txt and old.txt files
C:\Users\argusdusty\Code\src\github.com\argusdusty\sort2>benchstat old.txt new.txt
name                old time/op    new time/op        delta
SearchWrappers        85.0ns ± 2%       961.4ns ± 2%        +1031.41%  (p=0.001 n=10+5)
SortString1K          86.7µs ± 1%       493.4µs ± 1%         +468.95%  (p=0.001 n=9+5)
SortString1K_Slice    87.8µs ± 1%       721.6µs ± 6%         +721.67%  (p=0.001 n=10+5)
StableString1K         111µs ± 1%         833µs ± 1%         +648.98%  (p=0.001 n=10+5)
SortInt1K             39.6µs ± 1%       440.7µs ± 4%        +1012.55%  (p=0.001 n=10+5)
StableInt1K           52.1µs ± 1%       648.7µs ± 4%        +1144.15%  (p=0.001 n=10+5)
StableInt1K_Slice     51.9µs ± 0%       861.7µs ± 4%        +1558.91%  (p=0.002 n=8+5)
SortInt64K            4.35ms ± 0%       43.62ms ± 6%         +901.71%  (p=0.001 n=10+5)
SortInt64K_Slice      4.75ms ± 1%       64.80ms ± 3%        +1264.44%  (p=0.001 n=10+5)
StableInt64K          4.52ms ± 1%       54.33ms ± 4%        +1101.30%  (p=0.001 n=10+5)
Sort1e2               27.7µs ± 2%       178.2µs ± 4%         +543.47%  (p=0.001 n=10+5)
Stable1e2             52.8µs ± 3%       403.8µs ± 3%         +664.55%  (p=0.001 n=10+5)
Sort1e4               5.68ms ± 1%       38.43ms ± 2%         +576.39%  (p=0.001 n=10+5)
Stable1e4             16.8ms ± 2%       131.1ms ± 5%         +679.98%  (p=0.001 n=10+5)
Sort1e6                871ms ± 0%        6104ms ± 2%         +600.79%  (p=0.001 n=10+5)
Stable1e6              3.44s ± 0%        28.75s ± 4%         +736.31%  (p=0.001 n=10+5)

name                old alloc/op   new alloc/op       delta
SearchWrappers         0.00B            688.00B ± 0%            +Inf%  (p=0.001 n=10+5)
SortString1K           24.0B ± 0%     320232.0B ± 0%     +1334200.00%  (p=0.001 n=10+5)
SortString1K_Slice     56.0B ± 0%     427008.0B ± 0%      +762414.29%  (p=0.001 n=10+5)
StableString1K         24.0B ± 0%     551664.0B ± 0%     +2298500.00%  (p=0.001 n=10+5)
SortInt1K              24.0B ± 0%     319512.0B ± 0%     +1331200.00%  (p=0.001 n=10+5)
StableInt1K            24.0B ± 0%     445896.0B ± 0%     +1857800.00%  (p=0.001 n=10+5)
StableInt1K_Slice      56.0B ± 0%     594528.0B ± 0%     +1061557.14%  (p=0.001 n=10+5)
SortInt64K             24.0B ± 0%   32293440.0B ± 0%   +134555900.00%  (p=0.001 n=10+5)
SortInt64K_Slice       56.0B ± 0%   43057952.0B ± 0%    +76889100.00%  (p=0.001 n=10+5)
StableInt64K           24.0B ± 0%   40169208.0B ± 0%   +167371600.00%  (p=0.001 n=10+5)
Sort1e2                 168B ± 0%       121421B ± 0%       +72174.40%  (p=0.001 n=10+5)
Stable1e2               168B ± 0%       284151B ± 0%      +169037.26%  (p=0.001 n=10+5)
Sort1e4                 169B ± 0%     27099900B ± 0%    +16035343.91%  (p=0.001 n=10+5)
Stable1e4               171B ± 0%     86421198B ± 0%    +50538612.28%  (p=0.001 n=10+5)
Sort1e6                 288B ± 0%   4190006232B ± 0%  +1454863175.00%  (p=0.001 n=10+5)
Stable1e6               408B ± 0%  18409665024B ± 0%  +4512172700.00%  (p=0.001 n=10+5)

name                old allocs/op  new allocs/op      delta
SearchWrappers          0.00              19.00 ± 0%            +Inf%  (p=0.001 n=10+5)
SortString1K            1.00 ± 0%      13343.00 ± 0%     +1334200.00%  (p=0.001 n=10+5)
SortString1K_Slice      2.00 ± 0%      13344.00 ± 0%      +667100.00%  (p=0.001 n=10+5)
StableString1K          1.00 ± 0%      22986.00 ± 0%     +2298500.00%  (p=0.001 n=10+5)
SortInt1K               1.00 ± 0%      13313.00 ± 0%     +1331200.00%  (p=0.001 n=10+5)
StableInt1K             1.00 ± 0%      18579.00 ± 0%     +1857800.00%  (p=0.001 n=10+5)
StableInt1K_Slice       2.00 ± 0%      18579.00 ± 0%      +928850.00%  (p=0.001 n=10+5)
SortInt64K              1.00 ± 0%    1345560.00 ± 0%   +134555900.00%  (p=0.001 n=10+5)
SortInt64K_Slice        2.00 ± 0%    1345561.00 ± 0%    +67277950.00%  (p=0.001 n=10+5)
StableInt64K            1.00 ± 0%    1673717.00 ± 0%   +167371600.00%  (p=0.001 n=10+5)
Sort1e2                 7.00 ± 0%       5059.00 ± 0%       +72171.43%  (p=0.002 n=10+4)
Stable1e2               7.00 ± 0%      11839.40 ± 0%      +169034.29%  (p=0.001 n=10+5)
Sort1e4                 7.00 ± 0%    1129161.60 ± 0%    +16130780.00%  (p=0.001 n=10+5)
Stable1e4               7.00 ± 0%    3600883.00 ± 0%    +51441085.71%  (p=0.001 n=10+5)
Sort1e6                 8.00 ± 0%  174583593.00 ± 0%  +2182294812.50%  (p=0.001 n=10+5)
Stable1e6               9.00 ± 0%  767069376.00 ± 0%  +8522992966.67%  (p=0.001 n=10+5)
```