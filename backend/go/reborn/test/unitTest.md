
### 单元测试
在不使用 “单元测试” 的情况下，我们如何测试一个函数或方法的正确性？
```go
func  Add(num1, num2 int) int {
    return num1 + num2
}

func  main() {
    excepted := 5
    actual := Add(2, 3)
    if excepted == actual {
        fmt.Println("成功")
    } else {
        fmt.Println("失败")
    }
}
```
这样的测试方式，它有这些问题：测试代码和业务代码混乱、不分离；测试完后，测试代码必须删除；如果不删除，会参与编译。

**什么是单元测试**  
单元测试又称为模块测试，是针对程序模块（软件设计的最小单元）来进行正确性检验的测试工作。在 Go 语言中，测试的最小单元常常是函数和方法。

**测试文件**  
在很多语言中，常常把测试文件放在一个独立的目录下进行管理，而在 Go 语言中会和源文件放置在一块，即同一目录下。  
假如源文件的命名是 xxx.go, 那单元测试文件的命名则为 xxx_test.go。如果在编译阶段 xxx_test.go 文件会被忽略。
```go
// 对于上面的 Add 函数，所在文件是 add.go，那创建的测试文件也和它放在一块
// unitest 目录
//     add.go
//     add_test.go 单元测试
```

**单元测试文件-基本结构&内容**  
```go
// gobasic/unittest/add_test.go
package unittest
// 导入 testing 标准包
import  "testing"
// 创建一个 Test 开头的函数名 TestAdd，Test 是固定写法，后面的 Add 一般和你要测试的函数名对应，当然不对应也没有问题
// 参数类型 *tesing.T 用于打印测试结果，参数中也必须跟上
func  TestAdd(t *testing.T) {
    // excepted 函数期待的结果
    excepted := 4
    // actual 函数真实计算的结果
    actual := Add(2, 3)
    // 如果不相等，打印出错误
    if excepted != actual {
        t.Errorf("excepted：%d, actual:%d", excepted, actual)
    }
}
```
在 unittest 目录下运行 go test （或 go test ./）命令，表示运行 unittest 目录下的单元测试，不会再往下递归。如果想往下递归，即当前目录下还有目录，则运行 go test ./... 命令。
```bash
$ go test
--- FAIL: TestAdd (0.00s)
     add_test.go:11: excepted：4, actual:5
 FAIL
 FAIL    github.com/miaogaolin/gobasic/unittest  0.228s
 FAIL
```

**\*testing.T**  
参数类型 T 中的几个方法：
- Error 打印错误日志、标记为失败 FAIL，并继续往下执行。
- Errorf 格式化打印错误日志、标记为失败 FAIL，并继续往下执行。
- Fail 不打印日志，结果中只标记为失败 FAIL，并继续往下执行。
- FailNow 不打印日志，结果中只标记为失败 FAIL，但在当前测试函数中不继续往下执行。
- Fatal 打印日志、标记为失败，并且内部调用了 FaileNow 函数，也不往下执行。
- Fatalf 格式化打印错误日志、标记为失败，并且内部调用了 FaileNow 函数，也不往下执行。

`没有成功的方法，只要没有通知错误，那就说明是正确的。`

**测试资源**   
有时候在写单元测试时，可能需要读取文件，那这些相关的资源文件就放置在 testdata 目录下。
```go
// unitest 目录
//     add.go
//     add_test.go 单元测试
//     testdata 目录
```

**go test 和 go vet**  
在运行 go test 命令后，go vet 命令也会自动运行。go vet 命令用于代码的静态分析，检查编译器检查不出的错误。  
`在测试时无需单独运行 go vet 命令，一个 go test 命令就包含了。`
```go
package main

import  "fmt"

func  main() {
    // 占位符 % d 需要的是整数，但给的是字符串
    fmt.Printf("%d", "miao")
}
```
对于这种类似的错误，编译器是不会报错的，这时候就用到了 go vet 命令。
```bash
$ go vet
# github.com/miaogaolin/gobasic/vet
.\main.go:6:2: Printf format %d has arg "miao" of wrong type string
```

**表格驱动测试**  
对于多种情况，如何进行测试呢？
```go
package unittest

import "testing"

func TestAddTable(t *testing.T) {
    /*
    excepted := 5
	actual := Add(2, 3)
	if excepted != actual {
		t.Errorf("case1：excepted：%d, actual:%d", excepted, actual)

	}

	excepted = 10
	actual = Add(0, 10)
	if excepted != actual {
		t.Errorf("case2：excepted：%d, actual:%d", excepted, actual)
	}
    */
    // 改写如下
	type param struct {
		name string
		num1, num2, excepted int
	}

	testCases := []param{
		{name: "case1", num1: 2, num2: 3, excepted: 5},
		{name: "case2", num1: 0, num2: 10, excepted: 10},
	}

    // 通过切片保存每种想要测试的情况（测试用例），下来只需要通过循环判断即可
	for _, v := range testCases {
        // t.Run 方法，第一个参数是当前测试的名称，第二个是个匿名函数，用来写判断逻辑
		t.Run(v.name, func(t *testing.T) {
			actual := Add(v.num1, v.num2)
			if v.excepted != actual {
				t.Errorf("excepted:%d, actual:%d", v.excepted, actual)
			}
		})
	}
}
```
运行结果：
```bash
# go test 命令后的 add.go 和 add_test.go 文件是特意指定需要测试和依赖的文件
# -test.run 指明测试的函数名
# -v 展示详细的过程，如果不写，测试成功时，不会打印详细过程
$  go test add.go add_test.go -test.run TestAddTable -v
=== RUN   TestAddTable
=== RUN   TestAddTable/case1
=== RUN   TestAddTable/case2
--- PASS: TestAddTable (0.00s)
    --- PASS: TestAddTable/case1 (0.00s)
    --- PASS: TestAddTable/case2 (0.00s)
PASS
ok      command-line-arguments  0.041s
```

**缓存**  
当运行单元测试时，测试的结果会被缓存下来。如果更改了测试代码或源文件，则会重新运行测试，并再次缓存。但不是任何情况都可以缓存下来，只有当 go test 命令后跟着目录、指定的文件或包名才可以。比如 go test ./ 或 go test add.go add_test.go 或 go test ./pkg 等。  
```bash
# 第一次
$ go test ./
ok      github.com/miaogaolin/gobasic/unittest  0.228s
# 第二次
$ go test ./
ok      github.com/miaogaolin/gobasic/unittest  (cached)

# 如果想禁用缓存，可以使用如下命令运行
go test ./ -count=1
```
这两种情况不会影响测试文件和源文件的修改，但还是会重新缓存测试结果：  
- 读取环境变量的内容更改
- 读取文件的内容更改

**并发测试**  
为了提高多个单元测试的运行效率，我们可以采取并发测试。
```go
// 没有并发时，go test 结果是
//  ok      command-line-arguments  3.242s
func TestA(t *testing.T) {
	time.Sleep(time.Second)
}
func TestB(t *testing.T) {
	time.Sleep(time.Second)
}
func TestC(t *testing.T) {
	time.Sleep(time.Second)
}

// 加入并发，go test 结果是 
// ok      command-line-arguments  1.049s
func TestA(t *testing.T) {
    // 在每个测试函数前增加了 t.Parallel() 实现并发
	t.Parallel()
	time.Sleep(time.Second)
}
func TestB(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second)
}

func TestC(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second)
}
```

**代码覆盖率**  
代码覆盖率是一个指数，例如：20%、30% 、100% 等。它体现了你的项目代码是否得到了足够的测试，指数越大，说明测试的覆盖情况越全面。
```bash
# -cover 输出覆盖率的标识符
# 覆盖率为 100%，说明被测试的函数代码都有运行到，覆盖率 = 已执行语句数 / 总语句数
$ go test -cover
PASS
coverage: 100.0% of statements
ok      github.com/miaogaolin/gobasic/unittest  1.045s
```

计算覆盖率有三种模式，不同的模式在已执行语句的次数统计时存在差异性。  
1）模式 set  
默认的模式，它的计算方式是 “如果同一语句多次执行只记录一次”。  
```go
func GetSex(sex int) string {
	if sex == 1 {
		return "男"
	} else {
		return "女"
	}
}

func TestGetSex(t *testing.T) {
	excepted := "男"
	actual := GetSex(1)
	if actual != excepted {
		t.Errorf("excepted：%s, actual:%s", excepted, actual)
	}
}
```
运行覆盖率命令，
```bash
$ go test -cover
ok      command-line-arguments  0.228s  coverage: 66.7% of statements

# 运行后，会在当前目录生成一个覆盖率的采样文件 profile
$ go test -coverprofile profile
$ cat profile
mode: set
github.com/miaogaolin/gobasic/testcover/sex.go:3.29,4.14 1 1
github.com/miaogaolin/gobasic/testcover/sex.go:4.14,6.3 1 1
github.com/miaogaolin/gobasic/testcover/sex.go:6.8,8.3 1 0
# 用 profile 这个文件生成一个直观图
# -html profile 指明将 profile 文件在浏览器渲染出来，运行后会自动在浏览器展示（灰色不用管，绿色的已覆盖，红色的未覆盖）
$ go tool cover -html profile
```
如果想达到 100% 覆盖，只需要增加 else 的测试情况
```go
func TestGetSex2(t *testing.T) {
	excepted := "女"
	actual := GetSex(0)
	if actual != excepted {
		t.Errorf("excepted：%s, actual:%s", excepted, actual)
	}
}
```

2）模式 count  
该模式和 set 模式比较相似，唯一的区别是 count 模式对于相同的语句执行次数会进行累计。  
`count 模式下能看出哪些代码执行的次数多，而 set 模式下不能`。
```bash
$ go test -coverprofile profile -covermode count
$ cat profile
mode: count
github.com/miaogaolin/gobasic/testcover/sex.go:3.29,4.14 1 2
github.com/miaogaolin/gobasic/testcover/sex.go:4.14,6.3 1 1
github.com/miaogaolin/gobasic/testcover/sex.go:6.8,8.3 1 1
```

3）模式 atomic  
该模式和 count 类似，都是统计执行语句的次数，不同点是，在并发情况下 atomic 模式比 count 模式计数更精确。  
```bash
$ go test -coverprofile profile -covermode atomic
```

**testify 包**  
当对一个项目中写大量的单元测试时，如果按照上述的方式去写，就会产生大量的判断语句。  
推荐一个第三方包 testfiy，
```go
package unittest

import (
	"github.com/stretchr/testify/assert" // 导入 testify 包下的一个子包 assert
	"testing"
)

func TestAdd(t *testing.T) {
	excepted := 4
	actual := Add(2, 3)
	assert.Equal(t, excepted, actual) // 使用 assert.Equal 函数简化 if 语句和日志打印，该函数期待 excepted 和 actual 变量相同，如果不相同会打印失败日志
}
```
