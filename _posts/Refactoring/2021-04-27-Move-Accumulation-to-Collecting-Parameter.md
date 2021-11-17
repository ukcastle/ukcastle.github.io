---
layout: post
title: Move Accumulation to Collecting Parameter (415)
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

#### 개요

>지역 변수에 정보를 축적하는 매우 긴 메서드가 있다면
그것을 여러 메서드로 분해하고 각 메서드에 수집 파라미터를 넘겨 정보를 축적하도록 한다.

<br>

#### 동기

수집 파라미터(Collecting Parameter)란 **주어진 메서드로부터 정보를 수집하기 위해** 그 메서드에 **파라미터로 넘겨지는 개체**이다. 이 패턴은 [Composed Method](https://ukcastle.github.io/refactoring/2021/04/14/Compose-Method/)와 함께 사용되는 경우가 많다.  

다루기 어려울 정도로 비대한 메서드를 분해하여 Composed Method로 만드려면 Composed Method에 의해 호출되는 각 메서드로부터 정보를 얻어 **어떻게 축적할것인가** 결정해야 한다.  
각 메서드가 각각의 정보를 축적해두었다가 나중에 최종 형태로 합칠수도 있지만, 각 메서드에 **수집 파라미터**를 넘겨 점진적으로 축적할 수도 있다. 각 메서드는 자신의 정보를 수집 파라미터에 쓰고 그 결과로 정보가 축적된다.  
수집 파라미터는 여러 객체의 **메서드**에 파라미터로 주어질 수 있다. 이때 수집 파라미터가 정보를 축적하는 방법은 두가지가 있다.  
1. 각 메서드가 수집 파라미터의 콜백 메서드를 호출해 정보를 전달하기
2. 객체 자신을 수집 파라미터에 넘기고 그 콜백 메서드를 호출하여 정보를 얻기  

수집 파라미터는 특정 클래스의 특정 인터페이스를 통해 정보를 얻고 축적하도록 구현된다. 따라서 많은 대상으로부터, 그리고 다양한 인터페이스를 통해 정보를 수집하려는 경우 별로 적합하지 않다. 그럴 땐 Visitor 패턴을 사용하자.  

이 패턴은 [Composite](https://ukcastle.github.io/designpattern/2021/04/19/Composite/) 패턴과 궁합이 잘 맞는다. 수집 파라미터가 컴포짓 구조로부터 재귀적으로 정보를 축적할 수 있기 때문이다. 좋은 예시로 JUnit 프레임워크가 있다.  

<br>

#### 장점
- 다루기 어려울 정도로 비대한 메서드를 작고 간단하며 이해하기 쉬운 여러 개의 메서드로 분해하는 데 도움이 된다.  
- 코드의 실행 속도가 향상될 수 있다.

<br>

#### 절차

1. 정보를축적하여 하나의 결과로 만드는 **축적 메서드를 찾는다.** 그 결과를 담는 지역 변수를 수집 파라미터로 만들 것이다. 결과 변수의 타입이 여러 메서드를 통해 반복적으로 정보를 모으는 데 적합하지 않다면, 타입을 바꾼다. 예를 들어 Java의 String은 정보를 축적하기 부적절하므로 StringBuffer 클래스로 바꾼다.  

2. 축적 메서드의 내부에서 정보 축적의 한 과정을 골라 [Extract Method](https://ukcastle.github.io/refactoring/2021/04/07/RefactoringToPattern/#extract-method) 리팩터링을 적용해 별도의 메서드로 뽑아낸다. 접근 지정자는 `private`로, 리턴 타입은 `void`로 하고 **결과 변수를 파라미터**로 받도록 한다. 뽑아낸 메서드에서는 **결과 변수에 정보를 기록**하도록 만든다.

3. 정보 축적의 나머지 과정에 대해서도 단계 2를 반복하여 원래 코드가 결과 변수를 파라미터로 받아 거기에 정보를 기록하는 메서드로 바꾸도록 한다.
결과적으로 축적 메서드는 다음 세 단계로 이루어진다.  
    - 결과 객체를 생성한다.
    - 여러 메서드 중 첫 번쨰 메서드에 결과 객체를 파라미터로 넘긴다.
    - 결과 객체로부터 수집된 정보를 얻는다.

단계 2,3번은 여러 메서드에 Composed Method 리팩터링을 적용한 게 된다.

<br>

#### 구현

[저번에 쓴 코드](https://ukcastle.github.io/refactoring/2021/04/19/Replace-Implicit-Tree-with-Composite/#%EC%A0%88%EC%B0%A8-4)에서 가져오겠다.

```java

public class TagNode{
    ...

     public String toString(){
        String result;
        result = "<" + this.name + this.attribute + ">";
        Iterator it = this.child().iterator();
        while(it.hasNext()){
            TagNode node = (TagNode)it.next();
            result += node.toString();
        }
        result += this.value;
        result += "</" + this.name + ">";
        
        return result;
    }

    ...
} 

```

##### 절차 1

해당 메서드는, Composite 구조 내 각 태그의 정보를 **재귀**적으로 축적해 변수 `result`에 저장하고 있다.

result의 타입을 String에서 StringBuffer로 바꿔준다.  

```java

public class TagNode{
    ...

     public String toString(){
        StringBuffer result = new StringBuffer("");
        result = "<" + this.name + this.attribute + ">";
        Iterator it = this.child().iterator();
        while(it.hasNext()){
            TagNode node = (TagNode)it.next();
            result += node.toString();
        }
        result += this.value;
        result += "</" + this.name + ">";
        
        return result;
    }

    ...
} 

```

##### 절차 2

정보 축적의 첫 단계로서 `result` 변수에 XML의 시작 태그와 속성을 결합해 변수에 저장하는 코드를 찾아 Extract Method를 적용한다.  
즉 `result = "<" + this.name + this.attribute + ">";` 해당 코드를 메서드로 뽑아낸다.  

```java

public class TagNode{
    ...

     public String toString(){
        StringBuffer result = new StringBuffer("");
        
        this.writeOpenTagTo(result);

        Iterator it = this.child().iterator();
        while(it.hasNext()){
            TagNode node = (TagNode)it.next();
            result += node.toString();
        }
        result += this.value;
        result += "</" + this.name + ">";
        
        return result;
    }

    private void writeOpenTagTo(StringBuffer result){
        result.append("<");
        result.append(this.name);
        result.append(this.attributes.toString());
        result.append(">");
    }
    ...
} 

```

2단계가 끝났다. `writeOpenTagTo(result)` 함수가 toString()에 추가됐다.  

##### 절차 3

다음 목표를 찾아보자. `while` 루프로 트리 내 모든 노드들을 재귀호출하고있다.  
여기서 문제가 생겼다. `toString()`메서드는 이미 정의되어 result를 수집 파라미터로 건내줄 수 없다. 따라서 다른 방법이 필요하다. StringBuffer를 파라미터로 받아 수집 파라미터로 사용하며 toString 메서드가 하던 일을 그대로 수행하는 **도우미 메서드** 를 만들면 해결할 수 있다.  

```java

public class TagNode{
    ...

     public String toString(){
        StringBuffer result = new StringBuffer("");
        
        this.appendContentsTo(result);

        Iterator it = this.child().iterator();
        while(it.hasNext()){
            TagNode node = (TagNode)it.next();
            result += node.toString();
        }
        result += this.value;
        result += "</" + this.name + ">";
        
        return result;
    }

    private void appendContentsTo(StringBuffer result){ // 도우미 메서드
        this.writeOpenTagTo(result);
        ...
    }

    private void writeOpenTagTo(StringBuffer result){
        result.append("<");
        result.append(this.name);
        result.append(this.attributes.toString());
        result.append(">");
    }
} 

```

이렇게 만든 뒤, 재귀 호출 부분을 `appendContentsTo(result)`에 넣는다.  

```java

public class TagNode{
    ...

     public String toString(){
        StringBuffer result = new StringBuffer("");
        this.appendContentsTo(result);
        return result.toString();
    }

    private void appendContentsTo(StringBuffer result){
        this.writeOpenTagTo(result);

        this.writeChildrenTo(result);
  
    }

    private void writeOpenTagTo(StringBuffer result){
        result.append("<");
        result.append(this.name);
        result.append(this.attributes.toString());
        result.append(">");
    }

    private void writeChildrenTo(StringBuffer result){
        Iterator it = this.child().iterator();

        while(it.hasNext()){
            TagNode node = (TagNode)it.next();

            node.appendContentsTo(result);
        
        }
        result += this.value;
        result += "</" + this.name + ">";
    }
} 

```

이렇게 재귀 호출 부분을 도우미메서드를 통하여 잘 구현되도록 만들었다. 이제 남은 부분도 2번 방식을 적용하며 마무리해보자.  


```java

public class TagNode{
    ...

     public String toString(){
        StringBuffer result = new StringBuffer("");
        this.appendContentsTo(result);
        return result.toString();
    }

    private void appendContentsTo(StringBuffer result){
        this.writeOpenTagTo(result);
        this.writeChildrenTo(result);
        this.writeValueTo(result);
        this.writeEndTagTo(result);
    }

    private void writeOpenTagTo(StringBuffer result){
        result.append("<");
        result.append(this.name);
        result.append(this.attributes.toString());
        result.append(">");
    }

    private void writeChildrenTo(StringBuffer result){
        Iterator it = this.child().iterator();
        while(it.hasNext()){
            TagNode node = (TagNode)it.next();
            node.appendContentsTo(result);
        }
    }

    private void writeValueTo(StringBuffer result){
        if(!value.equals("")){
            result.append(value);
        }
    }

    private void writeEndTagTo(StringBuffer result){
        result.append("</");
        result.append(this.name);
        result.append(">");
    }
} 

```

이렇게 `toString()` 이 매우 단순해졌다. 그리고 `appendContentsTo()`는 Composed Method의 훌륭한 본보기가 된다. 
