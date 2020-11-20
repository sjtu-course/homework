# 使用ffmpeg混合两个视频并增强视觉效果

由于远程容器的ffmpeg版本不包括mix功能，使用的是从网盘下的本地代码和本地的ffmpeg，仅上传了有改动的部分。

代码在`code/notebook/04_homework_video_yyw.ipynb`，视频在`code/video/outputxxx.mp4`

> 由于播放视频的功能在容器中，所以在本地复制了一份`notebook_helpers.py`，稍微修改了代码确保正确导入。

**如果看不到图请访问[博客](https://www.cnblogs.com/smileglaze/p/13943236.html)或直接下载目录中的视频文件**

- `code/video/output.mp4`为直接使用ffmpeg混合两个输入得到的视频。可以看出视频内容为从8开始的倒计时，但不是很清楚。

![如果看不到图请访问博客或直接下载目录中的视频文件](https://img2020.cnblogs.com/blog/1507333/202011/1507333-20201117063806688-1380941637.gif)

[![如果看不到图请访问博客或直接下载目录中的视频文件](https://img2020.cnblogs.com/blog/1507333/202011/1507333-20201117063806688-1380941637.gif)](https://github.com/sjtu-course/homework/blob/main/YiweiYang/hw3/video/code/video/output.mp4)

- 首先观察到直接合成的视频对比度很低，先使用ffmpeg提高了对比度得到`code/video/output2.mp4`

[![如果看不到图请访问博客或直接下载目录中的视频文件](https://img2020.cnblogs.com/blog/1507333/202011/1507333-20201117063755438-1941509240.gif)](https://github.com/sjtu-course/homework/blob/main/YiweiYang/hw3/video/code/video/output2.mp4)

- 然后注意到视频中的噪点非常多，使用中值滤波和非局部平均去噪进行降噪，得到`code/video/output3.mp4`。
	- 该步骤非常耗时
	- 如果跳过前一步增加对比度的操作直接降噪，会使得画面几乎完全看不清，再增加对比度也只能得到马赛克图像

> 也可以使用ffmpeg带的[BM3D算法](https://zhuanlan.zhihu.com/p/92973703)

[![如果看不到图请访问博客或直接下载目录中的视频文件](https://img2020.cnblogs.com/blog/1507333/202011/1507333-20201117063745274-1583532204.gif)](https://github.com/sjtu-course/homework/blob/main/YiweiYang/hw3/video/code/video/output3.mp4)

- 视频经过滤波后，画面变得比较模糊，我希望通过锐化的方式让数字边缘看起来更清楚一些，得到`code/video/output4.mp4`

[![如果看不到图请访问博客或直接下载目录中的视频文件](https://img2020.cnblogs.com/blog/1507333/202011/1507333-20201117063733441-1027777838.gif)](https://github.com/sjtu-course/homework/blob/main/YiweiYang/hw3/video/code/video/output4.mp4)

- 锐化后噪点也看的更清楚了，使用中值滤波再处理一次，得到`code/video/output5.mp4`。由于最开始的噪声太强，如果不使用形状特征匹配，基本也就是这个效果了。

[![如果看不到图请访问博客或直接下载目录中的视频文件](https://img2020.cnblogs.com/blog/1507333/202011/1507333-20201117063724190-1213713227.gif)](https://github.com/sjtu-course/homework/blob/main/YiweiYang/hw3/video/code/video/output5.mp4)

**注意：** opencv保存的mp4文件需要跟原视频一样使用avci编码才可以，然而需要额外下载openh264.dll文件放在目录中，直接使用的话会报错

	Failed to load OpenH264 library: openh264-1.8.0-win64.dll
			Please check environment and/or download library: https://github.com/cisco/openh264/releases

	[libopenh264 @ 0000019f375b3cc0] Incorrect library version loaded
	Could not open codec 'libopenh264': Unspecified error
