package controller

import (
	"encoding/json"
	"net/http"
	"sync"
	"x-ui/database/model"
	"x-ui/util/json_util"
	"x-ui/web/service"

	"github.com/gin-gonic/gin"
)

type RouterController struct {
	xrayService        service.XrayService
	xraySettingService service.XraySettingService
	lock               *sync.Mutex
}

func NewRouterController(g *gin.RouterGroup) *RouterController {
	c := &RouterController{}
	c.initRouter(g)
	c.lock = &sync.Mutex{}
	return c
}

func (a *RouterController) initRouter(g *gin.RouterGroup) {
	g = g.Group("/router")
	g.POST("/add", a.AddRouter)
	g.POST("/delete", a.DeleteRouter)
	g.GET("/xray/config", a.GetRouter)
}

// err := a.XraySettingService.SaveXraySetting(xraySetting)
func (routerController *RouterController) AddRouter(c *gin.Context) {
	routerRule := &model.RouterRule{}
	err := c.ShouldBind(routerRule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "bind router error"+err.Error())
	}

	routerRuleObj, err := json.MarshalIndent(routerRule, "", " ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "marshal error"+err.Error())
	}

	routerController.lock.Lock()
	defer routerController.lock.Unlock()
	config, err := routerController.xrayService.GetXrayConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "get config error")
	}
	xrayConfig, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "xrayConfig marshal error"+err.Error())
	}

	res, err := json_util.XrayConfigUtil{}.AddRouterRules(string(xrayConfig), string(routerRuleObj))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "XrayConfigUtil AddRouterRules error"+err.Error())
	}

	err = routerController.xraySettingService.SaveXraySetting(res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "SaveXraySetting error"+err.Error())
	}
	routerController.xrayService.SetToNeedRestart()
	c.JSON(http.StatusOK, routerRule)
}

func (routerController *RouterController) GetRouter(c *gin.Context) {

	config, err := routerController.xrayService.GetXrayConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "get config error")
	}

	c.JSON(http.StatusOK, config)
}

func (routerController *RouterController) DeleteRouter(c *gin.Context) {

	routerRule := &model.RouterRule{}
	err := c.ShouldBind(routerRule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "bind router error"+err.Error())
	}

	routerRuleObj, err := json.MarshalIndent(routerRule, "", " ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "marshal error"+err.Error())
	}

	routerController.lock.Lock()
	defer routerController.lock.Unlock()
	config, err := routerController.xrayService.GetXrayConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "get config error")
	}
	xrayConfig, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "xrayConfig marshal error"+err.Error())
	}

	res, err := json_util.XrayConfigUtil{}.DeleteRouterRules(string(xrayConfig), string(routerRuleObj))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "XrayConfigUtil AddRouterRules error"+err.Error())
	}

	err = routerController.xraySettingService.SaveXraySetting(res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "SaveXraySetting error"+err.Error())
	}
	routerController.xrayService.SetToNeedRestart()
	c.JSON(http.StatusOK, routerRule)
}
