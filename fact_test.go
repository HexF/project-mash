package main

import (
	"html/template"
	"testing"
)

func TestFacts(t *testing.T) {
	table := []struct {
		mash   string
		expect string
	}{
		{
			mash:   "hexf",
			expect: "<a href=\"https://hexf.me\">Visit my website</a>",
		},
	}

	for _, test := range table {
		fact := getFact(test.mash)
		if fact != template.HTML(test.expect) {
			t.Errorf("Fact for '%v' was incorrect, got: %v, want: %v", test.mash, fact, test.expect)
		}
	}

}
