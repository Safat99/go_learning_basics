# Some Important theories

## Concurrency and Parallelism
main difference is 

concurrency means ready for doing multiple task. But the server/system can do only one task at a time. So it switch between tasks and doing the multiple task. But it works so fast that human eyes can't feel the difference. So ultimately multiple task will be done. But note that, the server/system will only do one task at a time.

On the other hand parallesim is same. But it is slightly more advance. In this case we have multiple servers/threads/resources. So we can fire up different threads and do multiple task at the same time.

---
## Goroutines in Golang
* Goroutines is the way how we can achieve parallesim

### Difference/Advantage of Goroutine vs Thread(in other languages)

|Thread | Goroutine|
|---|---|
|Managed by OS| Managed by runtime|
|Fixed Stack - 1MB | Flexible stack - 2KB |
---

**Important Slogan** <br>
`Do not communicate by sharing memory; Instead, share memory by communicating`

Since the stack size is so tiny in golang, so we can fireup multiple threads and distribute works. 

Steps for Go routine:
* Adding the `go` keyword in front of any method/function
* use time.sleep 
* also there is another method sync package