使用`ffmpeg`提供的`mix`命令，将两个视频合并，即

```
ffmpeg -i video_1.mp4 -i video_2.mp4 -filter_complex mix output.mp4
```

得到的结果为一个从8开始的倒计时视频。