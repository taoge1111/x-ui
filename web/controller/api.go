package controller

import (
	"x-ui/web/service"

	"github.com/gin-gonic/gin"
)

type APIController struct {
	BaseController
	inboundController  *InboundController
	outboundController *OutboundController
	routerController   *RouterController
	Tgbot              service.Tgbot
}

func NewAPIController(g *gin.RouterGroup) *APIController {
	a := &APIController{}
	a.initRouter(g)
	return a
}

func (a *APIController) initRouter(g *gin.RouterGroup) {
	g.Use(a.checkLogin)
	inboundGroup := g.Group("/xui/API/inbounds")
	{
		inboundGroup.GET("/", a.inbounds)
		inboundGroup.GET("/get/:id", a.inbound)
		inboundGroup.GET("/getClientTraffics/:email", a.getClientTraffics)
		inboundGroup.POST("/add", a.addInbound)
		inboundGroup.POST("/del/:id", a.delInbound)
		inboundGroup.POST("/update/:id", a.updateInbound)
		inboundGroup.POST("/addClient", a.addInboundClient)
		inboundGroup.POST("/:id/delClient/:clientId", a.delInboundClient)
		inboundGroup.POST("/updateClient/:clientId", a.updateInboundClient)
		inboundGroup.POST("/:id/resetClientTraffic/:email", a.resetClientTraffic)
		inboundGroup.POST("/resetAllTraffics", a.resetAllTraffics)
		inboundGroup.POST("/resetAllClientTraffics/:id", a.resetAllClientTraffics)
		inboundGroup.POST("/delDepletedClients/:id", a.delDepletedClients)
		inboundGroup.GET("/createbackup", a.createBackup)
		inboundGroup.POST("/onlines", a.onlines)
		a.inboundController = NewInboundController(inboundGroup)
	}

	outboundGroup := g.Group("/xui/API/outbounds")
	{
		outboundGroup.POST("/add", a.AddOutbound)
		outboundGroup.POST("/delete", a.DeleteOutbound)
		a.outboundController = NewOutboundController(outboundGroup)
	}

	routerGroup := g.Group("/xui/API/routers")
	{
		routerGroup.POST("/add", a.AddRouter)
		routerGroup.GET("/list", a.GetRouter)
		routerGroup.POST("/delete", a.DeleteRouter)
		a.routerController = NewRouterController(routerGroup)
	}

}

func (a *APIController) inbounds(c *gin.Context) {
	a.inboundController.getInbounds(c)
}

func (a *APIController) inbound(c *gin.Context) {
	a.inboundController.getInbound(c)
}

func (a *APIController) getClientTraffics(c *gin.Context) {
	a.inboundController.getClientTraffics(c)
}

func (a *APIController) addInbound(c *gin.Context) {
	a.inboundController.addInbound(c)
}

func (a *APIController) delInbound(c *gin.Context) {
	a.inboundController.delInbound(c)
}

func (a *APIController) updateInbound(c *gin.Context) {
	a.inboundController.updateInbound(c)
}

func (a *APIController) addInboundClient(c *gin.Context) {
	a.inboundController.addInboundClient(c)
}

func (a *APIController) delInboundClient(c *gin.Context) {
	a.inboundController.delInboundClient(c)
}

func (a *APIController) updateInboundClient(c *gin.Context) {
	a.inboundController.updateInboundClient(c)
}

func (a *APIController) resetClientTraffic(c *gin.Context) {
	a.inboundController.resetClientTraffic(c)
}

func (a *APIController) resetAllTraffics(c *gin.Context) {
	a.inboundController.resetAllTraffics(c)
}

func (a *APIController) resetAllClientTraffics(c *gin.Context) {
	a.inboundController.resetAllClientTraffics(c)
}

func (a *APIController) delDepletedClients(c *gin.Context) {
	a.inboundController.delDepletedClients(c)
}

func (a *APIController) createBackup(c *gin.Context) {
	a.Tgbot.SendBackupToAdmins()
}

func (a *APIController) onlines(c *gin.Context) {
	a.inboundController.onlines(c)
}

func (a *APIController) AddOutbound(c *gin.Context) {
	a.outboundController.addOutbound(c)
}

func (a *APIController) DeleteOutbound(c *gin.Context) {
	a.outboundController.deleteOutbound(c)
}

func (a *APIController) AddRouter(c *gin.Context) {
	a.routerController.AddRouter(c)
}

func (a *APIController) GetRouter(c *gin.Context) {
	a.routerController.GetRouter(c)
}

func (a *APIController) DeleteRouter(c *gin.Context) {
	a.routerController.DeleteRouter(c)
}
