package infrastructure

import "github.com/gin-gonic/gin"

// ServerInterface represents all server handlers.
type ServerInterface interface {
	Ping(c *gin.Context)
	GetReferralLink(c *gin.Context)
	ProcessRequest(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(c *gin.Context)

// Ping operation middleware
func (siw *ServerInterfaceWrapper) Ping(c *gin.Context) {
	siw.Handler.Ping(c)
}

func (siw *ServerInterfaceWrapper) GetReferralLink(c *gin.Context) {
	siw.Handler.GetReferralLink(c)
}

func (siw *ServerInterfaceWrapper) ProcessRequest(c *gin.Context) {
	siw.Handler.ProcessRequest(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.Engine, si ServerInterface) *gin.Engine {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions) *gin.Engine {
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	router.GET(options.BaseURL+"/ping", wrapper.Ping)
	router.POST(options.BaseURL+"/referral_link", wrapper.GetReferralLink)
	router.GET(options.BaseURL+"/open_position", wrapper.ProcessRequest)
	return router
}
