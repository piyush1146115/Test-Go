package SOLID

//Ref: https://levelup.gitconnected.com/practical-solid-in-golang-liskov-substitution-principle-e0d2eb9dd39

// If S is a subtype of T, then objects of type T in a program may be replaced with objects of type S without altering any of the desirable properties of that program
// If ObjectA is an instance of ClassA, and ObjectB is an instance of ClassB, and ClassB is a subtype of ClassA — if we use ObjectB instead
// ObjectA somewhere in the code, the functionality of the application must not be broken.

//Here is the point of LSP in Go: struct must not violate the purpose of the interface.

type ConvexQuadrilateral interface {
	GetArea() int
}

type EquilateralQuadrilateral interface {
	ConvexQuadrilateral
	SetA(a int)
}

type NonEquilateralQuadrilateral interface {
	ConvexQuadrilateral
	SetA(a int)
	SetB(b int)
}

type NonEquiangularQuadrilateral interface {
	ConvexQuadrilateral
	SetAngle(angle float64)
}

type Oblong struct {
	NonEquilateralQuadrilateral
	a int
	b int
}

type Square struct {
	EquilateralQuadrilateral
	a int
}

type Parallelogram struct {
	NonEquilateralQuadrilateral
	NonEquiangularQuadrilateral
	a     int
	b     int
	angle float64
}

type Rhombus struct {
	EquilateralQuadrilateral
	NonEquiangularQuadrilateral
	a     int
	angle float64
}

//bad code
//type User struct {
//	ID uuid.UUID
//	//
//	// some fields
//	//
//}
//
//type UserRepository interface {
//	Create(ctx context.Context, user User) (*User, error)
//	Update(ctx context.Context, user User) error
//}
//
//type DBUserRepository struct {
//	db *gorm.DB
//}
//
//func (r *DBUserRepository) Create(ctx context.Context, user User) (*User, error) {
//	err := r.db.WithContext(ctx).Create(&user).Error
//	return &user, err
//}
//
//func (r *DBUserRepository) Update(ctx context.Context, user User) error {
//	return r.db.WithContext(ctx).Save(&user).Error
//}
//
//type MemoryUserRepository struct {
//	users map[uuid.UUID]User
//}
//
//func (r *MemoryUserRepository) Create(_ context.Context, user User) (*User, error) {
//	if r.users == nil {
//		r.users = map[uuid.UUID]User{}
//	}
//	user.ID = uuid.New()
//	r.users[user.ID] = user
//
//	return &user, nil
//}
//
//func (r *MemoryUserRepository) Update(_ context.Context, user User) error {
//	if r.users == nil {
//		r.users = map[uuid.UUID]User{}
//	}
//	r.users[user.ID] = user
//
//	return nil
//}
