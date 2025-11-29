package csvmum_test

import (
	"bytes"
	"fmt"
	"io"

	"github.com/seannyphoenix/bogie/pkg/csvmum"
)

func ExampleNewUnmarshaler() {
	type person struct {
		Name string `csv:"name"`
		Age  int    `csv:"age"`
	}

	r := bytes.NewBuffer([]byte("name,age\nNobody,0\nSpot,2\n"))
	csvu, err := csvmum.NewUnmarshaler[person](r)
	if err != nil {
		panic(err)
	}

	pp := []person{}
	for {
		var p person
		err = csvu.Unmarshal(&p)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		pp = append(pp, p)
	}

	fmt.Println(pp)

	// Output: [{Nobody 0} {Spot 2}]
}
