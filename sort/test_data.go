package sort

type testCaseInt struct {
	in   []int
	want []int
}

type testCaseString struct {
	in   []string
	want []string
}

type testCaseAny struct {
	in   []interface{}
	want []interface{}
}

var intCases = []testCaseInt{
	{[]int{1, 0}, []int{0, 1}},
	{[]int{5, 4, 2, 3}, []int{2, 3, 4, 5}},
	{[]int{15, 12, 5, 7, 4, 9, 2, 3, 0}, []int{0, 2, 3, 4, 5, 7, 9, 12, 15}},
	{[]int{15, 12, -5, 7, 0, 9, -2, 3, 0}, []int{-5, -2, 0, 0, 3, 7, 9, 12, 15}},
}

var strCases = []testCaseString{
	{
		[]string{"Fedora", "visual", "C++"},
		[]string{"C++", "Fedora", "visual"},
	},
	{
		[]string{"command", "keyboard", "angular", "mern", "browserify"},
		[]string{"angular", "browserify", "command", "keyboard", "mern"},
	},
}

func loadIntCases() []testCaseAny {
	cases := make([]testCaseAny, len(intCases))
	for i, c := range intCases {
		cases[i].in = make([]interface{}, len(c.in))
		for j, u := range c.in {
			cases[i].in[j] = u
		}
		cases[i].want = make([]interface{}, len(c.want))
		for j, u := range c.want {
			cases[i].want[j] = u
		}
	}
	return cases
}

func loadStringCases() []testCaseAny {
	cases := make([]testCaseAny, len(strCases))
	for i, c := range strCases {
		cases[i].in = make([]interface{}, len(c.in))
		for j, u := range c.in {
			cases[i].in[j] = u
		}
		cases[i].want = make([]interface{}, len(c.want))
		for j, u := range c.want {
			cases[i].want[j] = u
		}
	}
	return cases
}
