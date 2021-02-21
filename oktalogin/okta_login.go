package oktalogin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (co Credentials) MarshalJSON() ([]byte, error) {
	type credentials Credentials
	cn := credentials(co)
	cn.Password = "[REDACTED]"
	return json.Marshal((*credentials)(&cn))
}

func getPassword() string {
	fmt.Println("\nPassword: ")
	passwd, e := terminal.ReadPassword(int(os.Stdin.Fd()))
	if e != nil {
		log.Fatal(e)
	}
	return string(passwd)
}

func OktaLogin(profile_name string) {
	profile, err := GetProfile(profile_name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Login using username:", profile.Username, " and url ", profile.Oktaurl)
	pass := getPassword()

	data := Credentials{
		Username: profile.Username,
		Password: pass,
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", profile.Oktaurl+"/api/v1/authn", body)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	//debug
	jsonString, _ := json.Marshal(resp)
	ioutil.WriteFile("output.json", jsonString, os.ModePerm)

	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

}
