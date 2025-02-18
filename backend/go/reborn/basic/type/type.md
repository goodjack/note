
### Go 数据类型
Go 语言的类型大体可分为基本数据类型、复合数据类型和其他类型这三种。
- 基本数据类型：整型、浮点型、复数、字符串、字符型、布尔型
- 复合（容器）数据类型：数组、切片 (slice)、通道 (chan)、映射 (map)
- 其它数据类型：函数、结构体、接口

[cast 支持所有内置类型之间的转换](https://github.com/spf13/cast)

**基本数据类型之整形**  
int 和 uint 类型长度由操作系统类型决定，如果系统是 32 位的，那么它们都是 32 位（4 字节）；如果是 64 位的，那么它们都是 64 位（8 字节）。
- 无符号：uint、uint8，uint16，uint32，uint64
- 有符号：int、int8，int16，int32，int64
- 三种特殊类型：uint、int、uintptr 依赖因具体的计算机架构不同而不同，在 32 位的计算机架构上是 32 位，64 位的计算机架构则是 64 位

| 类型 | 精度 |  
| :---: | :---: |  
| uint8	 | 无符号 8 位整型 (0 到 2^8-1) |  
| uint16 | 无符号 16 位整型 (0 到 2^16-1) |  
| uint32 | 无符号 32 位整型 (0 到 2^32-1) |  
| uint64 | 无符号 64 位整型 (0 到 2^64-1) |  
| int8 | 有符号 8 位整型 (-2^7 到 2^7-1) |  
| int16 | 有符号 16 位整型 (-2^15 到 2^15-1) |  
| int32 | 有符号 32 位整型 (-2^31 到 2^31) |  
| int64 | 有符号 64 位整型 (-2^64 到 2^64) |  
| uintptr | 长度4或8字节 | 存储指针的 uint32 或 uint64 整数 |  

```go
// 取值范围
package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	fmt.Println("各int类型的大小: ")
	var i1 int = 1
	var i2 int8 = 2
	var i3 int16 = 3
	var i4 int32 = 4
	var i5 int64 = 5
	var i6 uint64 = 6
	fmt.Printf("int       : %v\n", unsafe.Sizeof(i1)) // int       : 8
	fmt.Printf("int8  : %v\n", unsafe.Sizeof(i2)) // int8  : 1
	fmt.Printf("int16 : %v\n", unsafe.Sizeof(i3)) // int16 : 2
	fmt.Printf("int32 : %v\n", unsafe.Sizeof(i4)) // int32 : 4
	fmt.Printf("int64 : %v\n", unsafe.Sizeof(i5)) // int64 : 8
	fmt.Printf("uint64: %v\n", unsafe.Sizeof(i6)) // uint64: 8

	// 输出各int类型的取值范围
	fmt.Println("int8:", math.MinInt8, "~", math.MaxInt8) // int8: -128 ~ 127
	fmt.Println("int16:", math.MinInt16, "~", math.MaxInt16) // int16: -32768 ~ 32767
	fmt.Println("int32:", math.MinInt32, "~", math.MaxInt32) // int32: -2147483648 ~ 2147483647
	fmt.Println("int64:", math.MinInt64, "~", math.MaxInt64) // int64: -9223372036854775808 ~ 9223372036854775807
	fmt.Println()

	// n是自动推导类型
	n := 1234567890
	fmt.Printf("n := 1234567890 的默认类型为: %T\n", n) // n := 1234567890 的默认类型为: int
	fmt.Printf("int类型的字节数为: %v\n\n", unsafe.Sizeof(n)) // int类型的字节数为: 8

	// 初始化一个32位整型值
	var a int32 = 987654321
	fmt.Println("var a int32 = 987654321") // var a int32 = 987654321

	// 输出变量的十六进制形式和十进制
	fmt.Printf("int32: 十六进制为0x%x, 十进制为%d\n", a, a) // int32: 十六进制为0x3ade68b1, 十进制为987654321

	// 将a转换为int8类型, 发生数值截断
	b := int8(a)
	fmt.Printf("int8: 十六进制0x%x, 十进制为%d\n", b, b) // int8: 十六进制0x-4f, 十进制为-79

	// 将a转换为int16类型, 发生数值截断
	c := int16(a)
	fmt.Printf("int16: 十六进制为0x%x, 十进制%d\n", c, c) // int16: 十六进制为0x68b1, 十进制26801

	// 将a转换为int64类型
	d := int64(a)
	fmt.Printf("int64: 十六进制为0x%x, 十进制%d\n", d, d) // int64: 十六进制为0x3ade68b1, 十进制987654321
}
```

**基本数据类型之浮点形**  
对于浮点类型，只有 float32 和 float64，没有 float 类型，使用 IEEE-754 标准。  
float32 精确到小数点后 7 位，float64 精确到小数点后 15 位。由于精度的问题，在进行数据比对的时候，就要考虑精度损失。

| 类型 | 精度 |  
| :---: | :---: |  
| float32 | ±1.18×10^-38 到 ±3.4×10^38 |  
| float64 | 2.23×10^-308 到 ±1.8×10^308 |  

**基本数据类型之复数形**  
复数类型使用 complex64 和 complex128 声明，complex64 是 32 位实数和虚数，complex128 是 64 位实数和虚数。
```go
var c1 complex64 = 5 + 10i
// 或
c1 := 5 + 10i
```

| 类型 | 精度 |  
| :---: | :---: |  
| complex64 | 含 float32 位实数和 float32 位虚数 |  
| complex128 | 含 float64 位实数和 float64 位虚数 |  

**基本数据类型之布尔形**  
布尔型使用 bool 声明，要么是 true, 要么是 false, 默认是 false。  
if 和 for 语句的条件部分都是布尔类型的值，并且 == 和 < 等比较运算符也会产生布尔型的值。一元操作符 ! 对应逻辑非操作，因此 !true 的值为 false。


| 类型 | 精度 |  
| :---: | :---: |  
| bool | true 或 false |  

**基本数据类型之字符串形**  
字符串类型使用 string 声明，在初始化时分为单行和多行。单行使用双引号；多行使用反引号，中间如果出现转义字符时，不会进行解析，原样输出。
```go
str := "我最棒"
str := `
你好，\n

Hello
`
fmt.Println(str)
// 你好，\n
// 
// Hello
```

| 类型 | 精度 |  
| :---: | :---: |  
| string | UTF-8 编码标识的 Unicode 文本 |  


**基本数据类型之字符形**  
byte 和 rune 是字符型，但其实不是新增的类型，byte 是 uint8 的别名，占用 1 个字节，rune 是 int32 的别名，占用 4 个字节。  
byte 代表了 ASCII 码的一个字符，rune 代表了 UTF-8 的一个字符。
```go
// 对应ASCII码表中的97
var a byte = 'a'
// 或
var a uint8 = 'a'
var b rune = '嗨'
// 或
var b int32 = '嗨'
```

**类型零值**
Go 语言的每种原生类型都有它的默认值，这个默认值就是这个类型的零值。  
> 整形默认 0  
> 浮点默认 0.0  
> 布尔默认 FALSE  
> 字符串默认 ""   
> 指针、接口、切片、channel、map 和函数默认 nil  


### Go 数值类型
Go 编码中使用最多的就是基本数据类型，而基本数据类型中使用占比最大的又是数值类型。Go 语言原生支持的数值类型包括整型、浮点型以及复数类型。  
> Go 语言的整型，主要用来表示现实世界中整型数量，比如：人的年龄、班级人数等；  
> 整型可以分为平台无关整型和平台相关整型这两种，它们的区别主要就在，这些整数类型在不同 CPU 架构或操作系统下面，它们的长度是否是一致的；  
> 平台无关的整型也可以分成两类：有符号整型（int8~int64）和无符号整型（uint8~uint64）。两者的本质差别在于最高二进制位（bit 位）是否被解释为符号位，这点会影响到无符号整型与有符号整型的取值范围；  
> Go 语言原生提供了三个平台相关整型，它们是 int、uint 与 uintptr；由于这三个类型的长度是平台相关的，所以我们在编写有移植性要求的代码时，千万不要强依赖这些类型的长度。如果你不知道这三个类型在目标运行平台上的长度，可以通过 unsafe 包提供的 SizeOf 函数来获取。  
```go
var a, b = int(5), uint(6)
var p uintptr = 0x12345678
fmt.Println("signed integer a's length is", unsafe.Sizeof(a)) // 8
fmt.Println("unsigned integer b's length is", unsafe.Sizeof(b)) // 8
fmt.Println("uintptr's length is", unsafe.Sizeof(p)) // 8
```

**整型的溢出问题**  
无论哪种整型，都有它的取值范围，也就是有它可以表示的值边界。  
如果这个整型因为参与某个运算，导致结果超出了这个整型的值边界，我们就说发生了整型溢出的问题。由于整型无法表示它溢出后的那个 “结果”，所以出现溢出情况后，对应的整型变量的值依然会落到它的取值范围内，只是结果值与我们的预期不符，导致程序逻辑出错。
```go
var s int8 = 127
s += 1 // 预期128，实际结果-128

var u uint8 = 1
u -= 2 // 预期-1，实际结果255
```

**字面值与格式化输出**  
Go 语言在设计开始，就继承了 C 语言关于数值字面值（Number Literal）的语法形式。  
早期 Go 版本支持十进制、八进制、十六进制的数值字面值形式；Go 1.13 版本中，Go 又增加了对二进制字面值的支持和两种八进制字面值的形式。  
```go
a := 53        // 十进制
b := 0700      // 八进制，以"0"为前缀
c1 := 0xaabbcc // 十六进制，以"0x"为前缀
c2 := 0Xddeeff // 十六进制，以"0X"为前缀


d1 := 0b10000001 // 二进制，以"0b"为前缀
d2 := 0B10000001 // 二进制，以"0B"为前缀
e1 := 0o700      // 八进制，以"0o"为前缀
e2 := 0O700      // 八进制，以"0O"为前缀
```
为提升字面值的可读性，Go 1.13 版本还支持在字面值中增加数字分隔符 “\_”，分隔符可以用来将数字分组以提高可读性。  
```go
a := 5_3_7   // 十进制: 537
b := 0b_1000_0111  // 二进制位表示为10000111 
c1 := 0_700  // 八进制: 0700
c2 := 0o_700 // 八进制: 0700
d1 := 0x_5c_6d // 十六进制：0x5c6d
```
我们也可以通过标准库 fmt 包的格式化输出函数，将一个整型变量输出为不同进制的形式。  
```go
var a int8 = 59
fmt.Printf("%b\n", a) //输出二进制：111011
fmt.Printf("%d\n", a) //输出十进制：59
fmt.Printf("%o\n", a) //输出八进制：73
fmt.Printf("%O\n", a) //输出八进制(带0o前缀)：0o73
fmt.Printf("%x\n", a) //输出十六进制(小写)：3b
fmt.Printf("%X\n", a) //输出十六进制(大写)：3B
```

**浮点型**  
浮点型的使用场景，主要集中在科学数值计算、图形图像处理和仿真、多媒体游戏以及人工智能等领域。  
IEEE 754 是 IEEE 制定的二进制浮点数算术标准，它是 20 世纪 80 年代以来最广泛使用的浮点数运算标准，被许多 CPU 与浮点运算器采用。现存的大部分主流编程语言，包括 Go 语言，都提供了符合 IEEE 754 标准的浮点数格式与算术运算。IEEE 754 标准规定了四种表示浮点数值的方式：单精度（32 位）、双精度（64 位）、扩展单精度（43 比特以上）与扩展双精度（79 比特以上，通常以 80 位实现）。后两种其实很少使用，我们重点关注前面两个就好了。  

Go 语言提供了 float32 与 float64 两种浮点类型，它们分别对应的就是 IEEE 754 中的单精度与双精度浮点数值类型。不过，这里要注意，Go 语言中没有提供 float 类型。这不像整型那样，Go 既提供了 int16、int32 等类型，又有 int 类型。换句话说，*Go 提供的浮点类型都是平台无关的*。  
无论是 float32 还是 float64，它们的变量的默认值都为 0.0，不同的是它们占用的内存空间大小是不一样的，可以表示的浮点数的范围与精度也不同。  

Go 浮点类型字面值大体可分为两类，一类是直白地用十进制表示的浮点值形式；另一类则是科学计数法形式。采用科学计数法表示的浮点字面值，我们需要通过一定的换算才能确定其浮点值。而且，科学计数法形式又分为十进制形式表示的，和十六进制形式表示的两种。  
```go
3.1415
.15  // 整数部分如果为0，整数部分可以省略不写
81.80
82. // 小数部分如果为0，小数点后的0可以省略不写

// 十进制科学计数法形式的浮点数字面值，这里字面值中的 e/E 代表的幂运算的底数为 10
6674.28e-2 // 6674.28 * 10^(-2) = 66.742800
.12345E+5  // 0.12345 * 10^5 = 12345.000000

// 十六进制科学计数法形式的浮点数
0x2.p10  // 2.0 * 2^10 = 2048.000000
0x1.Fp+0 // 1.9375 * 2^0 = 1.937500
```

和整型一样，fmt 包也提供了针对浮点数的格式化输出。我们最常使用的格式化输出形式是 % f。通过 % f，我们可以输出浮点数最直观的原值形式。
```go
var f float64 = 123.45678
fmt.Printf("%f\n", f) // 123.456780

fmt.Printf("%e\n", f) // 1.234568e+02
fmt.Printf("%x\n", f) // 0x1.edd3be22e5de1p+06
```

**复数类型**  
形如 z=a+bi（a、b 均为实数，a 称为实部，b 称为虚部）的数称为复数。  
复数类型在 Go 中的应用很局限和小众，主要用于专业领域的计算，比如矢量计算等。  

Go 提供两种复数类型，它们分别是 complex64 和 complex128，complex64 的实部与虚部都是 float32 类型，而 complex128 的实部与虚部都是 float64 类型。如果一个复数没有显示赋予类型，那么它的默认类型为 complex128。  
```go
// 1，可以通过复数字面值直接初始化一个复数类型变量
var c = 5 + 6i
var d = 0o123 + .12345E+5i // 83+12345i

// 2，Go 还提供了 complex 函数，方便我们创建一个 complex128 类型值
var c = complex(5, 6) // 5 + 6i
var d = complex(0o123, .12345E+5) // 83+12345i

// 3，还可以通过 Go 提供的预定义的函数 real 和 imag，来获取一个复数的实部与虚部，返回值为一个浮点类型
var c = complex(5, 6) // 5 + 6i
r := real(c) // 5.000000
i := imag(c) // 6.000000
```
至于复数形式的格式化输出的问题，由于 complex 类型的实部与虚部都是浮点类型，所以我们可以直接运用浮点型的格式化输出方法，来输出复数类型。  

**创建自定义的数值类型**  
通过 Go 提供的类型定义语法，来创建自定义的数值类型，我们可以通过 type 关键字基于原生数值类型来声明一个新类型。  
但是自定义的数值类型，在和其他类型相互赋值时容易出现一些问题。  
```go
type MyInt int32


var m int = 5
var n int32 = 6
var a MyInt = m // 错误：在赋值中不能将m（int类型）作为MyInt类型使用
var a MyInt = n // 错误：在赋值中不能将n（int32类型）作为MyInt类型使用
```
要避免这个错误，我们需要借助显式转型，让赋值操作符左右两边的操作数保持类型一致。  
```go
var m int = 5
var n int32 = 6
var a MyInt = MyInt(m) // ok
var a MyInt = MyInt(n) // ok

// 也可以通过 Go 提供的类型别名（Type Alias）语法来自定义数值类型
// 和上面使用标准 type 语法的定义不同的是，通过类型别名语法定义的新类型与原类型别无二致，可以完全相互替代
type MyInt = int32
var n int32 = 6
var a MyInt = n // ok
```

**类型别名**  
和自定义类型是不一样的，类型别名和原类型是完全等价的，不需要类型转化，只是名称不一样而已。在内置类型中，byte 类型就是 uint8 类型的别名。
```go
type byte = uint8
```

### Go 类型转换
在某些特定的场景下，我们需要对数据进行类型转换才能继续后面的逻辑（如某个函数需要 float64 类型参数，需要将现有 int64 类型值传入其中时）。在 Go 语言中，进行类型转换有两个要注意的地方，分别是：
- 只能进行相同类别的转换，如将 int32 转换为 int64。不同类别的转换将引发编译时错误，如将 bool 转换为 string；
- 若将取值范围较大的类型转换为取值范围较小的类型，且实际值超过取值范围较小的类型时，将发生精度丢失的情况。
```go
// 显式转换
var b = int32(13)

//声明float32型变量exampleFloat32并赋值
var exampleFloat32 float32 = 150.25
//将exampleFloat32转换为float64类型，并将结果赋值给exampleFloat64
exampleFloat64 := float64(exampleFloat32)
//输出exampleFloat64的类型和值
fmt.Println(reflect.TypeOf(exampleFloat64), exampleFloat64) // float64 150.25
//将exampleFloat32转换为int32类型，exampleInt32
exampleInt32 := int32(exampleFloat32)
//输出exampleInt32的类型和值
fmt.Println(reflect.TypeOf(exampleInt32), exampleInt32) // int32 150
```

通过包 strconv 的 Itoa 函数可以把一个 int 类型转为 string，Atoi 函数则用来把 string 转为 int。
```go
i2s:=strconv.Itoa(i)
s2i,err:=strconv.Atoi(i2s)
fmt.Println(i2s,s2i,err)
```

对于浮点数、布尔型，Go 语言提供了 strconv.ParseFloat、strconv.ParseBool、strconv.FormatFloat 和 strconv.FormatBool 进行互转。

对于数字类型之间，可以通过强制转换的方式。
```go
i2f:=float64(i)
f2i:=int(f64)
fmt.Println(i2f,f2i)
```

