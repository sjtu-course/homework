import cv2
import numpy as np
from ffmpy3 import FFmpeg


def play_video(video_name):
    cap = cv2.VideoCapture(video_name)  # 打开视频
    while (True):
        ret, frame = cap.read()  # 捕获一帧图像
        if ret:
            cv2.imshow('frame', frame)
            cv2.waitKey(25)
        else:
            break

    cap.release()  # 关闭视频
    cv2.destroyAllWindows()  # 关闭窗口



# play video 1
play_video('../video/video_1.mp4')

# play video 2
play_video('../video/video_2.mp4')

# use ffmpeg to mix 2 videos

ff = FFmpeg(
    inputs={'../video/video_1.mp4': None, '../video/video_2.mp4': None},
    outputs={'output.mp4': '-filter_complex mix=inputs=2:duration=longest'}
)
print(ff.cmd)
ff.run()


# play output video
play_video('output.mp4')

# de-noise and play the video
cap = cv2.VideoCapture('output.mp4')  # 打开视频
fourcc = cv2.VideoWriter_fourcc(*'avc1')
length = cap.get(3)
height = cap.get(4)
out_video = cv2.VideoWriter('output_improment.mp4',fourcc, 20.0,(320,240))
while (True):
    ret, frame = cap.read()  # 捕获一帧图像
    if ret:
        frame_blur = cv2.pyrMeanShiftFiltering(frame, 10, 50)


        frame_sharp = cv2.addWeighted(frame, 2, frame_blur, -1, 0)
        out = 2*frame
        out[out > 255] = 255
        # 数据类型转换
        out = np.around(out)
        out = out.astype(np.uint8)

        out_blur_gaussian = cv2.GaussianBlur(out, (15, 15), 0)  # 高斯模糊

        out = cv2.addWeighted(out, 2, out_blur_gaussian, -1, 0)
        out_blur_median = cv2.medianBlur(out, 5)
        out_video.write(out_blur_median)  # 保存帧
        cv2.imshow('frame', out_blur_median)
        if cv2.waitKey(1) & 0xFF == ord('q'):
            break
        cv2.waitKey(25)
    else:
        break
cap.release()
out_video.release()
cv2.destroyAllWindows()#关闭窗口