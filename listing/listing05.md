```Golang
package main

type customError struct {
    msg string
}

func (e *customError) Error() string {
    return e.msg
}

func test() *customError {
    {
        // do something
    }
    return nil
}

func main() {
    var err error
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
}

```
Ответ: error

Функция возвращает nil указатель на customError, что не соответствует nil значению