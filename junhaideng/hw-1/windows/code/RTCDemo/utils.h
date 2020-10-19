#pragma once
#include "afxdialogex.h"

// 获取当前路径
void getCurrentPath(CString& path);

// 判断当前路径中是否存在
bool hasInitFile(const CString path);

// 配置信息，保存房间名和用户名
struct Info
{
	CString room;
	CString	name;
};

// 从配置文件中读取信息
Info readInitFile(const CString path);

// 将相关信息写入配置文件
void writeInitFile(const CString path, const CString room_id, const CString username);