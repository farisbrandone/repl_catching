package main

import (
	"testing"
)

func TestCleanInput(t *testing.T){
cases:=[]struct{
	input string
	expected []string
}{
	{
		input: "hello word",
		expected:[]string{
			"hello",
			"word",
		},
	},
	{
		input: "HEllo word",
		expected:[]string{
			"hello",
			"word",
		},
	},
}
for _,cse:=range cases{
	actual:=cleanInput(cse.input)
	if len(actual)!=len(cse.expected){
		t.Errorf("the length are not equal: %v vs %v", 
		len(actual), 
		len(cse.expected),
	)
	
		continue
	}
	for i := range actual {
		actualWord:=actual[i]
		expectedWord:=cse.expected[i]
		if actualWord!=expectedWord {
			t.Errorf("%v does not equal %v", 
			actualWord,
			 expectedWord,)
		}

		}
}

}