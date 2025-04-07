// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package captcha

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/mojocn/base64Captcha"
)

// configJSONBody json request body.
type configJSONBody struct {
	ID            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var store = base64Captcha.DefaultMemStore

// base64Captcha create http handler
func generateCaptchaHandler(w http.ResponseWriter, r *http.Request) {
	// parse request parameters
	decoder := json.NewDecoder(r.Body)
	var param configJSONBody
	err := decoder.Decode(&param)
	if err != nil {
		log.Println(err)
	}
	defer func() {
		_ = r.Body.Close()
	}()
	var driver base64Captcha.Driver

	// create base64 encoding captcha
	switch param.CaptchaType {
	case "audio":
		driver = param.DriverAudio
	case "string":
		driver = param.DriverString.ConvertFonts()
	case "math":
		driver = param.DriverMath.ConvertFonts()
	case "chinese":
		driver = param.DriverChinese.ConvertFonts()
	default:
		driver = param.DriverDigit
	}
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, answer, err := c.Generate()
	body := map[string]interface{}{"code": 1, "data": b64s, "captchaId": id, "msg": "success", "answer": answer}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(body)
}

// base64Captcha verify http handler
func captchaVerifyHandle(w http.ResponseWriter, r *http.Request) {
	// parse request json body
	decoder := json.NewDecoder(r.Body)
	var param configJSONBody
	err := decoder.Decode(&param)
	if err != nil {
		log.Println(err)
	}
	defer func() {
		_ = r.Body.Close()
	}()
	// verify the captcha
	body := map[string]interface{}{"code": 0, "msg": "failed"}
	if store.Verify(param.ID, param.VerifyValue, true) {
		body = map[string]interface{}{"code": 1, "msg": "ok"}
	}

	// set json response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	_ = json.NewEncoder(w).Encode(body)
}

// TestCaptcha start a net/http server
func TestCaptcha(t *testing.T) {
	t.Log("start a net/http server")
	// serve Vuejs+ElementUI+Axios Web Application
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// api for create captcha
	http.HandleFunc("/api/getCaptcha", generateCaptchaHandler)

	// api for verify captcha
	http.HandleFunc("/api/verifyCaptcha", captchaVerifyHandle)

	if err := http.ListenAndServe(":8777", nil); err != nil {
		log.Fatal(err)
	}
	t.Log("end a net/http server")
}
