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

package mongo

import (
	"time"

	"gopkg.in/mgo.v2"

	"JCMS/log"
)

var (
	MDSession *mgo.Session
)

const (
	MDJCMS  = "JCMS"
	User    = "user"
	Article = "article"
)

// 初始化 MongoDB 连接、文档类型初始化
func InitUserMD(url string) {
	var err error
	MDSession, err = mgo.DialWithTimeout(url+"/"+MDJCMS, time.Second)

	if err != nil {
		panic(err)
	}

	log.Logger.Debug("the MongoDB of %s connected!", MDJCMS)

	// 尽可能利用 MongoDB 分布性特性
	MDSession.SetMode(mgo.Monotonic, true)
}
