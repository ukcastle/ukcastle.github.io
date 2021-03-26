---
layout: post
title: PyQT 에 OpenCV 이미지 보여주기
category: Python
tag: [Python, PyQT, OpenCV]  
---

원리는 이렇다.  

PyQT의 Label 클래스에는 Pixmap 형식으로 이미지를 넣어줄 수 있다.  

근데 이 방식은 무조건 **멀티스레드**로 이루어져야 한다.  
왜냐면 메인 프레임도 띄워줘야하기 때문에..  

[이전 포스팅](https://jo631.github.io/python/2021/03/25/PyQT/)의 클래스를 또 쓸 것인데, 쓸모 없는 부분은 생략하겠다.  

#### GUI

라벨을 만들고, 보여줄 준비를 하고, 스레드를 돌린다.

```python
class Ui_Main(object):

    def __init__(self, W, H):
        
        super().__init__()
        #...

        self.vd, th = self.setupVideo(W,H)
        th.start()
        #...

    def setupVideo(self,W,H):
        half_W = int(W / 2)
        half_H = int(H / 2)
        tick_W = int(W / TICK)
        tick_H = int(H / TICK)

        size = half_W - 2 * (tick_W)

        self.FR_Camera = OM.makeFrame(self.__centralwidget,
            startX= tick_W,
            startY = tick_H,
            W = size,
            H = size)

        self.LB_Camera_Main = OM.makeLabel(self.FR_Camera,
            0,0,size,size)

        vd = Video(self.LB_Camera_Main)
        vd.setRunning(True)

        th = threading.Thread(target=vd.run)

        return vd, th

```

프레임을 만들고(경계선이 필요 없으면 안만들어도 된다), 프레임 안에 같은 크기의 라벨을 만든다.  