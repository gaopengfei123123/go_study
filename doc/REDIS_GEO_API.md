> 最近需要用到LBS距离测绘方面的东西，用于搜索附近的人，
在网上搜了一下有一个比较通用的方案，就是采用 GEOHASH 算法

[相关的参考资料](http://blog.jobbole.com/80633/)

这套算法在 mongoDB 当中已经集成了，但是公司的业务中没有用到 mongodb 而是用了redis，
本来打算自己造个轮子去实现，但是万幸翻看了一下redis的文档，它也在 3.* 版本后提供了GEO相关的api

[相关文章](http://cristian.regolo.cc/2015/07/07/introducing-the-geo-api-in-redis.html)



首先是添加命令 `GEOADD`
```
GEOADD key long lat name [long lat name1,long lat name2,long lat name3...]
```
1. key  `GEOSET`的名字
2. long 地点的经度
3. lat  地点的纬度
4. name `GEOSET`内的值的名字
可以批量输入

举个栗子
```
GEOADD chaoyangqu 116.43620200729366 39.916880160714435 "yabaolu" 116.4356870231628 39.908560377800676 "jianguomen" 116.4345336732864 39.924466658329585 "chaoyangmen" 116.4335788068771 39.921372916981106 "yinhesoho" 116.43564410781856 39.92024564137184 "zhongliang"
(integer) 5
```
返回值为1 说明是新添，返回0 说明是更新 错误的话就报错了

查看当前插入的内容
```
zrange chaoyangqu 0 -1
1) "jianguomen"
2) "yabaolu"
3) "zhongliang"
4) "yinhesoho"
5) "chaoyangmen"
```
从这条命令也可以看出来 geoset 是属于有序集合


命令 `GEORADIUS`
```
GEORADIUS key long lat radius unit [WITHDIST] [WITHHASH] [WITHCOORD] [ASC|DESC] [COUNT count]
```
1. key 集合的名字
2. long 指定的经度
3. lat  指定的纬度
4. radius 半径范围
5. unit 单位 可以说 m km ft(英尺) mi(英里)
```
GEORADIUS chaoyangqu 116.435182767868 39.91811857809279 300 m
1) "zhongliang"
2) "yabaolu"
```
带上参数可以显示更多位置信息，通常业务需要显示下距离和进行排序，添加 `WITHDIST` `ASC`
```
GEORADIUS chaoyangqu 116.435182767868 39.91811857809279 300 m WITHDIST ASC
1) 1) "yabaolu"
   2) "162.8454"
2) 1) "zhongliang"
   2) "239.7758"
```

命令`GEORADIUSBYMEMBER`
```
GEORADIUSBYMEMBER key member radius unit [WITHDIST] [WITHHASH] [WITHCOORD] [ASC|DESC] [COUNT count]
```
1. key 集合的名字
2. long 指定的经度
3. lat  指定的纬度
4. radius 半径范围
5. unit 单位 可以说 m km ft(英尺) mi(英里)

这个和 `GEORADIUS` 的区别就是 前者是指定坐标，后者是可以直接从集合中指定一个name做计算
栗子
```
GEORADIUSBYMEMBER chaoyangqu yinhesoho 300 m WITHDIST ASC
GEORADIUSBYMEMBER chaoyangqu yinhesoho 300 m WITHDIST ASC
1) 1) "yinhesoho"
   2) "0.0000"
2) 1) "zhongliang"
   2) "216.2817"
```
排在第一位的肯定是自己啊，而且需要确定你要搜寻的名字已经存在集合中，否则会报错

命令`GEOPOS`
```
GEOPOS key elem1 elem2 ... elemN
```
列出集合中指定name的坐标
栗子：
```
GEOPOS chaoyangqu yabaolu yinhesoho xxx
1) 1) "116.43620342016220093"
   2) "39.9168813345921123"
2) 1) "116.43358021974563599"
   2) "39.92137286048646416"
3) (nil)
```
说明如果 name 不存在set中就会返回nil


命令`GEODIST`
```
GEODIST key elem1 elem2 [unit]
1. unit 单位可以说m lm ft mi
```
只能是两个元素之间的啊，不然三个元素就是语法错误，两个元素中有一个不存在则返回 `nil`
栗子
```
GEODIST chaoyangqu yabaolu yinhesoho m
"547.4030"
```

下面的几个命令做个了解，平时不常用到

* GEOHASH key elem1 elem2 ... elemN   //获取指定name的hash值

* GEOENCODE long lat [radius unit]    //坐标转geohash

* GEODECODE hash    //geohash转坐标
