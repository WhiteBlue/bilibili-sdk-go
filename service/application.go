package service

import (
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/log"
	"github.com/go-playground/log/handlers/console"
	"github.com/whiteblue/bilibili-go/client"
	"time"
)

const (
	INDEX_CACHE        = "index"
	ALL_RANK_CACHE     = "all_rank"
	BANGUMI_CACHE      = "bangumi"
	BANGUMI_LIST_CACHE = "bangumi_list"
	SORT_TOP_CACHE     = "sort-"
	LIVE_INDEX_CACHE   = "live_index"
)

var (
	ProdLevels = []log.Level{
		log.InfoLevel,
		log.NoticeLevel,
		log.WarnLevel,
		log.ErrorLevel,
		log.PanicLevel,
		log.AlertLevel,
		log.FatalLevel,
	}
)

type BiliBiliApplication struct {
	Router *gin.Engine
	Corn   *CornService
	Conf   *Config
	Client *client.BCli
	Cache  *CacheManager
}

func NewApplication(configFile string) (*BiliBiliApplication, error) {
	conf, err := ReadConfigFromFile(configFile)

	if err != nil {
		return nil, err
	}

	cLog := console.New()

	if conf.Debug {
		log.RegisterHandler(cLog, log.AllLevels...)
		gin.SetMode(gin.DebugMode)
	} else {
		log.RegisterHandler(cLog, ProdLevels...)
		gin.SetMode(gin.ReleaseMode)
	}

	log.Info("conform config file")

	r := gin.New()

	//use gzip
	r.Use(gin.Recovery())
	r.Use(gzip.Gzip(gzip.BestCompression))

	//corn service
	corn := NewCornService()

	//bilibili client
	cli := client.NewClient(conf.Appkey, conf.Secret)

	cache := NewCacheManager()

	app := &BiliBiliApplication{Router: r, Corn: corn, Conf: conf, Client: cli, Cache: cache}

	ConformRoute(app)

	log.Info("conform route")

	conformTask(app)
	corn.Start()

	log.Info("conform task")

	log.Info("init complete, start listen...")

	return app, nil
}

func conformTask(app *BiliBiliApplication) {
	app.Corn.RegisterTask(&IndexInfoTask{CornTask: CornTask{Name: "index_info", Duration: 2 * time.Hour}, app: app})
	app.Corn.RegisterTask(&BangumiInfoTask{CornTask: CornTask{Name: "bangumi_info", Duration: 6 * time.Hour}, app: app})
	app.Corn.RegisterTask(&BangumiListTask{CornTask: CornTask{Name: "bangumi_list", Duration: 6 * time.Hour}, app: app})
	app.Corn.RegisterTask(&TopRankTask{CornTask: CornTask{Name: "top_rank", Duration: 2 * time.Hour}, app: app})
	app.Corn.RegisterTask(&LiveIndexTask{CornTask: CornTask{Name: "alive_index", Duration: 2 * time.Hour}, app: app})
}
