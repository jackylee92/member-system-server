package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackylee92/rgo/core/rgrequest"
	"net/http"
)

//func SystemHandle(c *gin.Context) {
//	this := rgrequest.Get(c)
//	data := getSystemListPageData()
//	log.Println("this.Response.View(\"home/system_list.html\", data)")
//	this.Response.View("home/system_list.html", data)
//	return
//}

func MakeOrderHandle(c *gin.Context) {
	this := rgrequest.Get(c)
	system := this.Ctx.Query("sys")
	viewPath, exist := getViewPath(system)
	if !exist {
		this.Ctx.Redirect(http.StatusFound, "/public/system")
		return
	}
	logoUrl := getLogoPath(system)
	productList := getProductList(system)
	data := map[string]interface{}{
		"title":       "订单列表",
		"disclaimer":  "声明：本程序仅供学习交流,正规等合法使用。利用本程序从事任何违反中国法律法规的行为产生法律纠纷自行承担后果,由此引起的一切后果与（程序作者）无关！",
		"logo":        logoUrl,
		"productList": productList,
		"systemName":  getSystemName(system),
	}
	this.Response.View(viewPath, data)
	return
}

func getProductList(sys string) (data []string) {
	allData := map[string][]string{
		"tx": {
			"腾讯视频会员周卡",
			"腾讯视频会员周卡",
			"腾讯视频会员周卡",
			"腾讯视频会员周卡",
			"腾讯视频会员月卡",
			"腾讯视频会员季卡",
			"腾讯视频会员年卡",
			"==================",
			"腾讯视频超级会员月卡",
			"腾讯视频超级会员月卡",
			"腾讯视频超级会员季卡",
			"腾讯视频超级会员年卡",
			"==================",
			"腾讯网游加速器3天",
			"腾讯网游加速器7卡",
			"腾讯网游加速器月卡",
			"腾讯网游加速器季卡",
			"腾讯网游加速器年卡",
			"==================",
			"腾讯体育会员月卡",
			"腾讯体育会员季卡",
			"腾讯体育会员年卡",
			"==================",
			"QQ会员月卡",
			"QQ会员季卡",
			"QQ会员年卡",
			"==================",
			"QQ超级会员月卡",

			"QQ超级会员月卡",
			"QQ超级会员月卡",
			"QQ超级会员季卡",
			"QQ超级会员年卡",
			"==================",
			"QQ大会员月卡",
			"QQ大会员季卡",
			"QQ大会员年卡",
			"==================",
			"QQ黄钻月卡",
			"QQ黄钻季卡",
			"QQ黄钻年卡",
			"==================",
			"QQ豪华黄钻月卡",
			"QQ豪华黄钻月卡",
			"QQ豪华黄钻季卡",
			"QQ豪华黄钻年卡",
			"==================",
			"QQ绿钻月卡",
			"QQ绿钻月卡",
			"QQ绿钻季卡",
			"QQ绿钻年卡",
			"==================",
			"QQ豪华绿钻月卡",
			"QQ豪华绿钻月卡",
			"QQ豪华绿钻季卡",
			"QQ豪华绿钻年卡",
			"==================",
			"QQ蓝钻月卡",
			"QQ蓝钻季卡",
			"QQ蓝钻年卡",
			"==================",
			"QQ红钻月卡",
			"QQ红钻季卡",
			"QQ红钻年卡",
			"==================",
			"QQ紫钻月卡",
			"QQ紫钻季卡",
			"QQ紫钻年卡",
			"==================",
			"QQ黑钻月卡",
			"QQ黑钻季卡",
			"QQ黑钻年卡",
			"==================",
			"QQ音乐包月卡",
			"QQ音乐包季卡",
			"QQ音乐包年卡",
			"==================",
			"QQ微云会员月卡",
			"QQ微云会员季卡",
			"QQ微云会员年卡",
			"==================",
			"QQ微云超级会员月卡",
			"QQ微云超级会员季卡",
			"QQ微云超级会员年卡",
			"==================",
		},
		"iqiyi": {
			"爱奇艺黄金会员天卡",
			"爱奇艺黄金会员周卡",
			"爱奇艺黄金会员月卡",
			"爱奇艺黄金会员季卡",
			"爱奇艺黄金会员半年卡",
			"爱奇艺黄金会员年卡",
			"爱奇艺白金会员年卡",
			"爱奇艺星钻会员月卡",
			"爱奇艺星钻会员季卡",
			"爱奇艺星钻会员半年卡",
			"爱奇艺星钻会员年卡",
		},
		"youku": {
			"优酷黄金会员周卡",
			"优酷黄金会员月卡",
			"优酷黄金会员年卡",
		},
		"mangguo": {
			"芒果TV视频会员周卡",
			"芒果TV视频会员月卡",
			"芒果TV视频会员季卡",
			"芒果TV视频会员年卡",
		},
		"letv": {
			"乐视视频会员周卡",
			"乐视视频会员月卡",
			"乐视视频会员年卡",
		},
		"sohu": {
			"搜狐视频会员周卡",
			"搜狐视频会员月卡",
			"搜狐视频会员年卡",
		},
		"bili": {
			"哔哩哔哩会员三天卡",
			"哔哩哔哩会员周卡",
			"哔哩哔哩会员月卡",
			"哔哩哔哩会员季卡",
			"哔哩哔哩会员年卡",
		},
		"meituan": {
			"美团外卖5元红包",
			"美团外卖10元红包",
			"美团会员周卡",
			"美团会员月卡",
			"美团会员季卡",
			"美团会员年卡",
		},
		"eleme": {
			"饿了么会员周卡",
			"饿了么会员月卡",
			"饿了么会员季卡",
			"饿了么会员年卡",
		},
		"wangyiyun": {
			"网易云音乐黑胶会员周卡",
			"网易云音乐黑胶会员月卡",
			"网易云音乐黑胶会员季卡",
			"网易云音乐黑胶会员年卡",
		},
		"ximalaya": {
			"喜马拉雅巅峰会员周卡",
			"喜马拉雅巅峰会员月卡",
			"喜马拉雅巅峰会员季卡",
			"喜马拉雅巅峰会员年卡",
		},
		"kugou": {
			"酷狗音乐会员周卡",
			"酷狗音乐会员月卡",
			"酷狗音乐会员季卡",
			"酷狗音乐会员年卡",
		},
		"baidupan": {
			"百度网盘超级会员周卡",
			"百度网盘超级会员月卡",
			"百度网盘超级会员季卡",
			"百度网盘超级会员年卡",
			"百度网盘普通会员周卡",
			"百度网盘普通会员月卡",
			"百度网盘普通会员季卡",
			"百度网盘普通会员年卡",
		},
		"ucpan": {
			"UC网盘会员周卡",
			"UC网盘会员月卡",
			"UC网盘会员季卡",
			"UC网盘会员年卡",
		},
		"baiduwenku": {
			"百度文库会员周卡",
			"百度文库会员月卡",
			"百度文库会员季卡",
			"百度文库会员年卡",
		},
		"wps": {
			"WPS会员周卡",
			"WPS会员月卡",
			"WPS会员季卡",
			"WPS会员年卡",
		},
		"tianyancha": {
			"天眼查会员周卡",
			"天眼查会员月卡",
			"天眼查会员季卡",
			"天眼查会员年卡",
		},
		"zhihu": {
			"知乎会员周卡",
			"知乎会员月卡",
			"知乎会员季卡",
			"知乎会员年卡",
		},
		"xunlei": {
			"迅雷会员周卡",
			"迅雷会员月卡",
			"迅雷会员季卡",
			"迅雷会员年卡",
		},
		"keep": {
			"Keep会员周卡",
			"Keep会员月卡",
			"Keep会员季卡",
			"Keep会员年卡",
		},
		"kuaikan": {
			"快看会员周卡",
			"快看会员月卡",
			"快看会员季卡",
			"快看会员年卡",
		},
		"kuake": {
			"夸克会员周卡",
			"夸克会员月卡",
			"夸克会员季卡",
			"夸克会员年卡",
		},
		"zuoyebang": {
			"作业帮会员周卡",
			"作业帮会员月卡",
			"作业帮会员季卡",
			"作业帮会员年卡",
		},
		"migu": {
			"咪咕会员周卡",
			"咪咕会员月卡",
			"咪咕会员季卡",
			"咪咕会员年卡",
		},
	}
	return allData[sys]
}

func getLogoPath(system string) (data string) {
	allData := map[string]string{
		"tx":         "logo_tx.png",
		"iqiyi":      "logo_iqiyi.png",
		"youku":      "logo_youku.png",
		"mangguo":    "logo_mangguo.png",
		"letv":       "logo_letv.png",
		"sohu":       "logo_sohu.png",
		"bili":       "logo_bili.png",
		"meituan":    "logo_meituan.png",
		"eleme":      "logo_eleme.png",
		"wangyiyun":  "logo_wangyiyun.png",
		"ximalaya":   "logo_ximalaya.png",
		"kugou":      "logo_kugou.png",
		"baidupan":   "logo_baidupan.png",
		"ucpan":      "logo_ucpan.png",
		"baiduwenku": "logo_baiduwenku.png",
		"wps":        "logo_wps.png",
		"tianyancha": "logo_tianyancha.png",
		"zhihu":      "logo_zhihu.png",
		"xunlei":     "logo_xunlei.png",
		"keep":       "logo_keep.png",
		"kuaikan":    "logo_kuaikan.png",
		"kuake":      "logo_kuake.png",
		"zuoyebang":  "logo_zuoyebang.png",
		"migu":       "logo_migu.png",
	}
	return allData[system]
}

func getViewPath(system string) (data string, res bool) {
	allData := map[string]string{
		"tx":         "home/tx.html",
		"iqiyi":      "home/no_tx.html",
		"youku":      "home/no_tx.html",
		"mangguo":    "home/no_tx.html",
		"letv":       "home/no_tx.html",
		"sohu":       "home/no_tx.html",
		"bili":       "home/no_tx.html",
		"meituan":    "home/no_tx.html",
		"eleme":      "home/no_tx.html",
		"wangyiyun":  "home/no_tx.html",
		"ximalaya":   "home/no_tx.html",
		"kugou":      "home/no_tx.html",
		"baidupan":   "home/no_tx.html",
		"ucpan":      "home/no_tx.html",
		"baiduwenku": "home/no_tx.html",
		"wps":        "home/no_tx.html",
		"tianyancha": "home/no_tx.html",
		"zhihu":      "home/no_tx.html",
		"xunlei":     "home/no_tx.html",
		"keep":       "home/no_tx.html",
		"kuaikan":    "home/no_tx.html",
		"kuake":      "home/no_tx.html",
		"zuoyebang":  "home/no_tx.html",
		"migu":       "home/no_tx.html",
	}
	data, ok := allData[system]
	return data, ok
}

func getSystemName(system string) (data string) {
	allData := getAllSystemName()
	return allData[system]
}
func getAllSystemName() (allData map[string]string) {
	allData = map[string]string{
		"tx":         "腾讯",
		"iqiyi":      "爱奇艺",
		"youku":      "优酷",
		"mangguo":    "芒果",
		"letv":       "乐视",
		"sohu":       "搜狐",
		"bili":       "B站",
		"meituan":    "美团",
		"eleme":      "饿了么",
		"wangyiyun":  "网易云",
		"ximalaya":   "喜马拉雅",
		"baidupan":   "百度网盘",
		"ucpan":      "UC网盘",
		"baiduwenku": "百度文库",
		"wps":        "WPS",
		"tianyancha": "天眼查",
		"zhihu":      "知乎",
		"xunlei":     "迅雷",
		"keep":       "Keep",
		"kuaikan":    "快看",
		"kuake":      "夸克",
		"zuoyebang":  "作业帮",
		"migu":       "咪咕",
	}
	return allData
}

func getSystemListPageData() (data map[string]interface{}) {
	data = make(map[string]interface{})
	list := []map[string]interface{}{
		{
			"title": "",
			"url":   "tx",
			"image": "/assets/fictitious_order/img/logo/tx.png",
		},
		{
			"title": "",
			"url":   "iqiyi",
			"image": "/assets/fictitious_order/img/logo/iqiyi.png",
		},
		{
			"title": "",
			"url":   "youku",
			"image": "/assets/fictitious_order/img/logo/youku.png",
		},
		{
			"title": "",
			"url":   "mangguo",
			"image": "/assets/fictitious_order/img/logo/mangguo.png",
		},
		{
			"title": "",
			"url":   "letv",
			"image": "/assets/fictitious_order/img/logo/letv.png",
		},
		{
			"title": "",
			"url":   "sohu",
			"image": "/assets/fictitious_order/img/logo/sohu.png",
		},
		{
			"title": "",
			"url":   "bili",
			"image": "/assets/fictitious_order/img/logo/bili.png",
		},
		{
			"title": "",
			"url":   "meituane",
			"image": "/assets/fictitious_order/img/logo/meituane.png",
		},
		{
			"title": "",
			"url":   "eleme",
			"image": "/assets/fictitious_order/img/logo/eleme.png",
		},
		{
			"title": "",
			"url":   "wangyiyun",
			"image": "/assets/fictitious_order/img/logo/wangyiyun.png",
		},
		{
			"title": "",
			"url":   "ximalaya",
			"image": "/assets/fictitious_order/img/logo/ximalaya.png",
		},
		{
			"title": "",
			"url":   "kugou",
			"image": "/assets/fictitious_order/img/logo/kugou.png",
		},
		{
			"title": "",
			"url":   "baidupan",
			"image": "/assets/fictitious_order/img/logo/baidupan.png",
		},
		{
			"title": "",
			"url":   "ucpan",
			"image": "/assets/fictitious_order/img/logo/ucpan.png",
		},
		{
			"title": "",
			"url":   "baiduwenku",
			"image": "/assets/fictitious_order/img/logo/baiduwenku.png",
		},
		{
			"title": "",
			"url":   "wps",
			"image": "/assets/fictitious_order/img/logo/wps.png",
		},
		{
			"title": "",
			"url":   "tianyancha",
			"image": "/assets/fictitious_order/img/logo/tianyancha.png",
		},
		{
			"title": "",
			"url":   "zhihu",
			"image": "/assets/fictitious_order/img/logo/zhihu.png",
		},
		{
			"title": "",
			"url":   "xunlei",
			"image": "/assets/fictitious_order/img/logo/xunlei.png",
		},
		{
			"title": "",
			"url":   "keep",
			"image": "/assets/fictitious_order/img/logo/keep.png",
		},
		{
			"title": "",
			"url":   "kuaikan",
			"image": "/assets/fictitious_order/img/logo/kuaikan.png",
		},
		{
			"title": "",
			"url":   "kuake",
			"image": "/assets/fictitious_order/img/logo/kuake.png",
		},
		{
			"title": "",
			"url":   "zuoyebang",
			"image": "/assets/fictitious_order/img/logo/zuoyebang.png",
		},
		{
			"title": "",
			"url":   "migu",
			"image": "/assets/fictitious_order/img/logo/migu.png",
		},
	}
	data["list"] = list
	return data
}
