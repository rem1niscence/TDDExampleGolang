package person

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPersonFails(t *testing.T) {
	c := require.New(t)

	person, err := NewPerson("", 20, 1520)
	c.Equal(ErrMissingName, err)
	c.Empty(person)

	person, err = NewPerson("Roniel", -10, 1520)
	c.Equal(ErrInvalidAge, err)
	c.Empty(person)

	person, err = NewPerson("Roniel", 20, -1520)
	c.Equal(ErrInvalidBalance, err)
	c.Empty(person)
}

func TestNewPerson(t *testing.T) {
	c := require.New(t)

	person, err := NewPerson("Roniel", 20, 1520)
	c.Nil(err)
	c.Equal(person.Name, "Roniel")
	c.Equal(person.Age, 20)
	c.Equal(person.Balance, 1520.00)
}

func TestAddFriend(t *testing.T) {
	c := require.New(t)

	person, _ := NewPerson("Roniel", 20, 1520)
	friend, _ := NewPerson("Albin", 20, 3300)

	err := person.AddFriend(friend)
	c.Nil(err)
	c.True(person.isFriendsWith(friend))
	c.True(friend.isFriendsWith(person))

	err = person.AddFriend(friend)
	c.Equal(ErrFriendAlreadyAdded, err)

	err = person.AddFriend(person)
	c.Equal(ErrCannotBefriendItself, err)

}

func TestRemoveFriend(t *testing.T) {
	c := require.New(t)

	person, _ := NewPerson("Roniel", 20, 1520)
	friend, _ := NewPerson("Albin", 20, 3300)
	notFriend, _ := NewPerson("Enrique", 21, 5540)

	err := person.AddFriend(friend)
	c.Nil(err)
	c.True(person.isFriendsWith(friend))
	c.True(friend.isFriendsWith(person))

	err = person.RemoveFriend(friend)
	c.Nil(err)
	c.False(person.isFriendsWith(friend))
	c.False(friend.isFriendsWith(person))

	err = person.RemoveFriend(notFriend)
	c.Equal(ErrNotFriend, err)

}

func TestGiveMoneyToFriend(t *testing.T) {
	c := require.New(t)

	person, _ := NewPerson("Roniel", 20, 1520)
	friend, _ := NewPerson("Albin", 20, 3300)
	notFriend, _ := NewPerson("Enrique", 21, 5540)

	amount := 500.00
	personInitialBalance := person.Balance
	friendInitialBalance := friend.Balance

	err := person.AddFriend(friend)
	c.Nil(err)

	err = person.GiveMoneyToFriend(amount, friend)
	c.Nil(err)
	c.Equal(person.Balance, personInitialBalance-amount)
	c.Equal(friend.Balance, friendInitialBalance+amount)

	err = person.GiveMoneyToFriend(amount, notFriend)
	c.Equal(ErrNotFriend, err)
	c.Equal(person.Balance, personInitialBalance-amount)
}

func TestGiveMoneyToFriendNoBalance(t *testing.T) {
	c := require.New(t)

	person, _ := NewPerson("Roniel", 20, 300)
	friend, _ := NewPerson("Albin", 20, 3300)

	amount := 500.00
	personInitialBalance := person.Balance
	friendInitialBalance := friend.Balance

	err := person.AddFriend(friend)
	c.Nil(err)

	err = person.GiveMoneyToFriend(amount, friend)
	c.Equal(ErrInvalidBalance, err)
	c.Equal(personInitialBalance, person.Balance)
	c.Equal(friendInitialBalance, friend.Balance)
}

func TestReceiveMoneyFroMFriendNoBalance(t *testing.T) {
	c := require.New(t)

	person, _ := NewPerson("Roniel", 20, 1200)
	friend, _ := NewPerson("Albin", 20, 300)

	amount := 500.00
	personInitialBalance := person.Balance
	friendInitialBalance := friend.Balance

	err := person.AddFriend(friend)
	c.Nil(err)

	err = person.ReceiveMoneyFromFriend(amount, friend)
	c.Equal(ErrInvalidBalance, err)
	c.Equal(personInitialBalance, person.Balance)
	c.Equal(friendInitialBalance, friend.Balance)
}

func TestReceiveMoneyFromFriend(t *testing.T) {
	c := require.New(t)

	person, _ := NewPerson("Roniel", 20, 1520)
	friend, _ := NewPerson("Albin", 20, 3300)
	notFriend, _ := NewPerson("Enrique", 21, 5540)

	amount := 500.00
	personInitialBalance := person.Balance
	friendInitialBalance := friend.Balance

	err := person.AddFriend(friend)
	c.Nil(err)

	err = person.ReceiveMoneyFromFriend(amount, friend)
	c.Nil(err)
	c.Equal(person.Balance, personInitialBalance+amount)
	c.Equal(friend.Balance, friendInitialBalance-amount)

	err = person.ReceiveMoneyFromFriend(amount, notFriend)
	c.Equal(ErrNotFriend, err)
	c.Equal(person.Balance, personInitialBalance+amount)
}
