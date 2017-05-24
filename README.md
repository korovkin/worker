# Sync / Async serial worker queue (GCD)

basic functionality you would expect from: 
https://developer.apple.com/reference/dispatch

## builds

[![Build Status](https://travis-ci.org/korovkin/worker.svg)](https://travis-ci.org/korovkin/worker)

## example

GCD style serial execution of function (lambdas / blocks)

```
  import "github.com/korovkin/worker"
  ...

  w := worker.NewWorkerDefault("test")
  for i := 0; i < 1000; i++ {
    i := i
    worker.Enqueue(func() {
        fmt.Println(" async: i:", i)
    })
  }
  worker.EnqueueSync(func() {
      fmt.Println("sync!")
  })
```

