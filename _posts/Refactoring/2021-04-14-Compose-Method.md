---
layout: post
title: Compose Method (179)
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

#### 개요

>어떤 메서드의 내부 로직이 한 눈에 이해하기 어렵다면,  
그 로직을 의도가 잘 드러나며 동등한 수준의 작업을 하는 여러 단계로 나눈다.  

```c++
void add(Object* element){
    if(!element->readOnly){
        int newSize = element->size + 1;
        if (newSize > element->length){
            Object[] newElements = new Object[elements.length + 10];
        }
        for (int i=0;i < element->size;i++){
            newElements[i] = elements[i];
        }
        element = newElements;
    }
    elements[size++] = element;
}
```
물론 못 읽을정돈 아니지만, 해석하기 귀찮다.  

```c++
void add(Object* element){
    if (element->readOnly){
        return;
    }
    if (element->atCapacity()){
        element->grow();
    }
    addElement(element);
}
```

주석 없이도 한 눈에 읽어진다.

#### 동기

그렇게 대단한 방법은 아니다. 오히려 엄청 사소한 방법이다.  
그럼에도 이 리팩토링 방식은 최고의 리팩토링 방식 중 하나라고 불려진다.  
**작고 단순하여 몇초만에 이해할 수 있는** 메소드 이다.  
여러 메소드를 동등한 수준에서 나눈다.  

이 리팩터링은 여러번의 [Extract Method](https://jo631.github.io/refactoring/2021/04/09/RefactoringToPattern/#extract-method)를 포함한다. 이 때 가장 어려운 것은 어떤 부분을 메소드 내부에 남기고, 어떤 부분을 다른 메소드로 뽑아낼 것 인지 결정하는 것이다.  
너무 많은 코드를 뽑아내려 할 때 적절하게 작명하기 어려울것이다. 이런 경우 [Inline Method](https://jo631.github.io/refactoring/2021/04/09/RefactoringToPattern/#inline-method)를 적용하고 다시 분리할 방법을 찾아야 한다.  

만약 메소드가 너무 많아져 성능이 떨어진다고 생각되는가?  
여태 경험상 큰 문제는 없었고, 이는 BigO 방식으로 생각해봐도 큰 문제가 없다.  


#### 장점

- 어떤 메소드가 무슨 일을 하고, 또 그 일을 어떻게 수행하는지 효과적으로 표현한다.  
- 동등한 수준의 작업을 하며 이름이 적절하게 붙은 몇 단계로 내부를 나눔으로써, 메소드를 단순하게 만든다.
- 어떤 부분에서 문제가 생겼는지 쉽게 알 수 있어 디버깅이 쉬워진다.  

#### 단점

- 작은 메소드가 지나치게 많이 생길 수 있다.  
    > 이런 경우 [Extract Class](https://jo631.github.io/refactoring/2021/04/09/RefactoringToPattern/#extract-class)를 적용할 수 있다.  
- 로직이 여러 곳에 흩어지기 때문에 디버깅이 어려울 수 있다.  

#### 절차  
이 방식은 매우 간단한듯 보이지만, 막상 해보면 느끼는것은 **정답이 없다**.  
다만 다음과 같은 지침은 지시할 수 있다.  

- 작게 만든다
    > Composed Method의 코드는 10줄을 잘 넘기지 않는다. 보통 5줄 정도이다. 
- 사용되지 않거나 중복된 코드를 제거한다
- 코드의 의도가 잘 드러나도록 한다
    > 변수와 메서드, 파라미터 이름은 그 목적을 잘 표현하도록 짓는다
- 단순화 한다
    > 가능한 한 단순하게 바꾼다. 보기좋게.
- 동등한 수준으로 단계를 나눈다
    > 예를들어 세부 조건을 검사하는 로직과 몇개의 고수준 메소드를 호출하는 코드가 섞어있다면, 실패한것이다. 세부 조건 로직을 이름이 잘 지어진 별도의 메소드로 뽑아내 다른 고수준 메소드와 동등한 수준으로 맞춰야 한다.  