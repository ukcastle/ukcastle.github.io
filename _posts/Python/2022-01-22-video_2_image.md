---
layout: post
title: OpenCV을 이용한 영상에서 이미지 추출하여 데이터셋 만들기  
category: [Python]
tag: [OpenCV] 
---  

## 개요  

데이터셋을 제작하며 영상에서 이미지로 데이터를 만들 때가 있습니다.  
이 때 주의해야될 점은 다음과 같습니다.  
- Train과 Valid에 같은 영상에서 나온 데이터가 들어가면 안된다.  

만약 같은 데이터가 Valid에도 들어갔다면, 해당 Valid Set은 오염된 상태이므로 실 생활에서 좋은 결과를 내지 못 할 확률이 높습니다.  

그리고 본 포스트에서는 약 30기가의 영상을 이미지로 추출하였기때문에, 빠른 속도를 위해 멀티프로세싱을 이용하였습니다.  

전체 코드는 [여기](https://github.com/boostcampaitech2/final-project-level3-cv-04/blob/main/input/make_input.py)있습니다.    

## 영상에서 이미지 추출하기  

가장 기본적이고 핵심이되는 코드는 다음과 같습니다.  

일단 비디오 파일을 `VideoCapture` 객체로 받아옵니다.  
```py
cap = cv2.VideoCapture(f"{VideoPath}")
```

```py
def readFrame(cap, frame):
	cap.set(cv2.CAP_PROP_POS_FRAMES, frame)
	ok, image = cap.read()

	if not ok:
		raise "프레임 없음"
	return image
```

openCV 라이브러리의 api들은 flag 변수를 첫 번째 변수로 포함하여 반환해줍니다.  

제가 작업한 데이터셋은 폴더 안에 비디오들이 있고, 이름이 매칭되는 csv파일이 있는 형태였습니다.  

코드에 주석을 달아뒀습니다.  

```py
def makeInput(datasetNum):
	rootDir = DATASET_NAME+str(datasetNum)
	videoRoot = os.path.join(rootDir,DATASET_VIDEO)
	videoFiles = os.listdir(videoRoot)
	labelName = "train.csv"
	cnt = 0
	
	# video 파일들의 리스트를 videoFiles에 저장 후 반복문

	for idx, video in enumerate(tqdm(videoFiles,position=datasetNum-1,leave=True,desc=f"datasetNum : {datasetNum}")):
		labelList = []

		# train과 valid의 차이를 8:2로 분류
		if labelName=="train.csv" and idx > len(videoFiles) * 0.8:
			labelName = "valid.csv"

		csv = ".".join(video.split(".")[:-1])+".csv"

		# 여러개의 Annotation을 동일한 Index로 존재하는 2차원 리스트로 만들도록 저장
		annotations = [pd.read_csv(os.path.join(rootDir,DATASET_ANNOTATION,x,csv)).to_numpy().tolist() 
			for x in os.listdir(os.path.join(rootDir,DATASET_ANNOTATION))
			if os.path.isfile(os.path.join(rootDir,DATASET_ANNOTATION,x,csv))
			]

		cap = cv2.VideoCapture(os.path.join(videoRoot,video))

		# 한 비디오에 대한 Annotation들을 추출
		for j, line in enumerate(zip(*annotations)):
			
			# 한 비디오에 있는 여러개의 Annotations를 병렬적으로 추출
			vertical = list(zip(*line))

			if 0.0 in vertical[2] or 7.0 in vertical[2]: #0,7번 라벨 거름
				continue

			# 한 프레임에 대한 최대 4개의 Annotation들을 리스트로 추출
			ls = []
			for i in vertical[2]:
				ls.append(int(i)) 
			
			# Annotations이 만장일치가 아니라면 Continue
			if len(set(ls)) > 1:
				continue
			
			label = int(ls[0])

			image = readFrame(cap, j)
			if not isHand(image):
				continue
			image = cv2.resize(image,RESIZE)
			fileName = f"{str(label)}/{cnt:07d}{IMAGE_FORMAT}"


			# CSV화를 위한 단계
			fullFileName = os.path.join(SAVE_IMAGE_ROOT,fileName)
			if not os.path.isfile(fullFileName):
				cv2.imwrite(fullFileName,image)

			oneLabel = [fullFileName]
			oneLabel.append(label)
			oneLabel.append(video)
			labelList.append(oneLabel)
			cnt+=1
		
		# Multiprocessing을 위한 처리, 모드는 덮어쓰기고 Header는 파일이 존재하지 않을시만 생성
		pd.DataFrame(labelList,columns=["file_name","label","video_name"]).to_csv(labelName, mode="a",header=not os.path.isfile(labelName),index=False)
		cap.release()
```

## 멀티프로세싱

이후 `makeInput` 함수를 병렬적으로 돌려줍니다.  

```py
def start(datasetNum):
	makeInput(datasetNum)

if __name__=="__main__":
	initFile()
	with Pool(len(DATASET_NUM)) as p:
		p.map(start,DATASET_NUM)
```

멀티프로세싱에 인자가 필요하다면 `p.map(func, *args)` 형태로 넣어주면 됩니다.  