## 概述
对于一个项目来说，一个功能的实现往往是需要通过庞大的代码来完成的，而对于程序来说，性能是必不可少的关注话题。

因此，在如此庞大的代码量中分析性能问题是非常重要的。

## Go测试种类

golang支持的测试种类有：

| 类型                           | 格式                  | 作用                                                 |
| ------------------------------ | --------------------- | ---------------------------------------------------- |
| 单元测试                       | 函数名前缀为Test      | 测试程序的一些逻辑行为是否正确                       |
| 基准（压力测试）               | 函数名前缀为Benchmark | 测试函数的性能                                       |
| 示例测试                       | 函数名前缀为Example   | 为文档提供示例文档                                   |
| 模糊（随机）测试（高版本才有） | 函数名前缀为Fuzz      | 生成一个随机测试用例，去覆盖人为测不到的各种复杂场景 |

### benchmark基准测试

#### 1.前言

优化代码性能之前，首先要了解当前性能怎么样。

Go 语言标准库内置的 testing 测试框架提供了基准测试(benchmark)的能力，能让我们很容易地对某一段代码进行性能测试。

但需要注意的是，性能测试受环境的影响很大，为了保证测试的可重复性，在进行性能测试时，尽可能地保持测试环境的稳定：

- 机器处于闲置状态，测试时不要执行其他任务，也不要和其他人共享硬件资源；
- 机器是否关闭了节能模式，一般笔记本会默认打开这个模式，测试时关闭；
- 避免使用虚拟机和云主机进行测试。一般情况下，为了尽可能提高资源的利用率，虚拟机和云主机 CPU 和内存会进行超分配，超分机器的性能表现会非常地不稳定。

> 超分配是针对硬件资源来说的，商业上对应的就是云主机的超卖。
>
> 虚拟化技术带来的最大直接收益是服务器整合，通过 CPU、内存、存储、网络的超分配（Overcommitment）技术，最大化服务器的使用率。
>
> 例如，虚拟化的技能之一就是随心所欲的操控 CPU，例如一台 32U(物理核心)的服务器可能会创建出 128 个 1U(虚拟核心)的虚拟机，当物理服务器资源闲置时，CPU 超分配一般不会对虚拟机上的业务产生明显影响，但如果大部分虚拟机都处于繁忙状态时，那么各个虚拟机为了获得物理服务器的资源就要相互竞争，相互等待。
>
> Linux 上专门有一个指标，Steal Time(st)，用来衡量被虚拟机监视器(Hypervisor)偷去给其它虚拟机使用的 CPU 时间所占的比例。

#### 2.测试命令

`go test <module name>/<package name>` 用来运行某个 package 内的所有测试用例：

- 运行当前 package 内的用例：`go test example` 或 `go test .`
- 运行子 package 内的用例： `go test example/<package name>` 或 `go test ./<package name>`
- 如果想递归测试当前目录下的所有的 package：`go test ./...` 或 `go test example/..`

`go test` 命令默认不运行 benchmark 用例的，如果我们想运行 benchmark 用例，则需要加上 `-bench` 参数。例如:

```go
$ go test -bench .
goos: darwin
goarch: amd64
pkg: example
BenchmarkFib-8               200           5865240 ns/op
PASS
ok      example 1.782s
```

`-bench` 参数支持传入一个正则表达式，匹配到的用例才会得到执行，例如，只运行以 `Fib` 结尾的 benchmark 用例：

```go
$ go test -bench='Fib$' .
goos: darwin
goarch: amd64
pkg: example
BenchmarkFib-8               202           5980669 ns/op
PASS
ok      example 1.813s
```

#### 3.benchmark是如何工作的

benchmark 用例的参数 `b *testing.B`，有个属性 `b.N` 表示这个用例需要运行的次数。`b.N` 对于每个用例都是不一样的。

那这个值是如何决定的呢？`b.N` 从 1 开始，如果该用例能够在 1s 内完成，`b.N` 的值便会增加，再次执行。`b.N` 的值大概以 1, 2, 3, 5, 10, 20, 30, 50, 100 这样的序列递增，越到后面，增加得越快。

运行 ./benchmark/fib_test.go 中的用例，观察输出:

```go
BenchmarkFib-12              291           3993668 ns/op
```

BenchmarkFib-12 中的 `-12` 即 `GOMAXPROCS`，默认等于 CPU 核数。可以通过 `-cpu` 参数改变 `GOMAXPROCS`，`-cpu` 支持传入一个列表作为参数，例如：

```go
$ go test -bench='Fib$' -cpu=2,4 .
goos: darwin
goarch: amd64
pkg: example
BenchmarkFib-2               206           5774888 ns/op
BenchmarkFib-4               205           5799426 ns/op
PASS
ok      example 3.563s
```

在这个例子中，改变 CPU 的核数对结果几乎没有影响，因为这个 Fib 的调用是串行的。

`291` 和 `3993668 ns/op` 表示用例执行了 291 次，每次花费约 0.004s。总耗时比 1s 略多。

#### 4.提升准确度

对于性能测试来说，提升测试准确度的一个重要手段就是增加测试的次数。我们可以使用 `-benchtime` 和 `-count` 两个参数达到这个目的。

- `-benchtime`

  benchmark 的默认时间是 1s，那么我们可以使用 `-benchtime` 指定为 5s。例如：

  ```go
  $ go test -bench='Fib$' -benchtime=5s .
  goos: windows
  goarch: amd64                                
  pkg: yiriGo/go-Native/pkg/test/benchmark     
  cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
  BenchmarkFib-12             1489           4133421 ns/op
  PASS
  ok      yiriGo/go-Native/pkg/test/benchmark     6.594s
  ```

  实际执行的时间是 6.5s，比 benchtime 的 5s 要长，测试用例编译、执行、销毁等是需要时间的。

  将 `-benchtime` 设置为 5s，用例执行次数也变成了原来的 5倍，每次函数调用时间仍为 0.004s，几乎没有变化。

  `-benchtime` 还可以指定具体执行的次数。例如，执行 30 次可以用 `-benchtime=30x`，可参考上述修改测试时间的方式运行测试。

- `-count`

  这个参数可以用来设置 benchmark 执行的轮数。例如，进行 3 轮 benchmark：

  ```go
  go test -bench='Fib$' -benchtime=5s -count=3 .
  ```

#### 5.内存分配

`-benchmem` 参数可以度量内存分配的次数。内存分配次数与性能也是息息相关的，例如不合理的切片容量，将导致内存重新分配，带来不必要的开销。

参考 .\benchmark\generate_test.go 文件中的测试用例。

测试结果分析：

```go
goos: windows
goarch: amd64
pkg: yiriGo/go-Native/pkg/test/benchmark
cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
BenchmarkGenerateWithCap-12          373          15916586 ns/op         8003589 B/op          1 allocs/op
BenchmarkGenerate-12                 301          21310610 ns/op        45188480 B/op         41 allocs/op
PASS
ok      yiriGo/go-Native/pkg/test/benchmark     16.015s
```

`Generate` 分配的内存是 `GenerateWithCap` 的 6 倍，设置了切片容量，内存只分配一次，而不设置切片容量，内存分配了 40 次。

#### 6.测试不同的输入

不同的函数复杂度不同，O(1)，O(n)，O(n^2) 等，利用 benchmark 验证复杂度一个简单的方式，是构造不同的输入。

参考 .\benchmark\complex_test.go 文件中的测试用例。

测试结果分析：

```go
goos: windows
goarch: amd64
pkg: yiriGo/go-Native/pkg/test/benchmark
cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
BenchmarkGenerate1000-12           45516             26033 ns/op
BenchmarkGenerate10000-12           6000            200992 ns/op
BenchmarkGenerate100000-12           588           1999295 ns/op
BenchmarkGenerate1000000-12           60          20085595 ns/op
PASS
ok      yiriGo/high_performance_go/1_performanceAnalysis/1_benchmark    5.486s
```

通过测试结果可以发现，输入变为原来的 10 倍，函数每次调用的时长也差不多是原来的 10 倍，这说明复杂度是线性的。

#### 7.注意事项

- ResetTimer

  在 benchmark 开始前，如果需要一些准备工作且比较耗时，则需要将这部分代码的耗时忽略掉。

  参考 .\benchmark\warn_test.go 文件中的BenchmarkFib2测试用例。

  测试结果分析：

  ```go
  goos: windows
  goarch: amd64
  pkg: yiriGo/go-Native/pkg/test/benchmark
  cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
  BenchmarkFib2-12              50          64579358 ns/op
  PASS
  ok      yiriGo/go-Native/pkg/test/benchmark     6.277s
  ```

  50次调用，每次调用约 0.064s，是之前的 0.004s 的16倍。究其原因，受到了耗时准备任务的干扰。我们需要用 `ResetTimer` 屏蔽掉。

- StopTimer & StartTimer

  如果每次函数调用前后都需要一些准备工作和清理工作，可以使用 `StopTimer` 暂停计时，以及 `StartTimer` 开始计时。

  参考 .\benchmark\warn_test.go 文件中的BenchmarkBubbleSort测试用例。

## 日志方法

单元测试框架提供了以下常用的日志方法：

| 方法   | 备注                             |
| ------ | -------------------------------- |
| Log    | 打印日志，同时结束测试           |
| Logf   | 格式化打印日志，同时结束测试     |
| Error  | 打印错误日志，同时结束测试       |
| Errorf | 格式化打印错误日志，同时结束测试 |
| Fatal  | 打印致命日志，同时结束测试       |
| Fatalf | 格式化打印致命日志，同时结束测试 |

## 子测试

自测试是 Go 语言内置支持的。可以在某个测试用例中，根据测试场景使用 t.Run 创建不同的子测试用例。

```go
// 使用案例:
TestMul()
TestDiv()
```



## 辅助函数

对于一些重复的逻辑，抽取出来作为公共的帮助函数（helpers），可以增加测试代码的可读性和可维护性。

借助帮助函数，可以使测试用例的主逻辑看起来更清晰。

```go
// 使用案例：
// 我们可以将创建子测试的逻辑抽取出来
createSubTestCase()
```

关于helper函数的两个建议：

1. 不要返回错误，帮助函数内部直接使用 t.Error 或 t.Fatal 即可，在用例主逻辑中不会因为太多的错误处理代码，影响可读性。
2. 调用 t.Helper() 让报错信息更准确，有助于定位。

## setup 和 teardown

如果在同一个测试文件中，每一个测试用例运行前后的逻辑是相同的，一般会写在 setup 和 teardown 函数中。例如执行前需要实例化待测试的对象，如果这个对象比较复杂，很适合将这一部分逻辑提取出来；执行后，可能会做一些资源回收类的工作。

例如关闭网络连接，释放文件等。

标准库 testing 提供了这样的机制。

使用案例见 setup_test.go

## 网络测试

### TCP/IP

假设需要测试某个 API 接口的 handler 能否正常工作，例如 helloHandler

案例见 net_test.go

可以创建真实的网络链接进行测试