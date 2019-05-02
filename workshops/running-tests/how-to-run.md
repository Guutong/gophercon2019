# Run the following commands:

- Run all tests in the package.
```bash
$go test

PASS
ok      github.com/guutong/gophercon2019/workshops/running-tests        3.007s
```
- Run all tests with verbose output.
```bash
$go test -v

=== RUN   TestTitle
=== RUN   TestTitle/rob
=== RUN   TestTitle/sandy
=== RUN   TestTitle/margaret
=== RUN   TestTitle/bill
=== RUN   TestTitle/tom
--- PASS: TestTitle (0.00s)
    --- PASS: TestTitle/rob (0.00s)
    --- PASS: TestTitle/sandy (0.00s)
    --- PASS: TestTitle/margaret (0.00s)
    --- PASS: TestTitle/bill (0.00s)
    --- PASS: TestTitle/tom (0.00s)
=== RUN   TestLong
--- PASS: TestLong (3.00s)
PASS
ok      github.com/guutong/gophercon2019/workshops/running-tests        3.006s
```
- Run only the "TestTitle" test with the only value being tested being "bill"
```bash
$go test -v -run TestTitle/bill

=== RUN   TestTitle
=== RUN   TestTitle/bill
--- PASS: TestTitle (0.00s)
    --- PASS: TestTitle/bill (0.00s)
PASS
ok      github.com/guutong/gophercon2019/workshops/running-tests        0.007s
```
- Run the tests with caching disabled
```bash
$go test -v -count=1

=== RUN   TestTitle
=== RUN   TestTitle/rob
=== RUN   TestTitle/sandy
=== RUN   TestTitle/margaret
=== RUN   TestTitle/bill
=== RUN   TestTitle/tom
--- PASS: TestTitle (0.00s)
    --- PASS: TestTitle/rob (0.00s)
    --- PASS: TestTitle/sandy (0.00s)
    --- PASS: TestTitle/margaret (0.00s)
    --- PASS: TestTitle/bill (0.00s)
    --- PASS: TestTitle/tom (0.00s)
=== RUN   TestLong
--- PASS: TestLong (3.00s)
PASS
ok      github.com/guutong/gophercon2019/workshops/running-tests        3.007s
```
- Bypass the long tests by turning on short testing
```bash
$go test -v -short

=== RUN   TestTitle
=== RUN   TestTitle/rob
=== RUN   TestTitle/sandy
=== RUN   TestTitle/margaret
=== RUN   TestTitle/bill
=== RUN   TestTitle/tom
--- PASS: TestTitle (0.00s)
    --- PASS: TestTitle/rob (0.00s)
    --- PASS: TestTitle/sandy (0.00s)
    --- PASS: TestTitle/margaret (0.00s)
    --- PASS: TestTitle/bill (0.00s)
    --- PASS: TestTitle/tom (0.00s)
=== RUN   TestLong
--- SKIP: TestLong (0.00s)
    strings_test.go:36: Skipping test due to short testing detected
PASS
ok      github.com/guutong/gophercon2019/workshops/running-tests        0.006s
```
- Run only the integration tests
```bash
$go test -v -tags=integration

=== RUN   TestIntegration
--- PASS: TestIntegration (3.00s)
    strings_integration_test.go:13: running long integration test...
PASS
ok      github.com/guutong/gophercon2019/workshops/running-tests        3.010s
```
Hint: It's easier to check your results by always using the verbose flag to see all the testing output.
