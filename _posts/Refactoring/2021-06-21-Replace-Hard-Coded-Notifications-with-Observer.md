---
layout: post
title: Replace Hard-Coded Notifications with Observer
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

#### 개요

>어떤 상속 구조 내의 서브클래스들이 **자신과 관련된 클래스에 통보하는 기능**을 **하드 코딩**으로 각자 구현하고 있다면,
Observer 인터페이스를 통해 그 **수퍼클래스**가 임의의 **다른 클래스에 통보할 수 있도록** 일반적인 통보기능을 만들고 **서브클래스를 제거**한다.

<br>

#### 동기

[Observer 패턴](https://jo631.github.io/designpattern/2021/06/20/Observer/)은 통보의 주체와 관찰자의 결합이 느슨해진다. 모든 관찰자를 위한 관찰자 인터페이스를 정의하여 사용하기 떄문이다. 변경된 상태를 통보 받기 위한 클래스는 **관찰자 인터페이스를 구현**하고 **자신을 통보 주체에 등록**만 하면 된다. 통보 주체 클래스는 관찰자 인터페이스를 구현한 객체의 컬렉션을 보관하고 있다가 상태의 변화가 생겼을 때 그들에게 통보만 하면 된다.  
통보 주체는 관찰자 객체를 컬렉션에 추가하고 삭제하는 메서드를 포함해야 한다.  
옵저버 패턴을 구현할 때 흔히 생기는 문제점은 두 가지가 있다.  
1. 통보 체인
2. 메모리 누수

통보 체인은 한 관찰자가 통보를 받았을 때 자신이 다시 그 주체가 되어 또 다른 관찰자에게 통보하고... 반복되는 체인이다. 이런 체인이 불가피한 상황에서는 중재자 패턴을 도입하는 것을 추천한다.  

메모리 누수는 제대로 Garbage Colletion이 되지 않은 현상인데, 그 이유는 통보 주체가 관찰자 객체의 참조를 가지고 있기 때문이다.  

##### 장점

- 통보 주체 클래스와 관찰자 클래스의 결합을 느슨하게 한다.  
- 관찰자가 여럿인 경우도 지원한다.

##### 단점

- 필요하지 않은 상황에서 적용한다면 설계만 복잡해진다.
- 통보 체인이 불가피한 상황에서는 설계가 더 복잡해진다.  
- 관찰자 객체에 대한 참조를 제때 삭제하지 않으면 메모리 누수가 발생한다.  

<br>

#### 절차

다른 객체의 참조를 갖고 있다가 어떤 통보를 보내는 클래스를 **통보자**, 통보자에 자신을 등록하고 통보를 받는 클래스를 **수령자**라고 하자. 이 리팩터링을 통해 쓸대없는 통보자 클래스를 없애고, 수령자를 **관찰자**로 변경하는 과정을 설명한다.

1. 통보자가 수령자를 대신해 어떤 기능을 수행하고 있다면, [Move Method](https://jo631.github.io/refactoring/2021/04/07/RefactoringToPattern/#move-method)를 적용해 그 기능을 수령자로 옮긴다. 작업이 끝나면 통보자에게는 순수한 통보 메서드만 남는다.  

2. 수령자의 메서드 중 통보자가 호출하는 메서드에 [Extract Interface](https://jo631.github.io/refactoring/2021/04/07/RefactoringToPattern/#extract-interface)를 적용해 **관찰자 인터페이스**를 만든다. 다른 수령자에게 이 인터페이스에 없는 메서드가 있다면, 그 메서드도 추가한다.

3. 모든 수령자가 앞서 만든 관찰자 인터페이스를 구현하도록 수정한다. 그리고 모든 통보자가 관찰자 인터페이스를 통해 수령자에게 통보하도록 수정한다.  

4. 통보자를 하나 고른 후, 통보 메서드에 [Pull Up Method](https://jo631.github.io/refactoring/2021/04/07/RefactoringToPattern/#pull-up-method)를 적용한다. 이 과정에서 관찰자 인터페이스 타입의 필드를 선언하고 참조를 등록하는 코드도 함께 옮긴다. 이 통보자의 수퍼클래스는 이제 통보 주체가 된다.

5. 이제 통보자 대신 통보 주체에 대한 모든 관찰자를 등록하고 그와 통신하도록 수정한다. 그리고 통보자를 제거한다.  

6. 통보 주체가 한 개의 관찰자에 대한 참조 대신 **관찰자 객체의 컬렉션**을 가지도록 리팩터링 한다. 이렇게 하고 나면 관찰자가 통보 주체에 자신을 등록하는 방식도 바꿔야 하는데, 관찰자 객체 하나를 추가하는 메서드로 바꾸면 된다. 마지막으로 컬렉션을 순회하며 통보하도록 메서드를 고친다.

<br>

#### 구현

자바의 단위테스트 도구인 `JUnit`의 초기버전 코드의 일부분이다.

```java
class UITestResult extends TestResult{
    private TestRunner fRunner;
    
    UiTestResult(TestRunner runner){
        this.fRunner = runner;
    }

    public synchronized void addFailure(Test test, Throwable t){
        super.addFailure(test, t);
        fRunner.addFailure(this, test, t)
    }
    ...
}

package ui;
public class TestRunner extends Frame{
    private TestResult fTestResult;

    protected TestResult createTestResult(){
        return new UITestResult(this);
    }

    synchronized void addFailure(TestResult result, Test test, Throwable t){
        ...
    }
}

public class TextTestResult extends TestResult{
    public synchronized void addError(Test test, Throwable t){
        super.addError(test, t);
        System.out.println("E");
    }
    public synchronized void addFailure(Test test, Throwable t){
        super.addError(test, t);
        System.out.println("F");
    }
}

```

기존의 2.x 버전에서는 문제 없이 구동되었지만, 여러 객체가 한 TestResult 객체를 동시에 관찰할 수 있게 해달라는 요구가 생기면서 상황이 달라졌다.  
그러면서 해당 리팩터링을 적용한다.  

##### 과정 1

모든 통보자에 순수한 통보 기능만을 남기고 다른 기능을 제거하는 것이다. `UITestResult`에는 통보 기능 뿐이지만 `TextTestResult`에는 그렇지 않다. 해당 클래스는 결과를 통보하는 대신 콘솔에 메시지를 직접 출력한다.

해당 클래스에 Move Method 리팩터링을 적용해 화면 출력 부분을 `TestRunner`로 옮긴다.

```java
public class TextTestResult extends TestResult{
    private TestRunner fRunner;
    TextTestResult(TestRunner runner){
        this.fRunner = runner;
    }

    public synchronized void addError(Test test, Throwable t){
        super.addError(test, t);
        this.fRunner.addError(this, test, t);
    }
    public synchronized void addFailure(Test test, Throwable t){
        super.addError(test, t);
        System.out.println("F");
    }
}

public class TestRunner{
    protected TextTestResult createTestResult(){
        return new TextTestResult(this);
    }

    //옮겨온 메서드
    public void addError(TestResult testResult, Test test, Throwable t){
        System.out.println("E");
    }
}
```

##### 과정 2

이제 관찰자 인터페이스를 만들 차례다. `TextTestResult`에 대응하는 `TestRunner` 클래스에 Extract Interface를 적용해 `TestListner` 인터페이스를 만든다. 새로 만든 인터페이스에 포함시킬 메서드를 정하려면, `TextTestResult`에서 호출하는 메서드가 어떤 것들인지 알아야 한다. 위의  `this.fRunner.addError()`와 같은 메서드들이다.

따라서 인터페이스는 다음과 같다.

```java
public Interface TestListener{
    public void addError(TestResult testResult, Test test, Throwable t);
    public void addFailure(TestResult testResult, Test test, Throwable t);
    public void startTest(TestResult testResult, Test test);
    public void endTest(TestResult testResult, Test test);  //uiTest에만 있는 메서드
} 

public class TestRunner implement TestListener{ 
    ...
    public void endTest(TestResult testResult, Test test){} // 구현만 해논다.
}

```

##### 과정 3

uiTestClass가 인터페이스를 구현하도록 수정한다. 

```java
public class TestRunner extends Frame implements TestListener{...}

class UITestResult extends TestResult{
    protected TestListener fRunner;

    UITestResult(TestListener runner){
        this.fRunner = runner;
    }
}

public class TextTestResult extends TestResult{
    protected TestListener fRunner;

    TextTestResult(TestListener runner){
        this.fRunner = runner;
    }
}
```


##### 과정 4

두 TestResult의 모든 통보 메서드에 대해 Pull Up Method를 적용할 차례다. 

```java
public class TestResult{
    protected TestListener fRunner;

    public TestResult(TestListener runner){
        this();
        this.fRunner= runner;
    }

    public TestResult(){
        fFailures = new Vector(10);
        fErrors = new Vector(10);
        fRunTests = 0;
        fStop = false;
    }

    public void addError(TestResult testResult, Test test, Throwable t){
        fErrors.addElements(new TestFailure(test, t));
        fRunner.addError(this, test, t);
    }
    public void addFailure(TestResult testResult, Test test, Throwable t){
        fFailures.addElements(new TestFailure(test, t));
        fRunner.addFailure(this, test, t);
    }
    public void startTest(TestResult testResult, Test test){
        fRunner.endTest(this, test);
    }
    public void endTest(TestResult testResult, Test test){
        fRunTests++;
        fRunner.startTest(this, test);
    }
}

package ui;
class UITestResult extends TestResult {}
class TextTestResult extends TestResult {}
```

##### 과정 5

이제 `TestRunner`가 `TestResult`와 직접적인 관계가 되도록 고칠 수 있다.

```java
package textui;
public class TestRunner implements TestListener{
    protected TestResult createTestResult(){
        return new TestResult(this);
    }

    protected void doRun(Test suite, boolean wait){
        ...
        TestResult result = createTestResult();
    }
}

```

이로써 두 `TestRunner` 클래스가 `TestResult`의 관찰자가 되었다.

6. 마지막으로 `TestResult` 객체 하나에 대한 관찰자가 동시에 여러 개 존재할 수 있도록 만들자.

```java
public class TestResult{
    ...
    private List observers = new ArrayList();

    public void addObserver(TestListener testListener){
        observers.add(testListener);
    }

    public void addError(TestResult testResult, Test test, Throwable t){
        fError.addElement(new TestFailure(test, t));
        for (Iterator i=observers.iterator();i.hasNext();){
            TestListener observer = (TestListener)i.next();
            observer.addError(this, test, t);
        }
    }
}

package textui;

public class TestRunner implements TestListener{
    protected TestResult createTestResult(){
        TestResult testResult = new TestResult();
        testResult.addObserver(this);
        return testResult;
    }

    protected void doRun(Test suite, boolean wait){
        ...
        TestResult result = createTestResult();
    }
}
```

이걸로 리팩터링이 완료되었다. 전체 코드는 다음과 같다.

```java
public Interface TestListener{
    public void addError(TestResult testResult, Test test, Throwable t);
    public void addFailure(TestResult testResult, Test test, Throwable t);
    public void startTest(TestResult testResult, Test test);
    public void endTest(TestResult testResult, Test test);  //uiTest에만 있는 메서드
} 

public class TestResult{
    protected TestListener fRunner;
    private List observers = new ArrayList();

    public void addObserver(TestListener testListener){
        observers.add(testListener);
    }

    public TestResult(TestListener runner){
        this();
        this.fRunner= runner;
    }

    public TestResult(){
        fFailures = new Vector(10);
        fErrors = new Vector(10);
        fRunTests = 0;
        fStop = false;
    }

    public void addError(TestResult testResult, Test test, Throwable t){
        fError.addElement(new TestFailure(test, t));
        for (Iterator i=observers.iterator();i.hasNext();){
            TestListener observer = (TestListener)i.next();
            observer.addError(this, test, t);
        }
    }
    public void addFailure(TestResult testResult, Test test, Throwable t){
        ...
    }
    public void startTest(TestResult testResult, Test test){
        ...
    }
    public void endTest(TestResult testResult, Test test){
        ...
    }
}


package ui;
public class TestRunner implements TestListener{
    protected TestResult createTestResult(){
        TestResult testResult = new TestResult();
        testResult.addObserver(this);
        return testResult;
    }

    protected void doRun(Test suite, boolean wait){
        ...
        TestResult result = createTestResult();
    }
}
package textui;
public class TestRunner implements TestListener{
    protected TestResult createTestResult(){
        TestResult testResult = new TestResult();
        testResult.addObserver(this);
        return testResult;
    }

    protected void doRun(Test suite, boolean wait){
        ...
        TestResult result = createTestResult();
    }
}
```

