# Spark 笔记

**Spark 是一种基于内存的快速、通用、可扩展的大数据分析计算引擎**

2009 - 至今



## 前置知识



### Hadoop HDFS 知识补充

架构:

  ![image-20220917112209545](https://ccurj.oss-cn-beijing.aliyuncs.com/mac-typoraimage-20220917112209545.png)



**常用shell 命令**

> 命令行中 `hadoop fs 具体命令` 等同于 `hadoop dfs 具体命令`

-cat 查看

-copyFromLocal -copyToLocal

-mv

-put

-rm

-rmdir







## 了解Spark



### 对比Hadoop

Hadoop 是一个批处理工具，由 java 语言编写

Spark 是一个批流一体的处理工具，基于内存，所以会比hadoop快

![image-20220915214022209](https://ccurj.oss-cn-beijing.aliyuncs.com/mac-typoraimage-20220915214022209.png)

**一次性数据计算例图**

> 框架在处理数据的时候，会从存储设备中读取数据，进行逻辑操作，然后将处理的结果重新存储到介质中。但是如此操作一旦遇到数据要重复修改，大量的IO操作便会影响整体的速度。Spark 其实就是在重复操作的过程中除去了保存到磁盘的那一步，直接放到内存进行处理。

![image-20220915200803410](https://ccurj.oss-cn-beijing.aliyuncs.com/mac-typoraimage-20220915200803410.png)



**不能完全替代 Hadoop**

- 在计算层面，Spark 相比较 Mr(MapReduce) 有巨大的性能优势，但至今仍然有许多计算工具基于 MR 框架，比如非常成熟的 Hive。
- Spark 仅做计算，而 Hadoop 生态圈不仅有计算( Mr ) 也有存储( HDFS ) 和 资源管理调度 (YARN) , HDFS和YARN仍是许多大多数体系的核心架构。





### 核心模块

![image-20220915201703053](https://ccurj.oss-cn-beijing.aliyuncs.com/mac-typoraimage-20220915201703053.png)

- **Spark Core** 提供最基础核心的功能，其他模块都在此基础上进行扩展
- **Spark SQL** 操作结构化数据的模块
- **Spark Streaming** 对实时数据/流式数据进行操作的模块
- **Spark MLlib** 机器学习操作
- **Spark Graphx** 图形挖掘计算



### **运行模式**

- 本地模式 (单机) - **Local**  一般用来开发和测试 

  > 本地模式就是以一个独立的进程，通过其内部的多个线程来模拟整个 Spark 运行时环境

- Standalone模式 (集群) - **HDFS**

  > Spark 中的各个角色以独立进程的形式存在，并组成Spark集群环境

- Hadoop YARN模式 (集群)

  > Spark 中的各个角色运行在Kubernetes的容器内部，并组成 Spark 集群环境

  

  Kubernetes 模式，云服务模式



### Spark 与 Yarn 架构角色对比

Spark 中由4类角色组成整个Spark 的运行时环境

- Master -- 【ResourceManager】集群管理者，分配整个集群的资源。
- Worker -- 【NodeManager】单个机器的管理者，负责在单个服务器上提供运行容器，管理当前机器的资源。
- Driver  单个 Spark任务的管理者，管理 Executor 的任务执行和任务分解分配，类似 YARN 的ApplicationMaster。
- Executor 具体执行任务的进程，Spark 的工作任务( Task )都由 Executor 来负责执行。



### 学习环境搭建

三台Linux 虚拟机服务器 - 有云的，不需要搭建

​	node1: Master(HDFS\YARN\Spark) 和 Worker (HDFS\YARN\Spark)

​	node2: Worker(HDFS\YARN\Spark)

​	node3:Worker(HDFS\YARN\Spark) 和 Hive

 网段192.168.88.0

搭建环境教程：【Spark python】https://www.bilibili.com/video/BV1Jq4y1z7VP?p=13&vd_source=daf77da3a61c83fe6e450deea76785f6



### PySpark 和 Spark

PySpark 是 Spark 的一种基于 python 的操作方式，程序调试好要上传到 Spark 去运行。



### 本机开发环境搭建

```python
pip install pyspark pyhive jieba
```

开发入口

```python
conf = SparkConf().setAppName(appName).setMaster(master)
sc = SparkContext(conf= conf)
```





### Spark On Yarn 模式搭建

**1 本质**

- Master 角色由 YARN 的 ResourceManage 担任。【管理】

- Worker 角色由 YARN 的 NodeManageer 担任。 【管理】

- Driver 角色运行在 YARN容器 内 或 提交任务的客户端进程。【任务计算】

- 真正干活的 Executor 运行在 YARN 提供的容器内。【运行任务】



**2 准备**

1. 需要 Yarn 集群
2. 需要 Spark 客户端工具，比如 spark-submit, 可以将Spark 程序提交到 YARN 中
3. 需要被提交的代码程序， 要自己写。



**3 环境搭建**

指路：👉[【黑马程序员Spark全套视频教程】](https://www.bilibili.com/video/BV1Jq4y1z7VP?p=24&vd_source=daf77da3a61c83fe6e450deea76785f6)

 

**4 两种部署模式**

Cluster 模式(运行效率高) 与 Client模式(可以查看日志)



### Spark 任务操作流程

![image-20220917135752430](https://ccurj.oss-cn-beijing.aliyuncs.com/mac-typoraimage-20220917135752430.png)



### Spark 广播变量

> 标记成 **广播变量** 可以避免一个对象被重复调用，占用 **io资源**



### Spark 累加器

> 因为数据是分区进行计算的，所以直接 count++ 可能导致加不到。7



### 内核调度



## RDD 

**—— 弹性分布式数据集**

> 分布式计算需要:分区控制，**Shuffle** 控制，数据存储｜序列化｜发送、数据计算API、等等。这些功能很难或无法通过 **Python** 内置的本地集合对象(List\字典)等去完成，在分布式框架中，需要一个统一的数据抽象对象，来实现分布式的功能，这个对象--就是 **RDD**
>
> **RDD** 是 **Spark** 中最基本的数据抽象，代表一个不可变、可分区，里面的元素可以并行计算的集合
>
> - **Dataset**：一个数据集合，用于存放数据
> - **Distributed**：**RDD** 中的数据是 **分布式存储** 的，可用于分布式计算
> - **Resilient**：**RDD** 中的数据可以存储在内存或磁盘中
> - - Key-Value 型的 RDD 可以有分区器
> - - RDD 的分区规划，会尽量靠近数据所在的服务器

![IMG_75C5ED0E79FC-1](https://ccurj.oss-cn-beijing.aliyuncs.com/mac-typoraIMG_75C5ED0E79FC-1.jpeg)



**SparkContext对象**

> Spark RDD 编程的程序入口对象。(不论何种语言)



**RDD** 的两种创建方式

- 通过并行化集合创建 即 本地对象 转 分布式RDD  *parallelize*

  `rdd = spark_context.parallelize(data, part_num)`

- 读取外部数据源 如 读取文件

  `rdd2 = spark_context.textFile(data_path, part_num)` *--part_num 不是100%生效的*

  *读取本地文件要加上：`file://`*

  `rdd3 = spark_context.wholeTextFiles(dir_path,part_num)` *此api应用于小文件特别多的场景*



**RDD 算子**

> 分布式集合对象上的 **api** 被称为 **算子**
>
> **算子** 分为 **转换算子** 和 **动作算子**
>
> - **Transformation 算子**
>
>   返回值仍然是 **rdd** ，这类算子是 **懒加载** 的，没有 **acion 算子** 将不会工作
>
> - **Action 算子**
>
>   返回值不是 **rdd**



### **常用转换算子**

- **map** 

  > 将 **RDD** 的数据 **分条进行处理** ，处理的逻辑 基于map算子中接收的处理函数，返回新的 **RDD**

  语法：

  ```python
  rdd.map(func)
  # func : f:(T) -> U
  # example: rdd.map(lambda x: x + 2)
  ```

- flatMap

  > 对 **RDD** 执行 **map操作** ，然后进行解除嵌套操作。
  >
  > 有时 **map** 中 **计算的结果可能为一个list对象**，结果就是 **map** 方法返回的数据是**嵌套了的**如 **[[1,2],[,3,4,5]]**，flatMap 负责**降低一个维度**，结果为 【1，2，3，4，5，6】

  ```python
  rdd.flatMap(lambda x: [x + 1, x + 2]).collect()
  # [2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8]
  ```

- reduceByKey

  > 针对 **KV型** 的 **RDD**，自动按照 **key** 分组，然后根据你提供的聚合逻辑，完成  **组内数据(value)** 的聚合操作

  ```python
  rdd,reduceByKey(func)
  # func : f:(V,V) -> V
  # rdd_key_value.reduceByKey(lambda value1, value2: value1 + value2).collect()
  #[('a', 1), ('a', 2), ('b', 2), ('c', 3), ('b', 5)] -> [('b', 7), ('c', 3), ('a', 3)]
  ```

- mapValues

  > 针对 二元元组RDD ，对其内部的 Value 进行 map 操作

  语法：

  ```python
  rdd.mapValues(func)
  # func : f:(V) -> U
  ```

- groupBy

  > 将 **rdd** 的数据进行分组，适用于 **元组数据** 

  语法：

  ```python
  rdd.groupBy(func)
  # func: f(T) -> K
  ```

- filter

  > 保留想要的数据，结果为 **TRUE** 则保留

  语法：

  ```python
  rdd.filter(func)
  # func: (T) -> bool
  # rdd.filter(lambda x: x % 2 == 0).collect()
  ```

- distinct

  > 对 RDD 进行去重，返回新 RDD

  语法：

  ```python
  rdd.distinct(参数1)
  # 参数1: 去重分区数量，一般不用传
  ```

- union

  > 两个 **RDD** 合并成1个 **RDD**，类型不同也可以合并

  语法：

  ```python
  rdd1.union(rdd2)
  ```

- join

  > 对两个 **RDD** 执行join操作(可以实现sql 的内/外连接) **只能用于二元组即KVRDD** ，就是通过 **key** 去join

  语法：

  ```python
  rdd1.join(rdd2)
  rdd1.leftOuterJoin(rdd2)
  rdd1.rightOuterJoin(rdd2)
  # [(1001,"zhangsan")]，[(1001,"销售部")] -> [(1001,("张三","销售部"))]
  ```

- intersection

  > 求两个 **RDD** 的交集，返回一个新 **RDD**

    语法：

    ```python
    rdd1.intersection(rdd2)
    ```

- glom

  > 将 **RDD** 通过分区进行嵌套

  语法:

  ```python
  rdd.glom()
  ```

- groupByKey

  > 自动按照 **key** 进行分组

  语法：

  ```python
  rdd.groupByKey()
  # [('a',1),('a',2)] -> [('a',(1,1))]
  ```

- sortBy

  > 对 RDD 数据进行排序，基于你指定的排序依据

  语法：

  ```python
  rdd.sortBy(func,ascending=False,numPartitioins=1)
  # func: (T) -> U 告知使用什么进行排序，如 lambda x: x[1]
  ```

- sortByKey

  > 针对 KV型RDD， 按照key 进行排序

  语法：

  ```python
  sourtByKey(ascending=True,numPartitions=None,keyfunc=...)
  # ascending 默认升序
  # numPartitions 按照几个分区进行排序如果全局有序，设置1
  # func: (K) -> U 可以对key进行一定的修饰操作，方便排序，同上
  ```



### 常用执行算子

- countByKey

  > 统计key出现的次数(一般适用于 KV型 RDD)

  ```python
  rdd.countByKey()
  ```

- collect 

  > 将各个分区的数据整合起来，统一到 Driver 中，形成一个 List 对象

  ```python
  rdd.collect()
  ```

- flod

  > 和 reduce 一样，接收传入逻辑进行聚合

  ```python
  rdd.fold(10,lambda a, b: a + b)
  # range(1,10) ->  85
  ```

- first

  > 取出 RDD 的第一个元素

  ```python
  rdd.first()
  ```

- takeSample

  > 随机抽样 RDD 的数据

- takeOrdered

  > 随机取 RDD排序后 的前N改个

- foreach

  > 对 RDD 的每一个元素执行你提供的操作，但是没有返回值

- saveAsTextFile

  > 将 RDD 的数据写入文本文件中

  ```python
  rdd.saveAsTextFile("hdfs://node1:8020/output/11111")
  ```

- mapPartitions

  > 传递一整个分区的数据

- foreachPartition

  > 一次处理一整个分区

- partitionBy

  > 对 RDD 进行自定义分区操作

- repartition

  > 对 RDD 进行重新分区



### RDD 是过程数据

> RDD 的数据是过程数据，只在处理的过程存在，一旦处理完成，就不见了。



### RDD 的缓存

> Add.cache() # 可以让

![IMG_0A2CF152530C-1](https://ccurj.oss-cn-beijing.aliyuncs.com/mac-typoraIMG_0A2CF152530C-1.jpeg)

### RDD 的CheckPoint

> 用于将 RDD 的数据保存起来。 **但是只允许磁盘存储**。与缓存不同的是，CheckPoint并不是分散存储，而是类似于saveAs.. 将数据聚合起来创建到一个地方。

![IMG_268D38AE8284-1](https://ccurj.oss-cn-beijing.aliyuncs.com/mac-typoraIMG_268D38AE8284-1.jpeg)





## Spark 案例练习



### 搜索引擎日志分析

> 使用搜狗实验室提供的用户查询日志数据，使用 Spark 框架，进行分析。

// todo





## Spark SQL

> SparkSQL 是非常成熟的 海量结构化数据处理框架
>
> - SparkSQL 本身十分优秀，支持SQL 语言\性能强\可以自动优化\API简单\兼容HIVE等
> - 企业大面积在使用 SparkSQL 处理业务数据
>
>  特点：
>
> 1. **融合** 可以和python代码无缝集成
> 2. **统一数据访问**，同一套 读写不同的数据源
> 3. **Hive 兼容** 可以使用 SparkSQL 直接计算并生成 **Hive** 数据表
> 4. **标准化链接** 支持标准化 JDBC \ ODBC 连接，方便和各种数据库进行数据交互 



### 数据抽象

- Pandas - DataFrame				二维表数据结构
- SparkCore - RDD                      无标准数据结构，存储什么数据均可
- SparkSQL - DataFrame。       二维表数据结构



### DataFrame VS RDD

|                  | RDD              | DataFrame                  |
| ---------------- | ---------------- | -------------------------- |
| 有分区的         | 是               | 是                         |
| 分布式的         | 是               | 是                         |
| 弹性的           | 是               | 是                         |
| 可保存的数据结构 | **任意数据结构** | **只能存储二维表结构数据** |

> 在结构方面: **DataFrame**[**python**] 中 **StructType** 对象描述整个 **DataFrame** 的表结构,**StructField** 对象描述一个列的信息
>
> ```python
>     struct_type = StructType(). \							# StructType
>         add("id", IntegerType(), False). \		# StructField
>         add("name", StringType(), True). \		# StructField
>         add("age", IntegerType(), False)			# StructField
> ```
>
> 



### sparkSession 执行对象

```python
spark_session = SparkSession.builder.appName("test").master("local[*]").getOrCreate()
# 获取context
spark_content = spark_session.sparkContext
```



### DataFrame 四种创建方式

- 通过 RDD 转换 1

  > ```python
  > # 通过RDD 转换
  > spark_context = spark_session.sparkContext
  > rdd = spark_context.textFile("file:///tmp/pycharm_project_93/data/stu_score.txt") \
  >     .map(lambda x: x.split(",")).map(lambda x: (x[0], int(x[1])))
  > # 构建 dataFrame
  > spark_session.createDataFrame(rdd, schema=['id', 'subject', 'score'])
  > ```

- 通过 RDD 转换 2 **StructType**

  > ```python
  > schema = StructType() \
  >         .add("id", IntegerType(), nullable=False) \
  >         .add("name", StringType(), nullable=False) \
  >         .add("score", IntegerType(), nullable=False)
  >     
  > spark_session.createDataFrame(rdd,schema=schema)
  > ```

- 通过 RDD 转换3 **toDF**

  > ```python
  > data_frame2 =  rdd.toDF(["id", "subject", "score"]) # 可以指定schema
  > ```





## SparkSql test

```sql
DROP TABLE IF EXISTS data_source;
CREATE TABLE data_source
USING tablestore
OPTIONS(
endpoint="https://c00qovcof0zs.cn-shenzhen.ots-internal.aliyuncs.com",
access.key.id="LTAI5tBvXCnXJgD1JRMhAaFs",
access.key.secret="yaIOiAIR3Mj8g48uTga5t77QQ5wjAT",
instance.name="c00qovcof0zs",
table.name="daily_flow",
catalog='{"columns": {
        "date": {"type":"long"},
        "asin": {"type":"string"},
        "keyword": {"type":"string"},
        "ad_flow": {"type":"long"},
        "nature_flow": {"type":"long"},
        "ac_flow": {"type":"long"},
        "er_flow": {"type":"long"},
        "na_flow": {"type":"long"},
        "sb_flow": {"type":"long"},
        "sp_flow": {"type":"long"},
        "sd_flow": {"type":"long"}
    }}'
);    


DROP TABLE IF EXISTS data_source;
CREATE TABLE data_source
USING tablestore
OPTIONS(
endpoint="",
access.key.id="",
access.key.secret="",
instance.name="",
table.name="daily_flow",
catalog='{"columns": {
        "date": {"type":"long"},
        "asin": {"type":"string"},
        "keyword": {"type":"string"},
        "ad_flow": {"type":"long"}
        "nature_flow": {"type":"long"}
        "ac_flow": {"type":"long"}
        "er_flow": {"type":"long"}
        "na_flow": {"type":"long"}
        "sb_flow": {"type":"long"}
        "sp_flow": {"type":"long"}
        "sd_flow": {"type":"long"}
    }}'
);    


```

