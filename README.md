# bilibili-service
B站API的Golang版本，提供视频源解析，排行获取等常用接口，[BiliBili-Html5](http://www.shiroblue.cn)的一部分

## API地址
===

URL : http://bilibili-service.daoapp.io

基于DaoCloud

## 接口文档
===

* 数据格式: applation/json
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




### 1. 首页内容获取
---
取得主要分类下的前9个热门视频

* URL: /topinfo
* 请求方式: GET

参数:


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

参数:

	page				[int]  页码
	count              	[int]  分页容量


成功返回:

```
{
  "code": 0,
  "list": {
    "0": {
      "aid": "3448965",
      "author": "哔哩哔哩番剧",
      "coins": 9094,
      "comment": 57078,
      "copyright": "Copy",
      "create": "2015-12-25 19:12",
      "credit": 0,
      "description": "#13 福神的传说（完）",
      "duration": "24:14",
      "favorites": 1177,
      "mid": 928123,
      "pic": "http://i2.hdslb.com/320_200/video/83/8389492cdceb85cdade3bc43c17bb699.jpg",
      "play": 677010,
      "review": 4300,
      "subtitle": "",
      "title": "【10月/完结】野良神 ARAGOTO 第二季 13【独家正版】",
      "typeid": 33,
      "typename": "连载动画",
      "video_review": 57078
    },
    ...
    }
}
```

### 3. 视频信息获取
---
各分类下的排行

* URL: /view/ { aid }
 	* aid => av号
* 请求方式: GET

参数:


成功返回:

```
{
  "allow_bp": 1,
  "allow_download": 1,
  "allow_feed": 0,
  "author": "哔哩哔哩番剧",
  "bangumi": {
    "allow_download": "1",
    "bangumi_id": "418",
    "season_id": "2725",
    "title": "K RETURN OF KINGS（第二季）"
  },
  "coins": "4697",
  "created": 1451068200,
  "created_at": "2015-12-26 02:30",
  "credit": "0",
  "description": "#13 Kings ",
  "face": "http://i0.hdslb.com/user/9281/928123/myface.png",
  "favorites": "1250",
  "instant_server": "chat.bilibili.com",
  "list": {
    "0": {
      "cid": 5472643,
      "has_alias": false,
      "page": 1,
      "part": "",
      "type": "vupload",
      "vid": "vupload_5472643"
    }
  },
  "mid": "928123",
  "pages": 1,
  "pic": "http://i2.hdslb.com/video/2d/2db3aaaba87c598e2166a8def95a3dfe.jpg",
  "play": "494592",
  "review": "5554",
  "season_episode": "13",
  "season_id": 2039,
  "season_index": "13",
  "sp_title": "「K」",
  "spid": 5188,
  "src": "c",
  "tag": "BILIBILI独家正版,王的回归,GORA×GOHANDS,K 回归的王权者,K 第二季,K RETURN OF KINGS,K,TV动画,BILIBILI正",
  "tid": 33,
  "title": "【10月/完结】K RETURN OF KINGS 第二季 13【独家正版】",
  "typename": "连载动画",
  "video_review": "49007"
}
```


### 4. 视频地址解析
---
mp4/flv视频源取得，（注意某些老视频没有mp4源）

* URL: (cid => 从视频信息接口取得)
	* /video/ { cid }  (mp4格式)
	* /videoflv/ { cid }  (flv格式)
* 请求方式: GET

参数:

	quailty				[int]  清晰度(1~3，根据视频有不同)


成功返回:

```
{
  "accept": "mp4,hdmp4",
  "size": 184992542,
  "url": "http://xxx.xxxx",
  "backup": [
    "http://xxx.xxxx"
  ]
}
```

### 5. 番剧更新列表
---
目前B站版权二次元/三次元新番

* URL: /bangumi
* 请求方式: GET

参数:

	btype				[int]  二次元新番:2，三次元:3


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

* URL: /spinfo/ { spid }
* 请求方式: GET

参数:


成功返回:

```
{
  "alias": "",
  "alias_spid": 18831,
  "attention": 179094,
  "bangumi_date": "2015-10-01",
  "count": 39,
  "cover": "http://i0.hdslb.com/sp/03/034df09a0c48e977473246d4765f8b4e.jpg",
  "create_at": "2013-12-18 17:27",
  "description": "这里是简介",
  "favourite": 52208,
  "isbangumi": 1,
  "isbangumi_end": 0,
  "lastupdate": 1446345978,
  "lastupdate_at": "2015-11-01 10:46",
  "pubdate": 1387358842,
  "season": [
    {
      "default": false,
      "index_cover": "http://i1.hdslb.com/sp/92/92f6aa4bf437c5ca2e6f357f0240678e.jpg",
      "last_episode": null,
      "season_id": 2050,
      "season_name": "第一季",
      "video_view": 0
    },
    {
      "default": false,
      "index_cover": "http://i2.hdslb.com/sp/cb/cb5ccf5e39045fd83d18ec31851f36c6.jpg",
      "last_episode": null,
      "season_id": 2051,
      "season_name": "第二季",
      "video_view": 1383367
    }
  ],
  "season_id": 2051,
  "spid": 56747,
  "title": "请问您今天要来点兔子吗？？（第二季）",
  "video_view": 11496158,
  "view": 736531
}
```

### 7. 专题视频获取
---
取得专题下的所有视频

* URL: /spvideos/ { spid }
* 请求方式: GET

参数:

	bangumi				[int]  取得番剧视频:1，其他视频:0


成功返回:

```
{
  "code": 0,
  "count": 17,
  "list": [
    {
      "aid": 3147596,
      "cid": 4954881,
      "click": 132329,
      "cover": "http://i1.hdslb.com/video/64/64a3c078c640faad3862208e367303f4.jpg",
      "from": "vupload",
      "page": 0,
      "title": "【Kyle钢琴】十月番其实是一首歌（传颂K野良点兔物语高达黑杰克FFF超人假面舞会）"
    },
    {
      "aid": 1186141,
      "cid": 1756783,
      "click": 14407,
      "cover": "http://i2.hdslb.com/u_f/df9a71b1232afc1a3a299b7fddfd532e.jpeg",
      "from": "sina",
      "page": 0,
      "title": "【東方】请问您今天要来点优昙华吗？"
    },
   	...
  ],
  "results": 17,
  "spid": 56747
}
```


### 8. 全站搜索
---

* URL: /search
* 请求方式: GET

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




