#!/bin/bash

#修改修改本机网络名称
# shellcheck disable=SC2162
read -p "Please Input Your ServerName is:" NAME
echo "$NAME" > /etc/hostname

#Server查看ip 网关
# shellcheck disable=SC2006
IP=`ip a sh | grep ens33 | awk 'NR==2 {print $2}'`
# shellcheck disable=SC2006
GATEWAY=`ip rou sh | awk 'NR==1 {print $3}'`
echo "Server IP/Netmask is:" $IP
echo "Server GateWay is:" $GATEWAY

#修改YUM镜像服务
yum -y install wget
cp -a /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.bak
wget -O /etc/yum.repos.d/CentOS-Base.repo https://repo.huaweicloud.com/repository/conf/CentOS-7-reg.repo
sed -i "s/#baseurl/baseurl/g" /etc/yum.repos.d/CentOS-Base.repo
sed -i "s/mirrorlist=http/#mirrorlist=http/g" /etc/yum.repos.d/CentOS-Base.repo
sed -i "s@http://mirror.centos.org@https://repo.huaweicloud.com@g" /etc/yum.repos.d/CentOS-Base.repo
#刷新页面
yum clean all && yum makecache

#安装软件工具
yum -y install vim net-tools curl gcc make lrzsz chrony telnet nmap nmap-ncat

#配置chrony时间服务
systemctl enable chronyd && systemctl start chronyd
#基础系统内核优化
more >> /etc/sysctl.conf << EOF
net.ipv4.tcp_syncookies = 1
net.ipv4.tcp_syn_retries = 1
net.ipv4.tcp_tw_recycle= 1
net.ipv4.tcp_tw_reuse = 1
net.ipv4.tcp_fin_timeout = 1
net.ipv4.tcp_keepalive_time = 1200
net.ipv4.ip_local_port_range = 1024 65535
net.ipv4.tcp_max_syn_backlog = 16384
net.ipv4.tcp_max_tw_buckets = 36000
net.ipv4.routegc_timeout = 100
net.ipv4.tcp_syn_retries = 1
net.ipv4.tcp_synack_retries = 1
net.core.somaxconn = 16384
net.core.netdev_max_backlog = 16384
net.ipv4.tcp_max_orphans = 16384
EOF
/sbin/sysctl -p
#关闭SELINUX
sed -i "s/SELINUX=enforcing/SELINUX=disabled/g" /etc/selinux/config
#关闭IPTABLES
systemctl disable firewalld
#重启服务器
#shutdown -r 0
# 安装nohup
yum install coreutils


#安装jdk1.8
        #查找可用安装，这个是带描述的 yum list available | grep jdk
yum search java | grep -i --color jdk
#建议先卸载无关jdk
#开始安装jdk
yum -y install java-1.8.0-openjdk-devel.x86_64

#开始安装tomcat
yum -y install tomcat
    #安装目录在 "/usr/share/tomcat"
    #systemctl start tomcat.service启动tomcat
    #systemctl start tomcat
    #systemctl status tomcat
    #systemctl enable tomcat开机自启
    #systemctl restart tomcat重启
#配置默认访问页面，就是tomcat默认页面，否则会404因为webapp下是空的
yum -y install tomcat-webapps tomcat-admin-webapps
systemctl start tomcat
systemctl status tomcat


#安装mysql5.7
wget https://dev.mysql.com/get/mysql57-community-release-el7-9.noarch.rpm
rpm -ivh mysql57-community-release-el7-9.noarch.rpm
cd /etc/yum.repos.d/
yum -y install mysql-server
systemctl start mysqld
echo "mysql初始密码"
grep 'temporary password' /var/log/mysqld.log
