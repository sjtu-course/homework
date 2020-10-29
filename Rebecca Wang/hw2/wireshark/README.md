实验一：UDP报文格式抓包分析和IP层协议格式分析

首先开始抓包，之后以UDP为过滤器进行筛选，得到如下图所示的结果：
![Image text](https://raw.githubusercontent.com/sjtu-course/homework/main/Rebecca%20Wang/hw2/wireshark/photo/QQ%E5%9B%BE%E7%89%8720201028202615.png)
UDP用户数据的报文格式如图所示：
![Image text](https://raw.githubusercontent.com/sjtu-course/homework/main/Rebecca%20Wang/hw2/wireshark/photo/)(R%7D~S0F%24FG4YLVWIUR8O4V.png)
选中其中的一个，得到的报文结果如下图所示：
![Image text](https://raw.githubusercontent.com/sjtu-course/homework/main/Rebecca%20Wang/hw2/wireshark/photo/QNP3K(HT3X1%5DJWLWTOH%24%400J.png)
其中源端口、目标端口、长度和检验和已经在图中标出。


实验二：任意网站三次握手过程和四次挥手过程抓包分析

TCP网络具有三次握手过程和四次挥手过程，打开一个网站，IP接口为116.128.163.67。由此筛选出全部TCP结果，如下图所示：
![Image text](https://raw.githubusercontent.com/sjtu-course/homework/main/Rebecca%20Wang/hw2/wireshark/photo/TCP.png)
TCP的报文格式如下图所示：
![Image text](https://raw.githubusercontent.com/sjtu-course/homework/main/Rebecca%20Wang/hw2/wireshark/photo/TCP原理.png)
三次握手的SYN、SYN-ACK和ACK分别如下图所示：
![Image text](https://raw.githubusercontent.com/sjtu-course/homework/main/Rebecca%20Wang/hw2/wireshark/photo/三次握手SYN.png)
![Image text](https://raw.githubusercontent.com/sjtu-course/homework/main/Rebecca%20Wang/hw2/wireshark/photo/三次握手SYN-ACK.png)
![Image text](https://raw.githubusercontent.com/sjtu-course/homework/main/Rebecca%20Wang/hw2/wireshark/photo/三次握手ACK.png)
四次挥手的FIN、ACK、FIN和ACK分别如下图所示：
![Image text](https://raw.githubusercontent.com/sjtu-course/homework/main/Rebecca%20Wang/hw2/wireshark/photo/四次挥手FIN1.png)
![Image text](https://raw.githubusercontent.com/sjtu-course/homework/main/Rebecca%20Wang/hw2/wireshark/photo/四次挥手ACK1.png)
![Image text](https://raw.githubusercontent.com/sjtu-course/homework/main/Rebecca%20Wang/hw2/wireshark/photo/四次挥手FIN2.png)
![Image text](https://raw.githubusercontent.com/sjtu-course/homework/main/Rebecca%20Wang/hw2/wireshark/photo/四次回收ACK2.png)
其中源端口、目标端口、序号、确认号、保留、窗口、检验和等都在图中显示出。


实验三：分析ping下的ICMP报文格式分析

以www.baidu.com 为例，进行此网站的ICMP报文抓取，首先在cmd中使用ping命令
![Image text](https://raw.githubusercontent.com/sjtu-course/homework/main/Rebecca%20Wang/hw2/wireshark/photo/ping命令.png)
得到ICMP的所有结果
![Image text](https://raw.githubusercontent.com/sjtu-course/homework/main/Rebecca%20Wang/hw2/wireshark/photo/ICMP.png)
选择一个打开其报文和他的IP首部：
![Image text](https://raw.githubusercontent.com/sjtu-course/homework/main/Rebecca%20Wang/hw2/wireshark/photo/ICMP报文.png)
![Image text](https://raw.githubusercontent.com/sjtu-course/homework/main/Rebecca%20Wang/hw2/wireshark/photo/IP首部.png)
