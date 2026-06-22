// Package plugin is a starter togo plugin shaped like a mini app: a
// self-registering backend provider, an injectable frontend (web/), its own
// internal packages, and a .claude/ agent kit. Copy this repo or run
// `togo make:plugin <name>` to scaffold a new one in the same shape.
//
// On `togo install <owner>/<repo>` the package is blank-imported into the app, so
// init() registers it with the kernel — no manual wiring.
package plugin

import (
	"github.com/togo-framework/togo"

	"github.com/togo-framework/plugin-template/internal/example"
)

// Name is the plugin's stable identifier.
const Name = "example"

func init() {
	togo.RegisterProviderFunc(Name, togo.PriorityLate, func(k *togo.Kernel) error {
		// Mount backend routes on the kernel router.
		svc := example.New(k)
		k.Router.Get("/api/example/ping", svc.Ping)
		// Expose the service to the rest of the app via the kernel container.
		k.Set(Name, svc)
		if k.Log != nil {
			k.Log.Info("plugin active", "plugin", Name)
		}
		return nil
	})
}
