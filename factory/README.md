# factory

在书中的factory模式分为三种，simple factory method，factory method 和 abstract factory method。

simple factory method 中是将类的实例化方法，createPizza 抽离到 simple factory 中。由简单工厂来生产 pizza，可能有一个疑问，尽管我们把这个从原有的orderPizza中抽离出来了，按时如果需要增加 pizza类型，还是需要修改 simple factory中的代码，但是要注意，因为这里只是作为一个示例代码，这个 simple factory 只生产了pizza，但是实际过程中，这个simple factory可能生产的不仅仅是 pizza，还会有其他的功能。这样就体现出了它的好处了，当发生改变的时候，只需要修改这个类就可以了。

simple factory method 存在的问题在于如果有多个地域的工厂店，那么我们可能需要多个 simple factory method，然后初始化工厂传给 PizzaStore去生产pizza。仔细对比一下有点类似于策略模式，只不过策略模式封装的范围更小一点，策略模式封装的是算法集合，而在这里则需要封装生成类的方法。

那么更好的办法是什么呢？将所有的simple factory 继承于一个工厂类就好了。这样在使用的是子类直接创建对应地域的工厂，然后去生产pizza就可以了，不需要像simple factory一样创建工厂，然后把工厂传给pizzaStore去制作。

书本配套的代码很糟糕，没有目录结构分开都放在一个目录下，看起来非常费劲。而且代码的例子举得也不好，比如说 ingredient 中，举个例子就可以了，两三个 ingredient就可以说明问题了，非要整好多个，很多余。

翻译一些词汇(有这么多地域关联性很强的词汇让原本一个知识点增加更多未知的元素)：

`clam` 蛤蜊，`pepperoni` 意大利辣肉肠，`veggie` 蔬菜，`spinach` 菠菜，`eggplant`茄子，`mozzarella` 马苏里拉，`parmesan`帕尔马干酪，`reggiano`巴马干酪，`pepper`胡椒粉，`marinara`马瑞那拉调味汁
