package proxy

// TODO: Probably can just use regular Caddy via Docker with no custom code, but some custom modules could come in handy later
// https://caddyserver.com/docs/extending-caddy
// https://github.com/caddyserver/xcaddy

import "github.com/caddyserver/caddy/v2"

func init() {
	caddy.RegisterModule(Gizmo{})
}

// Gizmo is an example; put your own type here.
type Gizmo struct {
}

// CaddyModule returns the Caddy module information.
func (Gizmo) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "foo.gizmo",
		New: func() caddy.Module { return new(Gizmo) },
	}
}
