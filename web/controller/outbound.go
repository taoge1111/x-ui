package controller

import (
	"net/http"
	"x-ui/database/model"
	"x-ui/web/service"

	"github.com/gin-gonic/gin"
)

type OutboundController struct {
	outboundService service.OutboundService
	xrayService     service.XrayService
}

func NewOutboundController(g *gin.RouterGroup) *OutboundController {
	a := &OutboundController{}
	a.initRouter(g)
	return a
}

func (a *OutboundController) initRouter(g *gin.RouterGroup) {
	g = g.Group("/outbound")
	g.POST("/add", a.addOutbound)
}

func (a *OutboundController) addOutbound(c *gin.Context) {
	outbound := &model.Outbound{}
	err := c.ShouldBind(outbound)
	if err != nil {
		jsonMsg(c, I18nWeb(c, "pages.outbounds.create"), err)
		return
	}

	needRestart := false
	outbound, needRestart, err = a.outboundService.AddOutbound(outbound)
	jsonMsgObj(c, I18nWeb(c, "pages.outbounds.create"), outbound, err)
	if err == nil && needRestart {
		a.xrayService.SetToNeedRestart()
	}
	//c.JSON(http.StatusOK, outbound)
}

func (a *OutboundController) deleteOutbound(c *gin.Context) {
	outbound := &model.Outbound{}
	err := c.ShouldBind(outbound)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	needRestart, err := a.outboundService.DelOutbound(outbound.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err == nil && needRestart {
		a.xrayService.SetToNeedRestart()
	}
	c.JSON(http.StatusOK, outbound)

}

// func (a *OutboundController) delInbound(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		jsonMsg(c, I18nWeb(c, "delete"), err)
// 		return
// 	}
// 	needRestart := true
// 	needRestart, err = a.inboundService.DelInbound(id)
// 	jsonMsgObj(c, I18nWeb(c, "delete"), id, err)
// 	if err == nil && needRestart {
// 		a.xrayService.SetToNeedRestart()
// 	}
// }

// func (a *OutboundController) updateInbound(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		jsonMsg(c, I18nWeb(c, "pages.inbounds.update"), err)
// 		return
// 	}
// 	inbound := &model.Inbound{
// 		Id: id,
// 	}
// 	err = c.ShouldBind(inbound)
// 	if err != nil {
// 		jsonMsg(c, I18nWeb(c, "pages.inbounds.update"), err)
// 		return
// 	}
// 	needRestart := true
// 	inbound, needRestart, err = a.inboundService.UpdateInbound(inbound)
// 	jsonMsgObj(c, I18nWeb(c, "pages.inbounds.update"), inbound, err)
// 	if err == nil && needRestart {
// 		a.xrayService.SetToNeedRestart()
// 	}
// }

// func (a *OutboundController) getInbound(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		jsonMsg(c, I18nWeb(c, "get"), err)
// 		return
// 	}
// 	inbound, err := a.inboundService.GetInbound(id)
// 	if err != nil {
// 		jsonMsg(c, I18nWeb(c, "pages.inbounds.toasts.obtain"), err)
// 		return
// 	}
// 	jsonObj(c, inbound, nil)
// }
// func (a *OutboundController) getClientTraffics(c *gin.Context) {
// 	email := c.Param("email")
// 	clientTraffics, err := a.inboundService.GetClientTrafficByEmail(email)
// 	if err != nil {
// 		jsonMsg(c, "Error getting traffics", err)
// 		return
// 	}
// 	jsonObj(c, clientTraffics, nil)
// }
