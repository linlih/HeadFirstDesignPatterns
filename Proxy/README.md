# Proxy

代理模式顾名思义了，就是由某个代理去完成某件事情，而不是由对象自己完成。

在书中GumBall Machine的例子使用的是Java提供的RMI接口去完成的，RMI是Java中特有的，所以在Go中文无法实现。但是在Go中可以RPC的机制来实现这个例子。其他的例子也是用到了Java的一些特性，所以这里就无法转化成Go语言来实现。

书中配套各个例子介绍：

`gumball`：使用RMI实现了gumball monitor能够远程打印gumball machine的状态信息

`gumballmonitor`：实现了本地的gumball monitor

`javaproxy`：使用java.lang.reflect中提供的动态代理实现，被代理的类需要实现InvocationHandler中的invoke方法

`virtualproxy`：实现的是虚拟代理，区别于gumball的远程代理和java proxy的动态代理，是一个显示CD封面的例子，将下载封面的任务委托给ImageProxy来完成，下载完成后交给ImageIcon绘制。

由于这几个Java的例子在Go中都不太好实现，虽然可以实现，但是由于涉及到很多其他的知识点，偏离了设计模式本身的内容，所以这里我们选择一个相对简单的Go例子来进行实现。这里例子模拟了Nginx这个代理服务器，它代理了用户的请求，可以对用户的请求进行速率的限制，也可以对请求进行缓存等。


该例子参考自：https://refactoringguru.cn/design-patterns/proxy/go/example





