# Dump of table Account
# ------------------------------------------------------------

DROP TABLE IF EXISTS `Account`;

CREATE TABLE `Account` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `accountType` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '账号类型',
  `accountKey` varchar(64) NOT NULL COMMENT '账号标识，不变',
  `username` varchar(50) NOT NULL COMMENT '账号名',
  `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `mobilePhone` varchar(64) DEFAULT NULL COMMENT '手机号码，加密',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码，签名处理',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `accountKey` (`accountKey`) USING BTREE,
  UNIQUE KEY `username` (`accountType`,`username`) USING BTREE,
  UNIQUE KEY `email` (`accountType`,`email`) USING BTREE,
  UNIQUE KEY `mobilePhone` (`accountType`,`mobilePhone`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

LOCK TABLES `Account` WRITE;
/*!40000 ALTER TABLE `Account` DISABLE KEYS */;

INSERT INTO `Account` (`id`, `accountType`, `accountKey`, `username`, `email`, `mobilePhone`, `password`, `createTime`, `updateTime`)
VALUES
	(3,X'61646D696E','5efe837e54621af52bf8ae2b','dengzw1',NULL,NULL,'00a1f187721c63501356bf791e69382c','2020-07-03 09:01:48','2020-07-03 09:01:48'),
	(1002,X'61646D696E','5efdb43c35f1190858e19b34','dengzw','','','00a1f187721c63501356bf791e69382c','2020-07-02 18:17:31','2020-07-02 18:17:31');

/*!40000 ALTER TABLE `Account` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table Account_Role
# ------------------------------------------------------------

DROP TABLE IF EXISTS `Account_Role`;

CREATE TABLE `Account_Role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `accountKey` varchar(30) NOT NULL COMMENT '账户标识',
  `roleID` int(11) NOT NULL COMMENT '角色ID',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `accountRole` (`accountKey`,`roleID`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

LOCK TABLES `Account_Role` WRITE;
/*!40000 ALTER TABLE `Account_Role` DISABLE KEYS */;

INSERT INTO `Account_Role` (`id`, `accountKey`, `roleID`, `createTime`, `updateTime`)
VALUES
	(1,'5efdb43c35f1190858e19b34',1,'2020-07-03 15:10:30','2020-07-03 23:03:33');

/*!40000 ALTER TABLE `Account_Role` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table AccountSystem
# ------------------------------------------------------------

DROP TABLE IF EXISTS `AccountSystem`;

CREATE TABLE `AccountSystem` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `accountType` varchar(30) NOT NULL COMMENT '账户类型，唯一',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `accountType` (`accountType`) USING BTREE COMMENT '账户类型唯一'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

LOCK TABLES `AccountSystem` WRITE;
/*!40000 ALTER TABLE `AccountSystem` DISABLE KEYS */;

INSERT INTO `AccountSystem` (`id`, `accountType`, `createTime`, `updateTime`)
VALUES
	(1,'admin','2020-07-02 16:32:22','2020-07-02 16:32:22');

/*!40000 ALTER TABLE `AccountSystem` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table App
# ------------------------------------------------------------

DROP TABLE IF EXISTS `App`;

CREATE TABLE `App` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `appKey` varchar(30) NOT NULL COMMENT '应用标识，唯一',
  `name` varchar(30) NOT NULL COMMENT '应用名称',
  `accountType` varchar(30) NOT NULL COMMENT '账号类型',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `appKey` (`appKey`) USING BTREE COMMENT '应用标识唯一'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

LOCK TABLES `App` WRITE;
/*!40000 ALTER TABLE `App` DISABLE KEYS */;

INSERT INTO `App` (`id`, `appKey`, `name`, `accountType`, `createTime`, `updateTime`)
VALUES
	(1,'Auth','权限服务','admin','2020-07-03 09:03:33','2020-07-03 09:03:33'),
	(2,'EAP','EAP','admin','2020-07-03 12:32:56','2020-07-03 12:32:56');

/*!40000 ALTER TABLE `App` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table Resource
# ------------------------------------------------------------

DROP TABLE IF EXISTS `Resource`;

CREATE TABLE `Resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `appKey` varchar(30) NOT NULL COMMENT '应用标识',
  `name` varchar(30) NOT NULL COMMENT '资源名称',
  `type` varchar(30) NOT NULL COMMENT '资源类型：api',
  `feature` varchar(100) NOT NULL COMMENT '特征值：如：uri',
  `method` varchar(20) DEFAULT NULL COMMENT 'restful api请求类型',
  `pid` int(11) NOT NULL DEFAULT '0' COMMENT '父级资源ID',
  `sequence` int(11) NOT NULL DEFAULT '0' COMMENT '资源排序，值越大越靠前',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `appResource` (`appKey`,`feature`,`type`) USING BTREE COMMENT '同应用同类型特征值唯一'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

LOCK TABLES `Resource` WRITE;
/*!40000 ALTER TABLE `Resource` DISABLE KEYS */;

INSERT INTO `Resource` (`id`, `appKey`, `name`, `type`, `feature`, `method`, `pid`, `sequence`, `createTime`, `updateTime`)
VALUES
	(1,'EAP','查询','api','/query','GET',0,0,'2020-07-03 15:11:27','2020-07-03 15:11:27');

/*!40000 ALTER TABLE `Resource` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table Role
# ------------------------------------------------------------

DROP TABLE IF EXISTS `Role`;

CREATE TABLE `Role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `appKey` varchar(30) NOT NULL COMMENT '应用标识',
  `name` varchar(30) NOT NULL COMMENT '角色名称',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `appRole` (`appKey`,`name`) USING BTREE COMMENT '应用角色唯一'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

LOCK TABLES `Role` WRITE;
/*!40000 ALTER TABLE `Role` DISABLE KEYS */;

INSERT INTO `Role` (`id`, `appKey`, `name`, `createTime`, `updateTime`)
VALUES
	(1,'EAP','超管','2020-07-03 15:10:17','2020-07-03 15:10:17');

/*!40000 ALTER TABLE `Role` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table Role_Resource
# ------------------------------------------------------------

DROP TABLE IF EXISTS `Role_Resource`;

CREATE TABLE `Role_Resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `roleID` int(11) NOT NULL COMMENT '角色ID',
  `resourceID` int(11) NOT NULL COMMENT '资源ID',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `roleResource` (`roleID`,`resourceID`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

LOCK TABLES `Role_Resource` WRITE;
/*!40000 ALTER TABLE `Role_Resource` DISABLE KEYS */;

INSERT INTO `Role_Resource` (`id`, `roleID`, `resourceID`, `createTime`, `updateTime`)
VALUES
	(1,1,1,'2020-07-03 22:02:15','2020-07-03 22:02:15');

/*!40000 ALTER TABLE `Role_Resource` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
