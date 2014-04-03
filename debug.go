package jdh

import (
    "fmt"
    "sync/atomic"
    "time"
    "runtime"
)

// prints message every second until returned thunk is executed.
// useful for determining what function calls are hanging.
func CheckHang(message string, args ...interface{}) func() {
  done := int32(0)
  go func() {

    // do nothing if done is called quickly; reduces printed messages
    time.Sleep(1000 * time.Millisecond)
    if atomic.LoadInt32(&done) != 0 {
      return
    }

    // loop until done, printing message at each interval
    for atomic.LoadInt32(&done) == 0 {
      fmt.Printf(message, args...)
      fmt.Printf("\t There are %v live goroutines.\n", runtime.NumGoroutine())
      time.Sleep(1000 * time.Millisecond)
    }

  }()

  // thunk to be called by client when done
  return func() {
    atomic.StoreInt32(&done, 1)
  }

}
