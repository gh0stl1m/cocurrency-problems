package pkg

import (
  "fmt"
  "sync"
)

var wg sync.WaitGroup

func SequentialExecution() {

  evenCh, oddCh := make(chan bool, 1), make(chan bool, 1)
  defer close(evenCh)
  defer close(oddCh)

  wg = sync.WaitGroup{}
  wg.Add(2)

  go even(evenCh, oddCh)
  go odd(oddCh, evenCh)

  evenCh <- true

  wg.Wait()
}

func even(evenCh, oddCh chan bool) {

  defer wg.Done()
  for i := 2; i <= 5; i += 2 {
    <-oddCh
    fmt.Println(i)
    evenCh <- true
  }
}

func odd(oddCh, evenCh chan bool) {

  defer wg.Done()
  for i := 1; i <= 5; i += 2 {
    <-evenCh
    fmt.Println(i)
    oddCh <- true
  }
}
