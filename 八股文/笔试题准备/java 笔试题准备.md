# Java 笔试题 2022年8月10日

--  她会魔法吧。

## 一、选择题(共50题，每题1.5分，共75分。多选题选不全或选错都不得分。)

　　**1. 以下属于面向对象的特征的是(C,D)。(两项)**

　　A) 重载

　　B) 重写

　　C) 封装

　　D) 继承
 

　　**2. 以下代码运行输出是(C)**

```java
public class Person{

    private String name=”Person”;

    int age=0;
    
}

public class Child extends Person{

    public String grade;

    public static void main(String[] args){

        Person p = new Child();

        System.out.println(p.name);
        
    }
}
```



　　A) 输出：Person

　　B) 没有输出

　　C) 编译出错

　　D) 运行出错
 

　　**3. 在使用super 和this关键字时，以下描述正确的是(A)**

　　A) 在子类构造方法中使用super()显示调用父类的构造方法，super()必须写在子类构造方法的第一行，否则编译不通过

　　B) super()和this()不一定要放在构造方法内第一行

　　C) this()和super()可以同时出现在一个构造函数中

　　D) this()和super()可以在static环境中使用，包括static方法和static语句块
 

　　**4. 以下对封装的描述正确的是(D)**

　　A) 只能对一个类中的方法进行封装，不能对属性进行封装

　　B) 如果子类继承了父类，对于父类中进行封装的方法，子类仍然可以直接调用

　　C) 封装的意义不大，因此在编码时尽量不要使用

　　D) 封装的主要作用在于对外隐藏内部实现细节，增强程序的安全性
 

　　**5. 以下对继承的描述错误的是(A)**

　　A) Java中的继承允许一个子类继承多个父类

　　B) 父类更具有通用性，子类更具体

　　C) Java中的继承存在着传递性

　　D) 当实例化子类时会递归调用父类中的构造方法
 

　　**6. 以下程序的运行结果是(D)**

```java
class Person{

    public Person(){

        System.out.println(“this is a Person”);

    }

}

public class Teacher extends Person{

    private String name=”tom”;

    public Teacher(){

        System.out.println(“this is a teacher”);

        super();

    }

    public static void main(String[] args){

        Teacher teacher = new Teacher();

        System.out.println(this.name); // 静态不能调用非静态

    }

}
```



　　A) this is a Person

　　this is a teacher

　　tom

　　B) this is a teacher

　　this is a Person

　　tom

　　C) 运行出错

　　D) 编译有两处错误
 

　　**7. 以下说法错误的是(B)**

　　A) super.方法()可以调用父类的所有非私有方法

　　B) super()可以调用父类的所有非私有构造函数

　　C) super.属性可以调用父类的所有非私有属性

　　D) this和super关键字可以出现在同一个构造函数中
 

　　**8. 以下关于final关键字说法错误的是(A,C)(两项)**

　　A) final是java中的修饰符，可以修饰类、接口、抽象类、方法和属性

　　B) final修饰的类肯定不能被继承

　　C) final修饰的方法不能被重载

　　D) final修饰的变量不允许被再次赋值
 

　　**9. 访问修饰符作用范围由大到小是(D)**

　　A) private-default-protected-public

　　B) public-default-protected-private

　　C) private-protected-default-public

　　D) public-protected-default-private
 

　　**10. 以下(D)不是Object类的方法**

　　A) clone()

　　B) finalize()

　　C) toString()

　　D) hasNext()
 

　　**11. 多态的表现形式有(A)**

　　A) 重写

　　B) 抽象

　　C) 继承

　　D) 封装
 

　　**12. 以下对重载描述错误的是(B)**

　　A) 方法重载只能发生在一个类的内部

　　B) 构造方法不能重载

　　C) 重载要求方法名相同，参数列表不同

　　D) 方法的返回值类型不是区分方法重载的条件
 

　　**13. 以下(D)添加到ComputerBook中不会出错**

```java
class Book{

    protected int getPrice(){

        return 30;

    }

}

public class ComputerBook extends Book{

}
```

　　A) protected float getPrice(){}

　　B) protected int getPrice(int page){}

　　C) int getPrice(){}

　　D) public int getPrice(){return 10;}
 

　　**14. 以下对抽象类的描述正确的是(C)**

　　A) 抽象类没有构造方法

　　B) 抽象类必须提供抽象方法

　　C) 有抽象方法的类一定是抽象类

　　D) 抽象类可以通过new关键字直接实例化
 

　　**15. 以下对接口描述错误的有(D)**

　　A) 接口没有提供构造方法

　　B) 接口中的方法默认使用public、abstract修饰

　　C) 接口中的属性默认使用public、static、final修饰

　　D) 接口不允许多继承
 

　　**16. 以下代码，描述正确的有(A)**

```java
interface IDemo{

    public static final String name; // 1

        void print(); // 2

        public void getInfo(); // 3

}

abstract class Person implements IDemo{ // 4

    public void print(){
	}
}
```



　　

　　A) 第1行错误，没有给变量赋值

　　B) 第2行错误，方法没有修饰符

　　C) 第4行错误，没有实现接口的全部方法

　　D) 第3行错误，没有方法的实现
 

　　**17. 接口和抽象类描述正确的有(B,C)(两项)**

　　A) 抽象类没有构造函数

　　B) 接口没有构造函数

　　C) 抽象类不允许多继承

　　D) 接口中的方法可以有方法体
 

　　**18. 以下描述错误的有(C)**

　　A) abstract 可以修饰类、接口、方法

　　B) abstract修饰的类主要用于被继承

　　C) abstract 可以修饰变量

　　D) abstract修饰的类，其子类也可以是abstract修饰的
 

　　**19. 以下描述正确的有(B)**

　　A) 方法的重写应用在一个类的内部

　　B) 方法的重载与返回值类型无关

　　C) 构造方法不能重载

　　D) 构造方法可以重写
 

### 　　**20. 以下程序运行结果是(A)** 🌑

```java
public class Test extends Father{

    private String name=”test”;

    public static void main(String[] args){

        Test test = new Test();

        System.out.println(test.getName());

    }

}

class Father{

    private String name=”father”;

    public String getName() {

        return name;

    }

}
```

　　A) father

　　B) test

　　C) 编译出错

　　D) 运行出错，无输出
 

　　**21. 以下对异常的描述不正确的有(C)**

　　A) 异常分为Error和Exception

　　B) Throwable是所有异常类的父类

　　C) Exception是所有异常类父类

　　D) Exception包括RuntimeException和RuntimeException之外的异常
 

　　**22. 在try-catch-finally语句块中，以下可以单独与finally一起使用的是(B)**

　　A) catch

　　B) try

　　C) throws

　　D) throw
 

### 　　**23. 下面代码运行结果是(B)** 🌑

```java
public class Demo{

    public int add(int a,int b){

        try{
            return a+b;
        }catch(Exception e){
            System.out.println(“catch 语句块”);
        }finally{
            System.out.println(“finally 语句块”);
        }
        return 0;
    }

    public static void main(String[] args){

        Demo demo = new Demo();

        System.out.println(“和是：”+demo.add(9,34));
    }
}
```

　　A) 编译异常

　　B) finally语句块 和是：43

　　C) 和是：43 finally语句块

　　D) catch语句块 和是：43
 

### 　　**24. 以下描述不正确的有(D)** 🌑

　　A) try块不可以省略

　　B) 可以使用多重catch块

　　C) finally块可以省略

　　D) catch块和finally块可以同时省略
 

　　**25. 以下对自定义异常描述正确的是(C)**

　　A) 自定义异常必须继承Exception

　　B) 自定义异常可以继承自Error

　　C) 自定义异常可以更加明确定位异常出错的位置和给出详细出错信息

　　D) 程序中已经提供了丰富的异常类，使用自定义异常没有意义
 

### 　　**26. 以下程序运行结果是(D)** 🌑

```java
public class Test {

    public int div(int a, int b) {

        try {
            return a / b;
        }catch(Exception e){
            System.out.println(“Exception”);
        }catch(NullPointerException e){
            System.out.println(“ArithmeticException”);
        }catch (ArithmeticException e) {
            System.out.println(“ArithmeticException”);
        }finally {
            System.out.println(“finally”);
        }
        return 0;
    }

    public static void main(String[] args) {

        Test demo = new Test();

        System.out.println(“商是：” + demo.div(9, 0));

    }

}
```

　　A) Exception finally 商是：0

　　B) ArithmeticException finally 商是：0

　　C) finally商是：0

　　D) 编译报错
 

### 　　**27. 以下对TCP和UDP描述正确的是(D)** 🌑

　　A) TCP不能提供数据的可靠性

　　B) UDP能够保证数据库的可靠性

　　C) TCP数据传输效率高于UDP

　　D) UDP数据传输效率高于TCP
 

　　**28. 在Java中，下面对于构造函数的描述正确的是(D)。(选择一项)**

　　A) 类必须显示定义构造函数

　　B) 构造函数的返回类型是void

　　C) 构造函数和类有相同的名称，并且不能带任何参数

　　D) 一个类可以定义多个构造函数
 

　　**29. 根据下面的代码，**

　　String s = null;

　　会抛出NullPointerException异常的有(A,C)。[两项]

　　A) if( (s!=null) & (s.length()>0) )

　　B) if( (s!=null) & & (s.length()>0) )

　　C) if( (s==null) | (s.length()==0) )

　　D) if( (s==null) || (s.length()==0) )
 

　　**30. .在Java中，关于HashMap类的描述，以下廉洁错误的是( B )。**

　　A) HashMap使用键/值得形式保存数据

　　B) HashMap 能够保证其中元素的顺序

　　C) HashMap允许将null用作键

　　D) HashMap允许将null用作值
 

　　**31. 下列选项中关于java中super关键字的说法错误的是( B )**

　　A) super关键字是在子类对象内部指代其父类对象的引用

　　B) super关键字不仅可以指代子类的直接父类，还可以指代父类的父类

　　C) 子类可以通过super关键字调用父类的方法

　　D) 子类可以通过super关键字调用父类的属性
 

　　**32. 在Java中，以下代码( A )正确地创建了一个InputStreamReader对象。**

　　A) InuptStreamReader(new FileInputStream(“1.dat”));

　　B) InuptStreamReader(new FileReader(“1.dat”));

　　C) InuptStreamReader(new BufferReader(“1.dat”));

　　D) InuptStreamReader (“1.dat”);
 

　　**33. 在Java中，( D )类提供定位本地文件系统，对文件或目录及其属性进行基本操作。**

　　A) FileInputStream

　　B) FileReader

　　C) FileWriter

　　D) File
 

　　**34. Java中的集合类包括ArrayList、LinkedList、HashMap等类，下列关于集合类描述错误的是(C)(选择一项)**

　　A) ArrayList和LinkedList均实现了List接口

　　B) ArrayList的访问速度比LinkedList快

　　C) 添加和删除元素时，ArrayList的表现更佳

　　D) HashMap实现Map接口，它允许任何类型的键和值对象，并允许将null用作键或值
 

　　**35. 在Java中开发JDBC应用程序时，使用DriverManager类的getConnection()方法建立与数据源的连接语句为：**

　　Connection con = DriverManager.getConnection(“jdbc:odbc:news”);

　　URL连接中的“news”表示的是(C)(选择一项)

　　A) 数据库中表的名称

　　B) 数据库服务器的机器名

　　C) 数据源的名称

　　D) 用户名
 

　　**36. 在Java中,JDBCAPI定义了一组用于与数据库进行通信的接口和类，它们包括在(B)包中。**

　　A) java.lang

　　B) java.sql

　　C) java.util

　　D) java.math
 

　　**37. Java中，以下( B )接口以键_值对的方式存储对象。**

　　A) java.util.Collection

　　B) java.util.Map

　　C) java.util.List

　　D) java.util.Set
 

　　**38. 以下关于对象序列化描述正确的是( C,D )[两项]**

　　A) 使用FileOutputStream可以将对象进行传输

　　B) 使用PrintWriter可以将对象进行传输

　　C) 使用ObjectOutputStream类完成对象存储，使用ObjectInputStream类完成对象读取

　　D) 对象序列化的所属类需要实现Serializable接口
 

　　**39. 在Java中，( A )类可用于创建链表数据结构的对象。**

　　A) LinkedList

　　B) ArrayList

　　C) Collection

　　D) HashMap
 

　　**40. 分析下面这段Java代码，它的运行结果是( C )。**

　　import java.io.*;

　　public class B{

　　public static void main(string [] args){

　　int i=12;

　　System.out.println(i+=i-=i*=i);

　　}

　　}

　　A) 100

　　B) 0

　　C) -120

　　D) 程序无法编译
 

　　**41. 使用JDBC事务的步骤是(C,A,B,D)(多选)**

　　A) 取消Connection的事务自动提交方式

　　B) 发生异常回滚事务

　　C) 获取Connection对象

　　D) 操作完毕提交事务
 

　　**42. 以下对JDBC事务描述错误的是( B )**

　　A) JDBC事务属于JAVA事务的一种

　　B) JDBC事务属于容器事务类型

　　C) JDBC事务可以保证操作的完整性和一致性

　　D) JDBC事务是由Connection发起的，并由Connection控制
 

　　**43. 要通过可滚动的结果集更新数据，以下正确的是(A**

　　A) pst=con.prepareStatement(sql, ResultSet.TYPE_SCROLL_SENSITIVE,ResultSet.CONCUR_UPDATABLE)

　　B) pst=con.prepareStatement(sql, ResultSet.TYPE_SCROLL_SENSITIVE,ResultSet.CONCUR_READ_ONLY)

　　C) pst=con.prepareStatement(sql, Resu ltSet.TYPE_SCROLL_SENSITIVE)

　　D) pst=con.prepareStatement(sql, ResultSet.CONCUR_UPDATABLE)
 

　　**44. 存储过程pro有两个参数，第一个为输入参数，第二个为输出参数，以下代码正确的是(C)**

　　A) CallableStatement cst=con.prepareCall(“(call pro(?,?))”);

　　B) CallableStatement cst=con.prepareCall(“(call pro(?))”);

　　C) CallableStatement cst=con.prepareCall(“{call pro(?,?)}”);

　　D) CallableStatement cst=con.prepareCall(“{call pro(?,?,?)}”);
 

　　**45. 以下描述正确的是(B)**

　　A) CallableStatement是PreparedStatement的父接口

　　B) PreparedStatement是CallableStatement的父接口

　　C) CallableStatement是Statement的子接口

　　D) PreparedStatement是Statement的父接口
 

　　**46. 要删除book表中书籍(bookName)是”java”的记录，以下代码正确的是(A)**

　　String sql=”delete from book where bookName=?”;

　　PreparedStatement pst=con.preparedStatement(sql);

　　______________________________

　　pst.execute();

　　A) pst.setString(1,”java”);

　　B) pst.setString(0,”java”);

　　C) pst.setInt(0,”java”);

　　D) 以上选项都不正确
 

　　**47. 获取ResutlSet对象rst的第一行数据，以下正确的是(B)**

　　A) rst.hashNext();

　　B) rst.next();

　　C) rst.first();

　　D) rst.nextRow();
 

　　**48. 以下可以正确获取结果集的有(AD)(多选)**

　　A) Statement sta=con.createStatement();

　　ResultSet rst=sta.executeQuery(“select * from book”);

　　B) Statement sta=con.createStatement(“select * from book”);

　　ResultSet rst=sta.executeQuery();

　　C) PreparedStatement pst=con.preparedStatement();

　　ResultSet rst=pst.executeQuery(“select * from book”);

　　D) PreparedStatement pst=con.preparedStatement(“select * from book”);

　　ResultSet rst=pst.executeQuery();
 

　　**49. 以下负责建立与数据库连接的是(D)**

　　A) Statement

　　B) PreparedStatement

　　C) ResultSet

　　D) DriverManager
 

　　**50. 使用JDBC连接数据库的顺序是(B,A,D,C,E)(多选)**

　　A) 加载驱动

　　B) 导入驱动包

　　C) 发送并处理SQL语句

　　D) 建立于数据库的连接

　　E 关闭连接
 

## **二、简答题(各5分，共25分)**

　　**1、在java中如果声明一个类为final，表示什么意思? (不计分)**

　　答：final是最终的意思，final可用于定义变量、方法和类但含义不同，声明为final的类不能被继承。
 

　　**1、父类的构造方法是否可以被子类覆盖(重写)?**

　　答：父类的构造方法不可以被子类覆盖，因为父类和子类的类名是不可能一样的。
 

　　**2、请讲述String 和StringBuffer的区别。**

　　答：String 类所定义的对象是用于存放”长度固定”的字符串。

　　StringBuffer类所定义的对象是用于存放”长度可变动”的字符串。
 

　　**3、如果有两个类A、B(注意不是接口)，你想同时使用这两个类的功能，那么你会如何编写这个C类呢?**

　　答：因为类A、B不是接口，所以是不可以直接继承的，但可以将A、B类定义成父子类，那么C类就能实现A、B类的功能了。假如A为B的父类，B为C的父类，此时C就能实现A、B的功能。
 

　　**4、分析sleep()和wait()方法的区别。**

　　答： sleep睡眠的意思 : sleep() 方法用来暂时中止执行的线程。在睡眠后，线程将进入就绪状态。

　　wait等待的意思: 如果调用了 wait() 方法，线程将处于等待状态。用于在两个或多个线程并发运行时。
 

　　**5、谈谈你对抽象类和接口的理解。**

　　答：定义抽象类的目的是提供可由其子类共享的一般形式、子类可以根据自身需要扩展抽象类、抽象类不能实例化、抽象方法没有函数体、抽象方法必须在子类中给出具体实现。他使用extends来继承。

　　接口：一个接口允许一个类从几个接口继承而来，Java 程序一次只能继承一个类但可以实现几个接口，接口不能有任何具体的方法，接口也可用来定义可由类使用的一组常量。其实现方式是interface来实现。