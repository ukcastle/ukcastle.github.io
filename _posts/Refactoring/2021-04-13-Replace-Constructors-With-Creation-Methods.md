---
layout: post
title: Replace Constructors with Creation Methods (97)
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

#### 개요

어떤 클래스의 **인스턴스를 생성할 때** 그것이 제공하는 **여러 생성자** 중 어떤 것을 호출해야 할 지 결정하기 어렵다면, 인스턴스를 생성해 리턴하는 **생성 메서드**로 **각 생성자를 대체**하여 그 용도가 명확히 드러나도록 한다.  

```c++
class Loan{
public:
    Loan(commitment, riskRating, maturity);
    Loan(commitment, riskRating, expiry);
    Loan(commitment, outStanding, riskRating, maturity, expiry);
    Loan(capitalStrategy, commitment, riskRating, maturity, expiry);
    Loan(capitalStrategy, commitment, outStanding, riskRating, maturity, expiry);
}
```

```c++
class Loan{
private:
    Loan(capitalStrategy, commitment, outStanding, riskRating, maturity, expiry);
public:
    Loan createTermLoan(commitment, riskRating, maturity);
    Loan createTermLoan(capitalStrategy, commitment, outStanding, riskRating, maturity, expiry);
    Loan createRevolver(commitment, outStanding, riskRating, maturity, expiry);
    Loan createRevolver(capitalStrategy, commitment, riskRating, maturity, expiry);
    Loan creaRCTL(commitment, outStanding, riskRating, maturity, expiry);
    Loan creaRCTL(capitalStrategy, commitment, outStanding, riskRating, maturity, expiry);
}
```

#### 동기

어떤 언어에서는 클래스 이름에 상관 없이 생성자의 이름을 마음대로 정할 수 있다. 하지만 Java나 C++같은 언어에서는 이것이 허용되지 않는다. 각각의 생성자는 클래스의 이름과 같아야한다.  
생성자가 한개가 아니고 여러개라면, 프로그래머가 생성자 코드를 직접 살펴본 후 어떤 생성자를 호출할 지 선택해야 한다.  
이런 방식은 프로그래머가 해당 클래스를 사용할 때 더 오랜 시간이 걸릴 수 있으며, 위의 예시를 통하여 이를 극복할 수 있다.  
다만 생성자가 너무 많은 클래스를 발견했다면, 이 리팩터링을 적용하기 전에 먼저 [Extract Class](https://jo631.github.io/refactoring/2021/04/09/RefactoringToPattern/#extract-superclass) 또는 [Extract Subclass](https://jo631.github.io/refactoring/2021/04/09/RefactoringToPattern/#extract-subclass) 리팩터링을 고려하는 것이 좋다.  
클래스가 지나치게 많은 일을 하고 있다면, Extract Class 리팩터링이 좋은 선택이고, 클래스가 인스턴스 변수의 일부만 사용하고 있다면 Extract Subclass 리팩터링을 적용할 만 하다.   


#### 장점

- 그 용도를 생성자보다 **명확하게** 드러낼 수 있다.  
- 동일한 수와 타입의 파라미터를 받는 생성자를 두 개 이상 만들 수 없었던 **제한**이 사라진다.  
- 사용되지 않는 생성 코드를 찾기가 쉬워진다.  

#### 단점

- 객체를 생성할 때 표준이 아닌 방식을 사용한다. 어떤 클래스에 대해서는 `new`연산자를 사용하고, 또 어떤 클래스에 대해서는 생성 메서드를 통하게 된다.   

#### 절차

이 리팩터링을 시작하기 전에, *실질 생성자가 존재하는지 찾아본다. 실질 생성자가 없다면 Chain Constructors(448) 리팩터링을 적용해 하나 만든다.

*실질 생성자: 실질적인 생성 기능을 모두 구현하는 생성자로서, 다른 생성자들은 이 실질 생성자에게 작업을 위임하는 역할만 하는 경우를 뜻한다.

1. 여러 생성자 중 하나를 선택하여 그것을 호출하는 클라이언트 코드를 찾는다. 그리고 그 코드에 [Extract Method](https://jo631.github.io/refactoring/2021/04/09/RefactoringToPattern/#extract-method)를 적용해 별도의 메소드(`public static`로 지정)로 뽑아낸다. 이렇게 만든 메서드를 생성 메서드라 하고, [Move Method](https://jo631.github.io/refactoring/2021/04/09/RefactoringToPattern/#move-method)를 적용해 이 생성 메서드를 해당 생성자를 포함하고 있는 클래스로 옮긴다.  

2. 선택한 생성자를 사용하는 곳(즉 동일한 종류의 인스턴스를 사용하는 곳)을 모두 찾아 앞에서 만든 생성 메서드를 호출하도록 수정한다.  

3. 만약 선택한 생성자가 다른 생성자를 호출하고 있다면, 생성 메소드에서 선택한 생성자 대신 호출되는 생성자를 사용하도록 고친다. [Inline Method](https://jo631.github.io/refactoring/2021/04/09/RefactoringToPattern/#inline-method) 리팩터링을 적용할 때 처럼 생성자를 인라인화하면 된다.  

4. 생성 메서드로 바꾸고 싶은 다른 모든 생성자에 대해 단계 1~3을 반복한다.  

5. 클래스의 생성자가 해당 클래스 밖에서 더 이상 사용되지 않는다면, `private`하게 바꾼다.  


#### 변형  

객체 구분 파라미터를 사용하는 생성 메소드(생성 메소드에 파라미터를 추가해 그 값에 따라 생성할 객체의 종류를 정하는 메소드) 리팩토링을 적용하려고 생각하면서 머릿속으로 따져보니, 엄청난 수의 메소드가 필요할 수도 있다.  
이 때는, 유연하게 특수하게 사용되는 일부 생성자는 `public`하게 냅두고, 공용적으로 사용되는 생성자만 해당 방법을 따르게 하면 된다.  

#### Extract Factory (팩토리 추출)

생성 메소드가 지나치게 많으면, 클래스의 주요 책임이 잘 드러나지 않을 수 있다. 취향 차이지만 이 경우 리팩토링을 통해 생성 메소드를 하나의 팩토리로 모으면 된다.  

```c++
class Loan{
private:
    Loan(capitalStrategy, commitment, outStanding, riskRating, maturity, expiry);
public:
    // constructor
    Loan createTermLoan(commitment, riskRating, maturity);
    Loan createTermLoan(capitalStrategy, commitment, outStanding, riskRating, maturity, expiry);
    Loan createRevolver(commitment, outStanding, riskRating, maturity, expiry);
    Loan createRevolver(capitalStrategy, commitment, riskRating, maturity, expiry);
    Loan creaRCTL(commitment, outStanding, riskRating, maturity, expiry);
    Loan creaRCTL(capitalStrategy, commitment, outStanding, riskRating, maturity, expiry);

    // method
    void calcCapital();
    void calcIncome();
    void calcROC();
    void setOustanding();
}
```

```c++
class LoanFactory{
public:
    static Loan createTermLoan(commitment, riskRating, maturity);
    static Loan createTermLoan(capitalStrategy, commitment, outStanding, riskRating, maturity, expiry);
    static Loan createRevolver(commitment, outStanding, riskRating, maturity, expiry);
    static Loan createRevolver(capitalStrategy, commitment, riskRating, maturity, expiry);
    static Loan creaRCTL(commitment, outStanding, riskRating, maturity, expiry);
    static Loan creaRCTL(capitalStrategy, commitment, outStanding, riskRating, maturity, expiry);
}

class Loan{
private:
public:
    Loan(capitalStrategy, commitment, outStanding, riskRating, maturity, expiry);
    void calcCapital();
    void calcIncome();
    void calcROC();
    void setOustanding();
}

int main(){
    Loan termLoan = LoanFactory.createTermLoan(1,2,3);
    termLoan.calcCapital();
}
```