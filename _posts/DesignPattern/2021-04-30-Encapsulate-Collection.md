---
layout: post
title: Encapsulate Collection
category: DesignPattern
tag: [DesignPattern] 
---

#### 개요

>메소드가 리스트를 반환한다면
읽기전용 뷰를 반환하도록 만들고 해당 리스트를 추가/제거하는 메소드를 제공하게 한다.  


#### 예제

```java
class Person{
    int[] list;

    int[] getList(){ return this.list; }
    void setList(int[] list){ this.list = list;} 
}
```

```java
class Person{
    int[] list;

    final int[] getList(){ return this.list; }
    addList(int num){ ... }
    deleteList(int num){ ... } 
}
```

