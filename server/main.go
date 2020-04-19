package main

import (
	"errors"
	"fmt"
	"github.com/angelogeminiani/vanilla.websocket/server/application"
	"github.com/botikasm/lygo/base/lygo_conv"
	"github.com/botikasm/lygo/base/lygo_json"
	"github.com/botikasm/lygo/base/lygo_paths"
	"github.com/botikasm/lygo/base/lygo_strings"
	"github.com/botikasm/lygo/ext/lygo_logs"
)

// ---------------------------------------------------------------------------------------------------------------------
//		c o n s t
// ---------------------------------------------------------------------------------------------------------------------

const version = "1.0.0"
const name = "WEBSOCKET TEST SERVER"

const modeProduction = "production"
const modeDebug = "debug"

// TODO: Change this before deploy in production!!!!!!!!!!!!!
//--  REAM THIS ---//
const MODE = modeDebug

// ---------------------------------------------------------------------------------------------------------------------
//		l a u n c h e r
// ---------------------------------------------------------------------------------------------------------------------

func main() {

	// PANIC RECOVERY
	defer func() {
		if r := recover(); r != nil {
			// recovered from panic
			message := lygo_strings.Format("APPLICATION %s ERROR: %s", name, r)
			fmt.Println(message)
			lygo_logs.Error(message)
		}
	}()

	initWorkspace()
	initLogger(MODE)

	// log start
	message := lygo_strings.Format("APPLICATION %s v.%s mode '%s' START: %s", name, version, MODE, lygo_paths.GetWorkspacePath())
	lygo_logs.Info(message)
	fmt.Println(message)

	application, err := application.NewApplication(MODE)
	if nil != err {
		panic(err)
	}
	application.Start()

	err = application.Join()
	if nil != err {
		panic(err)
	}

	message = lygo_strings.Format("APPLICATION %s v.%s mode '%s' QUIT.", name, version, MODE)
	lygo_logs.Info(message)
	fmt.Println(message)
}

// ---------------------------------------------------------------------------------------------------------------------
//		p r i v a t e
// ---------------------------------------------------------------------------------------------------------------------

func initWorkspace() {
	m, err := lygo_json.ReadMapFromFile("./init.json")
	if nil != err {
		panic(err)
	}
	path := lygo_conv.ToString(m["workspace"])
	if len(path) == 0 {
		panic(errors.New("invalid_configuration"))
	}
	lygo_paths.SetWorkspacePath(lygo_paths.Absolute(path))
}

func initLogger(mode string) {
	if mode == modeDebug {
		lygo_logs.SetLevel(lygo_logs.LEVEL_DEBUG)
	} else {
		lygo_logs.SetLevel(lygo_logs.LEVEL_INFO)
	}
	lygo_logs.SetOutput(lygo_logs.OUTPUT_FILE)
}



