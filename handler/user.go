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

	"GoTemp/model"
	"GoTemp/general"
)

func Login(c echo.Context) error {
	var User struct {
		Name  string `json:"username"  validate:"get=6, lte=15"`
		Pass  string `json:"password"  validate:"gte=6, lte=30"`
	}

	if err := c.Bind(&User); err != nil {
		return general.NewErrorWithMessage(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(User); err != nil {
		return general.NewErrorWithMessage(http.StatusBadRequest, err.Error())
	}

	err := model.UserService.Login(User.Name, User.Pass)

	if err != nil {
		return general.NewErrorWithMessage(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}
