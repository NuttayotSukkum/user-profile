package httpserv

import (
	"log"

	Log "github.com/NuttayotSukkum/user-profile/infrastructure"
	app "github.com/NuttayotSukkum/user-profile/infrastructure"
	"github.com/spf13/viper"
)

func Run() {

	a := app.NewApp()

	bindHealthRouter(a)
	bindCustomerProfileRouteCreate(a)

	name := viper.GetString("app.name")
	port := viper.GetString("app.port")

	if port == "" {
		port = "1323"
	}

	Log.Infof("Starting %s on port %s", name, port)

	if err := a.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
