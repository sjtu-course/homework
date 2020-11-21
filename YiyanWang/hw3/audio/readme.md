## 1.1 评估48khz采样立体声序列，不同码率条件(码率: 32kbps/64kpbs/128kbps)下的质量差异
1. 通过改变encode_bitrate的值为320000/640000/128000，码率越高，音频越清晰。
## 1.2 评估48khz采样立体声序列，在audio模式@48kbps、CBR模式和voip模式@ 32kbps、CVBR模式的质量差异
* 参数配置
```C++
#define OPUS_SET_VBR(x) OPUS_SET_VBR_REQUEST, __opus_check_int(x)
/* 
 * <dt>0</dt><dd>Hard CBR.</dd>
 * <dt>1</dt><dd>VBR (default). The exact type of VBR may be retrieved via #OPUS_GET_VBR_CONSTRAINT.</dd>
*/

#define OPUS_SET_VBR_CONSTRAINT(x) OPUS_SET_VBR_CONSTRAINT_REQUEST, __opus_check_int(x)
/* 
 * <dt>0</dt><dd>Unconstrained VBR.</dd>
 * <dt>1</dt><dd>Constrained VBR (default).</dd>
 */

#define OPUS_APPLICATION_VOIP                2048
/** Best for broadcast/high-fidelity application where the decoded audio should be as close as possible to the input
 * @hideinitializer */
#define OPUS_APPLICATION_AUDIO               2049
/** Only use when lowest-achievable latency is what matters most. Voice-optimized modes cannot be used.
 * @hideinitializer */
```

1. audio模式@48kbps、CBR模式：
```C++
encode_bitrate = 48000;
opus_encoder_handle = opus_encoder_create(encode_sample_rate, encode_channel, OPUS_APPLICATION_AUDIO, &error);
opus_encoder_ctl(opus_encoder_handle, OPUS_SET_VBR(0));
```

2. voip模式@ 32kbps、CVBR模式：
```C++
encode_bitrate = 32000;
opus_encoder_handle = opus_encoder_create(encode_sample_rate, encode_channel, OPUS_APPLICATION_VOIP, &error);
opus_encoder_ctl(opus_encoder_handle, OPUS_SET_VBR_CONSTRAINT(1));
```
## 1.3 模拟丢包场景，对比64kbps、cvbr模式、不同丢包率条件下(丢包率: 0%/10%/30%)解码质量的差异
* 64kbps、cvbr模式的设置类似上面一个问题
* 不同丢包率条件下(丢包率: 0%/10%/30%)解码质量的差异
设置丢包率：
```C++
int packet_loss_perc = 30; // 0/10/30
opus_encoder_ctl(opus_encoder_handle, OPUS_SET_PACKET_LOSS_PERC(packet_loss_perc));

int output_samples;
lost = (res == 0) || (packet_loss_perc > 0 && rand() % 100 < packet_loss_perc);
if (lost)
  opus_decoder_ctl(opus_decoder_handle, OPUS_GET_LAST_PACKET_DURATION(&output_samples));
else
  output_samples = INPUT_FRAME_SIZE;
res = opus_decode(opus_decoder_handle, encoded_buffer, res, decoded_buffer, output_samples, 0);
```
## 2.1 PCM数据编码之前，将数据混音成单声道数据，送入编码器进行编码
* 调整参数
```C++
/* 
 * <dt>1</dt>         <dd>Forced mono</dd>
 * <dt>2</dt>         <dd>Forced stereo</dd>
*/
#define OPUS_SET_FORCE_CHANNELS(x) OPUS_SET_FORCE_CHANNELS_REQUEST, __opus_check_int(x)
```
* 加入/更改代码
```C++
opus_encoder_ctl(opus_encoder_handle, OPUS_SET_FORCE_CHANNELS(1)); // 不能直接将INPUT_CHANNEL改为1

decode_channel = 1;
```
## 2.2 通过opus接口将opus的编码复杂度分度设置为1和10，并做运算复杂度统计（可通过统计代码运行时间实现）
* 代码
```C++
opus_encoder_ctl(opus_encoder_handle, OPUS_SET_COMPLEXITY(10));
```
* 用windows.h的GetTickCount()函数计时，得到复杂度为1/10的输出如下:
```
complex = 10: 7.813s
complex = 1: 4.421s
```
