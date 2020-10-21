package main

import (
	"testing"
)

func TestUserBuilder(t *testing.T) {
	t.Run("Successfully create free user", func(t *testing.T) {
		userBuilder := NewUserBuilder(0)

		user := userBuilder.SetName("mashu_matrix").SetGender("men").SetAge(29).Build()

		expected := &User{
			Kind:   0,
			Name:   "mashu_matrix",
			Gender: "men",
			Age:    29,
		}

		if user.Kind != expected.Kind {
			t.Errorf("got %v expected %v", user, expected)
		}
	})

	t.Run("Successfully create monthly user", func(t *testing.T) {
		userBuilder := NewUserBuilder(1)

		user := userBuilder.SetName("mashu_matrix").SetGender("men").SetAge(29).Build()

		expected := &User{
			Kind:   1,
			Name:   "mashu_matrix",
			Gender: "men",
			Age:    29,
		}

		if user.Kind != expected.Kind {
			t.Errorf("got %v expected %v", user, expected)
		}
	})

	t.Run("Successfully create Yearly user", func(t *testing.T) {
		userBuilder := NewUserBuilder(2)

		user := userBuilder.SetName("mashu_matrix").SetGender("men").SetAge(29).Build()

		expected := &User{
			Kind:   2,
			Name:   "mashu_matrix",
			Gender: "men",
			Age:    29,
		}

		if user.Kind != expected.Kind {
			t.Errorf("got %v expected %v", user, expected)
		}
	})
}
