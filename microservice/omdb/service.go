package omdb

type Service interface {
	Get(id string) (*Movie, error)
	Search(param Param) (*SearchResult, error)
}

type service struct {
	client Client
}

func New(client Client) Service {
	return &service{
		client: client,
	}
}

func (s *service) Get(id string) (*Movie, error) {
	return nil, nil
}

func (s *service) Search(param Param) (*SearchResult, error) {
	return nil, nil
}
