---
layout: post
title: 리눅스 Could not get lock /var/lib/dpkg/lock-frontend
category: RaspberryPi
tag: [RaspberryPi, OpenCV, Keras] 
---

Odroid에 리눅스 설치하자마자 에러가 났다.  

```bash
$sudo killall apt apt-get

진행중인 프로세스가 없다고 뜨면

$sudo rm /var/lib/apt/lists/lock
$sudo rm /var/cache/apt/archives/lock
$sudo rm /var/lib/dpkg/lock*
```

실행한 다음  

```bash
$sudo dbkg --configure -a
$sudo apt update 
```

를 하면 정상적으로 해결이 된다.  