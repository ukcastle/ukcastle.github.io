---
layout: post
title: 21-01-12-TiL 
category: TiL
tag: [RaspberryPi, OpenCV]
---

# Using OpenCV and Face Recognition in RaspberryPi

이제 본격적으로 클라이언트 설계를 해보려고 한다.  
클라이언트의 요구사항은 다음과 같다.  
1. 얼굴인식 기능
> 1-1. 얼굴 인식에 성공하면, Face Recognition 라이브러리를 사용하여 특징값을 추출한다.
> 1-2. 서버에 전송한 뒤 결과값을 받고, 성공 실패 시나리오를 만든다.  
> 1-3. 그 외 기타 기능들을 만든다...  

2. 적외선 열 감지 기능을 만든다. 
> 2-1. 얼굴인식이 성공하고 Face Recognition 라이브러리를 사용할 때 열 체크를 한다.
> 2-2. 졸작 세션 교수님의 말에 따르면, 피부의 온도를 측정할 때에 많은 오류들이 있다고 한다.  
> 2-3. 체온을 날씨, 환경에 따라 **보정**해서 측정해야하는가.. 고민을 해야된다.  

3. 시각적인 기능을 만든다.
> 3-1. 실패 시나리오에서 QR코드를 표시해야 한다.
> 3-2. 추가로 여러가지 액션을 생각할 수 있을 것이다. 

크게보면 이렇게 3가지 기능이 있다. 일단 첫번째부터 시작해보자. 아무리못해도 이번달 안으로 1번은 끝낼 예정이다.  

첫 번째에서 주의해야 할 점은, 일단 중복된 데이터를 전송하면 안된다. 서버에 송신한 뒤로는 대기상태에 있어야겠다. 

일단 설계는 여기까지 하고 본론으로 들어가보자. 

#### Install OpenCV

전체적으로 [여기](https://blog.xcoda.net/97)를 참조했다.
미리 빌드된 .deb 파일로 OpenCV3 버전을 금방 설치할 수 있는 장점이 있다. 와우  
아무튼 [GitHub](https://github.com/dltpdn/opencv-for-rpi)페이지를 보고 따라해봤다.  

```
sudo apt update
sudo apt upgrade

git clone https://github.com/dltpdn/opencv-for-rpi.git

cd opencv-for-rpi/stetch/(원하는 버전)
sudo apt install -y ./OpenCV*.deb

pkg-config —modversion opencv 
```
이렇게 따라오면, 성공적으로 OpenCV가 설치된다.


#### 이후 고민점
이제, Face Recognition을 사용해야 하는데... 고민을 해봐야 한다.
내가 테스트해본건 전부 **한 프레임의 이미지**이다.
비디오로, 만약 초당 10프레임의 이미지라고 하면 이 것을 전부 서버로 보내는건 매우매우 비효율적이다.  
그래서 일단 Face Recognition 라이브러리를 살펴보고, 이 문제점을 해결할 수 없을때를 대비하여 이미지를 처리할 수 있는 다른 라이브러리를 찾아봤다.  
바로 [haar feature](https://github.com/opencv/opencv/tree/master/data/haarcascades) 라이브러리를 사용하는 것  
이 라이브러리는, 실시간 비디오 이미지에서 얼굴이 존재하는지 유무를 판별하는 라이브러리이다.  
따라서 이 라이브러리로 얼굴이 있을 때 얼굴의 이미지만 사진으로 남겨(데이터를 최소화하기 위하여) Face Recognition 라이브러리에 접목시키면 효과적이지 않을까? 생각을 해봤다.  

공부한것을 모두 내일 테스트해볼 예정이다. 화이팅!  


