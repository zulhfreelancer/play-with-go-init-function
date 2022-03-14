Main

```
$ go run main.go
Hello from init
10
```

---

Test

```
$ go test -v ./controller
Do stuff BEFORE the tests!
=== RUN   TestSum
--- PASS: TestSum (0.00s)
PASS
Do stuff AFTER the tests!
ok  	init-with-testmain.io/controller	1.278s
```
