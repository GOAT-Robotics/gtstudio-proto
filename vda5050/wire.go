package vda5050

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// WireAdapter is the MQTT wire-format boundary for a robot protocol revision
// or vendor profile. Fleet control works only with the canonical VDA 3 model;
// adapters translate messages at ingress and egress. A future protocol can
// implement this interface without leaking its wire model into backend or
// herdIQ.
type WireAdapter interface {
	Name() string
	SupportedVersions() []string
	Subscriptions() []WireSubscription
	VehicleSubscriptions(Identity) []WireSubscription
	Decode(wireTopic string, payload []byte) (DecodedWireMessage, error)
	Encode(id Identity, version string, topic Topic, canonical any) (WireMessage, error)
}

type WireSubscription struct {
	Filter string
	Topic  Topic
	QoS    byte
}

type WireMessage struct {
	Topic    string
	Payload  []byte
	QoS      byte
	Retained bool
}

type DecodedWireMessage struct {
	Identity Identity
	Topic    Topic
	Version  string
	Payload  []byte // canonical JSON
}

// V3WireAdapter implements the native VDA 5050 3.0.0 wire format.
type V3WireAdapter struct{ Scheme TopicScheme }

func NewV3WireAdapter(prefix string) *V3WireAdapter {
	return &V3WireAdapter{Scheme: NewTopicScheme(prefix)}
}

func (a *V3WireAdapter) Name() string                { return "vda5050-v3" }
func (a *V3WireAdapter) SupportedVersions() []string { return []string{ProtocolVersion} }

func (a *V3WireAdapter) Subscriptions() []WireSubscription {
	return wireSubscriptions(a.Scheme, Identity{}, true, RobotTopics, nil)
}

func (a *V3WireAdapter) VehicleSubscriptions(id Identity) []WireSubscription {
	return wireSubscriptions(a.Scheme, id, false, RobotTopics, nil)
}

func (a *V3WireAdapter) Decode(wireTopic string, payload []byte) (DecodedWireMessage, error) {
	id, topic, err := a.Scheme.Parse(wireTopic)
	if err != nil {
		return DecodedWireMessage{}, err
	}
	var h struct {
		Version string `json:"version"`
	}
	if err := json.Unmarshal(payload, &h); err != nil {
		return DecodedWireMessage{}, err
	}
	if h.Version != ProtocolVersion {
		return DecodedWireMessage{}, fmt.Errorf("vda5050: v3 adapter received version %q", h.Version)
	}
	return DecodedWireMessage{Identity: id, Topic: topic, Version: h.Version, Payload: payload}, nil
}

func (a *V3WireAdapter) Encode(id Identity, version string, topic Topic, canonical any) (WireMessage, error) {
	if version != "" && version != ProtocolVersion {
		return WireMessage{}, fmt.Errorf("vda5050: v3 adapter cannot encode version %q", version)
	}
	if err := Validate(topic, canonical); err != nil {
		return WireMessage{}, err
	}
	payload, err := json.Marshal(canonical)
	if err != nil {
		return WireMessage{}, err
	}
	return WireMessage{Topic: a.Scheme.Build(id, topic), Payload: payload, QoS: topic.QoS(), Retained: topic.Retained()}, nil
}

// V2WireOptions describes a standards-compliant VDA 2.x topic layout or a
// vendor dialect. The canonical model is never changed by these options.
type V2WireOptions struct {
	Name                   string
	Prefix                 string
	InterfaceName          string
	VersionLevel           string
	DefaultVersion         string
	TopicAliases           map[Topic]string
	HeaderVersionPrefix    string
	NumericConnectionState map[int]ConnectionState
}

// V2WireAdapter supports both 2.0.0 and the backward-compatible 2.1.0.
// The exact version is retained per vehicle and used for outbound messages.
type V2WireAdapter struct{ opts V2WireOptions }

func NewV2WireAdapter(opts V2WireOptions) *V2WireAdapter {
	if opts.Name == "" {
		opts.Name = "vda5050-v2"
	}
	if opts.InterfaceName == "" {
		opts.InterfaceName = "uagv"
	}
	if opts.VersionLevel == "" {
		opts.VersionLevel = "v2"
	}
	opts.Prefix = strings.Trim(opts.Prefix, "/")
	return &V2WireAdapter{opts: opts}
}

func (a *V2WireAdapter) Name() string { return a.opts.Name }
func (a *V2WireAdapter) SupportedVersions() []string {
	return []string{"2.0.0", "2.1.0"}
}

var v2RobotTopics = []Topic{TopicState, TopicVisualization, TopicConnection, TopicFactsheet}

func (a *V2WireAdapter) Subscriptions() []WireSubscription {
	return a.subscriptions(Identity{}, true)
}

func (a *V2WireAdapter) VehicleSubscriptions(id Identity) []WireSubscription {
	return a.subscriptions(id, false)
}

func (a *V2WireAdapter) subscriptions(id Identity, wildcard bool) []WireSubscription {
	out := make([]WireSubscription, 0, len(v2RobotTopics))
	for _, topic := range v2RobotTopics {
		wireName := a.wireTopicName(topic)
		parts := a.leadingLevels()
		if wildcard {
			parts = append(parts, "+", "+", wireName)
		} else {
			parts = append(parts, id.Manufacturer, id.SerialNumber, wireName)
		}
		out = append(out, WireSubscription{Filter: strings.Join(parts, "/"), Topic: topic, QoS: topic.QoS()})
	}
	return out
}

func (a *V2WireAdapter) leadingLevels() []string {
	levels := []string{}
	if a.opts.Prefix != "" {
		levels = append(levels, strings.Split(a.opts.Prefix, "/")...)
	}
	return append(levels, a.opts.InterfaceName, a.opts.VersionLevel)
}

func (a *V2WireAdapter) parseTopic(wireTopic string) (Identity, Topic, error) {
	parts := strings.Split(strings.Trim(wireTopic, "/"), "/")
	lead := a.leadingLevels()
	if len(parts) != len(lead)+3 {
		return Identity{}, "", errors.New("topic does not match adapter")
	}
	for i := range lead {
		if parts[i] != lead[i] {
			return Identity{}, "", errors.New("topic does not match adapter")
		}
	}
	wireName := parts[len(parts)-1]
	topic := Topic(wireName)
	for canonical, alias := range a.opts.TopicAliases {
		if wireName == alias {
			topic = canonical
			break
		}
	}
	known := false
	for _, candidate := range v2RobotTopics {
		known = known || candidate == topic
	}
	if !known {
		return Identity{}, "", fmt.Errorf("unsupported VDA 2 topic %q", wireName)
	}
	id := Identity{Manufacturer: parts[len(lead)], SerialNumber: parts[len(lead)+1]}
	return id, topic, id.Valid()
}

func (a *V2WireAdapter) Decode(wireTopic string, payload []byte) (DecodedWireMessage, error) {
	id, topic, err := a.parseTopic(wireTopic)
	if err != nil {
		return DecodedWireMessage{}, err
	}
	var body map[string]any
	if err := json.Unmarshal(payload, &body); err != nil {
		return DecodedWireMessage{}, fmt.Errorf("vda5050: decoding %s: %w", topic, err)
	}
	version := normalizeV2Version(stringValue(body["version"]))
	if version == "" {
		// A configured dialect may omit the header version. Its explicit
		// default supplies the version; standard VDA messages remain strict.
		version = normalizeV2Version(a.opts.DefaultVersion)
	}
	if version != "2.0.0" && version != "2.1.0" {
		return DecodedWireMessage{}, fmt.Errorf("vda5050: unsupported VDA 2 version %q", body["version"])
	}
	body["version"] = version
	body["manufacturer"] = id.Manufacturer
	body["serialNumber"] = id.SerialNumber
	if err := a.toCanonical(topic, body); err != nil {
		return DecodedWireMessage{}, err
	}
	canonical, err := json.Marshal(body)
	if err != nil {
		return DecodedWireMessage{}, err
	}
	return DecodedWireMessage{Identity: id, Topic: topic, Version: version, Payload: canonical}, nil
}

func (a *V2WireAdapter) toCanonical(topic Topic, body map[string]any) error {
	switch topic {
	case TopicState:
		rename(body, "agvPosition", "mobileRobotPosition")
		rename(body, "batteryState", "powerSupply")
		if p := object(body["mobileRobotPosition"]); p != nil {
			rename(p, "positionInitialized", "localized")
		}
		if p := object(body["powerSupply"]); p != nil {
			rename(p, "batteryCharge", "stateOfCharge")
			rename(p, "reach", "range")
		}
		if s := object(body["safetyState"]); s != nil {
			rename(s, "eStop", "activeEmergencyStop")
		}
		for _, item := range array(body["actionStates"]) {
			if action := object(item); action != nil {
				rename(action, "actionDescription", "actionDescriptor")
				rename(action, "resultDescription", "actionResult")
			}
		}
		for _, item := range array(body["nodeStates"]) {
			if state := object(item); state != nil {
				rename(state, "nodeDescription", "nodeDescriptor")
			}
		}
		for _, item := range array(body["edgeStates"]) {
			if state := object(item); state != nil {
				rename(state, "edgeDescription", "edgeDescriptor")
			}
		}
		for _, item := range array(body["information"]) {
			if info := object(item); info != nil {
				rename(info, "infoDescription", "infoDescriptor")
			}
		}
	case TopicVisualization:
		rename(body, "agvPosition", "mobileRobotPosition")
		if p := object(body["mobileRobotPosition"]); p != nil {
			rename(p, "positionInitialized", "localized")
		}
	case TopicConnection:
		if n, ok := numberAsInt(body["connectionState"]); ok {
			state, exists := a.opts.NumericConnectionState[n]
			if !exists {
				return fmt.Errorf("vda5050: %s unknown numeric connectionState %d", a.Name(), n)
			}
			body["connectionState"] = string(state)
		}
		if _, ok := body["headerId"]; !ok {
			body["headerId"] = float64(0)
		}
	case TopicFactsheet:
		rename(body, "agvGeometry", "mobileRobotGeometry")
		rename(body, "vehicleConfig", "mobileRobotConfiguration")
		if ts := object(body["typeSpecification"]); ts != nil {
			rename(ts, "agvKinematic", "mobileRobotKinematics")
			rename(ts, "agvClass", "mobileRobotClass")
			rename(ts, "maxLoadMass", "maximumLoadMass")
		}
		if physical := object(body["physicalParameters"]); physical != nil {
			rename(physical, "speedMin", "minimumSpeed")
			rename(physical, "speedMax", "maximumSpeed")
			rename(physical, "accelerationMax", "maximumAcceleration")
			rename(physical, "decelerationMax", "maximumDeceleration")
			rename(physical, "heightMin", "minimumHeight")
			rename(physical, "heightMax", "maximumHeight")
		}
		if limits := object(body["protocolLimits"]); limits != nil {
			rename(limits, "maxStringLens", "maximumStringLengths")
			rename(limits, "maxArrayLens", "maximumArrayLengths")
			if strings := object(limits["maximumStringLengths"]); strings != nil {
				rename(strings, "msgLen", "maximumMessageLength")
				rename(strings, "topicSerialLen", "maximumTopicSerialLength")
				rename(strings, "topicElemLen", "maximumTopicElementLength")
				rename(strings, "idLen", "maximumIDLength")
				rename(strings, "loadIdLen", "maximumLoadIDLength")
			}
			if timing := object(limits["timing"]); timing != nil {
				rename(timing, "minOrderInterval", "minimumOrderInterval")
				rename(timing, "minStateInterval", "minimumStateInterval")
			}
		}
		if pf := object(body["protocolFeatures"]); pf != nil {
			rename(pf, "agvActions", "mobileRobotActions")
			for _, item := range array(pf["mobileRobotActions"]) {
				if action := object(item); action != nil {
					rename(action, "resultDescription", "actionResult")
					if _, ok := action["pauseAllowed"]; !ok {
						action["pauseAllowed"] = false
					}
					if _, ok := action["cancelAllowed"]; !ok {
						action["cancelAllowed"] = false
					}
				}
			}
		}
	}
	return nil
}

func (a *V2WireAdapter) Encode(id Identity, version string, topic Topic, canonical any) (WireMessage, error) {
	version = normalizeV2Version(version)
	if version != "2.0.0" && version != "2.1.0" {
		return WireMessage{}, fmt.Errorf("vda5050: VDA 2 adapter cannot encode version %q", version)
	}
	if topic == TopicZoneSet || topic == TopicResponses {
		return WireMessage{}, fmt.Errorf("vda5050: topic %s is unavailable before VDA 3.0.0", topic)
	}
	raw, err := json.Marshal(canonical)
	if err != nil {
		return WireMessage{}, err
	}
	var body map[string]any
	if err := json.Unmarshal(raw, &body); err != nil {
		return WireMessage{}, err
	}
	body["version"] = a.opts.HeaderVersionPrefix + version
	body["manufacturer"], body["serialNumber"] = id.Manufacturer, id.SerialNumber
	if err := fromCanonicalV2(version, topic, body); err != nil {
		return WireMessage{}, err
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return WireMessage{}, err
	}
	wireName := a.wireTopicName(topic)
	parts := append(a.leadingLevels(), id.Manufacturer, id.SerialNumber, wireName)
	return WireMessage{Topic: strings.Join(parts, "/"), Payload: payload, QoS: topic.QoS(), Retained: topic.Retained()}, nil
}

func (a *V2WireAdapter) wireTopicName(topic Topic) string {
	if alias := a.opts.TopicAliases[topic]; alias != "" {
		return alias
	}
	return string(topic)
}

func fromCanonicalV2(version string, topic Topic, body map[string]any) error {
	switch topic {
	case TopicOrder:
		nodes := array(body["nodes"])
		for _, item := range nodes {
			if node := object(item); node != nil {
				rename(node, "nodeDescriptor", "nodeDescription")
				v2Actions(array(node["actions"]), version, false)
			}
		}
		for i, item := range array(body["edges"]) {
			edge := object(item)
			if edge == nil {
				continue
			}
			rename(edge, "edgeDescriptor", "edgeDescription")
			rename(edge, "maximumSpeed", "maxSpeed")
			rename(edge, "maximumMobileRobotHeight", "maxHeight")
			rename(edge, "minimumLoadHandlingDeviceHeight", "minHeight")
			rename(edge, "maximumRotationSpeed", "maxRotationSpeed")
			delete(edge, "corridor")
			delete(edge, "orientationType")
			delete(edge, "reachOrientationBeforeEntering")
			if i < len(nodes)-1 {
				edge["startNodeId"] = object(nodes[i])["nodeId"]
				edge["endNodeId"] = object(nodes[i+1])["nodeId"]
			}
			v2Actions(array(edge["actions"]), version, false)
		}
	case TopicInstantActions:
		v2Actions(array(body["actions"]), version, true)
	}
	return nil
}

func v2Actions(items []any, version string, instant bool) {
	for _, item := range items {
		action := object(item)
		if action == nil {
			continue
		}
		rename(action, "actionDescriptor", "actionDescription")
		if action["actionType"] == ActionInitializePosition {
			action["actionType"] = "initPosition"
		}
		if instant && version == "2.0.0" {
			rename(action, "actionType", "actionName")
		}
		delete(action, "retriable")
	}
}

func wireSubscriptions(s TopicScheme, id Identity, wildcard bool, topics []Topic, aliases map[Topic]string) []WireSubscription {
	out := make([]WireSubscription, 0, len(topics))
	for _, topic := range topics {
		filter := s.Build(id, topic)
		if wildcard {
			filter = s.SubscribeAll(topic)
		}
		if alias := aliases[topic]; alias != "" {
			filter = strings.TrimSuffix(filter, string(topic)) + alias
		}
		out = append(out, WireSubscription{Filter: filter, Topic: topic, QoS: topic.QoS()})
	}
	return out
}

func normalizeV2Version(v string) string {
	v = strings.TrimSpace(strings.TrimPrefix(v, "v"))
	if v == "2" || v == "2.0" {
		return "2.0.0"
	}
	if v == "2.1" {
		return "2.1.0"
	}
	return v
}

func rename(m map[string]any, old, next string) {
	if v, ok := m[old]; ok {
		if _, exists := m[next]; !exists {
			m[next] = v
		}
		delete(m, old)
	}
}
func object(v any) map[string]any { m, _ := v.(map[string]any); return m }
func array(v any) []any           { a, _ := v.([]any); return a }
func stringValue(v any) string    { s, _ := v.(string); return s }
func numberAsInt(v any) (int, bool) {
	switch n := v.(type) {
	case float64:
		return int(n), n == float64(int(n))
	case int:
		return n, true
	case json.Number:
		i, err := strconv.Atoi(n.String())
		return i, err == nil
	default:
		return 0, false
	}
}

// DefaultWireAdapters enables only standards-compliant wire formats.
// Deployment-specific dialects are supplied as additional V2WireAdapters.
func DefaultWireAdapters(prefix string) []WireAdapter {
	return []WireAdapter{
		NewV3WireAdapter(prefix),
		NewV2WireAdapter(V2WireOptions{Prefix: prefix}),
	}
}
