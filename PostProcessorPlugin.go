package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/plugin"
	"github.com/hashicorp/packer-plugin-sdk/version"
	"github.com/mitchellh/mapstructure"
	"github.com/zclconf/go-cty/cty"
)

var (
	// Version is the plugin version - set by GoReleaser
	Version = "dev"
	// VersionPrerelease is the plugin prerelease version
	VersionPrerelease = ""
	// PluginVersion is the common plugin version
	PluginVersion = version.InitializePluginVersion(Version, VersionPrerelease)
)

type Config struct {
	Name string `mapstructure:"name" required:"true"`
}

// FlatMapstructure returns a new FlatConfig.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

type FlatConfig struct {
	Name *string `mapstructure:"name" required:"true" cty:"name" hcl:"name"`
}

// HCL2Spec returns the hcl spec of a Config.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"name": &hcldec.AttrSpec{Name: "name", Type: cty.String, Required: true},
	}
	return s
}

type PostProcessor struct {
	config Config
}

func (p *PostProcessor) ConfigSpec() hcldec.ObjectSpec {
	return p.config.FlatMapstructure().HCL2Spec()
}

func (p *PostProcessor) Configure(raws ...interface{}) error {
	// Use simple mapstructure decode
	if len(raws) == 0 {
		return fmt.Errorf("no configuration provided")
	}

	err := mapstructure.Decode(raws[0], &p.config)
	if err != nil {
		return fmt.Errorf("error decoding config: %v", err)
	}

	if p.config.Name == "" {
		return fmt.Errorf("name parameter is required")
	}

	return nil
}

func (p *PostProcessor) PostProcess(ctx context.Context, ui packer.Ui, artifact packer.Artifact) (packer.Artifact, bool, bool, error) {
	ui.Say(fmt.Sprintf("Hello from post-processor! Name: %s", p.config.Name))

	// Return the original artifact unchanged
	return artifact, false, false, nil
}

func main() {
	pps := plugin.NewSet()
	pps.RegisterPostProcessor("syft", new(PostProcessor))
	pps.SetVersion(PluginVersion)
	err := pps.Run()
	if err != nil {
		log.Fatalf("Failed to run Syft plugin: %v", err)
	}
}
