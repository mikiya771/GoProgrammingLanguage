package pack

import "testing"

type example struct {
	Labels     []string `http:"1"`
	MaxResults int      `http:"max"`
	Exact      bool     `http:"x"`
}

func TestPack(t *testing.T) {
	table := []struct {
		input  example
		expect string
	}{
		{
			example{
				Labels:     []string{"a", "b", "c"},
				MaxResults: 10,
				Exact:      true,
			},
			"1=a&1=b&1=c&max=10&x=true",
		},
	}
	for _, td := range table {
		act := Pack(td.input)
		if td.expect != act {
			t.Errorf("Pack(td.input)=%s, but want %s ", act, td.expect)
		}
	}
}
