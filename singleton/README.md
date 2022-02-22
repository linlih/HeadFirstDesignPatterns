# singleton

单例模式相对比较简单，在《Head First设计模式》书本中也给了较多的例子，核心的要点就是将构造函数私有。一些适合Java语言提供的特性相关的，所以就没有办法在Go中实现该例子了。所以这里只实现了`chocolate`这个例子。

但是在Go语言中没有构造函数的说法，所以就需要通过控制一个变量是否被初始化的方式来实现单例模式，一个是通过锁来判断，一个通过Go提供的sync.Once来完成。
（代码参考来自：https://refactoring.guru/design-patterns/singleton/go/example）


