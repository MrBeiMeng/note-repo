# Java EE

## 第一章 spring 的基本应用

> spring 的概念和优点
>
> spring中的ioc 和 DI 思想
>
> APplicationContext容器的使用
>
> 属性setter方法注入的实现

## 第一节课运行spring测试环境

`配置applicationContext.xml 配置java方法 在applicationContext中注入。`

**spring框架的优点**

**spring和springMVC 的区别与关系**

> Spring 是IOC和APO的容器框架，springMVC是基于spring功能之上添加的web框架，想用springMVC 必须先依赖Spring
>
> 详细点来说，spring是一个管理bean的容器，也可以说是包括很多开源项目的总成，springmvc是其中一个开源项目。
>
> springMVC是一个MVC模式的web开发框架。

**IOC控制反转**

> ioc-inversion of COntrol, 即控制反转，不是什么技术，而是一种设计思想。在java开发中，ioc意味着将你设计好的对象交由容器控制，而不是传统的在你的对象内部直接控制。
>
> 也就是对于之前没有引入ioc之前，你需要的每个对象都是通过自己new出来的。但是引入ioc之后，所有的对象不用在去new，而会有ioc容器主动创建一个对象并注入。

IoC和DI

> DI - Dependency Injection,即“依赖注入”: 是组件之间以来关系由容器在于型器件决定，形象地说，即有容器动态的将某个以来祝融到组件之中。
>
> IoC和DI有什么关系呢？其实i他们是同一个概念的不同角度描述。
>
> **控制反转IOC(Inversion of Control)是说创建对象的控制权进行转移。以前创建对象的主动权和创建实际是由自己把控的，现在这种权力转移到第三方。**比如转移交给了ioc容器，他就是一个专门用来创建对象的工厂，你要什么对象，他就给你什么对象，有了ioc容器，依赖关系就变了，原先的依赖关系就没了，他们都依赖ioc容器了，通过ioc容器来建立他们之间的关系。

## 第二章Spring中的bean



## 第三章SpringAOP

**什么是AOP**

> AOP的全称是Aspect_Oriented Programming , 即面向切面编程(也称面向方面编程)。它是面向对象编程的一种补充。
>
> [面试官：什么是AOP？Spring AOP和AspectJ的区别是什么？ - 墨尘无雪 - 博客园 (cnblogs.com)](https://www.cnblogs.com/chaoesha/p/13037368.html)

**AOP的一些专业术语**

> Aspect(切面):切面通常是指封装的用于横向插入系统功能(如事务，日志)的类。
>
> Joinpoint(连接点):在程序执行过程中的某个阶段点，他实际上是对象的一个操作。在springAOP中，连接点就是方法的调用。
>
> Pointcut(切入点):是指切面与程序流程的交叉点，即那些需要处理的连接点。
>
> Advice(通知/增强处理):AOP框架在特定的切入点执行的增强处理，即在定义好的切入点出所要执行的程序代码。
>
> Target Object(目标对象):是指所有被通知的对象，也成为被增强对象。
>
> Proxy(代理):将通知应用到目标对象之后，被动态创建的对象。
>
> Weaving(织入):将切面代码插入到目标对象上，从而生成代理对象的过程。



