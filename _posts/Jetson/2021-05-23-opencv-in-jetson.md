---
layout: post
title: Jetson OpenCV 4.5 Contrib 설치하기
category: Jetson
tag: [Jetson] 
---

#### 들어가기 전

Jetson은 기본적으로 CUDA GPGPU기능이 내장되어있다.  
그리고 JetPack의 버전에 따라 다르지만, 보통 OpenCV 3.x 버전이 내장되어있다.  

하지만, OpenCV에서 CUDA를 지원하는 것은 **4.2.0** 이상 버전부터다.  
사실 Jetson Nano에서 GPU를 사용하지 않는다면 라즈베리파이랑 다른점이 없기에, 필수불가결적으로 높은버전의 OpenCV를 사용해야 한다.  

일단, 설치되어있는 OpenCV를 제거해야 한다.  

다른 곳에서 낮은버전의 OpenCV를 `cmake`를 이용한 방식으로 `make install`을 사용해서 설치했다면,  

높은 버전의 OpenCV를 받아도 경로가 꼬여 에러가 발생한다. (이전버전을 깔끔하게 못지워 경로가 남아있으면, 최신버전을 불러오지 못한다.)  

가장 좋은 방법은 낮은버전의 OpenCV를 설치한 폴더의 `build`에서 `make uninstall`을 하면 된다.  

만약 설치폴더를 지웠다면? 많은 시도를 해봤기는 했는데... 그냥 다시 구버전을 설치하고 `make uninstall`을 하는게 빠르다..  

```bash
$ pkg-config --modversion opencv
$ pkg-config --modversion opencv2
```

해당 명령어를 입력하면 opencv가 설치되어있는지 확인할 수 있다.

```
Package opencv was not found in the pkg-config search path.
Perhaps you should add the directory containing `opencv.pc'
to the PKG_CONFIG_PATH environment variable
No package 'opencv' found
```
이렇게 나온다면 없다고 생각할 수 있지만, path가 삭제가 안되고 구버전으로 남아있을 경우가 있다. 그래서 사실 다시 install -> uninstall이 가장 맘편하다. path까지 다 지워주기 때문에...  

jtop의 GPU 탭에서도 확인할 수 있다.  

#### 시작하기

모든 절차는 [여기](https://qengineering.eu/install-opencv-4.5-on-jetson-nano.html)를 참고했다.


```
$ sudo apt-get update
$ sudo apt-get upgrade
```

가장 기본적인 절차다.  

#### 스왑 메모리 확장

Jetson Nano의 기존 스왑메모리는 2GB이다.  

기존의 4GB RAM + 2GB Swap Memory로는 OpenCV 4.x를 설치하는데 문제가 있다.  

따라서, 스왑메모리를 4GB까지 확장시켜주어야 한다.  

```
$ sudo apt-get install dphys-swapfile
$ sudo vim /etc/dphys-swapfile
```

![이미지](https://qengineering.eu/gallery/swap2jetson.webp)

이렇게 바꿔주면 된다.  

그 다음 재부팅을 해준다.  

```
$ sudo reboot
```

그 다음 스왑메모리가 잘 늘어났는지 확인하는 절차다.  

```
$ free -m
```

![img](https://qengineering.eu/images/Free-m.webp) 

4GB인 것을 확인하자.  

#### 의존성

```
$ sudo sh -c "echo '/usr/local/cuda/lib64' >> /etc/ld.so.conf.d/nvidia-tegra.conf"
```

CUDA 라이브러리 경로를 등록한다.  

그 다음, 아래의 라이브러리들을 모두 설치해준다.

```
$ sudo apt-get install build-essential cmake git unzip pkg-config
$ sudo apt-get install libjpeg-dev libpng-dev libtiff-dev
$ sudo apt-get install libavcodec-dev libavformat-dev libswscale-dev
$ sudo apt-get install libgtk2.0-dev libcanberra-gtk*
$ sudo apt-get install python3-dev python3-numpy python3-pip
$ sudo apt-get install libxvidcore-dev libx264-dev libgtk-3-dev
$ sudo apt-get install libtbb2 libtbb-dev libdc1394-22-dev
$ sudo apt-get install libv4l-dev v4l-utils
$ sudo apt-get install libgstreamer1.0-dev libgstreamer-plugins-base1.0-dev
$ sudo apt-get install libavresample-dev libvorbis-dev libxine2-dev
$ sudo apt-get install libfaac-dev libmp3lame-dev libtheora-dev
$ sudo apt-get install libopencore-amrnb-dev libopencore-amrwb-dev
$ sudo apt-get install libopenblas-dev libatlas-base-dev libblas-dev
$ sudo apt-get install liblapack-dev libeigen3-dev gfortran
$ sudo apt-get install libhdf5-dev protobuf-compiler
$ sudo apt-get install libprotobuf-dev libgoogle-glog-dev libgflags-dev
```

순서대로 다 설치해줘야 한다!!

#### OpenCV 설치

cmake의 path가 기본적으로 HOME 디렉토리로 설정되어있어, 가능한 한 홈 디렉토리에서 작업을 하는 것을 추천한다.  

```
$ cd ~
$ wget -O opencv.zip https://github.com/opencv/opencv/archive/4.5.2.zip
$ wget -O opencv_contrib.zip https://github.com/opencv/opencv_contrib/archive/4.5.2.zip

$ unzip opencv.zip
$ unzip opencv_contrib.zip

$ mv opencv-4.5.2 opencv
$ mv opencv_contrib-4.5.2 opencv_contrib

$ rm opencv.zip
$ rm opencv_contrib.zip
```

opencv 4.5.2버전을 설치하고, opencv와 opencv_contrib로 이름을 바꿔준다.  

```
$ cd ~/opencv
$ mkdir build
$ cd build
```

다음 opencv/build 폴더를 만들고 이동한다.  

```
$ cmake -D CMAKE_BUILD_TYPE=RELEASE \
-D CMAKE_INSTALL_PREFIX=/usr \
-D OPENCV_EXTRA_MODULES_PATH=~/opencv_contrib/modules \
-D EIGEN_INCLUDE_PATH=/usr/include/eigen3 \
-D WITH_OPENCL=OFF \
-D WITH_CUDA=ON \
-D CUDA_ARCH_BIN=5.3 \
-D CUDA_ARCH_PTX="" \
-D WITH_CUDNN=ON \
-D WITH_CUBLAS=ON \
-D ENABLE_FAST_MATH=ON \
-D CUDA_FAST_MATH=ON \
-D OPENCV_DNN_CUDA=ON \
-D ENABLE_NEON=ON \
-D WITH_QT=OFF \
-D WITH_OPENMP=ON \
-D WITH_OPENGL=ON \
-D BUILD_TIFF=ON \
-D WITH_FFMPEG=ON \
-D WITH_GSTREAMER=ON \
-D WITH_TBB=ON \
-D BUILD_TBB=ON \
-D BUILD_TESTS=OFF \
-D WITH_EIGEN=ON \
-D WITH_V4L=ON \
-D WITH_LIBV4L=ON \
-D OPENCV_ENABLE_NONFREE=ON \
-D INSTALL_C_EXAMPLES=OFF \
-D INSTALL_PYTHON_EXAMPLES=OFF \
-D BUILD_opencv_python3=TRUE \
-D OPENCV_GENERATE_PKGCONFIG=ON \
-D BUILD_EXAMPLES=OFF ..
```

**해당 작업은 대략 5분정도 걸린다.**

QT GUI를 사용하려면 `-D WITH_QT=ON`으로 바꾸라고 하던데, PyQT5는 별 차이 없는거같아서.. 이 부분은 잘 모르겠다.  

여기서 만약 홈디렉토리에서 작업을 안하면 오류가 발생할수도 있는데, 꼭 아래와 같은 화면을 보아야 한다.  

![이미지](https://qengineering.eu/images/cmake-opencv.webp)

#### Make

이제, 시간이 오래걸리는 부분이다.

```
$ make -j4
```

`-j4` 옵션은 코어를 4개 사용한다는 의미다.  
만약 본인의 Jetson에 **쿨링팬이 없다면** `-j2`를 하길 추천한다.  

이 부분에서 **4코어 기준 1~2시간정도** 걸렸던거같다.  

![이미지](https://qengineering.eu/images/Make-ready.webp)

이렇게 표시된다면, 컴파일이 성공한것이다.  

```
$ sudo rm -r /usr/include/opencv4/opencv2
```

혹시나 남아있을수도 있는, 이전의 OpenCV 파일들을 삭제한다.  

```
$ sudo make install
```

그 다음 OpenCV를 설치한다. 이 부분에서 **메모리를 상당히 많이 사용**한다. 가급적 다른 프로그램의 실행은 자제하자. 또한 이 부분에서 10분정도 걸렸던것같다.  

```
$ sudo ldconfig
```

공유 라이브러리 캐시를 다시 설정해준다. 다 왔다!!  

#### 확인

이제 OpenCV가 잘 설치되었는지 보자

```
$ Python3
>>import cv2
>>print(cv2.__version__)
```

4.5.2가 출력되면 성공이다.  


#### 마무리

이제 늘려줬던 스왑메모리를 해제해주자.  

```
$ sudo /etc/init.d/dphys-swapfile stop
$ sudo apt-get remove --purge dphys-swapfile
```

다음 설치했던 파일들을 제거하던 말던, 선택이다.  
나는 이전버전을 지우는데 너무 고생을 많이해서.. 남겨둘 예정이다.  

```
$ make clean
$ sudo rm -rf ~/opencv
$ sudo rm -rf ~/opencv_contrib
```

이제 모든 과정이 끝났다!  

#### 구현

이걸 이제 어떻게 사용하느냐...  

자신의 코드에 `net = cv2.dnn.readNet~~(~~~)` 과 같은 코드가 있을것이다.  

```python
net.setPreferableBackend(cv2.dnn.DNN_BACKEND_CUDA)
net.setPreferableTarget(cv2.dnn.DNN_TARGET_CUDA)
```

그 부분에 해당 코드를 적용시켜주면 된다.  

![img](https://github.com/ukcastle/ukcastle.github.io/blob/main/_posts/Jetson/_postimg/210523-opencv/cuda.png?raw=true)

간단한 테스트코드를 돌려봤는데, 기존의 CPU를 사용한 모델 처리는 0.3초가 걸렸는데, CUDA GPGPU 환경에서는 단 0.05초가 걸렸다  

현질 최고