// file: middleware/basicauth.go

package middleware

import (
	"time"

	"github.com/kataras/iris/middleware/basicauth"
)

// BasicAuth middleware sample.
var BasicAuth = basicauth.New(basicauth.Config{
	Users:   map[string]string{"agusjim": "Icui4culoc.", "mySecondusername": "mySecondpassword"},
	Realm:   "Authorization Required", // defaults to "Authorization Required"
	Expires: time.Duration(30) * time.Minute,
},
)
