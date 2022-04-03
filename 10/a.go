package main

func a() error {
	return nil
}

func b() error {
	err := a()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	_ := b()

}
