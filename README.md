This is a proof-of-concept demonstration of a userland check for
not-in-heap types.

```
$ go run pass/check.go -- pass/testdata/good.go
$ go run pass/check.go -- pass/testdata/bad.go
/home/mdempsky/wd/notinheap/pass/testdata/bad.go:9:25: heap allocation of x: new x (new)
/home/mdempsky/wd/notinheap/pass/testdata/bad.go:11:24: heap allocation of x: new x (complit)
/home/mdempsky/wd/notinheap/pass/testdata/bad.go:15:27: heap allocation of [10]x: new [10]x (makeslice)
exit status 3
```
