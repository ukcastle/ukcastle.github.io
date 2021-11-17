---
layout: post
title: Extract Composite (291)
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

#### 개요

>한 상속 구조 내의 서브클래스가 동일한 [Composite](https://ukcastle.github.io/designpattern/2021/04/19/Composite/)기능을 **각자** 구현하고 있다면
컴포짓 기능을 수퍼클래스로 옮겨 구현한다.  

<br>

#### 동기

해당 리팩터링은 [Extract Superclass](https://ukcastle.github.io/refactoring/2021/04/07/RefactoringToPattern/#extract-superclass)와 유사하다. 다른점은 Composite 기능에 대한 것을 다룬다는 점이다.  

어떤 상속 구조 내의 서브클래스들이 자신의 자식 객체를 컬렉션에 저장하고 그 **자식 객체들의 기능에 접근하기 위한 메서드를 각자 구현**하는 것은 자주 볼 수 있다. 이 때 만약 자식 객체의 타입이 부모 객체와 **동일한 상속 구조 내의 타입**이라면, Composite 패턴으로 리팩터링해 많은 코드 중복을 제거할 수 있다.  

이 리팩터링은 Extract Superclass와 근본적으로 동일하지만 **대상**이 다르다. 자식 객체 처리 로직을 수퍼클래스로 올리는 일을 할 때는 해당 리팩터링을 이용하고 그 후에도 중복된 로직이 남아있다면 Extract Superclass 리팩터링을 적용하면 된다. 

<br>

#### 장점

- 중복된 자식 객체 저장/처리 로직을 제거한다. 
- 자식 객체 처리 로직을 상속 받아 그대로 사용할 수 있음이 명확히 드러난다.  

<br>

#### 절차

이 리팩터링은 Extracat Superclass 리팩터링의 절차를 기반으로 한다.  

1. 처음에는 Composite 기능이 없는 상태지만 **리팩터링 후에는 Composite가 될 클래스**를 하나 만든다. 그 클래스의 이름에는 **추후에 다룰 객체의 종류를 표시**하는 것이 좋다. (예를 들어 CompositeTag)  

2. 자식 객체 컨테이너(상속 구조에서 **중복된 자식 객체 처리 로직을 포함하는 클래스**)를 앞서 만든 **컴포짓 클래스의 서브 클래스**로 만든다.  

3. 자식 객체 컨테이너 사이에 **중복된 자식 객체 처리 메서드를** 하나 찾아낸다.이 때 두가지 경우가 있는데, 하나는 **메서드 몸체 구현이 완전히 동일한 경우**이고 **다른 하나는 메서드 몸체 구현에 공통되는 부분과 다른 부분이 혼재하는 경우**이다. 전자의 경우 **완전 중복 메서드** 라고 부르고 후자의 경우를 **부분 중복 메서드**라고 한다.  
찾아낸 중복 메서드의 이름이 통일되어있지 않다면 [Rename Method](https://ukcastle.github.io/refactoring/2021/04/07/RefactoringToPattern/#rename-method)를 적용한다. 
- **완전 중복 메서드**의 경우 그 메서드에서 사용하는 자식 객체 컬렉션 필드를 [Pull Up Field](https://ukcastle.github.io/refactoring/2021/04/07/RefactoringToPattern/#pull-up-method) 리팩터링을 적용해 **Composite 클래스로 옮긴다.** 이 때 그 필드는 **모든 자식 객체 컨테이너에게 의미 있는 일반적인 이름**으로 바꾼다. 이 후 메서드도 동일한 방법으로 Composite 클래스로 옮긴다. 이 떄 그 메서드가 클래스의 생성자 코드에 **의존**하는 부분이 있다면, 그 **생성자 코드도 Composite 클래스의 생성자로** 올려야 한다. 
- 부분 중복 메서드의 경우 일단 [Substitute Algorithm](https://ukcastle.github.io/refactoring/2021/04/28/substitute-Algorithm/)을 통해 **메서드 구현을 모두 동일하게 만들 수 있는지** 살핀다. 만약 그럴 수 있다면 리팩터링을 통해 완전 중복 메서드로 만든다. 그렇지 않다면 [Extract Method](https://ukcastle.github.io/refactoring/2021/04/07/RefactoringToPattern/#extract-method)를 통해 **공통의 부분을 별도의 메서드를 뽑아낸 뒤 Composite 클래스로 옮긴다.** 만약 메서드의 내부 구현이 동일한 단계를 따르지만 각 단계가 조금씩 다르다면 [Form Template Method](https://ukcastle.github.io/refactoring/2021/04/16/Form-Template-Method/)를 고려해본다.

4. 다른 중복된 자식 객체 처리 메서드를 찾아 각각 단계 3을 적용한다.  

5. 가능하다면 클라이언트 코드에서 각 자식 객체 컨테이너를 사용할 때 일률적으로 Composite 클래스 타입만을 통하도록 만든다.  

<br>

#### 구현

예를 들어 다음과 같은 HTML 문서가 있다고 하자.

```html
<HTML>
    <BODY>
        Hello, World!
        <A HREF="https://ukcastle.github.io/">
            <IMG SRC="https://ukcastle.github.io/public/img/moon.jpg">
        </A>
    </BODY>
</HTML>
```

이런 HTML 문서의 파서를 작성할 때, HTML 문서를 읽어가며 태그나 문자열을 만나면 그에 대응하는 HTML 요소 객체를 생성한다. 예를 들면 이런 객체들을 생성한다.  
- \<body> 등의 태그를 위한 각종 `Tag` 객체
- 'Hello, World!' 와 같은 문자열을 위한 `StringNode` 객체
- \<A HREF="..."> 태그에 대응되는 `LinkTag` 객체
- \<IMG> 태그에 대응되는 `ImageTag` 객체

이런 태그들을 관리하는 코드가 있다.  

```java

public class LinkTag extends Tag{
    private Vector nodeVector;

    @Override
    public String toPlainTextStirng(){
       StringBuffer sb = new StringBuffer();
       Node node;
       for (Enumeration e=linkData() ; e.hasMoreElements(); ){
           node = (Node)e.nextElemnet();
           sb.append(node.toPlainTextString());
       } 
       return sb.toString()
    }
}

public class FormTag extends Tag{
    private Vector allNodesVector;

    @Override
    public String toPlainTextStirng(){
       StringBuffer stringRepresentation = new StringBuffer();
       Node node;
       for (Enumeration e=getAllNodesVector() ; e.hasMoreElements(); ){
           node = (Node)e.nextElemnet();
           stringRepresentation.append(node.toPlainTextString());
       } 
       return stringRepresentation.toString()
    }
}
```

두 태그는 모두 자식 객체를 가질 수 있기에 `Vector`라는 필드를 가진다. 필드의 이름은 다르지만 말이다. 그리고 두 클래스의 `toPlainTextString()`을 보면 자식 객체들을 순회하며 각각의 정보를 바탕으로 평문을 생성하는 로직을 **각자** 구현하고 있다. 그런데 그 코드가 거의 같다. 이럴 때 Extract Composite 리팩터링을 적용할 수 있다.  

##### 절차 1

자식 컨테이너 클래스들의 수퍼 클래스가 될 추상 클래스를 만든다. 그런데 두 클래스는 모두 `Tag`의 서브클래스이기 때문에 다음과 같이 만들 수 있다.  

```java

public abstract class CompositeTag extends Tag{
    public CompositeTag(int tagBegin, int tagEnd, String tagContents, String tagLine){
        super(tagBegin, tagEnd, tagContents, tagLine);
    }
}

```

##### 절차 2

자식 컨테이너 클래스들을 앞에서 만든 `CompositeTag`의 서브클래스가 되도록 수정한다.  

```java

public class LinkTag extends CompositeTag

public class FormTag extends CompositeTag
```

##### 절차 3

모든 자식 컨테이너를 조사해 **완전 중복 메서드**를 찾는다. 이번 예제의 경우 `toPlainTextString()`이다.  
첫 단계는 자식 객체를 저장하는 `Vector` 필드를 수퍼클래스로 옮기는 것이다. 

```java

public abstract class CompositeTag extends Tag{
    protected Vector children;
}

public class LinkTag extends CompositeTag{
    // 필드 없음!

    public Enumeration linkData(){
        return this.children.element();
    }

    @Override
    public String toPlainTextStirng(){
       StringBuffer sb = new StringBuffer();
       Node node;
       for (Enumeration e=linkData() ; e.hasMoreElements(); ){
           node = (Node)e.nextElemnet();
           sb.append(node.toPlainTextString());
       } 
       return sb.toString()
    }
}

public class FormTag extends CompositeTag{
    // 필드 없음!

    public Vector getAllNodesVector(){
        return this.children;
    }

    @Override
    public String toPlainTextStirng(){
       StringBuffer stringRepresentation = new StringBuffer();
       Node node;
       for (Enumeration e=getAllNodesVector().elements() ; e.hasMoreElements(); ){
           node = (Node)e.nextElemnet();
           stringRepresentation.append(node.toPlainTextString());
       } 
       return stringRepresentation.toString()
    }
}
```

좀 더 명료하게 하기 위해 필드 이름을 children 으로 바꿨다.  

다음은 toPlainTextString() 메서드를 바꿀 차례다.
위의 두 클래스는 해당 메서드에서 비슷한 내용을 구현하지만, 살짝 다르다. 이를 동일하게 하기 위해 하나하나 고쳐나가야 한다.  

```java

public abstract class CompositeTag extends Tag{
    protected Vector children;

    public Enemeration children(){
        return this.children.elements();
    }

    public String toPlainTextStirng(){
        StringBuffer sb = new StringBuffer();
        Node node;

        for (Enemeration e = this.children(); e.hasMoreElements(); ){
            node = (Node)e.nextElement();
            sb.append(node.toPlainTextString());
        }
        return sb.toString();
    }
}

```

##### 절차 4

나머지 중복 메서드에 대해서도 CompositeTag로 옮긴다. 완전 중복 메서드에 대해 다뤄봤으니, 이젠 부분 중복 메서드에 대해서도 알아보자.  

```java

public class LinkTag extends CompositeTag{
    @Override
    public String toHTML(){
        StringBuffer sb = new StringBuffer();
        this.putLinkStartTagInto(sb);
        Node node;
        for (Enemeration e = this.children(); e.hasMoreElements(); ){
            node = (Node)e.nextElement();
            sb.append(node.toHTML());
        }
        sb.append("</A>");
        return sb.toString();
    }

    public void putLinkStartTagInto(StringBuffer sb){
        sb.append("<A ");
        String key,value;
        int i=0;
        for (Enemeration e = parsed.keys(), e.hasMoreElements(); ){
            key = (String)e.nextElement();
            i++;
            if(key!=TAGNAME){
                value = getParameter(key);
                sb.append(key+"=\""+value+"\"");
                if (i(parsed.size()-1){
                    sb.append(" ");
                }
            }
        }
        sb.append(">");
    }
}

```

LinkTag 클래스는 버퍼를 생성한 뒤 `putLinkStartTagInto(..)` 메서드를 통해 시작 태그 속성 내용을 버퍼에 채운다. 시작 태그가 만약 \<A HREF="...">라면 속성은 HREF에 해당한다. 그 다음 태그 사이의 문자열에 해당하는 `StringNode` 또는 `ImageTag` 객체를 자식으로 가질 수 있으므로 자식 객체를 하나씩 순회하며 그 내용을 처리한 뒤 마지막으로 `</A>`를 버퍼에 쓴다. 


```java

public class FormTag extends CompositeTag{
    @Override
    public String toHTML(){
        StringBuffer rawBuffer = new StringBuffer();
        this.putLinkStartTagInto(sb);
        Node node, prevNode=null;
        rawBuffer.append("<FORM METHOD=\""+formMethod+"\" ACTION=\""+formURL+"\"");
        if (formName!=null && formName.length() >0){
            rawBuffer.append(" NAME=\""+formName+"\"");
        }
        for (Enemeration en = table.keys(); en.hasMoreElements(); ){

            key=(String)en.nextElement();

            List keys = Arrays.asList(new String[] {"METHOD","ACTION","NAME",tag.TAGNAME})

            for ( i : keys ){
                if(!key.equals(i)){
                    value = (String)table.get(key);
                    rawBuffer.append(" "+key+"="+"\""+value+"\"");
                    break;
                }
            }

            rawBuffer.append(">");
            rawBuffer.append(lineSeparator);
            for(;e.hasMoreElements();){
                node = (Node)e.nextElement();
                if(prevNode!=null){
                    if(prevNode.elementEnd()>node.elementBegin()){
                        rawBuffer.append(lineSeparator);
                    }
                }
                rawBuffer.append(node.toHTML());
            }
            prevNode=node;
        }
        return rawBuffer.toString();
    }
}

```

`FormTag`의 toHTML() 구현은 `LinkTag`와 비교할 때 일부는 비슷하지만 그렇지 않은 부분도 있다. 따라서 `toHTML()`은 **부분 중복 메서드** 이다. 따라서 절차에서 말한대로 Substitute Algorithm 리팩터링으로 완전 중복 메서드로 만들 수 있는지 살펴봐야 한다.  
결론부터 말하면, 가능하다. 심지어 보기보다 쉽다! `toHTML()`의 구조를 살펴보면 다음과 같다.  
- 시작 태그 및 그 속성에 대한 처리
- 자식 태그에 대한 처리
- 종료 태그에 대한 처리

위의 사실을 고려하면 각 과정에 해당하는 공통 메서드를 `CompositeTag`에 만들고, 두 서브클래스에 그 메서드를 사용하여 toHTML()을 구현하면 된다는 것을 알 수 있다. 한번 해보자.  


```java

public abstract class CompositeTag extends Tag{
    protected Vector children;

    public Enemeration children(){
        return this.children.elements();
    }

    public String toPlainTextStirng(){
        StringBuffer sb = new StringBuffer();
        Node node;

        for (Enemeration e = this.children(); e.hasMoreElements(); ){
            node = (Node)e.nextElement();
            sb.append(node.toPlainTextString());
        }
        return sb.toString();
    }

    public void putStartTagInto(StringBuffer sb){
        sb.append("<" + this.getTagName() + " ");
        String key, value;
        int i=0;
        for (Enumeration e = parsed.keys(); e.hasMoreElements(); ){
            key = (String)e.nextElement();
            i++;
            if(key != TAGNAME){
                value = this.getParameter(key);
                sb.append(key+"=\""+value+"\"");
                if(i<parsed.size()){
                    sb.append(" ");
                }
            }
        }
        sb.append(">");
    }
    ...

    public void putChildrenTagInto(StringBuffer sb){ ... }
    public void putEndTagInto(StringBuffer sb){ ... }
}

public class LinkTag extends CompositeTag{
    @Override
    public String toHTML(){
        StringBuffer sb = new StringBuffer();
        this.putStartTagInto(sb);
        ...
    }
}

public class FormTag extends CompositeTag{
    @Override
    public String toHTML(){
        StringBuffer rasBuffer = new StringBuffer();
        this.putStartTagInto(rawBuffer);
        ...
    }
}

```

자식 태그와 종료 태그에서도 같은 작업을 반복하면, toHTML()을 `CompositeTag`로 옮길 수 있다.

```java
public abstract class CompositeTag extends Tag{
    public abstract class CompositeTag extends Tag{
    protected Vector children;

    public Enemeration children(){
        return this.children.elements();
    }

    public String toPlainTextStirng(){
        StringBuffer sb = new StringBuffer();
        Node node;

        for (Enemeration e = this.children(); e.hasMoreElements(); ){
            node = (Node)e.nextElement();
            sb.append(node.toPlainTextString());
        }
        return sb.toString();
    }

    public void toHTML(){
        StringBuffer htmlContents = new StringBuffer("");
        this.putStartTagInto(htmlContents);
        this.putChildrenTagInto(htmlContents);
        this.putEndTagInto(htmlContents);
        return htmlContents.toString();
    }

    public void putStartTagInto(StringBuffer sb){
        sb.append("<" + this.getTagName() + " ");
        String key, value;
        int i=0;
        for (Enumeration e = parsed.keys(); e.hasMoreElements(); ){
            key = (String)e.nextElement();
            i++;
            if(key != TAGNAME){
                value = this.getParameter(key);
                sb.append(key+"=\""+value+"\"");
                if(i<parsed.size()){
                    sb.append(" ");
                }
            }
        }
        sb.append(">");
    }
    ...

    public void putChildrenTagInto(StringBuffer sb){ ... }
    public void putEndTagInto(StringBuffer sb){ ... }
    }
} 
```

##### 절차 5

마지막으로 자식 컨테이너 클래스에 대한 클라이언트 코드가 `CompositeTag` 타입만을 사용하도록 하는지 확인해야 하는데, 파서 자체에는 그럴 일이 없으므로 이 리팩토링은 여기서 끝낸다.  

