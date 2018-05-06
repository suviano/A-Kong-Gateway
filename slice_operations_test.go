package main

import "testing"

func Test_emptyValuesInSlice(t *testing.T) {
	t.Run("not empty value return list", func(t *testing.T) {
		hasEmpty, result, err := emptyValuesInSlice("lol", "he")
		if err != nil {
			t.Errorf("should not have any error for given values '%+v'\n", err)
		}

		if len(result) != 2 {
			t.Error("result should have 2 itens")
		}

		if hasEmpty {
			t.Error("not empty itens on the received params")
		}
	})
}

func Test_forEach(t *testing.T) {
	t.Run("iterator applied function with success", func(t *testing.T) {
		dummyFunc := func(item string) bool {
			return true
		}
		result, err := forEach(dummyFunc, "aaa", "aaaa", "item")
		if err != nil {
			t.Errorf("while executing iterator throwed error [%+v]", err)
		}
		if len(result) != 3 {
			t.Errorf("iterator have no applied right function for each item,"+
				" expected %d itens returned %d", len(result), 2)
		}
	})

	t.Run("iterator applied function without success", func(t *testing.T) {
		dummyFunc := func(item string) bool {
			return item == ""
		}
		result, err := forEach(dummyFunc, "", "", "item")
		if err == nil {
			t.Errorf("iterator should return error")
		}

		if len(result) != 2 {
			t.Errorf("iterator have no applied right function for each item,"+
				" expected %d itens returned %d", len(result), 2)
		}
	})
}
