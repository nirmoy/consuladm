package consul

import (
	"fmt"

	"github.com/hashicorp/consul/agent"
	"github.com/hashicorp/consul/agent/config"
	consulAPI "github.com/hashicorp/consul/api"
)

func GetMember(addr string) *consulAPI.AgentMember {
	config := consulAPI.DefaultConfig()
	config.Address = addr
	client, err := consulAPI.NewClient(config)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	members, err := client.Agent().Members(false)
	if err != nil {
		fmt.Print(err)
		return nil
	}

	return members[0]
}

func GetMemberDC(addr string) string {
	member := GetMember(addr)

	if member != nil {
		return member.Tags["dc"]
	}
	return ""
}

func AgentRun(flagArgs config.Flags) {

	fmt.Print(*flagArgs.Config.AdvertiseAddrLAN)
	b, err := config.NewBuilder(flagArgs)
	if err != nil {
		fmt.Print(err)
	}
	cfg, err := b.BuildAndValidate()
	if err != nil {
		fmt.Print(err)
	}
	agent, err := agent.New(&cfg)
	if err != nil {
		fmt.Printf("Error creating agent: %s", err)
		return
	}
	if err := agent.Start(); err != nil {
		fmt.Printf("Error starting agent: %s", err)
	}

}

func AgentJoin(addr string) error {
	config := consulAPI.DefaultConfig()
	client, err := consulAPI.NewClient(config)
	if err != nil {
		fmt.Print(err)
		return err
	}

	err = client.Agent().Join(addr, false)
	if err != nil {
		fmt.Print(err)
		return err
	}

	return nil
}
