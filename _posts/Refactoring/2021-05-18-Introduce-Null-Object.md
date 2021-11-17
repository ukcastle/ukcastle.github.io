---
layout: post
title: Introduce Null Object (402)
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

#### 개요

> 어떤 필드나 변수의 값이 Null인지를 검사하는 로직이 여기저기에 중복되어 있다면
값이 Null일 경우 행할 작업을 대신하는 Null Object를 사용하도록 수정한다.

<br>

#### 동기

클라이언트가 어떤 변수의 메서드를 호출하는데, 그 **값이 Null이라면** Exception이 발생하거나 문제가 발생할 수 있다. 그런 상황으로부터 시스템을 보호하기 위해 값이 널인지를 검사하여 별도의 동작으로 분기하는 코드를 작성하는 것이 보통이다.

```java
if (someObject != null){
    someObject.doSomething();
} else {
    doOtherthing();
}
```

위와 같은 검사가 한두곳에서 반복되는 것은 별 문제가 되지 않는다. 하지만 **여러 곳에서 자주 반복**된다면 말이 다르다. 오류가 생길 가능성이 많고, 관리하기 힘들어진다.  

Null Object 패턴은 이런 문제를 위한 해결책으로, 필드나 변수가 Null이 아니도록 유지하여 별도의 검사를 하지 않아도 된다. 이 때 Null 객체의 해당 메서드는 아무 일도 하지 않거나 디폴트 동작을 하는 둥 전체 동작에 영향이 없는 작업을 한다. 이런 방식을 사용하면 변수가 Null일 가능성을 걱정하지 않아도 된다.  

시스템에 Null 객체를 도입하면, **코드 크기가 줄어들거나 적어도 그대로 유지**되어야 한다. 만약 **그렇지 않다면, 해당 리팩토링을 사용할 필요가 없다.**   

Null 객체를 도입한다고 해서 검사 로직이 자동으로 사라지지는 않는다. Null 객체 덕분에 **Null 값으로부터 이미 보호되고 있다는 사실을 모르는 프로그래머도 있을것이다.** 그렇다면 그는 아마도 **절대 Null이 될 일이 없는 경우에 대한 Null 검사 로직을 작성할 것**이다. 또한 **특정 상황에서 Null 값이 리턴되기를 기대**하고 그에 맞춘 코드를 작성한다면, **원하지 않는 결과**가 발생할 것이다.   

**서브클래싱을 이용해 Null Object 패턴을 구현**하는 경우, Null 값을 위한 적절한 동작을 부여하기 위해 **상속된 모든 public 메서드를 오버라이드**해야한다. 따라서 수퍼클래스에 새로운 메서드를 추가할 때 반드시 Null 객체 클래스에 그 메서드를 오버라이드해야 한다는 단점이 있다. 만약 이것을 잊는다면, Null 객체는 상속된 기능대로 동작할 것이고, 런타임에 예상치 못한 문제가 발생할 것이다. 반면 **인터페이스를 이용하면 그럴 위험이 없다.**  

##### 장점

- 수 많은 Null 검사 로직 없이도 Null 값으로 인한 에러를 막을 수 있다.
- Null 검사 로직이 최소화되어 코드가 간단해진다

##### 단점

- 시스템에 Null 검사 로직이 별로 필요하지 않은 상황에서는 설계만 복잡해진다.  
- 프로그래머가 Null 객체의 존재를 모른다면, 쓸데없는 Null 검사를 반복하게된다.  
- 유지보수가 복잡해진다.  

<br>

#### 절차

여기서 제시할 절차는 어떤 필드나 변수 값이 Null일때 이를 참조하는 것을 막기 위해 코드 여기저기에 Null 검사 로직이 존재하는 상황을 가정한다. 그리고 이후 사용하는 **'원천 클래스'** 라는 용어는 해당 타입의 필드나 변수를 Null 값으로부터 보호해야 할 클래스를 자칭한다.  

1. 원천 클래스에 [Extract Subclass](https://ukcastle.github.io/refactoring/2021/04/07/RefactoringToPattern/#extract-subclass)를 적용하거나 그 클래스가 구현하고 있는 인터페이스를 구현하여 Null 객체 클래스를 만든다, 인터페이스를 이용하고 싶은데 원천 클래스가 구현하는 인터페이스가 없다면 [Extract Interface](https://ukcastle.github.io/refactoring/2021/04/07/RefactoringToPattern/#extract-interface)를 적용하여 인터페이스를 직접 만들어도 좋다.

2. 원천 클래스를 사용하는 클라이언트 코드에서 Null 검사 로직을 찾는다. 그리고 그때 호출되는 메서드는 Null 객체 클래스가 오버라이드하게 만들고, 값이 Null일 경우 동작을 수행하도록 구현한다.  

3. 원천 클래스와 관련된 다른 모든 Null 검사 로직에 대해 단계 2를 반복한다.

4. Null 검사 로직이 하나 이상 존재하는 클래스를 찾아, Null 검사 로직에서 참초하는 필드나 변수를 앞서 만든 Null 객체로 초기화한다. 단 이 초기화 작업은 클래스 인스턴스의 생존 기간중 되도록 이른 시기에(예를 들면, 생성될 떄) 이뤄지도록 해야 한다.  

5. 단계 4에서 작업한 클래스에 있는 Null 검사 로직을 모두 제거한다. 

6. 단계 4와 5의 작업을 Null 검사 로직이 있는 모든 클래스에 적용한다.  

<br>

#### 구현

java의 `MouseEventHandler` 객체를 아는가? 특정 마우스 이벤트가 있을 때 작동하는 객체이다. 다음 코드를 보자.  

```java

public class NavigationApplet extends Applet{
    public boolean mouseMove(Event e, int x, int y){
        if (MouseEventHandler != null){
            return mouseEventHandler.mouseMove(graphicsContext, event, x, y);
        }
        return true;
    }
    public boolean mouseDown(Event e, int x, int y){
        if (MouseEventHandler != null){
            return mouseEventHandler.mouseDown(graphicsContext, event, x, y);
        }
        return true;
    }
    public boolean mouseUp(Event e, int x, int y){
        if (MouseEventHandler != null){
            return mouseEventHandler.mouseUp(graphicsContext, event, x, y);
        }
        return true;
    }
    public boolean mouseExit(Event e, int x, int y){
        if (MouseEventHandler != null){
            return mouseEventHandler.mouseExit(graphicsContext, event, x, y);
        }
        return true;
    }
}
```

Null 검사 로직을 제거하기 위해 우리는 `Applet`을 리팩터링해 초기화가 끝나기 전에는 `NullMouseEventHandler` 객체로 사용하다 준비가 끝나면 `MouseEventHandler`로 사용하는 과정을 볼 것이다.  


##### 절차 1

`MouseEventHandler` 클래스에 Extract Subclass를 적용해 `NullMouseEventHandler`를 생성한다.  

```java
public class NullMouseEventHandler extends MouseEventHandler{
    public NullMouseEventHandelr(Context context){
        super(context);
    }
}
```

##### 절차 2

다음, 아래와 같은 Null 검사 로직을 찾았다.  

```java
public boolean mouseMove(Event e, int x, int y){
        if (MouseEventHandler != null){
            return mouseEventHandler.mouseMove(graphicsContext, event, x, y);
        }
        return true;
    }
```

Null 검사 로직에서 호출되는 메서드는 `mouseEventHandler.mouseMove` 이다. 이렇다면, 다음과 같이 `NullMouseEventHandler`를 수정해준다.  

```java
public class NullMouseEventHandler extends MouseEventHandler{
    public NullMouseEventHandelr(Context context){
        super(context);
    }

    @Override
    public boolean mouseMove(Event e, int x, int y){
        return true;
    }
}
```

##### 절차 3

모든 Null 검사 로직이 있는 함수에 대하여 반복한다.  

```java
public class NullMouseEventHandler extends MouseEventHandler{
    public NullMouseEventHandelr(Context context){
        super(context);
    }

    @Override
    public boolean mouseMove(Event e, int x, int y){
        return true;
    }
    @Override
    public boolean mouseDown(Event e, int x, int y){
        return true;
    }
    @Override
    public boolean mouseUp(Event e, int x, int y){
        return true;
    }
    @Override
    public boolean mouseExit(Event e, int x, int y){
        return true;
    }
}
```

##### 절차 4

다음 원천클래스 내의 Null 검사 로직이 참조하는 필드인 `mouseEventHandelr`를 `NullMouseEventHandler`로 초기화한다.  

```java
public class NavigationApplet extends Applet{
    private MouseEventHandler mouseEventHandler = new NullMouseEventHandelr();
}
```

다음 Null Object도 생성자를 바꿔준다.  

```java
public class NullMouseEventHandler extends MouseEventHandler{
    public NullMouseEventHandelr(){
        super(null);
    }
    ...
}
```

##### 절차 5

다음, 기존 코드들을 전부 바꿔준다.  

```java

public class NavigationApplet extends Applet{
    private MouseEventHandler mouseEventHandler = new NullMouseEventHandelr();

    public boolean mouseMove(Event e, int x, int y){
        // if (MouseEventHandler != null){
        return mouseEventHandler.mouseMove(graphicsContext, event, x, y);
        // }
        // return true;
    }
    public boolean mouseDown(Event e, int x, int y){
        // if (MouseEventHandler != null){
        return mouseEventHandler.mouseDown(graphicsContext, event, x, y);
        // }
        // return true;
    }
    public boolean mouseUp(Event e, int x, int y){
        // if (MouseEventHandler != null){
        return mouseEventHandler.mouseUp(graphicsContext, event, x, y);
        // }
        // return true;
    }
    public boolean mouseExit(Event e, int x, int y){
        // if (MouseEventHandler != null){
        return mouseEventHandler.mouseExit(graphicsContext, event, x, y);
        // }
        // return true;
    }
}
```

##### 절차 6

다른 클래스에도 모든 과정을 반복한다.  

아래는 전체 코드이다. 

```java

public class NullMouseEventHandler extends MouseEventHandler{
    public NullMouseEventHandelr(){
        super(null);
    } 
    @Override
    public boolean mouseMove(Event e, int x, int y){
        return true;
    }
    @Override
    public boolean mouseDown(Event e, int x, int y){
        return true;
    }
    @Override
    public boolean mouseUp(Event e, int x, int y){
        return true;
    }
    @Override
    public boolean mouseExit(Event e, int x, int y){
        return true;
    }
}

public class NavigationApplet extends Applet{
    private MouseEventHandler mouseEventHandler = new NullMouseEventHandelr();

    public boolean mouseMove(Event e, int x, int y){
        return mouseEventHandler.mouseMove(graphicsContext, event, x, y);
    }
    public boolean mouseDown(Event e, int x, int y){
        return mouseEventHandler.mouseDown(graphicsContext, event, x, y);
    }
    public boolean mouseUp(Event e, int x, int y){
        return mouseEventHandler.mouseUp(graphicsContext, event, x, y);
    }
    public boolean mouseExit(Event e, int x, int y){
        return mouseEventHandler.mouseExit(graphicsContext, event, x, y);
    }
}
```