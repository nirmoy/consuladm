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

package apis

import (
	"fmt"
	"os"

	"github.com/nirmoy/consuladm/constants"
	"github.com/nirmoy/consuladm/pkg/netutils"
)

// ConsulAdmConfig holds consuladm configuration
type ConsulAdmConfig struct {
	Version         string
	CertificatesDir string
	Advertise       string
	Name            string
	DataDir         string
	DataCenter      string
	ServerMode      bool
	ClientAddr      string
}

func DefaultAdvertise(cfg *ConsulAdmConfig) error {
	firstAddress, err := netutils.FirstGlobalV4Addr("")
	if err != nil {
		return fmt.Errorf("failed to set default Advertise: %s", err)
	}

	fmt.Printf("consuladm: advertise IP %s\n", firstAddress)
	cfg.Advertise = firstAddress.String()
	return nil
}

func SetDefaults(cfg *ConsulAdmConfig) error {
	if len(cfg.Name) == 0 {
		name, err := os.Hostname()
		if err != nil {
			return fmt.Errorf("unable to use hostname as default name: %s", err)
		}
		cfg.Name = name
	}

	DefaultAdvertise(cfg)
	cfg.Version = constants.DefaultVersion
	cfg.DataDir = constants.DefaultDataDir
	cfg.DataCenter = constants.DefaultDataCenter
	cfg.ServerMode = true
	cfg.ClientAddr = "0.0.0.0"

	return nil
}
