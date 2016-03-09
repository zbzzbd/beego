# 前端环境搭建
暂时未加入一些其他有用工具环境。还有工作量、选型、团队学习、接受方式等问题。后面视情况考虑。

## 1. node安装
https://nodejs.org/en/
安装完成后，能执行下列命令，则安装正确：
node -v
npm -v

淘宝源： http://npm.taobao.org/    用cnpm代替npm就可以了

## 2. 安装gulp
npm install -g gulp

安装完成后，能执行gulp -v命令，则安装正确.

## 3到前端static目录下面执行：
npm install




## 4. 安装浏览器livereload插件
chrome可以在应用商店搜索livereload
其他浏览器请google

注意：此为开发环境必须的。生产环境跳过。

## 5. 在前端目录下运行

```
gulp
```

上面命令自动构建前端所需静态文件。

开发环境下，可继续执行

```
gulp watch
```

监控文件改动，自动编译。配合livereload，可以自动刷新浏览器页面，实时查看页面修改效果。
