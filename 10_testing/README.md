# Testing

The code contained into this folder is a collection of 
tests, benchmarks, and examples (test examples, not code examples...
 more details in the book). 

To execute a test from the source folder you can run this:
```
go test testing/example_01/example01_test.go -v
=== RUN   TestMe
--- PASS: TestMe (0.00s)
PASS
ok  	command-line-arguments	0.135s
```
To run a benchmark:

```
go test -bench . benchmarking/example_01/example01_test.go -v
goos: darwin
goarch: amd64
BenchmarkSum
b.N: 1
b.N: 100
b.N: 4008
BenchmarkSum-16    	    4008	    259175 ns/op
PASS
ok  	command-line-arguments	1.153s
```
To run an example:
```
go test -v example/example_01/example01_test.go
=== RUN   ExampleUser
--- PASS: ExampleUser (0.00s)
=== RUN   ExampleCommonFriend
--- PASS: ExampleCommonFriend (0.00s)
=== RUN   ExampleUser_GetUserId
--- PASS: ExampleUser_GetUserId (0.00s)
=== RUN   ExampleUser_CountFriends
--- PASS: ExampleUser_CountFriends (0.00s)
PASS
ok  	command-line-arguments	(cached)
```
