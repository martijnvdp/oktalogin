/*
Copyright © 2021 M van der Ploeg

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/martijnxd/oktalogin/oktalogin"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

type Profiledata struct {
	Name     string `mapstructure:"name:"`
	Username string `mapstructure:"username"`
	Oktaurl  string `mapstructure:"oktaurl"`
}

type Oktaprofiles struct {
	Profiles []Profiledata `mapstructure:"profiles"`
}

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "oktalogin",
	Short: "A cli app to login to okta and aws",
	Long: `A cli app to login to okta and updte tokens for aws cli

Oktalogin is a cli app to login to okta from the command line and get access tokens for aws to be used for other cli apps like the awscli`,
	Run: func(cmd *cobra.Command, args []string) {
		profile, _ := cmd.Flags().GetString("profiles")
		oktalogin.OktaLogin(profile)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.oktalogin.yaml)")
	rootCmd.Flags().StringP("profiles", "p", "", "Specify oktalogin profile to use for login")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	var okta_profiles Oktaprofiles
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		configName := ".oktalogin"
		configType := "yaml"
		viper.AddConfigPath(home)
		viper.SetConfigName(configName)
		viper.SetConfigType(configType)
		cfgFile = filepath.Join(home, configName+"."+configType)
	}
	viper.AllowEmptyEnv(true)
	viper.Unmarshal(&okta_profiles)
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		_, err := os.Stat(cfgFile)
		if !os.IsExist(err) {
			if _, err := os.Create(cfgFile); err != nil {
			}
		}
		if err := viper.SafeWriteConfig(); err != nil {
		}
	}
}
