
// RTCDemoDlg.h : header file
//

#pragma once


// CRTCDemoDlg dialog
class CRTCDemoDlg : public CDialogEx
{
// Construction
public:
	CRTCDemoDlg(CWnd* pParent = nullptr);	// standard constructor

// Dialog Data
#ifdef AFX_DESIGN_TIME
	enum { IDD = IDD_RTCDEMO_DIALOG };
#endif

	protected:
	virtual void DoDataExchange(CDataExchange* pDX);	// DDX/DDV support


// Implementation
protected:
	HICON m_hIcon;

	// Generated message map functions
	virtual BOOL OnInitDialog();
	afx_msg void InitName();               //<Jixiang: read init
	afx_msg void OnSysCommand(UINT nID, LPARAM lParam);
	afx_msg void OnPaint();
	afx_msg HCURSOR OnQueryDragIcon();
	DECLARE_MESSAGE_MAP()
public:
	CEdit m_name;
	CEdit m_room;
	afx_msg void OnBnClickedLogin();
	afx_msg void ClickedLoginCached(const CString &, const CString &); //<Jixiang: write
};
