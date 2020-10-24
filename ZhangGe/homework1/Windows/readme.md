关键函数如下。登陆页面初始化时，读取`Config.int`文件中`room`和`user`信息并显示在界面上(初次登陆时为空)。当用户点击登陆按钮时，输入的`room`和`user`信息被保存在`Config.int`文件中，在下次登陆页面初始化时自动读取。
```c++

BOOL CRTCDemoDlg::OnInitDialog()
{
	CDialogEx::OnInitDialog();

	...
	// TODO: Add extra initialization here
	//======== load the room and user info saved last time ======== //
	loadRoomUser();

	return TRUE;  // return TRUE  unless you set the focus to a control
}

// ======== load the room and user info saved last time ======== //
void CRTCDemoDlg::loadRoomUser() 
{
	LPTSTR configPath = _T("..\\Config.ini");
	LPTSTR room = new TCHAR[100];
	LPTSTR user = new TCHAR[100];

	GetPrivateProfileStringW(_T("registerInfo"), _T("room"), _T(" "), room, 10, configPath);
	GetPrivateProfileStringW(_T("registerInfo"), _T("user"), _T(" "), user, 10, configPath);

	m_room.SetWindowText(room);
	m_name.SetWindowText(user);

}

// ======== save the room and user info when click button login ======== //
void CRTCDemoDlg::saveRoomUser(CString room, CString user)
{
	LPTSTR PATH = _T("..\\Config.ini");

	m_room.GetWindowText(room);
	m_name.GetWindowText(user);

	WritePrivateProfileStringW(_T("registerInfo"), _T("room"), room, PATH);
	WritePrivateProfileStringW(_T("registerInfo"), _T("user"), user, PATH);
}

void CRTCDemoDlg::OnBnClickedLogin()
{
	CString room_name, user_name;
	m_room.GetWindowText(room_name);
	m_name.GetWindowText(user_name);
	// ======== save the room and user info when click button login ======== //
	saveRoomUser(room_name, user_name);

	CChatDialog chat_dlg;
	chat_dlg.Init(user_name, room_name);
	chat_dlg.DoModal();
}

```

