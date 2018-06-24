package gkin

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/moby/moby/client"
)

// DockerPayload is container pulling payload.
type DockerPayload struct {
	ID             string `json:"id"`
	Status         string `json:"status"`
	Progress       string `json:"progress"`
	ProgressDetail struct {
		Current uint16 `json:"current"`
		Total   uint16 `json:"total"`
	} `json:"progressDetail"`
}

// Build is docker container build.
// return image name
func Build(pipe Pipe) (string, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return "", err
	}

	// pull
	ctx, cancelFunc := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancelFunc()
	fmt.Printf("Container Image pulling: %v\n", pipe.Image)
	rp, err := cli.ImagePull(ctx, pipe.Image, types.ImagePullOptions{})
	if err != nil {
		return "", err
	}
	defer rp.Close()

	// wait
	fmt.Println("Waiting...")
	payload := DockerPayload{}
	scanner := bufio.NewScanner(rp)
	for scanner.Scan() {
		json.Unmarshal(scanner.Bytes(), &payload)
		fmt.Printf("%+v\n", payload)
	}

	return "", nil
}

// Run is docker container run.
func Run(image, name string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}
	cport, err := nat.NewPort("tcp", "7777")
	if err != nil {
		return err
	}
	// container config
	cc := &container.Config{
		Image: image,
		ExposedPorts: nat.PortSet{
			cport: struct{}{},
		},
		// Entrypoint: pipe.Commands,
	}
	// host config
	hc := &container.HostConfig{
		PortBindings: nat.PortMap{
			cport: []nat.PortBinding{nat.PortBinding{HostPort: "8888"}},
		},
		AutoRemove: true,
	}

	fmt.Println("Container create")
	ctx := context.Background()
	body, err := cli.ContainerCreate(ctx, cc, hc, &network.NetworkingConfig{}, name)
	if err != nil {
		return err
	}

	fmt.Println("Container start")
	if err = cli.ContainerStart(ctx, body.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	r, err := cli.ContainerLogs(ctx, body.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return err
	}
	defer r.Close()
	io.Copy(os.Stdout, r)

	if _, err = cli.ContainerWait(ctx, body.ID); err != nil {
		return err
	}

	return nil
}
