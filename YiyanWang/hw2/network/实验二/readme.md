## TCP
* 三次握手
![](https://github.com/sjtu-course/homework/blob/main/YiyanWang/hw2/network/%E5%AE%9E%E9%AA%8C%E4%BA%8C/image/pic1.PNG)
* 四次挥手
![](https://github.com/sjtu-course/homework/blob/main/YiyanWang/hw2/network/%E5%AE%9E%E9%AA%8C%E4%BA%8C/image/pic2.PNG)
服务端主动发FIN表示它的数据已经全部发完，进入被动关闭状态。这和主动关闭有一定的区别。
1. 服务器发送FIN到客户端
2. 客户端收到FIN，关闭读通道，设置自己状态为TIME_WAIT，发送一个ACK给服务器
3. 服务器收到ACK，关闭写通道，并将自己的状态设为CLOSE
4. 客户端等待两个最大传输时间，然后将自己状态设置成CLOSED
![](https://github.com/sjtu-course/homework/blob/main/YiyanWang/hw2/network/%E5%AE%9E%E9%AA%8C%E4%BA%8C/image/pic3.PNG)
![](https://github.com/sjtu-course/homework/blob/main/YiyanWang/hw2/network/%E5%AE%9E%E9%AA%8C%E4%BA%8C/image/pic4.PNG)

* 三次握手：
1.	SYN = 1 seq = 4283929355
2.	SYN = 1 ACK = 1 seq = 405355663 ack = 4283929356
3.	ACK = 1 seq = 4283929356 ack = 405355664
