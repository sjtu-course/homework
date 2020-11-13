# PCM数据编码之前，将数据混音成单声道数据，送入编码器进行编码

> demo中的源PCM数据是48k采样率，16bit量化，双声道数据
双声道数据是以交织的方式进行存储，也就是是说，按左声道、右声道、左声道、右声道…的方式存放
16bit如果大小不够的话需要截断，注意跳变。

由于输入是双声道的，我先尝试如果编码时将声道设置为1，解码声道也设置为1

    //channels: Number of channels (1 or 2) in input signal
    encode_channel = 1;
    decode_channel = 1;

直接这么合并后声音变得十分恐怖可怕，有杂音，女声都变成了男声！因为两个声道的数据直接被合并到了一个声道中。

我采取的是`OPUS_SET_FORCE_CHANNELS()`函数，编解码后声音很正常。

    #define INPUT_CHANNEL 2
    encode_channel = INPUT_CHANNEL;
    decode_channel = 1;
    opus_encoder_ctl(opus_encoder_handle, OPUS_SET_FORCE_CHANNELS(1));

可以在`opus_defines.h`中查看到此函数的说明

    /** Configures mono/stereo forcing in the encoder.
    * This can force the encoder to produce packets encoded as either mono or
    * stereo, regardless of the format of the input audio. This is useful when
    * the caller knows that the input signal is currently a mono source embedded
    * in a stereo stream.
    * @see OPUS_GET_FORCE_CHANNELS
    * @param[in] x <tt>opus_int32</tt>: Allowed values:
    * <dl>
    * <dt>#OPUS_AUTO</dt><dd>Not forced (default)</dd>
    * <dt>1</dt>         <dd>Forced mono</dd>
    * <dt>2</dt>         <dd>Forced stereo</dd>
    * </dl>
    * @hideinitializer */
    #define OPUS_SET_FORCE_CHANNELS(x) OPUS_SET_FORCE_CHANNELS_REQUEST, __opus_check_int(x)

# opus的编码复杂度统计

> 通过opus接口将opus的编码复杂度分度设置为1和10，并做运算复杂度统计

相关函数为


    /** Configures the encoder's computational complexity.
    * The supported range is 0-10 inclusive with 10 representing the highest complexity.
    * @see OPUS_GET_COMPLEXITY
    * @param[in] x <tt>opus_int32</tt>: Allowed values: 0-10, inclusive.
    *
    * @hideinitializer */
    #define OPUS_SET_COMPLEXITY(x) OPUS_SET_COMPLEXITY_REQUEST, __opus_check_int(x)

通过循环设定`COMPLEXITY`分别为0~10，使用计时方法为`#include<ctime>`的`clock()`进行计时，详见代码。计时仅包括编码过程，最终结果为

    Complexity=0. The run time is: 2.218s
    Complexity=1. The run time is: 2.606s
    Complexity=2. The run time is: 3.101s
    Complexity=3. The run time is: 3.306s
    Complexity=4. The run time is: 3.339s
    Complexity=5. The run time is: 4.179s
    Complexity=6. The run time is: 4.159s
    Complexity=7. The run time is: 5.221s
    Complexity=8. The run time is: 6.368s
    Complexity=9. The run time is: 6.387s
    Complexity=10. The run time is: 6.307s

# 评估48khz采样立体声序列不同码率条件下的质量

修改demo中的下面代码可以将48khz的opus文件解码为不同码率

    encode_bitrate = 128000;

- 码率: 32kbps
  - 声音的塑料感很严重，声音听起来有在颤抖的感觉。
- 码率: 64kpbs
  - 较为清晰，基本感觉良好。
- 码率: 128kbps
  - 非常清晰，犹如临场视听的感觉

ps：这个按键自动播放音频真的吓人。

# 评估48khz采样立体声序列不同模式条件下的质量

> 个人理解：opus支持恒定码率CBR和变码率VBR以及受约束码率CVBR三种编码方式。一般流媒体使用CBR，voip场景使用VBR。VBR编码效率较高，在网络条件好的情况下可以选用CVBR模式。

- `audio模式@48kbps、CBR模式`对应设置为

        encode_bitrate = 48000;
        //set audio
        opus_encoder_handle = opus_encoder_create(encode_sample_rate, encode_channel, OPUS_APPLICATION_AUDIO, &error);
        //set cbr bitrate
        /*0 Hard CBR.
        * 1 VBR (default).The exact type of VBR may be retrieved via 
        * #OPUS_GET_VBR_CONSTRAINT.*/
        opus_encoder_ctl(opus_encoder_handle, OPUS_SET_VBR(0));
    
- `voip模式@ 32kbps、CVBR模式`对应设置为

        encode_bitrate = 32000;
        //set voip
        opus_encoder_handle = opus_encoder_create(encode_sample_rate, encode_channel, OPUS_APPLICATION_VOIP, &error);
        //set cvbr
        /*0 Hard CBR.
        * 1 VBR (default).The exact type of VBR may be retrieved via 
        * #OPUS_GET_VBR_CONSTRAINT.*/
        opus_encoder_ctl(opus_encoder_handle, OPUS_SET_VBR(1));
        /*0 Unconstrained VBR.
        * 1 Constrained VBR (default).*/
        opus_encoder_ctl(opus_encoder_handle, OPUS_SET_VBR_CONSTRAINT(1));

    比较起来确实是第一种配置更为清晰，毕竟第二种更注重压缩效率。

# 评价不同丢包率条件下解码质量的差异

> 使用64kbps、cvbr模式

    encode_bitrate = 64000;
    //set cvbr
    /*0 Hard CBR.
    * 1 VBR (default).The exact type of VBR may be retrieved via 
    * #OPUS_GET_VBR_CONSTRAINT.*/
    opus_encoder_ctl(opus_encoder_handle, OPUS_SET_VBR(1));
    /*0 Unconstrained VBR.
    * 1 Constrained VBR (default).*/
    opus_encoder_ctl(opus_encoder_handle, OPUS_SET_VBR_CONSTRAINT(1));

参照opus源码中[opus_demo.c](https://github.com/gcp/opus/blob/master/src/opus_demo.c)代码逻辑，设置不同的`packet_loss_perc`即可

    fprintf(stderr, "-loss <perc>         : simulate packet loss, in percent (0-100); default: 0\n" );
    int packet_loss_perc;
    packet_loss_perc = 0; //丢包率0%/10%/30%
    opus_encoder_ctl(enc, OPUS_SET_PACKET_LOSS_PERC(packet_loss_perc));
    lost = len[toggle]==0 || (packet_loss_perc>0 && rand()%100 < packet_loss_perc);
    if (lost)
        opus_decoder_ctl(dec, OPUS_GET_LAST_PACKET_DURATION(&output_samples));
    else
        output_samples = max_frame_size;

- 丢包率: 0%

    正常情况

- 丢包率:10%

    轻微的声音缺失，有顿感

- 丢包率:30%
  
  能够听出来卡卡的，不过还是能靠人类自身的能力想象缺失的部分（人脑插值，毕竟听了那么多遍了）。这也说明opus抗丢包能力还不错。

查到的一些说明：

> Opus和Silk的编码器提出一种新方法，采用了下降码率的做法，类似于两个8kbps。在16kbps的音频流中，有4kbps的小包来对前一帧补偿。一旦大的包丢了，就使用小包来进行恢复，但是带来的问题是音频质量下降了。FEC是一种很好的抗丢包方法，但是它的问题是有可能会浪费带宽。使用FEC之后，确实能提高包的到达率，能在有限的延时下把通信的质量提高。
[来源：音视频抗丢包技术综述，面向不可靠传输网络的抗丢包编解码器](https://www.sohu.com/a/198552141_458408)