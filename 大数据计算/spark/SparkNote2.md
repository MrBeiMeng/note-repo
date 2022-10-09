# Spark-Scala 开发

> 环境搭建
>
> `spark3X` + `scala2.12` + `jdk1.8` + `intellj idea` + `maven`
>
> 为**project** / **model** 添加项目支持，**idea** 要安装 **scala** 插件
>
> **开发定式**
>
> ```scala
> // 建立和Spark 框架的连接
> val sparkConf: SparkConf = new SparkConf().setMaster("local[*]").setAppName("WordCount")
> val sc = new SparkContext(sparkConf)
> // 执行业务作业
> wordCount(sc)
> // 关闭连接
> sc.stop()
> ```
>
> **补全代码配置**
>
> ![image-20220924170150793](https://ccurj.oss-cn-beijing.aliyuncs.com/mac-typoraimage-20220924170150793.png)





## Hello world

**[gitee[scala_spark_learn]](https://gitee.com/beimengclub/scala_spark_learn/settings#index)**



## Spark Core -> RDD



### 创建

> 在 **Spark** 中创建 **RDD** 的方式可以分为四种
>
> - 从集合创建
>
>   ```scala
>   // 1
>   val seq = Seq[Int](elems = 1, 2, 3, 4)
>   val rdd: RDD[Int] = sc.parallelize(seq)
>   // 2 
>   val rdd: RDD[Int] = sc.makeRDD(
>     List(1, 2, 34)
>   )
>   ```
>
> - 从外部文件中创建 [hadoop/file/...]
>
>   ```scala
>   val lines: RDD[String] = sc.textFile("datas")
>   ```
>
> - 从其他 RDD 创建
>
>   略
>
> - 自己创建



### 算子

> 就是算子



#### 转换型

> 功能的补充和封装 map,flatmap



##### 1 map

> 函数签名;
>
> `T => U`
>
> 说明： 将处理的数据 **逐条** 进行映射转换，这里的转换可以是类型的转换，也可以是值的转换。

```scala
val rdd2 = rdd.map(_ * 2) // 等效于 rdd.map(item => item*2)
```



##### 2 mapPartitions

> 函数签名：
>
> `iter[T] => iter[U]`
>
> 说明:可以以分区为单位进行数据转换操作,但是会将整个分区的数据加载到内存进行引用,如果处理完的数据时不会被释放的

```scala
val rdd: RDD[Int] = sc.makeRDD(List(1, 2, 3, 4),2)
val rdd2 = rdd.mapPartitions(iter => {
  iter.map(_*2)
})
```



##### 3 flatMap

> 函数签名：
>
> `T =>U`
>
> 说明：操作和 **map** 相同 ，但是将 **map** 后的数据打散成为一维数据

```scala
val rdd: RDD[List[Int]] = sc.makeRDD(List(
  List(1, 2),
  List(2, 3)
))
val rdd2: RDD[Int] = rdd.flatMap(list => list)
```



##### 4 glom 格拉!姆\

> 说明：将同一个分区的数据直接转换为相同类型的内存数组进行处理，分区不变



##### 5 groupBy

> 函数签名；
>
> `T => K`
>
> 说明：根据执行的规则进行分组，分区默认不变，但是数据会被 打乱重新组合， 我们将这样的操作称之为 shuffle。极限情况下，数据可能被分在同一个分区中

```scala
val rdd: RDD[Int] = sc.makeRDD(List(
  1,2,3,4,5,6,7,8,9,10
))
val rdd2: RDD[(Int, Iterable[Int])] = rdd.groupBy(item => item % 2)
```



##### 6 filter

> 函数签名:
>
> `T => Bool`
>
> 说明：根据过滤的逻辑，删除不需要的数据[获得需要的数据]



##### 7 sample

> 说明：根据指定的规则从数据集中 **抽取** 数据



##### 8 distinct

> 说明：去除重复的项目



##### 9 coalesce

> 根据数据量 **缩减分区**， 用于大数据集过滤后，提高小数据集的执行效率



##### 10 repartition

> 扩大分区



##### 11 sortBy

> 说明：根据规则进行排序



#### 作业型

> 触发作业的执行和任务的调动 collect 



##### 1 reduce



##### 2 collect



##### 3 foreach





## Spark Sql

> **SparkSql** 是 **Spark** 用于结构化数据( **structured data** )处理的 **Spark** 模块
>
> - **DataFram**e
>
>   DataFrame => Rdd
>
>   `val rdd = df.rdd`
>
>   Rdd => DataFrame
>
>   `rdd.toDF("name","age","sex")`
>
> - **DataSet**
>
>   一种强类型数据结构，可以自定义类。而 DataFrame 只支持基本的类型结构
>
>   DataFrame => DataSet
>
>   `val ds = df.as[类型]`
>
>   DataSet => DataFrame
>
>   `ds.toDF`



### 编程 ｜ 变化

> 使用 **SparkSessioin** 作为 Spark 最新的SQL 查询起始点。
>
> SparkSession 有三种创建方式
>
> 1. 从Spark 数据源创建
> 2. 从存在的 RDD 进行转换
> 3. 可以从 HiveTable 进行查询返回



### DSL 语法

> **DataFrame** 提供一个特定领域语言( **domain-specific language**, **DSL** ) 去管理结构化的数据。可以在 **Scala**，**Java**，**Python** 和 R 中使用 **DSL**， 使用 **DSL** 语法风格不必去创建临时视图了





### UDF

> 自定义函数，类似 **avg** 创建之后便可以在sql 语句中直接使用
>
> 170集



### 加载

> `spark.read.load()`



## 存储

> `df.write.format("json").save(path)`