package main

func main() {
	m := newIO()
	defer func() {
		err := m.closeAll()
		if err != nil {
			panic(err)
		}
	}()

	if err := m.process(); err != nil {
		panic(err)
	}
}
