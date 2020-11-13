首先使用了ffmpeg合成出`compose.mp4`。
之后使用了[FastDVDnet-Blind](https://github.com/Forsworns/FastDVDnet-Blind)中的降噪器处理出`denoiser.mp4`，`compose.mp4`影像中的帧被拆分成图像序列后进行降噪。
