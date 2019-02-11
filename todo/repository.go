package todo

// Repository ...
type Repository interface {
	Find(id string) (*Todo, error)
	FindAll() ([]*Todo, error)
}

// InMemoryRepository ...
type InMemoryRepository struct {
	todos  map[string]*Todo
	nextID string
}

// NewInMemoryRepository ...
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		todos:  make(map[string]*Todo),
		nextID: "1",
	}
}
