---
layout: post
title: 라즈베리파이 비디오에서 모델을 적용시키기
category: TiL
tag: [RaspberryPi, OpenCV, Keras] 
---

실제로 라즈베리파이의 카메라로 영상을 찍으며 얼굴 인식을 실험해봤다.   
[링크](https://github.com/jo631/frames-client)에 현재까지의 진행상황을 저장하고있다.  
현재 진행상황으로는 비디오의 프레임당 모델을 적용시켜 실행했는데, 라즈베리파이의 CPU로는 도저히 따라갈 수 없는 성능이다. 

일단 여태까지 만든 파이썬 코드들을 파헤쳐보려고 한다.  

```python
import time
import cv2
import numpy as np

from src.common.package.config import application
from src.opencv.package.config import application as _application
from src.common.package.http import server as _server
from src.common.package.http.handler import Handler
from src.common.package.camera.capture import Capture as _capture
from src.common.package.frame.action import Action as _frame
from src.common.package.frame.draw import Draw as _draw
from src.opencv.package.opencv.opencv import OpenCV

from tensorflow.keras.models import load_model
from tensorflow.keras.applications.mobilenet_v2 import preprocess_input
from tensorflow.keras.preprocessing.image import img_to_array

# Constant
_opencv = OpenCV()
model = load_model("mask_detector.model")


# Imutils 라이브러리를 사용해 HTTP 스트리밍을 한다.
# StreamHandler 클래스를 오버라이드하여 적용시킨다
class StreamHandler(Handler):

    
    # Handler.stream() 함수 오버라이드
    def stream(self):
        Handler.stream(self)
        print('[INFO] Overriding stream method...')

        # Imutils 라이브러리의 VideoStream 클래스 초기화
        capture = _capture(src=application.CAPTURING_DEVICE,
                           use_pi_camera=application.USE_PI_CAMERA,
                           resolution=application.RESOLUTION,
                           frame_rate=application.FRAME_RATE)

        if application.USE_PI_CAMERA:
            print('[INFO] Warming up pi camera...')
        else:
            print('[INFO] Warming up camera...')

        time.sleep(2.0)

        print('[INFO] Start capturing...')

        while True:
            # 한 순간을 캡처한다
            frame = capture.read()

            # 프레임 크기 50퍼센트 감소 ( 라즈베리파이의 성능 향상 )
            #frame = _frame.scale(frame=frame, scale=0.5)

            # 프레임을 흑백으로 변환( 1채널밖에 없어서 오류가 발생함, 고쳐야 
            #frame = cv2.cvtColor(frame, cv2.COLOR_BGR2GRAY)

            # 프레임의 height, width 추출
            (height, width) = frame.shape[:2]
            
            # OpenCV의 얼굴인식 모델 사용, src/dnn/ 안에 정의되어있음
            detections = _opencv.dnn_face_detector(frame=frame,
                                                   scale_factor=1.0,
                                                   size=(300, 300),
                                                   mean=(104.0, 177.0, 123.0))

            # Up size frame to 50% (how the frame was before down sizing)
            #frame = _frame.scale(frame=frame, scale=2)

            # If returns any detection
            for i in range(0, detections.shape[2]):
                
                
                # Get confidence associated with the detection
                confidence = detections[0, 0, i, 2]

                # Filter weak detection
                if confidence < _application.CONFIDENCE:
                    continue
                
                # Calculate coordinates
                box = detections[0, 0, i, 3:7] * np.array([width,
                                                           height,
                                                           width,
                                                           height])

                """
                (left, top, right, bottom) = box.astype('int')
                coordinates = {'left': left,
                               'top': top,
                               'right': right,
                               'bottom': bottom}
                text = "{:.2f}%".format(confidence * 100)
                frame = _draw.rectangle(frame=frame,
                                        coordinates=coordinates,
                                        text=text)
                """

                # 마스크 확인
                # compute the (x, y)-coordinates of the bounding box for
			    # the object
                (startX, startY, endX, endY) = box.astype("int")

			    # ensure the bounding boxes fall within the dimensions of
			    # the frame
                (startX, startY) = (max(0, startX), max(0, startY))
                (endX, endY) = (min(width - 1, endX), min(height- 1, endY))

		    	# extract the face ROI, convert it from BGR to RGB channel
		    	# ordering, resize it to 224x224, and preprocess it
                face = frame[startY:endY, startX:endX]
                face = cv2.cvtColor(face, cv2.COLOR_BGR2RGB)
                face = cv2.resize(face, (224, 224))
                face = img_to_array(face)
                face = preprocess_input(face)
                face = np.expand_dims(face, axis=0)

		    	# pass the face through the model to determine if the face
		    	# has a mask or not
                (mask, withoutMask) = model.predict(face)[0]

		    	# determine the class label and color we'll use to draw
		    	# the bounding box and text
                label = "Mask" if mask > withoutMask else "No Mask"
                color = (0, 255, 0) if label == "Mask" else (0, 0, 255)

		    	# include the probability in the label
                label = "{}: {:.2f}%".format(label, max(mask, withoutMask) * 100)

		    	# display the label and bounding box rectangle on the output
		    	# frame
                cv2.putText(frame, label, (startX, startY - 10),
                	cv2.FONT_HERSHEY_SIMPLEX, 0.45, color, 2)
                cv2.rectangle(frame, (startX, startY), (endX, endY), color, 2)
            


            """
            # Write date time on the frame
            frame = _draw.text(frame=frame,
                               coordinates={'left': application.WIDTH - 150, 'top': application.HEIGHT - 20},
                               text=time.strftime('%d/%m/%Y %H:%M:%S', time.localtime()),
                               font_color=(0, 0, 255))
            """
            
            
            # Convert frame into buffer for streaming
            retval, buffer = cv2.imencode('.jpg', frame)
            #buffer = cv2.imencode(frame)
            # Write buffer to HTML Handler
            self.wfile.write(b'--FRAME\r\n')
            self.send_header('Content-Type', 'frame/jpeg')
            self.send_header('Content-Length', len(buffer))
            self.end_headers()
            self.wfile.write(buffer)
            self.wfile.write(b'\r\n')



##
# Method main()
##
def main():
    try:
        address = ('', application.HTTP_PORT)
        server = _server.Server(address, StreamHandler)
        print('[INFO] HTTP server started successfully at %s' % str(server.server_address))
        print('[INFO] Waiting for client to connect to port %s' % str(application.HTTP_PORT))
        server.serve_forever()
    except Exception as e:
        server.socket.close()
        print('[INFO] HTTP server closed successfully.')
        print('[ERROR] Exception: %s' % str(e))


if __name__ == '__main__':
    main()


```