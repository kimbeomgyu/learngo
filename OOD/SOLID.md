# OOD object 중심의 설계

- 설계가 중요하다

- SOLID 법칙

- OOP를 할때에 지향해야 될 목표

- 5가지 원칙은 O를 위해서고 좋은 프로그래밍을 위해서다

- 의존성을 낮추고 응집성을 높히는게 목표

### S - Single Responsibility Principle - 단일 책임 원칙

- 하나의 객체가 여러 책임을 가지고 있으면 기존의 기능을 확장할 때마다 수정을 해줘야 한다.
- 하나의 객체가 하나의 책임을 가지고 있으면 새로운 기능이 확장될 때 새로운 기능를 만들기만 하면 된다.

### O - Open Closed Principle - 개방 패쇄 원칙

- 확장에는 열려있고 변경에 닫혀있는 것

### L - Liskov substitution Principle - 리스코프 치환 원칙

- 기존의 타입과 확장된 타입의 동작이 같아야 한다.

- ...어렵다..

### I - Interface Segregation Principle - 인터페이스 분리 원칙

- 여러개의 관계를 모은 인터페이스보다 하나씩 정의하는게 더 좋다.

### D - Dependency Inversion Principle - 의존 역전 원칙

- 관계를 인터페이스에 의존해야지 객체에 의존하면 ~~안된다.~~

- 관계를 인터페이스에 의존하는게 더 좋다.
