---
layout: post
title: Form Template Method (281)
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

#### 개요  

>한 상속 구조 내의 어떤 두 서브 클래스가 유사한 단위 작업을 같은 순서로 실행하는 메서드를 각자 구현하고 있다면,  
각 단위 작업을 별도의 메서드를 뽑아내어 두 메서드를 일반화하고 이렇게 일반화된 메서드를 수퍼클래스로 올려 [Template Method](https://jo631.github.io/designpattern/2021/04/14/TemplateMethod/)로 만든다.  

```java
class CapitalStrategy{
    public abstract double riskAmountFor(string name, double duration);
}

class PlanAStrategy extends CapitalStrategy{

    @Override
    public double riskAmountFor(string name, double duration) {
        double value =  getPrice(getRisk(name),duration);
        return value;
    }
}

class PlanBStrategy extends CapitalStrategy{
    @Override
    public double riskAmountFor(string name, double duration) {
        double value =  getPrice(getRisk(name),duration);
        value = value / 2;
        return value;
    }
}
```

간단하게, 똑같이 `getPrice`를 실행한 뒤 PlanB는 value 변수를 2로 나눠서 리턴한다.  
이렇게 유사한 단위 작업을 반복하는 메소드에 대해서 이 리팩토링을 실행할 수 있다.  

```java
class CapitalStrategy{

    protected double riskAmountFor(string name, double duration){
        return getPrice(getRisk(name),duration);
    }
}

class PlanAStrategy extends CapitalStrategy{
    @Override
    public double riskAmountFor(string name, double duration){
        return super();
    }
}

class PlanBStrategy : public CapitalStrategy{
    @Override
    public double riskAmountFor(string name, double duration){
        return super() / 2 ;
    }
}
```

c++은 super 호출하기가 왜이리 힘든가..  

이렇게 해서 중복된 코드를 없애는 리팩토링이다.  

#### 동기

[Template Method](https://jo631.github.io/designpattern/2021/04/14/TemplateMethod/)는 알고리즘에서 불변적인 부분은 한 번만 구현하고 가변적인 동작은 서브클래스에서 구현할 수 있또록 남겨둔 것 이다.  
서브클래스에 불변적인 부분과 가변적인 부분이 뒤섞여있다면 여러 서브클래스에서 중복될 것이다.  

템플릿 메서드의 불변적 동작은 다음을 포함한다.  
- 알고리즘을 구성하는 메서드 목록과 그 호출 순서  
- 서브클래스가 꼭 오버라이드해야 할 추상 메서드
- 서브클래스가 오버라이드 해도 되는 훅 메서드(기본적인 내용만 구현되어 있거나 아무 코드도 들어가있지 않은 메서드)  

예를 들어 다음 코드를 살펴보자.  

```java
public abstract class Game{
    public void initialize(){
        deck = createDeck();
        shuffle(deck);
        drawGameBoard();
        dealCardFrom(deck);
    }
    protected abstract Deck createDeck();

    protected void shuffle(Deck deck){
        // shuffle ~~
    }
    protected abstract void drawGameBoard();
    protected abstract void dealCardFrom(Deck deck);
}

public class BoardGame extends Game{
    //~~~
}
```

여기서 `initialize()`는 메서드의 목록을 정의하고 호출 순서도 규정한다.
> 알고리즘을 구성하는 메서드 목록과 그 호출 순서

서브 클래스인 BoardGame은 shuffle을 제외한 3가지 함수를 무조건 정의해야 한다. 
> 서브 클래스가 꼭 오버라이드 해야 할 추상 메서드

하지만 shuffle은 불변적 부분이 아니라 그대로 사용을 하던, 변경을 하던 선택할 수 있다.  
> 브클래스가 오버라이드 해도 되는 훅 메서드  


템플릿 메서드는 종종 팩토리 메서드를 호출하기도 한다. [Introduce Polymorphic Creation with Factory Method](https://jo631.github.io/refactoring/2021/04/13/Introduce-Polymorphic-Creation-with-Factory-Method/)에 예제가 있다.  
자바같은 경우는 Template Method를 `final`로 선언하여 서브클래스에서 실수로 `override`하는것을 방지할 수 있다. 단 이런 방법은 클라이언트 코드에서 Template Method의 불변적인 부분을 전혀 변경할 필요가 없는 것이 **확실**할 때만 사용해야 한다.  

#### 장점

- 서브 클래스들의 공통 기능을 수퍼클래스로 옮겨, **중복 코드가 제거**된다
- 알고리즘의 과정이 단순해지고 쉽게 알아볼 수 있다.  
- 서브클래스에서 알고리즘의 구현을 재정의하는것이 쉬워진다.

#### 단점  

- 서브클래스가 꼭 구현해야 하는 메서드의 개수가 많다면 설계가 복잡해진다.  
    > 서브클래스에서 `override` 해야하는 추상 메서드의 개수를 최소화하는 것을 추천한다. 그렇지 않으면 Template Method의 내용을 자세히 살펴보지 않고는 프로그래머가 클라이언트에서 어떤 메서드를 `override`해야 할 지 쉽게 알 수 없을 것이다.    

#### 절차

1. 주어진 상속 구조 내에서 두 서브클래스 사이 유사 메서드가 존재하는지 확인한다. 유사 메서드가 확인되면 [Compose Method](https://jo631.github.io/refactoring/2021/04/14/Compose-Method/)를 적용한다. 이 과정에서 동일한 시그니처와 내용을 가지는 메서드(이하 공동 메서드)와 그렇지 않은 메서드(이하 특수 메서드)가 새로 생성될 수 있다.  
메서드를 `공동 메서드` 또는 `특수 메서드` 중 하나로 결정하기 전 고려해야 할 것이 있다. `특수 메서드`로 만든다면 나중에 결국 이것을 수퍼클래스의 **추상 메서드 또는 훅 메서드**로 만들어야 한다. **제 3의 다른 서브클래스**가 이 특수 메서드를 상속하거나 오버라이드 할 필요가 있는가? 그렇지 않다면 처음부터 공통 메서드로 만들어야 한다.  
<br>
2. 공통 메서드를 [Pull Up Method](https://jo631.github.io/refactoring/2021/04/09/RefactoringToPattern/#pull-up-method)를 통해 수퍼클래스로 올린다.  
<br>

3. 양쪽 서브클래스에서 유사 메서드의 내용이 서로 같아지도록 각 `특수 메서드`에 [Rename Method](https://jo631.github.io/refactoring/2021/04/09/RefactoringToPattern/#rename-method)를 적용한다.  
<br>

4. 혹시 두 유사 메서드의 시그니처가 동일하지 않다면 [Rename Method](https://jo631.github.io/refactoring/2021/04/09/RefactoringToPattern/#rename-method)를 적용한다.  
<br>

5. 이제 양쪽 유사 메서드를 [Pull Up Method](https://jo631.github.io/refactoring/2021/04/09/RefactoringToPattern/#pull-up-method)를 통해 수퍼클래스로 올린다. 그리고 각각 특수 메서드에 대응하는 추상 메서드를 수퍼클래스에 정의한다. 수퍼클래스로 올린 유사 메서드는 이제 Template Method가 되었다.  