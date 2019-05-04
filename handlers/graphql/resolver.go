//go:generate go run github.com/99designs/gqlgen
package graphql

import (
	"context"
	"github.com/hunterlong/statping/core"

	"github.com/hunterlong/statping/types"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Checkin() CheckinResolver {
	return &checkinResolver{r}
}
func (r *Resolver) Core() CoreResolver {
	return &coreResolver{r}
}
func (r *Resolver) Group() GroupResolver {
	return &groupResolver{r}
}
func (r *Resolver) Message() MessageResolver {
	return &messageResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Service() ServiceResolver {
	return &serviceResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

type checkinResolver struct{ *Resolver }

func (r *checkinResolver) Service(ctx context.Context, obj *types.Checkin) (*types.Service, error) {
	service := core.SelectService(obj.ServiceId)
	return service.Service, nil
}
func (r *checkinResolver) Failures(ctx context.Context, obj *types.Checkin) ([]*types.Failure, error) {
	all := obj.Failures
	var objs []*types.Failure
	for _, v := range all {
		objs = append(objs, v.Select())
	}
	return objs, nil
}

type coreResolver struct{ *Resolver }

func (r *coreResolver) Footer(ctx context.Context, obj *types.Core) (string, error) {
	panic("not implemented")
}
func (r *coreResolver) Timezone(ctx context.Context, obj *types.Core) (string, error) {
	panic("not implemented")
}
func (r *coreResolver) UsingCdn(ctx context.Context, obj *types.Core) (bool, error) {
	panic("not implemented")
}

type messageResolver struct{ *Resolver }

func (r *messageResolver) NotifyUsers(ctx context.Context, obj *types.Message) (bool, error) {
	panic("not implemented")
}
func (r *messageResolver) NotifyMethod(ctx context.Context, obj *types.Message) (bool, error) {
	panic("not implemented")
}
func (r *messageResolver) NotifyBefore(ctx context.Context, obj *types.Message) (int, error) {
	panic("not implemented")
}

type groupResolver struct{ *Resolver }

func (r *groupResolver) Public(ctx context.Context, obj *types.Group) (bool, error) {
	return obj.Public.Bool, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Core(ctx context.Context) (*types.Core, error) {
	c := core.CoreApp
	return c.Core, nil
}

func (r *queryResolver) Message(ctx context.Context, id int64) (*types.Message, error) {
	message, err := core.SelectMessage(id)
	return message.Message, err
}
func (r *queryResolver) Messages(ctx context.Context) ([]*types.Message, error) {
	all, err := core.SelectMessages()
	var objs []*types.Message
	for _, v := range all {
		objs = append(objs, v.Message)
	}
	return objs, err
}

func (r *queryResolver) Group(ctx context.Context, id int64) (*types.Group, error) {
	group := core.SelectGroup(id)
	return group.Group, nil
}
func (r *queryResolver) Groups(ctx context.Context) ([]*types.Group, error) {
	all := core.SelectGroups(true, true)
	var objs []*types.Group
	for _, v := range all {
		objs = append(objs, v.Group)
	}
	return objs, nil
}

func (r *queryResolver) Checkin(ctx context.Context, id int64) (*types.Checkin, error) {
	panic("not implemented")
}
func (r *queryResolver) Checkins(ctx context.Context) ([]*types.Checkin, error) {
	panic("not implemented")
}

func (r *queryResolver) User(ctx context.Context, id int64) (*types.User, error) {
	user, err := core.SelectUser(id)
	return user.User, err
}
func (r *queryResolver) Users(ctx context.Context) ([]*types.User, error) {
	all, err := core.SelectAllUsers()
	var objs []*types.User
	for _, v := range all {
		objs = append(objs, v.User)
	}
	return objs, err
}

type userResolver struct{ *Resolver }

func (r *userResolver) Admin(ctx context.Context, obj *types.User) (bool, error) {
	return obj.Admin.Bool, nil
}

type serviceResolver struct{ *Resolver }

func (r *queryResolver) Service(ctx context.Context, id int64) (*types.Service, error) {
	service := core.SelectService(id)
	return service.Service, nil
}
func (r *queryResolver) Services(ctx context.Context) ([]*types.Service, error) {
	all := core.Services()
	var objs []*types.Service
	for _, v := range all {
		objs = append(objs, v.Select())
	}
	return objs, nil
}

func (r *serviceResolver) Expected(ctx context.Context, obj *types.Service) (string, error) {
	return obj.Expected.String, nil
}
func (r *serviceResolver) PostData(ctx context.Context, obj *types.Service) (string, error) {
	return obj.PostData.String, nil
}
func (r *serviceResolver) AllowNotifications(ctx context.Context, obj *types.Service) (bool, error) {
	return obj.AllowNotifications.Bool, nil
}
func (r *serviceResolver) Public(ctx context.Context, obj *types.Service) (bool, error) {
	return obj.Public.Bool, nil
}
func (r *serviceResolver) Headers(ctx context.Context, obj *types.Service) (string, error) {
	return obj.Headers.String, nil
}
func (r *serviceResolver) Permalink(ctx context.Context, obj *types.Service) (string, error) {
	return obj.Permalink.String, nil
}
func (r *serviceResolver) Online24Hours(ctx context.Context, obj *types.Service) (float64, error) {
	return float64(obj.Online24Hours), nil
}
func (r *serviceResolver) Failures(ctx context.Context, obj *types.Service) ([]*types.Failure, error) {
	all := obj.Failures
	var objs []*types.Failure
	for _, v := range all {
		objs = append(objs, v.Select())
	}
	return objs, nil
}
func (r *serviceResolver) Group(ctx context.Context, obj *types.Service) (*types.Group, error) {
	group := core.SelectGroup(int64(obj.GroupId))
	return group.Group, nil
}
