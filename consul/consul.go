package consul

import (
	"strings"

	"github.com/hashicorp/consul/api"
)

// Deregister -
func Deregister(host, key string) {
	c := &api.Config{
		Address: host,
	}
	// Get a new client
	client, err := api.NewClient(c)
	if err != nil {
		panic(err)
	}

	services, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}

	for _, v := range services {
		if strings.Contains(v.ID, key) {
			if err := client.Agent().ServiceDeregister(v.ID); err != nil {
				panic(err)
			}
		}
	}
}

// DeregisterAll -
func DeregisterAll(host string) {
	c := &api.Config{
		Address: host,
	}
	// Get a new client
	client, err := api.NewClient(c)
	if err != nil {
		panic(err)
	}

	services, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}

	for _, v := range services {
		if err := client.Agent().ServiceDeregister(v.ID); err != nil {
			panic(err)
		}
	}
}
