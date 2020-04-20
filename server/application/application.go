package application

import (
	"github.com/angelogeminiani/vanilla.websocket/server/application/application_runtime"
	"github.com/angelogeminiani/vanilla.websocket/server/application/application_settings"
	"github.com/angelogeminiani/vanilla.websocket/server/application/application_types"
	"github.com/botikasm/lygo/base/lygo_events"
)

// ---------------------------------------------------------------------------------------------------------------------
//		t y p e
// ---------------------------------------------------------------------------------------------------------------------

type Application struct {

	//-- private --//
	settings *application_settings.ApplicationSettings

	// internal controllers
	events  *lygo_events.Emitter
	runtime *application_runtime.ApplicationRuntime
	// database *application_settings.ApplicationDatabase
	// server   *application_settings.ApplicationServer
}

// ---------------------------------------------------------------------------------------------------------------------
//		c o n s t r u c t o r
// ---------------------------------------------------------------------------------------------------------------------

func NewApplication(mode string) (*Application, error) {
	instance := new(Application)
	instance.events = lygo_events.NewEmitter()
	settings, err := application_settings.NewApplicationSettings(mode)
	if nil == err {
		instance.settings = settings
		if nil != settings {
			instance.runtime, err = application_runtime.NewApplicationRuntime(settings)
		} else {
			err = application_types.MismatchConfigurationError
		}
	}
	return instance, err
}

// ---------------------------------------------------------------------------------------------------------------------
//		p u b l i c
// ---------------------------------------------------------------------------------------------------------------------

func (instance *Application) Start() error {
	if nil != instance && nil != instance.runtime {
		// start application server
		return instance.runtime.Start()
	}
	return application_types.PanicSystemError
}

func (instance *Application) Close() error {
	if nil != instance && nil != instance.runtime {
		return instance.runtime.Close()
	}
	return application_types.PanicSystemError
}

func (instance *Application) Join() error {
	if nil != instance && nil != instance.runtime {
		return instance.runtime.Join()
	}
	return application_types.PanicSystemError
}

// ---------------------------------------------------------------------------------------------------------------------
// 		p r i v a t e
// ---------------------------------------------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------------------------------------------
//  	S T A T I C
// ---------------------------------------------------------------------------------------------------------------------
