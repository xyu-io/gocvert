# Gocvert

## Description

data converter implement of golang.

+ Supports update data field value based on struct data types.

+ Supports exchanging data field values based on tags of two struct data types.


## Exp
```go
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

	data := TestCase{
		ID:   1,
		Name: "test",
		Age:  23,
	}
	data2 := TestCaseTo{
		UID:   101,
		UName: "default",
		UAge:  10,
	}
	var tmp1 = data

	err := RewriteFields(&data, opt)
	if err != nil {
		t.Error(err)
		return
	}

	err = SwapTags(&tmp1, &data2, opto)
	if err != nil {
		t.Error(err)
		return
	}
```
