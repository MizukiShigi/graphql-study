package graph

import (
	"context"

	"github.com/MizukiShigi/graphql-study/graph/model"
	"github.com/graph-gophers/dataloader/v7"
	"github.com/MizukiShigi/graphql-study/graph/service"
)

type Loaders struct {
	UserLoader dataloader.Interface[string, *model.User]
}

func NewLoaders(srv service.Services) *Loaders {
	userBatcher := &userBatcher{Srv: srv}

	return &Loaders{
		UserLoader: dataloader.NewBatchedLoader(userBatcher.BatchGetUsers),
	}
}

type userBatcher struct {
	Srv service.Services
}

func (u *userBatcher) BatchGetUsers(ctx context.Context, IDs []string) []*dataloader.Result[*model.User] {
	results := make([]*dataloader.Result[*model.User], len(IDs))
	for i := range results {
		results[i] = &dataloader.Result[*model.User]{
			Error: nil,
		}
	}

	// 検索条件であるIDが、引数でもらったIDsスライスの何番目のインデックスに格納されていたのか検索できるようにmap化する
	indexs := make(map[string]int, len(IDs))
	for i, ID := range IDs {
		indexs[ID] = i
	}

	users, err := u.Srv.GetUsersByIDs(ctx, IDs)

	// 取得結果を、戻り値resultの中の適切な場所に格納する
	for _, user := range users {
		var rsl *dataloader.Result[*model.User]
		if err != nil {
			rsl = &dataloader.Result[*model.User]{
				Error: err,
			}
		} else {
			rsl = &dataloader.Result[*model.User]{
				Data: user,
			}
		}
		// 引数でもらった条件と順序を保ったまま戻り値のスライスを作る
		results[indexs[user.ID]] = rsl
	}
	return results
}
