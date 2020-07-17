package consul

import (
	"context"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/hashicorp/consul/api"
)

// CloneKV -
func CloneKV(target, self string) error {
	targetC := &api.Config{
		Address: target,
	}

	// Get a new client
	targetClient, err := api.NewClient(targetC)
	if err != nil {
		panic(err)
	}

	selfC := &api.Config{
		Address: self,
	}

	// Get a new client
	selfClient, err := api.NewClient(selfC)
	if err != nil {
		panic(err)
	}

	keys := GetAllKV(target)
	for _, k := range keys {
		pair, _, _ := targetClient.KV().Get(k, nil)
		p := &api.KVPair{
			Key:   k,
			Value: pair.Value,
		}

		selfClient.KV().Put(p, nil)
	}

	return nil

}

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

// GetAllKV -
func GetAllKV(host string) []string {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var nodes []*cdp.Node
	var buf []byte
	err := chromedp.Run(ctx,
		// 訪問頁面
		chromedp.Navigate(host+"/ui/dc1/kv"),
		// 等待該元件顯示
		chromedp.WaitVisible(`.type-create`),
		chromedp.Nodes(".file a", &nodes, chromedp.ByQueryAll),
	)

	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("fullScreenshot.png", buf, 0644); err != nil {
		log.Fatal(err)
	}

	keys := make([]string, len(nodes))
	for i, v := range nodes {
		href := v.AttributeValue("href")

		keys[i] = strings.Split(href, "/")[4]
	}

	return keys
}
