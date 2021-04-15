---
layout: post
title: 패턴을 활용한 리팩토링
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

현재 [패턴을 활용한 리팩토링](http://www.yes24.com/Product/Goods/14752528)이라는 책을 읽고있다.  

가볍게 읽어보려고 대출했는데, 읽으면서 감탄이 나와서 독후감, 요점 정리 노트 형식으로 써보려고 한다.  

#### 목적  

이 책의 목적은 다음 사항을 돕는 것이다.  

- 리팩터링과 패턴을 어떻게 결합하는지 이해
- 패턴을 고려한 리팩터링으로 기존 코드의 설계를 개선
- 코드에서 패턴을 고려한 리팩터링이 필요한 영역을 식별
- 새로운 설계의 초기부터 패턴을 사용하는 것 보다 기존 코드를 개선하기 위해 패턴을 사용하는 것이 좋은 이유를 이해  

#### 배경지식

이 책은 [Refactoring](http://www.yes24.com/Product/Goods/7951038)을 보완재로 이용하고 있다.  
아래와 같은 개념이 필요하다.  
- Extract Method
- Extract Interface
- Extract Superclass
- Extract Subclass
- Pull Up Method
- Move Method
- Rename Method  
<br>
- Replace Inheritance with Delegation
- Replace COnditional with Polymorphism
- Replace Type Code with Subclasses  

들어가기 전에 해당 개념들을 하나하나씩 공부해 볼 예정이다.  


#### Extract Method

메소드를 정리하는 기술이다.  

**그룹으로 함께 묶을 수 있는 코드 조각이 있으면  
코드의 목적이 잘 드러나도록 메소드의 이름을 지어 별도의 메소드로 뽑아낸다.**

```java
void printOwing(double amount)
{
  printBanner();
  //상세 정보 표시
  System.out.println( "name:" + _name );
  System.out.println( "amount:" + amount );
}

===========================================

void printOwing(double amount)
{
  printBanner();
  printDetails(amount);
}
void printDetails(double amount)
{
  System.out.println( "name:" + _name );
  System.out.println( "amount:" + amount );
}
```

**사용 동기**  
1. 메소드가 잘 쪼개어져 있으면 다른 메소드에서도 활용 할 확률이 높아진다.  
2. 고수준의 메소드(아래의 printOwing())를 읽을 때 일련의 주석을 읽는 것 같은 효과를 준다.  

**절차**  
- 메소드를 새로 만들고 의도를 잘 나타낼 수 있도록 이름을 정한다.  
    - 뽑아내고자 하는 부분이 한두줄의 메세지나 함수 호출과 같은 간단한 경우, **새로운 메소드의 이름**이 그 코드의 **의도**를 더 잘 나타낼 수 있을때만 뽑아낸다.  
- 코드를 뽑아낸 뒤 다음과 같은 과정을 거친다.  
    - 뽑아내기 전 함수에서 사용되는 매개변수가 있을 시 뽑아낸 함수의 매개변수로 이용된다.
    - 뽑아낸 함수에서만 사용되는 변수가 있으면, 지역 변수로 사용한다.  
    - 뽑아낸 함수 안에서 변수의 값이 변경되는지 본다. 값이 이상하게 남으면 **Split Temporary Variable** 이나 **Replace Temp with Query** 방식을 이용한다.  


#### Extract Interface

몇몇 클라이언트가 클래스 인터페이스의 같은 부분집합을 사용하거나 두개의 클래스들이 인터페이스의 일부가 중복된다면 해당 부분을 인터페이스로 추출한다.

``` c++
class Employee {
    void getRate();
    void hasSpecialSkill();
    void getName();
    void getDepartment();
}
```

``` c++
class Ibillable{
public:
    virtual void getRate();
    virtual void hasSpecialSkill();
}

class Employee : public Ibillable{
    void getRate();
    void hasSpecialSkill();
    void getName();
    void getDepartment();
}
```


#### Extract Superclass
비슷한 기능을 가진 두개의 클래스들을 가지고 있다면 SuperClass를 만들고, 공통되는 기능을 새로 생성한 SuperClass로 이동한다.

``` python
class Department :
  def totalAnnualCost()
  def name()
  def headCount()
}

class Employee :
  def annualCost()
  def name()
  def id()
}
```

``` python
class Party :
  def name()
  def annualCost()
}

class Department(Party) :
  def annualCost()
  def headCount()
}

class Employee(Party) :
  def annualCost()
  def id()
}
```

#### Extract Subclass

클래스가 특정 인스턴스에서만 사용되는 기능들을 가지고 있다면 해당 기능들로 이루어진 서브클래스를 만들어라.

```python
def createEmployee(name, type):
    return Employee(name, type)
```

```python
def createEmployee(name, type):
    switch (type) :
        case "engineer": return Engineer(name)
        case "salesman": return Salesman(name)
        case "manager":  return Manager (name)
  
```

#### Extract Class

두 개의 클래스가 해야 할 일을 하나의 클래스가 하고 있는 경우  
새로운 클래스를 만들어서 관련 있는 필드와 메소드를 예전 클래스에서 새로운 클래스로 옮겨라


#### Pull Up Method

서브클래스들에 같은 결과를 반환하는 메소드들을 가지고 있다면 해당 메소드들을 슈퍼클래스로 옮기시오.

``` c++
class Employee {...}

class Salesman : public Employee {
  string name() {...}
}

class Engineer : public Employee {
  string name() {...}
}
```

``` c++
class Employee{
public:
    string name() {...}
}

class Salesmans : public Employee { ... }
class Engineer : public Employee { ... }
```

#### Move Method

메소드가 정의된 클래스보다 다른 클래스에서 더 많이 사용된다면 해당 클래스로 메소드를 옮기고, 이전 메소드에는 단순 대리자를 만들거나 전체에서 제거하시오.

#### Inline Method

메소드 몸체가 메소드의 이름 만큼이나 명확할 때는 호출하는 곳에 메소드의 몸체를 넣고, 메소드를 삭제하라

#### Rename Method  

메소드의 이름이 메소드의 목적을 나타내지 못한다면 메소드의 이름을 변경하시오.

#### Replace Inheritance with Delegation

서브클래스가 슈퍼클래스 인터페이스의 일부만 사용하거나 상속된 데이터를 사용하지 않는다면 슈퍼클래스를 위한 변수를 생성하고, 슈퍼클래스를 대신하기위한 메소드를 적용하고 서브클래스화를 제거하시오.

```c++
class List {...}
class Stack : public List {...}
```

```c++
class List {...}
class Stack {
private:
    List* l;
public:
    Stack(List* l){
        this.l = l;
    }
}
```

#### Replace COnditional with Polymorphism

객체의 타입에 따라서 다른 행동을 하는 조건이 있다면 각 조건의 내용을 서브클래스의 메소드로 오버라이딩하도록 옮기고, 원본 메소드는 추상메소드로 변경하시오.


```c++
double getSpeed() { 
    switch (_type) { 
    case EUROPEAN: 
        return getBaseSpeed(); 
    case AFRICAN: 
        return getBaseSpeed() - getLoadFactor() * _numberOfCoconuts;
    } 
```

```c++
class Bird{
    virtual int getSpeed();
}

class European : public Bird{
    int getSpeed(){...}
}

class African : public Bird{
    int getSpeed(){...}
}
```

#### Replace Type Code with Subclasses  

클래스의 행동에 영향을 미치면서 변경할수 없는 타입코드가 있다면 해당 타입코드들을 서브클래스로 변경하시오.

```c++
Employee* createEmployee(name, type) {
  return new Employee(name, type);
}
```

```c++
Employee* createEmployee(name, type) {
  switch (type) {
    case "engineer": return new Engineer(name);
    case "salesman": return new Salesman(name);
    case "manager":  return new Manager (name);
  }
```

#### Replace Type Code with State/Strategy

클래스의 행동에 영향을 미치는 타입 코드가 있지만, 서브 클래스화 할 수 없다면 해당 타입코드를 State/Strategy 객체로 변경하라

