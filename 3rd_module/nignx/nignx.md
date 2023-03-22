## 1.简介

## 1.1nginx由来

2004年10月4号由俄罗斯人创作。

当时的web server满足不了需求，阿帕奇解决不了，单枪匹马开发nginx。

## 1.2nginx三大功能

- web服务器

  作为web服务器，nginx是轻量级的而且能够处理的并发量更大。

- 反向代理服务器

  何为反向代理服务器？

  ![负载均衡](.\imgs\负载均衡.png)

  如图所示，nginx作为反向代理服务器，完成的是负载均衡的工作。

  它可以将用户端的请求，透明的转送给应用服务器。这样所有的客户端只需要访问同一个nginx服务器就可以了。

  nginx内布会有一些负载均衡的算法和规则，来平均给身后的server分发链接，达到每个服务器负载量均衡。

- 邮件服务器

  nginx也可以用来充当一个IMAP/POP3/SMTP服务器。

## 1.3下载&安装

### 1.3.1官网安装

官网地址：

http://nginx.org

企业一般是用稳定版，stable.

或者直接进入下载页面:

http://nginx.org/download/

### 1.3.2docker安装

本文着重讲解docker安装nginx的方法。

- 拉取1.22.0版本的nginx，其他版本可以去DockerHub查询

  ```go
  docker pull nginx:1.22.0
  ```

- 查看镜像库，获取本地nginx镜像信息

  ```go
  docker images
  ```

  记录下IMAGE ID前4位，后边使用。

  ![image-20230322153924650](.\imgs\image-20230322153924650.png)

- 简单启动容器，测试nginx镜像是否可用

  这一步一般不要跳过。除了测试镜像是否可用，我们还需要复制容器内的相关文件，后续挂载到宿主机的管理目录。

  ```go
  docker run --name nginx -p 80:80 -d nginx:1.22.0
  -d 指定容器以守护进程方式在后台运行
  –name 指定容器名称，此处指定的是nginx
  -p 指定主机与容器内部的端口号映射关系，格式 -p [宿主机端口号]：[容器内部端口]，此处我使用了主机80端口，映射容器80端口
  ```

  命令执行后返回了容器ID,此时我们已经成功创建了nginx容器

  使用docker ps命令查看正在运行的nginx容器， 映射端口为80

  浏览器访问http://host:80 可以看到nginx页面，其中host为宿主机ip。例如我的host就是wsl的ip：http://172.21.228.106/

- 修改nginx配置文件

  修改配置文件有两种方式：

  *方式一：*每次都进入到nginx容器内部修改，适合改动少，简单使用的情况：

  1. 执行命令：docker exec -it ef /bin/bash 进入到nginx容器内部

     exec命令代表附着到运行着的容器内部

     -it 是 -i 与 -t两个参数合并写法， -i -t 标志着为我们指定的容器创建了TTY并捕捉了STDIN

     ef 是我们要进入的容器的id，可以只写前几位，只要唯一即可。

     /bin/bash 指定了执行命令的shell

  2. 进入到nginx容器内部后，可以 cd /etc/nginx, 看到相关的nginx配置文件都在该目录下。

     而nginx容器内的默认首页html文件目录为 /usr/share/nginx/html

     日志文件位于 /var/log/nginx

  3. 执行`exit`命令可以从容器内部退出

  *方式二：*将nginx容器内部配置文件挂在到主机

  这里说的内部配置文件，主要是默认页面、日志、配置文件。挂载到宿主机上的好处是：可以直接在主机对应目录修改即可，在需要频繁修改Nginx配置的场景中，不需要关注容器内的数据，直接删除重建容器就好。步骤如下:

  1. 主机上创建挂载文件夹:

     ```go
     mkdir /home/lsx01/docker-nginx
     mkdir /home/lsx01/docker-nginx/conf
     mkdir /home/lsx01/docker-nginx/logs
     mkdir /home/lsx01/docker-nginx/html
     ```

  2. 将容器中的相应文件copy到刚创建的管理目录中

     ```go
     docker cp dbc:/etc/nginx/nginx.conf /home/lsx01/docker-nginx/
     
     docker cp dbc:/etc/nginx/conf.d/ /home/lsx01/docker-nginx/conf/
     
     docker cp dbc:/usr/share/nginx/html/ /home/lsx01/docker-nginx/html/
     
     docker cp dbc:/var/log/nginx/ /home/lsx01/docker-nginx/logs/
     ```

     注意，上述指令中的`dbc`为容器前缀，只要唯一即可。

  一般推荐使用方法二来修改nginx配置文件。

- 停止并移除刚刚创建的简单容器

  ```go
  docker stop nginx
  docker rm nginx
  ```

- 再次启动容器并作目录挂载

  注意！这个启动命令可能是错的，因为我这边启动之后无法访问nginx首页。

  但是上面的简单启动是没问题的，修改配置的话，就用方法一。

  ```go
  docker run -p 80:80 \
  -v /home/lsx01/docker-nginx/nginx.conf:/etc/nginx/nginx.conf \
  -v /home/lsx01/docker-nginx/logs:/var/log/nginx \
  -v /home/lsx01/docker-nginx/html:/usr/share/nginx/html \
  -v /home/lsx01/docker-nginx/conf:/etc/nginx/conf.d \
  --name nginx \
  -e TZ=Asia/Shanghai \
  --restart=always \
  -d nginx:1.22.0
  
  参数说明：
  -p 映射端口，格式为“宿主机端口:容器端口”
  -m 200m 分配内存空间
  -v 挂载文件
  第一个-v 表示将你本地的nginx.conf覆盖你要起启动的容器的nginx.conf文件，
  第二个-v 表示将日志文件进行挂载，就是把nginx服务器的日志写到你docker宿主机的/home/docker-nginx/log/下面
  第三个-v 表示的和第一个-v意思一样的。
  
  -e TZ=Asia/Shanghai  设置时区
  
  --privileged=true 让容器中的root用户拥有真正的root权限
  
  --name  容器名字，以后可以使用这个名字启动或者停止容器
  
  --restart=always docker启动时自动启动容器
  -d 指定要启动的镜像名
  ```

- 浏览器访问，确保nginx容器已启动

  大功告成！

- 修改Nginx配置

  在宿主机上修改html目录下的文件是即时生效的。

  在宿主机上修改nginx.conf和conf目录下的配置文件后，需要重启容器重新加载配置。

  `注意：`修改配置文件时，文件中的路径要使用容器中的路径。

- 新增/删除映射端口

  最简单的方法就是停止、删除当前的容器，修改docker run命令中参数-p对应的端口映射值后再重新创建容器。

- 容器跨主机访问

  Nginx用于负载均衡时，需要访问宿主机以外的其它主机，最简单的做法是使用host模式创建容器，这时候容器将会共用使用宿主机的IP和端口**。**这种方式性能高，但无法自定义容器的网络配置和管理。

  ```go
  #改造之前的docker run，添加--net=host，去掉-p。
  docker run \
  -v /home/用户/docker-nginx/nginx.conf:/etc/nginx/nginx.conf \
  -v /home/用户/docker-nginx/logs:/var/log/nginx \
  -v /home/用户/docker-nginx/html:/usr/share/nginx/html \
  -v /home/用户/docker-nginx/conf:/etc/nginx/conf.d \
  -v /etc/localtime:/etc/localtime \
  --net=host \
  --name nginx \
  --restart=always \
  -itd nginx:1.22.0
  ```

  sed -i -e 's#Welcome to nginx#what a beautiful aurora!#g' -e 's#<head>#<head><meta charset="utf-8">#g' index.html