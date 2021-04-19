---
layout: post
title: Encapsulate Composite with Builder (145)
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

#### 개요

>[Composite](https://jo631.github.io/designpattern/2021/04/19/Composite/)구조를 생성하는 과정이 반복적으로 수행되고 복잡하며 에러 발생 가능성도 많은 상태라면,
그 세부사항을 처리하는 별도의 Builder를 제공하여 Composite 구조를 쉽게 생성할 수 있도록 한다.  

#### 동기 

Builder 패턴은 어떤 객체 구조를 생성하는 성가시고 복잡한 과정을 클라이언트 대신 Builder 객체가 맡도록 하는 것이다.  
귀찮은 부분을 빌더에게 맡기면, 클라이언트는 객체 생성 과정에 대한 세부 내용을 알 필요 없이 빌더에게 지시만 하면 된다.  

빌더가 캡슐화하는 객체 구조는 Composite 구조일 확률이 높다. 이를 생성하는 작업은 반복적이고, 복잡하며 에러를 유발하기 쉽기 때문이다.  
예를 들어 부모 노드에 자식 노드를 추가하려면 클라이언트는 다음과 같은 과정을 거쳐야 한다.  

1. 새 노드 객체 생성
2. 새 노드 객체 초기화
3. 새 노드 객체를 적당한 부모 노드에 **정확히** 추가  

만약 클라이언트가 한 노드를 부모 노드에 추가하는 것을 잊거나 엉뚱한 노드에 추가하면 어떻게 될까? 찾아내는 것도 귀찮을것이다.  
그런 코드를 이 리팩터링을 적용하면, 간편하고 에러를 줄이며 단순해진다.  

또 다른 이유로는 클라이언트 코드와 Composite 구조의 **결합 관계**를 제거하는 것이다.  
클라이언트 코드가 직접적으로 결합되어 있으면, 구현한 Composite 구조를 **수정하기 어려워**진다.    

이 때 Builder를 만들면, 추상화된 한 코드만 수정하면 된다.  

이는 Builder패턴의 의도로
> 복잡한 객체를 생성하는 방법과 표현하는 방법을 정의하는 클래스를 별도로 분리하여
서로 다른 표현이라도 동일한 과정을 통해 생성할 수 있도록 하는 것

이라고 설명한다.

복잡한 객체의 서로 다른 표현을 동일한 과정으로 생성하는게 유용한 기능이긴 하지만, 이 것이 빌더의 유일한 기능은 아니다. 생성과정을 단순화 하고 복잡한 과정으로부터 클라이언트 코드를 분리하는 것도 빌더를 사용할 충분한 이유가 된다.  

빌더의 인터페이스는 **누가 보더라도 무슨 일을 수행하는지 알 수 있을 만큼 명확**해야 한다.  
물론 실질적으로 빌더의 인터페이스 전부 또는 일부가 그 정도로 명확하지 않을수도 있다. 내부적으로도 많은 일을 해야 하기 떄문이다. 따라서 빌더의 인터페이스를 잘 이해하려면, 코드 또는 테스트 코드를 참고하거나 문서를 읽어보는 등 노력이 필요하다.  

#### 장점

- Composite 구조를 생성하는 클라이언트의 코드를 단순화한다.  
- 반복적이고 에러 발생 가능성이 높은 Composite 구조 생성 작업의 단점을 개선한다.  
- 클라이언트 코드와 Composite 사이의 결합을 느긋하게 한다.  
- 캡슐화된 Composite 또는 복잡한 객체의 여러 다른 표현이 가능하게 한다.  

#### 단점

- 인터페이스의 의도가 덜 명확해질 수 있다.  

#### 절차

1. 빌더로 삼을 새 클래스를 만든다. 노드가 난 하나인 Composite 구조를 생성할 수 있도록 구현하고, 생성한 결과를 리턴하는 함수도 추가한다.  

2. 빌더에 자식 노드를 생성하는 기능을 추가한다. 클라이언트가 자식을 생성하고 삽입 위치를 지정할 때 사용할 몇 개의 메서드를 구현하는 과정이 포함 될 것이다.  

3. 기존의 Composite 생성 코드에 노드의 속성이나 값을 변경하는 부분이 있다면, 빌더에도 그런 속성이나 값을 지정할 수 있는 기능을 구현한다.  

4. 빌더의 현재 상태가 적합한지 생각해보고, 필요하다면 개선한다.  

5. 기존의 Composite 생성 코드를 리팩터링하여 앞서 만든 빌더를 사용하도록 한다.  

#### 구현

일전에 Composite 패턴으로 만든 코드가 하나 있다.  
이번 리팩토링에 딱 적합한 예제이다.  

```java
public class TagNode{
    private String name = "";
    private String value = "";
    private StringBuffer attributes;

    private List child; 

    public TagNode(String name){
        this.name = name;
        this.attributes = new StringBuffer("");
    }
    public void addAttribute(String attribute, String value){
        this.attributes.append(" ");
        this.attributes.append(attribute);
        this.attributes.append("='");
        this.attributes.append(value);
        this.attributes.append("'");
    }
    public void addValue(String value){
        this.value = value;
    }

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

    private List child(){
        if(this.child == null){
            this.child == new ArrayList();
        }
        return this.child;
    }

    public void add(TagNode n){
        this.child().add(n)
    }
}
```

```java
TagNode priceTag_f1234 = new TagNode("price");
priceTag_f1234.addAttribute("currency", "USD");
priceTag_f1234.addValue("8.95");

TagNode productTag_f1234 = new TagNode("product");
productTag_f1234.addAttribute("id", "f1234");
productTag_f1234.addAttribute("color", "red");
productTag_f1234.addAttribute("size", "medium");
productTag_f1234.addValue("Fire Truck");

productTag_f1234.add(priceTag_f1234);


TagNode priceTag_p3345 = new TagNode("price");
priceTag_p3345.addAttribute("currency", "USD");
priceTag_p3345.addValue("13.23");

TagNode productTag_p3345 = new TagNode("product");
productTag_f1234.addAttribute("id", "p3345");

productTag_f1234.add(priceTag_p3345);


TagNode orderTag_321 = new TagNode("order");
orderTag_321.addAttribute("id", "321");

orderTag_321.add(productTag_f1234);
orderTag_321.add(productTag_p3345);


TagNode ordersTag = new TagNode("orders");
ordersTag.add(orderTag_321);
```

`TagNode`는 XML 생성을 쉽게 해주는 클래스로, Composite의 3요소인 `Component, Leaf, Composite`의 역할을 모두 수행한다.  

`toString()`은 모든 TreeNode에 대한 XML 문자열을 `return` 한다.  
우리의 목표는 클라이언트가 TagNode를 좀 더 쉽게 생성할 수 있도록 `TagNode`를 `TagBuilder`로 **캡슐화**하는 것이다.  
이번에도 **TDD** 방식으로 해보겠다.  


##### 절차 1

>빌더로 삼을 새 클래스를 만든다. 노드가 난 하나인 Composite 구조를 생성할 수 있도록 구현하고, 생성한 결과를 리턴하는 함수도 추가한다. 

우선 한개 의 노드로 구성된 Composite 구조를 생성할 수 있는 빌더를 만든다.  
즉, 하나의 TagNode 객체만을 포함한 트리로부터 정확한 XML을 생성하는 TagBuilder 클래스를 구현하는 것이다.  
먼저 다음과 같은 테스트 코드를 작성한다.  

```java
public void testBuildOneNode(){
    String expectedXml = 
        "<flavors>" +
        "</flavors>";
    String actualXml = new TagBuilder("flavors").toXml();

    system.out.println(equalTo(actualXml, expectedXml));
}
```

이를 통과하는 코드를 만든다.  

```java
public class TagBuilder{
    private TagNode rootNode;

    public TagBuilder(String rootTagName){
        this.rootNode = new TagNode(rootTagName);
    }

    public String toXml(){
        return this.rootNode.toString();
    }
}
```
##### 절차 2

>빌더에 자식 노드를 생성하는 기능을 추가한다. 클라이언트가 자식을 생성하고 삽입 위치를 지정할 때 사용할 몇 개의 메서드를 구현하는 과정이 포함 될 것이다. 

만약 자식 객체를 다루는 **다양한** 경우, 모두 처리할 수 있어야 하므로 각 경우마다 `TagBuilder`에 **별도의 함수**를 추가해야만 한다.   

루트에 자식 **하나**를 추가하는 경우를 먼저 처리하자. 자식 노드를 하나 생성하여 캡슐화된 Composite 구조 내 정확한 위치에 삽입할 수 있어야 하므로, 그런 작업을 하는 `addChild()` 함수를 구현한다. 이를 테스트할 코드는 다음과 같다.  

```java
public void testBuildOneChild(){
    String expectedXml = 
        "<flavors>" +
            "<flavor>" +
                "<requirments>" +
                    "<requirment>" +
                    "</requirment>" +
                "</requirments>" +
            "</flavor>" +
        "</flavors>";

    TagBuilder builder = new TagBuilder("flavors");
    builder.addChild("flavor");
    builder.addChild("requirments");
    builder.addChild("requirment");
    String actualXml = new builder.toXml();
}
```

이 테스트를 통과하게 구현해보자.  

```java
public class TagBuilder{
    private TagNode rootNode;
    private TagNode currentNode;

    public TagBuilder(String rootTagName){
        this.rootNode = new TagNode(rootTagName);
        this.currentNode = this.rootNode;
    }

    public void addChild(String childTagName){
        TagNode parentNode = this.currentNode;
        this.currentNode = new TagNode(childTagName);
        parentNode.add(this.currentNode);
    }

    public String toXml(){
        return this.rootNode.toString();
    }
}
```
여기까진 쉽다.  

이제 형제 노드를 추가하는 테스트코드를 만들어보자.  

기존의 자식 노드에 대한 형제 노드를 추가한다는 것은, 공통의 부모 노드를 TagBuilder가 식별할 수 있도록 해야 한다. 하지만 현재 `TagNode` 객체가 자신의 부모 노드에 대한 참조를 갖고 있찌 않으므로, 현재로써는 그럴 방법이 없다.  
따라서 TagNode에 부모에 대한 참조를 갖게 만드는 코드를 설정한다.  

그에 대한 테스트 코드이다.  
```java
public void testTagNode(){
    TagNode root = new TagNode("root");

    TagNode child = new TagNode("child");
    root.add(child);

    system.out.println(equalTo(root.toString(), child.getParent().toString()));
}
```

이를 통과하기 위한 코드를 만들어보자.  

```java
public class TagNode{
    private String name = "";
    private String value = "";
    private StringBuffer attributes;

    private TagNode parent;
    private List child; 
    
    ...

    public void add(TagNode child){
        child.setParent(this);
        this.child().add(child);
    }

    private void setParent(TagNode parent){
        this.parent = parent;
    }

    public TagNode getParent(){
        return this.parent;
    }
```

우아하게 성공했다.  

이제 본론으로 돌아가자.  

```java
public void testBuildSibling(){
    String expectedXml = 
        "<flavors>" +
            "<flavor>" +
            "</flavor>" +
            "<flavor>" +
            "</flavor>" +
        "</flavors>";

    TagBuilder builder = new TagBuilder("flavors");
    builder.addChild("flavors");
    builder.addSibling("flavors");

    String actualXml = new builder.toXml();
    system.out.println(equalTo(expectedXml, actualXml));
}
```
형제 노드를 추가하는 테스트코드 이다.  

이를 통과시켜보자.  

```java
public class TagBuilder{
    private TagNode rootNode;
    private TagNode currentNode;

    public TagBuilder(String rootTagName){
        this.rootNode = new TagNode(rootTagName);
        this.currentNode = this.rootNode;
    }

    public void addChild(String childTagName){
        this.addTo(this.currentNode, childTagName);
    }

    public void addSibling(String siblingTagName){
        this.addTo(this.currentNode.getParent(), siblingTagName);
    }

    public void addTo(TagNode parentNode, String tagName){
        this.currentNode = new TagNode(tagName);
        parentNode.add(currentNode);
    }

    public String toXml(){
        return this.rootNode.toString();
    }
}
```

성공했다.  근데 뭔가 부족하다. 아래 테스트코드도 실행시켜볼까?  

```java
public void testBuildSibling(){
    String expectedXml = 
        "<flavors>" +
            "<flavor>" +
                "<requirments>" +
                    "<requirment>" +
                    "</requirment>" +
                "</requirments>" +
            "</flavor>" +
            "<flavor>" +
                "<requirments>" +
                    "<requirment>" +
                    "</requirment>" +
                "</requirments>" +
            "</flavor>" +
        "</flavors>";

    TagBuilder builder = new TagBuilder("flavors");
    for (int i=0; i<2; i++){
        builder.addChild("flavor");
        builder.addChild("requirments");
        builder.addChild("requirment");
    }


    String actualXml = new builder.toXml();
    system.out.println(equalTo(expectedXml, actualXml));
}
```

무사히 컴파일되는가? 당연히 **안된다.**  
만약 반복문이 아니고 하드코딩으로 `addSibling()`으로 한다면?  
그것도 아니다. requirments 위에 생길것이다.  

이 것은 addChild()를 고쳐서 해결할 문제가 아니다. **특정 부모 노드를 지정해 자식을 추가하는 함수**가 필요하다.  

그에 대한 테스트 코드를 작성하자.  

```java
public void testBuildSibling(){
    String expectedXml = 
        "<flavors>" +
            "<flavor>" +
                "<requirments>" +
                    "<requirment>" +
                    "</requirment>" +
                "</requirments>" +
            "</flavor>" +
            "<flavor>" +
                "<requirments>" +
                    "<requirment>" +
                    "</requirment>" +
                "</requirments>" +
            "</flavor>" +
        "</flavors>";

    TagBuilder builder = new TagBuilder("flavors");
    for (int i=0; i<2; i++){
        builder.addToParent("flavors", "flavor");
        builder.addChild("requirments");
        builder.addChild("requirment");
    }

    String actualXml = new builder.toXml();
    system.out.println(equalTo(expectedXml, actualXml));
}
```

이제, 이를 통과하게 바꿔보자.  

```java
public class TagBuilder{
    private TagNode rootNode;
    private TagNode currentNode;

    public TagBuilder(String rootTagName){
        this.rootNode = new TagNode(rootTagName);
        this.currentNode = this.rootNode;
    }

    public void addChild(String childTagName){
        this.addTo(this.currentNode, childTagName);
    }

    public void addSibling(String siblingTagName){
        this.addTo(this.currentNode.getParent(), siblingTagName);
    }

    public void addToParent(String parentTagName, String childTagName){
        addTo(findParentBy(parentTagName), childTagName);
    }

    private TagNode findParentBy(String parentName){
        TagNode parentNode = currentNode;

        while(parentNode != null){
            if(parentNode.equals(parentNode.getName())){
                return parentNode;
            }
            parentNode = parentNode.getParent();
        }
        return null;
    }

    public void addTo(TagNode parentNode, String tagName){
        this.currentNode = new TagNode(tagName);
        parentNode.add(currentNode);
    }

    public String toXml(){
        return this.rootNode.toString();
    }
}
```
이제 테스트를 잘 통과한다.  

하지만 예외 처리를 위한 테스트 코드도 추가하자.  

```java
public void testParentNameNotFound(){
    TagBuilder builder = new TagBuilder("flavors");
    try{
        for (int i=0; i<2; i++){
            builder.addToParent("NotFound","flavor");
        }
    } catch (RuntimeException e){
        String expectErrorMsg = ("NotFound"+" 태그를 찾을 수 없습니다.");
        system.out.println(equalTo(expectedErrorMsg, e.getMessage()));   
    }
}
```

```java
public class TagBuilder{
    private TagNode rootNode;
    private TagNode currentNode;

    public TagBuilder(String rootTagName){
        this.rootNode = new TagNode(rootTagName);
        this.currentNode = this.rootNode;
    }

    public void addChild(String childTagName){
        this.addTo(this.currentNode, childTagName);
    }

    public void addSibling(String siblingTagName){
        this.addTo(this.currentNode.getParent(), siblingTagName);
    }

    public void addToParent(String parentTagName, String childTagName){
        TagNode parentNode = findParentBy(parentTagName);
        if (parentNode == null){
            throw new RuntimeException(parentTagName+" 태그를 찾을 수 없습니다.");
        }
        addTo(findParentBy(parentTagName), childTagName);
    }
    
    private TagNode findParentBy(String parentName){
        TagNode parentNode = currentNode;

        while(parentNode != null){
            if(parentNode.equals(parentNode.getName())){
                return parentNode;
            }
            parentNode = parentNode.getParent();
        }
        return null;
    }

    public void addTo(TagNode parentNode, String tagName){
        this.currentNode = new TagNode(tagName);
        parentNode.add(currentNode);
    }

    public String toXml(){
        return this.rootNode.toString();
    }
}
```

예외 처리까지 성공했다!  

##### 절차 3

>기존의 Composite 생성 코드에 노드의 속성이나 값을 변경하는 부분이 있다면, 빌더에도 그런 속성이나 값을 지정할 수 있는 기능을 구현한다.  

거의 다 왔다.  

이미 `TagNode`에 속성을 처리하고 있으므로, 쉽게 추가할 수 있다.  

이에 대한 테스트 코드이다.  

```java
public void testBuildSibling(){
    String expectedXml = 
        "<flavors>" +
            "<flavor name='Test-Driven Development'>" + //속성을 가짐
                "<requirments>" +
                    "<requirment type='hardware'>" +
                        "1 Computer for Every 2 Participants" + //값을 가짐
                    "</requirment>" +
                "</requirments>" +
            "</flavor>" +
            "<flavor>" +
                "<requirments>" +
                    "<requirment type='software'>" +
                        "IDE" +
                    "</requirment>" +
                "</requirments>" +
            "</flavor>" +
        "</flavors>";

    TagBuilder builder = new TagBuilder("flavors");
        builder.addChild("flavor");
        builder.addAttribute("name", "Test-Driven Development");
            builder.addChild("requirments");
                builder.addChild("requirment");
                builder.addAttribute("type", "hardware");
                builder.addValue("1 Computer for Every 2 Participants");
        builder.addToParent("flavor");
            builder.addChild("requirments");
                builder.addChild("requirment");
                builder.addAttribute("type", "software");
                builder.addValue("IDE");
        
    String actualXml = new builder.toXml();
    system.out.println(equalTo(expectedXml, actualXml));
}
```

간단하게도 2개의 함수만 넣으면 된다. `addAttrivute`와 `addValue` 이다.  

```java
public class TagBuilder{
    private TagNode rootNode;
    private TagNode currentNode;

    public TagBuilder(String rootTagName){
        this.rootNode = new TagNode(rootTagName);
        this.currentNode = this.rootNode;
    }

    public void addChild(String childTagName){
        this.addTo(this.currentNode, childTagName);
    }

    public void addSibling(String siblingTagName){
        this.addTo(this.currentNode.getParent(), siblingTagName);
    }

    public void addToParent(String parentTagName, String childTagName){
        TagNode parentNode = findParentBy(parentTagName);
        if (parentNode == null){
            throw new RuntimeException(parentTagName+" 태그를 찾을 수 없습니다.");
        }
        addTo(findParentBy(parentTagName), childTagName);
    }
    
    private TagNode findParentBy(String parentName){
        TagNode parentNode = currentNode;

        while(parentNode != null){
            if(parentNode.equals(parentNode.getName())){
                return parentNode;
            }
            parentNode = parentNode.getParent();
        }
        return null;
    }

    public void addTo(TagNode parentNode, String tagName){
        this.currentNode = new TagNode(tagName);
        parentNode.add(currentNode);
    }

    public void addAttribute(String name, String value){
        this.currentNode.addAttribute(name,value);
    }

    public void addValue(String value){
        this.currentNode.addValue(value);
    }

    public String toXml(){
        return this.rootNode.toString();
    }
}
```

##### 절차 4

>빌더의 현재 상태가 적합한지 생각해보고, 필요하다면 개선한다.  

이는 코드만 보고 정할 수 있는 절차가 아니다.  
추후 계속 테스트해보면서 알아보자!  

##### 절차 5

>기존의 Composite 생성 코드를 리팩터링하여 앞서 만든 빌더를 사용하도록 한다.  

기존의 생성 코드는 이렇다.

```java
TagNode priceTag_f1234 = new TagNode("price");
priceTag_f1234.addAttribute("currency", "USD");
priceTag_f1234.addValue("8.95");

TagNode productTag_f1234 = new TagNode("product");
productTag_f1234.addAttribute("id", "f1234");
productTag_f1234.addAttribute("color", "red");
productTag_f1234.addAttribute("size", "medium");
productTag_f1234.addValue("Fire Truck");

productTag_f1234.add(priceTag_f1234);


TagNode priceTag_p3345 = new TagNode("price");
priceTag_p3345.addAttribute("currency", "USD");
priceTag_p3345.addValue("13.23");

TagNode productTag_p3345 = new TagNode("product");
productTag_f1234.addAttribute("id", "p3345");

productTag_f1234.add(priceTag_p3345);


TagNode orderTag_321 = new TagNode("order");
orderTag_321.addAttribute("id", "321");

orderTag_321.add(productTag_f1234);
orderTag_321.add(productTag_p3345);


TagNode ordersTag = new TagNode("orders");
ordersTag.add(orderTag_321);

String result = ordersTag.toString();
```

확실히, 오류가 나기 매우 쉬울 것 같다.  

고쳐보자.  

위의 트리구조는 다음과 같다.

```
orders
ㄴ order
    ㄴ product 
        ㄴ price
    ㄴ product
        ㄴ price
```

```java
TagBuilder builder = new TagBuilder("orders");
    builder.addChild("order");
    builder.addAttribute("id","321");
        builder.addToParent("order","product");
        builder.addAttribute("id", "f1234");
        builder.addAttribute("color", "red");
        builder.addAttribute("size", "medium");
        builder.addValue("Fire Truck");
            builder.addChild("price");
            builder.addAttribute("currency", "USD");
            builder.addValue("8.95");
        builder.addToParent("order","product");
        builder.addAttribute("id", "p3345");
            builder.addChild("price");
            builder.addAttribute("currency", "USD");
            builder.addValue("13.23");
```

기존의 코드보다 정말 보기 쉽고, 아름다운 코드로 리팩토링됐다.  



