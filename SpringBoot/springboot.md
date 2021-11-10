# SpringBoot🏆

## 1.Spring

1.微服务，2响应式编程，3分布式云开发，4web开发，5serverless开发，6事件驱动，7批处理。。。。。

### Spring的生态

覆盖了：

web开发

数据访问

安全控制

分布式

消息服务

移动开发

批处理

### Spring5重大升级

**响应式编程**
![image-20211108120848487](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211108120848487.png)
## 2.SpringBoot

### 	2.1SpringBoot优点

- Create stand-alone Spring applications

- - **创建独立Spring应用**

- Embed Tomcat, Jetty or Undertow directly (no need to deploy WAR files)

- - **内嵌web服务器（tomcat......）**

- Provide opinionated 'starter' dependencies to simplify your build configuration

- - **自动starter依赖，简化构建配置(jar包等配置环境)**

- Automatically configure Spring and 3rd party libraries whenever possible

- - **自动配置Spring以及第三方功能**

- Provide production-ready features such as metrics, health checks, and externalized configuration

- - **提供生产级别的监控、健康检查及外部化配置**

- Absolutely no code generation and no requirement for XML configuration

- - **无代码生成、无需编写XML**
- **SpringBoot是整合Spring技术栈的一站式框架**

- **SpringBoot是简化Spring技术栈的快速开发脚手架**

### 2.2、SpringBoot缺点

- 人称版本帝，迭代快，需要时刻关注变化
- 封装太深，内部原理复杂，不容易精通

### 官方文档
![image-20211108133352166](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211108133352166.png)
![image-20211108133404664](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211108133404664.png)
![image-20211108145130516](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211108145130516.png)
## SpringBoot2入门

### 1、系统要求

- [Java 8](https://www.java.com/) & 兼容java14 .
- Maven 3.3+

- idea 2019.1.2

- ### 1.1、maven设置

    - 进入安装maven的配置文件，->  conf  ->settings.xml

```xml
<mirrors>
      <mirror>
        <id>nexus-aliyun</id>
        <mirrorOf>central</mirrorOf>
        <name>Nexus aliyun</name>
  <url>http://maven.aliyun.com/nexus/content/groups/public</url>
      </mirror>
  </mirrors>

  <profiles>
         <profile>
              <id>jdk-1.8</id>
              <activation>
                <activeByDefault>true</activeByDefault>
                <jdk>1.8</jdk>
              </activation>
              <properties>
 <maven.compiler.source>1.8</maven.compiler.source>             		       		<maven.compiler.target>1.8</maven.compiler.target>    <maven.compiler.compilerVersion>1.8</maven.compiler.compilerVersion>
              </properties>
         </profile>
  </profiles>
```

### 2.HelloWorld

#### 2.1、创建maven工程

#### 2.2、引入依赖

```xml
<parent>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-parent</artifactId>
    <version>2.3.4.RELEASE</version>
</parent>


<dependencies>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-web</artifactId>
    </dependency>

</dependencies>
```

#### 2.3、创建主程序

* ```java
   /**
   
    * 主程序类
   
    * @SpringBootApplication：这是一个SpringBoot应用
      */
      @SpringBootApplication
      public class MainApplication {
   
      public static void main(String[] args) {
          SpringApplication.run(MainApplication.class,args);
      }
      }
   ```



#### 2.4、编写业务

```java
//@ResponseBody
//@Controller

@RestController
public class HelloController {

    @RequestMapping("/hello")
    public String handle01(){
        return "Hello,Spring Boot 2";
    }
}
```

#### 2.5、测试

直接运行main方法

#### 2.6、简化配置

application.properties            server.port=8888

#### 2.7、简化部署

```xml
<build>
        <plugins>
            <plugin>
                <groupId>org.springframework.boot</groupId>
                <artifactId>spring-boot-maven-plugin</artifactId>
            </plugin>
        </plugins>
    </build>
```

把项目打成jar包，直接在目标服务器执行即可。

![image-20211108160858209](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211108160858209.png)

![image-20211108160922023](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211108160922023.png)

![image-20211108160948764](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211108160948764.png)

注意点：

- 取消掉cmd的快速编辑模式

（右键属性->取消快速编辑模式）

## 了解自动配置原理

#### 1、SpringBoot特点

#### 1.1、依赖管理

- 父项目做依赖管理

```xml
依赖管理    
<parent>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-parent</artifactId>
        <version>2.3.4.RELEASE</version>
</parent>

他的父项目
 <parent>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-dependencies</artifactId>
    <version>2.3.4.RELEASE</version>
  </parent>

几乎声明了所有开发中常用的依赖的版本号,自动版本仲裁机制
```

- 开发导入starter场景启动器

- ```xml
  1、见到很多 spring-boot-starter-* ： *就某种场景
  2、只要引入starter，这个场景的所有常规需要的依赖我们都自动引入
  3、SpringBoot所有支持的场景
  https://docs.spring.io/spring-boot/docs/current/reference/html/using-spring-boot.html#using-boot-starter
  4、见到的  *-spring-boot-starter： 第三方为我们提供的简化开发的场景启动器。
  5、所有场景启动器最底层的依赖
  <dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter</artifactId>
    <version>2.3.4.RELEASE</version>
    <scope>compile</scope>
  </dependency>
  ```

  ![image-20211108163859724](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211108163859724.png)

    - 无需关注版本号，自动版本仲裁

1、引入依赖默认都可以不写版本

![image-20211108164353606](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211108164353606.png)2、引入非版本仲裁的jar，要写版本号。

![image-20211108164404894](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211108164404894.png)

- 可以修改默认版本号

- ```xml
  1、查看spring-boot-dependencies里面规定当前依赖的版本 用的 key。
  2、在当前项目里面重写配置
      <properties>
          <mysql.version>5.1.43</mysql.version>
      </properties>
  ```

### 1.2、自动配置

- 自动配好Tomcat

- - 引入Tomcat依赖。

- 配置Tomcat

- ```xml
    <dependency>
          <groupId>org.springframework.boot</groupId>
          <artifactId>spring-boot-starter-tomcat</artifactId>
          <version>2.3.4.RELEASE</version>
          <scope>compile</scope>
        </dependency>
    ```

- 自动配好SpringMVC

- - 引入SpringMVC全套组件
- 自动配好SpringMVC常用组件（功能）

- 自动配好Web常见功能，如：字符编码问题

- - SpringBoot帮我们配置好了所有web开发的常见场景

- 默认的包结构

- - 主程序所在包及其下面的所有子包里面的组件都会被默认扫描进来
- 无需以前的包扫描配置

- - 想要改变扫描路径，@SpringBootApplication(scanBasePackages=**"com.atguigu"**)

- - - 或者@ComponentScan 指定扫描路径

```java
@SpringBootApplication
等同于
@SpringBootConfiguration
@EnableAutoConfiguration
@ComponentScan("com.atguigu.boot")
```

- 各种配置拥有默认值

- - 默认配置最终都是映射到某个类上，如：MultipartProperties
- 配置文件的值最终会绑定每个类上，这个类会在容器中创建对象

- 按需加载所有自动配置项

- - 非常多的starter
- 引入了哪些场景这个场景的自动配置才会开启

- - SpringBoot所有的自动配置功能都在 spring-boot-autoconfigure 包里面

    ![image-20211108171107514](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211108171107514.png)

- - **虽然引入了很多类，但是有些是不生效的需要自己在进行引入批处理场景**

- ```xml
      <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-batch</artifactId>
      </dependency>
      ```
## 容器功能

### 组件添加

#### 1、@Configuration

基本使用（告诉SpringBoot这是一个配置类==配置文件）

配置类本身也是个文件

```java
//com.atguigu.boot.config.MyConfig$$EnhancerBySpringCGLIB$$64b915d9@3163987e
MyConfig bean = run.getBean(MyConfig.class);
System.out.println(bean);
```

**●**Full模式与Lite模式

○示例

○最佳实战

■配置 类组件之间无依赖关系用Lite模式加速容器启动过程，减少判断（**lite(proxyBeanMethods = false)每一次调用都会产生一个新的对象**）

■配置类组件之间有依赖关系，方法会被调用得到之前单实例组件，用Full模式（**Full(proxyBeanMethods = true)容器中所有配置的组件方法，外部的调用的时候都会在容器中找组件**）

```java
//总结：使用代理，容器实例唯一，不使用代理，多次调用多个实例
```

#### 2、@Bean

```
@Bean//给容器中添加组件，以方法名作为组件的id，返回值类型作为组件类型，返回的值，就是组件在容器中的实例
```

![image-20211109093420161](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211109093420161.png)

```java
/**    👩Configuration的实例
 *1.配置类里面使用@Bean标注在方法上给容器注册组件，默认也是单实例的
 *2. 配置类本身也是组件
 ❗❗❗❗📵*3.proxyBeanMethods:代理bean的方法
 * Full(proxyBeanMethods = true)容器中所有配置的组件方法，外部的调用的时候都会在容器中找组件
 * lite(proxyBeanMethods = false)每一次调用都会产生一个新的对象
 * 组件依赖
 *
 */
@Configuration(proxyBeanMethods = true)//告诉SpringBoot这是一个配置类==配置文件
public class MyConfig {
    /**
     * FUll：外部无论对配置类中的这个组件注册方法调用多少次获取的都之前注册容器中的单实例对象
     * @return
     */

    @Bean//给容器中添加组件，以方法名作为组件的id，返回值类型作为组件类型，返回的值，就是组件在容器中的实例
    public User user01(){
        User zhangsan=new User("zhangsan",18);
        //user组件依赖了Pet组件
        zhangsan.setPet(tomcatPet());
        return zhangsan;
    }
    @Bean("tom")
    public Pet tomcatPet(){
        return new Pet("tomcat",23);
    }
}

👨Configuration的测试
    
/**主程序类
 * @SpringBootApplication:这是一个SpringBoot应用
 */
@SpringBootApplication(scanBasePackages = "com.atguigu")
public class MainApplication {
    public static void main(String[] args) {
        //返回我们IOC容器
        ConfigurableApplicationContext run = SpringApplication.run(MainApplication.class, args);
        //2查看容器里面的组件
        String[] names=run.getBeanDefinitionNames();
        for (String name : names) {
            System.out.println(name);
        }
        //从容器中获取组件
        Pet tom1 = run.getBean("tom", Pet.class);
        Pet tom2 = run.getBean("tom", Pet.class);
System.out.println("组件"+(tom1==tom2));//组件true
        
        //com.atguigu.boot.config.MyConfig$$EnhancerBySpringCGLIB$$64b915d9@3163987e
        MyConfig bean = run.getBean(MyConfig.class);
        System.out.println(bean);
        //如果@Configuration(proxyBeanMethods=true)代理对象调用方法，SpringBoot总会检查这个组件是否在容器中
        //保持组件单实例
        🦀//总结：使用代理，容器实例唯一，不使用代理，多次调用多个实例
        User user1 = bean.user01();
        User user2= bean.user01();
        System.out.println(user1==user2);//true

        User user01 = run.getBean("user01", User.class);
        Pet tom=run.getBean("tom",Pet.class);
        System.out.println("用户的宠物："+(user01.getPet()==tom));//用户的宠物：true但使用代理的时候，是从容器中取的宠物。
    }
}

```

dao调用pojo，service调用dao，controller调用service

#### 3、**@Component//代表他是一个组件**

**标注一个类为Spring容器的Bean，（把普通pojo实例化到spring容器中，相当于配置文件中的<bean id="" class=""/>）**

#### 4、**@Controller 控制器（注入服务）**

**用于标注控制层，相当于struts中的action层**

#### 5、**@Service 服务（注入dao）**

**用于标注服务层，主要用来进行业务的逻辑处理**

#### 6、**@Repository（实现dao访问）**

**用于标注数据访问层，也可以说用于标注数据访问组件，即DAO组件**

**@Compo**nentScan主要就是定义**扫描的路径**从中找出标识了**需要装配**的类自动装配到spring的bean容器中

```java
@ComponentScan(value="com.zhang")
@Configuration
public class MainScanConfig {
}
```

#### 7、@**Import**

* ```java
   * 4、@Import({User.class, DBHelper.class})
   * 给容器中自动创建出这两个类型的组件、默认组件的名字就是全类名
   //  com.atguigu.boot.bean.User
    // ch.qos.logback.core.db.DBHelper@40147317
       */
   @Import({User.class, DBHelper.class})
   @Configuration(proxyBeanMethods = false) //告诉SpringBoot这是一个配置类 == 配置文件
   public class MyConfig {
   }
   ```

#### 8、@Conditional

**条件装配：满足Conditional指定的条件，则进行组件注入**

![image-20211109103112435](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211109103112435.png)




```java
=====================测试条件装配==========================
@Configuration(proxyBeanMethods = false) //告诉SpringBoot这是一个配置类 == 配置文件
//@ConditionalOnBean(name = "tom")
@ConditionalOnMissingBean(name = "tom")
public class MyConfig {
/**
 * Full:外部无论对配置类中的这个组件注册方法调用多少次获取的都是之前注册容器中的单实例对象
 * @return
 */

@Bean //给容器中添加组件。以方法名作为组件的id。返回类型就是组件类型。返回的值，就是组件在容器中的实例
public User user01(){
    User zhangsan = new User("zhangsan", 18);
    //user组件依赖了Pet组件
    zhangsan.setPet(tomcatPet());
    return zhangsan;
}

@Bean("tom22")
public Pet tomcatPet(){
    return new Pet("tomcat");
}
}

public static void main(String[] args) {
        //1、返回我们IOC容器
        ConfigurableApplicationContext run = SpringApplication.run(MainApplication.class, args);
         //2、查看容器里面的组件
    String[] names = run.getBeanDefinitionNames();
    for (String name : names) {
        System.out.println(name);
    }

    boolean tom = run.containsBean("tom");
    System.out.println("容器中Tom组件："+tom);
//false
    boolean user01 = run.containsBean("user01");
    System.out.println("容器中user01组件："+user01);
//true
    boolean tom22 = run.containsBean("tom22");
    System.out.println("容器中tom22组件："+tom22);
//true
    }
```

### 原生配置文件引入

#### 1、@ImportResource

======================beans.xml=========================

```xml
<?xml version="1.0" encoding="UTF-8"?>
<beans xmlns="http://www.springframework.org/schema/beans"
       xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
       xmlns:context="http://www.springframework.org/schema/context"
       xsi:schemaLocation="http://www.springframework.org/schema/beans http://www.springframework.org/schema/beans/spring-beans.xsd http://www.springframework.org/schema/context https://www.springframework.org/schema/context/spring-context.xsd">

    <bean id="haha" class="com.atguigu.boot.bean.User">
        <property name="name" value="zhangsan"></property>
        <property name="age" value="18"></property>
    </bean>
    
    <bean id="hehe" class="com.atguigu.boot.bean.Pet">
        <property name="name" value="tomcat"></property>
    </bean>
</beans>
```

```java
@ImportResource("classpath:beans.xml")//导入spring的原生文件
public class MyConfig {}

======================测试=================
        boolean haha = run.containsBean("haha");
        boolean hehe = run.containsBean("hehe");
        System.out.println("haha："+haha);//true
        System.out.println("hehe："+hehe);//true
```

### 配置绑定

就是将application.properties里面的配置进行绑定

![image-20211109134308782](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211109134308782.png)

将这两个绑定到类上

@ConfigurationProperties

两种方式：

- **@Component + @ConfigurationProperties**

- ### @EnableConfigurationProperties + @ConfigurationProperties

```java
只有在容器中的组件，才会拥有SpringBoot提供的强大功能
 */
 @Component
 @ConfigurationProperties(prefix = "mycar")
👩‍🚀 //有这两个就相当于这个类的所有属性是和application.prperties这个核心配置文件里面mycar这个前缀绑定的在一起的
public class Car {
private String brand;
    private Integer price;
    
public String getBrand() {
 return brand;
}

public void setBrand(String brand) {
 this.brand = brand;
}

public Integer getPrice() {
 return price;
}

public void setPrice(Integer price) {
 this.price = price;
}

@Override
public String toString() {
    return "Car{" +
            "brand='" + brand + '\'' +
            ", price=" + price +
         '}';
          }
 }
  
@EnableConfigurationProperties(Car.class)
//1、开启Car配置绑定功能
//2、把这个Car这个组件自动注册到容器中
public class MyConfig {
}
```
# yaml配置注入

## 配置文件

SpringBoot使用一个全局的配置文件 ， 配置文件名称是固定的

- application.properties

- - 语法结构 ：key=value

- application.yml

- - 语法结构 ：key：空格 value

**配置文件的作用 ：**修改SpringBoot自动配置的默认值，因为SpringBoot在底层都给我们自动配置好了；

比如我们可以在配置文件中修改Tomcat 默认启动的端口号！测试一下！

```
server.port=8081
```

## yaml概述

**这种语言以数据作为中心，而不是以标记语言为重点！**

以前的配置文件，大多数都是使用xml来配置；比如一个简单的端口配置，我们来对比下yaml和xml

传统xml配置：

```xml
<server>   
		<port>8081<port>
</server>
```

yaml配置：

```yaml
server:  
	prot: 8080
```

## yaml基础语法

说明：语法要求严格！

1、空格不能省略

2、以缩进来控制层级关系，只要是左边对齐的一列数据都是同一个层级的。

3、属性和值的大小写都是十分敏感的。

**字面量：普通的值  [ 数字，布尔值，字符串  ]**

字面量直接写在后面就可以 ， 字符串默认不用加上双引号或者单引号；

```
k: v
```

注意：

- “ ” 双引号，不会转义字符串里面的特殊字符 ， 特殊字符会作为本身想表示的意思；

  比如 ：name: "kuang \n shen"  输出 ：kuang  换行  shen

- '' 单引号，会转义特殊字符 ， 特殊字符最终会变成和普通字符一样输出

  比如 ：name: ‘kuang \n shen’  输出 ：kuang  \n  shen

**对象、Map（键值对）**

```
#对象、Map格式k:     v1:    v2:
```

在下一行来写对象的属性和值得关系，注意缩进；比如：

```yaml
student:  
    name: qinjiang 
    age: 3
```

行内写法

```yaml
student: {name: qinjiang,age: 3}
```

**数组（ List、set ）**

用 - 值表示数组中的一个元素,比如：

```yaml
pets: 
    - cat 
    - dog 
    - pig
```

行内写法

```yaml
pets: [cat,dog,pig]
```

**修改SpringBoot的默认端口号**

配置文件中添加，端口号的参数，就可以切换端口；

```yaml
server:
	port: 8082
```

## yaml注入配置文件

yaml文件更强大的地方在于，他可以给我们的实体类直接注入匹配值！

1、在springboot项目中的resources目录下新建一个文件 application.yml

### 2、编写一个实体类 person；

```java
@Component //注册bean到容器中
public class Person {
    private String name;
    private Integer age;
    private Boolean happy;
    private Date birth;
    private Map<String,Object> maps;
    private List<Object> lists;
    private Dog dog;

    //有参无参构造、get、set方法、toString()方法  
}
```

### 3、我们来使用yaml配置的方式进行注入

```yaml
person:
  name: qinjiang
  age: 3
  happy: false
  birth: 2000/01/01
  maps: {k1: v1,k2: v2}
  lists:
   - code
   - girl
   - music
  dog:
    name: 旺财
    age: 1
```

### 4、person这个对象的所有值都写好了，注入到类中！

```java
/*
@ConfigurationProperties作用：
将配置文件中配置的每一个属性的值，映射到这个组件中；
告诉SpringBoot将本类中的所有属性和配置文件中相关的配置进行绑定
参数 prefix = “person” : 将配置文件中的person下面的所有属性一一对应
*/
@Component //注册bean
@ConfigurationProperties(prefix = "person")
public class Person {
    private String name;
    private Integer age;
    private Boolean happy;
    private Date birth;
    private Map<String,Object> maps;
    private List<Object> lists;
    private Dog dog;
}
```

### 5、IDEA 提示，springboot配置注解处理器没有找到，让我们看文档，我们可以查看文档，找到一个依赖！

![image-20211110164721418](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211110164721418.png)

```xml
<!-- 导入配置文件处理器，配置文件进行绑定就会有提示，需要重启 -->
<dependency>
  <groupId>org.springframework.boot</groupId>
  <artifactId>spring-boot-configuration-processor</artifactId>
  <optional>true</optional>
</dependency>
```

### 6、确认以上配置都OK之后，我们去测试类中测试一下：

```java
@SpringBootTest
class SpringBoot2ApplicationTests {
    @Autowired
    private Person person;
    @Test
    void contextLoads() {
        System.out.println(person);
    }

}
```

### **yaml配置注入到实体类完全OK！**

## 配置绑定

就是将application.properties里面的配置进行绑定

![image-20211109134308782](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211109134308782.png)

将这两个绑定到类上

@ConfigurationProperties

两种方式：

- **@Component + @ConfigurationProperties**

- ### @EnableConfigurationProperties + @ConfigurationProperties

```java
只有在容器中的组件，才会拥有SpringBoot提供的强大功能
 */
 @Component
 @ConfigurationProperties(prefix = "mycar")
👩‍🚀 //有这两个就相当于这个类的所有属性是和application.prperties这个核心配置文件里面mycar这个前缀绑定的在一起的
public class Car {
private String brand;
    private Integer price;
    //get set 
}

@EnableConfigurationProerties(Car.class)
//1、开启Car配置绑定功能
//2、把这个Car这个组件自动注册到容器中
public class MyConfig {
}
```

## 加载指定的配置文件

**@PropertySource ：**加载指定的配置文件；

**@configurationProperties**：默认从全局配置文件中获取值；

1、我们去在resources目录下新建一个**person.properties**文件

```
name=kuangshen
```

2、然后在我们的代码中指定加载person.properties文件

```java
@PropertySource(value = "classpath:person.properties")
@Component //注册bean
public class Person {

    @Value("${name}")
    private String name;

    ......  
}
```

## 配置文件占位符

配置文件还可以编写占位符生成随机数

```yaml
person:
    name: qinjiang${random.uuid} # 随机uuid
    age: ${random.int}  # 随机int
    happy: false
    birth: 2000/01/01
    maps: {k1: v1,k2: v2}
    lists:
      - code
      - girl
      - music
    dog:
      name: ${person.hello:other}_旺财
      age: 1
```

![image-20211110170501426](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211110170501426.png)

## 回顾properties配置

【注意】properties配置文件在写中文的时候，会有乱码 ， 我们需要去IDEA中设置编码格式为UTF-8；

settings-->FileEncodings 中配置UTF-8；

```java
@Component //注册bean
@PropertySource(value = "classpath:user.properties")
public class User {
    //直接使用@value
    @Value("${user.name}") //从配置文件中取值
    private String name;
    @Value("#{9*2}")  // #{SPEL} Spring表达式
    private int age;
    @Value("男")  // 字面量
    private String sex;
}
```

## 对比小结

![image-20211110171408610](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211110171408610.png)

、@ConfigurationProperties只需要写一次即可 ， @Value则需要每个字段都添加

2、松散绑定：这个什么意思呢? 比如我的yml中写的last-name，这个和lastName是一样的， - 后面跟着的字母默认是大写的。这就是松散绑定。

3、JSR303数据校验 ， 这个就是我们可以在字段是增加一层过滤器验证 ， 可以保证数据的合法性

4、复杂类型封装，yml中可以封装对象 ， 使用value就不支持

**结论：**

配置yml和配置properties都可以获取到值 ， 强烈推荐 yml；

如果我们在某个业务中，只需要获取配置文件中的某个值，可以使用一下 @value；

如果说，我们专门编写了一个JavaBean来和配置文件进行一一映射，就直接@configurationProperties。

# JSR303数据校验和多环境切换

Springboot中可以用**@validated**来校验数据，如果数据异常则会统一抛出异常，方便异常中心统一处理。我们这里来写个注解让我们的name只能支持Email格式；

```java
@Component //注册bean
@ConfigurationProperties(prefix = "person")
@Validated  //数据校验
public class Person {

    @Email(message="邮箱格式错误") //name必须是邮箱格式
    private String name;
}
```

**使用数据校验，可以保证数据的正确性；**

## 常见参数

```java
@NotNull(message="名字不能为空")
private String userName;
@Max(value=120,message="年龄最大不能查过120")
private int age;
@Email(message="邮箱格式错误")
private String email;

空检查
@Null       验证对象是否为null
@NotNull    验证对象是否不为null, 无法查检长度为0的字符串
@NotBlank   检查约束字符串是不是Null还有被Trim的长度是否大于0,只对字符串,且会去掉前后空格.
@NotEmpty   检查约束元素是否为NULL或者是EMPTY.
    
Booelan检查
@AssertTrue     验证 Boolean 对象是否为 true  
@AssertFalse    验证 Boolean 对象是否为 false  
    
长度检查
@Size(min=, max=) 验证对象（Array,Collection,Map,String）长度是否在给定的范围之内  
@Length(min=, max=) string is between min and max included.

日期检查
@Past       验证 Date 和 Calendar 对象是否在当前时间之前  
@Future     验证 Date 和 Calendar 对象是否在当前时间之后  
@Pattern    验证 String 对象是否符合正则表达式的规则

.......等等
除此以外，我们还可以自定义一些数据校验规则
```

## 多环境切换

**profile是Spring对不同环境提供不同配置功能的支持，可以通过激活不同的环境版本，实现快速切换环境；**

## properties的多配置文件

我们在主配置文件编写的时候，文件名可以是 application-{profile}.properties/yml , 用来指定多个环境版本；

**例如：**

application-test.properties 代表测试环境配置

application-dev.properties 代表开发环境配置

但是Springboot并不会直接启动这些配置文件，它**默认使用application.properties主配置文件**；

我们需要通过一个配置来选择需要激活的环境：

```
#比如在配置文件中指定使用dev环境，我们可以通过设置不同的端口号进行测试；#我们启动SpringBoot，就可以看到已经切换到dev下的配置了；spring.profiles.active=dev
```

## yaml的多文档块

和properties配置文件中一样，但是使用yml去实现不需要创建多个配置文件，更加方便了 !

```yaml
server:
  port: 8081
#选择要激活那个环境块
spring:
  profiles:
    active: prod

---
server:
  port: 8083
spring:
  profiles: dev #配置环境的名称


---

server:
  port: 8084
spring:
  profiles: prod  #配置环境的名称
```

**注意：如果yml和properties同时都配置了端口，并且没有激活其他环境 ， 默认会使用properties配置文件的！**

## 配置文件加载位置

**外部加载配置文件的方式十分多，我们选择最常用的即可，在开发的资源文件中进行配置！**

官方外部配置文件说明参考文档

![image-20211110202708300](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211110202708300.png)

springboot 启动会扫描以下位置的application.properties或者application.yml文件作为Spring boot的默认配置文件：

```
优先级1：项目路径下的config文件夹配置文件
优先级2：项目路径下配置文件
优先级3：资源路径下的config文件夹配置文件
优先级4：资源路径下配置文件
优先级由高到底，高优先级的配置会覆盖低优先级的配置；
```

![image-20211110202924921](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211110202924921.png)

**SpringBoot会从这四个位置全部加载主配置文件；互补配置；**

我们在最低级的配置文件中设置一个项目访问路径的配置来测试互补问题；

```
#配置项目的访问路径server.servlet.context-path=/kuang
```

## 运维小技巧

指定位置加载配置文件

我们还可以通过spring.config.location来改变默认的配置文件位置

项目打包好以后，我们可以使用命令行参数的形式，启动项目的时候来指定配置文件的新位置；这种情况，一般是后期运维做的多，相同配置，外部指定的配置文件优先级最高

```
java -jar spring-boot-config.jar --spring.config.location=F:/application.properties
```


      