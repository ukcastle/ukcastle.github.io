---
layout: post
title: Composite Pattern
category: DesignPattern
tag: [DesignPattern] 
---

#### 개요

Composite 패턴이란, 객체들의 관계를 트리구조로 구성해 부분-전체 계층을 표현하는 패턴으로, 사용자가 단일 객체와 복합 객체 모두 동일하게 다루도록 하는 방법이다.  

```java

interface Component{
    public operation();
}

class File implements Componenet{
    public operation(){
        print("operaion !");
    }
}

class Composite implements Component{
    private List<File> files = new ArrayList<Component>();

    public operation(){
        for (File i : self.files){
            i.operation();
        }
    }

    public add(File file){
        self.files.add(file);
    }   

    public remove(File file){
        self.files.remove(file);
    }
}
```

가장 간단하게, 이런 구조다.  

`Composite`과 `File`는 `Component`를 구현하며  
`File`에서는 인터페이스인 `Component`를 실질적으로 구현하며  
`Composite`는 `File`들을 관리하는 구조이고 `File`의 함수를 호출하는 식으로 구현한다.  
또한 `Composite`는 File들을 **추가하고 삭제하는 등 관리**하는 함수들도 있다.  

#### 용도

복합 객체와 단일 객체의 처리 방법이 다르지 않을 경우  
부분-전체 계층으로 나타낼 수 있다.  
이런 형태를 가장 가장 잘 나타내는 부분은 **Directory-File** 시스템이다.  

정리하면 
1. 부분-전체 관계를 트리 구조로 표현하고 싶을 때
2. 부분-전체 관계를 클라이언트에서 동일하게 처리하고 싶을 때


#### 예시  

파일과 폴더는, 물론 공유하지 않는 행동도 있지만, 둘 다 가지고있는 기능이 있다.  
예를들어 폴더를 삭제하면, 폴더만 삭제되는가? 혹은 폴더를 이동하면, 폴더만 이동하는가?  
아니다. 안에 있는 파일들도 삭제되거나 이동된다.   
이를 코드로, 간단한 예제로 만들어보겠다.  


```java
interface Component{
    public delete();
    public move();
}

class File implements Component{
    public delete(){
        System.out.println("File Deleted");
    }

    public move(){
        System.out.println("File Moved");        
    }
}

class Directory implements Component{
    private List<File> files = new ArrayList<Component>();

    public move(){
        System.out.println("Directory Moved!");
        for (File i : self.files){
            i.move();
        }
    }

    public delete(){
        for (File i : self.files){
            i.delete();
        }
        System.out.println("Directory Deleted!");
    }

    public add(Component component){
        self.files.add(component);
    }   

    public remove(Component component){
        self.files.remove(component);
    }

    public moveAnotherDirectory(Component component, Directory to){
        to.add(component);
        self.remove(component);
    }
}
```

예를들어, 앨범1 폴더에는 사진1,2,3이 있다.  
앨범 2 폴더에는 사진 4,5,6이 있다.  

트리구조로 나타내면,  
```
앨범1
ㄴ 사진1
ㄴ 사진2
ㄴ 사진3

앨범2
ㄴ 사진4
ㄴ 사진5
ㄴ 사진6
```

이를 코드로 나타내면 이렇다.

```java
Directory d1 = Directory("앨범1");
Directory d2 = Directory("앨범2");

File p1 = File("사진1");
File p2 = File("사진2");
File p3 = File("사진3");
File p4 = File("사진4");
File p5 = File("사진5");
File p6 = File("사진6");

d1.add(p1);
d1.add(p2);
d1.add(p3);

d2.add(p4);
d2.add(p5);
d2.add(p6);
```

여기서 d2를 삭제한다면? 

```java
d2.delete();

/* 출력
File4 Deleted!
File5 Deleted!
File6 Deleted!
Directory2 Deleted!
```

또한, d1 안에 d2를 집어넣을수도 있고, d2안의 파일을 d1으로 보낼수도있다.  

```java
d1.add(d2);
d2.moveAnotherDirectory(p4,d1);
d1.delete();
```

2번째 줄까지 진행을 한다면 트리 구조는 이렇게 될 것이다.  

```
앨범1
ㄴ 앨범2
    ㄴ 사진5
    ㄴ 사진6
ㄴ 사진1
ㄴ 사진2
ㄴ 사진3
ㄴ 사진4
```

마지막으로 3번째 코드를 작동한다면, 리프노드부터 순서대로 실행이 된다.  

```
File5 Deleted!
File6 Deleted!
Directory2 Deleted!
File1 Deleted!
File2 Deleted!
File3 Deleted!
File4 Deleted!
Directory1 Deleted!
```

