package conv

import (
	"math"
	"reflect"
	"testing"
)

func TestFarhenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
	}

	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		got = (math.Round(got*100) / 100)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 56.67, want: 134.01},
	}

	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
    got = (math.Round(got*100) / 100)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 200, want: -73.15},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
    got = (math.Round(got*100) / 100)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: -73.15, want: 200},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
    got = (math.Round(got*100) / 100)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestKelvinToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 300, want: 80.33},
	}

	for _, tc := range tests {
		got := KelvinToFahrenheit(tc.input)
    got = (math.Round(got*100) / 100)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestFarhenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 80.33, want: 360.14},
	}

	for _, tc := range tests {
		got := FahrenheitToKelvin(tc.input)
    got = (math.Round(got*100) / 100)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}


