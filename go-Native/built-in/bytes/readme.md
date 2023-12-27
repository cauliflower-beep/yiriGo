## 1.简介

### 1.1概念

bytes.buffer为go语言内置提供的一个 字节缓存区。

可以理解为一个集装箱，可以存取数据。

源码在 *src/bytes/buffer.go*

```go
// A Buffer 提供了可见的 读，写方法
// 零值是一个空字节slice, 但是可以使用
type Buffer struct {
    buf      []byte // contents are the bytes buf[off : len(buf)]
    off      int    // read at &buf[off], write at &buf[len(buf)]
    lastRead readOp // last read operation, so that Unread* can work correctly.
}
```

![image-20221031095153787](..\imgs\image-20221031095153787.png)

### 1.2原理

go字节缓冲区底层以字节切片做存储，切片存在长度len与容量cap, 缓冲区写从长度len的位置开始写，当len>cap时，会自动扩容。缓冲区读会从内置标记off位置开始读(off始终记录读的起始位置)，当off==len时，表明缓冲区已全部读完

并重置缓冲区(len=off=0),此外当将要内容长度+已写的长度(即len) <= cap/2时，缓冲区前移覆盖掉已读的内容(off=0，len-=off)，从避免缓冲区不断扩容

## 2.常用的方法

### 2.1申明一个buffer

```go
var b bytes.Buffer                    //申明一个Buffer变量
b := new(bytes.Buffer)                //使用New申明Buffer变量，new的时候是空的，但可以直接write
b := bytes.NewBuffer(s []byte)        //从一个[]byte切片，构造一个Buffer
b := bytes.NewBufferString(s string)  //从一个string变量，构造一个Buffer
```

可以通过 b.Grow来初始化一个 容量大小为n的buffer

```go
// Grow 增加一个n长度的 bytes, 当使用Grow(n)后，至少可以有b个自己，可以写入到buffer里面，如果n是一个负数，Grow会panic，如果buffer不能Grow，会导致panic然后报错 ErrTooLarge
b.Grow(n int)
```

使用案例

```go
func TestNew(t *testing.T) {
    b := bytes.Buffer{}
    b.Grow(100)
    t.Log(b) //空字符

    b1 := bytes.NewBufferString("hello")
    t.Log(b1) //hello

    b2 := bytes.NewBuffer([]byte{'h', 'e', 'l', 'l', 'o'})
    t.Log(b2) //hello
}
```

### 2.2写入一个buffer

```go
b.Write(d []byte) (n int, err error)               //将切片d写入Buffer尾部
b.WriteString(s string) (n int, err error)         //将字符串s写入Buffer尾部
b.WriteByte(c byte) error                          //将字符c写入Buffer尾部
b.WriteRune(r rune) (n int, err error)            //将一个rune写入Buffer尾部
```

使用案例

```go
func TestWrite(t *testing.T) {
    var b bytes.Buffer
    for i := 0; i < 10; i++ {
        b.WriteString("hello")
    }
    t.Log(b.String())
}
```

## 2.3读取一个buffer

```go
//从off偏移量开始读取len(p)长度的字节到p中
b.Read(p []byte) (n int, err error)

//读取下一个byte并返回，如果没有字节了，然后错误io.EOF
b.ReadByte() (byte, error)

//读取下一个 UTF8 编码的字符
b.ReadRune() (r rune, size int, err error)

//读取缓冲区第一个分隔符delim内容，并且返回从off到分隔符delim的内容
b.ReadBytes(delim byte) (line []byte, err error)

//读取缓冲区第一个分隔符delim内容，并且返回从off到分隔符delim的内容
b.ReadString(delim byte) (line string, err error)
```

使用案例

```go
func TestRead(t *testing.T) {
    var b bytes.Buffer
    for i := 0; i < 10; i++ {
        b.WriteString("hello")
    }
    //读取到第一个'o'之间的字符
    s, err := b.ReadString('o')
    t.Log(s, err) //hello <nil>
}
```

### 2.4读入到缓存区/从缓存区写入

buffer 缓冲器的实现原理就是，将文件读取进缓冲（内存）之中，再次读取的时候就可以避免文件系统的 I/O 从而提高速度。同理在进行写操作时，先把文件写入缓冲（内存），然后由缓冲写入文件系统。

```go
//从实现了io.Reader接口的可读取对象写入Buffer尾部
b.ReadFrom(r io.Reader) (n int64, err error)    

//将 Buffer 中的内容输出到实现了 io.Writer 接口的可写入对象中，成功返回写入的字节数，失败返回错误
b.WriteTo(w io.Writer) (n int64, err error)
```

使用案例

```go
//写入到文件中
func TestWriteTo(t *testing.T) {
    b := bytes.NewBufferString("hello,world")
    f, _ := os.Create("hi.txt")
    defer f.Close()
    _, _ = b.WriteTo(f)
}

//从文件中读取
func TestReadFrom(t *testing.T) {
    b := bytes.NewBufferString("")
    f, _ := os.Open("hi.txt")
    defer f.Close()
    _, _ = b.ReadFrom(f)
    t.Log(b.String())   //hello,world
}
```

