# sort2
A mostly drop-in replacement for "sort" that uses the new [generics](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md), via the current tip version of go (soon to be Go 1.18).

This is intended as an experimental demonstration of the more performant and memory-efficient sorting that will come with Go 1.18 generics. In order to achieve best performance it is receommended that you avoid Sort for now and stick to Slice functions, which appears to work efficiently. Hopefully the release of Go 1.18 will resolve that.

Go generics are not yet stable and should not be used in production, however, this is good for getting a feel for how generic sorting may work in Go 2.

This follows [the June 16, 2020 generics design draft](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md).

### How?

When using Go 1.18, change `import "sort"` to `import "github.com/argusdusty/sort2"` and `sort.` calls to `sort2.` calls. That's it. If you're running a development version of go1.18 you may need to supply `-gcflags="-lang=go1.18"`. You may need to instantiate some sort functions to your generic type to use them.

### Benchmarks

```
C:\Users\argusdusty\Code\src\github.com\argusdusty\sort2>go version
go version devel go1.18-a1938435d6 Mon Sep 6 19:10:15 2021 +0000 windows/amd64
C:\Users\argusdusty\Code\src\github.com\argusdusty\sort2>go test -run=None -bench=. -benchmem -cpu=1 -count=20 sort > old.txt
C:\Users\argusdusty\Code\src\github.com\argusdusty\sort2>go test -run=None -bench=. -benchmem -cpu=1 -gcflags="-lang=go1.18" -count=10 > new.txt
C:\Users\argusdusty\Code\src\github.com\argusdusty\sort2>benchstat old.txt new.txt
name                old time/op    new time/op        delta
SearchWrappers        75.1ns ± 2%        36.6ns ± 1%          -51.29%  (p=0.000 n=20+10)
SortString1K          85.0µs ± 1%        67.5µs ± 1%          -20.54%  (p=0.000 n=20+10)
SortString1K_Slice    82.5µs ± 1%        80.6µs ± 1%           -2.34%  (p=0.000 n=19+10)
StableString1K         110µs ± 1%         784µs ± 0%         +613.39%  (p=0.000 n=20+8)
SortInt1K             39.0µs ± 2%        11.6µs ± 2%          -70.31%  (p=0.000 n=20+10)
StableInt1K           51.4µs ± 2%       561.8µs ± 0%         +992.89%  (p=0.000 n=20+10)
StableInt1K_Slice     49.0µs ± 1%        32.7µs ± 1%          -33.34%  (p=0.000 n=19+10)
SortInt64K            4.29ms ± 1%        1.95ms ± 2%          -54.50%  (p=0.000 n=19+10)
SortInt64K_Slice      4.08ms ± 1%        3.96ms ± 1%           -2.75%  (p=0.000 n=20+10)
StableInt64K          4.50ms ± 1%       52.00ms ± 1%        +1055.11%  (p=0.000 n=20+10)
Sort1e2               27.3µs ± 3%       165.4µs ± 5%         +506.01%  (p=0.000 n=20+9)
Stable1e2             52.1µs ± 4%       385.5µs ± 3%         +639.26%  (p=0.000 n=20+9)
Sort1e4               5.63ms ± 1%       36.27ms ± 2%         +543.77%  (p=0.000 n=19+10)
Stable1e4             16.8ms ± 2%       120.8ms ± 1%         +619.75%  (p=0.000 n=20+9)
Sort1e6                867ms ± 1%        5569ms ± 1%         +542.18%  (p=0.000 n=20+10)
Stable1e6              3.43s ± 1%        28.53s ± 9%         +731.46%  (p=0.000 n=20+9)

name                old alloc/op   new alloc/op       delta
SearchWrappers         0.00B              0.00B                  ~     (all equal)
SortString1K           24.0B ± 0%          0.0B              -100.00%  (p=0.000 n=20+10)
SortString1K_Slice     56.0B ± 0%          0.0B              -100.00%  (p=0.000 n=20+10)
StableString1K         24.0B ± 0%     551664.0B ± 0%     +2298500.00%  (p=0.000 n=20+10)
SortInt1K              24.0B ± 0%          0.0B              -100.00%  (p=0.000 n=20+10)
StableInt1K            24.0B ± 0%     445896.0B ± 0%     +1857800.00%  (p=0.000 n=20+10)
StableInt1K_Slice      56.0B ± 0%          0.0B              -100.00%  (p=0.000 n=20+10)
SortInt64K             24.0B ± 0%          0.0B              -100.00%  (p=0.000 n=20+10)
SortInt64K_Slice       56.0B ± 0%          0.0B              -100.00%  (p=0.000 n=20+10)
StableInt64K           24.0B ± 0%   40169208.0B ± 0%   +167371600.00%  (p=0.000 n=20+10)
Sort1e2                 168B ± 0%       121421B ± 0%       +72174.23%  (p=0.000 n=20+10)
Stable1e2               168B ± 0%       284134B ± 0%      +169027.26%  (p=0.000 n=20+10)
Sort1e4                 169B ± 0%     27095558B ± 0%    +16032774.62%  (p=0.000 n=20+10)
Stable1e4               171B ± 0%     86415189B ± 0%    +50535098.25%  (p=0.000 n=18+10)
Sort1e6                 288B ± 0%   4190006232B ± 0%  +1454863175.00%  (p=0.000 n=20+10)
Stable1e6               408B ± 0%  18409665024B ± 0%  +4512172700.00%  (p=0.000 n=20+9)

name                old allocs/op  new allocs/op      delta
SearchWrappers          0.00               0.00                  ~     (all equal)
SortString1K            1.00 ± 0%          0.00              -100.00%  (p=0.000 n=20+10)
SortString1K_Slice      2.00 ± 0%          0.00              -100.00%  (p=0.000 n=20+10)
StableString1K          1.00 ± 0%      22986.00 ± 0%     +2298500.00%  (p=0.000 n=20+10)
SortInt1K               1.00 ± 0%          0.00              -100.00%  (p=0.000 n=20+10)
StableInt1K             1.00 ± 0%      18579.00 ± 0%     +1857800.00%  (p=0.000 n=20+10)
StableInt1K_Slice       2.00 ± 0%          0.00              -100.00%  (p=0.000 n=20+10)
SortInt64K              1.00 ± 0%          0.00              -100.00%  (p=0.000 n=20+10)
SortInt64K_Slice        2.00 ± 0%          0.00              -100.00%  (p=0.000 n=20+10)
StableInt64K            1.00 ± 0%    1673717.00 ± 0%   +167371600.00%  (p=0.000 n=20+10)
Sort1e2                 7.00 ± 0%       5058.70 ± 0%       +72167.14%  (p=0.000 n=20+10)
Stable1e2               7.00 ± 0%      11838.60 ± 0%      +169022.86%  (p=0.000 n=20+10)
Sort1e4                 7.00 ± 0%    1128980.70 ± 0%    +16128195.71%  (p=0.000 n=20+10)
Stable1e4               7.00 ± 0%    3600632.00 ± 0%    +51437500.00%  (p=0.000 n=20+10)
Sort1e6                 8.00 ± 0%  174583593.00 ± 0%  +2182294812.50%  (p=0.000 n=20+10)
Stable1e6               9.00 ± 0%  767069376.00 ± 0%  +8522992966.67%  (p=0.000 n=20+9)
```
Unfortunately due to how generics are implements right now, Go will convert between the input type and an interface type at every .Less/.Swap function call, each time requiring a new allocation. For some reason the Slice methods are smart enough to not do this, and thus run faster and require zero allocations.

Hopefully before the release of Go 1.18 this will be fixed and this library will be more performant than a non-generic sorting library.