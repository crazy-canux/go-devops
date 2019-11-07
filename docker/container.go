package docker

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"log"
)

var cli *client.Client

func init() {
	var err error
	cli, err = client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		log.Fatal("New Client failed.")
	}
}

func ServiceCheck(name string) bool {
	containers, err := cli.ContainerList(
		context.Background(),
		types.ContainerListOptions{
			All: true,
		})
	if err != nil {
		log.Fatalf("List containers failed: %s", err.Error())
	}
	for _, container := range containers {
		if container.Names[0] == ("/" + name) {
			log.Printf("Container %s, status: %s, state: %s.", container.Names[0], container.Status, container.State)
			if container.State == "running" {
				return true
			} else {
				return false
			}
		}
	}
	log.Printf("Container %s not exist.", name)
	return false
}
