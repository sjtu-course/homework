// CChatDialog.cpp : implementation file
// 聊天框

#include "pch.h"
#include "RTCDemo.h"
#include "CChatDialog.h"
#include "afxdialogex.h"


// CChatDialog dialog
// 动态创建  IMPLEMENT_DYNCREATE(class_name, base_class_name )
// class_name 派生类（子）名
// base_class_name 基类（父）名
IMPLEMENT_DYNAMIC(CChatDialog, CDialogEx)

// 构造函数
// 用父类的构造函数来创建自己的窗体
CChatDialog::CChatDialog(CWnd* pParent /*=nullptr*/)
	: CDialogEx(IDD_CHATDIALOG, pParent)
{

}

// 析构函数
CChatDialog::~CChatDialog()
{
}

// 变量与其关联控件交换数据
void CChatDialog::DoDataExchange(CDataExchange* pDX)
{
	CDialogEx::DoDataExchange(pDX);
	// 把变量和控件本身关联起来
	// pDX是指向CDataExchange对象的指针
	// 这种结构为指定对象提供了建立数据交换的环境，也包括结构的指向
	// IDC_EMSG是控件ID
	// m_strMessage是为IDC_EMSG控件关联的一个CString型变量
	DDX_Control(pDX, IDC_CHAT, m_chat);
	DDX_Control(pDX, IDC_INPUT, m_input);
}

// 初始化
void CChatDialog::Init(const CString& name, const CString& room)
{
	m_name = name;
	m_room = room;
}

// 宏定义
// 添加消息响应函数,为每个消息处理函数加入一个入口
BEGIN_MESSAGE_MAP(CChatDialog, CDialogEx)
	ON_BN_CLICKED(IDC_SEND, &CChatDialog::OnBnClickedSend)
END_MESSAGE_MAP()


// CChatDialog message handlers

// 初始化对话框
BOOL CChatDialog::OnInitDialog()
{
	CDialogEx::OnInitDialog();

	// TODO:  Add extra initialization here
	CString winTitle;
	winTitle.Format(L"RoomID - %s", m_room);
	// 改变指定窗口的标题栏的文本内容
	// 如果指定窗口是一个控件，则改变控件的文本内容。
	// SetWindowText函数不改变其他应用程序中的控件的文本内容
	SetWindowText(winTitle);

	CString helloText;
	helloText.Format(L"Hello, %s\r\n", m_name);
	m_chat.SetWindowText(helloText);

	return TRUE;  // return TRUE unless you set the focus to a control
				  // EXCEPTION: OCX Property Pages should return FALSE
}

// 对话内容编辑
void CChatDialog::OnBnClickedSend()
{
	CString strText;
	m_chat.GetWindowText(strText);
	strText += "\r\n";

	CString strInput;
	m_input.GetWindowText(strInput);
	strText += strInput;

	m_input.SetWindowText(L"");
	m_chat.SetWindowText(strText);
}

