package main

import (
	"github.com/NuttayotSukkum/user-profile/httpserv"
	infrastrcture "github.com/NuttayotSukkum/user-profile/infrastructure"
	"github.com/spf13/viper"
)

func init() {
	infrastrcture.InitConfig()
}

func main() {

	infrastrcture.InitLogger(viper.GetString("app.env"))
	defer infrastrcture.Log.Sync()

	infrastrcture.InitDb()
	httpserv.Run()

}
