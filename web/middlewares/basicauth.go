// file: middleware/basicauth.go

package middleware

import (
	"time"

	"github.com/kataras/iris/middleware/basicauth"
)

var BasicAuth = basicauth.New(basicauth.Config{
	Users:   map[string]string{"agusjim": "123.", "alfre": "123"},
	Realm:   "Authorization Required", // defaults to "Authorization Required"
	Expires: time.Duration(30) * time.Minute,
},
)
