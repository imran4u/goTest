package calculator

import (
	"fmt"
	"imran/calculator"
	"testing" // go inbuild library
)

func TestAdd(t *testing.T) {
	result := calculator.Add(6, 3)
	if result != 9 {
		t.Errorf("Add(6,3) = %d; want 9", result)
	}
	fmt.Println("Add(6,3) = ", result)
}

func TestSub(t *testing.T) {
	result := calculator.Sub(6, 3)
	if result != 3 {
		t.Errorf("Sub(6,3) = %d; want 3", result)
	}
	fmt.Println("Sub(6,3) = ", result)
}
func TestMul(t *testing.T) {
	result := calculator.Mul(6, 3)
	if result != 18 {
		t.Errorf("Mul(6,3) = %d; want 18", result)
	}
	fmt.Println("Mul(6,3) = ", result)
}
func TestDiv(t *testing.T) {
	result := calculator.Div(6, 3)
	if result != 2 {
		t.Errorf("Div(6,3) = %d; want 2", result)
	}
	fmt.Println("Div(6,3) = ", result)
}

func TestDivForZeroResult(t *testing.T) {
	result := calculator.Div(2, 3)
	if result != 0 {
		t.Errorf("Div(2,3) = %d; want 0", result)
	}
	fmt.Println("Div(2,3) = ", result)
}
