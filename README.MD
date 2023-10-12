# GoFrame 开发的v2ray面板

开发文档:  
- https://goframe.org/pages/viewpage.action?pageId=1114399  


节点地址   
``` json 
$uuid$ 订阅时会替换成用户的uuid   
```

数据示例：
``` json  
tcp:
协议:v2ray/vmess
{
  "port": "12345",
  "scy": "auto",
  "net": "tcp",
  "type": "none",
  "host": "",
  "path": "",
  "tls": "",
  "sni": "",
  "alpn": ""
}

tcp+tls:
协议:v2ray/vmess
{
  "port": "12345",
  "scy": "auto",
  "net": "tcp",
  "type": "none",
  "host": "",
  "path": "",
  "tls": "tls",
  "sni": "xxxx.com",
  "alpn": ""
}

tcp+http:
协议:v2ray/vmess
{
  "port": "12345",
  "scy": "auto",
  "net": "tcp",
  "type": "http",
  "host": "iqiyi.com",
  "path": "/test",
  "tls": "",
  "sni": "",
  "alpn": "",
  "header": {
    "type": "http",
    "request": {},
    "response": {}
  }
}

ws:
协议:v2ray/vmess
{
  "port": "12345",
  "scy": "auto",
  "net": "ws",
  "type": "none",
  "host": "iqiyi.com",
  "path": "/test",
  "tls": "",
  "sni": "",
  "alpn": ""
}

ws+tls:
协议:v2ray/vmess
{
  "port": "12345",
  "scy": "auto",
  "net": "ws",
  "type": "none",
  "host": "xxxx.com",
  "path": "/test",
  "tls": "tls",
  "sni": "xxxx.com",
  "alpn": ""
}


协议:v2ray/vless
{
  "port": "12345",
  "encryption": "", //加密方式
  "security": "", //传输层安全
  "flow": "", //流控
  "sni": "",
  "alpn": "",
  "fp": "",
  "pbk": "",
  "sid": "",
  "spx": "",
  "type": "", //传输协议 (订阅用)
  "net": "", //传输协议 (后端用)
  "serviceName": "", //grpc serviceName
  "mode": "", //grpc 模式
  "headerType": "",
  "quicSecurity": "", //quid 加密方式
  "key": "", //quid 加密密钥
  "host": "",
  "path": "",
  "seed": "" //kcp seed
}

vless+tcp+reality:
协议:v2ray/vless
{
	"port": "45555",
	"encryption": "none",
	"security": "reality",
	"flow": "xtls-rprx-vision",
	"sni": "www.smzdm.com",
	"alpn": "",
	"fp": "chrome",
	"pbk": "3embN-YWuuKUd4Lg-7YsbNJwLisZZx2ZF2omf5YvjTs",
	"sid": "0123456789abcdef",
	"type": "tcp",
	"net": "tcp",
	"headerType": "none",
	"host": ""
}

shadowsocks2022
协议:shadowsocks/ss2022
{
	"port": "9996",
	"cypher_method": "aes-128-gcm",
	"obfs": "none",
	"server_key":"",
	"obfs_settings": {
		"path": "",
		"host": ""
	}
}

协议:trojan/trojan
trojan
{
	"host":"",
	"port":"",
  "sni":"",
  "security":"",
  "alpn":"",
  "fp":"",
  "type":"tcp",
  "headerType":""
}

``` 

# 支付数据 当前仅支持易支付  
``` json  
{
	"url": "",
	"pid": "",
	"key": ""
}
```


# 当前阶段任务  
测试和BUG  
订阅开发  
xrayr后端对接  

# 计划任务  
windows 客户端  
安卓 客户端  



# 安装教程  
xrayr后端对接以开发，但尚未编译并且还在测试用：https://github.com/pingProMax/XrayR , PanelType填写"GoV2Panel"  

1、创建数据库导入数据  

2、修改config.yaml文件  
manifest/config/config.yaml  
修改数据库连接信息  
修改jwtkey 随机填写  
修改admini_path 后台管理地址  

3、启动项目  
``` code
cd 项目目录
nohup ./gov2panel &
检查日志启动无误
```

4、配置nginx  
反向代理，仅供参考  

``` nginx 
location / {
  proxy_pass              http://127.0.0.1:8080;
  proxy_redirect          off;
  proxy_set_header        X-Real-IP       $remote_addr;
  proxy_set_header        X-Forwarded-For $proxy_add_x_forwarded_for;
}
```

5、登录后台修改 账号/密码  
adminx/123456789  

# 后台展示  
![后台界面](https://github.com/pingProMax/gov2panel/blob/master/admin.png?raw=true)