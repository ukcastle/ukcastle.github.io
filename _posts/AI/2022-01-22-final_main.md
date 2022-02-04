---
layout: post
title: Handwash Detection 회고
category: [AI, BC]
tag: [Pytorch] 
---

## 개요

- Naver BoostCamp AI Tech 2nd. 교육 중 진행한 프로젝트  
- 21.11.27.~21.12.24.
- 기획부터 데이터수집, 모델 연구, 서비스까지 End-to-End로 진행된 프로젝트  
- 정부 권장 손씻기 6단계를 분류하여 실시간으로 손씻기를 얼마나 잘 했는지 판단해주는 서비스 개발  
- 추후 AR을 접목하여 아이들 교육용으로 사용하거나, 정확도를 더 향상시켜 병원, 식당 내 검수용 프로그램 개발 계획 

## 결과  

- [Github repo](https://github.com/boostcampaitech2/final-project-level3-cv-04)  
- Demo  
  ![image](https://github.com/boostcampaitech2/final-project-level3-cv-04/raw/main/src/demo_2x.gif)
- Overall Architecture  
  ![image](https://github.com/boostcampaitech2/final-project-level3-cv-04/raw/main/src/service_architecture.png)  
- [발표 자료](https://github.com/boostcampaitech2/final-project-level3-cv-04/blob/main/docs/Networking%20Day%20%EB%B0%9C%ED%91%9C%EC%9E%90%EB%A3%8C%20%EC%B5%9C%EC%A2%85%EB%B3%B8.pdf)
  

## 문제 정의, 해결

!! 추가중 !!

- 영상에서 이미지를 추출한 뒤 Train, Valid를 나누자. 용량이 너무 커서 오래걸린다. 
  - [멀티프로세싱을 이용한 데이터셋 제작](https://ukcastle.github.io/python/2022/01/22/video_2_image/)

- 훈련 파이프라인 구축
  - [깔쌈하고 아름다운 Pytorch 훈련 파이프라인](https://ukcastle.github.io/ai/2022/01/03/pytorch_baseline/)
  - 파이썬 데코레이터를 이용한 WandB 연결하기 

## 회고