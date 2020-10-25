### 方式一 [main.go](main.go)

命令行参数使用`flag`包即可实现解析

json 配置文件的读取使用`encoding/json`包即可

日志系统使用`log`包即可

### 方式二 [main_viper.go](main_viper.go)

命令行参数使用`flag`包实现解析

json 配置文件的读取使用`viper`包

日志系统使用`logrus`包
