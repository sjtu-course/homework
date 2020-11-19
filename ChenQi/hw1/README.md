# Windows RTC

主要修改了 `OnInitDialog()` 和 `OnBnClickedLogin()` 这两个函数：

- 功能实现：在 `OnInitDialog()`函数中应该读取已经保存的上次登录的房间名/用户名文件 `Login.ini`，在

 `OnBnClickedLogin()` 函数中应该在每次点击 `login` 的时候把这次登录的房间名/用户名保存到 `Login.ini`。

- 具体实现：课件里推荐使用的是 `GetProfileString` 和 `WriteProfileString` 的接口，但是查阅了资料后发现这些函数是操作系统配置文件 `Win.ini` 的函数：

  | 函数名              | 功能                                                         |
  | ------------------- | ------------------------------------------------------------ |
  | GetProfileSection   | 读取 `win.ini` 中指定节 `lpAppName` 中所有键名及其值。`lpReturnedString`字符串形式如下：Key1=Value1/0Key2=Value2/0…KeyN=ValueN/0/0 |
  | GetProfileString    | 读取 `win.ini` 中指定节 `lpAppName` 中键名为 `lpKeyName` 对应变量的字符串值。 |
  | GetProfileInt       | 读取 `win.ini` 中指定节 `lpAppName` 中键名为 `lpKeyName` 对应变量的整数值 |
  | WriteProfileSection | 写（替换）`win.ini` 中指定节 `lpAppName` 中的键值。`lpString`字符串形式同`GetProfileSection` 中的 `lpReturnedString`。 |
  | WriteProfileString  | 写（替换）`win.ini` 中指定节 `lpAppName` 中键名为 `lpKeyName` 对应变量的字符串值。 |

  为了让我们的Login信息更加灵活保存、同时避免修改系统的配置文件，我们可以使用操作用户自定义配置文件的函数：

  | 函数名                        | 功能                                                         |
  | ----------------------------- | ------------------------------------------------------------ |
  | GetPrivateProfileSectionNames | 读取lpFileName指定的配置文件中所有的节名，lpszReturnBuffer字符串形式如下：Section1/0Section2/0…SectionN/0/0 |
  | GetPrivateProfileSection      | 同GetProfileSection。                                        |
  | ⭐GetPrivateProfileString      | 同GetProfileString。                                         |
  | GetPrivateProfileInt          | 同GetProfileInt                                              |
  | GetPrivateProfileStruct       | 须同WritePrivateProfileStruct配套使用。                      |
  | WritePrivateProfileSection    | 同WriteProfileSection                                        |
  | ⭐WritePrivateProfileString    | 同WriteProfileString                                         |
  | WritePrivateProfileStruct     | 不常用。                                                     |

  其中，`GetPrivateProfileString` 的返回值是 `DWORD` ---------接收缓冲区的大小（long类型）

  > 型参：

  `LPCTSTR lpAppName` ---------- INI文件中的一个字段名，这个字串不区分大小写。
  `LPCTSTR lpKeyName` ---------- lpAppName 下的一个键名，也就是里面具体的变量名，这个字串不区分大小写。

  `LPCTSTR lpDefaut` ----------------如果没有其前两个参数值，则将此值赋给变量。指定的条目没有找到时返回的默认值。可设为空（""）。
  `LPSTR lpReturnedString` -------接收INI文件中的值的CString对象，指定一个字串缓冲区，长度至少为nSize。

  `DWORD nSize` ---------------------指定装载到lpReturnedString缓冲区的最大字符数量
  `LPCTSTR lpFileName` -----------完整的INI文件路径名

  在本实验中我将 `DWORD buffer` 设置为1024。

  最终的保存得到的 `Login.ini` 如下：

  ```ini
  [Login]
  RoomName=room
  UserName=chenqi
  ```

# Android Demo

