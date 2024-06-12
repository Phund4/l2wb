```Golang
package main
 
import (
    "fmt"
    "math/rand"
    "time"
)
 
func asChan(vs ...int) <-chan int {
   c := make(chan int)
 
   go func() {
       for _, v := range vs {
           c <- v
           time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
      }
 
      close(c)
  }()
  return c
}
 
func merge(a, b <-chan int) <-chan int {
   c := make(chan int)
   go func() {
       for {
           select {
               case v := <-a:
                   c <- v
              case v := <-b:
                   c <- v
           }
      }
   }()
 return c
}
 
func main() {
 
   a := asChan(1, 3, 5, 7)
   b := asChan(2, 4 ,6, 8)
   c := merge(a, b )
   for v := range c {
       fmt.Println(v)
   }
}

```

Ответ: 1 2 3 4 6 8 5 7 0 0 0 0 0

Закрываем каналы a и b, при этом продолжая считывать из них значения (по умолчанию 0). 
И не закрываем канал c, поэтому мы видим бесконечный вывод 0.
Чтобы программа корректно смерджила каналы, необходимо делать проверку на закрытие каналов a и b.