#!/bin/zsh

inpcm="resource/origin.pcm"
outdir="resource/out"

if [ ! -d "$outdir/comp" ]; then
    mkdir -p "$outdir/comp"
fi

for comp in {0..10}; do
    outbase="${outdir}/comp/comp_${comp}@48kbps"
    /usr/bin/time opusenc --raw $inpcm --bitrate 48 --comp ${comp} "$outbase.opus" 2>&1 |tee "$outbase.log"
done

if [ ! -d "$outdir/bitrate" ]; then
    mkdir -p "$outdir/bitrate"
fi

for bitrate in 32 64 128; do
    outbase="${outdir}/bitrate/bitrate@${bitrate}kbps"
    /usr/bin/time opusenc --raw $inpcm --bitrate $bitrate "$outbase.opus" 2>&1 |tee "$outbase.log"
done

if [ ! -d "$outdir/mode" ]; then
    mkdir -p "$outdir/mode"
fi

outbase="${outdir}/mode/audio_vbr@48kbps"
/usr/bin/time opusenc --raw $inpcm --audio --bitrate 48 "$outbase.opus" 2>&1 |tee "$outbase.log"
outbase="${outdir}/mode/audio_cbr@48kbps"
/usr/bin/time opusenc --raw $inpcm --audio --bitrate 48 --hard-cbr "$outbase.opus" 2>&1 |tee "$outbase.log"
outbase="${outdir}/mode/voip_vbr@32kbps"
/usr/bin/time opusenc --raw $inpcm --voip --bitrate 32 "$outbase.opus" 2>&1 |tee "$outbase.log"
outbase="${outdir}/mode/voip_cvbr@32kbps"
/usr/bin/time opusenc --raw $inpcm --voip --bitrate 32 --cvbr "$outbase.opus" 2>&1 |tee "$outbase.log"

if [ ! -d "$outdir/loss" ]; then
    mkdir -p "$outdir/loss"
fi

for loss in 0 10 30; do
    outbase="${outdir}/loss/loss_${loss}@64kbps"
    /usr/bin/time opusenc --raw $inpcm --voip --bitrate 64 --cvbr "$outbase.opus" 2>&1 |tee "$outbase.log"
done