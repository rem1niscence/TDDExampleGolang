package person

import "errors"

var (
	// ErrMissingName error when the name parameter is missing
	ErrMissingName = errors.New("missing name field")
	// ErrInvalidAge error when the age parameter is lower than 0
	ErrInvalidAge = errors.New("invalid age: age cannot be lower than 0")
	// ErrInvalidBalance error when the balance lower than 0
	ErrInvalidBalance = errors.New("invalid balance: balance cannot be lower than 0")
	// ErrFriendAlreadyAdded error when trying to add an existing friend
	ErrFriendAlreadyAdded = errors.New("friends has already been added")
	// ErrCannotBefriendItself error when trying to be friend with itself
	ErrCannotBefriendItself = errors.New("cannot be friends with itself")
	// ErrNotFriend error when the person has not that other person has a friend
	ErrNotFriend = errors.New("is not a friend")
)

// Person is the struct handler for person
type Person struct {
	Name    string
	Age     int
	Balance float64
	Friends map[string]*Person
}

// NewPerson is the Factory function for person
func NewPerson(name string, age int, balance float64) (*Person, error) {
	if name == "" {
		return nil, ErrMissingName
	}

	if age < 0 {
		return nil, ErrInvalidAge
	}

	if balance < 0.0 {
		return nil, ErrInvalidBalance
	}

	return &Person{
		Name:    name,
		Age:     age,
		Balance: balance,
		Friends: make(map[string]*Person),
	}, nil
}

// AddFriend adds a friend to a person
func (p *Person) AddFriend(person *Person) error {
	if p.isFriendsWith(person) {
		return ErrFriendAlreadyAdded
	}

	if p.Name == person.Name {
		return ErrCannotBefriendItself
	}

	p.Friends[person.Name] = person
	person.Friends[p.Name] = p

	return nil
}

// RemoveFriend removes a friend from a person
func (p *Person) RemoveFriend(friend *Person) error {
	if !p.isFriendsWith(friend) {
		return ErrNotFriend
	}

	delete(p.Friends, friend.Name)
	delete(friend.Friends, p.Name)

	return nil
}

func (p *Person) isFriendsWith(person *Person) bool {
	_, found := p.Friends[person.Name]
	return found
}

// GiveMoneyToFriend gives money to a friend
func (p *Person) GiveMoneyToFriend(amount float64, friend *Person) error {
	if !p.isFriendsWith(friend) {
		return ErrNotFriend
	}

	if p.Balance-amount < 0 {
		return ErrInvalidBalance
	}

	friend.Balance += amount
	p.Balance -= amount

	return nil
}

// ReceiveMoneyFromFriend receives money from a friend
func (p *Person) ReceiveMoneyFromFriend(amount float64, friend *Person) error {
	if !p.isFriendsWith(friend) {
		return ErrNotFriend
	}

	if friend.Balance-amount < 0 {
		return ErrInvalidBalance
	}

	p.Balance += amount
	friend.Balance -= amount

	return nil
}
