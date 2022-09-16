## Spark 笔记

**Spark 是一种基于内存的快速、通用、可扩展的大数据分析计算引擎**

2009 - 至今





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



### **架构角色**

Spark 中由4类角色组成整个Spark 的运行时环境

- Master
- Worker
- Driver
- Executor



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