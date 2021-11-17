---
layout: post
title: 라즈베리파이에서 Face-Mask-Detection 라이브러리 사용하기
category: RaspberryPi
tag: [RaspberryPi, OpenCV]
---
 
# Using Face-Mask-Detection Library in Raspberry Pi 3 B

이 글은 기본적인 설치와 예제 코드만을 다룬다.

[Face-Mask-Detection 라이브러리](https://github.com/chandrikadeb7/Face-Mask-Detection) 를 라즈베리파이에서 사용하려고 했는데 많은 에로사항이 있었다.

예제대로 하면 모델 훈련 -> 검증 순서로 하는데, 모델 훈련을 라즈베리파이의 열악한 환경 속에서 하는것도 무리가 있고...  
또한 텐서플로의 버전도 라즈베리파이 OS 안에서는 1.4까지밖에 릴리즈가 안되어있다. 따라서 `pip3 install -r requirements.txt`를 하면 처음부터 뻑난다..  
이제부터 이 문제들의 해결법을 다룰 계획이다.  

#### 모델 만들기
성능 좋은 데스크탑 환경에서 만들었다.
![train](https://github.com/ukcastle/ukcastle.github.io/blob/main/postimg/210118TiL/trainModel.png?raw=true)

 
![avg](https://github.com/ukcastle/ukcastle.github.io/blob/main/postimg/210118TiL/avg1.png?raw=true)

매우 정확한 모델이 만들어졌다.

[다운로드 링크](https://github.com/ukcastle/frames-client/raw/main/mask_detector.model)  
뭐 굳이 다시 돌릴 필요가 있는가? 잘 만들어진거 다운로드 받아서 쓰자  
이걸로 21-01-18 기준 데이터셋을 이용한 99% 정확도의 모델이 만들어졌다.  
<br>

#### 라즈베리파이에서 라이브러리 다운로드하기

실행하기 전 ***가상환경***을 만들고 진행하는걸 추천한다.  
라즈베리파이에 Python3가 깔려있고, 아무것도 없는 상태라고 가정한다.  
설치해야할 라이브러리들은 다음과 같다.   
```
tensorflow>=1.15.2
keras==2.3.1
imutils==0.5.3
numpy==1.18.2
opencv-python==4.2.0.*
matplotlib==3.2.1
argparse==1.1
scipy==1.4.1
scikit-learn==0.23.1
pillow==7.2.0
streamlit==0.65.2
```

여기서 tensorflow가 오류가난다. 왜냐면 라즈베리파이OS에서 정식 릴리즈된게 1.4까지밖에 없어서...  
그래서 [이 사이트](https://towardsdatascience.com/3-ways-to-install-tensorflow-2-on-raspberry-pi-fe1fa2da9104)를 보고 따라했다.  

```
pip3 install https://github.com/bitsy-ai/tensorflow-arm-bin/releases/download/v2.4.0-rc2/tensorflow-2.4.0rc2-cp37-none-linux_armv7l.whl
```

이러고 `pip3 list` 를 쳐보면 `tensorflow==2.4.0rc2`인 것을 확인할 수 있다. 
![1](https://github.com/ukcastle/ukcastle.github.io/blob/main/postimg/210118TiL/111.png?raw=true)
![2](https://github.com/ukcastle/ukcastle.github.io/blob/main/postimg/210118TiL/222.png?raw=true)
![3](https://github.com/ukcastle/ukcastle.github.io/blob/main/postimg/210118TiL/333.png?raw=true)

이런 다음 `requirements.txt`에 적혀있는 내용을 수정해준다.
```
keras==2.3.1
imutils==0.5.3
numpy==1.18.2
opencv-python==4.2.0.*
matplotlib==3.2.1
argparse==1.1
scipy==1.4.1
scikit-learn==0.23.1
pillow==7.2.0
streamlit==0.65.2
```

다음 `pip3 install -r requirements.txt` 를 실행해준다.

numpy같은 경우 이미 1.9.*가 깔려있어 그냥 넘어가기도 하고.. 경고가 몇몇 뜨는데 넘어가도 되는 것 같다.  

다음 detect_mask_image.py 파일을 돌려볼건데 OpenCV의 파이썬 문법에서 되는가?싶은것들이 있다.  
오류가 좀 나던데..  np.copy, np.shape(image)[:2] 이렇게 바꿔주면 잘 된다.  
[여기](https://github.com/ukcastle/frames-client/blob/main/detect_mask_image.py) 에 있는 py 파일을 다운받아도 되고 저것만 바꿔줘도 된다.  
![in](https://github.com/ukcastle/ukcastle.github.io/blob/main/postimg/210118TiL/in.jpg?raw=true)
![out](https://github.com/ukcastle/ukcastle.github.io/blob/main/postimg/210118TiL/out.jpg?raw=true)

이런식으로 라즈베리파이에서도 잘 나온다.

다음엔 이걸 실제로 카메라와 연동시켜봐야겠다.

#### 여담

설치때문에 발이 너무 묶여있었다.. 