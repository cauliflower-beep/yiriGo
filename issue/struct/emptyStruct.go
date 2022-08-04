package _struct

/*
空结构体的用途：
1.go中没有set，不像python，如何利用map来实现一个set？（面试题）
	map和set是两个抽象的数据结构，map存储一个键值对集合，其中键不重复，set存储一个不重复的元素集合。
	本质上set可以视为一种特殊的map，set其实就是map中的键。
	用map模拟一个set，就要把值置为struct{},struct{}本身是不占任何空间的，可以避免任何多余的内存分配；
2.有时候给通道发送一个空结构体，channel<-struct{}{}，也是为了节省空间
3.仅有方法的结构体
*/
