//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico -manifest=res/papp.manifest
package main

import (
	"os"

	"github.com/portapps/portapps/v3"
	"github.com/portapps/portapps/v3/pkg/log"
	"github.com/portapps/portapps/v3/pkg/utl"
)

var (
	app *portapps.App
)

func init() {
	var err error

	// Init app
	if app, err = portapps.New("borderless-gaming-portable", "Borderless Gaming"); err != nil {
		log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
	}
}

func main() {
	utl.CreateFolder(utl.PathJoin(app.DataPath, "AppData"))
	app.Process = utl.PathJoin(app.AppPath, "BorderlessGaming.exe")

	os.Setenv("APPDATA", utl.PathJoin(app.DataPath, "AppData"))

	defer app.Close()
	app.Launch(os.Args[1:])
}
