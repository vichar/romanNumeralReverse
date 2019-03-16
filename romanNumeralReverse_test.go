package romanNumeralReverse_test

import (
	"testing"

	"github.com/vichar/romanNumeralReverse"
)

func Test(t *testing.T) {
	t.Run("input number as string I should return int 1", func(t *testing.T) {
		get := romanNumeralReverse.Parse("I")
		want := 1
		if want != get {
			t.Errorf("It should say %d but get %d", want, get)
		}
	})
	t.Run("input number as string II should return int 2", func(t *testing.T) {
		get := romanNumeralReverse.Parse("II")
		want := 2
		if want != get {
			t.Errorf("It should say %d but get %d", want, get)
		}
	})
	t.Run("input number as string III should return int 3", func(t *testing.T) {
		get := romanNumeralReverse.Parse("III")
		want := 3
		if want != get {
			t.Errorf("It should say %d but get %d", want, get)
		}
	})
	t.Run("input number as string IV should return int 4", func(t *testing.T) {
		get := romanNumeralReverse.Parse("IV")
		want := 4
		if want != get {
			t.Errorf("It should say %d but get %d", want, get)
		}
	})
	t.Run("input number as string V should return int 5", func(t *testing.T) {
		get := romanNumeralReverse.Parse("V")
		want := 5
		if want != get {
			t.Errorf("It should say %d but get %d", want, get)
		}
	})
	t.Run("input number as string VI should return int 6", func(t *testing.T) {
		get := romanNumeralReverse.Parse("VI")
		want := 6
		if want != get {
			t.Errorf("It should say %d but get %d", want, get)
		}
	})
}
