package pdd

import (
	"encoding/json"
	"fmt"
)

//MallCoupon MallCoupon
type MallCoupon struct {
	CouponQuantity       int64 `json:"coupon_quantity"`
	MinOrderAmount       int64 `json:"min_order_amount"`
	CouponID             int64 `json:"coupon_id"`
	CouponRemainQuantity int64 `json:"coupon_remain_quantity"`
	MallID               int64 `json:"mall_id"`
	Discount             int64 `json:"discount"`
	CouponType           int64 `json:"coupon_type"`
	MaxDiscountAmount    int64 `json:"max_discount_amount"`
	CouponStartTime      int64 `json:"coupon_start_time"`
	CouponEndTime        int64 `json:"coupon_end_time"`
}

//Mall /Mall
type Mall struct {
	MallID             int64        `json:"mall_id"`
	ImgURL             string       `json:"img_url"`
	MallCouponInfoList []MallCoupon `json:"mall_coupon_info_list"`
	GoodsDetailVoList  []Goods      `json:"goods_detail_vo_list"`
	MallName           string       `json:"mall_name"`
	GoodsNum           string       `json:"goods_num"`
	MallRate           int64        `json:"mall_rate"`
	MerchantType       int64        `json:"merchant_type"`
	DescPct            float64      `json:"desc_pct"`
	AvgDesc            float64      `json:"avg_desc"`
	AvgServ            float64      `json:"avg_serv"`
	AvgLgst            float64      `json:"avg_lgst"`
	ServPct            float64      `json:"serv_pct"`
	SoldQuantity       float64      `json:"sold_quantity"`
	CatIDList          []int        `json:"cat_id_list"`
	LgstPct            float64      `json:"lgst_pct"`
}

//MallResponse MallResponse
type MallResponse struct {
	Total                int64  `json:"total"`
	MallSearchInfoVoList []Mall `json:"mall_search_info_vo_list"`
}

//MallDetail 多多客查店铺详情接口
func (d *DDK) MallDetail(mallid int, notMustparams ...Params) (res *Mall, err error) {
	res = new(Mall)
	params := NewParamsWithType(DDK_MerchantListGet, notMustparams...)
	params.Set("mall_id_list", fmt.Sprintf("[%d]", mallid))

	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseArrayIndexBytes(r, 0, "merchant_list_response", "mall_search_info_vo_list")
	err = json.Unmarshal(bytes, res)
	return
}

//MallCouponGenerateURLResponse MallCouponGenerateURLResponse
type MallCouponGenerateURLResponse struct {
	URL                  string `json:"url"`
	ShortURL             string `json:"short_url"`
	MobileURL            string `json:"mobile_url"`
	MobileShortURL       string `json:"mobile_short_url"`
	WeAppWebViewURL      string `json:"we_app_web_view_url"`
	WeAppWebViewShortURL string `json:"we_app_web_view_short_url"`
}

//MallURLGen 多多客生成店铺推广链接API
func (d *DDK) MallURLGen(mallid int64, pid string, notMustparams ...Params) (res *MallCouponGenerateURLResponse, err error) {
	res = new(MallCouponGenerateURLResponse)
	params := NewParamsWithType(DDK_MallUrlGen, notMustparams...)
	params.Set("mall_id", mallid)
	params.Set("pid", pid)

	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseArrayIndexBytes(r, 0, "mall_coupon_generate_url_response", "list")
	err = json.Unmarshal(bytes, res)
	return
}
