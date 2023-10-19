# Locks

Another way to synchronize goroutines is using a **mutex**, also known as a **lock**.
It's a simpler concept than channels.

Go provides locks as [`sync.Mutex`](https://pkg.go.dev/sync#Mutex).
When you call its `Lock()` method, all subsequent `Lock()` calls will block until `Unlock()` is called.

Locks come helpful when multiple goroutines access a single resource that is not *thread-safe* (i.e., can't be used by multiple goroutines at once).

A common example is a `map`. If multiple goroutines try to write to the same map, you'll get a panic:

```bash
fatal error: concurrent map writes
```

Another example could be a file on disk. If multiple goroutines try to write to a file, you might end up with inconsistent data.

To avoid these problems, use locks around the code that uses the common resource:

```go
var (
	lock sync.Mutex
	metrics = map[string]int{}
)

func SaveMetric(name string, value int) {
	lock.Lock()
	
	metrics[name] = value
	
	lock.Unlock()
}
```

When you need to handle errors, you need to be very careful to call `Unlock()` in every scenario.
If you miss one, you might end up with a deadlock and your application will freeze or crash.

This makes `defer` shine when used together with locks. Simply follow the `Lock()` call with a deferred `Unlock()`.
Whatever happens in the rest of the function, the lock releases automatically.

```go
func SaveMetricsToFile(content []byte) error {
	lock.Lock()
	defer lock.Unlock()
	
	return os.WriteFile("metrics.csv", content, 0644)	
}
```

The `sync.Mutex`'s zero-value is ready to use.

You can't copy a `Mutex` after it's initialized. In Go, structs are copied when passed as functions arguments.
If you pass a `Mutex` to a function or keep it as a struct field, store it as a pointer.

Remember that the zero-value of `*sync.Mutex` is `nil`, so you have to initialize it in this case.

```go
type MetricSaver struct {
	lock *sync.Mutex
}

func (m MetricSaver) SaveMetricsToFile(content []byte) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	
	return os.WriteFile("metrics.csv", content, 0644)
}

func main() {
	saver := MetricSaver{
		lock: &sync.Mutex{},
	}
	
	_ = saver.SaveMetricsToFile([]byte("my metrics"))
}
```


<div class="alert alert-dismissible bg-light-primary d-flex flex-column flex-sm-row p-7 mb-10">
    <div class="d-flex flex-column">
        <h3 class="mb-5 text-dark">
			<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-lightbulb text-primary" viewBox="0 0 16 16">
			  <path d="M2 6a6 6 0 1 1 10.174 4.31c-.203.196-.359.4-.453.619l-.762 1.769A.5.5 0 0 1 10.5 13a.5.5 0 0 1 0 1 .5.5 0 0 1 0 1l-.224.447a1 1 0 0 1-.894.553H6.618a1 1 0 0 1-.894-.553L5.5 15a.5.5 0 0 1 0-1 .5.5 0 0 1 0-1 .5.5 0 0 1-.46-.302l-.761-1.77a1.964 1.964 0 0 0-.453-.618A5.984 5.984 0 0 1 2 6zm6-5a5 5 0 0 0-3.479 8.592c.263.254.514.564.676.941L5.83 12h4.342l.632-1.467c.162-.377.413-.687.676-.941A5 5 0 0 0 8 1z"/>
			</svg>
			Tip
		</h3>
        <span>

Use constructors to initialize structs like this.

</span>
	</div>
	</div>

## Exercise

File: `15-concurrency/03-locks/main.go`

Add a locking mechanism to the `AddUser()` method, so the `users` map can't be accessed concurrently.

You can declare the lock as a global variable.
