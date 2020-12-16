import cv2
import numpy as np
import time
from inference_usbCam_face import TensoflowFaceDector
from inference_usbCam_face import PATH_TO_CKPT
from utils import visualization_utils_color as vis_util
from inference_usbCam_face import category_index
import os
import sys
import multiprocessing
from threading import Thread
from collections import deque


class HiddenPrints:
    def __enter__(self):
        self._original_stdout = sys.stdout
        sys.stdout = open(os.devnull, "w")

    def __exit__(self, exc_type, exc_val, exc_tb):
        sys.stdout.close()
        sys.stdout = self._original_stdout


def main():
    # Globals
    scale_factor = 5
    mog2 = cv2.cuda.createBackgroundSubtractorMOG2(
        120, 5, False
    )  # TODO: Optimize first two arguments
    detector = TensoflowFaceDector(PATH_TO_CKPT)

    cpus = multiprocessing.cpu_count()
    pool = multiprocessing.Pool(processes=cpus - 1)

    cameras = [
        (
            "Main",
            'filesrc location="Main_Entrance___Entry_Camera_01_20201123084117.mp4" ! qtdemux ! h264parse ! nvh264dec ! appsink',
        ),
        (
            "Side",
            'filesrc location="Side_Entrance___Entry_Camera_01_20201123124307.mp4" ! qtdemux ! h264parse ! nvh264dec ! appsink',
        ),
        (
            "Back",
            'filesrc location="Back_Entrance___Entry_Camera_01_20201123091410.mp4" ! qtdemux ! h264parse ! nvh264dec ! appsink',
        ),
        # ('Main',
        # 'rtspsrc location="rtsp://admin:xxxxxxxxx@192.168.1.207:554//h264Preview_01_main" ! rtph264depay ! h264parse ! nvh264dec ! appsink'),
        # ('Side',
        # 'rtspsrc location="rtsp://admin:xxxxxxxxx@192.168.1.200:554//h264Preview_01_main" ! rtph264depay ! h264parse ! nvh264dec ! appsink'),
        # ('Back',
        # 'rtspsrc location="rtsp://admin:xxxxxxxxx@192.168.1.203:554//h264Preview_01_main" ! rtph264depay ! h264parse ! nvh264dec ! appsink')
    ]

    capture_devices = []

    queue = deque()
    threads = []

    for camera_name, camera_stream in cameras:
        capture_device = cv2.VideoCapture(camera_stream, cv2.CAP_GSTREAMER)

        if capture_device.isOpened():
            thread = Thread(target=get_frame, args=(capture_device, queue, camera_name))
            thread.daemon = True
            thread.start()
            threads.append(thread)
            cv2.namedWindow(camera_name, cv2.WINDOW_NORMAL)
            cv2.resizeWindow(camera_name, 800, 600)

    while True:
        if not len(queue):
            break  # switch to continue if using rtsp streams
        print(len(queue))
        camera_name, frame = queue.popleft()
        # loop_start = time.time()

        # GStreamer outputs planar YUV420_NV12
        gray = frame[0:1920, 0:2560]

        # CPU: Planar YUV420_NV12 makes it difficult to crop before converting, so we will convert the whole frame
        rgb = cv2.cvtColor(frame, cv2.COLOR_YUV2RGB_NV12)

        # GPU: resize and background subtraction
        frame_gpu = cv2.cuda_GpuMat()
        frame_gpu.upload(gray)
        frame_gpu = cv2.cuda.resize(
            frame_gpu,
            (0, 0),
            fx=1.0 / scale_factor,
            fy=1.0 / scale_factor,
            interpolation=cv2.INTER_CUBIC,
        )
        frame_gpu = mog2.apply(frame_gpu, 0, None)  # TODO: Optimize second argument
        frame_gpu = frame_gpu.download()

        # CPU, but no cuda version exists for this function
        contours, hierarchy = cv2.findContours(
            frame_gpu, cv2.RETR_TREE, cv2.CHAIN_APPROX_SIMPLE
        )  # TODO: Optimize last two args

        if len(contours):
            contour = max(contours, key=cv2.contourArea)
            area = cv2.contourArea(contour)
            if 500 < area:
                (x, y, w, h) = cv2.boundingRect(contour)
                x *= scale_factor
                y *= scale_factor
                w *= scale_factor
                h *= scale_factor

                roi = rgb[y : y + h, x : x + w]
                cv2.rectangle(rgb, (x, y), (x + w, y + h), (0, 255, 0), 3)

                with HiddenPrints():
                    (boxes, scores, classes, num_detections) = detector.run(roi)

                vis_util.visualize_boxes_and_labels_on_image_array(
                    roi,
                    np.squeeze(boxes),
                    np.squeeze(classes).astype(np.int32),
                    np.squeeze(scores),
                    category_index,
                    use_normalized_coordinates=True,
                    line_thickness=4,
                )

        bgr = cv2.cvtColor(rgb, cv2.COLOR_RGB2BGR)
        print(camera_name)
        cv2.imshow(camera_name, bgr)
        # print('fps', 1 / (time.time() - loop_start))

        if cv2.waitKey(1) & 0xFF == ord("q"):
            break

    for capture_device in capture_devices:
        capture_device.release()

    for thread in threads:
        thread.join(timeout=1)


def process_image():
    return


def get_frame(capture_device, queue, camera_name):
    while True:
        try:
            if capture_device.isOpened():
                status, frame = capture_device.read()
                if status:
                    queue.append((camera_name, frame))
                else:
                    # disconnected
                    print("read(): empty frame", file=sys.stderr)
                    return
        except Exception:
            # disconnected
            print("read(): error.", file=sys.stderr)
            return


if __name__ == "__main__":
    main()
