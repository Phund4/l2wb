```Golang
package main
 
import (
    "fmt"
    "os"
)
 
func Foo() error {
    var err *os.PathError = nil
    return err
}
 
func main() {
    err := Foo()
    fmt.Println(err)
    fmt.Println(err == nil)
}

```

Ответ: <nil> false

Интерфейс под капотом:

```Golang
type iface struct {
	tab  *itab
	data unsafe.Pointer
}
```

```Golang
type itab struct {
	inter *interfacetype
	_type *_type
	hash  uint32 // copy of _type.hash. Used for type switches.
	_     [4]byte
	fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
}
```

Интерфейс представляет собой структуру из 2 полей: динамического типа и значения и статического типа.
Динамическое значение, то есть data, будет nil. Интерфейс будет nil, когда inter и _type будут nil одновременно. 
_type это присвоенное нами nil значение. inter это статический тип интерфейса, то есть *os.PathError.
Получаем, что интерфейс не будет равняться nil.
