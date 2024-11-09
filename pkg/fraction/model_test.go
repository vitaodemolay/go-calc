package fraction

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFraction(t *testing.T) {
	//Arrange
	tests := []struct {
		numerator   int
		denominator int
		expectError bool
	}{
		{1, 2, false},
		{3, 4, false},
		{1, 0, true},
	}

	for i, test := range tests {
		//Act
		_, err := NewFraction(test.numerator, test.denominator)

		//Assert
		assert.Equal(t, test.expectError, (err != nil), fmt.Sprintf("TestNewFraction fail on case test: %d. (case values = %+v) ", i+1, test))
	}
}

func TestAddWhenDenominatorIsDifferent(t *testing.T) {
	//Arrange
	f1, _ := NewFraction(1, 2)
	f2, _ := NewFraction(1, 3)
	expected, _ := NewFraction(5, 6)

	//Act
	result, err := f1.Add(f2)

	//Assert
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, result)
}

func TestAddWhenDenominatorIsTheSame(t *testing.T) {
	//Arrange
	f1, _ := NewFraction(1, 8)
	f2, _ := NewFraction(3, 8)
	expected, _ := NewFraction(4, 8)

	//Act
	result, err := f1.Add(f2)

	//Assert
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, result)
}

func TestSubtractWhenDenominatorIsDifferent(t *testing.T) {
	//Arrange
	f1, _ := NewFraction(1, 2)
	f2, _ := NewFraction(1, 3)
	expected, _ := NewFraction(1, 6)

	//Act
	result, err := f1.Subtract(f2)

	//Assert
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, result)
}

func TestSubtractWhenDenominatorIsTheSame(t *testing.T) {
	//Arrange
	f1, _ := NewFraction(3, 3)
	f2, _ := NewFraction(1, 3)
	expected, _ := NewFraction(2, 3)

	//Act
	result, err := f1.Subtract(f2)

	//Assert
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, result)
}

func TestMultiply(t *testing.T) {
	//Arrange
	f1, _ := NewFraction(1, 2)
	f2, _ := NewFraction(1, 3)
	expected, _ := NewFraction(1, 6)

	//Act
	result, err := f1.Multiply(f2)

	//Assert
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, result)
}

func TestDivide(t *testing.T) {
	//Arrange
	f1, _ := NewFraction(4, 12)
	f2, _ := NewFraction(8, 3)
	expected, _ := NewFraction(12, 96)

	//Act
	result, err := f1.Divide(f2)

	//Assert
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, result)
}

func TestToString(t *testing.T) {
	//Arrange
	f, _ := NewFraction(1, 2)
	expected := "1/2"

	//Act
	result := f.ToString()

	//Assert
	assert.Equal(t, expected, result)
}

func TestToFloat(t *testing.T) {
	//Arrange
	f, _ := NewFraction(1, 2)
	var expected float32 = 0.5

	//Act
	result := f.ToFloat()

	//Assert
	assert.Equal(t, expected, result)
}
