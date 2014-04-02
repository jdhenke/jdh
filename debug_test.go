package jdh

import (
  "testing"
  "time"
  "fmt"
  "sync"
)


// example where different go routines sleep for different lengths.
// illustrates use of CheckHang to identify which ones are still going.
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

  // Would produce something like this.
  //
  // Start: 0
  // Start: 1
  // Start: 2
  // Start: 3
  // End: 0
  // End: 1
  // 	 2 still running...
  // 	 3 still running...
  // End: 2
  // 	 3 still running...
  // End: 3
}



func TestCheckHang(t *testing.T) {
  ExampleCheckHang()
}
