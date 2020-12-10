extern crate serde_json;

use serde::{Serialize, Deserialize};
use std::collections::HashMap;
/// FlowVersion
///
/// **REQUIRED**. This string MUST be the semantic version number of the Specification that the document uses.
///
#[derive(Serialize, Deserialize)]
pub enum FlowVersion {
    #[serde(rename = "0.0.0-development")]
    00Development,
}
/// FlowName
///
/// **REQUIRED** Name of the Flow
///
/// # Example
///
/// `foobarbaz`
///
pub type FlowName = String;
pub type Namespace = String;
#[derive(Serialize, Deserialize)]
pub enum Transport {
    #[serde(rename = "postmessageiframe")]
    Postmessageiframe,
    #[serde(rename = "postmessagewindow")]
    Postmessagewindow,
    #[serde(rename = "http")]
    Http,
    #[serde(rename = "websocket")]
    Websocket,
}
pub type Url = String;
#[derive(Serialize, Deserialize)]
pub struct StepProvider {
    pub(crate) namespace: Namespace,
    pub(crate) transport: Transport,
    pub(crate) url: Url,
}
pub type StepProviders = Vec<StepProvider>;
/// FlowDescription
///
/// Markdown description describing the specification extension.
///
pub type FlowDescription = String;
/// FlowSummary
///
/// A short summary of what the flow is.
///
pub type FlowSummary = String;
pub type $Id = String;
pub type $Schema = String;
pub type $Ref = String;
pub type $Comment = String;
pub type Title = String;
pub type Description = String;
type AlwaysTrue = serde_json::Value;
pub type ReadOnly = bool;
pub type Examples = Vec<AlwaysTrue>;
pub type MultipleOf = f64;
pub type Maximum = f64;
pub type ExclusiveMaximum = f64;
pub type Minimum = f64;
pub type ExclusiveMinimum = f64;
pub type NonNegativeInteger = i64;
pub type NonNegativeIntegerDefaultZero = i64;
pub type Pattern = String;
pub type SchemaArray = Vec<JSONSchema>;
#[derive(Serialize, Deserialize)]
pub enum Items {
    JSONSchema,
    SchemaArray
}
pub type UniqueItems = bool;
pub type StringDoaGddGA = String;
/// StringArray
///
/// # Default
///
/// []
///
pub type StringArray = Vec<StringDoaGddGA>;
/// Definitions
///
/// # Default
///
/// {}
///
pub type Definitions = HashMap<String, Option<serde_json::Value>>;
/// Properties
///
/// # Default
///
/// {}
///
pub type Properties = HashMap<String, Option<serde_json::Value>>;
/// PatternProperties
///
/// # Default
///
/// {}
///
pub type PatternProperties = HashMap<String, Option<serde_json::Value>>;
#[derive(Serialize, Deserialize)]
pub enum DependenciesSet {
    JSONSchema,
    StringArray
}
pub type Dependencies = HashMap<String, Option<serde_json::Value>>;
pub type Enum = Vec<AlwaysTrue>;
pub type SimpleTypes = serde_json::Value;
pub type ArrayOfSimpleTypes = Vec<SimpleTypes>;
#[derive(Serialize, Deserialize)]
pub enum Type {
    SimpleTypes,
    ArrayOfSimpleTypes
}
pub type Format = String;
pub type ContentMediaType = String;
pub type ContentEncoding = String;
#[derive(Serialize, Deserialize)]
pub struct JSONSchemaObject {
    pub(crate) $id: Option<$Id>,
    pub(crate) $schema: Option<$Schema>,
    pub(crate) $ref: Option<$Ref>,
    pub(crate) $comment: Option<$Comment>,
    pub(crate) title: Option<Title>,
    pub(crate) description: Option<Description>,
    pub(crate) default: Option<AlwaysTrue>,
    pub(crate) readOnly: Option<ReadOnly>,
    pub(crate) examples: Option<Examples>,
    pub(crate) multipleOf: Option<MultipleOf>,
    pub(crate) maximum: Option<Maximum>,
    pub(crate) exclusiveMaximum: Option<ExclusiveMaximum>,
    pub(crate) minimum: Option<Minimum>,
    pub(crate) exclusiveMinimum: Option<ExclusiveMinimum>,
    pub(crate) maxLength: Option<NonNegativeInteger>,
    pub(crate) minLength: Option<NonNegativeIntegerDefaultZero>,
    pub(crate) pattern: Option<Pattern>,
    pub(crate) additionalItems: Option<JSONSchema>,
    pub(crate) items: Option<Items>,
    pub(crate) maxItems: Option<NonNegativeInteger>,
    pub(crate) minItems: Option<NonNegativeIntegerDefaultZero>,
    pub(crate) uniqueItems: Option<UniqueItems>,
    pub(crate) contains: Option<JSONSchema>,
    pub(crate) maxProperties: Option<NonNegativeInteger>,
    pub(crate) minProperties: Option<NonNegativeIntegerDefaultZero>,
    pub(crate) required: Option<StringArray>,
    pub(crate) additionalProperties: Option<JSONSchema>,
    pub(crate) definitions: Option<Definitions>,
    pub(crate) properties: Option<Properties>,
    pub(crate) patternProperties: Option<PatternProperties>,
    pub(crate) dependencies: Option<Dependencies>,
    pub(crate) propertyNames: Option<JSONSchema>,
    pub(crate) const: Option<AlwaysTrue>,
    pub(crate) enum: Option<Enum>,
    pub(crate) type: Option<Type>,
    pub(crate) format: Option<Format>,
    pub(crate) contentMediaType: Option<ContentMediaType>,
    pub(crate) contentEncoding: Option<ContentEncoding>,
    pub(crate) if: Option<JSONSchema>,
    pub(crate) then: Option<JSONSchema>,
    pub(crate) else: Option<JSONSchema>,
    pub(crate) allOf: Option<SchemaArray>,
    pub(crate) anyOf: Option<SchemaArray>,
    pub(crate) oneOf: Option<SchemaArray>,
    pub(crate) not: Option<JSONSchema>,
}
/// JSONSchemaBoolean
///
/// Always valid if true. Never valid if false. Is constant.
///
pub type JSONSchemaBoolean = bool;
#[derive(Serialize, Deserialize)]
pub enum JSONSchema {
    JSONSchemaObject,
    JSONSchemaBoolean
}
/// FlowContext
///
/// Additional context provided to the flow
///
pub type FlowContext = HashMap<String, Option<serde_json::Value>>;
pub type ObjectHAgrRKSz = HashMap<String, Option<serde_json::Value>>;
pub type SetOfAnykhpArGaq = Vec<AlwaysTrue>;
#[derive(Serialize, Deserialize)]
pub enum OneOfObjectHAgrRKSzSetOfAnykhpArGaqQYzyujLV {
    ObjectHAgrRKSz,
    SetOfAnykhpArGaq
}
#[derive(Serialize, Deserialize)]
pub struct FlowObject {
    pub(crate) method: StringDoaGddGA,
    pub(crate) params: OneOfObjectHAgrRKSzSetOfAnykhpArGaqQYzyujLV,
}
pub type Flow = HashMap<String, Option<serde_json::Value>>;
/// FlowExternalDocumentationObjectDescription
///
/// external documentation description.
///
pub type FlowExternalDocumentationObjectDescription = String;
/// FlowExternalDocumentationObjectUrl
///
/// external documentation description.
///
pub type FlowExternalDocumentationObjectUrl = String;
/// FlowExternalDocumentationObject
///
/// Information about external documentation.
///
#[derive(Serialize, Deserialize)]
pub struct FlowExternalDocumentationObject {
    pub(crate) description: Option<FlowExternalDocumentationObjectDescription>,
    pub(crate) url: FlowExternalDocumentationObjectUrl,
}
/// FlowMetaSchema
///
/// A standard way to define JSON-RPC flows
///
#[derive(Serialize, Deserialize)]
pub struct FlowMetaSchema {
    pub(crate) rpcflow: FlowVersion,
    pub(crate) name: FlowName,
    pub(crate) stepProviders: StepProviders,
    pub(crate) description: Option<FlowDescription>,
    pub(crate) summary: Option<FlowSummary>,
    pub(crate) contextSchema: Option<JSONSchema>,
    pub(crate) context: Option<FlowContext>,
    pub(crate) flow: Flow,
    pub(crate) externalDocumentation: Option<FlowExternalDocumentationObject>,
}