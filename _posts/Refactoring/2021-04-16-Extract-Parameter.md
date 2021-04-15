---
layout: post
title: Extract Parameter (456)
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

#### 개요

>메소드의 생성자 내에서 생성한 값을 필드에 저장하고 있다면
대입문의 우변을 새 파라미터로 대체해 클라이언트가 그 값을 지정할 수 있도록 한다.

```c++
class DecodingNode{
private:
    Node delegate;

public:
    DecodingNode(StringBuffer textBuffer, int textBegin, int textEnd){
        this.delegate = new StringNode(textBuffer, textBegin, textEnd);
    }
}


int main(){
    DecodingNode a = new DecodingNode(textBuffer, text, textEnd);
}
```

`delegate` 변수는 생성자 내에서 생성된다. 이 것을 리팩터링 해보자.  

```c++
class DecodingNode{
private:
    Node delegate;

public:
    DecodingNode(Node delegate){
        this.delegate = delegate;
    }
}


int main(){
    DecodingNode a = new DecodingNode(new StringNode(textBuffer, textBegin, textEnd));
}
```

#### 동기

때로는 객체 내의 필드에 대입할 값을 다른 객체가 지정하도록 하고 싶을 수 있다.  
만약 그 필드에 지역값이 할당되고 있다면, 대입문의 우변을 새로운 파라미터로 대체해 클라이언트가 그 값을 지정하도록 할 수 있다.  

#### 잘치

1. 이 리팩터링을 하기 전에 해당 필드에 대한 대입문은 생성자나 메서드 내에 있어야 한다. 그렇지 않다면, 해당 대입문을 생성자나 메서드 안으로 옮긴다.  
2. 생성자의 파라미터 값을 적절히 바꿔준다. 다음 대입해준다.  