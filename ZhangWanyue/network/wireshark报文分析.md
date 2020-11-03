#### UDP报文分析

```shell
go run udpS.go 1234
go run udpC.go 127.0.0.1:1234
```

从udpC的终端发送消息，可抓包如下图：

![img0](/Users/winnieaz/Documents/techNotebook/homework/ZhangWanyue/network/img0.png)

客户端发送消息到服务端，源端口63765，目的端口1234，校验和0xfe29。

![img01](/Users/winnieaz/Documents/techNotebook/homework/ZhangWanyue/network/img01.png)服务端返回请求源端口号1234，目的端口号54321.

#### TCP报文分析

1. 三次握手![img1](/Users/winnieaz/Documents/techNotebook/homework/ZhangWanyue/network/img1.png)

   第一次，请求客户端发送一个 [SYN] 包，序列号是 0（seq=0）。

   第二次，服务器收到 [SYN] 包，然后会发送一个 [SYN, ACK] 包，序列号是 0，ACK 置 1（seq=0，ack=1）。

   第三次，客户端请求服务器，客户端会发送一个 [ACK] 包，序列号是 1，Ack 置 1（seq=1，ack=1）来回复服务器。

2. 四次挥手

![img2](/Users/winnieaz/Documents/techNotebook/homework/ZhangWanyue/network/img2.png)

​	服务端发起关闭连接请求，

​	第一次挥手，发送FIN和ACK,ack = 680, seq = 910

​	第二次挥手，客户端发送ACK,ack = 910, seq = 680

​	第三次挥手，客户端发送ACK,ack = 911, seq = 680

​	第四次挥手，客户端发送ACK,ack = 681, seq = 911

#### ping下的ICMP报文分析

```
ping www.baidu.com
```

wireshark抓包结果如图：![img4](/Users/winnieaz/Documents/techNotebook/homework/ZhangWanyue/network/img4.png)

ICMP的协议号为1，报文包括类型、代码、校验和、序列号。