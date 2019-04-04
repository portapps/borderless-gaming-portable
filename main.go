//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico -manifest=res/papp.manifest
package main

import (
	"os"

	. "github.com/portapps/portapps"
	"github.com/portapps/portapps/pkg/utl"
)

var (
	app *App
)

func init() {
	var err error

	// Init app
	if app, err = New("borderless-gaming-portable", "Borderless Gaming"); err != nil {
		Log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
	}
}

func main() {
	utl.CreateFolder(utl.PathJoin(app.DataPath, "AppData"))
	app.Process = utl.PathJoin(app.AppPath, "BorderlessGaming.exe")

	utl.OverrideEnv("APPDATA", utl.PathJoin(app.DataPath, "AppData"))

	app.Launch(os.Args[1:])
}
