
### Tinachain-tx
Tinachain-tx.js 目的是为了让用户非常方便的在客户端（钱包端）发起交易，并在本地进行签名后将交易发送到Tianchain上。
由于Tinachain不建议将区块链的默认端口8545对外开发，因此需要安装Nginx进行交易数据中转。后期可以在Nginx上进行特定业务的开发。

### 安装Nginx
* gcc 安装
    yum install gcc-c++

* PCRE pcre-devel 安装
    yum install -y pcre pcre-devel

* zlib 安装
    yum install -y zlib zlib-devel

* OpenSSL 安装
    yum install -y openssl openssl-devel

* 下载Nginx
    wget -c https://nginx.org/download/nginx-1.10.1.tar.gz

* 解压
    tar -zxvf nginx-1.10.1.tar.gz
    cd nginx-1.10.1

* 配置
    ./configure

* 安装
    make && make install


### 配置Nginx
* cd /usr/local/nginx/conf
* vi nginx.conf

    server {
        listen      9001 ;
        location / {
                proxy_pass   http://127.0.0.1:8545;
                add_header 'Access-Control-Allow-Credentials' 'true';
                add_header Access-Control-Allow-Origin *;
                add_header Access-Control-Allow-Headers *;
        }
    }