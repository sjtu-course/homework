# RTC音频技术作业

> 一. PCM数据编码之前，将数据混音成单声道数据，送入编码器进行编码

对照 `PPT` 中 `go` 语言中的 `OPUS_SET_FORCE_CHANNELS(force_channel)` 接口，在 `C++` 中实现如下：

```C++
//set channels
opus_encoder_ctl(opus_encoder_handle, OPUS_SET_FORCE_CHANNELS(encode_bitrate));
```

听起来还是很清晰，单声道数据编码后 `wav` 文件大小为 `20.9MB` ，是之前双声道的一半。

> 二. 通过opus接口将opus的编码复杂度分度设置为1和10，并做运算复杂度统计（可通过统计代码运行时间实现）

设置不同配置复杂度的接口对照 `PPT` 中 `go` 语言中的 `OPUS_SET_COMPLEXITY(complexity)` 接口，在 `C++` 中实现如下：

```C++
time_t start = clock();
while (1) {
    int read_size = fread(input_buffer, sizeof(short), INPUT_FRAME_SIZE, origin_file_id);
    if (read_size != INPUT_FRAME_SIZE) {
        break;
    }
    process_cnt++;
    read_samples += read_size;
    if (process_cnt % 1000 == 0) {
        std::cout << "processing...., frame count:" << process_cnt << std::endl;
    }
    //encode process
    res = opus_encode(opus_encoder_handle, input_buffer, INPUT_FRAME_SIZE / 2, encoded_buffer, 1024);
    if (res > 0) {
        fwrite(encoded_buffer, 1, res, opus_file_id);
    }
}

time_t end = clock();
```



刚开始跑的结果如下：

```
Complexity 1 running time: 9seconds
Complexity 2 running time: 11seconds
Complexity 3 running time: 12seconds
Complexity 4 running time: 12seconds
Complexity 5 running time: 16seconds
Complexity 6 running time: 14seconds
Complexity 7 running time: 26seconds
Complexity 8 running time: 17seconds
Complexity 9 running time: 16seconds
Complexity 10 running time: 17seconds
```

不同复杂度的总时间并没有太大区别，后来仔细查看题目发现是要看编码时间，应该将解码的部分代码注释掉🤦‍

```
Complexity 1 running time: 6.825seconds
Complexity 2 running time: 9.358seconds
Complexity 3 running time: 9.125seconds
Complexity 4 running time: 9.221seconds
Complexity 5 running time: 10.265seconds
Complexity 6 running time: 10.796seconds
Complexity 7 running time: 13.264seconds
Complexity 8 running time: 18.094seconds
Complexity 9 running time: 15.055seconds
Complexity 10 running time: 15.695seconds
```

可以看出，除了复杂度 `2` 和 `3` 之间的编码时间比较相近，整体的编码时间是随着复杂度上升而上升的。

>三. 评估48khz采样立体声序列，不同码率条件(码率: 32kbps/64kpbs/128kbps)下的质量差异

通过修改 `encode_bitrate = 32000/64000/128000` 即可。

 `32kbps` 有点像老磁带CD的感觉，质量一般；`64kbps` 的效果就跟日常听音乐差不多；`128kbps` 明显感觉音质很饱满，质量很好。

> 四. 评估48khz采样立体声序列，在audio模式@48kbps、CBR模式和voip模式@ 32kbps、CVBR模式的质量差异

`CVBR模式` 参考PPT中给出的 `go` 接口找到对应的 `C++` 接口为 `OPUS_SET_VBR_CONSTRAINT(vbr_constraint)`

`VBR模式` 的接口为 `opus_encoder_ctl(opus_encoder_handle, OPUS_SET_VBR(vbr))`

我们可以设置 `audio模式@48kbps、CBR模式` 如下：

```C++
encode_bitrate = 48000;
//set CBR
opus_encoder_ctl(opus_encoder_handle, OPUS_SET_VBR(0));
```

播放质量很清晰。

设置`voip模式@ 32kbps、CVBR模式`如下，每帧分配不同编码比特数，总体保证平均码率不变，故应该设置

```C++
encode_bitrate = 32000;
//set CVBR
opus_encoder_ctl(opus_encoder_handle, OPUS_SET_VBR_CONSTRAINT(1));
```

感觉播放质量也很清晰啊...

> 五. 模拟丢包场景，对比64kbps、cvbr模式、不同丢包率条件下(丢包率: 0%/10%/30%)解码质量的差异；

利用 `opus_encoder_ctl(opus_encoder_SET_PACKET_LOSS_PERC(pkt_loss))` 来实现编码端丢包，在 `64kbps、cvbr` 模式模式下，10%丢包我几乎听不出来有区别，30%的丢包能稍微听出来像网卡了一样。