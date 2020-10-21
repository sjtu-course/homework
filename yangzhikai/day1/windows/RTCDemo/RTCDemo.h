
// RTCDemo.h : main header file for the PROJECT_NAME application
//

#pragma once

#ifndef __AFXWIN_H__
	#error "include 'pch.h' before including this file for PCH"
#endif

#include <afxwin.h>

#include "resource.h"		// main symbols
#include <string>


// CRTCDemoApp:
// See RTCDemo.cpp for the implementation of this class
//

class CRTCDemoApp : public CWinApp
{
public:
	CRTCDemoApp();
	//std::string GetModuleFilePath();
	//std::string GetProfileString();

// Overrides
public:
	virtual BOOL InitInstance();


// Implementation

	DECLARE_MESSAGE_MAP()
};

extern CRTCDemoApp theApp;
