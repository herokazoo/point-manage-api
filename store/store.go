package store

import (
	"errors"

	"github.com/herokazoo/point-manage-api/entity"
)

var Histories = &HistoryStore{Histories: map[entity.UserPointHistoryID]*entity.UserPointHistory{}}
var ErrNotFound = errors.New("not found")

type HistoryStore struct {
	LastID    entity.UserPointHistoryID
	Histories map[entity.UserPointHistoryID]*entity.UserPointHistory
}

func (hs *HistoryStore) Add(h *entity.UserPointHistory) (entity.UserPointHistoryID, error) {
	hs.LastID++
	h.ID = hs.LastID
	hs.Histories[h.ID] = h
	return h.ID, nil
}

func (hs *HistoryStore) Get(id entity.UserPointHistoryID) (*entity.UserPointHistory, error) {
	if hs, ok := hs.Histories[id]; ok {
		return hs, nil
	}
	return nil, ErrNotFound
}

func (hs *HistoryStore) All() entity.UserPointHistories {
	histories := make([]*entity.UserPointHistory, len(hs.Histories))
	for i, h := range hs.Histories {
		histories[i-1] = h
	}
	return histories
}
