package tickstack

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
)

type Response struct {
	Enabled bool
	From string
	Global bool
	Host string
	IdleTimeout string
	NoVerify bool
	Port int
	StateChangesOnly bool
	To []string
}

type SmtpResp struct {
	Link struct {
		Rel string `json:"rel"`
		Href string `json:"href"`
	} `json:"link"`
	Elements []struct {
		Link struct {
			Rel string `json:"rel"`
			Href string `json:"href"`
		} `json:"link"`
		Options struct {
			Enabled bool `json:"enabled"`
			From string `json:"from"`
			Global bool `json:"global"`
			Host string `json:"host"`
			IdleTimeout string `json:"idle-timeout"`
			NoVerify bool `json:"no-verify"`
			Password bool `json:"password"`
			Port int `json:"port"`
			StateChangesOnly bool `json:"state-changes-only"`
			To []string `json:"to"`
			Username string `json:"username"`
		} `json:"options"`
		Redacted []string `json:"redacted"`
	} `json:"elements"`
}

func GetSmtp(host, port string) (*Response, error) {
	var response Response
	url := "http://" + host + ":" + port + "/kapacitor/v1/config/smtp"

	res, err := http.Get(url)
	if err != nil {
		log.Println("GET smtp failed")
		return &response, err
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Println("Read smtp failed")
		return &response, err
	}

	var resp SmtpResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Println("Parse smtp failed")
		return &response, err
	}

	option := resp.Elements[0].Options
	response.Enabled = option.Enabled
	response.From = option.From
	response.Global = option.Global
	response.Host = option.Host
	response.IdleTimeout = option.IdleTimeout
	response.NoVerify = option.NoVerify
	response.Port = option.Port
	response.StateChangesOnly = option.StateChangesOnly
	response.To = option.To
	return &response, nil
}

