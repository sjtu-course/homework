# Windows



---

1.使用`GetModuleFileName`获取当前目录（实际获取的是当前目录下exe文件的绝对路径）
2.在每次按下登陆的时候使用`WritePrivateProfileString`将房间号与姓名写入当前目录下的ini文件
3.在每次初始化窗口的时候使用`GetPrivateProfileString`读取当前目录下的ini文件




