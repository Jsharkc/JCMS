/*
 * MIT License
 *
 * Copyright (c) 2017 九光年.
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
 *     Initial: 2017/07/21        九光年
 */

package model

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"GoTemp/mongo"
)

type UserProvider struct{}

type User struct {
	UserId       bson.ObjectId   `bson:"_id,omitempty"  json:"id"`
	UserName     string          `bson:"username"       json:"username" validate:"required"  validate:"gte=6, lte=15"`
	PassWord     string          `bson:"password"       json:"password" validate:"required"  validate:" gte=6, lte=20"`
}

var (
	UserService *UserProvider

	// Collection
	RefUser *mgo.Collection

	UserError = errors.New("user not exist")
	PassError = errors.New("password is error")
)

func PrepareUser() {
	RefUser = mongo.MDSession.DB(mongo.MDGoTemp).C(mongo.User)
	UserService = &UserProvider{}
}

func(this *UserProvider) Login(username, password string ) error {
	var pass User

	if mongo.GetUniqueOne(RefUser, bson.M{"username": username}, &pass) != nil{
		return UserError
	}
	
	if pass.PassWord != password {
		return  PassError
	}

	return nil
}