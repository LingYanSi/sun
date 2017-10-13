mac:
	gox -osarch="darwin/amd64"

linux:
	gox -osarch="linux/amd64"

run:
	lywatch --cmd="go run *.go" --port="8965"

pro:
	nohup /root/sun/sun_linux_amd64 &

cp:
	scp sun_linux_amd64 root@108.61.160.163:/root/sun

redis: 
	yum install -y gcc-c++
	yum install -y tcl
	yum install -y wget

	# 下载 https://redis.io/ 下载最新稳定版本
	wget http://download.redis.io/releases/redis-4.0.2.tar.gz
	tar -xzvf redis-4.0.2.tar.gz 

	# 安装
	cd redis-4.0.2
	make
	make PREFIX=/usr/local/ install

	# 创建配置
	mkdir -p /etc/redis
	cp redis.conf /etc/redis

	# 开机自启动
	echo "/usr/local/redis/bin/redis-server /etc/redis/redis.conf &" >> /etc/rc.local

	# 仅修改： daemonize yes （no-->yes）
	# 启动守护进程
	# vi /etc/redis/redis.conf
	sed -i "s/daemonize\s\+no/daemonize yes/g" /etc/redis/redis.conf

	# 启动redis 并制定配置文件
	redis-server /etc/redis/redis.conf 

depend:
	go get github.com/fatih/color
	go get github.com/go-redis/redis
	gi get github.com/go-sql-driver/mysql
