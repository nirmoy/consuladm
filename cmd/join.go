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

var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "Join an existing consul cluster",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var flagArgs config.Flags
		addr := args[0]

		flags := flag.NewFlagSet("", flag.ContinueOnError)
		config.AddFlags(flags, &flagArgs)

		flagArgs.Config.AdvertiseAddrLAN = &consulAdmConfig.Advertise
		flagArgs.Config.DataDir = &consulAdmConfig.DataDir
		consulAdmConfig.ServerMode = false
		flagArgs.Config.ServerMode = &consulAdmConfig.ServerMode
		flagArgs.Config.NodeName = &consulAdmConfig.Name

		consulAdmConfig.DataCenter = consul.GetMemberDC(addr + ":" + constants.DefaultHttpPort)
		flagArgs.Config.Datacenter = &consulAdmConfig.DataCenter
		consul.AgentRun(flagArgs)
		consul.AgentJoin(addr)
		fmt.Print("consulAdmConfig: [join] init was successful, looping for ever\n")
		for {
			time.Sleep(time.Second)
		}

	},
}

func init() {
	rootCmd.AddCommand(joinCmd)
	joinCmd.PersistentFlags().StringVar(&consulAdmConfig.Name, "name", "", "consul node name")
	joinCmd.PersistentFlags().StringVar(&consulAdmConfig.Advertise, "advertise", "", "Advertise address")
	joinCmd.PersistentFlags().StringVar(&consulAdmConfig.DataCenter, "datacenter", constants.DefaultDataCenter, "Advertise address")

}
