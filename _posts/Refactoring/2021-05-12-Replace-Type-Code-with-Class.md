---
layout: post
title: Replace Type Code with Class (383)
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

#### 개요

>어떤 필드 타입(예를 들어, String 또는 int 등)이 부적합한 값의 대입이나 유효하지 않은 동일성 검사(비교)를 방지하지 못한다면,
필드의 **타입을 클래스로 바꿔** 값의 **대입과 동일성 검사에 제약 조건을 부여**한다

<br>

#### 동기

타입 코드로 리팩터링하는 주된 이유는 코드의 **타입 안정성**을 보장하려는 것이다.  
예를 들어, Config파일에서 상태값을 `String`으로 저장하는데, "GRANTED"라는 문자열로 저장한다고 하자.  
이 떄 클라이언트 코드에서 실수로, "GRNATED"라고 오타가 났다고 하면, IDE는 오류를 잡아주지 못하고, 원하는데로 돌아가지 않아 에러가 난 부분을 찾는데 고생을 할 것이다. 단지 한 글자 차이로 말이다.  
이런 문제를 처리하기 위해, 보통 `enum 열거자`를 지정한다. 하지만 이럴 경우 더 많은 기능을 추가하는 것이 불가능하다.  

##### 장점

- 부적절한 값의 대입이나 유효하지 않은 동일성 검사로부터 코드를 보호한다.  

##### 단점

- 타입 안정성이 결여된 경우보다 더 많은 코드가 필요하다.

<br>

#### 절차

타입 안정성이 없는 상수는 int나 String같은 기본 타입 또는 String으로 정의된 상수를 의미한다.  

1. **타입 안정성이 없는 필드를 확인**한다. 즉, 타입 안정성이 없는 상수를 대입하거나 그 상수와 동일성 검사를 하는 필드를 찾는 것이다. 찾은 필드에 Self Encapsulate Field(필드 자체를 클래스로 만들어 get/set으로 접근) 리팩토링을 적용해 자체 캡슐화 한다.  <br>

2. 새로운 클래스를 하나 만든다. 이 클래스는 나중에 앞에서 찾은 필드의 타입을 대체할 것이다. **클래스의 이름은 관련 상수의 의미를 참고하여 짓는다.** 당장은 생성자를 별도로 선언하지 않는다.  <br>

3. 타입 안정성이 없는 필드에 대입되거나 이와 비교되는 **상수를 하나 선택**해, 이에 대응하는 새로운 상수를 **앞에서 만든 새 클래스를 선언**하는데, 새로운 상수는 이 클래스의 **인스턴스**가 되어야 한다. `final static`과 같은 선언을 해주는것이 일반적이다.  <br>

4. 타입 안정성이 없는 필드가 선언된 클래스에 **앞에서 만든 새 클래스 타입의 필드를 선언**한다(이 필드는 타입 안정성이 보장된다). 그리고 그에 대한 `setter`를 구현한다.  <br>

5. **타입 안정성이 없는 필드에 값을 대입하는 코드를 모두 찾아**, **타입 안정성이 보장된 필드에 대한 적절한 대입문을 추가**한다. 이 때 대입값은 새 클래스에 정의한 상수 중 하나다.

6. 타입 안정성이 없는 필드에 대한 **`getter` 메서드를 수정**해, **타입 안정성이 보장된 필드로부터 얻은 값을 리턴**하도록 한다. 물론, 새 클래스도 올바른 상수값을 리턴할 수 있도록 수정해야 한다.  

7. 타입 안정성이 없는 필드와 그에 대한 **setter 메서드, 그리고 그 메서드를 호출하던 코드를 모두 제거**한다.  

8. 타입 안정성이 없는 상수를 참조하던 코드를 모두 찾아 새 클레스에 있는 상수 가운데 **그에 대응하는 것으로 치환**한다. 이 때 **타입안정성이 없는 필드에 대한 getter의 리턴타입을 새 클래스로 변경**하고, 그 getter를 호출하는 모든 코드도 적절히 수정한다.  

결과적으로 기본 타입을 사용하던 동일한 검사 로직이 **새 클래스의 인스턴스를 비교하는 방식**으로 바뀐다. 프로그래밍 언어가 객체 동일성 검사 로직을 기본적으로 제공할 수 있다. 그렇지 않으면 새 클래스의 객체 동일성 검사가 제대로 이뤄질 수 있도록 코드를 추가해야 한다.  

<br>

#### 구현

```java

public class SystemPermission{

    private String state;
    private boolean granted;

    public final static String REQUESTED = "REQUESTED";
    public final static String CLAMED = "CLAMED";
    public final static String DENIED = "DENIED";
    public final static String GRANTED = "GRANTED";

    public StstemPermission(){
        this.state = REQUESTED;
        this.granted = false;
    }

    public void clamed(){
        if(state.equals(REQUESTED)){
            this.state = CLAMED;
        }
    }
    public void denied(){
        if(state.equals(REQUESTED)){
            this.state = DENIED;
        }
    }
    public void granted(){
        if(!state.equals(CLAMED)){
            return;
        }
        this.state = GRNATED;
        this.granted = true;
    }
    public boolean isGranted(){
        return this.greanted;
    }
    public String getState(){
        return this.state;
    }
} 
```

##### 절차 1

해당 클래스에는 `state`라는 타입 안정성이 없는 필드가 있다. 이 필드에는 `String`타입의 상수가 대입된다. 따라서 이 필드의 타입을 `String`이 아닌 다른 클래스로 바꾸어 타입 안정성을 확보하는 것이 리팩터링의 목표이다.  
일단 `state`필드를 자체 캡슐화 한다.  

```java

public class SystemPermission{

    private String state;
    private boolean granted;

    public final static String REQUESTED = "REQUESTED";
    public final static String CLAMED = "CLAMED";
    public final static String DENIED = "DENIED";
    public final static String GRANTED = "GRANTED";

    public StstemPermission(){
        setState(REQUESTED);
        this.granted = false;
    }

    public void clamed(){
        if(getState().equals(REQUESTED)){
            setState(CLAMED);
        }
    }
    public void denied(){
        if(state.equals(REQUESTED)){
            this.state = DENIED;
        }
    }
    ...

    private void setState(String state) { this.state = state; }
    public String getState(){ return this.state; }
} 
```

##### 절차 2

`PermissionState`라는 새 클래스를 만든다. 이 클래스가 SystemPermission 객체의 상태를 표현하게 될 것이다.

```java
public class PermissionState{}
```

##### 절차 3

`state`필드에 대입되거나 또는 비교되는 상수를 하나 골라 그에 대응하는 상수를 `PermissionState`에 정의한다. 이 때 새로 만드는 상수는 `PermissionState` 타입으로 한다.  

```java
public class PermissionState{
    public final static PermissionState REQUESTED = new PermissionState();
    public final static PermissionState CLAMED = new PermissionState();
    public final static PermissionState GRANTED = new PermissionState();
    public final static PermissionState DENIED = new PermissionState();
}
```

다른 상수에 대해서도 같은 작업을 반복한다.

이 때 클래스의 인스턴스가 위의 4개 이상 존재할 수 없게 제한하여 더 엄격한 수준의 타입 안정성을 챙길 수 있다. 이번 예제에서는 하지 않겠다.  


##### 절차 4

`SystemPermission`에 `PermissionState` 타입의 필드를 만든다. 이 필드는 안정성이 보장된다. 그리고 그에 대한 set 메서드를 구현한다.

```java

public class SystemPermission{

    private String state;
    private PermissionState permission;

    private void setState(PermissionState permission){
        this.permission = permissionl
    }
} 
```

##### 절차 5

타입 안정성이 없는 `state`필드에 값을 대입하는 곳을 찾아 그에 대응하여 `permission`필드에 대한 대입문을 적절히 추가한다.  


```java

public class SystemPermission{

    private String state;
    private PermissionState permission;
    private boolean granted;

    private void setState(PermissionState permission){
        this.permission = permissionl
    }

    public StstemPermission(){
        // setState(REQUEST);
        setState(PermissionState.REQUEST);
        this.granted = false;
    }

    public void clamed(){
        if(state.equals(REQUESTED)){
            // this.state = CLAMED;
            setState(PermissionState.CLAEMED);
        }
    }
    public void denied(){
        if(state.equals(REQUESTED)){
            setState(PermissionState.DENIED);
        }
    }
    public void granted(){
        if(!state.equals(CLAMED)){
            return;
        }
        // this.state = GRNATED;
        setState(PermissionState.GRANTED);
        this.granted = true;
    }
    public boolean isGranted(){
        return this.greanted;
    }
    public String getState(){
        return this.state;
    }
} 
```

##### 절차 6

`state` 필드에 대한 `getter` 메서드가 타입 안정성이 보장된 `permission`필드 값을 리턴하도록 수정할 차례다. `state`에 대한 `getter`가 `String`을 리턴하므로 `permission`또한 `String`을 리턴할 수 있도록 해야할 것이다. 그 첫 단계는 각 상수의 이름을 리턴하는 `toString()` 메서드를 `PermissionState`에 추가하는 것이다.  

```java
public class PermissionState{

    private final String name;

    private PermissionState(String name){
        this.name = name;
    }

    public String toString(){
        return name;
    }

    public final static PermissionState REQUESTED = new PermissionState("REQUESTED");
    public final static PermissionState CLAMED = new PermissionState("CLAMED");
    public final static PermissionState GRANTED = new PermissionState("GRANTED");
    public final static PermissionState DENIED = new PermissionState("DENIED");
}
```

다음 `state`필드에 대한 `getter`를 수정한다.

```java

public class SystemPermission{

    ...
    public String getState(){
        return this.permission.toString();
    }
} 
```

##### 절차 7

다음 `SystemPermission`에서 타입 안정성이 없는 `state`필드를 제거한다. 그에 대한 메서드들도 제거한다.  

```java

public class SystemPermission{

    private PermissionState permission;
    private boolean granted;

    public StstemPermission(){
        setState(PermissionState.REQUEST);
        this.granted = false;
    }

    public void clamed(){
        if(getState().equals("REQUESTED")){
            setState(PermissionState.CLAEMED);
        }
    }
    public void denied(){
        if(getState().equals("REQUESTED")){
            setState(PermissionState.DENIED);
        }
    }
    public void granted(){
        if(!getState().equals("CLAMED")){
            return;
        }
        setState(PermissionState.GRANTED);
        this.granted = true;
    }

    public String getState(){
        return this.permission.toString();
    }
} 
```

##### 절차 8

다음 `SystemPermission`에 정의된 타입 안정성이 없는 상수를 참조하는 모든 코드를 `PermissionState`에 정의된 상수를 참조하도록 고친다. 예를 들어, `clamed()`에는 타입 안정성이 없는 `REQUESTED` 상수를 참조한다.

최종 코드다.

```java
public class PermissionState{

    private final String name;

    private PermissionState(String name){
        this.name = name;
    }

    public String toString(){
        return name;
    }

    public final static PermissionState REQUESTED = new PermissionState("REQUESTED");
    public final static PermissionState CLAMED = new PermissionState("CLAMED");
    public final static PermissionState GRANTED = new PermissionState("GRANTED");
    public final static PermissionState DENIED = new PermissionState("DENIED");
}

public class SystemPermission{

    private PermissionState permission;
    private boolean granted;

    public StstemPermission(){
        setState(PermissionState.REQUEST);
        this.granted = false;
    }

    public void clamed(){
        if(getState().equals(PermissionState.REQUESTED)){
            setState(PermissionState.CLAEMED);
        }
    }
    public void denied(){
        if(getState().equals(PermissionState.REQUESTED)){
            setState(PermissionState.DENIED);
        }
    }
    public void granted(){
        if(!getState().equals(PermissionState.CLAMED)){
            return;
        }
        setState(PermissionState.GRANTED);
        this.granted = true;
    }

    public String getState(){
        return this.permission;
    }
} 
```