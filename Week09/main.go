package main

import (
	"context"
	"fmt"
	"net"
)

func main() {
	// 服务端端口
	service := ":5000"
	// 绑定
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {
		panic(err)
	}
	// 监听
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}
	for {
		// 接受
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)

			continue
		}

		// 创建 Goroutine
		go handleClient(context.Background(), conn)
	}

}
func handleClient(ctx context.Context, conn net.Conn) {
	// 逆序调用 Close() 保证连接能正常关闭
	defer func() {
		conn.Close()
		ctx.Done()
	}()

	var buf [512]byte
	for {
		// 接收数据
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		rAddr := conn.RemoteAddr()
		fmt.Println("Receive from client", rAddr.String(), string(buf[0:n]))

		go write2Client(ctx, conn)
	}
}

func write2Client(ctx context.Context, conn net.Conn) {
	for {
		select {
		case <-ctx.Done():
			conn.Close()
			return
		default:
			defer conn.Close()
			var b = make([]byte, 512)
			fmt.Scan(&b)
			_, _ = conn.Write(b)
		}

	}

}
