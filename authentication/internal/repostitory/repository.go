package repostitory

type repo struct {
	database mockUsers
}

func NewRepository() controller.repositoryInterfaces {
	return &repo{database: mockUsers{}}
}
