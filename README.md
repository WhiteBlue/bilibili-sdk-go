# BiliBiliGo

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
go get -u github.com/WhiteBlue/bilibiligo
```

## Usage

```
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

URL : http://bilibili-service.daoapp.io

基于DaoCloud免费容器

## 接口文档
===

* 数据格式: application/json
* 请求方式: post/get


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
取得主要分类下的前9个热门视频

* URL: /allrank
* 请求方式: GET
* 示例: GET http://bilibili-service.daoapp.io/allrank

参数: 无


成功返回:

```
{
  "动画": [
    {
      "aid": "3448677",
      "author": "名作之壁吧",
      "coins": 2036,
      "comment": 17736,
      "copyright": "Original",
      "create": "2015-12-25 18:33",
      "credit": 0,
      "description": "这里是简介",
      "duration": "53:00",
      "favorites": 4954,
      "mid": 2859372,
      "pic": "http://i1.hdslb.com/320_200/video/42/42bf663ddae99db7c22d7a4c093ab25b.jpg",
      "play": 97149,
      "review": 2528,
      "subtitle": "",
      "title": "2015年日本TV动画销量排行榜",
      "typeid": 27,
      "typename": "综合",
      "video_review": 17736
    },
    ...
  ],
  "娱乐":[
  ...
  ]
}
```
	
### 2. 分类排行获取
---
各分类下的排行

* URL: /sort/ { tid }
	* tid为分类id(例如“动画”=>13)
* 请求方式: GET
* 示例: GET http://bilibili-service.daoapp.io/sort/13

参数:

	page				[int]  页码
	count              	[int]  分页容量
	order               [str]  排序方式(new,hot)


成功返回:

```
{
    "code": 0,
    "list": {
        "0": {
            "aid": "3580782",
            "author": "哔哩哔哩番剧",
            "badgepay": false,
            "coins": 5581,
            "comment": 46962,
            "copyright": "Copy",
            "create": "2016-01-14 19:09",
            "credit": 0,
            "description": "#02 枫的时间 ",
            "duration": "23:17",
            "favorites": 4618,
            "mid": 928123,
            "pic": "http://i1.hdslb.com/320_200/video/9a/9a886536a807d717fe9bfe9bd811ed5b.jpg",
            "play": 978653,
            "review": 4629,
            "subtitle": "",
            "title": "【1月】极速老师 第二季 02【独家正版】",
            "typeid": 33,
            "typename": "连载动画",
            "video_review": 46962
        },
        "1": {
            "aid": "3580514",
            "author": "TV-TOKYO",
            "badgepay": false,
            "coins": 4575,
            "comment": 19437,
            "copyright": "Copy",
            "create": "2016-01-14 18:31",
            "credit": 0,
            "description": "#664 自来也忍法帐 鸣人豪杰物语「叛逃」【应官方要求，本片字幕全网统一为优土译制版本。】",
            "duration": "23:24",
            "favorites": 4051,
            "mid": 21453565,
            "pic": "http://i0.hdslb.com/320_200/video/c8/c89497a44e4d05e812c2c61066d42484.jpg",
            "play": 555915,
            "review": 6915,
            "subtitle": "",
            "title": "【1月】火影忍者 疾风传 664",
            "typeid": 33,
            "typename": "连载动画",
            "video_review": 19437
        },
        ...
        "num": "665"
    },
    "name": "番剧",
    "num": 665,
    "pages": 34,
    "results": 665
}

```

### 3. 视频信息获取
---
各分类下的排行

* URL: /view/ { aid }
 	* aid => av号
* 请求方式: GET
* 示例: GET http://bilibili-service.daoapp.io/view/3580782

参数: 无


成功返回:

```
{
    "allow_bp": 1,
    "allow_download": 1,
    "allow_feed": 0,
    "author": "哔哩哔哩番剧",
    "bangumi": {
        "allow_download": "1",
        "bangumi_id": "1033",
        "season_id": "3271",
        "title": "极速老师"
    },
    "coins": "5584",
    "created": 1452804600,
    "created_at": "2016-01-15 04:50",
    "credit": "0",
    "description": "#02 枫的时间 ",
    "face": "http://i0.hdslb.com/user/9281/928123/myface.png",
    "favorites": "4621",
    "instant_server": "chat.bilibili.com",
    "list": {
        "0": {
            "cid": 5710510,
            "has_alias": false,
            "page": 1,
            "part": "",
            "type": "vupload",
            "vid": "vupload_5710510"
        }
    },
    "mid": "928123",
    "pages": 1,
    "pic": "http://i0.hdslb.com/video/9a/9a886536a807d717fe9bfe9bd811ed5b.jpg",
    "play": "979055",
    "review": "4630",
    "spid": null,
    "src": "c",
    "tag": "TV动画,洲崎绫,田中美海,冈本信彦,伊藤静,BILIBILI正版,黄老师,福山润,极速老师 第二季",
    "tid": 33,
    "title": "【1月】极速老师 第二季 02【独家正版】",
    "typename": "连载动画",
    "video_review": "46974"
}
```


### 4. 视频地址解析
---
mp4/flv视频源取得，（注意某些老视频没有mp4源）

* URL: (cid => 从视频信息接口取得)
	* /video/ { cid }  (mp4格式)
* 请求方式: GET
* 示例: 
	* GET http://bilibili-service.daoapp.io/video/3580782?quality=2

参数:

	quailty				[int]  清晰度(1~2，根据视频有不同)
	type				[int]  0:flv,1:hdmp4,2:mp4


成功返回:

mp4
```
{
    "accept": "mp4,hdmp4",
    "backup": [
        "http://cc.acgvideo.com/201601191329/77fcfd7934552b0e2cf974e84d7d92ba/b/81/3580782-1.mp4",
        "http://ws.acgvideo.com/2/e8/3580782-1hd.mp4?wsTime=1453224424&wsSecret2=bdd28d4da66692521875ce8be36a2807&oi=2021932405&appkey=4ebafd7c4951b366&or=987503882"
    ],
    "url": "http://ws.acgvideo.com/2/e8/3580782-1.mp4?wsTime=1453224424&wsSecret2=cc4ff05eabce23fb761d568caf3c85db&oi=2021932405&appkey=4ebafd7c4951b366&or=987503882"
}
```


### 5. 番剧更新列表
---
目前B站版权二次元新番

* URL: /bangumi
* 请求方式: GET
* 示例: GET http://bilibili-service.daoapp.io/bangumi

参数:


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
    * /spinfo/ { spid } 
    * /spinfo   (当没有spid时可用"title"参数代替请求)
* 请求方式: GET
* 示例:
	* { spid } GET http://bilibili-service.daoapp.io/spinfo/56749
	* title GET http://bilibili-service.daoapp.io/spinfo/title=[极速老师]

参数:
    
    title       [str]      专题标题(可选)


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

* URL: /spvideos/ { spid }
* 请求方式: GET
* 示例: GET http://bilibili-service.daoapp.io/spvideos/56749


参数:

	bangumi				[int]  取得番剧视频:1，其他视频:0


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

* URL: /search
* 请求方式: POST

参数:

	content				[string]  搜索内容
	page				[int]  页码
	count				[int]  分页大小

成功返回:

```
{
  "code": 0,
  "cost": {
    "bgm_format": "0.000025",
    "bgm_parse": "0.000003",
    "bgm_read_db": "0.000007",
    "bgm_score": "0.000028",
    "bgm_sengine": "0.000580",
    "bgm_sort": "0.000002",
    "chk_params": "0.000411",
    "locate_ip": "0.000001",
    "prepare": "0.000771",
    "rcache": "0.000387",
    "special_parse": "0.000666",
    "special_read_db": "0.000071",
    "special_score": "0.000002",
    "special_sengine": "0.000840",
    "special_sort": "0.000002",
    "split words": "0.000605",
    "state": "0.000001",
    "suggest words": "0.000006",
    "timer": "all",
    "total": "0.136029",
    "video_parse": "0.000012",
    "video_read_db": "0.007094",
    "video_score": "0.001366",
    "video_sengine": "0.121594",
    "video_sort": "0.000214",
    "wcache": "0.001710"
  },
  "numPages": 9,
  "numResults": 167,
  "page": 9,
  "pagesize": 20,
  "result": [
    {
      "arcurl": "/bangumi/i/2747/",
      "attention": 246827,
      "author": "",
      "bgmcount": "12",
      "click": 5453715,
      "count": 13,
      "description": "我是简介",
      "favourite": 246827,
      "id": 2747,
      "is_bangumi": 1,
      "is_bangumi_end": 0,
      "ischeck": 1,
      "lastupdate": 1443891600,
      "mid": "",
      "pic": "http://i2.hdslb.com/u_user/d2fc60535637c670e6467329dcc02acf.jpg",
      "postdate": 1443891600,
      "pubdate": 1443891600,
      "season_id": 2747,
      "spcount": 0,
      "spid": 5002747,
      "tag": "",
      "thumb": "http://i2.hdslb.com/u_user/d2fc60535637c670e6467329dcc02acf.jpg",
      "title": "终物语",
      "type": "special",
      "typename": "专题",
      "typeurl": "http://www.bilibili.com/"
    },
    {
      "aid": "3303938",
      "arcrank": "0",
      "arcurl": "http://www.bilibili.com/video/av3303938/",
      "author": "你看不到我哟吼",
      "description": "自制 自截 终物语第九话 阿良良木历的笑声",
      "favorites": 13,
      "id": 3303938,
      "mid": 650824,
      "pic": "http://i0.hdslb.com/video/c7/c7c09d3522bebae499bfbd19ea9c1089.jpg",
      "play": 1552,
      "pubdate": 1448804881,
      "review": 1,
      "tag": "神谷浩史,物语系列,终物语",
      "title": "【终物语】垃圾君的笑声",
      "type": "video",
      "typename": "综合",
      "video_review": 11
    },
    {
      "aid": "3309295",
      "arcrank": "0",
      "arcurl": "http://www.bilibili.com/video/av3309295/",
      "author": "矢口晶",
      "description": "我是简介",
      "favorites": 2,
      "id": 3309295,
      "mid": 8159773,
      "pic": "http://i2.hdslb.com/video/75/758a0213215e3bf79bdeee5150a91a68.jpg",
      "play": 481,
      "pubdate": 1448893778,
      "review": 5,
      "tag": "新番,OP替换,终物语,开心就好",
      "title": "【终物语】画面一下子就骚气起来了,不是么?",
      "type": "video",
      "typename": "综合",
      "video_review": 9
    },
    ...
  ],
  "seid": "11235589006635531144",
  "sengine": {
    "count": "0",
    "usage": "0.000"
  },
  "suggest_keyword": "",
  "total": 167
}
```


### 9. 新番推荐
---

* URL: /bangumiindex
* 请求方式: GET
* 示例: GET http://bilibili-service.daoapp.io/bangumiindex

参数: 无


成功返回:

```
{
    "code": "0",
    "result": {
        "banners": [
            {
                "aid": 579138,
                "img": "http://i1.hdslb.com/u_user/f4170a4a141da0ec2fee8f1ef03c2978.jpg",
                "link": "http://www.bilibili.com/video/av579138/?br",
                "pid": 1,
                "platform": 3,
                "simg": "",
                "title": "备长炭",
                "type": "video"
            },
						...
        ],
        "catalogys": [
            {
                "hots": [
                    {
                        "aid": "3594386",
                        "author": "哔哩哔哩番剧",
                        "badgepay": false,
                        "coins": 870,
                        "create": "2016-01-16 18:15",
                        "description": "#02 炎之巨人",
                        "duration": "23:40",
                        "favorites": 1108,
                        "mid": 928123,
                        "pic": "http://i0.hdslb.com/320_200/video/53/5334088c93f4907ddbad826fc1122bfd.jpg",
                        "play": 320032,
                        "pts": 338006,
                        "review": 2403,
                        "subtitle": "",
                        "title": "【1月】舞武器舞乱伎 02",
                        "video_review": 13834
                    },
										...
                ],
                "tid": 33
            },
            {
                "hots": [
                    {
                        "aid": "3608429",
                        "author": "哔哩哔哩番剧",
                        "badgepay": false,
                        "coins": 21,
                        "create": "2016-01-18 16:05",
                        "description": "“机器人少女”的企划是东映动画公司在2009年时的一项企划，《机器人少女z》则是把位于“机器人动画金字塔顶端”的“魔神系列三部曲”《魔神Z》《大魔神 GREAT MAZINGER》《UFO魔神古兰戴萨》中的机体进行美少女化，并组成最强最凶猛的“机器人少女Z”简称“Z组合”。",
                        "duration": "83:00",
                        "favorites": 428,
                        "mid": 928123,
                        "pic": "http://i2.hdslb.com/320_200/video/1a/1ac92893de90b57316ee1bd224c0e485.jpg",
                        "play": 9770,
                        "pts": 14959,
                        "review": 98,
                        "subtitle": "",
                        "title": "【合集】机器人少女Z",
                        "video_review": 244
                    },
										...
                ],
                "tid": 32
            },
            {
                "hots": [
                    {
                        "aid": "3613184",
                        "author": "腾讯动漫",
                        "badgepay": false,
                        "coins": 31,
                        "create": "2016-01-19 09:53",
                        "description": "自制 中国惊奇先生第47集：小狐狸脱险，王小二直捣贼窝！\r\n小伙伴们~惊奇先生改为每周二更新了哟~周周都更~约吗？！",
                        "duration": "12:16",
                        "favorites": 31,
                        "mid": 732364,
                        "pic": "http://i1.hdslb.com/320_200/video/5f/5f32f083c6d7e932b72e13fbfb4a54d6.jpg",
                        "play": 10044,
                        "pts": 11123,
                        "review": 81,
                        "subtitle": "",
                        "title": "中国惊奇先生 47",
                        "video_review": 569
                    },
										...
                ],
                "tid": 153
            },
            {
                "hots": [
                    {
                        "aid": "3589718",
                        "author": "东映动画",
                        "badgepay": false,
                        "coins": 55,
                        "create": "2016-01-16 00:23",
                        "description": "https://www.youtube.com/watch?v=TGFWXJ0F4JM 片源：東映アニメーション公式YouTubeチャンネル／原文·翻译：思言／校对：翼尔，みなもと楓／时间轴·后期：Geemon，野龙／制作统筹：AGUMON／制作监督：みなもと楓／字幕：驯兽师联盟／制作：驯兽师联盟／数码兽大冒险tri. 第2章「决意」2016年3月12日剧场上映／数码兽大冒险tri. 第3章「告白」2016年夏剧场上映\r\n",
                        "duration": "1:27",
                        "favorites": 248,
                        "mid": 3923048,
                        "pic": "http://i1.hdslb.com/320_200/video/fc/fcc480f8369fb279c4527aa7fe6b7c79.jpg",
                        "play": 31576,
                        "pts": 32593,
                        "review": 82,
                        "subtitle": "",
                        "title": "【剧场版】数码兽大冒险tri.第2章 决意 PV1【驯兽师联盟】",
                        "video_review": 154
                    },
										...
                ],
                "tid": 51
            },
            {
                "hots": [
                    {
                        "aid": "3612711",
                        "author": "南條小鹿",
                        "badgepay": false,
                        "coins": 986,
                        "create": "2016-01-19 05:41",
                        "description": "二次创作 LoveLive！Nico生课外活动系列的限定复活！！！\r\n\r\n在找视频里的素材时翻到不少以前的东西觉得好怀念 \r\n想起自己当初入μ's坑做的第一档字幕就是课外活动的生放送 \r\n一转眼也过去那么久了\r\n（姨妈痛的缘故躺了两天 所以迟发了两天_(:зゝ∠)_）",
                        "duration": "82:00",
                        "favorites": 2042,
                        "mid": 554122,
                        "pic": "http://i1.hdslb.com/320_200/video/66/66de11a229cb7db3bf95cb25cfb1d83e.jpg",
                        "play": 19015,
                        "pts": 47728,
                        "review": 418,
                        "subtitle": "",
                        "title": "[中字] LoveLive! μ's课外活动 一夜限定的鸟果海姬复活 感谢特番",
                        "video_review": 5636
                    },
										...
                ],
                "tid": 152
            }
        ],
        "recommends": [
            {
                "aid": "3612711",
                "author": "南條小鹿",
                "badgepay": false,
                "coins": 986,
                "create": "2016-01-19 05:41",
                "description": "二次创作 LoveLive！Nico生课外活动系列的限定复活！！！\r\n\r\n在找视频里的素材时翻到不少以前的东西觉得好怀念 \r\n想起自己当初入μ's坑做的第一档字幕就是课外活动的生放送 \r\n一转眼也过去那么久了\r\n（姨妈痛的缘故躺了两天 所以迟发了两天_(:зゝ∠)_）",
                "duration": "82:00",
                "favorites": 2042,
                "mid": 554122,
                "pic": "http://i1.hdslb.com/320_200/video/66/66de11a229cb7db3bf95cb25cfb1d83e.jpg",
                "play": 19015,
                "review": 418,
                "subtitle": "",
                "title": "[中字] LoveLive! μ's课外活动 一夜限定的鸟果海姬复活 感谢特番",
                "video_review": 5636
            },
						...
        ]
    }
}
```



### 9. APP首页推荐
---
> 有大量直播内容,还没想到有什么卵用


* URL: /appindex
* 请求方式: GET
* 示例: GET http://bilibili-service.daoapp.io/appindex

参数: 无


成功返回:略