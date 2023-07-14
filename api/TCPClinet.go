package api

import "github.com/praveenmahasena647/bg-client/internal/helpers"

func StartClient() error {
	return helpers.DialClient()
}
