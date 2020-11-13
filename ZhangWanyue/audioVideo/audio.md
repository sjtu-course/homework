#### 3-1

1. PCM数据编码之前，将数据混音成单声道数据，送入编码器进行编码

```shell
ffmpeg -y -f s16le -ar 48000 -ac 2 -i origin.pcm -f s16le  -ar 48000 -ac 1 origin_mono.pcm
```



2. 通过opus接口将opus的编码复杂度分度设置为1和10，并做运算复杂度统计（可通过统计代码运行时间实现）

```{cpp}
opus_encoder_ctl(opus_encoder_handle, OPUS_SET_COMPLEXITY(complexity));

DWORD start_time = GetTickCount();
DWORD end_time = GetTickCount();
std::cout << "The run time is:" << (end_time - start_time) << "ms!" << std::endl;
```

complexity 设置为10，编码时间9875ms

complexity 设置为1，编码时间5625ms

#### 3-2

1. 评估48khz采样立体声序列，不同码率条件(码率: 32kbps/64kpbs/128kbps)下的质量差异；

   ```c++
   opus_encoder_ctl(opus_encoder_handle, OPUS_SET_BITRATE(encode_bitrate));
   ```

    相同采样频率下，码率越高，声音质量越好。



2. 评估48khz采样立体声序列，在audio模式@48kbps、CBR模式和voip模式@ 32kbps、CVBR模式的质量差异；

   ```c++
   
     //audio@48kbps
   encode_bitrate = 48000;
   opus_encoder_handle = opus_encoder_create(encode_sample_rate, encode_channel, OPUS_APPLICATION_AUDIO, &error);
   
     //cbr bitrate@48kbps
   encode_bitrate = 48000;
   opus_encoder_ctl(opus_encoder_handle, OPUS_SET_VBR(0));
   
         
     //voip@32kbps
   encode_bitrate = 32000;
   opus_encoder_handle = opus_encoder_create(encode_sample_rate, encode_channel, OPUS_APPLICATION_VOIP, &error);
         
     //cvbr@32kbps
   encode_bitrate = 32000;
   opus_encoder_ctl(opus_encoder_handle, OPUS_SET_VBR_CONSTRAINT(1));
   ```

文档和使用场景解释：

[OPUS_APPLICATION_VOIP](https://www.opus-codec.org/docs/html_api/group__ctlvalues.html#ga07884aa018303a419d1f7acb2f3fa669) gives best quality at a given bitrate for voice signals. It enhances the input signal by high-pass filtering and emphasizing formants and harmonics. Optionally it includes in-band forward error correction to protect against packet loss. Use this mode for typical VoIP applications. Because of the enhancement, even at high bitrates the output may sound different from the input.

[OPUS_APPLICATION_AUDIO](https://www.opus-codec.org/docs/html_api/group__ctlvalues.html#ga5909f7cb35c04f1110026c6889edd345) gives best quality at a given bitrate for most non-voice signals like music. Use this mode for music and mixed (music/voice) content, broadcast, and applications requiring less than 15 ms of coding delay.

[CVBR](https://www.opus-codec.org/docs/opus_api-1.2/group__opus__encoderctls.html#ga34d09ae06cab7e1a6c49876249b67892) can be useful for real-time communication and streaming. It means that the rate is almost constant, but can have short variations as long as they're compensated by opposite variations soon after.



3. 模拟丢包场景，对比64kbps、cvbr模式、不同丢包率条件下(丢包率: 0%/10%/30%)解码质量的差异；
   	TIPS：参照opus源码中opus_demo.c代码逻辑;

```
opus_encoder_ctl(opus_encoder_handle, OPUS_SET_PACKET_LOSS_PERC(pkt_loss));
```

评估效果：丢包率低，解码质量高。