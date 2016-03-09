#!/bin/sh 
#######1. 修改数据库连接配置
ACHILLES="/root/Achilles/src"
cd $ACHILLES/consoles/scripts
#ruby  update_Achilles_conf.rb
######2.初始化建表格
cd $ACHILLES/consoles/dbseeds
chmod +x dbseeds
#./dbseeds --env=test  //构建数据改成手动，防止初始化生产数据

#######4.安装静态资源文件
cd $ACHILLES/assets
cnpm install
gulp

########5. 修改执行权限
cd  /root/Achilles
chmod +x *.sh

#####查找最新的编译文件
cd $ACHILLES
path=`find . -name 'achilles_*' |sort -n -r |head -1`
filename=${path#./}
echo $filename
######杀掉进程，更新可执行权限
pgrep achilles | xargs kill -9
mv $filename  achilles
chmod +x achilles

