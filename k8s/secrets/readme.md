```shell

openssl genrsa -out idefav.key 2048
openssl req -new -x509 -key idefav.key -out idefav.crt -subj /C=CN/ST=Shanghai/L=Shanghai/O=DevOps/CN=*.idefav.local

kubectl create secret tls idefav-secret --cert=idefav.crt --key=idefav.key
```