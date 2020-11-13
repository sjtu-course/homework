import numpy as np
import cv2
import os

# os.system("ffmpeg -i video_1.mp4 -i video_2.mp4 -filter_complex mix video.mp4")
cap = cv2.VideoCapture('video.mp4')

fourcc = cv2.VideoWriter_fourcc(*'XVID')
fps = cap.get(cv2.CAP_PROP_FPS)
size = (int(cap.get(cv2.CAP_PROP_FRAME_WIDTH)), int(cap.get(cv2.CAP_PROP_FRAME_HEIGHT)))
out = cv2.VideoWriter('output6.mp4', fourcc, fps, size)
img, img2 = [], []
while(cap.isOpened()):
    ret, frame = cap.read()
    if ret==True:
        frame = cv2.fastNlMeansDenoising(frame, 3, 3, 9, 15)
        frame = cv2.blur(frame, (5,5))
        kernel = np.array([[0, -1, 0], [-1, 5, -1], [0, -1, 0]], np.float32)
        dst = cv2.filter2D(frame, -1, kernel=kernel)
        # piece = dst[:,:,0]
        # piece = cv2.equalizeHist(piece)
        # dst = cv2.cvtColor(piece, cv2.COLOR_GRAY2RGB)
        dst[dst > 160] = 255
        dst[dst < 90] = 0
        out.write(dst)
        # cv2.imshow('frame',dst)
        if cv2.waitKey(1) & 0xFF == ord('q'):
            break
    else:
        break

cap.release()
out.release()