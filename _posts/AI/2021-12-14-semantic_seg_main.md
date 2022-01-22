---
layout: post
title: 재활용 쓰레기 Semantic Segmentation 회고
category: [AI, BC]
tag: [Semantic Segmentation, Pytorch] 
---

## 개요

- Naver BoostCamp AI Tech 2nd. 교육 중 진행한 프로젝트  
- 21.10.19.~21.11.04.
- 거리 위 쓰레기 사진를 Semantic Segmentation 하여 쓰레기의 분리수거 유무를 AI가 판단하여 사회적 환경 부담 문제를 줄일 수 있는 방법 연구, 개발  
- Metric : mIoU
- **(개인)** Pytorch 기반 연구 파이프라인 생성, 모델 성능 향상을 위한 가설 및 검증
- [Github](https://github.com/boostcampaitech2/semantic-segmentation-level2-cv-04)

## 결과

- 부스트캠프 내 18팀 중 3위
- 검증 사진은 공개 불가능

## 문제 정의, 해결
- 투박한 연구 프로세스는 그만...
  - [Pytorch 훈련 베이스라인 생성](https://ukcastle.github.io/ai/2022/01/03/pytorch_baseline/)  

- WandB 차트와 mIoU 계산식을 보고 난 뒤의 고민사항  
  - 어떤 부분때문에 정확도가 낮은지 찾아봤다.  
  - 위치 자체를 잘 못잡기보다는, 종이와 종이팩과 같은 비슷한 클래스들에서 아예 잘못 예측하여 정확도가 많이 떨어졌다.  
  - 따라서 성능이 좋은 백본을 사용해보자 라고 생각했다.    

- Obj Det때부터 느낀점은 우리의 데이터셋에 Swin Transformer가 잘 맞는다. 하지만 smp에는 swinT를 지원해주지 않는다.  
  - [smp에서 swin Transformer를 이식해보자](https://ukcastle.github.io/ai/2022/01/22/smp-swin/)  
  
- 앙상블이 정답인가?  
  - objdet에서의 mAP는 낮은 정확도의 박스들은 엄청나게 큰 영향을 주지 않는다. 하지만 SMP에서 mIoU는 이미지를 2차원적으로 보며 점수를 매기기때문에, 결국 많은 박스들이 Hard-voting 된 값이라고 생각하면 된다.  
  따라서 hard-voting된 값들을 또 다시 앙상블한다면 이는 상당히 위험한 방법이다.  
  앙상블은 하는게 좋지만, 되도록 soft-voting 방식을 이용하는게 좋다.  


## 회고

이번 대회를 진행하면서 Pytorch 기반 코드를 개발자스럽게 활용했다고 생각했고, 멘토님한테도 창의적이고 좋은 시도라고 칭찬받았다!  
Semantic segmentation의 timm이라 할 수 있는 smp 라이브러리 내에 없는 모델들도 백본으로 이식하는 방법도 익혔고, wandb도 잘 이용하여 모델을 연구했고 여러 메트릭에 대한 이해가 늘어나 좋은 경험을 가진 대회라고 생각했다.  