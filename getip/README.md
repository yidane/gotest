# 自动设置环境变量LOCALIP

## 脚本意义

开机自动获取网卡IP地址，并将此IP地址设置为环境变量，以便供其他脚本调用

## 测试数据准备

在当前目录下执行以下脚本

```` bash
cp /etc/skel/.bashrc .bashrc
````