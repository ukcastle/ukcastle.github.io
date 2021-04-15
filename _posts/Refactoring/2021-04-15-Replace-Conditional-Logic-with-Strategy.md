---
layout: post
title: Replace Conditional Logic with Strategy (187)
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

#### 개요

>메서드 내의 조건문을 통해 여러 개의 서로 다른 로직(계산법) 가운데 어떤 것을 실행할지 선택하고 있다면
>각 계산법에 대응하는 스트레티지 클래스를 만들고 해당 스트레티지 인스턴스에 계산을 위임하도록 메서드를 수정한다.  

#### 동기

이 리팩토링의 동기는 한 클래스가 조건문으로 인해 복잡하고, 많은 역할을 수행하고 있을 때 이용된다.  
많은 알고리즘에서 조건문으로 인해 매우 복잡해지고 길어진 함수를 볼 수 있다. 이를 해결하기 위해 [Composed Method](https://jo631.github.io/refactoring/2021/04/14/Compose-Method/)같은 방식을 이용하면, 한 클래스 안에 너무 많은 메소드가 생긴다. 이것을 해결하기 위해 서브클래스를 만들거나, 새로운 클래스로 옮길 수 있는데, 서브클래스로 옮기면 **상속**을 이용하는 방식이고 새로운 클래스로 옮기면 **객체 조합**을 이용하는 것이다.  
이 리팩토링은 **객체 조합**을 이용하는 방식이고, [Replace Conditional with Polymorphism](https://jo631.github.io/refactoring/2021/04/09/RefactoringToPattern/#replace-conditional-with-polymorphism)은 상속을 이용하는 방법이다.  
이 리팩터링의 전제 조건은 알고리즘을 구현한 호스트 클래스가 서브 클래스를 가지고 있어야 한다는 것이다.  서브 클래스가 이미 존재하여 알고리즘들을 각각 하나의 서브클래스에 대응시키기 쉽다면, 이 리팩토링이 제격이다.  
그러나 리팩토링에 앞서 서브클래스를 먼저 만들어야 하는 경우라면, 객체 조합을 이용한 [Strategy](https://jo631.github.io/designpattern/2021/04/14/Strategy/)가 더 쉬운 길이 아닐지 생각해봐야 한다.  
알고리즘 내의 조건문이 타입 코드를 사용하는 경우엔 타입 코드의 값 하나마다 호스트 클래스의 서브 클래스를 하나씩 만드는 [Replace Type Code with Subclasses](https://jo631.github.io/refactoring/2021/04/09/RefactoringToPattern/#replace-type-code-with-subclasses) 방식을 이용하면 쉽다. 타입 코드가 없는 경우 Strategy 패턴으로 리팩터링하는 것이 좋다.  
마지막으로 클라이언트에서 런타임에 한 알고리즘을 다른 알고리즘으로 교체할 필요가 있다면, 상속을 이용한 방법은 피하는 것이 좋다. 이것은 한 Strategy 인스턴스를 다른 인스턴스를 바꾸는 것이 아니라 클라이언트가 사용하는 객체의 타입을 바꾸는 것이 되기 때문이다.  

Strategy 패턴을 **목표로(to)** 리팩터링 할것인지, **향해서(toward)** 리팩터링 할 것인지 결정하려면, 각 Strategy 내의 로직에서 실행에 필요한 데이터에 어떻게 접근하도록 할지 고려해야 한다.  
여기엔 두가지 방법이 있는데,  
- 첫번째로는 호스트 클래스의 **인스턴스**를 Strategy에 **직접** 넘겨서 호스트 클래스의 메서드 호출을 통해 필요한 데이터를 얻게하는 것이다. 이때 호스트 클래스를 **Context**라고 부른다.  
- 두번째로는 필요한 데이터를 파라미터로 Strategy에 넘기는 방법이다.

Strategy 패턴을 사용한 설계를 구현할 때 Context 클래스가 Strategy 클래스 객체를 어떻게 얻게 할 것인지 고민이 될 것이다. 객체의 조합이 별로 많이 않다면, Strategy 객체를 생성하고 이를 Context에 넘겨주는 과정을 클라이언트가 신경쓰지 않게 하는것이 좋다. 그렇게 하는데는 [Encapsulate Classes with Factory](https://jo631.github.io/refactoring/2021/04/13/Encapsulate-Classes-with-Factory/) 방식이 도움이 된다.  


#### 장점

- 조건 로직을 줄이거나 제거해 알고리즘을 명확하게 한다.  
- 알고리즘 내의 여러 로직을 상속 구조로 옮겨 클래스를 단순화한다.  
- 런타임에 어떤 알고리즘을 다른 알고리즘으로 변경하기가 쉬워진다.  

#### 단점 

- 상속을 이용하거나 Simplifying Conditional Expressions 리팩터링을 적용하는 것이 더 쉬울 만한 상황에서는 설계를 복잡하게 만들 뿐이다
- 알고리즘이 컨텍스트 클래스와 데이터를 주고받는 방식이 복잡해진다. 

(본 예제에서 뼈저리게 느꼈다.)


#### 절차

**복잡한 조건 로직으로 이루어진 계산 메소드**를 가지고 있는 클래스. 즉 Context 클래스를 찾는 것이 첫 단계다.  

1. Strategy로 쓸 클래스를 하나 만든다. 클래스의 이름은 수행하는 작업에 맞게 붙인다. 예를들어 `함수이름+Strategy()`와 같은 방식으로.. 
<br> 
2. [Move Method](https://jo631.github.io/refactoring/2021/04/09/RefactoringToPattern/#move-method) 리팩터링으로 계산 메소드를 Strategy 클래스로 옮긴다.  
함수를 옮겼으면, **2가지 상황**이 발생한다.  
    1. Context 클래스에서 쓰던 변수들을 같이 가져오지 않아 에러가 발생하고,  
    2. Context 클래스에서는 더이상 안쓰이는 함수들이 존재한다.  

    일단 1번을 처리해야 한다. 이 때도 두가지 방법이 있다.  
    - 변수의 양이 얼마 되지 않으면, 매개변수로 가져온다.  
        > 이렇게 할 시 서로의 의존성은 더 줄어들어 멋진 코드가 된다.  
    - 변수의 양이 많거나, Context객체의 필드 값을 수정시킬 경우, 레퍼런스를 참조하여 가져온다.  
        > 이렇게 할 시 더 편하게 문제를 해결할 수 있지만, 여러가지 `getter` 메소드를 public으로 둬야할 수도 있어서 **정보 은닉을 위반**할 수도 있다. 보통 같은 패키지 안에 넣어놓아 접근을 한다.  

    다음 2번을 처리하는데, **Context에서 더이상 쓰이지 않는 함수들을 Strategy 클래스로 이동**시킨다.
<br>
3. Context 클래스의 코드 중 Strategy 객체를 생성하고 필드에 대입하는 부분에 [Extract Parameter](https://jo631.github.io/refactoring/2021/04/16/Extract-Parameter/)을 적용하여 클라이언트가 **Context에 Strategy를 넘겨주는 모양새**가 되도록 한다.  
<br>
4. Strategy의 계산 메소드에 [Replace Conditional with Polymorphism](https://jo631.github.io/refactoring/2021/04/09/RefactoringToPattern/#replace-conditional-with-polymorphism)을 적용한다. 이 때 [Replace Type Code with Subclasses](https://jo631.github.io/refactoring/2021/04/09/RefactoringToPattern/#replace-type-code-with-subclasses)를 사용할 것인지, 아니면  [Replace Type Code with State/strategy](https://jo631.github.io/refactoring/2021/04/09/RefactoringToPattern/#replace-type-code-with-statestrategy)를 사용할 것이지 결정해야 한다. 어지간해선 **전자**를 선택하는 것이 좋다. 타입 코드가 명확하게 존재하지 않더라도 전자를 적용할 수 있다. **계산 메서드의 조건 로직의 검사 조건을 타입 코드**라고 생각하면 된다.  

한번에 하나의 서브클래스를 만드는 데 집중해야 한다. 이 과정을 마치고 나면 계산 메서드에 있던 조건 로직이 훨씬 간단해지거나, 심지어 없어질 수도 있다. 그 뒤엔 각 알고리즘에 해당하는 Strategy를 여러 개 얻게 된다. 가능하면 이들의 수퍼클래스를 `abstract`하게 만들면 좋다.  

#### 예제  

가정해보자. 고객의 상황에 맞춰 카메라를 대여해주는 카메라숍이 있다.  

카메라샵은 주기에 맞춰 새로운 카메라를 가져온다. 이 카메라는 풀프레임인지 아닌지, DSLR인지 아닌지 총 4가지의 경우의 수를 가지고있다.    
그리고 경우의 수에 따라 **플랜A,B,C,D**로 나뉘어진다.
그리고 기간을 정하면, 일일마다 계산하여 가격이 측정되는데,  
풀프레임이면 기간이 오래될수록 가격이 더 비싸디. 
또한 똑딱이면 렌즈 보증금이 제외되므로 가격이 더 싸다.    
이를 클래스로 나타내보았다.  

```c++
class CameraLoan{
private:
    bool fullFrame;
    bool DSLR;

public:
    double getFullFramePrice(string n);
    double getPrice(string n);
    double compactPrice(double d);
    double DSLR_Price(double d);

    double loanCamera(string name, double duration){

        if (fullFrame && DSLR){ //Plan A
            return new LoanPlanA(getFullFramePrice(name) * DSLR_Price(duration));
        } else if (fullFrame && !DSLR) { //Plan B
            return new LoanPlanB(getFullFramePrice(name) * compactPrice(duration));
        } else if (!fullFrame && DSLR) { //Plan C
            return new LoanPlanC(getPrice(name) * DSLR_Price(duration));
        } else{ //Plan D
            return new LoanPlanD(getPrice(name) * compactPrice(duration));
        }
        return 0.0;           
    }
};

class LoanPlanA : public CameraLoan {
public:
    LoanPlanA(string name, double duration){
        // ...
    }
}

class LoanPlanB : public CameraLoan{
    //... 이하 생략
}
```

써보니 예시가 적합하진 않은것같은데 아주 간단하게 4가지 상태에 따라 다른 값을 출력하는 함수 `loanCamera()`로 설정하였다.  
또한 나머지 4개를 함수는 단 하나의 함수를 위해 만들어져있다.   
이럴 경우 이 리팩토링이 매우 효과적이다. 이 리팩토링을 거치면 클래스에게 너무 많은 책임이 가해지는 문제를 해결할 수도 있다.   
```c++
class loanCameraStrategy{

public:
    double loanCamera(string name, double duration){
        return 0.0;
    }
}
```

`loanCamera()`함수에 대한 클래스이니, 뒤에 Strategy와 같은 이름을 붙여 선언해주는것이 좋다.  
다음 Move Method를 이용하여 함수를 옮길것이다.  
**다만!** 기존의 Context 클래스에서도 `loanCamera()`는 남겨두어야 한다.  


```c++
class loanCameraStrategy{

public:
    double loanCamera(string name, double duration){
        if (fullFrame && DSLR){ // Plan A
            return new LoanPlanA(getFullFramePrice(name) * DSLR_Price(duration));
        } else if (fullFrame && !DSLR) { // Plan B
            return new LoanPlanB(getFullFramePrice(name) * compactPrice(duration));
        } else if (!fullFrame && DSLR) { // Plan C
            return new LoanPlanC(getPrice(name) * DSLR_Price(duration));
        } else{ // Plan D
            return new LoanPlanD(getPrice(name) * compactPrice(duration));
        }
        return 0.0;  
    }
}
```

일단 함수를 옮겼으면, **2가지 상황**이 발생한다.  
1. Context 클래스에서 쓰던 변수들을 같이 가져오지 않아 에러가 발생하고,  
2. Context 클래스에서는 더이상 안쓰이는 함수들이 존재한다.  

일단 1번을 처리해야 한다. 이 때도 두가지 방법이 있다.  
- 변수의 양이 얼마 되지 않으면, 매개변수로 가져온다.  
    > 이렇게 할 시 서로의 의존성은 더 줄어들어 멋진 코드가 된다.  
- 변수의 양이 많거나, Context객체의 필드 값을 수정시킬 경우, 레퍼런스를 참조하여 가져온다.  
    > 이렇게 할 시 더 편하게 문제를 해결할 수 있지만, 여러가지 `getter` 메소드를 public으로 둬야할 수도 있어서 **정보 은닉을 위반**할 수도 있다. 보통 같은 패키지 안에 넣어놓아 접근을 한다.  

책의 예제에서는 2번 방식으로 접근했지만, 나는 매개변수 양이 얼마 되지 않으므로 1번 방식으로 접근하겠다.  

다음 2번을 처리하는데, 이는 이 리팩토링의 매력포인트이다.  

```c++
    double getFullFramePrice(string n);
    double getPrice(string n);
    double compactPrice(double d);
    double DSLR_Price(double d);
```
이 함수들은, `loanCamera()`를 수행하기 위해 존재한다. 따라서 이 함수들도 MoveMethod 방식으로 이동시킬 수 있다.  

그렇게 1,2번 방식을 적용시키면... 이렇게된다.   

```c++
class CameraLoan{
private:
    bool fullFrame;
    bool DSLR;

    LoanCameraStrategy loanCameraStrategy;

public:

    CameraLoan(LoanCameraStrategy loanCameraStrategy){
        this.loanCameraStrategy = loanCameraStrategy;
    }

    double loanCamera(string name, double duration){
       return this.loanCameraStrategy.loanCamera(name, duration, fullFrame, DSLR);    
    }
};


class LoanCameraStrategy{

public:
    double getFullFramePrice(string n);
    double getPrice(string n);
    double compactPrice(double d);
    double DSLR_Price(double d);

    double loanCamera(string name, double duration, bool fullFrame, bool DSLR){
        if (fullFrame && DSLR){ // Plan A
            return new LoanPlanA(getFullFramePrice(name) * DSLR_Price(duration));
        } else if (fullFrame && !DSLR) { // Plan B
            return new LoanPlanB(getFullFramePrice(name) * compactPrice(duration));
        } else if (!fullFrame && DSLR) { // Plan C
            return new LoanPlanC(getPrice(name) * DSLR_Price(duration));
        } else{ // Plan D
            return new LoanPlanD(getPrice(name) * compactPrice(duration));
        }
        return 0.0;  
    }
};



```

이렇게 하고... 기존에 있던 class들인 PlanA~D를 위한 Strategy Class도 만들어줘야 한다.  

```c++
class CameraLoan{
private:
    bool fullFrame;
    bool DSLR;

    LoanCameraStrategy loanCameraStrategy;

public:

    CameraLoan(LoanCameraStrategy loanCameraStrategy){
        this.loanCameraStrategy = loanCameraStrategy;
    }

    double loanCamera(string name, double duration){
       return this.loanCameraStrategy.loanCamera(name, duration, fullFrame, DSLR);    
    }
};

class LoanCameraStrategy{

public:
    double getFullFramePrice(string n);
    double getPrice(string n);
    double compactPrice(double d);
    double DSLR_Price(double d);

    virtual double loanCamera(string name, double duration);
};

class LoanPlanAStrategy : public LoanCameraStrategy{
    double loanCamera(string name, double duration){
        return getFullFramePrice(name) * DSLR_Price(duration);
    }
}

class LoanPlanBStrategy : public LoanCameraStrategy{
    double loanCamera(string name, double duration){
        return getFullFramePrice(name) * compactPrice(duration);
    }
}
// 이후 생략..

```

```c++
CameraLoan a = new CameraLoan(new LoanPlanAStrategy));
```

요약하자면 메인 Strategy 클래스를 추상화시킨점이 있다.  
또한 클라이언트가 Strategy를 지정한다.
아래 과정에서 조건문을 이용하여도 된다. 