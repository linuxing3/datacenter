// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	memberaddress "datacenter/internal/handler/member/address"
	membergrowthchangehistory "datacenter/internal/handler/member/growthchangehistory"
	memberintegrationchangehistory "datacenter/internal/handler/member/integrationchangehistory"
	memberintegrationconsumesetting "datacenter/internal/handler/member/integrationconsumesetting"
	memberlevel "datacenter/internal/handler/member/level"
	memberloginlog "datacenter/internal/handler/member/loginlog"
	membermember "datacenter/internal/handler/member/member"
	memberrulesetting "datacenter/internal/handler/member/rulesetting"
	memberstatistics "datacenter/internal/handler/member/statistics"
	membertag "datacenter/internal/handler/member/tag"
	membertask "datacenter/internal/handler/member/task"
	ordercart "datacenter/internal/handler/order/cart"
	ordercompayaddress "datacenter/internal/handler/order/compayaddress"
	orderoperatehistory "datacenter/internal/handler/order/operatehistory"
	orderorder "datacenter/internal/handler/order/order"
	orderreturnapply "datacenter/internal/handler/order/returnapply"
	orderreturnreason "datacenter/internal/handler/order/returnreason"
	ordersetting "datacenter/internal/handler/order/setting"
	productbrand "datacenter/internal/handler/product/brand"
	productcategory "datacenter/internal/handler/product/category"
	productcomment "datacenter/internal/handler/product/comment"
	productfeighttemplate "datacenter/internal/handler/product/feighttemplate"
	productmemberprice "datacenter/internal/handler/product/memberprice"
	productproduct "datacenter/internal/handler/product/product"
	productskustock "datacenter/internal/handler/product/skustock"
	smscoupon "datacenter/internal/handler/sms/coupon"
	smscouponhistory "datacenter/internal/handler/sms/couponhistory"
	smsflashpromotion "datacenter/internal/handler/sms/flashpromotion"
	smsflashpromotionlog "datacenter/internal/handler/sms/flashpromotionlog"
	smsflashpromotionsession "datacenter/internal/handler/sms/flashpromotionsession"
	smshomeadvertise "datacenter/internal/handler/sms/homeadvertise"
	smshomebrand "datacenter/internal/handler/sms/homebrand"
	smshomenewproduct "datacenter/internal/handler/sms/homenewproduct"
	smshomerecommendproduct "datacenter/internal/handler/sms/homerecommendproduct"
	smshomerecommendsubject "datacenter/internal/handler/sms/homerecommendsubject"
	// sysconfig "datacenter/internal/handler/sys/config"
	// sysdept "datacenter/internal/handler/sys/dept"
	// sysdict "datacenter/internal/handler/sys/dict"
	// syslog "datacenter/internal/handler/sys/log"
	// sysmenu "datacenter/internal/handler/sys/menu"
	// sysrole "datacenter/internal/handler/sys/role"
	// sysuser "datacenter/internal/handler/sys/user"
	"datacenter/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterGoZeroHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	// engine.AddRoutes(
	// 	[]rest.Route{
	// 		{
	// 			Method:  http.MethodGet,
	// 			Path:    "/api/sys/user/currentUser",
	// 			Handler: sysuser.UserInfoHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/user/add",
	// 			Handler: sysuser.UserAddHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/user/list",
	// 			Handler: sysuser.UserListHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/user/update",
	// 			Handler: sysuser.UserUpdateHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/user/delete",
	// 			Handler: sysuser.UserDeleteHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/user/reSetPassword",
	// 			Handler: sysuser.ReSetPasswordHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/user/UpdateUserStatus",
	// 			Handler: sysuser.UpdateUserStatusHandler(serverCtx),
	// 		},
	// 	},
	// 	rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	// )

	// engine.AddRoutes(
	// 	[]rest.Route{
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/user/login",
	// 			Handler: sysuser.UserLoginHandler(serverCtx),
	// 		},
	// 	},
	// )

	// engine.AddRoutes(
	// 	[]rest.Route{
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/role/add",
	// 			Handler: sysrole.RoleAddHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/role/list",
	// 			Handler: sysrole.RoleListHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/role/update",
	// 			Handler: sysrole.RoleUpdateHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/role/delete",
	// 			Handler: sysrole.RoleDeleteHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/role/roleMenuIds",
	// 			Handler: sysrole.RoleMenuIdsHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/role/queryMenuByRoleId",
	// 			Handler: sysrole.QueryMenuByRoleIdHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/role/updateRoleMenu",
	// 			Handler: sysrole.UpdateRoleMenuHandler(serverCtx),
	// 		},
	// 	},
	// 	rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	// )

	// engine.AddRoutes(
	// 	[]rest.Route{
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/menu/add",
	// 			Handler: sysmenu.MenuAddHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/menu/list",
	// 			Handler: sysmenu.MenuListHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/menu/update",
	// 			Handler: sysmenu.MenuUpdateHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/menu/delete",
	// 			Handler: sysmenu.MenuDeleteHandler(serverCtx),
	// 		},
	// 	},
	// 	rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	// )

	// engine.AddRoutes(
	// 	[]rest.Route{
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/dict/add",
	// 			Handler: sysdict.DictAddHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/dict/list",
	// 			Handler: sysdict.DictListHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/dict/update",
	// 			Handler: sysdict.DictUpdateHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/dict/delete",
	// 			Handler: sysdict.DictDeleteHandler(serverCtx),
	// 		},
	// 	},
	// 	rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	// )

	// engine.AddRoutes(
	// 	[]rest.Route{
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/dept/add",
	// 			Handler: sysdept.DeptAddHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/dept/list",
	// 			Handler: sysdept.DeptListHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/dept/update",
	// 			Handler: sysdept.DeptUpdateHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/dept/delete",
	// 			Handler: sysdept.DeptDeleteHandler(serverCtx),
	// 		},
	// 	},
	// 	rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	// )

	// engine.AddRoutes(
	// 	[]rest.Route{
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/loginLog/list",
	// 			Handler: syslog.LoginLogListHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/loginLog/delete",
	// 			Handler: syslog.LoginLogDeleteHandler(serverCtx),
	// 		},
	// 	},
	// 	rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	// )

	// engine.AddRoutes(
	// 	[]rest.Route{
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/sysLog/list",
	// 			Handler: syslog.SysLogListHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/sysLog/delete",
	// 			Handler: syslog.SysLogDeleteHandler(serverCtx),
	// 		},
	// 	},
	// 	rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	// )

	// engine.AddRoutes(
	// 	[]rest.Route{
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/config/add",
	// 			Handler: sysconfig.ConfigAddHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/config/list",
	// 			Handler: sysconfig.ConfigListHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/config/update",
	// 			Handler: sysconfig.ConfigUpdateHandler(serverCtx),
	// 		},
	// 		{
	// 			Method:  http.MethodPost,
	// 			Path:    "/api/sys/config/delete",
	// 			Handler: sysconfig.ConfigDeleteHandler(serverCtx),
	// 		},
	// 	},
	// 	rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	// )

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/order/cart/add",
				Handler: ordercart.CartItemAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/cart/list",
				Handler: ordercart.CartItemListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/cart/update",
				Handler: ordercart.CartItemUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/cart/delete",
				Handler: ordercart.CartItemDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/order/order/add",
				Handler: orderorder.OrderAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/order/list",
				Handler: orderorder.OrderListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/order/update",
				Handler: orderorder.OrderUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/order/delete",
				Handler: orderorder.OrderDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/order/compayaddress/add",
				Handler: ordercompayaddress.CompayAddressAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/compayaddress/list",
				Handler: ordercompayaddress.CompayAddressListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/compayaddress/update",
				Handler: ordercompayaddress.CompayAddressUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/compayaddress/delete",
				Handler: ordercompayaddress.CompayAddressDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/order/operatehistory/add",
				Handler: orderoperatehistory.OperateHistoryAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/operatehistory/list",
				Handler: orderoperatehistory.OperateHistoryListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/operatehistory/update",
				Handler: orderoperatehistory.OperateHistoryUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/operatehistory/delete",
				Handler: orderoperatehistory.OperateHistoryDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/order/returnapply/add",
				Handler: orderreturnapply.ReturnApplyAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/returnapply/list",
				Handler: orderreturnapply.ReturnApplyListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/returnapply/update",
				Handler: orderreturnapply.ReturnApplyUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/returnapply/delete",
				Handler: orderreturnapply.ReturnApplyDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/order/returnreason/add",
				Handler: orderreturnreason.ReturnResonAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/returnreason/list",
				Handler: orderreturnreason.ReturnResonListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/returnreason/update",
				Handler: orderreturnreason.ReturnResonUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/returnreason/delete",
				Handler: orderreturnreason.ReturnResonDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/order/setting/add",
				Handler: ordersetting.OrderSettingAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/setting/list",
				Handler: ordersetting.OrderSettingListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/setting/update",
				Handler: ordersetting.OrderSettingUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/setting/delete",
				Handler: ordersetting.OrderSettingDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/product/product/add",
				Handler: productproduct.ProductAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/product/list",
				Handler: productproduct.ProductListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/product/update",
				Handler: productproduct.ProductUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/product/delete",
				Handler: productproduct.ProductDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/product/brand/add",
				Handler: productbrand.ProductBrandAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/brand/list",
				Handler: productbrand.ProductBrandListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/brand/update",
				Handler: productbrand.ProductBrandUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/brand/delete",
				Handler: productbrand.ProductBrandDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/product/category/add",
				Handler: productcategory.ProductCategoryAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/category/list",
				Handler: productcategory.ProductCategoryListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/category/update",
				Handler: productcategory.ProductCategoryUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/category/delete",
				Handler: productcategory.ProductCategoryDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/product/comment/add",
				Handler: productcomment.ProductCommentAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/comment/list",
				Handler: productcomment.ProductCommentListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/comment/update",
				Handler: productcomment.ProductCommentUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/comment/delete",
				Handler: productcomment.ProductCommentDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/product/feighttemplate/add",
				Handler: productfeighttemplate.FeightTemplateAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/feighttemplate/list",
				Handler: productfeighttemplate.FeightTemplateListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/feighttemplate/update",
				Handler: productfeighttemplate.FeightTemplateUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/feighttemplate/delete",
				Handler: productfeighttemplate.FeightTemplateDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/product/memberprice/add",
				Handler: productmemberprice.MemberPriceAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/memberprice/list",
				Handler: productmemberprice.MemberPriceListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/memberprice/update",
				Handler: productmemberprice.MemberPriceUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/memberprice/delete",
				Handler: productmemberprice.MemberPriceDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/product/skustock/add",
				Handler: productskustock.SkuStockAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/skustock/list",
				Handler: productskustock.SkuStockListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/skustock/update",
				Handler: productskustock.SkuStockUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/skustock/delete",
				Handler: productskustock.SkuStockDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/coupon/add",
				Handler: smscoupon.CouponAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/coupon/list",
				Handler: smscoupon.CouponListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/coupon/update",
				Handler: smscoupon.CouponUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/coupon/delete",
				Handler: smscoupon.CouponDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/couponhistory/add",
				Handler: smscouponhistory.CouponHistoryAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/couponhistory/list",
				Handler: smscouponhistory.CouponHistoryListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/couponhistory/update",
				Handler: smscouponhistory.CouponHistoryUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/couponhistory/delete",
				Handler: smscouponhistory.CouponHistoryDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/flashpromotion/add",
				Handler: smsflashpromotion.FlashPromotionAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/flashpromotion/list",
				Handler: smsflashpromotion.FlashPromotionListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/flashpromotion/update",
				Handler: smsflashpromotion.FlashPromotionUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/flashpromotion/delete",
				Handler: smsflashpromotion.FlashPromotionDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/flashpromotionlog/add",
				Handler: smsflashpromotionlog.FlashPromotionLogAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/flashpromotionlog/list",
				Handler: smsflashpromotionlog.FlashPromotionLogListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/flashpromotionlog/update",
				Handler: smsflashpromotionlog.FlashPromotionLogUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/flashpromotionlog/delete",
				Handler: smsflashpromotionlog.FlashPromotionLogDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/flashpromotionsession/add",
				Handler: smsflashpromotionsession.FlashPromotionSessionAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/flashpromotionsession/list",
				Handler: smsflashpromotionsession.FlashPromotionSessionListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/flashpromotionsession/update",
				Handler: smsflashpromotionsession.FlashPromotionSessionUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/flashpromotionsession/delete",
				Handler: smsflashpromotionsession.FlashPromotionSessionDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/homeadvertise/add",
				Handler: smshomeadvertise.HomeAdvertiseAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/homeadvertise/list",
				Handler: smshomeadvertise.HomeAdvertiseListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/homeadvertise/update",
				Handler: smshomeadvertise.HomeAdvertiseUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/homeadvertise/delete",
				Handler: smshomeadvertise.HomeAdvertiseDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/homebrand/add",
				Handler: smshomebrand.HomeBrandAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/homebrand/list",
				Handler: smshomebrand.HomeBrandListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/homebrand/update",
				Handler: smshomebrand.HomeBrandUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/homebrand/delete",
				Handler: smshomebrand.HomeBrandDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/homenewproduct/add",
				Handler: smshomenewproduct.HomeNewProductAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/homenewproduct/list",
				Handler: smshomenewproduct.HomeNewProductListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/homenewproduct/update",
				Handler: smshomenewproduct.HomeNewProductUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/homenewproduct/delete",
				Handler: smshomenewproduct.HomeNewProductDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/homerecommendproduct/add",
				Handler: smshomerecommendproduct.HomeRecommendProductAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/homerecommendproduct/list",
				Handler: smshomerecommendproduct.HomeRecommendProductListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/homerecommendproduct/update",
				Handler: smshomerecommendproduct.HomeRecommendProductUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/homerecommendproduct/delete",
				Handler: smshomerecommendproduct.HomeRecommendProductDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/homerecommendsubject/add",
				Handler: smshomerecommendsubject.HomeRecommendSubjectAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/homerecommendsubject/list",
				Handler: smshomerecommendsubject.HomeRecommendSubjectListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/apisms/homerecommendsubject/update",
				Handler: smshomerecommendsubject.HomeRecommendSubjectUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sms/homerecommendsubject/delete",
				Handler: smshomerecommendsubject.HomeRecommendSubjectDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/member/member/add",
				Handler: membermember.MemberAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/member/list",
				Handler: membermember.MemberListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/member/update",
				Handler: membermember.MemberUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/member/delete",
				Handler: membermember.MemberDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/member/address/add",
				Handler: memberaddress.MemberAddressAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/address/list",
				Handler: memberaddress.MemberAddressListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/address/update",
				Handler: memberaddress.MemberAddressUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/address/delete",
				Handler: memberaddress.MemberAddressDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/member/growthchangehistory/add",
				Handler: membergrowthchangehistory.GrowthChangeHistoryAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/growthchangehistory/list",
				Handler: membergrowthchangehistory.GrowthChangeHistoryListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/growthchangehistory/update",
				Handler: membergrowthchangehistory.GrowthChangeHistoryUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/growthchangehistory/delete",
				Handler: membergrowthchangehistory.GrowthChangeHistoryDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/member/integrationchangehistory/add",
				Handler: memberintegrationchangehistory.IntegrationChangeHistoryAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/integrationchangehistory/list",
				Handler: memberintegrationchangehistory.IntegrationChangeHistoryListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/integrationchangehistory/update",
				Handler: memberintegrationchangehistory.IntegrationChangeHistoryUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/integrationchangehistory/delete",
				Handler: memberintegrationchangehistory.IntegrationChangeHistoryDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/member/integrationconsumesetting/add",
				Handler: memberintegrationconsumesetting.IntegrationConsumeSettingAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/integrationconsumesetting/list",
				Handler: memberintegrationconsumesetting.IntegrationConsumeSettingListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/integrationconsumesetting/update",
				Handler: memberintegrationconsumesetting.IntegrationConsumeSettingUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/integrationconsumesetting/delete",
				Handler: memberintegrationconsumesetting.IntegrationConsumeSettingDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/member/level/add",
				Handler: memberlevel.MemberLevelAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/level/list",
				Handler: memberlevel.MemberLevelListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/level/update",
				Handler: memberlevel.MemberLevelUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/level/delete",
				Handler: memberlevel.MemberLevelDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/member/loginlog/add",
				Handler: memberloginlog.MemberLoginLogAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/loginlog/list",
				Handler: memberloginlog.MemberLoginLogListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/loginlog/update",
				Handler: memberloginlog.MemberLoginLogUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/loginlog/delete",
				Handler: memberloginlog.MemberLoginLogDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/member/rulesetting/add",
				Handler: memberrulesetting.MemberRuleSettingAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/rulesetting/list",
				Handler: memberrulesetting.MemberRuleSettingListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/rulesetting/update",
				Handler: memberrulesetting.MemberRuleSettingUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/rulesetting/delete",
				Handler: memberrulesetting.MemberRuleSettingDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/member/statistics/add",
				Handler: memberstatistics.MemberStatisticsInfoAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/statistics/list",
				Handler: memberstatistics.MemberStatisticsInfoListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/statistics/update",
				Handler: memberstatistics.MemberStatisticsInfoUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/statistics/delete",
				Handler: memberstatistics.MemberStatisticsInfoDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/member/tag/add",
				Handler: membertag.MemberTagAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/tag/list",
				Handler: membertag.MemberTagListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/tag/update",
				Handler: membertag.MemberTagUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/tag/delete",
				Handler: membertag.MemberTagDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/member/task/add",
				Handler: membertask.MemberTaskAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/task/list",
				Handler: membertask.MemberTaskListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/task/update",
				Handler: membertask.MemberTaskUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/member/task/delete",
				Handler: membertask.MemberTaskDeleteHandler(serverCtx),
			},
		},
	)
}
