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

package model

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"JCMS/mongo"
)

type ArticleServerProvider struct{}

var (
	ArticleServer *ArticleServerProvider
	RefArticle    *mgo.Collection
)

// PrepareArticle is init mongo connection.
func PrepareArticle() {
	RefArticle = mongo.MDSession.DB(mongo.MDJCMS).C(mongo.Article)
	nameIndex := mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}
	if err := RefArticle.EnsureIndex(nameIndex); err != nil {
		panic(err)
	}

	ArticleServer = &ArticleServerProvider{}
}

// Article contain all article information
type Article struct {
	ID      bson.ObjectId `bson:"_id" json:"id"`
	Title   string        `json:"title"`
	Content string        `json:"content"`
	Author  string        `json:"author"`
	Images  []string      `json:"images"`
	Created time.Time     `json:"created"`
}

// Create use for create an article
type Create struct {
	Title   string   `json:"title"   validate:"gte=6,lte=20"`
	Content string   `json:"content" validate:"gte=6,lte=500"`
	Images  []string `json:"images"`
}

// Create create an article and return nil if it success
func (as *ArticleServerProvider) Create(c *Create, author string) error {
	article := Article{
		ID:      bson.NewObjectId(),
		Title:   c.Title,
		Content: c.Content,
		Author:  author,
		Images:  c.Images,
		Created: time.Now(),
	}

	return RefArticle.Insert(&article)
}

// GetArticleByID return an Article{} and nil if id article exists
func (as *ArticleServerProvider) GetArticleByID(id bson.ObjectId) (Article, error) {
	var article Article

	err := RefArticle.FindId(id).One(&article)

	return article, err
}

//GetArticleByAuthor return []Article{} and nil if bd has any article written by authorID
func (as *ArticleServerProvider) GetArticleByAuthor(authorID string) ([]Article, error) {
	var article []Article

	err := RefArticle.Find(bson.M{"author": authorID}).All(&article)

	return article, err
}

//GetAll return []Article{} and nil id db has any article
func (as *ArticleServerProvider) GetAll() ([]Article, error) {
	var article []Article

	err := RefArticle.Find(bson.M{}).All(&article)

	return article, err
}
