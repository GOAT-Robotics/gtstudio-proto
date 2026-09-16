package vda5050

import (
	"encoding/json"
	"strings"
	"testing"
)

var dialectTestID = Identity{Manufacturer: "ACME", SerialNumber: "R-001"}

func customV2Dialect() *V2WireAdapter {
	return NewV2WireAdapter(V2WireOptions{
		Name: "custom-v2", VersionLevel: "v2.0.0", DefaultVersion: "2.0.0",
		TopicAliases:        map[Topic]string{TopicConnection: "connect"},
		HeaderVersionPrefix: "v", NumericConnectionState: map[int]ConnectionState{2: ConnectionStateOnline},
	})
}

func TestConfiguredV2DialectNormalizesState(t *testing.T) {
	a := customV2Dialect()
	payload := []byte(`{"actionStates":[],"agvPosition":{"localizationScore":1,"mapId":"L1","positionInitialized":true,"theta":1.2,"x":3,"y":4},"batteryState":{"batteryCharge":14,"batteryHealth":99,"batteryVoltage":43.5,"charging":false},"driving":false,"edgeStates":[],"errors":[],"headerId":565,"information":[],"lastNodeId":"","lastNodeSequenceId":0,"loads":[],"manufacturer":"ACME","nodeStates":[],"operatingMode":"MANUAL","orderId":"","orderUpdateId":0,"safetyState":{"eStop":"NONE","fieldViolation":false},"serialNumber":"R-001","timestamp":"2026-09-16T17:31:29.059Z","velocity":{"omega":0,"vx":0,"vy":0},"version":"v2.0.0"}`)
	decoded, err := a.Decode("uagv/v2.0.0/ACME/R-001/state", payload)
	if err != nil {
		t.Fatal(err)
	}
	if decoded.Version != "2.0.0" || decoded.Identity != dialectTestID {
		t.Fatalf("unexpected envelope: %+v", decoded)
	}
	var state State
	if err := json.Unmarshal(decoded.Payload, &state); err != nil {
		t.Fatal(err)
	}
	if state.MobileRobotPosition == nil || !state.MobileRobotPosition.Localized || state.MobileRobotPosition.MapID != "L1" {
		t.Fatalf("position was not normalized: %+v", state.MobileRobotPosition)
	}
	if state.PowerSupply.StateOfCharge != 14 || state.SafetyState.ActiveEmergencyStop != ActiveEmergencyStopNone {
		t.Fatalf("state was not normalized: %+v", state)
	}
}

func TestConfiguredConnectionDialect(t *testing.T) {
	a := customV2Dialect()
	decoded, err := a.Decode("uagv/v2.0.0/ACME/R-001/connect", []byte(`{"connectionState":2}`))
	if err != nil {
		t.Fatal(err)
	}
	var connection Connection
	if err := json.Unmarshal(decoded.Payload, &connection); err != nil {
		t.Fatal(err)
	}
	if connection.ConnectionState != ConnectionStateOnline || connection.Version != "2.0.0" {
		t.Fatalf("connection was not normalized: %+v", connection)
	}
}

func TestV2OrderEgressAddsLegacyShape(t *testing.T) {
	a := NewV2WireAdapter(V2WireOptions{})
	order := &Order{
		HeaderID: 1, Timestamp: "2026-09-16T00:00:00Z", Version: ProtocolVersion,
		Manufacturer: "KIT", SerialNumber: "1", OrderID: "o1", Nodes: []Node{
			{NodeID: "n1", SequenceID: 0, Released: true, Actions: []Action{}},
			{NodeID: "n2", SequenceID: 2, Released: true, Actions: []Action{}},
		}, Edges: []Edge{{EdgeID: "e1", SequenceID: 1, Released: true, Actions: []Action{}}},
	}
	wire, err := a.Encode(Identity{Manufacturer: "KIT", SerialNumber: "1"}, "2.0.0", TopicOrder, order)
	if err != nil {
		t.Fatal(err)
	}
	if wire.Topic != "uagv/v2/KIT/1/order" {
		t.Fatalf("unexpected topic %q", wire.Topic)
	}
	var body map[string]any
	if err := json.Unmarshal(wire.Payload, &body); err != nil {
		t.Fatal(err)
	}
	if body["version"] != "2.0.0" {
		t.Fatalf("unexpected version %v", body["version"])
	}
	edge := object(array(body["edges"])[0])
	if edge["startNodeId"] != "n1" || edge["endNodeId"] != "n2" {
		t.Fatalf("legacy endpoints missing: %v", edge)
	}
}

func TestV2InstantActionVersionDifference(t *testing.T) {
	a := NewV2WireAdapter(V2WireOptions{})
	msg := &InstantActions{Actions: []InstantAction{{ActionType: ActionInitializePosition, ActionID: "a1", BlockingType: InstantActionBlockingTypeNone}}}
	for _, tc := range []struct{ version, key string }{{"2.0.0", "actionName"}, {"2.1.0", "actionType"}} {
		wire, err := a.Encode(Identity{Manufacturer: "KIT", SerialNumber: "1"}, tc.version, TopicInstantActions, msg)
		if err != nil {
			t.Fatal(err)
		}
		var body map[string]any
		if err := json.Unmarshal(wire.Payload, &body); err != nil {
			t.Fatal(err)
		}
		action := object(array(body["actions"])[0])
		if action[tc.key] != "initPosition" {
			t.Fatalf("%s action not converted: %v", tc.version, action)
		}
		other := "actionType"
		if tc.key == other {
			other = "actionName"
		}
		if _, exists := action[other]; exists {
			t.Fatalf("%s retained wrong key: %v", tc.version, action)
		}
	}
}

func TestDefaultWireAdaptersExposeSupportedVersions(t *testing.T) {
	var got []string
	for _, adapter := range DefaultWireAdapters("") {
		got = append(got, adapter.Name()+":"+strings.Join(adapter.SupportedVersions(), ","))
	}
	joined := strings.Join(got, " ")
	for _, want := range []string{"vda5050-v3:3.0.0", "vda5050-v2:2.0.0,2.1.0"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing %q in %q", want, joined)
		}
	}
}
