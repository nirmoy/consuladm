/**
 *   Copyright 2019 Nirmoy Das.
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

package cmd

import (
	"flag"
	"fmt"
	"time"

	"github.com/hashicorp/consul/agent/config"

	"github.com/nirmoy/consuladm/constants"
	"github.com/nirmoy/consuladm/consul"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new consul agent",
	Run: func(cmd *cobra.Command, args []string) {
		var flagArgs config.Flags
		flags := flag.NewFlagSet("", flag.ContinueOnError)
		config.AddFlags(flags, &flagArgs)

		flagArgs.Config.AdvertiseAddrLAN = &consulAdmConfig.Advertise
		flagArgs.Config.DataDir = &consulAdmConfig.DataDir
		flagArgs.Config.Datacenter = &consulAdmConfig.DataCenter
		flagArgs.Config.ServerMode = &consulAdmConfig.ServerMode
		flagArgs.Config.Bootstrap = &consulAdmConfig.Bootstrap
		flagArgs.Config.ClientAddr = &consulAdmConfig.ClientAddr
		flagArgs.Config.NodeName = &consulAdmConfig.Name
		if err := consul.AgentRun(flagArgs); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("consulAdmConfig: init was successful, looping for ever\n")
		for {
			time.Sleep(time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().StringVar(&consulAdmConfig.Name, "name", "", "consul node name")
	initCmd.PersistentFlags().StringVar(&consulAdmConfig.Advertise, "advertise", "", "Advertise address")
	initCmd.PersistentFlags().StringVar(&consulAdmConfig.DataCenter, "datacenter", constants.DefaultDataCenter, "Advertise address")

}
