package main

/*
方法名
sort.Search()

使用模板
index := sort.Search(n int,f func(i int) bool) int

该函数使用二分查找的方法，会从[0, n)中取出一个值index，index为[0, n)中最小的使函数f(index)为True的值，并且f(index+1)也为True。
如果无法找到该index值，则该方法为返回n。

常用场景
该方法一般用于从一个已经排序的数组中找到某个值所对应的索引。
或者从字符串数组中，找到满足某个条件的最小索引值，比如etcd中的键值范围查询就用到了该方法。
本项目中的算法部分，一致性hash算法的实现也用到了这个方法。
*/
