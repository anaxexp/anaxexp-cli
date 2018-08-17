package build

import (
	"github.com/anaxexp/anaxexp-cli/pkg/docker"
	"github.com/anaxexp/anaxexp-cli/pkg/types"
)

// Builder is builder representation.
type Builder struct {
	Client *docker.Client
}

// Build builds docker images for anaxexp services by config.
func (b *Builder) Build(config *types.BuildConfig, context string) error {
	return nil
}

func NewBuilder() *Builder {
	return &Builder{Client: docker.NewClient()}
}
