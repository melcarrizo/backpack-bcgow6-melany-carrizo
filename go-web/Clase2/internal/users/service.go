package users

type Service interface {
	GetAll() ([]User, error)
	Store(nombre, apellido, email, fecha string, edad int, altura float64, activo bool) (User, error)
	Update(id int, nombre, apellido, email, fecha string, edad int, altura float64, activo bool) (User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]User, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Store(nombre, apellido, email, fecha string, edad int, altura float64, activo bool) (User, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return User{}, err
	}

	lastID++

	producto, err := s.repository.Store(lastID, nombre, apellido, email, fecha, edad, altura, activo)
	if err != nil {
		return User{}, err
	}

	return producto, nil
}

func (s *service) Update(id int, nombre, apellido, email, fecha string, edad int, altura float64, activo bool) (User, error) {

	return s.repository.Update(id, nombre, apellido, email, fecha, edad, altura, activo)
}
