package movie

type Service interface {
	Get(id IMDBID) (*MovieDetailResult, error)
	Search(param Param) (*SearchResult, error)
}

type service struct {
	client Client
}

func NewService(client Client) Service {
	return &service{
		client: client,
	}
}

func (s *service) Get(id IMDBID) (*MovieDetailResult, error) {
	result, err := s.client.Get(id.ToMap())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *service) Search(param Param) (*SearchResult, error) {
	result, err := s.client.Search(param.ToMap())
	if err != nil {
		return nil, err
	}
	return result, nil
}
