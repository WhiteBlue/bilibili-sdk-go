# bilibili-sdk-go

BiliBili Open API & SDK written in Go



## Open API

Docs: [docs](docs/api_doc.md)

* api.bilibilih5.club
* api.prprpr.me/bilibili ( support by [DIYgod](https://github.com/DIYgod))

Deploy:

* ```docker build -t bilibili-go```
* ```docekr run -d -p 80:8080 bilibili-go```


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

## Related Projects

* [BiliBili-Html5](http://bilibilih5.club)



## License

MIT License

Copyright (c) 2016 Castaway Consulting LLC

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.


