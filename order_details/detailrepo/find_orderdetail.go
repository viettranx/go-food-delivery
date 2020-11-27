package detailrepo

// interface
type FindOrderDetailStorage interface {
}

// struct => store
type findOrderDetailRepo struct {
	store FindOrderDetailStorage
}

// implementation
