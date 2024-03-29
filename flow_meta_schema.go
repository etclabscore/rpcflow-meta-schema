package flow_meta_schema


import "encoding/json"
import "errors"
// **REQUIRED**. This string MUST be the semantic version number of the Specification that the document uses.
type FlowVersion string
const (
	FlowVersionEnum0 FlowVersion = "0.0.0-development"
)
// **REQUIRED**. Name of the Flow.
//
// --- Example ---
//
// `foobarbaz`
type FlowName string
type Namespace string
type Transport string
const (
	TransportEnum0 Transport = "postmessageiframe"
	TransportEnum1 Transport = "postmessagewindow"
	TransportEnum2 Transport = "http"
	TransportEnum3 Transport = "websocket"
)
type Url string
type StepProvider struct {
	Namespace *Namespace `json:"namespace"`
	Transport *Transport `json:"transport"`
	Url       *Url       `json:"url"`
}
type StepProviders []StepProvider
// Markdown description describing the specification extension.
type FlowDescription string
// A short summary of what the flow is.
type FlowSummary string
type Id string
type Schema string
type Ref string
type Comment string
type Title string
type Description string
type AlwaysTrue interface{}
type ReadOnly bool
type Examples []AlwaysTrue
type MultipleOf float64
type Maximum float64
type ExclusiveMaximum float64
type Minimum float64
type ExclusiveMinimum float64
type NonNegativeInteger int64
type NonNegativeIntegerDefaultZero int64
type Pattern string
type SchemaArray []JSONSchema
//
// --- Default ---
//
// true
type Items struct {
	JSONSchema  *JSONSchema
	SchemaArray *SchemaArray
}
func (a *Items) UnmarshalJSON(bytes []byte) error {
	var ok bool
	var myJSONSchema JSONSchema
	if err := json.Unmarshal(bytes, &myJSONSchema); err == nil {
		ok = true
		a.JSONSchema = &myJSONSchema
	}
	var mySchemaArray SchemaArray
	if err := json.Unmarshal(bytes, &mySchemaArray); err == nil {
		ok = true
		a.SchemaArray = &mySchemaArray
	}
	if ok {
		return nil
	}
	return errors.New("failed to unmarshal any of the object properties")
}
func (o Items) MarshalJSON() ([]byte, error) {
	out := []interface{}{}
	if o.JSONSchema != nil {
		out = append(out, o.JSONSchema)
	}
	if o.SchemaArray != nil {
		out = append(out, o.SchemaArray)
	}
	return json.Marshal(out)
}
type UniqueItems bool
type StringDoaGddGA string
//
// --- Default ---
//
// []
type StringArray []StringDoaGddGA
//
// --- Default ---
//
// {}
type Definitions map[string]interface{}
//
// --- Default ---
//
// {}
type Properties map[string]interface{}
//
// --- Default ---
//
// {}
type PatternProperties map[string]interface{}
type DependenciesSet struct {
	JSONSchema  *JSONSchema
	StringArray *StringArray
}
func (a *DependenciesSet) UnmarshalJSON(bytes []byte) error {
	var ok bool
	var myJSONSchema JSONSchema
	if err := json.Unmarshal(bytes, &myJSONSchema); err == nil {
		ok = true
		a.JSONSchema = &myJSONSchema
	}
	var myStringArray StringArray
	if err := json.Unmarshal(bytes, &myStringArray); err == nil {
		ok = true
		a.StringArray = &myStringArray
	}
	if ok {
		return nil
	}
	return errors.New("failed to unmarshal any of the object properties")
}
func (o DependenciesSet) MarshalJSON() ([]byte, error) {
	out := []interface{}{}
	if o.JSONSchema != nil {
		out = append(out, o.JSONSchema)
	}
	if o.StringArray != nil {
		out = append(out, o.StringArray)
	}
	return json.Marshal(out)
}
type Dependencies map[string]interface{}
type Enum []AlwaysTrue
type SimpleTypes interface{}
type ArrayOfSimpleTypes []SimpleTypes
type Type struct {
	SimpleTypes        *SimpleTypes
	ArrayOfSimpleTypes *ArrayOfSimpleTypes
}
func (a *Type) UnmarshalJSON(bytes []byte) error {
	var ok bool
	var mySimpleTypes SimpleTypes
	if err := json.Unmarshal(bytes, &mySimpleTypes); err == nil {
		ok = true
		a.SimpleTypes = &mySimpleTypes
	}
	var myArrayOfSimpleTypes ArrayOfSimpleTypes
	if err := json.Unmarshal(bytes, &myArrayOfSimpleTypes); err == nil {
		ok = true
		a.ArrayOfSimpleTypes = &myArrayOfSimpleTypes
	}
	if ok {
		return nil
	}
	return errors.New("failed to unmarshal any of the object properties")
}
func (o Type) MarshalJSON() ([]byte, error) {
	out := []interface{}{}
	if o.SimpleTypes != nil {
		out = append(out, o.SimpleTypes)
	}
	if o.ArrayOfSimpleTypes != nil {
		out = append(out, o.ArrayOfSimpleTypes)
	}
	return json.Marshal(out)
}
type Format string
type ContentMediaType string
type ContentEncoding string
type JSONSchemaObject struct {
	Id                   *Id                            `json:"$id,omitempty"`
	Schema               *Schema                        `json:"$schema,omitempty"`
	Ref                  *Ref                           `json:"$ref,omitempty"`
	Comment              *Comment                       `json:"$comment,omitempty"`
	Title                *Title                         `json:"title,omitempty"`
	Description          *Description                   `json:"description,omitempty"`
	Default              *AlwaysTrue                    `json:"default,omitempty"`
	ReadOnly             *ReadOnly                      `json:"readOnly,omitempty"`
	Examples             *Examples                      `json:"examples,omitempty"`
	MultipleOf           *MultipleOf                    `json:"multipleOf,omitempty"`
	Maximum              *Maximum                       `json:"maximum,omitempty"`
	ExclusiveMaximum     *ExclusiveMaximum              `json:"exclusiveMaximum,omitempty"`
	Minimum              *Minimum                       `json:"minimum,omitempty"`
	ExclusiveMinimum     *ExclusiveMinimum              `json:"exclusiveMinimum,omitempty"`
	MaxLength            *NonNegativeInteger            `json:"maxLength,omitempty"`
	MinLength            *NonNegativeIntegerDefaultZero `json:"minLength,omitempty"`
	Pattern              *Pattern                       `json:"pattern,omitempty"`
	AdditionalItems      *JSONSchema                    `json:"additionalItems,omitempty"`
	Items                *Items                         `json:"items,omitempty"`
	MaxItems             *NonNegativeInteger            `json:"maxItems,omitempty"`
	MinItems             *NonNegativeIntegerDefaultZero `json:"minItems,omitempty"`
	UniqueItems          *UniqueItems                   `json:"uniqueItems,omitempty"`
	Contains             *JSONSchema                    `json:"contains,omitempty"`
	MaxProperties        *NonNegativeInteger            `json:"maxProperties,omitempty"`
	MinProperties        *NonNegativeIntegerDefaultZero `json:"minProperties,omitempty"`
	Required             *StringArray                   `json:"required,omitempty"`
	AdditionalProperties *JSONSchema                    `json:"additionalProperties,omitempty"`
	Definitions          *Definitions                   `json:"definitions,omitempty"`
	Properties           *Properties                    `json:"properties,omitempty"`
	PatternProperties    *PatternProperties             `json:"patternProperties,omitempty"`
	Dependencies         *Dependencies                  `json:"dependencies,omitempty"`
	PropertyNames        *JSONSchema                    `json:"propertyNames,omitempty"`
	Const                *AlwaysTrue                    `json:"const,omitempty"`
	Enum                 *Enum                          `json:"enum,omitempty"`
	Type                 *Type                          `json:"type,omitempty"`
	Format               *Format                        `json:"format,omitempty"`
	ContentMediaType     *ContentMediaType              `json:"contentMediaType,omitempty"`
	ContentEncoding      *ContentEncoding               `json:"contentEncoding,omitempty"`
	If                   *JSONSchema                    `json:"if,omitempty"`
	Then                 *JSONSchema                    `json:"then,omitempty"`
	Else                 *JSONSchema                    `json:"else,omitempty"`
	AllOf                *SchemaArray                   `json:"allOf,omitempty"`
	AnyOf                *SchemaArray                   `json:"anyOf,omitempty"`
	OneOf                *SchemaArray                   `json:"oneOf,omitempty"`
	Not                  *JSONSchema                    `json:"not,omitempty"`
}
// Always valid if true. Never valid if false. Is constant.
type JSONSchemaBoolean bool
//
// --- Default ---
//
// {}
type JSONSchema struct {
	JSONSchemaObject  *JSONSchemaObject
	JSONSchemaBoolean *JSONSchemaBoolean
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *JSONSchema) UnmarshalJSON(bytes []byte) error {
	var myJSONSchemaObject JSONSchemaObject
	if err := json.Unmarshal(bytes, &myJSONSchemaObject); err == nil {
		o.JSONSchemaObject = &myJSONSchemaObject
		return nil
	}
	var myJSONSchemaBoolean JSONSchemaBoolean
	if err := json.Unmarshal(bytes, &myJSONSchemaBoolean); err == nil {
		o.JSONSchemaBoolean = &myJSONSchemaBoolean
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o JSONSchema) MarshalJSON() ([]byte, error) {
	if o.JSONSchemaObject != nil {
		return json.Marshal(o.JSONSchemaObject)
	}
	if o.JSONSchemaBoolean != nil {
		return json.Marshal(o.JSONSchemaBoolean)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
// Additional context provided to the flow. If `contextSchema` exists, `context` **MUST** validate against it. Context can be used via the `\"${context.<path>}\"` [Runtime Expression](https://github.com/etclabscore/json-template-language#what-is-it-used-for)
type FlowContext map[string]interface{}
// **REQUIRED**. Canonical name for the method call.
type FlowMethod string
type ObjectHAgrRKSz map[string]interface{}
type SetOfAnykhpArGaq []AlwaysTrue
// A list of parameters to be called for this method.
type FlowParams struct {
	ObjectHAgrRKSz   *ObjectHAgrRKSz
	SetOfAnykhpArGaq *SetOfAnykhpArGaq
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *FlowParams) UnmarshalJSON(bytes []byte) error {
	var myObjectHAgrRKSz ObjectHAgrRKSz
	if err := json.Unmarshal(bytes, &myObjectHAgrRKSz); err == nil {
		o.ObjectHAgrRKSz = &myObjectHAgrRKSz
		return nil
	}
	var mySetOfAnykhpArGaq SetOfAnykhpArGaq
	if err := json.Unmarshal(bytes, &mySetOfAnykhpArGaq); err == nil {
		o.SetOfAnykhpArGaq = &mySetOfAnykhpArGaq
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o FlowParams) MarshalJSON() ([]byte, error) {
	if o.ObjectHAgrRKSz != nil {
		return json.Marshal(o.ObjectHAgrRKSz)
	}
	if o.SetOfAnykhpArGaq != nil {
		return json.Marshal(o.SetOfAnykhpArGaq)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
type FlowObject struct {
	Method *FlowMethod `json:"method"`
	Params *FlowParams `json:"params"`
}
// Describes how to call JSON-RPC methods at runtime and use their results via [Runtime Expressions](https://github.com/etclabscore/json-template-language#what-is-it-used-for). Flow results can be used via \"${results.<keyName>}\" where `keyName` refers to the key of the flow object to use the results from.
type Flow map[string]interface{}
// A verbose explanation of the target documentation. GitHub Flavored Markdown syntax MAY be used for rich text representation.
type FlowExternalDocumentationObjectDescription string
// **REQUIRED**. The URL for the target documentation. Value MUST be in the format of a URL.
type FlowExternalDocumentationObjectUrl string
// Allows referencing an external resource for extended documentation.
type FlowExternalDocumentationObject struct {
	Description *FlowExternalDocumentationObjectDescription `json:"description,omitempty"`
	Url         *FlowExternalDocumentationObjectUrl         `json:"url"`
}
// A standard way to define JSON-RPC flows.
type FlowMetaSchema struct {
	Rpcflow               *FlowVersion                     `json:"rpcflow"`
	Name                  *FlowName                        `json:"name"`
	StepProviders         *StepProviders                   `json:"stepProviders"`
	Description           *FlowDescription                 `json:"description,omitempty"`
	Summary               *FlowSummary                     `json:"summary,omitempty"`
	ContextSchema         *JSONSchema                      `json:"contextSchema,omitempty"`
	Context               *FlowContext                     `json:"context,omitempty"`
	Flow                  *Flow                            `json:"flow"`
	ExternalDocumentation *FlowExternalDocumentationObject `json:"externalDocumentation,omitempty"`
}

const RawFlow_meta_schema = "{\"$schema\":\"https://meta.json-schema.tools/\",\"title\":\"Flow Meta Schema\",\"description\":\"A standard way to define JSON-RPC flows.\",\"type\":\"object\",\"required\":[\"name\",\"stepProviders\",\"flow\",\"rpcflow\"],\"properties\":{\"rpcflow\":{\"$ref\":\"#/definitions/FlowVersion\"},\"name\":{\"$ref\":\"#/definitions/FlowName\"},\"stepProviders\":{\"$ref\":\"#/definitions/StepProviders\"},\"description\":{\"$ref\":\"#/definitions/FlowDescription\"},\"summary\":{\"$ref\":\"#/definitions/FlowSummary\"},\"contextSchema\":{\"$ref\":\"#/definitions/JSONSchema\"},\"context\":{\"$ref\":\"#/definitions/FlowContext\"},\"flow\":{\"$ref\":\"#/definitions/Flow\"},\"externalDocumentation\":{\"$ref\":\"#/definitions/FlowExternalDocumentationObject\"}},\"definitions\":{\"FlowVersion\":{\"title\":\"FlowVersion\",\"description\":\"**REQUIRED**. This string MUST be the semantic version number of the Specification that the document uses.\",\"type\":\"string\",\"enum\":[\"0.0.0-development\"]},\"FlowName\":{\"title\":\"FlowName\",\"description\":\"**REQUIRED**. Name of the Flow.\",\"type\":\"string\",\"examples\":[\"foobarbaz\"]},\"Namespace\":{\"title\":\"Namespace\",\"type\":\"string\"},\"Transport\":{\"title\":\"Transport\",\"type\":\"string\",\"enum\":[\"postmessageiframe\",\"postmessagewindow\",\"http\",\"websocket\"]},\"Url\":{\"title\":\"Url\",\"type\":\"string\",\"format\":\"uri\"},\"StepProvider\":{\"title\":\"StepProvider\",\"type\":\"object\",\"required\":[\"namespace\",\"transport\",\"url\"],\"properties\":{\"namespace\":{\"$ref\":\"#/definitions/Namespace\"},\"transport\":{\"$ref\":\"#/definitions/Transport\"},\"url\":{\"$ref\":\"#/definitions/Url\"}}},\"StepProviders\":{\"title\":\"StepProviders\",\"type\":\"array\",\"items\":{\"$ref\":\"#/definitions/StepProvider\"}},\"FlowDescription\":{\"title\":\"FlowDescription\",\"type\":\"string\",\"description\":\"Markdown description describing the specification extension.\"},\"FlowSummary\":{\"title\":\"FlowSummary\",\"type\":\"string\",\"description\":\"A short summary of what the flow is.\"},\"$id\":{\"title\":\"$id\",\"type\":\"string\",\"format\":\"uri-reference\"},\"$schema\":{\"title\":\"$schema\",\"type\":\"string\",\"format\":\"uri\"},\"$ref\":{\"title\":\"$ref\",\"type\":\"string\",\"format\":\"uri-reference\"},\"$comment\":{\"title\":\"$comment\",\"type\":\"string\"},\"title\":{\"title\":\"title\",\"type\":\"string\"},\"description\":{\"title\":\"description\",\"type\":\"string\"},\"AlwaysTrue\":true,\"readOnly\":{\"title\":\"readOnly\",\"type\":\"boolean\",\"default\":false},\"examples\":{\"title\":\"examples\",\"type\":\"array\",\"items\":true},\"multipleOf\":{\"title\":\"multipleOf\",\"type\":\"number\",\"exclusiveMinimum\":0},\"maximum\":{\"title\":\"maximum\",\"type\":\"number\"},\"exclusiveMaximum\":{\"title\":\"exclusiveMaximum\",\"type\":\"number\"},\"minimum\":{\"title\":\"minimum\",\"type\":\"number\"},\"exclusiveMinimum\":{\"title\":\"exclusiveMinimum\",\"type\":\"number\"},\"nonNegativeInteger\":{\"title\":\"nonNegativeInteger\",\"type\":\"integer\",\"minimum\":0},\"nonNegativeIntegerDefaultZero\":{\"title\":\"nonNegativeIntegerDefaultZero\",\"type\":\"integer\",\"minimum\":0,\"default\":0},\"pattern\":{\"title\":\"pattern\",\"type\":\"string\",\"format\":\"regex\"},\"schemaArray\":{\"title\":\"schemaArray\",\"type\":\"array\",\"minItems\":1,\"items\":{\"$ref\":\"#/definitions/JSONSchema\"}},\"items\":{\"title\":\"items\",\"anyOf\":[{\"$ref\":\"#/definitions/JSONSchema\"},{\"$ref\":\"#/definitions/schemaArray\"}],\"default\":true},\"uniqueItems\":{\"title\":\"uniqueItems\",\"type\":\"boolean\",\"default\":false},\"string_doaGddGA\":{\"type\":\"string\",\"title\":\"string_doaGddGA\"},\"stringArray\":{\"title\":\"stringArray\",\"type\":\"array\",\"items\":{\"$ref\":\"#/definitions/string_doaGddGA\"},\"uniqueItems\":true,\"default\":[]},\"definitions\":{\"title\":\"definitions\",\"type\":\"object\",\"additionalProperties\":{\"$ref\":\"#/definitions/JSONSchema\"},\"default\":{}},\"properties\":{\"title\":\"properties\",\"type\":\"object\",\"additionalProperties\":{\"$ref\":\"#/definitions/JSONSchema\"},\"default\":{}},\"patternProperties\":{\"title\":\"patternProperties\",\"type\":\"object\",\"additionalProperties\":{\"$ref\":\"#/definitions/JSONSchema\"},\"propertyNames\":{\"title\":\"propertyNames\",\"format\":\"regex\"},\"default\":{}},\"dependenciesSet\":{\"title\":\"dependenciesSet\",\"anyOf\":[{\"$ref\":\"#/definitions/JSONSchema\"},{\"$ref\":\"#/definitions/stringArray\"}]},\"dependencies\":{\"title\":\"dependencies\",\"type\":\"object\",\"additionalProperties\":{\"$ref\":\"#/definitions/dependenciesSet\"}},\"enum\":{\"title\":\"enum\",\"type\":\"array\",\"items\":true,\"minItems\":1,\"uniqueItems\":true},\"simpleTypes\":{\"title\":\"simpleTypes\",\"enum\":[\"array\",\"boolean\",\"integer\",\"null\",\"number\",\"object\",\"string\"]},\"arrayOfSimpleTypes\":{\"title\":\"arrayOfSimpleTypes\",\"type\":\"array\",\"items\":{\"$ref\":\"#/definitions/simpleTypes\"},\"minItems\":1,\"uniqueItems\":true},\"type\":{\"title\":\"type\",\"anyOf\":[{\"$ref\":\"#/definitions/simpleTypes\"},{\"$ref\":\"#/definitions/arrayOfSimpleTypes\"}]},\"format\":{\"title\":\"format\",\"type\":\"string\"},\"contentMediaType\":{\"title\":\"contentMediaType\",\"type\":\"string\"},\"contentEncoding\":{\"title\":\"contentEncoding\",\"type\":\"string\"},\"JSONSchemaObject\":{\"title\":\"JSONSchemaObject\",\"type\":\"object\",\"properties\":{\"$id\":{\"$ref\":\"#/definitions/$id\"},\"$schema\":{\"$ref\":\"#/definitions/$schema\"},\"$ref\":{\"$ref\":\"#/definitions/$ref\"},\"$comment\":{\"$ref\":\"#/definitions/$comment\"},\"title\":{\"$ref\":\"#/definitions/title\"},\"description\":{\"$ref\":\"#/definitions/description\"},\"default\":true,\"readOnly\":{\"$ref\":\"#/definitions/readOnly\"},\"examples\":{\"$ref\":\"#/definitions/examples\"},\"multipleOf\":{\"$ref\":\"#/definitions/multipleOf\"},\"maximum\":{\"$ref\":\"#/definitions/maximum\"},\"exclusiveMaximum\":{\"$ref\":\"#/definitions/exclusiveMaximum\"},\"minimum\":{\"$ref\":\"#/definitions/minimum\"},\"exclusiveMinimum\":{\"$ref\":\"#/definitions/exclusiveMinimum\"},\"maxLength\":{\"$ref\":\"#/definitions/nonNegativeInteger\"},\"minLength\":{\"$ref\":\"#/definitions/nonNegativeIntegerDefaultZero\"},\"pattern\":{\"$ref\":\"#/definitions/pattern\"},\"additionalItems\":{\"$ref\":\"#/definitions/JSONSchema\"},\"items\":{\"$ref\":\"#/definitions/items\"},\"maxItems\":{\"$ref\":\"#/definitions/nonNegativeInteger\"},\"minItems\":{\"$ref\":\"#/definitions/nonNegativeIntegerDefaultZero\"},\"uniqueItems\":{\"$ref\":\"#/definitions/uniqueItems\"},\"contains\":{\"$ref\":\"#/definitions/JSONSchema\"},\"maxProperties\":{\"$ref\":\"#/definitions/nonNegativeInteger\"},\"minProperties\":{\"$ref\":\"#/definitions/nonNegativeIntegerDefaultZero\"},\"required\":{\"$ref\":\"#/definitions/stringArray\"},\"additionalProperties\":{\"$ref\":\"#/definitions/JSONSchema\"},\"definitions\":{\"$ref\":\"#/definitions/definitions\"},\"properties\":{\"$ref\":\"#/definitions/properties\"},\"patternProperties\":{\"$ref\":\"#/definitions/patternProperties\"},\"dependencies\":{\"$ref\":\"#/definitions/dependencies\"},\"propertyNames\":{\"$ref\":\"#/definitions/JSONSchema\"},\"const\":true,\"enum\":{\"$ref\":\"#/definitions/enum\"},\"type\":{\"$ref\":\"#/definitions/type\"},\"format\":{\"$ref\":\"#/definitions/format\"},\"contentMediaType\":{\"$ref\":\"#/definitions/contentMediaType\"},\"contentEncoding\":{\"$ref\":\"#/definitions/contentEncoding\"},\"if\":{\"$ref\":\"#/definitions/JSONSchema\"},\"then\":{\"$ref\":\"#/definitions/JSONSchema\"},\"else\":{\"$ref\":\"#/definitions/JSONSchema\"},\"allOf\":{\"$ref\":\"#/definitions/schemaArray\"},\"anyOf\":{\"$ref\":\"#/definitions/schemaArray\"},\"oneOf\":{\"$ref\":\"#/definitions/schemaArray\"},\"not\":{\"$ref\":\"#/definitions/JSONSchema\"}}},\"JSONSchemaBoolean\":{\"title\":\"JSONSchemaBoolean\",\"description\":\"Always valid if true. Never valid if false. Is constant.\",\"type\":\"boolean\"},\"JSONSchema\":{\"$schema\":\"https://meta.json-schema.tools/\",\"$id\":\"https://meta.json-schema.tools/\",\"title\":\"JSONSchema\",\"default\":{},\"oneOf\":[{\"$ref\":\"#/definitions/JSONSchemaObject\"},{\"$ref\":\"#/definitions/JSONSchemaBoolean\"}]},\"FlowContext\":{\"title\":\"FlowContext\",\"description\":\"Additional context provided to the flow. If `contextSchema` exists, `context` **MUST** validate against it. Context can be used via the `\\\\"${context.<path>}\\\\"` [Runtime Expression](https://github.com/etclabscore/json-template-language#what-is-it-used-for)\",\"type\":\"object\"},\"FlowMethod\":{\"title\":\"FlowMethod\",\"type\":\"string\",\"description\":\"**REQUIRED**. Canonical name for the method call.\"},\"object_HAgrRKSz\":{\"type\":\"object\",\"title\":\"object_HAgrRKSz\"},\"setOfAnykhpArGaq\":{\"type\":\"array\",\"items\":true,\"title\":\"setOfAnykhpArGaq\"},\"FlowParams\":{\"title\":\"FlowParams\",\"description\":\"A list of parameters to be called for this method.\",\"oneOf\":[{\"$ref\":\"#/definitions/object_HAgrRKSz\"},{\"$ref\":\"#/definitions/setOfAnykhpArGaq\"}]},\"FlowObject\":{\"title\":\"FlowObject\",\"type\":\"object\",\"required\":[\"method\",\"params\"],\"properties\":{\"method\":{\"$ref\":\"#/definitions/FlowMethod\"},\"params\":{\"$ref\":\"#/definitions/FlowParams\"}}},\"Flow\":{\"title\":\"Flow\",\"description\":\"Describes how to call JSON-RPC methods at runtime and use their results via [Runtime Expressions](https://github.com/etclabscore/json-template-language#what-is-it-used-for). Flow results can be used via \\\\"${results.<keyName>}\\\\" where `keyName` refers to the key of the flow object to use the results from.\",\"type\":\"object\",\"patternProperties\":{\"\":{\"$ref\":\"#/definitions/FlowObject\"}}},\"FlowExternalDocumentationObjectDescription\":{\"title\":\"FlowExternalDocumentationObjectDescription\",\"description\":\"A verbose explanation of the target documentation. GitHub Flavored Markdown syntax MAY be used for rich text representation.\",\"summary\":\"External documentation description.\",\"type\":\"string\"},\"FlowExternalDocumentationObjectUrl\":{\"title\":\"FlowExternalDocumentationObjectUrl\",\"description\":\"**REQUIRED**. The URL for the target documentation. Value MUST be in the format of a URL.\",\"summary\":\"External Documentation URI.\",\"type\":\"string\",\"format\":\"uri\"},\"FlowExternalDocumentationObject\":{\"title\":\"FlowExternalDocumentationObject\",\"type\":\"object\",\"additionalProperties\":false,\"summary\":\"Information about external documentation.\",\"description\":\"Allows referencing an external resource for extended documentation.\",\"required\":[\"url\"],\"properties\":{\"description\":{\"$ref\":\"#/definitions/FlowExternalDocumentationObjectDescription\"},\"url\":{\"$ref\":\"#/definitions/FlowExternalDocumentationObjectUrl\"}}}}}"