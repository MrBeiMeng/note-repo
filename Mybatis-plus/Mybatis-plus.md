# MyBatisPlus概述

[(35条消息) 狂神说 MyBatisPlus 学习笔记_DDDDeng_的博客-CSDN博客_狂神说mybatisplus笔记](https://blog.csdn.net/DDDDeng_/article/details/107707545)

需要的基础：把我的MyBatis、Spring、SpringMVC就可以学习这个了！ 为什么要学习它呢？MyBatisPlus可以节省我们大量工作时间，所有的CRUD代码它都可以自动化完成！ JPA 、 tk-mapper、MyBatisPlus 偷懒的！

> 简介

是什么？ MyBatis 本来就是简化 JDBC 操作的！ 官网：https://mp.baomidou.com/ MyBatis Plus，简化 MyBatis

![image-20211123110455432](https://i.loli.net/2021/11/23/hKEY3HRDArsdJOC.png)

# 快速入门

文档：https://mp.baomidou.com/

使用第三方组件：

1. 导入对应依赖
2. 研究依赖如何配置
3. 代码如何编写
4. 提高扩展技术能力

> 步骤

1、创建数据库 `mybatis_plus`

2、创建user表DROP TABLE IF EXISTS user;

```sql
CREATE TABLE user
(
	id BIGINT(20) NOT NULL COMMENT '主键ID',
	name VARCHAR(30) NULL DEFAULT NULL COMMENT '姓名',
	age INT(11) NULL DEFAULT NULL COMMENT '年龄',
	email VARCHAR(50) NULL DEFAULT NULL COMMENT '邮箱',
	PRIMARY KEY (id)
);
-- 真实开发中，version(乐观锁)、deleted(逻辑删除)、gmt_create、gmt_modified
```

3、创建项目springboot

4、导入依赖

```xml
<!--数据库驱动-->
<dependency>
    <groupId>mysql</groupId>
    <artifactId>mysql-connector-java</artifactId>
    <version>8.0.20</version>
</dependency>
<!--mybatis-plus-->
<dependency>
    <groupId>com.baomidou</groupId>
    <artifactId>mybatis-plus-boot-starter</artifactId>
    <version>3.0.5</version>
</dependency>
<!--lombok-->
<dependency>
    <groupId>org.projectlombok</groupId>
    <artifactId>lombok</artifactId>
    <optional>true</optional>
</dependency>

```

说明：我们使用mybatis-plus可以节省我们大量的代码，尽量不要同时导入mybatis和mybatis-plus！

5、连接数据库，这一步和mybatis相同

```properties
# mysql
spring.datasource.username=root
spring.datasource.password=root
spring.datasource.url=jdbc:mysql://localhost:3306/mybatis_plus?userSSL=true&useUnicode=true&characterEncoding=UTF-8&serverTimezone=UTC
spring.datasource.driver-class-name=com.mysql.cj.jdbc.Driver

```

==6、传统方式：pojo-dao（连接mybatis，配置mapper.xml文件）- service - controller==

6、使用了mybatis-plus 之后

- pojo

```java
@Data
@AllArgsConstructor
@NoArgsConstructor
public class User {
    
    private Long id;
    private String name;
    private Integer age;
    private String email;
}
```

- mapper接口

```java
// 在对应的Mapper上面继承基本的接口 BaseMapper
@Repository // 代表持久层
public interface UserMapper extends BaseMapper<User> {
    // 所有的CRUD操作都已经基本完成了
    // 你不需要像以前的配置一大堆文件了
}
```

注意点：我们需要在主启动类上去扫描我们的mapper包下的所有接口`MapperScan("com.kuang.mapper")`

- 测试类中测试

  ```java
  @SpringBootTest
  class MybatisPlusApplicationTests {
  // 继承了BaseMapper,所有的方法都来自父类
  // 我们也可以编写自己的扩展方法
  @Autowired
  UserMapper userMapper;
  
  @Test
  void contextLoads() {
          // 参数是一个 Wrapper , 条件构造器，这里我们先不用 null
          // 查询全部用户
          List<User> users = userMapper.selectList(null);
          users.forEach(System.out::println);//相当于增强for循环
      }
  }
  ```

> 思考问题：

1、SQL是谁帮我们写的？MyBatis-Plus都写好了

2、方法哪里来的？MyBatis-Plus 都写好了

二、配置日志
我们所有的sqld现在都是不可见的，我们希望知道它是怎么执行的，所以我们必须要看日志！

# 配置日志

(系统自带的，控制台输出)我们所有的sql现在是不可见的，我们希望知道它是怎么执行的，所以我们必须要看日志

```
# 配置日志 mybatis-plus.configuration.log-impl=org.apache.ibatis.logging.stdout.StdOutImpl
```

![image-20211123121141641](https://i.loli.net/2021/11/23/cAzVsyoXOl3DQ5L.png)

配置完毕日志之后，后面的学习就需要注意这个自动生成的SQL，你们就会喜欢上 MyBatis-Plus！

# CRUD扩展

## 插入操作

> insert插入

```java
// 测试插入
@Test
public void testInsert(){
    User user = new User();
    user.setName("枫语");
    user.setAge(3);
    user.setEmail("1789164015@qq.com");
    int result = userMapper.insert(user); // 帮我们自动生成id
    System.out.println(result); // 受影响的行数
    System.out.println(user); // 发现，id会自动回填
}
```

![image-20211123140602988](https://i.loli.net/2021/11/23/PlVb6hSf8NMc4io.png)

> 数据库插入的id的默认值为：全局的唯一id

## 主键生成策略

> 默认 ID_WORKER 全局唯一id

分布式系统唯一id生成：https://www.cnblogs.com/haoxinyue/p/5208136.html

**雪花算法：**

snowflake是Twitter开源的分布式ID生成算法，结果是一个long型的ID。其核心思想是：使用41bit作为毫秒数，10bit作为机器的ID（5个bit是数据中心，5个bit的机器ID），12bit作为毫秒内的流水号（意味着每个节点在每毫秒可以产生 4096 个 ID），最后还有一个符号位，永远是0。可以保证几乎全球唯一！

> 主键自增

我们需要配置主键自增：

- 实体类字段上`@TableId(type=IdType.AUTO)`
- 数据库字段一定要是自增的

> 其与的源码解释

```java
public enum IdType {    AUTO(0), // 数据库id自增    NONE(1), // 未设置主键    INPUT(2), // 手动输入    ID_WORKER(3), // 默认的全局唯一id    UUID(4), // 全局唯一id uuid    ID_WORKER_STR(5); //ID_WORKER 字符串表示法}
```

## 更新操作

```java
// 测试更新@Test    public void testUpdate(){    User user = new User();    // 通过条件自动拼接动态sql    user.setId(6L);    user.setName("关注公众号：狂神说");    user.setAge(18);    // 注意：updateById 但是参数是一个 对象！    int i = userMapper.updateById(user);    System.out.println(i);}
```

![image-20211123142433496](https://i.loli.net/2021/11/23/xj1nORvloEKqwgQ.png)

## 自动填充

创建时间、修改时间！这些个操作一般都是自动化完成的，我们不希望手动更新！

**阿里巴巴开发手册：**所有的数据库表：gmt_create、gmt_modified几乎所有的表都要配置上！而且需要自动化！

> 方式一：数据库级别（工作中不允许修改数据库）

1、在表中新增字段 create_time, update_time

![image-20211124111320499](https://i.loli.net/2021/11/24/mqgvbQRONax91D5.png)

2、再次测试插入方法，我们需要先把实体类同步！

```java
private Date createTime; private Date updateTime;
```

3、再次更新查看结果即可

![image-20211124111355807](https://i.loli.net/2021/11/24/GU8BqpoVLDleYwn.png)

> 方式二：代码级别

1、删除数据库的默认值、更新操作！

![image-20211124111452453](https://i.loli.net/2021/11/24/aUJiBKhW5FZ3edq.png)

2、实体类字段属性上需要增加注解

```java
//字段自动填充内容@TableField(fill= FieldFill.INSERT)private Date createTime;@TableField(fill= FieldFill.INSERT_UPDATE)private Date updateTime;
```

3、编写处理器来处理这个注解即可！

```java
@Slf4j@Component//一定不要忘记吧处理器加到IOC容器中public class MyMetaObjectHandler implements MetaObjectHandler {    //插入时候的填充策略    @Override    public void insertFill(MetaObject metaObject) {        log.info("开始插入0");        this.setFieldValByName("createTime",new Date(),metaObject);        this.setFieldValByName("updateTime",new Date(),metaObject);    }    //更新的填充策略    @Override    public void updateFill(MetaObject metaObject) {        log.info("开始更新");        this.setFieldValByName("updateTime",new Date(),metaObject);    }}
```

4、测试插入 5、测试更新、观察时间即可！

## 乐观锁

在面试过程中，我们经常会被问道乐观锁，悲观锁！这个其实非常简单！

> 乐观锁 : 故名思意十分乐观，它总是认为不会出现问题，无论干什么不去上锁！如果出现了问题， 再次更新值测试
>
> 悲观锁：故名思意十分悲观，它总是认为总是出现问题，无论干什么都会上锁！再去操作！

![image-20211124183534152](https://i.loli.net/2021/11/24/chluLQ87gMxHjdp.png)

乐观锁实现方式： 取出记录时，获取当前 version

- 更新时，带上这个version

- 执行更新时， set version = newVersion where version = oldVersion

- 如果version不对，就更新失败

```sql
乐观锁：先查询，获得版本号-- Aupdate user set name = "wsk",version = version+1 where id = 1 and version = 1-- B  （B线程抢先完成，此时version=2，会导致A线程修改失败！）update user set name = "wsk",version = version+1 where id = 1 and version = 1
```

> 测试一下Mybatis-Plus乐观锁插件

1、给数据库中增加version字段

![image-20211124120524908](https://i.loli.net/2021/11/24/qhHCbiazseFGJ2V.png)

2、实体类加对应的字段

```
@Version//乐观锁version注解private Integer version;
```

3、注册组件

```java
//扫描mapper文件夹@MapperScan("com.kuang.mapper")@EnableTransactionManagement@Configurationpublic class MyBatisPlusConfig {    // 注册乐观锁插件    @Bean    public OptimisticLockerInterceptor optimisticLockerInterceptor() {        return new OptimisticLockerInterceptor();    }}
```

4、测试一下

```java
@Test//测试乐观锁成功public void testOptimisticLocker1(){    //1、查询用户信息    User user = userMapper.selectById(1L);    //2、修改用户信息    user.setAge(18);    user.setEmail("2803708553@qq.com");    //3、执行更新操作    userMapper.updateById(user);}
```

![image-20211124120630204](https://i.loli.net/2021/11/24/oBZ9CaJycOfRzlg.png)

![image-20211124120639443](https://i.loli.net/2021/11/24/xTntPgbkClHOKMF.png)

- 失败

- ```java
  @Test//测试乐观锁失败  多线程下public void testOptimisticLocker2(){    //线程1    User user1 = userMapper.selectById(1L);    user1.setAge(1);    user1.setEmail("2803708553@qq.com");    //模拟另外一个线程执行了插队操作    User user2 = userMapper.selectById(1L);    user2.setAge(2);    user2.setEmail("2803708553@qq.com");    userMapper.updateById(user2);    //自旋锁来多次尝试提交！    userMapper.updateById(user1);//如果没有乐观锁就会覆盖插队线程的值}
  ```

![image-20211124120738909](https://i.loli.net/2021/11/24/2DzfCWilEnu3s56.png)

![image-20211124121236186](https://i.loli.net/2021/11/24/IUjMdBpLeT25Gvi.png)

## 查询操作

```java
//测试查询@Testpublic void testSelectById(){    User user = userMapper.selectById(1L);    System.out.println(user);}//通过多个id查询用户@Testpublic void testSelectBatchId(){    List<User> users = userMapper.selectBatchIds(Arrays.asList(1, 2, 3));    users.forEach(System.out::println);}//条件查询map@Testpublic void testSelectByBatchIds(){    HashMap<String,Object> map=new HashMap<>();    //自定义要查询    map.put("name","Tom");    map.put("age",3);    List<User> users = userMapper.selectByMap(map);    users.forEach(System.out::println);}
```

## 分页查询

分页在网站的使用十分之多！

1、原始的limit分页

2、pageHelper第三方插件

3、MybatisPlus其实也内置了分页插件！

> 如何使用：

1、配置拦截器组件

```java
//分页插件@Beanpublic PaginationInterceptor paginationInterceptor() {    return new PaginationInterceptor();}
```

2、直接使用page对象即可

```java
@Test//测试分页查询public void testPage(){    //参数一current：当前页   参数二size：页面大小    //使用了分页插件之后，所有的分页操作都变得简单了    Page<User> page = new Page<>(2,5);    userMapper.selectPage(page,null);    page.getRecords().forEach(System.out::println);    System.out.println("总页数==>"+page.getTotal());}
```

![image-20211124160230104](https://i.loli.net/2021/11/24/4PmCwrIlkHTGWdb.png)

## 删除操作

> 基本的删除任务：

![image-20211124161514145](https://i.loli.net/2021/11/24/lvQxsTnihYyu6MB.png)

```java
//测试删除@Testpublic void testDeleteById(){    userMapper.deleteById(1463336868539412482L);    }//批量删除@Testpublic void testDeleteBatchId(){    userMapper.deleteBatchIds(Arrays.asList(1463336868539412481L,6L));}//通过map删除@Testpublic void testDeleteMap(){    HashMap<String,Object> map=new HashMap<>();    map.put("name","放地了口感");    userMapper.deleteByMap(map);}
```

## 逻辑删除

> 物理删除：从数据库中直接删除

> 逻辑删除：在数据库中没有被删除，而是通过一个变量来使他失效！ deleted=0 ==> deleted=1

**管理员可以查看被删除的记录！防止数据的丢失，类似于回收站！**

***测试一下：\***

> 1、在数据表中增加一个deleted字段

![image-20211124163911620](https://i.loli.net/2021/11/24/tuYv8LXSxmKVOca.png)

> 2、实体类中添加对应属性

```
@TableLogic//逻辑删除注解private Integer deleted;
```

> 3、配置！

```java
//逻辑删除组件@Beanpublic ISqlInjector sqlInjector(){    return new LogicSqlInjector();}
```

1. `#配置逻辑删除  没删除的为0 删除的为1`
2. `mybatis-plus.global-config.db-config.logic-delete-value=1`
3. `mybatis-plus.global-config.db-config.logic-not-delete-value=0`

> 4、测试一下删除

![image-20211124164002663](https://i.loli.net/2021/11/24/cANiuy3HmEOh4B6.png)

发现： 记录还在，deleted变为1

再次测试查询被删除的用户，发现查询为空

![image-20211124164017774](https://i.loli.net/2021/11/24/ygXVfILNvTPdebM.png)

## 性能分析插件

我们在平时的开发中，会遇到一些满Sql。测试、druid···

MybatisPlus也提供了性能分析插件，如果超过这个时间就停止运行！

**性能分析拦截器作用：用于输出每条sql语句及其执行时间**

1、导入插件

```java
 //性能分析插件@Bean@Profile({"dev","test"})//设置dev开发、test测试 环境开启  保证我们的效率public PerformanceInterceptor performanceInterceptor(){    PerformanceInterceptor performanceInterceptor = new PerformanceInterceptor();    performanceInterceptor.setMaxTime(100);//设置sql最大执行时间*ms，如果超过了则不执行    performanceInterceptor.setFormat(true);//开启sql格式化    return performanceInterceptor;}
```

**注意： 要在SpringBoot中配置环境为dev或test环境！**

```
#设置开发环境spring.profiles.active=dev
```

```java
2、测试使用@Testvoid contextLoads() {    //参数是一个wrapper ，条件构造器，这里我们先不用 null    //查询全部的用户    List<User> userList = userMapper.selectList(null);    userList.forEach(System.out::println);}
```

![image-20211125132527383](https://i.loli.net/2021/11/25/ThfYRqH53yVpBtr.png)

使用性能分析插件，可以帮助我们提高效率！、

## 条件构造器

![image-20211126173027860](https://i.loli.net/2021/11/26/UkVqpIQNROJ7ucv.png)

![image-20211126173049810](https://i.loli.net/2021/11/26/7AxhlmfpQjoYBsV.png)

**十分重要：Wrapper** 记住查看输出的SQL进行分析

```java
1、测试一@Testpublic void testWrapper1() {    //参数是一个wrapper ，条件构造器，和刚才的map对比学习！    //查询name不为空，email不为空，age大于18的用户    QueryWrapper<User> wrapper = new QueryWrapper<>();    wrapper        .isNotNull("name")        .isNotNull("email")        .ge("age",18);    List<User> userList = userMapper.selectList(wrapper);    userList.forEach(System.out::println);}测试二@Testpublic void testWrapper2() {    //查询name=wsk的用户    QueryWrapper<User> wrapper = new QueryWrapper<>();    wrapper.eq("name","wsk");    //查询一个数据selectOne，若查询出多个会报错    //Expected one result (or null) to be returned by selectOne(), but found: *    //若出现多个结果使用list或map    User user = userMapper.selectOne(wrapper);//查询一个数据，若出现多个结果使用list或map    System.out.println(user);}测试三@Testpublic void testWrapper3() {    //查询age在10-20之间的用户    QueryWrapper<User> wrapper = new QueryWrapper<>();    wrapper.between("age", 10, 20);//区间    Integer count = userMapper.selectCount(wrapper);//输出查询的数量selectCount    System.out.println(count);}测试四@Testpublic void testWrapper4() {    //模糊查询    QueryWrapper<User> wrapper = new QueryWrapper<>();    wrapper        .notLike("name","s")        .likeRight("email","t");//qq%  左和右？    List<Map<String, Object>> maps = userMapper.selectMaps(wrapper);    maps.forEach(System.out::println);}测试五@Testpublic void testWrapper5() {    //模糊查询    // SELECT id,name,age,email,version,deleted,create_time,update_time     //FROM user     //WHERE deleted=0 AND id IN     //(select id from user where id<5)    QueryWrapper<User> wrapper = new QueryWrapper<>();    //id 在子查询中查出来    wrapper.inSql("id","select id from user where id<5");    List<Object> objects = userMapper.selectObjs(wrapper);    objects.forEach(System.out::println);}测试六@Testpublic void testWrapper6() {    QueryWrapper<User> wrapper = new QueryWrapper<>();    //通过id进行降序排序    wrapper.orderByDesc("id");    List<User> userList = userMapper.selectList(wrapper);    userList.forEach(System.out::println);}Mysql => JDBC => Mybatis => MybatisPlus
```

# 代码自动生成器

`AutoGenerator` 是 MyBatis-Plus 的代码生成器，通过 `AutoGenerator` 可以快速生成 Entity、Mapper、Mapper XML、Service、Controller 等各个模块的代码，极大的提升了开发效率。

```java
package com.wsk;import com.baomidou.mybatisplus.annotation.DbType;import com.baomidou.mybatisplus.annotation.FieldFill;import com.baomidou.mybatisplus.annotation.IdType;import com.baomidou.mybatisplus.generator.AutoGenerator;import com.baomidou.mybatisplus.generator.config.DataSourceConfig;import com.baomidou.mybatisplus.generator.config.GlobalConfig;import com.baomidou.mybatisplus.generator.config.PackageConfig;import com.baomidou.mybatisplus.generator.config.StrategyConfig;import com.baomidou.mybatisplus.generator.config.po.TableFill;import com.baomidou.mybatisplus.generator.config.rules.DateType;import com.baomidou.mybatisplus.generator.config.rules.NamingStrategy;import java.util.ArrayList;//代码自动生成器public class WskCode {    public static void main(String[] args) {        //我们需要构建一个代码生成器对象        AutoGenerator mpg = new AutoGenerator();        //怎么样去执行，配置策略        //1、全局配置        GlobalConfig gc = new GlobalConfig();        String projectPath = System.getProperty("user.dir");//获取当前目录        gc.setOutputDir(projectPath+"/src/main/java");//输出到哪个目录        gc.setAuthor("wsk");        gc.setOpen(false);        gc.setFileOverride(false);//是否覆盖        gc.setServiceName("%sService");//去Service的I前缀        gc.setIdType(IdType.ID_WORKER);        gc.setDateType(DateType.ONLY_DATE);        gc.setSwagger2(true);        mpg.setGlobalConfig(gc);        //2、设置数据源        DataSourceConfig dsc = new DataSourceConfig();        dsc.setUsername("root");        dsc.setPassword("root");        dsc.setUrl("jdbc:mysql://localhost:3306/wuye?useSSL=false&serverTimezone=GMT%2B8&useUnicode=true&characterEncoding=utf-8");        dsc.setDriverName("com.mysql.cj.jdbc.Driver");        dsc.setDbType(DbType.MYSQL);        mpg.setDataSource(dsc);        //3、包的配置        PackageConfig pc = new PackageConfig();        pc.setModuleName("study");        pc.setParent("com.wsk");        pc.setEntity("pojo");        pc.setMapper("mapper");        pc.setService("service");        pc.setController("controller");        mpg.setPackageInfo(pc);        //4、策略配置        StrategyConfig strategy = new StrategyConfig();        strategy.setInclude("admin","danyuan","building","room");//设置要映射的表名,只需改这里即可        strategy.setNaming(NamingStrategy.underline_to_camel);        strategy.setColumnNaming(NamingStrategy.underline_to_camel);        strategy.setEntityLombokModel(true);//是否使用lombok开启注解        strategy.setLogicDeleteFieldName("deleted");        //自动填充配置        TableFill gmtCreate = new TableFill("gmt_create", FieldFill.INSERT);        TableFill gmtUpdate = new TableFill("gmt_update", FieldFill.INSERT_UPDATE);        ArrayList<TableFill> tableFills = new ArrayList<>();        tableFills.add(gmtCreate);        tableFills.add(gmtUpdate);        strategy.setTableFillList(tableFills);        //乐观锁配置        strategy.setVersionFieldName("version");        strategy.setRestControllerStyle(true);//开启驼峰命名        strategy.setControllerMappingHyphenStyle(true);//localhost:8080/hello_id_2        mpg.setStrategy(strategy);        mpg.execute();//执行    }}
```

```xml
<!--模板引擎 依赖:mybatis-plus代码生成的时候报异常--><dependency>    <groupId>org.apache.velocity</groupId>    <artifactId>velocity-engine-core</artifactId>    <version>2.0</version></dependency><!--配置ApiModel在实体类中不生效--><dependency>    <groupId>com.spring4all</groupId>    <artifactId>spring-boot-starter-swagger</artifactId>    <version>1.5.1.RELEASE</version></dependency><!--freemarker--><dependency>    <groupId>org.freemarker</groupId>    <artifactId>freemarker</artifactId>    <version>2.3.30</version></dependency><!--beetl--><dependency>    <groupId>com.ibeetl</groupId>    <artifactId>beetl</artifactId>    <version>3.3.2.RELEASE</version></dependency>
```

