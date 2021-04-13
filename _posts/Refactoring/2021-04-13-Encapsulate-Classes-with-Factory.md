---
layout: post
title: Encapsulate-Classes-with-Factory (124)
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

#### 개요

클라이언트가 한 패키지 내의, 공통 인터페이스를 가지는 클래스들의 인스턴스를 직접 생성하고 있다면, 그 클래스의 생성자를 클라이언트가 직접 볼 수 없게 바꾸고 클라이언트는 팩토리를 통해 그 인스턴스를 얻도록 한다.  

```c++
class AttributeDescriptor{
protected:
    AttributeDescriptor(...);
public:
    virtual void printStatus()=0;
};

class BooleanDescriptor : public AttributeDescriptor{
private:
    int status;
public:
    BooleanDescriptor(int s) : AttributeDescriptor(){
        this.status = s;
    }
    void printStatus(...)
};

class DefaultDescriptor : public AttributeDescriptor{
public:
    DefaultDescriptor() : AttributeDescriptor(){}

    void printStatus(...)
};
```

```c++
class AttributeDescriptor{
public:
    AttributeDescriptor(){}
    static AttributeDescriptor* forBoolean(int s){
        return new BooleanDescriptor(s);
    }
    static AttributeDescriptor* forDefault(){
        return new DefaultDescriptor();
    }

    virtual void printStatus()=0;
};

class BooleanDescriptor : public AttributeDescriptor{
private:
    int status;
protected:
    friend class AttributeDescriptor;
    BooleanDescriptor(int s) : AttributeDescriptor(){
        this->status = s;
    }
public:
    void printStatus() override {  }
};

class DefaultDescriptor : public AttributeDescriptor{
protected:
    friend class AttributeDescriptor;
    DefaultDescriptor() : AttributeDescriptor(){}
public:
    void printStatus() override {  }
};

int main(){
    AttributeDescriptor* a = AttributeDescriptor::forDefault();
    AttributeDescriptor* b = AttributeDescriptor::forBoolean(3);
}

```

이번엔 직접 코딩해봤는데, 예제가 자바 형식으로 돼 있어서 좀 고생을 많이했다.  
같은 패키지 안에 있으면 protected 취급이 되는데, c++은 그게 안된다.  
그래서 `friend class` 키워드로 접근을 허가해줬다.  
또 따로 헤더파일에 선언해줘야 하는데... 여기까지만 하자.. 자바로할껄..  

#### 동기

클라이언트가 사용할 객체의 클래스를 직접적으로 알아야 한다면, 첫 번째 방식대로 해도 좋다. 하지만 직접 알아야 할 필요가 없다면? 그리고 같은 패키지 안에 있고 조건이 변할 일도 없다면 어떻게 할까? 그렇다면 인스턴스 생성을 팩토리에 맡기고 그 클래스 자체를 정보 은닉을 하는 효과를 가질 수 있다.  

#### 장점

- 용도를 쉽게 알아볼 수 있는 생성 메서드를 제공하여, 클라이언트가 원하는 종류의 객체를 쉽게 생성할 수 있도록 한다.  
- 공개될 필요가 없는 클래스들을 숨겨 패키지의 **개념적 무게**를 줄인다  
    >클라이언트가 굳이 알아야 하지 않는 정보는 알 필요가 없다.  
- 클라이언트가 **구현에 대해서가 아닌 인터페이스에 대한 프로그래밍**을 하게 된다.  

#### 단점

- 새로운 종류의 객체가 필요할 경우 생성 메서드를 추가하거나 수정해야 한다.  

    > 이 리팩토링의 문제는 종속성에 있다.  
    새로운 생성자를 만들어야 하는 일이 잦으면, 이 리팩토링을 포기할 수도 있다.  

- 팩토리의 소스 코드가 아닌 바이너리만 배포할 경우 클라이언트가 쉽게 수정할 수 없게 된다.  

    >이런 방법을 탈피하기 위해서 고유 기능을 가지는 동시에 팩토리의 역할까지 하게 될 수도 있는데, 이렇게 한 클래스에 여러 책임을 부과하는 것에 거부감을 가질 수 있다. 이땐 [Extract Factory](https://jo631.github.io/refactoring/2021/04/13/Replace-Constructors-With-Creation-Methods/#extract-factory-%ED%8C%A9%ED%86%A0%EB%A6%AC-%EC%B6%94%EC%B6%9C)를 고려해봄직 하다.  

#### 절차  

어떤 클래스들이 하나의 인터페이스를 공유하거나, 같은 부모클래스를 가지면서 같은 패키지에 있을 때 이 리팩토링이 필요할 수도 있다. 이런 클래스들을 **대상 클래스** 라고 부르자.  

1. 대상 클래스 중 하나를 선택하고 그 생성자 중 하나를 골라 그 생성자를 호출하는 클라이언트 코드를 찾는다. [Extract Method](https://jo631.github.io/refactoring/2021/04/09/RefactoringToPattern/#extract-method) 리팩토링을 통해 그 코드를 `public static` 메소드로 만든다. 그 메소드가 **생성 메소드**다.  
그 다음 그 메소드를 [Move Method](https://jo631.github.io/refactoring/2021/04/09/RefactoringToPattern/#move-method)를 통해 대상의 부모클래스로 옮긴다.  

2. 앞에서 선택한 생성자를 호출하는 곳 중 단계 1에서 만든 생성 메소드와 **같은 종류**의 객체를 생성하는 코드를 모두 찾아 생성 메소드를 호출하도록 수정한다.  

3. 앞에서 선택한 생성자로 생성할 수 있는 모든 종류의 객체에 대해서도 1,2를 반복한다.  

4. 앞에서 선택한 생성자의 접근 지정자를 public 이외의 것으로 바꿔 클라이언트로부터 **숨긴다**.  

5. 나머지 대상 클래스에 대해 1~4를 반복한다.  


#### 내부 클래스의 캡슐화
`java.util.Collections`클래스는 생성 메소드를 가진 클래스를 캡슐화하는 것이 어떤것인가를 보여주는 훌륭한 예제이다.  
이 클래스를 사용하는 프로그래머가 컬렉션 객체를 수정 불가 또는 동기화 상태로 만들 수 있는 기능을 제공하기 위해 [Proxy](https://jo631.github.io/refactoring/2021/04/13/ProxyPattern/) 패턴을 도입했다. 게다가 그 프록시 클래스를 `public`으로 만들어 프로그래머가 자신의 컬렉션 객체를 직접 보호하지 않고 Collections 클래스의 내부 클래스로 정의한 다음 Collections 클래스에 생성 메서드를 추가해 프로그래머 자신이 필요한 프록시를 얻는 방법을 제공했다.  
`java.util.Collections`의 내부 클래스들도 상속 구조를 이루고 있는데, 각각의 내부 클래스에 컬렉션을 받아 수정할 수 없거나 동기화하도록 보호한 다음 그 객체를 List나 Set과 같은 일반적인 인터페이스 타입으로 리턴하는 방식을 사용한다.  
그 결과 프로그래머가 알아야 할 클래스 수는 늘리지 않으면서 필요한 기능을 제공한 결과가 되었다. 이는 팩토리의 좋은 예 이기도 한다.  