package vda5050

import (
	"encoding/json"
	"testing"
)

func TestFactsheetActionFlagsAcceptBooleanAndString(t *testing.T) {
	for _, tc := range []struct {
		name, pause, cancel string
		wantPause, wantCancel bool
	}{
		{"booleans (3.0 PDF)", `true`, `false`, true, false},
		{"strings (3.0 JSON schema)", `"false"`, `"true"`, false, true},
		{"upper-case strings", `"TRUE"`, `"FALSE"`, true, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			payload := `{"protocolFeatures":{"mobileRobotActions":[{"actionType":"pick","actionScopes":["NODE"],` +
				`"pauseAllowed":` + tc.pause + `,"cancelAllowed":` + tc.cancel + `}]}}`
			var fs Factsheet
			if err := json.Unmarshal([]byte(payload), &fs); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			a := fs.ProtocolFeatures.MobileRobotActions[0]
			if bool(a.PauseAllowed) != tc.wantPause || bool(a.CancelAllowed) != tc.wantCancel {
				t.Fatalf("pauseAllowed=%v cancelAllowed=%v, want %v %v",
					a.PauseAllowed, a.CancelAllowed, tc.wantPause, tc.wantCancel)
			}
		})
	}
}

func TestLenientBoolRejectsOtherValuesAndMarshalsAsBoolean(t *testing.T) {
	for _, in := range []string{`"yes"`, `1`, `null`, `{}`} {
		var b LenientBool
		if err := json.Unmarshal([]byte(in), &b); err == nil && in != `null` {
			t.Errorf("accepted %s", in)
		}
	}
	out, err := json.Marshal(MobileRobotAction{ActionType: "pick", PauseAllowed: true})
	if err != nil {
		t.Fatal(err)
	}
	var doc map[string]any
	if err := json.Unmarshal(out, &doc); err != nil {
		t.Fatal(err)
	}
	if doc["pauseAllowed"] != true || doc["cancelAllowed"] != false {
		t.Fatalf("not marshalled as booleans: %s", out)
	}
}
