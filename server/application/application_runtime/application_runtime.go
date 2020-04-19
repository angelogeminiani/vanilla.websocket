package application_runtime

import (
	"github.com/angelogeminiani/vanilla.websocket/server/application/application_settings"
	"time"
)

// ---------------------------------------------------------------------------------------------------------------------
//		t y p e
// ---------------------------------------------------------------------------------------------------------------------

type ApplicationRuntime struct {

	//-- private --//
	settings *application_settings.ApplicationSettings
}

// ---------------------------------------------------------------------------------------------------------------------
//		c o n s t r u c t o r
// ---------------------------------------------------------------------------------------------------------------------

func NewApplicationRuntime(settings *application_settings.ApplicationSettings) *ApplicationRuntime {
	instance := new(ApplicationRuntime)
	instance.settings = settings

	return instance
}

// ---------------------------------------------------------------------------------------------------------------------
//		p u b l i c
// ---------------------------------------------------------------------------------------------------------------------

func (instance *ApplicationRuntime) Start() {
	time.Sleep(10*time.Second)
}

// ---------------------------------------------------------------------------------------------------------------------
//		p r i v a t e
// ---------------------------------------------------------------------------------------------------------------------