package application_www

import (
	"encoding/json"
	"errors"
	"github.com/angelogeminiani/vanilla.websocket/server/application/application_types"
	"github.com/botikasm/lygo/base/lygo_conv"
	"github.com/botikasm/lygo/base/lygo_events"
	"github.com/botikasm/lygo/ext/lygo_http/lygo_http_server"
)

// ---------------------------------------------------------------------------------------------------------------------
// 		t y p e s
// ---------------------------------------------------------------------------------------------------------------------

type MessageHandler struct {
	//database *meeteese_app_database.ApplicationDatabase
	events *lygo_events.Emitter
}

type Message struct {
	RequestUUID string           `json:"request_uuid"`
	Lang        string           `json:"lang"`
	UID         string           `json:"uid"` // user_id
	Payload     *MessagePayload  `json:"payload"`
	Response    *MessageResponse `json:"response"`
}

type MessagePayload struct {
	AppToken  string      `json:"app_token"`
	Namespace string      `json:"namespace"`
	Function  string      `json:"function"`
	Params    interface{} `json:"params"`
}
type MessageResponse struct {
	Error string        `json:"error"`
	Data  []interface{} `json:"data"`
}

func (instance *Message) Marshal() []byte {
	data, _ := json.Marshal(instance)
	return data
}

func (instance *Message) ToString() string {
	return string(instance.Marshal())
}

func (instance *Message) IsValid() bool {
	if nil != instance {
		if len(instance.UID) > 0 && nil != instance.Payload && len(instance.Payload.AppToken) > 0 {
			return instance.Payload.AppToken == application_types.AppToken
		}
	}
	return false
}

func (instance *Message) SetResponse(val interface{}) {
	if nil != instance {
		if v, b := val.(error); b {
			instance.Response = &MessageResponse{
				Error: v.Error(),
				Data:  nil,
			}
		} else {
			instance.Response = &MessageResponse{
				Data: lygo_conv.ToArray(val),
			}
		}
	}
}

func (instance *Message) GetPayloadParam(name string) interface{} {
	if nil != instance && nil != instance.Payload && nil != instance.Payload.Params {
		params := lygo_conv.ToMap(instance.Payload.Params)
		if v, b := params[name]; b {
			return v
		}
	}
	return nil
}

func (instance *Message) GetPayloadParamAsString(name string) string {
	if nil != instance && nil != instance.Payload && nil != instance.Payload.Params {
		return lygo_conv.ToString(instance.GetPayloadParam(name))
	}
	return ""
}

func (instance *Message) GetPayloadParamAsInt(name string) int {
	if nil != instance && nil != instance.Payload && nil != instance.Payload.Params {
		return lygo_conv.ToInt(instance.GetPayloadParam(name))
	}
	return 0
}

func (instance *Message) GetPayloadParamAsMap(name string) map[string]interface{} {
	if nil != instance && nil != instance.Payload && nil != instance.Payload.Params {
		return lygo_conv.ForceMap(instance.GetPayloadParam(name))
	}
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------
// 		p u b l i c
// ---------------------------------------------------------------------------------------------------------------------

func NewMessageHandler(events *lygo_events.Emitter) *MessageHandler {
	instance := new(MessageHandler)
	// instance.database = database
	instance.events = events

	return instance
}

func (instance *MessageHandler) Handle(payload interface{}) {
	if nil != payload {
		if v, b := payload.(*lygo_http_server.HttpWebsocketEventPayload); b {
			instance.handleWs(v)
		}
	}
}

// ---------------------------------------------------------------------------------------------------------------------
// 		p r i v a t e
// ---------------------------------------------------------------------------------------------------------------------

func (instance *MessageHandler) handleWs(payload *lygo_http_server.HttpWebsocketEventPayload) {
	if nil != instance && nil != instance.events {
		ws := payload.Websocket
		if nil != ws && ws.IsAlive() && len(payload.Message.Data) > 0 {
			var m Message
			err := json.Unmarshal(payload.Message.Data, &m)
			if nil == err && m.IsValid() {
				response := instance.execute(&m)
				m.SetResponse(response)
				if ws.IsAlive() {
					ws.SendData(m.Marshal())
				}
			} else {
				instance.events.Emit(application_types.EventError, application_types.ContextWebsocket, payload, err)
			}
		}
	}
}

func (instance *MessageHandler) execute(message *Message) interface{} {
	method := message.Payload.Namespace + "." + message.Payload.Function
	switch method {
	// ACCOUNT
	case "user.init":

	case "user.login":

	case "user.people":

	case "user.friends":

	case "user.activities":

	case "user.update_business_card":

	default:
		return errors.New("Unknown command: " + method)

	}
	return errors.New("Unknown command: " + method)
}
