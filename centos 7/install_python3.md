# centos 7 安装 python3

先看这里：

https://blog.csdn.net/qq_39715000/article/details/125009276

> 首先要确保机器安装好了 **openssl**  

```shell
yum install openssl 
```

> 下载python安装包

```shell
wget https://www.python.org/ftp/python/3.10.7/Python-3.10.7.tgz
```

> 解压并进入对应的文件夹

```shell
tar -zxvf Python-3.10.7.tgz
cd Python-3.10.7
```

> 进行编译

```shell
./configure --with-ssl --prefix=/usr/local/python3
make && make install
```

> 创建软连接

```shell
ln -s /usr/local/python3/bin/python3 /usr/bin/python3
ln -s /usr/local/python3/bin/pip3 /usr/bin/python3
```

> 修改pip 国内源

```shell
 config set global.index-url https://pypi.mirrors.ustc.edu.cn/simple/    
pip3 config set install.trusted-host mirrors.ustc.edu.cn
pip3 config list
# 显示结果：
# pip config set global.index-url http://mirrors.aliyun.com/pypi/simple/
# pip config set install.trusted-host mirrors.aliyun.com
```

- 