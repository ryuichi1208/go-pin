if err != nil {
	return err
}


func TestHogeee(t *testing.T) {
	
}

var tests = []struct {
	name string
	expected string
	given string
}{
	{"", "aaa", "",},
}
for _, tt := range tests {
	tt := tt
	t.Run(tt.name, func(t *testing.T){
		actual := (tt.given)
		if actual != tt.expected {
				t.Errorf("(%s): expected %s, actual %s", tt.given, tt.expected, actual)
		}

	})
}

go func() {
	
}()

