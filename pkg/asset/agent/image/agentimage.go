package image

import (
	"encoding/json"
	"io"
	"os"

	"github.com/openshift/assisted-image-service/pkg/isoeditor"
	"github.com/openshift/installer/pkg/asset"
)

const (
	// TODO: Make this relative to the directory passed as --dir rather than
	// the current working directory
	agentISOFilename = "agent.iso"
)

// AgentImage is an asset that generates the bootable image used to install clusters.
type AgentImage struct {
	File *asset.File
}

var _ asset.WritableAsset = (*AgentImage)(nil)

// Dependencies returns the assets on which the Bootstrap asset depends.
func (a *AgentImage) Dependencies() []asset.Asset {
	return []asset.Asset{
		&Ignition{},
		&BaseIso{},
	}
}

// Generate generates the image file for to ISO asset.
func (a *AgentImage) Generate(dependencies asset.Parents) error {
	ignition := &Ignition{}
	dependencies.Get(ignition)

	baseImage := &BaseIso{}
	dependencies.Get(baseImage)

	ignitionByte, err := json.Marshal(ignition.Config)
	if err != nil {
		return err
	}

	ignitionContent := &isoeditor.IgnitionContent{Config: ignitionByte}
	custom, err := isoeditor.NewRHCOSStreamReader(baseImage.File.Filename, ignitionContent, nil)
	if err != nil {
		return err
	}
	defer custom.Close()

	// Remove symlink if it exists
	os.Remove(agentISOFilename)

	output, err := os.Create(agentISOFilename)
	if err != nil {
		return err
	}
	defer output.Close()

	_, err = io.Copy(output, custom)
	return err
}

// Name returns the human-friendly name of the asset.
func (a *AgentImage) Name() string {
	return "Agent Installer ISO"
}

// Load returns the ISO from disk.
func (a *AgentImage) Load(f asset.FileFetcher) (bool, error) {
	// The ISO will not be needed by another asset so load is noop.
	// This is implemented because it is required by WritableAsset
	return false, nil
}

// Files returns the files generated by the asset.
func (a *AgentImage) Files() []*asset.File {
	// Return empty array because File will never be loaded.
	return []*asset.File{}
}