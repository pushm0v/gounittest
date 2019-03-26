package service

type StorageService interface {
	Save(key string, data interface{})
	Fetch(key string) interface{}
}

type storageService struct {
	memoryStore map[string]interface{}
}

func NewStorage() StorageService {
	return &storageService{
		memoryStore: make(map[string]interface{}),
	}
}

func (s *storageService) Save(key string, data interface{}) {
	s.memoryStore[key] = data
}

func (s *storageService) Fetch(key string) interface{} {
	if exist, ok := s.memoryStore[key]; ok {
		return exist
	}

	return nil
}
