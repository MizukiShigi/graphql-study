package paginationutil

import "errors"

const maxLimit = 100

type ListParams struct {
	After  *string
	Before *string
	First  *int32
	Last   *int32
}

func (p *ListParams) Validate() error {
	// 排他的なパラメータの検証
	if p.After != nil && p.Before != nil {
		return errors.New("after and before cannot be set at the same time")
	}
	if p.First != nil && p.Last != nil {
		return errors.New("first and last cannot be set at the same time")
	}
    // afterはfirstとセットでのみ使用可能
    if p.After != nil && p.First == nil {
        return errors.New("after must be used with first")
    }
    // beforeはlastとセットでのみ使用可能
    if p.Before != nil && p.Last == nil {
        return errors.New("before must be used with last")
    }
	// 値の検証
	if p.First != nil && *p.First < 0 {
			return errors.New("first must be positive")
	}
	if p.Last != nil && *p.Last < 0 {
		return errors.New("last must be positive")
	}
	return nil
}

func (p *ListParams) GetLimit() int {
    if p.First != nil {
		if *p.First > maxLimit {
			return maxLimit
		}
        return int(*p.First)
    }
    if p.Last != nil {
		if *p.Last > maxLimit {
			return maxLimit
		}
        return int(*p.Last)
    }
    return 10
}
