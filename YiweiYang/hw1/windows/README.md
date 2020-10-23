# Windows更新内容和功能：
- [x] 保存上次登录的房间名和用户名。
保存文件为`LoginInfo.ini`，位置在项目相同目录中。
- [x] 界面生成时尝试读取`LoginInfo.ini`文件，如果存在则将上次的登陆信息预填入登录界面
- 使用的是`WritePrivateProfileString()`这一类可以自定义ini文件的函数，而不是使用系统的`win.ini`文件，方便修改位置
