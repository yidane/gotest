#!/usr/bin/env bash
# 环境变量设置方法：

# 1、/etc/profile:在登录时,操作系统定制用户环境时使用的第一个文件,此 文件为系统的每个用户设置环境信息,当用户第一次登录时,该文件被执行。
# 2、/etc/environment:在登录时操作系统使用的第二个文件,系统在 读取你自己的profile前,设置环境文件的环境变量。
# 3、~/.bash_profile:在登录时用到的第三个文件是.profile文 件,每个用户都可使用该文件输入专用于自己使用的shell信息,当用户登录时,
# 该文件仅仅执行一次!默认情况下,他设置一些环境变游戏量,执 行用户的.bashrc文件。/etc/bashrc:为每一个运行bash shell的用户执行此文件.当bash shell被打开时,该文件被读取.
# 4、~/.bashrc:该文件包含专用于你的bash shell的bash信 息,当登录时以及每次打开新的shell时,该该文件被读取。

# 几个环境变量的优先级
# 1>2>3

# 设置永久环境变量
# 1.环境变量配置中，要先删除.bash_profile中的三行关于.bashrc的 定义，然后把环境变量配置在.bashrc中
# 2.选择要使用的java环境：update-alternatives –config java
# 3.要使得刚修改的环境变量生效：source .bashrc
# 4.查看环境变量：env

# 可以放到/etc/bash/bashrc，这样就是系统级的

# 获取网卡wlp2s0所获取的IP地址
envKey='LOCALIP';
bashPath='.bashrc'
netCard='wlp2s0'

ip=`ifconfig $netCard | grep "inet" | awk '{ print $2}' | awk 'NR==1{print}'`;

echo $ip

c=1;
while [ -z $ip ] && [ $c -le 10 ];
do
    echo 'ERROR: IP of '$netCard' is empty '$c;
    let c+=1;
    sleep 1;
    ip=`ifconfig $netCard | grep "inet" | awk '{ print $2}' | awk 'NR==1{print}'`;
done

if [ -z $ip ]
then
    echo 'ERROR: Get IP of '$netCard' faild';
    ip='NOTHING'
fi

num=$(grep -nr $envKey  $bashPath | awk -F ':' '{print $1}');
if [ ${#num} -eq 0 ] # 判断长度是否大于0
then
    echo 'WARN: no '$envKey' in file '$bashPath
    sudo sed -i '$a export '$envKey'='$ip $bashPath;
    echo 'INFO: add '$envKey' to the end line of '$bashPath
    env | grep $envKey;
    exit 0;
fi

sudo sed -i $num'c export '$envKey'='$ip $bashPath;
echo 'INFO: update value of '$envKey' to '$ip' in file '$bashPath
source $bashPath;
env | grep $envKey;
exit 0;