package fraction

import (
	"errors"
	"fmt"
)

// Fraction represents a fraction with a numerator and a denominator
// (example 1/2 => numerator=1, denominator=2 in float 0.5)
type Fraction struct {
	numerator   int
	denominator int
}

func (c *Fraction) validator() error {
	if c.denominator == 0 {
		return errors.New("Denominator cannot be zero")
	}
	return nil
}

func create(numerator, denominator int) *Fraction {
	f := &Fraction{
		numerator:   numerator,
		denominator: denominator,
	}

	return f
}

func (f *Fraction) shiftRight() *Fraction {
	return create(f.denominator, f.numerator)
}

// NewFraction creates a new Fraction instance.
// It returns an error if the denominator is zero.
func NewFraction(numerator, denominator int) (*Fraction, error) {
	f := create(numerator, denominator)

	err := f.validator()
	if err != nil {
		return nil, err
	}

	return f, nil
}

// Numerator returns the numerator of the fraction
// (example 1/2 => 1)
func (f *Fraction) Numerator() int {
	return f.numerator
}

// Denominator returns the denominator of the fraction
// (example 1/2 => 2)
func (f *Fraction) Denominator() int {
	return f.denominator
}

// Negativate returns the negativated fraction
// (example 1/2 => -1/2)
func (f *Fraction) Negativate() *Fraction {
	return create(f.numerator*-1, f.denominator)
}

// Add adds two fractions and returns the result
// It returns an error if result with the denominator is zero.
func (fright *Fraction) Add(fleft *Fraction) (*Fraction, error) {
	newNumerator, newDenominator := func() (int, int) {
		if fright.denominator == fleft.denominator {
			return (fright.numerator + fleft.numerator), fright.denominator
		}
		return (fright.numerator*fleft.denominator + fleft.numerator*fright.denominator), (fright.denominator * fleft.denominator)
	}()
	return NewFraction(newNumerator, newDenominator)
}

// Subtract subtracts two fractions and returns the result
// It returns an error if result with the denominator is zero.
func (fright *Fraction) Subtract(fleft *Fraction) (*Fraction, error) {
	return fright.Add(fleft.Negativate())
}

// Multiply multiplies two fractions and returns the result
// It returns an error if result with the denominator is zero.
func (fright *Fraction) Multiply(fleft *Fraction) (*Fraction, error) {
	newNumerator := fright.numerator * fleft.numerator
	newDenominator := fright.denominator * fleft.denominator
	return NewFraction(newNumerator, newDenominator)
}

// Divide divides two fractions and returns the result
// It returns an error if result with the denominator is zero.
func (fright *Fraction) Divide(fleft *Fraction) (*Fraction, error) {
	return fright.Multiply(fleft.shiftRight())
}

// ToString returns the string representation of the fraction
// (example 1/2 => "1/2")
func (f *Fraction) ToString() string {
	return fmt.Sprintf("%d/%d", f.numerator, f.denominator)
}

// ToFloat returns the float representation of the fraction
// (example 1/2 => 0.5)
func (f *Fraction) ToFloat() float32 {
	return float32(f.numerator) / float32(f.denominator)
}
