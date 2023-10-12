
```sql
show tables;
```

## 搭建项目的骨架

1. 建库建表

新建发号器表
```sql
CREATE TABLE `sequence` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `stub` varchar(1) NOT NULL,
  `timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_uniq_stub` (`stub`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT = '序号表';
```

新建长链接短链接映射表：
```sql
CREATE TABLE `short_url_map` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_by` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '创建者',
    `is_del` tinyint UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否删除：0正常1删除',
    
    `lurl` varchar(2048) DEFAULT NULL COMMENT '长链接',
    `md5` char(32) DEFAULT NULL COMMENT '长链接MD5',
    `surl` varchar(11) DEFAULT NULL COMMENT '短链接',
    PRIMARY KEY (`id`),
    INDEX(`is_del`),
    UNIQUE(`md5`),
    UNIQUE(`surl`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COMMENT = '长短链映射表';
```

```sql
desc shorturl;
    
desc short_url_map;
```

2. 根据api文件生成go代码

```bash
goctl api go -api shortener.api -dir .
```

3. 根据数据表生成model层代码
```bash
goctl model mysql datasource -url="root:571400yst@tcp(127.0.0.1:3306)/shorturl" -table="short_url_map"  -dir="./model" -c

goctl model mysql datasource -url="root:571400yst@tcp(127.0.0.1:3306)/shorturl" -table="sequence"  -dir="./model" -c
```

### 发号器
- 基于Mysql主键
- 基于redis Key的自增
- SnowFlake


### bloom 过滤器
- "github.com/zeromicro/go-zero/core/bloom"
- 使用go-zero自带布隆过滤器

在布隆过滤器中查询是否存在
```go
exists, err := l.svcCtx.Filter.Exists([]byte(req.ShortUrl))
if err != nil || !exists {
    logx.Errorw("Bloom Filter failed", logx.LogField{Value: err.Error(), Key: "err"})
    return nil, Err404
}
```
存入布隆过滤器中
```go
if err := l.svcCtx.Filter.Add([]byte(sUrl)); err != nil {
    logx.Errorw("Filter.Add() failed", logx.LogField{Key: "err", Value: err.Error()})
    return nil, err
}
```

### 敏感词汇屏蔽

```go
if _, ok := l.svcCtx.BlackList[sUrl]; ok {
    continue
}
```

### EFK的使用

待完....