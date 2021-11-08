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

      