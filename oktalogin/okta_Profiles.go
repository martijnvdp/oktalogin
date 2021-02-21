package oktalogin

import (
	"fmt"

	"github.com/spf13/viper"
)

type Profiledata struct {
	Name     string `mapstructure:"name"`
	Username string `mapstructure:"username"`
	Oktaurl  string `mapstructure:"oktaurl"`
}

type Oktaprofiles struct {
	Profiles []Profiledata `mapstructure:"profiles"`
}

func ListProfiles() {
	var profiles []Profiledata
	viper.UnmarshalKey("profiles.Profiles", &profiles)
	fmt.Println("List existing oktalogin profiles:")
	for _, p := range profiles {
		fmt.Println(" name: ", p.Name)
		fmt.Println(" login e-mail: ", p.Username)
		fmt.Println(" Okta url: ", p.Oktaurl)
		fmt.Println("---")
	}
}

func GetProfile(name string) (profile *Profiledata, err error) {
	var profiles []Profiledata
	viper.UnmarshalKey("profiles.Profiles", &profiles)
	for _, p := range profiles {
		if p.Name == name {
			profile = &p
		}
	}

	return profile, err
}

func FindProfile(name string) (exist bool) {
	var profiles []Profiledata
	viper.UnmarshalKey("profiles.Profiles", &profiles)
	for _, p := range profiles {
		if p.Name == name {
			exist = true
		}
	}
	return exist
}

func AddProfile() *Profiledata {
	var profile Profiledata
	fmt.Println("add user")
	fmt.Println("Enter alias for this user: ")
	_, err := fmt.Scanln(&profile.Name)
	fmt.Println("Enter username/e-mail: ")
	_, err = fmt.Scanln(&profile.Username)
	fmt.Println("Enter okta url: ")
	_, err = fmt.Scanln(&profile.Oktaurl)
	if err != nil {
		fmt.Printf("%v", err)
	}
	return &profile
}

func AddProfiles() {
	var okta_profiles Oktaprofiles
	var err error
	var input string
	var profiles []Profiledata
	viper.UnmarshalKey("profiles.Profiles", &profiles)

	okta_profiles.Profiles = append(okta_profiles.Profiles, profiles...)
	for addanother := true; addanother != false; {
		prof := AddProfile()
		if !FindProfile(prof.Name) {
			okta_profiles.Profiles = append(okta_profiles.Profiles, *prof)
		} else {
			fmt.Println("Profile already exists")
		}
		println("Add another profile(yes/no)")
		fmt.Scan(&input)
		fmt.Scanln()
		if input != "yes" && input != "y" {
			addanother = false
		}
	}
	viper.Set("profiles", okta_profiles)
	viper.WriteConfig()

	if err != nil {
		error.Error(err)
	}
}
