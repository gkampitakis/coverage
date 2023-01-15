# Coverage

Coverage is a simple tool that checks per package coverage and fails if is not met.

```bash
go get github.com/gkampitakis/coverage
```

## Usage

You can use it by calling in your `TestMain`. For more information around [TestMain](https://pkg.go.dev/testing#hdr-Main).


```go
func TestMain(m *testing.M) {
	coverage.Run(m, 95)
}
```

What `coverage.Run` does, it runes the tests with `t.Run()` and then depending on the coverage fails the tests and calls `os.Exit` with correct exit code.

In case you have "clean" up code after running your tests `coverage.Run` support passing a callback function `func(t *testing.M)` for running your code. 

e.g. 

```go
func TestMain(m *testing.M) {
	coverage.Run(m, 91, func(t *testing.M) {
		fmt.Println("tests are done")
	})
}
```

## Example

```
PASS
coverage: 89.5% of statements

FAIL    Coverage threshold not met 91.0 >= 85.0 for my-package/example

FAIL    my-package/example        0.185s
```

> `coverage: 89.5% of statements` reports slightly different number from the
`testing.Coverage`. Couldn't find a way to fix this or an explanation.

`coverage.Run` will also fail when tests are run with no `cover` flag.

e.g.

```
PASS

FAIL    coverage is not enabled. You can enable by using `-cover`
```
