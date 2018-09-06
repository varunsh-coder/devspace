package configutil

import "github.com/covexo/devspace/pkg/devspace/config/v1"

func makeConfig() *v1.Config {
	return &v1.Config{
		Cluster: &v1.Cluster{
			User: &v1.User{},
		},
		DevSpace: &v1.DevSpaceConfig{
			PortForwarding: []*v1.PortForwardingConfig{},
			Release:        &v1.Release{},
			Sync:           []*v1.SyncConfig{},
		},
		Image: &v1.ImageConfig{},
		Services: &v1.ServiceConfig{
			Registry: &v1.RegistryConfig{
				Internal: &v1.InternalRegistry{
					Release: &v1.Release{},
				},
				User: &v1.RegistryUser{},
			},
			Tiller: &v1.TillerConfig{
				AppNamespaces: []*string{},
				Release:       &v1.Release{},
			},
		},
	}
}