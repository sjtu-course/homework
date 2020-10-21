#include "pch.h"
#include "utils.h"

// 获取当前exe文件的所在目录
void getCurrentPath(CString& path) {
	HMODULE module = GetModuleHandle(0);
	CString temp;
	GetModuleFileName(module, temp.GetBufferSetLength(MAX_PATH + 1), MAX_PATH);
	// fullpath 是exe文件的路径
	CString fullpath(temp);
	int nPos = fullpath.ReverseFind(TEXT('\\'));
	
	if (nPos < 0)
		path = L"";
	else
		path = fullpath.Left(nPos);
}

// 判断是否存在配置文件
bool hasInitFile(const CString path) {
	if (PathFileExists(path)) {
		return true;
	}
	return false;
}

// 将值写入配置文件
void writeInitFile(const CString path,  const CString room_id, const CString username) {
	// 打开文件的时候清空处理
	CStdioFile file(path, CStdioFile::modeWrite|CStdioFile::modeCreate);
	file.WriteString(L"room=");
	file.WriteString(room_id);
	file.WriteString(L"\n");
	file.WriteString(L"name=");
	file.WriteString(username);
	file.Close();
}

// 根据赋值给info
void assignInfo(const CString content, const int pos, Info& info) {
	if (content.Left(pos).Compare(L"room")==0) {
		info.room = content.Right(content.GetLength()-pos-1);
	}
	else if (content.Left(pos).Compare(L"name")==0) {
		info.name = content.Right(content.GetLength()-pos-1);
	}
}


// 读取配置文件中的内容，并且返回
Info readInitFile(const CString path) {
	CStdioFile file(path, CFile::modeRead);
	CString content;
	// 创建info变量
	Info info;
	while (file.ReadString(content)) {
		int pos = content.Find(L"=", 0);
		assignInfo(content, pos, info);
	}
	file.Close();
	return info;
}
