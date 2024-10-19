package main

import (
	"context"
	"dagger/spellcheck/internal/dagger"
)

type Spellcheck struct {
	// +private
	Ctr *dagger.Container
}

func New(
	// Spellcheck image to use.
	// +optional
	// renovate image: datasource=docker depName=jonasbn/github-action-spellcheck versioning=docker
	// +default="jonasbn/github-action-spellcheck:0.43.1"
	Image string,
) *Spellcheck {
	return &Spellcheck{
		Ctr: dag.Container().From(Image),
	}
}

// Spellcheck runs spellcheck.
//
// Example usage: dagger call spellcheck --source /path/to/your/repo
func (m *Spellcheck) Spellcheck(
	ctx context.Context,
	// The directory of the repository.
	source *dagger.Directory,
) *dagger.Container {
	return m.Ctr.
		WithMountedDirectory("/src", source).
		WithWorkdir("/src").
		WithExec(nil, dagger.ContainerWithExecOpts{UseEntrypoint: true})
}
