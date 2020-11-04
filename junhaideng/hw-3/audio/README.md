### 下午第一节课作业

#### 1. PCM数据编码之前，将数据混音成单声道数据，送入编码器进行编码

将`encode_channel`修改成1即可

```c++
//encode_channel = INPUT_CHANNEL; 
encode_channel = 1;
```



#### 2. 通过opus接口将opus的编码复杂度分度设置为1和10，并做运算复杂度统计

使用`<ctime>`库，创建两个`clock_t`类型`start`, `end`，然后插入到相应的位置，即可实现计时效果

编码复杂度设置可以使用下面的代码

```c++
opus_encoder_ctl(opus_encoder_handle, OPUS_SET_COMPLEXITY(0)); // 设置复杂度0
opus_encoder_ctl(opus_encoder_handle, OPUS_SET_COMPLEXITY(0)); // 设置复杂度为10
```

通过测试我们可以获取到对应的运行时间：

如果我们使用编码复杂度0，需要的时间大致为`5.088000s`，如果我们使用编码复杂度10，那么需要的时间为`20.943000s`，可以知道，编码复杂度越高，那么相应所需要的时间就越多。

<hr/>

### 下午第二次课作业

#### 1. 评估48khz采样立体声序列，不同码率条件(码率: 32kbps/64kpbs/128kbps)下的质量差异

码率越高，所获取到的音频质量越好，但是当码率达到一定水平的时候，对于大多数人来说区别并不是很明显，比如我 :cry:



#### 2. 评估48khz采样立体声序列，在audio模式@48kbps、CBR模式和voip模式@ 32kbps、CVBR模式的质量差异

使用下面的方式可以生成对应的音频，由于本人对音频不是很敏感，很难判断哪一个质量好:cry:

```c++
// audio模式@48kbps、CBR模式
opus_encoder_handle = opus_encoder_create(encode_sample_rate, encode_channel, OPUS_APPLICATION_AUDIO, &error); // 设置为audio编码
encode_bitrate = 48000;
opus_encoder_ctl(opus_encoder_handle, OPUS_SET_VBR(0)); // CBR码率控制

// voip模式@32kbps、CVBR模式
opus_encoder_handle = opus_encoder_create(encode_sample_rate, encode_channel, OPUS_APPLICATION_VOIP, &error); // 设置为audio编码
encode_bitrate = 32000;
opus_encoder_ctl(opus_encoder_handle, OPUS_SET_VBR_CONSTRAINT(1)); // CVBR码率控制
```



#### 3. 模拟丢包场景，对比64kbps、cvbr模式、不同丢包率条件下(丢包率: 0%/10%/30%)解码质量的差异

在丢包率为0%时，音频播放流畅，

在丢包率为10%时，出现部分卡顿现象，

在丢包率为30%的时候，编码出来的音频出现了明显的卡顿现象

很明显可以知道，丢包率越低，那么解码的质量越好