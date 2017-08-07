/*
* MIT License
*
* Copyright (c) 2017 SmartestEE Co.,Ltd..
*
* Permission is hereby granted, free of charge, to any person obtaining a copy of
* this software and associated documentation files (the "Software"), to deal
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
* Revision History
*     Initial: 2017/08/07          Yusan Kurban
 */

package handler

import (
	"github.com/labstack/echo"

	"JCMS/general"
	"JCMS/general/errcode"
	"JCMS/model"
)

// Create create an article
func Create(c echo.Context) error {
	var (
		create model.Create
		err    error
	)

	if err = c.Bind(&create); err != nil {
		return general.NewErrorWithMessage(errcode.ErrInvalidParams, err.Error())
	}

	if err = c.Validate(create); err != nil {
		return general.NewErrorWithMessage(errcode.ErrInvalidParams, err.Error())
	}

	uid := c.Get().(string)
	err = model.ArticleServer.Create(*create, uid)
	if err != nil {
		return general.NewErrorWithMessage(errcode.ErrDBOperationFailed, err.Error())
	}

	return c.JSON(errcode.ErrSucceed, err.Error())
}
