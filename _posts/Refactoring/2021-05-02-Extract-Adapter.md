---
layout: post
title: Extract Adapter (347)
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

#### 개요

>하나의 클래스가 컴포넌트, 라이브러리, API 등의 **여러 버전**을 **동시에 지원**하기 위한 어댑터 역할을 하고 있다면
각 버전을 위한 기능을 **별도의 어댑터**로 뽑아낸다.  

<br>

#### 동기

소프트웨어를 개발하다 보면 컴포넌트, 라이브러리, API를 동시에 여러 버전으로 지원해야 할 때가 있지만 이런 버전 처리 코드가 굳이 복잡해질 필요는 없다. 그러나 **특정 버전만을 위한 상태 변수, 메서드**를 한 클래스에 오버로딩해 구현한 경우를 자주 볼 수 있다. 이 경우 주석으로 `# 버전 Y로 이동하면 반드시 삭제할것!` 이라고 적혀있지만, 대부분의 프로그래머들이 그 코드를 못 볼 가능성도 있다.  
이제, 이 리팩토링을 사용해 각 버전을 지원하는 별도 클래스를 만든다고 생각해보자. 클래스 이름에 버전을 명시하는 것도 좋다. 이런 클래스를 [Adapter](https://ukcastle.github.io/designpattern/2021/05/01/Adapter/)라고 부른다. 어댑터는 **공통 인터페이스**를 구현하고 특정 버전의 코드에 대해 정확히 동작해야 한다. 어댑터를 사용하면 클라이언트 코드에서 버전을 변경하기가 매우 쉬워진다.  

##### 장점

- 컴포넌트, 라이브러리, API의 버전에 따른 차이점을 격리한다.
- 클래스가 하나의 버전만 책임지도록 한다.
- 자주 변하는 코드를 시스템과 분리할 수 있다.

##### 단점

- 원래 있던 주요 기능을 어댑터에서 제공하지 못하면, 클라이언트가 그런 주요 기능에 접근하는 데 장벽이 될 수 있다.
    > 이럴 경우 어댑터를 재설계해야한다.  

<br>

#### 절차

이 리팩토링의 절차는 상황마다 다르다. 어떤 외부 코드의 여러 버전을 지원하기 위해 조건 로직을 사용하고 있다면 [Replace Conditional with Polymorphism](https://ukcastle.github.io/refactoring/2021/04/07/RefactoringToPattern/#replace-conditional-with-polymorphism)을 적용해 각 버전을 위한 어댑터를 만들 수 있다. 만약 어댑터 하나가 라이브러리의 **여러 버전을 지원하기 위해 버전 종속적인 변수나 메서드를 여러 개 포함**하고 있다면, 다른 방법을 사용해 어댑터를 여러 개 뽑아놔야 한다. 그 방법은 다음과 같다.  

1. 여러 버전의 코드를 어댑팅하기 위해 **과중한 책임을 떠맡고 있는 어댑터 클래스**를 찾는다.  

2. 과중한 책임을 맡고 있는 어댑터 클래스에 [Extract Subclass](https://ukcastle.github.io/refactoring/2021/04/07/RefactoringToPattern/#extract-subclass) 또는 [Extract Class](https://ukcastle.github.io/refactoring/2021/04/07/RefactoringToPattern/#extract-class)를 적용해 특정 버전에 종속적인 부분을 각각 별도의 클래스로 뽑아낸다. 특정 버전을 지원하기 위해 **배타적으로 사용되는 인스턴스 변수와 메서드를 새로 만든 어댑터로 모두 복사하거나 옮긴다.**  
이 과정에서 기존 어댑터 클래스의 `private` 필드나 메서드 중 일부를 `protected`로 수정해야할 수도 있다.  

3. 기존의 어댑터 클래스에 버전 종속적인 코드가 모두 사라질 때 까지 단계 2를 반복한다.  

4. 새로 만든 어댑터 클래스들 사이에 존재하는 **중복 코드**는 [Pull Up Method](https://ukcastle.github.io/refactoring/2021/04/07/RefactoringToPattern/#pull-up-method) 또는 [Form Template Method](https://ukcastle.github.io/refactoring/2021/04/16/Form-Template-Method/)를 사용해 제거한다.  

<br>

#### 구현

이번 코드는 써드파티 라이브러리를 이용해 DB 쿼리를 처리하는 코드이다.

```java
public Class Query{
    private SDLogin sdLogin; // SD 5.1
    private SDSession sdSession; // SD 5.1
    private SDLoginSession sdLoginSession; //SD5.2
    private boolean sd52; //52로 동작하고 있음을 나타내는 플래그
    private SDQuery sdQuery; //1,2 둘다

    //SD5.1을 위한 로그인 메서드
    //Warning!!! 모든 애플리케이션이 5.2로 전환되면 이 코드를 삭제할것!!
    public void login(String server, String user, String pw) throws QueryException{
        this.sd52 = false;
        // 5.1방식 DB 세션 로그인
    }

    //SD5.2를 위한 메서드
    public void login(String server, String user, String pw, String configName) throws QueryException{
        this.sd52 = true;
        // 5.2방식 DB 세션 로그인
    }

    public void doQuery() throws QueryException{
        if(this.sd52){
            this.sdQuery = this.sdLoginSession.createQuery(~~~);
        } else {
            this.sdQuery = this.sdSession.createQuery(~~~);
        }
        executeQuery();
    }
}
```

##### 절차 1
과중한 임무를 맡고 있는 클래스를 찾았다. 이 것을 어떻게 처리할지 고민해보자.  

##### 절차 2

`Query` 클래스는 아직 서브클래스가 없으므로 Extract Subclass로 **SD5.1**을 위한 코드를 분리하기로 하자.

```java
class QuerySD51 extends Query{
    public QuerySD51(){
        super();
    }
}
```

다음 적절하게 기존 코드를 바꿔준다.  

```java
public Class Query{
    ...

    public void doQuery() throws QueryException{
        if(this.sd52){
            this.sdQuery = this.sdLoginSession.createQuery(~~~);
        } else {
            QuerySD51 query = new QuerySD51();
            query.sdQuery = query.sdSession.createQuery(~~~);
        }
        executeQuery();
    }
}
```

다음 Push Down Method, Field를 적용하여 서브클래스를 바꿔준다.  

```java
class QuerySD51 extends Query{
    private SDLogin sdLogin;
    private SDSession sdSession;

    public QuerySD51(){
        super();
    }

    public login(~~~){
        //SD51 버전 로그인
    }

    public void doQuery() throws QueryException{
        this.sdQuery = this.sdSession.createQuery(~~~);
        executeQuery();
    }
}

public class Query{
    private SDLoginSession sdLoginSession; //SD5.2
    private boolean sd52; //52로 동작하고 있음을 나타내는 플래그
    private SDQuery sdQuery; //1,2 둘다

    public void login(String server, String user, String pw) throws QueryException{ 
        //pass    
    }

    //SD5.2를 위한 메서드
    public void login(String server, String user, String pw, String configName) throws QueryException{
        this.sd52 = true;
        // 5.2방식 DB 세션 로그인
    }

    public void doQuery() throws QueryException{
        
        this.sdQuery = this.sdLoginSession.createQuery(~~~);
        executeQuery();
    }
}
```

정상적으로 작동함을 알 수 있다.  

##### 절차 3

이제 `QuerySD52` 클래스를 만들고, Query를 추상클래스로 만들어줄 차례다.  

```java
public abstract class Query{
    private SDQuery sdQuery; //1,2 둘다


    public void login(String server, String user, String pw) throws QueryException;

    public void login(String server, String user, String pw, String configName) throws QueryException;

    public void doQuery() throws QueryException;
}

class QuerySD51 extends Query{
    private SDLogin sdLogin;
    private SDSession sdSession;

    public QuerySD51(){
        super();
    }

    public login(String server, String user, String pw){
        //SD51 버전 로그인
    }

    public void doQuery() throws QueryException{
        this.sdQuery = this.sdSession.createQuery(~~~);
        executeQuery();
    }
}

class QuerySD52 extends Query{
    private SDLoginSession sdLoginSession;

    public QuerySD52(){
        super();
    }
    
    @Override
    public login(String server, String user, String pw, String configName){
        //SD52 버전 로그인
    }

    @Override
    public void doQuery() throws QueryException{ 
        this.sdQuery = this.sdLoginSession.createQuery(~~~);
        executeQuery();
    }
}

```

만들고 보니, doQuery()의 `executeQuery()`라는 중복되는 함수가 있다. 예제라 한 줄 이지만, 이게 과연 100줄이라면? 당연히 리팩토링을 해야된다. [Introduce Polymorphic Creation with Factory Method](https://ukcastle.github.io/refactoring/2021/04/13/Introduce-Polymorphic-Creation-with-Factory-Method/)와 [Form Template Method](https://ukcastle.github.io/refactoring/2021/04/16/Form-Template-Method/)를 통해 다음과 같이 `doQuery()`를 수퍼클래스로 옮기자.  

```java
public abstract class Query{
    protected abstract SDQuery createQuery(); //팩터리 메서드

    public void login(String server, String user, String pw) throws QueryException;

    public void login(String server, String user, String pw, String configName) throws QueryException;

    public void doQuery() throws QueryException{
        SDQuery sdQuery = createQuery();
        executeQuery();
    }        
}

class QuerySD51 extends Query{
    private SDLogin sdLogin;
    private SDSession sdSession;

    @Override
    protected SDQuery createQuery(){
        return sdSession.createQuery(~~);
    }

    @Override
    public login(String server, String user, String pw){
        //SD51 버전 로그인
    }
}

class QuerySD52 extends Query{
    private SDLoginSession sdLoginSession;
    
    @Override
    protected SDQuery createQuery(){
        return sdLoginSession.createQuery(~~);
    }
    
    @Override
    public login(String server, String user, String pw, String configName){
        //SD52 버전 로그인
    }
}
```

다음 이상한 점은 모두 알 것이다.  
추상 메소드 `login()` 이 두 가지로 나뉘어져있다.  
SD52버전에선 `configName`이라는 변수를 하나 더받는다. 이를 생성자로 바꿔주면 된다.  

```java
public abstract class Query{
    protected abstract SDQuery createQuery(); //팩터리 메서드

    public abstract void login(String server, String user, String pw) throws QueryException;

    public void doQuery() throws QueryException{
        SDQuery sdQuery = createQuery();
        executeQuery();
    }        
}

class QuerySD51 extends Query{
    private SDLogin sdLogin;
    private SDSession sdSession;

    @Override
    protected SDQuery createQuery(){
        return sdSession.createQuery(~~);
    }

    @Override
    public login(String server, String user, String pw){
        //SD51 버전 로그인
    }
}

class QuerySD52 extends Query{
    private SDLoginSession sdLoginSession;
    String configName;

    public QuerySD52(String configName){
        super()
        this.configName = configName;
    }

    @Override
    protected SDQuery createQuery(){
        return sdLoginSession.createQuery(~~);
    }
    
    @Override
    public login(String server, String user, String pw){
        //SD52 버전 로그인
    }
}
```

이제 막바지에 이르렀다. `Query`는 추상클래스가 되었으므로 이름을 `AbstractQuery`로 바꾸어 그 특성을 명확하게 드러나게 하는 것이 좋다. 그런데 이름을 바꾸면 이미 코딩된 클라이언트 코드에서는 Query 타입 변수를 선언한 곳을 모두 찾아 이름을 바꿔주어야 한다. 그렇게 하고 싶지 않으므로, `AbstractQuery`에 Extract Interface를 적용시켜 Query 인터페이스를 만들고 이를 구현하게 한다.  

```java
interface Query{
    public abstract void login(String server, String user, String pw) throws QueryException;
    public void doQuery() throws QueryException;
}

public abstract class AbstractQuery implements Query{
    protected abstract SDQuery createQuery(); //팩터리 메서드

    public void doQuery() throws QueryException{
        SDQuery sdQuery = createQuery();
        executeQuery();
    }        
}

class QuerySD51 extends AbstractQuery{
    private SDLogin sdLogin;
    private SDSession sdSession;

    @Override
    protected SDQuery createQuery(){
        return sdSession.createQuery(~~);
    }

    @Override
    public login(String server, String user, String pw){
        //SD51 버전 로그인
    }
}

class QuerySD52 extends AbstractQuery{
    private SDLoginSession sdLoginSession;
    String configName;

    public QuerySD52(String configName){
        super()
        this.configName = configName;
    }

    @Override
    protected SDQuery createQuery(){
        return sdLoginSession.createQuery(~~);
    }
    
    @Override
    public login(String server, String user, String pw){
        //SD52 버전 로그인
    }
}
```

##### 완성

이렇게 기존 코드와 동일한 기능을 하며 완전히 어댑팅되었다. 코드가 원래보다 간단해졌고 두 버전을 동일한 방식으로 다룰 수 있게 되었다. 더 나아가 다음과 같은 이득을 얻게 되었다.  

- 각 버전 간의 유사점과 차이점을 쉽게 알아볼 수 있다.  
- 오래되어 사용하지 않는 버전을 위한 코드를 쉽게 제거할 수 있게 되었다.
- 새 버전을 지원하는 일이 쉬워졌다.  


#### 변형 - 익명 내부 클래스를 사용해 어댑팅하기

Java에는 기존 `Enumeration`이라는 인터페이스가 있었다. 그런데 점점 발전하며 `Iterator` 인터페이스가 추가되며, 역할을 대신하게 되었다. 그러나 기존 코드와도 상호 동작이 가능해야 하므로, JDK에는 다음과 같은 익명 클래스 기능을 이용해 `Iterator`을 어댑팅하는 생성 메서드를 제공한다.  

```java
public class Collections...{
    public static Enumeration enumeration(final Collection c){
        return new Enumeration(){
            Iterator i = c.iterator();

            public boolean hasMoreElements(){
                return i.hasNext();
            }
            public Object nextElement(){
                return i.next()
            }
        };
    }
    ...
}
```
익명 클래스에 `Iterator` 변수를 넣어 반환하는 방식이다.   

