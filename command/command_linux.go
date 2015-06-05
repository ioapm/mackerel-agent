package command

import (
	"time"

	"github.com/mackerelio/mackerel-agent/config"
	"github.com/mackerelio/mackerel-agent/metrics"
	metricsLinux "github.com/mackerelio/mackerel-agent/metrics/linux"
	"github.com/mackerelio/mackerel-agent/spec"
	specLinux "github.com/mackerelio/mackerel-agent/spec/linux"
)

func specGenerators() []spec.Generator {
	specs := []spec.Generator{
		&specLinux.KernelGenerator{},
		&specLinux.CPUGenerator{},
		&specLinux.MemoryGenerator{},
		&specLinux.BlockDeviceGenerator{},
		&specLinux.FilesystemGenerator{},
	}
	cloudGenerator, err := specLinux.NewCloudGenerator("")
	if err != nil {
		logger.Errorf("Failed to create cloudGenerator: %s", err.Error())
	} else {
		specs = append(specs, cloudGenerator)
	}
	return specs
}

func interfaceGenerator() spec.Generator {
	return &specLinux.InterfaceGenerator{}
}

func metricsGenerators(conf *config.Config) []metrics.Generator {
	generators := []metrics.Generator{
		&metricsLinux.Loadavg5Generator{},
		&metricsLinux.CPUUsageGenerator{Interval: time.Duration(metricsInterval)},
		&metricsLinux.MemoryGenerator{},
		&metricsLinux.UptimeGenerator{},
		&metricsLinux.InterfaceGenerator{Interval: time.Duration(metricsInterval)},
		&metricsLinux.DiskGenerator{Interval: time.Duration(metricsInterval)},
		&metricsLinux.FilesystemGenerator{},
	}

	return generators
}

func pluginGenerators(conf *config.Config) []metrics.PluginGenerator {
	generators := []metrics.PluginGenerator{}

	for _, pluginConfig := range conf.Plugin["metrics"] {
		generators = append(generators, metrics.NewPluginGenerator(pluginConfig))
	}

	return generators
}
