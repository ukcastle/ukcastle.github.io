---
layout: post
title: 라즈베리파이 카메라 설정과 OpenCV4 설치
category: RaspberryPi
tag: [RaspberryPi, OpenCV]
---

# Install OpenCV4 in RaspBerryPi 3

오늘은 라즈베리에 카메라 연결을 한 뒤, OpenCV4를 설치했다.

#### Connect Camera Module

![카메라찍음](https://github.com/ukcastle/ukcastle.github.io/blob/main/postimg/%EC%B9%B4%EB%A9%94%EB%9D%BC%EC%B0%8D%EC%9D%8C.jpg?raw=true){: width="50%" heigth="50%"}

이런식으로 `Camera` 포트에 잘 끼워주면 된다.  
이제 설정을 해주자  
다음 터미널에 `sudo raspi-config`을 입력하고 `Interface Option`에 들어간 뒤 Camera를 **Enable** 해준다.  
이제 잘 연결이 되었는지 확인만 해보자.  
```
raspistill -o 이미지이름.jpg
```
잠시동안 디스플레이에 카메라 화면이 표시된 후 설정한 이름으로 저장이 된다.  
![이미지](https://github.com/ukcastle/ukcastle.github.io/blob/main/postimg/%EC%9D%B4%EB%AF%B8%EC%A7%80.jpg?raw=true){: width="50%" heigth="50%"}

연결이 잘 됐는지만 확인했다. 근데 내 카메라 모듈은... 초점이 좀 멀리잡히는 것 같다.  
당연하겠지만 오토포커싱 기능이 없네  
<br>

#### Install OpenCV

[GitHub 링크](https://github.com/dltpdn)를 보고 적용했다.  
deb파일을 만들어서 거의 원클릭수준으로 매우 쉽게 설치할 수 있다. 매우매우 좋다.  

일단 업데이트부터 해주자
```
sudo apt-get update
sudo apt-get upgrade
```

다음 HTTP 프로토콜을 이용하여 다운로드 받을수 있는 wget 명령어를 사용해 [GitHub](https://github.com/dltpdn/opencv-for-rpi/releases) 링크에서 다운로드 한다.
```
wget https://github.com/dltpdn/opencv-for-rpi/releases/download/4.2.0_buster_pi3b/opencv4.2.0.deb.tar
```

다운로드받은 tar 파일의 압축을 풀어준다.
```
tar -xvf opencv4.2.0.deb.tar
```

다음, deb파일들로 OpenCV4를 설치한다. 
```
sudo apt-get install -y ./OpenCV*.deb
```
-y 옵션은 설치할 때 Y/N을 고르는 단계를 없애는것이다. 여러개를 한번에 인스톨할 때 좋다.  

이러면 설치가 완료됐다. 잘 됐는지 확인해보자.
```
pkg-config --modversion opencv4
>> 4.2.0
```

```
python
import cv2
cv2.__version__

>> 4.2.0
```

이러면 잘 설치가 됐다.  
내일은 Mask Detection 라이브러리에 대해 알아보려고 한다.  
그리고 라즈베리파이->컴퓨터로 이미지를 옮기는 데 조금 귀찮음이 있다.  
내 블로그에 이미지 업로드시스템을 만들어볼까 생각중이다.  

#### 여담
카메라 성능 별로인거같다...
