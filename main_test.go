package main

import (
	"sec/calculate"
	"testing"
)

func TestPlus(t *testing.T) {
	tests := []struct {
		name      string
		a         int
		b         int
		expResult int
	}{
		{"Положительные числа", 5, 5, 10},
		{"Отрицательные числа", -100, -500, -600},
		{"Сложение с цифрой 0", 1, 0, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculate.Add(tt.a, tt.b)
			if result != tt.expResult {
				t.Errorf("Ошибка! %d + %d = ?? Ожидался ответ: %d, получили: %d", tt.a, tt.b, tt.expResult, result)
			}
		})
	}
}

func TestMinus(t *testing.T) {
	tests := []struct {
		name      string
		a, b      int
		expResult int
	}{
		{"Положительные числа", 10, 5, 5},
		{"Отрицательые числа", -10, -5, -5},
		{"Разные знаки", 10, -5, 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := calculate.Subtract(tt.a, tt.b)
			if res != tt.expResult {
				t.Errorf("Ошибка! %d - %d = ?? Ожидался ответ: %d, получили: %d", tt.a, tt.b, tt.expResult, res)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name            string
		a, b, expResult int
	}{
		{"Положительные числа", 10, 5, 50},
		{"Отрицательые числа", -10, -5, 50},
		{"Разные знаки", 10, -5, -50},
		{"Умножение на 0", 10, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := calculate.Multiply(tt.a, tt.b)
			if res != tt.expResult {
				t.Errorf("Ошибка! %d * %d = ?? Ожидался ответ: %d, получили: %d", tt.a, tt.b, tt.expResult, res)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name      string
		a, b      int
		expResult float64
		expError  bool
	}{
		{"Положительные числа", 10, 5, 2.0, false},
		{"Отрицательые числа", -10, -5, 2.0, false},
		{"Разные знаки", 10, -5, -2.0, false},
		{"Деление на 0", 10, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := calculate.Divide(tt.a, tt.b)
			if (err != nil) != tt.expError {
				t.Errorf("%d / %d = ?? Неожиданная ошибка: %v", tt.a, tt.b, err)
			}

			if tt.expError {
				return
			}

			if res != tt.expResult {
				t.Errorf("Ошибка! %d / %d = ?? Ожидался ответ: %.2f, получили: %.2f", tt.a, tt.b, tt.expResult, res)
			}
		})
	}
}
