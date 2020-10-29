作者：黄韦嘉
邮箱：wjhuang@sjtu.edu.cn

功能：在RTCDemo中添加了用户名和房间号保存的功能。每次运行程序，界面显示上次输入的信息。

修改：
1、在CRTCDemoDlg::OnPaint()中添加读取数据功能
使用GetPrivateProfileString()函数从INI文件读取数据，并使用SetWindowTextW在用户界面上显示。

2、CRTCDemoDlg::OnBnClickedLogin()中添加数据写入功能
使用WritePrivateProfileString()函数将数据写入INI文件。