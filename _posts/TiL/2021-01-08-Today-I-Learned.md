---
layout: post
title: 21.01.08 기록
category: TiL
tag: [github, blog,WSL, jekyll]
---

# GitHub blog with Jekyll in Window 10
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

#### Install WSL

WSL을 사용하기위해 설정해야 할 점이 몇가지 있다.
> 1. Windows 기능 켜기
> 2. Microsoft Store에서 우분투 다운로드
> 3. 계정 생성 후 권한 주기
> 4. 기초 설정


1. 첫번째로 Windows 기능 켜기/끄기 탭에서 **Linux용 Windows 하위 시스템** 을 체크 해줘야 한다.
![기능 켜기](.docsImage/210108_1.jpg)  
그 다음 재부팅을 하라고 알림이 뜨면 하고 오면 된다.

2. 두번째로 


[WSL Bash Shell](https://forbes.tistory.com/543)

