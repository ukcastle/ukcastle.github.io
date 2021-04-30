---
layout: post
title: Replace One/Many Distinctions with Composite (303)
category: Refactoring
tag: [DesignPattern, Refactoring] 
---

#### 개요

>어떤 클래스에서 주어진 객체를 처리할 때 그 **객체의 개수에 따라 서로 다른 로직**을 사용하고 있다면,
컴포짓을 사용해 **객체의 개수에 상관없이 한 로직으로** 처리할 수 있도록 만든다.

<br>

#### 동기

만약 어떤 클래스에 하나의 객체를 처리하는 메서드가 있는데, 로직이 이와 거의 비슷하면서 여러 객체(즉, 객체의 리스트)를 처리하는 메서드가 있다면, 이는 **한 객체와 여러 객체를 구별**하고 있는 것이다. 이럴 땐 다음과 같은 문제가 발생할 수 있다.  
- 코드가 중복된다.
    >코드가 하나일 때와 여러개일 때 처리 로직은 기본적으로 동일한 경우가 많아 두 메서드 사이 코드가 중복되기 마련이다.
- 클라이언트 코드가 복잡해진다.
    >클라이언트에서도 객체의 개수에 따라 메서드가 별도로 존재한다면, 어쩔수 없이 호출을 다르게 해야 한다.
- 처리 결과를 취합하기 위해서 추가적인 처리가 필요하다.
    >예를 들어 빨간색이며 5달러 이하이거나, 파란색이며 10달러 이하인 상품을 검색하는 코드를 만든다면? 두 개의 조건을 합쳐야 한다. 

이런 상황에서는 [Composite](https://jo631.github.io/designpattern/2021/04/19/Composite/) 패턴을 사용하는 것이 훨씬 좋으며, 다음과 같은 장점이 있다.  
- 객체의 개수와 무관하게 **메서드 하나로 처리**할 수 있어 코드의 중복이 없어진다.
- 클라이언트 코드 내에서도 객체의 개수에 상관없이 동일한 메서드를 사용할 수 있다.  
- 객체의 트리를 처리할 때에도 클라이언트는 처리 메서드를 한번만 호출할 수 있다.

하지만, 아래 두 가지의 장점이 그렇게 중요하다고 생각되지 않을 경우 굳이 Composite 패턴을 사용하지 않는 것도 방법이다.  

<br>

#### 장점

- 객체가 하나일 때와 여러 개일 때를 처리하기 위해 각각 존재했던 중복된 처리 코드를 제거할 수 있다.  
- 객체 하나를 처리하는 방법과 객체 여러 개를 처리하는 방법이 하나로 통일된다.  
- 여러 개의 객체를 처리하기 위해 필요한 추가적인 기능을 부여할 수 있다.
    > 예를 들면 AND, OR 조건 처리


#### 단점

- 컴포짓을 생성하는 동안 **타입 안전성**을 위해 런타임 타입 검사가 필요할 수도 있다.
    > 컴포짓에 유효하지 않은 타입의 객체를 넣는 것을 예방하기 위해, 검사하는 절차가 필요할 수도 있다.  

<br>

#### 절차

하나의 객체를 처리하는 메서드를 **단수 객체 메서드**, 여러개 처리 메서드를 **복수 객체 메서드**라고 하자.  

1. 복수 객체 메서드는 리스트를 인자로 받는다. 이때 새 클래스를 하나 만들어 **생성자가 리스트를 인자**로 받도록 하고 리스트를 리턴하는 get 메서드를 구현한다. 이 클래스가 이후 Composite 클래스가 될 것이다. 복수 객체 메서드 안에서 새로 만들어진 Composite 클래스 변수를 선언하고 객체를 하나 생성한다. 그리고 복수 객체 메서드의 기존 코드 가운데 **리스트에 접근하는 부분**을 모두 앞서 만든 **get 메서드**를 통해 접근하도록 수정한다.

2. 복수 객체 메서드 내부에서 **리스트를 다루는 코드**에 [Extract Method](https://jo631.github.io/refactoring/2021/04/07/RefactoringToPattern/#extract-method)를 정용해 별도의 메서드를 뽑아내고 public으로 만든다. 이 후 Move Method를 적용해 **Composite 클래스로 옮긴다.**  

3. 이제 복수와 단수 메서드의 구현이 거의 동일할것이다. 주된 차이점은, 복수 객체 메서드는 Composite을 생성해 사용한다는 것이다. 그 외의 차이점이 있으면 리팩터링을 이용해 제거한다.  

4. Composite 객체를 파라미터로 하여 단수 객체 메서드를 호출하는 한 줄의 코드만을 포함하도록 복수 객체 메서드를 수정한다. 이렇게 하면 Composite 클래스가 단수 객체 메서드에 사용하는 타입과 인터페이스 또는 수퍼클래스를 공유하도록 수정해야 하는데, 컴포짓 클래스를 대상 타입의 서브클래스로 만들어도 되고, [Extract Interface](https://jo631.github.io/refactoring/2021/04/07/RefactoringToPattern/#extract-interface)를 통해 둘을 포괄하는 새로운 인터페이스를 만들어도 된다.  

5. 이제 복수 객체 메서드는 단 한줄의 코드만으로 구현되어 있으므로 [Inline Method](https://jo631.github.io/refactoring/2021/04/07/RefactoringToPattern/#inline-method)를 통해 인라인화 한다.  

6. Composite 클래스에 [Encapsulate Collection](https://jo631.github.io/designpattern/2021/04/30/Encapsulate-Collection/)을 적용한다. 그 결과 Composite 클래스에 `add()` 메서드가 생길 것이다. 클라이언트가 Composite 클래스의 생성자에 객체를 파라미터로 넣어주는 대신, `add()`를 호출하도록 수정한다.  

<br>

#### 구현

상품 검색 시스템을 만든다고 가정해보자. `ProductRepo` 객체로 부터 `Product` 객체의 목록을 얻기 위해 `Spec` 객체를 이용하는 코드이다. 이 예제는 Specifacation 패턴과 관련이 있다.  

```java

public class ProductRepo{
    private List products = new ArrayList();

    public Iterator iterator(){
        return this.products.iterator();
    }

    public List selectBy(Spec spec){
        List foundedProducts = new ArrayList();
        Iterator products = this.iterator();

        while(products.hasNext()){
            Product product = (Product)products.next();
            if(spec.isSatisfiedBy(product)){
                foundedProducts.add(product);
            }
        }
        return foundedProducts;
    }
    
    public List selectBy(List specs){
        List foundedProducts = new ArrayList();
        Iterator products = this.iterator();

        while(products.hasNext()){
            Product product = (Product)products.next();
            Iterator specifications = specs.iterator();
            boolean satisfiesAllSpecs = true;
            while(specifications.hasNext()){
                Spec productSpec = ((Spec)specifications.next());
                satisfiesAllSpecs &= productSpec.isSatisfiedBy(product);
            }
            if(satisfiesAllSpecs){
                foundedProducts.add(product);
            }
        }
        return foundedProducts;
    }
}
```

위를 보면, List를 받는 selectBy의 코드가 훨씬 복잡하다.  
Composite 패턴을 사용하지 않고 바꾼다면?

```java
public class ProductRepo{
    private List products = new ArrayList();

    public Iterator iterator(){
        return this.products.iterator();
    }

    public List selectBy(Spec spec){
        Spec[] specs = {spec};
        return this.selectBy(Arrays.asList(specs));
    }
    
    public List selectBy(List specs){ ... }
```

이렇게 수정하면, 복잡한 코드는 그대로지만 코드 중복이 사라진다.  
그렇다면 Composite 패턴으로 리팩터링하는것 보다 위 방법이 좋을까? 아닐수도 있고 그럴수도 있다. 답은 상황마다 다르다.
이번 예제에선 OR, AND, NOT 조건을 사용할 수 있다고 생각해보자.  
`product.getColor() != targetColor || product.getPrice() < targetPrice`

위와 같은 조건을 selectBy(List ..)는 처리할 수 없다. 또한 클라이언트 측에서도 해당 조건을 위해 여러개의 배열을 만들어 받고, AND OR 연산을 해야 한다. 따라서 해당 기능을 구현하기 위해서는 Composite 패턴으로 리팩터링하는 것이 더 좋은 경우에 해당된다.  

##### 절차 1

`selectBy(List specs)`는 복수 객체 메서드이고 이 파라미터 값을 보관하며 그에 대한 `get()` 메서드를 제공하는 클래스를 만드는 것이 목표이다.  

```java
class CompositeSpec{
    private List specs;

    public CompositeSpec(List specs){
        this.specs = specs;
    }

    public List getSpecs(){
        return this.specs;
    }
}
```

다음 기존 코드를 수정한다.  

```java
public class ProductRepo{
    private List products = new ArrayList();

    public Iterator iterator(){
        return this.products.iterator();
    }

    public List selectBy(Spec spec){
        List foundedProducts = new ArrayList();
        Iterator products = this.iterator();

        while(products.hasNext()){
            Product product = (Product)products.next();
            if(spec.isSatisfiedBy(product)){
                foundedProducts.add(product);
            }
        }
        return foundedProducts;
    }
    
    public List selectBy(List specs){
        //추가
        CompositeSpec spec = new CompositeSpec(specs);
        
        List foundedProducts = new ArrayList();
        Iterator products = this.iterator();

        while(products.hasNext()){
            Product product = (Product)products.next();
            //추가
            Iterator specifications = spec.getSpecs().iterator();
            
            boolean satisfiesAllSpecs = true;
            while(specifications.hasNext()){
                Spec productSpec = ((Spec)specifications.next());
                satisfiesAllSpecs &= productSpec.isSatisfiedBy(product);
            }
            if(satisfiesAllSpecs){
                foundedProducts.add(product);
            }
        }
        return foundedProducts;
    }
}
```

##### 절차 2

`selectBy(List specs)`에 Extract Method를 적용해 List 객체 내 `Spec` 객체들을 처리하는 코드를 별도의 메서드로 분리한다.  


```java
public class ProductRepo{
    private List products = new ArrayList();

    public Iterator iterator(){
        return this.products.iterator();
    }

    public List selectBy(Spec spec){
        List foundedProducts = new ArrayList();
        Iterator products = this.iterator();

        while(products.hasNext()){
            Product product = (Product)products.next();
            if(spec.isSatisfiedBy(product)){
                foundedProducts.add(product);
            }
        }
        return foundedProducts;
    }
    
    public List selectBy(List specs){
        CompositeSpec spec = new CompositeSpec(specs);
        List foundedProducts = new ArrayList();
        Iterator products = this.iterator();

        while(products.hasNext()){
            Product product = (Product)products.next();
            
            if(this.isSatisfiedBy(spec, product))){
                foundedProducts.add(product);
            }
        }
        return foundedProducts;
    }

    //추가
    public boolean isSatisfiedBy(CompositeSpec spec, Product product){
        Iterator specifications = spec.getSpecs().iterator();
        boolean satisfiesAllSpecs = true;
        while(specifications.hasNext()){
            Spec productSpec = ((Spec)specifications.next());
            satisfiesAllSpecs &= productSpec.isSatisfiedBy(product);
        }
        return satisfiesAllSpecs;
    }
}
```

다음, `isSatisfiedBy(spec, product)`를 Spec 클래스로 옮긴다.  

```java
class CompositeSpec{
    private List specs;

    public CompositeSpec(List specs){
        this.specs = specs;
    }

    public List getSpecs(){
        return this.specs;
    }

    public boolean isSatisfiedBy(Product product){
        Iterator specifications = this.getSpecs().iterator();
        boolean satisfiesAllSpecs = true;
        while(specifications.hasNext()){
            Spec productSpec = ((Spec)specifications.next());
            satisfiesAllSpecs &= productSpec.isSatisfiedBy(product);
        }
        return satisfiesAllSpecs;
    }
}
```

##### 절차 3

이제 두 `selectBy(...)`의 코드가 동일해졌다. 유일하게 다른 점은 `selectBy(List Spec)`에서는 `CompositeSPec` 객체를 생성한다는 점이다.  

##### 절차 4

복수 객체 메서드에서 다음과 같이 Spec을 호출하도록 수정한다.  

```java
public class ProductRepo{
    public List selectBy(Spec spec) { ... }
    public List selectBy(List specs) {
        return this.selectBy(new CompositeSpec(specs))
    }
}

class CompositeSpec extends Spec { ... }
```

사실 Spec은 CompositeSpec의 부모 클래스로 나타낼 수 있다. 따라서 위와 같이 설정해준다면 코드가 정상적으로 작동 할 것이다. 

##### 절차 5

`selectBy(List specs)`의 코드는 단 한줄이므로, Inline Method를 적용한다.  

`List foundProducts = repo.selectBy(specs)`
만약 다음과 같은 테스트코드가 있다면

`List foundProducts = repo.selectBy(new CompositeSpec(specs))`
로 바꿀 수 있다.

이제 마지막 단계만 남았다.  

##### 절차 6

`List foundProducts = repo.selectBy(new CompositeSpec(specs))`
위의 코드는 좀 미심쩍다. 클라이언트 코드에서의 안정성이 필요하다.  

```java
class CompositeSpec extends Spec{
    //수정됨
    private List specs = new ArrayList();

    //삭제, 더이상 필요 없음
    public CompositeSpec(List specs){
        this.specs = specs;
    }

    //수정됨
    public List getSpecs(){
        return Collections.unmodifiableList(this.specs);
    }

    public boolean isSatisfiedBy(Product product){
        Iterator specifications = this.getSpecs().iterator();
        boolean satisfiesAllSpecs = true;
        while(specifications.hasNext()){
            Spec productSpec = ((Spec)specifications.next());
            satisfiesAllSpecs &= productSpec.isSatisfiedBy(product);
        }
        return satisfiesAllSpecs;
    }

    //추가됨
    public void add(Spec spec){
        this.specs.add(spec);
    }
}
```

다음 클라이언트에서 호출을 할 땐, spec을 빈 객체로 초기화한다.  
```java
CompositeSpec specs = new CompositeSpec();
specs.add(new ColorSpec(Color.red));
specs.add(new BelowPriceSpec(10.00));
List foundProducts = productRepo.selectBy(specs);
```

이렇게 하면 리팩토링에 성공했다.  

아래는 최종 코드들이다.  

```java
class CompositeSpec extends Spec{
    private List specs = new ArrayList();

    public List getSpecs(){
        return Collections.unmodifiableList(this.specs);
    }

    public boolean isSatisfiedBy(Product product){
        Iterator specifications = this.getSpecs().iterator();
        boolean satisfiesAllSpecs = true;
        while(specifications.hasNext()){
            Spec productSpec = ((Spec)specifications.next());
            satisfiesAllSpecs &= productSpec.isSatisfiedBy(product);
        }
        return satisfiesAllSpecs;
    }

    public void add(Spec spec){
        this.specs.add(spec);
    }
}

public class ProductRepo{
    private List products = new ArrayList();

    public Iterator iterator(){
        return this.products.iterator();
    }

    public List selectBy(Spec spec){
        List foundedProducts = new ArrayList();
        Iterator products = this.iterator();

        while(products.hasNext()){
            Product product = (Product)products.next();
            if(spec.isSatisfiedBy(product)){
                foundedProducts.add(product);
            }
        }
        return foundedProducts;
    }
}
```

##### 추가

해당 패턴에서는 , AND와 OR연산에 대해 다루지 않았다. 이는 추후 Interpreter 리팩토링에 대해 다룰 때 해당 예제를 사용하겠다.