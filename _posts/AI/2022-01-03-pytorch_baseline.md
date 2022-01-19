---
layout: post
title: Pytorch 훈련 베이스라인 생성
category: [AI]
tag: [Pytorch] 
---

## 개요

전체적인 코드는 [여기](https://github.com/boostcampaitech2/final-project-level3-cv-04/tree/main/model_lab/frame_classification)에 있습니다.  
해당 코드를 풀어보는 과정을 적어보려 합니다.  
짧은 개발 기간으로 인해 구상은 했지만 시도하지 않은 점도 있습니다. 해당 내용은 글로만 푸려고 합니다.  

들어가기 앞서, Pytorch의 훈련 프로세스에는 다음과 같은 속성들이 필요합니다.  
- Dataset, Dataloader
- Optimizer, Criterion
- Model
- etc...

ML 프로젝트를 하다보면, 특히 Competition을 목표로 하는 프로젝트를 하는 경우에는, 파이프라인 내 많은 코드를 수정하게됩니다.  
물론 짧은 기간동안 진행하는 프로젝트일 경우, 그리고 혼자 진행할 경우 파이프라인을 대충 만들어도 괜찮을 경우가 있습니다.  
하지만 대규모의 인원으로 여러 방면에서 접근할 경우 이런 방법은 모델의 재구현을 힘들게 만들 수 있습니다.  
물론 mlflow 등 여러 방면으로 해결할 수 있지만, 본 포스트에서는 파이프라인 자체를 잘 모듈화하는 방법을 소개하겠습니다.  

큰 틀은 다음과 같습니다.  
`importlib`을 이용하여 코드의 변화가 잦은 부분은 **동적으로 자동 저장**하고, 코드의 수정이 거의 필요가 없는 부분은 정적으로 이용하는 접근법입니다.  

코드의 수정이 거의 필요가 없는 부분은 **Main, Dataset, Others**이라고 정의했습니다.  
또한, 코드의 수정이 잦은 부분은 **Config, Transform, Dataloader, Optimizer, Model**이라고 정의했습니다. 

## Config (동적)

많은 딥러닝 프레임워크를 사용해보면 항상 Config의 중요성이 부각됩니다. 저는 Config를 [yolov5](https://github.com/ultralytics/yolov5/tree/master/data/hyps)의 형태를 꽤나 본떠 사용했습니다.  
이유는 yaml파일과 Python 라이브러리의 호환성에 있었습니다. 가장 대충 막써도 찰떡같이 알아듣더군요...  

```yaml
root : "../../input"
train_csv : "full_train.csv"
valid_csv : "full_valid.csv"
batch : 128
lr : 0.01
epoch : 50
num_classes : 6

seed : 42
```

간단하게 hyp의 일부만 긁어왔습니다. 전체는 [여기](https://github.com/boostcampaitech2/final-project-level3-cv-04/blob/main/model_lab/frame_classification/custom/full/config.yaml)있습니다.  

이를 `main.py`에서 이렇게 불러옵니다.  
```py
with open(os.path.join("custom",RECIPE,"config.yaml"), "r") as f:
	config = yaml.load(f, Loader=yaml.FullLoader)
```
이러면 config는 `dict`형태로 매핑이 됩니다.  

## Dataset (정적)

데이터셋은 정적인 코드라고 말했습니다.  
물론 변동이 있을수도 있습니다만, 핵심은 이렇습니다.  
**같은 데이터일 경우, 변인을 통제하는 목적에서는 데이터셋이 변화하면 안된다.**  

다만 데이터셋의 경우 마땅한 정답이 없습니다.  
`__init__()`, `__len()__`, `__getitem(index)__` 은 결국 각 태스크마다 다르게 적용해야하니까요.  
아무튼 해당 세 함수를 잘 적용했다 하고 넘어가겠습니다.  

단 여기서, 데이터셋이 변화된다면 어떻게 해야하나? 이는 코드 변화가 많은 Dataloader 내에서 처리할 수 있습니다. 간단하게 미리 예를 들자면 다음과 같습니다.  

```py
def _buildDataloader(self):
		from torch.utils.data import DataLoader
		from src.kaggle_dataset import WashingDataset
		from torch.utils.data.sampler import WeightedRandomSampler
		trainT, validT = self._getTransform()

		tDataset = WashingDataset("train",self.config["root"],self.config["train_csv"],trainT)
```

3번째 줄에서 보면, `src.kaggle_dataset` 이라는 부분을 import 했습니다.  

해당 프로젝트의 데이터셋은 두개가 있었으며([데이터셋 1](https://github.com/boostcampaitech2/final-project-level3-cv-04/blob/main/model_lab/frame_classification/src/dataset.py), [데이터셋 2](https://github.com/boostcampaitech2/final-project-level3-cv-04/blob/main/model_lab/frame_classification/src/kaggle_dataset.py)) 이를 동적으로 잘 입력했습니다.  

## Recipe

훈련에 필요한 클래스들(모델, 옵티마이저, 데이터로더 등)을 저장하고 훈련까지 진행하는 가장 General한 클래스입니다.  
**모든 동적인 부분은 해당 클래스를 상속하는 형태**로 만들었고, `main.py` 내에서는 Recipe을 이용하여 모든 소스들을 생성하게 만들었습니다.  

추상 클래스의 전체적인 [코드](https://github.com/boostcampaitech2/final-project-level3-cv-04/blob/main/model_lab/frame_classification/src/abc.py)는 여기에 있습니다.  

```py
from abc import ABC, abstractmethod
import torch

class AbstractRecipe(ABC):
	
	trainDataloader : torch.utils.data.dataloader.DataLoader = None
	validDataloader : torch.utils.data.dataloader.DataLoader = None
	model : torch.nn.Module = None
	optimizer : torch.optim.Optimizer = None
	loss : torch.nn.Module = None
	scheduler : torch.optim.lr_scheduler._LRScheduler = None

	def __init__(self, config):	
		self.config = config
		self.build()
		self.checkNull()

	@abstractmethod
	def build(self):
		'''
		dataloader, model, optimizer, loss, scheduler(선택) 을 구현해주는 함수입니다.
		해당 함수를 Override 해서 작성해주세요.
		'''
		pass

	def checkNull(self):
		if not self.trainDataloader or not self.validDataloader:
			raise NotImplementedError("Dataloader 구현 안됨")

		if not self.model:
			raise NotImplementedError("Model 구현 안됨")

		if not self.optimizer:
			raise NotImplementedError("Optimizer 구현 안됨")

		if not self.loss:
			raise NotImplementedError("Loss 구현 안됨")

	def getDataloader(self):
		return self.trainDataloader, self.validDataloader

	def getModel(self):
		return self.model
		
	def getOptimizer(self):
		return self.optimizer
	
	def getScheduler(self):
		return self.scheduler

	def getLoss(self):
		return self.loss
```

가장 추상적인 클래스로 많은 필드들의 `getter`를 선언하고, Null Object를 체크하는 단계를 거칩니다.  
다시 한번 말하지만 모든 동적 소스들은 해당 클래스를 상속하여 만들어집니다.  


## Trainer (정적)

나름 생소한 개념입니다. 해당 부분은 [Detectron2](https://github.com/facebookresearch/detectron2)에서 차용했습니다.  
요약하자면, 필요한 소스를 받아 훈련하는 프로세스를 담당하고 평가 Metric과 Logging 등을 담당합니다.  

ML/DL 프로세스에서 상당히 자주 보이는 코드들이므로 자세한 설명은 생략하고 [전체 코드의 링크](https://github.com/boostcampaitech2/final-project-level3-cv-04/blob/main/model_lab/frame_classification/src/trainer.py)를 올리겠습니다.


## Main (정적)

ML/DL 파이프라인에서의 `main.py`의 역할은 대부분 하나일것입니다.  

- 인자를 받아 각종 세팅에 매핑  

메인은 늘 느끼지만, 짧은게 좋습니다.  

전체 코드는 다음과 같습니다.  

```py
import os
from importlib import import_module
from src.set_seed import setSeed
import yaml

RECIPE = "full"

with open(os.path.join("custom",RECIPE,"config.yaml"), "r") as f:
	config = yaml.load(f, Loader=yaml.FullLoader)
config["custom_name"] = RECIPE
setSeed(config["seed"])
recipe = getattr(import_module(f"custom.{RECIPE}.recipe"),"Recipe")(config)

model = recipe.getModel()
trainDataloader, validDataloader = recipe.getDataloader()
optimizer = recipe.getOptimizer()
criterion = recipe.getLoss()
scheduler = recipe.getScheduler()

from src.trainer import Trainer

t = Trainer(config, trainDataloader, validDataloader, model, optimizer,criterion,scheduler)
```

argparse로 `RECIPE` 부분의 인자를 받아야하지만, 혼자 연구하는 특성 상 설정하지 않았습니다. 귀찮아서...   

정리하면 모든 동적 소스들을 `importlib`으로 불러옵니다. 자연스레 파일 구조는 `custom/{동적인 파일 이름}/recipe.py`, `custom/{동적인 파일 이름}/config.yaml` 두 개가 기본이 됩니다.  

그렇게 recipe를 통하여 대부분의 소스를 불러온 뒤, 트레이너에 인자로 사용하여 훈련합니다.  

## Recipe 구현

이제 동적인 코드들의 전부인 Recipe을 구현합니다.  
Recipe은 바로 위 `AbstractRecipe`를 상속하여 만들어집니다.

## etc..