##  go test工具

Go语言中的测试依赖 go test命令。

go test命令是一个按照一定约定和组织的测试代码的驱动程序。在包目录内，所有以_test.go为后缀名的源代码文件都是go test测试的一部分，不会被go build编译到最终的可执行文件中。

在*_test.go文件中有三种类型的函数，单元测试函数、基准测试函数和示例函数。

| 类型     | 格式                  | 作用                           |
| -------- | --------------------- | ------------------------------ |
| 测试函数 | 函数名前缀为Test      | 测试程序的一些逻辑行为是否正确 |
| 基准函数 | 函数名前缀为Benchmark | 测试函数的性能                 |
| 示例函数 | 函数名前缀为Example   | 为文档提供实例文档             |

go test命令会遍历所有的*_test.go文件中符合上述命名规则的函数，然后生成一个临时的main包用于调用相应的测试函数，然后构建并运行、报告测试结果，最后清理测试中生成的临时文件。

## 单元测试函数

### 格式

每个测试函数必须导入`testing`包，测试函数的基本格式（签名）如下：

```go
func TestName(t *testing.T){
	...
}
```

测试函数的名字必须以Test开头，可选的后缀名必须以大写字母开头，举几个例子：

```go
func TestAdd(t *testing.T){...}
func TestSum(t *testing.T){...}
func TestLog(t *testing.T){...}
```

其中参数t用于报告测试失败和附加的日志信息。testing.T的拥有的方法如下：

```go
func (c *T) Cleanup(func())
func (c *T) Error(args ...interface{})
func (c *T) Errorf(format string, args ...interface{})
func (c *T) Fail()
func (c *T) FailNow()
func (c *T) Failed() bool
func (c *T) Fatal(args ...interface{})
func (c *T) Fatalf(format string, args ...interface{})
func (c *T) Helper()
func (c *T) Log(args ...interface{})
func (c *T) Logf(format string, args ...interface{})
func (c *T) Name() string
func (c *T) Skip(args ...interface{})
func (c *T) SkipNow()
func (c *T) Skipf(format string, args ...interface{})
func (c *T) Skipped() bool
func (c *T) TempDir() string
```

### 单元测试示例

就像细胞是构成我们身体的基本单位，一个软件程序也是由很多单元组件构成的。单元组件可以是函数、结构体、方法和最终用户可能依赖的任意东西。总之我们需要确保这些组件是能够正常运行的。`单元测试是一些利用各种方法测试单元组件的程序`。它会将结果与预期输出进行比较。

Demo详见  ./0.base/Split.go. 

测试函数详见 ./0.base/splite_test.go

在 ./0.base 路径下执行 go test 命令，可以看到输出结果如下：

```go
$ go test
PASS
ok      beautifulGo/unit_test/0.base    0.001s

```

## go test -v

一个测试用例有些单薄，我们再编写一个测试使用多个字符切割字符串的例子，在split_test.go中添加如下测试函数：`TestSplitWithComplexSep(t *testing.T)`````

有了多个测试用例之后，为了能更好的在输出结果中看到每个测试用例的执行情况，我们可以为 go test命令添加 -v 参数，让它完整输出测试结果：

```go
$go test -v
=== RUN   TestSplit
--- PASS: TestSplit (0.00s)
=== RUN   TestSplitWithComplexSep
    split_test.go:20: expected:[a d],got:[a cd]
--- FAIL: TestSplitWithComplexSep (0.00s)
FAIL
exit status 1
FAIL    beautifulGo/unit_test/0.base    0.002s
```

从输出结果可以清楚的看到，`TestSplitWithComplexSep`这个测试用例并没有通过测试。

## go test -run

上述单元测试的结果表明，`split` 函数的实现并不可靠，没有考虑到传入的sep参数是多个字符的情况，下面我们来修复下这个bug：

./0.base/Split.go --fix bug

在执行 go test 命令的时候可以添加 -run 参数，它对应一个正则表达式，只有函数名匹配上的测试函数才会被 go test 命令执行。

例如给 go test 添加 -run=Sep（注意区分大小写）参数来告诉它本次测试只运行 `TestSplitWithComplexSep`这个测试用例：

```go
$ go test -run=Sep -v
=== RUN   TestSplitWithComplexSep
--- PASS: TestSplitWithComplexSep (0.00s)
PASS
ok      beautifulGo/unit_test/0.base    0.001s
```

## 回归测试

我们修改了代码（fix bug）之后，仅仅执行那些失败的测试用例或新引入的测试用例是错误且危险的，正确的做法应该是完整运行所有的测试用例，保证不会因为修改代码而引入新的问题。

```go
$ go test -v
=== RUN   TestSplit
--- PASS: TestSplit (0.00s)
=== RUN   TestSplitWithComplexSep
--- PASS: TestSplitWithComplexSep (0.00s)
PASS
ok      beautifulGo/unit_test/0.base    0.001s
```

测试结果表明我们的单元测试全部通过。

通过这个示例我们可以看到，有了单元测试就能够在代码改动后快速进行回归测试，极大的提高开发效率并保证代码的质量。

## 跳过某些测试用例

为了节省时间支持在单元测试时跳过某些耗时的测试用例。

```go
func TestTimeConsuming(t *testing.T) {
    if testing.Short() {
        t.Skip("short模式下会跳过该测试用例")
    }
    ...
}
```

当执行 go test -short 时就不会执行上面的 TestTimeConsuming 测试用例。

## 子测试

在上面的示例中我们为每一个测试数据编写了一个测试函数，而通常单元测试中需要多组测试数据保证测试的效果。Go1.7+中新增了子测试，支持在测试函数中使用t.Run执行一组测试用例，这样就不需要为不同的测试数据定义多个测试函数了。

```go
func TestXXX(t *testing.T){
  t.Run("case1", func(t *testing.T){...})
  t.Run("case2", func(t *testing.T){...})
  t.Run("case3", func(t *testing.T){...})
}
```

