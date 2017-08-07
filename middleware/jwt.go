/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Co.,ltd..
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2017/08/07        Liu JiaChang
 */

package middleware


import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"

	"JCMS/config"
)

var urlMap map[string]struct{}

func CustomJWT() echo.MiddlewareFunc {
	jwtconf := middleware.JWTConfig {
		Skipper:     CustomSkipper,
		SigningKey:  []byte(config.Configuration.JwtKey),
	}

	return middleware.JWTWithConfig(jwtconf)
}

func CustomSkipper(c echo.Context) bool {
	if _, ok := urlMap[c.Request().RequestURI]; ok {
		return true
	}

	return false
}

func init() {
	urlMap = make(map[string]struct{})

	urlMap["/login"] = struct{}{}
	urlMap["/register"] = struct{}{}
}

func MustLoginIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		strUid := claims[config.Configuration.JwtUid].(string)

		c.Set(config.Configuration.JwtUid, strUid)

		return next(c)
	}
}
