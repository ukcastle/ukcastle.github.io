---
layout: post
title: Unify Interface with Adapter (333)
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

#### 개요

>클라이언트가 두 개의 유사한 클래스를 사용하고 있는데 그 중 한 인터페이스가 다른 하나보다 더 좋아보이면
[어댑터](https://jo631.github.io/designpattern/2021/05/01/Adapter/)를 도입해 인터페이스를 통합한다.

<br>

#### 동기

Adapter 패턴을 사용하는 경우는 다음과 같다.  

- 두 클래스가 동일하거나 유사한 작업을 수행하지만 **인터페이스가 서로 다른 경우**
- 두 클래스가 공통 인터페이스를 가지면, 클라이언트 **코드가 더 간단하고 명료**해질 수 있는 경우
- 외부 라이브러리라서 인터페이스를 바꾸고 싶어도 **쉽게 바꿀 수 없는 경우**, 또는 인터페이스가 프레임워크의 일부라서 **이미 많은 클라이언트에서 사용**되고 있는 경우, 또는 **소스 코드를 갖고 있지 않은 경우**

위와 같이, 비슷한 일을 하는 클래스지만 공통 인터페이스가 없어 각각을 별도의 방식으로 사용해야 하는 경우를 [인터페이스가 서로 다른 대체 클래스](https://jo631.github.io/refactoring/2021/04/12/RF-Ch4/#%EC%9D%B8%ED%84%B0%ED%8E%98%EC%9D%B4%EC%8A%A4%EA%B0%80-%EC%84%9C%EB%A1%9C-%EB%8B%A4%EB%A5%B8-%EB%8C%80%EC%B2%B4-%ED%81%B4%EB%9E%98%EC%8A%A4)의 냄새가 난다고 표현한다.  
이 냄새를 제거하는 가장 간단한 방법은 메서드의 이름을 바꾸거나 메서드 자체를 옮겨 인터페이스를 동일하게 만드는 것이다. 하지만 위의 설명과 같은 이유로 그렇게 할 수 없다면, Adapter 패턴의 도입을 고려해야 한다.  
Adapter 패턴으로 리팩토링하면 코드가 **일반화**되는 경우가 있다. 그리고 이 리팩터링은 **코드 중복을 제거하기 위한 다른 리팩터링의 토대**가 된다. Adapter패턴을 도입하여 대체 관계에 있는 클래스의 인터페이스를 하나로 통합하면, 클라이언트가 대체 클래스를 사용하는 방식 또한 일반화된다. 그 이후 [Form Templater Method](https://jo631.github.io/refactoring/2021/04/16/Form-Template-Method/)를 적용하면 클라이언트 코드의 **중복된 처리 로직을 제거**할 수 있다. 따라서 클라이언트 코드가 더 간결해진다.  

##### 장점

- 클라이언트가 **대체 클래스들을 하나의 인터페이스를 통해 사용하도록 통합**함으로써, 코드 중복을 없애거나 줄인다.
- 클라이언트 **코드가 간결**해진다.
- 클라이언트가 **대체 클래스들을 사용하는 방식이 통합**된다.

##### 단점

- 해당 클래스의 인터페이스를 **직접 바꾸는 것이 가능한 상황**에서 Adapter 패턴을 구현하면, 쓸데없이 **설계만 복잡**해진다.  

<br>

#### 절차

1. 대체 클래스 중 **가장 일반적이고 적합한 인터페이스를 가진 클래스**에 [Extract Method](https://jo631.github.io/refactoring/2021/04/07/RefactoringToPattern/#extract-method)를 적용시켜 공통 인터페이스를 만든다. 그리고 뽑아낸 메서드의 파라미터를 조사해, 대체 클래스 타입을 쓰는 것이 있으면 새로 정의한 공통 인터페이스 타입을 사용하도록 변경한다.  
이 후 단계에서는, 클라이언트가 어댑팅의 대상이 되는 Adapter 클래스를 사용할 때 이 단계에서 만든 공동 인터페이스를 통하도록 수정할 것이다.  

2. Adapter 클래스를 **사용하는 클라이언트 클래스**를 찾는다. 그리고 [Extract Class](https://jo631.github.io/refactoring/2021/04/07/RefactoringToPattern/#extract-class)를 적용해, 원시 어댑터를 만든다. **원시 어댑터**란, **어댑티 객체를 저장하는 필드를 선언하고 그에 대한 getter/setter를 제공하는 클래스**를 말한다.

3. 클라이언트 코드 중 어댑티 클래스 타입의 **필드 또는 지역 변수, 메서드 파라미터가 있다면**, 모두 **원시 어댑터 타입으로 치환**한다.

4. 클라이언트 코드 중 어댑티의 메서드를 호출하는(어댑터의 **get 메서드를 경유**하는) 부분을 모두 **별도의 메서드**로 뽑아낸다. 즉, 어댑팅되어야 할 메서드에 Extract Method 리팩터링을 적용하는 것이다. 이 때 뽑아낸 메서드 안에서 사용할, 어댑팅되는 객체에 대한 참조는 파라미터를 통해 받도록 한다.  
말이 좀 어려운데 예시를 들면 다음과 같다.  
    ```java
    Adapter childNode = new Adapter(...);
    currentNode.getElement().appendChild(childNode.getElement());
    ```
    이 것을
    ```java
    private void appendChild(Adapter parent, Adapter child){
        parent.getElement().appendChild(child.getElement());
    }

    appendChild(currentNode, childNode);
    ```
    이렇게 바꾼다.

5. 단계 4에서 뽑아낸 메서드 중 하나에 [Move Method](https://jo631.github.io/refactoring/2021/04/07/RefactoringToPattern/#move-method)를 적용해 단계 2에서 만든 원시 어댑터 클래스로 옮긴다. 즉, **클라이언트가 어댑티 클래스의 메서드를 호출할 때 항상 어댑터를 통하도록**만드는 것이다.  
이 때 주의할 점이 있는데, 단계 1에서 만든 공통 인터페이스를 살펴보면 지금 옮기려는 메서드에 대응하는 메서드가 하나씩 있을텐데, 메서드를 옮긴 후 시그니처가 대응 메서드의 시그니처와 **최대한 비슷**해야 한다. 그리고 만약 옮겨진 메서드의 내부 코드에서 클라이언트로부터 얻어야 하는 **추가적인 정보**가 있다고 해도, **파라미터를 추가하는 것은 피해야 한다.** 메서드의 시그니처가 이에 대응하는 공통 인터페이스의 메서드와 달라질 것이기 때문이다. 가능하면 메서드의 시그니처를 바꾸지 않고 해결할 수 있는 방법을 찾아야 한다. 어댑터 클래스 생성자의 파라미터로 전달하거나, 어댑터에 객체를 넘겨 런타임에 그 객체를 통해 값을 얻을 수 있도록 할 수 있다. 그러나 꼭 메서드의 파라미터로 넘길 수 밖에 없다면, 공통 인터페이스에 있는 대응 메서드의 시그니처를 적절히 수정해 두 메서드를 일치시켜야 한다.

6. 어댑터 클래스가 공통 인터페이스를 구현하도록 수정한다. 이 떄, 메서드의 파라미터 중 어댑터 타입인 것이 있다면 모두 공통 인터페이스 타입으로 변경한다.

7. 클라이언트 코드에서 어댑터 타입을 사용하는 모든 부분을 공통 인터페이스타입으로 변경한다.  

<br>

#### 구현

[원본](https://jo631.github.io/refactoring/2021/04/13/Introduce-Polymorphic-Creation-with-Factory-Method/#%EA%B0%9C%EC%9A%94)의 예제이다.  

```java
interface IBuilderAction{
    ...
}

abstract class AbstractBuilder implements IBuilderAction{
    ...
}


class DOMBuilder extends AbstractBuilder{

    private Document document;
    private Element root;
    private Element parent;
    private Element current;

    public void addAtrribute(String name, String value){
        this.current.setAttribute(name, value);
    }

    public void addBelow(String child){
        Element chileNode = document.createElement(child);
        this.current.appendChild(child);
        this.parent = this.current;
        this.current = childNode;
        history.push(this.current);
    }

    public void addBeside(String sibling){
        if (this.current == this.root){
            // Exception
        }
        Element siblingNode = this.document.createElement(sibling);
        this.parent.appendChild(siblingNode);
        this.current = siblingNode;
        history.pop();
        history.push(current);
    }

    public void addValue(String value){
        this.current.appendChild(document.createTextNode(value));
    }
    ...
}

class XMLBuilder extends AbstractBuilder{
    private TagNode rootNode;
    private TagNode currentNode;

    public void addChild(String childTagNode){
        this.addTo(this.currentNode, childTagNode);
    }

    public void addSibling(String siblingTagNode){
        this.addTo(this.currentNode.getParent(), siblingTagNode);
    }

    private void addTo(TagNode parentNode, String tagName){
        this.currentNode = new TagNode(tagName);
        parentNode.add(this.currentNode);
    }

    public void addAttribute(String name, String value){
        this.currentNode.addAttribute(name, value);
    }

    public void addValue(String value){
        this.currentNode.addValue(value);
    }
    ...
}
```

위 코드를 보면 두 빌더가 거의 동일한 메서드를 구현하고 있다. 차이점이 있다면 한 쪽은 `TagNode`를, 한 쪽은 `Element`를 사용한다는 것 뿐이다. 따라서 두 클래스에 대한 **공통 인터페이스**를 만들어 **두 빌더 클래스 사이에 중복된 코드를 제거하는 것**이 이 리팩터링의 목표이다.  

##### 절차 1

첫 번째로 할 일은 **공통 인터페이스를 만드는 것**이다. `TagNode`의 인터페이스가 더 좋아보이므로 이를 기준으로 인터페이스를 만들 것이다. `TagNode`에는 메서드가 10개 있고, 그 중 public 메서드는 5개이다. 공통 인터페이스에는 그 중 3개만 포함시키면 된다. `TagNode`에 [Extract Interface](https://jo631.github.io/refactoring/2021/04/07/RefactoringToPattern/#extract-interface)를 적용해보자.

```java
public interface XMLNode{
    public abstract void add(XMLNode childNode);
    public abstract void addAttribute(String attribute, String value);
    public abstract void addValue(String value);
}

public class TagNode implements XMLNode{
    public void add(XMLNode childNOde){
        ...
    }
    ...
}
```

##### 절차 2

이제 `DOMBuilder`를 수정할 차례이다. Extract Class를 적용해 `Element`를 위한 **원시 어댑터 클래스**를 만든다. 

```java
public class ElementAdapter{
    Element element;

    public ElementAdapter(Element element){
        this.element = element;
    }
    public Element getElement(){
        return element;
    }
}
```

getter와 (생성자에 포함된)setter를 만들었다.  

##### 절차 3

`DOMBuilder`에 있는 `Element`타입의 필드를 모두 원시 어댑터인 `ElementAdapter`타입으로 변경한다.

```java
class DOMBuilder extends AbstractBuilder{

    private Document document;
    private ElementAdapter root;
    private ElementAdapter parent;
    private ElementAdapter current;

    public void addAtrribute(String name, String value){
        this.current.getElement().setAttribute(name, value);
    }

    public void addBelow(String child){
        ElementAdapter chileNode = new ElementAdapter(document.createElement(child));
        this.current.getElement().appendChild(child.getElement());
        this.parent = this.current;
        this.current = childNode;
        history.push(this.current);
    }

    public void addBeside(String sibling){
        if (this.current == this.root){
            // Exception
        }
        ElementAdapter siblingNode = new ElementAdapter(this.document.createElement(sibling));
        this.parent.getElement().appendChild(siblingNode.getElement());
        this.current = siblingNode;
        history.pop();
        history.push(current);
    }

    public void addValue(String value){
        this.current.appendChild(document.createTextNode(value));
    }
    ...
}
```

##### 절차 4

`DOMBuilder`에서 `Element` 인터페이스의 메서드를 호출하는 부분을 **Extract Method를 통해 별도의 메서드**로 뽑아낸다. 이 때 핵심은 **메서드 호출의 대상이 되는 `Element` 객체에 대한 참조**를 **파라미터를 통해** 얻도록 만드는 것이다.

```java
class DOMBuilder extends AbstractBuilder{

    private Document document;
    private ElementAdapter root;
    private ElementAdapter parent;
    private ElementAdapter current;

    public void addAttribute(String name, String value){
        addAttribute(this.current, name, value);
        //this.current.getElement().setAttribute(name, value);
    }

    private void addAttribute(ElementAdapter current, String name, String value){
        current.getElement().setAttribute(name,value);
    }

    //public void addBelow(String child){
    public void addChild(String child){
        ElementAdapter chileNode = new ElementAdapter(document.createElement(child));
        //this.current.getElement().appendChild(child.getElement());
        add(this.current, childNode);
        this.parent = this.current;
        this.current = childNode;
        history.push(this.current);
    }

    //public void addBeside(String sibling){
    public void addSibling(String siblingName){
        if (this.current == this.root){
            // Exception
        }
        ElementAdapter siblingNode = new ElementAdapter(this.document.createElement(sibling));
        //this.parent.getElement().appendChild(siblingNode.getElement());
        add(this.parent, siblingNode);
        this.current = siblingNode;
        history.pop();
        history.push(current);
    }

    private void add(ElementAdapter parent, ElementAdapter child){
        this.parent.getElement().appendChild(child.getElement());
    }

    public void addValue(String value){
        //this.current.appendChild(document.createTextNode(value));
        addValue(this.currentNode, value);
    }

    private void addValue(ElementAdapter current, String value){
        this.current.appendChild(document.createTextNode(value));
    }
    ...
}
```

##### 절차 5

앞에서 뽑아낸 메서드에 Move Method를 적용해 `ElementAdapter`로 옮기는데, 절차에서 설명했듯이 공통 인터페이스 `XMLNode`의 **대응 메서드와 가능한 한 유사**해야 한다. `addValue(...)`를 제외한 대부분의 메서드에 대해서는 이렇게 인터페이스를 통합하는 데 별 문제가 없다. `addValue(...)`는 조금 뒤로 미루고 다른 것 부터 옮긴다.  

```java
public class ElementAdapter{
    Element element;

    public ElementAdapter(Element element){
        this.element = element;
    }
    
    public Element getElement(){
        return element;
    }

    public void addAttribute(String name, String value){
        this.getElement().setAttribute(name, value);
    }

    public void add(ElementAdapter child){
        this.getElement().appendChild(child.getElement);
    }
}
```

이러면 `DOMBuilder`는 다음과 같이 바꿔야 한다.  

```java
class DOMBuilder extends AbstractBuilder{

    private Document document;
    private ElementAdapter root;
    private ElementAdapter parent;
    private ElementAdapter current;

    public void addAttribute(String name, String value){
        // addAttribute(this.current, name, value);
        this.current.addAttribute(name, value);
    }

    // private void addAttribute(ElementAdapter current, String name, String value){
    //     current.getElement().setAttribute(name,value);
    // }

    public void addChild(String child){
        ElementAdapter chileNode = new ElementAdapter(document.createElement(child));
        // add(this.current, childNode);
        this.current.add(childNode);
        this.parent = this.current;
        this.current = childNode;
        history.push(this.current);
    }

    public void addSibling(String siblingName){
        if (this.current == this.root){
            // Exception
        }
        ElementAdapter siblingNode = new ElementAdapter(this.document.createElement(sibling));
        // add(this.parent, siblingNode);
        this.parent.add(siblingNode);
        this.current = siblingNode;
        history.pop();
        history.push(current);
    }

    // private void add(ElementAdapter parent, ElementAdapter child){
    //     this.parent.getElement().appendChild(child.getElement());
    // }

    public void addValue(String value){
        addValue(this.currentNode, value);
    }

    private void addValue(ElementAdapter current, String value){
        this.current.appendChild(document.createTextNode(value));
    }
    ...
}
```

이 다음 `addValue(...)`를 옮길 차례인데, 이 메서드는 `Document` 필드를 참조하여 조금 까다롭다. 따라서 `ElementAdapter`의 생성자에 `Document` 객체를 전달한다.  

```java
public class ElementAdapter{
    Element element;
    Document document;

    public ElementAdapter(Element element, Document document){
        this.element = element;
        this.document = document;
    }
    
    public Element getElement(){
        return element;
    }

    public void addAttribute(String name, String value){
        this.getElement().setAttribute(name, value);
    }

    public void add(ElementAdapter child){
        this.getElement().appendChild(child.getElement);
    }

    public void addValue(String value){
        this.getElement().appendChild(this.document.createTextNode(value));
    }
}
```

그러면 `DOMBuilder`도 다음과 같이 바뀐다.  

```java
class DOMBuilder extends AbstractBuilder{

    private Document document;
    private ElementAdapter root;
    private ElementAdapter parent;
    private ElementAdapter current;

    public void addAttribute(String name, String value){
        this.current.addAttribute(name, value);
    }

    public void addChild(String child){
        ElementAdapter chileNode = new ElementAdapter(document.createElement(child),document);
        this.current.add(childNode);
        this.parent = this.current;
        this.current = childNode;
        history.push(this.current);
    }

    public void addSibling(String siblingName){
        if (this.current == this.root){
            // Exception
        }
        ElementAdapter siblingNode = new ElementAdapter(this.document.createElement(sibling),document);
        this.parent.add(siblingNode);
        this.current = siblingNode;
        history.pop();
        history.push(current);
    }

    public void addValue(String value){
        this.current.addValue(value);
    }

    // private void addValue(ElementAdapter current, String value){
    //     this.current.appendChild(document.createTextNode(value));
    // }
    ...
}
```

##### 절차 6

`ElementAdapter`가 `XMLNode`를 구현하게 만든다.

```java
public class ElementAdapter implements XMLNode{
    Element element;
    Document document;

    public ElementAdapter(Element element, Document document){
        this.element = element;
        this.document = document;
    }
    
    public Element getElement(){
        return element;
    }

    public void addAttribute(String name, String value){
        this.getElement().setAttribute(name, value);
    }

    public void add(ElementAdapter child){
        ElementsAdapter childElement = (ElementAdapter)child;
        this.getElement().appendChild(child.getElement);
    }

    public void addValue(String value){
        this.getElement().appendChild(this.document.createTextNode(value));
    }
}
```

형변환만 잘 해주면 된다.

##### 절차 7

마지막으로 `DOMBuilder`의 코드 중 `ElementAdapter`타입으로 되어있는 필드 또는 지역 변수, 메서드 파라미터를 모두 `XMLNode` 타입으로 바꾼다.  

```java
class DOMBuilder extends AbstractBuilder{

    private Document document;
    private XMLNode root;
    private XMLNode parent;
    private XMLNode current;

    public void addAttribute(String name, String value){
        this.current.addAttribute(name, value);
    }

    public void addChild(String child){
        XMLNode chileNode = new ElementAdapter(document.createElement(child),document);
        this.current.add(childNode);
        this.parent = this.current;
        this.current = childNode;
        history.push(this.current);
    }

    public void addSibling(String siblingName){
        if (this.current == this.root){
            // Exception
        }
        XMLNode siblingNode = new ElementAdapter(this.document.createElement(sibling),document);
        this.parent.add(siblingNode);
        this.current = siblingNode;
        history.pop();
        history.push(current);
    }

    public void addValue(String value){
        this.current.addValue(value);
    }
    ...
}
```

위 과정을 거쳐 `DOMBuilder`가 사용하는 `Element` 인터페이스를 어댑팅하면 `XMLBuilder`와 `DOMBuilder`의 코드가 매우 비슷해진다. 따라서 [Form Template Method](https://jo631.github.io/refactoring/2021/04/16/Form-Template-Method/)나 [Introduce Polymorphic Creation with Factory Method](https://jo631.github.io/refactoring/2021/04/13/Introduce-Polymorphic-Creation-with-Factory-Method/)를 통해 공통 부분을 수퍼클래스인 `AbstractBuilder`로 옮길 수 있다.

```java
interface XMLNode{
    public void add(XMLNode child);
    public void addAttribute(String name, String value);
    public void addValue(String value);
}

abstract class AbstractBuilder implements XMLNode{
    private XMLNode rootNode;
    private XMLNode currentNode;

    public void addChild(String child){ //Template Method
        XMLNode childNode = this.createNode(child);
        this.currentNode.add(childNode);
        this.currentNode = childNode;
    }

    public void addSibling(String sibling){...} //Template Method

    protected XMLNode createNode(String name); // Factory Method
}

public class ElementAdapter implements XMLNode{
    Element element;
    Document document;

    public ElementAdapter(Element element, Document document){
        this.element = element;
        this.document = document;
    }
    
    public Element getElement(){
        return element;
    }

    public void addAttribute(String name, String value){
        this.getElement().setAttribute(name, value);
    }

    public void add(ElementAdapter child){
        ElementsAdapter childElement = (ElementAdapter)child;
        this.getElement().appendChild(child.getElement);
    }

    public void addValue(String value){
        this.getElement().appendChild(this.document.createTextNode(value));
    }
}

class DOMBuilder extends AbstractBuilder{

    private Document document;
    private XMLNode parent;

    @Override
    protected XMLNode createNode(String name){
        return new ElementAdapter(this.document.createElement(name),document);
    }

    @Override
    public void addAttribute(String name, String value){...}
    @Override
    public void addValue(String value){...}
    
}

class XMLBuilder extends AbstractBuilder{

    @Override
    protected XMLNode createNode(String name){
        return new TagNode(name);
    }

    @Override
    public void addAttribute(String name, String value){...}
    @Override
    public void addValue(String value){...}

}
```