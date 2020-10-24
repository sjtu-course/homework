#include "pch.h"
#include "ConfigUtils.h"
#include <atlstr.h>
		
#define MAX_BUFFER_SIZE 256
#define Init_FILE _T("\\rtc.ini")

CString getPath();

void readRoomAndUser(CString& room, CString& user) {
	CString path = getPath() + Init_FILE;
	TCHAR charBuffer[MAX_BUFFER_SIZE];
	CString section = CString(_T("config"));

	CString key = CString(_T("room"));
	GetPrivateProfileString(section, key, NULL, charBuffer, MAX_BUFFER_SIZE / 16, path);
	room = CString(charBuffer);

	key = _T("user");
	GetPrivateProfileString(section, key, NULL, charBuffer, MAX_BUFFER_SIZE / 16, path);
	user = CString(charBuffer);
}


bool writeRoomAndUser(CString room, CString user)
{
	CString path = getPath() + Init_FILE;
	bool res = TRUE;
	CString section = CString(_T("config"));
	CString key = CString(_T("room"));
	res = res && WritePrivateProfileString(section, key, room, path);

	key = CString("user");
	res = res && WritePrivateProfileString(section, key, user, path);

	return res;
}

CString getPath() {
	HMODULE module = GetModuleHandle(0);
	CString temp;
	GetModuleFileName(module, temp.GetBufferSetLength(MAX_PATH + 1), MAX_PATH);
	CString path(temp);
	int nPos = path.ReverseFind(TEXT('\\'));

	if (nPos < 0) {
		return _T("");
	}
	else {
		return path.Left(nPos);
	}
}
