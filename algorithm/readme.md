## 动态规划

### 定义

动态规划算法是通过拆分问题，定义问题状态和状态之间的关系，使得问题能够以递推（或者说分治）的方式去解决。

动态规划算法的基本思想与分治法类似，也是将待求解的问题分解为若干个子问题（阶段），按顺序求解子阶段，前一子问题的解，为后一子问题的求解提供了有用的信息。在求解任一子问题时，列出各种可能的局部解，通过决策保留那些有可能达到最优的局部解，丢弃其他局部解。依次解决各子问题，最后一个子问题就是初始问题的解。

### 基本思想与策略编辑

由于动态规划解决的问题多数有重叠子问题这个特点，为减少重复计算，对每一个子问题只解一次，将其不同阶段的不同状态保存在一个二维数组中。

上面的描述没有动态规划的基础很难看懂，但是也能从中看出一些信息：

1. `拆分问题`

   根据问题的可能性划分，一步一步通过递推或者递归来实现。关键就是这个步骤。动态规划有一类问题就是从后往前推导，有时候我们很容易知道：如果只有一种情况时，最佳的选择应该怎么做，然后根据这个最佳选择往前一步推导，得到前一步的最佳选择。

2. `定义问题状态和状态之间的关系`

   把前面拆分的步骤之间的关系，用一种量化的形式表现出来。类似于高中学的推导公式，因为这种式子很容易用程序写出来，也可以说对程序比较亲和(也就是最后所说的状态转移方程式)。

总结一下，假如可以找到最优解，应该将最优解保存下来，为了往前推导时能够使用前一步的最优解。在这个过程中难免有一些相比于最优解差的解，此时我们应该放弃，`只保存最优解`(动态规划的精髓，减少时间复杂度的关键)，这样我们每一次都把最优解保存了下来，大大降低了时间复杂度。

### 举个栗子

### 解题思路

首先递归应该是我们解决动态规划问题最常用的方法。

递归到动规的一般转化方法为：

如果该递归函数有n个参数,那么就定义一个n维数组；数组下标是递归函数参数的取值范围(也就是数组每一维的大小)。数组元素的值就是递归函数的返回值(初始化为一个标志值,表明还未被填充)；这样就可以从边界值开始逐步的填充数组。相当于计算递归函数的逆过程(这和前面所说的推导过程应该是相同的)。
原文链接:https://blog.csdn.net/ailaojie/article/details/83014821

动规解题的一般思路(标准官方,不过经过前边讲解应该就能理解了)：

1. 将原问题分解为子问题。 

   子问题与原问题形式相同或类似，只是问题规模变小了，从而变简单了。子问题一旦求出就要保存下来，保证每个子问题只求解一遍。

2. **确定状态**。

   在动规解题中，和子问题相关的各个变量的一组取值，称之为一个"状态"。一个状态对应一个或多个子问题所谓的在某个状态的值，这个就是状态所对应的子问题的解，所有状态的集合称为"状态空间"。

   通俗来说，状态就是某个问题某组变量，状态空间就是该问题的所有组变量。 `整个问题的时间复杂度就是状态数目乘以每个状态所需要的时间`。

3. 确定一些初始状态(边界条件)的值

   这个视情况而定，不要以为就是最简单的那个子问题解。真正实践动规千变万化。

4. 确定状态转移方程 

   这一步和第三步是最关键的。由已知推未知。


适合使用动规求解的问题:：

1. 问题具有最优子结构；
2. 无后效性 

其实一般遇到求最优解问题，一般适合使用动态规划。