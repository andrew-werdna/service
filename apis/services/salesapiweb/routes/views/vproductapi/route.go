package vproductapi

import (
	"net/http"

	"github.com/ardanlabs/service/app/core/views/vproductapp"
	"github.com/ardanlabs/service/business/api/auth"
	midhttp "github.com/ardanlabs/service/business/api/mid/http"
	"github.com/ardanlabs/service/business/core/views/vproduct"
	"github.com/ardanlabs/service/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	VProductCore *vproduct.Core
	Auth         *auth.Auth
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := midhttp.Authenticate(cfg.Auth)
	ruleAdmin := midhttp.Authorize(cfg.Auth, auth.RuleAdminOnly)

	api := newAPI(vproductapp.NewCore(cfg.VProductCore))
	app.Handle(http.MethodGet, version, "/vproducts", api.query, authen, ruleAdmin)
}