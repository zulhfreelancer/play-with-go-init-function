Main

```
$ go run main.go
Hello from init
10
```

---

Test

```
$ go test -v .
Hello from init
Do stuff BEFORE the tests!
=== RUN   TestSum
--- PASS: TestSum (0.00s)
PASS
Do stuff AFTER the tests!
ok  	init-with-testmain.io	0.646s
```
