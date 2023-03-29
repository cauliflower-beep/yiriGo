package main

import (
	"fmt"
	"net"
)

/*
	学习 aurora 框架的时候，第一节服务器接口需要启动一个tcp监听器
	查阅资料发现有如下两种方式可以创建：
*/

// createTCPListener1
//  @Description: 更加灵活，可以支持更多的地址格式，也更为严谨可靠
func createTCPListener1(IPVersion, IP string, Port int) {
	/*
		这里使用net.ResolveTCPAddr函数来解析TCP地址，可以解析IP和端口号并返回一个TCPAddr类型的对象，这个对象可以用于创建一个TCP监听器
		会进行一些额外的验证和转换操作，确保返回的TCPAddr对象是有效的
	*/
	addr, err := net.ResolveTCPAddr(IPVersion, fmt.Sprintf("%s:%d", IP, Port))
	if err != nil {
		fmt.Println("resolve tcp addr err: ", err)
		return
	}

	listener, err := net.ListenTCP(IPVersion, addr) //监听服务器地址
	if err != nil {
		panic(err)
	}

	fmt.Println("start Aurora server succ, now listenning...") //已经监听成功

	//启动server网络连接业务
	for {
		// 阻塞等待客户端建立连接请求,如果有客户端链接过来，阻塞会返回
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err ", err)
			continue
		}
		fmt.Println("Get conn remote addr = ", conn.RemoteAddr().String())

		//链接建立，做一个最基本的内容回显业务
		go func() {
			for {
				buf := make([]byte, 512)
				cnt, err := conn.Read(buf)
				if err != nil {
					fmt.Println("recv buf err", err)
					continue
				}
				//若读取成功，则回显
				if _, err := conn.Write(buf[:cnt]); err != nil {
					fmt.Println("write back buf err", err)
				}
			}
		}()
	}
}

// createTCPlISTENER2
//  @Description: 优化建议问 cursor
//  @Description: 更加简洁
func createTCPlISTENER2(ip string, port int) {
	/*
		直接使用net.Listen函数来创建TCP监听器，它只需要传入一个字符串形式的地址和端口号，不需要先解析地址。
		如果字符串格式不正确，可能会导致TCP监听器创建失败
	*/
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port)) //监听服务器地址
	if err != nil {
		panic(err)
	}

	fmt.Println("start Aurora server succ, now listenning...") //已经监听成功

	//启动server网络连接业务
	for {
		// 阻塞等待客户端建立连接请求,如果有客户端链接过来，阻塞会返回
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept err ", err)
			continue
		}
		fmt.Println("Get conn remote addr = ", conn.RemoteAddr().String())

		//链接建立，做一个最基本的内容回显业务
		go func() {
			for {
				buf := make([]byte, 512)
				cnt, err := conn.Read(buf)
				if err != nil {
					fmt.Println("recv buf err", err)
					continue
				}
				//若读取成功，则回显
				if _, err := conn.Write(buf[:cnt]); err != nil {
					fmt.Println("write back buf err", err)
				}
			}
		}()
	}
}

func main() {

}
