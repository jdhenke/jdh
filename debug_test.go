package jdh

import (
  "testing"
  "time"
  "fmt"
  "sync"
)

func ExampleCheckHang() {
  var wg sync.WaitGroup
  for i := 0; i < 4; i += 1 {
    wg.Add(1)
    go func(i int) {
      defer wg.Done()
      done := CheckHang("\t %v still running...\n", i)
      fmt.Printf("Start: %v\n", i)
      time.Sleep(time.Duration(i) * time.Second)
      done()
      fmt.Printf("End: %v\n", i)
    }(i)
  }
  wg.Wait()
}

func TestExample(t *testing.T) {
  ExampleCheckHang()
}
