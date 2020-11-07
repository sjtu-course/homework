# OPUS

## 实验环境

- 尝试了给的sjtu_audio_demo，因为参数都是hardcode写在代码里了，为了方便后续的实验需要写argparse，感觉麻烦而划不来。毕竟都只是调用`ope_encoder_ctl`来修改配置。
- 选择[**opus-tools**](https://github.com/xiph/opus-tools)中的`opusenc`作为基础，添加application mode的选项以支持audio/voip的切换。
  - 修改后的代码在https://github.com/spartazhc/opus-tools
- 使用[**peaqb-fast**](https://github.com/akinori-ito/peaqb-fast)，作为PEAQ实现[ITU-R BS.1387](http://www.itu.int/rec/R-REC-BS.1387/en)
  - 使用ODG(Objective difference grade) 作为评价指标。
  - 该peaq程序使用时感觉有点问题，一个wav自己和自己的对比都有-0.2+的ODG，编码后的和编码前的比有-3.+。
  - 不过似乎相对值还能一看。

## 实验项目

### stereo to mono

- 目标：在PCM数据编码前，混音成单声道数据。
- 使用`ffmpeg`实现：

```bash
ffmpeg -f s16le -ar 48000 -ac 2 -i origin.pcm  -f s16le -ar 48000 -ac 1 mono.pcm
```

- add wav header to PCM file

```bash
ffmpeg -f s16le -ar 48000 -ac 2 -i origin.pcm origin.wav
```

### complexity

- 目标：设置opus的编码复杂度，做运算复杂度分析

设定`ope_encoder_ctl(enc, OPUS_SET_COMPLEXITY(complexity));`以调节复杂度。

使用GNU time对程序运行计时：

```bash
for comp in {0..10}; do
    outbase="${outdir}/comp/comp_${comp}@48kbps"
    /usr/bin/time opusenc --raw $inpcm --bitrate 48 --comp ${comp} "$outbase.opus" 2>&1 |tee "$outbase.log"
done
```

得到的log如下：

```bash
Encoding using libopus 1.3.1 (audio)
-----------------------------------------------------
   Input: 48 kHz, 2 channels
  Output: 2 channels (2 coupled)
          20ms packets, 48 kbit/s VBR
 Preskip: 312

Encoding complete                            
-----------------------------------------------------
       Encoded: 3 minutes and 49.08 seconds
       Runtime: 2 seconds
                (114.5x realtime)
         Wrote: 1602716 bytes, 11454 packets, 232 pages
       Bitrate: 55.3243 kbit/s (without overhead)
 Instant rates: 1.2 to 96.4 kbit/s
                (3 to 241 bytes per packet)
      Overhead: 1.15% (container+metadata)

1.81user 0.00system 0:01.82elapsed 99%CPU (0avgtext+0avgdata 4160maxresident)k
0inputs+3136outputs (0major+438minor)pagefaults 0swaps
```

分析处理log：

```bash
grep -oP "^(\d+\.\d+)(?=user)" *.log |sed -re 's/comp_([0-9]+)@48kbps.log:/\1|/g'
```

表：user mode 时间对比

| complexity | user mode time (s) |
| ---------- | ------------------ |
| 0          | 0.65               |
| 1          | 0.87               |
| 2          | 0.98               |
| 3          | 0.99               |
| 4          | 1.04               |
| 5          | 1.28               |
| 6          | 1.30               |
| 7          | 1.63               |
| 8          | 1.82               |
| 9          | 1.82               |
| 10         | 1.81               |

- 分析
  - 复杂度配置0-10，数字越高复杂度越高；
  - 在本测试序列中，复杂度10的运算时间是复杂度0的三倍左右；
  - 复杂度8-10运算时间差不多。

### bitrate

- 目标：评估48khz采样立体声序列，不同码率条件(码率: 32kbps/64kpbs/128kbps)下的质量差异；
- 编码不同bitrate：

```bash
for bitrate in 32 64 128; do
    outbase="${outdir}/bitrate/bitrate@${bitrate}kbps"
    /usr/bin/time opusenc --raw $inpcm --bitrate $bitrate "$outbase.opus" 2>&1 |tee "$outbase.log"
done
```

- 对比质量差异：

  - 人耳：对比128kbps，勉强能感觉到32kbps在乐器声音上有些欠缺，128kbps与64kbps无法感受到差异。
  - PEAQ：数据有点问题，但是比特率越高，ODG越大

  ```bash
  # ODG 
  bitrate@128kbps.peaq: -3.30925
  bitrate@64kbps.peaq: -3.35684
  bitrate@32kbps.peaq: -3.51749
  ```

### application & rate control

- 目标：评估48khz采样立体声序列，在audio模式@48kbps、CBR模式和voip模式@ 32kbps、CVBR模式的质量差异；

```bash
opusenc --raw $inpcm --audio --bitrate 48 "${outdir}/mode/audio_vbr@48kbps"
opusenc --raw $inpcm --audio --bitrate 48 --hard-cbr "${outdir}/mode/audio_cbr@48kbps"
opusenc --raw $inpcm --voip --bitrate 32 "${outdir}/mode/voip_vbr@32kbps"
opusenc --raw $inpcm --voip --bitrate 32 --cvbr "${outdir}/mode/voip_cvbr@32kbps"
```

- 对比质量差异：

  - 人耳：
    - 在每一组application内部，仅码控策略变化的情况下无法分辨。
    - audio和voip的差别很大，voip感觉像是频谱带限了，高频与低频部分都有缺失。

  - PEAQ
    - cbr好于vbr，猜测是因为取平均值的方式引起的，不一定正确
    - audio好于voip

  ```bash
  # ODG 
  audio_cbr@48kbps.opus.wav.peaq: -3.45136
  audio_vbr@48kbps.opus.wav.peaq: -3.41826
  voip_cvbr@32kbps.opus.wav.peaq: -3.55057
  voip_vbr@32kbps.opus.wav.peaq: -3.51462
  ```

### packet loss

- 目标：模拟丢包场景，对比64kbps、cvbr模式、不同丢包率条件下(丢包率: 0%/10%/30%)解码质量的差异；

- 这里没有搞清楚的是，模拟丢包是在编码端还是解码端，因为编码端有一个`expect loss`的选项

  ```bash
  # opusenc
  --expect-loss n    Set expected packet loss in percent (default: 0)
  # opusdec
  --packet-loss n       Simulate n % random packet loss
  ```

- 解码质量对比

  - 人耳：
    - 丢包的音质受到很大影响，10%勉强还算流畅，30%勉强能听出内容
  - PEAQ
    - PEAQ的测试让我发现了或许opusenc的`--expect-loss`选项并没有用，和听觉感受一致
    - `loss_0`和`loss_30`分别是设置`--expect-loss`的结果，但是他们的peaq值是一样的
    - 丢包10的ODG比30还差，有点异常。

  ```bash
  loss_0_0@64kbps.wav.peaq: -3.37574
  loss_0_10@64kbps.wav.peaq: -3.65235
  loss_0_30@64kbps.wav.peaq: -3.4826
  loss_30_0@64kbps.wav.peaq: -3.37574
  loss_30_10@64kbps.wav.peaq: -3.65235
  loss_30_30@64kbps.wav.peaq: -3.4826
  ```

  

