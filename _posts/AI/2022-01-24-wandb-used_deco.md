---
layout: post
title: 데코레이터 패턴을 활용한 WandB 연결하기   
category: [AI]
tag: [Pytorch] 
---  

## 개요

협업하며 연구 결과를 기록하거나, 여러가지 시각화를 위하여 WandB를 사용할 때가 많습니다.  

하지만 새로 만든 코드를 테스트하는 과정에서 WandB가 연결되어 있다면, 심지어 에러가 계속 난다면, 돌릴 때 마다 남는 로그가 고통스러울 것입니다.  

그런다고 주석처리하기엔 로그를 남기는 특성상 한 부분에 몰려있지 않죠.  

그런 이유로 Config에 환경변수로 저장해놓고, 변수에 따라 실행하는 방식을 사용합니다.  

하지만 모든 로그를 남기는 함수에 환경변수 체크 코드를 중복하는 것은 멋지지 않습니다.  

들어가기 앞서 전체 코드는 [여기](https://github.com/boostcampaitech2/final-project-level3-cv-04/blob/main/model_lab/frame_classification/src/wandb_helper.py) 있습니다.  

## 구현

일단, Config를 입맛대로 적어줍니다.  

```py
wandb : True
wandb_project : "{Project}"
wandb_entity : "{Entity}"
wandb_group : "{Group}"
```

다음 WandBHelper 클래스를 작성해줍니다.  
```py
import wandb
from src.metric import getScore
class WandB:
	def __init__(self, config):

		self.isRun = config["wandb"]
		self.config = config
		self.nowF1 = 0
		self.init()

	# Decorator
	def decorator_checkRun(originalFn):
		def wrapper(*args):
			if args[0].isRun:
				return originalFn(*args)
			else:
				return 
		return wrapper	

	@decorator_checkRun
	def init(self):
		wandb.login()
		wandb.init(
			project = self.config["wandb_project"],
			entity = self.config["wandb_entity"],
			name = self.config["output_dir"],
			group = self.config["wandb_group"],
			config = self.config,
		)

	@decorator_checkRun
	def trainLog(self,loss,acc,lr):
		wandb.log({
			"train/loss" : loss,
			"train/acc" : acc,
			"info/lr" : lr
		})
	
	@decorator_checkRun
	def validLog(self, preds, y_true, validLoss, validAcc, confusionMatrix, time):
		precision, recall, f1, f1List = getScore(confusionMatrix)
		self.nowF1 = f1
		conf_matrix = wandb.plot.confusion_matrix(probs=None,
			preds=preds, y_true=y_true,
			class_names= self.config["class_names"]
		)
		
		wandb.log({
			"valid/loss" : validLoss,
			"valid/acc" : validAcc,
			"valid/f1" : f1,
			"valid/precision" : precision,
			"valid/recall" : recall,
			"valid/sub/move1_f1" : f1List[0],
			"valid/sub/move2_f1" : f1List[1],
			"valid/sub/move3_f1" : f1List[2],
			"valid/sub/move4_f1" : f1List[3],
			"valid/sub/move5_f1" : f1List[4],
			"valid/sub/move6_f1" : f1List[5],
			"info/valid_time" : time,
			"confusion_matrix" : conf_matrix,
		})

	def getF1(self):
		return self.nowF1

```

첫 번째로 생성자 부분을 봅시다.  

```py
	def __init__(self, config):

		self.isRun = config["wandb"]
		self.config = config
		self.nowF1 = 0
		self.init()
```

`config` 내에 wandb의 Boolean 값이 있지만, 좀 더 보기 편하게 따로 변수를 지정했습니다.  

다음, 데코레이터 부분입니다.  


```py
# Decorator
	def decorator_checkRun(originalFn):
		def wrapper(*args):
			if args[0].isRun:
				return originalFn(*args)
			else:
				return 
		return wrapper
```

함수의 데코레이터와 비슷하지만, 클래스 내부의 데코레이터는 중요한 점이 하나 있습니다.  
바로 `self` 키워드를 잘 컨트롤해줘야 합니다.  
`*args` 중, `args[0]`가 self를 나타냅니다.  
이 외에는 예제로 많이 풀려있는 Decorator 패턴을 참조하면 됩니다.  

그 다음, 어떤 함수를 지정하던
```py
@decorator_checkRun
	def init(self):
```

이렇게 데코레이터를 선언해주기만 한다면, 완성입니다.  

