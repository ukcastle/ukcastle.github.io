---
layout: post
title: TemplateMethod
category: DesignPattern
tag: [DesignPattern] 
---

Template Method가 뭔지 몇번 다뤄봤지만, 제대로 정리하려고 한다.  

> Defines the skeleton of an algorithm in a method, deferring some steps to subclasses. Template Method lets subclasses redefine certain steps of an algorithm without changing the algorithms structure. – GoF Design Patterns

알고리즘의 구조를 메소드에 정의하고, 하위 클래스에서 알고리즘 구조의 변경 없이 알고리즘을 재정의하는 패턴. 알고리즘이 단계적으로 나누어지거나 같은 역할을 하는 메소드지만 여러 군데에서 필요할 때 주로 사용된다.  

#### 간단한 예제

```java
class Template{
    protected void algorithm1(){
        System.out.println("algorithm - 1");
    }
    protected void algorithm2(){
        System.out.println("algorithm - 2");
    }

    public void mainAlgorithm(){
        algorithm1();
        algorithm2();
    }
}

class Test {
    public static void main(String[ ] args) {
        Template a = new Template();

        a.mainAlgorithm();
    }
}
```

매우 간단하다. 이렇게 큰 알고리즘 안에 알고리즘들을 조각조각 시켜놓는 방법이다.  
지금은 한 클래스 안에 넣어놨지만, 여러 클래스에서 필요한 알고리즘이면 상속을 하던가 변수를 불러오던가 해서 사용하면 된다.  