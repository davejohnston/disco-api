package resolver

import (
	"context"

	"github.com/davejohnston/disco-api/internal/data"
	"github.com/davejohnston/disco-api/internal/model"
	"github.com/davejohnston/disco-api/internal/server"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() server.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) AllHosts(ctx context.Context, after *string, first *int, before *string, last *int) (*model.HostConnection, error) {
	hosts, count := data.Hosts(first)

	return &model.HostConnection{
		TotalCount: count,
		Hosts:      hosts,
	}, nil
}
func (r *queryResolver) AllVulnerabilties(ctx context.Context, after *string, first *int, before *string, last *int) (*model.VulnerabilityConnection, error) {
	panic("not implemented")
}
