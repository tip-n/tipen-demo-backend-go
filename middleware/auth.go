package middleware

import "tipen-demo/pkg"

func (m *Middleware) IsAuthorized(jwt string) error {
	return pkg.ValidateJWT(jwt)
}
