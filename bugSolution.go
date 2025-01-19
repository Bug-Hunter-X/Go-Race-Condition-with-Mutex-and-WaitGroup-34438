```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var m sync.Mutex
        m.Lock()
        wg.Add(1)
        go func() {
                defer wg.Done()
                fmt.Println("goroutine 1")
                m.Unlock()
        }()
        fmt.Println("main")
        wg.Wait()
}
```