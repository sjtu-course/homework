#include "pch.h"
#include "utils.h"

CString GetModuleFolder(HMODULE hModule) {
	TCHAR szPath[MAX_PATH] = { 0 };
	GetModuleFileName(hModule, szPath, MAX_PATH);
	ZeroMemory(_tcsrchr(szPath, _T('\\')), _tcslen(_tcsrchr(szPath, _T('\\'))) * sizeof(TCHAR));
	return CString(szPath);
}