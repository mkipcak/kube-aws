package model

import (
	"testing"
)

func TestRef(t *testing.T) {
	i := Identifier{ID: "vpc-xxxx"}
	if ref := (i.Ref(func() string { return "" })); ref != `"vpc-xxxx"` {
		t.Errorf("Ref incorrect for id: %s", ref)
	}

	i = Identifier{IDFromStackOutput: "output-key"}
	if ref := (i.Ref(func() string { return "" })); ref != `{ "Fn::ImportValue" : "output-key" }` {
		t.Errorf("Ref incorrect for idfromStackOutput: %s", ref)
	}

	i = Identifier{IDFromFn: `{ "Fn::Ref" : foo }`}
	if ref := (i.Ref(func() string { return "" })); ref != `{ "Fn::Ref" : foo }` {
		t.Errorf("Ref incorrect for idFromFn: %s", ref)
	}
}
