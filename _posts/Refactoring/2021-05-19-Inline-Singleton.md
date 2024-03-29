---
layout: post
title: Inline Singleton (168)
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

#### 개요

> 코드의 여러 곳에서 접근할 수 있어야 하지만 전역적인 필요까지는 없는 객체가 싱글톤으로 구현되어 있다면
싱글톤 객체를 저장하고 그에 대한 접근 경로를 제공하는 클래스로 싱글톤의 기능을 옮긴다. 그리고 제거한다.  

<br>

#### 동기

싱글톤 패턴의 의도는 '어떤 클래스의 인스턴스를 단 하나만 허용하고, 그에 대한 전역적인 접근을 가능하게 하려는 것' 이다. 싱글톤 패턴에 대한 생각이 뼈 속 깊이 침투해 있어 다른 패턴 또는 더 단순한 설계보다 싱글톤이 좋다고 믿기 시작해 지나치게 많은 싱글톤을 만들고 있다면, 싱글톤에 중독된 것이다.  
Inline Singleton은 이런 중독으로부터 해방될 수 있는 유용한 수단이다.  
**싱글톤이 불필요한 상황은 언제인가?** 짧게 말하면, 대부분의 상황에서 불필요하다.  
조금 길게 답하면, 어떤 객체에 전역적인 접근이 가능하도록 만드는 것 보다 필요한 곳에 그 참조를 넘겨주는 것이 더 간단한 상황에서는, 싱글톤이 불필요하다.  
**전역 데이터에 대해서는 유죄 추정의 원칙이 적용된다.** 정말 필요한지 증명하기 전 까지는 그 필요성을 의심해야 한다.  

##### 장점

- 객체 간의 협력 관계를 좀 더 명확하게 만든다.
- 싱글턴 객체를 보호하기 위한 특수 코드가 필요 없어진다.  

##### 단점

- 객체의 참조를 호출 트리의 여러 계층에 넘겨야 해서 불편하고 힘들어졌다면, 설계를 좀 더 복잡하게 만드는 것이다.  

<br>

#### 절차

흡수 클래스 라는 용어는 기존의 싱글턴이 맡고 있던 역할을 대신할 클래스를 지칭한다.  

1. 싱글턴이 구현하고 있는 `public` 메서드를 흡수 클래스에 선언한다. 그리고 이 새 메서드의 구현은 기존의 싱글톤에 위임하도록 한다. 이 때 그 메서드 중 `static` 메서드가 있다면, 흡수 클래스에 그에 **대응하는 메서드를 선언할 때 `static` 키워드를 제거**한다.  
만약 기존의 싱글톤 클래스를 흡수 클래스로 삼을 생각이라면, static 메서드를 그대로 놔두어도 무방하다.  

2. 클라이언트 코드에서 싱글톤을 참조하는 부분을 모두 흡수 클래스를 참조하도록 수정한다.  

3. 싱글톤에 아무 기능도 남아있지 않도록 Move Method와 Move Field를 적용해 모든 기능을 흡수 클래스로 옮긴다.

4. 싱글톤을 제거한다.  
<br>