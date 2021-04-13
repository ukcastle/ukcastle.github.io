---
layout: post
title: Decorator
category: DesignPattern
tag: [DesignPattern] 
---

#### 개요

이에 대해 PPT도 만들어서 발표해서, 지식이 꽤나 있다.  

쉽게말하여, **토핑 추가**의 개념이다.  

**객체의 결합을 통해 기능을 동적으로 유연하게 확장**하는 패턴이다.  

#### 사례 

도로를 표시하는 방법을 정의할 떄

```java
// 기본 도로 표시 클래스
public class RoadDisplay {
    public void draw() { System.out.println("기본 도로 표시"); }
}
// 기본 도로 표시 + 차선 표시 클래스
public class RoadDisplayWithLane extends RoadDisplay {
  public void draw() {
      super.draw(); // 상위 클래스, 즉 RoadDisplay 클래스의 draw 메서드를 호출해서 기본 도로 표시
      drawLane(); // 추가적으로 차선 표시
  }
  private void drawLane() { System.out.println("차선 표시"); }
}
public class Client {
  public static void main(String[] args) {
      RoadDisplay road = new RoadDisplay();
      road.draw(); // 기본 도로만 표시

      RoadDisplay roadWithLane = new RoadDisplayWithLane();
      roadWithLane.draw(); // 기본 도로 표시 + 차선 표시
  }
}
// 출처 : https://gmlwjd9405.github.io/2018/07/09/decorator-pattern.html
```

이렇게 하나하나 상속을 통해 표시한다면 문제점이 생긴다.  
만약 신호등이 있는지 없는지도 표시하려면, 또 새로운 클래스를 상속해야한다.  
또 다른 옵션을 붙이고싶으면 또 상속해야한다.  
이렇게 상속되는 클래스의 가짓수가 2^n개로 증가하게 된다.  


#### 해결책

클래스 내에서 추상화된 기본 클래스를 변수로 삼아 멤버로 가진다.  
향후 줄줄이 소세지처럼 생성할 때 생성자에 이전 소세지의 주소를 넘겨주면 된다.  

```java
public abstract class Display { public abstract void draw(); }

/* 기본 도로 표시 클래스 */
public class RoadDisplay extends Display {
  @Override
  public void draw() { System.out.println("기본 도로 표시"); }
}

/* 차선 표시를 추가하는 클래스 */
public class LaneDecorator extends DisplayDecorator {
  // 기존 표시 클래스의 설정
  public LaneDecorator(Display decoratedDisplay) { super(decoratedDisplay); }
  @Override
  public void draw() {
      super.draw(); // 설정된 기존 표시 기능을 수행
      drawLane(); // 추가적으로 차선을 표시
  }
  // 차선 표시 기능만 직접 제공
  private void drawLane() { System.out.println("\t차선 표시"); }
}
/* 교통량 표시를 추가하는 클래스 */
public class TrafficDecorator extends DisplayDecorator {
  // 기존 표시 클래스의 설정
  public TrafficDecorator(Display decoratedDisplay) { super(decoratedDisplay); }
  @Override
  public void draw() {
      super.draw(); // 설정된 기존 표시 기능을 수행
      drawTraffic(); // 추가적으로 교통량을 표시
  }
  // 교통량 표시 기능만 직접 제공
  private void drawTraffic() { System.out.println("\t교통량 표시"); }
}

public class Client {
  public static void main(String[] args) {
      Display road = new RoadDisplay();
      road.draw(); // 기본 도로 표시
      Display roadWithLane = new LaneDecorator(new RoadDisplay());
      roadWithLane.draw(); // 기본 도로 표시 + 차선 표시
      Display roadWithTraffic = new TrafficDecorator(new RoadDisplay());
      roadWithTraffic.draw(); // 기본 도로 표시 + 교통량 표시
  }
}
// 출처 : https://gmlwjd9405.github.io/2018/07/09/decorator-pattern.html

```

이런 방식으로 진행되고, 이는 추후 팩토리패턴과 결합하면 클라이언트에서 `new`를 사용하지 않고도 더 강력하게 사용할 수 있다.  