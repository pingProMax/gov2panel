# GoFrame 开发的v2ray面板

开发文档:  
- https://goframe.org/pages/viewpage.action?pageId=1114399  



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

shadowsocks2022
协议:shadowsocks/ss2022
{
  "cypher_method": "12345",
  "obfs": "",
  "server_key":"",
  "obfs_settings": {
	"path":"",
	"host":""
  }
}

协议:trojan/trojan
trojan
{
	"host":"",
	"port":"",
	"server_name":"sni.com",
}

``` 



# 当前阶段任务  
Web 开发  
xrayr 后端对接  

# 计划任务  
windows 客户端  
安卓 客户端  


# 安装教程  
修改jwtkey
manifest/config/config.yaml  jwtkey 随机填写


