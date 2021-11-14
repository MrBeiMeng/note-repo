## 一、下载安装IDEA

官方下载地址: https://www.jetbrains.com/zh-cn/idea/download/

![image-20211113193928482](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113193928482.png)

![image-20211113193955627](https://gitee.com/xie-zhi-bin/image/raw/master/typro/image-20211113193955627.png)

注：安装IDEA之前需要我们机器上有JDK环境！！！



双击打开安装即可：

![image-20211113194017323](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113194017323.png)

![image-20211113194025706](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113194025706.png)

![image-20211113194034642](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113194034642.png)

![image-20211113194043738](https://gitee.com/xie-zhi-bin/image/raw/master/typro/image-20211113194043738.png)

![image-20211113194052041](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113194052041.png)

![image-20211113194105027](https://gitee.com/xie-zhi-bin/image/raw/master/typro/image-20211113194105027.png)

![image-20211113194113965](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113194113965.png)

双击打开`IntelliJ IDEA 2021.1.1 x64`:

![image-20211113194138168](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113194138168.png)

![image-20211113194145162](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113194145162.png)

![image-20211113194152938](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113194152938.png)

![image-20211113194200232](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113194200232.png)

#### 













## 二、创建Java工程

### 第一步：

Create New Project：创建一个新的工程。
Import Project：导入一个现有的工程。
Open：打开一个已有工程。比如：可以打开 Eclipse 项目。

**方式一：**

![image-20211113192324152](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113192324152.png)

**或者使用这种方式二：**

![image-20211113192438656](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113192438656.png)

### **第二步：选择指定目录下的 JDK 作为 Project SDK。**

其中，选择【New…】，选择 jdk 的安装路径所在位置：

![image-20211113192630407](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113192630407.png)

![image-20211113192707479](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113192707479.png)

上图是选择jdk的安装路径。

点击【OK】以后，选择【Next】

![image-20211113192802109](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113192802109.png)

这里不用勾选。选择【Next】，进入下一个页面：

![image-20211113192933066](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113192933066.png)

给创建的工程起一个名字和工程所在位置，点击【finish】

### **创建 Package **

**接着在 src 目录下创建一个 package：**
右键项目下的 src文件夹–>new -->Package

![image-20211113194037075](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113194037075.png)

点击后：输入包名

![image-20211113194125663](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113194125663.png)

### 创建一个class

右键创建

![image-20211113194332399](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113194332399.png)

![image-20211113194508198](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113194508198.png)

**不管是创建 Class，还是 Interface，还是 Annotation，都是选择 New --> Java Class，然后在 Kind 下拉框中选择创建的结构的类型**。接着在类 HelloWorld 里声明主方法，输出 Hello World!，完成测试

### 运行方法

**方法1：**

![image-20211113201155141](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113201155141.png)

**方法2：**

![image-20211113194818213](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113194818213.png)

**方法3：**

![image-20211113194944846](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113194944846.png)

![image-20211113195301544](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113195301544.png)

说明：在 IDEA 里要说的是，写完代码，不用点击保存。IDEA 会自动保存代码。

### **创建模块：**

相比较于多 Module 项目，小项目就无需搞得这么复杂。只有一个 Module 的结构 IntelliJ IDEA 也是支持的，并且 IntelliJ IDEA 创建项目的时候，默认就是单 Module 的结构的。

![image-20211113195437406](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113195437406.png)

点击【Next】

![image-20211113195459322](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113195459322.png)

![image-20211113195546395](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113195546395.png)

![image-20211113195727672](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113195727672.png)

### **如何删除模块**

![image-20211113200050212](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113200050212.png)

![image-20211113200105549](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113200105549.png)

**注意移除模块后，文件夹还保留，但是不能在其中创建java程序了，然后我们再次右键此工程，有delete选项，点击可以从磁盘彻底删除此模块。**

## 三、常用配置

### **查看项目配置**

![image-20211113200234821](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113200234821.png)

**查看项目结构后，显示这样**

![image-20211113200358150](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113200358150.png)

### 常用配置

**进入设置界面：**

![image-20211113200513124](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113200513124.png)

目录结构如下：

![image-20211113200837824](https://gitee.com/xie-zhiqing1/image/raw/master/typora/image-20211113200837824.png)

## IDEA快捷键：

### 1、编辑类快捷键：

```
        编辑类快捷键	                 介绍
        
        psvm + Tab					生成main方法
        sout + tab					生成输出语句
        Ctrl+X / Ctrl + Y			 删除一行
        Ctrl+D						复制一行
        Ctrl+/ 或 Ctrl+Shift+/		 注释代码
        Ctrl + Z					撤销
        Ctrl + Shift + Z			 取消撤销
        Ctrl + C					复制
        Ctrl + V					粘贴
        Ctrl + O					重写方法
        Ctrl + I					实现方法
        Ctr + shift + U				 大小写转化
        Ctrl + Shift + J			 整合两行为一行
        Ctrl + Shift + space		 自动补全代码
        Alt + 回车				   导入包,自动修正
        Alt + /	代码提示
        Alt + Insert				生成代码(如GET,SET方法,构造函数等)
        Ctrl + Alt + L				格式化代码
        Ctrl + Alt + I				自动缩进
        Ctrl + Alt + T				生成try catch
        Ctrl + Alt + O				优化导入的类和包
        fori					   生成for循环for (int i = 0; i< ; i++) { }
        iter					   生成增强for循环
        itar					   生成array for代码块
        itit					   生成iterator 迭代
        itli					   生成List的遍历
        itco					   生成Collection迭代
```

### 2、查找，替换类快捷键：

```
        查找、替换类快捷键			 介绍
        
        Ctrl + F					在当前文件中查找
        Ctrl + Shift + F			 在整个项目或者指定窗口中查找文本
        Ctrl + N					在项目中查找类
        Ctrl + Shift + N			 查找文件
        Ctrl + R					在当前文件进行文本替换
        Ctrl + Shift+R				 在指定窗口替换文本
        Ctrl + W					自动按语法选中代码
        Ctrl + Shift + W			 反向自动按语法选中代码
        Ctrl + G					定位行
        Ctrl＋Shift＋Backspace		跳转到上一次编辑的位置
        Ctrl + alt + ←/→	    	  前后跳转编辑过的地方
        Ctrl + Shift + Alt + N		  查找 变量 / 方法
        Alt + F7				     找到你的函数或者变量或者类的所有引用到的地方
        Alt + F3					高亮显示所有该选中文本，按 Enter 选中下一个，按 Esc 高亮消失
        F4							在当前类中查找变量的来源
        Ctrl + Shift + F7			 高亮显示所有该选中文本，按 Esc 高亮消失
        双击Shift					   查找任何内容

```

### 3、编译，运行类快捷键：

```
        编译、运行类快捷键				 介绍

        Ctrl + F9						编译项目
        Ctrl + Shift + F9				 编译当前文件
        Shift + F10						正常启动
        Alt + Shift + F10				 弹出 Run 的可选择菜单
        Shift + F9						debug模式启动
        Alt + Shift + F9				 选择 Debug
```

### 4、Debug快捷键：

```
        Debug快捷键					介绍

        F7							  在Debug模式下，步入，如果当前行断点是一个方法，则进入当前方法体内，如果该方法体还有方法，则不会进入该内嵌的方法中
        Shift + F7					   智能步入
        Alt + Shift + F7	             强制步入
        F8	在 Debug 				  模式下，步过，如果当前行断点是一个方法，则不进入当前方法体内
        Shift + F8					   步出
        Alt + Shift + F8	             强制步过
        alt + F8					   在 Debug 模式下，选中查看值
        Ctrl + Shift + F8				查看断点
        F9	在 Debug 				  模式下，恢复程序运行。如果该断点下面代码还有断点则停在下一个断点上
        Alt +F9						   运行至光标的位置
        Ctrl + Alt+ F9				   强制运行至光标处
        Alt + F10					   定位到断点
```

### 5、重构快捷键：

```
        重构快捷键						介绍

        Shift + F6						重命名
        Ctrl + Alt + C					抽取常量
        Ctrl + Alt + F					抽取字段
        Ctrl + Alt + M					抽取方法
        Ctrl + Alt + P					抽取参数
        Ctrl + Alt + V					抽取变量
```

### 6、其他类快捷键：

```
        一个普通标题						一个普通标题

        Ctrl + C						复制文件名
        Ctrl + Shift + C				 复制文件的完整路径
        Ctrl + E						显示最近打开的文件
        Ctrl + Shift + E				 显示最近修改的文件列表的弹出层
        Ctrl + P						方法参数提示
        Ctrl + Q						可以看到当前方法的声明
        Ctrl + Alt + Space				 类名或接口名提示
        Ctrl + F12						显示当前文件的结构
        Ctrl + H						显示当前类的结构图
        Ctrl + Q						显示注释文档信息
        连按两次Shift					 弹出 Search Everywhere 弹出层，查找任任内容
        Ctrl + [						移动光标到当前所在代码的花括号开始位置
        Ctrl + ]						移动光标到当前所在代码的花括号结束位置
        Ctrl + K						版本控制提交项目，需要此项目有加入到版本控制才能够使用
        Ctrl + T						版本控制更新项目，需要此项目有加入到版本控制才能够使用
        Ctrl + Tab						切换编辑窗口。如果在切换的过程又按Delete键，则是关闭对应选中的窗口
```

## 四、JetBrains校园邮箱申请，IDEA激活

#### 1.到JetBrains申请

#### 2.填写相应信息

#### 3.邮箱会收到验证信息，邮件标题为“JetBrains Educational Pack Confirmation”

#### 4.注册JetBrains邮箱

#### 5.使 JetBrains Account 帐号登录

#### 6.到IDEA激活

学生和教师可免费申请JetBrains开发工具使用。 取得授权后只需要使用相同的 JetBrains 帐号就可以激活其他JetBrains产品。

##### 1.到JetBrains申请

申请官网页面为https://www.jetbrains.com/student/点击申请按钮开始申请。

![image-20211113194617372](https://gitee.com/fhfhhfffds/ssj/raw/master/ssj/20211113194624.png)

##### 2、

![image-20211113194905274](https://gitee.com/fhfhhfffds/ssj/raw/master/ssj/20211113194905.png)

3、



![image-20211113195222684](https://gitee.com/xie-zhiqing1/image/raw/master/typora/20211113195222.png)

4、

![image-20211113194939848](https://gitee.com/fhfhhfffds/ssj/raw/master/ssj/20211113194939.png)