package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		err := cli.ContainerKill(context.Background(), container.ID, "SIGKILL")
		if err != nil {
			fmt.Errorf("could not kill container with ID %s", container.ID)
		}
	}
}
