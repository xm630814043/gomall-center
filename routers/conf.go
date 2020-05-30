package routers

import (
	"gomall-center/api"
	"gomall-center/pkg/web"

	"github.com/gin-gonic/gin"
)

// Register 配置路由
func Register(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	// 栏目管理
	v1.GET("/cms/channel/list", web.Handler(api.FindChannelList))     // 频道专题列表
	v1.GET("/cms/subject/list", web.Handler(api.FindBySubjectList))   // 专题内容列表
	v1.DELETE("/cms/subject/:id", web.Handler(api.RemoveSubject))     // 删除专题内容
	v1.POST("/cms/subject/create", web.Handler(api.AddSubject))       // 添加主题内容
	v1.GET("/cms/subject/detail/:id", web.Handler(api.FindBySubject)) // 专题内容详情

	// 企业管理
	v1.GET("/company/list", web.Handler(api.FindCompanyList))     // 企业列表
	v1.GET("/company/detail/:id", web.Handler(api.FindByCompany)) // 企业详情
	v1.PUT("/company/status/:id", web.Handler(api.ModifyCompany)) // 修改企业状态
	v1.GET("/account/list", web.Handler(api.FindAccountList))     // 用户列表
	v1.GET("/account/detail/:id", web.Handler(api.FindByAccount)) // 用户详情

	//商品审核
	v1.GET("/product/basic/status", web.Handler(api.FindProductBasicByStatus)) //根据产品状态查找产品基本信息
	v1.GET("/product/info/:id", web.Handler(api.FindProductDetails))           //根据产品id展示需要审核的信息
	v1.PUT("/product/status/:id", web.Handler(api.ModifyProductStatus))        //修改商品发布状态

	//商品属性管理
	v1.GET("/categorys", web.Handler(api.FindCategoryList))                        //种类树
	v1.GET("/category/propertys/:id", web.Handler(api.FindPropByCategoryID))       //根据种类ID获取属性
	v1.GET("/category/property/value/:id", web.Handler(api.FindPropValueByPropID)) //根据属性ID获取属性值
	v1.POST("/category/property/value", web.Handler(api.AddPropValByPropID))       //根据属性ID增加属性值
	v1.PUT("/category/property/value/:id", web.Handler(api.ModifyPropValue))       //修改分类属性的值prop_value
	v1.DELETE("/category/property/value/:id", web.Handler(api.RemovePropValue))    //根据propValueID删除属性值
	v1.POST("/category/property", web.Handler(api.AddProp))                        //添加商品分类属性根据分类Id
	v1.DELETE("/category/property/remove/:id", web.Handler(api.RemoveProp))        //删除商品分类属性根据分类Id
	v1.PUT("/category/property/edit/:id", web.Handler(api.ModifyProp))             //修改属性根据ID

	//测试
	v1.POST("/controlled/create", web.Handler(api.AddControlSell))
	v1.GET("/controlled/detail/:id", web.Handler(api.FindByControlSell))
	v1.GET("/controlled/list/:id", web.Handler(api.FindControlSellList))
	v1.DELETE("/controlled/:id", web.Handler(api.RemoveControlSell))
	v1.PUT("/controlled/edit/:id", web.Handler(api.ModifyControlSell))

	v1.POST("/article/create", web.Handler(api.AddArticle))
	v1.GET("/article/detail/:id", web.Handler(api.FindByArticle))
	v1.GET("/article/list", web.Handler(api.FindArticleList))
	v1.DELETE("/article/:id", web.Handler(api.RemoveArticle))
	v1.PUT("/article/edit/:id", web.Handler(api.ModifyArticle))

	v1.POST("/promotionActivity/create", web.Handler(api.AddPromotionActivity))
	v1.POST("/promotionActivity/create/:id", web.Handler(api.AddPromotionActivityAbs))
	v1.DELETE("/promotionActivity/:id", web.Handler(api.RemovePromotionActivity))
	v1.PUT("/promotionActivity/edit/:id", web.Handler(api.ModifyPromotionActivity))
	v1.GET("/promotionActivity/detail/:id", web.Handler(api.FindByPromotionActivity))
	v1.GET("/promotionActivity/list/:id", web.Handler(api.FindPromotionActivityList))
	v1.GET("/promotionProduct/list", web.Handler(api.FindPromotionProductList))
	v1.GET("/discountCoupon/list", web.Handler(api.FindDiscountCouponList))


	v1.GET("/promotionPattern/list", web.Handler(api.FIndDiscounts))



}
