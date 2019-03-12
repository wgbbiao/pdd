package pdd

type MethodType string

const (
	// 多多客
	DDK_ThemeListGet              MethodType = "pdd.ddk.theme.list.get"
	DDK_ThemeGoodsSearch          MethodType = "pdd.ddk.theme.goods.search"
	DDK_RPPromUrlGenerate         MethodType = "pdd.ddk.rp.prom.url.generate" // 生成红包推广链接
	DDK_OrderListIncrementGet     MethodType = "pdd.ddk.order.list.increment.get"
	DDK_ColorOrderIncrementGet    MethodType = "pdd.ddk.color.order.increment.get"
	DDK_GoodsDetail               MethodType = "pdd.ddk.goods.detail"
	DDK_GoodsSearch               MethodType = "pdd.ddk.goods.search"
	DDK_GoodsPidQuery             MethodType = "pdd.ddk.goods.pid.query"
	DDK_GoodsPidGenerate          MethodType = "pdd.ddk.goods.pid.generate"
	DDK_GoodsPromotionUrlGenerate MethodType = "pdd.ddk.goods.promotion.url.generate"
	DDK_LotteryUrlGen             MethodType = "pdd.ddk.lottery.url.gen"
	DDK_CMSPromUrlGenerate        MethodType = "pdd.ddk.cms.prom.url.generate"
	DDK_TopGoodsListQuery         MethodType = "pdd.ddk.top.goods.list.query"       //多多客获取爆款排行商品接口
	DDK_GoodsZsUnitUrlGen         MethodType = "pdd.ddk.goods.zs.unit.url.gen"      //多多进宝转链接口
	DDK_GoodsRecommendGet         MethodType = "pdd.ddk.goods.recommend.get"        //运营频道商品查询
	DDK_WeappQrcodeUrlGen         MethodType = "pdd.ddk.oauth.weapp.qrcode.url.gen" //多多客工具生成单品推广小程序二维码

	//商品 API
	GoodsCatsGet MethodType = "pdd.goods.cats.get" //商品标准类目接口
	GoodsOptGet  MethodType = "pdd.goods.opt.get"  //查询商品标签列表
)
