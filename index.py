from typing import TypedDict
from typing import Optional
from enum import Enum
from typing import NewType
from typing import List
from typing import Any
from typing import Union
from typing import Mapping
"""**REQUIRED**. This string MUST be the semantic version number of the Specification that the document uses.
"""
class FlowVersion(Enum):
    0.0.0-DEVELOPMENT = 0
"""**REQUIRED** Name of the Flow
"""
FlowName = NewType("FlowName", str)

Namespace = NewType("Namespace", str)

class Transport(Enum):
    POSTMESSAGEIFRAME = 0
    POSTMESSAGEWINDOW = 1
    HTTP = 2
    WEBSOCKET = 3

Url = NewType("Url", str)

class StepProvider(TypedDict):
    namespace: Optional[Namespace]
    transport: Optional[Transport]
    url: Optional[Url]

StepProviders = NewType("StepProviders", List[StepProvider])
"""Markdown description describing the specification extension.
"""
FlowDescription = NewType("FlowDescription", str)
"""A short summary of what the flow is.
"""
FlowSummary = NewType("FlowSummary", str)

$Id = NewType("$Id", str)

$Schema = NewType("$Schema", str)

$Ref = NewType("$Ref", str)

$Comment = NewType("$Comment", str)

Title = NewType("Title", str)

Description = NewType("Description", str)
AlwaysTrue = NewType("AlwaysTrue", Any)

ReadOnly = NewType("ReadOnly", bool)

Examples = NewType("Examples", List[AlwaysTrue])

MultipleOf = NewType("MultipleOf", float)

Maximum = NewType("Maximum", float)

ExclusiveMaximum = NewType("ExclusiveMaximum", float)

Minimum = NewType("Minimum", float)

ExclusiveMinimum = NewType("ExclusiveMinimum", float)

NonNegativeInteger = NewType("NonNegativeInteger", int)

NonNegativeIntegerDefaultZero = NewType("NonNegativeIntegerDefaultZero", int)

Pattern = NewType("Pattern", str)

SchemaArray = NewType("SchemaArray", List[JSONSchema])

Items = NewType("Items", Union[JSONSchema, SchemaArray])

UniqueItems = NewType("UniqueItems", bool)

StringDoaGddGA = NewType("StringDoaGddGA", str)

StringArray = NewType("StringArray", List[StringDoaGddGA])

Definitions = NewType("Definitions", Mapping[Any, Any])

Properties = NewType("Properties", Mapping[Any, Any])

PatternProperties = NewType("PatternProperties", Mapping[Any, Any])

DependenciesSet = NewType("DependenciesSet", Union[JSONSchema, StringArray])

Dependencies = NewType("Dependencies", Mapping[Any, Any])

Enum = NewType("Enum", List[AlwaysTrue])

SimpleTypes = NewType("SimpleTypes", Any)

ArrayOfSimpleTypes = NewType("ArrayOfSimpleTypes", List[SimpleTypes])

Type = NewType("Type", Union[SimpleTypes, ArrayOfSimpleTypes])

Format = NewType("Format", str)

ContentMediaType = NewType("ContentMediaType", str)

ContentEncoding = NewType("ContentEncoding", str)

class JSONSchemaObject(TypedDict):
    $id: Optional[$Id]
    $schema: Optional[$Schema]
    $ref: Optional[$Ref]
    $comment: Optional[$Comment]
    title: Optional[Title]
    description: Optional[Description]
    default: Optional[AlwaysTrue]
    readOnly: Optional[ReadOnly]
    examples: Optional[Examples]
    multipleOf: Optional[MultipleOf]
    maximum: Optional[Maximum]
    exclusiveMaximum: Optional[ExclusiveMaximum]
    minimum: Optional[Minimum]
    exclusiveMinimum: Optional[ExclusiveMinimum]
    maxLength: Optional[NonNegativeInteger]
    minLength: Optional[NonNegativeIntegerDefaultZero]
    pattern: Optional[Pattern]
    additionalItems: Optional[JSONSchema]
    items: Optional[Items]
    maxItems: Optional[NonNegativeInteger]
    minItems: Optional[NonNegativeIntegerDefaultZero]
    uniqueItems: Optional[UniqueItems]
    contains: Optional[JSONSchema]
    maxProperties: Optional[NonNegativeInteger]
    minProperties: Optional[NonNegativeIntegerDefaultZero]
    required: Optional[StringArray]
    additionalProperties: Optional[JSONSchema]
    definitions: Optional[Definitions]
    properties: Optional[Properties]
    patternProperties: Optional[PatternProperties]
    dependencies: Optional[Dependencies]
    propertyNames: Optional[JSONSchema]
    const: Optional[AlwaysTrue]
    enum: Optional[Enum]
    type: Optional[Type]
    format: Optional[Format]
    contentMediaType: Optional[ContentMediaType]
    contentEncoding: Optional[ContentEncoding]
    if: Optional[JSONSchema]
    then: Optional[JSONSchema]
    else: Optional[JSONSchema]
    allOf: Optional[SchemaArray]
    anyOf: Optional[SchemaArray]
    oneOf: Optional[SchemaArray]
    not: Optional[JSONSchema]
"""Always valid if true. Never valid if false. Is constant.
"""
JSONSchemaBoolean = NewType("JSONSchemaBoolean", bool)

JSONSchema = NewType("JSONSchema", Union[JSONSchemaObject, JSONSchemaBoolean])
"""Additional context provided to the flow
"""
FlowContext = NewType("FlowContext", Mapping[Any, Any])

ObjectHAgrRKSz = NewType("ObjectHAgrRKSz", Mapping[Any, Any])

SetOfAnykhpArGaq = NewType("SetOfAnykhpArGaq", List[AlwaysTrue])

OneOfObjectHAgrRKSzSetOfAnykhpArGaqQYzyujLV = NewType("OneOfObjectHAgrRKSzSetOfAnykhpArGaqQYzyujLV", Union[ObjectHAgrRKSz, SetOfAnykhpArGaq])

class FlowObject(TypedDict):
    method: Optional[StringDoaGddGA]
    params: Optional[OneOfObjectHAgrRKSzSetOfAnykhpArGaqQYzyujLV]

Flow = NewType("Flow", Mapping[Any, Any])
"""external documentation description.
"""
FlowExternalDocumentationObjectDescription = NewType("FlowExternalDocumentationObjectDescription", str)
"""external documentation description.
"""
FlowExternalDocumentationObjectUrl = NewType("FlowExternalDocumentationObjectUrl", str)
"""Information about external documentation.
"""
class FlowExternalDocumentationObject(TypedDict):
    description: Optional[FlowExternalDocumentationObjectDescription]
    url: Optional[FlowExternalDocumentationObjectUrl]
"""A standard way to define JSON-RPC flows
"""
class FlowMetaSchema(TypedDict):
    rpcflow: Optional[FlowVersion]
    name: Optional[FlowName]
    stepProviders: Optional[StepProviders]
    description: Optional[FlowDescription]
    summary: Optional[FlowSummary]
    contextSchema: Optional[JSONSchema]
    context: Optional[FlowContext]
    flow: Optional[Flow]
    externalDocumentation: Optional[FlowExternalDocumentationObject]