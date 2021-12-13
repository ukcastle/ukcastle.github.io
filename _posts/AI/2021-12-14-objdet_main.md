---
layout: post
title: Object Detection 회고
category: AI
tag: [Object Detection, Pytorch, Detectron] 
---

## 개요

- Naver BoostCamp AI Tech 2nd. 교육 중 진행한 프로젝트  
- 거리 위 쓰레기 사진를 Object Detection 하여 쓰레기의 분리수거 유무를 AI가 판단하여 사회적 환경 부담 문제를 줄일 수 있는 방법 연구, 개발  
- Metric : mAP50
- **(개인)** Detectron2, Yolo v5 프레임워크를 이용한 연구 파이프라인 생성, 모델 성능 향상을 위한 가설 및 검증
- [Github](https://github.com/boostcampaitech2/object-detection-level2-cv-04)
## 결과

- 부스트캠프 내 18팀 중 3위
- 검증 사진은 공개 불가능

## 문제 정의, 해결

- 추후 작성 예정

## 회고

회고는 중요하고 프로젝트를 진행한 지는 2개월이 지났지만 바쁜 일정으로 정리를 하지 못했다. 그냥 잠을 아껴서 작성하고 자야겠다.  
Pytorch 기반, 그리고 Obj Det 모델을 연구하기 위한 고수준의 프레임워크는 크게 MM-det, Detectron2, Yolo v5 등이 있다.  
고수준의 프레임워크의 공통점은 **개발을, MLOps를 잘 몰라도 사용하기 쉽다**, 그리고 **제공되는 기능이 아닌 기능을 추가적으로 사용하려면, 시각화하려면 힘든 경향이 있다** 라고 생각한다.  
이번 대회에서는 AI쪽으로는 큰 발전은 없었다고 생각했지만, 이후 여러 고수준의 프레임워크를 사용하며 느낀 점을 나의 연구 파이프라인 개발 과정에 잘 녹여 MLOps적으로 큰 발전이 있었다고 생각한다.   
