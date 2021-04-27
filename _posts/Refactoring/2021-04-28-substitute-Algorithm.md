---
layout: post
title: Substitute Algorithm
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

#### 개요

>알고리즘을 더욱 명료하게 하나로 변경하고 싶으면
메소드의 내용을 새로운 알고리즘으로 변경하시오.

<br>

#### 예시

```java
String foundPerson(String[] people){ 
    for (int i = 0; i < people.length; i++) { 
        if (people[i].equals ("Don")){ 
            return "Don";
        } 
        if (people[i].equals ("John")){ 
            return "John"; 
        } 
        if (people[i].equals ("Kent")){ 
            return "Kent"; 
        } 
    } 
    return ""; 
}
```

```java

String foundPerson(String[] people){ 
    List candidates = Arrays.asList(new String[] {"Don", "John", "Kent"}); 
    for (int i=0; i<people.length; i++) 
        if (candidates.contains(people[i])) 
            return people[i]; 

    return ""; 
}

//출처: https://arisu1000.tistory.com/27669 [아리수]
```