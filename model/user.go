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

package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"JCMS/mongo"
	"JCMS/utility"
)

type UserProvider struct{}

var (
	UserService *UserProvider

	// Collection
	RefUser *mgo.Collection
)

func PrepareUser() {
	RefUser = mongo.MDSession.DB(mongo.MDJCMS).C(mongo.User)
	nameIndex := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}

	if err := RefUser.EnsureIndex(nameIndex); err != nil {
		panic(err)
	}

	UserService = &UserProvider{}
}

type User struct {
	UserId       bson.ObjectId   `bson:"_id,omitempty"  json:"id"`
	UserName     string          `bson:"name"           json:"username" validate:"required"  validate:"gte=6, lte=20"`
	PassWord     string          `bson:"pass"           json:"password" validate:"required"  validate:"gte=6, lte=20"`
}

func(this *UserProvider) Login(username string) (*User, error) {
	var user User

	mongo.MDSession.Refresh()

	err := RefUser.Find(bson.M{"name": username}).One(&user)

	return &user, err
}

func(this *UserProvider) Register(username, password string) (string, error) {
	result, err := util.GenerateHash(password)
	if err != nil {
		return "", err
	}

	pass := string(result)
	create := User{
		UserId:     bson.NewObjectId(),
		UserName:   username,
		PassWord:   pass,
	}

	mongo.MDSession.Refresh()
	err = RefUser.Insert(&create)
	if err != nil {
		return "", err
	}

	return create.UserId.Hex(), nil
}
