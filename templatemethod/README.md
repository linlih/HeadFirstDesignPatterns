# template method

模板方法就是说设计的一个模板的方法过程，子类按照这个模板实现要求的方法即可。然后模板方法调用统一的步骤去完成任务。

比如这里就是咖啡因饮料的制作过程抽象成一个模板方法。子类完成自己的Brew和AddCondiments即可，其他可以使用模板方法默认的方法。但是由于go不能支持子类实现的时候覆盖父类的方法，所以这里需要抽象出接口，将实际的对象传回去。也就是：tea.PrepareRecipe(tea)，需要将tea对象也传回去。

同时模板方法可以开放一些钩子函数，让子类可以通过钩子函数来控制模板方法的执行过程，也就是示例代码中的CustomerWantsCondiments函数。



