# GoFrame 开发的v2ray面板

开发文档:  
- https://goframe.org/pages/viewpage.action?pageId=1114399  

更新日志:   
- https://github.com/pingProMax/gov2panel/blob/master/UpdateLog.MD  

* * *  
高性能的面板，巨量用户和巨量节点用起来cpu和内存占用率极低
* * *  

# 节点
节点地址   
``` json 
$uuid$ 订阅时会替换成用户的uuid   
```

数据示例(使用时请删除注释)：
``` json  

协议:v2ray/vmess
tcp:
{
  "port": "8888",
  "transport_protocol": "tcp",

  "scy": "auto",
  "net": "tcp"
}

协议:v2ray/vmess
tcp + tls:
{
  "port": "8888",
  "transport_protocol": "tcp",
  "enable_t_l_s":true,

  "scy": "auto",
  "net": "tcp",
  "type": "none",
  "tls": "tls",
  "sni": "local.id.seewo.com",
  "alpn": "http/1.1",
  "fp": "chrome"
}


协议:v2ray/vmess
ws:
{
  "port": "8888",
  "transport_protocol": "ws",

  "scy": "auto",
  "net": "ws",
  "type": "none",
  "host": "iqiyi.com",
  "path": "/test"
}

协议:v2ray/vmess
ws+tls:
{
  "port": "8888",
  "transport_protocol": "ws",
  "enable_t_l_s":true,

  "scy": "auto",
  "net": "ws",
  "type": "none",
  "host": "local.id.seewo.com",
  "path": "/test",
  "tls": "tls",
  "sni": "local.id.seewo.com",
  "alpn": "http/1.1",
  "fp": "chrome"
}


协议:v2ray/vmess
httpupgrade:
{
  "port": "8888",
  "transport_protocol": "httpupgrade",

  "scy": "auto",
  "net": "httpupgrade",
  "type": "none",
  "host": "iqiyi.com",
  "path": "/test"
}

协议:v2ray/vmess
httpupgrade+tls:
{
  "port": "8888",
  "transport_protocol": "httpupgrade",
  "enable_t_l_s":true,

  "scy": "auto",
  "net": "httpupgrade",
  "type": "none",
  "host": "iqiyi.com",
  "path": "/test",
  "tls": "tls",
  "sni": "local.id.seewo.com",
  "alpn": "http/1.1",
  "fp": "chrome"
}

协议:v2ray/vmess
splithttp
{
  "port": "8888",
  "transport_protocol": "splithttp",

  "scy": "auto",
  "net": "splithttp",
  "type": "none",
  "host": "iqiyi.com",
  "path": "/test"
}

协议:v2ray/vmess
splithttp + tls
{
  "port": "8888",
  "transport_protocol": "splithttp",
  "enable_t_l_s":true,

  "scy": "auto",
  "net": "splithttp",
  "type": "none",
  "host": "iqiyi.com",
  "path": "/test",
  "tls": "tls",
  "sni": "local.id.seewo.com",
  "alpn": "http/1.1",
  "fp": "chrome"
}


协议:v2ray/vmess
h2 + tls
{
  "port": "8888",
  "transport_protocol": "h2",
  "enable_t_l_s":true,

  "scy": "auto",
  "net": "h2",
  "type": "none",
  "host": "iqiyi.com",
  "path": "/test",
  "tls": "tls",
  "sni": "local.id.seewo.com",
  "alpn": "http/1.1",
  "fp": "chrome"
}


协议:v2ray/vmess
quic + tls
{
  "port": "8888",
  "transport_protocol": "quic",
  "enable_t_l_s":true,
  "header": {
		"type": "srtp"
  },

  "scy": "auto",
  "net": "quic",
  "type": "srtp",
  "host": "iqiyi.com",
  "path": "/test",
  "tls": "tls",
  "sni": "local.id.seewo.com",
  "alpn": "http/1.1",
  "fp": "chrome"
}


协议:v2ray/vmess
quic + srtp + tls
{
  "port": "8888",
  "transport_protocol": "quic",
  "enable_t_l_s":true,
  "header": {
		"type": "srtp"
  },

  "scy": "auto",
  "net": "quic",
  "type": "srtp",
  "host": "iqiyi.com",
  "path": "/test",
  "tls": "tls",
  "sni": "local.id.seewo.com",
  "alpn": "http/1.1",
  "fp": "chrome"
}

协议:v2ray/vmess
grpc + tls
{
  "port": "8888",
  "transport_protocol": "grpc",
  "enable_t_l_s":true,
  "authority": "local.id.seewo.com",
  "service_name": "local.id.seewo.com",

  "scy": "auto",
  "net": "grpc",
  "type": "gun",
  "host": "local.id.seewo.com",
  "path": "local.id.seewo.com",
  "tls": "tls",
  "sni": "local.id.seewo.com",
  "alpn": "http/1.1",
  "fp": "chrome"
}


协议:v2ray/vless 注释描述  
说明：xrayr 开启vless协议需要去config.yml 设置NodeType: Vless/EnableVless: true  
{
  "port": "8888",
  "transport_protocol": "tcp", //传输协议 (后端用)
  "security": "tls", //传输层安全 (后端用)  (订阅用)
  "reject_unknown_sni": true, //拒绝未知Sni
  "enable_t_l_s": true,//启用 tls

  "encryption": "", //加密方式
  "security": "", //传输层安全
  "flow": "xtls-rprx-vision", //流控
  "sni": "",  //sni
  "alpn": "", //alpn
  "fp": "", //fp
  "pbk": "",  //PublicKey
  "sid": "",  //ShortId
  "spx": "",  //SpiderX
  "type": "tcp", //传输协议 (订阅用)
  "serviceName": "", //grpc serviceName
  "mode": "", //grpc 模式
  "headerType": "",
  "quicSecurity": "", //quid 加密方式
  "key": "", //quid 加密密钥
  "host": "",
  "path": "",
  "seed": "" //kcp seed
}


协议:v2ray/vless
tcp + tls
{
  "port": "8888",
  "transport_protocol": "tcp", 
  "reject_unknown_sni": true, 
  "enable_t_l_s": true,     
  "enable_vless": true, 
  "security": "tls", 

  "encryption": "none",
  "flow": "xtls-rprx-vision", 
  "headerType": "none",
  "type": "tcp",
  "sni": "local.id.seewo.com",
  "alpn": "http/1.1",
  "fp": "chrome"
}

协议:v2ray/vless
tcp + reality
{
  "port": "8888",
  "transport_protocol": "tcp", 
  "reject_unknown_sni": true, 
  "enable_r_e_a_l_i_t_y": true,    
  "security": "reality", 
 

  "encryption": "none",
  "flow": "xtls-rprx-vision", 
  "headerType": "none",
  "type": "tcp",
  "sni": "local.id.seewo.com",
  "alpn": "http/1.1",
  "fp": "chrome",
  "pbk": "z7JzZnpZIUjr2FZXouK3ij11l-PiL0pnSVmYUxpv_Ss",
	"sid": "0123456789abcdef",
  "headerType": "none",
  "spx": "/$uuid$/cccc"
}

......
其他vless协议 参考上面vmess协议


shadowsocks2022
协议:shadowsocks/ss2022
{
	"port": "8888",
  "transport_protocol": "tcp", 
  "cypher_method":"aes-128-gcm",
	"server_key":"", // 2022-blake3-aes-128-gcm 16个字符->base64 || 2022-blake3-aes-256-gcm 32个字符->base64
  "header":{}
}

协议:trojan/trojan
trojan + tcp + tls
{
	"port":"8888",
  "transport_protocol":"tcp",
  "enable_t_l_s": true,     
	"host":"local.id.seewo.com",
	"service_name":"local.id.seewo.com",
  
  "sni":"local.id.seewo.com",
  "security":"tls", 
  "alpn":"h2",
  "fp":"chrome", 
  "type":"tcp",
  "headerType":""
}

trojan + tcp + http + tls
{
	"port":"8888",
  "transport_protocol":"tcp",
  "enable_t_l_s": true,     
	"host":"local.id.seewo.com",
	"service_name":"local.id.seewo.com",
  "header": {
		"type": "http",
		"request": {
		},
		"response": {
		}
	},
  
  "sni":"local.id.seewo.com",
  "security":"tls", 
  "alpn":"h2",
  "fp":"chrome", 
  "type":"tcp",
  "headerType":"http"
}

协议:trojan/trojan
trojan + ws + tls
{
	"port":"8888",
  "transport_protocol":"ws",
  "enable_t_l_s": true,     
	"host":"local.id.seewo.com",
	"service_name":"local.id.seewo.com",
  
  "sni":"local.id.seewo.com",
  "security":"tls", 
  "alpn":"h2",
  "fp":"chrome", 
  "type":"ws",
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


# clash订阅自定义配置  
``` json  
项目文件夹/manifest/config/clash.yaml
```

# 当前阶段任务  
测试和修复BUG  

# 计划任务  
windows 客户端 √ 
安卓 客户端   √



# 安装教程  
xrayr后端：
https://github.com/XrayR-project/XrayR , PanelType填写"GoV2Panel"  

建议自行拉取编译，打包好的可能比较落后...  
还需要以下两个静态资源文件夹  
resource  
manifest  

1、创建数据库导入数据  
sql下载： https://github.com/pingProMax/gov2panel/releases/download/gov2panel/gov2panel.sql  

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

# 客户端 
https://github.com/pingProMax/GoV2App  
https://github.com/pingProMax/gov2app_flutter  


# 后台展示  
![后台界面](https://github.com/pingProMax/gov2panel/blob/master/admin.png?raw=true)

# 适配 v2board 用户BCRYPT密码的加密方式  
如要导入本项目,password_algo字段固定值`BCRYPT`,其他的请看表注释,后面用户修改密码会自动重置为MD5,以下是导入示例  
1、把v2board v2_user用户表复制或导入到gov2panel数据库中，命名为：v2board_user;   

2、执行插入语句,把v2board_user数据插入到v2_user表中： 
``` sql
INSERT INTO `v2_user`(
  invite_user_id,user_name,password,password_algo,password_salt,balance,discount,commission_type,commission_rate,commission_balance,commission_code,u,d,transfer_enable,banned,is_admin,is_staff,uuid,group_id,token,remarks,expired_at,created_at,updated_at) 
SELECT 
  invite_user_id,email,password,"BCRYPT" as password_algo,password_salt,balance / 100,discount,commission_type,commission_rate,commission_balance / 100,
  CONCAT(
    SUBSTRING('AaBbCcDdEdFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz0123456789', FLOOR(RAND() * 36) + 1, 1),
    SUBSTRING('AaBbCcDdEdFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz0123456789', FLOOR(RAND() * 36) + 1, 1),
    SUBSTRING('AaBbCcDdEdFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz0123456789', FLOOR(RAND() * 36) + 1, 1),
    SUBSTRING('AaBbCcDdEdFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz0123456789', FLOOR(RAND() * 36) + 1, 1),
    SUBSTRING('AaBbCcDdEdFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz0123456789', FLOOR(RAND() * 36) + 1, 1),
    SUBSTRING('AaBbCcDdEdFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz0123456789', FLOOR(RAND() * 36) + 1, 1),
    SUBSTRING('AaBbCcDdEdFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz0123456789', FLOOR(RAND() * 36) + 1, 1)
  ) AS commission_code,u,d,transfer_enable,banned,is_admin,is_staff,uuid,group_id,token,"v2board_user" as remarks,expired_at,created_at,updated_at  FROM `v2board_user`;    
```

3、手动创建订阅权限，注意订阅权限类型要为覆盖；  
4、更新用户的订阅组(这里用remarks字段做下区分,表示是v2board_user刚导入的):    
``` sql
update `v2_user` set group_id = 填手动创建订阅权限ID, remarks = '' where group_id = v2board项目中的id and remarks = "v2board_user"

如果v2board_user原套餐时间为长期有效，需要手动设置下时间，否则无法使用: 
update `v2_user` set expired_at = "2034-12-12" where (expired_at is null or expired_at < '1970-01-01') and group_id = 填手动创建订阅权限ID;

```

