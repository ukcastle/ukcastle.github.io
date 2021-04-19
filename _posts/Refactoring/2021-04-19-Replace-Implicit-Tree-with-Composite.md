---
layout: post
title: Replace Implicit Tree with Composite (249)
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

#### 개요

>실질적으로 트리 구조인 데이터를 String과 같은 기본 타입으로 표현하고 있다면,
그 기본 타입의 표현을 [Composite](https://jo631.github.io/designpattern/2021/04/19/Composite/) 구조로 바꾼다.

#### 동기

데이터나 코드가 명시적으로 트리 구조를 가지는 것은 아니지만, 트리 형태로 표현되는 경우를 **묵시적 트리**라고 한다.  

XML 데이터를 생성하는 예시다.  
```java
String expectedResult = 
    "<orders>" +
        "<order id='321'>" +
            "<product id='f1234'>" +
                "<price=8>"+
                "</price>" +
            "</product>"
            "<product id='p3345'>" +
                "<price=13>"+
                "</price>" +
            "</product>" +
        "</order>"
    "</orders>";
```
##### 대표적인 예시 1
대표적인 이 리팩토링이 필요한 코드이다. 이를 트리 구조로 표현하면 다음과 같다.  
```
orders
ㄴ order = 321
    ㄴ product = f1234
        ㄴ price = 9
    ㄴ product = p3345
        ㄴ price = 13
```


##### 대표적인 예시 2
또한, 위의 예제와 다른 상황도 적용할 수 있다. **조건 로직**이다.  

```java
if(product.getPrice() < price && product.getColor() != color){
    ~~~
}
```

이 조건로직을 트리 구조로 표현하면 다음과 같다.  

```
and(&&)
ㄴ price < target price
ㄴ not(!=)
    ㄴ product color == color
```

두 예시의 본질은 다르지만 Composite 패턴을 이용하여 모델화 할 수 있다는 공통점이 있다.  그렇게 해서 얻는 이점이 무엇일까? **코드를 더 단순하게** 만드는 것이다.  

예를 들어 1번 예시에 Composit을 이용하면 XML을 생성할 때 태그나 속성을 추가하기 위해 코드를 **반복할 필요가 없기 때문에** 코드가 단순해지고 코드의 양이 줄어든다.  
2번 예시에 Composit을 이용하면 위와 비슷한 효과가 있는데, **비슷한 조건 로직이 여러 곳에서 사용되는 경우** 에만 의미가 있다.  

두번째 예시를 리팩토링한다면, 이렇게 될 것이다.  

```java
class ProductFinder{

    public Product[] allProduct;

    public List byColor(Color color){
        List returnList;
        for (Product p : self.allProduct){
            if (p.getColor() == color){
                returnList.append(p);
            }
        }
        return returnList;
    }

    public List byColorAndBelowPrice(Color color, float price){
        List returnList;
        for (Product p : self.allProduct){
            if (p.getColor() == color && p.getPrice < price){
                returnList.append(p);
            }
        }
        return returnList;
    }

    ...
}
```

이렇게 설정하면, 복잡한 조건식을 단순화시키고, 조건식들을 트리구조로 만들 수 있다.  


첫번째 예시를 Composite으로 리팩터링하면 그 결합도는 줄어들지만 클라이언트 코드와 Composite이 꼬이게 된다. 때로는 이런 결합도를 줄이기 위해 완전히 다른 수준의 **인디렉션**이 필요할 수도 있다. 예를 들어 한 프로젝트 내 클라이언트의 코드가 XML을 생성하기 위해 어떤 때는 Composite을 사용하고, 어떤 떄는 DOM(Document Object Model)을 사용할 수도 있다.

#### 장점

- 노드를 추가/삭제/포매팅 하는 등의 반복적인 코드를 캡슐화
- 빈번하게 사용되는 유사한 로직을 다루기 위한 일반화된 방법을 제공
- 클라이언트가 데이터를 생성하는 방법의 단순화

#### 단점

- 묵시적인 트리로도 충분한 경우 괜히 설계가 복잡해짐 

#### 절차

이 리팩터링을 적용하는 경로는 두가지이다. 하나는 표준 방법대로 묵시적 트리를 조금씩 리팩터링해 Composite으로 바꾸는 것이고, 다른 하나는 여기에 테스트 주도 개발(TDD)를 포함시키는 것이다. 묵시적 트리에 [Extract Class](https://jo631.github.io/refactoring/2021/04/07/RefactoringToPattern/#extract-class)같은 리팩터링을 적용하는 것이 여의치 않을 때 TDD를 사용한다.  

1. 묵시적 트리 중 새로운 클래스로 모델화할 수 있는 부분인 **묵시적 종단(Leaf)**를 찾는다. 이 때 새로운 클래스는 종단 노드를 나타내는 것으로, **Composite:Leaf**에 해당한다. 종단 노드 클래스에 Extract Class 또는 TDD를 통해 생성한다.  
묵시적 종단에 속성이 있다면, **각 속성에 해당하는 변수**를 종단 노드 클래스에 만들어 결과적으로 새로 만든 종단 노드 클래스가 묵시적인 종단과 동일한 정보를 나타내게 해야 한다.  

2. 묵시적 종단이 쓰인 곳은 모두 종단 노드 클래스의 인스턴스로 치환해 묵시적 트리가 묵시적인 종단 대신 종단 노드로 구성되도록 한다.  

3. 묵시적 트리에서 묵시적 종단을 나타내는 다른 부분에 대해서도 1,2를 반복한다. 이 때 모든 클래스는 **동일한 인터페이스를 구현**해야 한다.  

4. 묵시적 종단의 부모 역할을 하는 **묵시적인 부모**를 찾는다. 이를 **부모 노드 클래스**라 부를 것이다. 부모 노드 클래스에는 종단 노드를 추가할 수 있는 `add()`와 같은 함수를 추가할 수 잇어야 하고, 부모 노드는 자식 노드를 공통 인터페이스를 통해 다룰 수 있어야 한다. 부모의 부모 노드도 있을 가능성이 다분하므로 부모 노드도 또한 **동일한 인터페이스**를 구현하는 것을 추천한다.  

5. 묵시적 부모가 쓰인 곳을 모두 부모 노드 클래스로 치환한다.  

6. 다른 묵시적 부모에 대해서도 3,4를 반복한다. 묵시적인 부모가 다른 부모의 자식이 될 수 있을 때만, 부모 클래스에서도 그런 동작이 가능하도록 만든다.(같은 인터페이스를 구현)


#### 구현

##### 절차 1

1번 예제 코드를 TDD 방식으로 구현해보겠다.  

```java
String expectedResult = 
    "<orders>" +
        "<order id='321'>" +
            "<product id='f1234' color='red' size='medium'>" +
                "<price currency='USD'>"+
                    "8.95" +
                "</price>" +
                "Fire Truck" +
            "</product>"
            "<product id='p3345'>" +
                "<price=currency='USD'>"+
                    "13.23" +
                "</price>" +
            "</product>" +
        "</order>"
    "</orders>";
```

여기서 첫번째로 결정해야 한다.  
`"8.95"`를 묵시적인 종단으로 볼 것인가? 아니면 `<price> </price>`를 묵시적 종단으로 볼 것인가?  
태그의 값인 "8.95"를 쉽게 표현할 수 있을 것이기 떄문에, 후자를 **묵시적인 종단**으로 선택했다.  
좀 더 살펴보니 모든 XML 태그에는 **이름**이 반드시 있고, **속성(이름/값의 쌍)**, **자식**, **값** 4개의 형태로 이루어져 있다.  
이제 모든 묵시적 종단을 나타내는 **쫑단 노드**의 일반 타입을 만들 수 있을 것이다.  
TDD를 통해 TagNode라는 이름의 클래스를 만든 뒤, 간단한 테스트를 통과시켜 본다.  

간단한 테스트를 만들었다.  

```java
public class TagTest{
    private static final String SAMPLE_PRICE = "8.95";

    public void testTagWithOneAttrAndValue(){
        TagNode priceTag = new TagNode("price");
        priceTag.addAttribute("currency", "USD");
        priceTag.addValue(SAMPLE_PRICE);

        string expected = 
            "<price currency='USD'>"+
                "8.95" +
            "</price>";

        system.out.println(equalTo(priceTag.toString(), expected));
    }
}
```

이를 통과하는 테스트이다.  

```java
public class TagNode{
    private String name = "";
    private String value = "";
    private StringBuffer attributes;

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

    public string toString(){
        String result;
        result = 
            "<" + self.name + self.attribute + ">" +
            self.value +
            "</" + self.name + ">";
        
        return result;
    }
}
```

##### 절차 2

이런 다음, 예제 코드를 수정한다.  

```java

TagNode priceTag_f1234 = new TagNode("price");
priceTag.addAttribute("currency", "USD");
priceTag.addValue("8.95");

TagNode priceTag_p3345 = new TagNode("price");
priceTag.addAttribute("currency", "USD");
priceTag.addValue("13.23");

String result = 
    "<orders>" +
        "<order id='321'>" +
            "<product id='f1234' color='red' size='medium'>" +
                priceTag_f1234.toString() +
                "Fire Truck" +
            "</product>"
            "<product id='p3345'>" +
                priceTag_p3345.toString() +
            "</product>" +
        "</order>"
    "</orders>";
```

##### 절차 3

`TagNode`는 모든 묵시적 종단을 태표할 수 있기 때문에, 단계 1,2를 반복할 필요가 없다.  

##### 절차 4

이제 `묵시적 부모`를 찾을 차례이다.  
예제를 살펴보면  `product` 태그가 `price` 태그의 부모이며, 그 위로 `order`, `orders`로 올라감을 확인할 수 있다.  

부모 노드를 만들기 위한 테스트 코드를 만들었다.  

```java
public void testCompositeTagOneChild(){
    TagNode productTag = new TagNode("product");
    productTag.add(new TagNode("price"));

    String excepted = 
        "<product>" +
            "<price>" +
            "</price>" +
        "<product>";
    
    system.out.println(equalTo(priceTag.toString(), expected));
}
```

 이 테스트를 통과하기 위해 기존 `TagNode`클래스를 수정해보자.  

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

    public string toString(){
        String result;
        result = "<" + self.name + self.attribute + ">";
        Iterator it = self.child().iterator();
        while(it.hasNext()){
            TagNode node = (TagNode)it.next();
            result += node.toString();
        }
        result += self.value;
        result += "</" + self.name + ">";
        
        return result;
    }

    private List child(){
        if(self.child == null){
            self.child == new ArrayList();
        }
        return self.child;
    }

    public void add(TagNode n){
        self.child().add(n)
    }
}
 ``` 

위를 통과하는 코드이다.  

 ##### 절차 5

 이를 모든 부모노드에 적용시킨다.  

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

완성!

반복되는 코드들에 대해 리팩토링을 추가로 진행할 수도 있을 것이다.  
하지만 여기서는 충분하다고 생각되어 종료해야겠다.  

위의 리팩토링은, `TagNode` 단 한개의 클래스로 모든 행동이 가능하여 인터페이스를 추가하지 않았다. 만약 TagNode만으로 해결이 안 될 코드면, 인터페이스를 만들고 `toString()` 메서드를 구현해야 할 것이다.  

