package gkin

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/moby/moby/client"
)

// Build is docker container build.
// return image name
func Build(pipe Pipe) (string, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return "", err
	}

	// pull
	fmt.Println("Container pulling")
	ctx := context.Background()
	rp, err := cli.ImagePull(ctx, pipe.Image, types.ImagePullOptions{})
	if err != nil {
		return "", err
	}
	defer rp.Close()

	// wait
	var s string
	for {
		n, err := fmt.Scanf("%s", &s)
		if err != nil {
			return "", err
		}
		if n == 0 {
			break
		}
	}

	cport, err := nat.NewPort("http", "80")
	if err != nil {
		return "", err
	}
	// container config
	cc := &container.Config{
		Image: pipe.Image,
		ExposedPorts: nat.PortSet{
			cport: struct{}{},
		},
		Entrypoint: pipe.Commands,
	}
	// host config
	hc := &container.HostConfig{
		PortBindings: nat.PortMap{
			cport: []nat.PortBinding{nat.PortBinding{HostPort: "8888"}},
		},
		AutoRemove: true,
	}

	body, err := cli.ContainerCreate(ctx, cc, hc, &network.NetworkingConfig{}, pipe.Name)
	if err != nil {
		return "", err
	}
	fmt.Println("Container start")
	if err = cli.ContainerStart(ctx, body.ID, types.ContainerStartOptions{}); err != nil {
		return "", err
	}
	if _, err = cli.ContainerWait(ctx, body.ID); err != nil {
		return "", err
	}

	r, err := cli.ContainerLogs(ctx, body.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return "", err
	}
	defer r.Close()

	io.Copy(os.Stdout, r)
	return "", nil
}
