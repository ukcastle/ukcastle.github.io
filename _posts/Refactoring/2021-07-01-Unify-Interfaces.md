---
layout: post
title: Unify Interfaces (453)
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

#### 개요

>수퍼클래스(또는 인터페이스)가 서브클래스와 동일한 인터페이스를 가질 필요가 있다면,  
서브 클래스에서 수퍼 크래스에 없는 모든 public 메서드를 찾아 이를 수퍼클래스에 추가한다. 이때 메서드 몸체는 비워놓아 아무 일도 하지 않도록 만든다.

<br>

#### 동기

여러 객체를 다형적으로 처리하려면, 그 클래스들은 수퍼 클래스가든 인터페이스든 동일한 인터페이스를 공유해야 한다. 이 리팩터링은 수퍼클래스나 인터페이스가 자신의 서브클래스와 동일한 인터페이스를 가져야 할 필요가 있는 경우를 위한 것이다.  

수퍼클래스와 서브클래스에 이 리팩터링을 적용한 후 독자적인 인터페이스를 만들기 위해 수퍼클래스에 Extract Interface를 적용하는 경우도 있다. 그럴때는 보통 추상 클래스에 상태 필드가 있지만 이를 상속해 구현하는 (데코레이터 같은)서브클래스가 이 필드를 상속하는 것을 원하지 않을 때다.

<br>

#### 절차

서브 클래스의 public 메서드 중 수퍼클래스나 인터페이스에 선언되지 않은 것을 찾는다. 이런 메서드를 간단히 **누락 메서드** 라 하자.  

1. 누락 메서드를 복사해 수퍼클래스/인터페이스에 추가한다. 수퍼클래스에 추가하는 경우라면 메서드 몸체를 비워 아무 일도 하지 않게 만든다.
2. 수퍼클래스/인터페이스가 서브클래스와 동일한 인터페이스를 갖게 될 때 까지 반복한다. 

#### 구현

```java
public class StringNode extends AbstractNode{
    public void accept(TextExtractor: textExtractor){
        ...
    }
}
```

##### 과정 1
accept(...)를 복사해 수퍼클래스에 추가한 다음, 메서드가 아무 일도 하지 않도록 몸체를 비운다.  

```java
public abstract class AbstractNode{
    public void accpet(TextExtractor: textExtractor){}
}
```