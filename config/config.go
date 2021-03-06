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

package config

import (
	"github.com/spf13/viper"
)

type workServerConfig struct {
	Address			        string
	IsDebug			        bool
	MgoUrl                  string
	JwtUid                  string
	JwtExp                  int
	JwtKey                  string
}

var (
	Configuration *workServerConfig
)

func ReadConfiguration() {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	Configuration = &workServerConfig{
		Address:		        viper.GetString("server.address"),
		IsDebug:		        viper.GetBool("server.debug"),
		MgoUrl:                 viper.GetString("mongodb.url"),
		JwtUid:                 viper.GetString("jwt.uid"),
		JwtExp:                 viper.GetInt("jwt.exp"),
		JwtKey:                 viper.GetString("jwt.tokenkey"),
	}
}
