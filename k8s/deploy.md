# 生产化部署操作步骤

## 编译并上传 docker 镜像

```shell
make IMAGE=httpserver:3.0.0 push
```

## 部署镜像

###  http 访问普通方式部署
#### 使用 helm 安装 httpserver 到 k8s 集群

```shell
helm install idefav-httpserver ./idefav-httpserver -f ./idefav-httpserver/values.yaml
```

#### 查看 k8s

```shell
kubectl get po
NAME                                 READY   STATUS    RESTARTS       AGE
deathstar-5df8d659b6-jppjj           1/1     Running   1 (103m ago)   5h58m
idefav-httpserver-659cfbdc69-5zcql   1/1     Running   1 (103m ago)   3h11m
idefav-httpserver-659cfbdc69-6rg9s   1/1     Running   1 (103m ago)   3h15m
idefav-httpserver-659cfbdc69-82pwr   1/1     Running   1 (103m ago)   3h11m
nginx-564769bf66-s6dql               1/1     Running   1 (103m ago)   5h58m


kubectl get deploy
NAME                READY   UP-TO-DATE   AVAILABLE   AGE
deathstar           1/1     1            1           22d
idefav-httpserver   3/3     3            3           4h22m
nginx               1/1     1            1           22d

kubectl get cm
NAME                DATA   AGE
idefav-httpserver   1      4h22m
kube-root-ca.crt    1      22d

kubectl get ingress
NAME                CLASS    HOSTS               ADDRESS         PORTS   AGE
idefav-httpserver   <none>   demo.idefav.local   10.110.34.166   80      4h23m
```

#### 验证 域名访问

```shell
curl -H"host:demo.idefav.local" -i 192.168.0.112:30080/healthz
HTTP/1.1 200 OK
Date: Sun, 28 Nov 2021 12:05:27 GMT
Content-Type: application/json
Content-Length: 39
Connection: keep-alive

{"code":0,"message":"","data":"Health"}
```

### https 访问方式部署

#### 制作证书
```shell
openssl genrsa -out idefav.key 2048
openssl req -new -x509 -key idefav.key -out idefav.crt -subj /C=CN/ST=Shanghai/L=Shanghai/O=DevOps/CN=*.idefav.local
```
#### 上传证书到 k8s
```shell
kubectl create secret tls idefav-secret --cert=idefav.crt --key=idefav.key
```
#### 部署 tls 访问的 ingress 到 k8s
```shell
helm upgrade idefav-httpserver ./idefav-httpserver/ -f ./idefav-httpserver/valuesWithTls.yaml
```
#### 验证访问
```shell
curl -H"host:demo.idefav.local" -i -k https://192.168.0.112:30443/healthz
HTTP/1.1 200 OK
Date: Sun, 28 Nov 2021 12:11:29 GMT
Content-Type: application/json
Content-Length: 39
Connection: keep-alive
Strict-Transport-Security: max-age=15724800; includeSubDomains

{"code":0,"message":"","data":"Health"}
```

## 应用高可用参照项目Readme.md