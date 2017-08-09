/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Co.,Ltd..
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
 *     Initial: 2017/07/21        Liu JiaChang
 */

package handler

import (
	"net/http"

	"github.com/labstack/echo"
	jwt "github.com/dgrijalva/jwt-go"

	"JCMS/model"
	"JCMS/general"
	"JCMS/general/errcode"
	"JCMS/utility"
	"JCMS/config"
	"JCMS/log"
)

// Login user login
func Login(c echo.Context) error {
	var (
		User        model.User
		tokenStr    string
	)

	if err := c.Bind(&User); err != nil {
		return general.NewErrorWithMessage(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(User); err != nil {
		return general.NewErrorWithMessage(http.StatusBadRequest, err.Error())
	}

	user, err := model.UserService.Login(User.UserName)

	if err != nil {
		return general.NewErrorWithMessage(http.StatusInternalServerError, err.Error())
	}

	if !util.CompareHash([]byte(user.PassWord), User.PassWord) {
		return general.NewErrorWithMessage(errcode.ErrInvalidParams, errcode.ErrLoginPassErr)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims[config.Configuration.JwtUid] = user.UserId.Hex()

	if tokenStr, err = token.SignedString([]byte(config.Configuration.JwtKey)); err != nil {
		log.Logger.Error("Signe string error:", err)

		return general.NewErrorWithMessage(errcode.ErrJwtSignString, err.Error())
	}

	return c.JSON(http.StatusOK, tokenStr)
}

// Register user register
func Register(c echo.Context) error {
	var (
		User        model.User
		userId      string
		tokenStr    string
		err         error
	)

	if err := c.Bind(&User); err != nil {
		return general.NewErrorWithMessage(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(User); err != nil {
		return general.NewErrorWithMessage(http.StatusBadRequest, err.Error())
	}

	if userId, err = model.UserService.Register(User.UserName, User.PassWord); err != nil {
		return general.NewErrorWithMessage(http.StatusNotAcceptable, err.Error())
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims[config.Configuration.JwtUid] = userId

	if tokenStr, err = token.SignedString([]byte(config.Configuration.JwtKey)); err != nil {
		log.Logger.Error("Signe string error:", err)

		return general.NewErrorWithMessage(errcode.ErrJwtSignString, err.Error())
	}

	return c.JSON(http.StatusOK, tokenStr)
}
