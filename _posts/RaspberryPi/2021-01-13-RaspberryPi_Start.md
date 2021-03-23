---
layout: post
title: 라즈베리파이 한글 설치와 와이파이 연결 
category: RaspberryPi
tag: [RaspberryPi]
---

# Install RaspberryPI and Search WiFi

라즈베리파이를 실제로 실행해보고 집에서 혼자 쉐도우복싱 해본채로 실전으로 들어갔는데 여러가지 문제점이 발생했다.

#### LCD 패널 액정이 작동을 안한다  
 연구실에 있던 모니터 하나를 사용했다...

#### WiFi 설정에서 문제가 좀 있었다.  
 가장 좋은 방법은 일단 랜선을 연결하는 것이다. 하지만 사정상 그러지 못하고 핸드폰으로 핫스팟을 켰는데 어째서인지 라즈베리파이에서 잡지를 못하였다.  
 연구실에 랜선 포트가 고장나서 어떻게 할까 고민하다가 그냥 지원팀을 불러서 해결했다  
 그래서 결국 WiFi SSID와 비밀번호를 알고난 뒤 연결했다.  
 <br>
 일단 첫번째로 `설정 -> Raspberry Pi Configuration -> Localisation` 에 들어간다.  
 `Set WiFi Country`를 누르고 `US`로 바꿔준다. KR로 하니까 와이파이가 안잡히더라.  
 <br>
 두번째로 터미널에서 `sudo iwlist wlan0 scan` 을 입력한다.  
 여러가지가 뜨는데, `Cell 숫자 - Address:...` 로 구분을 한다.  
 ESSID쪽을 보며 내가 연결할 와이파이가 있나 찾아본다. 2.4GHz 주파수의 와이파이만 검색이 된다. 동글쓰면 5GHz로 설정할 수 있을까?  
 그 다음 ESSID를 잘 기억해놓는다.  
 <br>
 마지막으로 터미널에 
 `sudo nano /etc/wpa_supplicant/wpa_supplicant.conf` 를 입력한다.  
 파일 안에는 이런 문자열들이 있을 것이다.  
```
ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev
update_config=1
country=US
```

이렇게 돼있을 것이다. 아래에 코드를 추가해준다
    
```
ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev
update_config=1
country=US

network={
    ssid="와이파이 이름"
    psk="와이파이 비밀번호"
}
```

다음 재부팅을 하고 와이파이가 정상적으로 연결이 되었는지 확인한다. 
가장 좋은 방법은 우측 위에 와이파이 표시가 있는지 확인하는 것이고...  
![이미지](https://github.com/jo631/jo631.github.io/blob/main/postimg/%EC%99%80%EC%9D%B4%ED%8C%8C%EC%9D%B4.jpg?raw=true)  
 터미널에 `ping 8.8.8.8`을 입력해도 된다.  
또는 터미널에 `iwconfig`를 입력해도 된다.
![이미지](https://github.com/jo631/jo631.github.io/blob/main/postimg/iwconfig.jpg?raw=true)  
연결이 정상적으로 됐다면, **평문**으로 쓰인 비밀번호를 바꿔주자. (귀찮으면 안해도 된다)   
터미널에 `wpa_passphrase 와이파이이름 와이파이비밀번호` 를 입력한다.
그럼 익숙한 형식이 보인다.
```
network={
    ssid="와이파이 이름"
    #psk="와이파이 비밀번호" (주석이니 삭제해도 됨)
    psk=해싱된 비밀번호
}
```
이걸 복붙해서 아까 경로의 wpa_supplicant.conf에 넣어주면 된다. 다음 재부팅해서 확인해보자.  
사실 바로 이 과정으로 넘어와도 됐는데, 나는 이상하게 이 부분에서 에러가 나서 저렇게 한번 테스트해보고 하는게 좋더라.  
    
#### 한글이 안된다.
와이파이까지 연결 성공했으면, 이제 간단하다. 일단 인터넷에 연결됐으니 업데이트부터 해주자.
```
sudo apt update
sudo apt upgrade
```
다음 `설정->Raspberry Pi Configuration Localisation` 에 들어가 `Set Locate` 를 클릭하고 `언어`를 ko, `Chracter Set`을 UTF-8로 설정해준다.  
또한 `Set Keyboard` 에서 `Variant`에서 `Korean(101/104 key compatible)`로 설정을 해준다.   
그리고 재부팅을 하면 한글이 깨진다. 휴지통이 무슨 이상한 언어로 되어있다.  
터미널을 킨다.
```
sudo apt-get install ibus
sudo apt-get install ibus-hangul
sudo apt-get install fonts-unfonts-core

reboot
```

다음, 설정에 보면 `iBus 환경 설정` 이라는게 생겼을 것이다. 이게 한글을 입력하게 해주는건데, 가끔 이런 에러가 뜰 수가 있다.  
` IBus 데몬이 5초 이내에 시작하지 못했습니다. `
권한이 루트유저로 되어있어서 그렇다. 터미널에서 이를 입력하자.  
```
sudo rm -rf .config/ibus
```
rm -rf를 칠땐 항상 조심조심하면서...  
그 다음 설정을 눌러보면 잘 실행된다.
설정에서 다음 명령어도 입력해주자.  
```
im-config -n ibus
```
부팅시 자동으로 iBus 데몬을 실행시키는 것이다.

그러면 우측 위에 키보드모양 어플리케이션이 생겼는데, 그걸 눌러주고 태극문양으로 바꿔주자.
다음 한영 또는 Shift+Space를 눌러주면 한글 입력이 잘 되는 것을 알 수 있다.


#### 여담
생각보다 분량이 많아, OpenCV는 다음 포스트에 다뤄야겠다.