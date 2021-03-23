---
layout: post
title: 21-01-08-TiL
category: Jekyll
tag: [WSL, Jekyll]
---

# Intall Jekyll in Window 10
깃허브 블로그를 만들어보려고 사이트 개발 툴을 알아보던 도중 **Jekyll**이란 툴이 유용하고 편하다고 들어서 사용해보기로 결심했다.  

#### What is Jekyll?
[Jekyll Official GitHub](https://jekyllrb-ko.github.io/)  

지킬은 **정적**(static) 사이트로 PHP 언어와 같은 서버 소프트웨어를 사용하지 않고 오직 HTML, CSS 등의 정적 파일만을 사용하여 사이트 생성이 가능한 툴이다.  
동적(dynamic) 사이트로 대표적인 **워드프레스**는 현재에도 많이 사용중인데, 둘의 차이점은 다음과 같다.
1. 워드프레스는 지킬에 비해 기능이 많지만 무겁다
2. 워드프레스는 과한 트래픽에 약하다
3. 워드프레스는 느리고 비싸다

다음과 같은 차이점으로 오직 **블로깅** 에만 초점을 맞추면 워드프레스보다 지킬을 사용하는 것이 개발자들에게 유리하다고 볼 수 있다.  

물론 지킬은 정적인 사이트인 만큼 동적인 사이트보다 기능적인 측면에서 불리할 수도 있다. 이런 장단점을 생각하고 결정을 하는것이 좋을 것 같다.

#### Start

개발을 시작해보려고 한 뒤 처음으로 봉착한 난관은 지킬은 **Ruby**라는 언어를 사용한 프레임워크이고 루비는 (물론 윈도우에서도 잘 되지만)리눅스 환경에서 더 편한 언어이다. 따라서 불편하게 윈도우용으로 설치하는가 리눅스를 설치하여 쓰는가 두가지 선택지가 있었는데, Windows Subsystem for Linux(일명 **WSL**)를 사용하여 개발하기로 결정했다.

### Install WSL

WSL을 사용하기위해 설정해야 할 점이 몇가지 있다.
> 1. Windows 기능 켜기
> 2. Microsoft Store에서 우분투 다운로드
> 3. 계정 생성 후 권한 주기
> 4. 기초 설정


##### 1.  Windows 기능 켜기

첫번째로 Windows 기능 켜기/끄기 탭에서 **Linux용 Windows 하위 시스템** 을 체크 해줘야 한다.
![Linux용 Windows 하위 시스템](https://github.com/jo631/jo631.github.io/blob/main/postimg/%EA%B8%B0%EB%8A%A5%EC%BC%9C%EA%B8%B0.jpg?raw=true)

그 다음 **재부팅**을 하라고 알림이 뜨면 하고 오면 된다.        
<br/>

##### 2. Microsoft Store에서 우분투 다운로드

두번째로 **Microsoft Store**에서 **Ubuntu**를 검색하자
![우분투 설치](https://github.com/jo631/jo631.github.io/blob/main/postimg/%EC%9A%B0%EB%B6%84%ED%88%AC%EC%84%A4%EC%B9%98.jpg?raw=true)
다음 설치를 해준 뒤 실행을 해주면 된다.
<br/> 

##### 3. 계정 생성 후 권한 주기
설치가 된 뒤에 Ubuntu를 실행하면 익숙한 Bash 콘솔 창이 표시된다. 다음 id와 password를 지정하라고 나오고, 이 때 **root**나 **admin** 등 이미 사용중인 계정은 id 설정이 **불가능**하다  
일단 자주 사용하는 ID와 PW를 설정한 뒤 CMD를 **관리자 권한** 으로 실행한 뒤 다음과 같은 커맨드를 입력한다.   
```
ubuntu.exe config --default-user 설정한 아이디
```
 
그 뒤 **서비스** 에 들어가 **LxssManager**을 다시 시작 한다.
![다시 시작](https://github.com/jo631/jo631.github.io/blob/main/postimg/%EB%8B%A4%EC%8B%9C%EC%8B%9C%EC%9E%91.jpg?raw=true)

이 다음 Ubuntu를 재시작한다. 이제 sudo 커맨드를 사용할 수 있다.
<br/>

##### 4. 기초 설정
WSL은 윈도우 위에 설치되어 있다. 하지만 그렇다고 윈도우의 앱을 사용할 수 있지는 않다.  
일단 기초 설정부터 하자.
`sudo apt update`
`sudo apt install gcc` 
이것으로 WSL 설치와 기초 설정을 끝냈다. 

[참고 자료: 윈도우10 Linux Bash Shell 설치 및 사용법](https://forbes.tistory.com/543)

### Install Jekyll

Jekyll을 설치 전 Ruby를 설치해야 한다.
```
sudo apt-get install rubygems
또는
sudo apt-get install libgemplugin-ruby
```

이것으로 **gem** 키워드를 사용할 수 있다.

다음으로 Jekyll을 설치하는 명령어다.
```
sudo gem install bundler
sudo gem install bundler jekyll
```

설치한 다음 GitHub Blog가 설치된 경로로 간다.  
WSL에서는 `./mnt/드라이브`에 경로가 저장되어있기는 한데, 이렇게 가긴 귀찮다.  
다른 방법이 없을까?


### Using PowerShell
윈도우에 깔려있는 PowerShell을 켜보자. 다음 `bash` 명령어를 친다.  
![배쉬](https://github.com/jo631/jo631.github.io/blob/main/postimg/bash.gif?raw=true)
그럼 자동으로 BashShell로 전환된다!

해당 기능을 이용하여 편하게 BashShell을 실행해보자

![파워쉘](https://github.com/jo631/jo631.github.io/blob/main/postimg/powershell.jpg?raw=true)

여기서 실행한 뒤 `bash` 커맨드를 입력한다. 다음 본격적으로 Jekyll을 이용해보자.

해당 과정은 템플릿을 다운받지 않았을 때 초기 설정하는 단계이다. 이미 템플릿을 받았다면 넘어가도 된다.
```
bundle init
bundle add jekyll
```
폴더에 **Gemfile**이라는 파일과 뒤에 Lock이 붙은 파일 총 두개가 있을 것이다.

여기까지 하면 기본 설정이 끝났다!  
이제 서버를 실행시켜보자
```
jekyll serve
```
![서버실행](https://github.com/jo631/jo631.github.io/blob/main/postimg/%EC%84%9C%EB%B2%84%EC%8B%A4%ED%96%89.jpg?raw=true)
문제가 없다면 위와 같이 표시된다. 
다음 웹 브라우저를 실행하고 **localhost:4000** 에 접속하면, 홈페이지가 나온다.

여기까지 Jekyll 프레임워크의 기본 설정이 끝났다.  
이 뒤로는 원하는대로 커스터마이징 하면 된다.
<br/>
<br/>

### 여담
분명 처음 목표는 라즈베리파이와 카메라, 적외선카메라를 이용한 클라이언트 개발이였는데...  
아직 홈페이지도 고칠게 많다. 내일은 시작할 수 있을까?