/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80100 (8.1.0)
 Source Host           : localhost:3306
 Source Schema         : gov2panel

 Target Server Type    : MySQL
 Target Server Version : 80100 (8.1.0)
 File Encoding         : 65001

 Date: 14/09/2023 14:59:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for v2_coupon
-- ----------------------------
DROP TABLE IF EXISTS `v2_coupon`;
CREATE TABLE `v2_coupon`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '优惠码',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `type` tinyint(1) NOT NULL COMMENT '类型 1金额优惠 2百分比优惠',
  `value` decimal(10, 2) NOT NULL COMMENT '优惠多少',
  `enable` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否启用',
  `limit_use` int NULL DEFAULT NULL COMMENT '每个用户可使用次数',
  `limit_use_with_user` int NULL DEFAULT NULL COMMENT '最大使用次数',
  `limit_plan_id` int NULL DEFAULT NULL COMMENT '指定订阅',
  `started_at` timestamp NULL DEFAULT NULL COMMENT '有效期开始时间',
  `ended_at` timestamp NULL DEFAULT NULL COMMENT '有效期结束时间',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `remarks` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of v2_coupon
-- ----------------------------
INSERT INTO `v2_coupon` VALUES (9, '2023code', '2023优惠', 2, 30.00, 1, -1, -1, 0, '2023-09-06 18:16:00', '2024-12-31 18:16:00', '2023-09-07 18:16:38', '2023-09-07 18:17:37', '2023常规优惠码');

-- ----------------------------
-- Table structure for v2_coupon_use
-- ----------------------------
DROP TABLE IF EXISTS `v2_coupon_use`;
CREATE TABLE `v2_coupon_use`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `coupon_id` int NOT NULL,
  `user_id` int NOT NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  `plan_id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `coupon_id`(`coupon_id` ASC) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  INDEX `plan_id`(`plan_id` ASC) USING BTREE,
  CONSTRAINT `v2_coupon_use_ibfk_1` FOREIGN KEY (`coupon_id`) REFERENCES `v2_coupon` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `v2_coupon_use_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `v2_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `v2_coupon_use_ibfk_3` FOREIGN KEY (`plan_id`) REFERENCES `v2_plan` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of v2_coupon_use
-- ----------------------------
INSERT INTO `v2_coupon_use` VALUES (1, 9, 3782, '2023-09-07 18:17:53', '2023-09-07 18:17:53', 5);
INSERT INTO `v2_coupon_use` VALUES (2, 9, 3782, '2023-09-07 18:21:16', '2023-09-07 18:21:16', 5);
INSERT INTO `v2_coupon_use` VALUES (3, 9, 3782, '2023-09-07 18:27:25', '2023-09-07 18:27:25', 3);
INSERT INTO `v2_coupon_use` VALUES (4, 9, 3782, '2023-09-07 18:28:45', '2023-09-07 18:28:45', 3);
INSERT INTO `v2_coupon_use` VALUES (5, 9, 3782, '2023-09-07 18:35:10', '2023-09-07 18:35:10', 3);
INSERT INTO `v2_coupon_use` VALUES (6, 9, 3782, '2023-09-07 18:52:35', '2023-09-07 18:52:35', 5);
INSERT INTO `v2_coupon_use` VALUES (7, 9, 3785, '2023-09-12 15:24:37', '2023-09-12 15:24:37', 5);
INSERT INTO `v2_coupon_use` VALUES (8, 9, 3785, '2023-09-12 15:45:37', '2023-09-12 15:45:37', 3);

-- ----------------------------
-- Table structure for v2_invitation_records
-- ----------------------------
DROP TABLE IF EXISTS `v2_invitation_records`;
CREATE TABLE `v2_invitation_records`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `amount` decimal(10, 2) NOT NULL COMMENT '金额',
  `user_id` int NOT NULL COMMENT '邀请者',
  `from_user_id` int NOT NULL COMMENT '被邀请者',
  `commission_rate` int NOT NULL COMMENT '佣金比例',
  `recharge_records_id` int NOT NULL COMMENT '订单id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `operate_type` tinyint(1) NOT NULL COMMENT '1邀请 2提现',
  `state` tinyint(1) NOT NULL COMMENT '状态 0未审核 1审核 2拒绝',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  INDEX `from_user_id`(`from_user_id` ASC) USING BTREE,
  CONSTRAINT `v2_invitation_records_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `v2_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `v2_invitation_records_ibfk_2` FOREIGN KEY (`from_user_id`) REFERENCES `v2_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of v2_invitation_records
-- ----------------------------

-- ----------------------------
-- Table structure for v2_knowledge
-- ----------------------------
DROP TABLE IF EXISTS `v2_knowledge`;
CREATE TABLE `v2_knowledge`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `category` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '分類名',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '標題',
  `body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '內容',
  `order_id` int NULL DEFAULT NULL COMMENT '排序',
  `show` tinyint(1) NOT NULL DEFAULT 0 COMMENT '顯示',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '知識庫' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of v2_knowledge
-- ----------------------------
INSERT INTO `v2_knowledge` VALUES (2, '手机教程', '安卓手机v2rayNG使用教程', '### v2rayNG官网下载地址：https://github.com/2dust/v2rayNG/releases\n---\n### 你的v2rayNG订阅地址：{{subscribeUrl}}&flag=v2rayng\n---\n\nV2RayNG添加订阅  \n打开 V2RayNG ，点击左上角三道杠标志，点击\"**订阅设置**\"进入订阅设置；  \n点击 \"**+**\" 以添加新的订阅。请粘贴自己的订阅地址。  \n<br>\n\n如何启动 V2RayNG  \n点击右下角 ”V“ 字图标启动代理。首次配置会提示是否创建代理，请允许。  \n<br>\n\nV2RayNG更多设置  \n进入\"**设置**\"页面，找到路由模式，可以设置代理模式。', 2, 1, '2023-08-26 14:41:45', '2023-08-26 14:41:45');
INSERT INTO `v2_knowledge` VALUES (3, '电脑教程', 'Windows-v2rayN使用教程', '### v2rayN官网下载地址：https://github.com/2dust/v2rayN/releases\n---\n### 你的v2rayN订阅地址：{{subscribeUrl}}&flag=v2rayn\n---\n\n\n##### ①下载后解压压缩包，打开\"v2rayN.exe\"应用程序；  \n![图片](/img/Windows-v2rayN%E4%BD%BF%E7%94%A8%E6%95%99%E7%A8%8B/1.png)  \n<br><br>\n\n##### ②点击右下角任务栏“v2ray图标”打开软件界面；  \n![图片](/img/Windows-v2rayN%E4%BD%BF%E7%94%A8%E6%95%99%E7%A8%8B/1.1.png)  \n<br><br>\n\n##### ③设置订阅，如图所示；  \n![图片](/img/Windows-v2rayN%E4%BD%BF%E7%94%A8%E6%95%99%E7%A8%8B/2.png)  \n<br><br>\n\n##### ④更新订阅，如图所示；  \n![图片](/img/Windows-v2rayN%E4%BD%BF%E7%94%A8%E6%95%99%E7%A8%8B/3.png)  \n<br><br>\n\n##### ⑤选择节点，如图所示；   \n![图片](/img/Windows-v2rayN%E4%BD%BF%E7%94%A8%E6%95%99%E7%A8%8B/4.png)  \n<br><br>\n\n##### ⑥选择“自动配置系统代理”启用代理；  \n路由\"全局\"是指代理所有；  \n路由\"全局\"是指代理所有； \n路由\"绕过大陆\"是指大陆网站不使用代理，非大陆网站使用代理。  \n![图片](/img/Windows-v2rayN%E4%BD%BF%E7%94%A8%E6%95%99%E7%A8%8B/5.png)', 4, 1, '2023-08-26 14:41:45', '2023-08-26 14:41:45');
INSERT INTO `v2_knowledge` VALUES (4, '手机教程', '苹果手机Shadowrocket小火箭使用教程', '### Shadowrocket小火箭下载：抱歉本站不提供下载\n---\n### 你的Shadowrocket小火箭订阅地址：{{subscribeUrl}}&flag=shadowrocket\n---\n节点订阅设置  \n打开Shadowrocket，点击右上角加号 + ，在添加节点页面，将类型改为第三个 Subscribe，复制订阅地址粘贴到URL中，然后点击右上角完成即可。  \n![](/img/苹果手机Shadowrocket小火箭教程/1.png)  \n<br><br>\n\n规则的设置  \n一般全局路由选择默认 配置 即可。  \n全局路由的功能：全局路由有四个选项，每个选项都对应着不同的功能，以下将对这些功能进行一一的列举说明。  \n配置：小火箭内置的代理，国内网站直连，国外常用网站走代理，这样可以保证安全性。  \n代理：上网产生的流量都走代理的路线，这样不但可以隐蔽网络，还可以保障用户的隐私安全。  \n直连：如果点击了直连，那么手机当前就是出于脱离小火箭的状态。  \n场景：可以进行自定场景的定位，还可以根据当前网速来决定使用何种服务节点。  \n如果要实现部分网站国内走直连，国外走代理的话，就需要用到规则。', 3, 1, '2023-08-26 14:41:45', '2023-08-26 14:41:45');
INSERT INTO `v2_knowledge` VALUES (5, '电脑教程', 'Mac-v2rayU使用教程', '### 官网下载地址：https://github.com/yanue/V2rayU/releases\r\n---', 5, 0, '2023-08-26 14:41:45', '2023-08-26 16:03:45');
INSERT INTO `v2_knowledge` VALUES (6, '常见问题', '已经成功更新订阅，但是无法使用？', '1、检查时间，时区应为\"**东八区**\"（北京/上海时间）；    \n2、检查是否装有浏览器插件，或者其他代理软件没有退出；  \n3、检查本地网络环境，是否屏蔽了海外IP；', 7, 1, '2023-08-26 14:41:45', '2023-08-26 14:41:45');
INSERT INTO `v2_knowledge` VALUES (8, '常见问题', 'cf节点自选IP', '<h1>1111111</h1><h1>1111111</h1><h1>1111111</h1>', 10, 1, '2023-08-26 14:41:45', '2023-09-01 11:43:23');
INSERT INTO `v2_knowledge` VALUES (9, '常见问题', '无法更新订阅？', '### 1、老订阅地址失效，尝试重新复制订阅；\n### 2、重新复制订阅还是无法更新，请耐心等待修复；', 6, 1, '2023-08-26 14:41:45', '2023-08-26 14:41:45');

-- ----------------------------
-- Table structure for v2_payment
-- ----------------------------
DROP TABLE IF EXISTS `v2_payment`;
CREATE TABLE `v2_payment`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `uuid` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'uuid',
  `payment` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '支付类型',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名字',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图标地址',
  `config` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '配置json数',
  `notify_domain` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '回调域名',
  `handling_fee_fixed` decimal(10, 2) NULL DEFAULT NULL COMMENT '固定手续费',
  `handling_fee_percent` int NULL DEFAULT NULL COMMENT '百分比手续费',
  `enable` int NOT NULL DEFAULT 0 COMMENT '是否启用',
  `order_id` int NULL DEFAULT NULL COMMENT '顺序',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `remarks` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of v2_payment
-- ----------------------------
INSERT INTO `v2_payment` VALUES (6, '', 'epay', '易pay', 'https://gw.alipayobjects.com/mdn/member_frontWeb/afts/img/A*h7o9Q4g2KiUAAAAAAAAAAABkARQnAQ', '{\r\n	\"url\": \"https://go.etopay.top/\",\r\n	\"pid\": 1060,\r\n	\"key\": \"QC2w4NAQwaf88Ca2AU0Z20qaZqNf0Z68\"\r\n}', 'https://127.0.0.1:8080', 1.00, 10, 1, 0, '2023-09-12 16:50:23', '2023-09-14 09:24:32', '');
INSERT INTO `v2_payment` VALUES (7, '', 'epay', '支付宝|微信|USDT', 'https://td.cdn-go.cn/enterprise_payment/v0.1.2/logo.svg', '', '', 0.00, 0, 1, 0, '2023-09-13 13:36:52', '2023-09-13 13:36:52', '');

-- ----------------------------
-- Table structure for v2_plan
-- ----------------------------
DROP TABLE IF EXISTS `v2_plan`;
CREATE TABLE `v2_plan`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `transfer_enable` decimal(10, 2) NOT NULL COMMENT '流量(GB)',
  `speed_limit` int NOT NULL COMMENT '速度限制',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `show` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否显示',
  `order_id` int NOT NULL COMMENT '顺序',
  `renew` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否允许续购',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '描述',
  `expired` int NOT NULL COMMENT '有效期 day',
  `price` decimal(10, 2) NOT NULL COMMENT '价格',
  `reset_traffic_method` tinyint(1) NOT NULL COMMENT '套餐类型，1 覆盖、2 叠加',
  `capacity_limit` int NOT NULL COMMENT '最大用户',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `remarks` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of v2_plan
-- ----------------------------
INSERT INTO `v2_plan` VALUES (3, 100.00, 30, '月包', 1, 10, 1, '', 30, 10.00, 1, 0, '2023-08-20 21:51:59', '2023-08-20 21:51:59', '每月100G流量');
INSERT INTO `v2_plan` VALUES (4, 200.00, 0, '月包 Plus', 1, 20, 1, '', 30, 20.00, 1, 0, '2023-08-21 16:37:57', '2023-08-21 16:37:57', '一个月200G流量');
INSERT INTO `v2_plan` VALUES (5, 100.00, 0, '特价优惠', 1, 30, 1, '特价优惠套餐手动阀ccc', 30, 5.00, 1, 0, '2023-08-21 16:58:27', '2023-09-09 21:28:54', 'aaaaa');
INSERT INTO `v2_plan` VALUES (6, 10.00, 0, '流量叠加包', 1, 0, 1, '', 10, 10.00, 2, 0, '2023-08-21 17:02:12', '2023-09-07 18:37:38', '');

-- ----------------------------
-- Table structure for v2_proxy_service
-- ----------------------------
DROP TABLE IF EXISTS `v2_proxy_service`;
CREATE TABLE `v2_proxy_service`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `agreement` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '协议',
  `service_json` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '服务器json数据',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '显示名称',
  `plan_id` int NULL DEFAULT NULL COMMENT '所属订阅组',
  `show` tinyint(1) NULL DEFAULT NULL COMMENT '是否显示',
  `host` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '服务器地址',
  `port` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '服务器端口',
  `rate` int NULL DEFAULT NULL COMMENT '倍率',
  `order_id` int NOT NULL COMMENT '顺序',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of v2_proxy_service
-- ----------------------------
INSERT INTO `v2_proxy_service` VALUES (1, 'vmell', 'test', 'testname', 3, 1, '127.0.0.1', '1231', 1, 1, '2023-08-27 21:48:27', '2023-08-27 21:48:27');
INSERT INTO `v2_proxy_service` VALUES (2, '5', '7', '1', 0, 1, '2', '3', 4, 6, '2023-08-27 23:57:24', '2023-09-08 18:17:29');
INSERT INTO `v2_proxy_service` VALUES (3, '55', '77', '11', 4, 0, '22', '33', 44, 66, '2023-08-31 15:24:25', '2023-09-08 18:14:05');

-- ----------------------------
-- Table structure for v2_recharge_records
-- ----------------------------
DROP TABLE IF EXISTS `v2_recharge_records`;
CREATE TABLE `v2_recharge_records`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `amount` decimal(10, 2) NOT NULL COMMENT '金额',
  `user_id` int NULL DEFAULT NULL COMMENT '用户id',
  `operate_type` tinyint(1) NULL DEFAULT NULL COMMENT '1充值 2消费',
  `recharge_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '充值类型 operate_type=1才有',
  `consumption_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '消费类型 operate_type=2才有',
  `remarks` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `transaction_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '订单号 规则看程序注释',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  CONSTRAINT `v2_recharge_records_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `v2_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 36 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of v2_recharge_records
-- ----------------------------

-- ----------------------------
-- Table structure for v2_setting
-- ----------------------------
DROP TABLE IF EXISTS `v2_setting`;
CREATE TABLE `v2_setting`  (
  `code` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `order_id` int NULL DEFAULT NULL COMMENT '顺序',
  `remarks` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '备注',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of v2_setting
-- ----------------------------
INSERT INTO `v2_setting` VALUES ('admin_path', 'admintest', 20, '后台地址，设置完需要重启程序才生效', '2023-08-19 11:57:02', '2023-09-12 16:46:52');
INSERT INTO `v2_setting` VALUES ('commission_rate', '10', 40, '系统默认佣金比例，百分比', '2023-09-09 15:23:07', '2023-09-12 16:47:07');
INSERT INTO `v2_setting` VALUES ('commission_type', '2', 45, '系统默认佣金类型；1：循环，2：一次性', '2023-09-09 15:25:10', '2023-09-12 16:47:16');
INSERT INTO `v2_setting` VALUES ('login_page_html', '公告：如发现无法使用，请尝试更新订阅！', 2, '登录页面提示html', '2023-08-20 10:49:12', '2023-09-12 16:47:27');
INSERT INTO `v2_setting` VALUES ('recharge_page_html', '提示：如果充值未到账请提交工单或者tg群联系我们！\r\n充值价格 = 商品价格 + (商品价格 * 百分比手续费/100) + 固定手续费', 4, '充值页面提示', '2023-09-13 15:08:52', '2023-09-13 15:16:55');
INSERT INTO `v2_setting` VALUES ('register_page_html', '公告：请牢记账号密码，不支持密码找回，只能修改密码，切勿分享账号！', 1, '注册页面提示html', '2023-08-20 20:49:08', '2023-09-12 16:47:20');
INSERT INTO `v2_setting` VALUES ('title', 'v2ray订阅', 30, '站点名字', '2023-08-27 10:41:04', '2023-09-12 16:47:01');
INSERT INTO `v2_setting` VALUES ('user_page_html', '欢迎使用', 3, '用户登录后必看公告', '2023-09-12 16:47:54', '2023-09-12 16:48:05');

-- ----------------------------
-- Table structure for v2_ticket
-- ----------------------------
DROP TABLE IF EXISTS `v2_ticket`;
CREATE TABLE `v2_ticket`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `subject` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `level` tinyint(1) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '0:已开启 1:已关闭',
  `reply_status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '0:待回复 1:已回复',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  CONSTRAINT `v2_ticket_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `v2_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 176 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of v2_ticket
-- ----------------------------
INSERT INTO `v2_ticket` VALUES (1, 3782, '我才刚买不久，这个月买的，帮我恢复下', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-27 16:42:19');
INSERT INTO `v2_ticket` VALUES (2, 3782, '订阅消失', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-27 16:42:13');
INSERT INTO `v2_ticket` VALUES (3, 3782, '帐号换了', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-27 16:42:05');
INSERT INTO `v2_ticket` VALUES (4, 3782, '订单无法付款', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (5, 3782, '支付方式出错', 1, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (6, 3782, '没有教程使用文档', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (7, 3782, '我怎么付不成款', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (8, 3782, '无法购买套餐', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (9, 3782, '不能支付', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (10, 3782, '订阅了没线路', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (11, 3782, '已成功付费，但找不到订阅地址', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (12, 3782, '没节点 骗钱呢', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (13, 3782, '显示无效订阅', 0, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (14, 3782, '我的账号为啥没了？', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (15, 3782, '账号丢失', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (16, 3782, '111', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (17, 3782, '我说怎么登陆不上', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (18, 3782, '无法更新节点', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (19, 3782, '账号出现问题', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (20, 3782, '以前订阅的没了？', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (21, 3782, '以前的订阅没了？', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (22, 3782, '才冲的年费', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (23, 3782, '无法登录 & 无法收到忘记密码邮件', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (24, 3782, '之前购买的订阅不见了', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (25, 3782, '4/23充值240G不限时套餐，账号里面订单丢失', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (26, 3782, 'shadow订阅请求超时', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (27, 3782, '账号清空', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (28, 3782, '您好，我之前购买的是三个月的套餐。', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (29, 3782, '为啥我买了2次还是240G', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (30, 3782, '使用文档打不开', 0, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (31, 3782, '不小心买了两次', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (32, 3782, '购买之后连接不上clash', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (33, 3782, '重复购买', 2, 1, 0, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (34, 3782, '怎么恢复原订单', 0, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (35, 3782, '补偿套餐到账', 2, 1, 0, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (36, 3782, '一次性订阅未到账', 0, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (37, 3782, '账号没了', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (38, 3782, '更新订阅失败', 0, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (39, 3782, '订阅号在哪查？', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (40, 3782, '为什么我之前的账号没有了？被注销了？真6', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (41, 3782, '多次订阅补偿套餐，未生效', 1, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (42, 3782, '重复付款', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (43, 3782, '连续两天购买流量包，只显示一个订阅地址，要求退费一个流量包购买费用或者提供两个订阅地址。', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (44, 3782, '买完用不了', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (45, 3782, '我的账户怎么被注销了', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (46, 3782, '电脑端无效的订阅内容', 0, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (47, 3782, '旧账号都无法恢复了吗？', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (48, 3782, '用不了', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (49, 3782, '我的另一账号登录不上了', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (50, 3782, '怎么更新不了订阅', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (51, 3782, '订阅地址', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (52, 3782, '2023042714044549855524080', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (53, 3782, '使用文档没了', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (54, 3782, '支付了现实未支付', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (55, 3782, '怎样使用', 0, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (56, 3782, '找不到我的订阅地址', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (57, 3782, '新购买补偿套餐没有提供订阅url', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (58, 3782, '买了两个用完为止套餐，第二个套餐没效果', 1, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (59, 3782, '免费订单会累积到期时间&一直重置100G', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (60, 3782, '订阅失败', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (61, 3782, '刚缴了一年费用就没了？', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (62, 3782, '了解到您的硬盘问题，那我没用完的年套餐怎么办', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (63, 3782, '无法订阅', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (64, 3782, 'vpn', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (65, 3782, '误操作导致用完即止套餐付款两次', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (66, 3782, '被攻击前的token失效', 2, 1, 0, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (67, 3782, '突然账号咋没了', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (68, 3782, '网络原因买了两次补偿套餐导致覆盖', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (69, 3782, '网络原因买了两次补偿套餐导致覆盖', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (70, 3782, '找不到订阅链接', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (71, 3782, '支付方式少', 1, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (72, 3782, '吞我流量了', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (73, 3782, '订阅地址无法使用', 1, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (74, 3782, '给的链接无法更新出服务器', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (75, 3782, '无教程文档', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (76, 3782, '咨询', 2, 1, 0, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (77, 3782, '使用文档空白', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (78, 3782, '梯子没法用', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (79, 3782, '我刚4.2号买的一年套餐，才用了几次，还推荐了两个人？', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (80, 3782, '我前面买的一年套餐，才用了几次？还推荐了两个人？刚买的套餐两个没叠加？', 1, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (81, 3782, '我前面买的一年套餐，才用了几次？还推荐了两个人？刚买的套餐两个没叠加？', 2, 1, 0, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (82, 3782, '我9天前买的用完即止套餐，我的账号直接就登陆不上去了', 2, 1, 0, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (83, 3782, '不小心买了两次不限时套餐，可以退吗', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (84, 3782, '之前的号登不上去，用完为止套餐还有230多g，可以转移到这个号吗', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (85, 3782, '流量扣费计算', 1, 1, 0, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (86, 3782, '全局代理', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (87, 3782, '购买了3个月的套餐现在用不了', 1, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (88, 3782, '免费月付覆盖了补偿套餐', 2, 1, 0, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (89, 3782, 'Connection problem', 2, 1, 0, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (90, 3782, '付款后没有显示支付成功', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (91, 3782, 'v2rayclub是傻逼', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (92, 3782, 'v2rayclub是傻逼', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (93, 3782, 'v2rayclub是傻逼', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (94, 3782, 'v2rayclub是傻逼', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (95, 3782, 'v2rayclub是傻逼', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (96, 3782, 'v2rayclub是傻逼', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (97, 3782, 'v2rayclub是傻逼', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (98, 3782, 'v2rayclub是傻逼', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (99, 3782, 'v2rayclub是傻逼', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (100, 3782, 'v2rayclub是傻逼', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (101, 3782, 'v2rayclub是傻逼', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (102, 3782, 'v2rayclub是傻逼', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (103, 3782, 'v2rayclub是傻逼', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (104, 3782, 'v2rayclub是傻逼', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (105, 3782, '1111', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (106, 3782, '我这个id才刚刚续费就清了？？？', 0, 1, 0, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (107, 3782, '我以前的账号怎么没了？', 2, 1, 0, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (108, 3782, '之前购买过服务的账号无法使用', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (109, 3782, '更新订阅失败', 0, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (110, 3782, '链接无效', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (111, 3782, '今后的使用是否受影响', 1, 1, 0, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (112, 3782, '付款两笔怎么不生效', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (113, 3782, '节点寄了？', 0, 1, 0, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (114, 3782, '线路数量只有一个，而且24小时一半都掉线', 2, 1, 0, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (115, 3782, '已经支付宝支付 却没有订阅是怎么回事', 2, 1, 0, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (116, 3782, '二次购买用完为止套餐流量没有增加', 1, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (117, 3782, '节点显示问题', 2, 1, 0, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (118, 3782, '已支付未生效', 2, 1, 0, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (119, 3782, '关于本次补偿问题', 0, 1, 0, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (120, 3782, '无法更新订阅', 2, 1, 0, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (121, 3782, '不同订阅的区别', 0, 1, 0, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (122, 3782, '不是问题。。就是老用户来看下', 0, 1, 0, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (123, 3782, '流量无法使用', 2, 1, 0, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (124, 3782, '苹果上怎么找不到app了', 1, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (125, 3782, 'alien', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (126, 3782, 'can\'t access internet', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (127, 3782, '订单问题', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (128, 3782, '无法使用', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (129, 3782, 'can‘t visit openai/chatgpt', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (130, 3782, '流量问题', 1, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (131, 3782, '申请退费', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (132, 3782, '节点都不可用了？', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (133, 3782, '买完之后报错 没有v2ray地址', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (134, 3782, '订单号 2023050220054215401114737', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (135, 3782, '节点失效，-1ms延迟', 1, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (136, 3782, '提供 clash 可用的订阅链接', 1, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (137, 3782, '我之前购买的订阅不见了', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (138, 3782, '原账号', 0, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (139, 3782, 'cs', 1, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (140, 3782, '网络不稳定', 0, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (141, 3782, '我买的一直连不上，会不会也是数据被删除的原因？', 1, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (142, 3782, '我购买的套餐以及邮箱注册全没了，还有170G的流量', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (143, 3782, '想请问3.9套餐的安全问题', 1, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (144, 3782, 'macbook如何使用', 1, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (145, 3782, '补偿套餐', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (146, 3782, '套餐没了', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (147, 3782, '已经购买了一年的货', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (148, 3782, 'mac 电脑怎么使用', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (149, 3782, 'سیا', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (150, 3782, '买第二个用完即止后，之前没用完的包被覆盖掉了', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (151, 3782, 'کانفیگ', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (152, 3782, '账号无法登录', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (153, 3782, '我的订阅地址无效了', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (154, 3782, '多买了一个', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (155, 3782, '不能使用', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (156, 3782, '支付后订单仍是未支付', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (157, 3782, 'clash 订阅节点失败', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (158, 3782, '给的链接无法完成订阅更新', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (159, 3782, '无法上网', 2, 1, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (160, 3782, 'window10上使用clash 和v2ray都无法使用', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (161, 3782, '节点经常断网', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (162, 3782, '支付后套餐未生效', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (163, 3782, '我之前的账户没有了', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (164, 3782, '补偿套餐一次性问题', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (165, 3782, '无法下载app', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (166, 3782, 'Mac使用', 2, 0, 1, '2023-08-26 16:46:46', '2023-09-05 15:50:30');
INSERT INTO `v2_ticket` VALUES (167, 3782, 'MacBook设置', 0, 0, 1, '2023-08-26 16:46:46', '2023-09-05 15:38:23');
INSERT INTO `v2_ticket` VALUES (168, 3782, '节点不可使用', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (169, 3782, '无法科学上网', 2, 1, 1, '2023-08-26 16:46:46', '2023-09-08 14:53:59');
INSERT INTO `v2_ticket` VALUES (170, 3782, '连不上', 2, 0, 1, '2023-08-26 16:46:46', '2023-08-26 16:46:46');
INSERT INTO `v2_ticket` VALUES (171, 3782, '无法上网', 2, 0, 1, '2023-08-26 16:46:46', '2023-09-08 17:16:55');
INSERT INTO `v2_ticket` VALUES (172, 3782, '[提现申请] 本工单由系统发出', 2, 0, 0, '2023-08-26 16:46:46', '2023-09-11 15:00:37');
INSERT INTO `v2_ticket` VALUES (173, 3782, '提交工单啦', 1, 0, 0, '2023-09-08 15:56:23', '2023-09-08 15:56:23');
INSERT INTO `v2_ticket` VALUES (174, 3782, 'test', 1, 0, 0, '2023-09-08 16:32:57', '2023-09-08 16:32:57');
INSERT INTO `v2_ticket` VALUES (175, 3785, 'qweqwe测试', 3, 0, 0, '2023-09-08 17:17:59', '2023-09-11 15:10:11');

-- ----------------------------
-- Table structure for v2_ticket_message
-- ----------------------------
DROP TABLE IF EXISTS `v2_ticket_message`;
CREATE TABLE `v2_ticket_message`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `ticket_id` int NOT NULL,
  `message` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  INDEX `ticket_id`(`ticket_id` ASC) USING BTREE,
  CONSTRAINT `v2_ticket_message_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `v2_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `v2_ticket_message_ibfk_2` FOREIGN KEY (`ticket_id`) REFERENCES `v2_ticket` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 228 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of v2_ticket_message
-- ----------------------------
INSERT INTO `v2_ticket_message` VALUES (1, 3782, 1, '支付宝订单号2023040922001470241423169186', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (2, 3782, 2, '你们数据库损坏，导致我的账号都被清空了，之前的订阅也已经消失，新的订阅无法订阅，提示我没有有效的支付渠道，也没有支付渠道给我选择', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (3, 3785, 3, '账号换绑，之前购买的帐号突然不能登陆上去，现在上不了tg群', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (4, 3785, 4, '哎，丢失完了，重新购买但无法付款，请查', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (5, 3785, 5, '无法完成支付，因为支付方式是空的，没有可用选项', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (6, 3782, 6, '使用文档中没有任何', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (7, 3782, 7, '结账的时候，显示所选货币无效？', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (8, 3782, 8, '购买3.9元套餐时显示invalid currency', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (9, 3782, 9, '选择套餐之后不能支付，报错了', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (10, 3782, 10, '你们能把数据弄丢就离谱啊，我还有推广的余额都没用，而且免费订阅没有线路啊，是必须要付费订阅吗', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (11, 3782, 11, '已成功付费，但找不到订阅地址', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (12, 3782, 12, '？？？', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (13, 3782, 13, '重新购买一次性套餐以后  拷贝订阅链接  显示无效订阅', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (14, 3782, 14, '为啥我的账号啥都没了？我购买了订阅的，还有拉了很多的人过来。咋回事啊？', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (15, 3782, 15, '之前的账号2365260132@qq.com密码突然不对了..找回qq收不到…我还有大半年啊啊啊啊啊', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (16, 3782, 16, '没有美国节点了嘛', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (17, 3782, 17, '上一个套餐才买没多久还有200多G没有用，这样就没拉', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (18, 3782, 18, '无法更新节点', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (19, 3782, 19, '340243153@qq.com 这是我之前的账号 登入不进去了 也找不回密码 邮箱收不到验证码 随手注册同账号有注册进去了  我还有100多兆流量https://dy.v2ray.club/api/v1/client/subscribe?token=5b2560f69abb73aff3901ae30d376b82 这是我之前的URL 一直无法更新订阅 才登陆看看 这是什么问题 我确定原先也是这个账号 因为浏览器有保存', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (20, 3782, 20, '订阅还没过期', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (21, 3782, 21, '以前订单号2022020822001413351401829767\n还没过期', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (22, 3782, 22, '才冲的年费  4月3号才冲的，这样就不管了？', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (23, 3782, 23, '我尝试使用我的账户，应该是 chenf@surest.cn \n\n但是我始终无法收到忘记密码的邮件', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (24, 3782, 24, '之前在 v2ray.run 这个网站上购买的39块钱240G流量的订阅不见了，很急，很急，希望能帮忙处理一下', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (25, 3782, 25, '4/23充值240G不限时套餐，账号里面订单丢失，请补偿到我的账号！！！\n前一天充值，后一天你们就丢失所有数据？', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (26, 3782, 26, 'shadow导入订阅链接显示请求超时', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (27, 3782, 27, '账号之前购买了一个月，被清空数据。帮忙查一下', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (28, 3782, 28, '您好，我之前购买的是三个月套餐，你这个补偿只有一个月，而且还需要我再花钱，请问我之前的套餐怎么办？', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (29, 3782, 29, '为啥我买了2次还是240G', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (30, 3782, 30, '使用文档打不开', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (31, 3782, 31, '不小心买了两次流量，流量能叠加吗？', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (32, 3782, 32, '我的url如下：https://dy.v2ray.ws/api/v1/client/subscribe?token=c02477c9ec8bbdb16508000281a5dfa1\n显示Network Error', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (33, 3782, 33, '重复购买了一次性的套餐', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (34, 3782, 34, '怎么恢复原订单需要提供信息吗', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (35, 3782, 35, '这个补偿套餐只能买一个吗，我下单了三个，只有240G。', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (36, 3782, 36, '一次性订阅未到账', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (37, 3782, 37, '账号没了？？？？', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (38, 3782, 38, '更新订阅失败', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (39, 3782, 39, '买了补偿套餐，订阅号查不了。', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (40, 3782, 40, '为什么我之前的账号没有了？被注销了？真6', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (41, 3782, 41, '订单号2：2023042709042310984649888\n订单号1：2023042709041266833558069', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (42, 3782, 42, '购买流量的时候下了两个订单，但是流量并没有累计', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (43, 3782, 43, '我在2023年04月26购买了一个240GB流量包订单号：2023042617042333695291290，第二天2023年04月27日用我的账号帮朋友购买了一次性流量包，显示流量被重置，昨天买的流量包消失了，2023042708043558571649165，要求退费一个流量包购买费用或者提供两个订阅地址。', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (44, 3782, 44, '订阅导入只有0个条目，不行用不了', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (45, 3782, 33, '已经退款', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (46, 3782, 35, '只能买一个', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (47, 3782, 45, '账户被注销了，现在重新注册 账户里原来的套餐也失效了', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (48, 3782, 46, '我在pc端更新订阅内容（不通过代理）系统显示为无效的订阅内容', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (49, 3782, 47, '我购买服务的账号spike.vector@gmail.com，还有不限时套餐240g，基本没怎么用啊？就不补偿了？', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (50, 3782, 48, '我之前花了39买了一次。这次又花了3.9买了一次。都用不了怎么回事', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (51, 3782, 49, '我之前有一个账号chendi2874@163.com ，但是登录不上提示密码或账号错误。怎么回事啊', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (52, 3782, 50, '更新不出来节点', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (53, 3782, 51, '在使用文档里找不到我的订阅地址了，请问在哪可以找到？', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (54, 3782, 52, '怎么没有地址了，怎么使用阿', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (55, 3782, 53, '找不到节点配置', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (56, 3782, 54, '支付了但是显示我未支付', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (57, 3782, 55, '我已经购买了订单，但是使用说明的页面加载不出来，不知道如何使用，需要使用教程', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (58, 3782, 56, '订阅地址看不到', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (59, 3782, 57, '新购买补偿套餐没有提供订阅url 无法使用', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (60, 3782, 58, '买了两个用完为止套餐，第二个套餐没效果', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (61, 3782, 59, '哈哈', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (62, 3782, 60, '重新购买节点后，导入订阅，更新失败。', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (63, 3782, 61, '除了我的大批邀请外，刚缴了一年的费用啊，我邀请的朋友也有很多刚交费的。有办法补偿吗？', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (64, 3782, 62, '了解到您遇到了硬盘问题，那我没用完的年套餐怎么办，目前还剩余180G。以前买的还是这个邮箱，重新注册了一下才联系到您。希望能够获得合理的处理方式，谢谢！', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (65, 3782, 63, '按照提示购买了补偿订单，目前无法订阅成功。麻烦尽快处理，感谢', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (66, 3782, 64, '为什么我的在电脑无法正常使用', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (67, 3782, 65, '误操作导致用完即止套餐付款两次，申请退回一次付款或增加流量', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (68, 3782, 66, '你们的被攻击后账号数据被删除了，但是一开始我的token还能照常用，现在token也失效了，请问有办法补吗（有之前的支付凭证）', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (69, 3782, 67, '原来的账号', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (70, 3782, 68, '所以申请退其中一次费', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (71, 3782, 69, '所以申请退掉其中一次的费用', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (72, 3782, 70, '使用文档没有内容，找不到订阅信息', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (73, 3782, 71, '能不能增加微信支付', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (74, 3782, 72, '买了2次240G 但是仪表盘只显示240G', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (75, 3782, 73, '订阅地址导入后未增加服务器', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (76, 3782, 74, '给的链接无法更新出服务器，试了手机和电脑扫码和不扫码一个个都试过了', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (77, 3782, 75, '无教程文档  无url地址', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (78, 3782, 76, '我以前订阅的密码登陆不上了，还剩下的未使用完的流量还能用吗？', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (79, 3782, 77, '使用文档界面是空白，看不到我的订阅地址', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (80, 3782, 78, '重新订阅了那个一折的，都没法用，不能上网', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (81, 3782, 79, '我刚4.2号买的一年套餐，才用了几次，还推荐了两个人，账号密码一直登不上，重新注册进来的，这么说全部没了？就只有一个用完即止？我可以提供支付截图', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (82, 3782, 80, '1、我刚4.2号买的一年套餐，才用了几次，还推荐了两个人，账号密码这两天一直登不上，重新注册进来的，这么说全部没了？就只有一个用完即止？我可以提供支付截图的？\n2、那个一次性的用完即止是只能买一次吗？我刚买了两个？流量没叠加？这个流量不叠加那我两个订单的流量怎么算的？知道你们的损失，还请在广大网友支持的份儿上回复一下？我做主播推荐的也是你们啊，别让大家失望，兄弟', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (83, 3782, 81, '1、我刚4.2号买的一年套餐，才用了几次，还推荐了两个人，账号密码这两天一直登不上，重新注册进来的，这么说全部没了？就只有一个用完即止？我可以提供支付截图的？ 2、那个一次性的用完即止是只能买一次吗？我刚买了两个？流量没叠加？这个流量不叠加那我两个订单的流量怎么算的？知道你们的损失，还请在广大网友支持的份儿上回复一下？我做主播推荐的也是你们啊，别让大家失望，兄弟', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (84, 3782, 82, '用完即止套餐我才用没7，8天，我的账号就没有了。账号：947091741@qq.com这个是我买了用完即止套餐的账号，我还尝试过忘记密码都没用，压根没有验证码发过来。\n我还有之前的截图和配置更新日期，同时我还能提供9天前的支付宝付款凭证。希望尽快把这个套餐能恢复到我目前这个账号（944327742@qq.com）上', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (85, 3782, 83, '不小心买了两次不限时套餐，可以退一次吗', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (86, 3782, 84, '账号与密码：\n1149769964@qq,com\nlxdbb1212.', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (87, 3782, 85, '请问如果同时购买了流量卡和月卡，优先扣除哪一个的流量呢', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (88, 3782, 86, '配置了绕过IP，但是最终还是会全局代理', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (89, 3782, 87, '4月16日办理的三个月包月，前几天使用正常，现在IP都无法使用，更新出来的订阅都是无效的订阅内容，麻烦帮忙进行修复一下。账号466307346@qq.com', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (90, 3782, 88, '先买了补偿套餐，不小心点了0元的月份，覆盖掉了', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (91, 3782, 89, 'I subscribed to semiannual package in 14 April, from past few days it was not working well, so I delete the server to re enter again, but after that I can’t even sign in to my account, they mentioned your account email or password is wrong, although which are correct. But still I tired to change password but not even sending email to my registered account. And when enter the url link to the server it showed “forbidden” I don’t know why. Now I registered my another email to send this complaint', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (92, 3782, 90, '你好我下单了支付宝付款后订单显示待支付，麻烦处理一下', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (93, 3782, 91, 'v2rayclub是傻逼', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (94, 3782, 92, 'v2rayclub是傻逼', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (95, 3782, 93, 'v2rayclub是傻逼', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (96, 3782, 94, 'v2rayclub是傻逼', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (97, 3782, 95, 'v2rayclub是傻逼', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (98, 3782, 96, 'v2rayclub是傻逼', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (99, 3782, 97, 'v2rayclub是傻逼', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (100, 3782, 98, 'v2rayclub是傻逼', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (101, 3782, 99, 'v2rayclub是傻逼', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (102, 3782, 100, 'v2rayclub是傻逼', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (103, 3782, 101, 'v2rayclub是傻逼', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (104, 3782, 102, 'v2rayclub是傻逼', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (105, 3782, 103, 'v2rayclub是傻逼', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (106, 3782, 104, 'v2rayclub是傻逼', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (107, 3782, 105, '<html><head>\n    <meta charset=\"utf-8\">\n    <title>Telegram: Contact @V2rayrun</title>\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n    <script>try{if(window.parent!=null&&window!=window.parent){window.parent.postMessage(JSON.stringify({eventType:\'web_app_open_tg_link\',eventData:{path_full:\"\\/V2rayrun\"}}),\'https://web.telegram.org\');}}catch(e){}</script>\n    \n<meta property=\"og:title\" content=\"v2ray.run\">\n<meta property=\"og:image\" content=\"data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMzIwIiBoZWlnaHQ9IjMyMCIgcHJlc2VydmVBc3BlY3RSYXRpbz0ibm9uZSIgdmlld0JveD0iMCAwIDEwMCAxMDAiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI%2BPGRlZnM%2BPGxpbmVhckdyYWRpZW50IGlkPSJnIiB4MT0iMCUiIHgyPSIwJSIgeTE9IjAlIiB5Mj0iMTAwJSI%2BPHN0b3Agb2Zmc2V0PSIwJSIgc3RvcC1jb2xvcj0iIzcyZDVmZCIvPjxzdG9wIG9mZnNldD0iMTAwJSIgc3RvcC1jb2xvcj0iIzJhOWVmMSIvPjwvbGluZWFyR3JhZGllbnQ%2BPC9kZWZzPjxzdHlsZT50ZXh0e2ZvbnQ6NjAwIDQ0cHggLWFwcGxlLXN5c3RlbSxCbGlua01hY1N5c3RlbUZvbnQsJ1NlZ29lIFVJJyxSb2JvdG8sSGVsdmV0aWNhLEFyaWFsLHNhbnMtc2VyaWYsJ0FwcGxlIENvbG9yIEVtb2ppJywnU2Vnb2UgVUkgRW1vamknLCdTZWdvZSBVSSBTeW1ib2wnOy13ZWJraXQtdXNlci1zZWxlY3Q6bm9uZTt1c2VyLXNlbGVjdDpub25lfTwvc3R5bGU%2BPHJlY3Qgd2lkdGg9IjEwMCIgaGVpZ2h0PSIxMDAiIGZpbGw9InVybCgjZykiLz48dGV4dCB0ZXh0LWFuY2hvcj0ibWlkZGxlIiB4PSI1MCIgeT0iNjYiIGZpbGw9IiNmZmYiPlY8L3RleHQ%2BPC9zdmc%2B\">\n<meta property=\"og:site_name\" content=\"Telegram\">\n<meta property=\"og:description\" content=\"网址：	netvog.com		备用：(订阅不了请尝试使用备用地址！)	dy.v2ray.club	机器人	@v2rayrun_bot\">\n\n<meta property=\"twitter:title\" content=\"v2ray.run\">\n<meta property=\"twitter:image\" content=\"data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMzIwIiBoZWlnaHQ9IjMyMCIgcHJlc2VydmVBc3BlY3RSYXRpbz0ibm9uZSIgdmlld0JveD0iMCAwIDEwMCAxMDAiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI%2BPGRlZnM%2BPGxpbmVhckdyYWRpZW50IGlkPSJnIiB4MT0iMCUiIHgyPSIwJSIgeTE9IjAlIiB5Mj0iMTAwJSI%2BPHN0b3Agb2Zmc2V0PSIwJSIgc3RvcC1jb2xvcj0iIzcyZDVmZCIvPjxzdG9wIG9mZnNldD0iMTAwJSIgc3RvcC1jb2xvcj0iIzJhOWVmMSIvPjwvbGluZWFyR3JhZGllbnQ%2BPC9kZWZzPjxzdHlsZT50ZXh0e2ZvbnQ6NjAwIDQ0cHggLWFwcGxlLXN5c3RlbSxCbGlua01hY1N5c3RlbUZvbnQsJ1NlZ29lIFVJJyxSb2JvdG8sSGVsdmV0aWNhLEFyaWFsLHNhbnMtc2VyaWYsJ0FwcGxlIENvbG9yIEVtb2ppJywnU2Vnb2UgVUkgRW1vamknLCdTZWdvZSBVSSBTeW1ib2wnOy13ZWJraXQtdXNlci1zZWxlY3Q6bm9uZTt1c2VyLXNlbGVjdDpub25lfTwvc3R5bGU%2BPHJlY3Qgd2lkdGg9IjEwMCIgaGVpZ2h0PSIxMDAiIGZpbGw9InVybCgjZykiLz48dGV4dCB0ZXh0LWFuY2hvcj0ibWlkZGxlIiB4PSI1MCIgeT0iNjYiIGZpbGw9IiNmZmYiPlY8L3RleHQ%2BPC9zdmc%2B\">\n<meta property=\"twitter:site\" content=\"@Telegram\">\n\n<meta property=\"al:ios:app_store_id\" content=\"686449807\">\n<meta property=\"al:ios:app_name\" content=\"Telegram Messenger\">\n<meta property=\"al:ios:url\" content=\"tg://resolve?domain=V2rayrun\">\n\n<meta property=\"al:android:url\" content=\"tg://resolve?domain=V2rayrun\">\n<meta property=\"al:android:app_name\" content=\"Telegram\">\n<meta property=\"al:android:package\" content=\"org.telegram.messenger\">\n\n<meta name=\"twitter:card\" content=\"summary\">\n<meta name=\"twitter:site\" content=\"@Telegram\">\n<meta name=\"twitter:description\" content=\"网址：	netvog.com		备用：(订阅不了请尝试使用备用地址！)	dy.v2ray.club	机器人	@v2rayrun_bot\n\">\n<meta name=\"twitter:app:name:iphone\" content=\"Telegram Messenger\">\n<meta name=\"twitter:app:id:iphone\" content=\"686449807\">\n<meta name=\"twitter:app:url:iphone\" content=\"tg://resolve?domain=V2rayrun\">\n<meta name=\"twitter:app:name:ipad\" content=\"Telegram Messenger\">\n<meta name=\"twitter:app:id:ipad\" content=\"686449807\">\n<meta name=\"twitter:app:url:ipad\" content=\"tg://resolve?domain=V2rayrun\">\n<meta name=\"twitter:app:name:googleplay\" content=\"Telegram\">\n<meta name=\"twitter:app:id:googleplay\" content=\"org.telegram.messenger\">\n<meta name=\"twitter:app:url:googleplay\" content=\"https://t.me/V2rayrun\">\n\n<meta name=\"apple-itunes-app\" content=\"app-id=686449807, app-argument: tg://resolve?domain=V2rayrun\">\n    <script>window.matchMedia&&window.matchMedia(\'(prefers-color-scheme: dark)\').matches&&document.documentElement&&document.documentElement.classList&&document.documentElement.classList.add(\'theme_dark\');</script>\n    <link rel=\"icon\" type=\"image/svg+xml\" href=\"//telegram.org/img/website_icon.svg?4\">\n<link rel=\"apple-touch-icon\" sizes=\"180x180\" href=\"//telegram.org/img/apple-touch-icon.png\">\n<link rel=\"icon\" type=\"image/png\" sizes=\"32x32\" href=\"//telegram.org/img/favicon-32x32.png\">\n<link rel=\"icon\" type=\"image/png\" sizes=\"16x16\" href=\"//telegram.org/img/favicon-16x16.png\">\n<link rel=\"alternate icon\" href=\"//telegram.org/img/favicon.ico\" type=\"image/x-icon\">\n    <link href=\"//telegram.org/css/font-roboto.css?1\" rel=\"stylesheet\" type=\"text/css\">\n    <!--link href=\"/css/myriad.css\" rel=\"stylesheet\"-->\n    <link href=\"//telegram.org/css/bootstrap.min.css?3\" rel=\"stylesheet\">\n    <link href=\"//telegram.org/css/telegram.css?236\" rel=\"stylesheet\" media=\"screen\">\n  </head>\n  <body class=\"\">\n      <div class=\"tgme_background_wrap\">\n    <canvas id=\"tgme_background\" class=\"tgme_background default\" width=\"50\" height=\"50\" data-colors=\"dbddbb,6ba587,d5d88d,88b884\"></canvas>\n    <div class=\"tgme_background_pattern default\"></div>\n  </div>\n    <div class=\"tgme_page_wrap\">\n      <div class=\"tgme_head_wrap\">\n        <div class=\"tgme_head\">\n          <a href=\"//telegram.org/\" class=\"tgme_head_brand\">\n            <svg class=\"tgme_logo\" height=\"34\" viewBox=\"0 0 133 34\" width=\"133\" xmlns=\"http://www.w3.org/2000/svg\">\n              <g fill=\"none\" fill-rule=\"evenodd\">\n                <circle cx=\"17\" cy=\"17\" fill=\"var(--accent-btn-color)\" r=\"17\"></circle><path d=\"m7.06510669 16.9258959c5.22739451-2.1065178 8.71314291-3.4952633 10.45724521-4.1662364 4.9797665-1.9157646 6.0145193-2.2485535 6.6889567-2.2595423.1483363-.0024169.480005.0315855.6948461.192827.1814076.1361492.23132.3200675.2552048.4491519.0238847.1290844.0536269.4231419.0299841.65291-.2698553 2.6225356-1.4375148 8.986738-2.0315537 11.9240228-.2513602 1.2428753-.7499132 1.5088847-1.2290685 1.5496672-1.0413153.0886298-1.8284257-.4857912-2.8369905-1.0972863-1.5782048-.9568691-2.5327083-1.3984317-4.0646293-2.3321592-1.7703998-1.0790837-.212559-1.583655.7963867-2.5529189.2640459-.2536609 4.7753906-4.3097041 4.755976-4.431706-.0070494-.0442984-.1409018-.481649-.2457499-.5678447-.104848-.0861957-.2595946-.0567202-.3712641-.033278-.1582881.0332286-2.6794907 1.5745492-7.5636077 4.6239616-.715635.4545193-1.3638349.6759763-1.9445998.6643712-.64024672-.0127938-1.87182452-.334829-2.78737602-.6100966-1.12296117-.3376271-1.53748501-.4966332-1.45976769-1.0700283.04048-.2986597.32581586-.610598.8560076-.935815z\" fill=\"#fff\"></path><path d=\"m49.4 24v-12.562h-4.224v-2.266h11.198v2.266h-4.268v12.562zm16.094-4.598h-7.172c.066 1.936 1.562 2.772 3.3 2.772 1.254 0 2.134-.198 2.97-.484l.396 1.848c-.924.396-2.2.682-3.74.682-3.476 0-5.522-2.134-5.522-5.412 0-2.97 1.804-5.764 5.236-5.764 3.476 0 4.62 2.86 4.62 5.214 0 .506-.044.902-.088 1.144zm-7.172-1.892h4.708c.022-.99-.418-2.618-2.222-2.618-1.672 0-2.376 1.518-2.486 2.618zm9.538 6.49v-15.62h2.706v15.62zm14.84-4.598h-7.172c.066 1.936 1.562 2.772 3.3 2.772 1.254 0 2.134-.198 2.97-.484l.396 1.848c-.924.396-2.2.682-3.74.682-3.476 0-5.522-2.134-5.522-5.412 0-2.97 1.804-5.764 5.236-5.764 3.476 0 4.62 2.86 4.62 5.214 0 .506-.044.902-.088 1.144zm-7.172-1.892h4.708c.022-.99-.418-2.618-2.222-2.618-1.672 0-2.376 1.518-2.486 2.618zm19.24-1.144v6.072c0 2.244-.462 3.85-1.584 4.862-1.1.99-2.662 1.298-4.136 1.298-1.364 0-2.816-.308-3.74-.858l.594-2.046c.682.396 1.826.814 3.124.814 1.76 0 3.08-.924 3.08-3.234v-.924h-.044c-.616.946-1.694 1.584-3.124 1.584-2.662 0-4.554-2.2-4.554-5.236 0-3.52 2.288-5.654 4.862-5.654 1.65 0 2.596.792 3.102 1.672h.044l.11-1.43h2.354c-.044.726-.088 1.606-.088 3.08zm-2.706 2.948v-1.738c0-.264-.022-.506-.088-.726-.286-.99-1.056-1.738-2.2-1.738-1.518 0-2.64 1.32-2.64 3.498 0 1.826.924 3.3 2.618 3.3 1.012 0 1.892-.66 2.2-1.65.088-.264.11-.638.11-.946zm5.622 4.686v-7.26c0-1.452-.022-2.508-.088-3.454h2.332l.11 2.024h.066c.528-1.496 1.782-2.266 2.948-2.266.264 0 .418.022.638.066v2.53c-.242-.044-.484-.066-.814-.066-1.276 0-2.178.814-2.42 2.046-.044.242-.066.528-.066.814v5.566zm16.05-6.424v3.85c0 .968.044 1.914.176 2.574h-2.442l-.198-1.188h-.066c-.638.836-1.76 1.43-3.168 1.43-2.156 0-3.366-1.562-3.366-3.19 0-2.684 2.398-4.07 6.358-4.048v-.176c0-.704-.286-1.87-2.178-1.87-1.056 0-2.156.33-2.882.792l-.528-1.76c.792-.484 2.178-.946 3.872-.946 3.432 0 4.422 2.178 4.422 4.532zm-2.64 2.662v-1.474c-1.914-.022-3.74.374-3.74 2.002 0 1.056.682 1.54 1.54 1.54 1.1 0 1.87-.704 2.134-1.474.066-.198.066-.396.066-.594zm5.6 3.762v-7.524c0-1.232-.044-2.266-.088-3.19h2.31l.132 1.584h.066c.506-.836 1.474-1.826 3.3-1.826 1.408 0 2.508.792 2.97 1.98h.044c.374-.594.814-1.034 1.298-1.342.616-.418 1.298-.638 2.2-.638 1.76 0 3.564 1.21 3.564 4.642v6.314h-2.64v-5.918c0-1.782-.616-2.838-1.914-2.838-.924 0-1.606.66-1.892 1.43-.088.242-.132.594-.132.902v6.424h-2.64v-6.204c0-1.496-.594-2.552-1.848-2.552-1.012 0-1.694.792-1.958 1.518-.088.286-.132.594-.132.902v6.336z\" fill=\"var(--tme-logo-color)\" fill-rule=\"nonzero\"></path>\n              </g>\n            </svg>\n          </a>\n          <a class=\"tgme_head_right_btn\" href=\"//telegram.org/dl?tme=45ecc5c7460e78941d_13041178241482869553\">\n            Download\n          </a>\n        </div>\n      </div>\n      <div class=\"tgme_body_wrap\">\n        <div class=\"tgme_page\">\n          <div class=\"tgme_page_photo\">\n  <a href=\"tg://resolve?domain=V2rayrun\"><img class=\"tgme_page_photo_image\" src=\"data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMzIwIiBoZWlnaHQ9IjMyMCIgcHJlc2VydmVBc3BlY3RSYXRpbz0ibm9uZSIgdmlld0JveD0iMCAwIDEwMCAxMDAiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI%2BPGRlZnM%2BPGxpbmVhckdyYWRpZW50IGlkPSJnIiB4MT0iMCUiIHgyPSIwJSIgeTE9IjAlIiB5Mj0iMTAwJSI%2BPHN0b3Agb2Zmc2V0PSIwJSIgc3RvcC1jb2xvcj0iIzcyZDVmZCIvPjxzdG9wIG9mZnNldD0iMTAwJSIgc3RvcC1jb2xvcj0iIzJhOWVmMSIvPjwvbGluZWFyR3JhZGllbnQ%2BPC9kZWZzPjxzdHlsZT50ZXh0e2ZvbnQ6NjAwIDQ0cHggLWFwcGxlLXN5c3RlbSxCbGlua01hY1N5c3RlbUZvbnQsJ1NlZ29lIFVJJyxSb2JvdG8sSGVsdmV0aWNhLEFyaWFsLHNhbnMtc2VyaWYsJ0FwcGxlIENvbG9yIEVtb2ppJywnU2Vnb2UgVUkgRW1vamknLCdTZWdvZSBVSSBTeW1ib2wnOy13ZWJraXQtdXNlci1zZWxlY3Q6bm9uZTt1c2VyLXNlbGVjdDpub25lfTwvc3R5bGU%2BPHJlY3Qgd2lkdGg9IjEwMCIgaGVpZ2h0PSIxMDAiIGZpbGw9InVybCgjZykiLz48dGV4dCB0ZXh0LWFuY2hvcj0ibWlkZGxlIiB4PSI1MCIgeT0iNjYiIGZpbGw9IiNmZmYiPlY8L3RleHQ%2BPC9zdmc%2B\"></a>\n</div>\n<div class=\"tgme_page_title\" dir=\"auto\">\n  <span dir=\"auto\">v2ray.run</span>\n</div>\n<div class=\"tgme_page_extra\">1 069 members, 17 online</div>\n<div class=\"tgme_page_description\" dir=\"auto\">网址：<br>netvog.com<br><br>备用：(订阅不了请尝试使用备用地址！)<br>dy.v2ray.club<br>机器人<br><a href=\"https://t.me/v2rayrun_bot\">@v2rayrun_bot</a></div>\n<div class=\"tgme_page_action\">\n  <a class=\"tgme_action_button_new shine\" href=\"tg://resolve?domain=V2rayrun\">View in Telegram</a>\n</div>\n<!-- WEBOGRAM_BTN -->\n\n<div class=\"tgme_page_additional\">\n  If you have <strong>Telegram</strong>, you can view and join <br><strong>v2ray.run</strong> right away.\n</div>\n        </div>\n        \n      </div>\n    </div>\n\n    <div id=\"tgme_frame_cont\"></div>\n\n    <script src=\"//telegram.org/js/tgwallpaper.min.js?3\"></script>\n\n    <script type=\"text/javascript\">\n\nvar protoUrl = \"tg:\\/\\/resolve?domain=V2rayrun\";\nif (false) {\n  var iframeContEl = document.getElementById(\'tgme_frame_cont\') || document.body;\n  var iframeEl = document.createElement(\'iframe\');\n  iframeContEl.appendChild(iframeEl);\n  var pageHidden = false;\n  window.addEventListener(\'pagehide\', function () {\n    pageHidden = true;\n  }, false);\n  window.addEventListener(\'blur\', function () {\n    pageHidden = true;\n  }, false);\n  if (iframeEl !== null) {\n    iframeEl.src = protoUrl;\n  }\n  !false && setTimeout(function() {\n    if (!pageHidden) {\n      window.location = protoUrl;\n    }\n  }, 2000);\n}\nelse if (protoUrl) {\n  setTimeout(function() {\n    window.location = protoUrl;\n  }, 100);\n}\n\nvar tme_bg = document.getElementById(\'tgme_background\');\nif (tme_bg) {\n  TWallpaper.init(tme_bg);\n  TWallpaper.animate(true);\n  window.onfocus = function(){ TWallpaper.update(); };\n}\ndocument.body.classList.remove(\'no_transition\');\n\nfunction toggleTheme(dark) {\n  document.documentElement.classList.toggle(\'theme_dark\', dark);\n  window.Telegram && Telegram.setWidgetOptions({dark: dark});\n}\nif (window.matchMedia) {\n  var darkMedia = window.matchMedia(\'(prefers-color-scheme: dark)\');\n  toggleTheme(darkMedia.matches);\n  darkMedia.addListener(function(e) {\n    toggleTheme(e.matches);\n  });\n}\n\n    \n    </script>\n  \n\n\n</body></html>', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (108, 3782, 106, '如果需要我的付费截图请联系我，', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (109, 3782, 107, '我以前账号是3108000678@qq.com，无法登录一直提示密码错误，但我确定密码是对的，怎么回事？', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (110, 3782, 108, '之前购买服务的账号登入的时候显示邮箱或密码错误切重置密码无法收到邮件。服务器无法更新订阅，显示无效订阅。目前原订阅的服务器均已无法使用。', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (111, 3782, 109, '之前的到期了，今天来看到公告购买了一次性套餐，手机V2rayng更新订阅失败，重置了订阅链接也不行。', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (112, 3782, 110, '订阅链接在v2rayN中提示无效', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (113, 3782, 111, '看到公告后，想咨询下，1是今后是否能正常使用；2是原先购买服务的账号是否能恢复或者按购买记录继承服务，我朋友因为觉得用的还不错，4月刚买了一年的服务，这个有没有一个协调的方案。', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (114, 3782, 112, '付款两笔怎么不生效', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (115, 3782, 113, '节点几乎全部无效。延迟-1ms，尝试刷新缓存重新订阅，无效', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (116, 3782, 114, '尽快维护下啊', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (117, 3782, 115, '已经支付了的 却没给订阅 还剩显示的未支付', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (118, 3782, 115, '修复了', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (119, 3782, 114, '免费节点后续维护', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (120, 3782, 116, '刚才发现账号无法登陆后发现是服务器受到影响，重新购买了3.9之后有240G再次购买一份后还是240G', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (121, 3782, 117, '免费用户是只能看到免费节点吗，还是目前只有一个节点？', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (122, 3782, 118, '实际已完成支付了，订单状态还是未支付', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (123, 3782, 119, '请问本次补偿的1折套餐是5月25日失效吗？届时需要重新购买套餐？', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (124, 3782, 120, '我购买后，根据教程中的订阅地址设置完毕后，在软件中无法更新订阅', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (125, 3782, 121, '请问免费节点和其他付费节点在数量上、速度上、倍率上会有什么区别吗？', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (126, 3782, 122, '我说我的账号怎么怎么输入都不对。。竟然还能重新注册。。。。原来是数据没了。。。', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (127, 3782, 123, '刚下单的所有流量节点均无法使用，请加急处理，体验极差', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (128, 3782, 122, '感谢', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (129, 3782, 119, '不是，是优惠截至时间', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (130, 3782, 119, '套餐是流量用完为止', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (131, 3782, 121, '免费只有美国', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (132, 3782, 118, '已经补了，非常抱歉', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (133, 3782, 111, '1是今后是否能正常使用：已经做了 异地备份，不会再数据丢失了给你的账号补了200g，已经做了 异地备份，', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (134, 3782, 111, '1是今后是否能正常使用：已经做了 异地备份，不会再数据丢失了', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (135, 3782, 111, '2是原先购买服务的账号是...  给你的账号补了200g', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (136, 3782, 107, '非常抱歉，请查看公共', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (137, 3782, 107, '非常抱歉，请查看公告', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (138, 3782, 106, '非常抱歉，请查看公告', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (139, 3782, 89, '非常抱歉，请查看公告', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (140, 3782, 88, '已经更换，下次注意', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (141, 3782, 123, '可能刚好在维护', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (142, 3782, 120, '尝试重新复制订阅地址', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (143, 3782, 117, '目前只有一个节点', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (144, 3782, 113, '修复了，还是不能用吗？', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (145, 3782, 85, '只能买一个，再买回覆盖', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (146, 3782, 82, '给老板加了100g', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (147, 3782, 81, '大哥，购买会覆盖', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (148, 3782, 81, '给大哥加了200g', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (149, 3782, 76, '抱歉，请查看公告', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (150, 3782, 66, '给老板加了100g', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (151, 3782, 124, '以前那个app怎么在苹果店里搜不到了呢', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (152, 3782, 125, '你好，我这边更新的节点有问题，节点名字都叫Alien', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (153, 3782, 126, 'can not access google and speed is so slowly', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (154, 3782, 127, '流量问题', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (155, 3782, 128, '测试延迟100多，点击使用马上变成-1ms', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (156, 3782, 129, 'visit https://chat.openai.com/ error info:Sorry, you have been blocked', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (157, 3782, 130, '请问我付了两个流量套餐\n单号2023050412053238274491725，流量仍然是240g\n为什么不是480g', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (158, 3782, 131, '购买了两次流量套餐，没有叠加，申请退一个单。\n2023050412051126201510211\n2023050412053238274491725', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (159, 3782, 132, '付费节点免费节点都不能用了', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (160, 3782, 133, '买完之后报错 没有v2ray地址', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (161, 3782, 134, '之前订单没了，补偿下单这个根本用不了，没一个可用，不信自己试试\nhttps://by.v2ray.ws/api/v1/client/subscribe?token=277b59c73633221e601802219ace29cb&flag=v2rayn，请赶紧处理，连续几天都无法使用', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (162, 3782, 135, '不行...偶尔有几个几百延迟的，再测试还是寄了，有推断说是时间同步的问题，但是试过了还是8行', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (163, 3782, 136, '目前主用 clash 作为手机和电脑客户端，订阅多个机场，希望提供 clash 可用的订阅链接，方便使用', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (164, 3782, 137, '我之前购买了套餐，现在全不见了，之前的账号9668014@qq.com也没了，现在是新注册的9668104@qq.com', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (165, 3782, 138, '我原来的账号是252768150@qq.com，\n我朋友的账号是sukime0120@shphantom.com\n不知道后续怎么操作', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (166, 3782, 139, 'cs', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (167, 3782, 140, '网络时而连接时而中断', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (168, 3782, 141, '之前一直在用你们家的，4月18日买了11块钱一个月的，开始几天还能用，结果后面一直连不上，电脑也是。这咋办？不会还要重新买吧？我登录账号也登不上了，直接注册了一个新账号，支付宝账单记录和订单号还在的。可以处理一下吗？蟹蟹', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (169, 3782, 142, '我购买的用完即止套餐以及邮箱注册全没了，还有170G的流量，怎么处理......', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (170, 3782, 143, '你们这个补偿套餐的节点安全嘛，会不会用了被叫去喝茶', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (171, 3782, 144, '有没有适用于mac的安装包啊', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (172, 3782, 145, '补偿套餐可以重复购买嘛', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (173, 3782, 146, '因为数据库问题我的套餐消失了', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (174, 3782, 147, '通过之前服务器已经购买了1年。可以处理什么？\n希望确认...\n[月:100G，有效期:~2024.4.9], 以前ID: asoka78@naver.com\n不能上传附件，这样留下邮件...我的手机上盖', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (175, 3782, 148, 'mac 电脑怎么使用', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (176, 3782, 149, 'بالا', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (177, 3782, 150, '买了第二个用完即止，把第一个给覆盖了，可否退款或增加对应流量。', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (178, 3782, 151, 'کانفیگ v2ray می خواهم', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (179, 3782, 152, '你好，我之前的号无法登录了，但是剩余的流量还有很多呢', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (180, 3782, 153, '我的订阅地址无效了，请排查下', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (181, 3782, 154, '多买了一个，第二次买的退一下', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (182, 3782, 155, '使用的是Windows系统，按照说明书进行配置好后，启动成功后，不能访问到外网', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (183, 3782, 156, '支付过去了，但是显示订单没有支付', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (184, 3782, 157, '使用购买的订阅节点，在clash上订阅失败', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (185, 3782, 158, '如题：无法正常的更新出订阅服务器。试过各种方式也去问了，但无济于事。用了其他平台的测试连接是可以正常更新的，用本平台的连接也试过重置，但都只会给出：->Unable to read data from the transport connection: 远程主机强迫关闭了一个现有的连接。. 2->The SSL connection could', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (186, 3782, 159, '能更新订阅，但是不能上网，尝试了很多办法都解决不了', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (187, 3782, 160, '模式使用全局和绕过大陆都试过，系统代理也试过，clash上测试均 超时，v2ray测试真连接延迟均为 -1，这是为什么', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (188, 3782, 161, '节点经常使用不了，间歇性断网', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (189, 3782, 162, '我支付了两次用完即止套餐后仪表盘中未更新相应流量，若该套餐不能重复购买，请退回相应款项；反之请更新剩余流量', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (190, 3782, 163, '我在之前还充钱了 今天发现账户没了  又注册了一个 还买了一点流量 这个不会在关闭把', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (191, 3782, 164, '补偿套餐一次性购买了2次无法叠加到480G', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (192, 3782, 165, '无法下载，速度很慢或者直接下载失败', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (193, 3782, 166, '请问Mac怎么使用，软件哪里下载谢谢', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (194, 3782, 167, 'MacBook下载v2ray U，设置订阅后，无法科学上网', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (195, 3782, 167, 'https://github.com/Qv2ray/Qv2ray', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (196, 3782, 167, 'https://github.com/v2rayA/v2rayA', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (197, 3782, 166, 'https://github.com/v2rayA/v2rayA', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (198, 3782, 166, 'https://github.com/Qv2ray/Qv2ray', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (199, 3782, 168, '全部超时', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (200, 3782, 169, '测速后显示有网速，实际不能科学', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (201, 3782, 170, '我这边订阅了地址，https://by.v2ray.ws/api/v1/client/subscribe?token=7f1b8de1fba7e9ea704337e13abb66f6&flag=v2rayng   但是各个节点都不行', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (202, 3782, 171, '启用代理，但是无法链接到外网', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (203, 3782, 172, '提现方式：USDT\r\n提现账号：test', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (204, 3782, 172, '1', '2023-08-26 16:46:59', '2023-08-26 16:46:59');
INSERT INTO `v2_ticket_message` VALUES (207, 3782, 1, '123123', '2023-08-27 16:37:15', '2023-08-27 16:37:15');
INSERT INTO `v2_ticket_message` VALUES (208, 3782, 1, '1111', '2023-08-27 16:38:58', '2023-08-27 16:38:58');
INSERT INTO `v2_ticket_message` VALUES (209, 3782, 167, 'hao', '2023-09-05 15:38:23', '2023-09-05 15:38:23');
INSERT INTO `v2_ticket_message` VALUES (210, 3782, 172, 'aaa', '2023-09-05 15:38:56', '2023-09-05 15:38:56');
INSERT INTO `v2_ticket_message` VALUES (211, 3782, 172, '', '2023-09-05 15:38:58', '2023-09-05 15:38:58');
INSERT INTO `v2_ticket_message` VALUES (212, 3782, 172, '', '2023-09-05 15:38:58', '2023-09-05 15:38:58');
INSERT INTO `v2_ticket_message` VALUES (213, 3782, 172, '', '2023-09-05 15:38:59', '2023-09-05 15:38:59');
INSERT INTO `v2_ticket_message` VALUES (214, 3782, 172, 'asada', '2023-09-05 15:39:00', '2023-09-05 15:39:00');
INSERT INTO `v2_ticket_message` VALUES (215, 3782, 166, 'q', '2023-09-05 15:50:30', '2023-09-05 15:50:30');
INSERT INTO `v2_ticket_message` VALUES (216, 3782, 171, 'oo', '2023-09-08 17:16:55', '2023-09-08 17:16:55');
INSERT INTO `v2_ticket_message` VALUES (217, 3785, 175, '为什么用不了', '2023-09-08 17:18:15', '2023-09-08 17:18:15');
INSERT INTO `v2_ticket_message` VALUES (218, 3785, 175, 'qq', '2023-09-08 17:18:41', '2023-09-08 17:18:41');
INSERT INTO `v2_ticket_message` VALUES (219, 3782, 175, '我看看', '2023-09-08 17:19:00', '2023-09-08 17:19:00');
INSERT INTO `v2_ticket_message` VALUES (220, 3785, 175, '哦哦', '2023-09-08 17:29:29', '2023-09-08 17:29:29');
INSERT INTO `v2_ticket_message` VALUES (221, 3785, 175, '？？？', '2023-09-08 17:33:02', '2023-09-08 17:33:02');
INSERT INTO `v2_ticket_message` VALUES (222, 3782, 175, '可以了', '2023-09-08 17:34:19', '2023-09-08 17:34:19');
INSERT INTO `v2_ticket_message` VALUES (223, 3785, 175, '11111', '2023-09-08 17:35:55', '2023-09-08 17:35:55');
INSERT INTO `v2_ticket_message` VALUES (224, 3782, 175, '好的', '2023-09-08 17:39:30', '2023-09-08 17:39:30');
INSERT INTO `v2_ticket_message` VALUES (225, 3785, 175, '6666', '2023-09-11 15:00:29', '2023-09-11 15:00:29');
INSERT INTO `v2_ticket_message` VALUES (226, 3785, 172, '777', '2023-09-11 15:00:37', '2023-09-11 15:00:37');
INSERT INTO `v2_ticket_message` VALUES (227, 3785, 175, 'test', '2023-09-11 15:10:11', '2023-09-11 15:10:11');

-- ----------------------------
-- Table structure for v2_user
-- ----------------------------
DROP TABLE IF EXISTS `v2_user`;
CREATE TABLE `v2_user`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `invite_user_id` int NULL DEFAULT NULL COMMENT '邀请id',
  `telegram_id` bigint NULL DEFAULT NULL COMMENT '电报id',
  `user_name` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '账号',
  `password` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '密码',
  `password_algo` char(10) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '加密方式',
  `password_salt` char(10) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '加密盐',
  `balance` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '账户余额',
  `discount` decimal(10, 2) NULL DEFAULT NULL COMMENT '专享折扣',
  `commission_type` tinyint NOT NULL DEFAULT 0 COMMENT '0: system 1: period 2: onetime',
  `commission_rate` int NULL DEFAULT NULL COMMENT '返利比例',
  `commission_balance` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT 'aff余额',
  `commission_code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '邀请码',
  `t` int NOT NULL DEFAULT 0 COMMENT '最后在线时间戳',
  `u` bigint NOT NULL DEFAULT 0 COMMENT '上传',
  `d` bigint NOT NULL DEFAULT 0 COMMENT '下载',
  `transfer_enable` bigint NOT NULL DEFAULT 0 COMMENT '流量',
  `banned` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否禁用',
  `is_admin` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否管理员',
  `is_staff` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否员工',
  `last_login_at` int NULL DEFAULT NULL COMMENT '最后登入时间',
  `last_login_ip` int NULL DEFAULT NULL COMMENT '最后登入ip',
  `uuid` varchar(36) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'uuid',
  `group_id` int NULL DEFAULT NULL COMMENT '权限组',
  `token` char(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'token 订阅用',
  `remarks` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '备注',
  `expired_at` timestamp NULL DEFAULT NULL COMMENT '到期时间',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_name`(`user_name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3789 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of v2_user
-- ----------------------------
INSERT INTO `v2_user` VALUES (3782, 0, 0, 'admina', '', 'MD5', 'e7c78ddd', 0.00, 67.00, 0, 0, 0.00, 'yqmasda', 0, 66, 77, 128849018880, 0, 0, 0, 0, 0, '', 5, '', '', '2023-10-27 18:53:21', '2023-08-21 22:01:46', '2023-09-11 15:24:48');
INSERT INTO `v2_user` VALUES (3785, 0, 0, 'qweqwe', 'a3264b7021966c9e48efa0cf6f223731', 'MD5', '0acfeadb', 0.00, 30.00, 0, 0, 0.00, '2023code', 0, 0, 0, 107374182400, 0, 0, 0, 0, 0, 'b4d5e1e8-13a5-49c9-9ba2-bade1ec2a15d', 3, '', '', '2023-10-12 15:45:37', '2023-08-21 22:32:05', '2023-09-14 12:50:40');
INSERT INTO `v2_user` VALUES (3786, 3785, 0, 'zxczxc', '0cc5478b27f547239a05bb6e70ce47dc', 'MD5', '1e941569', 0.00, 20.00, 0, 0, 0.00, '2bdaf375', 0, 0, 0, 0, 0, 0, 0, 0, 0, 'af1c4204-5e34-4a61-9907-0f5710082cd9', 0, '60c818871dc24e749b562973c5b798d8', '', NULL, '2023-09-12 14:25:55', '2023-09-12 14:25:55');
INSERT INTO `v2_user` VALUES (3787, 0, 0, 'asdasd', '6941268457aa36a9588fc69ca94da646', 'MD5', '3596183d', 0.00, 0.00, 0, 0, 0.00, 'a5c2697b', 0, 0, 0, 0, 0, 0, 0, 0, 0, 'f9a1231b-d166-42e7-8e70-73a9b98ab997', 0, 'bfbee7505d974c46ba5aaf9ba0db0595', '', NULL, '2023-09-12 16:02:47', '2023-09-12 16:02:47');
INSERT INTO `v2_user` VALUES (3788, 3787, 0, '123123', 'f13f89a89302450d5abc4e60f2758c8a', 'MD5', 'f06e6176', 0.00, 0.00, 0, 0, 0.00, 'da576164', 0, 0, 0, 107374182400, 0, 0, 0, 0, 0, 'a99ef502-eec3-484c-8517-28d2ebdc31d8', 5, 'dc7192318fdf4aa8b38a659dde793503', '', '2023-10-12 16:27:42', '2023-09-12 16:04:54', '2023-09-14 12:50:52');

SET FOREIGN_KEY_CHECKS = 1;
