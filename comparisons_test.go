package main

import (
	"reflect"
	"testing"
)

func TestSortMatchers(t *testing.T) {
	testMatchers := []matcher{
		{Name: "G"},
		{Name: "Z"},
		{Name: "A"},
	}

	expected := []matcher{
		{Name: "A"},
		{Name: "G"},
		{Name: "Z"},
	}

	sortMatchers(testMatchers)

	if !reflect.DeepEqual(testMatchers, expected) {
		t.Errorf("got %+v but want %+v", testMatchers, expected)
	}
}

func TestMatchersCompare(t *testing.T) {
	t.Run("unequal slices", func(t *testing.T) {
		a := []matcher{
			{Name: "a"},
			{Name: "b"},
		}

		b := []matcher{
			{Name: "a"},
		}

		got := matchersCompare(a, b)
		if got != false {
			t.Errorf("got %v but want %v", got, false)
		}
	})
	t.Run("regex false", func(t *testing.T) {
		a := []matcher{
			{IsRegex: false, Name: "a"},
		}

		b := []matcher{
			{IsRegex: true, Name: "a"},
		}

		got := matchersCompare(a, b)
		if got != false {
			t.Errorf("got %v but want %v", got, false)
		}
	})

	t.Run("name false", func(t *testing.T) {
		a := []matcher{
			{IsRegex: false, Name: "a"},
		}

		b := []matcher{
			{IsRegex: false, Name: "b"},
		}

		got := matchersCompare(a, b)
		if got != false {
			t.Errorf("got %v but want %v", got, false)
		}
	})

	t.Run("value false", func(t *testing.T) {
		a := []matcher{
			{IsRegex: true, Name: "a", Value: "a value"},
		}
		c := []matcher{
			{IsRegex: true, Name: "a", Value: "b value"},
		}

		got := matchersCompare(a, c)
		if got != false {
			t.Errorf("got %v want %v", got, true)
		}
	})

	t.Run("true", func(t *testing.T) {
		a := []matcher{
			{IsRegex: true, Name: "a", Value: "a value"},
		}
		c := []matcher{
			{IsRegex: true, Name: "a", Value: "a value"},
		}

		got := matchersCompare(a, c)
		if got != true {
			t.Errorf("got %v want %v", got, true)
		}
	})
}

func TestCompareSilences(t *testing.T) {
	t.Run("not found", func(t *testing.T) {
		d := []scheduledSilence{
			{Matchers: []matcher{
				{IsRegex: false, Name: "a", Value: "a value"},
				{IsRegex: true, Name: "b", Value: "b value"}},
			},
		}

		a := []alertmanagerSilence{
			{Matchers: []matcher{
				{IsRegex: false, Name: "a", Value: "a value"}},
			},
		}

		want := []alertmanagerSilence{
			{Matchers: []matcher{
				{IsRegex: false, Name: "a", Value: "a value"},
				{IsRegex: true, Name: "b", Value: "b value"},
			}},
		}

		got := compareSilences(a, d)
		for i := range got {
			if !reflect.DeepEqual(got[i].Matchers, want[i].Matchers) {
				t.Errorf("got %v but want %v", got, want)
			}
		}
	})

	t.Run("found", func(t *testing.T) {
		d := []scheduledSilence{
			{Matchers: []matcher{
				{IsRegex: false, Name: "a", Value: "a value"}},
			},
		}
		a := []alertmanagerSilence{
			{Matchers: []matcher{
				{IsRegex: false, Name: "a", Value: "a value"}},
			}}

		want := []alertmanagerSilence{
			{Matchers: []matcher{
				{IsRegex: false, Name: "a", Value: "a value"},
			}},
		}

		got := compareSilences(a, d)
		for i := range got {
			if !reflect.DeepEqual(got[i].Matchers, want[i].Matchers) {
				t.Errorf("got %v but want %v", got, want)
			}
		}
	})
}
