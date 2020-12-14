/**
 *
 * **REQUIRED**. This string MUST be the semantic version number of the Specification that the document uses.
 *
 */
export type FlowVersion = "0.0.0-development";
/**
 *
 * **REQUIRED**. Name of the Flow.
 *
 * @example
 * `foobarbaz`
 *
 */
export type FlowName = string;
export type Namespace = string;
export type Transport = "postmessageiframe" | "postmessagewindow" | "http" | "websocket";
export type Url = string;
export interface StepProvider {
  namespace: Namespace;
  transport: Transport;
  url: Url;
  [k: string]: any;
}
export type StepProviders = StepProvider[];
/**
 *
 * Markdown description describing the specification extension.
 *
 */
export type FlowDescription = string;
/**
 *
 * A short summary of what the flow is.
 *
 */
export type FlowSummary = string;
export type $Id = string;
export type $Schema = string;
export type $Ref = string;
export type $Comment = string;
export type Title = string;
export type Description = string;
type AlwaysTrue = any;
export type ReadOnly = boolean;
export type Examples = AlwaysTrue[];
export type MultipleOf = number;
export type Maximum = number;
export type ExclusiveMaximum = number;
export type Minimum = number;
export type ExclusiveMinimum = number;
export type NonNegativeInteger = number;
export type NonNegativeIntegerDefaultZero = number;
export type Pattern = string;
export type SchemaArray = JSONSchema[];
/**
 *
 * @default true
 *
 */
export type Items = JSONSchema | SchemaArray;
export type UniqueItems = boolean;
export type StringDoaGddGA = string;
/**
 *
 * @default []
 *
 */
export type StringArray = StringDoaGddGA[];
/**
 *
 * @default {}
 *
 */
export interface Definitions { [key: string]: any; }
/**
 *
 * @default {}
 *
 */
export interface Properties { [key: string]: any; }
/**
 *
 * @default {}
 *
 */
export interface PatternProperties { [key: string]: any; }
export type DependenciesSet = JSONSchema | StringArray;
export interface Dependencies { [key: string]: any; }
export type Enum = AlwaysTrue[];
export type SimpleTypes = any;
export type ArrayOfSimpleTypes = SimpleTypes[];
export type Type = SimpleTypes | ArrayOfSimpleTypes;
export type Format = string;
export type ContentMediaType = string;
export type ContentEncoding = string;
export interface JSONSchemaObject {
  $id?: $Id;
  $schema?: $Schema;
  $ref?: $Ref;
  $comment?: $Comment;
  title?: Title;
  description?: Description;
  default?: AlwaysTrue;
  readOnly?: ReadOnly;
  examples?: Examples;
  multipleOf?: MultipleOf;
  maximum?: Maximum;
  exclusiveMaximum?: ExclusiveMaximum;
  minimum?: Minimum;
  exclusiveMinimum?: ExclusiveMinimum;
  maxLength?: NonNegativeInteger;
  minLength?: NonNegativeIntegerDefaultZero;
  pattern?: Pattern;
  additionalItems?: JSONSchema;
  items?: Items;
  maxItems?: NonNegativeInteger;
  minItems?: NonNegativeIntegerDefaultZero;
  uniqueItems?: UniqueItems;
  contains?: JSONSchema;
  maxProperties?: NonNegativeInteger;
  minProperties?: NonNegativeIntegerDefaultZero;
  required?: StringArray;
  additionalProperties?: JSONSchema;
  definitions?: Definitions;
  properties?: Properties;
  patternProperties?: PatternProperties;
  dependencies?: Dependencies;
  propertyNames?: JSONSchema;
  const?: AlwaysTrue;
  enum?: Enum;
  type?: Type;
  format?: Format;
  contentMediaType?: ContentMediaType;
  contentEncoding?: ContentEncoding;
  if?: JSONSchema;
  then?: JSONSchema;
  else?: JSONSchema;
  allOf?: SchemaArray;
  anyOf?: SchemaArray;
  oneOf?: SchemaArray;
  not?: JSONSchema;
  [k: string]: any;
}
/**
 *
 * Always valid if true. Never valid if false. Is constant.
 *
 */
export type JSONSchemaBoolean = boolean;
/**
 *
 * @default {}
 *
 */
export type JSONSchema = JSONSchemaObject | JSONSchemaBoolean;
/**
 *
 * Additional context provided to the flow. If `contextSchema` exists, `context` **MUST** validate against it. Context can be used via the `\"${context.<path>}\"` [Runtime Expression](https://github.com/etclabscore/json-template-language#what-is-it-used-for)
 *
 */
export interface FlowContext { [key: string]: any; }
/**
 *
 * **REQUIRED**. Canonical name for the method call.
 *
 */
export type FlowMethod = string;
export interface ObjectHAgrRKSz { [key: string]: any; }
export type SetOfAnykhpArGaq = AlwaysTrue[];
/**
 *
 * A list of parameters to be called for this method.
 *
 */
export type FlowParams = ObjectHAgrRKSz | SetOfAnykhpArGaq;
export interface FlowObject {
  method: FlowMethod;
  params: FlowParams;
  [k: string]: any;
}
/**
 *
 * Describes how to call JSON-RPC methods at runtime and use their results via [Runtime Expressions](https://github.com/etclabscore/json-template-language#what-is-it-used-for). Flow results can be used via \"${results.<keyName>}\" where `keyName` refers to the key of the flow object to use the results from.
 *
 */
export interface Flow { [key: string]: any; }
/**
 *
 * A verbose explanation of the target documentation. GitHub Flavored Markdown syntax MAY be used for rich text representation.
 *
 */
export type FlowExternalDocumentationObjectDescription = string;
/**
 *
 * **REQUIRED**. The URL for the target documentation. Value MUST be in the format of a URL.
 *
 */
export type FlowExternalDocumentationObjectUrl = string;
/**
 *
 * Allows referencing an external resource for extended documentation.
 *
 */
export interface FlowExternalDocumentationObject {
  description?: FlowExternalDocumentationObjectDescription;
  url: FlowExternalDocumentationObjectUrl;
}
/**
 *
 * A standard way to define JSON-RPC flows.
 *
 */
export interface FlowMetaSchema {
  rpcflow: FlowVersion;
  name: FlowName;
  stepProviders: StepProviders;
  description?: FlowDescription;
  summary?: FlowSummary;
  contextSchema?: JSONSchema;
  context?: FlowContext;
  flow: Flow;
  externalDocumentation?: FlowExternalDocumentationObject;
  [k: string]: any;
}