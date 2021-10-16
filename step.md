# 步骤
## 编译并推送镜像
```shell
make build
```
## 启动
```shell
make run
```
## 查看容器信息
```shell
docker ps
```
显示结果如下: 
```text
CONTAINER ID   IMAGE                     COMMAND                  CREATED         STATUS         PORTS                                       NAMES
9da42d4b92ed   idefav/httpserver:0.0.1   "./idefav-httpserver…"   8 minutes ago   Up 8 minutes   0.0.0.0:8080->8080/tcp, :::8080->8080/tcp   amazing_kapitsa
```
执行下面命令查看 docker 容器详情
```shell
docker inspect 9da42d4b92ed|grep Pid
```
结果如下:
```text
"Pid": 37437,
"PidMode": "",
"PidsLimit": null,
```

## 进入容器网络
```shell
nsenter -t 37437 -n
```
执行网络命令
```shell
ifconfig
```
显示结果
```text
eth0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 172.17.0.2  netmask 255.255.0.0  broadcast 172.17.255.255
        ether 02:42:ac:11:00:02  txqueuelen 0  (Ethernet)
        RX packets 50  bytes 5128 (5.1 KB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 10  bytes 947 (947.0 B)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

lo: flags=73<UP,LOOPBACK,RUNNING>  mtu 65536
        inet 127.0.0.1  netmask 255.0.0.0
        loop  txqueuelen 1000  (Local Loopback)
        RX packets 0  bytes 0 (0.0 B)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 0  bytes 0 (0.0 B)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
```
退出容器网络
```shell
exit
```
再执行网络命令
```shell
ifconfig
```
结果如下
```text
docker0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 172.17.0.1  netmask 255.255.0.0  broadcast 172.17.255.255
        inet6 fe80::42:29ff:fe56:a0fa  prefixlen 64  scopeid 0x20<link>
        ether 02:42:29:56:a0:fa  txqueuelen 0  (Ethernet)
        RX packets 10  bytes 807 (807.0 B)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 27  bytes 2592 (2.5 KB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

ens33: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 192.168.0.111  netmask 255.255.255.0  broadcast 192.168.0.255
        inet6 fe80::ce84:fdc2:a9e:440f  prefixlen 64  scopeid 0x20<link>
        ether 00:0c:29:90:cc:d7  txqueuelen 1000  (Ethernet)
        RX packets 249051  bytes 232109693 (232.1 MB)
        RX errors 0  dropped 10  overruns 0  frame 0
        TX packets 72018  bytes 8270815 (8.2 MB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

lo: flags=73<UP,LOOPBACK,RUNNING>  mtu 65536
        inet 127.0.0.1  netmask 255.0.0.0
        inet6 ::1  prefixlen 128  scopeid 0x10<host>
        loop  txqueuelen 1000  (Local Loopback)
        RX packets 6672  bytes 2656549 (2.6 MB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 6672  bytes 2656549 (2.6 MB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

veth4f16903: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet6 fe80::8a9:5fff:fef6:420  prefixlen 64  scopeid 0x20<link>
        ether 0a:a9:5f:f6:04:20  txqueuelen 0  (Ethernet)
        RX packets 10  bytes 947 (947.0 B)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 50  bytes 5128 (5.1 KB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
```