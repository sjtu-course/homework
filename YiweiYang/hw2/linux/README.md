# Golang综合实战作业
## [demo项目](https://github.com/sjtu-course/homework/tree/main/YiweiYang/hw2/linux/demo)
最开始理解错了作业的意思，想着把老师给的demo获取的response解析一下，踩了不少坑，特别是接口的方面，把踩的坑写了一个[博客](https://www.cnblogs.com/smileglaze/p/13893794.html)。

运行`main.go`即可获得解析的内容，如：

	IP: 202.120.38.248
	Port: 2967
	ID: 8rBc1gGDNQ4i_FBnP13r
	Room ID: roomid
	Name: rtcname0

## [LocalIO项目](https://github.com/sjtu-course/homework/tree/main/YiweiYang/hw2/linux/LocalIO)

- [x] 从config.json配置读取IP，port，ID，roomID，name信息
- [x] 命令行传入config.json参数，即 ./main -config config.json
- [x] 将日志写入文件系统，chat.log

**运行方式**
`go run main.go -config config.json`

**输出**

	IP: 127.0.0.1
	Port: 8848
	ID: KHKLJPHO
	Room ID: 7749
	Name: Yiwei Yang

**chat.log**

	[log]2020/11/03 13:07:43 main.go:67: IP: 127.0.0.1
	[log]2020/11/03 13:07:43 main.go:68: Port: 8848
	[log]2020/11/03 13:07:43 main.go:69: ID: KHKLJPHO
	[log]2020/11/03 13:07:43 main.go:70: Room ID: 7749
	[log]2020/11/03 13:07:43 main.go:71: Name: Yiwei Yang
