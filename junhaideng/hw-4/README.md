### 四人视频

主要修改`RtcWinDemo.cpp`文件中的`AllocateWindow`方法以及`RemoveTrack`方法，可以保证多人视频时，有人退出聊天室，那么视频窗口可以移动，使得视频窗口连续的，中间不会穿插其他的空视频窗口

另外实现了信息发送的功能，比较类似于`JoinRoom`一类的方法，直接在调试台`network`中的`ws`中找到对应需要发送的数据格式即可

