package application_runtime

import (
	"github.com/angelogeminiani/vanilla.websocket/server/application/application_settings"
	"github.com/angelogeminiani/vanilla.websocket/server/application/application_types"
	"github.com/angelogeminiani/vanilla.websocket/server/application/application_www"
	"github.com/botikasm/lygo/base/lygo_events"
)

// ---------------------------------------------------------------------------------------------------------------------
//		t y p e
// ---------------------------------------------------------------------------------------------------------------------

type ApplicationRuntime struct {

	//-- private --//
	settings *application_settings.ApplicationSettings

	// internal controllers
	events   *lygo_events.Emitter
	server   *application_www.ApplicationServer
}

// ---------------------------------------------------------------------------------------------------------------------
//		c o n s t r u c t o r
// ---------------------------------------------------------------------------------------------------------------------

func NewApplicationRuntime(settings *application_settings.ApplicationSettings) (*ApplicationRuntime, error) {
	var err error
	instance := new(ApplicationRuntime)
	instance.events = lygo_events.NewEmitter()
	instance.settings = settings
	if nil != settings && nil != settings.Server {

		// application server
		instance.server = application_www.NewApplicationServer(
			instance.events, instance.settings.Server)

	} else {
		err = application_types.MismatchConfigurationError
	}
	return instance, err
}

// ---------------------------------------------------------------------------------------------------------------------
//		p u b l i c
// ---------------------------------------------------------------------------------------------------------------------

func (instance *ApplicationRuntime) Start() error{
	if nil != instance {
		// start application server
		return instance.server.Start()
	}
	return application_types.PanicSystemError
}

func (instance *ApplicationRuntime) Close() error {
	if nil != instance && nil != instance.server {
		var err error
		if nil != instance.server {
			err = instance.server.Close()
			instance.server = nil
		}
		return err
	}
	return application_types.PanicSystemError
}

func (instance *ApplicationRuntime) Join() error {
	if nil != instance {
		if nil != instance.server {
			instance.server.Join()
		}
		return nil
	}
	return application_types.PanicSystemError
}


// ---------------------------------------------------------------------------------------------------------------------
//		p r i v a t e
// ---------------------------------------------------------------------------------------------------------------------
