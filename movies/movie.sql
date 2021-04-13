-- ----------------------------
-- Table structure for base_member
-- ----------------------------
DROP TABLE IF EXISTS `movies`;
CREATE TABLE `movies` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '电影id',
  `title` varchar(100) NOT NULL COMMENT '标题',
  `description` varchar(250) NOT NULL COMMENT '描述',
  `url` varchar(100) NOT NULL COMMENT '链接',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `title_index` (`title`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='中台用户表';

