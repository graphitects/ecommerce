package handler

import "testing"

func TestHandler_Create(t *testing.T) {
	t.Run("success to save a product", func(t *testing.T) {
		// arrange

		// act

		// assert
	})

	t.Run("failed to save a product", func(t *testing.T) {
		// arrange

		// act

		// assert
	})

	t.Run("failed to save a product with invalid request", func(t *testing.T) {
		// arrange

		// act

		// assert
	})
}

func TestHandler_GetAll(t *testing.T) {
	type TableTest struct {
		input []any
		set func()
		expected []any
	}

	funcTest := func(t *testing.T, c TableTest) {
			// arrange

			// act

			// assert
	}

	t.Run("success to get all products", func(t *testing.T) {
		// act
		funcTest(t, TableTest{
			input: []any{1},
			set: func() {
				// arrange
			},
			expected: []any{
				`{"id":"1","name":"product 1","price":1000}`,
			},
		})
	})
}


func TestHandler_Get(t *testing.T) {
	type TableTest struct {
		name string
		input []any
		set func()
		expected []any
	}
	cases := []TableTest{
		{
			name: "success to get a product",
			input: []any{"1"},
			set: func() {
				// arrange
			},
			expected: []any{
				`{"id":"1","name":"product 1","price":1000}`,
			},
		},
		{
			name: "success to get a product",
			input: []any{"1"},
			set: func() {
				// arrange
			},
			expected: []any{
				`{"id":"1","name":"product 1","price":1000}`,
			},
		},
		{
			name: "success to get a product",
			input: []any{"1"},
			set: func() {
				// arrange
			},
			expected: []any{
				`{"id":"1","name":"product 1","price":1000}`,
			},
		},
		{
			name: "success to get a product",
			input: []any{"1"},
			set: func() {
				// arrange
			},
			expected: []any{
				`{"id":"1","name":"product 1","price":1000}`,
			},
		},
		{
			name: "success to get a product",
			input: []any{"1"},
			set: func() {
				// arrange
			},
			expected: []any{
				`{"id":"1","name":"product 1","price":1000}`,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.set()

			// act


			// assert

		})
	}
}