## 1.外观模式

当一个系统中存在复杂的`子系统`，而客户端需要与这些子系统进行交互时，外观模式

（Facade Pattern）可以提供一个简化的接口，将子系统的复杂性隐藏起来，为客户端提供

一个更简单、更统一的接口。

外观模式是一种`结构型`设计模式，它通过一个外观类来封装一组子系统的接口，使得客

户端可以通过外观类来间接访问子系统，从而简化客户端与子系统之间的交互过程。

## 2.主要参与者

### 2.1外观类

外观模式的`核心`。它知道哪些子系统类负责处理客户端的请求，并将这些请求委托给相应

的子系统进行处理。

### 2.2子系统类

处理客户端的具体请求。

## 3.优缺点

外观模式的`优点`包括：

- 简化客户端与子系统之间的交互，使得客户端更加简单、易于使用。
- 将子系统的复杂性封装起来，减少了客户端与子系统之间的耦合，提高了系统的松耦合性和灵活性。如果底层子系统发生变化，只需要修改外观类的实现，而不会影响客户端代码。
- 提供了一个统一的接口，使得客户端不需要了解子系统的具体实现细节。

然而，外观模式也有一些`缺点`：

- 外观模式增加了一个额外的抽象层，有可能导致性能上的一定损失。但在现代计算机系统中，这种损失通常可以忽略不计。
- 如果子系统的接口发生变化，可能需要修改外观类，这`违反了开闭原则`。

## 4.适用场景

1. 当一个系统中的`子系统非常复杂`，且它们之间存在较多的依赖关系，直接与这些子系统进行交互会增加客户端的复杂性时，可以考虑使用外观模式。
2. 当需要提供一个简化的接口来代替复杂的子系统接口，隐藏子系统的实现细节，从而`简化客户端的使用`时，可以使用外观模式。
3. 当希望将`子系统和客户端之间解耦`，降低客户端与子系统之间的依赖性，使得子系统可以独立演化而不影响客户端时，可以采用外观模式。