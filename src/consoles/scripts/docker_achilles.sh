#!/bin/sh
#######1. 修改数据库连接配置
ACHILLES="/root/docker/src_code/achilles/src"
cd $ACHILLES/consoles/scripts
#ruby  update_Achilles_conf.rb
######2.初始化建表格
cd $ACHILLES/consoles/dbseeds
chmod +x dbseeds
#./dbseeds --env=prod

########3. 修改执行权限
cd  /root/docker/src_code/achilles/
chmod +x *.sh

##4.安装资源文件
cd $ACHILLES/assets
cnpm install
gulp
#####5.增加可运行权限
cd $ACHILLES
mv achilles  src
chmod +x src


##6.杀掉旧容器
echo '>>> Get old container id'

CID=$(docker ps | grep "4b397e671d3a" | awk '{print $1}')

echo $CID
echo '>>>restart docker'
if [ "$CID" != "" ]
then
   docker stop $CID
fi

###7.运行 docker
docker run -it -d -p 8081:8080 -v /root/docker/src_code/achilles:/mnt 4b397e671d3a