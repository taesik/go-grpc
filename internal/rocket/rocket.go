package rocket

import "context"

type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}

type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

type Service struct {
	Store Store
}

func New(store Store) Service {
	return Service{
		Store: store,
	}
}

func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}
func (s Service) InsertRocket(rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}
func (s Service) DeleteRocket(id string) error {
	err := s.Store.DeleteRocket(id)

	if err != nil {
		return err
	}

	return nil
}
