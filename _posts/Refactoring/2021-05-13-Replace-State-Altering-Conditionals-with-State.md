---
layout: post
title: Replace State-Altering Conditionals with State (234)
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

#### 개요

>어떤 객체의 상태 전이를 제어하는 조건 로직이 **복잡**하다면,
각 상태에 해당하는 [State](https://ukcastle.github.io/designpattern/2021/05/13/State/) 클래스를 하나씩 만들고 그들이 스스로 다른 상태로 전이하는 것을 책임지도록 하여 복잡한 조건 로직을 제거한다.  

<br>

#### 동기

State 패턴으로 리팩터링하는 주된 목적은 **상태 전이를 위한 조건 로직이 지나치게 복잡한 경우** 이를 해소하는 것이다. 상태 전이 로직이란 객체의 상태와 이들 간의 전이 방법을 제어하는 것으로, 클래스 내부 여기저기에 흩어져 존재하는 경향이 있다. State 패턴을 구현한다는 것은 각 상태에 대응하는 별도의 클래스를 만들고 상태 전이 로직을 그 클래스들로 옮기는 작업을 뜻한다. 이 때 원래의 호스트인 `Context` 객체는 상태와 관련된 기능을 State 객체에 위임한다. 그리고 **상태 전이**는 Context 객체의 대리 객체를 한 State 객체에서 **다른 State 객체로 바꾸는 일**이 된다.  
이 리팩터링을 시작하기 전에, [Extract Method](https://ukcastle.github.io/refactoring/2021/04/07/RefactoringToPattern/#extract-method)와 같은 단순한 리팩터링으로도 충분한지 보는것이 좋다. 이는 복잡한 코드에 적합한 리팩터링이기 때문이다.  

만약 State 객체에 필드가 없다면, Context 객체가 State 객체를 공유하게 만들어서 메모리를 절약할 수 있다. 공유를 위해 자주 사용되는 패턴에는 Flyweight와 Singleton이 있다. 그러나 너무 앞서 State 객체의 공유를 구현하기 보단, 추후 State의 생성 코드가 주요 병목지점임을 확인한 후 적용해도 늦지 않는다.  

##### 장점

- 상태 전이를 위한 조건 로직을 줄이거나 제거할 수 있다.
- 복잡한 상태 전이 로직이 단순해진다. 
- 상태 전이 로직을 더 쉽게 알아볼 수 있다.

##### 단점

- 원래의 상태 전이 로직이 **별로 복잡하지 않았다면**, 괜히 설계만 복잡하게 만드는 것이다.  

<br>

#### 절차

1. 컨텍스트 클래스는 원래의 상태 필드를 갖고 있는 클래스다. 상태 필드에는 상태 전이가 일어나는 동안에 상태를 나타내는 상수 가운데 하나가 대입되고, 그 값이 비교되기도 한다. [Replace Type Code with Class](https://ukcastle.github.io/refactoring/2021/05/12/Replace-Type-Code-with-Class/)를 적용해준다.  

2. 이제 State 수퍼 클래스에서 정의된 각 상수는 State 인스턴스를 하나씩 참조하고 있다. [Extract Subclass](https://ukcastle.github.io/refactoring/2021/04/07/RefactoringToPattern/#extract-subclass)를 통해 **각 상수에 대한 서브클래스를 하나씩 만든 후, 수퍼 클래스에서는 그에 대한 서브 클래스 인스턴스들을 참조**하도록 수정한다. 그리고 마무리 작업으로 **수퍼 클래스를 추상화**시킨다.  

3. 컨텍스트 클래스에서 **상태 전이 로직에 따라 원래의 상태 필드의 값을 변경하는 작업을 수행하는 메서드**를 찾는다. 그리고 이 메서드를 State 수퍼 클래스로 복사하는데, 단순히 복사만 해서는 코드가 동작하지 않을 수 있다. **가장 간단한 해결법은 Context 객체를 새 메서드에 파라미터로 넘겨** 해결한다. 마지막으로, 컨텍스트 클래스의 원본 메서드는 작업을 새로 만든 메서드에게 위임하도록 수정한다.  

4. Context 클래스가 가질 수 있는 특정 상태를 하나 선택한 다음, State 수퍼 클래스의 메서드 중 **선택한 상태에서 다른 상태로 상태를 전이하는 코드가 있는지 확인**한다. 만약 있다면, 해당 메서드를 그 상태에 대응하는 서브클래스로 복사한 다음 상태 전이와 관련 없는 코드는 제거한다. **현재의 상태를 확인**하는 코드나 **현재와 관련이 없는 상태로 전이**하는 로직은 특정 상태를 나타내는 **서브 클래스에서 의미 없는 로직**이다.    
모든 서브클래스에 대해 이 작업을 반복한다.  

5. 앞의 단계 3에서 State 수퍼 클래스로 복사한 메서드의 내부 코드를 제거해 빈 메서드로 만든다.   


<br>

#### 구현

State 패턴으로 리팩터링하는 것이 타당한지를 이해하려면, 역으로 State 패턴이 필요한 만큼 복잡하지 않은 상태 관리 로직을 가진 클래스를 살펴보는 것이 도움이 된다. 다음 코드를 보자. 

```java

public class SystemPermission{
    private SystemProfile profile;
    private SystemUser requestor;
    private SystemAdmin admin;
    private String state;
    private boolean isGranted;

    public final static String REQUESTED = "REQUESTED";
    public final static String CLAMED = "CLAMED";
    public final static String DENIED = "DENIED";
    public final static String GRANTED = "GRANTED";

    public StstemPermission(SystemUser requestor, SystemProfile profile){}
        this.requestor = requestor;
        this.profile = profile;
        this.state = REQUESTED;
        this.isGranted = false;
        notifyAdminOfPermissionRequest();
    }

    public void clamedBy(SystemAdmin admin){
        if(!state.equals(REQUESTED)){
            return
        }
        willBeHandledBy(admin);
        this.state = CLAMED;
    }

    public void deniedBy(SystemAdmin admin){
        if(!state.equals(CLAMED)){
            return;
        }
        if(!this.admin.equlas(admin)){
            return;
        }

        isGranted = false;
        state = DENIED;
        notifyAdminOfPermissionRequest();        
    }

    public void grantedBy(SystemAdmin admin){
        if(!state.equals(CLAMED)){
            return;
        }
        if(!this.admin.equlas(admin)){
            return;
        }
        this.state = GRNATED;
        this.isGranted = true;
        notifyAdminOfPermissionRequest();  
    }
    public boolean isGranted(){
        return this.greanted;
    }
    public String getState(){
        return this.state;
    }
} 

public class Test{
    private SystemPermission permission;

    public void setUp(){
        this.permission = new SystemPermission(user, profile);
    }

    public void testGrantedBy(){
        permission.grantedBy(admin);
        assertEquals("requested", permission.REQUESTED, permission.state());
        assertEquals("not granted", false, permission.isGranted());
        permission.claimedBy(admin);
        permission.grantedBy(admin);
        assertEquals("requested", permission.GRANTED, permission.state());
        assertEquals("not granted", true, permission.isGranted());
    }
}
```

아래 클라이언트에서 여러 메서드를 호출할 때 `state`의 값이 바뀌는 방식을 보자. 그리고 클래스 전체에 퍼져있는 상태 전이 로직을 보자. 사실 그렇게 복잡한 로직은 아니다.  
그러나 추후 살을 더 붙여나가다 보면 상태 전이 로직이 따라갈 수 없을 정도로 복잡해지는 것은 한순간이다. 예를 들어, 시스템에 더 권한을 얻기 위해 UNIX/DB 접근 권한을 얻기 위한다면, 모든 상태 로직에 해당 코드가 들어가야 한다.  
```java
if(!state.equals(UNIX_REQUESTED){
    return;
}
if(!state.equals(DB_REQUESTED){
    return;
}
```

등등 이럴 경우를 생각하여, 리팩토링을 해보자.  

##### 절차 1

첫번째로 [Replace Type Code with Class](https://ukcastle.github.io/refactoring/2021/05/12/Replace-Type-Code-with-Class/)를 적용시킨다. 결과적으로 다음과 같은 클래스가 생긴다.  

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
    public final static PermissionState UNIX_REQUESTED = new PermissionState("UNIX_REQUESTED");
    public final static PermissionState UNIX_CLAMED = new PermissionState("UNIX_CLAMED");
}

public class SystemPermission{
    private SystemProfile profile;
    private SystemUser requestor;
    private SystemAdmin admin;
    private String state;
    private boolean isGranted;

    private PermissionState permissionState;

    public StstemPermission(SystemUser requestor, SystemProfile profile){}
        this.requestor = requestor;
        this.profile = profile;
        setState(PermissionState.REQUEST);
        notifyAdminOfPermissionRequest();
    }

    public void clamedBy(SystemAdmin admin){
        if(!state.equals(PermissionState.REQUESTED)){
            return
        }
        if(!state.equals(PermissionState.UNIX_REQUESTED)){
            return
        }
        willBeHandledBy(admin);
        setState(PermissionState.CLAMED);
    }

    public void deniedBy(SystemAdmin admin){
        if(!state.equals(PermissionState.CLAMED)){
            return;
        }
        if(!state.equals(PermissionState.UNIX_CLAMED)){
            return;
        }
        if(!this.admin.equlas(admin)){
            return;
        }

        isGranted = false;
        setState(PermissionState.DENIED);
        notifyAdminOfPermissionRequest();        
    }

    public void grantedBy(SystemAdmin admin){
        if(!state.equals(PermissionState.CLAMED)){
            return;
        }
        if(!state.equals(PermissionState.UNIX_CLAMED)){
            return;
        }
        if(!this.admin.equlas(admin)){
            return;
        }
        setState(PermissionState.GRNATED);
        this.isGranted = true;
        notifyAdminOfPermissionRequest();  
    }
    public boolean isGranted(){
        return this.greanted;
    }

    public setState(PermissionState state){
        permissionState = state;
    }

    public String getState(){
        return permissionState.toString();
    }
} 

```

##### 절차 2

`PermissionState`에는 각각의 상태를 나타내는 6개의 상수가 있다. 이를 전부 서브클래스화 시키고, 수퍼 클래스는 추상화 시켜준다.  
```java
public abstract class PermissionState{

    private final String name;

    private PermissionState(String name){
        this.name = name;
    }

    public String toString(){
        return name;
    }

    public final static PermissionState REQUESTED = new PermissionRequested();
    public final static PermissionState CLAMED = new PermissionRequested();
    public final static PermissionState GRANTED = new PermissionRequested();
    public final static PermissionState DENIED = new PermissionRequested();
    public final static PermissionState UNIX_REQUESTED = new PermissionRequested();
    public final static PermissionState UNIX_CLAMED = new PermissionRequested();
}

public class PermissionRequested extends PermissionState{
    public PermissionRequested(){
        super("REQUESTED");
    }
}
...

```

##### 절차 3

이제 `SystemPermission`에서 상태 전이 로직에 따라 State값을 바꾸는 메서드를 찾을 차례이다. 하나씩 처리해보자.  

```java
public class SystemPermission{
    private SystemProfile profile;
    private SystemUser requestor;
    private SystemAdmin admin;
    private boolean isGranted;

    private PermissionState permissionState;

    public StstemPermission(SystemUser requestor, SystemProfile profile){}
        this.requestor = requestor;
        this.profile = profile;
        setState(PermissionState.REQUEST);
        notifyAdminOfPermissionRequest();
    }

    public void clamedBy(SystemAdmin admin){
        // if(!state.equals(PermissionState.REQUESTED)){
        //     return
        // }
        // if(!state.equals(PermissionState.UNIX_REQUESTED)){
        //     return
        // }
        // willBeHandledBy(admin);
        // setState(PermissionState.CLAMED);
        permissionState.claimedBy(admin, this);
    }

    void willBeHandledBy(SystemAdmin admin){
        this.admin = admin;
    }

    public void deniedBy(SystemAdmin admin){
        // if(!state.equals(PermissionState.CLAMED)){
        //     return;
        // }
        // if(!state.equals(PermissionState.UNIX_CLAMED)){
        //     return;
        // }
        // if(!this.admin.equlas(admin)){
        //     return;
        // }

        // isGranted = false;
        // setState(PermissionState.DENIED);
        // notifyAdminOfPermissionRequest();        
        permissionState.deniedBy(admin, this);
    }

    public void grantedBy(SystemAdmin admin){
        // if(!state.equals(PermissionState.CLAMED)){
        //     return;
        // }
        // if(!state.equals(PermissionState.UNIX_CLAMED)){
        //     return;
        // }
        // if(!this.admin.equlas(admin)){
        //     return;
        // }
        // setState(PermissionState.GRNATED);
        // this.isGranted = true;
        // notifyAdminOfPermissionRequest();  
        permissionState.grantedBy(admin, this);
    }
    public boolean isGranted(){
        return this.greanted;
    }

    public setState(PermissionState state){
        permissionState = state;
    }

    public String getState(){
        return permissionState.toString();
    }
} 


public abstract class PermissionState{

    private final String name;

    private PermissionState(String name){
        this.name = name;
    }

    public String toString(){
        return name;
    }

    public void claimedBy(SystemAdmin admin, SystemPermission permission){
        if(!state.equals(PermissionState.REQUESTED)){
            return
        }
        if(!state.equals(PermissionState.UNIX_REQUESTED)){
            return
        }
        permission.willBeHandledBy(admin);
        permission.setState(PermissionState.CLAMED);
    }

    public void deniedBy(SystemAdmin admin, SystemPermission permission){
        if(!state.equals(PermissionState.CLAMED)){
            return;
        }
        if(!state.equals(PermissionState.UNIX_CLAMED)){
            return;
        }
        if(!this.admin.equlas(admin)){
            return;
        }

        permission.isGranted = false;
        permission.setState(PermissionState.DENIED);
    }

    public void grantedBy(SystemAdmin admin, SystemPermission permission){
        if(!state.equals(PermissionState.CLAMED)){
            return;
        }
        if(!state.equals(PermissionState.UNIX_CLAMED)){
            return;
        }
        if(!this.admin.equlas(admin)){
            return;
        }
        permission.setState(PermissionState.GRNATED);
        permission.isGranted = true;
    }

    public final static PermissionState REQUESTED = new PermissionRequested();
    public final static PermissionState CLAMED = new PermissionRequested();
    public final static PermissionState GRANTED = new PermissionRequested();
    public final static PermissionState DENIED = new PermissionRequested();
    public final static PermissionState UNIX_REQUESTED = new PermissionRequested();
    public final static PermissionState UNIX_CLAMED = new PermissionRequested();
}
```

##### 절차 4

이제 `SystemPermission`이 가질 수 있는 상태를 하나 고르고, `PermissionState`에서 그 상태를 다른 상태로 바꾸는 일을 하는 메서드를 찾을 단계이다. 우선 `REQUESTED` 상태부터 시작하자. 이 상태에서는 `CLAMED` 상태로만 갈 수 있고, `claimedBy()`에서 그 전이가 일어난다.

```java
public class PermissionRequested extends PermissionState{
    public PermissionRequested(){
        super("REQUESTED");
    }

    public void claimedBy(SystemAdmin admin, SystemPermission permission){
        // if(!state.equals(PermissionState.REQUESTED)){
        //     return
        // }
        // if(!state.equals(PermissionState.UNIX_REQUESTED)){
        //     return
        // }
        permission.willBeHandledBy(admin);
        permission.setState(CLAMED);
    }
}
```

if 문을 다 제거할 수 있다, 왜냐? 당연히 이미 권한이 지정되어있기 떄문이다. 첫 단계는 매우 간단했다. 다음 단계로 가자. `PermissionClaimed`에 관한 내용이다.  
이 클래스는 `DENIED`, `GRANTED`, `UNIX_REQUESTED` 상태로 전이할 수 있고, `deniedBy()`와 `grantedBy()`에 전이 코드가 있다.  

```java
public class PermissionClaimed extends PermissionState{
    public PermissionRequested(){
        super("CLAIMED");
    }

    public void deniedBy(SystemAdmin admin, SystemPermission permission){
        // if(!state.equals(PermissionState.CLAMED)){
        //     return;
        // }
        // if(!state.equals(PermissionState.UNIX_CLAMED)){
        //     return;
        // }
        if(!permission.admin.equlas(admin)){
            return;
        }

        permission.isGranted = false;
        permission.isUnixPermissionGranted = false;
        permission.setState(DENIED);
        permission.notifyAdminOfPermissionRequest();
    }

    public void grantedBy(SystemAdmin admin, SystemPermission permission){
        // if(!state.equals(PermissionState.CLAMED)){
        //     return;
        // }
        // if(!state.equals(PermissionState.UNIX_CLAMED)){
        //     return;
        // }
        if(!permission.admin.equlas(admin)){
            return;
        }
        permission.setState(GRNATED);
        permission.isGranted = true;
        permission.notifyAdminOfPermissionRequest();
    }

}
```

##### 절차 5

이제 `PermissionState`의 세가지 상태 전이 함수의 내부 코드를 모두 삭제할 수 있다.  
최종적으로 코드는 이렇게 된다.  

```java

public abstract class PermissionState{

    private final String name;

    private PermissionState(String name){ this.name = name; }

    public String toString(){ return name; }

    public void claimedBy(SystemAdmin admin, SystemPermission permission);
    public void deniedBy(SystemAdmin admin, SystemPermission permission);
    public void grantedBy(SystemAdmin admin, SystemPermission permission);

    public final static PermissionState REQUESTED = new PermissionRequested();
    public final static PermissionState CLAMED = new PermissionRequested();
    public final static PermissionState GRANTED = new PermissionRequested();
    public final static PermissionState DENIED = new PermissionRequested();
    public final static PermissionState UNIX_REQUESTED = new PermissionRequested();
    public final static PermissionState UNIX_CLAMED = new PermissionRequested();
}

public class PermissionRequested extends PermissionState{
    public PermissionRequested(){
        super("REQUESTED");
    }
    public void claimedBy(SystemAdmin admin, SystemPermission permission){
        permission.willBeHandledBy(admin);
        permission.setState(CLAMED);
    }
}

public class PermissionClaimed extends PermissionState{
    public PermissionRequested(){
        super("CLAIMED");
    }

    public void deniedBy(SystemAdmin admin, SystemPermission permission){
        if(!permission.admin.equlas(admin)){
            return;
        }
        permission.isGranted = false;
        permission.isUnixPermissionGranted = false;
        permission.setState(DENIED);
        permission.notifyAdminOfPermissionRequest();
    }

    public void grantedBy(SystemAdmin admin, SystemPermission permission){
        if(!permission.admin.equlas(admin)){
            return;
        }
        permission.setState(GRNATED);
        permission.isGranted = true;
        permission.notifyAdminOfPermissionRequest();
    }
}

...

public class SystemPermission{
    private SystemProfile profile;
    private SystemUser requestor;
    private SystemAdmin admin;
    private boolean isGranted;

    private PermissionState permissionState;

    public StstemPermission(SystemUser requestor, SystemProfile profile){}
        this.requestor = requestor;
        this.profile = profile;
        setState(PermissionState.REQUEST);
        notifyAdminOfPermissionRequest();
    }

    public void clamedBy(SystemAdmin admin){
        permissionState.claimedBy(admin, this);
    }

    void willBeHandledBy(SystemAdmin admin){
        this.admin = admin;
    }

    public void deniedBy(SystemAdmin admin){ 
        permissionState.deniedBy(admin, this);
    }

    public void grantedBy(SystemAdmin admin){
        permissionState.grantedBy(admin, this);
    }
    public boolean isGranted(){
        return this.greanted;
    }

    public setState(PermissionState state){
        permissionState = state;
    }

    public String getState(){
        return permissionState.toString();
    }
} 
```




