# Linux：Golang综合实战

- [x] 从 `config.json` 配置读取 `IP` ，`port`，`ID`，`roomID`，`name` 信息
- [x] 命令行传入 `config.json` 参数，即 `./main -config config.json`
- [x] 将日志写入文件系统，`chat.log`

主要用到了 `Golang` 如下的包：

- **encoding/json**：用 `json.Unmarshal` 来实现 `json` 的解码，`Unmarshal` 输入是 `[]byte, v` ，输出是 `error` 。将 `ioutil.ReadAll(jsonFile)` 读出来的 `data` 进行解码，解码到我们自定义的 `json` 结构体`UserFromJson` 中。

- **flag**：用来解析命令行参数，注意一定要通过 `flag.Parse()` 函数接下命令行参数，解析函数将会在碰到第一个非 `flag` 命令行参数时停止：

  ```go
  flag.Parse()
  ```

- **io/ioutil**：I/O文件读取

- **log**：日志包

- **os**：文件打开

设置的 `config.json` 文件如下：

```json
{
  "id": "111111",
  "port": "3001",
  "room_id": "test",
  "name": "cq",
  "ip": "124.236.22.20"
}
```

得到的 `log` 如下：

```
2020/11/19 17:26:05 Successfully opened config.json
2020/11/19 17:26:05 IP: 124.236.22.20
2020/11/19 17:26:05 Port: 3001
2020/11/19 17:26:05 ID: 111111
2020/11/19 17:26:05 Room_ID: test
2020/11/19 17:26:05 Name: chenqi
2020/11/19 17:26:05 connect success


2020/11/19 17:26:05 success: [map[address:183.192.10.52:6460 avatarUrl:http://q.qlogo.cn/headimg_dl?dst_uin=5684277&spec=100 id:Ov5Xxuay1HCRbBP1IAde login_time:2020-11-19T09:26:05.849915771Z name:rtcname0 room_id:roomid type:user] [map[address: avatarUrl:http://q.qlogo.cn/headimg_dl?dst_uin=1000000&spec=100 id:roomid login_time:2020-10-23T02:16:27.256163327Z name:【群聊】roomid room_id:roomid type:group]]]
2020/11/19 17:26:10 finished
```

