## ICMP
![](https://github.com/sjtu-course/homework/tree/main/YiyanWang/hw2/network/实验三/image/pic1.png)
![](https://github.com/sjtu-course/homework/tree/main/YiyanWang/hw2/network/实验三/image/pic2.png)
ICMP报文分为两种，即ICMP差错报告报文和ICMP询问报文  

ping使用了ICMP回送请求与回送回答报文，是应用层直接使用网络层ICMP的例子，并没有通过传输层UDP或TCP  

ICMP报文包括ICMP报文的类型和代码，这样源主机收到该报文就知道是什么问题导致的需要重传。但是ICMP的类型是不够的，需要将出错的数据报的首部也放在该ICMP报文中。
