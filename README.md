# bilibili-go

BiliBili API written in Go



## Progress

* Rank
	* ```SortRank``` (order by danmu/comment/hot)
* Video
	* ```VideoInfo```
	* ```VideoLin``` (mp4/flv)
* User
	* ```UserInfo```
	* ```UserVideos```
* Special
	* ```SpecialInfo```
	* ```SpecialVideos```
* Bangumi
	* ```BangumiList```
	* ```BangumiRecommend```
* Others
	* ```Search```(search user/video/bangumi)



## Install 

```
go get github.com/WhiteBlue/bilibili-go
```

## Usage

```
import "github.com/whiteblue/bilibili-go/client"

c := client.NewClient("APPKEY", "SECRET")
back, err := c.Bangumi.GetWeekList("2")

if err != nil {
    log.Error(err)
    return
}

log.Info(result)

```






## 接口地址
===

URL : ```http://bilibili-service.daoapp.io```

基于DaoCloud免费容器

## 接口文档
===

* 数据格式: ```application/json```
* 请求方式: ```post/get```


#### 基本状态码(HTTP)约定:
	
	500: 服务器错误（API返回异常）
	404: 请求的资源不可得
	200: 成功
	400: 参数异常


#### 错误返回格式:
	
	{
    "code": "PARAM_ERROR",
    "message": "request valitdate error"
	}


#### 分类(sort)约定:
        '1' => '动画',
        '3' => '音乐',
        '4' => '游戏',
        '5' => '娱乐',
        '11' => '电视剧'
        '13' => '番剧',
        '23' => '电影',
        '36' => '科技',
        '119' => '鬼畜',
        '129' => '舞蹈',


### 1. 首页内容获取
---
取得主要分类下的前10个热门视频

* URL: ```/allrank```
* 请求方式: GET
* 示例:
	*  ```curl -X GET -H "Content-Type: application/json" -H "Cache-Control: no-cache" "http://bilibili-service.daoapp.io/allrank"```

* 参数: 无


成功返回:

```
[
  {
    "sort_name": "动画",
    "videos": [
      {
        "aid": "5697545",
        "mid": 28965086,
        "copyright": "Original",
        "typeid": 27,
        "typename": "综合",
        "title": "【妈的智障】笑到胃疼的动画片段（第七期）片尾洗澡＆足控福利",
        "subtitle": "",
        "play": 1398924,
        "review": 3435,
        "video_review": 24134,
        "favorites": 40908,
        "author": "噗汪汪",
        "description": "看你们谁还敢说我短！嗯？不知道怎么抽奖的请看上期视频。番名按照顺序分别是：银魂、超元气三姐妹、男子高中生的日常、妄想学生会、潜行吧奈亚子、濑户的花嫁、我们大家的河合庄和悠哉日常大王，片尾是银魂OAD。谢谢支持！",
        "create": "2016-08-07 19:12",
        "pic": "http://i2.hdslb.com/bfs/archive/7b2aced4ba0e5924b12fcf5a0ad36d8c2728b73d.jpg_320x200.jpg",
        "credit": 0,
        "coins": 6264,
        "duration": "24:28",
        "comment": 24134,
        "badgepay": false
      },
      ...
]
```
	
### 2. 分类排行获取
---
各分类下的排行

* URL: ```/sort/{tid}```
	* ```tid```为分类id(例如“动画”=>13)
* 请求方式: GET
* 示例: 	
	* ```curl -X GET -H "Content-Type: application/json" -H "Cache-Control: no-cache" "http://bilibili-service.daoapp.io/sort/13"```

* 参数:

	* ```page```: 页码
	* ```count```: 分页容量
	* ```order```: 排序方式(new,hot)


成功返回:

```
{
  "name": "番剧",
  "list": {
    "0": {
      "aid": "5698105",
      "mid": 21453565,
      "copyright": "Copy",
      "typeid": 33,
      "typename": "连载动画",
      "title": "【4月】Re：从零开始的异世界生活 19",
      "subtitle": "",
      "play": 2493341,
      "review": 42346,
      "video_review": 191673,
      "favorites": 2157,
      "author": "TV-TOKYO",
      "description": "#19 白鲸攻略战",
      "create": "2016-08-07 19:46",
      "pic": "http://i1.hdslb.com/bfs/archive/a0656101763a68a4bcb3fe603496037c253e106d.jpg_320x200.jpg",
      "credit": 0,
      "coins": 15908,
      "duration": "24:35",
      "comment": 191673,
      "badgepay": false
    },
    "1":{...},
    ...
    }
}
```

### 3. 视频信息获取
---
各分类下的排行

* URL: ```/view/{aid}```
 	* aid => av号
* 请求方式: GET
* 示例:
	* ```curl -X GET -H "Content-Type: application/json" -H "Cache-Control: no-cache" "http://bilibili-service.daoapp.io/view/5698105"```

* 参数: 无


成功返回:

```
{
  "tid": 33,
  "typename": "连载动画",
  "arctype": "Copy",
  "play": "2493735",
  "review": "42400",
  "video_review": "191632",
  "favorites": "2159",
  "title": "【4月】Re：从零开始的异世界生活 19",
  "description": "#19 白鲸攻略战",
  "tag": "TV动画,BILIBILI正版,RE：从零开始的异世界生活,从零开始的异世界生活",
  "pic": "http://i0.hdslb.com/bfs/archive/a0656101763a68a4bcb3fe603496037c253e106d.jpg",
  "author": "TV-TOKYO",
  "mid": "21453565",
  "face": "http://i0.hdslb.com/bfs/face/69ef6861067d6ef637b7c73b77d71c3414996745.jpg",
  "pages": 1,
  "created_at": "2016-08-08 01:05",
  "coins": "15911",
  "list": {
    "0": {
      "page": 1,
      "type": "vupload",
      "part": "ReZERO_19",
      "cid": 9253164,
      "vid": 0
    }
  }
}
```


### 4. 视频地址解析
---
mp4/flv视频源取得，（注意某些老视频没有mp4源）

* URL: ```/video/{cid}```
* 请求方式: GET
* 示例: 
	* ```curl -X GET -H "Cache-Control: no-cache" "http://bilibili-service.daoapp.io/video/9253164?quality=2"```
* 参数:
	* ```quailty```:清晰度(1~2，根据视频有不同)
	* ```type```: 格式(mp4/flv)


成功返回:

```
{
  "format": "hdmp4",
  "timelength": 1476160,
  "accept_format": "mp4,hdmp4",
  "accept_quality": [
    2,
    1
  ],
  "durl": [
    {
      "length": 1476160,
      "size": 206950377,
      "url": "http://cn-tj1-cu.acgvideo.com/vg123/3/2e/9253164-1-hd.mp4?expires=1471021200&ssig=1NyYrtPpFZmm4zHVClIHzA&oi=2067479167&rate=0",
      "backup_url": [
        "http://cn-sddz2-cu.acgvideo.com/vg1/3/6d/9253164-1-hd.mp4?expires=1471021200&ssig=actJTZUGft5yN6fSQGL9Kw&oi=2067479167&rate=0",
        "http://cn-sdjn-cu-v-01.acgvideo.com/vg6/e/d5/9253164-1-hd.mp4?expires=1471021200&ssig=7I1sNpH6szF5CtNP4bwfXA&oi=2067479167&rate=0"
      ]
    }
  ]
}
```


### 5. 番剧更新列表
---
目前B站版权二次元新番

* URL: ```/bangumi```
* 请求方式: GET
* 示例:
	* ```curl -X GET -H "Cache-Control: no-cache" "http://bilibili-service.daoapp.io/bangumi"``` 

* 参数: 无


成功返回:

```
{
  "0": [
    {
      "area": "日本",
      "arealimit": 0,
      "attention": 471585,
      "bangumi_id": 1070,
      "bgmcount": "24",
      "brief": null,
      "cover": "http://i1.hdslb.com/u_user/e6835e74ce9d6f63bc44d4f42dfc82e4.jpg",
      "danmaku_count": 439806,
      "favorites": 471585,
      "is_finish": 0,
      "lastupdate": 1451152800,
      "lastupdate_at": "2015-12-27 02:00:00",
      "new": true,
      "play_count": 9646746,
      "pub_time": "",
      "season_id": 2760,
      "spid": 56749,
      "square_cover": "http://i0.hdslb.com/sp/1e/1e21c6a6e17f5419eb1e10fadc53e6eb.jpg",
      "title": "终结的炽天使 第二季",
      "url": "/bangumi/i/2760/",
      "weekday": 0
    },
    ...
 ],
 ...
}
```


### 6. 专题信息查看
---
例如番剧专题

* URL: 
    * ```/spinfo/{spid}```
* 请求方式: GET
* 示例:
	* ```curl -X GET -H "Cache-Control: no-cache" "http://bilibili-service.daoapp.io/spinfo/56749"```

* 参数: 无


成功返回:

```
{
    "alias": "终わりのセラフ",
    "alias_spid": 41465,
    "attention": 367438,
    "bangumi_date": "2015-10-01",
    "count": 40,
    "cover": "http://i0.hdslb.com/sp/1e/1e21c6a6e17f5419eb1e10fadc53e6eb.jpg",
    "create_at": "2014-12-12 21:19",
    "description": "电视动画《终结的炽天使》改编自日本轻小说家镜贵也原作、漫画家山本大和作画的同名漫画。\r\n2014年8月28日，发表了《终结的炽天使》电视动画化的决定。\r\n2014年12月20日，在日本千叶县幕张展览馆开幕的“Jump Festa 2015”会场上，宣布电视动画《终结的炽天使》会被分割成两个季度播出。\r\n第1期的播送时间为2015年4月4日－6月20日。\r\n第2期则是同年的10月至12月。",
    "favourite": 172114,
    "isbangumi": 1,
    "isbangumi_end": 1,
    "lastupdate": 1450364026,
    "lastupdate_at": "2015-12-17 22:53",
    "pubdate": 1418390386,
    "season": [
        {
            "default": false,
            "index_cover": "http://i2.hdslb.com/sp/5c/5c7dbad52d522b6a5cbc8fe383ed92fe.jpg",
            "last_episode": null,
            "season_id": 2052,
            "season_name": "第一季",
            "video_view": 826087
        },
        {
            "default": false,
            "index_cover": "http://i1.hdslb.com/sp/87/87d650e53a8d50302a369365063c45a4.jpg",
            "last_episode": null,
            "season_id": 2053,
            "season_name": "第二季",
            "video_view": 4314354
        }
    ],
    "season_id": 2053,
    "spid": 56749,
    "title": "终结的炽天使 第二季",
    "video_view": 28723627,
    "view": 1350787
}
```

### 7. 专题视频获取
---
取得专题下的所有视频

* URL: ```/spvideos/{spid}```
* 请求方式: GET
* 示例: 
	* ```curl -X GET -H "Cache-Control: no-cache" "http://bilibili-service.daoapp.io/spvideos/56749?bangumi=0"```

* 参数:
	* ```bangumi```: 取得番剧视频:1，其他视频:0


成功返回:

```
{
    "code": 0,
    "count": 17,
    "list": [
        {
            "aid": 2330598,
            "cid": 3638258,
            "click": 498347,
            "cover": "http://i2.hdslb.com/video/51/512fc7fce5bb04a42fe116eb5500af20.jpg",
            "from": "vupload",
            "page": 0,
            "title": "「终结的炽天使」OP ED专辑"
        },
        {
            "aid": 2425245,
            "cid": 3796297,
            "click": 101788,
            "cover": "http://i0.hdslb.com/video/4c/4cf6be151a858c200e4aa5ac07c2ccd1.jpg",
            "from": "vupload",
            "page": 0,
            "title": "让我们的炽天使燃起来吧Answer is near【MAD】"
        },
				...
    ],
    "results": 17,
    "spid": 56749
}
```


### 8. 全站搜索
---

* URL: ```/search```
* 请求方式: POST
* 示例:
	* ```curl -X POST -H "Cache-Control: no-cache" -H "Content-Type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW" -F "content=Fate" "http://bilibili-service.daoapp.io/search"```

* 参数:
	* ```content```: 搜索内容
	* ```page```: 页码
	* ```count```: 分页大小

成功返回:

```
{
  "page": 1,
  "pagesize": 20,
  "pageinfo": {
    "bangumi": {
      "total": 11,
      "numResults": 11,
      "pages": 4
    },
    "movie": {
      "total": 2,
      "numResults": 2,
      "pages": 1
    },
    "pgc": {
      "total": 2,
      "numResults": 2,
      "pages": 1
    },
    "special": {
      "total": 27,
      "numResults": 24,
      "pages": 8
    },
    "topic": {
      "total": 10,
      "numResults": 10,
      "pages": 4
    },
    "tvplay": {
      "total": 0,
      "numResults": 0,
      "pages": 1
    },
    "upuser": {
      "total": 86,
      "numResults": 86,
      "pages": 29
    },
    "video": {
      "total": 24742,
      "numResults": 999,
      "pages": 50
    }
  },
  "result": {
    "video": [
      {
        "aid": "4912937",
        "mid": 777536,
        "copyright": "",
        "typeid": 0,
        "typename": "综合",
        "title": "【灵魂配音】10分钟演完fate stay night UBW",
        "subtitle": "",
        "play": 1121120,
        "review": 6922,
        "video_review": 18656,
        "favorites": 32058,
        "author": "LexBurner",
        "description": "自制 试水作，感谢新月冰冰配以及小鹤儿的帮忙，希望以后参与的人能越来越多，做的越来越好玩，这次还不是很到位，以后继续努力啦\r\nlex的零食铺：http://lexzhils.taobao.com\r\nlex的新浪微博：http://weibo.com/lexburner\r\n新月冰冰视频空间：http://space.bilibili.com/3295/#!/index\r\n小鹤儿视频空间：http://space.bilibili.com/6719190/#!/index",
        "create": "",
        "pic": "http://i0.hdslb.com/bfs/archive/7edf866255ae2d9a8f31c176c0873769d6451243.jpg_320x200.jpg",
        "credit": 0,
        "coins": 0,
        "duration": "10:33",
        "comment": 0,
        "badgepay": false
      },
      ...
      ]
  }   
}
```

### 9. 用户信息
---
* URL: ```/user/{mid}```
* 请求方式: GET
* 示例:
	* ```curl -X GET -H "Cache-Control: no-cache" "http://bilibili-service.daoapp.io/user/116683"```
	
* 参数: 无

成功返回:

```
{
  "mid": 116683,
  "name": "=咬人猫=",
  "sex": "女",
  "rank": 10000,
  "face": "http://i1.hdslb.com/bfs/face/8fad84a4470f3d894d8f0dc95555ab8f2cb10a83.jpg",
  "coins": 63211.8,
  "regtime": 1301718879,
  "birthday": "0000-00-00",
  "place": "",
  "description": "bilibili 知名舞见",
  "attentions": [
    179628,
    271126,
    622863,
    5055,
    433715,
    6870383,
    4350178,
    8084905
  ],
  "fans": 627933,
  "friend": 8,
  "attention": 8,
  "sign": "面瘫女仆酱~小粗腿~事业线什么的！！吐槽你就输了！喵~"
}
```

### 10. 用户视频
---
* URL: ```/uservideos/{mid}```
* 请求方式: GET
* 示例:
	* ```curl -X GET -H "Cache-Control: no-cache" "http://bilibili-service.daoapp.io/uservideos/116683"```
	
* 参数: 
	* ```page```: 页码
	* ```count```: 分页容量


成功返回:

```
{
  "vlist": [
    {
      "aid": 5682645,
      "copyright": "Original",
      "typeid": 154,
      "title": "【咬人猫】落花情❤ o(*≧▽≦)ツ",
      "subtitle": "",
      "play": 752435,
      "review": 6402,
      "video_review": 9549,
      "favorites": 25734,
      "mid": 116683,
      "author": "=咬人猫=",
      "description": "也许是有史以来最忐忑的一次投稿，这次尝试的风格对我来说挑战很大，完全没有这类舞蹈的基础，所以当时这支舞挖坑填了一段时间因为实在是能力有限，被我搁置了很久～后来因为确实是太喜欢这种感觉的舞蹈了，又重新鼓起劲去学，也尝试了不同的服装版本录制了很多次，非常想完成这一支舞，虽然还有很多地方不够好，也希望大家多多包涵，谢谢大家的支持和等待～\n服装：七秀萝莉定国套\n舞蹈歌曲：七朵组合（一代）的作品《落花情》。",
      "created": "2016-08-06 20:20:45",
      "pic": "http://i0.hdslb.com/bfs/archive/a96729ac9e0a14544d1fcdb1471d23cf7ac1e61b.jpg",
      "comment": 9549,
      "length": "03:46"
    },
    ...
   ]
}
```



### 11. 新番推荐
---

* URL: ```/bangumiindex```
* 请求方式: GET
* 示例:
	* ```curl -X GET -H "Cache-Control: no-cache" "http://bilibili-service.daoapp.io/bangumiindex"```

参数: 无


成功返回:

```
{
  "banners": [
    {
      "title": "美术社大有问题",
      "link": "http://www.bilibili.com/bangumi/i/5043/",
      "img": "http://i0.hdslb.com/bfs/archive/c69c196266bfe6fb52e05d5752f4687e501d83e8.jpg",
      "simg": "",
      "aid": 0,
      "type": "link",
      "platform": 0,
      "pid": 0
    },
    {
      "title": "魔法战争",
      "link": "http://www.bilibili.com/bangumi/i/4367/",
      "img": "http://i0.hdslb.com/bfs/archive/3324a9fc99275ef22d0364bd3221e009020d9567.jpg",
      "simg": "",
      "aid": 0,
      "type": "link",
      "platform": 0,
      "pid": 0
    },
    {
      "title": "月歌",
      "link": "http://bangumi.bilibili.com/anime/5038",
      "img": "http://i0.hdslb.com/bfs/archive/dfbc9098218c7deffed53dd9c14619d88ac8f180.jpg",
      "simg": "",
      "aid": 0,
      "type": "link",
      "platform": 0,
      "pid": 0
    },
    {
      "title": "灵能百分百",
      "link": "http://bangumi.bilibili.com/anime/5058",
      "img": "http://i0.hdslb.com/bfs/archive/35e322a660aa83ada2a6ae94923c27510f40fd26.jpg",
      "simg": "",
      "aid": 0,
      "type": "link",
      "platform": 0,
      "pid": 0
    }
  ],
  "recommends": [
    {
      "aid": "5753187",
      "title": "【蒼氏甜品坊】初恋怪兽「…这是我的广播、要怎么办呢？」",
      "subtitle": "",
      "play": 631,
      "review": 0,
      "video_review": 82,
      "favorites": 139,
      "mid": 628114,
      "author": "祈妹",
      "description": "TV动画初恋怪兽的应援广播，配信日为每周五，主持人为苍井翔太。\n一个与动画内容相比在污和hentai的层面毫不逊色的广播节目。\n这是甜品坊第一次做广播，由于人手、经验和水平的不足不能保证翻译完全准确，有错误的地方欢迎指正~",
      "create": "2016-08-10 22:24",
      "pic": "http://i0.hdslb.com/bfs/archive/8770dd5682edb0170b5d3bfe08b8763aff9f7e4d.jpg_320x200.jpg",
      "coins": 40,
      "duration": "128:36"
    },
    ...
  ]
}
```


