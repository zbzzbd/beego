package notify

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"

	"utils"
)

var (
	apiUser            string
	apiKey             string
	from               string
	sendEmailByTempUrl string
)

func getParams() {
	apiUser = utils.GetConf().String("email::api_user")
	apiKey = utils.GetConf().String("email::api_key")
	from = utils.GetConf().String("email::from")
	sendEmailByTempUrl = utils.GetConf().String("email::email_tmp_url")
}

type Email struct {
	notifier EmailNotifier
	data     url.Values
}

var once sync.Once

func NewEmail(en EmailNotifier) *Email {
	once.Do(getParams)

	email := new(Email)

	email.notifier = en
	email.data = url.Values{
		"api_user": {apiUser},
		"api_key":  {apiKey},
		"from":     {from},
		"fromname": {from},
	}

	return email
}

func (e *Email) IsEnable() bool {
	isEnable, _ := utils.GetConf().Bool("email::enable")
	return isEnable
}

// make email sending as a queue, never block the progress
func (e *Email) Send() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("notify.Email.Send err=%s", err.Error())
		} else {
			utils.GetLog().Debug("notify.Email.Send obj=%s", utils.Sdump(e))
		}
	}()

	if !e.IsEnable() {
		err = errors.New("email is not enabled")
		return nil
	}

	e.data.Set("template_invoke_name", string(e.notifier.GetTemplateName()))

	var subs string
	if subs, err = e.handleSubs(); err != nil {
		return
	}
	e.data.Set("substitution_vars", subs)

	postData := e.data.Encode()
	postBody := bytes.NewBufferString(postData)

	resp, err := httpSend(sendEmailByTempUrl, postBody)
	if err != nil {
		return
	}
	var result map[string]interface{}
	err = json.Unmarshal(resp, &result)
	if result["message"] != "success" || err != nil {
		return err
	}
	return
}

func (e Email) formatTemplateKey(key string) string {
	return "%" + key + "%"
}

func (e *Email) handleSubs() (string, error) {
	vars := e.notifier.GetVariables()
	subs := make(map[string][]interface{})

	for _, val := range vars {
		for k, v := range val {
			formattedKey := e.formatTemplateKey(k)
			subs[formattedKey] = append(subs[formattedKey], v)
		}
	}
	combine := map[string]interface{}{
		"to":  e.notifier.GetTo(),
		"sub": subs,
	}
	combineBytes, err := json.Marshal(combine)
	if err != nil {
		return "", err
	}

	return string(combineBytes), nil
}

func httpSend(url string, body io.Reader) (result []byte, err error) {
	responseHandler, err := http.Post(url, "application/x-www-form-urlencoded", body)
	if err != nil {
		utils.GetLog().Error("httpSend error : %s", err.Error())
		return nil, err

	}
	defer responseHandler.Body.Close()
	bodyByte, err := ioutil.ReadAll(responseHandler.Body)
	if err != nil {
		return nil, err
	}
	return bodyByte, nil
}
