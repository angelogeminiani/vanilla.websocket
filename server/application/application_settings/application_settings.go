package application_settings

import (
	"github.com/botikasm/lygo/base/lygo_io"
	"github.com/botikasm/lygo/base/lygo_json"
	"github.com/botikasm/lygo/base/lygo_paths"
	"github.com/botikasm/lygo/ext/lygo_db/lygo_db_arango"
	"github.com/botikasm/lygo/ext/lygo_http/lygo_http_server/lygo_http_server_config"
	"github.com/cbroglie/mustache"
)

// ---------------------------------------------------------------------------------------------------------------------
//		t y p e
// ---------------------------------------------------------------------------------------------------------------------

type ApplicationSettings struct {
	Mode     string                                    `json:"mode"`
	Server   *lygo_http_server_config.HttpServerConfig `json:"server"`
	Database *lygo_db_arango.ArangoConfig              `json:"database"`
}

// ---------------------------------------------------------------------------------------------------------------------
//		c o n s t r u c t o r
// ---------------------------------------------------------------------------------------------------------------------

func NewApplicationSettings(mode string) (*ApplicationSettings, error) {
	path := lygo_paths.WorkspacePath("settings." + mode + ".json")
	settings := new(ApplicationSettings)
	text, err := lygo_io.ReadTextFromFile(path)
	if nil != err {
		return settings, err
	}
	context := make(map[string]interface{})
	context["workspace"] = lygo_paths.GetWorkspacePath()
	text, err = mustache.Render(text, context)
	if nil != err {
		return settings, err
	}
	err = lygo_json.Read(text, &settings)
	return settings, err
}

//----------------------------------------------------------------------------------------------------------------------
//	MeeteeseSettings
//----------------------------------------------------------------------------------------------------------------------

func (instance *ApplicationSettings) IsDebug() bool {
	return instance.Mode == "debug"
}
