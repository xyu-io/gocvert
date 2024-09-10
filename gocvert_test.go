package gocvert

import "testing"

type TestCase0 struct {
	ID   int
	Name string
	Age  int
}

type TestCase1 struct {
	UID   int
	UName string
	UAge  int
}

func TestGoCvert(t *testing.T) {
	opt := []FieldOption{
		{
			Tag:      "Name",
			NewField: "test-convert",
		},
		{
			Tag:      "Age",
			NewField: 18,
		},
	}

	opto := []SwapOption{
		{
			Tag:    "Name",
			NewTag: "UName",
		},
		{
			Tag:    "Age",
			NewTag: "UAge",
		},
	}

	data := TestCase0{
		ID:   1,
		Name: "test",
		Age:  23,
	}
	data2 := TestCase1{
		UID:   101,
		UName: "default",
		UAge:  10,
	}

	var tmp1 = data
	var tmp2 = data2

	err := RewriteFields(&data, opt)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v convert values  >> %+v \n", tmp1, data)

	err = SwapWithTags(&tmp1, &data2, opto)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v swap tags to %+v >> %+v \n", tmp1, tmp2, data2)
}

func BenchmarkFieldGoCvert(b *testing.B) {
	opt := []FieldOption{
		{
			Tag:      "Name",
			NewField: "test-convert",
		},
		{
			Tag:      "Age",
			NewField: 18,
		},
	}

	data := TestCase0{
		ID:   1,
		Name: "test",
		Age:  23,
	}

	for i := 0; i < b.N; i++ {
		err := RewriteFields(&data, opt)
		if err != nil {
			b.Error(err)
			return
		}
	}
}

func BenchmarkSwapGoCvert(b *testing.B) {
	opto := []SwapOption{
		{
			Tag:    "Name",
			NewTag: "UName",
		},
		{
			Tag:    "Age",
			NewTag: "UAge",
		},
	}

	data1 := TestCase0{
		ID:   1,
		Name: "test",
		Age:  23,
	}

	data2 := TestCase1{
		UID:   101,
		UName: "default",
		UAge:  10,
	}
	for i := 0; i < b.N; i++ {
		err := SwapWithTags(&data1, &data2, opto)
		if err != nil {
			b.Error(err)
			return
		}
	}
}
