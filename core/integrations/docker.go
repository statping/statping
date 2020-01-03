// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package integrations

import (
	"context"
	dTypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/hunterlong/statping/types"
	"os"
)

type dockerIntegration struct {
	*types.Integration
}

var dockerIntegrator = &dockerIntegration{&types.Integration{
	ShortName: "docker",
	Name:      "Docker",
	Icon:      "<i class=\"fab fa-docker\"></i>",
	Description: `Import multiple services from Docker by attaching the unix socket to Statping. 
You can also do this in Docker by setting <u>-v /var/run/docker.sock:/var/run/docker.sock</u> in the Statping Docker container. 
All of the containers with open TCP/UDP ports will be listed for you to choose which services you want to add. If you running Statping inside of a container, 
this container must be attached to all networks you want to communicate with.`,
	Fields: []*types.IntegrationField{
		{
			Name:        "path",
			Description: "The absolute path to the Docker unix socket",
			Type:        "text",
			Value:       client.DefaultDockerHost,
		},
		{
			Name:        "version",
			Description: "Version number of Docker server",
			Type:        "text",
			Value:       client.DefaultVersion,
		},
	},
}}

var cli *client.Client

func (t *dockerIntegration) Get() *types.Integration {
	return t.Integration
}

func (t *dockerIntegration) List() ([]*types.Service, error) {
	var err error
	path := Value(t, "path").(string)
	version := Value(t, "version").(string)
	os.Setenv("DOCKER_HOST", path)
	os.Setenv("DOCKER_VERSION", version)
	cli, err = client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	var services []*types.Service

	containers, err := cli.ContainerList(context.Background(), dTypes.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	for _, container := range containers {
		if container.State != "running" {
			continue
		}

		for _, v := range container.Ports {
			if v.IP == "" {
				continue
			}

			service := &types.Service{
				Name:     container.Names[0][1:],
				Domain:   v.IP,
				Type:     v.Type,
				Port:     int(v.PublicPort),
				Interval: 60,
				Timeout:  2,
			}

			services = append(services, service)
		}

	}
	return services, nil
}
