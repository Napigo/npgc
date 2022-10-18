package rest

import (
	"github.com/Napigo/npglogger"
	"github.com/gofiber/fiber/v2"
)

func CreateRestHook(app *fiber.App) {
	app.Hooks().OnGroup(func(g fiber.Group) error {
		npglogger.Info("Group route added, path :" + g.Prefix)
		return nil
	})

	app.Hooks().OnRoute(func(r fiber.Route) error {
		npglogger.Info("Route added, path : " + r.Path + " method: " + r.Method)
		return nil
	})
}
