---
layout: post
title: 재활용 쓰레기 Object Detection 회고
category: [AI, BC]
tag: [Object Detection, Pytorch, Detectron2] 
---

## 개요

- Naver BoostCamp AI Tech 2nd. 교육 중 진행한 프로젝트  
- 21.09.30~21.10.16
- 거리 위 쓰레기 사진를 Object Detection 하여 쓰레기의 분리수거 유무를 AI가 판단하여 사회적 환경 부담 문제를 줄일 수 있는 방법 연구, 개발  
- Metric : mAP50
- **(개인)** Detectron2, Yolo v5 프레임워크를 이용한 연구 파이프라인 생성, 모델 성능 향상을 위한 가설 및 검증
- [Github](https://github.com/boostcampaitech2/object-detection-level2-cv-04)
## 결과

- 부스트캠프 내 18팀 중 3위
- 검증 사진은 공개 불가능

## 문제 정의, 해결

- **Detectron2 사용법을 알아보기**  
  기본 예제를 보면서 사용법을 [여기](https://ukcastle.github.io/bc/2021/09/28/w9d2/)와 [저기](https://ukcastle.github.io/bc/2021/09/29/w9d3/)에 정리했다.  
	다만 후술할 내용이지만 입력받는 변수를 확정지어버리면 위험하다.  

- **Detectron2 Repeated Sampler 사용하기**  
  torch의 weighted sampler와 같은 개념인것같다.  
	기본적으로 제공하고 있지만 사용법이 별로 없어 처음 사용할 때 난항을 겪었다. 이를 [여기](https://ukcastle.github.io/bc/2021/09/30/w9d4/)에 정리하였다.  
	

- **Detectron2 관점에서 MM-Detection을 같이 사용하며 연구할 때 WandB 연동하기**  
	Detectron2는 21.10. 기준 WandB를 제공해주지 않는다. 하지만 연구의 효율을 높이기 위하여 직접 연동하여 [여기](https://ukcastle.github.io/bc/2021/10/05/w10d1/)에 정리해놨다.  
	
- **Detectron2 출력창 깨끗이 하기**
	Detectron2는 Logger를 집중적으로 사용하며 `torch.nn.Module`의 멋진 기능인 Hook을 자유자재로 사용하며 구현해놨다. 하지만 솔직히 말하면 터미널이 너무 지저분해져서 싫었다.  
	그래서 출력을 최대한 없애고 `tqdm`을 설정해논 내용을 [여기](https://ukcastle.github.io/bc/2021/10/06/w10d2/)에 정리해놨다.

- **Detectron2 Optimizer Custom하기**  
  고수준 API의 단점이 명백히 있다. 새로운 기능이 나와도 업데이트하지 않으면 잘 사용할수가 없다. 이를 수정하는 메소드를 [여기](https://ukcastle.github.io/bc/2021/10/07/w10d3/)에 정리해놨다.  

- **Detectron2를 추상화할때 주의할 점**  
  프레임워크의 전체적인 구조는 각 모델 별 yaml 파일이 있고 이를 불러와서 사용하는 흔한 개념이다.  
	그래서 좀 더 사용하기 편하게 커스텀하는 과정에서 하나의 yaml 파일을 잡고 [이렇게](https://ukcastle.github.io/bc/2021/09/27/w9d1/)구조를 확정지어버렸는데, 다른 모델을 돌릴 때 엄청난 에러가 쏟아졌다. 하나하나 디버깅해가며 찾아 본 결과 yaml파일마다 저장된 key 값이 달랐다.  
	그래서.. 이를 **하나의 yaml 파일로 매개변수를 고정해버리면 대참사**가 일어난다.  

- **1 Stage 와 2 Stage의 Ensemble**  
  다른 팀들에 비해 가장 차별화됐던 우리 팀만의 기법이였다.  
	Inference의 시간이 제한이 없고 csv파일을 제출하는 대회의 특성상 **가장 무겁고 깊은 모델들의 앙상블**이 대회의 해답이라고 생각했다.  
	따라서 대회의 중후반, 팀은 2분류로 나뉘어 연구를 진행했고 앙상블을 했다.  
	1 Stage의 mAP 점수가 평균 0.1정도 낮았지만 앙상블을 진행했을 때, 기존 2 Stage만을 앙상블했던 때 보다 5배는 높은 효과를 보여줬다.  
	이로써 **1 Stage와 2 Stage Detector가 서로 다른 방향으로 물체를 측정하고 이를 앙상블 할 때 좋은 효과를 보일 가능성이 크다**라고 추론했다.  

## 회고

회고는 중요하고 프로젝트를 진행한 지는 2개월이 지났지만 바쁜 일정으로 정리를 하지 못했다. 그냥 잠을 아껴서 작성하고 자야겠다.  
Pytorch 기반, 그리고 Obj Det 모델을 연구하기 위한 고수준의 프레임워크는 크게 MM-det, Detectron2, Yolo v5 등이 있다.  
고수준의 프레임워크의 공통점은 **개발을, MLOps를 잘 몰라도 사용하기 쉽다**, 그러나 **제공되는 기능이 아닌 기능을 추가적으로 사용하려면, 시각화하려면 힘든 경향이 있다** 라고 생각한다.  
이번 대회에서는 AI쪽으로는 큰 발전은 없었다고 생각했지만, 이후 여러 고수준의 프레임워크를 사용하며 느낀 점을 나의 연구 파이프라인 개발 과정에 잘 녹여 MLOps적으로 큰 발전이 있었다고 생각한다.   
