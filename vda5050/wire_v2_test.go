package vda5050

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

var v2TestID = Identity{Manufacturer: "KIT", SerialNumber: "1"}

// validateV2 checks an encoded message against an official VDA 2.x schema kept
// under testdata/<schemas>. "2.0.0" holds the 2024 corrected schemas and
// "2.0.0-original" the ones first published with 2.0.0.
func validateV2(t *testing.T, schemas string, topic Topic, payload []byte) {
	t.Helper()
	raw, err := os.ReadFile("testdata/" + schemas + "/" + string(topic) + ".schema")
	if err != nil {
		t.Fatal(err)
	}
	doc, err := jsonschema.UnmarshalJSON(bytes.NewReader(raw))
	if err != nil {
		t.Fatal(err)
	}
	c := jsonschema.NewCompiler()
	if err := c.AddResource("schema.json", doc); err != nil {
		t.Fatal(err)
	}
	schema, err := c.Compile("schema.json")
	if err != nil {
		t.Fatal(err)
	}
	inst, err := jsonschema.UnmarshalJSON(bytes.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	if err := schema.Validate(inst); err != nil {
		t.Fatalf("%s %s does not match the official schema: %v\n%s", schemas, topic, err, payload)
	}
}

func decodeV2State(t *testing.T, a *V2WireAdapter, payload string) State {
	t.Helper()
	decoded, err := a.Decode("uagv/v2/KIT/1/state", []byte(payload))
	if err != nil {
		t.Fatal(err)
	}
	var st State
	if err := json.Unmarshal(decoded.Payload, &st); err != nil {
		t.Fatal(err)
	}
	return st
}

const v2StateTemplate = `{"headerId":1,"timestamp":"2026-09-16T00:00:00Z","version":"%s","manufacturer":"KIT","serialNumber":"1",
"orderId":"o1","orderUpdateId":0,"lastNodeId":"n1","lastNodeSequenceId":0,"nodeStates":[],"edgeStates":[],"driving":false,
"operatingMode":"%s","errors":[],"information":[],"loads":[],
"batteryState":{"batteryCharge":50,"charging":false},"safetyState":{"eStop":"%s","fieldViolation":false},
"actionStates":%s}`

func v2State(version, mode, eStop, actionStates string) string {
	return fmt.Sprintf(v2StateTemplate, version, mode, eStop, actionStates)
}

func v2Order() *Order {
	return &Order{
		HeaderID: 1, Timestamp: "2026-09-16T00:00:00Z", Version: ProtocolVersion,
		Manufacturer: "KIT", SerialNumber: "1", OrderID: "o1", OrderDescription: Str("demo"),
		Nodes: []Node{
			{NodeID: "n1", SequenceID: 0, Released: true, Actions: []Action{},
				NodePosition: &NodePosition{X: 1, Y: 2, MapID: "L1", AllowedDeviationXY: &AllowedDeviationXY{A: 0.5, B: 0.3}}},
			{NodeID: "n2", SequenceID: 2, Released: true, Actions: []Action{
				{ActionType: ActionPick, ActionID: "a1", BlockingType: BlockingTypeSingle, Retriable: Bool(true)},
			}},
		},
		Edges: []Edge{{EdgeID: "e1", SequenceID: 1, Released: true, Actions: []Action{}}},
	}
}

// v2Schema pairs a wire profile with the official schemas it must satisfy.
type v2Schema struct {
	schemas, version string
	original200      bool
	deviationKey     string
	instantKey       string
}

var v2Schemas = []v2Schema{
	{"2.0.0", "2.0.0", false, "allowedDeviationXY", "actionType"},
	{"2.0.0-original", "2.0.0", true, "allowedDeviationXy", "actionName"},
	{"2.1.0", "2.1.0", false, "allowedDeviationXY", "actionType"},
}

func (s v2Schema) adapter() *V2WireAdapter {
	return NewV2WireAdapter(V2WireOptions{OriginalV200Schema: s.original200})
}

func encodeV2(t *testing.T, s v2Schema, topic Topic, msg any) map[string]any {
	t.Helper()
	wire, err := s.adapter().Encode(v2TestID, s.version, topic, msg)
	if err != nil {
		t.Fatal(err)
	}
	validateV2(t, s.schemas, topic, wire.Payload)
	var body map[string]any
	if err := json.Unmarshal(wire.Payload, &body); err != nil {
		t.Fatal(err)
	}
	return body
}

func TestV2OrderEgressMatchesOfficialSchemas(t *testing.T) {
	for _, s := range v2Schemas {
		body := encodeV2(t, s, TopicOrder, v2Order())
		if _, ok := body["orderDescription"]; ok {
			t.Fatalf("%s kept VDA 3 orderDescription", s.schemas)
		}
		position := object(object(array(body["nodes"])[0])["nodePosition"])
		if position[s.deviationKey] != 0.3 {
			t.Fatalf("%s deviation ellipse not reduced to its smaller axis: %v", s.schemas, position)
		}
		action := object(array(object(array(body["nodes"])[1])["actions"])[0])
		if action["blockingType"] != "HARD" {
			t.Fatalf("%s SINGLE not mapped to HARD: %v", s.schemas, action)
		}
	}
}

func TestV2InstantActionsMatchOfficialSchemas(t *testing.T) {
	msg := &InstantActions{HeaderID: 1, Timestamp: "2026-09-16T00:00:00Z", Actions: []InstantAction{
		InitializePosition(1, 2, 0, "L1", "n1"), CancelOrder("o1"),
	}}
	for _, s := range v2Schemas {
		action := object(array(encodeV2(t, s, TopicInstantActions, msg)["actions"])[0])
		if action[s.instantKey] != "initPosition" {
			t.Fatalf("%s instant action not keyed by %s: %v", s.schemas, s.instantKey, action)
		}
	}
}

func TestV21KeepsCorridorAndGlobalOrientation(t *testing.T) {
	order := v2Order()
	order.Edges[0].Orientation = Float(1.57)
	global, kinematic := OrientationTypeGlobal, CorridorReferencePoint("KINEMATIC_CENTER")
	order.Edges[0].OrientationType = &global
	order.Edges[0].Corridor = &Corridor{LeftWidth: 0.4, RightWidth: 0.6, CorridorReferencePoint: &kinematic}

	edge := object(array(encodeV2(t, v2Schemas[2], TopicOrder, order)["edges"])[0])
	corridor := object(edge["corridor"])
	if edge["orientationType"] != "GLOBAL" || corridor == nil || corridor["corridorRefPoint"] != "KINEMATICCENTER" {
		t.Fatalf("2.1.0 lost edge constraints: %v", edge)
	}

	// The corrected 2.0.0 schema has orientationType but no corridor.
	edge = object(array(encodeV2(t, v2Schemas[0], TopicOrder, order)["edges"])[0])
	if _, ok := edge["corridor"]; ok || edge["orientationType"] != "GLOBAL" {
		t.Fatalf("2.0.0 edge constraints wrong: %v", edge)
	}

	// The original 2.0.0 schema cannot express a GLOBAL orientation; sending
	// it tangentially would turn the vehicle the wrong way.
	if _, err := v2Schemas[1].adapter().Encode(v2TestID, "2.0.0", TopicOrder, order); err == nil {
		t.Fatal("original 2.0.0 schema accepted a GLOBAL orientation")
	}
}

func TestV21RejectsCorridorRelease(t *testing.T) {
	order := v2Order()
	order.Edges[0].Corridor = &Corridor{LeftWidth: 0.4, RightWidth: 0.6, ReleaseRequired: Bool(true)}
	if _, err := NewV2WireAdapter(V2WireOptions{}).Encode(v2TestID, "2.1.0", TopicOrder, order); err == nil {
		t.Fatal("corridor release silently dropped for 2.1.0")
	}
}

func TestV2InstantActionStatesAreSplitOut(t *testing.T) {
	a := NewV2WireAdapter(V2WireOptions{})
	cancel := CancelOrder("o1")
	if _, err := a.Encode(v2TestID, "2.1.0", TopicInstantActions, &InstantActions{Actions: []InstantAction{cancel}}); err != nil {
		t.Fatal(err)
	}
	st := decodeV2State(t, a, v2State("2.1.0", "AUTOMATIC", "NONE", `[
		{"actionId":"a1","actionType":"pick","actionStatus":"RUNNING"},
		{"actionId":"`+cancel.ActionID+`","actionType":"cancelOrder","actionStatus":"FINISHED"}]`))
	if len(st.ActionStates) != 1 || st.ActionStates[0].ActionID != "a1" {
		t.Fatalf("order action states wrong: %+v", st.ActionStates)
	}
	if len(st.InstantActionStates) != 1 || st.InstantActionStates[0].ActionID != cancel.ActionID {
		t.Fatalf("instant action states wrong: %+v", st.InstantActionStates)
	}

	// Another vehicle's state is not affected by this vehicle's instant actions.
	other, err := a.Decode("uagv/v2/KIT/2/state", []byte(strings.Replace(v2State("2.1.0", "AUTOMATIC", "NONE",
		`[{"actionId":"`+cancel.ActionID+`","actionType":"cancelOrder","actionStatus":"FINISHED"}]`), `"serialNumber":"1"`, `"serialNumber":"2"`, 1)))
	if err != nil {
		t.Fatal(err)
	}
	var otherState State
	if err := json.Unmarshal(other.Payload, &otherState); err != nil {
		t.Fatal(err)
	}
	if len(otherState.InstantActionStates) != 0 || len(otherState.ActionStates) != 1 {
		t.Fatalf("instant actions leaked across vehicles: %+v", otherState)
	}
}

func TestV2StateEnumsAreCanonical(t *testing.T) {
	a := NewV2WireAdapter(V2WireOptions{})
	st := decodeV2State(t, a, v2State("2.0.0", "TEACHIN", "AUTOACK",
		`[{"actionId":"p","actionType":"initPosition","actionStatus":"FINISHED"}]`))
	if st.OperatingMode != OperatingModeTeachIn {
		t.Fatalf("operatingMode = %q", st.OperatingMode)
	}
	// AUTOACK is a self-clearing protective stop, not an emergency stop.
	if st.SafetyState.ActiveEmergencyStop != ActiveEmergencyStopNone || !st.SafetyState.FieldViolation {
		t.Fatalf("AUTOACK not mapped to a field violation: %+v", st.SafetyState)
	}
	if st.ActionStates[0].ActionType == nil || *st.ActionStates[0].ActionType != ActionInitializePosition {
		t.Fatalf("initPosition not mapped: %+v", st.ActionStates[0])
	}
	if st.InstantActionStates == nil {
		t.Fatal("instantActionStates must be present in canonical state")
	}
}

func TestV2ConnectionBrokenIsCanonical(t *testing.T) {
	decoded, err := NewV2WireAdapter(V2WireOptions{}).Decode("uagv/v2/KIT/1/connection",
		[]byte(`{"headerId":3,"timestamp":"2026-09-16T00:00:00Z","version":"2.1.0","manufacturer":"KIT","serialNumber":"1","connectionState":"CONNECTIONBROKEN"}`))
	if err != nil {
		t.Fatal(err)
	}
	var c Connection
	if err := json.Unmarshal(decoded.Payload, &c); err != nil {
		t.Fatal(err)
	}
	if c.ConnectionState != ConnectionStateConnectionBroken {
		t.Fatalf("connectionState = %q", c.ConnectionState)
	}
}

func TestV2FactsheetActionsAreCanonical(t *testing.T) {
	decoded, err := NewV2WireAdapter(V2WireOptions{}).Decode("uagv/v2/KIT/1/factsheet", []byte(`{"version":"2.1.0",
		"protocolLimits":{"maxStringLens":{"idLen":40,"loadIdLen":20},"maxArrayLens":{},"timing":{"minOrderInterval":1,"minStateInterval":1}},
		"protocolFeatures":{"optionalParameters":[],"agvActions":[
			{"actionType":"initPosition","actionScopes":["INSTANT"]},
			{"actionType":"lift","actionScopes":["NODE"],"actionParameters":[{"key":"height","valueDataType":"FLOAT"}]}]}}`))
	if err != nil {
		t.Fatal(err)
	}
	var fs Factsheet
	if err := json.Unmarshal(decoded.Payload, &fs); err != nil {
		t.Fatal(err)
	}
	strs := fs.ProtocolLimits.MaximumStringLengths
	if strs.MaximumIDLength == nil || *strs.MaximumIDLength != 40 || strs.MaximumLoadIDLength == nil || *strs.MaximumLoadIDLength != 20 {
		t.Fatalf("string limits lost: %+v", strs)
	}
	actions := fs.ProtocolFeatures.MobileRobotActions
	if actions[0].ActionType != ActionInitializePosition {
		t.Fatalf("initPosition not mapped: %q", actions[0].ActionType)
	}
	if actions[1].ActionParameters[0].ValueDataType != ValueDataTypeNumber {
		t.Fatalf("FLOAT not mapped: %q", actions[1].ActionParameters[0].ValueDataType)
	}
}
