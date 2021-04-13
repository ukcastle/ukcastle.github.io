---
layout: post
title: Proxy Pattern
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

Refactoring 책을 읽으면서 설명이 부족했던 디자인패턴들을 공부하는 공간이다.  

#### 개요

프록시는 **대변인, 대리자** 라는 의미이다.  
프로그램 내에서 봤을때도 똑같이, 누군가에게 어떤 일을 대신 시키는 것이다.  
중요한 것은 **흐름제어만 할 뿐 결괏값을 조작하거나 변경하면 안된다**라는 것이다.  

```java
Interface IImage{
    public void displayImage();
}

class RealImage implements IImage {
    private String filename;
    public RealImage(String filename) {
        this.filename = filename;
        loadImageFromDisk();
    }

    private void loadImageFromDisk() {
        System.out.println("Loading   " + filename);
        //오랜 시간...
    }

    @Override
    public void displayImage() {
        System.out.println("Displaying " + filename);
    }
}

class ProxyImage implements IImage {
    private String filename;
    private IImage image;

    public ProxyImage(String filename) {
        this.filename = filename;
    }

    @Override
    public void displayImage() {
        if (image == null)
           image = new RealImage(filename);

        image.displayImage();
    }
}

class ProxyExample {
    public static void main(String[] args) {
        IImage image1 = new ProxyImage("HiRes_10MB_Photo1");
        IImage image2 = new ProxyImage("HiRes_10MB_Photo2");

        image1.displayImage(); // loading necessary
        image2.displayImage(); // loading necessary
    }
}
```

#### 동기 

처음 배우면서, 이걸 왜 써? 라고 생각했다.  

일단, 객체지향의 SOLID 원칙 중 **OCP**, **DIP** 원칙이 스며들어있다.  
*OCP: 개방 폐쇠 원칙. 자신의 확장엔 열려있고, 주변의 변화엔 닫혀 있어야 한다.  
*DIP: 의존 역전 원칙. 고수준 모듈은 저수준 모듈에 의존하면 안되고, 구체적인 것이 추상화된 것에 의존해야 한다.  
[여기](https://limkydev.tistory.com/77)를 검색해가면서 공부했다.  

또한 프록시에도 여러 종류가 있는데, 후술하겠다.  
 

#### 특징

- 원래 하려던 기능을 수행하며 그외의 부가적인 작업(로깅, 인증, 네트워크 통신 등)을 수행하기에 좋다.  
- 비용이 많이 드는 연산(DB 쿼리, 대용량 텍스트 파일 등)을 실제로 필요한 시점에 수행할 수 있다.  
- 사용자 입장에서는 프록시 객체나 실제 객체나 사용법은 유사하므로 사용성이 좋다.  
- 또한 프로그래머 입장에서 자신이 객체를 다루고있는지, 프록시를 다루고 있는지 모른다. (같은 인터페이스를 구현하고 있기 때문에) 

#### 가상 프록시

예를들어, 생성하는 데 **많은 비용**이 드는 반면 사용률은 별로 안되는 클래스가 있다고 치자.  
이런 클래스의 수명이 프로그램 시작부터 끝까지 간다고 생각하면? 매우 잘못된 설계일 것이다.  
이 때 가상 프록시를 사용하여 **늦은 초기화** 기능을 제공할 수 있다. 프록시 클래스만 만들어 놓고, 직접 사용할 떄만 클래스를 호출하는 방식으로.  

#### 보호 프록시  

객체에 따른 접근 권한을 제어하는 방식이다.  
예를들어 내 개인정보를 열람할 수 있는 코드가 있을 때, 아무나 이 코드를 접근해도 될까?  

#### 방화벽 프록시  

네트워크 자원과 같은 **파이**가 정해진 자원에 접근을 제어함으로써 파이를 독식하려는 **나쁜** 클라이언트로부터 보호하는 방식

#### 스마트 레퍼런스 프록시

주 객체(realClass)가 참조될 때 추가적인 행동을 제공하는 프록시  

#### 캐싱 프록시  

연산 비용이 많이드는 작업의 결과를 일시적으로 저장한 뒤, 추후에 실제 작업 대신 보여주는 프록시  

#### 동기화 프록시

여러 스레드에서 주 객체에 접근을 하는 경우 안전하게 처리할 수 있도록 도와주는 프록시, 자바 스페이스에서 사용된다.  

#### 복잡도 숨김 프록시  

복잡한 클래스들의 집합에 대한 접근을 제어하고 복잡도를 은폐한다.  

#### 지연 복사 프록시  

클라이언트에서 필요로 할 때 까지 객체가 복사되는 것을 지연시킴으로써 객체의 복사를 제어한다.  
Java5의 CopyOnWriteArrayList에서 사용된다.  