**��װNginx

gcc ��װ
yum install gcc-c++

PCRE pcre-devel ��װ
yum install -y pcre pcre-devel

zlib ��װ
yum install -y zlib zlib-devel

OpenSSL ��װ
yum install -y openssl openssl-devel

����Nginx
wget -c https://nginx.org/download/nginx-1.10.1.tar.gz

��ѹ
tar -zxvf nginx-1.10.1.tar.gz
cd nginx-1.10.1

����
./configure

��װ
make && make install

**����Nginx
cd /usr/local/nginx/conf
vi nginx.conf

server {
        listen      9001 ;
        location / {
                proxy_pass   http://127.0.0.1:8545;
                add_header 'Access-Control-Allow-Credentials' 'true';
                add_header Access-Control-Allow-Origin *;
                add_header Access-Control-Allow-Headers *;
        }
    }
	
	