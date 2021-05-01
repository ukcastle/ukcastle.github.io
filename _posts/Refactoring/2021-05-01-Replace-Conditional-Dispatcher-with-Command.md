---
layout: post
title: Replace Conditional Dispatcher with Command (265)
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

#### 개요

>요청에 대한 디스패처가 조건 로직으로 구현되어 있다면
각 액션에 대한 커맨드 객체를 만들어 컬렉션에 저장해두고, 조건 로직은 컬렉션에서 원하는 커맨드 객체를 찾아 실행하는 코드로 대체한다.

<br>

#### 동기

많은 코드에서, 어떠한 상황에 따라 여러가지 로직으로 분류되는 (switch와 같은)조건문을 보았을것이다. 이를 **조건적 디스패처**라고 한다.  
처리해야 할 요청의 종류가 적고 이를 처리하는 로직도 얼마 되지 않는다면 기존의 형태대로 구현해도 무방하지만, 조건 로직 부분이 모니터 한 화면에서 보지 못 할 정도로 방대하다면 기존의 조건적 디스패처를 [Command](https://jo631.github.io/designpattern/2021/05/01/Command/) 패턴 으로 바꾸는 리팩토링이 권장된다.  
하지만 이는 해당 리팩토링을 하는 대표적인 이유는 아닌데, 대표적인 이유 두가지는 다음과 같다.  
1. 런타임에 충분히 유동적이지 못하다.  
    >요청이나 처리 로직이 동적으로 구성될 필요가 있는 경우, 조건적 디스패처는 적절하지 않다. 이는 처리 로직이 하드 코딩되기 때문에 로직의 동적 구성을 지원할 수 없다.  
2. 코드가 비대해진다.
    >새로운 종류의 요청을 처리하기 위한 로직이 추가되거나, 로직이 복잡해지면 코드는 무지막지하게 커질 수 밖에 없다. 별도의 메서드로 분리한다고 해도 클래스의 크기는 똑같아 큰 도움이 되지는 못 할 것이다.  

Command 패턴은 각 요청을 처리하는 로직을 `execute()` 또는 `run()`과 같은 공통 메서드를 가진 별도의 커맨드 클래스로 옮겨 캡슐화한다. 이렇게 커맨드의 집합을 만들고 나면 리스트를 이용해 그 명령들를 추가하고 삭제하고 변경하는 등 조작을 할 수 있다.  
요청을 분배하고 다양한 액션을 동일한 방싯으로 실행시킬 수 있께 하는 것은 설계에 있어 매우 일반적이기 떄문에, 나중에 리팩터링 하기 보단 **개발 초기부터** 이 패턴을 사용하는 경우가 많다.  

<br>

#### 장점

- 다양한 액션을 단일한 방식으로 실행하는 단순한 구조를 제공한다.  
- 요청을 처리하는 로직의 구성을 런타임에서 변경할 수 있다.  
- 간단한 코드로 구현할 수 있다.  

#### 단점

- 조건적 디스패처로도 충분한 상황에는 괜히 설계만 복잡하게 만드는 것 이다.  

<br>

#### 절차

1. 조건적 디스패처를 포함한 클래스에서 **요청을 실행하는 코드를 찾고**, [Extract Method](https://jo631.github.io/refactoring/2021/04/07/RefactoringToPattern/#extract-method)를 적용해 별도의 실행 메서드로 뽑아낸다.  

2. 요청을 실행하는 나머지 다른 코드에 대해서도 단계 1을 반복해 모두 별도의 실행 메서드로 바꾼다.  

3. 각각의 실행 메서드에 [Extract Class](https://jo631.github.io/refactoring/2021/04/07/RefactoringToPattern/#extract-class) 를 적용해 요청을 처리하는 구체 커맨드 클래스로 만든다. 이 과정에서 구체 커맨드 클래스로 옮긴 실행 메서드들은 보통 `public` 메서드가 될 것이다. 만약 옮긴 실행 메서드가 너무 크거나 쉽게 이해할 수 없다면 [Compose Method](https://jo631.github.io/refactoring/2021/04/14/Compose-Method/)를 적용하라. 구체 클래스를 모두 만들고 난 후 중복된 코드가 없는지 확인해보고, 만약 있다면 [Form Template Method](https://jo631.github.io/refactoring/2021/04/16/Form-Template-Method/)를 적용하라.   

4. 앞서 만든 모든 구체 커맨드 클래스에 공통으로 적용될 수 있는 메서드를 선언하는 **인터페이스**를 정의한다. 이 과정에서 커맨드 클래스들의 공통점과 차이점을 찾아야 한다. 다음 질문에 대해 답을 찾아보자.  

- 공통 실행 메서드에는 어떤 파라미터를 넘겨야 할까?
    >기존 코드 실행에 필요한 지역변수들
- 구체 커맨드 인스턴스를 만들 때에는 어떤 파라미터를 넘겨줄 수 있을까?
    >조건적 디스패처의 조건을 구성하는 변수 또는 상황
- 실행 메서드에 직접 넘기기보단 구체 커맨드 클래스에서 파라미터에 대한 콜백을 통해 얻도록 하는 것이 나은 정보에는 어떤것이 있을까?
    >추후 사용될 가능성이 있는 변수
- 모든 구체 커맨드 클래스에 동일하게 적용할 수 있는 실행 메서드의 가장 간단한 시그니처는 무엇인가?
    > execute() or run()

구체 커맨드 클래스에 대해 [Extract Superclass](https://jo631.github.io/refactoring/2021/04/07/RefactoringToPattern/#extract-superclass) 또는 [Extract Interface](https://jo631.github.io/refactoring/2021/04/07/RefactoringToPattern/#extract-interface)를 적용해 초기 버전의 커맨드를 만드는 것을 고려한다.  

5. 모든 구체 커맨드 클래스가 단계 4에서 만든 커맨드 타입을 **구현하거나 상속**하도록 수정한다.  

6. 조건적 디스패처가 있는 클래스에 커맨드 **맵(딕셔너리)**를 만든다. 즉, 각 구체 커맨드 클래스 인스턴스를 맵에 저장하는데, 클래스 이름 등의 유일한 식별자를 Key로 사용한다. 유일한 식별자는 런타임에 커맨드 객체를 찾는데 사용할 것이다.  

7. 조건적 디스패처가 있는 클래스에서 요청을 부냅하는 코드를 제거하고 커맨드 객체를 맵에서 찾아 그 실행 메서드를 호출하느 코드로 대체한다. 이제 이 클래스는 Invoker가 된다.  

<br>

#### 구현

워크샵의 카탈로그를 HTML로 생성하는 코드이다.

```java

public class CatalogApp{
    private Response executeAction(String actionName, Map param){

        if(actionName.equals(NEW_WORKSHOP)){
            String nextWorkshopID = workshopManager.getNextWorkshopID();
            StringBuffer newWorkshopContents = 
                workshopManager.createNewFileFromTemplate(
                    nextWorkshopID,
                    workshopManager.getWorkshopDir(),
                    workshopManager.getWorkshopTemplate()
                );
            workshopManager.addWorkshop(newWorkshopContents);
            param.put("id",nextWorkshopID);
            this.executeAction(ALL_WORKSHOPS, param);
        }
        else if(actionName.equals(ALL_WORKSHOPS)){
            //HTML을 만드는 코드들...
        }
        ...
    }
}
```

아주 많지만, 간소화했다. 첫 분기는 새로운 워크샵을 생성하는 것이고, 두 번째는 모든 워크샵에 대한 정보로 HTML을 만드는 것이다. 잘 보면 첫 분기에서 두번째 분기를 호출해 간소화하고있다.  

##### 절차 1

첫 분기부터 시작하자. Extract Metho를 적용해 `getNewWorkshopResponse()`라는 실행 메소드를 만든다.  

```java

public class CatalogApp{
    private Response executeAction(String actionName, Map param){

        if(actionName.equals(NEW_WORKSHOP)){
            this.getNewWorkshopResponse(param)
        }
        else if(actionName.equals(ALL_WORKSHOPS)){
            //HTML을 만드는 코드들...
        }
        ...
    }
    private void getNewWorkshopResponse(Map param){
        String nextWorkshopID = workshopManager.getNextWorkshopID();
        StringBuffer newWorkshopContents = 
                workshopManager.createNewFileFromTemplate(
                    nextWorkshopID,
                    workshopManager.getWorkshopDir(),
                    workshopManager.getWorkshopTemplate()
                );
        workshopManager.addWorkshop(newWorkshopContents);
        param.put("id",nextWorkshopID);
        this.executeAction(ALL_WORKSHOPS, param);
    }
}
```

##### 절차 2

나머지 분기들에서도 절차 1을 반복한다.  

```java

public class CatalogApp{
    private Response executeAction(String actionName, Map param){

        if(actionName.equals(NEW_WORKSHOP)){
            this.getNewWorkshopResponse(param);
        }
        else if(actionName.equals(ALL_WORKSHOPS)){
            this.getAllWorkshopResponse();
        }
        ...
    }
    private void getNewWorkshopResponse(Map param){
        String nextWorkshopID = workshopManager.getNextWorkshopID();
        StringBuffer newWorkshopContents = 
                workshopManager.createNewFileFromTemplate(
                    nextWorkshopID,
                    workshopManager.getWorkshopDir(),
                    workshopManager.getWorkshopTemplate()
                );
        workshopManager.addWorkshop(newWorkshopContents);
        param.put("id",nextWorkshopID);
        this.executeAction(ALL_WORKSHOPS, param);
    }

    private Response getAllWorkshopResponse(){
        //HTML을 만드는 코드들..
    }
}
```

##### 절차 3

구체 커맨드 클래스를 만들기 시작한다. 

```java

public class NewWorkshopHandler{
    CatalogApp ctApp;
    
    public NewWorkshopHandler(CatalogApp ctApp){
        this.ctApp = ctApp;
    }

    public Response getNewWorkshopResponse(Map param){
        ...
        this.ctApp.executeAction(ALL_WORKSHOPS, param);
    }
}

public class AllWorkshopHandler{
    CatalogApp ctApp;
    
    public AllWorkshopHandler(CatalogApp ctApp){
        this.ctApp = ctApp;
    }

    public Response getAllWorkshopResponse(){
        ...
    }
}

public class CatalogApp{
    Response executeAction(String actionName, Map param){

        if(actionName.equals(NEW_WORKSHOP)){
            return new NewWorkshopHandler(this).getNewWorkshopResponse(param);
        }
        else if(actionName.equals(ALL_WORKSHOPS)){
            return new ALlWorkshopHandler(this).getAllWorkshopResponse()
        }
        ...
    }
}
```


##### 절차 4

타음 커맨드들의 수퍼 타입을 정의할 차례이다. 현재 상태를 보면, 구체 커맨드 클래스의 실행 메서드는 **이름도 다르고**, **파라미터의 개수**도 다르다.  
커맨드 수퍼 타입을 만드려면, 다음 사항을 결정해야 한다.  
- 공통 실행 메서드의 이름  
    >`execute()`를 사용하기로 결정했다.  
- 실행 메서드로 넘겨야 할 정보와 실행 메서드로부터 받을 정보
    >`Response`를 반환하고, `param`을 받는다

이제 추상클래스를 만들어보자.  

```java

public abstract class Handler{
    protected CatalogApp catalogApp;

    public Handler(CatalogApp catalogApp){
        this.catalogApp = catalogApp;
    }
    Response execute(Map param);
}

public class NewWorkshopHandler extends Handler{
    public NewWorkshopHandler(CatalogApp ctApp){
        super(ctApp);
    }

    @Override
    public Response execute(Map param){
        ...
        this.ctApp.executeAction(ALL_WORKSHOPS, param);
    }
}

public class AllWorkshopHandler{
    public AllWorkshopHandler(CatalogApp ctApp){
        super(ctApp);
    }

    @Override
    public Response execute(Map param){
        ...
    }
}
```

##### 절차 5

수퍼 타입을 만들었으므로, 클라이언트에서도 코드를 바꿔준다.

```java
public class CatalogApp{
    Response executeAction(String actionName, Map param){

        if(actionName.equals(NEW_WORKSHOP)){
            return new NewWorkshopHandler(this).execute(param);
        }
        else if(actionName.equals(ALL_WORKSHOPS)){
            return new ALlWorkshopHandler(this).execute(param)
        }
        ...
    }
}
```

##### 절차 6

이제 흥미로운 부분이다. `CatalogApp` 클래스의 조건 로직은 단지 매핑의 역할만 하고 있다. 이 것을 커맨드 인스턴스로 저장하는 진짜 맵으로 대체하자. handler이라는 `Map` 객체를 만들고 액션 이름을 키로 해서 커맨드 객체를 Map에 넣는다. 

```java

public class CatalogApp{
    private Map handlers;

    public CatalogApp(...){
        ...
        createHandlers();
        ...
    }

    public void createHandlers(){
        handlers = new HaspMap();
        handlers.put(NEW_WORKSHOP), new NewWorkshopHandler(this));
        handlers.put(ALL_WORKSHOPS), new AllWorkshopHandler(this));
    }

    Response executeAction(String actionName, Map param){

        if(actionName.equals(NEW_WORKSHOP)){
            return new NewWorkshopHandler(this).execute(param);
        }
        else if(actionName.equals(ALL_WORKSHOPS)){
            return new ALlWorkshopHandler(this).execute(param)
        }
        ...
    }
}
```

##### 절차 7

마지막 단계다. 클라이언트 코드에서 조건문을 제거하자.  

```java

    Response executeAction(String handlerName, Map param){
        return lookupHandlerBy(handlerName).execute(param);
    }

    private Handler lookupHandlerBy(String handlerName){
        return (Handler)handlers.get(handlerName)
    }
```

이제 새로운 커맨드를 추가하려면, 새로운 커맨드 클래스르 만들고 커맨드 맵에 그를 등록하기만 하면, 런타임에서 알아서 실행 될 것 이다.  

최종 코드다.  

```java
public class CatalogApp{
    private Map handlers;

    public CatalogApp(...){
        ...
        createHandlers();
        ...
    }

    public void createHandlers(){
        handlers = new HaspMap();
        handlers.put(NEW_WORKSHOP), new NewWorkshopHandler(this));
        handlers.put(ALL_WORKSHOPS), new AllWorkshopHandler(this));
    }

    Response executeAction(String handlerName, Map param){
        return lookupHandlerBy(handlerName).execute(param);
    }

    private Handler lookupHandlerBy(String handlerName){
        return (Handler)handlers.get(handlerName)
    }
}

public abstract class Handler{
    protected CatalogApp catalogApp;

    public Handler(CatalogApp catalogApp){
        this.catalogApp = catalogApp;
    }
    Response execute(Map param);
}

public class NewWorkshopHandler extends Handler{
    public NewWorkshopHandler(CatalogApp ctApp){
        super(ctApp);
    }

    @Override
    public Response execute(Map param){
        ...
        this.ctApp.executeAction(ALL_WORKSHOPS, param);
    }
}

public class AllWorkshopHandler{
    public AllWorkshopHandler(CatalogApp ctApp){
        super(ctApp);
    }

    @Override
    public Response execute(Map param){
        ...
    }
}
```