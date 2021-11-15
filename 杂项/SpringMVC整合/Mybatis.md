# Mybatis层

**环境要求**

`idea mysql5.7 tomcat9 Maven3.6`

------

## 数据库准备

### sql 代码

```mysql
create database ssmbuild;

use ssmbuild;
create table if not exists book(
    book_id int(10) not null auto_increment comment '书籍id',
    book_name varchar(100) not null comment '书籍名称',
    book_counts int(11) not null comment '数量',
    detail varchar(200) not null comment '描述',
    primary key (book_id)
) engine = INNODB default charset = utf8;

insert into book (book_id, book_name, bookCounts, detail)
values
       (1,'Java',1,'从入门到精通'),
       (2,'MySQL',10,'数据库基础入门'),
       (3,'Linux',5,'Linux基本运维');
```

运行sql代码

![image-20211115140151702](https://i.loli.net/2021/11/15/BUgLjc1E9dFyNAs.png)

### maven 依赖

```xml
 <dependencies>
     <dependency>
         <groupId>junit</groupId>
         <artifactId>junit</artifactId>
         <version>4.13.2</version>
         <scope>test</scope>
     </dependency>
     <dependency>
         <groupId>mysql</groupId>
         <artifactId>mysql-connector-java</artifactId>
         <version>5.1.47</version>
     </dependency>


     <dependency>
         <groupId>javax.servlet</groupId>
         <artifactId>servlet-api</artifactId>
         <version>2.5</version>
     </dependency>
     <dependency>
         <groupId>javax.servlet</groupId>
         <artifactId>jstl</artifactId>
         <version>1.2</version>
     </dependency>

     <!--        mybatis-->
     <dependency>
         <groupId>org.mybatis</groupId>
         <artifactId>mybatis</artifactId>
         <version>3.5.7</version>
     </dependency>
     <dependency>
         <groupId>org.mybatis</groupId>
         <artifactId>mybatis-spring</artifactId>
         <version>2.0.6</version>
     </dependency>

     <!--        spring-->
     <dependency>
         <groupId>org.springframework</groupId>
         <artifactId>spring-webmvc</artifactId>
         <version>5.3.10</version>
     </dependency>
     <dependency>
         <groupId>org.springframework</groupId>
         <artifactId>spring-jdbc</artifactId>
         <version>5.3.10</version>
     </dependency>
</dependencies>
```

### build

```xml
<build>
    <resources>
        <resource>
            <directory>src/main/java</directory>
            <includes>
                <include>**/*.properties</include>
                <include>**/*.xml</include>
            </includes>
            <filtering>false</filtering>
        </resource>
        <resource>
            <directory>src/main/resources</directory>
            <includes>
                <include>**/*.properties</include>
                <include>**/*.xml</include>
            </includes>
            <filtering>false</filtering>
        </resource>
    </resources>
</build>
```

### entity

```java
@Data
public class Book {
    private int bookId;
    private String bookName;
    private int bookCounts;
    private String detail;
}
```

### applicationContext.xml

```xml
```

### BookMapper.xml

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper
        PUBLIC "-//mybatis.org//DTD Config 3.0//EN"
        "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper namespace="club.beimeng.mapper.BookMapper">
    
    <insert id="addBook" parameterType="Book">
        insert into ssmbuild.book (book_name, book_counts, detail)
        values (#{bookName},#{bookCounts},#{detail});
    </insert>

    <delete id="deleteBookById" >
        delete
        from ssmbuild.book
                 where book_id=#{bookId};
    </delete>

    <update id="updateBookById" parameterType="Book">
        update ssmbuild.book
        set book_name = #{bookName},
            book_counts = #{bookCounts},
            detail = #{detail}
            where book_id = #{bookId};
    </update>

    <select id="getBookById" resultType="Book">
        select *
        from ssmbuild.book
        where book_id = #{bookId};
    </select>

    <select id="getAllBooks" resultType="Book">
        select *
        from ssmbuild.book;
    </select>
</mapper>
```

