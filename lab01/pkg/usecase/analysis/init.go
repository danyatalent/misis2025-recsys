package analysis

type UseCase struct {
	apis []API
}

func New(apis ...API) (*UseCase, error) {
	return &UseCase{apis: apis}, nil
}
