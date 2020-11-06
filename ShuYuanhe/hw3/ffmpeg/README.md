#ffmpeg



---



1. 使用`-filter_complex mix`对两个视频进行合并，可以比较明显得看到这是一个从8开始倒计时的视频

2. 调整视频对比度、亮度和饱和度增加提升了视觉上的观感（`eq=contrast=3:brightness=0.2:saturation=2`）

3. 对视频进行了`hqdn3d`降噪，肉眼效果不明显（也尝试了`denoise3d`与直接设置噪声取值`noise reduction`等方法，效果均不明显）




