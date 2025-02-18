
### error 接口
错误处理在每个编程语言中都是一项重要内容，通常开发中遇到的分为异常与错误两种，Go 语言中也不例外。

在 C 语言中通过返回 -1 或者 NULL 之类的信息来表示错误，但是对于使用者来说，如果不查看相应的 API 说明文档，根本搞不清楚这个返回值究竟代表什么意思，比如返回 0 是成功还是失败？  
针对这样的情况，Go 语言中引入 error 接口类型作为错误处理的标准模式，如果函数要返回错误，则返回值类型列表中肯定包含 error。error 处理过程类似于 C 语言中的错误码，可逐层返回，直到被处理。  

Go 语言中返回的 error 类型究竟是什么呢？查看 Go 语言的源码就会发现 error 类型是一个非常简单的接口类型。  error 接口有一个签名为 Error () string 的方法，所有实现该接口的类型都可以当作一个错误类型。Error () 方法给出了错误的描述，在使用 fmt.Println 打印错误时，会在内部调用 Error () string 方法来得到该错误的描述。
```go
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
    Error() string
}
```

一般情况下，如果函数需要返回错误，就将 error 作为多个返回值中的最后一个（但这并非是强制要求）。  
创建一个 error 最简单的方法就是调用 errors.New 函数，它会根据传入的错误信息返回一个新的 error。  
```go
package main
import (
    "errors"
    "fmt"
    "math"
)
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return -1, errors.New("math: square root of negative number")
    }
    return math.Sqrt(f), nil
}
func main() {
    result, err := Sqrt(-13)
    if err != nil {
        fmt.Println(err) // math: square root of negative number
    } else {
        fmt.Println(result)
    }
}
```

除了 errors.New 用法之外，我们还可以使用 error 接口自定义一个 Error () 方法，来返回自定义的错误信息。
```go
package main
import (
    "fmt"
    "math"
)
type dualError struct {
    Num     float64
    problem string
}
func (e dualError) Error() string {
    return fmt.Sprintf("Wrong!!!,because \"%f\" is a negative number", e.Num)
}
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return -1, dualError{Num: f}
    }
    return math.Sqrt(f), nil
}
func main() {
    result, err := Sqrt(-13)
    if err != nil {
        fmt.Println(err) // Wrong!!!,because "-13.000000" is a negative number
    } else {
        fmt.Println(result)
    }
}
```

