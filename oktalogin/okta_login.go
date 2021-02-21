package oktalogin

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/bodytype"
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

type Result struct {
	URL        string            `json:url`
	Origin     string            `json:origin`
	Headers    map[string]string `json:headers`
	Status     string            `json:status`
	stateToken string            `json:stateToken`
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

	reqBody := `{"username":"` + profile.Username + `","password":"` + pass + `"}`
	cli := gentleman.New()
	cli.Use(body.String(reqBody))
	cli.Use(bodytype.Type("json"))
	res, err := cli.Request().Method("POST").URL(profile.Oktaurl + "/api/v1/authn").Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}

	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}
	//debug
	result := &Result{}
	// Parse the body and map into a struct
	res.JSON(result)
	fmt.Printf("Body: %#v\n", result)

	//	ioutil.WriteFile("big_marhsall.json", result, os.ModePerm)

	if result.Status == "MFA_REQUIRED" {
		fmt.Println("mfa required..")
	}
	fmt.Printf("Status: %d\n", res.StatusCode)
	//fmt.Printf("Body: %s", res.String())
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		// handle err
	}

}
