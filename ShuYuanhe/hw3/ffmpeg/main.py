import subprocess

cmd1 = 'ffmpeg -i video_1.mp4 -i video_2.mp4 -filter_complex mix=inputs=2:duration=longest output1.mp4'
cmd2 = 'ffmpeg -i output1.mp4 -vf eq=contrast=3:brightness=0.2:saturation=2 output2.mp4'
cmd3 = 'ffmpeg -i output2.mp4 -vf hqdn3d output.mp4'
subprocess.call(cmd1, shell=True)
subprocess.call(cmd2, shell=True)
subprocess.call(cmd3, shell=True)



