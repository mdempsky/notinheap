This is a proof-of-concept demonstration of a userland check for
not-in-heap types.

```
$ go run pass/check.go -- pass/testdata/good.go
$ go run pass/check.go -- pass/testdata/bad.go
/home/mdempsky/wd/notinheap/pass/testdata/bad.go:9:25: heap allocation of x: new x (new)
/home/mdempsky/wd/notinheap/pass/testdata/bad.go:11:24: heap allocation of x: new x (complit)
/home/mdempsky/wd/notinheap/pass/testdata/bad.go:15:32: heap allocation of x: make []x n n
exit status 3
```

It relies on the
[buildssa](https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/buildssa)
analyzer pass, but in retrospect it could probably just as easily
achieve the same result using just go/ast and go/types directly.

The interesting cases to detect are:

1. Named parameters and variables of NIH type.
2. Composite literals whose address is taken (e.g., `&T{...}`).
3. Dynamic slice construction (e.g., `make([]T, ...)`).

There are other cases where values can end up on the Go heap (e.g.,
map and channel backing stores, non-pointer-shaped values converted to
interface type, large spilled temporaries), but none of these memory
locations are user-addressable (i.e., users cannot construct a pointer
to them).
