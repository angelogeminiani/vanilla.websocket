package application_www

import (
	"github.com/angelogeminiani/vanilla.websocket/server/application/application_types"
	"github.com/botikasm/lygo/base/lygo_events"
	"github.com/botikasm/lygo/ext/lygo_http/lygo_http_server"
	"github.com/botikasm/lygo/ext/lygo_http/lygo_http_server/lygo_http_server_config"
	"github.com/botikasm/lygo/ext/lygo_http/lygo_http_server/lygo_http_server_types"
	"github.com/botikasm/lygo/ext/lygo_logs"
	"github.com/gofiber/fiber"
)

// ---------------------------------------------------------------------------------------------------------------------
// 		c o n s t
// ---------------------------------------------------------------------------------------------------------------------

const (
	wsPath = "/websocket"
)

// ---------------------------------------------------------------------------------------------------------------------
// 		t y p e s
// ---------------------------------------------------------------------------------------------------------------------

type ApplicationServer struct {
	events   *lygo_events.Emitter
	settings *lygo_http_server_config.HttpServerConfig
	server   *lygo_http_server.HttpServer
	handler  *MessageHandler
}

// ---------------------------------------------------------------------------------------------------------------------
// 		p u b l i c
// ---------------------------------------------------------------------------------------------------------------------

func NewApplicationServer(events *lygo_events.Emitter,
	settings *lygo_http_server_config.HttpServerConfig) *ApplicationServer {

	instance := new(ApplicationServer)
	instance.events = events
	instance.settings = settings
	instance.handler = NewMessageHandler(events)

	return instance
}

func (instance *ApplicationServer) Start() error {
	if nil != instance && nil == instance.server && nil != instance.settings {
		instance.server = lygo_http_server.NewHttpServer(instance.settings)
		return instance.openServer(instance.server)
	}
	return application_types.PanicSystemError
}

func (instance *ApplicationServer) Close() error {
	if nil != instance && nil != instance.server {
		if nil != instance.server {
			instance.server.Stop()
			instance.server = nil
		}
		return nil
	}
	return application_types.PanicSystemError
}

func (instance *ApplicationServer) Join() error {
	if nil != instance {
		if nil != instance.server {
			instance.server.Join()
		}
		return nil
	}
	return application_types.PanicSystemError
}

// ---------------------------------------------------------------------------------------------------------------------
// 		p r i v a t e
// ---------------------------------------------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------------------------------------------
//  	S T A T I C
// ---------------------------------------------------------------------------------------------------------------------

func (instance *ApplicationServer) openServer(server *lygo_http_server.HttpServer) error {
	if nil != instance && nil != instance.server && nil != instance.events {
		server.CallbackError = instance.onError
		server.CallbackLimitReached = instance.onLimit

		server.Websocket(wsPath, func(ws *lygo_http_server.HttpWebsocketConn) {
			ws.OnMessage(instance.handleWs)
		})

		return server.Start()
	}
	return application_types.PanicSystemError
}

func (instance *ApplicationServer) onError(errCtx *lygo_http_server_types.HttpServerError) {
	if nil != instance && nil != instance.server && nil != instance.events {
		// fmt.Println(errCtx.Message, errCtx.Error.Error())
		lygo_logs.Error(errCtx.Message, errCtx.Error.Error())
		instance.events.Emit(application_types.EventError, application_types.ContextWebsocket, errCtx, errCtx.Error)
	}
}

func (instance *ApplicationServer) onLimit(c *fiber.Ctx) {
	if nil != instance && nil != instance.server && nil != instance.events {
		c.Send("too many requests: limit exceeded")
		instance.events.Emit(application_types.EventError, application_types.ContextWebsocket, "too many requests: limit exceeded", c.Error().Error())
	}
}

func (instance *ApplicationServer) handleWs(payload *lygo_http_server.HttpWebsocketEventPayload) {
	if nil != instance && nil != instance.server && nil != instance.handler {
		instance.handler.Handle(payload)
	}
}
