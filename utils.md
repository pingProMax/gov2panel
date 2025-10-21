# 初始化项目sql语句

-- 删除表格中所有数据
DELETE from v2_coupon;
DELETE from v2_coupon_use;
DELETE from v2_invitation_records;
DELETE from v2_knowledge;
DELETE from v2_payment;
DELETE from v2_plan;
DELETE from v2_proxy_service;
DELETE from v2_recharge_records;
DELETE from v2_server_route;
DELETE from v2_ticket;
DELETE from v2_ticket_message;
DELETE from v2_user;

-- 修改表格AUTO_INCREMENT的初始值
ALTER TABLE v2_coupon AUTO_INCREMENT = 1;
ALTER TABLE v2_coupon_use AUTO_INCREMENT = 1;
ALTER TABLE v2_invitation_records AUTO_INCREMENT = 1;
ALTER TABLE v2_knowledge AUTO_INCREMENT = 1;
ALTER TABLE v2_payment AUTO_INCREMENT = 1;
ALTER TABLE v2_plan AUTO_INCREMENT = 1;
ALTER TABLE v2_proxy_service AUTO_INCREMENT = 1;
ALTER TABLE v2_recharge_records AUTO_INCREMENT = 1;
ALTER TABLE v2_server_route AUTO_INCREMENT = 1;
ALTER TABLE v2_ticket AUTO_INCREMENT = 1;
ALTER TABLE v2_ticket_message AUTO_INCREMENT = 1;
ALTER TABLE v2_user AUTO_INCREMENT = 1;

# 创建管理员
INSERT INTO `v2_user`(`invite_user_id`,`telegram_id`,`user_name`,`password`,`password_algo`,`password_salt`,`balance`,`discount`,`commission_type`,`commission_rate`,`commission_balance`,`commission_code`,`t`,`u`,`d`,`transfer_enable`,`banned`,`is_admin`,`is_staff`,`last_login_at`,`last_login_ip`,`uuid`,`group_id`,`token`,`remarks`,`expired_at`,`created_at`,`updated_at`) VALUES(0,0,'adminx','6639b10c5cbcd478da4010d97a3c74de','MD5','ebfbb20e',0,0,3,0,0,'422a3927',0,0,0,0,-1,1,1,0,"",'68d2c51f-5311-4761-aa3a-9e4852f30f25',0,'ab2e2cd2ff064e1990f9b3d60157026e','system init admin',null,'2023-09-20 17:28:53','2023-09-20 17:28:53')


SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -trimpath main.go

bug修复记录
vless 订阅flow和security位置相反
订阅fp 字段



用户每天的流量使用情况 (记录7天的)
USER_%s_%s_FLOW_UPLOAD,UID,2023922

服务器最后在线时间
SERVER_%s_LAST_PUSH_AT,服务器id

服务器当前用户在线数量
SERVER_%s_ONLINE_USER,服务器id

服务器当天的流量使用情况 (记录两天的)
SERVER_%s_%s_FLOW,服务器id,2023922


gf gen ctrl
gf gen service