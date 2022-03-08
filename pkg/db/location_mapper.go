package db

import (
	"errors"
	"time"

	"github.com/zeurd/location/pkg/models"
)

type LocationMapper struct {
	// K: orderId, V: LocationHistory
	db map[string]*models.LocationHistory
}

func NewLocationMapper() *LocationMapper {
	return &LocationMapper{
		make(map[string]*models.LocationHistory),
	}
}

func (m *LocationMapper) GetLocationHistory(orderId string) (*models.LocationHistory, error) {
	lh, ok := m.db[orderId]
	if !ok {
		return nil, errors.New("location history not found")
	}
	return lh, nil
}

func (m *LocationMapper) GetLocationHistoryLimit(orderId string, limit int) (*models.LocationHistory, error) {
	lh, ok := m.db[orderId]
	if !ok {
		return nil, errors.New("location history not found")
	}
	if len(lh.History) <= limit {
		return lh, nil
	}
	start := len(lh.History) - limit
	return &models.LocationHistory{
		OrderId: orderId,
		History: lh.History[start:],
	}, nil
}

func (m *LocationMapper) Insert(orderId string, location models.Location) error {
	location.Timestamp = time.Now().Unix()
	his, ok := m.db[orderId]
	if !ok {
		m.db[orderId] = models.NewLocationHistory(orderId)
	}
	his.History = append(his.History, location)
	return nil
}

func (m *LocationMapper) Delete(orderId string) error {
	delete(m.db, orderId)
	return nil
}

// we append the History in chrono order
// so when limit timestamp is reached, we can drop all older entries
// this has to run in its own routine each interval sec
func (m *LocationMapper) Clean(limit int64, interval int) {
	for {
		time.Sleep(time.Duration(interval) * time.Second)
		for _, lh := range m.db {
			for i, h := range lh.History {
				if h.Timestamp >= time.Now().Unix()-limit {
					lh.History = lh.History[i:]
				}
			}
		}
	}
}
