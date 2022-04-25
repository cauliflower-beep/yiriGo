package main

/*
多线程并发读写map时，由于map是非线程安全的，程序会直接报错，而且是Go源码调用 throw 方法所导致的致命错误，Go进程会中断。
错误提示：fatal error: concurrent map writes
这点与slice不一样，slice也是非线程安全的，但是多线程读写的时候不会报错，而是隐式的，多次运行结果会不一样。
如果想要map支持并发读写，可以对map进行上锁
 */
/*
参考链接：
https://zhuanlan.zhihu.com/p/412693892
 */