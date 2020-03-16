/*
Copyright © 2020 reeve0930 <reeve0930@gmail.com>

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
	"log"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type config struct {
	fHost string
	fPath string
	cDest string
}

func (c config) String() string {
	return fmt.Sprintf("foltia_host = \"%s\"\nfoltia_path = \"%s\"\ncopy_dest = \"%s\"",
		c.fHost,
		c.fPath,
		c.cDest,
	)
}

var (
	conf       config
	configPath string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "foltia",
	Short: "A Command Line Tool for foltia ANIME LOCKER",
	Long:  `This is a command line tool for foltia ANIME LOCKER. This tool is made by reeve0930.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
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
}

func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	configPath = filepath.Join(home, ".config", "foltia")
	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		os.Mkdir(configPath, 0777)
	}

	viper.AddConfigPath(configPath)
	viper.SetConfigName("config.toml")
	viper.SetConfigType("toml")

	configPath = filepath.Join(configPath, "config.toml")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	conf.fHost = viper.GetString("foltia_host")
	conf.fPath = viper.GetString("foltia_path")
	conf.cDest = viper.GetString("copy_dest")
}
