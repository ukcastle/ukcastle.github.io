---
layout: post
title: Replace Constructors with Creation Methods (448)
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

#### 개요

중복된 코드를 갖는 생성자가 여러 개 있다면, 중복을 최소화하기 위해 생성자들이 서로 호출하게 된다.  

```c++
class Loan{
public:
    Loan(commitment, riskRating, expiry){
        ...
        this.commitment = commitment;
        this.riskRating = riskRating;
        this.expiry = expiry;
        ...
    }
    Loan(commitment, outStanding, riskRating, maturity, expiry){
        ...
        this.commitment = commitment;
        this.outStanding = outstanding;
        this.riskRating = riskRating;
        this.maturity = maturity;
        this.expiry = expiry;
        ...
    }
    Loan(capitalStrategy, commitment, outStanding, riskRating, maturity, expiry){
        this.capitalStrategy = capitalStrategy;
        this.commitment = commitment;
        this.outStanding = outstanding;
        this.riskRating = riskRating;
        this.maturity = maturity;
        this.expiry = expiry;
    }
}
```

```c++
class Loan{
public:
    Loan(commitment, riskRating, expiry){
        this(null, commitment, null, riskRating, null ,expiry);
    }
    Loan(commitment, outStanding, riskRating, maturity, expiry){
        this(null, commitment, outStanding, riskRating, maturity, expiry);
    }
    Loan(capitalStrategy, commitment, outStanding, riskRating, maturity, expiry){
        this.capitalStrategy = capitalStrategy;
        this.commitment = commitment;
        this.outStanding = outstanding;
        this.riskRating = riskRating;
        this.maturity = maturity;
        this.expiry = expiry;
    }
}
```

#### 동기

생성자에 중복된 코드는 문제가 일어날 소지가 많다.  
만약 한 개의 생성자만 코드를 수정하고 지나갔다면? **으악!** 아마 다른 결함을 생성할 것이다.  
따라서 우리는 중복 코드를 줄이거나 제거할 수 있는 좋은 아이디어가 필요하다.  

이 문제는 생성자 체인을 통해 해결할 수 있다. 특수 목적의 생성자가 더 일반적인 생성자를 호출하도록 생성자 전체를 수정하여 체인을 형성하는 것이다.  
각 체인이 결국 한 생성자로 연결된다면, 이 생성자가 모든 생성자 호출을 실질적으로 처리하는 것 이므로 이를 **실질 생성자** 라고 한다. 실질 생성자가 다른 생성자보다 **많은 파라미터**를 갖는 것이 보통이다.  

#### 절차

1. 중복된 코드를 갖는 두 생성자를 찾는다. 그 중 하나가 다른 하나를 호출하게 하는 것이 어떻게 중복된 코드를 **안전하게**(그리고 더 쉽게) 제거할 수 있을지 생각해본다.  
2. 이미 작업한 생성자를 포함해 클래스 내의 모든 생성자에 대해 단계 1을 반복한다. 이렇게 모든 생성자에 대해 대해 중복을 가능한 한 적게 만든다.  
3. `public` 으로 남겨둘 필요가 없는 생성자의 접근 지정자를 변경한다.  

