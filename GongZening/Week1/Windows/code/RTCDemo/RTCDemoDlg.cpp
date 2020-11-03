
// RTCDemoDlg.cpp : implementation file
//

#include "pch.h"
#include "framework.h"
#include "RTCDemo.h"
#include "RTCDemoDlg.h"
#include "afxdialogex.h"
#include "CChatDialog.h"
#include "resource.h"
#include "string"

#ifdef _DEBUG
#define new DEBUG_NEW
#endif

CString cur_path;
CFileFind finder;
std::string ini_path;

// CAboutDlg dialog used for App About

class CAboutDlg : public CDialogEx
{
public:
	CAboutDlg();

// Dialog Data
#ifdef AFX_DESIGN_TIME
	enum { IDD = IDD_ABOUTBOX };
#endif

	protected:
	virtual void DoDataExchange(CDataExchange* pDX);    // DDX/DDV support

// Implementation
protected:
	DECLARE_MESSAGE_MAP()
};

CAboutDlg::CAboutDlg() : CDialogEx(IDD_ABOUTBOX)
{
}

void CAboutDlg::DoDataExchange(CDataExchange* pDX)
{
	CDialogEx::DoDataExchange(pDX);
}

BEGIN_MESSAGE_MAP(CAboutDlg, CDialogEx)
END_MESSAGE_MAP()


// CRTCDemoDlg dialog



CRTCDemoDlg::CRTCDemoDlg(CWnd* pParent /*=nullptr*/)
	: CDialogEx(IDD_RTCDEMO_DIALOG, pParent)
{
	m_hIcon = AfxGetApp()->LoadIcon(IDR_MAINFRAME);
}

void CRTCDemoDlg::DoDataExchange(CDataExchange* pDX)
{
	CDialogEx::DoDataExchange(pDX);
	DDX_Control(pDX, IDC_NAME, m_room);
	DDX_Control(pDX, IDC_USER, m_name);
	DDX_Control(pDX, IDC_CHECK1, CCButton);
}

BEGIN_MESSAGE_MAP(CRTCDemoDlg, CDialogEx)
	ON_WM_SYSCOMMAND()
	ON_WM_PAINT()
	ON_WM_QUERYDRAGICON()
	ON_BN_CLICKED(IDC_LOGIN, &CRTCDemoDlg::OnBnClickedLogin)
	ON_BN_CLICKED(IDC_CHECK1, &CRTCDemoDlg::OnBnClickedCheck1)
	ON_EN_CHANGE(IDC_NAME, &CRTCDemoDlg::OnEnChangeName)
END_MESSAGE_MAP()


// CRTCDemoDlg message handlers

BOOL CRTCDemoDlg::OnInitDialog()
{
	CDialogEx::OnInitDialog();

	// Add "About..." menu item to system menu.

	// IDM_ABOUTBOX must be in the system command range.
	ASSERT((IDM_ABOUTBOX & 0xFFF0) == IDM_ABOUTBOX);
	ASSERT(IDM_ABOUTBOX < 0xF000);

	CMenu* pSysMenu = GetSystemMenu(FALSE);
	if (pSysMenu != nullptr)
	{
		BOOL bNameValid;
		CString strAboutMenu;
		bNameValid = strAboutMenu.LoadString(IDS_ABOUTBOX);
		ASSERT(bNameValid);
		if (!strAboutMenu.IsEmpty())
		{
			pSysMenu->AppendMenu(MF_SEPARATOR);
			pSysMenu->AppendMenu(MF_STRING, IDM_ABOUTBOX, strAboutMenu);
		}
	}

	// Set the icon for this dialog.  The framework does this automatically
	//  when the application's main window is not a dialog
	SetIcon(m_hIcon, TRUE);			// Set big icon
	SetIcon(m_hIcon, FALSE);		// Set small icon

	// TODO: Add extra initialization here
	// 读取登录信息
	CString init_user_name, init_room_name, init_status;
	//GetModuleFileName(NULL, cur_path. GetBufferSetLength(MAX_PATH + 1), MAX_PATH);
	//ini_path = CT2A(cur_path.GetString());
	//ini_path = ini_path + "\\data.ini";
	BOOL exist = finder.FindFile(_T("..\\data.ini"));
	if (exist) {
		::GetPrivateProfileString(_T("Account"), _T("room"), _T(""), init_room_name.GetBuffer(20), 20, _T("..\\data.ini"));
		::GetPrivateProfileString(_T("Account"), _T("user"), _T(""), init_user_name.GetBuffer(20), 20, _T("..\\data.ini"));
		::GetPrivateProfileString(_T("Account"), _T("status"), _T("0"), init_status.GetBuffer(20), 20, _T("..\\data.ini"));

		m_room.SetWindowTextW(init_room_name);
		m_name.SetWindowTextW(init_user_name);
		if (init_status == "1") {
			CCButton.SetCheck(BST_CHECKED);
		}
		//SetDlgItemText(IDC_NAME, _T());
	}
	return TRUE;  // return TRUE  unless you set the focus to a control
}

void CRTCDemoDlg::OnSysCommand(UINT nID, LPARAM lParam)
{
	if ((nID & 0xFFF0) == IDM_ABOUTBOX)
	{
		CAboutDlg dlgAbout;
		dlgAbout.DoModal();
	}
	else
	{
		CDialogEx::OnSysCommand(nID, lParam);
	}
}

// If you add a minimize button to your dialog, you will need the code below
//  to draw the icon.  For MFC applications using the document/view model,
//  this is automatically done for you by the framework.

void CRTCDemoDlg::OnPaint()
{
	if (IsIconic())
	{
		CPaintDC dc(this); // device context for painting

		SendMessage(WM_ICONERASEBKGND, reinterpret_cast<WPARAM>(dc.GetSafeHdc()), 0);

		// Center icon in client rectangle
		int cxIcon = GetSystemMetrics(SM_CXICON);
		int cyIcon = GetSystemMetrics(SM_CYICON);
		CRect rect;
		GetClientRect(&rect);
		int x = (rect.Width() - cxIcon + 1) / 2;
		int y = (rect.Height() - cyIcon + 1) / 2;

		// Draw the icon
		dc.DrawIcon(x, y, m_hIcon);
	}
	else
	{
		CDialogEx::OnPaint();
	}
}

// The system calls this function to obtain the cursor to display while the user drags
//  the minimized window.
HCURSOR CRTCDemoDlg::OnQueryDragIcon()
{
	return static_cast<HCURSOR>(m_hIcon);
}



void CRTCDemoDlg::OnBnClickedLogin()
{
	CString user_name, room_name;
	m_name.GetWindowText(user_name);
	m_room.GetWindowText(room_name);
	int cb_state = IsDlgButtonChecked(IDC_CHECK1);

	// 确认记住登录状态
	if (cb_state == 1){
		WritePrivateProfileString(_T("Account"), _T("room"), room_name.GetBuffer(0), _T("..\\data.ini"));
		WritePrivateProfileString(_T("Account"), _T("user"), user_name.GetBuffer(0), _T("..\\data.ini"));
		WritePrivateProfileString(_T("Account"), _T("status"), _T("1"), _T("..\\data.ini"));

	}
	// 取消记住登录状态
	else {
		WritePrivateProfileString(_T("Account"), _T("room"), _T(""), _T("..\\data.ini"));
		WritePrivateProfileString(_T("Account"), _T("user"), _T(""), _T("..\\data.ini"));
		WritePrivateProfileString(_T("Account"), _T("status"), _T("0"), _T("..\\data.ini"));
	}

	CChatDialog chat_dlg;
	chat_dlg.Init(user_name, room_name);
	chat_dlg.DoModal();
}


void CRTCDemoDlg::OnEnChangeName()
{
	// TODO:  如果该控件是 RICHEDIT 控件，它将不
	// 发送此通知，除非重写 CDialogEx::OnInitDialog()
	// 函数并调用 CRichEditCtrl().SetEventMask()，
	// 同时将 ENM_CHANGE 标志“或”运算到掩码中。

	// TODO:  在此添加控件通知处理程序代码
}

void CRTCDemoDlg::OnBnClickedCheck1()
{
	// TODO: 在此添加控件通知处理程序代码
}