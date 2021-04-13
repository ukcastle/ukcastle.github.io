---
layout: post
title: Introduce Polymorphic Creation with Factory Method (134)
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

#### 개요 

한 상속구조 내에서 클래스들이 어떤 메소드를 각자 구현하는데, 객체 생성 단계만 제외하고 나머지가 서로 유사하다면 그 메소드를 부모클래스로 옮기고 객체 생성은 팩터리 메소드에 맡기기로 한다.  

```java
class DOMBuilder{
    public DOMBuilder(){
        System.out.println("Make DOMBuilder");
    }

    public void action(){
        System.out.println("Do Action");
    }
}

class XMLBuilder{
    public XMLBuilder(){
        System.out.println("Make XMLBuilder");
    }

    public void action(){
        System.out.println("Do Action");
    }
}

class Test {
    public static void main(String[ ] args) {
        DOMBuilder a = new DOMBuilder();
        XMLBuilder b = new XMLBuilder();

        a.action();
        b.action();
    }
}
/*  ---output---
    Make DOMBuilder
    Make XMLBuilder
    Do Action
    Do Action
*/
```
매우 간단하게 표현했는데, 요약하면 둘다 똑같은 action()을 사용한다.  
이러면 해당 리팩토링을 할 수 있다.

```java
interface IBuilderAction{
    public void action();
}

abstract class AbstractBuilder implements IBuilderAction{
    public abstract IBuilderAction createBuilder();
    
    @Override
    public void action(){
        System.out.println("Do Action");
    }
}


class DOMBuilder extends AbstractBuilder{
    protected DOMBuilder(){}

    @Override
    public IBuilderAction createBuilder(){
        System.out.println("Make DOMBuilder");
        return new DOMBuilder();
    }
}

class XMLBuilder extends AbstractBuilder{
    protected XMLBuilder(){}

    @Override
    public IBuilderAction createBuilder(){
        System.out.println("Make XMLBuilder");
        return new XMLBuilder();
    }
}

class Test {
    public static void main(String[ ] args) {
        DOMBuilder a = new DOMBuilder();
        XMLBuilder b = new XMLBuilder();

        IBuilderAction[] c = {a.createBuilder(), b.createBuilder()};

        for(IBuilderAction i : c){
            i.action();
        }
    }
}
```

#### 동기

[Replace Constructors with Creation Methods](https://jo631.github.io/refactoring/2021/04/13/Replace-Constructors-With-Creation-Methods/) 리팩터링을 구현하려면 원하는 객체를 생성해서 리턴하는 메소드를 클래스에 추가하기만 하면 된다.  그 메소드는 `static`메법일 수도 있고 아닐수도 있다. 다만 Factory Method 패턴의 경우 다음과 같은 필수 요소가 필요하다.  
>1. 팩토리 메서드가 생성해 리턴하는 객체의 집합을 대표하는 하나의 타입(Interface, Abstract Class, Class)
>2. 위 타입을 구현하는 클래스들
>3. 팩토리 메소드를 구현하는 여러개의 클래스

팩토리 메소드는 단순히 공통 인터페이스를 가지는 클래스들을 통해 구현되기도 하지만, 실제로 어떤 상속 구조 내에서 구현되는 것이 보통이다.(사실 예시는 전자이다.)  

팩토리 메소드 패턴은 [Template Method](https://jo631.github.io/designpattern/2021/04/14/TemplateMethod/) 패턴과 함께 사용되는 경우가 많다.  즉 템플릿 메소드에서 팩토리 메소드를 호출하는 것이다.  

상속 구조 내의 중복 코드를 제거하는 리팩터링을 수행하다 보면 자연스럽게 이 두 패턴을 함께 사용하게 된다. 예를들어 어떤 메소드가 수퍼 클래스에도 있고 여러 서브클래스에서도 오버라이드 되어있는데, 그 구현은 객체 생성 단계를 제외하곤 거의 동일하다면? 이 상황에서는 그 메서드를 수퍼 클래스로 옮겨 일종의 템플릿 메서드를 만들면 중복 코드를 제거할 수 있다.  
그러나 수퍼 클래스에서 어떤 경우에 어느 객체를 생성해야 할지 알 수 없으므로 그 작업은 서브 클래스에 맡겨야 한다. 그리고 이런 상황에서는 Factory Method보다 적절한 패턴이 없다.  
Factory Method를 사용하는 것이 new 연산자나, 생성자를 사용하는 것보다 더 간단할까? 답은 아니다. 하지만 구현하고 나면, 이전보다 중복이 훨씬 줄어들고 안정적이다.  

#### 장점

- 객체를 생성하는 과정에서의 코드 **중복**을 줄인다.
- 객체를 생성하는 곳이 실제로는 어디고, 또 어떻게 오버라이드하면 되는지 잘 드러난다.  
- 팩토리 메소드에서 인스턴스로 만들 클래스가 특정 타입을 구현하도록 **강제**할 수 있다.  

#### 단점

- 일부 서브클래스가 구현하는 팩토리 메소드에서는 **불필요한 파라미터**를 어쩔 수 없이 남겨둬야 할 수도 있다.


#### 절차

이 리팩터링이 주로 사용되는 상황은 다음과 같다.  
> - 형제 서브클래스들이 어떤 메소드를 각각 구현하고 있는데, 객체를 생성하는 단계만 제외하고는 거의 유사한 경우
> - 수퍼클래스와 서브클래스가 어떤 메소드를 각각 구현하고 있는데, 객체를 생성하는 단계만 제외하고는 거의 유사한 경우  

둘 중 위의 상황에 대해 다루겠다. 하지만 방법은 거의 비슷하다.  

1. 유사 메소드 중 하나를 선택해 객체 생성 단계가 별도의 객체 생성 메소드에서 수행되도록 수정한다(생성자를 포장해 함수로 만든다). 객체를 생성하는 코드에 [Extrac Method](https://jo631.github.io/refactoring/2021/04/09/RefactoringToPattern/#extract-method)를 적용해도 되고, 이미 객체 생성 메소드가 있으면 그 것을 사용해도 된다.  
객체 생성에서는 일반화된 이름을 붙여야 한다. 모든 형제 클래스에도 같은 이름의 메소드가 필요하기 때문이다. 또한 리턴 타입도 유사 메소드들이 모두 포괄할 수 있는 타입이여야 한다.  

2. 나머지 유사 메소드에 대해서도 단계 1을 반복한다. 관련된 모든 형제 클래스에 객체 생성 메소드가 하나씩 생길텐데 각 메소드의 시그니처는 모두 동일해야한다.  

3. 다음은 수퍼클래스를 수정해야 한다. 만약 직접 수정할 수 없는 상황이거나, 수정하지 않는것이 더 좋다고 판단될 경우는 [Extract Superclass](https://jo631.github.io/refactoring/2021/04/09/RefactoringToPattern/#extract-superclass)를 사용하고, 원래의 수퍼클래스를 상속하는 서브클래스들은 새로 만들어진 수퍼클래스를 상속하게 바꾼다.  

4. 유사 메소드에 대해 From Template Method 를 적용한다. 이 과정에서는 [Pull Up Method](https://jo631.github.io/refactoring/2021/04/09/RefactoringToPattern/#pull-up-method)가 포함되는데, 이 때는 꼭 절차에 나온 조언을 따르는 것을 추천한다.
> 타입 검사를 엄격하게 하는 프로그래밍 언어를 사용하고, 옮기려는 메소드에서 두 서브클래스에는 있지만 수퍼클래스에는 없는 메소드를 호출한다면, 해당 메소드를 수퍼클래스에 추상 메소드로 선언하라.

이렇게 하면, 팩토리 메소드를 구현하게 된 것이다. 이제 형제클래스들은 각각 FactoryMethod:ConcreteCreator가 되었다.  

5. 관련 형제클래스에 또 다른 유사 메소드가 존재하고 그 메소드도 수정하는것이 좋다 판단되면 1~4를 반복한다.  

6. ConcreteCreator 클래스들이 팩터리 메소드 가운데 다수가 동일한 객체 생성 코드를 포함하고 있다면, 수퍼클래스에서 선언된 팩토리 메소드로 그 코드를 옮겨 객체 생성에 대한 디폴트 구현으로 삼는다. 

#### 예제 

위의 첫번째 -> 마지막으로 가기 위한 단계이다.  

##### 객체 생성 메소드 추출

```java
interface IBuilderAction{
    public void action();
    public IBuilderAction createBuilder();
}


class DOMBuilder implements IBuilderAction{
    public DOMBuilder(){}

    @Override
    public IBuilderAction createBuilder(){
        System.out.println("Make DOMBuilder");
        return new DOMBuilder();
    }

    @Override
    public void action(){
        System.out.println("Do Action");
    }
}

class XMLBuilder implements IBuilderAction{
    public XMLBuilder(){}

    @Override
    public IBuilderAction createBuilder(){
        System.out.println("Make XMLBuilder");
        return new XMLBuilder();
    }

    @Override
    public void action(){
        System.out.println("Do Action");
    }


}

class Test {
    public static void main(String[ ] args) {
        DOMBuilder a = new DOMBuilder();
        XMLBuilder b = new XMLBuilder();

        IBuilderAction[] c = {a.createBuilder(), b.createBuilder()};

        for(IBuilderAction i : c){
            i.action();
        }
    }
}
```

인터페이스 집합을 만든 뒤 action을 실행시킨다. 잘 실행 된다.  

##### 수퍼클래스 생성

```java
interface IBuilderAction{
    public void action();
}

abstract class AbstractBuilder implements IBuilderAction{
    public abstract IBuilderAction createBuilder();
    
    @Override
    public void action(){
        System.out.println("Do Action");
    }
}


class DOMBuilder extends AbstractBuilder{
    protected DOMBuilder(){}

    @Override
    public IBuilderAction createBuilder(){
        System.out.println("Make DOMBuilder");
        return new DOMBuilder();
    }
}

class XMLBuilder extends AbstractBuilder{
    protected XMLBuilder(){}

    @Override
    public IBuilderAction createBuilder(){
        System.out.println("Make XMLBuilder");
        return new XMLBuilder();
    }
}

class Test {
    public static void main(String[ ] args) {
        DOMBuilder a = new DOMBuilder();
        XMLBuilder b = new XMLBuilder();

        IBuilderAction[] c = {a.createBuilder(), b.createBuilder()};

        for(IBuilderAction i : c){
            i.action();
        }
    }
}
```

클래스 안에 마땅한 실행 코드를 넣지 못하여 가장 낮은 수준의 클래스들의 생성자를 `protected`로 설정했지만, 적당한 예제를 찾으면, 이를 `private`으로 놓아도 된다.  
그 다음 AbstractBuilder 클래스를 만든 뒤, 두 개가 중복되는 메소드를 AbstractBuilder 클래스에 넣는다.  
