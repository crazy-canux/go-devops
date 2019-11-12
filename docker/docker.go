package docker

import (
	"errors"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
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

func CreateNetwork(name, driver, subnet, gateway string) (string, error) {
	options := types.NetworkCreate{
		CheckDuplicate: true,
		Driver:         driver,
		IPAM: &network.IPAM{
			Driver: "default",
			Config: []network.IPAMConfig{
				{
					Subnet:  subnet,
					Gateway: gateway,
				},
			},
		},
		Options: map[string]string{
			"com.docker.network.bridge.name": name,
		},
	}
	resp, err := cli.NetworkCreate(context.Background(), name, options)
	if err != nil {
		log.Fatalf("Create Docker Network %s failed: %s", name, err.Error())
		return "", errors.New("create docker network failed")
	}
	return resp.ID, nil
}

func RemoveNetwork(name string) bool {
	id, err := NetworkId(name)
	if err != nil {
		log.Fatalf("Docker Network %s not exist.", name)
	} else {
		err := cli.NetworkRemove(context.Background(), id)
		if err != nil {
			log.Fatalf("Remove Docker Network %s failed: %s", name, err.Error())
			return false
		}
	}
	log.Fatalf("Remove Docker Network %s succeed.", name)
	return true
}

func NetworkId(name string) (string, error) {
	networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		log.Fatalf("List docker network failed: %s", err.Error())
	}
	for _, network := range networks {
		if name == network.Name {
			log.Printf("Docker network %s exist.", name)
			return network.ID, nil
		}
	}
	return "", errors.New("docker network not exist")
}
