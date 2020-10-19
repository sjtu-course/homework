## Homework 1 易兴睿 Windows

1、采用`GetPrivateProfileString`和`WritePrivateProfileString`函数的调用来实现房间名与用户名的存取。

2、具体实现：

​	将`WritePrivateProfileString`在 `RTCDemoDlg.cpp`文件中的`OnBnClickedLogin`步骤中进行调用，实现触发Login按键后即保存设置。

​	将`GetPrivateProfileString`在 `RTCDemoDlg.cpp`文件中的`OnInitDialog`步骤中进行调用，实现登录界面初始化即读取预存数据。



