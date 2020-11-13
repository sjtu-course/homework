import cv2
import time
import subprocess as sp
import os


def play_video(filename):
    orig = cv2.VideoCapture(filename)

    while orig.isOpened():
        res, f = orig.read()
        if res == True:
            cv2.imshow(filename, f)
            cv2.waitKey(1)
        else:
            break
    orig.release()


def seq2mp4():
    cap = cv2.VideoCapture('./video/compose.mp4')
    fourcc = cv2.VideoWriter_fourcc('m', 'p', '4', 'v')
    fps = cap.get(cv2.CAP_PROP_FPS)
    size = (int(cap.get(cv2.CAP_PROP_FRAME_WIDTH)),
            int(cap.get(cv2.CAP_PROP_FRAME_HEIGHT)))
    out = cv2.VideoWriter('./video/denoise.mp4', fourcc, fps, size)
    seqs = os.listdir("./FastDVDnet-Blind/results/compose")
    seqs = list(filter(lambda x: x.endswith('.png'), seqs))
    for s in range(len(seqs)):
        print(f"./FastDVDnet-Blind/results/compose/n0_FastDVDnet_{s}.png")
        out.write(cv2.imread(
            f"./FastDVDnet-Blind/results/compose/n0_FastDVDnet_{s}.png"))


if __name__ == "__main__":
    # sp.run(
    #    ["ffmpeg", "-i", "./video/video_1.mp4", "-i", "./video/video_2.mp4", "-filter_complex", "mix", "./video/compose.mp4"])

    # play_video("./video/compose.mp4")
    # play_video("./video/denoise.mp4")
    seq2mp4()
