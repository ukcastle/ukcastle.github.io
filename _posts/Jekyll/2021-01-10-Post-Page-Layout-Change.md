---
layout: post
title: 21-01-10-TiL
category: Jekyll
tag: [Jekyll]
---

# Post Page Layout Change
원래같았으면 라즈베리파이를 시작해보려고 했지만...  
일단 홈페이지 레이아웃을 좀 수정했다.
[졸작 팀원의 블로그](https://outstanding1301.github.io/)에서 많이 참고하였다.  
<br/>

#### 스크롤에 따른 목차 띄우기
오른쪽에 신기한게 생겼다.  

------
![move](https://github.com/jo631/jo631.github.io/blob/dev/postimg/move.gif?raw=true)

------
scroll spy라고 하던데 이름 적당한 것 같다.
솔직히 근데 정말 따라하기만 했다.  
[여기](https://outstanding1301.github.io/git/2021/01/08/table-of-contents-scroll-spy/)를 보고 참고했다. JS를 사용했고 좀 비효율적이라고 들었다.   
다음에 더 알아보자  
<br/>

#### 댓글 기능 추가

댓글기능이 고장나있길래 다른 방법으로 만들었다.  
[Utteranc](https://utteranc.es/?installation_id=13996114&setup_action=install)를 이용했다. 너무너무 편리해서 좋았다.  
자신의 계정과 저장소 이름을 입력한 뒤 스크립트 코드를 복붙해서 사용하면 된다

------
![comment](https://github.com/jo631/jo631.github.io/blob/dev/postimg/comment.gif?raw=true)  

------
내 템플릿같은 경우는 `layouts/post.html`에 넣으면 됐다.  
행복하다 이런게 오픈소스지  

#### 좋아요 기능 추가
------
![like](https://github.com/jo631/jo631.github.io/blob/dev/postimg/like.gif?raw=true)

------
직접 구현해보려했는데, 한 가지 간과한점이 있다.  
지킬은 **정적** 홈페이지 이다. 즉 **동적**인 행동을(좋아요를 누르면 추가되는 행동) DB와 연동하여 넣기가 힘들다더라. [출처: StackOverflow](https://stackoverflow.com/questions/39344219/like-button-for-posts-in-jekyll)  
대신 [LIKEBTN](https://likebtn.com/en/)이라는 사이트를 이용하면, 위의 댓글기능과 같은 맥락으로 설정이 가능했다.  


------
![likebtn](https://github.com/jo631/jo631.github.io/blob/dev/postimg/LIKEBTN.jpg?raw=true)

------  

이쁘게 커스터마이징을 하고 우측에 스크립트 코드 복사를 누르면 된다.  
Like 버튼을 누르면 광고가 표시되길래... 꼼수를 좀 부렸다. 

오늘 공부한 것은 여기까지이다. 사실 내부적으로 더 바꾸긴 했다.


#### 여담
좋아요 만들어봤자 나말고 누를 사람이 있을까?  
