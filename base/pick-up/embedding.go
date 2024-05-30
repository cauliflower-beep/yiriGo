package main

/* ----类型嵌入----*/
/*
	本文讨论的是golang中的嵌入类型(embedding types)，有时候也被叫做嵌入式字段(embedding fields)

	涉及到为什么使用嵌入类型，以及嵌入类型的一些“坑”。
*/

/*----1.什么是嵌入类型----*/
type FileSystem struct {
	MetaData []byte // 元数据
}

func (fs *FileSystem) Read()  {}
func (fs *FileSystem) Write() {}

type NTFS struct {
	*FileSystem
}

type EXT4 struct {
	*FileSystem
}

/*
	有一个FileSystem类型作为对文件系统的抽象，其中包含了所有文件系统都会存在的元数据和读写文件的方法.
	接着基于此定义了Windows的NTFS文件系统和广泛应用于Linux系统中的EXT4文件系统。其中的*FileSystem就是一个嵌入类型的字段

	一个更严谨的解释是：如果一个字段只含有字段类型而没有指定字段的名字，那么这个字段就是一个嵌入类型字段。
*/

/*----2.嵌入类型的使用----*/
/* 2.1嵌入类型字段引用
嵌入类型只有类型名而没有字段名，那么我们怎么引用它呢？

答案是嵌入类型字段的类型名会被当成该字段的名字。继续刚才的例子，想要在NTFS中引用FileSystem的函数，则需要这样写：
	ntfs.FileSystem.Read() // ntfs 是一个已经初始化了的NTFS实例
要注意，指针的*只是类型修饰符，并不是类型名的一部分，所以对于形如*Type和Type的嵌入类型，我们都只能通过Type这个名字进行引用。

通过Type这个名字，我们不仅可以引用Type里的方法，还可以引用其中的数据字段：
type A struct {
    Age int
    Name string
}

type B struct {
    A
}

b := B{}
fmt.Println(b.A.Age, b.A.Name)
*/
/* 2.2嵌入类型的初始化
嵌入类型字段只是普通的匿名字段，你可以放在类型的任意位置，也就是说嵌入类型可以不必作为类型的第一个字段：
type A struct {
    a int
    b int
}

type B struct {
    *A
    name string
}

type C struct {
    age int
    B
    address string
}
B和C都是合法的，如果想要初始化B和C，则只需要按字段出现的顺序给出相应的初始化值即可：
// 初始化B和C
b := &B{
    &A{1, 2},
    "B",
}
c := &C{
    30,
    B{
        &A{1, 2},
        "B in C",
    },
    "my address",
}
由于我们还可以使用对应的类型名来引用嵌入类型字段，所以初始化还可以写成这样：
// 使用字段名称初始化B和C
b := &B{
    A: &A{1, 2},
    name: "B",
}

c := &C{
    age: 30,
    B: B{
        A: &A{1, 2},
        name: "B in C",
    },
    address: "my address",
}
https://www.cnblogs.com/apocelipes/p/14090671.html#%E5%B5%8C%E5%85%A5%E7%B1%BB%E5%9E%8B%E7%9A%84%E5%AD%97%E6%AE%B5%E6%8F%90%E5%8D%87
*/
