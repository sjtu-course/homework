## 网络作业

本次共三次实验，分别是

- [x] [实验一: UDP抓包](./实验一)
- [x] [实验二: 三次握手四次挥手抓包](实验二)
- [x] [实验三: ICMP抓包](实验三)

<hr/>

在抓包过程中遇到了一些问题，比如某次抓包过程中分析四次挥手的时候，发现一直没有收到最后一个ACK包，类似下面这张图：

<img src="./wave.png">

从图中我们可以发现client(10.162.41.74)向server(182.61.200.6)发送了一个FIN包，server接收到了这个FIN包，并且返回ACK包进行确认，过了一小段时间，然后server发送FIN包给client，但是我们可以从图中看出，client并没有回复一个ACK=Seq+1的包给server，而是返回了一个RST包，并且ack为FIN包的seq，和明显感觉出和四次挥手不同。

经过老师的解答，造成这个原因的一个因素可能是TCP连接的复用。当最后server向本地发送FIN包之后，本地进行接收，但是这个连接没有选择关闭，而是将这个连接进行reset，然后发送一个ACK=Seq的RST包给server。

在后面又经过一些资料[^1]的查找，还有下面的一些可能：

- 连接的提前关闭

  > 在这个案例中，可能是本地TCP连接提前关闭，所以当server发送FIN包过来的时候，本地时间上没有对应的TCP连接，所以发送一个RST包给server

- 用于效率的提高（maybe）

  > 查找网上资源发现，使用RST包关闭连接可以快速释放已经完成数据交互的TCP连接，以提高业务交互的效率。
  >
  > 不过这个可能性感觉比较低，因为是这个RST是最后一次应该发送ACK的时候发送的，网上资源图片[^2]上的应该表示的是发送全部数据之后发送RST包

[^1]: https://blog.csdn.net/hik_zxw/article/details/50167703
[^2]:https://img-my.csdn.net/uploads/201512/03/1449155451_3152.png
