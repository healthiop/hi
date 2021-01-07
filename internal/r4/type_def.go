// Copyright (c) 2020-2021, Volker Schmidt (volker@volsch.eu)
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package r4

import (
	"github.com/healthiop/hi/internal/common"
	"regexp"
	"sync"
)

var typeDefsContainerOnce sync.Once
var cachedTypeDefContainer *common.TypeDefContainer

func TypeDefContainer() *common.TypeDefContainer {
	typeDefsContainerOnce.Do(func() {
		cachedTypeDefContainer = createTypeDefContainer()
	})
	return cachedTypeDefContainer
}

func createTypeDefContainer() *common.TypeDefContainer {
	typeDefs := make([]common.TypeDefAccessor, 682)
	typeDefs[0] = common.NewStructTypeDef("Element", nil, false)
	typeDefs[1] = common.NewStructTypeDef("Resource", nil, false)
	// with base type Element
	typeDefs[2] = common.NewStructTypeDef("Address", typeDefs[0], false)
	// with base type Element
	typeDefs[3] = common.NewStructTypeDef("Annotation", typeDefs[0], false)
	// with base type Element
	typeDefs[4] = common.NewStructTypeDef("Attachment", typeDefs[0], false)
	// with base type Element
	typeDefs[5] = common.NewStructTypeDef("BackboneElement", typeDefs[0], false)
	// with base type Resource
	typeDefs[6] = common.NewStructTypeDef("Binary", typeDefs[1], false)
	// with base type Resource
	typeDefs[7] = common.NewStructTypeDef("Bundle", typeDefs[1], false)
	// with base type Element
	typeDefs[8] = common.NewStructTypeDef("CodeableConcept", typeDefs[0], false)
	// with base type Element
	typeDefs[9] = common.NewStructTypeDef("Coding", typeDefs[0], false)
	// with base type Element
	typeDefs[10] = common.NewStructTypeDef("ContactDetail", typeDefs[0], false)
	// with base type Element
	typeDefs[11] = common.NewStructTypeDef("ContactPoint", typeDefs[0], false)
	// with base type Element
	typeDefs[12] = common.NewStructTypeDef("Contributor", typeDefs[0], false)
	// with base type Element
	typeDefs[13] = common.NewStructTypeDef("DataRequirement", typeDefs[0], false)
	// with base type Resource
	typeDefs[14] = common.NewStructTypeDef("DomainResource", typeDefs[1], false)
	// with base type Element
	typeDefs[15] = common.NewStructTypeDef("Expression", typeDefs[0], false)
	// with base type Element
	typeDefs[16] = common.NewStructTypeDef("Extension", typeDefs[0], false)
	// with base type Element
	typeDefs[17] = common.NewStructTypeDef("HumanName", typeDefs[0], false)
	// with base type Element
	typeDefs[18] = common.NewStructTypeDef("Identifier", typeDefs[0], false)
	// with base type Element
	typeDefs[19] = common.NewStructTypeDef("MarketingStatus", typeDefs[0], false)
	// with base type Element
	typeDefs[20] = common.NewStructTypeDef("Meta", typeDefs[0], false)
	// with base type Element
	typeDefs[21] = common.NewStructTypeDef("Money", typeDefs[0], false)
	// with base type Element
	typeDefs[22] = common.NewStructTypeDef("Narrative", typeDefs[0], false)
	// with base type Element
	typeDefs[23] = common.NewStructTypeDef("ParameterDefinition", typeDefs[0], false)
	// with base type Resource
	typeDefs[24] = common.NewStructTypeDef("Parameters", typeDefs[1], false)
	// with base type Element
	typeDefs[25] = common.NewStructTypeDef("Period", typeDefs[0], false)
	// with base type Element
	typeDefs[26] = common.NewStructTypeDef("Population", typeDefs[0], false)
	// with base type Element
	typeDefs[27] = common.NewStructTypeDef("ProdCharacteristic", typeDefs[0], false)
	// with base type Element
	typeDefs[28] = common.NewStructTypeDef("ProductShelfLife", typeDefs[0], false)
	// with base type Element
	typeDefs[29] = common.NewStructTypeDef("Quantity", typeDefs[0], false)
	// with base type Element
	typeDefs[30] = common.NewStructTypeDef("Range", typeDefs[0], false)
	// with base type Element
	typeDefs[31] = common.NewStructTypeDef("Ratio", typeDefs[0], false)
	// with base type Element
	typeDefs[32] = common.NewStructTypeDef("Reference", typeDefs[0], false)
	// with base type Element
	typeDefs[33] = common.NewStructTypeDef("RelatedArtifact", typeDefs[0], false)
	// with base type Element
	typeDefs[34] = common.NewStructTypeDef("SampledData", typeDefs[0], false)
	// with base type Element
	typeDefs[35] = common.NewStructTypeDef("Signature", typeDefs[0], false)
	// with base type Element
	typeDefs[36] = common.NewStructTypeDef("SubstanceAmount", typeDefs[0], false)
	// with base type Element
	typeDefs[37] = common.NewStructTypeDef("TriggerDefinition", typeDefs[0], false)
	// with base type Element
	typeDefs[38] = common.NewStructTypeDef("UsageContext", typeDefs[0], false)
	// with base type Element
	typeDefs[39] = common.NewPrimitiveTypeDef("base64Binary", typeDefs[0], nil, common.StringSimpleType)
	// with base type Element
	typeDefs[40] = common.NewPrimitiveTypeDef("boolean", typeDefs[0], regexp.MustCompile("^true|false$"), common.BoolSimpleType)
	// with base type Element
	typeDefs[41] = common.NewPrimitiveTypeDef("date", typeDefs[0], regexp.MustCompile("^([0-9]([0-9]([0-9][1-9]|[1-9]0)|[1-9]00)|[1-9]000)(-(0[1-9]|1[0-2])(-(0[1-9]|[1-2][0-9]|3[0-1]))?)?$"), common.StringSimpleType)
	// with base type Element
	typeDefs[42] = common.NewPrimitiveTypeDef("dateTime", typeDefs[0], regexp.MustCompile("^([0-9]([0-9]([0-9][1-9]|[1-9]0)|[1-9]00)|[1-9]000)(-(0[1-9]|1[0-2])(-(0[1-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:([0-5][0-9]|60)(\\.[0-9]+)?(Z|([+\\-])((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?$"), common.StringSimpleType)
	// with base type Element
	typeDefs[43] = common.NewPrimitiveTypeDef("decimal", typeDefs[0], regexp.MustCompile("^-?(0|[1-9][0-9]*)(\\.[0-9]+)?([eE][+-]?[0-9]+)?$"), common.NumberSimpleType)
	// with base type Element
	typeDefs[44] = common.NewPrimitiveTypeDef("instant", typeDefs[0], regexp.MustCompile("^([0-9]([0-9]([0-9][1-9]|[1-9]0)|[1-9]00)|[1-9]000)-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])T([01][0-9]|2[0-3]):[0-5][0-9]:([0-5][0-9]|60)(\\.[0-9]+)?(Z|([+\\-])((0[0-9]|1[0-3]):[0-5][0-9]|14:00))$"), common.StringSimpleType)
	// with base type Element
	typeDefs[45] = common.NewPrimitiveTypeDef("integer", typeDefs[0], regexp.MustCompile("^-?([0]|([1-9][0-9]*))$"), common.NumberSimpleType)
	// with base type Element
	typeDefs[46] = common.NewPrimitiveTypeDef("string", typeDefs[0], regexp.MustCompile("^[ \\r\\n\\t\\S]+$"), common.StringSimpleType)
	// with base type Element
	typeDefs[47] = common.NewPrimitiveTypeDef("time", typeDefs[0], regexp.MustCompile("^([01][0-9]|2[0-3]):[0-5][0-9]:([0-5][0-9]|60)(\\.[0-9]+)?$"), common.StringSimpleType)
	// with base type Element
	typeDefs[48] = common.NewPrimitiveTypeDef("uri", typeDefs[0], regexp.MustCompile("^\\S*$"), common.StringSimpleType)
	// with base type Element
	typeDefs[49] = common.NewPrimitiveTypeDef("xhtml", typeDefs[0], nil, common.StringSimpleType)
	// with base type DomainResource
	typeDefs[50] = common.NewStructTypeDef("Account", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[51] = common.NewStructTypeDef("Account_Coverage", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[52] = common.NewStructTypeDef("Account_Guarantor", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[53] = common.NewStructTypeDef("ActivityDefinition", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[54] = common.NewStructTypeDef("ActivityDefinition_DynamicValue", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[55] = common.NewStructTypeDef("ActivityDefinition_Participant", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[56] = common.NewStructTypeDef("AdverseEvent", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[57] = common.NewStructTypeDef("AdverseEvent_Causality", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[58] = common.NewStructTypeDef("AdverseEvent_SuspectEntity", typeDefs[5], true)
	// with base type Quantity
	typeDefs[59] = common.NewStructTypeDef("Age", typeDefs[29], false)
	// with base type DomainResource
	typeDefs[60] = common.NewStructTypeDef("AllergyIntolerance", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[61] = common.NewStructTypeDef("AllergyIntolerance_Reaction", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[62] = common.NewStructTypeDef("Appointment", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[63] = common.NewStructTypeDef("AppointmentResponse", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[64] = common.NewStructTypeDef("Appointment_Participant", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[65] = common.NewStructTypeDef("AuditEvent", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[66] = common.NewStructTypeDef("AuditEvent_Agent", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[67] = common.NewStructTypeDef("AuditEvent_Detail", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[68] = common.NewStructTypeDef("AuditEvent_Entity", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[69] = common.NewStructTypeDef("AuditEvent_Network", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[70] = common.NewStructTypeDef("AuditEvent_Source", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[71] = common.NewStructTypeDef("Basic", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[72] = common.NewStructTypeDef("BiologicallyDerivedProduct", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[73] = common.NewStructTypeDef("BiologicallyDerivedProduct_Collection", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[74] = common.NewStructTypeDef("BiologicallyDerivedProduct_Manipulation", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[75] = common.NewStructTypeDef("BiologicallyDerivedProduct_Processing", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[76] = common.NewStructTypeDef("BiologicallyDerivedProduct_Storage", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[77] = common.NewStructTypeDef("BodyStructure", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[78] = common.NewStructTypeDef("Bundle_Entry", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[79] = common.NewStructTypeDef("Bundle_Link", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[80] = common.NewStructTypeDef("Bundle_Request", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[81] = common.NewStructTypeDef("Bundle_Response", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[82] = common.NewStructTypeDef("Bundle_Search", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[83] = common.NewStructTypeDef("CapabilityStatement", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[84] = common.NewStructTypeDef("CapabilityStatement_Document", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[85] = common.NewStructTypeDef("CapabilityStatement_Endpoint", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[86] = common.NewStructTypeDef("CapabilityStatement_Implementation", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[87] = common.NewStructTypeDef("CapabilityStatement_Interaction", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[88] = common.NewStructTypeDef("CapabilityStatement_Interaction1", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[89] = common.NewStructTypeDef("CapabilityStatement_Messaging", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[90] = common.NewStructTypeDef("CapabilityStatement_Operation", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[91] = common.NewStructTypeDef("CapabilityStatement_Resource", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[92] = common.NewStructTypeDef("CapabilityStatement_Rest", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[93] = common.NewStructTypeDef("CapabilityStatement_SearchParam", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[94] = common.NewStructTypeDef("CapabilityStatement_Security", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[95] = common.NewStructTypeDef("CapabilityStatement_Software", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[96] = common.NewStructTypeDef("CapabilityStatement_SupportedMessage", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[97] = common.NewStructTypeDef("CarePlan", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[98] = common.NewStructTypeDef("CarePlan_Activity", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[99] = common.NewStructTypeDef("CarePlan_Detail", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[100] = common.NewStructTypeDef("CareTeam", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[101] = common.NewStructTypeDef("CareTeam_Participant", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[102] = common.NewStructTypeDef("CatalogEntry", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[103] = common.NewStructTypeDef("CatalogEntry_RelatedEntry", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[104] = common.NewStructTypeDef("ChargeItem", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[105] = common.NewStructTypeDef("ChargeItemDefinition", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[106] = common.NewStructTypeDef("ChargeItemDefinition_Applicability", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[107] = common.NewStructTypeDef("ChargeItemDefinition_PriceComponent", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[108] = common.NewStructTypeDef("ChargeItemDefinition_PropertyGroup", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[109] = common.NewStructTypeDef("ChargeItem_Performer", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[110] = common.NewStructTypeDef("Claim", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[111] = common.NewStructTypeDef("ClaimResponse", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[112] = common.NewStructTypeDef("ClaimResponse_AddItem", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[113] = common.NewStructTypeDef("ClaimResponse_Adjudication", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[114] = common.NewStructTypeDef("ClaimResponse_Detail", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[115] = common.NewStructTypeDef("ClaimResponse_Detail1", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[116] = common.NewStructTypeDef("ClaimResponse_Error", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[117] = common.NewStructTypeDef("ClaimResponse_Insurance", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[118] = common.NewStructTypeDef("ClaimResponse_Item", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[119] = common.NewStructTypeDef("ClaimResponse_Payment", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[120] = common.NewStructTypeDef("ClaimResponse_ProcessNote", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[121] = common.NewStructTypeDef("ClaimResponse_SubDetail", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[122] = common.NewStructTypeDef("ClaimResponse_SubDetail1", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[123] = common.NewStructTypeDef("ClaimResponse_Total", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[124] = common.NewStructTypeDef("Claim_Accident", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[125] = common.NewStructTypeDef("Claim_CareTeam", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[126] = common.NewStructTypeDef("Claim_Detail", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[127] = common.NewStructTypeDef("Claim_Diagnosis", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[128] = common.NewStructTypeDef("Claim_Insurance", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[129] = common.NewStructTypeDef("Claim_Item", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[130] = common.NewStructTypeDef("Claim_Payee", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[131] = common.NewStructTypeDef("Claim_Procedure", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[132] = common.NewStructTypeDef("Claim_Related", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[133] = common.NewStructTypeDef("Claim_SubDetail", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[134] = common.NewStructTypeDef("Claim_SupportingInfo", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[135] = common.NewStructTypeDef("ClinicalImpression", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[136] = common.NewStructTypeDef("ClinicalImpression_Finding", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[137] = common.NewStructTypeDef("ClinicalImpression_Investigation", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[138] = common.NewStructTypeDef("CodeSystem", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[139] = common.NewStructTypeDef("CodeSystem_Concept", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[140] = common.NewStructTypeDef("CodeSystem_Designation", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[141] = common.NewStructTypeDef("CodeSystem_Filter", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[142] = common.NewStructTypeDef("CodeSystem_Property", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[143] = common.NewStructTypeDef("CodeSystem_Property1", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[144] = common.NewStructTypeDef("Communication", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[145] = common.NewStructTypeDef("CommunicationRequest", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[146] = common.NewStructTypeDef("CommunicationRequest_Payload", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[147] = common.NewStructTypeDef("Communication_Payload", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[148] = common.NewStructTypeDef("CompartmentDefinition", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[149] = common.NewStructTypeDef("CompartmentDefinition_Resource", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[150] = common.NewStructTypeDef("Composition", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[151] = common.NewStructTypeDef("Composition_Attester", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[152] = common.NewStructTypeDef("Composition_Event", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[153] = common.NewStructTypeDef("Composition_RelatesTo", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[154] = common.NewStructTypeDef("Composition_Section", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[155] = common.NewStructTypeDef("ConceptMap", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[156] = common.NewStructTypeDef("ConceptMap_DependsOn", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[157] = common.NewStructTypeDef("ConceptMap_Element", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[158] = common.NewStructTypeDef("ConceptMap_Group", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[159] = common.NewStructTypeDef("ConceptMap_Target", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[160] = common.NewStructTypeDef("ConceptMap_Unmapped", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[161] = common.NewStructTypeDef("Condition", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[162] = common.NewStructTypeDef("Condition_Evidence", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[163] = common.NewStructTypeDef("Condition_Stage", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[164] = common.NewStructTypeDef("Consent", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[165] = common.NewStructTypeDef("Consent_Actor", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[166] = common.NewStructTypeDef("Consent_Data", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[167] = common.NewStructTypeDef("Consent_Policy", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[168] = common.NewStructTypeDef("Consent_Provision", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[169] = common.NewStructTypeDef("Consent_Verification", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[170] = common.NewStructTypeDef("Contract", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[171] = common.NewStructTypeDef("Contract_Action", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[172] = common.NewStructTypeDef("Contract_Answer", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[173] = common.NewStructTypeDef("Contract_Asset", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[174] = common.NewStructTypeDef("Contract_ContentDefinition", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[175] = common.NewStructTypeDef("Contract_Context", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[176] = common.NewStructTypeDef("Contract_Friendly", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[177] = common.NewStructTypeDef("Contract_Legal", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[178] = common.NewStructTypeDef("Contract_Offer", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[179] = common.NewStructTypeDef("Contract_Party", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[180] = common.NewStructTypeDef("Contract_Rule", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[181] = common.NewStructTypeDef("Contract_SecurityLabel", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[182] = common.NewStructTypeDef("Contract_Signer", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[183] = common.NewStructTypeDef("Contract_Subject", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[184] = common.NewStructTypeDef("Contract_Term", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[185] = common.NewStructTypeDef("Contract_ValuedItem", typeDefs[5], true)
	// with base type Quantity
	typeDefs[186] = common.NewStructTypeDef("Count", typeDefs[29], false)
	// with base type DomainResource
	typeDefs[187] = common.NewStructTypeDef("Coverage", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[188] = common.NewStructTypeDef("CoverageEligibilityRequest", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[189] = common.NewStructTypeDef("CoverageEligibilityRequest_Diagnosis", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[190] = common.NewStructTypeDef("CoverageEligibilityRequest_Insurance", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[191] = common.NewStructTypeDef("CoverageEligibilityRequest_Item", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[192] = common.NewStructTypeDef("CoverageEligibilityRequest_SupportingInfo", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[193] = common.NewStructTypeDef("CoverageEligibilityResponse", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[194] = common.NewStructTypeDef("CoverageEligibilityResponse_Benefit", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[195] = common.NewStructTypeDef("CoverageEligibilityResponse_Error", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[196] = common.NewStructTypeDef("CoverageEligibilityResponse_Insurance", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[197] = common.NewStructTypeDef("CoverageEligibilityResponse_Item", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[198] = common.NewStructTypeDef("Coverage_Class", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[199] = common.NewStructTypeDef("Coverage_CostToBeneficiary", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[200] = common.NewStructTypeDef("Coverage_Exception", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[201] = common.NewStructTypeDef("DataRequirement_CodeFilter", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[202] = common.NewStructTypeDef("DataRequirement_DateFilter", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[203] = common.NewStructTypeDef("DataRequirement_Sort", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[204] = common.NewStructTypeDef("DetectedIssue", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[205] = common.NewStructTypeDef("DetectedIssue_Evidence", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[206] = common.NewStructTypeDef("DetectedIssue_Mitigation", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[207] = common.NewStructTypeDef("Device", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[208] = common.NewStructTypeDef("DeviceDefinition", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[209] = common.NewStructTypeDef("DeviceDefinition_Capability", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[210] = common.NewStructTypeDef("DeviceDefinition_DeviceName", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[211] = common.NewStructTypeDef("DeviceDefinition_Material", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[212] = common.NewStructTypeDef("DeviceDefinition_Property", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[213] = common.NewStructTypeDef("DeviceDefinition_Specialization", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[214] = common.NewStructTypeDef("DeviceDefinition_UdiDeviceIdentifier", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[215] = common.NewStructTypeDef("DeviceMetric", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[216] = common.NewStructTypeDef("DeviceMetric_Calibration", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[217] = common.NewStructTypeDef("DeviceRequest", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[218] = common.NewStructTypeDef("DeviceRequest_Parameter", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[219] = common.NewStructTypeDef("DeviceUseStatement", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[220] = common.NewStructTypeDef("Device_DeviceName", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[221] = common.NewStructTypeDef("Device_Property", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[222] = common.NewStructTypeDef("Device_Specialization", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[223] = common.NewStructTypeDef("Device_UdiCarrier", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[224] = common.NewStructTypeDef("Device_Version", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[225] = common.NewStructTypeDef("DiagnosticReport", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[226] = common.NewStructTypeDef("DiagnosticReport_Media", typeDefs[5], true)
	// with base type Quantity
	typeDefs[227] = common.NewStructTypeDef("Distance", typeDefs[29], false)
	// with base type DomainResource
	typeDefs[228] = common.NewStructTypeDef("DocumentManifest", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[229] = common.NewStructTypeDef("DocumentManifest_Related", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[230] = common.NewStructTypeDef("DocumentReference", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[231] = common.NewStructTypeDef("DocumentReference_Content", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[232] = common.NewStructTypeDef("DocumentReference_Context", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[233] = common.NewStructTypeDef("DocumentReference_RelatesTo", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[234] = common.NewStructTypeDef("Dosage", typeDefs[5], false)
	// with base type BackboneElement
	typeDefs[235] = common.NewStructTypeDef("Dosage_DoseAndRate", typeDefs[5], true)
	// with base type Quantity
	typeDefs[236] = common.NewStructTypeDef("Duration", typeDefs[29], false)
	// with base type DomainResource
	typeDefs[237] = common.NewStructTypeDef("EffectEvidenceSynthesis", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[238] = common.NewStructTypeDef("EffectEvidenceSynthesis_Certainty", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[239] = common.NewStructTypeDef("EffectEvidenceSynthesis_CertaintySubcomponent", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[240] = common.NewStructTypeDef("EffectEvidenceSynthesis_EffectEstimate", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[241] = common.NewStructTypeDef("EffectEvidenceSynthesis_PrecisionEstimate", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[242] = common.NewStructTypeDef("EffectEvidenceSynthesis_ResultsByExposure", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[243] = common.NewStructTypeDef("EffectEvidenceSynthesis_SampleSize", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[244] = common.NewStructTypeDef("ElementDefinition", typeDefs[5], false)
	// with base type BackboneElement
	typeDefs[245] = common.NewStructTypeDef("ElementDefinition_Base", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[246] = common.NewStructTypeDef("ElementDefinition_Binding", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[247] = common.NewStructTypeDef("ElementDefinition_Constraint", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[248] = common.NewStructTypeDef("ElementDefinition_Discriminator", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[249] = common.NewStructTypeDef("ElementDefinition_Example", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[250] = common.NewStructTypeDef("ElementDefinition_Mapping", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[251] = common.NewStructTypeDef("ElementDefinition_Slicing", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[252] = common.NewStructTypeDef("ElementDefinition_Type", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[253] = common.NewStructTypeDef("Encounter", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[254] = common.NewStructTypeDef("Encounter_ClassHistory", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[255] = common.NewStructTypeDef("Encounter_Diagnosis", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[256] = common.NewStructTypeDef("Encounter_Hospitalization", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[257] = common.NewStructTypeDef("Encounter_Location", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[258] = common.NewStructTypeDef("Encounter_Participant", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[259] = common.NewStructTypeDef("Encounter_StatusHistory", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[260] = common.NewStructTypeDef("Endpoint", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[261] = common.NewStructTypeDef("EnrollmentRequest", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[262] = common.NewStructTypeDef("EnrollmentResponse", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[263] = common.NewStructTypeDef("EpisodeOfCare", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[264] = common.NewStructTypeDef("EpisodeOfCare_Diagnosis", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[265] = common.NewStructTypeDef("EpisodeOfCare_StatusHistory", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[266] = common.NewStructTypeDef("EventDefinition", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[267] = common.NewStructTypeDef("Evidence", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[268] = common.NewStructTypeDef("EvidenceVariable", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[269] = common.NewStructTypeDef("EvidenceVariable_Characteristic", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[270] = common.NewStructTypeDef("ExampleScenario", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[271] = common.NewStructTypeDef("ExampleScenario_Actor", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[272] = common.NewStructTypeDef("ExampleScenario_Alternative", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[273] = common.NewStructTypeDef("ExampleScenario_ContainedInstance", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[274] = common.NewStructTypeDef("ExampleScenario_Instance", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[275] = common.NewStructTypeDef("ExampleScenario_Operation", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[276] = common.NewStructTypeDef("ExampleScenario_Process", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[277] = common.NewStructTypeDef("ExampleScenario_Step", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[278] = common.NewStructTypeDef("ExampleScenario_Version", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[279] = common.NewStructTypeDef("ExplanationOfBenefit", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[280] = common.NewStructTypeDef("ExplanationOfBenefit_Accident", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[281] = common.NewStructTypeDef("ExplanationOfBenefit_AddItem", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[282] = common.NewStructTypeDef("ExplanationOfBenefit_Adjudication", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[283] = common.NewStructTypeDef("ExplanationOfBenefit_BenefitBalance", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[284] = common.NewStructTypeDef("ExplanationOfBenefit_CareTeam", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[285] = common.NewStructTypeDef("ExplanationOfBenefit_Detail", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[286] = common.NewStructTypeDef("ExplanationOfBenefit_Detail1", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[287] = common.NewStructTypeDef("ExplanationOfBenefit_Diagnosis", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[288] = common.NewStructTypeDef("ExplanationOfBenefit_Financial", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[289] = common.NewStructTypeDef("ExplanationOfBenefit_Insurance", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[290] = common.NewStructTypeDef("ExplanationOfBenefit_Item", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[291] = common.NewStructTypeDef("ExplanationOfBenefit_Payee", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[292] = common.NewStructTypeDef("ExplanationOfBenefit_Payment", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[293] = common.NewStructTypeDef("ExplanationOfBenefit_Procedure", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[294] = common.NewStructTypeDef("ExplanationOfBenefit_ProcessNote", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[295] = common.NewStructTypeDef("ExplanationOfBenefit_Related", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[296] = common.NewStructTypeDef("ExplanationOfBenefit_SubDetail", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[297] = common.NewStructTypeDef("ExplanationOfBenefit_SubDetail1", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[298] = common.NewStructTypeDef("ExplanationOfBenefit_SupportingInfo", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[299] = common.NewStructTypeDef("ExplanationOfBenefit_Total", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[300] = common.NewStructTypeDef("FamilyMemberHistory", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[301] = common.NewStructTypeDef("FamilyMemberHistory_Condition", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[302] = common.NewStructTypeDef("Flag", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[303] = common.NewStructTypeDef("Goal", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[304] = common.NewStructTypeDef("Goal_Target", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[305] = common.NewStructTypeDef("GraphDefinition", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[306] = common.NewStructTypeDef("GraphDefinition_Compartment", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[307] = common.NewStructTypeDef("GraphDefinition_Link", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[308] = common.NewStructTypeDef("GraphDefinition_Target", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[309] = common.NewStructTypeDef("Group", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[310] = common.NewStructTypeDef("Group_Characteristic", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[311] = common.NewStructTypeDef("Group_Member", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[312] = common.NewStructTypeDef("GuidanceResponse", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[313] = common.NewStructTypeDef("HealthcareService", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[314] = common.NewStructTypeDef("HealthcareService_AvailableTime", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[315] = common.NewStructTypeDef("HealthcareService_Eligibility", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[316] = common.NewStructTypeDef("HealthcareService_NotAvailable", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[317] = common.NewStructTypeDef("ImagingStudy", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[318] = common.NewStructTypeDef("ImagingStudy_Instance", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[319] = common.NewStructTypeDef("ImagingStudy_Performer", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[320] = common.NewStructTypeDef("ImagingStudy_Series", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[321] = common.NewStructTypeDef("Immunization", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[322] = common.NewStructTypeDef("ImmunizationEvaluation", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[323] = common.NewStructTypeDef("ImmunizationRecommendation", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[324] = common.NewStructTypeDef("ImmunizationRecommendation_DateCriterion", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[325] = common.NewStructTypeDef("ImmunizationRecommendation_Recommendation", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[326] = common.NewStructTypeDef("Immunization_Education", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[327] = common.NewStructTypeDef("Immunization_Performer", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[328] = common.NewStructTypeDef("Immunization_ProtocolApplied", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[329] = common.NewStructTypeDef("Immunization_Reaction", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[330] = common.NewStructTypeDef("ImplementationGuide", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[331] = common.NewStructTypeDef("ImplementationGuide_Definition", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[332] = common.NewStructTypeDef("ImplementationGuide_DependsOn", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[333] = common.NewStructTypeDef("ImplementationGuide_Global", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[334] = common.NewStructTypeDef("ImplementationGuide_Grouping", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[335] = common.NewStructTypeDef("ImplementationGuide_Manifest", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[336] = common.NewStructTypeDef("ImplementationGuide_Page", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[337] = common.NewStructTypeDef("ImplementationGuide_Page1", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[338] = common.NewStructTypeDef("ImplementationGuide_Parameter", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[339] = common.NewStructTypeDef("ImplementationGuide_Resource", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[340] = common.NewStructTypeDef("ImplementationGuide_Resource1", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[341] = common.NewStructTypeDef("ImplementationGuide_Template", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[342] = common.NewStructTypeDef("InsurancePlan", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[343] = common.NewStructTypeDef("InsurancePlan_Benefit", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[344] = common.NewStructTypeDef("InsurancePlan_Benefit1", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[345] = common.NewStructTypeDef("InsurancePlan_Contact", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[346] = common.NewStructTypeDef("InsurancePlan_Cost", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[347] = common.NewStructTypeDef("InsurancePlan_Coverage", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[348] = common.NewStructTypeDef("InsurancePlan_GeneralCost", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[349] = common.NewStructTypeDef("InsurancePlan_Limit", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[350] = common.NewStructTypeDef("InsurancePlan_Plan", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[351] = common.NewStructTypeDef("InsurancePlan_SpecificCost", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[352] = common.NewStructTypeDef("Invoice", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[353] = common.NewStructTypeDef("Invoice_LineItem", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[354] = common.NewStructTypeDef("Invoice_Participant", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[355] = common.NewStructTypeDef("Invoice_PriceComponent", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[356] = common.NewStructTypeDef("Library", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[357] = common.NewStructTypeDef("Linkage", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[358] = common.NewStructTypeDef("Linkage_Item", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[359] = common.NewStructTypeDef("List", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[360] = common.NewStructTypeDef("List_Entry", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[361] = common.NewStructTypeDef("Location", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[362] = common.NewStructTypeDef("Location_HoursOfOperation", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[363] = common.NewStructTypeDef("Location_Position", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[364] = common.NewStructTypeDef("Measure", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[365] = common.NewStructTypeDef("MeasureReport", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[366] = common.NewStructTypeDef("MeasureReport_Component", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[367] = common.NewStructTypeDef("MeasureReport_Group", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[368] = common.NewStructTypeDef("MeasureReport_Population", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[369] = common.NewStructTypeDef("MeasureReport_Population1", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[370] = common.NewStructTypeDef("MeasureReport_Stratifier", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[371] = common.NewStructTypeDef("MeasureReport_Stratum", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[372] = common.NewStructTypeDef("Measure_Component", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[373] = common.NewStructTypeDef("Measure_Group", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[374] = common.NewStructTypeDef("Measure_Population", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[375] = common.NewStructTypeDef("Measure_Stratifier", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[376] = common.NewStructTypeDef("Measure_SupplementalData", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[377] = common.NewStructTypeDef("Media", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[378] = common.NewStructTypeDef("Medication", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[379] = common.NewStructTypeDef("MedicationAdministration", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[380] = common.NewStructTypeDef("MedicationAdministration_Dosage", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[381] = common.NewStructTypeDef("MedicationAdministration_Performer", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[382] = common.NewStructTypeDef("MedicationDispense", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[383] = common.NewStructTypeDef("MedicationDispense_Performer", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[384] = common.NewStructTypeDef("MedicationDispense_Substitution", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[385] = common.NewStructTypeDef("MedicationKnowledge", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[386] = common.NewStructTypeDef("MedicationKnowledge_AdministrationGuidelines", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[387] = common.NewStructTypeDef("MedicationKnowledge_Cost", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[388] = common.NewStructTypeDef("MedicationKnowledge_Dosage", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[389] = common.NewStructTypeDef("MedicationKnowledge_DrugCharacteristic", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[390] = common.NewStructTypeDef("MedicationKnowledge_Ingredient", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[391] = common.NewStructTypeDef("MedicationKnowledge_Kinetics", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[392] = common.NewStructTypeDef("MedicationKnowledge_MaxDispense", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[393] = common.NewStructTypeDef("MedicationKnowledge_MedicineClassification", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[394] = common.NewStructTypeDef("MedicationKnowledge_MonitoringProgram", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[395] = common.NewStructTypeDef("MedicationKnowledge_Monograph", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[396] = common.NewStructTypeDef("MedicationKnowledge_Packaging", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[397] = common.NewStructTypeDef("MedicationKnowledge_PatientCharacteristics", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[398] = common.NewStructTypeDef("MedicationKnowledge_Regulatory", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[399] = common.NewStructTypeDef("MedicationKnowledge_RelatedMedicationKnowledge", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[400] = common.NewStructTypeDef("MedicationKnowledge_Schedule", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[401] = common.NewStructTypeDef("MedicationKnowledge_Substitution", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[402] = common.NewStructTypeDef("MedicationRequest", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[403] = common.NewStructTypeDef("MedicationRequest_DispenseRequest", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[404] = common.NewStructTypeDef("MedicationRequest_InitialFill", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[405] = common.NewStructTypeDef("MedicationRequest_Substitution", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[406] = common.NewStructTypeDef("MedicationStatement", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[407] = common.NewStructTypeDef("Medication_Batch", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[408] = common.NewStructTypeDef("Medication_Ingredient", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[409] = common.NewStructTypeDef("MedicinalProduct", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[410] = common.NewStructTypeDef("MedicinalProductAuthorization", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[411] = common.NewStructTypeDef("MedicinalProductAuthorization_JurisdictionalAuthorization", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[412] = common.NewStructTypeDef("MedicinalProductAuthorization_Procedure", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[413] = common.NewStructTypeDef("MedicinalProductContraindication", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[414] = common.NewStructTypeDef("MedicinalProductContraindication_OtherTherapy", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[415] = common.NewStructTypeDef("MedicinalProductIndication", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[416] = common.NewStructTypeDef("MedicinalProductIndication_OtherTherapy", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[417] = common.NewStructTypeDef("MedicinalProductIngredient", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[418] = common.NewStructTypeDef("MedicinalProductIngredient_ReferenceStrength", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[419] = common.NewStructTypeDef("MedicinalProductIngredient_SpecifiedSubstance", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[420] = common.NewStructTypeDef("MedicinalProductIngredient_Strength", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[421] = common.NewStructTypeDef("MedicinalProductIngredient_Substance", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[422] = common.NewStructTypeDef("MedicinalProductInteraction", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[423] = common.NewStructTypeDef("MedicinalProductInteraction_Interactant", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[424] = common.NewStructTypeDef("MedicinalProductManufactured", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[425] = common.NewStructTypeDef("MedicinalProductPackaged", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[426] = common.NewStructTypeDef("MedicinalProductPackaged_BatchIdentifier", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[427] = common.NewStructTypeDef("MedicinalProductPackaged_PackageItem", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[428] = common.NewStructTypeDef("MedicinalProductPharmaceutical", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[429] = common.NewStructTypeDef("MedicinalProductPharmaceutical_Characteristics", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[430] = common.NewStructTypeDef("MedicinalProductPharmaceutical_RouteOfAdministration", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[431] = common.NewStructTypeDef("MedicinalProductPharmaceutical_TargetSpecies", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[432] = common.NewStructTypeDef("MedicinalProductPharmaceutical_WithdrawalPeriod", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[433] = common.NewStructTypeDef("MedicinalProductUndesirableEffect", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[434] = common.NewStructTypeDef("MedicinalProduct_CountryLanguage", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[435] = common.NewStructTypeDef("MedicinalProduct_ManufacturingBusinessOperation", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[436] = common.NewStructTypeDef("MedicinalProduct_Name", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[437] = common.NewStructTypeDef("MedicinalProduct_NamePart", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[438] = common.NewStructTypeDef("MedicinalProduct_SpecialDesignation", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[439] = common.NewStructTypeDef("MessageDefinition", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[440] = common.NewStructTypeDef("MessageDefinition_AllowedResponse", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[441] = common.NewStructTypeDef("MessageDefinition_Focus", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[442] = common.NewStructTypeDef("MessageHeader", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[443] = common.NewStructTypeDef("MessageHeader_Destination", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[444] = common.NewStructTypeDef("MessageHeader_Response", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[445] = common.NewStructTypeDef("MessageHeader_Source", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[446] = common.NewStructTypeDef("MolecularSequence", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[447] = common.NewStructTypeDef("MolecularSequence_Inner", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[448] = common.NewStructTypeDef("MolecularSequence_Outer", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[449] = common.NewStructTypeDef("MolecularSequence_Quality", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[450] = common.NewStructTypeDef("MolecularSequence_ReferenceSeq", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[451] = common.NewStructTypeDef("MolecularSequence_Repository", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[452] = common.NewStructTypeDef("MolecularSequence_Roc", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[453] = common.NewStructTypeDef("MolecularSequence_StructureVariant", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[454] = common.NewStructTypeDef("MolecularSequence_Variant", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[455] = common.NewStructTypeDef("NamingSystem", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[456] = common.NewStructTypeDef("NamingSystem_UniqueId", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[457] = common.NewStructTypeDef("NutritionOrder", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[458] = common.NewStructTypeDef("NutritionOrder_Administration", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[459] = common.NewStructTypeDef("NutritionOrder_EnteralFormula", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[460] = common.NewStructTypeDef("NutritionOrder_Nutrient", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[461] = common.NewStructTypeDef("NutritionOrder_OralDiet", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[462] = common.NewStructTypeDef("NutritionOrder_Supplement", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[463] = common.NewStructTypeDef("NutritionOrder_Texture", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[464] = common.NewStructTypeDef("Observation", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[465] = common.NewStructTypeDef("ObservationDefinition", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[466] = common.NewStructTypeDef("ObservationDefinition_QualifiedInterval", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[467] = common.NewStructTypeDef("ObservationDefinition_QuantitativeDetails", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[468] = common.NewStructTypeDef("Observation_Component", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[469] = common.NewStructTypeDef("Observation_ReferenceRange", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[470] = common.NewStructTypeDef("OperationDefinition", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[471] = common.NewStructTypeDef("OperationDefinition_Binding", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[472] = common.NewStructTypeDef("OperationDefinition_Overload", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[473] = common.NewStructTypeDef("OperationDefinition_Parameter", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[474] = common.NewStructTypeDef("OperationDefinition_ReferencedFrom", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[475] = common.NewStructTypeDef("OperationOutcome", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[476] = common.NewStructTypeDef("OperationOutcome_Issue", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[477] = common.NewStructTypeDef("Organization", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[478] = common.NewStructTypeDef("OrganizationAffiliation", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[479] = common.NewStructTypeDef("Organization_Contact", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[480] = common.NewStructTypeDef("Parameters_Parameter", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[481] = common.NewStructTypeDef("Patient", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[482] = common.NewStructTypeDef("Patient_Communication", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[483] = common.NewStructTypeDef("Patient_Contact", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[484] = common.NewStructTypeDef("Patient_Link", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[485] = common.NewStructTypeDef("PaymentNotice", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[486] = common.NewStructTypeDef("PaymentReconciliation", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[487] = common.NewStructTypeDef("PaymentReconciliation_Detail", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[488] = common.NewStructTypeDef("PaymentReconciliation_ProcessNote", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[489] = common.NewStructTypeDef("Person", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[490] = common.NewStructTypeDef("Person_Link", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[491] = common.NewStructTypeDef("PlanDefinition", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[492] = common.NewStructTypeDef("PlanDefinition_Action", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[493] = common.NewStructTypeDef("PlanDefinition_Condition", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[494] = common.NewStructTypeDef("PlanDefinition_DynamicValue", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[495] = common.NewStructTypeDef("PlanDefinition_Goal", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[496] = common.NewStructTypeDef("PlanDefinition_Participant", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[497] = common.NewStructTypeDef("PlanDefinition_RelatedAction", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[498] = common.NewStructTypeDef("PlanDefinition_Target", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[499] = common.NewStructTypeDef("Practitioner", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[500] = common.NewStructTypeDef("PractitionerRole", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[501] = common.NewStructTypeDef("PractitionerRole_AvailableTime", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[502] = common.NewStructTypeDef("PractitionerRole_NotAvailable", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[503] = common.NewStructTypeDef("Practitioner_Qualification", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[504] = common.NewStructTypeDef("Procedure", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[505] = common.NewStructTypeDef("Procedure_FocalDevice", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[506] = common.NewStructTypeDef("Procedure_Performer", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[507] = common.NewStructTypeDef("Provenance", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[508] = common.NewStructTypeDef("Provenance_Agent", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[509] = common.NewStructTypeDef("Provenance_Entity", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[510] = common.NewStructTypeDef("Questionnaire", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[511] = common.NewStructTypeDef("QuestionnaireResponse", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[512] = common.NewStructTypeDef("QuestionnaireResponse_Answer", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[513] = common.NewStructTypeDef("QuestionnaireResponse_Item", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[514] = common.NewStructTypeDef("Questionnaire_AnswerOption", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[515] = common.NewStructTypeDef("Questionnaire_EnableWhen", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[516] = common.NewStructTypeDef("Questionnaire_Initial", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[517] = common.NewStructTypeDef("Questionnaire_Item", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[518] = common.NewStructTypeDef("RelatedPerson", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[519] = common.NewStructTypeDef("RelatedPerson_Communication", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[520] = common.NewStructTypeDef("RequestGroup", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[521] = common.NewStructTypeDef("RequestGroup_Action", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[522] = common.NewStructTypeDef("RequestGroup_Condition", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[523] = common.NewStructTypeDef("RequestGroup_RelatedAction", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[524] = common.NewStructTypeDef("ResearchDefinition", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[525] = common.NewStructTypeDef("ResearchElementDefinition", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[526] = common.NewStructTypeDef("ResearchElementDefinition_Characteristic", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[527] = common.NewStructTypeDef("ResearchStudy", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[528] = common.NewStructTypeDef("ResearchStudy_Arm", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[529] = common.NewStructTypeDef("ResearchStudy_Objective", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[530] = common.NewStructTypeDef("ResearchSubject", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[531] = common.NewStructTypeDef("RiskAssessment", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[532] = common.NewStructTypeDef("RiskAssessment_Prediction", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[533] = common.NewStructTypeDef("RiskEvidenceSynthesis", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[534] = common.NewStructTypeDef("RiskEvidenceSynthesis_Certainty", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[535] = common.NewStructTypeDef("RiskEvidenceSynthesis_CertaintySubcomponent", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[536] = common.NewStructTypeDef("RiskEvidenceSynthesis_PrecisionEstimate", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[537] = common.NewStructTypeDef("RiskEvidenceSynthesis_RiskEstimate", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[538] = common.NewStructTypeDef("RiskEvidenceSynthesis_SampleSize", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[539] = common.NewStructTypeDef("Schedule", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[540] = common.NewStructTypeDef("SearchParameter", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[541] = common.NewStructTypeDef("SearchParameter_Component", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[542] = common.NewStructTypeDef("ServiceRequest", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[543] = common.NewStructTypeDef("Slot", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[544] = common.NewStructTypeDef("Specimen", typeDefs[14], false)
	// with base type DomainResource
	typeDefs[545] = common.NewStructTypeDef("SpecimenDefinition", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[546] = common.NewStructTypeDef("SpecimenDefinition_Additive", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[547] = common.NewStructTypeDef("SpecimenDefinition_Container", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[548] = common.NewStructTypeDef("SpecimenDefinition_Handling", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[549] = common.NewStructTypeDef("SpecimenDefinition_TypeTested", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[550] = common.NewStructTypeDef("Specimen_Collection", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[551] = common.NewStructTypeDef("Specimen_Container", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[552] = common.NewStructTypeDef("Specimen_Processing", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[553] = common.NewStructTypeDef("StructureDefinition", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[554] = common.NewStructTypeDef("StructureDefinition_Context", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[555] = common.NewStructTypeDef("StructureDefinition_Differential", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[556] = common.NewStructTypeDef("StructureDefinition_Mapping", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[557] = common.NewStructTypeDef("StructureDefinition_Snapshot", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[558] = common.NewStructTypeDef("StructureMap", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[559] = common.NewStructTypeDef("StructureMap_Dependent", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[560] = common.NewStructTypeDef("StructureMap_Group", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[561] = common.NewStructTypeDef("StructureMap_Input", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[562] = common.NewStructTypeDef("StructureMap_Parameter", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[563] = common.NewStructTypeDef("StructureMap_Rule", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[564] = common.NewStructTypeDef("StructureMap_Source", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[565] = common.NewStructTypeDef("StructureMap_Structure", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[566] = common.NewStructTypeDef("StructureMap_Target", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[567] = common.NewStructTypeDef("Subscription", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[568] = common.NewStructTypeDef("Subscription_Channel", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[569] = common.NewStructTypeDef("Substance", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[570] = common.NewStructTypeDef("SubstanceAmount_ReferenceRange", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[571] = common.NewStructTypeDef("SubstanceNucleicAcid", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[572] = common.NewStructTypeDef("SubstanceNucleicAcid_Linkage", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[573] = common.NewStructTypeDef("SubstanceNucleicAcid_Subunit", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[574] = common.NewStructTypeDef("SubstanceNucleicAcid_Sugar", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[575] = common.NewStructTypeDef("SubstancePolymer", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[576] = common.NewStructTypeDef("SubstancePolymer_DegreeOfPolymerisation", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[577] = common.NewStructTypeDef("SubstancePolymer_MonomerSet", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[578] = common.NewStructTypeDef("SubstancePolymer_Repeat", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[579] = common.NewStructTypeDef("SubstancePolymer_RepeatUnit", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[580] = common.NewStructTypeDef("SubstancePolymer_StartingMaterial", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[581] = common.NewStructTypeDef("SubstancePolymer_StructuralRepresentation", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[582] = common.NewStructTypeDef("SubstanceProtein", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[583] = common.NewStructTypeDef("SubstanceProtein_Subunit", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[584] = common.NewStructTypeDef("SubstanceReferenceInformation", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[585] = common.NewStructTypeDef("SubstanceReferenceInformation_Classification", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[586] = common.NewStructTypeDef("SubstanceReferenceInformation_Gene", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[587] = common.NewStructTypeDef("SubstanceReferenceInformation_GeneElement", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[588] = common.NewStructTypeDef("SubstanceReferenceInformation_Target", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[589] = common.NewStructTypeDef("SubstanceSourceMaterial", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[590] = common.NewStructTypeDef("SubstanceSourceMaterial_Author", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[591] = common.NewStructTypeDef("SubstanceSourceMaterial_FractionDescription", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[592] = common.NewStructTypeDef("SubstanceSourceMaterial_Hybrid", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[593] = common.NewStructTypeDef("SubstanceSourceMaterial_Organism", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[594] = common.NewStructTypeDef("SubstanceSourceMaterial_OrganismGeneral", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[595] = common.NewStructTypeDef("SubstanceSourceMaterial_PartDescription", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[596] = common.NewStructTypeDef("SubstanceSpecification", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[597] = common.NewStructTypeDef("SubstanceSpecification_Code", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[598] = common.NewStructTypeDef("SubstanceSpecification_Isotope", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[599] = common.NewStructTypeDef("SubstanceSpecification_Moiety", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[600] = common.NewStructTypeDef("SubstanceSpecification_MolecularWeight", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[601] = common.NewStructTypeDef("SubstanceSpecification_Name", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[602] = common.NewStructTypeDef("SubstanceSpecification_Official", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[603] = common.NewStructTypeDef("SubstanceSpecification_Property", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[604] = common.NewStructTypeDef("SubstanceSpecification_Relationship", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[605] = common.NewStructTypeDef("SubstanceSpecification_Representation", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[606] = common.NewStructTypeDef("SubstanceSpecification_Structure", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[607] = common.NewStructTypeDef("Substance_Ingredient", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[608] = common.NewStructTypeDef("Substance_Instance", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[609] = common.NewStructTypeDef("SupplyDelivery", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[610] = common.NewStructTypeDef("SupplyDelivery_SuppliedItem", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[611] = common.NewStructTypeDef("SupplyRequest", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[612] = common.NewStructTypeDef("SupplyRequest_Parameter", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[613] = common.NewStructTypeDef("Task", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[614] = common.NewStructTypeDef("Task_Input", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[615] = common.NewStructTypeDef("Task_Output", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[616] = common.NewStructTypeDef("Task_Restriction", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[617] = common.NewStructTypeDef("TerminologyCapabilities", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[618] = common.NewStructTypeDef("TerminologyCapabilities_Closure", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[619] = common.NewStructTypeDef("TerminologyCapabilities_CodeSystem", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[620] = common.NewStructTypeDef("TerminologyCapabilities_Expansion", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[621] = common.NewStructTypeDef("TerminologyCapabilities_Filter", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[622] = common.NewStructTypeDef("TerminologyCapabilities_Implementation", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[623] = common.NewStructTypeDef("TerminologyCapabilities_Parameter", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[624] = common.NewStructTypeDef("TerminologyCapabilities_Software", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[625] = common.NewStructTypeDef("TerminologyCapabilities_Translation", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[626] = common.NewStructTypeDef("TerminologyCapabilities_ValidateCode", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[627] = common.NewStructTypeDef("TerminologyCapabilities_Version", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[628] = common.NewStructTypeDef("TestReport", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[629] = common.NewStructTypeDef("TestReport_Action", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[630] = common.NewStructTypeDef("TestReport_Action1", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[631] = common.NewStructTypeDef("TestReport_Action2", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[632] = common.NewStructTypeDef("TestReport_Assert", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[633] = common.NewStructTypeDef("TestReport_Operation", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[634] = common.NewStructTypeDef("TestReport_Participant", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[635] = common.NewStructTypeDef("TestReport_Setup", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[636] = common.NewStructTypeDef("TestReport_Teardown", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[637] = common.NewStructTypeDef("TestReport_Test", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[638] = common.NewStructTypeDef("TestScript", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[639] = common.NewStructTypeDef("TestScript_Action", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[640] = common.NewStructTypeDef("TestScript_Action1", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[641] = common.NewStructTypeDef("TestScript_Action2", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[642] = common.NewStructTypeDef("TestScript_Assert", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[643] = common.NewStructTypeDef("TestScript_Capability", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[644] = common.NewStructTypeDef("TestScript_Destination", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[645] = common.NewStructTypeDef("TestScript_Fixture", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[646] = common.NewStructTypeDef("TestScript_Link", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[647] = common.NewStructTypeDef("TestScript_Metadata", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[648] = common.NewStructTypeDef("TestScript_Operation", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[649] = common.NewStructTypeDef("TestScript_Origin", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[650] = common.NewStructTypeDef("TestScript_RequestHeader", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[651] = common.NewStructTypeDef("TestScript_Setup", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[652] = common.NewStructTypeDef("TestScript_Teardown", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[653] = common.NewStructTypeDef("TestScript_Test", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[654] = common.NewStructTypeDef("TestScript_Variable", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[655] = common.NewStructTypeDef("Timing", typeDefs[5], false)
	// with base type BackboneElement
	typeDefs[656] = common.NewStructTypeDef("Timing_Repeat", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[657] = common.NewStructTypeDef("ValueSet", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[658] = common.NewStructTypeDef("ValueSet_Compose", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[659] = common.NewStructTypeDef("ValueSet_Concept", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[660] = common.NewStructTypeDef("ValueSet_Contains", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[661] = common.NewStructTypeDef("ValueSet_Designation", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[662] = common.NewStructTypeDef("ValueSet_Expansion", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[663] = common.NewStructTypeDef("ValueSet_Filter", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[664] = common.NewStructTypeDef("ValueSet_Include", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[665] = common.NewStructTypeDef("ValueSet_Parameter", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[666] = common.NewStructTypeDef("VerificationResult", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[667] = common.NewStructTypeDef("VerificationResult_Attestation", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[668] = common.NewStructTypeDef("VerificationResult_PrimarySource", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[669] = common.NewStructTypeDef("VerificationResult_Validator", typeDefs[5], true)
	// with base type DomainResource
	typeDefs[670] = common.NewStructTypeDef("VisionPrescription", typeDefs[14], false)
	// with base type BackboneElement
	typeDefs[671] = common.NewStructTypeDef("VisionPrescription_LensSpecification", typeDefs[5], true)
	// with base type BackboneElement
	typeDefs[672] = common.NewStructTypeDef("VisionPrescription_Prism", typeDefs[5], true)
	// with base type uri
	typeDefs[673] = common.NewPrimitiveTypeDef("canonical", typeDefs[48], regexp.MustCompile("^\\S*$"), common.StringSimpleType)
	// with base type string
	typeDefs[674] = common.NewPrimitiveTypeDef("code", typeDefs[46], regexp.MustCompile("^[^\\s]+(\\s[^\\s]+)*$"), common.StringSimpleType)
	// with base type string
	typeDefs[675] = common.NewPrimitiveTypeDef("id", typeDefs[46], regexp.MustCompile("^[A-Za-z0-9\\-.]{1,64}$"), common.StringSimpleType)
	// with base type string
	typeDefs[676] = common.NewPrimitiveTypeDef("markdown", typeDefs[46], regexp.MustCompile("^[ \\r\\n\\t\\S]+$"), common.StringSimpleType)
	// with base type uri
	typeDefs[677] = common.NewPrimitiveTypeDef("oid", typeDefs[48], regexp.MustCompile("^urn:oid:[0-2](\\.(0|[1-9][0-9]*))+$"), common.StringSimpleType)
	// with base type integer
	typeDefs[678] = common.NewPrimitiveTypeDef("positiveInt", typeDefs[45], regexp.MustCompile("^[1-9][0-9]*$"), common.NumberSimpleType)
	// with base type integer
	typeDefs[679] = common.NewPrimitiveTypeDef("unsignedInt", typeDefs[45], regexp.MustCompile("^[0]|([1-9][0-9]*)$"), common.NumberSimpleType)
	// with base type uri
	typeDefs[680] = common.NewPrimitiveTypeDef("url", typeDefs[48], regexp.MustCompile("^\\S*$"), common.StringSimpleType)
	// with base type uri
	typeDefs[681] = common.NewPrimitiveTypeDef("uuid", typeDefs[48], regexp.MustCompile("^urn:uuid:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"), common.StringSimpleType)

	// properties of Account
	typeDefs[50].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Account_Coverage
		common.NewPropDef("coverage", typeDefs[51], "", true, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Account_Guarantor
		common.NewPropDef("guarantor", typeDefs[52], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("owner", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("partOf", typeDefs[32], "", false, nil),
		// with type Period
		common.NewPropDef("servicePeriod", typeDefs[25], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"active",
			"inactive",
			"entered-in-error",
			"on-hold",
			"unknown",
		}),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of Account_Coverage
	typeDefs[51].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("coverage", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type positiveInt
		common.NewPropDef("priority", typeDefs[678], "", false, nil),
	})
	// properties of Account_Guarantor
	typeDefs[52].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("onHold", typeDefs[40], "", false, nil),
		// with type Reference
		common.NewPropDef("party", typeDefs[32], "", false, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
	})
	// properties of ActivityDefinition
	typeDefs[53].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type date
		common.NewPropDef("approvalDate", typeDefs[41], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("author", typeDefs[10], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("bodySite", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type boolean
		common.NewPropDef("doNotPerform", typeDefs[40], "", false, nil),
		// with type Dosage
		common.NewPropDef("dosage", typeDefs[234], "", true, nil),
		// with type ActivityDefinition_DynamicValue
		common.NewPropDef("dynamicValue", typeDefs[54], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("editor", typeDefs[10], "", true, nil),
		// with type Period
		common.NewPropDef("effectivePeriod", typeDefs[25], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("endorser", typeDefs[10], "", true, nil),
		// with type boolean
		common.NewPropDef("experimental", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("intent", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("kind", typeDefs[674], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type date
		common.NewPropDef("lastReviewDate", typeDefs[41], "", false, nil),
		// with type canonical
		common.NewPropDef("library", typeDefs[673], "", true, nil),
		// with type Reference
		common.NewPropDef("location", typeDefs[32], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("observationRequirement", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("observationResultRequirement", typeDefs[32], "", true, nil),
		// with type ActivityDefinition_Participant
		common.NewPropDef("participant", typeDefs[55], "", true, nil),
		// with type code
		common.NewPropDef("priority", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("productCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("productReference", typeDefs[32], "", false, nil),
		// with type canonical
		common.NewPropDef("profile", typeDefs[673], "", false, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("purpose", typeDefs[676], "", false, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type RelatedArtifact
		common.NewPropDef("relatedArtifact", typeDefs[33], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("reviewer", typeDefs[10], "", true, nil),
		// with type Reference
		common.NewPropDef("specimenRequirement", typeDefs[32], "", true, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type CodeableConcept
		common.NewPropDef("subjectCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("subjectReference", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("subtitle", typeDefs[46], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type Age
		common.NewPropDef("timingAge", typeDefs[59], "", false, nil),
		// with type dateTime
		common.NewPropDef("timingDateTime", typeDefs[42], "", false, nil),
		// with type Duration
		common.NewPropDef("timingDuration", typeDefs[236], "", false, nil),
		// with type Period
		common.NewPropDef("timingPeriod", typeDefs[25], "", false, nil),
		// with type Range
		common.NewPropDef("timingRange", typeDefs[30], "", false, nil),
		// with type Timing
		common.NewPropDef("timingTiming", typeDefs[655], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("topic", typeDefs[8], "", true, nil),
		// with type canonical
		common.NewPropDef("transform", typeDefs[673], "", false, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type string
		common.NewPropDef("usage", typeDefs[46], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of ActivityDefinition_DynamicValue
	typeDefs[54].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Expression
		common.NewPropDef("expression", typeDefs[15], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("path", typeDefs[46], "", false, nil),
	})
	// properties of ActivityDefinition_Participant
	typeDefs[55].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("role", typeDefs[8], "", false, nil),
		// with type code
		common.NewPropDef("type", typeDefs[674], "", false, nil),
	})
	// properties of Address
	typeDefs[2].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("city", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("country", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("district", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("line", typeDefs[46], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type string
		common.NewPropDef("postalCode", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("state", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("text", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"postal",
			"physical",
			"both",
		}),
		// with type string
		common.NewPropDef("use", typeDefs[46], "", false, []string{
			"home",
			"work",
			"temp",
			"old",
			"billing",
		}),
	})
	// properties of AdverseEvent
	typeDefs[56].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("actuality", typeDefs[46], "", false, []string{
			"actual",
			"potential",
		}),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("contributor", typeDefs[32], "", true, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type dateTime
		common.NewPropDef("detected", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("event", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("location", typeDefs[32], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("outcome", typeDefs[8], "", false, nil),
		// with type dateTime
		common.NewPropDef("recordedDate", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("recorder", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("referenceDocument", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("resultingCondition", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("seriousness", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("severity", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("study", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("subjectMedicalHistory", typeDefs[32], "", true, nil),
		// with type AdverseEvent_SuspectEntity
		common.NewPropDef("suspectEntity", typeDefs[58], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of AdverseEvent_Causality
	typeDefs[57].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("assessment", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("author", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("method", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("productRelatedness", typeDefs[46], "", false, nil),
	})
	// properties of AdverseEvent_SuspectEntity
	typeDefs[58].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type AdverseEvent_Causality
		common.NewPropDef("causality", typeDefs[57], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("instance", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Age
	typeDefs[59].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("code", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("comparator", typeDefs[46], "", false, []string{
			"<",
			"<=",
			">=",
			">",
		}),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("system", typeDefs[48], "", false, nil),
		// with type string
		common.NewPropDef("unit", typeDefs[46], "", false, nil),
		// with type decimal
		common.NewPropDef("value", typeDefs[43], "", false, nil),
	})
	// properties of AllergyIntolerance
	typeDefs[60].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("asserter", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("category", typeDefs[46], "", true, []string{
			"food",
			"medication",
			"environment",
			"biologic",
		}),
		// with type CodeableConcept
		common.NewPropDef("clinicalStatus", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type string
		common.NewPropDef("criticality", typeDefs[46], "", false, []string{
			"low",
			"high",
			"unable-to-assess",
		}),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type dateTime
		common.NewPropDef("lastOccurrence", typeDefs[42], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type Age
		common.NewPropDef("onsetAge", typeDefs[59], "", false, nil),
		// with type dateTime
		common.NewPropDef("onsetDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("onsetPeriod", typeDefs[25], "", false, nil),
		// with type Range
		common.NewPropDef("onsetRange", typeDefs[30], "", false, nil),
		// with type string
		common.NewPropDef("onsetString", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("patient", typeDefs[32], "", false, nil),
		// with type AllergyIntolerance_Reaction
		common.NewPropDef("reaction", typeDefs[61], "", true, nil),
		// with type dateTime
		common.NewPropDef("recordedDate", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("recorder", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"allergy",
			"intolerance",
		}),
		// with type CodeableConcept
		common.NewPropDef("verificationStatus", typeDefs[8], "", false, nil),
	})
	// properties of AllergyIntolerance_Reaction
	typeDefs[61].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("exposureRoute", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("manifestation", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type dateTime
		common.NewPropDef("onset", typeDefs[42], "", false, nil),
		// with type string
		common.NewPropDef("severity", typeDefs[46], "", false, []string{
			"mild",
			"moderate",
			"severe",
		}),
		// with type CodeableConcept
		common.NewPropDef("substance", typeDefs[8], "", false, nil),
	})
	// properties of Annotation
	typeDefs[3].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("authorReference", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("authorString", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("text", typeDefs[676], "", false, nil),
		// with type dateTime
		common.NewPropDef("time", typeDefs[42], "", false, nil),
	})
	// properties of Appointment
	typeDefs[62].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("appointmentType", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("basedOn", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("cancelationReason", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("comment", typeDefs[46], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type dateTime
		common.NewPropDef("created", typeDefs[42], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type instant
		common.NewPropDef("end", typeDefs[44], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type positiveInt
		common.NewPropDef("minutesDuration", typeDefs[678], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Appointment_Participant
		common.NewPropDef("participant", typeDefs[64], "", true, nil),
		// with type string
		common.NewPropDef("patientInstruction", typeDefs[46], "", false, nil),
		// with type unsignedInt
		common.NewPropDef("priority", typeDefs[679], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("reasonCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("reasonReference", typeDefs[32], "", true, nil),
		// with type Period
		common.NewPropDef("requestedPeriod", typeDefs[25], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("serviceCategory", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("serviceType", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("slot", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("specialty", typeDefs[8], "", true, nil),
		// with type instant
		common.NewPropDef("start", typeDefs[44], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"proposed",
			"pending",
			"booked",
			"arrived",
			"fulfilled",
			"cancelled",
			"noshow",
			"entered-in-error",
			"checked-in",
			"waitlist",
		}),
		// with type Reference
		common.NewPropDef("supportingInformation", typeDefs[32], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of AppointmentResponse
	typeDefs[63].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("actor", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("appointment", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("comment", typeDefs[46], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type instant
		common.NewPropDef("end", typeDefs[44], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type code
		common.NewPropDef("participantStatus", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("participantType", typeDefs[8], "", true, nil),
		// with type instant
		common.NewPropDef("start", typeDefs[44], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of Appointment_Participant
	typeDefs[64].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("actor", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type string
		common.NewPropDef("required", typeDefs[46], "", false, []string{
			"required",
			"optional",
			"information-only",
		}),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"accepted",
			"declined",
			"tentative",
			"needs-action",
		}),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", true, nil),
	})
	// properties of Attachment
	typeDefs[4].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("contentType", typeDefs[674], "", false, nil),
		// with type dateTime
		common.NewPropDef("creation", typeDefs[42], "", false, nil),
		// with type base64Binary
		common.NewPropDef("data", typeDefs[39], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type base64Binary
		common.NewPropDef("hash", typeDefs[39], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type unsignedInt
		common.NewPropDef("size", typeDefs[679], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type url
		common.NewPropDef("url", typeDefs[680], "", false, nil),
	})
	// properties of AuditEvent
	typeDefs[65].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("action", typeDefs[46], "", false, []string{
			"C",
			"R",
			"U",
			"D",
			"E",
		}),
		// with type AuditEvent_Agent
		common.NewPropDef("agent", typeDefs[66], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type AuditEvent_Entity
		common.NewPropDef("entity", typeDefs[68], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("outcome", typeDefs[46], "", false, []string{
			"0",
			"4",
			"8",
			"12",
		}),
		// with type string
		common.NewPropDef("outcomeDesc", typeDefs[46], "", false, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("purposeOfEvent", typeDefs[8], "", true, nil),
		// with type instant
		common.NewPropDef("recorded", typeDefs[44], "", false, nil),
		// with type AuditEvent_Source
		common.NewPropDef("source", typeDefs[70], "", false, nil),
		// with type Coding
		common.NewPropDef("subtype", typeDefs[9], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type Coding
		common.NewPropDef("type", typeDefs[9], "", false, nil),
	})
	// properties of AuditEvent_Agent
	typeDefs[66].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("altId", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("location", typeDefs[32], "", false, nil),
		// with type Coding
		common.NewPropDef("media", typeDefs[9], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type AuditEvent_Network
		common.NewPropDef("network", typeDefs[69], "", false, nil),
		// with type uri
		common.NewPropDef("policy", typeDefs[48], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("purposeOfUse", typeDefs[8], "", true, nil),
		// with type boolean
		common.NewPropDef("requestor", typeDefs[40], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("role", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("who", typeDefs[32], "", false, nil),
	})
	// properties of AuditEvent_Detail
	typeDefs[67].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, nil),
		// with type base64Binary
		common.NewPropDef("valueBase64Binary", typeDefs[39], "", false, nil),
		// with type string
		common.NewPropDef("valueString", typeDefs[46], "", false, nil),
	})
	// properties of AuditEvent_Entity
	typeDefs[68].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type AuditEvent_Detail
		common.NewPropDef("detail", typeDefs[67], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Coding
		common.NewPropDef("lifecycle", typeDefs[9], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type base64Binary
		common.NewPropDef("query", typeDefs[39], "", false, nil),
		// with type Coding
		common.NewPropDef("role", typeDefs[9], "", false, nil),
		// with type Coding
		common.NewPropDef("securityLabel", typeDefs[9], "", true, nil),
		// with type Coding
		common.NewPropDef("type", typeDefs[9], "", false, nil),
		// with type Reference
		common.NewPropDef("what", typeDefs[32], "", false, nil),
	})
	// properties of AuditEvent_Network
	typeDefs[69].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("address", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"1",
			"2",
			"3",
			"4",
			"5",
		}),
	})
	// properties of AuditEvent_Source
	typeDefs[70].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("observer", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("site", typeDefs[46], "", false, nil),
		// with type Coding
		common.NewPropDef("type", typeDefs[9], "", true, nil),
	})
	// properties of BackboneElement
	typeDefs[5].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Basic
	typeDefs[71].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("author", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type date
		common.NewPropDef("created", typeDefs[41], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of Binary
	typeDefs[6].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("contentType", typeDefs[674], "", false, nil),
		// with type base64Binary
		common.NewPropDef("data", typeDefs[39], "", false, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Reference
		common.NewPropDef("securityContext", typeDefs[32], "", false, nil),
	})
	// properties of BiologicallyDerivedProduct
	typeDefs[72].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type BiologicallyDerivedProduct_Collection
		common.NewPropDef("collection", typeDefs[73], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type BiologicallyDerivedProduct_Manipulation
		common.NewPropDef("manipulation", typeDefs[74], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("parent", typeDefs[32], "", true, nil),
		// with type BiologicallyDerivedProduct_Processing
		common.NewPropDef("processing", typeDefs[75], "", true, nil),
		// with type string
		common.NewPropDef("productCategory", typeDefs[46], "", false, []string{
			"organ",
			"tissue",
			"fluid",
			"cells",
			"biologicalAgent",
		}),
		// with type CodeableConcept
		common.NewPropDef("productCode", typeDefs[8], "", false, nil),
		// with type integer
		common.NewPropDef("quantity", typeDefs[45], "", false, nil),
		// with type Reference
		common.NewPropDef("request", typeDefs[32], "", true, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"available",
			"unavailable",
		}),
		// with type BiologicallyDerivedProduct_Storage
		common.NewPropDef("storage", typeDefs[76], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of BiologicallyDerivedProduct_Collection
	typeDefs[73].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type dateTime
		common.NewPropDef("collectedDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("collectedPeriod", typeDefs[25], "", false, nil),
		// with type Reference
		common.NewPropDef("collector", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("source", typeDefs[32], "", false, nil),
	})
	// properties of BiologicallyDerivedProduct_Manipulation
	typeDefs[74].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type dateTime
		common.NewPropDef("timeDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("timePeriod", typeDefs[25], "", false, nil),
	})
	// properties of BiologicallyDerivedProduct_Processing
	typeDefs[75].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("additive", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("procedure", typeDefs[8], "", false, nil),
		// with type dateTime
		common.NewPropDef("timeDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("timePeriod", typeDefs[25], "", false, nil),
	})
	// properties of BiologicallyDerivedProduct_Storage
	typeDefs[76].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Period
		common.NewPropDef("duration", typeDefs[25], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("scale", typeDefs[46], "", false, []string{
			"farenheit",
			"celsius",
			"kelvin",
		}),
		// with type decimal
		common.NewPropDef("temperature", typeDefs[43], "", false, nil),
	})
	// properties of BodyStructure
	typeDefs[77].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("active", typeDefs[40], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type Attachment
		common.NewPropDef("image", typeDefs[4], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("location", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("locationQualifier", typeDefs[8], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("morphology", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("patient", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of Bundle
	typeDefs[7].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Bundle_Entry
		common.NewPropDef("entry", typeDefs[78], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Bundle_Link
		common.NewPropDef("link", typeDefs[79], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Signature
		common.NewPropDef("signature", typeDefs[35], "", false, nil),
		// with type instant
		common.NewPropDef("timestamp", typeDefs[44], "", false, nil),
		// with type unsignedInt
		common.NewPropDef("total", typeDefs[679], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"document",
			"message",
			"transaction",
			"transaction-response",
			"batch",
			"batch-response",
			"history",
			"searchset",
			"collection",
		}),
	})
	// properties of Bundle_Entry
	typeDefs[78].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type uri
		common.NewPropDef("fullUrl", typeDefs[48], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Bundle_Link
		common.NewPropDef("link", typeDefs[79], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Bundle_Request
		common.NewPropDef("request", typeDefs[80], "", false, nil),
		// with type Resource
		common.NewPropDef("resource", typeDefs[1], "", false, nil),
		// with type Bundle_Response
		common.NewPropDef("response", typeDefs[81], "", false, nil),
		// with type Bundle_Search
		common.NewPropDef("search", typeDefs[82], "", false, nil),
	})
	// properties of Bundle_Link
	typeDefs[79].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("relation", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
	})
	// properties of Bundle_Request
	typeDefs[80].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("ifMatch", typeDefs[46], "", false, nil),
		// with type instant
		common.NewPropDef("ifModifiedSince", typeDefs[44], "", false, nil),
		// with type string
		common.NewPropDef("ifNoneExist", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("ifNoneMatch", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("method", typeDefs[46], "", false, []string{
			"GET",
			"HEAD",
			"POST",
			"PUT",
			"DELETE",
			"PATCH",
		}),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
	})
	// properties of Bundle_Response
	typeDefs[81].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("etag", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type instant
		common.NewPropDef("lastModified", typeDefs[44], "", false, nil),
		// with type uri
		common.NewPropDef("location", typeDefs[48], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Resource
		common.NewPropDef("outcome", typeDefs[1], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, nil),
	})
	// properties of Bundle_Search
	typeDefs[82].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("mode", typeDefs[46], "", false, []string{
			"match",
			"include",
			"outcome",
		}),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("score", typeDefs[43], "", false, nil),
	})
	// properties of CapabilityStatement
	typeDefs[83].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type CapabilityStatement_Document
		common.NewPropDef("document", typeDefs[84], "", true, nil),
		// with type boolean
		common.NewPropDef("experimental", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("fhirVersion", typeDefs[46], "", false, []string{
			"0.01",
			"0.05",
			"0.06",
			"0.11",
			"0.0.80",
			"0.0.81",
			"0.0.82",
			"0.4.0",
			"0.5.0",
			"1.0.0",
			"1.0.1",
			"1.0.2",
			"1.1.0",
			"1.4.0",
			"1.6.0",
			"1.8.0",
			"3.0.0",
			"3.0.1",
			"3.3.0",
			"3.5.0",
			"4.0.0",
			"4.0.1",
		}),
		// with type code
		common.NewPropDef("format", typeDefs[674], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type CapabilityStatement_Implementation
		common.NewPropDef("implementation", typeDefs[86], "", false, nil),
		// with type canonical
		common.NewPropDef("implementationGuide", typeDefs[673], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type canonical
		common.NewPropDef("imports", typeDefs[673], "", true, nil),
		// with type canonical
		common.NewPropDef("instantiates", typeDefs[673], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type string
		common.NewPropDef("kind", typeDefs[46], "", false, []string{
			"instance",
			"capability",
			"requirements",
		}),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type CapabilityStatement_Messaging
		common.NewPropDef("messaging", typeDefs[89], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type code
		common.NewPropDef("patchFormat", typeDefs[674], "", true, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("purpose", typeDefs[676], "", false, nil),
		// with type CapabilityStatement_Rest
		common.NewPropDef("rest", typeDefs[92], "", true, nil),
		// with type CapabilityStatement_Software
		common.NewPropDef("software", typeDefs[95], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of CapabilityStatement_Document
	typeDefs[84].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type markdown
		common.NewPropDef("documentation", typeDefs[676], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("mode", typeDefs[46], "", false, []string{
			"producer",
			"consumer",
		}),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type canonical
		common.NewPropDef("profile", typeDefs[673], "", false, nil),
	})
	// properties of CapabilityStatement_Endpoint
	typeDefs[85].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type url
		common.NewPropDef("address", typeDefs[680], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Coding
		common.NewPropDef("protocol", typeDefs[9], "", false, nil),
	})
	// properties of CapabilityStatement_Implementation
	typeDefs[86].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("custodian", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type url
		common.NewPropDef("url", typeDefs[680], "", false, nil),
	})
	// properties of CapabilityStatement_Interaction
	typeDefs[87].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("code", typeDefs[46], "", false, []string{
			"read",
			"vread",
			"update",
			"patch",
			"delete",
			"history-instance",
			"history-type",
			"create",
			"search-type",
		}),
		// with type markdown
		common.NewPropDef("documentation", typeDefs[676], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of CapabilityStatement_Interaction1
	typeDefs[88].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("code", typeDefs[46], "", false, []string{
			"transaction",
			"batch",
			"search-system",
			"history-system",
		}),
		// with type markdown
		common.NewPropDef("documentation", typeDefs[676], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of CapabilityStatement_Messaging
	typeDefs[89].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type markdown
		common.NewPropDef("documentation", typeDefs[676], "", false, nil),
		// with type CapabilityStatement_Endpoint
		common.NewPropDef("endpoint", typeDefs[85], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type unsignedInt
		common.NewPropDef("reliableCache", typeDefs[679], "", false, nil),
		// with type CapabilityStatement_SupportedMessage
		common.NewPropDef("supportedMessage", typeDefs[96], "", true, nil),
	})
	// properties of CapabilityStatement_Operation
	typeDefs[90].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type canonical
		common.NewPropDef("definition", typeDefs[673], "", false, nil),
		// with type markdown
		common.NewPropDef("documentation", typeDefs[676], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
	})
	// properties of CapabilityStatement_Resource
	typeDefs[91].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("conditionalCreate", typeDefs[40], "", false, nil),
		// with type string
		common.NewPropDef("conditionalDelete", typeDefs[46], "", false, []string{
			"not-supported",
			"single",
			"multiple",
		}),
		// with type string
		common.NewPropDef("conditionalRead", typeDefs[46], "", false, []string{
			"not-supported",
			"modified-since",
			"not-match",
			"full-support",
		}),
		// with type boolean
		common.NewPropDef("conditionalUpdate", typeDefs[40], "", false, nil),
		// with type markdown
		common.NewPropDef("documentation", typeDefs[676], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CapabilityStatement_Interaction
		common.NewPropDef("interaction", typeDefs[87], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CapabilityStatement_Operation
		common.NewPropDef("operation", typeDefs[90], "", true, nil),
		// with type canonical
		common.NewPropDef("profile", typeDefs[673], "", false, nil),
		// with type boolean
		common.NewPropDef("readHistory", typeDefs[40], "", false, nil),
		// with type string
		common.NewPropDef("referencePolicy", typeDefs[46], "", true, []string{
			"literal",
			"logical",
			"resolves",
			"enforced",
			"local",
		}),
		// with type string
		common.NewPropDef("searchInclude", typeDefs[46], "", true, nil),
		// with type CapabilityStatement_SearchParam
		common.NewPropDef("searchParam", typeDefs[93], "", true, nil),
		// with type string
		common.NewPropDef("searchRevInclude", typeDefs[46], "", true, nil),
		// with type canonical
		common.NewPropDef("supportedProfile", typeDefs[673], "", true, nil),
		// with type code
		common.NewPropDef("type", typeDefs[674], "", false, nil),
		// with type boolean
		common.NewPropDef("updateCreate", typeDefs[40], "", false, nil),
		// with type string
		common.NewPropDef("versioning", typeDefs[46], "", false, []string{
			"no-version",
			"versioned",
			"versioned-update",
		}),
	})
	// properties of CapabilityStatement_Rest
	typeDefs[92].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type canonical
		common.NewPropDef("compartment", typeDefs[673], "", true, nil),
		// with type markdown
		common.NewPropDef("documentation", typeDefs[676], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CapabilityStatement_Interaction1
		common.NewPropDef("interaction", typeDefs[88], "", true, nil),
		// with type string
		common.NewPropDef("mode", typeDefs[46], "", false, []string{
			"client",
			"server",
		}),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CapabilityStatement_Operation
		common.NewPropDef("operation", typeDefs[90], "", true, nil),
		// with type CapabilityStatement_Resource
		common.NewPropDef("resource", typeDefs[91], "", true, nil),
		// with type CapabilityStatement_SearchParam
		common.NewPropDef("searchParam", typeDefs[93], "", true, nil),
		// with type CapabilityStatement_Security
		common.NewPropDef("security", typeDefs[94], "", false, nil),
	})
	// properties of CapabilityStatement_SearchParam
	typeDefs[93].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type canonical
		common.NewPropDef("definition", typeDefs[673], "", false, nil),
		// with type markdown
		common.NewPropDef("documentation", typeDefs[676], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"number",
			"date",
			"string",
			"token",
			"reference",
			"composite",
			"quantity",
			"uri",
			"special",
		}),
	})
	// properties of CapabilityStatement_Security
	typeDefs[94].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("cors", typeDefs[40], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("service", typeDefs[8], "", true, nil),
	})
	// properties of CapabilityStatement_Software
	typeDefs[95].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type dateTime
		common.NewPropDef("releaseDate", typeDefs[42], "", false, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of CapabilityStatement_SupportedMessage
	typeDefs[96].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type canonical
		common.NewPropDef("definition", typeDefs[673], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("mode", typeDefs[46], "", false, []string{
			"sender",
			"receiver",
		}),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of CarePlan
	typeDefs[97].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CarePlan_Activity
		common.NewPropDef("activity", typeDefs[98], "", true, nil),
		// with type Reference
		common.NewPropDef("addresses", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("author", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("basedOn", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("careTeam", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("contributor", typeDefs[32], "", true, nil),
		// with type dateTime
		common.NewPropDef("created", typeDefs[42], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("goal", typeDefs[32], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type canonical
		common.NewPropDef("instantiatesCanonical", typeDefs[673], "", true, nil),
		// with type uri
		common.NewPropDef("instantiatesUri", typeDefs[48], "", true, nil),
		// with type code
		common.NewPropDef("intent", typeDefs[674], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type Reference
		common.NewPropDef("partOf", typeDefs[32], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type Reference
		common.NewPropDef("replaces", typeDefs[32], "", true, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("supportingInfo", typeDefs[32], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
	})
	// properties of CarePlan_Activity
	typeDefs[98].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CarePlan_Detail
		common.NewPropDef("detail", typeDefs[99], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("outcomeCodeableConcept", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("outcomeReference", typeDefs[32], "", true, nil),
		// with type Annotation
		common.NewPropDef("progress", typeDefs[3], "", true, nil),
		// with type Reference
		common.NewPropDef("reference", typeDefs[32], "", false, nil),
	})
	// properties of CarePlan_Detail
	typeDefs[99].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Quantity
		common.NewPropDef("dailyAmount", typeDefs[29], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("doNotPerform", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("goal", typeDefs[32], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type canonical
		common.NewPropDef("instantiatesCanonical", typeDefs[673], "", true, nil),
		// with type uri
		common.NewPropDef("instantiatesUri", typeDefs[48], "", true, nil),
		// with type code
		common.NewPropDef("kind", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("location", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("performer", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("productCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("productReference", typeDefs[32], "", false, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("reasonCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("reasonReference", typeDefs[32], "", true, nil),
		// with type Period
		common.NewPropDef("scheduledPeriod", typeDefs[25], "", false, nil),
		// with type string
		common.NewPropDef("scheduledString", typeDefs[46], "", false, nil),
		// with type Timing
		common.NewPropDef("scheduledTiming", typeDefs[655], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"not-started",
			"scheduled",
			"in-progress",
			"on-hold",
			"completed",
			"cancelled",
			"stopped",
			"unknown",
			"entered-in-error",
		}),
		// with type CodeableConcept
		common.NewPropDef("statusReason", typeDefs[8], "", false, nil),
	})
	// properties of CareTeam
	typeDefs[100].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("managingOrganization", typeDefs[32], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type CareTeam_Participant
		common.NewPropDef("participant", typeDefs[101], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("reasonCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("reasonReference", typeDefs[32], "", true, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"proposed",
			"active",
			"suspended",
			"inactive",
			"entered-in-error",
		}),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type ContactPoint
		common.NewPropDef("telecom", typeDefs[11], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of CareTeam_Participant
	typeDefs[101].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("member", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("onBehalfOf", typeDefs[32], "", false, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("role", typeDefs[8], "", true, nil),
	})
	// properties of CatalogEntry
	typeDefs[102].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("additionalCharacteristic", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("additionalClassification", typeDefs[8], "", true, nil),
		// with type Identifier
		common.NewPropDef("additionalIdentifier", typeDefs[18], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("classification", typeDefs[8], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type dateTime
		common.NewPropDef("lastUpdated", typeDefs[42], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("orderable", typeDefs[40], "", false, nil),
		// with type Reference
		common.NewPropDef("referencedItem", typeDefs[32], "", false, nil),
		// with type CatalogEntry_RelatedEntry
		common.NewPropDef("relatedEntry", typeDefs[103], "", true, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type dateTime
		common.NewPropDef("validTo", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("validityPeriod", typeDefs[25], "", false, nil),
	})
	// properties of CatalogEntry_RelatedEntry
	typeDefs[103].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("item", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("relationtype", typeDefs[46], "", false, []string{
			"triggers",
			"is-replaced-by",
		}),
	})
	// properties of ChargeItem
	typeDefs[104].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("account", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("bodysite", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("context", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("costCenter", typeDefs[32], "", false, nil),
		// with type canonical
		common.NewPropDef("definitionCanonical", typeDefs[673], "", true, nil),
		// with type uri
		common.NewPropDef("definitionUri", typeDefs[48], "", true, nil),
		// with type dateTime
		common.NewPropDef("enteredDate", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("enterer", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("factorOverride", typeDefs[43], "", false, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type dateTime
		common.NewPropDef("occurrenceDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("occurrencePeriod", typeDefs[25], "", false, nil),
		// with type Timing
		common.NewPropDef("occurrenceTiming", typeDefs[655], "", false, nil),
		// with type string
		common.NewPropDef("overrideReason", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("partOf", typeDefs[32], "", true, nil),
		// with type ChargeItem_Performer
		common.NewPropDef("performer", typeDefs[109], "", true, nil),
		// with type Reference
		common.NewPropDef("performingOrganization", typeDefs[32], "", false, nil),
		// with type Money
		common.NewPropDef("priceOverride", typeDefs[21], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("productCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("productReference", typeDefs[32], "", false, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("reason", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("requestingOrganization", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("service", typeDefs[32], "", true, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"planned",
			"billable",
			"not-billable",
			"aborted",
			"billed",
			"entered-in-error",
			"unknown",
		}),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("supportingInformation", typeDefs[32], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of ChargeItemDefinition
	typeDefs[105].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ChargeItemDefinition_Applicability
		common.NewPropDef("applicability", typeDefs[106], "", true, nil),
		// with type date
		common.NewPropDef("approvalDate", typeDefs[41], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type uri
		common.NewPropDef("derivedFromUri", typeDefs[48], "", true, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type Period
		common.NewPropDef("effectivePeriod", typeDefs[25], "", false, nil),
		// with type boolean
		common.NewPropDef("experimental", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type Reference
		common.NewPropDef("instance", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type date
		common.NewPropDef("lastReviewDate", typeDefs[41], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type canonical
		common.NewPropDef("partOf", typeDefs[673], "", true, nil),
		// with type ChargeItemDefinition_PropertyGroup
		common.NewPropDef("propertyGroup", typeDefs[108], "", true, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type canonical
		common.NewPropDef("replaces", typeDefs[673], "", true, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of ChargeItemDefinition_Applicability
	typeDefs[106].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("expression", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("language", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of ChargeItemDefinition_PriceComponent
	typeDefs[107].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Money
		common.NewPropDef("amount", typeDefs[21], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("factor", typeDefs[43], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type code
		common.NewPropDef("type", typeDefs[674], "", false, nil),
	})
	// properties of ChargeItemDefinition_PropertyGroup
	typeDefs[108].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ChargeItemDefinition_Applicability
		common.NewPropDef("applicability", typeDefs[106], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type ChargeItemDefinition_PriceComponent
		common.NewPropDef("priceComponent", typeDefs[107], "", true, nil),
	})
	// properties of ChargeItem_Performer
	typeDefs[109].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("actor", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("function", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Claim
	typeDefs[110].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Claim_Accident
		common.NewPropDef("accident", typeDefs[124], "", false, nil),
		// with type Period
		common.NewPropDef("billablePeriod", typeDefs[25], "", false, nil),
		// with type Claim_CareTeam
		common.NewPropDef("careTeam", typeDefs[125], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type dateTime
		common.NewPropDef("created", typeDefs[42], "", false, nil),
		// with type Claim_Diagnosis
		common.NewPropDef("diagnosis", typeDefs[127], "", true, nil),
		// with type Reference
		common.NewPropDef("enterer", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("facility", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("fundsReserve", typeDefs[8], "", false, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type Claim_Insurance
		common.NewPropDef("insurance", typeDefs[128], "", true, nil),
		// with type Reference
		common.NewPropDef("insurer", typeDefs[32], "", false, nil),
		// with type Claim_Item
		common.NewPropDef("item", typeDefs[129], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("originalPrescription", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("patient", typeDefs[32], "", false, nil),
		// with type Claim_Payee
		common.NewPropDef("payee", typeDefs[130], "", false, nil),
		// with type Reference
		common.NewPropDef("prescription", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("priority", typeDefs[8], "", false, nil),
		// with type Claim_Procedure
		common.NewPropDef("procedure", typeDefs[131], "", true, nil),
		// with type Reference
		common.NewPropDef("provider", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("referral", typeDefs[32], "", false, nil),
		// with type Claim_Related
		common.NewPropDef("related", typeDefs[132], "", true, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("subType", typeDefs[8], "", false, nil),
		// with type Claim_SupportingInfo
		common.NewPropDef("supportingInfo", typeDefs[134], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type Money
		common.NewPropDef("total", typeDefs[21], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("use", typeDefs[46], "", false, []string{
			"claim",
			"preauthorization",
			"predetermination",
		}),
	})
	// properties of ClaimResponse
	typeDefs[111].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ClaimResponse_AddItem
		common.NewPropDef("addItem", typeDefs[112], "", true, nil),
		// with type ClaimResponse_Adjudication
		common.NewPropDef("adjudication", typeDefs[113], "", true, nil),
		// with type Reference
		common.NewPropDef("communicationRequest", typeDefs[32], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type dateTime
		common.NewPropDef("created", typeDefs[42], "", false, nil),
		// with type string
		common.NewPropDef("disposition", typeDefs[46], "", false, nil),
		// with type ClaimResponse_Error
		common.NewPropDef("error", typeDefs[116], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Attachment
		common.NewPropDef("form", typeDefs[4], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("formCode", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("fundsReserve", typeDefs[8], "", false, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type ClaimResponse_Insurance
		common.NewPropDef("insurance", typeDefs[117], "", true, nil),
		// with type Reference
		common.NewPropDef("insurer", typeDefs[32], "", false, nil),
		// with type ClaimResponse_Item
		common.NewPropDef("item", typeDefs[118], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type code
		common.NewPropDef("outcome", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("patient", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("payeeType", typeDefs[8], "", false, nil),
		// with type ClaimResponse_Payment
		common.NewPropDef("payment", typeDefs[119], "", false, nil),
		// with type Period
		common.NewPropDef("preAuthPeriod", typeDefs[25], "", false, nil),
		// with type string
		common.NewPropDef("preAuthRef", typeDefs[46], "", false, nil),
		// with type ClaimResponse_ProcessNote
		common.NewPropDef("processNote", typeDefs[120], "", true, nil),
		// with type Reference
		common.NewPropDef("request", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("requestor", typeDefs[32], "", false, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("subType", typeDefs[8], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type ClaimResponse_Total
		common.NewPropDef("total", typeDefs[123], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type code
		common.NewPropDef("use", typeDefs[674], "", false, nil),
	})
	// properties of ClaimResponse_AddItem
	typeDefs[112].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ClaimResponse_Adjudication
		common.NewPropDef("adjudication", typeDefs[113], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("bodySite", typeDefs[8], "", false, nil),
		// with type ClaimResponse_Detail1
		common.NewPropDef("detail", typeDefs[115], "", true, nil),
		// with type positiveInt
		common.NewPropDef("detailSequence", typeDefs[678], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("factor", typeDefs[43], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type positiveInt
		common.NewPropDef("itemSequence", typeDefs[678], "", true, nil),
		// with type Address
		common.NewPropDef("locationAddress", typeDefs[2], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("locationCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("locationReference", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("modifier", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Money
		common.NewPropDef("net", typeDefs[21], "", false, nil),
		// with type positiveInt
		common.NewPropDef("noteNumber", typeDefs[678], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("productOrService", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("programCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("provider", typeDefs[32], "", true, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type date
		common.NewPropDef("servicedDate", typeDefs[41], "", false, nil),
		// with type Period
		common.NewPropDef("servicedPeriod", typeDefs[25], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("subSite", typeDefs[8], "", true, nil),
		// with type positiveInt
		common.NewPropDef("subdetailSequence", typeDefs[678], "", true, nil),
		// with type Money
		common.NewPropDef("unitPrice", typeDefs[21], "", false, nil),
	})
	// properties of ClaimResponse_Adjudication
	typeDefs[113].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Money
		common.NewPropDef("amount", typeDefs[21], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("reason", typeDefs[8], "", false, nil),
		// with type decimal
		common.NewPropDef("value", typeDefs[43], "", false, nil),
	})
	// properties of ClaimResponse_Detail
	typeDefs[114].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ClaimResponse_Adjudication
		common.NewPropDef("adjudication", typeDefs[113], "", true, nil),
		// with type positiveInt
		common.NewPropDef("detailSequence", typeDefs[678], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type positiveInt
		common.NewPropDef("noteNumber", typeDefs[678], "", true, nil),
		// with type ClaimResponse_SubDetail
		common.NewPropDef("subDetail", typeDefs[121], "", true, nil),
	})
	// properties of ClaimResponse_Detail1
	typeDefs[115].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ClaimResponse_Adjudication
		common.NewPropDef("adjudication", typeDefs[113], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("factor", typeDefs[43], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("modifier", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Money
		common.NewPropDef("net", typeDefs[21], "", false, nil),
		// with type positiveInt
		common.NewPropDef("noteNumber", typeDefs[678], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("productOrService", typeDefs[8], "", false, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type ClaimResponse_SubDetail1
		common.NewPropDef("subDetail", typeDefs[122], "", true, nil),
		// with type Money
		common.NewPropDef("unitPrice", typeDefs[21], "", false, nil),
	})
	// properties of ClaimResponse_Error
	typeDefs[116].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type positiveInt
		common.NewPropDef("detailSequence", typeDefs[678], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type positiveInt
		common.NewPropDef("itemSequence", typeDefs[678], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type positiveInt
		common.NewPropDef("subDetailSequence", typeDefs[678], "", false, nil),
	})
	// properties of ClaimResponse_Insurance
	typeDefs[117].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("businessArrangement", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("claimResponse", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("coverage", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("focal", typeDefs[40], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type positiveInt
		common.NewPropDef("sequence", typeDefs[678], "", false, nil),
	})
	// properties of ClaimResponse_Item
	typeDefs[118].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ClaimResponse_Adjudication
		common.NewPropDef("adjudication", typeDefs[113], "", true, nil),
		// with type ClaimResponse_Detail
		common.NewPropDef("detail", typeDefs[114], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type positiveInt
		common.NewPropDef("itemSequence", typeDefs[678], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type positiveInt
		common.NewPropDef("noteNumber", typeDefs[678], "", true, nil),
	})
	// properties of ClaimResponse_Payment
	typeDefs[119].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Money
		common.NewPropDef("adjustment", typeDefs[21], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("adjustmentReason", typeDefs[8], "", false, nil),
		// with type Money
		common.NewPropDef("amount", typeDefs[21], "", false, nil),
		// with type date
		common.NewPropDef("date", typeDefs[41], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of ClaimResponse_ProcessNote
	typeDefs[120].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("language", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type positiveInt
		common.NewPropDef("number", typeDefs[678], "", false, nil),
		// with type string
		common.NewPropDef("text", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"display",
			"print",
			"printoper",
		}),
	})
	// properties of ClaimResponse_SubDetail
	typeDefs[121].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ClaimResponse_Adjudication
		common.NewPropDef("adjudication", typeDefs[113], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type positiveInt
		common.NewPropDef("noteNumber", typeDefs[678], "", true, nil),
		// with type positiveInt
		common.NewPropDef("subDetailSequence", typeDefs[678], "", false, nil),
	})
	// properties of ClaimResponse_SubDetail1
	typeDefs[122].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ClaimResponse_Adjudication
		common.NewPropDef("adjudication", typeDefs[113], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("factor", typeDefs[43], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("modifier", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Money
		common.NewPropDef("net", typeDefs[21], "", false, nil),
		// with type positiveInt
		common.NewPropDef("noteNumber", typeDefs[678], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("productOrService", typeDefs[8], "", false, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type Money
		common.NewPropDef("unitPrice", typeDefs[21], "", false, nil),
	})
	// properties of ClaimResponse_Total
	typeDefs[123].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Money
		common.NewPropDef("amount", typeDefs[21], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Claim_Accident
	typeDefs[124].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type date
		common.NewPropDef("date", typeDefs[41], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Address
		common.NewPropDef("locationAddress", typeDefs[2], "", false, nil),
		// with type Reference
		common.NewPropDef("locationReference", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of Claim_CareTeam
	typeDefs[125].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("provider", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("qualification", typeDefs[8], "", false, nil),
		// with type boolean
		common.NewPropDef("responsible", typeDefs[40], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("role", typeDefs[8], "", false, nil),
		// with type positiveInt
		common.NewPropDef("sequence", typeDefs[678], "", false, nil),
	})
	// properties of Claim_Detail
	typeDefs[126].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("factor", typeDefs[43], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("modifier", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Money
		common.NewPropDef("net", typeDefs[21], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("productOrService", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("programCode", typeDefs[8], "", true, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("revenue", typeDefs[8], "", false, nil),
		// with type positiveInt
		common.NewPropDef("sequence", typeDefs[678], "", false, nil),
		// with type Claim_SubDetail
		common.NewPropDef("subDetail", typeDefs[133], "", true, nil),
		// with type Reference
		common.NewPropDef("udi", typeDefs[32], "", true, nil),
		// with type Money
		common.NewPropDef("unitPrice", typeDefs[21], "", false, nil),
	})
	// properties of Claim_Diagnosis
	typeDefs[127].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("diagnosisCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("diagnosisReference", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("onAdmission", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("packageCode", typeDefs[8], "", false, nil),
		// with type positiveInt
		common.NewPropDef("sequence", typeDefs[678], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", true, nil),
	})
	// properties of Claim_Insurance
	typeDefs[128].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("businessArrangement", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("claimResponse", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("coverage", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("focal", typeDefs[40], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("preAuthRef", typeDefs[46], "", true, nil),
		// with type positiveInt
		common.NewPropDef("sequence", typeDefs[678], "", false, nil),
	})
	// properties of Claim_Item
	typeDefs[129].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("bodySite", typeDefs[8], "", false, nil),
		// with type positiveInt
		common.NewPropDef("careTeamSequence", typeDefs[678], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", false, nil),
		// with type Claim_Detail
		common.NewPropDef("detail", typeDefs[126], "", true, nil),
		// with type positiveInt
		common.NewPropDef("diagnosisSequence", typeDefs[678], "", true, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("factor", typeDefs[43], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type positiveInt
		common.NewPropDef("informationSequence", typeDefs[678], "", true, nil),
		// with type Address
		common.NewPropDef("locationAddress", typeDefs[2], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("locationCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("locationReference", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("modifier", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Money
		common.NewPropDef("net", typeDefs[21], "", false, nil),
		// with type positiveInt
		common.NewPropDef("procedureSequence", typeDefs[678], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("productOrService", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("programCode", typeDefs[8], "", true, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("revenue", typeDefs[8], "", false, nil),
		// with type positiveInt
		common.NewPropDef("sequence", typeDefs[678], "", false, nil),
		// with type date
		common.NewPropDef("servicedDate", typeDefs[41], "", false, nil),
		// with type Period
		common.NewPropDef("servicedPeriod", typeDefs[25], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("subSite", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("udi", typeDefs[32], "", true, nil),
		// with type Money
		common.NewPropDef("unitPrice", typeDefs[21], "", false, nil),
	})
	// properties of Claim_Payee
	typeDefs[130].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("party", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of Claim_Procedure
	typeDefs[131].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("procedureCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("procedureReference", typeDefs[32], "", false, nil),
		// with type positiveInt
		common.NewPropDef("sequence", typeDefs[678], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("udi", typeDefs[32], "", true, nil),
	})
	// properties of Claim_Related
	typeDefs[132].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("claim", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Identifier
		common.NewPropDef("reference", typeDefs[18], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("relationship", typeDefs[8], "", false, nil),
	})
	// properties of Claim_SubDetail
	typeDefs[133].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("factor", typeDefs[43], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("modifier", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Money
		common.NewPropDef("net", typeDefs[21], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("productOrService", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("programCode", typeDefs[8], "", true, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("revenue", typeDefs[8], "", false, nil),
		// with type positiveInt
		common.NewPropDef("sequence", typeDefs[678], "", false, nil),
		// with type Reference
		common.NewPropDef("udi", typeDefs[32], "", true, nil),
		// with type Money
		common.NewPropDef("unitPrice", typeDefs[21], "", false, nil),
	})
	// properties of Claim_SupportingInfo
	typeDefs[134].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("reason", typeDefs[8], "", false, nil),
		// with type positiveInt
		common.NewPropDef("sequence", typeDefs[678], "", false, nil),
		// with type date
		common.NewPropDef("timingDate", typeDefs[41], "", false, nil),
		// with type Period
		common.NewPropDef("timingPeriod", typeDefs[25], "", false, nil),
		// with type Attachment
		common.NewPropDef("valueAttachment", typeDefs[4], "", false, nil),
		// with type boolean
		common.NewPropDef("valueBoolean", typeDefs[40], "", false, nil),
		// with type Quantity
		common.NewPropDef("valueQuantity", typeDefs[29], "", false, nil),
		// with type Reference
		common.NewPropDef("valueReference", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("valueString", typeDefs[46], "", false, nil),
	})
	// properties of ClinicalImpression
	typeDefs[135].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("assessor", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type dateTime
		common.NewPropDef("effectiveDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("effectivePeriod", typeDefs[25], "", false, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type ClinicalImpression_Finding
		common.NewPropDef("finding", typeDefs[136], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type ClinicalImpression_Investigation
		common.NewPropDef("investigation", typeDefs[137], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type Reference
		common.NewPropDef("previous", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("problem", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("prognosisCodeableConcept", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("prognosisReference", typeDefs[32], "", true, nil),
		// with type uri
		common.NewPropDef("protocol", typeDefs[48], "", true, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("statusReason", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("summary", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("supportingInfo", typeDefs[32], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of ClinicalImpression_Finding
	typeDefs[136].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("basis", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("itemCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("itemReference", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of ClinicalImpression_Investigation
	typeDefs[137].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("item", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of CodeSystem
	typeDefs[138].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("caseSensitive", typeDefs[40], "", false, nil),
		// with type boolean
		common.NewPropDef("compositional", typeDefs[40], "", false, nil),
		// with type CodeSystem_Concept
		common.NewPropDef("concept", typeDefs[139], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type string
		common.NewPropDef("content", typeDefs[46], "", false, []string{
			"not-present",
			"example",
			"fragment",
			"complete",
			"supplement",
		}),
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type unsignedInt
		common.NewPropDef("count", typeDefs[679], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type boolean
		common.NewPropDef("experimental", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeSystem_Filter
		common.NewPropDef("filter", typeDefs[141], "", true, nil),
		// with type string
		common.NewPropDef("hierarchyMeaning", typeDefs[46], "", false, []string{
			"grouped-by",
			"is-a",
			"part-of",
			"classified-with",
		}),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type CodeSystem_Property
		common.NewPropDef("property", typeDefs[142], "", true, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("purpose", typeDefs[676], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type canonical
		common.NewPropDef("supplements", typeDefs[673], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type canonical
		common.NewPropDef("valueSet", typeDefs[673], "", false, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("versionNeeded", typeDefs[40], "", false, nil),
	})
	// properties of CodeSystem_Concept
	typeDefs[139].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("code", typeDefs[674], "", false, nil),
		// with type CodeSystem_Concept
		common.NewPropDef("concept", typeDefs[139], "", true, nil),
		// with type string
		common.NewPropDef("definition", typeDefs[46], "", false, nil),
		// with type CodeSystem_Designation
		common.NewPropDef("designation", typeDefs[140], "", true, nil),
		// with type string
		common.NewPropDef("display", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeSystem_Property1
		common.NewPropDef("property", typeDefs[143], "", true, nil),
	})
	// properties of CodeSystem_Designation
	typeDefs[140].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Coding
		common.NewPropDef("use", typeDefs[9], "", false, nil),
		// with type string
		common.NewPropDef("value", typeDefs[46], "", false, nil),
	})
	// properties of CodeSystem_Filter
	typeDefs[141].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("code", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type code
		common.NewPropDef("operator", typeDefs[674], "", true, nil),
		// with type string
		common.NewPropDef("value", typeDefs[46], "", false, nil),
	})
	// properties of CodeSystem_Property
	typeDefs[142].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("code", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"code",
			"Coding",
			"string",
			"integer",
			"boolean",
			"dateTime",
			"decimal",
		}),
		// with type uri
		common.NewPropDef("uri", typeDefs[48], "", false, nil),
	})
	// properties of CodeSystem_Property1
	typeDefs[143].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("code", typeDefs[674], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("valueBoolean", typeDefs[40], "value", false, nil),
		// with type code
		common.NewPropDef("valueCode", typeDefs[674], "value", false, nil),
		// with type Coding
		common.NewPropDef("valueCoding", typeDefs[9], "value", false, nil),
		// with type dateTime
		common.NewPropDef("valueDateTime", typeDefs[42], "value", false, nil),
		// with type decimal
		common.NewPropDef("valueDecimal", typeDefs[43], "value", false, nil),
		// with type integer
		common.NewPropDef("valueInteger", typeDefs[45], "value", false, nil),
		// with type string
		common.NewPropDef("valueString", typeDefs[46], "value", false, nil),
	})
	// properties of CodeableConcept
	typeDefs[8].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Coding
		common.NewPropDef("coding", typeDefs[9], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("text", typeDefs[46], "", false, nil),
	})
	// properties of Coding
	typeDefs[9].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("code", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("display", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("system", typeDefs[48], "", false, nil),
		// with type boolean
		common.NewPropDef("userSelected", typeDefs[40], "", false, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of Communication
	typeDefs[144].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("about", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("basedOn", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type Reference
		common.NewPropDef("inResponseTo", typeDefs[32], "", true, nil),
		// with type canonical
		common.NewPropDef("instantiatesCanonical", typeDefs[673], "", true, nil),
		// with type uri
		common.NewPropDef("instantiatesUri", typeDefs[48], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("medium", typeDefs[8], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type Reference
		common.NewPropDef("partOf", typeDefs[32], "", true, nil),
		// with type Communication_Payload
		common.NewPropDef("payload", typeDefs[147], "", true, nil),
		// with type code
		common.NewPropDef("priority", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("reasonCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("reasonReference", typeDefs[32], "", true, nil),
		// with type dateTime
		common.NewPropDef("received", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("recipient", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("sender", typeDefs[32], "", false, nil),
		// with type dateTime
		common.NewPropDef("sent", typeDefs[42], "", false, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("statusReason", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("topic", typeDefs[8], "", false, nil),
	})
	// properties of CommunicationRequest
	typeDefs[145].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("about", typeDefs[32], "", true, nil),
		// with type dateTime
		common.NewPropDef("authoredOn", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("basedOn", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type boolean
		common.NewPropDef("doNotPerform", typeDefs[40], "", false, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Identifier
		common.NewPropDef("groupIdentifier", typeDefs[18], "", false, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("medium", typeDefs[8], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type dateTime
		common.NewPropDef("occurrenceDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("occurrencePeriod", typeDefs[25], "", false, nil),
		// with type CommunicationRequest_Payload
		common.NewPropDef("payload", typeDefs[146], "", true, nil),
		// with type code
		common.NewPropDef("priority", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("reasonCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("reasonReference", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("recipient", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("replaces", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("requester", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("sender", typeDefs[32], "", false, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("statusReason", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of CommunicationRequest_Payload
	typeDefs[146].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Attachment
		common.NewPropDef("contentAttachment", typeDefs[4], "", false, nil),
		// with type Reference
		common.NewPropDef("contentReference", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("contentString", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Communication_Payload
	typeDefs[147].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Attachment
		common.NewPropDef("contentAttachment", typeDefs[4], "", false, nil),
		// with type Reference
		common.NewPropDef("contentReference", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("contentString", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of CompartmentDefinition
	typeDefs[148].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("code", typeDefs[46], "", false, []string{
			"Patient",
			"Encounter",
			"RelatedPerson",
			"Practitioner",
			"Device",
		}),
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type boolean
		common.NewPropDef("experimental", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("purpose", typeDefs[676], "", false, nil),
		// with type CompartmentDefinition_Resource
		common.NewPropDef("resource", typeDefs[149], "", true, nil),
		// with type boolean
		common.NewPropDef("search", typeDefs[40], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of CompartmentDefinition_Resource
	typeDefs[149].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("code", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("documentation", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("param", typeDefs[46], "", true, nil),
	})
	// properties of Composition
	typeDefs[150].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Composition_Attester
		common.NewPropDef("attester", typeDefs[151], "", true, nil),
		// with type Reference
		common.NewPropDef("author", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("confidentiality", typeDefs[674], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("custodian", typeDefs[32], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type Composition_Event
		common.NewPropDef("event", typeDefs[152], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Composition_RelatesTo
		common.NewPropDef("relatesTo", typeDefs[153], "", true, nil),
		// with type Composition_Section
		common.NewPropDef("section", typeDefs[154], "", true, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"preliminary",
			"final",
			"amended",
			"entered-in-error",
		}),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of Composition_Attester
	typeDefs[151].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("mode", typeDefs[46], "", false, []string{
			"personal",
			"professional",
			"legal",
			"official",
		}),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("party", typeDefs[32], "", false, nil),
		// with type dateTime
		common.NewPropDef("time", typeDefs[42], "", false, nil),
	})
	// properties of Composition_Event
	typeDefs[152].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("detail", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
	})
	// properties of Composition_RelatesTo
	typeDefs[153].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("code", typeDefs[674], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Identifier
		common.NewPropDef("targetIdentifier", typeDefs[18], "", false, nil),
		// with type Reference
		common.NewPropDef("targetReference", typeDefs[32], "", false, nil),
	})
	// properties of Composition_Section
	typeDefs[154].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("author", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("emptyReason", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("entry", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("focus", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type code
		common.NewPropDef("mode", typeDefs[674], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("orderedBy", typeDefs[8], "", false, nil),
		// with type Composition_Section
		common.NewPropDef("section", typeDefs[154], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
	})
	// properties of ConceptMap
	typeDefs[155].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type boolean
		common.NewPropDef("experimental", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type ConceptMap_Group
		common.NewPropDef("group", typeDefs[158], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("purpose", typeDefs[676], "", false, nil),
		// with type canonical
		common.NewPropDef("sourceCanonical", typeDefs[673], "", false, nil),
		// with type uri
		common.NewPropDef("sourceUri", typeDefs[48], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type canonical
		common.NewPropDef("targetCanonical", typeDefs[673], "", false, nil),
		// with type uri
		common.NewPropDef("targetUri", typeDefs[48], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of ConceptMap_DependsOn
	typeDefs[156].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("display", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type uri
		common.NewPropDef("property", typeDefs[48], "", false, nil),
		// with type canonical
		common.NewPropDef("system", typeDefs[673], "", false, nil),
		// with type string
		common.NewPropDef("value", typeDefs[46], "", false, nil),
	})
	// properties of ConceptMap_Element
	typeDefs[157].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("code", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("display", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type ConceptMap_Target
		common.NewPropDef("target", typeDefs[159], "", true, nil),
	})
	// properties of ConceptMap_Group
	typeDefs[158].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ConceptMap_Element
		common.NewPropDef("element", typeDefs[157], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type uri
		common.NewPropDef("source", typeDefs[48], "", false, nil),
		// with type string
		common.NewPropDef("sourceVersion", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("target", typeDefs[48], "", false, nil),
		// with type string
		common.NewPropDef("targetVersion", typeDefs[46], "", false, nil),
		// with type ConceptMap_Unmapped
		common.NewPropDef("unmapped", typeDefs[160], "", false, nil),
	})
	// properties of ConceptMap_Target
	typeDefs[159].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("code", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("comment", typeDefs[46], "", false, nil),
		// with type ConceptMap_DependsOn
		common.NewPropDef("dependsOn", typeDefs[156], "", true, nil),
		// with type string
		common.NewPropDef("display", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("equivalence", typeDefs[46], "", false, []string{
			"relatedto",
			"equivalent",
			"equal",
			"wider",
			"subsumes",
			"narrower",
			"specializes",
			"inexact",
			"unmatched",
			"disjoint",
		}),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type ConceptMap_DependsOn
		common.NewPropDef("product", typeDefs[156], "", true, nil),
	})
	// properties of ConceptMap_Unmapped
	typeDefs[160].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("code", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("display", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("mode", typeDefs[46], "", false, []string{
			"provided",
			"fixed",
			"other-map",
		}),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type canonical
		common.NewPropDef("url", typeDefs[673], "", false, nil),
	})
	// properties of Condition
	typeDefs[161].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Age
		common.NewPropDef("abatementAge", typeDefs[59], "", false, nil),
		// with type dateTime
		common.NewPropDef("abatementDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("abatementPeriod", typeDefs[25], "", false, nil),
		// with type Range
		common.NewPropDef("abatementRange", typeDefs[30], "", false, nil),
		// with type string
		common.NewPropDef("abatementString", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("asserter", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("bodySite", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("clinicalStatus", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type Condition_Evidence
		common.NewPropDef("evidence", typeDefs[162], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type Age
		common.NewPropDef("onsetAge", typeDefs[59], "", false, nil),
		// with type dateTime
		common.NewPropDef("onsetDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("onsetPeriod", typeDefs[25], "", false, nil),
		// with type Range
		common.NewPropDef("onsetRange", typeDefs[30], "", false, nil),
		// with type string
		common.NewPropDef("onsetString", typeDefs[46], "", false, nil),
		// with type dateTime
		common.NewPropDef("recordedDate", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("recorder", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("severity", typeDefs[8], "", false, nil),
		// with type Condition_Stage
		common.NewPropDef("stage", typeDefs[163], "", true, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("verificationStatus", typeDefs[8], "", false, nil),
	})
	// properties of Condition_Evidence
	typeDefs[162].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("detail", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Condition_Stage
	typeDefs[163].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("assessment", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("summary", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of Consent
	typeDefs[164].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type dateTime
		common.NewPropDef("dateTime", typeDefs[42], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("organization", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("patient", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("performer", typeDefs[32], "", true, nil),
		// with type Consent_Policy
		common.NewPropDef("policy", typeDefs[167], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("policyRule", typeDefs[8], "", false, nil),
		// with type Consent_Provision
		common.NewPropDef("provision", typeDefs[168], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("scope", typeDefs[8], "", false, nil),
		// with type Attachment
		common.NewPropDef("sourceAttachment", typeDefs[4], "", false, nil),
		// with type Reference
		common.NewPropDef("sourceReference", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"proposed",
			"active",
			"rejected",
			"inactive",
			"entered-in-error",
		}),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type Consent_Verification
		common.NewPropDef("verification", typeDefs[169], "", true, nil),
	})
	// properties of Consent_Actor
	typeDefs[165].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("reference", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("role", typeDefs[8], "", false, nil),
	})
	// properties of Consent_Data
	typeDefs[166].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("meaning", typeDefs[46], "", false, []string{
			"instance",
			"related",
			"dependents",
			"authoredby",
		}),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("reference", typeDefs[32], "", false, nil),
	})
	// properties of Consent_Policy
	typeDefs[167].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type uri
		common.NewPropDef("authority", typeDefs[48], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type uri
		common.NewPropDef("uri", typeDefs[48], "", false, nil),
	})
	// properties of Consent_Provision
	typeDefs[168].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("action", typeDefs[8], "", true, nil),
		// with type Consent_Actor
		common.NewPropDef("actor", typeDefs[165], "", true, nil),
		// with type Coding
		common.NewPropDef("class", typeDefs[9], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", true, nil),
		// with type Consent_Data
		common.NewPropDef("data", typeDefs[166], "", true, nil),
		// with type Period
		common.NewPropDef("dataPeriod", typeDefs[25], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type Consent_Provision
		common.NewPropDef("provision", typeDefs[168], "", true, nil),
		// with type Coding
		common.NewPropDef("purpose", typeDefs[9], "", true, nil),
		// with type Coding
		common.NewPropDef("securityLabel", typeDefs[9], "", true, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"deny",
			"permit",
		}),
	})
	// properties of Consent_Verification
	typeDefs[169].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type dateTime
		common.NewPropDef("verificationDate", typeDefs[42], "", false, nil),
		// with type boolean
		common.NewPropDef("verified", typeDefs[40], "", false, nil),
		// with type Reference
		common.NewPropDef("verifiedWith", typeDefs[32], "", false, nil),
	})
	// properties of ContactDetail
	typeDefs[10].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type ContactPoint
		common.NewPropDef("telecom", typeDefs[11], "", true, nil),
	})
	// properties of ContactPoint
	typeDefs[11].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type positiveInt
		common.NewPropDef("rank", typeDefs[678], "", false, nil),
		// with type string
		common.NewPropDef("system", typeDefs[46], "", false, []string{
			"phone",
			"fax",
			"email",
			"pager",
			"url",
			"sms",
			"other",
		}),
		// with type string
		common.NewPropDef("use", typeDefs[46], "", false, []string{
			"home",
			"work",
			"temp",
			"old",
			"mobile",
		}),
		// with type string
		common.NewPropDef("value", typeDefs[46], "", false, nil),
	})
	// properties of Contract
	typeDefs[170].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("alias", typeDefs[46], "", true, nil),
		// with type Period
		common.NewPropDef("applies", typeDefs[25], "", false, nil),
		// with type Reference
		common.NewPropDef("author", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("authority", typeDefs[32], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Contract_ContentDefinition
		common.NewPropDef("contentDefinition", typeDefs[174], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("contentDerivative", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("domain", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("expirationType", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Contract_Friendly
		common.NewPropDef("friendly", typeDefs[176], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type Reference
		common.NewPropDef("instantiatesCanonical", typeDefs[32], "", false, nil),
		// with type uri
		common.NewPropDef("instantiatesUri", typeDefs[48], "", false, nil),
		// with type dateTime
		common.NewPropDef("issued", typeDefs[42], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Contract_Legal
		common.NewPropDef("legal", typeDefs[177], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("legalState", typeDefs[8], "", false, nil),
		// with type Attachment
		common.NewPropDef("legallyBindingAttachment", typeDefs[4], "", false, nil),
		// with type Reference
		common.NewPropDef("legallyBindingReference", typeDefs[32], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("relevantHistory", typeDefs[32], "", true, nil),
		// with type Contract_Rule
		common.NewPropDef("rule", typeDefs[180], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("scope", typeDefs[8], "", false, nil),
		// with type Contract_Signer
		common.NewPropDef("signer", typeDefs[182], "", true, nil),
		// with type Reference
		common.NewPropDef("site", typeDefs[32], "", true, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("subType", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", true, nil),
		// with type string
		common.NewPropDef("subtitle", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("supportingInfo", typeDefs[32], "", true, nil),
		// with type Contract_Term
		common.NewPropDef("term", typeDefs[184], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("topicCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("topicReference", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of Contract_Action
	typeDefs[171].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("context", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("contextLinkId", typeDefs[46], "", true, nil),
		// with type boolean
		common.NewPropDef("doNotPerform", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("intent", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("linkId", typeDefs[46], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type dateTime
		common.NewPropDef("occurrenceDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("occurrencePeriod", typeDefs[25], "", false, nil),
		// with type Timing
		common.NewPropDef("occurrenceTiming", typeDefs[655], "", false, nil),
		// with type Reference
		common.NewPropDef("performer", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("performerLinkId", typeDefs[46], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("performerRole", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("performerType", typeDefs[8], "", true, nil),
		// with type string
		common.NewPropDef("reason", typeDefs[46], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("reasonCode", typeDefs[8], "", true, nil),
		// with type string
		common.NewPropDef("reasonLinkId", typeDefs[46], "", true, nil),
		// with type Reference
		common.NewPropDef("reasonReference", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("requester", typeDefs[32], "", true, nil),
		// with type string
		common.NewPropDef("requesterLinkId", typeDefs[46], "", true, nil),
		// with type unsignedInt
		common.NewPropDef("securityLabelNumber", typeDefs[679], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("status", typeDefs[8], "", false, nil),
		// with type Contract_Subject
		common.NewPropDef("subject", typeDefs[183], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of Contract_Answer
	typeDefs[172].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Attachment
		common.NewPropDef("valueAttachment", typeDefs[4], "value", false, nil),
		// with type boolean
		common.NewPropDef("valueBoolean", typeDefs[40], "value", false, nil),
		// with type Coding
		common.NewPropDef("valueCoding", typeDefs[9], "value", false, nil),
		// with type date
		common.NewPropDef("valueDate", typeDefs[41], "value", false, nil),
		// with type dateTime
		common.NewPropDef("valueDateTime", typeDefs[42], "value", false, nil),
		// with type decimal
		common.NewPropDef("valueDecimal", typeDefs[43], "value", false, nil),
		// with type integer
		common.NewPropDef("valueInteger", typeDefs[45], "value", false, nil),
		// with type Quantity
		common.NewPropDef("valueQuantity", typeDefs[29], "value", false, nil),
		// with type Reference
		common.NewPropDef("valueReference", typeDefs[32], "value", false, nil),
		// with type string
		common.NewPropDef("valueString", typeDefs[46], "value", false, nil),
		// with type time
		common.NewPropDef("valueTime", typeDefs[47], "value", false, nil),
		// with type uri
		common.NewPropDef("valueUri", typeDefs[48], "value", false, nil),
	})
	// properties of Contract_Asset
	typeDefs[173].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Contract_Answer
		common.NewPropDef("answer", typeDefs[172], "", true, nil),
		// with type string
		common.NewPropDef("condition", typeDefs[46], "", false, nil),
		// with type Contract_Context
		common.NewPropDef("context", typeDefs[175], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("linkId", typeDefs[46], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("periodType", typeDefs[8], "", true, nil),
		// with type Coding
		common.NewPropDef("relationship", typeDefs[9], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("scope", typeDefs[8], "", false, nil),
		// with type unsignedInt
		common.NewPropDef("securityLabelNumber", typeDefs[679], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("subtype", typeDefs[8], "", true, nil),
		// with type string
		common.NewPropDef("text", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("typeReference", typeDefs[32], "", true, nil),
		// with type Period
		common.NewPropDef("usePeriod", typeDefs[25], "", true, nil),
		// with type Contract_ValuedItem
		common.NewPropDef("valuedItem", typeDefs[185], "", true, nil),
	})
	// properties of Contract_ContentDefinition
	typeDefs[174].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type dateTime
		common.NewPropDef("publicationDate", typeDefs[42], "", false, nil),
		// with type code
		common.NewPropDef("publicationStatus", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("publisher", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("subType", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of Contract_Context
	typeDefs[175].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("reference", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("text", typeDefs[46], "", false, nil),
	})
	// properties of Contract_Friendly
	typeDefs[176].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Attachment
		common.NewPropDef("contentAttachment", typeDefs[4], "", false, nil),
		// with type Reference
		common.NewPropDef("contentReference", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Contract_Legal
	typeDefs[177].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Attachment
		common.NewPropDef("contentAttachment", typeDefs[4], "", false, nil),
		// with type Reference
		common.NewPropDef("contentReference", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Contract_Offer
	typeDefs[178].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Contract_Answer
		common.NewPropDef("answer", typeDefs[172], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("decision", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("decisionMode", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type string
		common.NewPropDef("linkId", typeDefs[46], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Contract_Party
		common.NewPropDef("party", typeDefs[179], "", true, nil),
		// with type unsignedInt
		common.NewPropDef("securityLabelNumber", typeDefs[679], "", true, nil),
		// with type string
		common.NewPropDef("text", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("topic", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of Contract_Party
	typeDefs[179].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("reference", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("role", typeDefs[8], "", false, nil),
	})
	// properties of Contract_Rule
	typeDefs[180].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Attachment
		common.NewPropDef("contentAttachment", typeDefs[4], "", false, nil),
		// with type Reference
		common.NewPropDef("contentReference", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Contract_SecurityLabel
	typeDefs[181].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Coding
		common.NewPropDef("category", typeDefs[9], "", true, nil),
		// with type Coding
		common.NewPropDef("classification", typeDefs[9], "", false, nil),
		// with type Coding
		common.NewPropDef("control", typeDefs[9], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type unsignedInt
		common.NewPropDef("number", typeDefs[679], "", true, nil),
	})
	// properties of Contract_Signer
	typeDefs[182].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("party", typeDefs[32], "", false, nil),
		// with type Signature
		common.NewPropDef("signature", typeDefs[35], "", true, nil),
		// with type Coding
		common.NewPropDef("type", typeDefs[9], "", false, nil),
	})
	// properties of Contract_Subject
	typeDefs[183].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("reference", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("role", typeDefs[8], "", false, nil),
	})
	// properties of Contract_Term
	typeDefs[184].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Contract_Action
		common.NewPropDef("action", typeDefs[171], "", true, nil),
		// with type Period
		common.NewPropDef("applies", typeDefs[25], "", false, nil),
		// with type Contract_Asset
		common.NewPropDef("asset", typeDefs[173], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Contract_Term
		common.NewPropDef("group", typeDefs[184], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type dateTime
		common.NewPropDef("issued", typeDefs[42], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Contract_Offer
		common.NewPropDef("offer", typeDefs[178], "", false, nil),
		// with type Contract_SecurityLabel
		common.NewPropDef("securityLabel", typeDefs[181], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("subType", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("text", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("topicCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("topicReference", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of Contract_ValuedItem
	typeDefs[185].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type dateTime
		common.NewPropDef("effectiveTime", typeDefs[42], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("entityCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("entityReference", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("factor", typeDefs[43], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type string
		common.NewPropDef("linkId", typeDefs[46], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Money
		common.NewPropDef("net", typeDefs[21], "", false, nil),
		// with type string
		common.NewPropDef("payment", typeDefs[46], "", false, nil),
		// with type dateTime
		common.NewPropDef("paymentDate", typeDefs[42], "", false, nil),
		// with type decimal
		common.NewPropDef("points", typeDefs[43], "", false, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type Reference
		common.NewPropDef("recipient", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("responsible", typeDefs[32], "", false, nil),
		// with type unsignedInt
		common.NewPropDef("securityLabelNumber", typeDefs[679], "", true, nil),
		// with type Money
		common.NewPropDef("unitPrice", typeDefs[21], "", false, nil),
	})
	// properties of Contributor
	typeDefs[12].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"author",
			"editor",
			"reviewer",
			"endorser",
		}),
	})
	// properties of Count
	typeDefs[186].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("code", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("comparator", typeDefs[46], "", false, []string{
			"<",
			"<=",
			">=",
			">",
		}),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("system", typeDefs[48], "", false, nil),
		// with type string
		common.NewPropDef("unit", typeDefs[46], "", false, nil),
		// with type decimal
		common.NewPropDef("value", typeDefs[43], "", false, nil),
	})
	// properties of Coverage
	typeDefs[187].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("beneficiary", typeDefs[32], "", false, nil),
		// with type Coverage_Class
		common.NewPropDef("class", typeDefs[198], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("contract", typeDefs[32], "", true, nil),
		// with type Coverage_CostToBeneficiary
		common.NewPropDef("costToBeneficiary", typeDefs[199], "", true, nil),
		// with type string
		common.NewPropDef("dependent", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("network", typeDefs[46], "", false, nil),
		// with type positiveInt
		common.NewPropDef("order", typeDefs[678], "", false, nil),
		// with type Reference
		common.NewPropDef("payor", typeDefs[32], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type Reference
		common.NewPropDef("policyHolder", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("relationship", typeDefs[8], "", false, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type boolean
		common.NewPropDef("subrogation", typeDefs[40], "", false, nil),
		// with type Reference
		common.NewPropDef("subscriber", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("subscriberId", typeDefs[46], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of CoverageEligibilityRequest
	typeDefs[188].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type dateTime
		common.NewPropDef("created", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("enterer", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("facility", typeDefs[32], "", false, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CoverageEligibilityRequest_Insurance
		common.NewPropDef("insurance", typeDefs[190], "", true, nil),
		// with type Reference
		common.NewPropDef("insurer", typeDefs[32], "", false, nil),
		// with type CoverageEligibilityRequest_Item
		common.NewPropDef("item", typeDefs[191], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("patient", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("priority", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("provider", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("purpose", typeDefs[46], "", true, []string{
			"auth-requirements",
			"benefits",
			"discovery",
			"validation",
		}),
		// with type date
		common.NewPropDef("servicedDate", typeDefs[41], "", false, nil),
		// with type Period
		common.NewPropDef("servicedPeriod", typeDefs[25], "", false, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type CoverageEligibilityRequest_SupportingInfo
		common.NewPropDef("supportingInfo", typeDefs[192], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of CoverageEligibilityRequest_Diagnosis
	typeDefs[189].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("diagnosisCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("diagnosisReference", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of CoverageEligibilityRequest_Insurance
	typeDefs[190].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("businessArrangement", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("coverage", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("focal", typeDefs[40], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of CoverageEligibilityRequest_Item
	typeDefs[191].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("detail", typeDefs[32], "", true, nil),
		// with type CoverageEligibilityRequest_Diagnosis
		common.NewPropDef("diagnosis", typeDefs[189], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("facility", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("modifier", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("productOrService", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("provider", typeDefs[32], "", false, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type positiveInt
		common.NewPropDef("supportingInfoSequence", typeDefs[678], "", true, nil),
		// with type Money
		common.NewPropDef("unitPrice", typeDefs[21], "", false, nil),
	})
	// properties of CoverageEligibilityRequest_SupportingInfo
	typeDefs[192].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("appliesToAll", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("information", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type positiveInt
		common.NewPropDef("sequence", typeDefs[678], "", false, nil),
	})
	// properties of CoverageEligibilityResponse
	typeDefs[193].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type dateTime
		common.NewPropDef("created", typeDefs[42], "", false, nil),
		// with type string
		common.NewPropDef("disposition", typeDefs[46], "", false, nil),
		// with type CoverageEligibilityResponse_Error
		common.NewPropDef("error", typeDefs[195], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("form", typeDefs[8], "", false, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CoverageEligibilityResponse_Insurance
		common.NewPropDef("insurance", typeDefs[196], "", true, nil),
		// with type Reference
		common.NewPropDef("insurer", typeDefs[32], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("outcome", typeDefs[46], "", false, []string{
			"queued",
			"complete",
			"error",
			"partial",
		}),
		// with type Reference
		common.NewPropDef("patient", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("preAuthRef", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("purpose", typeDefs[46], "", true, []string{
			"auth-requirements",
			"benefits",
			"discovery",
			"validation",
		}),
		// with type Reference
		common.NewPropDef("request", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("requestor", typeDefs[32], "", false, nil),
		// with type date
		common.NewPropDef("servicedDate", typeDefs[41], "", false, nil),
		// with type Period
		common.NewPropDef("servicedPeriod", typeDefs[25], "", false, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of CoverageEligibilityResponse_Benefit
	typeDefs[194].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Money
		common.NewPropDef("allowedMoney", typeDefs[21], "", false, nil),
		// with type string
		common.NewPropDef("allowedString", typeDefs[46], "", false, nil),
		// with type unsignedInt
		common.NewPropDef("allowedUnsignedInt", typeDefs[679], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type Money
		common.NewPropDef("usedMoney", typeDefs[21], "", false, nil),
		// with type string
		common.NewPropDef("usedString", typeDefs[46], "", false, nil),
		// with type unsignedInt
		common.NewPropDef("usedUnsignedInt", typeDefs[679], "", false, nil),
	})
	// properties of CoverageEligibilityResponse_Error
	typeDefs[195].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of CoverageEligibilityResponse_Insurance
	typeDefs[196].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Period
		common.NewPropDef("benefitPeriod", typeDefs[25], "", false, nil),
		// with type Reference
		common.NewPropDef("coverage", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("inforce", typeDefs[40], "", false, nil),
		// with type CoverageEligibilityResponse_Item
		common.NewPropDef("item", typeDefs[197], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of CoverageEligibilityResponse_Item
	typeDefs[197].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("authorizationRequired", typeDefs[40], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("authorizationSupporting", typeDefs[8], "", true, nil),
		// with type uri
		common.NewPropDef("authorizationUrl", typeDefs[48], "", false, nil),
		// with type CoverageEligibilityResponse_Benefit
		common.NewPropDef("benefit", typeDefs[194], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("excluded", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("modifier", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("network", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("productOrService", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("provider", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("term", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("unit", typeDefs[8], "", false, nil),
	})
	// properties of Coverage_Class
	typeDefs[198].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("value", typeDefs[46], "", false, nil),
	})
	// properties of Coverage_CostToBeneficiary
	typeDefs[199].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Coverage_Exception
		common.NewPropDef("exception", typeDefs[200], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type Money
		common.NewPropDef("valueMoney", typeDefs[21], "", false, nil),
		// with type Quantity
		common.NewPropDef("valueQuantity", typeDefs[29], "", false, nil),
	})
	// properties of Coverage_Exception
	typeDefs[200].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of DataRequirement
	typeDefs[13].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type DataRequirement_CodeFilter
		common.NewPropDef("codeFilter", typeDefs[201], "", true, nil),
		// with type DataRequirement_DateFilter
		common.NewPropDef("dateFilter", typeDefs[202], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type positiveInt
		common.NewPropDef("limit", typeDefs[678], "", false, nil),
		// with type string
		common.NewPropDef("mustSupport", typeDefs[46], "", true, nil),
		// with type canonical
		common.NewPropDef("profile", typeDefs[673], "", true, nil),
		// with type DataRequirement_Sort
		common.NewPropDef("sort", typeDefs[203], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("subjectCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("subjectReference", typeDefs[32], "", false, nil),
		// with type code
		common.NewPropDef("type", typeDefs[674], "", false, nil),
	})
	// properties of DataRequirement_CodeFilter
	typeDefs[201].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Coding
		common.NewPropDef("code", typeDefs[9], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("path", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("searchParam", typeDefs[46], "", false, nil),
		// with type canonical
		common.NewPropDef("valueSet", typeDefs[673], "", false, nil),
	})
	// properties of DataRequirement_DateFilter
	typeDefs[202].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("path", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("searchParam", typeDefs[46], "", false, nil),
		// with type dateTime
		common.NewPropDef("valueDateTime", typeDefs[42], "", false, nil),
		// with type Duration
		common.NewPropDef("valueDuration", typeDefs[236], "", false, nil),
		// with type Period
		common.NewPropDef("valuePeriod", typeDefs[25], "", false, nil),
	})
	// properties of DataRequirement_Sort
	typeDefs[203].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("direction", typeDefs[46], "", false, []string{
			"ascending",
			"descending",
		}),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("path", typeDefs[46], "", false, nil),
	})
	// properties of DetectedIssue
	typeDefs[204].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("author", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type string
		common.NewPropDef("detail", typeDefs[46], "", false, nil),
		// with type DetectedIssue_Evidence
		common.NewPropDef("evidence", typeDefs[205], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type dateTime
		common.NewPropDef("identifiedDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("identifiedPeriod", typeDefs[25], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type Reference
		common.NewPropDef("implicated", typeDefs[32], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type DetectedIssue_Mitigation
		common.NewPropDef("mitigation", typeDefs[206], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("patient", typeDefs[32], "", false, nil),
		// with type uri
		common.NewPropDef("reference", typeDefs[48], "", false, nil),
		// with type string
		common.NewPropDef("severity", typeDefs[46], "", false, []string{
			"high",
			"moderate",
			"low",
		}),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of DetectedIssue_Evidence
	typeDefs[205].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("detail", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of DetectedIssue_Mitigation
	typeDefs[206].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("action", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("author", typeDefs[32], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Device
	typeDefs[207].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ContactPoint
		common.NewPropDef("contact", typeDefs[11], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("definition", typeDefs[32], "", false, nil),
		// with type Device_DeviceName
		common.NewPropDef("deviceName", typeDefs[220], "", true, nil),
		// with type string
		common.NewPropDef("distinctIdentifier", typeDefs[46], "", false, nil),
		// with type dateTime
		common.NewPropDef("expirationDate", typeDefs[42], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("location", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("lotNumber", typeDefs[46], "", false, nil),
		// with type dateTime
		common.NewPropDef("manufactureDate", typeDefs[42], "", false, nil),
		// with type string
		common.NewPropDef("manufacturer", typeDefs[46], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type string
		common.NewPropDef("modelNumber", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type Reference
		common.NewPropDef("owner", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("parent", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("partNumber", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("patient", typeDefs[32], "", false, nil),
		// with type Device_Property
		common.NewPropDef("property", typeDefs[221], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("safety", typeDefs[8], "", true, nil),
		// with type string
		common.NewPropDef("serialNumber", typeDefs[46], "", false, nil),
		// with type Device_Specialization
		common.NewPropDef("specialization", typeDefs[222], "", true, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"active",
			"inactive",
			"entered-in-error",
			"unknown",
		}),
		// with type CodeableConcept
		common.NewPropDef("statusReason", typeDefs[8], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type Device_UdiCarrier
		common.NewPropDef("udiCarrier", typeDefs[223], "", true, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type Device_Version
		common.NewPropDef("version", typeDefs[224], "", true, nil),
	})
	// properties of DeviceDefinition
	typeDefs[208].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type DeviceDefinition_Capability
		common.NewPropDef("capability", typeDefs[209], "", true, nil),
		// with type ContactPoint
		common.NewPropDef("contact", typeDefs[11], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type DeviceDefinition_DeviceName
		common.NewPropDef("deviceName", typeDefs[210], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("languageCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("manufacturerReference", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("manufacturerString", typeDefs[46], "", false, nil),
		// with type DeviceDefinition_Material
		common.NewPropDef("material", typeDefs[211], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type string
		common.NewPropDef("modelNumber", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type uri
		common.NewPropDef("onlineInformation", typeDefs[48], "", false, nil),
		// with type Reference
		common.NewPropDef("owner", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("parentDevice", typeDefs[32], "", false, nil),
		// with type ProdCharacteristic
		common.NewPropDef("physicalCharacteristics", typeDefs[27], "", false, nil),
		// with type DeviceDefinition_Property
		common.NewPropDef("property", typeDefs[212], "", true, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("safety", typeDefs[8], "", true, nil),
		// with type ProductShelfLife
		common.NewPropDef("shelfLifeStorage", typeDefs[28], "", true, nil),
		// with type DeviceDefinition_Specialization
		common.NewPropDef("specialization", typeDefs[213], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type DeviceDefinition_UdiDeviceIdentifier
		common.NewPropDef("udiDeviceIdentifier", typeDefs[214], "", true, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", true, nil),
	})
	// properties of DeviceDefinition_Capability
	typeDefs[209].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("description", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of DeviceDefinition_DeviceName
	typeDefs[210].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"udi-label-name",
			"user-friendly-name",
			"patient-reported-name",
			"manufacturer-name",
			"model-name",
			"other",
		}),
	})
	// properties of DeviceDefinition_Material
	typeDefs[211].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("allergenicIndicator", typeDefs[40], "", false, nil),
		// with type boolean
		common.NewPropDef("alternate", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("substance", typeDefs[8], "", false, nil),
	})
	// properties of DeviceDefinition_Property
	typeDefs[212].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("valueCode", typeDefs[8], "", true, nil),
		// with type Quantity
		common.NewPropDef("valueQuantity", typeDefs[29], "", true, nil),
	})
	// properties of DeviceDefinition_Specialization
	typeDefs[213].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("systemType", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of DeviceDefinition_UdiDeviceIdentifier
	typeDefs[214].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("deviceIdentifier", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("issuer", typeDefs[48], "", false, nil),
		// with type uri
		common.NewPropDef("jurisdiction", typeDefs[48], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of DeviceMetric
	typeDefs[215].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type DeviceMetric_Calibration
		common.NewPropDef("calibration", typeDefs[216], "", true, nil),
		// with type string
		common.NewPropDef("category", typeDefs[46], "", false, []string{
			"measurement",
			"setting",
			"calculation",
			"unspecified",
		}),
		// with type string
		common.NewPropDef("color", typeDefs[46], "", false, []string{
			"black",
			"red",
			"green",
			"yellow",
			"blue",
			"magenta",
			"cyan",
			"white",
		}),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Timing
		common.NewPropDef("measurementPeriod", typeDefs[655], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("operationalStatus", typeDefs[46], "", false, []string{
			"on",
			"off",
			"standby",
			"entered-in-error",
		}),
		// with type Reference
		common.NewPropDef("parent", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("source", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("unit", typeDefs[8], "", false, nil),
	})
	// properties of DeviceMetric_Calibration
	typeDefs[216].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("state", typeDefs[46], "", false, []string{
			"not-calibrated",
			"calibration-required",
			"calibrated",
			"unspecified",
		}),
		// with type instant
		common.NewPropDef("time", typeDefs[44], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"unspecified",
			"offset",
			"gain",
			"two-point",
		}),
	})
	// properties of DeviceRequest
	typeDefs[217].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type dateTime
		common.NewPropDef("authoredOn", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("basedOn", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("codeCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("codeReference", typeDefs[32], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Identifier
		common.NewPropDef("groupIdentifier", typeDefs[18], "", false, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type canonical
		common.NewPropDef("instantiatesCanonical", typeDefs[673], "", true, nil),
		// with type uri
		common.NewPropDef("instantiatesUri", typeDefs[48], "", true, nil),
		// with type Reference
		common.NewPropDef("insurance", typeDefs[32], "", true, nil),
		// with type code
		common.NewPropDef("intent", typeDefs[674], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type dateTime
		common.NewPropDef("occurrenceDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("occurrencePeriod", typeDefs[25], "", false, nil),
		// with type Timing
		common.NewPropDef("occurrenceTiming", typeDefs[655], "", false, nil),
		// with type DeviceRequest_Parameter
		common.NewPropDef("parameter", typeDefs[218], "", true, nil),
		// with type Reference
		common.NewPropDef("performer", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("performerType", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("priorRequest", typeDefs[32], "", true, nil),
		// with type code
		common.NewPropDef("priority", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("reasonCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("reasonReference", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("relevantHistory", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("requester", typeDefs[32], "", false, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("supportingInfo", typeDefs[32], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of DeviceRequest_Parameter
	typeDefs[218].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("valueBoolean", typeDefs[40], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("valueCodeableConcept", typeDefs[8], "", false, nil),
		// with type Quantity
		common.NewPropDef("valueQuantity", typeDefs[29], "", false, nil),
		// with type Range
		common.NewPropDef("valueRange", typeDefs[30], "", false, nil),
	})
	// properties of DeviceUseStatement
	typeDefs[219].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("basedOn", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("bodySite", typeDefs[8], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("derivedFrom", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("device", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("reasonCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("reasonReference", typeDefs[32], "", true, nil),
		// with type dateTime
		common.NewPropDef("recordedOn", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("source", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"active",
			"completed",
			"entered-in-error",
			"intended",
			"stopped",
			"on-hold",
		}),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type dateTime
		common.NewPropDef("timingDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("timingPeriod", typeDefs[25], "", false, nil),
		// with type Timing
		common.NewPropDef("timingTiming", typeDefs[655], "", false, nil),
	})
	// properties of Device_DeviceName
	typeDefs[220].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"udi-label-name",
			"user-friendly-name",
			"patient-reported-name",
			"manufacturer-name",
			"model-name",
			"other",
		}),
	})
	// properties of Device_Property
	typeDefs[221].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("valueCode", typeDefs[8], "", true, nil),
		// with type Quantity
		common.NewPropDef("valueQuantity", typeDefs[29], "", true, nil),
	})
	// properties of Device_Specialization
	typeDefs[222].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("systemType", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of Device_UdiCarrier
	typeDefs[223].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type base64Binary
		common.NewPropDef("carrierAIDC", typeDefs[39], "", false, nil),
		// with type string
		common.NewPropDef("carrierHRF", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("deviceIdentifier", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("entryType", typeDefs[46], "", false, []string{
			"barcode",
			"rfid",
			"manual",
			"card",
			"self-reported",
			"unknown",
		}),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("issuer", typeDefs[48], "", false, nil),
		// with type uri
		common.NewPropDef("jurisdiction", typeDefs[48], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Device_Version
	typeDefs[224].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Identifier
		common.NewPropDef("component", typeDefs[18], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("value", typeDefs[46], "", false, nil),
	})
	// properties of DiagnosticReport
	typeDefs[225].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("basedOn", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("conclusion", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("conclusionCode", typeDefs[8], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type dateTime
		common.NewPropDef("effectiveDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("effectivePeriod", typeDefs[25], "", false, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type Reference
		common.NewPropDef("imagingStudy", typeDefs[32], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type instant
		common.NewPropDef("issued", typeDefs[44], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type DiagnosticReport_Media
		common.NewPropDef("media", typeDefs[226], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("performer", typeDefs[32], "", true, nil),
		// with type Attachment
		common.NewPropDef("presentedForm", typeDefs[4], "", true, nil),
		// with type Reference
		common.NewPropDef("result", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("resultsInterpreter", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("specimen", typeDefs[32], "", true, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"registered",
			"partial",
			"preliminary",
			"final",
			"amended",
			"corrected",
			"appended",
			"cancelled",
			"entered-in-error",
			"unknown",
		}),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of DiagnosticReport_Media
	typeDefs[226].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("comment", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("link", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Distance
	typeDefs[227].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("code", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("comparator", typeDefs[46], "", false, []string{
			"<",
			"<=",
			">=",
			">",
		}),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("system", typeDefs[48], "", false, nil),
		// with type string
		common.NewPropDef("unit", typeDefs[46], "", false, nil),
		// with type decimal
		common.NewPropDef("value", typeDefs[43], "", false, nil),
	})
	// properties of DocumentManifest
	typeDefs[228].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("author", typeDefs[32], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("content", typeDefs[32], "", true, nil),
		// with type dateTime
		common.NewPropDef("created", typeDefs[42], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Identifier
		common.NewPropDef("masterIdentifier", typeDefs[18], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("recipient", typeDefs[32], "", true, nil),
		// with type DocumentManifest_Related
		common.NewPropDef("related", typeDefs[229], "", true, nil),
		// with type uri
		common.NewPropDef("source", typeDefs[48], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"current",
			"superseded",
			"entered-in-error",
		}),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of DocumentManifest_Related
	typeDefs[229].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("ref", typeDefs[32], "", false, nil),
	})
	// properties of DocumentReference
	typeDefs[230].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("authenticator", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("author", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type DocumentReference_Content
		common.NewPropDef("content", typeDefs[231], "", true, nil),
		// with type DocumentReference_Context
		common.NewPropDef("context", typeDefs[232], "", false, nil),
		// with type Reference
		common.NewPropDef("custodian", typeDefs[32], "", false, nil),
		// with type instant
		common.NewPropDef("date", typeDefs[44], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type code
		common.NewPropDef("docStatus", typeDefs[674], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Identifier
		common.NewPropDef("masterIdentifier", typeDefs[18], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type DocumentReference_RelatesTo
		common.NewPropDef("relatesTo", typeDefs[233], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("securityLabel", typeDefs[8], "", true, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"current",
			"superseded",
			"entered-in-error",
		}),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of DocumentReference_Content
	typeDefs[231].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Attachment
		common.NewPropDef("attachment", typeDefs[4], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Coding
		common.NewPropDef("format", typeDefs[9], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of DocumentReference_Context
	typeDefs[232].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("event", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("facilityType", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("practiceSetting", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("related", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("sourcePatientInfo", typeDefs[32], "", false, nil),
	})
	// properties of DocumentReference_RelatesTo
	typeDefs[233].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("code", typeDefs[46], "", false, []string{
			"replaces",
			"transforms",
			"signs",
			"appends",
		}),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("target", typeDefs[32], "", false, nil),
	})
	// properties of DomainResource
	typeDefs[14].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Dosage
	typeDefs[234].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("additionalInstruction", typeDefs[8], "", true, nil),
		// with type boolean
		common.NewPropDef("asNeededBoolean", typeDefs[40], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("asNeededCodeableConcept", typeDefs[8], "", false, nil),
		// with type Dosage_DoseAndRate
		common.NewPropDef("doseAndRate", typeDefs[235], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Quantity
		common.NewPropDef("maxDosePerAdministration", typeDefs[29], "", false, nil),
		// with type Quantity
		common.NewPropDef("maxDosePerLifetime", typeDefs[29], "", false, nil),
		// with type Ratio
		common.NewPropDef("maxDosePerPeriod", typeDefs[31], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("method", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("patientInstruction", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("route", typeDefs[8], "", false, nil),
		// with type integer
		common.NewPropDef("sequence", typeDefs[45], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("site", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("text", typeDefs[46], "", false, nil),
		// with type Timing
		common.NewPropDef("timing", typeDefs[655], "", false, nil),
	})
	// properties of Dosage_DoseAndRate
	typeDefs[235].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Quantity
		common.NewPropDef("doseQuantity", typeDefs[29], "", false, nil),
		// with type Range
		common.NewPropDef("doseRange", typeDefs[30], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Quantity
		common.NewPropDef("rateQuantity", typeDefs[29], "", false, nil),
		// with type Range
		common.NewPropDef("rateRange", typeDefs[30], "", false, nil),
		// with type Ratio
		common.NewPropDef("rateRatio", typeDefs[31], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of Duration
	typeDefs[236].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("code", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("comparator", typeDefs[46], "", false, []string{
			"<",
			"<=",
			">=",
			">",
		}),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("system", typeDefs[48], "", false, nil),
		// with type string
		common.NewPropDef("unit", typeDefs[46], "", false, nil),
		// with type decimal
		common.NewPropDef("value", typeDefs[43], "", false, nil),
	})
	// properties of EffectEvidenceSynthesis
	typeDefs[237].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type date
		common.NewPropDef("approvalDate", typeDefs[41], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("author", typeDefs[10], "", true, nil),
		// with type EffectEvidenceSynthesis_Certainty
		common.NewPropDef("certainty", typeDefs[238], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("editor", typeDefs[10], "", true, nil),
		// with type EffectEvidenceSynthesis_EffectEstimate
		common.NewPropDef("effectEstimate", typeDefs[240], "", true, nil),
		// with type Period
		common.NewPropDef("effectivePeriod", typeDefs[25], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("endorser", typeDefs[10], "", true, nil),
		// with type Reference
		common.NewPropDef("exposure", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("exposureAlternative", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type date
		common.NewPropDef("lastReviewDate", typeDefs[41], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type Reference
		common.NewPropDef("outcome", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("population", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type RelatedArtifact
		common.NewPropDef("relatedArtifact", typeDefs[33], "", true, nil),
		// with type EffectEvidenceSynthesis_ResultsByExposure
		common.NewPropDef("resultsByExposure", typeDefs[242], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("reviewer", typeDefs[10], "", true, nil),
		// with type EffectEvidenceSynthesis_SampleSize
		common.NewPropDef("sampleSize", typeDefs[243], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type CodeableConcept
		common.NewPropDef("studyType", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("synthesisType", typeDefs[8], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("topic", typeDefs[8], "", true, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of EffectEvidenceSynthesis_Certainty
	typeDefs[238].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type EffectEvidenceSynthesis_CertaintySubcomponent
		common.NewPropDef("certaintySubcomponent", typeDefs[239], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("rating", typeDefs[8], "", true, nil),
	})
	// properties of EffectEvidenceSynthesis_CertaintySubcomponent
	typeDefs[239].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("rating", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of EffectEvidenceSynthesis_EffectEstimate
	typeDefs[240].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type EffectEvidenceSynthesis_PrecisionEstimate
		common.NewPropDef("precisionEstimate", typeDefs[241], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("unitOfMeasure", typeDefs[8], "", false, nil),
		// with type decimal
		common.NewPropDef("value", typeDefs[43], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("variantState", typeDefs[8], "", false, nil),
	})
	// properties of EffectEvidenceSynthesis_PrecisionEstimate
	typeDefs[241].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("from", typeDefs[43], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type decimal
		common.NewPropDef("level", typeDefs[43], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("to", typeDefs[43], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of EffectEvidenceSynthesis_ResultsByExposure
	typeDefs[242].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("exposureState", typeDefs[46], "", false, []string{
			"exposure",
			"exposure-alternative",
		}),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("riskEvidenceSynthesis", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("variantState", typeDefs[8], "", false, nil),
	})
	// properties of EffectEvidenceSynthesis_SampleSize
	typeDefs[243].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type integer
		common.NewPropDef("numberOfParticipants", typeDefs[45], "", false, nil),
		// with type integer
		common.NewPropDef("numberOfStudies", typeDefs[45], "", false, nil),
	})
	// properties of Element
	typeDefs[0].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
	})
	// properties of ElementDefinition
	typeDefs[244].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("alias", typeDefs[46], "", true, nil),
		// with type ElementDefinition_Base
		common.NewPropDef("base", typeDefs[245], "", false, nil),
		// with type ElementDefinition_Binding
		common.NewPropDef("binding", typeDefs[246], "", false, nil),
		// with type Coding
		common.NewPropDef("code", typeDefs[9], "", true, nil),
		// with type markdown
		common.NewPropDef("comment", typeDefs[676], "", false, nil),
		// with type id
		common.NewPropDef("condition", typeDefs[675], "", true, nil),
		// with type ElementDefinition_Constraint
		common.NewPropDef("constraint", typeDefs[247], "", true, nil),
		// with type uri
		common.NewPropDef("contentReference", typeDefs[48], "", false, nil),
		// with type Address
		common.NewPropDef("defaultValueAddress", typeDefs[2], "defaultValue", false, nil),
		// with type Age
		common.NewPropDef("defaultValueAge", typeDefs[59], "defaultValue", false, nil),
		// with type Annotation
		common.NewPropDef("defaultValueAnnotation", typeDefs[3], "defaultValue", false, nil),
		// with type Attachment
		common.NewPropDef("defaultValueAttachment", typeDefs[4], "defaultValue", false, nil),
		// with type base64Binary
		common.NewPropDef("defaultValueBase64Binary", typeDefs[39], "defaultValue", false, nil),
		// with type boolean
		common.NewPropDef("defaultValueBoolean", typeDefs[40], "defaultValue", false, nil),
		// with type canonical
		common.NewPropDef("defaultValueCanonical", typeDefs[673], "defaultValue", false, nil),
		// with type code
		common.NewPropDef("defaultValueCode", typeDefs[674], "defaultValue", false, nil),
		// with type CodeableConcept
		common.NewPropDef("defaultValueCodeableConcept", typeDefs[8], "defaultValue", false, nil),
		// with type Coding
		common.NewPropDef("defaultValueCoding", typeDefs[9], "defaultValue", false, nil),
		// with type ContactDetail
		common.NewPropDef("defaultValueContactDetail", typeDefs[10], "defaultValue", false, nil),
		// with type ContactPoint
		common.NewPropDef("defaultValueContactPoint", typeDefs[11], "defaultValue", false, nil),
		// with type Contributor
		common.NewPropDef("defaultValueContributor", typeDefs[12], "defaultValue", false, nil),
		// with type Count
		common.NewPropDef("defaultValueCount", typeDefs[186], "defaultValue", false, nil),
		// with type DataRequirement
		common.NewPropDef("defaultValueDataRequirement", typeDefs[13], "defaultValue", false, nil),
		// with type date
		common.NewPropDef("defaultValueDate", typeDefs[41], "defaultValue", false, nil),
		// with type dateTime
		common.NewPropDef("defaultValueDateTime", typeDefs[42], "defaultValue", false, nil),
		// with type decimal
		common.NewPropDef("defaultValueDecimal", typeDefs[43], "defaultValue", false, nil),
		// with type Distance
		common.NewPropDef("defaultValueDistance", typeDefs[227], "defaultValue", false, nil),
		// with type Dosage
		common.NewPropDef("defaultValueDosage", typeDefs[234], "defaultValue", false, nil),
		// with type Duration
		common.NewPropDef("defaultValueDuration", typeDefs[236], "defaultValue", false, nil),
		// with type Expression
		common.NewPropDef("defaultValueExpression", typeDefs[15], "defaultValue", false, nil),
		// with type HumanName
		common.NewPropDef("defaultValueHumanName", typeDefs[17], "defaultValue", false, nil),
		// with type id
		common.NewPropDef("defaultValueId", typeDefs[675], "defaultValue", false, nil),
		// with type Identifier
		common.NewPropDef("defaultValueIdentifier", typeDefs[18], "defaultValue", false, nil),
		// with type instant
		common.NewPropDef("defaultValueInstant", typeDefs[44], "defaultValue", false, nil),
		// with type integer
		common.NewPropDef("defaultValueInteger", typeDefs[45], "defaultValue", false, nil),
		// with type markdown
		common.NewPropDef("defaultValueMarkdown", typeDefs[676], "defaultValue", false, nil),
		// with type Meta
		common.NewPropDef("defaultValueMeta", typeDefs[20], "defaultValue", false, nil),
		// with type Money
		common.NewPropDef("defaultValueMoney", typeDefs[21], "defaultValue", false, nil),
		// with type oid
		common.NewPropDef("defaultValueOid", typeDefs[677], "defaultValue", false, nil),
		// with type ParameterDefinition
		common.NewPropDef("defaultValueParameterDefinition", typeDefs[23], "defaultValue", false, nil),
		// with type Period
		common.NewPropDef("defaultValuePeriod", typeDefs[25], "defaultValue", false, nil),
		// with type positiveInt
		common.NewPropDef("defaultValuePositiveInt", typeDefs[678], "defaultValue", false, nil),
		// with type Quantity
		common.NewPropDef("defaultValueQuantity", typeDefs[29], "defaultValue", false, nil),
		// with type Range
		common.NewPropDef("defaultValueRange", typeDefs[30], "defaultValue", false, nil),
		// with type Ratio
		common.NewPropDef("defaultValueRatio", typeDefs[31], "defaultValue", false, nil),
		// with type Reference
		common.NewPropDef("defaultValueReference", typeDefs[32], "defaultValue", false, nil),
		// with type RelatedArtifact
		common.NewPropDef("defaultValueRelatedArtifact", typeDefs[33], "defaultValue", false, nil),
		// with type SampledData
		common.NewPropDef("defaultValueSampledData", typeDefs[34], "defaultValue", false, nil),
		// with type Signature
		common.NewPropDef("defaultValueSignature", typeDefs[35], "defaultValue", false, nil),
		// with type string
		common.NewPropDef("defaultValueString", typeDefs[46], "defaultValue", false, nil),
		// with type time
		common.NewPropDef("defaultValueTime", typeDefs[47], "defaultValue", false, nil),
		// with type Timing
		common.NewPropDef("defaultValueTiming", typeDefs[655], "defaultValue", false, nil),
		// with type TriggerDefinition
		common.NewPropDef("defaultValueTriggerDefinition", typeDefs[37], "defaultValue", false, nil),
		// with type unsignedInt
		common.NewPropDef("defaultValueUnsignedInt", typeDefs[679], "defaultValue", false, nil),
		// with type uri
		common.NewPropDef("defaultValueUri", typeDefs[48], "defaultValue", false, nil),
		// with type url
		common.NewPropDef("defaultValueUrl", typeDefs[680], "defaultValue", false, nil),
		// with type UsageContext
		common.NewPropDef("defaultValueUsageContext", typeDefs[38], "defaultValue", false, nil),
		// with type uuid
		common.NewPropDef("defaultValueUuid", typeDefs[681], "defaultValue", false, nil),
		// with type markdown
		common.NewPropDef("definition", typeDefs[676], "", false, nil),
		// with type ElementDefinition_Example
		common.NewPropDef("example", typeDefs[249], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Address
		common.NewPropDef("fixedAddress", typeDefs[2], "fixed", false, nil),
		// with type Age
		common.NewPropDef("fixedAge", typeDefs[59], "fixed", false, nil),
		// with type Annotation
		common.NewPropDef("fixedAnnotation", typeDefs[3], "fixed", false, nil),
		// with type Attachment
		common.NewPropDef("fixedAttachment", typeDefs[4], "fixed", false, nil),
		// with type base64Binary
		common.NewPropDef("fixedBase64Binary", typeDefs[39], "fixed", false, nil),
		// with type boolean
		common.NewPropDef("fixedBoolean", typeDefs[40], "fixed", false, nil),
		// with type canonical
		common.NewPropDef("fixedCanonical", typeDefs[673], "fixed", false, nil),
		// with type code
		common.NewPropDef("fixedCode", typeDefs[674], "fixed", false, nil),
		// with type CodeableConcept
		common.NewPropDef("fixedCodeableConcept", typeDefs[8], "fixed", false, nil),
		// with type Coding
		common.NewPropDef("fixedCoding", typeDefs[9], "fixed", false, nil),
		// with type ContactDetail
		common.NewPropDef("fixedContactDetail", typeDefs[10], "fixed", false, nil),
		// with type ContactPoint
		common.NewPropDef("fixedContactPoint", typeDefs[11], "fixed", false, nil),
		// with type Contributor
		common.NewPropDef("fixedContributor", typeDefs[12], "fixed", false, nil),
		// with type Count
		common.NewPropDef("fixedCount", typeDefs[186], "fixed", false, nil),
		// with type DataRequirement
		common.NewPropDef("fixedDataRequirement", typeDefs[13], "fixed", false, nil),
		// with type date
		common.NewPropDef("fixedDate", typeDefs[41], "fixed", false, nil),
		// with type dateTime
		common.NewPropDef("fixedDateTime", typeDefs[42], "fixed", false, nil),
		// with type decimal
		common.NewPropDef("fixedDecimal", typeDefs[43], "fixed", false, nil),
		// with type Distance
		common.NewPropDef("fixedDistance", typeDefs[227], "fixed", false, nil),
		// with type Dosage
		common.NewPropDef("fixedDosage", typeDefs[234], "fixed", false, nil),
		// with type Duration
		common.NewPropDef("fixedDuration", typeDefs[236], "fixed", false, nil),
		// with type Expression
		common.NewPropDef("fixedExpression", typeDefs[15], "fixed", false, nil),
		// with type HumanName
		common.NewPropDef("fixedHumanName", typeDefs[17], "fixed", false, nil),
		// with type id
		common.NewPropDef("fixedId", typeDefs[675], "fixed", false, nil),
		// with type Identifier
		common.NewPropDef("fixedIdentifier", typeDefs[18], "fixed", false, nil),
		// with type instant
		common.NewPropDef("fixedInstant", typeDefs[44], "fixed", false, nil),
		// with type integer
		common.NewPropDef("fixedInteger", typeDefs[45], "fixed", false, nil),
		// with type markdown
		common.NewPropDef("fixedMarkdown", typeDefs[676], "fixed", false, nil),
		// with type Meta
		common.NewPropDef("fixedMeta", typeDefs[20], "fixed", false, nil),
		// with type Money
		common.NewPropDef("fixedMoney", typeDefs[21], "fixed", false, nil),
		// with type oid
		common.NewPropDef("fixedOid", typeDefs[677], "fixed", false, nil),
		// with type ParameterDefinition
		common.NewPropDef("fixedParameterDefinition", typeDefs[23], "fixed", false, nil),
		// with type Period
		common.NewPropDef("fixedPeriod", typeDefs[25], "fixed", false, nil),
		// with type positiveInt
		common.NewPropDef("fixedPositiveInt", typeDefs[678], "fixed", false, nil),
		// with type Quantity
		common.NewPropDef("fixedQuantity", typeDefs[29], "fixed", false, nil),
		// with type Range
		common.NewPropDef("fixedRange", typeDefs[30], "fixed", false, nil),
		// with type Ratio
		common.NewPropDef("fixedRatio", typeDefs[31], "fixed", false, nil),
		// with type Reference
		common.NewPropDef("fixedReference", typeDefs[32], "fixed", false, nil),
		// with type RelatedArtifact
		common.NewPropDef("fixedRelatedArtifact", typeDefs[33], "fixed", false, nil),
		// with type SampledData
		common.NewPropDef("fixedSampledData", typeDefs[34], "fixed", false, nil),
		// with type Signature
		common.NewPropDef("fixedSignature", typeDefs[35], "fixed", false, nil),
		// with type string
		common.NewPropDef("fixedString", typeDefs[46], "fixed", false, nil),
		// with type time
		common.NewPropDef("fixedTime", typeDefs[47], "fixed", false, nil),
		// with type Timing
		common.NewPropDef("fixedTiming", typeDefs[655], "fixed", false, nil),
		// with type TriggerDefinition
		common.NewPropDef("fixedTriggerDefinition", typeDefs[37], "fixed", false, nil),
		// with type unsignedInt
		common.NewPropDef("fixedUnsignedInt", typeDefs[679], "fixed", false, nil),
		// with type uri
		common.NewPropDef("fixedUri", typeDefs[48], "fixed", false, nil),
		// with type url
		common.NewPropDef("fixedUrl", typeDefs[680], "fixed", false, nil),
		// with type UsageContext
		common.NewPropDef("fixedUsageContext", typeDefs[38], "fixed", false, nil),
		// with type uuid
		common.NewPropDef("fixedUuid", typeDefs[681], "fixed", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("isModifier", typeDefs[40], "", false, nil),
		// with type string
		common.NewPropDef("isModifierReason", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("isSummary", typeDefs[40], "", false, nil),
		// with type string
		common.NewPropDef("label", typeDefs[46], "", false, nil),
		// with type ElementDefinition_Mapping
		common.NewPropDef("mapping", typeDefs[250], "", true, nil),
		// with type string
		common.NewPropDef("max", typeDefs[46], "", false, nil),
		// with type integer
		common.NewPropDef("maxLength", typeDefs[45], "", false, nil),
		// with type date
		common.NewPropDef("maxValueDate", typeDefs[41], "", false, nil),
		// with type dateTime
		common.NewPropDef("maxValueDateTime", typeDefs[42], "", false, nil),
		// with type decimal
		common.NewPropDef("maxValueDecimal", typeDefs[43], "", false, nil),
		// with type instant
		common.NewPropDef("maxValueInstant", typeDefs[44], "", false, nil),
		// with type integer
		common.NewPropDef("maxValueInteger", typeDefs[45], "", false, nil),
		// with type positiveInt
		common.NewPropDef("maxValuePositiveInt", typeDefs[678], "", false, nil),
		// with type Quantity
		common.NewPropDef("maxValueQuantity", typeDefs[29], "", false, nil),
		// with type time
		common.NewPropDef("maxValueTime", typeDefs[47], "", false, nil),
		// with type unsignedInt
		common.NewPropDef("maxValueUnsignedInt", typeDefs[679], "", false, nil),
		// with type markdown
		common.NewPropDef("meaningWhenMissing", typeDefs[676], "", false, nil),
		// with type unsignedInt
		common.NewPropDef("min", typeDefs[679], "", false, nil),
		// with type date
		common.NewPropDef("minValueDate", typeDefs[41], "", false, nil),
		// with type dateTime
		common.NewPropDef("minValueDateTime", typeDefs[42], "", false, nil),
		// with type decimal
		common.NewPropDef("minValueDecimal", typeDefs[43], "", false, nil),
		// with type instant
		common.NewPropDef("minValueInstant", typeDefs[44], "", false, nil),
		// with type integer
		common.NewPropDef("minValueInteger", typeDefs[45], "", false, nil),
		// with type positiveInt
		common.NewPropDef("minValuePositiveInt", typeDefs[678], "", false, nil),
		// with type Quantity
		common.NewPropDef("minValueQuantity", typeDefs[29], "", false, nil),
		// with type time
		common.NewPropDef("minValueTime", typeDefs[47], "", false, nil),
		// with type unsignedInt
		common.NewPropDef("minValueUnsignedInt", typeDefs[679], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("mustSupport", typeDefs[40], "", false, nil),
		// with type string
		common.NewPropDef("orderMeaning", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("path", typeDefs[46], "", false, nil),
		// with type Address
		common.NewPropDef("patternAddress", typeDefs[2], "pattern", false, nil),
		// with type Age
		common.NewPropDef("patternAge", typeDefs[59], "pattern", false, nil),
		// with type Annotation
		common.NewPropDef("patternAnnotation", typeDefs[3], "pattern", false, nil),
		// with type Attachment
		common.NewPropDef("patternAttachment", typeDefs[4], "pattern", false, nil),
		// with type base64Binary
		common.NewPropDef("patternBase64Binary", typeDefs[39], "pattern", false, nil),
		// with type boolean
		common.NewPropDef("patternBoolean", typeDefs[40], "pattern", false, nil),
		// with type canonical
		common.NewPropDef("patternCanonical", typeDefs[673], "pattern", false, nil),
		// with type code
		common.NewPropDef("patternCode", typeDefs[674], "pattern", false, nil),
		// with type CodeableConcept
		common.NewPropDef("patternCodeableConcept", typeDefs[8], "pattern", false, nil),
		// with type Coding
		common.NewPropDef("patternCoding", typeDefs[9], "pattern", false, nil),
		// with type ContactDetail
		common.NewPropDef("patternContactDetail", typeDefs[10], "pattern", false, nil),
		// with type ContactPoint
		common.NewPropDef("patternContactPoint", typeDefs[11], "pattern", false, nil),
		// with type Contributor
		common.NewPropDef("patternContributor", typeDefs[12], "pattern", false, nil),
		// with type Count
		common.NewPropDef("patternCount", typeDefs[186], "pattern", false, nil),
		// with type DataRequirement
		common.NewPropDef("patternDataRequirement", typeDefs[13], "pattern", false, nil),
		// with type date
		common.NewPropDef("patternDate", typeDefs[41], "pattern", false, nil),
		// with type dateTime
		common.NewPropDef("patternDateTime", typeDefs[42], "pattern", false, nil),
		// with type decimal
		common.NewPropDef("patternDecimal", typeDefs[43], "pattern", false, nil),
		// with type Distance
		common.NewPropDef("patternDistance", typeDefs[227], "pattern", false, nil),
		// with type Dosage
		common.NewPropDef("patternDosage", typeDefs[234], "pattern", false, nil),
		// with type Duration
		common.NewPropDef("patternDuration", typeDefs[236], "pattern", false, nil),
		// with type Expression
		common.NewPropDef("patternExpression", typeDefs[15], "pattern", false, nil),
		// with type HumanName
		common.NewPropDef("patternHumanName", typeDefs[17], "pattern", false, nil),
		// with type id
		common.NewPropDef("patternId", typeDefs[675], "pattern", false, nil),
		// with type Identifier
		common.NewPropDef("patternIdentifier", typeDefs[18], "pattern", false, nil),
		// with type instant
		common.NewPropDef("patternInstant", typeDefs[44], "pattern", false, nil),
		// with type integer
		common.NewPropDef("patternInteger", typeDefs[45], "pattern", false, nil),
		// with type markdown
		common.NewPropDef("patternMarkdown", typeDefs[676], "pattern", false, nil),
		// with type Meta
		common.NewPropDef("patternMeta", typeDefs[20], "pattern", false, nil),
		// with type Money
		common.NewPropDef("patternMoney", typeDefs[21], "pattern", false, nil),
		// with type oid
		common.NewPropDef("patternOid", typeDefs[677], "pattern", false, nil),
		// with type ParameterDefinition
		common.NewPropDef("patternParameterDefinition", typeDefs[23], "pattern", false, nil),
		// with type Period
		common.NewPropDef("patternPeriod", typeDefs[25], "pattern", false, nil),
		// with type positiveInt
		common.NewPropDef("patternPositiveInt", typeDefs[678], "pattern", false, nil),
		// with type Quantity
		common.NewPropDef("patternQuantity", typeDefs[29], "pattern", false, nil),
		// with type Range
		common.NewPropDef("patternRange", typeDefs[30], "pattern", false, nil),
		// with type Ratio
		common.NewPropDef("patternRatio", typeDefs[31], "pattern", false, nil),
		// with type Reference
		common.NewPropDef("patternReference", typeDefs[32], "pattern", false, nil),
		// with type RelatedArtifact
		common.NewPropDef("patternRelatedArtifact", typeDefs[33], "pattern", false, nil),
		// with type SampledData
		common.NewPropDef("patternSampledData", typeDefs[34], "pattern", false, nil),
		// with type Signature
		common.NewPropDef("patternSignature", typeDefs[35], "pattern", false, nil),
		// with type string
		common.NewPropDef("patternString", typeDefs[46], "pattern", false, nil),
		// with type time
		common.NewPropDef("patternTime", typeDefs[47], "pattern", false, nil),
		// with type Timing
		common.NewPropDef("patternTiming", typeDefs[655], "pattern", false, nil),
		// with type TriggerDefinition
		common.NewPropDef("patternTriggerDefinition", typeDefs[37], "pattern", false, nil),
		// with type unsignedInt
		common.NewPropDef("patternUnsignedInt", typeDefs[679], "pattern", false, nil),
		// with type uri
		common.NewPropDef("patternUri", typeDefs[48], "pattern", false, nil),
		// with type url
		common.NewPropDef("patternUrl", typeDefs[680], "pattern", false, nil),
		// with type UsageContext
		common.NewPropDef("patternUsageContext", typeDefs[38], "pattern", false, nil),
		// with type uuid
		common.NewPropDef("patternUuid", typeDefs[681], "pattern", false, nil),
		// with type string
		common.NewPropDef("representation", typeDefs[46], "", true, []string{
			"xmlAttr",
			"xmlText",
			"typeAttr",
			"cdaText",
			"xhtml",
		}),
		// with type markdown
		common.NewPropDef("requirements", typeDefs[676], "", false, nil),
		// with type string
		common.NewPropDef("short", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("sliceIsConstraining", typeDefs[40], "", false, nil),
		// with type string
		common.NewPropDef("sliceName", typeDefs[46], "", false, nil),
		// with type ElementDefinition_Slicing
		common.NewPropDef("slicing", typeDefs[251], "", false, nil),
		// with type ElementDefinition_Type
		common.NewPropDef("type", typeDefs[252], "", true, nil),
	})
	// properties of ElementDefinition_Base
	typeDefs[245].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("max", typeDefs[46], "", false, nil),
		// with type unsignedInt
		common.NewPropDef("min", typeDefs[679], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("path", typeDefs[46], "", false, nil),
	})
	// properties of ElementDefinition_Binding
	typeDefs[246].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("strength", typeDefs[46], "", false, []string{
			"required",
			"extensible",
			"preferred",
			"example",
		}),
		// with type canonical
		common.NewPropDef("valueSet", typeDefs[673], "", false, nil),
	})
	// properties of ElementDefinition_Constraint
	typeDefs[247].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("expression", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("human", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type id
		common.NewPropDef("key", typeDefs[675], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("requirements", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("severity", typeDefs[46], "", false, []string{
			"error",
			"warning",
		}),
		// with type canonical
		common.NewPropDef("source", typeDefs[673], "", false, nil),
		// with type string
		common.NewPropDef("xpath", typeDefs[46], "", false, nil),
	})
	// properties of ElementDefinition_Discriminator
	typeDefs[248].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("path", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"value",
			"exists",
			"pattern",
			"type",
			"profile",
		}),
	})
	// properties of ElementDefinition_Example
	typeDefs[249].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("label", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Address
		common.NewPropDef("valueAddress", typeDefs[2], "value", false, nil),
		// with type Age
		common.NewPropDef("valueAge", typeDefs[59], "value", false, nil),
		// with type Annotation
		common.NewPropDef("valueAnnotation", typeDefs[3], "value", false, nil),
		// with type Attachment
		common.NewPropDef("valueAttachment", typeDefs[4], "value", false, nil),
		// with type base64Binary
		common.NewPropDef("valueBase64Binary", typeDefs[39], "value", false, nil),
		// with type boolean
		common.NewPropDef("valueBoolean", typeDefs[40], "value", false, nil),
		// with type canonical
		common.NewPropDef("valueCanonical", typeDefs[673], "value", false, nil),
		// with type code
		common.NewPropDef("valueCode", typeDefs[674], "value", false, nil),
		// with type CodeableConcept
		common.NewPropDef("valueCodeableConcept", typeDefs[8], "value", false, nil),
		// with type Coding
		common.NewPropDef("valueCoding", typeDefs[9], "value", false, nil),
		// with type ContactDetail
		common.NewPropDef("valueContactDetail", typeDefs[10], "value", false, nil),
		// with type ContactPoint
		common.NewPropDef("valueContactPoint", typeDefs[11], "value", false, nil),
		// with type Contributor
		common.NewPropDef("valueContributor", typeDefs[12], "value", false, nil),
		// with type Count
		common.NewPropDef("valueCount", typeDefs[186], "value", false, nil),
		// with type DataRequirement
		common.NewPropDef("valueDataRequirement", typeDefs[13], "value", false, nil),
		// with type date
		common.NewPropDef("valueDate", typeDefs[41], "value", false, nil),
		// with type dateTime
		common.NewPropDef("valueDateTime", typeDefs[42], "value", false, nil),
		// with type decimal
		common.NewPropDef("valueDecimal", typeDefs[43], "value", false, nil),
		// with type Distance
		common.NewPropDef("valueDistance", typeDefs[227], "value", false, nil),
		// with type Dosage
		common.NewPropDef("valueDosage", typeDefs[234], "value", false, nil),
		// with type Duration
		common.NewPropDef("valueDuration", typeDefs[236], "value", false, nil),
		// with type Expression
		common.NewPropDef("valueExpression", typeDefs[15], "value", false, nil),
		// with type HumanName
		common.NewPropDef("valueHumanName", typeDefs[17], "value", false, nil),
		// with type id
		common.NewPropDef("valueId", typeDefs[675], "value", false, nil),
		// with type Identifier
		common.NewPropDef("valueIdentifier", typeDefs[18], "value", false, nil),
		// with type instant
		common.NewPropDef("valueInstant", typeDefs[44], "value", false, nil),
		// with type integer
		common.NewPropDef("valueInteger", typeDefs[45], "value", false, nil),
		// with type markdown
		common.NewPropDef("valueMarkdown", typeDefs[676], "value", false, nil),
		// with type Meta
		common.NewPropDef("valueMeta", typeDefs[20], "value", false, nil),
		// with type Money
		common.NewPropDef("valueMoney", typeDefs[21], "value", false, nil),
		// with type oid
		common.NewPropDef("valueOid", typeDefs[677], "value", false, nil),
		// with type ParameterDefinition
		common.NewPropDef("valueParameterDefinition", typeDefs[23], "value", false, nil),
		// with type Period
		common.NewPropDef("valuePeriod", typeDefs[25], "value", false, nil),
		// with type positiveInt
		common.NewPropDef("valuePositiveInt", typeDefs[678], "value", false, nil),
		// with type Quantity
		common.NewPropDef("valueQuantity", typeDefs[29], "value", false, nil),
		// with type Range
		common.NewPropDef("valueRange", typeDefs[30], "value", false, nil),
		// with type Ratio
		common.NewPropDef("valueRatio", typeDefs[31], "value", false, nil),
		// with type Reference
		common.NewPropDef("valueReference", typeDefs[32], "value", false, nil),
		// with type RelatedArtifact
		common.NewPropDef("valueRelatedArtifact", typeDefs[33], "value", false, nil),
		// with type SampledData
		common.NewPropDef("valueSampledData", typeDefs[34], "value", false, nil),
		// with type Signature
		common.NewPropDef("valueSignature", typeDefs[35], "value", false, nil),
		// with type string
		common.NewPropDef("valueString", typeDefs[46], "value", false, nil),
		// with type time
		common.NewPropDef("valueTime", typeDefs[47], "value", false, nil),
		// with type Timing
		common.NewPropDef("valueTiming", typeDefs[655], "value", false, nil),
		// with type TriggerDefinition
		common.NewPropDef("valueTriggerDefinition", typeDefs[37], "value", false, nil),
		// with type unsignedInt
		common.NewPropDef("valueUnsignedInt", typeDefs[679], "value", false, nil),
		// with type uri
		common.NewPropDef("valueUri", typeDefs[48], "value", false, nil),
		// with type url
		common.NewPropDef("valueUrl", typeDefs[680], "value", false, nil),
		// with type UsageContext
		common.NewPropDef("valueUsageContext", typeDefs[38], "value", false, nil),
		// with type uuid
		common.NewPropDef("valueUuid", typeDefs[681], "value", false, nil),
	})
	// properties of ElementDefinition_Mapping
	typeDefs[250].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("comment", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type id
		common.NewPropDef("identity", typeDefs[675], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("map", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of ElementDefinition_Slicing
	typeDefs[251].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type ElementDefinition_Discriminator
		common.NewPropDef("discriminator", typeDefs[248], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("ordered", typeDefs[40], "", false, nil),
		// with type string
		common.NewPropDef("rules", typeDefs[46], "", false, []string{
			"closed",
			"open",
			"openAtEnd",
		}),
	})
	// properties of ElementDefinition_Type
	typeDefs[252].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("aggregation", typeDefs[46], "", true, []string{
			"contained",
			"referenced",
			"bundled",
		}),
		// with type uri
		common.NewPropDef("code", typeDefs[48], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type canonical
		common.NewPropDef("profile", typeDefs[673], "", true, nil),
		// with type canonical
		common.NewPropDef("targetProfile", typeDefs[673], "", true, nil),
		// with type string
		common.NewPropDef("versioning", typeDefs[46], "", false, []string{
			"either",
			"independent",
			"specific",
		}),
	})
	// properties of Encounter
	typeDefs[253].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("account", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("appointment", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("basedOn", typeDefs[32], "", true, nil),
		// with type Coding
		common.NewPropDef("class", typeDefs[9], "", false, nil),
		// with type Encounter_ClassHistory
		common.NewPropDef("classHistory", typeDefs[254], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Encounter_Diagnosis
		common.NewPropDef("diagnosis", typeDefs[255], "", true, nil),
		// with type Reference
		common.NewPropDef("episodeOfCare", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Encounter_Hospitalization
		common.NewPropDef("hospitalization", typeDefs[256], "", false, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Duration
		common.NewPropDef("length", typeDefs[236], "", false, nil),
		// with type Encounter_Location
		common.NewPropDef("location", typeDefs[257], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("partOf", typeDefs[32], "", false, nil),
		// with type Encounter_Participant
		common.NewPropDef("participant", typeDefs[258], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("priority", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("reasonCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("reasonReference", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("serviceProvider", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("serviceType", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"planned",
			"arrived",
			"triaged",
			"in-progress",
			"onleave",
			"finished",
			"cancelled",
			"entered-in-error",
			"unknown",
		}),
		// with type Encounter_StatusHistory
		common.NewPropDef("statusHistory", typeDefs[259], "", true, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", true, nil),
	})
	// properties of Encounter_ClassHistory
	typeDefs[254].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Coding
		common.NewPropDef("class", typeDefs[9], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
	})
	// properties of Encounter_Diagnosis
	typeDefs[255].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("condition", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type positiveInt
		common.NewPropDef("rank", typeDefs[678], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("use", typeDefs[8], "", false, nil),
	})
	// properties of Encounter_Hospitalization
	typeDefs[256].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("admitSource", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("destination", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("dietPreference", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("dischargeDisposition", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("origin", typeDefs[32], "", false, nil),
		// with type Identifier
		common.NewPropDef("preAdmissionIdentifier", typeDefs[18], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("reAdmission", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("specialArrangement", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("specialCourtesy", typeDefs[8], "", true, nil),
	})
	// properties of Encounter_Location
	typeDefs[257].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("location", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("physicalType", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"planned",
			"active",
			"reserved",
			"completed",
		}),
	})
	// properties of Encounter_Participant
	typeDefs[258].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("individual", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", true, nil),
	})
	// properties of Encounter_StatusHistory
	typeDefs[259].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"planned",
			"arrived",
			"triaged",
			"in-progress",
			"onleave",
			"finished",
			"cancelled",
			"entered-in-error",
			"unknown",
		}),
	})
	// properties of Endpoint
	typeDefs[260].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type url
		common.NewPropDef("address", typeDefs[680], "", false, nil),
		// with type Coding
		common.NewPropDef("connectionType", typeDefs[9], "", false, nil),
		// with type ContactPoint
		common.NewPropDef("contact", typeDefs[11], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("header", typeDefs[46], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("managingOrganization", typeDefs[32], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type code
		common.NewPropDef("payloadMimeType", typeDefs[674], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("payloadType", typeDefs[8], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"active",
			"suspended",
			"error",
			"off",
			"entered-in-error",
			"test",
		}),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of EnrollmentRequest
	typeDefs[261].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("candidate", typeDefs[32], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("coverage", typeDefs[32], "", false, nil),
		// with type dateTime
		common.NewPropDef("created", typeDefs[42], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type Reference
		common.NewPropDef("insurer", typeDefs[32], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("provider", typeDefs[32], "", false, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of EnrollmentResponse
	typeDefs[262].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type dateTime
		common.NewPropDef("created", typeDefs[42], "", false, nil),
		// with type string
		common.NewPropDef("disposition", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("organization", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("outcome", typeDefs[46], "", false, []string{
			"queued",
			"complete",
			"error",
			"partial",
		}),
		// with type Reference
		common.NewPropDef("request", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("requestProvider", typeDefs[32], "", false, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of EpisodeOfCare
	typeDefs[263].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("account", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("careManager", typeDefs[32], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type EpisodeOfCare_Diagnosis
		common.NewPropDef("diagnosis", typeDefs[264], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("managingOrganization", typeDefs[32], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("patient", typeDefs[32], "", false, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type Reference
		common.NewPropDef("referralRequest", typeDefs[32], "", true, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"planned",
			"waitlist",
			"active",
			"onhold",
			"finished",
			"cancelled",
			"entered-in-error",
		}),
		// with type EpisodeOfCare_StatusHistory
		common.NewPropDef("statusHistory", typeDefs[265], "", true, nil),
		// with type Reference
		common.NewPropDef("team", typeDefs[32], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", true, nil),
	})
	// properties of EpisodeOfCare_Diagnosis
	typeDefs[264].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("condition", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type positiveInt
		common.NewPropDef("rank", typeDefs[678], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("role", typeDefs[8], "", false, nil),
	})
	// properties of EpisodeOfCare_StatusHistory
	typeDefs[265].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"planned",
			"waitlist",
			"active",
			"onhold",
			"finished",
			"cancelled",
			"entered-in-error",
		}),
	})
	// properties of EventDefinition
	typeDefs[266].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type date
		common.NewPropDef("approvalDate", typeDefs[41], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("author", typeDefs[10], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("editor", typeDefs[10], "", true, nil),
		// with type Period
		common.NewPropDef("effectivePeriod", typeDefs[25], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("endorser", typeDefs[10], "", true, nil),
		// with type boolean
		common.NewPropDef("experimental", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type date
		common.NewPropDef("lastReviewDate", typeDefs[41], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("purpose", typeDefs[676], "", false, nil),
		// with type RelatedArtifact
		common.NewPropDef("relatedArtifact", typeDefs[33], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("reviewer", typeDefs[10], "", true, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type CodeableConcept
		common.NewPropDef("subjectCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("subjectReference", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("subtitle", typeDefs[46], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("topic", typeDefs[8], "", true, nil),
		// with type TriggerDefinition
		common.NewPropDef("trigger", typeDefs[37], "", true, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type string
		common.NewPropDef("usage", typeDefs[46], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of Evidence
	typeDefs[267].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type date
		common.NewPropDef("approvalDate", typeDefs[41], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("author", typeDefs[10], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("editor", typeDefs[10], "", true, nil),
		// with type Period
		common.NewPropDef("effectivePeriod", typeDefs[25], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("endorser", typeDefs[10], "", true, nil),
		// with type Reference
		common.NewPropDef("exposureBackground", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("exposureVariant", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type date
		common.NewPropDef("lastReviewDate", typeDefs[41], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type Reference
		common.NewPropDef("outcome", typeDefs[32], "", true, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type RelatedArtifact
		common.NewPropDef("relatedArtifact", typeDefs[33], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("reviewer", typeDefs[10], "", true, nil),
		// with type string
		common.NewPropDef("shortTitle", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type string
		common.NewPropDef("subtitle", typeDefs[46], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("topic", typeDefs[8], "", true, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of EvidenceVariable
	typeDefs[268].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type date
		common.NewPropDef("approvalDate", typeDefs[41], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("author", typeDefs[10], "", true, nil),
		// with type EvidenceVariable_Characteristic
		common.NewPropDef("characteristic", typeDefs[269], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("editor", typeDefs[10], "", true, nil),
		// with type Period
		common.NewPropDef("effectivePeriod", typeDefs[25], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("endorser", typeDefs[10], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type date
		common.NewPropDef("lastReviewDate", typeDefs[41], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type RelatedArtifact
		common.NewPropDef("relatedArtifact", typeDefs[33], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("reviewer", typeDefs[10], "", true, nil),
		// with type string
		common.NewPropDef("shortTitle", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type string
		common.NewPropDef("subtitle", typeDefs[46], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("topic", typeDefs[8], "", true, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"dichotomous",
			"continuous",
			"descriptive",
		}),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of EvidenceVariable_Characteristic
	typeDefs[269].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type canonical
		common.NewPropDef("definitionCanonical", typeDefs[673], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("definitionCodeableConcept", typeDefs[8], "", false, nil),
		// with type DataRequirement
		common.NewPropDef("definitionDataRequirement", typeDefs[13], "", false, nil),
		// with type Expression
		common.NewPropDef("definitionExpression", typeDefs[15], "", false, nil),
		// with type Reference
		common.NewPropDef("definitionReference", typeDefs[32], "", false, nil),
		// with type TriggerDefinition
		common.NewPropDef("definitionTriggerDefinition", typeDefs[37], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("exclude", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("groupMeasure", typeDefs[46], "", false, []string{
			"mean",
			"median",
			"mean-of-mean",
			"mean-of-median",
			"median-of-mean",
			"median-of-median",
		}),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type dateTime
		common.NewPropDef("participantEffectiveDateTime", typeDefs[42], "", false, nil),
		// with type Duration
		common.NewPropDef("participantEffectiveDuration", typeDefs[236], "", false, nil),
		// with type Period
		common.NewPropDef("participantEffectivePeriod", typeDefs[25], "", false, nil),
		// with type Timing
		common.NewPropDef("participantEffectiveTiming", typeDefs[655], "", false, nil),
		// with type Duration
		common.NewPropDef("timeFromStart", typeDefs[236], "", false, nil),
		// with type UsageContext
		common.NewPropDef("usageContext", typeDefs[38], "", true, nil),
	})
	// properties of ExampleScenario
	typeDefs[270].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ExampleScenario_Actor
		common.NewPropDef("actor", typeDefs[271], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type boolean
		common.NewPropDef("experimental", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type ExampleScenario_Instance
		common.NewPropDef("instance", typeDefs[274], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type ExampleScenario_Process
		common.NewPropDef("process", typeDefs[276], "", true, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("purpose", typeDefs[676], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
		// with type canonical
		common.NewPropDef("workflow", typeDefs[673], "", true, nil),
	})
	// properties of ExampleScenario_Actor
	typeDefs[271].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("actorId", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"person",
			"entity",
		}),
	})
	// properties of ExampleScenario_Alternative
	typeDefs[272].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type ExampleScenario_Step
		common.NewPropDef("step", typeDefs[277], "", true, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
	})
	// properties of ExampleScenario_ContainedInstance
	typeDefs[273].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("resourceId", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("versionId", typeDefs[46], "", false, nil),
	})
	// properties of ExampleScenario_Instance
	typeDefs[274].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ExampleScenario_ContainedInstance
		common.NewPropDef("containedInstance", typeDefs[273], "", true, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("resourceId", typeDefs[46], "", false, nil),
		// with type code
		common.NewPropDef("resourceType", typeDefs[674], "", false, nil),
		// with type ExampleScenario_Version
		common.NewPropDef("version", typeDefs[278], "", true, nil),
	})
	// properties of ExampleScenario_Operation
	typeDefs[275].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("initiator", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("initiatorActive", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("number", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("receiver", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("receiverActive", typeDefs[40], "", false, nil),
		// with type ExampleScenario_ContainedInstance
		common.NewPropDef("request", typeDefs[273], "", false, nil),
		// with type ExampleScenario_ContainedInstance
		common.NewPropDef("response", typeDefs[273], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, nil),
	})
	// properties of ExampleScenario_Process
	typeDefs[276].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type markdown
		common.NewPropDef("postConditions", typeDefs[676], "", false, nil),
		// with type markdown
		common.NewPropDef("preConditions", typeDefs[676], "", false, nil),
		// with type ExampleScenario_Step
		common.NewPropDef("step", typeDefs[277], "", true, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
	})
	// properties of ExampleScenario_Step
	typeDefs[277].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ExampleScenario_Alternative
		common.NewPropDef("alternative", typeDefs[272], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type ExampleScenario_Operation
		common.NewPropDef("operation", typeDefs[275], "", false, nil),
		// with type boolean
		common.NewPropDef("pause", typeDefs[40], "", false, nil),
		// with type ExampleScenario_Process
		common.NewPropDef("process", typeDefs[276], "", true, nil),
	})
	// properties of ExampleScenario_Version
	typeDefs[278].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("versionId", typeDefs[46], "", false, nil),
	})
	// properties of ExplanationOfBenefit
	typeDefs[279].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ExplanationOfBenefit_Accident
		common.NewPropDef("accident", typeDefs[280], "", false, nil),
		// with type ExplanationOfBenefit_AddItem
		common.NewPropDef("addItem", typeDefs[281], "", true, nil),
		// with type ExplanationOfBenefit_Adjudication
		common.NewPropDef("adjudication", typeDefs[282], "", true, nil),
		// with type ExplanationOfBenefit_BenefitBalance
		common.NewPropDef("benefitBalance", typeDefs[283], "", true, nil),
		// with type Period
		common.NewPropDef("benefitPeriod", typeDefs[25], "", false, nil),
		// with type Period
		common.NewPropDef("billablePeriod", typeDefs[25], "", false, nil),
		// with type ExplanationOfBenefit_CareTeam
		common.NewPropDef("careTeam", typeDefs[284], "", true, nil),
		// with type Reference
		common.NewPropDef("claim", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("claimResponse", typeDefs[32], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type dateTime
		common.NewPropDef("created", typeDefs[42], "", false, nil),
		// with type ExplanationOfBenefit_Diagnosis
		common.NewPropDef("diagnosis", typeDefs[287], "", true, nil),
		// with type string
		common.NewPropDef("disposition", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("enterer", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("facility", typeDefs[32], "", false, nil),
		// with type Attachment
		common.NewPropDef("form", typeDefs[4], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("formCode", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("fundsReserve", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("fundsReserveRequested", typeDefs[8], "", false, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type ExplanationOfBenefit_Insurance
		common.NewPropDef("insurance", typeDefs[289], "", true, nil),
		// with type Reference
		common.NewPropDef("insurer", typeDefs[32], "", false, nil),
		// with type ExplanationOfBenefit_Item
		common.NewPropDef("item", typeDefs[290], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("originalPrescription", typeDefs[32], "", false, nil),
		// with type code
		common.NewPropDef("outcome", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("patient", typeDefs[32], "", false, nil),
		// with type ExplanationOfBenefit_Payee
		common.NewPropDef("payee", typeDefs[291], "", false, nil),
		// with type ExplanationOfBenefit_Payment
		common.NewPropDef("payment", typeDefs[292], "", false, nil),
		// with type string
		common.NewPropDef("preAuthRef", typeDefs[46], "", true, nil),
		// with type Period
		common.NewPropDef("preAuthRefPeriod", typeDefs[25], "", true, nil),
		// with type positiveInt
		common.NewPropDef("precedence", typeDefs[678], "", false, nil),
		// with type Reference
		common.NewPropDef("prescription", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("priority", typeDefs[8], "", false, nil),
		// with type ExplanationOfBenefit_Procedure
		common.NewPropDef("procedure", typeDefs[293], "", true, nil),
		// with type ExplanationOfBenefit_ProcessNote
		common.NewPropDef("processNote", typeDefs[294], "", true, nil),
		// with type Reference
		common.NewPropDef("provider", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("referral", typeDefs[32], "", false, nil),
		// with type ExplanationOfBenefit_Related
		common.NewPropDef("related", typeDefs[295], "", true, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"active",
			"cancelled",
			"draft",
			"entered-in-error",
		}),
		// with type CodeableConcept
		common.NewPropDef("subType", typeDefs[8], "", false, nil),
		// with type ExplanationOfBenefit_SupportingInfo
		common.NewPropDef("supportingInfo", typeDefs[298], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type ExplanationOfBenefit_Total
		common.NewPropDef("total", typeDefs[299], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type code
		common.NewPropDef("use", typeDefs[674], "", false, nil),
	})
	// properties of ExplanationOfBenefit_Accident
	typeDefs[280].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type date
		common.NewPropDef("date", typeDefs[41], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Address
		common.NewPropDef("locationAddress", typeDefs[2], "", false, nil),
		// with type Reference
		common.NewPropDef("locationReference", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of ExplanationOfBenefit_AddItem
	typeDefs[281].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ExplanationOfBenefit_Adjudication
		common.NewPropDef("adjudication", typeDefs[282], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("bodySite", typeDefs[8], "", false, nil),
		// with type ExplanationOfBenefit_Detail1
		common.NewPropDef("detail", typeDefs[286], "", true, nil),
		// with type positiveInt
		common.NewPropDef("detailSequence", typeDefs[678], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("factor", typeDefs[43], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type positiveInt
		common.NewPropDef("itemSequence", typeDefs[678], "", true, nil),
		// with type Address
		common.NewPropDef("locationAddress", typeDefs[2], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("locationCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("locationReference", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("modifier", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Money
		common.NewPropDef("net", typeDefs[21], "", false, nil),
		// with type positiveInt
		common.NewPropDef("noteNumber", typeDefs[678], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("productOrService", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("programCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("provider", typeDefs[32], "", true, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type date
		common.NewPropDef("servicedDate", typeDefs[41], "", false, nil),
		// with type Period
		common.NewPropDef("servicedPeriod", typeDefs[25], "", false, nil),
		// with type positiveInt
		common.NewPropDef("subDetailSequence", typeDefs[678], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("subSite", typeDefs[8], "", true, nil),
		// with type Money
		common.NewPropDef("unitPrice", typeDefs[21], "", false, nil),
	})
	// properties of ExplanationOfBenefit_Adjudication
	typeDefs[282].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Money
		common.NewPropDef("amount", typeDefs[21], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("reason", typeDefs[8], "", false, nil),
		// with type decimal
		common.NewPropDef("value", typeDefs[43], "", false, nil),
	})
	// properties of ExplanationOfBenefit_BenefitBalance
	typeDefs[283].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("excluded", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type ExplanationOfBenefit_Financial
		common.NewPropDef("financial", typeDefs[288], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("network", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("term", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("unit", typeDefs[8], "", false, nil),
	})
	// properties of ExplanationOfBenefit_CareTeam
	typeDefs[284].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("provider", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("qualification", typeDefs[8], "", false, nil),
		// with type boolean
		common.NewPropDef("responsible", typeDefs[40], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("role", typeDefs[8], "", false, nil),
		// with type positiveInt
		common.NewPropDef("sequence", typeDefs[678], "", false, nil),
	})
	// properties of ExplanationOfBenefit_Detail
	typeDefs[285].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ExplanationOfBenefit_Adjudication
		common.NewPropDef("adjudication", typeDefs[282], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("factor", typeDefs[43], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("modifier", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Money
		common.NewPropDef("net", typeDefs[21], "", false, nil),
		// with type positiveInt
		common.NewPropDef("noteNumber", typeDefs[678], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("productOrService", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("programCode", typeDefs[8], "", true, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("revenue", typeDefs[8], "", false, nil),
		// with type positiveInt
		common.NewPropDef("sequence", typeDefs[678], "", false, nil),
		// with type ExplanationOfBenefit_SubDetail
		common.NewPropDef("subDetail", typeDefs[296], "", true, nil),
		// with type Reference
		common.NewPropDef("udi", typeDefs[32], "", true, nil),
		// with type Money
		common.NewPropDef("unitPrice", typeDefs[21], "", false, nil),
	})
	// properties of ExplanationOfBenefit_Detail1
	typeDefs[286].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ExplanationOfBenefit_Adjudication
		common.NewPropDef("adjudication", typeDefs[282], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("factor", typeDefs[43], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("modifier", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Money
		common.NewPropDef("net", typeDefs[21], "", false, nil),
		// with type positiveInt
		common.NewPropDef("noteNumber", typeDefs[678], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("productOrService", typeDefs[8], "", false, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type ExplanationOfBenefit_SubDetail1
		common.NewPropDef("subDetail", typeDefs[297], "", true, nil),
		// with type Money
		common.NewPropDef("unitPrice", typeDefs[21], "", false, nil),
	})
	// properties of ExplanationOfBenefit_Diagnosis
	typeDefs[287].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("diagnosisCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("diagnosisReference", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("onAdmission", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("packageCode", typeDefs[8], "", false, nil),
		// with type positiveInt
		common.NewPropDef("sequence", typeDefs[678], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", true, nil),
	})
	// properties of ExplanationOfBenefit_Financial
	typeDefs[288].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Money
		common.NewPropDef("allowedMoney", typeDefs[21], "", false, nil),
		// with type string
		common.NewPropDef("allowedString", typeDefs[46], "", false, nil),
		// with type unsignedInt
		common.NewPropDef("allowedUnsignedInt", typeDefs[679], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type Money
		common.NewPropDef("usedMoney", typeDefs[21], "", false, nil),
		// with type unsignedInt
		common.NewPropDef("usedUnsignedInt", typeDefs[679], "", false, nil),
	})
	// properties of ExplanationOfBenefit_Insurance
	typeDefs[289].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("coverage", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("focal", typeDefs[40], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("preAuthRef", typeDefs[46], "", true, nil),
	})
	// properties of ExplanationOfBenefit_Item
	typeDefs[290].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ExplanationOfBenefit_Adjudication
		common.NewPropDef("adjudication", typeDefs[282], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("bodySite", typeDefs[8], "", false, nil),
		// with type positiveInt
		common.NewPropDef("careTeamSequence", typeDefs[678], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", false, nil),
		// with type ExplanationOfBenefit_Detail
		common.NewPropDef("detail", typeDefs[285], "", true, nil),
		// with type positiveInt
		common.NewPropDef("diagnosisSequence", typeDefs[678], "", true, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("factor", typeDefs[43], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type positiveInt
		common.NewPropDef("informationSequence", typeDefs[678], "", true, nil),
		// with type Address
		common.NewPropDef("locationAddress", typeDefs[2], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("locationCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("locationReference", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("modifier", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Money
		common.NewPropDef("net", typeDefs[21], "", false, nil),
		// with type positiveInt
		common.NewPropDef("noteNumber", typeDefs[678], "", true, nil),
		// with type positiveInt
		common.NewPropDef("procedureSequence", typeDefs[678], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("productOrService", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("programCode", typeDefs[8], "", true, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("revenue", typeDefs[8], "", false, nil),
		// with type positiveInt
		common.NewPropDef("sequence", typeDefs[678], "", false, nil),
		// with type date
		common.NewPropDef("servicedDate", typeDefs[41], "", false, nil),
		// with type Period
		common.NewPropDef("servicedPeriod", typeDefs[25], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("subSite", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("udi", typeDefs[32], "", true, nil),
		// with type Money
		common.NewPropDef("unitPrice", typeDefs[21], "", false, nil),
	})
	// properties of ExplanationOfBenefit_Payee
	typeDefs[291].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("party", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of ExplanationOfBenefit_Payment
	typeDefs[292].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Money
		common.NewPropDef("adjustment", typeDefs[21], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("adjustmentReason", typeDefs[8], "", false, nil),
		// with type Money
		common.NewPropDef("amount", typeDefs[21], "", false, nil),
		// with type date
		common.NewPropDef("date", typeDefs[41], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of ExplanationOfBenefit_Procedure
	typeDefs[293].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("procedureCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("procedureReference", typeDefs[32], "", false, nil),
		// with type positiveInt
		common.NewPropDef("sequence", typeDefs[678], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("udi", typeDefs[32], "", true, nil),
	})
	// properties of ExplanationOfBenefit_ProcessNote
	typeDefs[294].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("language", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type positiveInt
		common.NewPropDef("number", typeDefs[678], "", false, nil),
		// with type string
		common.NewPropDef("text", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"display",
			"print",
			"printoper",
		}),
	})
	// properties of ExplanationOfBenefit_Related
	typeDefs[295].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("claim", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Identifier
		common.NewPropDef("reference", typeDefs[18], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("relationship", typeDefs[8], "", false, nil),
	})
	// properties of ExplanationOfBenefit_SubDetail
	typeDefs[296].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ExplanationOfBenefit_Adjudication
		common.NewPropDef("adjudication", typeDefs[282], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("factor", typeDefs[43], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("modifier", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Money
		common.NewPropDef("net", typeDefs[21], "", false, nil),
		// with type positiveInt
		common.NewPropDef("noteNumber", typeDefs[678], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("productOrService", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("programCode", typeDefs[8], "", true, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("revenue", typeDefs[8], "", false, nil),
		// with type positiveInt
		common.NewPropDef("sequence", typeDefs[678], "", false, nil),
		// with type Reference
		common.NewPropDef("udi", typeDefs[32], "", true, nil),
		// with type Money
		common.NewPropDef("unitPrice", typeDefs[21], "", false, nil),
	})
	// properties of ExplanationOfBenefit_SubDetail1
	typeDefs[297].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ExplanationOfBenefit_Adjudication
		common.NewPropDef("adjudication", typeDefs[282], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("factor", typeDefs[43], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("modifier", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Money
		common.NewPropDef("net", typeDefs[21], "", false, nil),
		// with type positiveInt
		common.NewPropDef("noteNumber", typeDefs[678], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("productOrService", typeDefs[8], "", false, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type Money
		common.NewPropDef("unitPrice", typeDefs[21], "", false, nil),
	})
	// properties of ExplanationOfBenefit_SupportingInfo
	typeDefs[298].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Coding
		common.NewPropDef("reason", typeDefs[9], "", false, nil),
		// with type positiveInt
		common.NewPropDef("sequence", typeDefs[678], "", false, nil),
		// with type date
		common.NewPropDef("timingDate", typeDefs[41], "", false, nil),
		// with type Period
		common.NewPropDef("timingPeriod", typeDefs[25], "", false, nil),
		// with type Attachment
		common.NewPropDef("valueAttachment", typeDefs[4], "", false, nil),
		// with type boolean
		common.NewPropDef("valueBoolean", typeDefs[40], "", false, nil),
		// with type Quantity
		common.NewPropDef("valueQuantity", typeDefs[29], "", false, nil),
		// with type Reference
		common.NewPropDef("valueReference", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("valueString", typeDefs[46], "", false, nil),
	})
	// properties of ExplanationOfBenefit_Total
	typeDefs[299].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Money
		common.NewPropDef("amount", typeDefs[21], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Expression
	typeDefs[15].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("expression", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("language", typeDefs[46], "", false, []string{
			"text/cql",
			"text/fhirpath",
			"application/x-fhir-query",
		}),
		// with type id
		common.NewPropDef("name", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("reference", typeDefs[48], "", false, nil),
	})
	// properties of Extension
	typeDefs[16].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type Address
		common.NewPropDef("valueAddress", typeDefs[2], "value", false, nil),
		// with type Age
		common.NewPropDef("valueAge", typeDefs[59], "value", false, nil),
		// with type Annotation
		common.NewPropDef("valueAnnotation", typeDefs[3], "value", false, nil),
		// with type Attachment
		common.NewPropDef("valueAttachment", typeDefs[4], "value", false, nil),
		// with type base64Binary
		common.NewPropDef("valueBase64Binary", typeDefs[39], "value", false, nil),
		// with type boolean
		common.NewPropDef("valueBoolean", typeDefs[40], "value", false, nil),
		// with type canonical
		common.NewPropDef("valueCanonical", typeDefs[673], "value", false, nil),
		// with type code
		common.NewPropDef("valueCode", typeDefs[674], "value", false, nil),
		// with type CodeableConcept
		common.NewPropDef("valueCodeableConcept", typeDefs[8], "value", false, nil),
		// with type Coding
		common.NewPropDef("valueCoding", typeDefs[9], "value", false, nil),
		// with type ContactDetail
		common.NewPropDef("valueContactDetail", typeDefs[10], "value", false, nil),
		// with type ContactPoint
		common.NewPropDef("valueContactPoint", typeDefs[11], "value", false, nil),
		// with type Contributor
		common.NewPropDef("valueContributor", typeDefs[12], "value", false, nil),
		// with type Count
		common.NewPropDef("valueCount", typeDefs[186], "value", false, nil),
		// with type DataRequirement
		common.NewPropDef("valueDataRequirement", typeDefs[13], "value", false, nil),
		// with type date
		common.NewPropDef("valueDate", typeDefs[41], "value", false, nil),
		// with type dateTime
		common.NewPropDef("valueDateTime", typeDefs[42], "value", false, nil),
		// with type decimal
		common.NewPropDef("valueDecimal", typeDefs[43], "value", false, nil),
		// with type Distance
		common.NewPropDef("valueDistance", typeDefs[227], "value", false, nil),
		// with type Dosage
		common.NewPropDef("valueDosage", typeDefs[234], "value", false, nil),
		// with type Duration
		common.NewPropDef("valueDuration", typeDefs[236], "value", false, nil),
		// with type Expression
		common.NewPropDef("valueExpression", typeDefs[15], "value", false, nil),
		// with type HumanName
		common.NewPropDef("valueHumanName", typeDefs[17], "value", false, nil),
		// with type id
		common.NewPropDef("valueId", typeDefs[675], "value", false, nil),
		// with type Identifier
		common.NewPropDef("valueIdentifier", typeDefs[18], "value", false, nil),
		// with type instant
		common.NewPropDef("valueInstant", typeDefs[44], "value", false, nil),
		// with type integer
		common.NewPropDef("valueInteger", typeDefs[45], "value", false, nil),
		// with type markdown
		common.NewPropDef("valueMarkdown", typeDefs[676], "value", false, nil),
		// with type Meta
		common.NewPropDef("valueMeta", typeDefs[20], "value", false, nil),
		// with type Money
		common.NewPropDef("valueMoney", typeDefs[21], "value", false, nil),
		// with type oid
		common.NewPropDef("valueOid", typeDefs[677], "value", false, nil),
		// with type ParameterDefinition
		common.NewPropDef("valueParameterDefinition", typeDefs[23], "value", false, nil),
		// with type Period
		common.NewPropDef("valuePeriod", typeDefs[25], "value", false, nil),
		// with type positiveInt
		common.NewPropDef("valuePositiveInt", typeDefs[678], "value", false, nil),
		// with type Quantity
		common.NewPropDef("valueQuantity", typeDefs[29], "value", false, nil),
		// with type Range
		common.NewPropDef("valueRange", typeDefs[30], "value", false, nil),
		// with type Ratio
		common.NewPropDef("valueRatio", typeDefs[31], "value", false, nil),
		// with type Reference
		common.NewPropDef("valueReference", typeDefs[32], "value", false, nil),
		// with type RelatedArtifact
		common.NewPropDef("valueRelatedArtifact", typeDefs[33], "value", false, nil),
		// with type SampledData
		common.NewPropDef("valueSampledData", typeDefs[34], "value", false, nil),
		// with type Signature
		common.NewPropDef("valueSignature", typeDefs[35], "value", false, nil),
		// with type string
		common.NewPropDef("valueString", typeDefs[46], "value", false, nil),
		// with type time
		common.NewPropDef("valueTime", typeDefs[47], "value", false, nil),
		// with type Timing
		common.NewPropDef("valueTiming", typeDefs[655], "value", false, nil),
		// with type TriggerDefinition
		common.NewPropDef("valueTriggerDefinition", typeDefs[37], "value", false, nil),
		// with type unsignedInt
		common.NewPropDef("valueUnsignedInt", typeDefs[679], "value", false, nil),
		// with type uri
		common.NewPropDef("valueUri", typeDefs[48], "value", false, nil),
		// with type url
		common.NewPropDef("valueUrl", typeDefs[680], "value", false, nil),
		// with type UsageContext
		common.NewPropDef("valueUsageContext", typeDefs[38], "value", false, nil),
		// with type uuid
		common.NewPropDef("valueUuid", typeDefs[681], "value", false, nil),
	})
	// properties of FamilyMemberHistory
	typeDefs[300].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Age
		common.NewPropDef("ageAge", typeDefs[59], "", false, nil),
		// with type Range
		common.NewPropDef("ageRange", typeDefs[30], "", false, nil),
		// with type string
		common.NewPropDef("ageString", typeDefs[46], "", false, nil),
		// with type date
		common.NewPropDef("bornDate", typeDefs[41], "", false, nil),
		// with type Period
		common.NewPropDef("bornPeriod", typeDefs[25], "", false, nil),
		// with type string
		common.NewPropDef("bornString", typeDefs[46], "", false, nil),
		// with type FamilyMemberHistory_Condition
		common.NewPropDef("condition", typeDefs[301], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("dataAbsentReason", typeDefs[8], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type Age
		common.NewPropDef("deceasedAge", typeDefs[59], "", false, nil),
		// with type boolean
		common.NewPropDef("deceasedBoolean", typeDefs[40], "", false, nil),
		// with type date
		common.NewPropDef("deceasedDate", typeDefs[41], "", false, nil),
		// with type Range
		common.NewPropDef("deceasedRange", typeDefs[30], "", false, nil),
		// with type string
		common.NewPropDef("deceasedString", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("estimatedAge", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type canonical
		common.NewPropDef("instantiatesCanonical", typeDefs[673], "", true, nil),
		// with type uri
		common.NewPropDef("instantiatesUri", typeDefs[48], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type Reference
		common.NewPropDef("patient", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("reasonCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("reasonReference", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("relationship", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("sex", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"partial",
			"completed",
			"entered-in-error",
			"health-unknown",
		}),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of FamilyMemberHistory_Condition
	typeDefs[301].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type boolean
		common.NewPropDef("contributedToDeath", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type Age
		common.NewPropDef("onsetAge", typeDefs[59], "", false, nil),
		// with type Period
		common.NewPropDef("onsetPeriod", typeDefs[25], "", false, nil),
		// with type Range
		common.NewPropDef("onsetRange", typeDefs[30], "", false, nil),
		// with type string
		common.NewPropDef("onsetString", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("outcome", typeDefs[8], "", false, nil),
	})
	// properties of Flag
	typeDefs[302].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("author", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"active",
			"inactive",
			"entered-in-error",
		}),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of Goal
	typeDefs[303].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("achievementStatus", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("addresses", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("description", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("expressedBy", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("lifecycleStatus", typeDefs[46], "", false, []string{
			"proposed",
			"planned",
			"accepted",
			"active",
			"on-hold",
			"completed",
			"cancelled",
			"entered-in-error",
			"rejected",
		}),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("outcomeCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("outcomeReference", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("priority", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("startCodeableConcept", typeDefs[8], "", false, nil),
		// with type date
		common.NewPropDef("startDate", typeDefs[41], "", false, nil),
		// with type date
		common.NewPropDef("statusDate", typeDefs[41], "", false, nil),
		// with type string
		common.NewPropDef("statusReason", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Goal_Target
		common.NewPropDef("target", typeDefs[304], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of Goal_Target
	typeDefs[304].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("detailBoolean", typeDefs[40], "detail", false, nil),
		// with type CodeableConcept
		common.NewPropDef("detailCodeableConcept", typeDefs[8], "detail", false, nil),
		// with type integer
		common.NewPropDef("detailInteger", typeDefs[45], "detail", false, nil),
		// with type Quantity
		common.NewPropDef("detailQuantity", typeDefs[29], "detail", false, nil),
		// with type Range
		common.NewPropDef("detailRange", typeDefs[30], "detail", false, nil),
		// with type Ratio
		common.NewPropDef("detailRatio", typeDefs[31], "detail", false, nil),
		// with type string
		common.NewPropDef("detailString", typeDefs[46], "detail", false, nil),
		// with type date
		common.NewPropDef("dueDate", typeDefs[41], "", false, nil),
		// with type Duration
		common.NewPropDef("dueDuration", typeDefs[236], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("measure", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of GraphDefinition
	typeDefs[305].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type boolean
		common.NewPropDef("experimental", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type GraphDefinition_Link
		common.NewPropDef("link", typeDefs[307], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type canonical
		common.NewPropDef("profile", typeDefs[673], "", false, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("purpose", typeDefs[676], "", false, nil),
		// with type code
		common.NewPropDef("start", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of GraphDefinition_Compartment
	typeDefs[306].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("code", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("expression", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("rule", typeDefs[46], "", false, []string{
			"identical",
			"matching",
			"different",
			"custom",
		}),
		// with type string
		common.NewPropDef("use", typeDefs[46], "", false, []string{
			"condition",
			"requirement",
		}),
	})
	// properties of GraphDefinition_Link
	typeDefs[307].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("max", typeDefs[46], "", false, nil),
		// with type integer
		common.NewPropDef("min", typeDefs[45], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("path", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("sliceName", typeDefs[46], "", false, nil),
		// with type GraphDefinition_Target
		common.NewPropDef("target", typeDefs[308], "", true, nil),
	})
	// properties of GraphDefinition_Target
	typeDefs[308].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type GraphDefinition_Compartment
		common.NewPropDef("compartment", typeDefs[306], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type GraphDefinition_Link
		common.NewPropDef("link", typeDefs[307], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("params", typeDefs[46], "", false, nil),
		// with type canonical
		common.NewPropDef("profile", typeDefs[673], "", false, nil),
		// with type code
		common.NewPropDef("type", typeDefs[674], "", false, nil),
	})
	// properties of Group
	typeDefs[309].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("active", typeDefs[40], "", false, nil),
		// with type boolean
		common.NewPropDef("actual", typeDefs[40], "", false, nil),
		// with type Group_Characteristic
		common.NewPropDef("characteristic", typeDefs[310], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("managingEntity", typeDefs[32], "", false, nil),
		// with type Group_Member
		common.NewPropDef("member", typeDefs[311], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type unsignedInt
		common.NewPropDef("quantity", typeDefs[679], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"person",
			"animal",
			"practitioner",
			"device",
			"medication",
			"substance",
		}),
	})
	// properties of Group_Characteristic
	typeDefs[310].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type boolean
		common.NewPropDef("exclude", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type boolean
		common.NewPropDef("valueBoolean", typeDefs[40], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("valueCodeableConcept", typeDefs[8], "", false, nil),
		// with type Quantity
		common.NewPropDef("valueQuantity", typeDefs[29], "", false, nil),
		// with type Range
		common.NewPropDef("valueRange", typeDefs[30], "", false, nil),
		// with type Reference
		common.NewPropDef("valueReference", typeDefs[32], "", false, nil),
	})
	// properties of Group_Member
	typeDefs[311].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("entity", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("inactive", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
	})
	// properties of GuidanceResponse
	typeDefs[312].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type DataRequirement
		common.NewPropDef("dataRequirement", typeDefs[13], "", true, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("evaluationMessage", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type canonical
		common.NewPropDef("moduleCanonical", typeDefs[673], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("moduleCodeableConcept", typeDefs[8], "", false, nil),
		// with type uri
		common.NewPropDef("moduleUri", typeDefs[48], "", false, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type dateTime
		common.NewPropDef("occurrenceDateTime", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("outputParameters", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("performer", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("reasonCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("reasonReference", typeDefs[32], "", true, nil),
		// with type Identifier
		common.NewPropDef("requestIdentifier", typeDefs[18], "", false, nil),
		// with type Reference
		common.NewPropDef("result", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"success",
			"data-requested",
			"data-required",
			"in-progress",
			"failure",
			"entered-in-error",
		}),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of HealthcareService
	typeDefs[313].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("active", typeDefs[40], "", false, nil),
		// with type boolean
		common.NewPropDef("appointmentRequired", typeDefs[40], "", false, nil),
		// with type string
		common.NewPropDef("availabilityExceptions", typeDefs[46], "", false, nil),
		// with type HealthcareService_AvailableTime
		common.NewPropDef("availableTime", typeDefs[314], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("characteristic", typeDefs[8], "", true, nil),
		// with type string
		common.NewPropDef("comment", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("communication", typeDefs[8], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("coverageArea", typeDefs[32], "", true, nil),
		// with type HealthcareService_Eligibility
		common.NewPropDef("eligibility", typeDefs[315], "", true, nil),
		// with type Reference
		common.NewPropDef("endpoint", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type markdown
		common.NewPropDef("extraDetails", typeDefs[676], "", false, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("location", typeDefs[32], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type HealthcareService_NotAvailable
		common.NewPropDef("notAvailable", typeDefs[316], "", true, nil),
		// with type Attachment
		common.NewPropDef("photo", typeDefs[4], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("program", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("providedBy", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("referralMethod", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("serviceProvisionCode", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("specialty", typeDefs[8], "", true, nil),
		// with type ContactPoint
		common.NewPropDef("telecom", typeDefs[11], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", true, nil),
	})
	// properties of HealthcareService_AvailableTime
	typeDefs[314].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("allDay", typeDefs[40], "", false, nil),
		// with type time
		common.NewPropDef("availableEndTime", typeDefs[47], "", false, nil),
		// with type time
		common.NewPropDef("availableStartTime", typeDefs[47], "", false, nil),
		// with type string
		common.NewPropDef("daysOfWeek", typeDefs[46], "", true, []string{
			"mon",
			"tue",
			"wed",
			"thu",
			"fri",
			"sat",
			"sun",
		}),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of HealthcareService_Eligibility
	typeDefs[315].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type markdown
		common.NewPropDef("comment", typeDefs[676], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of HealthcareService_NotAvailable
	typeDefs[316].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Period
		common.NewPropDef("during", typeDefs[25], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of HumanName
	typeDefs[17].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("family", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("given", typeDefs[46], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type string
		common.NewPropDef("prefix", typeDefs[46], "", true, nil),
		// with type string
		common.NewPropDef("suffix", typeDefs[46], "", true, nil),
		// with type string
		common.NewPropDef("text", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("use", typeDefs[46], "", false, []string{
			"usual",
			"official",
			"temp",
			"nickname",
			"anonymous",
			"old",
			"maiden",
		}),
	})
	// properties of Identifier
	typeDefs[18].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("assigner", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type uri
		common.NewPropDef("system", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("use", typeDefs[46], "", false, []string{
			"usual",
			"official",
			"temp",
			"secondary",
			"old",
		}),
		// with type string
		common.NewPropDef("value", typeDefs[46], "", false, nil),
	})
	// properties of ImagingStudy
	typeDefs[317].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("basedOn", typeDefs[32], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("endpoint", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type Reference
		common.NewPropDef("interpreter", typeDefs[32], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("location", typeDefs[32], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Coding
		common.NewPropDef("modality", typeDefs[9], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type unsignedInt
		common.NewPropDef("numberOfInstances", typeDefs[679], "", false, nil),
		// with type unsignedInt
		common.NewPropDef("numberOfSeries", typeDefs[679], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("procedureCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("procedureReference", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("reasonCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("reasonReference", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("referrer", typeDefs[32], "", false, nil),
		// with type ImagingStudy_Series
		common.NewPropDef("series", typeDefs[320], "", true, nil),
		// with type dateTime
		common.NewPropDef("started", typeDefs[42], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"registered",
			"available",
			"cancelled",
			"entered-in-error",
			"unknown",
		}),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of ImagingStudy_Instance
	typeDefs[318].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type unsignedInt
		common.NewPropDef("number", typeDefs[679], "", false, nil),
		// with type Coding
		common.NewPropDef("sopClass", typeDefs[9], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type id
		common.NewPropDef("uid", typeDefs[675], "", false, nil),
	})
	// properties of ImagingStudy_Performer
	typeDefs[319].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("actor", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("function", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of ImagingStudy_Series
	typeDefs[320].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Coding
		common.NewPropDef("bodySite", typeDefs[9], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("endpoint", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type ImagingStudy_Instance
		common.NewPropDef("instance", typeDefs[318], "", true, nil),
		// with type Coding
		common.NewPropDef("laterality", typeDefs[9], "", false, nil),
		// with type Coding
		common.NewPropDef("modality", typeDefs[9], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type unsignedInt
		common.NewPropDef("number", typeDefs[679], "", false, nil),
		// with type unsignedInt
		common.NewPropDef("numberOfInstances", typeDefs[679], "", false, nil),
		// with type ImagingStudy_Performer
		common.NewPropDef("performer", typeDefs[319], "", true, nil),
		// with type Reference
		common.NewPropDef("specimen", typeDefs[32], "", true, nil),
		// with type dateTime
		common.NewPropDef("started", typeDefs[42], "", false, nil),
		// with type id
		common.NewPropDef("uid", typeDefs[675], "", false, nil),
	})
	// properties of Immunization
	typeDefs[321].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Quantity
		common.NewPropDef("doseQuantity", typeDefs[29], "", false, nil),
		// with type Immunization_Education
		common.NewPropDef("education", typeDefs[326], "", true, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type date
		common.NewPropDef("expirationDate", typeDefs[41], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("fundingSource", typeDefs[8], "", false, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type boolean
		common.NewPropDef("isSubpotent", typeDefs[40], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("location", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("lotNumber", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("manufacturer", typeDefs[32], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type dateTime
		common.NewPropDef("occurrenceDateTime", typeDefs[42], "", false, nil),
		// with type string
		common.NewPropDef("occurrenceString", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("patient", typeDefs[32], "", false, nil),
		// with type Immunization_Performer
		common.NewPropDef("performer", typeDefs[327], "", true, nil),
		// with type boolean
		common.NewPropDef("primarySource", typeDefs[40], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("programEligibility", typeDefs[8], "", true, nil),
		// with type Immunization_ProtocolApplied
		common.NewPropDef("protocolApplied", typeDefs[328], "", true, nil),
		// with type Immunization_Reaction
		common.NewPropDef("reaction", typeDefs[329], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("reasonCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("reasonReference", typeDefs[32], "", true, nil),
		// with type dateTime
		common.NewPropDef("recorded", typeDefs[42], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("reportOrigin", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("route", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("site", typeDefs[8], "", false, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("statusReason", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("subpotentReason", typeDefs[8], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("vaccineCode", typeDefs[8], "", false, nil),
	})
	// properties of ImmunizationEvaluation
	typeDefs[322].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("authority", typeDefs[32], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type positiveInt
		common.NewPropDef("doseNumberPositiveInt", typeDefs[678], "", false, nil),
		// with type string
		common.NewPropDef("doseNumberString", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("doseStatus", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("doseStatusReason", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type Reference
		common.NewPropDef("immunizationEvent", typeDefs[32], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("patient", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("series", typeDefs[46], "", false, nil),
		// with type positiveInt
		common.NewPropDef("seriesDosesPositiveInt", typeDefs[678], "", false, nil),
		// with type string
		common.NewPropDef("seriesDosesString", typeDefs[46], "", false, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("targetDisease", typeDefs[8], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of ImmunizationRecommendation
	typeDefs[323].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("authority", typeDefs[32], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("patient", typeDefs[32], "", false, nil),
		// with type ImmunizationRecommendation_Recommendation
		common.NewPropDef("recommendation", typeDefs[325], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of ImmunizationRecommendation_DateCriterion
	typeDefs[324].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type dateTime
		common.NewPropDef("value", typeDefs[42], "", false, nil),
	})
	// properties of ImmunizationRecommendation_Recommendation
	typeDefs[325].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("contraindicatedVaccineCode", typeDefs[8], "", true, nil),
		// with type ImmunizationRecommendation_DateCriterion
		common.NewPropDef("dateCriterion", typeDefs[324], "", true, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type positiveInt
		common.NewPropDef("doseNumberPositiveInt", typeDefs[678], "", false, nil),
		// with type string
		common.NewPropDef("doseNumberString", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("forecastReason", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("forecastStatus", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("series", typeDefs[46], "", false, nil),
		// with type positiveInt
		common.NewPropDef("seriesDosesPositiveInt", typeDefs[678], "", false, nil),
		// with type string
		common.NewPropDef("seriesDosesString", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("supportingImmunization", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("supportingPatientInformation", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("targetDisease", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("vaccineCode", typeDefs[8], "", true, nil),
	})
	// properties of Immunization_Education
	typeDefs[326].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("documentType", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type dateTime
		common.NewPropDef("presentationDate", typeDefs[42], "", false, nil),
		// with type dateTime
		common.NewPropDef("publicationDate", typeDefs[42], "", false, nil),
		// with type uri
		common.NewPropDef("reference", typeDefs[48], "", false, nil),
	})
	// properties of Immunization_Performer
	typeDefs[327].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("actor", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("function", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Immunization_ProtocolApplied
	typeDefs[328].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("authority", typeDefs[32], "", false, nil),
		// with type positiveInt
		common.NewPropDef("doseNumberPositiveInt", typeDefs[678], "", false, nil),
		// with type string
		common.NewPropDef("doseNumberString", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("series", typeDefs[46], "", false, nil),
		// with type positiveInt
		common.NewPropDef("seriesDosesPositiveInt", typeDefs[678], "", false, nil),
		// with type string
		common.NewPropDef("seriesDosesString", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("targetDisease", typeDefs[8], "", true, nil),
	})
	// properties of Immunization_Reaction
	typeDefs[329].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("detail", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("reported", typeDefs[40], "", false, nil),
	})
	// properties of ImplementationGuide
	typeDefs[330].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type ImplementationGuide_Definition
		common.NewPropDef("definition", typeDefs[331], "", false, nil),
		// with type ImplementationGuide_DependsOn
		common.NewPropDef("dependsOn", typeDefs[332], "", true, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type boolean
		common.NewPropDef("experimental", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("fhirVersion", typeDefs[46], "", true, []string{
			"0.01",
			"0.05",
			"0.06",
			"0.11",
			"0.0.80",
			"0.0.81",
			"0.0.82",
			"0.4.0",
			"0.5.0",
			"1.0.0",
			"1.0.1",
			"1.0.2",
			"1.1.0",
			"1.4.0",
			"1.6.0",
			"1.8.0",
			"3.0.0",
			"3.0.1",
			"3.3.0",
			"3.5.0",
			"4.0.0",
			"4.0.1",
		}),
		// with type ImplementationGuide_Global
		common.NewPropDef("global", typeDefs[333], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("license", typeDefs[46], "", false, []string{
			"not-open-source",
			"0BSD",
			"AAL",
			"Abstyles",
			"Adobe-2006",
			"Adobe-Glyph",
			"ADSL",
			"AFL-1.1",
			"AFL-1.2",
			"AFL-2.0",
			"AFL-2.1",
			"AFL-3.0",
			"Afmparse",
			"AGPL-1.0-only",
			"AGPL-1.0-or-later",
			"AGPL-3.0-only",
			"AGPL-3.0-or-later",
			"Aladdin",
			"AMDPLPA",
			"AML",
			"AMPAS",
			"ANTLR-PD",
			"Apache-1.0",
			"Apache-1.1",
			"Apache-2.0",
			"APAFML",
			"APL-1.0",
			"APSL-1.0",
			"APSL-1.1",
			"APSL-1.2",
			"APSL-2.0",
			"Artistic-1.0-cl8",
			"Artistic-1.0-Perl",
			"Artistic-1.0",
			"Artistic-2.0",
			"Bahyph",
			"Barr",
			"Beerware",
			"BitTorrent-1.0",
			"BitTorrent-1.1",
			"Borceux",
			"BSD-1-Clause",
			"BSD-2-Clause-FreeBSD",
			"BSD-2-Clause-NetBSD",
			"BSD-2-Clause-Patent",
			"BSD-2-Clause",
			"BSD-3-Clause-Attribution",
			"BSD-3-Clause-Clear",
			"BSD-3-Clause-LBNL",
			"BSD-3-Clause-No-Nuclear-License-2014",
			"BSD-3-Clause-No-Nuclear-License",
			"BSD-3-Clause-No-Nuclear-Warranty",
			"BSD-3-Clause",
			"BSD-4-Clause-UC",
			"BSD-4-Clause",
			"BSD-Protection",
			"BSD-Source-Code",
			"BSL-1.0",
			"bzip2-1.0.5",
			"bzip2-1.0.6",
			"Caldera",
			"CATOSL-1.1",
			"CC-BY-1.0",
			"CC-BY-2.0",
			"CC-BY-2.5",
			"CC-BY-3.0",
			"CC-BY-4.0",
			"CC-BY-NC-1.0",
			"CC-BY-NC-2.0",
			"CC-BY-NC-2.5",
			"CC-BY-NC-3.0",
			"CC-BY-NC-4.0",
			"CC-BY-NC-ND-1.0",
			"CC-BY-NC-ND-2.0",
			"CC-BY-NC-ND-2.5",
			"CC-BY-NC-ND-3.0",
			"CC-BY-NC-ND-4.0",
			"CC-BY-NC-SA-1.0",
			"CC-BY-NC-SA-2.0",
			"CC-BY-NC-SA-2.5",
			"CC-BY-NC-SA-3.0",
			"CC-BY-NC-SA-4.0",
			"CC-BY-ND-1.0",
			"CC-BY-ND-2.0",
			"CC-BY-ND-2.5",
			"CC-BY-ND-3.0",
			"CC-BY-ND-4.0",
			"CC-BY-SA-1.0",
			"CC-BY-SA-2.0",
			"CC-BY-SA-2.5",
			"CC-BY-SA-3.0",
			"CC-BY-SA-4.0",
			"CC0-1.0",
			"CDDL-1.0",
			"CDDL-1.1",
			"CDLA-Permissive-1.0",
			"CDLA-Sharing-1.0",
			"CECILL-1.0",
			"CECILL-1.1",
			"CECILL-2.0",
			"CECILL-2.1",
			"CECILL-B",
			"CECILL-C",
			"ClArtistic",
			"CNRI-Jython",
			"CNRI-Python-GPL-Compatible",
			"CNRI-Python",
			"Condor-1.1",
			"CPAL-1.0",
			"CPL-1.0",
			"CPOL-1.02",
			"Crossword",
			"CrystalStacker",
			"CUA-OPL-1.0",
			"Cube",
			"curl",
			"D-FSL-1.0",
			"diffmark",
			"DOC",
			"Dotseqn",
			"DSDP",
			"dvipdfm",
			"ECL-1.0",
			"ECL-2.0",
			"EFL-1.0",
			"EFL-2.0",
			"eGenix",
			"Entessa",
			"EPL-1.0",
			"EPL-2.0",
			"ErlPL-1.1",
			"EUDatagrid",
			"EUPL-1.0",
			"EUPL-1.1",
			"EUPL-1.2",
			"Eurosym",
			"Fair",
			"Frameworx-1.0",
			"FreeImage",
			"FSFAP",
			"FSFUL",
			"FSFULLR",
			"FTL",
			"GFDL-1.1-only",
			"GFDL-1.1-or-later",
			"GFDL-1.2-only",
			"GFDL-1.2-or-later",
			"GFDL-1.3-only",
			"GFDL-1.3-or-later",
			"Giftware",
			"GL2PS",
			"Glide",
			"Glulxe",
			"gnuplot",
			"GPL-1.0-only",
			"GPL-1.0-or-later",
			"GPL-2.0-only",
			"GPL-2.0-or-later",
			"GPL-3.0-only",
			"GPL-3.0-or-later",
			"gSOAP-1.3b",
			"HaskellReport",
			"HPND",
			"IBM-pibs",
			"ICU",
			"IJG",
			"ImageMagick",
			"iMatix",
			"Imlib2",
			"Info-ZIP",
			"Intel-ACPI",
			"Intel",
			"Interbase-1.0",
			"IPA",
			"IPL-1.0",
			"ISC",
			"JasPer-2.0",
			"JSON",
			"LAL-1.2",
			"LAL-1.3",
			"Latex2e",
			"Leptonica",
			"LGPL-2.0-only",
			"LGPL-2.0-or-later",
			"LGPL-2.1-only",
			"LGPL-2.1-or-later",
			"LGPL-3.0-only",
			"LGPL-3.0-or-later",
			"LGPLLR",
			"Libpng",
			"libtiff",
			"LiLiQ-P-1.1",
			"LiLiQ-R-1.1",
			"LiLiQ-Rplus-1.1",
			"Linux-OpenIB",
			"LPL-1.0",
			"LPL-1.02",
			"LPPL-1.0",
			"LPPL-1.1",
			"LPPL-1.2",
			"LPPL-1.3a",
			"LPPL-1.3c",
			"MakeIndex",
			"MirOS",
			"MIT-0",
			"MIT-advertising",
			"MIT-CMU",
			"MIT-enna",
			"MIT-feh",
			"MIT",
			"MITNFA",
			"Motosoto",
			"mpich2",
			"MPL-1.0",
			"MPL-1.1",
			"MPL-2.0-no-copyleft-exception",
			"MPL-2.0",
			"MS-PL",
			"MS-RL",
			"MTLL",
			"Multics",
			"Mup",
			"NASA-1.3",
			"Naumen",
			"NBPL-1.0",
			"NCSA",
			"Net-SNMP",
			"NetCDF",
			"Newsletr",
			"NGPL",
			"NLOD-1.0",
			"NLPL",
			"Nokia",
			"NOSL",
			"Noweb",
			"NPL-1.0",
			"NPL-1.1",
			"NPOSL-3.0",
			"NRL",
			"NTP",
			"OCCT-PL",
			"OCLC-2.0",
			"ODbL-1.0",
			"OFL-1.0",
			"OFL-1.1",
			"OGTSL",
			"OLDAP-1.1",
			"OLDAP-1.2",
			"OLDAP-1.3",
			"OLDAP-1.4",
			"OLDAP-2.0.1",
			"OLDAP-2.0",
			"OLDAP-2.1",
			"OLDAP-2.2.1",
			"OLDAP-2.2.2",
			"OLDAP-2.2",
			"OLDAP-2.3",
			"OLDAP-2.4",
			"OLDAP-2.5",
			"OLDAP-2.6",
			"OLDAP-2.7",
			"OLDAP-2.8",
			"OML",
			"OpenSSL",
			"OPL-1.0",
			"OSET-PL-2.1",
			"OSL-1.0",
			"OSL-1.1",
			"OSL-2.0",
			"OSL-2.1",
			"OSL-3.0",
			"PDDL-1.0",
			"PHP-3.0",
			"PHP-3.01",
			"Plexus",
			"PostgreSQL",
			"psfrag",
			"psutils",
			"Python-2.0",
			"Qhull",
			"QPL-1.0",
			"Rdisc",
			"RHeCos-1.1",
			"RPL-1.1",
			"RPL-1.5",
			"RPSL-1.0",
			"RSA-MD",
			"RSCPL",
			"Ruby",
			"SAX-PD",
			"Saxpath",
			"SCEA",
			"Sendmail",
			"SGI-B-1.0",
			"SGI-B-1.1",
			"SGI-B-2.0",
			"SimPL-2.0",
			"SISSL-1.2",
			"SISSL",
			"Sleepycat",
			"SMLNJ",
			"SMPPL",
			"SNIA",
			"Spencer-86",
			"Spencer-94",
			"Spencer-99",
			"SPL-1.0",
			"SugarCRM-1.1.3",
			"SWL",
			"TCL",
			"TCP-wrappers",
			"TMate",
			"TORQUE-1.1",
			"TOSL",
			"Unicode-DFS-2015",
			"Unicode-DFS-2016",
			"Unicode-TOU",
			"Unlicense",
			"UPL-1.0",
			"Vim",
			"VOSTROM",
			"VSL-1.0",
			"W3C-19980720",
			"W3C-20150513",
			"W3C",
			"Watcom-1.0",
			"Wsuipa",
			"WTFPL",
			"X11",
			"Xerox",
			"XFree86-1.1",
			"xinetd",
			"Xnet",
			"xpp",
			"XSkat",
			"YPL-1.0",
			"YPL-1.1",
			"Zed",
			"Zend-2.0",
			"Zimbra-1.3",
			"Zimbra-1.4",
			"zlib-acknowledgement",
			"Zlib",
			"ZPL-1.1",
			"ZPL-2.0",
			"ZPL-2.1",
		}),
		// with type ImplementationGuide_Manifest
		common.NewPropDef("manifest", typeDefs[335], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type id
		common.NewPropDef("packageId", typeDefs[675], "", false, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of ImplementationGuide_Definition
	typeDefs[331].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type ImplementationGuide_Grouping
		common.NewPropDef("grouping", typeDefs[334], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type ImplementationGuide_Page
		common.NewPropDef("page", typeDefs[336], "", false, nil),
		// with type ImplementationGuide_Parameter
		common.NewPropDef("parameter", typeDefs[338], "", true, nil),
		// with type ImplementationGuide_Resource
		common.NewPropDef("resource", typeDefs[339], "", true, nil),
		// with type ImplementationGuide_Template
		common.NewPropDef("template", typeDefs[341], "", true, nil),
	})
	// properties of ImplementationGuide_DependsOn
	typeDefs[332].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("packageId", typeDefs[675], "", false, nil),
		// with type canonical
		common.NewPropDef("uri", typeDefs[673], "", false, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of ImplementationGuide_Global
	typeDefs[333].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type canonical
		common.NewPropDef("profile", typeDefs[673], "", false, nil),
		// with type code
		common.NewPropDef("type", typeDefs[674], "", false, nil),
	})
	// properties of ImplementationGuide_Grouping
	typeDefs[334].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
	})
	// properties of ImplementationGuide_Manifest
	typeDefs[335].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("image", typeDefs[46], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("other", typeDefs[46], "", true, nil),
		// with type ImplementationGuide_Page1
		common.NewPropDef("page", typeDefs[337], "", true, nil),
		// with type url
		common.NewPropDef("rendering", typeDefs[680], "", false, nil),
		// with type ImplementationGuide_Resource1
		common.NewPropDef("resource", typeDefs[340], "", true, nil),
	})
	// properties of ImplementationGuide_Page
	typeDefs[336].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("generation", typeDefs[46], "", false, []string{
			"html",
			"markdown",
			"xml",
			"generated",
		}),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("nameReference", typeDefs[32], "", false, nil),
		// with type url
		common.NewPropDef("nameUrl", typeDefs[680], "", false, nil),
		// with type ImplementationGuide_Page
		common.NewPropDef("page", typeDefs[336], "", true, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
	})
	// properties of ImplementationGuide_Page1
	typeDefs[337].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("anchor", typeDefs[46], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
	})
	// properties of ImplementationGuide_Parameter
	typeDefs[338].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("code", typeDefs[46], "", false, []string{
			"apply",
			"path-resource",
			"path-pages",
			"path-tx-cache",
			"expansion-parameter",
			"rule-broken-links",
			"generate-xml",
			"generate-json",
			"generate-turtle",
			"html-template",
		}),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("value", typeDefs[46], "", false, nil),
	})
	// properties of ImplementationGuide_Resource
	typeDefs[339].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("exampleBoolean", typeDefs[40], "", false, nil),
		// with type canonical
		common.NewPropDef("exampleCanonical", typeDefs[673], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("fhirVersion", typeDefs[46], "", true, []string{
			"0.01",
			"0.05",
			"0.06",
			"0.11",
			"0.0.80",
			"0.0.81",
			"0.0.82",
			"0.4.0",
			"0.5.0",
			"1.0.0",
			"1.0.1",
			"1.0.2",
			"1.1.0",
			"1.4.0",
			"1.6.0",
			"1.8.0",
			"3.0.0",
			"3.0.1",
			"3.3.0",
			"3.5.0",
			"4.0.0",
			"4.0.1",
		}),
		// with type id
		common.NewPropDef("groupingId", typeDefs[675], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("reference", typeDefs[32], "", false, nil),
	})
	// properties of ImplementationGuide_Resource1
	typeDefs[340].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("exampleBoolean", typeDefs[40], "", false, nil),
		// with type canonical
		common.NewPropDef("exampleCanonical", typeDefs[673], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("reference", typeDefs[32], "", false, nil),
		// with type url
		common.NewPropDef("relativePath", typeDefs[680], "", false, nil),
	})
	// properties of ImplementationGuide_Template
	typeDefs[341].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("code", typeDefs[674], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("scope", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("source", typeDefs[46], "", false, nil),
	})
	// properties of InsurancePlan
	typeDefs[342].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("administeredBy", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("alias", typeDefs[46], "", true, nil),
		// with type InsurancePlan_Contact
		common.NewPropDef("contact", typeDefs[345], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type InsurancePlan_Coverage
		common.NewPropDef("coverage", typeDefs[347], "", true, nil),
		// with type Reference
		common.NewPropDef("coverageArea", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("endpoint", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("network", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("ownedBy", typeDefs[32], "", false, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type InsurancePlan_Plan
		common.NewPropDef("plan", typeDefs[350], "", true, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", true, nil),
	})
	// properties of InsurancePlan_Benefit
	typeDefs[343].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type InsurancePlan_Limit
		common.NewPropDef("limit", typeDefs[349], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("requirement", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of InsurancePlan_Benefit1
	typeDefs[344].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type InsurancePlan_Cost
		common.NewPropDef("cost", typeDefs[346], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of InsurancePlan_Contact
	typeDefs[345].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Address
		common.NewPropDef("address", typeDefs[2], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type HumanName
		common.NewPropDef("name", typeDefs[17], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("purpose", typeDefs[8], "", false, nil),
		// with type ContactPoint
		common.NewPropDef("telecom", typeDefs[11], "", true, nil),
	})
	// properties of InsurancePlan_Cost
	typeDefs[346].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("applicability", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("qualifiers", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type Quantity
		common.NewPropDef("value", typeDefs[29], "", false, nil),
	})
	// properties of InsurancePlan_Coverage
	typeDefs[347].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type InsurancePlan_Benefit
		common.NewPropDef("benefit", typeDefs[343], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("network", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of InsurancePlan_GeneralCost
	typeDefs[348].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("comment", typeDefs[46], "", false, nil),
		// with type Money
		common.NewPropDef("cost", typeDefs[21], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type positiveInt
		common.NewPropDef("groupSize", typeDefs[678], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of InsurancePlan_Limit
	typeDefs[349].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Quantity
		common.NewPropDef("value", typeDefs[29], "", false, nil),
	})
	// properties of InsurancePlan_Plan
	typeDefs[350].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("coverageArea", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type InsurancePlan_GeneralCost
		common.NewPropDef("generalCost", typeDefs[348], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("network", typeDefs[32], "", true, nil),
		// with type InsurancePlan_SpecificCost
		common.NewPropDef("specificCost", typeDefs[351], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of InsurancePlan_SpecificCost
	typeDefs[351].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type InsurancePlan_Benefit1
		common.NewPropDef("benefit", typeDefs[344], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Invoice
	typeDefs[352].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("account", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("cancelledReason", typeDefs[46], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type Reference
		common.NewPropDef("issuer", typeDefs[32], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Invoice_LineItem
		common.NewPropDef("lineItem", typeDefs[353], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type Invoice_Participant
		common.NewPropDef("participant", typeDefs[354], "", true, nil),
		// with type markdown
		common.NewPropDef("paymentTerms", typeDefs[676], "", false, nil),
		// with type Reference
		common.NewPropDef("recipient", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"issued",
			"balanced",
			"cancelled",
			"entered-in-error",
		}),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type Money
		common.NewPropDef("totalGross", typeDefs[21], "", false, nil),
		// with type Money
		common.NewPropDef("totalNet", typeDefs[21], "", false, nil),
		// with type Invoice_PriceComponent
		common.NewPropDef("totalPriceComponent", typeDefs[355], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of Invoice_LineItem
	typeDefs[353].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("chargeItemCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("chargeItemReference", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Invoice_PriceComponent
		common.NewPropDef("priceComponent", typeDefs[355], "", true, nil),
		// with type positiveInt
		common.NewPropDef("sequence", typeDefs[678], "", false, nil),
	})
	// properties of Invoice_Participant
	typeDefs[354].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("actor", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("role", typeDefs[8], "", false, nil),
	})
	// properties of Invoice_PriceComponent
	typeDefs[355].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Money
		common.NewPropDef("amount", typeDefs[21], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("factor", typeDefs[43], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"base",
			"surcharge",
			"deduction",
			"discount",
			"tax",
			"informational",
		}),
	})
	// properties of Library
	typeDefs[356].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type date
		common.NewPropDef("approvalDate", typeDefs[41], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("author", typeDefs[10], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Attachment
		common.NewPropDef("content", typeDefs[4], "", true, nil),
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type DataRequirement
		common.NewPropDef("dataRequirement", typeDefs[13], "", true, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("editor", typeDefs[10], "", true, nil),
		// with type Period
		common.NewPropDef("effectivePeriod", typeDefs[25], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("endorser", typeDefs[10], "", true, nil),
		// with type boolean
		common.NewPropDef("experimental", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type date
		common.NewPropDef("lastReviewDate", typeDefs[41], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type ParameterDefinition
		common.NewPropDef("parameter", typeDefs[23], "", true, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("purpose", typeDefs[676], "", false, nil),
		// with type RelatedArtifact
		common.NewPropDef("relatedArtifact", typeDefs[33], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("reviewer", typeDefs[10], "", true, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type CodeableConcept
		common.NewPropDef("subjectCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("subjectReference", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("subtitle", typeDefs[46], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("topic", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type string
		common.NewPropDef("usage", typeDefs[46], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of Linkage
	typeDefs[357].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("active", typeDefs[40], "", false, nil),
		// with type Reference
		common.NewPropDef("author", typeDefs[32], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type Linkage_Item
		common.NewPropDef("item", typeDefs[358], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of Linkage_Item
	typeDefs[358].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("resource", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"source",
			"alternate",
			"historical",
		}),
	})
	// properties of List
	typeDefs[359].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("emptyReason", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type List_Entry
		common.NewPropDef("entry", typeDefs[360], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type string
		common.NewPropDef("mode", typeDefs[46], "", false, []string{
			"working",
			"snapshot",
			"changes",
		}),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("orderedBy", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("source", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"current",
			"retired",
			"entered-in-error",
		}),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
	})
	// properties of List_Entry
	typeDefs[360].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type boolean
		common.NewPropDef("deleted", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("flag", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("item", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Location
	typeDefs[361].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Address
		common.NewPropDef("address", typeDefs[2], "", false, nil),
		// with type string
		common.NewPropDef("alias", typeDefs[46], "", true, nil),
		// with type string
		common.NewPropDef("availabilityExceptions", typeDefs[46], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("endpoint", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Location_HoursOfOperation
		common.NewPropDef("hoursOfOperation", typeDefs[362], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("managingOrganization", typeDefs[32], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type string
		common.NewPropDef("mode", typeDefs[46], "", false, []string{
			"instance",
			"kind",
		}),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type Coding
		common.NewPropDef("operationalStatus", typeDefs[9], "", false, nil),
		// with type Reference
		common.NewPropDef("partOf", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("physicalType", typeDefs[8], "", false, nil),
		// with type Location_Position
		common.NewPropDef("position", typeDefs[363], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"active",
			"suspended",
			"inactive",
		}),
		// with type ContactPoint
		common.NewPropDef("telecom", typeDefs[11], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", true, nil),
	})
	// properties of Location_HoursOfOperation
	typeDefs[362].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("allDay", typeDefs[40], "", false, nil),
		// with type time
		common.NewPropDef("closingTime", typeDefs[47], "", false, nil),
		// with type code
		common.NewPropDef("daysOfWeek", typeDefs[674], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type time
		common.NewPropDef("openingTime", typeDefs[47], "", false, nil),
	})
	// properties of Location_Position
	typeDefs[363].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type decimal
		common.NewPropDef("altitude", typeDefs[43], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type decimal
		common.NewPropDef("latitude", typeDefs[43], "", false, nil),
		// with type decimal
		common.NewPropDef("longitude", typeDefs[43], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of MarketingStatus
	typeDefs[19].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("country", typeDefs[8], "", false, nil),
		// with type Period
		common.NewPropDef("dateRange", typeDefs[25], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type dateTime
		common.NewPropDef("restoreDate", typeDefs[42], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("status", typeDefs[8], "", false, nil),
	})
	// properties of Measure
	typeDefs[364].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type date
		common.NewPropDef("approvalDate", typeDefs[41], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("author", typeDefs[10], "", true, nil),
		// with type markdown
		common.NewPropDef("clinicalRecommendationStatement", typeDefs[676], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("compositeScoring", typeDefs[8], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type markdown
		common.NewPropDef("definition", typeDefs[676], "", true, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type markdown
		common.NewPropDef("disclaimer", typeDefs[676], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("editor", typeDefs[10], "", true, nil),
		// with type Period
		common.NewPropDef("effectivePeriod", typeDefs[25], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("endorser", typeDefs[10], "", true, nil),
		// with type boolean
		common.NewPropDef("experimental", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Measure_Group
		common.NewPropDef("group", typeDefs[373], "", true, nil),
		// with type markdown
		common.NewPropDef("guidance", typeDefs[676], "", false, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("improvementNotation", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type date
		common.NewPropDef("lastReviewDate", typeDefs[41], "", false, nil),
		// with type canonical
		common.NewPropDef("library", typeDefs[673], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("purpose", typeDefs[676], "", false, nil),
		// with type string
		common.NewPropDef("rateAggregation", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("rationale", typeDefs[676], "", false, nil),
		// with type RelatedArtifact
		common.NewPropDef("relatedArtifact", typeDefs[33], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("reviewer", typeDefs[10], "", true, nil),
		// with type string
		common.NewPropDef("riskAdjustment", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("scoring", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type CodeableConcept
		common.NewPropDef("subjectCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("subjectReference", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("subtitle", typeDefs[46], "", false, nil),
		// with type Measure_SupplementalData
		common.NewPropDef("supplementalData", typeDefs[376], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("topic", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", true, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type string
		common.NewPropDef("usage", typeDefs[46], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of MeasureReport
	typeDefs[365].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("evaluatedResource", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type MeasureReport_Group
		common.NewPropDef("group", typeDefs[367], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("improvementNotation", typeDefs[8], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type canonical
		common.NewPropDef("measure", typeDefs[673], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type Reference
		common.NewPropDef("reporter", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"complete",
			"pending",
			"error",
		}),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"individual",
			"subject-list",
			"summary",
			"data-collection",
		}),
	})
	// properties of MeasureReport_Component
	typeDefs[366].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("value", typeDefs[8], "", false, nil),
	})
	// properties of MeasureReport_Group
	typeDefs[367].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Quantity
		common.NewPropDef("measureScore", typeDefs[29], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type MeasureReport_Population
		common.NewPropDef("population", typeDefs[368], "", true, nil),
		// with type MeasureReport_Stratifier
		common.NewPropDef("stratifier", typeDefs[370], "", true, nil),
	})
	// properties of MeasureReport_Population
	typeDefs[368].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type integer
		common.NewPropDef("count", typeDefs[45], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("subjectResults", typeDefs[32], "", false, nil),
	})
	// properties of MeasureReport_Population1
	typeDefs[369].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type integer
		common.NewPropDef("count", typeDefs[45], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("subjectResults", typeDefs[32], "", false, nil),
	})
	// properties of MeasureReport_Stratifier
	typeDefs[370].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type MeasureReport_Stratum
		common.NewPropDef("stratum", typeDefs[371], "", true, nil),
	})
	// properties of MeasureReport_Stratum
	typeDefs[371].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type MeasureReport_Component
		common.NewPropDef("component", typeDefs[366], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Quantity
		common.NewPropDef("measureScore", typeDefs[29], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type MeasureReport_Population1
		common.NewPropDef("population", typeDefs[369], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("value", typeDefs[8], "", false, nil),
	})
	// properties of Measure_Component
	typeDefs[372].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Expression
		common.NewPropDef("criteria", typeDefs[15], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Measure_Group
	typeDefs[373].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Measure_Population
		common.NewPropDef("population", typeDefs[374], "", true, nil),
		// with type Measure_Stratifier
		common.NewPropDef("stratifier", typeDefs[375], "", true, nil),
	})
	// properties of Measure_Population
	typeDefs[374].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Expression
		common.NewPropDef("criteria", typeDefs[15], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Measure_Stratifier
	typeDefs[375].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Measure_Component
		common.NewPropDef("component", typeDefs[372], "", true, nil),
		// with type Expression
		common.NewPropDef("criteria", typeDefs[15], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Measure_SupplementalData
	typeDefs[376].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Expression
		common.NewPropDef("criteria", typeDefs[15], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("usage", typeDefs[8], "", true, nil),
	})
	// properties of Media
	typeDefs[377].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("basedOn", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("bodySite", typeDefs[8], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Attachment
		common.NewPropDef("content", typeDefs[4], "", false, nil),
		// with type dateTime
		common.NewPropDef("createdDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("createdPeriod", typeDefs[25], "", false, nil),
		// with type Reference
		common.NewPropDef("device", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("deviceName", typeDefs[46], "", false, nil),
		// with type decimal
		common.NewPropDef("duration", typeDefs[43], "", false, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type positiveInt
		common.NewPropDef("frames", typeDefs[678], "", false, nil),
		// with type positiveInt
		common.NewPropDef("height", typeDefs[678], "", false, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type instant
		common.NewPropDef("issued", typeDefs[44], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("modality", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type Reference
		common.NewPropDef("operator", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("partOf", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("reasonCode", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("view", typeDefs[8], "", false, nil),
		// with type positiveInt
		common.NewPropDef("width", typeDefs[678], "", false, nil),
	})
	// properties of Medication
	typeDefs[378].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Ratio
		common.NewPropDef("amount", typeDefs[31], "", false, nil),
		// with type Medication_Batch
		common.NewPropDef("batch", typeDefs[407], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("form", typeDefs[8], "", false, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type Medication_Ingredient
		common.NewPropDef("ingredient", typeDefs[408], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("manufacturer", typeDefs[32], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of MedicationAdministration
	typeDefs[379].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("context", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("device", typeDefs[32], "", true, nil),
		// with type MedicationAdministration_Dosage
		common.NewPropDef("dosage", typeDefs[380], "", false, nil),
		// with type dateTime
		common.NewPropDef("effectiveDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("effectivePeriod", typeDefs[25], "", false, nil),
		// with type Reference
		common.NewPropDef("eventHistory", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type uri
		common.NewPropDef("instantiates", typeDefs[48], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("medicationCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("medicationReference", typeDefs[32], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type Reference
		common.NewPropDef("partOf", typeDefs[32], "", true, nil),
		// with type MedicationAdministration_Performer
		common.NewPropDef("performer", typeDefs[381], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("reasonCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("reasonReference", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("request", typeDefs[32], "", false, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("statusReason", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("supportingInformation", typeDefs[32], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of MedicationAdministration_Dosage
	typeDefs[380].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Quantity
		common.NewPropDef("dose", typeDefs[29], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("method", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Quantity
		common.NewPropDef("rateQuantity", typeDefs[29], "", false, nil),
		// with type Ratio
		common.NewPropDef("rateRatio", typeDefs[31], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("route", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("site", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("text", typeDefs[46], "", false, nil),
	})
	// properties of MedicationAdministration_Performer
	typeDefs[381].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("actor", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("function", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of MedicationDispense
	typeDefs[382].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("authorizingPrescription", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("context", typeDefs[32], "", false, nil),
		// with type Quantity
		common.NewPropDef("daysSupply", typeDefs[29], "", false, nil),
		// with type Reference
		common.NewPropDef("destination", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("detectedIssue", typeDefs[32], "", true, nil),
		// with type Dosage
		common.NewPropDef("dosageInstruction", typeDefs[234], "", true, nil),
		// with type Reference
		common.NewPropDef("eventHistory", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("location", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("medicationCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("medicationReference", typeDefs[32], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type Reference
		common.NewPropDef("partOf", typeDefs[32], "", true, nil),
		// with type MedicationDispense_Performer
		common.NewPropDef("performer", typeDefs[383], "", true, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type Reference
		common.NewPropDef("receiver", typeDefs[32], "", true, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("statusReasonCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("statusReasonReference", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type MedicationDispense_Substitution
		common.NewPropDef("substitution", typeDefs[384], "", false, nil),
		// with type Reference
		common.NewPropDef("supportingInformation", typeDefs[32], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type dateTime
		common.NewPropDef("whenHandedOver", typeDefs[42], "", false, nil),
		// with type dateTime
		common.NewPropDef("whenPrepared", typeDefs[42], "", false, nil),
	})
	// properties of MedicationDispense_Performer
	typeDefs[383].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("actor", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("function", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of MedicationDispense_Substitution
	typeDefs[384].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("reason", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("responsibleParty", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type boolean
		common.NewPropDef("wasSubstituted", typeDefs[40], "", false, nil),
	})
	// properties of MedicationKnowledge
	typeDefs[385].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type MedicationKnowledge_AdministrationGuidelines
		common.NewPropDef("administrationGuidelines", typeDefs[386], "", true, nil),
		// with type Quantity
		common.NewPropDef("amount", typeDefs[29], "", false, nil),
		// with type Reference
		common.NewPropDef("associatedMedication", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("contraindication", typeDefs[32], "", true, nil),
		// with type MedicationKnowledge_Cost
		common.NewPropDef("cost", typeDefs[387], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("doseForm", typeDefs[8], "", false, nil),
		// with type MedicationKnowledge_DrugCharacteristic
		common.NewPropDef("drugCharacteristic", typeDefs[389], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type MedicationKnowledge_Ingredient
		common.NewPropDef("ingredient", typeDefs[390], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("intendedRoute", typeDefs[8], "", true, nil),
		// with type MedicationKnowledge_Kinetics
		common.NewPropDef("kinetics", typeDefs[391], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("manufacturer", typeDefs[32], "", false, nil),
		// with type MedicationKnowledge_MedicineClassification
		common.NewPropDef("medicineClassification", typeDefs[393], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type MedicationKnowledge_MonitoringProgram
		common.NewPropDef("monitoringProgram", typeDefs[394], "", true, nil),
		// with type MedicationKnowledge_Monograph
		common.NewPropDef("monograph", typeDefs[395], "", true, nil),
		// with type MedicationKnowledge_Packaging
		common.NewPropDef("packaging", typeDefs[396], "", false, nil),
		// with type markdown
		common.NewPropDef("preparationInstruction", typeDefs[676], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("productType", typeDefs[8], "", true, nil),
		// with type MedicationKnowledge_Regulatory
		common.NewPropDef("regulatory", typeDefs[398], "", true, nil),
		// with type MedicationKnowledge_RelatedMedicationKnowledge
		common.NewPropDef("relatedMedicationKnowledge", typeDefs[399], "", true, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("synonym", typeDefs[46], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of MedicationKnowledge_AdministrationGuidelines
	typeDefs[386].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type MedicationKnowledge_Dosage
		common.NewPropDef("dosage", typeDefs[388], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("indicationCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("indicationReference", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type MedicationKnowledge_PatientCharacteristics
		common.NewPropDef("patientCharacteristics", typeDefs[397], "", true, nil),
	})
	// properties of MedicationKnowledge_Cost
	typeDefs[387].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Money
		common.NewPropDef("cost", typeDefs[21], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("source", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of MedicationKnowledge_Dosage
	typeDefs[388].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Dosage
		common.NewPropDef("dosage", typeDefs[234], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of MedicationKnowledge_DrugCharacteristic
	typeDefs[389].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type base64Binary
		common.NewPropDef("valueBase64Binary", typeDefs[39], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("valueCodeableConcept", typeDefs[8], "", false, nil),
		// with type Quantity
		common.NewPropDef("valueQuantity", typeDefs[29], "", false, nil),
		// with type string
		common.NewPropDef("valueString", typeDefs[46], "", false, nil),
	})
	// properties of MedicationKnowledge_Ingredient
	typeDefs[390].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("isActive", typeDefs[40], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("itemCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("itemReference", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Ratio
		common.NewPropDef("strength", typeDefs[31], "", false, nil),
	})
	// properties of MedicationKnowledge_Kinetics
	typeDefs[391].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Quantity
		common.NewPropDef("areaUnderCurve", typeDefs[29], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Duration
		common.NewPropDef("halfLifePeriod", typeDefs[236], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Quantity
		common.NewPropDef("lethalDose50", typeDefs[29], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of MedicationKnowledge_MaxDispense
	typeDefs[392].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Duration
		common.NewPropDef("period", typeDefs[236], "", false, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
	})
	// properties of MedicationKnowledge_MedicineClassification
	typeDefs[393].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("classification", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of MedicationKnowledge_MonitoringProgram
	typeDefs[394].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of MedicationKnowledge_Monograph
	typeDefs[395].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("source", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of MedicationKnowledge_Packaging
	typeDefs[396].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of MedicationKnowledge_PatientCharacteristics
	typeDefs[397].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("characteristicCodeableConcept", typeDefs[8], "", false, nil),
		// with type Quantity
		common.NewPropDef("characteristicQuantity", typeDefs[29], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("value", typeDefs[46], "", true, nil),
	})
	// properties of MedicationKnowledge_Regulatory
	typeDefs[398].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type MedicationKnowledge_MaxDispense
		common.NewPropDef("maxDispense", typeDefs[392], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("regulatoryAuthority", typeDefs[32], "", false, nil),
		// with type MedicationKnowledge_Schedule
		common.NewPropDef("schedule", typeDefs[400], "", true, nil),
		// with type MedicationKnowledge_Substitution
		common.NewPropDef("substitution", typeDefs[401], "", true, nil),
	})
	// properties of MedicationKnowledge_RelatedMedicationKnowledge
	typeDefs[399].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("reference", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of MedicationKnowledge_Schedule
	typeDefs[400].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("schedule", typeDefs[8], "", false, nil),
	})
	// properties of MedicationKnowledge_Substitution
	typeDefs[401].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("allowed", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of MedicationRequest
	typeDefs[402].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type dateTime
		common.NewPropDef("authoredOn", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("basedOn", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("courseOfTherapyType", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("detectedIssue", typeDefs[32], "", true, nil),
		// with type MedicationRequest_DispenseRequest
		common.NewPropDef("dispenseRequest", typeDefs[403], "", false, nil),
		// with type boolean
		common.NewPropDef("doNotPerform", typeDefs[40], "", false, nil),
		// with type Dosage
		common.NewPropDef("dosageInstruction", typeDefs[234], "", true, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("eventHistory", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Identifier
		common.NewPropDef("groupIdentifier", typeDefs[18], "", false, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type canonical
		common.NewPropDef("instantiatesCanonical", typeDefs[673], "", true, nil),
		// with type uri
		common.NewPropDef("instantiatesUri", typeDefs[48], "", true, nil),
		// with type Reference
		common.NewPropDef("insurance", typeDefs[32], "", true, nil),
		// with type code
		common.NewPropDef("intent", typeDefs[674], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("medicationCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("medicationReference", typeDefs[32], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type Reference
		common.NewPropDef("performer", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("performerType", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("priorPrescription", typeDefs[32], "", false, nil),
		// with type code
		common.NewPropDef("priority", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("reasonCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("reasonReference", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("recorder", typeDefs[32], "", false, nil),
		// with type boolean
		common.NewPropDef("reportedBoolean", typeDefs[40], "", false, nil),
		// with type Reference
		common.NewPropDef("reportedReference", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("requester", typeDefs[32], "", false, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("statusReason", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type MedicationRequest_Substitution
		common.NewPropDef("substitution", typeDefs[405], "", false, nil),
		// with type Reference
		common.NewPropDef("supportingInformation", typeDefs[32], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of MedicationRequest_DispenseRequest
	typeDefs[403].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Duration
		common.NewPropDef("dispenseInterval", typeDefs[236], "", false, nil),
		// with type Duration
		common.NewPropDef("expectedSupplyDuration", typeDefs[236], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type MedicationRequest_InitialFill
		common.NewPropDef("initialFill", typeDefs[404], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type unsignedInt
		common.NewPropDef("numberOfRepeatsAllowed", typeDefs[679], "", false, nil),
		// with type Reference
		common.NewPropDef("performer", typeDefs[32], "", false, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type Period
		common.NewPropDef("validityPeriod", typeDefs[25], "", false, nil),
	})
	// properties of MedicationRequest_InitialFill
	typeDefs[404].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Duration
		common.NewPropDef("duration", typeDefs[236], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
	})
	// properties of MedicationRequest_Substitution
	typeDefs[405].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("allowedBoolean", typeDefs[40], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("allowedCodeableConcept", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("reason", typeDefs[8], "", false, nil),
	})
	// properties of MedicationStatement
	typeDefs[406].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("basedOn", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("context", typeDefs[32], "", false, nil),
		// with type dateTime
		common.NewPropDef("dateAsserted", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("derivedFrom", typeDefs[32], "", true, nil),
		// with type Dosage
		common.NewPropDef("dosage", typeDefs[234], "", true, nil),
		// with type dateTime
		common.NewPropDef("effectiveDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("effectivePeriod", typeDefs[25], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type Reference
		common.NewPropDef("informationSource", typeDefs[32], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("medicationCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("medicationReference", typeDefs[32], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type Reference
		common.NewPropDef("partOf", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("reasonCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("reasonReference", typeDefs[32], "", true, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("statusReason", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of Medication_Batch
	typeDefs[407].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type dateTime
		common.NewPropDef("expirationDate", typeDefs[42], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("lotNumber", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Medication_Ingredient
	typeDefs[408].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("isActive", typeDefs[40], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("itemCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("itemReference", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Ratio
		common.NewPropDef("strength", typeDefs[31], "", false, nil),
	})
	// properties of MedicinalProduct
	typeDefs[409].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("additionalMonitoringIndicator", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("attachedDocument", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("clinicalTrial", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("combinedPharmaceuticalDoseForm", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("contact", typeDefs[32], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Identifier
		common.NewPropDef("crossReference", typeDefs[18], "", true, nil),
		// with type Coding
		common.NewPropDef("domain", typeDefs[9], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("legalStatusOfSupply", typeDefs[8], "", false, nil),
		// with type MedicinalProduct_ManufacturingBusinessOperation
		common.NewPropDef("manufacturingBusinessOperation", typeDefs[435], "", true, nil),
		// with type MarketingStatus
		common.NewPropDef("marketingStatus", typeDefs[19], "", true, nil),
		// with type Reference
		common.NewPropDef("masterFile", typeDefs[32], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type MedicinalProduct_Name
		common.NewPropDef("name", typeDefs[436], "", true, nil),
		// with type Reference
		common.NewPropDef("packagedMedicinalProduct", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("paediatricUseIndicator", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("pharmaceuticalProduct", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("productClassification", typeDefs[8], "", true, nil),
		// with type MedicinalProduct_SpecialDesignation
		common.NewPropDef("specialDesignation", typeDefs[438], "", true, nil),
		// with type string
		common.NewPropDef("specialMeasures", typeDefs[46], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of MedicinalProductAuthorization
	typeDefs[410].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("country", typeDefs[8], "", true, nil),
		// with type Period
		common.NewPropDef("dataExclusivityPeriod", typeDefs[25], "", false, nil),
		// with type dateTime
		common.NewPropDef("dateOfFirstAuthorization", typeDefs[42], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("holder", typeDefs[32], "", false, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type dateTime
		common.NewPropDef("internationalBirthDate", typeDefs[42], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type MedicinalProductAuthorization_JurisdictionalAuthorization
		common.NewPropDef("jurisdictionalAuthorization", typeDefs[411], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("legalBasis", typeDefs[8], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type MedicinalProductAuthorization_Procedure
		common.NewPropDef("procedure", typeDefs[412], "", false, nil),
		// with type Reference
		common.NewPropDef("regulator", typeDefs[32], "", false, nil),
		// with type dateTime
		common.NewPropDef("restoreDate", typeDefs[42], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("status", typeDefs[8], "", false, nil),
		// with type dateTime
		common.NewPropDef("statusDate", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type Period
		common.NewPropDef("validityPeriod", typeDefs[25], "", false, nil),
	})
	// properties of MedicinalProductAuthorization_JurisdictionalAuthorization
	typeDefs[411].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("country", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("legalStatusOfSupply", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Period
		common.NewPropDef("validityPeriod", typeDefs[25], "", false, nil),
	})
	// properties of MedicinalProductAuthorization_Procedure
	typeDefs[412].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type MedicinalProductAuthorization_Procedure
		common.NewPropDef("application", typeDefs[412], "", true, nil),
		// with type dateTime
		common.NewPropDef("dateDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("datePeriod", typeDefs[25], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of MedicinalProductContraindication
	typeDefs[413].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("comorbidity", typeDefs[8], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("disease", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("diseaseStatus", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type MedicinalProductContraindication_OtherTherapy
		common.NewPropDef("otherTherapy", typeDefs[414], "", true, nil),
		// with type Population
		common.NewPropDef("population", typeDefs[26], "", true, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type Reference
		common.NewPropDef("therapeuticIndication", typeDefs[32], "", true, nil),
	})
	// properties of MedicinalProductContraindication_OtherTherapy
	typeDefs[414].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("medicationCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("medicationReference", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("therapyRelationshipType", typeDefs[8], "", false, nil),
	})
	// properties of MedicinalProductIndication
	typeDefs[415].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("comorbidity", typeDefs[8], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("diseaseStatus", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("diseaseSymptomProcedure", typeDefs[8], "", false, nil),
		// with type Quantity
		common.NewPropDef("duration", typeDefs[29], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("intendedEffect", typeDefs[8], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type MedicinalProductIndication_OtherTherapy
		common.NewPropDef("otherTherapy", typeDefs[416], "", true, nil),
		// with type Population
		common.NewPropDef("population", typeDefs[26], "", true, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type Reference
		common.NewPropDef("undesirableEffect", typeDefs[32], "", true, nil),
	})
	// properties of MedicinalProductIndication_OtherTherapy
	typeDefs[416].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("medicationCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("medicationReference", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("therapyRelationshipType", typeDefs[8], "", false, nil),
	})
	// properties of MedicinalProductIngredient
	typeDefs[417].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("allergenicIndicator", typeDefs[40], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("manufacturer", typeDefs[32], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("role", typeDefs[8], "", false, nil),
		// with type MedicinalProductIngredient_SpecifiedSubstance
		common.NewPropDef("specifiedSubstance", typeDefs[419], "", true, nil),
		// with type MedicinalProductIngredient_Substance
		common.NewPropDef("substance", typeDefs[421], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of MedicinalProductIngredient_ReferenceStrength
	typeDefs[418].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("country", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("measurementPoint", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Ratio
		common.NewPropDef("strength", typeDefs[31], "", false, nil),
		// with type Ratio
		common.NewPropDef("strengthLowLimit", typeDefs[31], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("substance", typeDefs[8], "", false, nil),
	})
	// properties of MedicinalProductIngredient_SpecifiedSubstance
	typeDefs[419].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("confidentiality", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("group", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type MedicinalProductIngredient_Strength
		common.NewPropDef("strength", typeDefs[420], "", true, nil),
	})
	// properties of MedicinalProductIngredient_Strength
	typeDefs[420].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Ratio
		common.NewPropDef("concentration", typeDefs[31], "", false, nil),
		// with type Ratio
		common.NewPropDef("concentrationLowLimit", typeDefs[31], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("country", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("measurementPoint", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Ratio
		common.NewPropDef("presentation", typeDefs[31], "", false, nil),
		// with type Ratio
		common.NewPropDef("presentationLowLimit", typeDefs[31], "", false, nil),
		// with type MedicinalProductIngredient_ReferenceStrength
		common.NewPropDef("referenceStrength", typeDefs[418], "", true, nil),
	})
	// properties of MedicinalProductIngredient_Substance
	typeDefs[421].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type MedicinalProductIngredient_Strength
		common.NewPropDef("strength", typeDefs[420], "", true, nil),
	})
	// properties of MedicinalProductInteraction
	typeDefs[422].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("effect", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("incidence", typeDefs[8], "", false, nil),
		// with type MedicinalProductInteraction_Interactant
		common.NewPropDef("interactant", typeDefs[423], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("management", typeDefs[8], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of MedicinalProductInteraction_Interactant
	typeDefs[423].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("itemCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("itemReference", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of MedicinalProductManufactured
	typeDefs[424].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type Reference
		common.NewPropDef("ingredient", typeDefs[32], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("manufacturedDoseForm", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("manufacturer", typeDefs[32], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("otherCharacteristics", typeDefs[8], "", true, nil),
		// with type ProdCharacteristic
		common.NewPropDef("physicalCharacteristics", typeDefs[27], "", false, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("unitOfPresentation", typeDefs[8], "", false, nil),
	})
	// properties of MedicinalProductPackaged
	typeDefs[425].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type MedicinalProductPackaged_BatchIdentifier
		common.NewPropDef("batchIdentifier", typeDefs[426], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("legalStatusOfSupply", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("manufacturer", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("marketingAuthorization", typeDefs[32], "", false, nil),
		// with type MarketingStatus
		common.NewPropDef("marketingStatus", typeDefs[19], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type MedicinalProductPackaged_PackageItem
		common.NewPropDef("packageItem", typeDefs[427], "", true, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of MedicinalProductPackaged_BatchIdentifier
	typeDefs[426].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("immediatePackaging", typeDefs[18], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Identifier
		common.NewPropDef("outerPackaging", typeDefs[18], "", false, nil),
	})
	// properties of MedicinalProductPackaged_PackageItem
	typeDefs[427].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("alternateMaterial", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("device", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type Reference
		common.NewPropDef("manufacturedItem", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("manufacturer", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("material", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("otherCharacteristics", typeDefs[8], "", true, nil),
		// with type MedicinalProductPackaged_PackageItem
		common.NewPropDef("packageItem", typeDefs[427], "", true, nil),
		// with type ProdCharacteristic
		common.NewPropDef("physicalCharacteristics", typeDefs[27], "", false, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type ProductShelfLife
		common.NewPropDef("shelfLifeStorage", typeDefs[28], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of MedicinalProductPharmaceutical
	typeDefs[428].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("administrableDoseForm", typeDefs[8], "", false, nil),
		// with type MedicinalProductPharmaceutical_Characteristics
		common.NewPropDef("characteristics", typeDefs[429], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("device", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type Reference
		common.NewPropDef("ingredient", typeDefs[32], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type MedicinalProductPharmaceutical_RouteOfAdministration
		common.NewPropDef("routeOfAdministration", typeDefs[430], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("unitOfPresentation", typeDefs[8], "", false, nil),
	})
	// properties of MedicinalProductPharmaceutical_Characteristics
	typeDefs[429].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("status", typeDefs[8], "", false, nil),
	})
	// properties of MedicinalProductPharmaceutical_RouteOfAdministration
	typeDefs[430].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Quantity
		common.NewPropDef("firstDose", typeDefs[29], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Quantity
		common.NewPropDef("maxDosePerDay", typeDefs[29], "", false, nil),
		// with type Ratio
		common.NewPropDef("maxDosePerTreatmentPeriod", typeDefs[31], "", false, nil),
		// with type Quantity
		common.NewPropDef("maxSingleDose", typeDefs[29], "", false, nil),
		// with type Duration
		common.NewPropDef("maxTreatmentPeriod", typeDefs[236], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type MedicinalProductPharmaceutical_TargetSpecies
		common.NewPropDef("targetSpecies", typeDefs[431], "", true, nil),
	})
	// properties of MedicinalProductPharmaceutical_TargetSpecies
	typeDefs[431].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type MedicinalProductPharmaceutical_WithdrawalPeriod
		common.NewPropDef("withdrawalPeriod", typeDefs[432], "", true, nil),
	})
	// properties of MedicinalProductPharmaceutical_WithdrawalPeriod
	typeDefs[432].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("supportingInformation", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("tissue", typeDefs[8], "", false, nil),
		// with type Quantity
		common.NewPropDef("value", typeDefs[29], "", false, nil),
	})
	// properties of MedicinalProductUndesirableEffect
	typeDefs[433].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("classification", typeDefs[8], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("frequencyOfOccurrence", typeDefs[8], "", false, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Population
		common.NewPropDef("population", typeDefs[26], "", true, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("symptomConditionEffect", typeDefs[8], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of MedicinalProduct_CountryLanguage
	typeDefs[434].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("country", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("language", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of MedicinalProduct_ManufacturingBusinessOperation
	typeDefs[435].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Identifier
		common.NewPropDef("authorisationReferenceNumber", typeDefs[18], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("confidentialityIndicator", typeDefs[8], "", false, nil),
		// with type dateTime
		common.NewPropDef("effectiveDate", typeDefs[42], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("manufacturer", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("operationType", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("regulator", typeDefs[32], "", false, nil),
	})
	// properties of MedicinalProduct_Name
	typeDefs[436].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type MedicinalProduct_CountryLanguage
		common.NewPropDef("countryLanguage", typeDefs[434], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type MedicinalProduct_NamePart
		common.NewPropDef("namePart", typeDefs[437], "", true, nil),
		// with type string
		common.NewPropDef("productName", typeDefs[46], "", false, nil),
	})
	// properties of MedicinalProduct_NamePart
	typeDefs[437].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("part", typeDefs[46], "", false, nil),
		// with type Coding
		common.NewPropDef("type", typeDefs[9], "", false, nil),
	})
	// properties of MedicinalProduct_SpecialDesignation
	typeDefs[438].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("indicationCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("indicationReference", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("intendedUse", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("species", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("status", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of MessageDefinition
	typeDefs[439].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type MessageDefinition_AllowedResponse
		common.NewPropDef("allowedResponse", typeDefs[440], "", true, nil),
		// with type canonical
		common.NewPropDef("base", typeDefs[673], "", false, nil),
		// with type string
		common.NewPropDef("category", typeDefs[46], "", false, []string{
			"consequence",
			"currency",
			"notification",
		}),
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type Coding
		common.NewPropDef("eventCoding", typeDefs[9], "", false, nil),
		// with type uri
		common.NewPropDef("eventUri", typeDefs[48], "", false, nil),
		// with type boolean
		common.NewPropDef("experimental", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type MessageDefinition_Focus
		common.NewPropDef("focus", typeDefs[441], "", true, nil),
		// with type canonical
		common.NewPropDef("graph", typeDefs[673], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type canonical
		common.NewPropDef("parent", typeDefs[673], "", true, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("purpose", typeDefs[676], "", false, nil),
		// with type canonical
		common.NewPropDef("replaces", typeDefs[673], "", true, nil),
		// with type string
		common.NewPropDef("responseRequired", typeDefs[46], "", false, []string{
			"always",
			"on-error",
			"never",
			"on-success",
		}),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of MessageDefinition_AllowedResponse
	typeDefs[440].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type canonical
		common.NewPropDef("message", typeDefs[673], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type markdown
		common.NewPropDef("situation", typeDefs[676], "", false, nil),
	})
	// properties of MessageDefinition_Focus
	typeDefs[441].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("code", typeDefs[674], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("max", typeDefs[46], "", false, nil),
		// with type unsignedInt
		common.NewPropDef("min", typeDefs[679], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type canonical
		common.NewPropDef("profile", typeDefs[673], "", false, nil),
	})
	// properties of MessageHeader
	typeDefs[442].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("author", typeDefs[32], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type canonical
		common.NewPropDef("definition", typeDefs[673], "", false, nil),
		// with type MessageHeader_Destination
		common.NewPropDef("destination", typeDefs[443], "", true, nil),
		// with type Reference
		common.NewPropDef("enterer", typeDefs[32], "", false, nil),
		// with type Coding
		common.NewPropDef("eventCoding", typeDefs[9], "", false, nil),
		// with type uri
		common.NewPropDef("eventUri", typeDefs[48], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("focus", typeDefs[32], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("reason", typeDefs[8], "", false, nil),
		// with type MessageHeader_Response
		common.NewPropDef("response", typeDefs[444], "", false, nil),
		// with type Reference
		common.NewPropDef("responsible", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("sender", typeDefs[32], "", false, nil),
		// with type MessageHeader_Source
		common.NewPropDef("source", typeDefs[445], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of MessageHeader_Destination
	typeDefs[443].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type url
		common.NewPropDef("endpoint", typeDefs[680], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("receiver", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("target", typeDefs[32], "", false, nil),
	})
	// properties of MessageHeader_Response
	typeDefs[444].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("code", typeDefs[46], "", false, []string{
			"ok",
			"transient-error",
			"fatal-error",
		}),
		// with type Reference
		common.NewPropDef("details", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type id
		common.NewPropDef("identifier", typeDefs[675], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of MessageHeader_Source
	typeDefs[445].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ContactPoint
		common.NewPropDef("contact", typeDefs[11], "", false, nil),
		// with type url
		common.NewPropDef("endpoint", typeDefs[680], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("software", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of Meta
	typeDefs[20].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type instant
		common.NewPropDef("lastUpdated", typeDefs[44], "", false, nil),
		// with type canonical
		common.NewPropDef("profile", typeDefs[673], "", true, nil),
		// with type Coding
		common.NewPropDef("security", typeDefs[9], "", true, nil),
		// with type uri
		common.NewPropDef("source", typeDefs[48], "", false, nil),
		// with type Coding
		common.NewPropDef("tag", typeDefs[9], "", true, nil),
		// with type id
		common.NewPropDef("versionId", typeDefs[675], "", false, nil),
	})
	// properties of MolecularSequence
	typeDefs[446].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type integer
		common.NewPropDef("coordinateSystem", typeDefs[45], "", false, nil),
		// with type Reference
		common.NewPropDef("device", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("observedSeq", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("patient", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("performer", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("pointer", typeDefs[32], "", true, nil),
		// with type MolecularSequence_Quality
		common.NewPropDef("quality", typeDefs[449], "", true, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type integer
		common.NewPropDef("readCoverage", typeDefs[45], "", false, nil),
		// with type MolecularSequence_ReferenceSeq
		common.NewPropDef("referenceSeq", typeDefs[450], "", false, nil),
		// with type MolecularSequence_Repository
		common.NewPropDef("repository", typeDefs[451], "", true, nil),
		// with type Reference
		common.NewPropDef("specimen", typeDefs[32], "", false, nil),
		// with type MolecularSequence_StructureVariant
		common.NewPropDef("structureVariant", typeDefs[453], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"aa",
			"dna",
			"rna",
		}),
		// with type MolecularSequence_Variant
		common.NewPropDef("variant", typeDefs[454], "", true, nil),
	})
	// properties of MolecularSequence_Inner
	typeDefs[447].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type integer
		common.NewPropDef("end", typeDefs[45], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type integer
		common.NewPropDef("start", typeDefs[45], "", false, nil),
	})
	// properties of MolecularSequence_Outer
	typeDefs[448].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type integer
		common.NewPropDef("end", typeDefs[45], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type integer
		common.NewPropDef("start", typeDefs[45], "", false, nil),
	})
	// properties of MolecularSequence_Quality
	typeDefs[449].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type integer
		common.NewPropDef("end", typeDefs[45], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("fScore", typeDefs[43], "", false, nil),
		// with type decimal
		common.NewPropDef("gtFP", typeDefs[43], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("method", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("precision", typeDefs[43], "", false, nil),
		// with type decimal
		common.NewPropDef("queryFP", typeDefs[43], "", false, nil),
		// with type decimal
		common.NewPropDef("queryTP", typeDefs[43], "", false, nil),
		// with type decimal
		common.NewPropDef("recall", typeDefs[43], "", false, nil),
		// with type MolecularSequence_Roc
		common.NewPropDef("roc", typeDefs[452], "", false, nil),
		// with type Quantity
		common.NewPropDef("score", typeDefs[29], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("standardSequence", typeDefs[8], "", false, nil),
		// with type integer
		common.NewPropDef("start", typeDefs[45], "", false, nil),
		// with type decimal
		common.NewPropDef("truthFN", typeDefs[43], "", false, nil),
		// with type decimal
		common.NewPropDef("truthTP", typeDefs[43], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"indel",
			"snp",
			"unknown",
		}),
	})
	// properties of MolecularSequence_ReferenceSeq
	typeDefs[450].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("chromosome", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("genomeBuild", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("orientation", typeDefs[46], "", false, []string{
			"sense",
			"antisense",
		}),
		// with type CodeableConcept
		common.NewPropDef("referenceSeqId", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("referenceSeqPointer", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("referenceSeqString", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("strand", typeDefs[46], "", false, []string{
			"watson",
			"crick",
		}),
		// with type integer
		common.NewPropDef("windowEnd", typeDefs[45], "", false, nil),
		// with type integer
		common.NewPropDef("windowStart", typeDefs[45], "", false, nil),
	})
	// properties of MolecularSequence_Repository
	typeDefs[451].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("datasetId", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("readsetId", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"directlink",
			"openapi",
			"login",
			"oauth",
			"other",
		}),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type string
		common.NewPropDef("variantsetId", typeDefs[46], "", false, nil),
	})
	// properties of MolecularSequence_Roc
	typeDefs[452].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("fMeasure", typeDefs[43], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type integer
		common.NewPropDef("numFN", typeDefs[45], "", true, nil),
		// with type integer
		common.NewPropDef("numFP", typeDefs[45], "", true, nil),
		// with type integer
		common.NewPropDef("numTP", typeDefs[45], "", true, nil),
		// with type decimal
		common.NewPropDef("precision", typeDefs[43], "", true, nil),
		// with type integer
		common.NewPropDef("score", typeDefs[45], "", true, nil),
		// with type decimal
		common.NewPropDef("sensitivity", typeDefs[43], "", true, nil),
	})
	// properties of MolecularSequence_StructureVariant
	typeDefs[453].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("exact", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type MolecularSequence_Inner
		common.NewPropDef("inner", typeDefs[447], "", false, nil),
		// with type integer
		common.NewPropDef("length", typeDefs[45], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type MolecularSequence_Outer
		common.NewPropDef("outer", typeDefs[448], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("variantType", typeDefs[8], "", false, nil),
	})
	// properties of MolecularSequence_Variant
	typeDefs[454].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("cigar", typeDefs[46], "", false, nil),
		// with type integer
		common.NewPropDef("end", typeDefs[45], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("observedAllele", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("referenceAllele", typeDefs[46], "", false, nil),
		// with type integer
		common.NewPropDef("start", typeDefs[45], "", false, nil),
		// with type Reference
		common.NewPropDef("variantPointer", typeDefs[32], "", false, nil),
	})
	// properties of Money
	typeDefs[21].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("currency", typeDefs[674], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type decimal
		common.NewPropDef("value", typeDefs[43], "", false, nil),
	})
	// properties of NamingSystem
	typeDefs[455].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type string
		common.NewPropDef("kind", typeDefs[46], "", false, []string{
			"codesystem",
			"identifier",
			"root",
		}),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("responsible", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type NamingSystem_UniqueId
		common.NewPropDef("uniqueId", typeDefs[456], "", true, nil),
		// with type string
		common.NewPropDef("usage", typeDefs[46], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
	})
	// properties of NamingSystem_UniqueId
	typeDefs[456].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("comment", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type boolean
		common.NewPropDef("preferred", typeDefs[40], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"oid",
			"uuid",
			"uri",
			"other",
		}),
		// with type string
		common.NewPropDef("value", typeDefs[46], "", false, nil),
	})
	// properties of Narrative
	typeDefs[22].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type xhtml
		common.NewPropDef("div", typeDefs[49], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"generated",
			"extensions",
			"additional",
			"empty",
		}),
	})
	// properties of NutritionOrder
	typeDefs[457].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("allergyIntolerance", typeDefs[32], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type dateTime
		common.NewPropDef("dateTime", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type NutritionOrder_EnteralFormula
		common.NewPropDef("enteralFormula", typeDefs[459], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("excludeFoodModifier", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("foodPreferenceModifier", typeDefs[8], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type uri
		common.NewPropDef("instantiates", typeDefs[48], "", true, nil),
		// with type canonical
		common.NewPropDef("instantiatesCanonical", typeDefs[673], "", true, nil),
		// with type uri
		common.NewPropDef("instantiatesUri", typeDefs[48], "", true, nil),
		// with type code
		common.NewPropDef("intent", typeDefs[674], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type NutritionOrder_OralDiet
		common.NewPropDef("oralDiet", typeDefs[461], "", false, nil),
		// with type Reference
		common.NewPropDef("orderer", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("patient", typeDefs[32], "", false, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type NutritionOrder_Supplement
		common.NewPropDef("supplement", typeDefs[462], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of NutritionOrder_Administration
	typeDefs[458].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type Quantity
		common.NewPropDef("rateQuantity", typeDefs[29], "", false, nil),
		// with type Ratio
		common.NewPropDef("rateRatio", typeDefs[31], "", false, nil),
		// with type Timing
		common.NewPropDef("schedule", typeDefs[655], "", false, nil),
	})
	// properties of NutritionOrder_EnteralFormula
	typeDefs[459].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("additiveProductName", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("additiveType", typeDefs[8], "", false, nil),
		// with type NutritionOrder_Administration
		common.NewPropDef("administration", typeDefs[458], "", true, nil),
		// with type string
		common.NewPropDef("administrationInstruction", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("baseFormulaProductName", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("baseFormulaType", typeDefs[8], "", false, nil),
		// with type Quantity
		common.NewPropDef("caloricDensity", typeDefs[29], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Quantity
		common.NewPropDef("maxVolumeToDeliver", typeDefs[29], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("routeofAdministration", typeDefs[8], "", false, nil),
	})
	// properties of NutritionOrder_Nutrient
	typeDefs[460].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Quantity
		common.NewPropDef("amount", typeDefs[29], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("modifier", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of NutritionOrder_OralDiet
	typeDefs[461].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("fluidConsistencyType", typeDefs[8], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("instruction", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type NutritionOrder_Nutrient
		common.NewPropDef("nutrient", typeDefs[460], "", true, nil),
		// with type Timing
		common.NewPropDef("schedule", typeDefs[655], "", true, nil),
		// with type NutritionOrder_Texture
		common.NewPropDef("texture", typeDefs[463], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", true, nil),
	})
	// properties of NutritionOrder_Supplement
	typeDefs[462].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("instruction", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("productName", typeDefs[46], "", false, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type Timing
		common.NewPropDef("schedule", typeDefs[655], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of NutritionOrder_Texture
	typeDefs[463].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("foodType", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("modifier", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Observation
	typeDefs[464].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("basedOn", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("bodySite", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Observation_Component
		common.NewPropDef("component", typeDefs[468], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("dataAbsentReason", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("derivedFrom", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("device", typeDefs[32], "", false, nil),
		// with type dateTime
		common.NewPropDef("effectiveDateTime", typeDefs[42], "", false, nil),
		// with type instant
		common.NewPropDef("effectiveInstant", typeDefs[44], "", false, nil),
		// with type Period
		common.NewPropDef("effectivePeriod", typeDefs[25], "", false, nil),
		// with type Timing
		common.NewPropDef("effectiveTiming", typeDefs[655], "", false, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("focus", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("hasMember", typeDefs[32], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("interpretation", typeDefs[8], "", true, nil),
		// with type instant
		common.NewPropDef("issued", typeDefs[44], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("method", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type Reference
		common.NewPropDef("partOf", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("performer", typeDefs[32], "", true, nil),
		// with type Observation_ReferenceRange
		common.NewPropDef("referenceRange", typeDefs[469], "", true, nil),
		// with type Reference
		common.NewPropDef("specimen", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"registered",
			"preliminary",
			"final",
			"amended",
			"corrected",
			"cancelled",
			"entered-in-error",
			"unknown",
		}),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type boolean
		common.NewPropDef("valueBoolean", typeDefs[40], "value", false, nil),
		// with type CodeableConcept
		common.NewPropDef("valueCodeableConcept", typeDefs[8], "value", false, nil),
		// with type dateTime
		common.NewPropDef("valueDateTime", typeDefs[42], "value", false, nil),
		// with type integer
		common.NewPropDef("valueInteger", typeDefs[45], "value", false, nil),
		// with type Period
		common.NewPropDef("valuePeriod", typeDefs[25], "value", false, nil),
		// with type Quantity
		common.NewPropDef("valueQuantity", typeDefs[29], "value", false, nil),
		// with type Range
		common.NewPropDef("valueRange", typeDefs[30], "value", false, nil),
		// with type Ratio
		common.NewPropDef("valueRatio", typeDefs[31], "value", false, nil),
		// with type SampledData
		common.NewPropDef("valueSampledData", typeDefs[34], "value", false, nil),
		// with type string
		common.NewPropDef("valueString", typeDefs[46], "value", false, nil),
		// with type time
		common.NewPropDef("valueTime", typeDefs[47], "value", false, nil),
	})
	// properties of ObservationDefinition
	typeDefs[465].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("abnormalCodedValueSet", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("criticalCodedValueSet", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("method", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("multipleResultsAllowed", typeDefs[40], "", false, nil),
		// with type Reference
		common.NewPropDef("normalCodedValueSet", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("permittedDataType", typeDefs[46], "", true, []string{
			"Quantity",
			"CodeableConcept",
			"string",
			"boolean",
			"integer",
			"Range",
			"Ratio",
			"SampledData",
			"time",
			"dateTime",
			"Period",
		}),
		// with type string
		common.NewPropDef("preferredReportName", typeDefs[46], "", false, nil),
		// with type ObservationDefinition_QualifiedInterval
		common.NewPropDef("qualifiedInterval", typeDefs[466], "", true, nil),
		// with type ObservationDefinition_QuantitativeDetails
		common.NewPropDef("quantitativeDetails", typeDefs[467], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type Reference
		common.NewPropDef("validCodedValueSet", typeDefs[32], "", false, nil),
	})
	// properties of ObservationDefinition_QualifiedInterval
	typeDefs[466].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Range
		common.NewPropDef("age", typeDefs[30], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("appliesTo", typeDefs[8], "", true, nil),
		// with type string
		common.NewPropDef("category", typeDefs[46], "", false, []string{
			"reference",
			"critical",
			"absolute",
		}),
		// with type string
		common.NewPropDef("condition", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("context", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("gender", typeDefs[46], "", false, []string{
			"male",
			"female",
			"other",
			"unknown",
		}),
		// with type Range
		common.NewPropDef("gestationalAge", typeDefs[30], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Range
		common.NewPropDef("range", typeDefs[30], "", false, nil),
	})
	// properties of ObservationDefinition_QuantitativeDetails
	typeDefs[467].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type decimal
		common.NewPropDef("conversionFactor", typeDefs[43], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("customaryUnit", typeDefs[8], "", false, nil),
		// with type integer
		common.NewPropDef("decimalPrecision", typeDefs[45], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("unit", typeDefs[8], "", false, nil),
	})
	// properties of Observation_Component
	typeDefs[468].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("dataAbsentReason", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("interpretation", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Observation_ReferenceRange
		common.NewPropDef("referenceRange", typeDefs[469], "", true, nil),
		// with type boolean
		common.NewPropDef("valueBoolean", typeDefs[40], "value", false, nil),
		// with type CodeableConcept
		common.NewPropDef("valueCodeableConcept", typeDefs[8], "value", false, nil),
		// with type dateTime
		common.NewPropDef("valueDateTime", typeDefs[42], "value", false, nil),
		// with type integer
		common.NewPropDef("valueInteger", typeDefs[45], "value", false, nil),
		// with type Period
		common.NewPropDef("valuePeriod", typeDefs[25], "value", false, nil),
		// with type Quantity
		common.NewPropDef("valueQuantity", typeDefs[29], "value", false, nil),
		// with type Range
		common.NewPropDef("valueRange", typeDefs[30], "value", false, nil),
		// with type Ratio
		common.NewPropDef("valueRatio", typeDefs[31], "value", false, nil),
		// with type SampledData
		common.NewPropDef("valueSampledData", typeDefs[34], "value", false, nil),
		// with type string
		common.NewPropDef("valueString", typeDefs[46], "value", false, nil),
		// with type time
		common.NewPropDef("valueTime", typeDefs[47], "value", false, nil),
	})
	// properties of Observation_ReferenceRange
	typeDefs[469].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Range
		common.NewPropDef("age", typeDefs[30], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("appliesTo", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Quantity
		common.NewPropDef("high", typeDefs[29], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Quantity
		common.NewPropDef("low", typeDefs[29], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("text", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of OperationDefinition
	typeDefs[470].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("affectsState", typeDefs[40], "", false, nil),
		// with type canonical
		common.NewPropDef("base", typeDefs[673], "", false, nil),
		// with type code
		common.NewPropDef("code", typeDefs[674], "", false, nil),
		// with type markdown
		common.NewPropDef("comment", typeDefs[676], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type boolean
		common.NewPropDef("experimental", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type canonical
		common.NewPropDef("inputProfile", typeDefs[673], "", false, nil),
		// with type boolean
		common.NewPropDef("instance", typeDefs[40], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type string
		common.NewPropDef("kind", typeDefs[46], "", false, []string{
			"operation",
			"query",
		}),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type canonical
		common.NewPropDef("outputProfile", typeDefs[673], "", false, nil),
		// with type OperationDefinition_Overload
		common.NewPropDef("overload", typeDefs[472], "", true, nil),
		// with type OperationDefinition_Parameter
		common.NewPropDef("parameter", typeDefs[473], "", true, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("purpose", typeDefs[676], "", false, nil),
		// with type code
		common.NewPropDef("resource", typeDefs[674], "", true, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type boolean
		common.NewPropDef("system", typeDefs[40], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("type", typeDefs[40], "", false, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of OperationDefinition_Binding
	typeDefs[471].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("strength", typeDefs[46], "", false, []string{
			"required",
			"extensible",
			"preferred",
			"example",
		}),
		// with type canonical
		common.NewPropDef("valueSet", typeDefs[673], "", false, nil),
	})
	// properties of OperationDefinition_Overload
	typeDefs[472].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("comment", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("parameterName", typeDefs[46], "", true, nil),
	})
	// properties of OperationDefinition_Parameter
	typeDefs[473].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type OperationDefinition_Binding
		common.NewPropDef("binding", typeDefs[471], "", false, nil),
		// with type string
		common.NewPropDef("documentation", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("max", typeDefs[46], "", false, nil),
		// with type integer
		common.NewPropDef("min", typeDefs[45], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type code
		common.NewPropDef("name", typeDefs[674], "", false, nil),
		// with type OperationDefinition_Parameter
		common.NewPropDef("part", typeDefs[473], "", true, nil),
		// with type OperationDefinition_ReferencedFrom
		common.NewPropDef("referencedFrom", typeDefs[474], "", true, nil),
		// with type string
		common.NewPropDef("searchType", typeDefs[46], "", false, []string{
			"number",
			"date",
			"string",
			"token",
			"reference",
			"composite",
			"quantity",
			"uri",
			"special",
		}),
		// with type canonical
		common.NewPropDef("targetProfile", typeDefs[673], "", true, nil),
		// with type code
		common.NewPropDef("type", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("use", typeDefs[46], "", false, []string{
			"in",
			"out",
		}),
	})
	// properties of OperationDefinition_ReferencedFrom
	typeDefs[474].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("source", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("sourceId", typeDefs[46], "", false, nil),
	})
	// properties of OperationOutcome
	typeDefs[475].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type OperationOutcome_Issue
		common.NewPropDef("issue", typeDefs[476], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of OperationOutcome_Issue
	typeDefs[476].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("code", typeDefs[46], "", false, []string{
			"invalid",
			"structure",
			"required",
			"value",
			"invariant",
			"security",
			"login",
			"unknown",
			"expired",
			"forbidden",
			"suppressed",
			"processing",
			"not-supported",
			"duplicate",
			"multiple-matches",
			"not-found",
			"deleted",
			"too-long",
			"code-invalid",
			"extension",
			"too-costly",
			"business-rule",
			"conflict",
			"transient",
			"lock-error",
			"no-store",
			"exception",
			"timeout",
			"incomplete",
			"throttled",
			"informational",
		}),
		// with type CodeableConcept
		common.NewPropDef("details", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("diagnostics", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("expression", typeDefs[46], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("location", typeDefs[46], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("severity", typeDefs[46], "", false, []string{
			"fatal",
			"error",
			"warning",
			"information",
		}),
	})
	// properties of Organization
	typeDefs[477].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("active", typeDefs[40], "", false, nil),
		// with type Address
		common.NewPropDef("address", typeDefs[2], "", true, nil),
		// with type string
		common.NewPropDef("alias", typeDefs[46], "", true, nil),
		// with type Organization_Contact
		common.NewPropDef("contact", typeDefs[479], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("endpoint", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("partOf", typeDefs[32], "", false, nil),
		// with type ContactPoint
		common.NewPropDef("telecom", typeDefs[11], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", true, nil),
	})
	// properties of OrganizationAffiliation
	typeDefs[478].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("active", typeDefs[40], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("endpoint", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("healthcareService", typeDefs[32], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("location", typeDefs[32], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("network", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("organization", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("participatingOrganization", typeDefs[32], "", false, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("specialty", typeDefs[8], "", true, nil),
		// with type ContactPoint
		common.NewPropDef("telecom", typeDefs[11], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of Organization_Contact
	typeDefs[479].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Address
		common.NewPropDef("address", typeDefs[2], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type HumanName
		common.NewPropDef("name", typeDefs[17], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("purpose", typeDefs[8], "", false, nil),
		// with type ContactPoint
		common.NewPropDef("telecom", typeDefs[11], "", true, nil),
	})
	// properties of ParameterDefinition
	typeDefs[23].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("documentation", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("max", typeDefs[46], "", false, nil),
		// with type integer
		common.NewPropDef("min", typeDefs[45], "", false, nil),
		// with type code
		common.NewPropDef("name", typeDefs[674], "", false, nil),
		// with type canonical
		common.NewPropDef("profile", typeDefs[673], "", false, nil),
		// with type code
		common.NewPropDef("type", typeDefs[674], "", false, nil),
		// with type code
		common.NewPropDef("use", typeDefs[674], "", false, nil),
	})
	// properties of Parameters
	typeDefs[24].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Parameters_Parameter
		common.NewPropDef("parameter", typeDefs[480], "", true, nil),
	})
	// properties of Parameters_Parameter
	typeDefs[480].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type Parameters_Parameter
		common.NewPropDef("part", typeDefs[480], "", true, nil),
		// with type Resource
		common.NewPropDef("resource", typeDefs[1], "", false, nil),
		// with type Address
		common.NewPropDef("valueAddress", typeDefs[2], "value", false, nil),
		// with type Age
		common.NewPropDef("valueAge", typeDefs[59], "value", false, nil),
		// with type Annotation
		common.NewPropDef("valueAnnotation", typeDefs[3], "value", false, nil),
		// with type Attachment
		common.NewPropDef("valueAttachment", typeDefs[4], "value", false, nil),
		// with type base64Binary
		common.NewPropDef("valueBase64Binary", typeDefs[39], "value", false, nil),
		// with type boolean
		common.NewPropDef("valueBoolean", typeDefs[40], "value", false, nil),
		// with type canonical
		common.NewPropDef("valueCanonical", typeDefs[673], "value", false, nil),
		// with type code
		common.NewPropDef("valueCode", typeDefs[674], "value", false, nil),
		// with type CodeableConcept
		common.NewPropDef("valueCodeableConcept", typeDefs[8], "value", false, nil),
		// with type Coding
		common.NewPropDef("valueCoding", typeDefs[9], "value", false, nil),
		// with type ContactDetail
		common.NewPropDef("valueContactDetail", typeDefs[10], "value", false, nil),
		// with type ContactPoint
		common.NewPropDef("valueContactPoint", typeDefs[11], "value", false, nil),
		// with type Contributor
		common.NewPropDef("valueContributor", typeDefs[12], "value", false, nil),
		// with type Count
		common.NewPropDef("valueCount", typeDefs[186], "value", false, nil),
		// with type DataRequirement
		common.NewPropDef("valueDataRequirement", typeDefs[13], "value", false, nil),
		// with type date
		common.NewPropDef("valueDate", typeDefs[41], "value", false, nil),
		// with type dateTime
		common.NewPropDef("valueDateTime", typeDefs[42], "value", false, nil),
		// with type decimal
		common.NewPropDef("valueDecimal", typeDefs[43], "value", false, nil),
		// with type Distance
		common.NewPropDef("valueDistance", typeDefs[227], "value", false, nil),
		// with type Dosage
		common.NewPropDef("valueDosage", typeDefs[234], "value", false, nil),
		// with type Duration
		common.NewPropDef("valueDuration", typeDefs[236], "value", false, nil),
		// with type Expression
		common.NewPropDef("valueExpression", typeDefs[15], "value", false, nil),
		// with type HumanName
		common.NewPropDef("valueHumanName", typeDefs[17], "value", false, nil),
		// with type id
		common.NewPropDef("valueId", typeDefs[675], "value", false, nil),
		// with type Identifier
		common.NewPropDef("valueIdentifier", typeDefs[18], "value", false, nil),
		// with type instant
		common.NewPropDef("valueInstant", typeDefs[44], "value", false, nil),
		// with type integer
		common.NewPropDef("valueInteger", typeDefs[45], "value", false, nil),
		// with type markdown
		common.NewPropDef("valueMarkdown", typeDefs[676], "value", false, nil),
		// with type Meta
		common.NewPropDef("valueMeta", typeDefs[20], "value", false, nil),
		// with type Money
		common.NewPropDef("valueMoney", typeDefs[21], "value", false, nil),
		// with type oid
		common.NewPropDef("valueOid", typeDefs[677], "value", false, nil),
		// with type ParameterDefinition
		common.NewPropDef("valueParameterDefinition", typeDefs[23], "value", false, nil),
		// with type Period
		common.NewPropDef("valuePeriod", typeDefs[25], "value", false, nil),
		// with type positiveInt
		common.NewPropDef("valuePositiveInt", typeDefs[678], "value", false, nil),
		// with type Quantity
		common.NewPropDef("valueQuantity", typeDefs[29], "value", false, nil),
		// with type Range
		common.NewPropDef("valueRange", typeDefs[30], "value", false, nil),
		// with type Ratio
		common.NewPropDef("valueRatio", typeDefs[31], "value", false, nil),
		// with type Reference
		common.NewPropDef("valueReference", typeDefs[32], "value", false, nil),
		// with type RelatedArtifact
		common.NewPropDef("valueRelatedArtifact", typeDefs[33], "value", false, nil),
		// with type SampledData
		common.NewPropDef("valueSampledData", typeDefs[34], "value", false, nil),
		// with type Signature
		common.NewPropDef("valueSignature", typeDefs[35], "value", false, nil),
		// with type string
		common.NewPropDef("valueString", typeDefs[46], "value", false, nil),
		// with type time
		common.NewPropDef("valueTime", typeDefs[47], "value", false, nil),
		// with type Timing
		common.NewPropDef("valueTiming", typeDefs[655], "value", false, nil),
		// with type TriggerDefinition
		common.NewPropDef("valueTriggerDefinition", typeDefs[37], "value", false, nil),
		// with type unsignedInt
		common.NewPropDef("valueUnsignedInt", typeDefs[679], "value", false, nil),
		// with type uri
		common.NewPropDef("valueUri", typeDefs[48], "value", false, nil),
		// with type url
		common.NewPropDef("valueUrl", typeDefs[680], "value", false, nil),
		// with type UsageContext
		common.NewPropDef("valueUsageContext", typeDefs[38], "value", false, nil),
		// with type uuid
		common.NewPropDef("valueUuid", typeDefs[681], "value", false, nil),
	})
	// properties of Patient
	typeDefs[481].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("active", typeDefs[40], "", false, nil),
		// with type Address
		common.NewPropDef("address", typeDefs[2], "", true, nil),
		// with type date
		common.NewPropDef("birthDate", typeDefs[41], "", false, nil),
		// with type Patient_Communication
		common.NewPropDef("communication", typeDefs[482], "", true, nil),
		// with type Patient_Contact
		common.NewPropDef("contact", typeDefs[483], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type boolean
		common.NewPropDef("deceasedBoolean", typeDefs[40], "", false, nil),
		// with type dateTime
		common.NewPropDef("deceasedDateTime", typeDefs[42], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("gender", typeDefs[46], "", false, []string{
			"male",
			"female",
			"other",
			"unknown",
		}),
		// with type Reference
		common.NewPropDef("generalPractitioner", typeDefs[32], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Patient_Link
		common.NewPropDef("link", typeDefs[484], "", true, nil),
		// with type Reference
		common.NewPropDef("managingOrganization", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("maritalStatus", typeDefs[8], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("multipleBirthBoolean", typeDefs[40], "", false, nil),
		// with type integer
		common.NewPropDef("multipleBirthInteger", typeDefs[45], "", false, nil),
		// with type HumanName
		common.NewPropDef("name", typeDefs[17], "", true, nil),
		// with type Attachment
		common.NewPropDef("photo", typeDefs[4], "", true, nil),
		// with type ContactPoint
		common.NewPropDef("telecom", typeDefs[11], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of Patient_Communication
	typeDefs[482].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("language", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("preferred", typeDefs[40], "", false, nil),
	})
	// properties of Patient_Contact
	typeDefs[483].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Address
		common.NewPropDef("address", typeDefs[2], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("gender", typeDefs[46], "", false, []string{
			"male",
			"female",
			"other",
			"unknown",
		}),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type HumanName
		common.NewPropDef("name", typeDefs[17], "", false, nil),
		// with type Reference
		common.NewPropDef("organization", typeDefs[32], "", false, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("relationship", typeDefs[8], "", true, nil),
		// with type ContactPoint
		common.NewPropDef("telecom", typeDefs[11], "", true, nil),
	})
	// properties of Patient_Link
	typeDefs[484].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("other", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"replaced-by",
			"replaces",
			"refer",
			"seealso",
		}),
	})
	// properties of PaymentNotice
	typeDefs[485].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Money
		common.NewPropDef("amount", typeDefs[21], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type dateTime
		common.NewPropDef("created", typeDefs[42], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("payee", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("payment", typeDefs[32], "", false, nil),
		// with type date
		common.NewPropDef("paymentDate", typeDefs[41], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("paymentStatus", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("provider", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("recipient", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("request", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("response", typeDefs[32], "", false, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of PaymentReconciliation
	typeDefs[486].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type dateTime
		common.NewPropDef("created", typeDefs[42], "", false, nil),
		// with type PaymentReconciliation_Detail
		common.NewPropDef("detail", typeDefs[487], "", true, nil),
		// with type string
		common.NewPropDef("disposition", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("formCode", typeDefs[8], "", false, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("outcome", typeDefs[46], "", false, []string{
			"queued",
			"complete",
			"error",
			"partial",
		}),
		// with type Money
		common.NewPropDef("paymentAmount", typeDefs[21], "", false, nil),
		// with type date
		common.NewPropDef("paymentDate", typeDefs[41], "", false, nil),
		// with type Identifier
		common.NewPropDef("paymentIdentifier", typeDefs[18], "", false, nil),
		// with type Reference
		common.NewPropDef("paymentIssuer", typeDefs[32], "", false, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type PaymentReconciliation_ProcessNote
		common.NewPropDef("processNote", typeDefs[488], "", true, nil),
		// with type Reference
		common.NewPropDef("request", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("requestor", typeDefs[32], "", false, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of PaymentReconciliation_Detail
	typeDefs[487].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Money
		common.NewPropDef("amount", typeDefs[21], "", false, nil),
		// with type date
		common.NewPropDef("date", typeDefs[41], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("payee", typeDefs[32], "", false, nil),
		// with type Identifier
		common.NewPropDef("predecessor", typeDefs[18], "", false, nil),
		// with type Reference
		common.NewPropDef("request", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("response", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("responsible", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("submitter", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of PaymentReconciliation_ProcessNote
	typeDefs[488].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("text", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"display",
			"print",
			"printoper",
		}),
	})
	// properties of Period
	typeDefs[25].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type dateTime
		common.NewPropDef("end", typeDefs[42], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type dateTime
		common.NewPropDef("start", typeDefs[42], "", false, nil),
	})
	// properties of Person
	typeDefs[489].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("active", typeDefs[40], "", false, nil),
		// with type Address
		common.NewPropDef("address", typeDefs[2], "", true, nil),
		// with type date
		common.NewPropDef("birthDate", typeDefs[41], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("gender", typeDefs[46], "", false, []string{
			"male",
			"female",
			"other",
			"unknown",
		}),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Person_Link
		common.NewPropDef("link", typeDefs[490], "", true, nil),
		// with type Reference
		common.NewPropDef("managingOrganization", typeDefs[32], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type HumanName
		common.NewPropDef("name", typeDefs[17], "", true, nil),
		// with type Attachment
		common.NewPropDef("photo", typeDefs[4], "", false, nil),
		// with type ContactPoint
		common.NewPropDef("telecom", typeDefs[11], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of Person_Link
	typeDefs[490].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("assurance", typeDefs[46], "", false, []string{
			"level1",
			"level2",
			"level3",
			"level4",
		}),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("target", typeDefs[32], "", false, nil),
	})
	// properties of PlanDefinition
	typeDefs[491].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type PlanDefinition_Action
		common.NewPropDef("action", typeDefs[492], "", true, nil),
		// with type date
		common.NewPropDef("approvalDate", typeDefs[41], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("author", typeDefs[10], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("editor", typeDefs[10], "", true, nil),
		// with type Period
		common.NewPropDef("effectivePeriod", typeDefs[25], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("endorser", typeDefs[10], "", true, nil),
		// with type boolean
		common.NewPropDef("experimental", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type PlanDefinition_Goal
		common.NewPropDef("goal", typeDefs[495], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type date
		common.NewPropDef("lastReviewDate", typeDefs[41], "", false, nil),
		// with type canonical
		common.NewPropDef("library", typeDefs[673], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("purpose", typeDefs[676], "", false, nil),
		// with type RelatedArtifact
		common.NewPropDef("relatedArtifact", typeDefs[33], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("reviewer", typeDefs[10], "", true, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type CodeableConcept
		common.NewPropDef("subjectCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("subjectReference", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("subtitle", typeDefs[46], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("topic", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type string
		common.NewPropDef("usage", typeDefs[46], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of PlanDefinition_Action
	typeDefs[492].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type PlanDefinition_Action
		common.NewPropDef("action", typeDefs[492], "", true, nil),
		// with type string
		common.NewPropDef("cardinalityBehavior", typeDefs[46], "", false, []string{
			"single",
			"multiple",
		}),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", true, nil),
		// with type PlanDefinition_Condition
		common.NewPropDef("condition", typeDefs[493], "", true, nil),
		// with type canonical
		common.NewPropDef("definitionCanonical", typeDefs[673], "", false, nil),
		// with type uri
		common.NewPropDef("definitionUri", typeDefs[48], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type RelatedArtifact
		common.NewPropDef("documentation", typeDefs[33], "", true, nil),
		// with type PlanDefinition_DynamicValue
		common.NewPropDef("dynamicValue", typeDefs[494], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("goalId", typeDefs[675], "", true, nil),
		// with type string
		common.NewPropDef("groupingBehavior", typeDefs[46], "", false, []string{
			"visual-group",
			"logical-group",
			"sentence-group",
		}),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type DataRequirement
		common.NewPropDef("input", typeDefs[13], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type DataRequirement
		common.NewPropDef("output", typeDefs[13], "", true, nil),
		// with type PlanDefinition_Participant
		common.NewPropDef("participant", typeDefs[496], "", true, nil),
		// with type string
		common.NewPropDef("precheckBehavior", typeDefs[46], "", false, []string{
			"yes",
			"no",
		}),
		// with type string
		common.NewPropDef("prefix", typeDefs[46], "", false, nil),
		// with type code
		common.NewPropDef("priority", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("reason", typeDefs[8], "", true, nil),
		// with type PlanDefinition_RelatedAction
		common.NewPropDef("relatedAction", typeDefs[497], "", true, nil),
		// with type string
		common.NewPropDef("requiredBehavior", typeDefs[46], "", false, []string{
			"must",
			"could",
			"must-unless-documented",
		}),
		// with type string
		common.NewPropDef("selectionBehavior", typeDefs[46], "", false, []string{
			"any",
			"all",
			"all-or-none",
			"exactly-one",
			"at-most-one",
			"one-or-more",
		}),
		// with type CodeableConcept
		common.NewPropDef("subjectCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("subjectReference", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("textEquivalent", typeDefs[46], "", false, nil),
		// with type Age
		common.NewPropDef("timingAge", typeDefs[59], "", false, nil),
		// with type dateTime
		common.NewPropDef("timingDateTime", typeDefs[42], "", false, nil),
		// with type Duration
		common.NewPropDef("timingDuration", typeDefs[236], "", false, nil),
		// with type Period
		common.NewPropDef("timingPeriod", typeDefs[25], "", false, nil),
		// with type Range
		common.NewPropDef("timingRange", typeDefs[30], "", false, nil),
		// with type Timing
		common.NewPropDef("timingTiming", typeDefs[655], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type canonical
		common.NewPropDef("transform", typeDefs[673], "", false, nil),
		// with type TriggerDefinition
		common.NewPropDef("trigger", typeDefs[37], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of PlanDefinition_Condition
	typeDefs[493].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Expression
		common.NewPropDef("expression", typeDefs[15], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("kind", typeDefs[46], "", false, []string{
			"applicability",
			"start",
			"stop",
		}),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of PlanDefinition_DynamicValue
	typeDefs[494].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Expression
		common.NewPropDef("expression", typeDefs[15], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("path", typeDefs[46], "", false, nil),
	})
	// properties of PlanDefinition_Goal
	typeDefs[495].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("addresses", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("description", typeDefs[8], "", false, nil),
		// with type RelatedArtifact
		common.NewPropDef("documentation", typeDefs[33], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("priority", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("start", typeDefs[8], "", false, nil),
		// with type PlanDefinition_Target
		common.NewPropDef("target", typeDefs[498], "", true, nil),
	})
	// properties of PlanDefinition_Participant
	typeDefs[496].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("role", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"patient",
			"practitioner",
			"related-person",
			"device",
		}),
	})
	// properties of PlanDefinition_RelatedAction
	typeDefs[497].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type id
		common.NewPropDef("actionId", typeDefs[675], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Duration
		common.NewPropDef("offsetDuration", typeDefs[236], "", false, nil),
		// with type Range
		common.NewPropDef("offsetRange", typeDefs[30], "", false, nil),
		// with type string
		common.NewPropDef("relationship", typeDefs[46], "", false, []string{
			"before-start",
			"before",
			"before-end",
			"concurrent-with-start",
			"concurrent",
			"concurrent-with-end",
			"after-start",
			"after",
			"after-end",
		}),
	})
	// properties of PlanDefinition_Target
	typeDefs[498].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("detailCodeableConcept", typeDefs[8], "", false, nil),
		// with type Quantity
		common.NewPropDef("detailQuantity", typeDefs[29], "", false, nil),
		// with type Range
		common.NewPropDef("detailRange", typeDefs[30], "", false, nil),
		// with type Duration
		common.NewPropDef("due", typeDefs[236], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("measure", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Population
	typeDefs[26].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("ageCodeableConcept", typeDefs[8], "", false, nil),
		// with type Range
		common.NewPropDef("ageRange", typeDefs[30], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("gender", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("physiologicalCondition", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("race", typeDefs[8], "", false, nil),
	})
	// properties of Practitioner
	typeDefs[499].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("active", typeDefs[40], "", false, nil),
		// with type Address
		common.NewPropDef("address", typeDefs[2], "", true, nil),
		// with type date
		common.NewPropDef("birthDate", typeDefs[41], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("communication", typeDefs[8], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("gender", typeDefs[46], "", false, []string{
			"male",
			"female",
			"other",
			"unknown",
		}),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type HumanName
		common.NewPropDef("name", typeDefs[17], "", true, nil),
		// with type Attachment
		common.NewPropDef("photo", typeDefs[4], "", true, nil),
		// with type Practitioner_Qualification
		common.NewPropDef("qualification", typeDefs[503], "", true, nil),
		// with type ContactPoint
		common.NewPropDef("telecom", typeDefs[11], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of PractitionerRole
	typeDefs[500].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("active", typeDefs[40], "", false, nil),
		// with type string
		common.NewPropDef("availabilityExceptions", typeDefs[46], "", false, nil),
		// with type PractitionerRole_AvailableTime
		common.NewPropDef("availableTime", typeDefs[501], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("endpoint", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("healthcareService", typeDefs[32], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("location", typeDefs[32], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type PractitionerRole_NotAvailable
		common.NewPropDef("notAvailable", typeDefs[502], "", true, nil),
		// with type Reference
		common.NewPropDef("organization", typeDefs[32], "", false, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type Reference
		common.NewPropDef("practitioner", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("specialty", typeDefs[8], "", true, nil),
		// with type ContactPoint
		common.NewPropDef("telecom", typeDefs[11], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of PractitionerRole_AvailableTime
	typeDefs[501].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("allDay", typeDefs[40], "", false, nil),
		// with type time
		common.NewPropDef("availableEndTime", typeDefs[47], "", false, nil),
		// with type time
		common.NewPropDef("availableStartTime", typeDefs[47], "", false, nil),
		// with type code
		common.NewPropDef("daysOfWeek", typeDefs[674], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of PractitionerRole_NotAvailable
	typeDefs[502].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Period
		common.NewPropDef("during", typeDefs[25], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Practitioner_Qualification
	typeDefs[503].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type Reference
		common.NewPropDef("issuer", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
	})
	// properties of Procedure
	typeDefs[504].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("asserter", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("basedOn", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("bodySite", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("complication", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("complicationDetail", typeDefs[32], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Procedure_FocalDevice
		common.NewPropDef("focalDevice", typeDefs[505], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("followUp", typeDefs[8], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type canonical
		common.NewPropDef("instantiatesCanonical", typeDefs[673], "", true, nil),
		// with type uri
		common.NewPropDef("instantiatesUri", typeDefs[48], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("location", typeDefs[32], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("outcome", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("partOf", typeDefs[32], "", true, nil),
		// with type Age
		common.NewPropDef("performedAge", typeDefs[59], "", false, nil),
		// with type dateTime
		common.NewPropDef("performedDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("performedPeriod", typeDefs[25], "", false, nil),
		// with type Range
		common.NewPropDef("performedRange", typeDefs[30], "", false, nil),
		// with type string
		common.NewPropDef("performedString", typeDefs[46], "", false, nil),
		// with type Procedure_Performer
		common.NewPropDef("performer", typeDefs[506], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("reasonCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("reasonReference", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("recorder", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("report", typeDefs[32], "", true, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("statusReason", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("usedCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("usedReference", typeDefs[32], "", true, nil),
	})
	// properties of Procedure_FocalDevice
	typeDefs[505].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("action", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("manipulated", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of Procedure_Performer
	typeDefs[506].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("actor", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("function", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("onBehalfOf", typeDefs[32], "", false, nil),
	})
	// properties of ProdCharacteristic
	typeDefs[27].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("color", typeDefs[46], "", true, nil),
		// with type Quantity
		common.NewPropDef("depth", typeDefs[29], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Quantity
		common.NewPropDef("externalDiameter", typeDefs[29], "", false, nil),
		// with type Quantity
		common.NewPropDef("height", typeDefs[29], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Attachment
		common.NewPropDef("image", typeDefs[4], "", true, nil),
		// with type string
		common.NewPropDef("imprint", typeDefs[46], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Quantity
		common.NewPropDef("nominalVolume", typeDefs[29], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("scoring", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("shape", typeDefs[46], "", false, nil),
		// with type Quantity
		common.NewPropDef("weight", typeDefs[29], "", false, nil),
		// with type Quantity
		common.NewPropDef("width", typeDefs[29], "", false, nil),
	})
	// properties of ProductShelfLife
	typeDefs[28].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Quantity
		common.NewPropDef("period", typeDefs[29], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("specialPrecautionsForStorage", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of Provenance
	typeDefs[507].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("activity", typeDefs[8], "", false, nil),
		// with type Provenance_Agent
		common.NewPropDef("agent", typeDefs[508], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Provenance_Entity
		common.NewPropDef("entity", typeDefs[509], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("location", typeDefs[32], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type dateTime
		common.NewPropDef("occurredDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("occurredPeriod", typeDefs[25], "", false, nil),
		// with type uri
		common.NewPropDef("policy", typeDefs[48], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("reason", typeDefs[8], "", true, nil),
		// with type instant
		common.NewPropDef("recorded", typeDefs[44], "", false, nil),
		// with type Signature
		common.NewPropDef("signature", typeDefs[35], "", true, nil),
		// with type Reference
		common.NewPropDef("target", typeDefs[32], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of Provenance_Agent
	typeDefs[508].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("onBehalfOf", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("role", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("who", typeDefs[32], "", false, nil),
	})
	// properties of Provenance_Entity
	typeDefs[509].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Provenance_Agent
		common.NewPropDef("agent", typeDefs[508], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("role", typeDefs[46], "", false, []string{
			"derivation",
			"revision",
			"quotation",
			"source",
			"removal",
		}),
		// with type Reference
		common.NewPropDef("what", typeDefs[32], "", false, nil),
	})
	// properties of Quantity
	typeDefs[29].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("code", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("comparator", typeDefs[46], "", false, []string{
			"<",
			"<=",
			">=",
			">",
		}),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("system", typeDefs[48], "", false, nil),
		// with type string
		common.NewPropDef("unit", typeDefs[46], "", false, nil),
		// with type decimal
		common.NewPropDef("value", typeDefs[43], "", false, nil),
	})
	// properties of Questionnaire
	typeDefs[510].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type date
		common.NewPropDef("approvalDate", typeDefs[41], "", false, nil),
		// with type Coding
		common.NewPropDef("code", typeDefs[9], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type canonical
		common.NewPropDef("derivedFrom", typeDefs[673], "", true, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type Period
		common.NewPropDef("effectivePeriod", typeDefs[25], "", false, nil),
		// with type boolean
		common.NewPropDef("experimental", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type Questionnaire_Item
		common.NewPropDef("item", typeDefs[517], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type date
		common.NewPropDef("lastReviewDate", typeDefs[41], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("purpose", typeDefs[676], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type code
		common.NewPropDef("subjectType", typeDefs[674], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of QuestionnaireResponse
	typeDefs[511].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("author", typeDefs[32], "", false, nil),
		// with type dateTime
		common.NewPropDef("authored", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("basedOn", typeDefs[32], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type QuestionnaireResponse_Item
		common.NewPropDef("item", typeDefs[513], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("partOf", typeDefs[32], "", true, nil),
		// with type canonical
		common.NewPropDef("questionnaire", typeDefs[673], "", false, nil),
		// with type Reference
		common.NewPropDef("source", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"in-progress",
			"completed",
			"amended",
			"entered-in-error",
			"stopped",
		}),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of QuestionnaireResponse_Answer
	typeDefs[512].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type QuestionnaireResponse_Item
		common.NewPropDef("item", typeDefs[513], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Attachment
		common.NewPropDef("valueAttachment", typeDefs[4], "value", false, nil),
		// with type boolean
		common.NewPropDef("valueBoolean", typeDefs[40], "value", false, nil),
		// with type Coding
		common.NewPropDef("valueCoding", typeDefs[9], "value", false, nil),
		// with type date
		common.NewPropDef("valueDate", typeDefs[41], "value", false, nil),
		// with type dateTime
		common.NewPropDef("valueDateTime", typeDefs[42], "value", false, nil),
		// with type decimal
		common.NewPropDef("valueDecimal", typeDefs[43], "value", false, nil),
		// with type integer
		common.NewPropDef("valueInteger", typeDefs[45], "value", false, nil),
		// with type Quantity
		common.NewPropDef("valueQuantity", typeDefs[29], "value", false, nil),
		// with type Reference
		common.NewPropDef("valueReference", typeDefs[32], "value", false, nil),
		// with type string
		common.NewPropDef("valueString", typeDefs[46], "value", false, nil),
		// with type time
		common.NewPropDef("valueTime", typeDefs[47], "value", false, nil),
		// with type uri
		common.NewPropDef("valueUri", typeDefs[48], "value", false, nil),
	})
	// properties of QuestionnaireResponse_Item
	typeDefs[513].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type QuestionnaireResponse_Answer
		common.NewPropDef("answer", typeDefs[512], "", true, nil),
		// with type uri
		common.NewPropDef("definition", typeDefs[48], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type QuestionnaireResponse_Item
		common.NewPropDef("item", typeDefs[513], "", true, nil),
		// with type string
		common.NewPropDef("linkId", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("text", typeDefs[46], "", false, nil),
	})
	// properties of Questionnaire_AnswerOption
	typeDefs[514].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("initialSelected", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Coding
		common.NewPropDef("valueCoding", typeDefs[9], "", false, nil),
		// with type date
		common.NewPropDef("valueDate", typeDefs[41], "", false, nil),
		// with type integer
		common.NewPropDef("valueInteger", typeDefs[45], "", false, nil),
		// with type Reference
		common.NewPropDef("valueReference", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("valueString", typeDefs[46], "", false, nil),
		// with type time
		common.NewPropDef("valueTime", typeDefs[47], "", false, nil),
	})
	// properties of Questionnaire_EnableWhen
	typeDefs[515].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("answerBoolean", typeDefs[40], "answer", false, nil),
		// with type Coding
		common.NewPropDef("answerCoding", typeDefs[9], "answer", false, nil),
		// with type date
		common.NewPropDef("answerDate", typeDefs[41], "answer", false, nil),
		// with type dateTime
		common.NewPropDef("answerDateTime", typeDefs[42], "answer", false, nil),
		// with type decimal
		common.NewPropDef("answerDecimal", typeDefs[43], "answer", false, nil),
		// with type integer
		common.NewPropDef("answerInteger", typeDefs[45], "answer", false, nil),
		// with type Quantity
		common.NewPropDef("answerQuantity", typeDefs[29], "answer", false, nil),
		// with type Reference
		common.NewPropDef("answerReference", typeDefs[32], "answer", false, nil),
		// with type string
		common.NewPropDef("answerString", typeDefs[46], "answer", false, nil),
		// with type time
		common.NewPropDef("answerTime", typeDefs[47], "answer", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("operator", typeDefs[46], "", false, []string{
			"exists",
			"=",
			"!=",
			">",
			"<",
			">=",
			"<=",
		}),
		// with type string
		common.NewPropDef("question", typeDefs[46], "", false, nil),
	})
	// properties of Questionnaire_Initial
	typeDefs[516].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Attachment
		common.NewPropDef("valueAttachment", typeDefs[4], "value", false, nil),
		// with type boolean
		common.NewPropDef("valueBoolean", typeDefs[40], "value", false, nil),
		// with type Coding
		common.NewPropDef("valueCoding", typeDefs[9], "value", false, nil),
		// with type date
		common.NewPropDef("valueDate", typeDefs[41], "value", false, nil),
		// with type dateTime
		common.NewPropDef("valueDateTime", typeDefs[42], "value", false, nil),
		// with type decimal
		common.NewPropDef("valueDecimal", typeDefs[43], "value", false, nil),
		// with type integer
		common.NewPropDef("valueInteger", typeDefs[45], "value", false, nil),
		// with type Quantity
		common.NewPropDef("valueQuantity", typeDefs[29], "value", false, nil),
		// with type Reference
		common.NewPropDef("valueReference", typeDefs[32], "value", false, nil),
		// with type string
		common.NewPropDef("valueString", typeDefs[46], "value", false, nil),
		// with type time
		common.NewPropDef("valueTime", typeDefs[47], "value", false, nil),
		// with type uri
		common.NewPropDef("valueUri", typeDefs[48], "value", false, nil),
	})
	// properties of Questionnaire_Item
	typeDefs[517].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Questionnaire_AnswerOption
		common.NewPropDef("answerOption", typeDefs[514], "", true, nil),
		// with type canonical
		common.NewPropDef("answerValueSet", typeDefs[673], "", false, nil),
		// with type Coding
		common.NewPropDef("code", typeDefs[9], "", true, nil),
		// with type uri
		common.NewPropDef("definition", typeDefs[48], "", false, nil),
		// with type string
		common.NewPropDef("enableBehavior", typeDefs[46], "", false, []string{
			"all",
			"any",
		}),
		// with type Questionnaire_EnableWhen
		common.NewPropDef("enableWhen", typeDefs[515], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Questionnaire_Initial
		common.NewPropDef("initial", typeDefs[516], "", true, nil),
		// with type Questionnaire_Item
		common.NewPropDef("item", typeDefs[517], "", true, nil),
		// with type string
		common.NewPropDef("linkId", typeDefs[46], "", false, nil),
		// with type integer
		common.NewPropDef("maxLength", typeDefs[45], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("prefix", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("readOnly", typeDefs[40], "", false, nil),
		// with type boolean
		common.NewPropDef("repeats", typeDefs[40], "", false, nil),
		// with type boolean
		common.NewPropDef("required", typeDefs[40], "", false, nil),
		// with type string
		common.NewPropDef("text", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"group",
			"display",
			"boolean",
			"decimal",
			"integer",
			"date",
			"dateTime",
			"time",
			"string",
			"text",
			"url",
			"choice",
			"open-choice",
			"attachment",
			"reference",
			"quantity",
		}),
	})
	// properties of Range
	typeDefs[30].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Quantity
		common.NewPropDef("high", typeDefs[29], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Quantity
		common.NewPropDef("low", typeDefs[29], "", false, nil),
	})
	// properties of Ratio
	typeDefs[31].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Quantity
		common.NewPropDef("denominator", typeDefs[29], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Quantity
		common.NewPropDef("numerator", typeDefs[29], "", false, nil),
	})
	// properties of Reference
	typeDefs[32].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("display", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type string
		common.NewPropDef("reference", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("type", typeDefs[48], "", false, nil),
	})
	// properties of RelatedArtifact
	typeDefs[33].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type markdown
		common.NewPropDef("citation", typeDefs[676], "", false, nil),
		// with type string
		common.NewPropDef("display", typeDefs[46], "", false, nil),
		// with type Attachment
		common.NewPropDef("document", typeDefs[4], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("label", typeDefs[46], "", false, nil),
		// with type canonical
		common.NewPropDef("resource", typeDefs[673], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"documentation",
			"justification",
			"citation",
			"predecessor",
			"successor",
			"derived-from",
			"depends-on",
			"composed-of",
		}),
		// with type url
		common.NewPropDef("url", typeDefs[680], "", false, nil),
	})
	// properties of RelatedPerson
	typeDefs[518].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("active", typeDefs[40], "", false, nil),
		// with type Address
		common.NewPropDef("address", typeDefs[2], "", true, nil),
		// with type date
		common.NewPropDef("birthDate", typeDefs[41], "", false, nil),
		// with type RelatedPerson_Communication
		common.NewPropDef("communication", typeDefs[519], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("gender", typeDefs[46], "", false, []string{
			"male",
			"female",
			"other",
			"unknown",
		}),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type HumanName
		common.NewPropDef("name", typeDefs[17], "", true, nil),
		// with type Reference
		common.NewPropDef("patient", typeDefs[32], "", false, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type Attachment
		common.NewPropDef("photo", typeDefs[4], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("relationship", typeDefs[8], "", true, nil),
		// with type ContactPoint
		common.NewPropDef("telecom", typeDefs[11], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of RelatedPerson_Communication
	typeDefs[519].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("language", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("preferred", typeDefs[40], "", false, nil),
	})
	// properties of RequestGroup
	typeDefs[520].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type RequestGroup_Action
		common.NewPropDef("action", typeDefs[521], "", true, nil),
		// with type Reference
		common.NewPropDef("author", typeDefs[32], "", false, nil),
		// with type dateTime
		common.NewPropDef("authoredOn", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("basedOn", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Identifier
		common.NewPropDef("groupIdentifier", typeDefs[18], "", false, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type canonical
		common.NewPropDef("instantiatesCanonical", typeDefs[673], "", true, nil),
		// with type uri
		common.NewPropDef("instantiatesUri", typeDefs[48], "", true, nil),
		// with type code
		common.NewPropDef("intent", typeDefs[674], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type code
		common.NewPropDef("priority", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("reasonCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("reasonReference", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("replaces", typeDefs[32], "", true, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of RequestGroup_Action
	typeDefs[521].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type RequestGroup_Action
		common.NewPropDef("action", typeDefs[521], "", true, nil),
		// with type code
		common.NewPropDef("cardinalityBehavior", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", true, nil),
		// with type RequestGroup_Condition
		common.NewPropDef("condition", typeDefs[522], "", true, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type RelatedArtifact
		common.NewPropDef("documentation", typeDefs[33], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type code
		common.NewPropDef("groupingBehavior", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("participant", typeDefs[32], "", true, nil),
		// with type code
		common.NewPropDef("precheckBehavior", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("prefix", typeDefs[46], "", false, nil),
		// with type code
		common.NewPropDef("priority", typeDefs[674], "", false, nil),
		// with type RequestGroup_RelatedAction
		common.NewPropDef("relatedAction", typeDefs[523], "", true, nil),
		// with type code
		common.NewPropDef("requiredBehavior", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("resource", typeDefs[32], "", false, nil),
		// with type code
		common.NewPropDef("selectionBehavior", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("textEquivalent", typeDefs[46], "", false, nil),
		// with type Age
		common.NewPropDef("timingAge", typeDefs[59], "", false, nil),
		// with type dateTime
		common.NewPropDef("timingDateTime", typeDefs[42], "", false, nil),
		// with type Duration
		common.NewPropDef("timingDuration", typeDefs[236], "", false, nil),
		// with type Period
		common.NewPropDef("timingPeriod", typeDefs[25], "", false, nil),
		// with type Range
		common.NewPropDef("timingRange", typeDefs[30], "", false, nil),
		// with type Timing
		common.NewPropDef("timingTiming", typeDefs[655], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of RequestGroup_Condition
	typeDefs[522].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Expression
		common.NewPropDef("expression", typeDefs[15], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type code
		common.NewPropDef("kind", typeDefs[674], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of RequestGroup_RelatedAction
	typeDefs[523].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type id
		common.NewPropDef("actionId", typeDefs[675], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Duration
		common.NewPropDef("offsetDuration", typeDefs[236], "", false, nil),
		// with type Range
		common.NewPropDef("offsetRange", typeDefs[30], "", false, nil),
		// with type code
		common.NewPropDef("relationship", typeDefs[674], "", false, nil),
	})
	// properties of ResearchDefinition
	typeDefs[524].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type date
		common.NewPropDef("approvalDate", typeDefs[41], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("author", typeDefs[10], "", true, nil),
		// with type string
		common.NewPropDef("comment", typeDefs[46], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("editor", typeDefs[10], "", true, nil),
		// with type Period
		common.NewPropDef("effectivePeriod", typeDefs[25], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("endorser", typeDefs[10], "", true, nil),
		// with type boolean
		common.NewPropDef("experimental", typeDefs[40], "", false, nil),
		// with type Reference
		common.NewPropDef("exposure", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("exposureAlternative", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type date
		common.NewPropDef("lastReviewDate", typeDefs[41], "", false, nil),
		// with type canonical
		common.NewPropDef("library", typeDefs[673], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("outcome", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("population", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("purpose", typeDefs[676], "", false, nil),
		// with type RelatedArtifact
		common.NewPropDef("relatedArtifact", typeDefs[33], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("reviewer", typeDefs[10], "", true, nil),
		// with type string
		common.NewPropDef("shortTitle", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type CodeableConcept
		common.NewPropDef("subjectCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("subjectReference", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("subtitle", typeDefs[46], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("topic", typeDefs[8], "", true, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type string
		common.NewPropDef("usage", typeDefs[46], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of ResearchElementDefinition
	typeDefs[525].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type date
		common.NewPropDef("approvalDate", typeDefs[41], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("author", typeDefs[10], "", true, nil),
		// with type ResearchElementDefinition_Characteristic
		common.NewPropDef("characteristic", typeDefs[526], "", true, nil),
		// with type string
		common.NewPropDef("comment", typeDefs[46], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("editor", typeDefs[10], "", true, nil),
		// with type Period
		common.NewPropDef("effectivePeriod", typeDefs[25], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("endorser", typeDefs[10], "", true, nil),
		// with type boolean
		common.NewPropDef("experimental", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type date
		common.NewPropDef("lastReviewDate", typeDefs[41], "", false, nil),
		// with type canonical
		common.NewPropDef("library", typeDefs[673], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("purpose", typeDefs[676], "", false, nil),
		// with type RelatedArtifact
		common.NewPropDef("relatedArtifact", typeDefs[33], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("reviewer", typeDefs[10], "", true, nil),
		// with type string
		common.NewPropDef("shortTitle", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type CodeableConcept
		common.NewPropDef("subjectCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("subjectReference", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("subtitle", typeDefs[46], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("topic", typeDefs[8], "", true, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"population",
			"exposure",
			"outcome",
		}),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type string
		common.NewPropDef("usage", typeDefs[46], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("variableType", typeDefs[46], "", false, []string{
			"dichotomous",
			"continuous",
			"descriptive",
		}),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of ResearchElementDefinition_Characteristic
	typeDefs[526].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type canonical
		common.NewPropDef("definitionCanonical", typeDefs[673], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("definitionCodeableConcept", typeDefs[8], "", false, nil),
		// with type DataRequirement
		common.NewPropDef("definitionDataRequirement", typeDefs[13], "", false, nil),
		// with type Expression
		common.NewPropDef("definitionExpression", typeDefs[15], "", false, nil),
		// with type boolean
		common.NewPropDef("exclude", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type dateTime
		common.NewPropDef("participantEffectiveDateTime", typeDefs[42], "", false, nil),
		// with type string
		common.NewPropDef("participantEffectiveDescription", typeDefs[46], "", false, nil),
		// with type Duration
		common.NewPropDef("participantEffectiveDuration", typeDefs[236], "", false, nil),
		// with type string
		common.NewPropDef("participantEffectiveGroupMeasure", typeDefs[46], "", false, []string{
			"mean",
			"median",
			"mean-of-mean",
			"mean-of-median",
			"median-of-mean",
			"median-of-median",
		}),
		// with type Period
		common.NewPropDef("participantEffectivePeriod", typeDefs[25], "", false, nil),
		// with type Duration
		common.NewPropDef("participantEffectiveTimeFromStart", typeDefs[236], "", false, nil),
		// with type Timing
		common.NewPropDef("participantEffectiveTiming", typeDefs[655], "", false, nil),
		// with type dateTime
		common.NewPropDef("studyEffectiveDateTime", typeDefs[42], "", false, nil),
		// with type string
		common.NewPropDef("studyEffectiveDescription", typeDefs[46], "", false, nil),
		// with type Duration
		common.NewPropDef("studyEffectiveDuration", typeDefs[236], "", false, nil),
		// with type string
		common.NewPropDef("studyEffectiveGroupMeasure", typeDefs[46], "", false, []string{
			"mean",
			"median",
			"mean-of-mean",
			"mean-of-median",
			"median-of-mean",
			"median-of-median",
		}),
		// with type Period
		common.NewPropDef("studyEffectivePeriod", typeDefs[25], "", false, nil),
		// with type Duration
		common.NewPropDef("studyEffectiveTimeFromStart", typeDefs[236], "", false, nil),
		// with type Timing
		common.NewPropDef("studyEffectiveTiming", typeDefs[655], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("unitOfMeasure", typeDefs[8], "", false, nil),
		// with type UsageContext
		common.NewPropDef("usageContext", typeDefs[38], "", true, nil),
	})
	// properties of ResearchStudy
	typeDefs[527].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ResearchStudy_Arm
		common.NewPropDef("arm", typeDefs[528], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("condition", typeDefs[8], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type Reference
		common.NewPropDef("enrollment", typeDefs[32], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("focus", typeDefs[8], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("keyword", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("location", typeDefs[8], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type ResearchStudy_Objective
		common.NewPropDef("objective", typeDefs[529], "", true, nil),
		// with type Reference
		common.NewPropDef("partOf", typeDefs[32], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("phase", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("primaryPurposeType", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("principalInvestigator", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("protocol", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("reasonStopped", typeDefs[8], "", false, nil),
		// with type RelatedArtifact
		common.NewPropDef("relatedArtifact", typeDefs[33], "", true, nil),
		// with type Reference
		common.NewPropDef("site", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("sponsor", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"active",
			"administratively-completed",
			"approved",
			"closed-to-accrual",
			"closed-to-accrual-and-intervention",
			"completed",
			"disapproved",
			"in-review",
			"temporarily-closed-to-accrual",
			"temporarily-closed-to-accrual-and-intervention",
			"withdrawn",
		}),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
	})
	// properties of ResearchStudy_Arm
	typeDefs[528].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of ResearchStudy_Objective
	typeDefs[529].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of ResearchSubject
	typeDefs[530].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("actualArm", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("assignedArm", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("consent", typeDefs[32], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type Reference
		common.NewPropDef("individual", typeDefs[32], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"candidate",
			"eligible",
			"follow-up",
			"ineligible",
			"not-registered",
			"off-study",
			"on-study",
			"on-study-intervention",
			"on-study-observation",
			"pending-on-study",
			"potential-candidate",
			"screening",
			"withdrawn",
		}),
		// with type Reference
		common.NewPropDef("study", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of Resource
	typeDefs[1].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
	})
	// properties of RiskAssessment
	typeDefs[531].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("basedOn", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("basis", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("condition", typeDefs[32], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("method", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("mitigation", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type dateTime
		common.NewPropDef("occurrenceDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("occurrencePeriod", typeDefs[25], "", false, nil),
		// with type Reference
		common.NewPropDef("parent", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("performer", typeDefs[32], "", false, nil),
		// with type RiskAssessment_Prediction
		common.NewPropDef("prediction", typeDefs[532], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("reasonCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("reasonReference", typeDefs[32], "", true, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of RiskAssessment_Prediction
	typeDefs[532].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("outcome", typeDefs[8], "", false, nil),
		// with type decimal
		common.NewPropDef("probabilityDecimal", typeDefs[43], "", false, nil),
		// with type Range
		common.NewPropDef("probabilityRange", typeDefs[30], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("qualitativeRisk", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("rationale", typeDefs[46], "", false, nil),
		// with type decimal
		common.NewPropDef("relativeRisk", typeDefs[43], "", false, nil),
		// with type Period
		common.NewPropDef("whenPeriod", typeDefs[25], "", false, nil),
		// with type Range
		common.NewPropDef("whenRange", typeDefs[30], "", false, nil),
	})
	// properties of RiskEvidenceSynthesis
	typeDefs[533].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type date
		common.NewPropDef("approvalDate", typeDefs[41], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("author", typeDefs[10], "", true, nil),
		// with type RiskEvidenceSynthesis_Certainty
		common.NewPropDef("certainty", typeDefs[534], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("editor", typeDefs[10], "", true, nil),
		// with type Period
		common.NewPropDef("effectivePeriod", typeDefs[25], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("endorser", typeDefs[10], "", true, nil),
		// with type Reference
		common.NewPropDef("exposure", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type date
		common.NewPropDef("lastReviewDate", typeDefs[41], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type Reference
		common.NewPropDef("outcome", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("population", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type RelatedArtifact
		common.NewPropDef("relatedArtifact", typeDefs[33], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("reviewer", typeDefs[10], "", true, nil),
		// with type RiskEvidenceSynthesis_RiskEstimate
		common.NewPropDef("riskEstimate", typeDefs[537], "", false, nil),
		// with type RiskEvidenceSynthesis_SampleSize
		common.NewPropDef("sampleSize", typeDefs[538], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type CodeableConcept
		common.NewPropDef("studyType", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("synthesisType", typeDefs[8], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("topic", typeDefs[8], "", true, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of RiskEvidenceSynthesis_Certainty
	typeDefs[534].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type RiskEvidenceSynthesis_CertaintySubcomponent
		common.NewPropDef("certaintySubcomponent", typeDefs[535], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("rating", typeDefs[8], "", true, nil),
	})
	// properties of RiskEvidenceSynthesis_CertaintySubcomponent
	typeDefs[535].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("rating", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of RiskEvidenceSynthesis_PrecisionEstimate
	typeDefs[536].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("from", typeDefs[43], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type decimal
		common.NewPropDef("level", typeDefs[43], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("to", typeDefs[43], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of RiskEvidenceSynthesis_RiskEstimate
	typeDefs[537].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type integer
		common.NewPropDef("denominatorCount", typeDefs[45], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type integer
		common.NewPropDef("numeratorCount", typeDefs[45], "", false, nil),
		// with type RiskEvidenceSynthesis_PrecisionEstimate
		common.NewPropDef("precisionEstimate", typeDefs[536], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("unitOfMeasure", typeDefs[8], "", false, nil),
		// with type decimal
		common.NewPropDef("value", typeDefs[43], "", false, nil),
	})
	// properties of RiskEvidenceSynthesis_SampleSize
	typeDefs[538].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type integer
		common.NewPropDef("numberOfParticipants", typeDefs[45], "", false, nil),
		// with type integer
		common.NewPropDef("numberOfStudies", typeDefs[45], "", false, nil),
	})
	// properties of SampledData
	typeDefs[34].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("data", typeDefs[46], "", false, nil),
		// with type positiveInt
		common.NewPropDef("dimensions", typeDefs[678], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type decimal
		common.NewPropDef("factor", typeDefs[43], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type decimal
		common.NewPropDef("lowerLimit", typeDefs[43], "", false, nil),
		// with type Quantity
		common.NewPropDef("origin", typeDefs[29], "", false, nil),
		// with type decimal
		common.NewPropDef("period", typeDefs[43], "", false, nil),
		// with type decimal
		common.NewPropDef("upperLimit", typeDefs[43], "", false, nil),
	})
	// properties of Schedule
	typeDefs[539].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("active", typeDefs[40], "", false, nil),
		// with type Reference
		common.NewPropDef("actor", typeDefs[32], "", true, nil),
		// with type string
		common.NewPropDef("comment", typeDefs[46], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Period
		common.NewPropDef("planningHorizon", typeDefs[25], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("serviceCategory", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("serviceType", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("specialty", typeDefs[8], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of SearchParameter
	typeDefs[540].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("base", typeDefs[674], "", true, nil),
		// with type string
		common.NewPropDef("chain", typeDefs[46], "", true, nil),
		// with type code
		common.NewPropDef("code", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("comparator", typeDefs[46], "", true, []string{
			"eq",
			"ne",
			"gt",
			"lt",
			"ge",
			"le",
			"sa",
			"eb",
			"ap",
		}),
		// with type SearchParameter_Component
		common.NewPropDef("component", typeDefs[541], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type canonical
		common.NewPropDef("derivedFrom", typeDefs[673], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type boolean
		common.NewPropDef("experimental", typeDefs[40], "", false, nil),
		// with type string
		common.NewPropDef("expression", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type string
		common.NewPropDef("modifier", typeDefs[46], "", true, []string{
			"missing",
			"exact",
			"contains",
			"not",
			"text",
			"in",
			"not-in",
			"below",
			"above",
			"type",
			"identifier",
			"ofType",
		}),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("multipleAnd", typeDefs[40], "", false, nil),
		// with type boolean
		common.NewPropDef("multipleOr", typeDefs[40], "", false, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("purpose", typeDefs[676], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type code
		common.NewPropDef("target", typeDefs[674], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"number",
			"date",
			"string",
			"token",
			"reference",
			"composite",
			"quantity",
			"uri",
			"special",
		}),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("xpath", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("xpathUsage", typeDefs[46], "", false, []string{
			"normal",
			"phonetic",
			"nearby",
			"distance",
			"other",
		}),
	})
	// properties of SearchParameter_Component
	typeDefs[541].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type canonical
		common.NewPropDef("definition", typeDefs[673], "", false, nil),
		// with type string
		common.NewPropDef("expression", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of ServiceRequest
	typeDefs[542].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("asNeededBoolean", typeDefs[40], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("asNeededCodeableConcept", typeDefs[8], "", false, nil),
		// with type dateTime
		common.NewPropDef("authoredOn", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("basedOn", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("bodySite", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type boolean
		common.NewPropDef("doNotPerform", typeDefs[40], "", false, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type canonical
		common.NewPropDef("instantiatesCanonical", typeDefs[673], "", true, nil),
		// with type uri
		common.NewPropDef("instantiatesUri", typeDefs[48], "", true, nil),
		// with type Reference
		common.NewPropDef("insurance", typeDefs[32], "", true, nil),
		// with type code
		common.NewPropDef("intent", typeDefs[674], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("locationCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("locationReference", typeDefs[32], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type dateTime
		common.NewPropDef("occurrenceDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("occurrencePeriod", typeDefs[25], "", false, nil),
		// with type Timing
		common.NewPropDef("occurrenceTiming", typeDefs[655], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("orderDetail", typeDefs[8], "", true, nil),
		// with type string
		common.NewPropDef("patientInstruction", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("performer", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("performerType", typeDefs[8], "", false, nil),
		// with type code
		common.NewPropDef("priority", typeDefs[674], "", false, nil),
		// with type Quantity
		common.NewPropDef("quantityQuantity", typeDefs[29], "", false, nil),
		// with type Range
		common.NewPropDef("quantityRange", typeDefs[30], "", false, nil),
		// with type Ratio
		common.NewPropDef("quantityRatio", typeDefs[31], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("reasonCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("reasonReference", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("relevantHistory", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("replaces", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("requester", typeDefs[32], "", false, nil),
		// with type Identifier
		common.NewPropDef("requisition", typeDefs[18], "", false, nil),
		// with type Reference
		common.NewPropDef("specimen", typeDefs[32], "", true, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("supportingInfo", typeDefs[32], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of Signature
	typeDefs[35].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type base64Binary
		common.NewPropDef("data", typeDefs[39], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("onBehalfOf", typeDefs[32], "", false, nil),
		// with type code
		common.NewPropDef("sigFormat", typeDefs[674], "", false, nil),
		// with type code
		common.NewPropDef("targetFormat", typeDefs[674], "", false, nil),
		// with type Coding
		common.NewPropDef("type", typeDefs[9], "", true, nil),
		// with type instant
		common.NewPropDef("when", typeDefs[44], "", false, nil),
		// with type Reference
		common.NewPropDef("who", typeDefs[32], "", false, nil),
	})
	// properties of Slot
	typeDefs[543].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("appointmentType", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("comment", typeDefs[46], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type instant
		common.NewPropDef("end", typeDefs[44], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("overbooked", typeDefs[40], "", false, nil),
		// with type Reference
		common.NewPropDef("schedule", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("serviceCategory", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("serviceType", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("specialty", typeDefs[8], "", true, nil),
		// with type instant
		common.NewPropDef("start", typeDefs[44], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"busy",
			"free",
			"busy-unavailable",
			"busy-tentative",
			"entered-in-error",
		}),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of Specimen
	typeDefs[544].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Identifier
		common.NewPropDef("accessionIdentifier", typeDefs[18], "", false, nil),
		// with type Specimen_Collection
		common.NewPropDef("collection", typeDefs[550], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("condition", typeDefs[8], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Specimen_Container
		common.NewPropDef("container", typeDefs[551], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type Reference
		common.NewPropDef("parent", typeDefs[32], "", true, nil),
		// with type Specimen_Processing
		common.NewPropDef("processing", typeDefs[552], "", true, nil),
		// with type dateTime
		common.NewPropDef("receivedTime", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("request", typeDefs[32], "", true, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"available",
			"unavailable",
			"unsatisfactory",
			"entered-in-error",
		}),
		// with type Reference
		common.NewPropDef("subject", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of SpecimenDefinition
	typeDefs[545].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("collection", typeDefs[8], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("patientPreparation", typeDefs[8], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("timeAspect", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("typeCollected", typeDefs[8], "", false, nil),
		// with type SpecimenDefinition_TypeTested
		common.NewPropDef("typeTested", typeDefs[549], "", true, nil),
	})
	// properties of SpecimenDefinition_Additive
	typeDefs[546].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("additiveCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("additiveReference", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of SpecimenDefinition_Container
	typeDefs[547].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type SpecimenDefinition_Additive
		common.NewPropDef("additive", typeDefs[546], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("cap", typeDefs[8], "", false, nil),
		// with type Quantity
		common.NewPropDef("capacity", typeDefs[29], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("material", typeDefs[8], "", false, nil),
		// with type Quantity
		common.NewPropDef("minimumVolumeQuantity", typeDefs[29], "", false, nil),
		// with type string
		common.NewPropDef("minimumVolumeString", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("preparation", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of SpecimenDefinition_Handling
	typeDefs[548].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("instruction", typeDefs[46], "", false, nil),
		// with type Duration
		common.NewPropDef("maxDuration", typeDefs[236], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("temperatureQualifier", typeDefs[8], "", false, nil),
		// with type Range
		common.NewPropDef("temperatureRange", typeDefs[30], "", false, nil),
	})
	// properties of SpecimenDefinition_TypeTested
	typeDefs[549].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type SpecimenDefinition_Container
		common.NewPropDef("container", typeDefs[547], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type SpecimenDefinition_Handling
		common.NewPropDef("handling", typeDefs[548], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("isDerived", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("preference", typeDefs[46], "", false, []string{
			"preferred",
			"alternate",
		}),
		// with type CodeableConcept
		common.NewPropDef("rejectionCriterion", typeDefs[8], "", true, nil),
		// with type string
		common.NewPropDef("requirement", typeDefs[46], "", false, nil),
		// with type Duration
		common.NewPropDef("retentionTime", typeDefs[236], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of Specimen_Collection
	typeDefs[550].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("bodySite", typeDefs[8], "", false, nil),
		// with type dateTime
		common.NewPropDef("collectedDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("collectedPeriod", typeDefs[25], "", false, nil),
		// with type Reference
		common.NewPropDef("collector", typeDefs[32], "", false, nil),
		// with type Duration
		common.NewPropDef("duration", typeDefs[236], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("fastingStatusCodeableConcept", typeDefs[8], "", false, nil),
		// with type Duration
		common.NewPropDef("fastingStatusDuration", typeDefs[236], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("method", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
	})
	// properties of Specimen_Container
	typeDefs[551].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("additiveCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("additiveReference", typeDefs[32], "", false, nil),
		// with type Quantity
		common.NewPropDef("capacity", typeDefs[29], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Quantity
		common.NewPropDef("specimenQuantity", typeDefs[29], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of Specimen_Processing
	typeDefs[552].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("additive", typeDefs[32], "", true, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("procedure", typeDefs[8], "", false, nil),
		// with type dateTime
		common.NewPropDef("timeDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("timePeriod", typeDefs[25], "", false, nil),
	})
	// properties of StructureDefinition
	typeDefs[553].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("abstract", typeDefs[40], "", false, nil),
		// with type canonical
		common.NewPropDef("baseDefinition", typeDefs[673], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type StructureDefinition_Context
		common.NewPropDef("context", typeDefs[554], "", true, nil),
		// with type string
		common.NewPropDef("contextInvariant", typeDefs[46], "", true, nil),
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type string
		common.NewPropDef("derivation", typeDefs[46], "", false, []string{
			"specialization",
			"constraint",
		}),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type StructureDefinition_Differential
		common.NewPropDef("differential", typeDefs[555], "", false, nil),
		// with type boolean
		common.NewPropDef("experimental", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("fhirVersion", typeDefs[46], "", false, []string{
			"0.01",
			"0.05",
			"0.06",
			"0.11",
			"0.0.80",
			"0.0.81",
			"0.0.82",
			"0.4.0",
			"0.5.0",
			"1.0.0",
			"1.0.1",
			"1.0.2",
			"1.1.0",
			"1.4.0",
			"1.6.0",
			"1.8.0",
			"3.0.0",
			"3.0.1",
			"3.3.0",
			"3.5.0",
			"4.0.0",
			"4.0.1",
		}),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type Coding
		common.NewPropDef("keyword", typeDefs[9], "", true, nil),
		// with type string
		common.NewPropDef("kind", typeDefs[46], "", false, []string{
			"primitive-type",
			"complex-type",
			"resource",
			"logical",
		}),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type StructureDefinition_Mapping
		common.NewPropDef("mapping", typeDefs[556], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("purpose", typeDefs[676], "", false, nil),
		// with type StructureDefinition_Snapshot
		common.NewPropDef("snapshot", typeDefs[557], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("type", typeDefs[48], "", false, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of StructureDefinition_Context
	typeDefs[554].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("expression", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"fhirpath",
			"element",
			"extension",
		}),
	})
	// properties of StructureDefinition_Differential
	typeDefs[555].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ElementDefinition
		common.NewPropDef("element", typeDefs[244], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of StructureDefinition_Mapping
	typeDefs[556].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("comment", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type id
		common.NewPropDef("identity", typeDefs[675], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("uri", typeDefs[48], "", false, nil),
	})
	// properties of StructureDefinition_Snapshot
	typeDefs[557].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ElementDefinition
		common.NewPropDef("element", typeDefs[244], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of StructureMap
	typeDefs[558].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type boolean
		common.NewPropDef("experimental", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type StructureMap_Group
		common.NewPropDef("group", typeDefs[560], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type canonical
		common.NewPropDef("import", typeDefs[673], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("purpose", typeDefs[676], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type StructureMap_Structure
		common.NewPropDef("structure", typeDefs[565], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of StructureMap_Dependent
	typeDefs[559].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("name", typeDefs[675], "", false, nil),
		// with type string
		common.NewPropDef("variable", typeDefs[46], "", true, nil),
	})
	// properties of StructureMap_Group
	typeDefs[560].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("documentation", typeDefs[46], "", false, nil),
		// with type id
		common.NewPropDef("extends", typeDefs[675], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type StructureMap_Input
		common.NewPropDef("input", typeDefs[561], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("name", typeDefs[675], "", false, nil),
		// with type StructureMap_Rule
		common.NewPropDef("rule", typeDefs[563], "", true, nil),
		// with type string
		common.NewPropDef("typeMode", typeDefs[46], "", false, []string{
			"none",
			"types",
			"type-and-types",
		}),
	})
	// properties of StructureMap_Input
	typeDefs[561].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("documentation", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("mode", typeDefs[46], "", false, []string{
			"source",
			"target",
		}),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("name", typeDefs[675], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, nil),
	})
	// properties of StructureMap_Parameter
	typeDefs[562].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("valueBoolean", typeDefs[40], "value", false, nil),
		// with type decimal
		common.NewPropDef("valueDecimal", typeDefs[43], "value", false, nil),
		// with type id
		common.NewPropDef("valueId", typeDefs[675], "value", false, nil),
		// with type integer
		common.NewPropDef("valueInteger", typeDefs[45], "value", false, nil),
		// with type string
		common.NewPropDef("valueString", typeDefs[46], "value", false, nil),
	})
	// properties of StructureMap_Rule
	typeDefs[563].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type StructureMap_Dependent
		common.NewPropDef("dependent", typeDefs[559], "", true, nil),
		// with type string
		common.NewPropDef("documentation", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("name", typeDefs[675], "", false, nil),
		// with type StructureMap_Rule
		common.NewPropDef("rule", typeDefs[563], "", true, nil),
		// with type StructureMap_Source
		common.NewPropDef("source", typeDefs[564], "", true, nil),
		// with type StructureMap_Target
		common.NewPropDef("target", typeDefs[566], "", true, nil),
	})
	// properties of StructureMap_Source
	typeDefs[564].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("check", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("condition", typeDefs[46], "", false, nil),
		// with type id
		common.NewPropDef("context", typeDefs[675], "", false, nil),
		// with type Address
		common.NewPropDef("defaultValueAddress", typeDefs[2], "defaultValue", false, nil),
		// with type Age
		common.NewPropDef("defaultValueAge", typeDefs[59], "defaultValue", false, nil),
		// with type Annotation
		common.NewPropDef("defaultValueAnnotation", typeDefs[3], "defaultValue", false, nil),
		// with type Attachment
		common.NewPropDef("defaultValueAttachment", typeDefs[4], "defaultValue", false, nil),
		// with type base64Binary
		common.NewPropDef("defaultValueBase64Binary", typeDefs[39], "defaultValue", false, nil),
		// with type boolean
		common.NewPropDef("defaultValueBoolean", typeDefs[40], "defaultValue", false, nil),
		// with type canonical
		common.NewPropDef("defaultValueCanonical", typeDefs[673], "defaultValue", false, nil),
		// with type code
		common.NewPropDef("defaultValueCode", typeDefs[674], "defaultValue", false, nil),
		// with type CodeableConcept
		common.NewPropDef("defaultValueCodeableConcept", typeDefs[8], "defaultValue", false, nil),
		// with type Coding
		common.NewPropDef("defaultValueCoding", typeDefs[9], "defaultValue", false, nil),
		// with type ContactDetail
		common.NewPropDef("defaultValueContactDetail", typeDefs[10], "defaultValue", false, nil),
		// with type ContactPoint
		common.NewPropDef("defaultValueContactPoint", typeDefs[11], "defaultValue", false, nil),
		// with type Contributor
		common.NewPropDef("defaultValueContributor", typeDefs[12], "defaultValue", false, nil),
		// with type Count
		common.NewPropDef("defaultValueCount", typeDefs[186], "defaultValue", false, nil),
		// with type DataRequirement
		common.NewPropDef("defaultValueDataRequirement", typeDefs[13], "defaultValue", false, nil),
		// with type date
		common.NewPropDef("defaultValueDate", typeDefs[41], "defaultValue", false, nil),
		// with type dateTime
		common.NewPropDef("defaultValueDateTime", typeDefs[42], "defaultValue", false, nil),
		// with type decimal
		common.NewPropDef("defaultValueDecimal", typeDefs[43], "defaultValue", false, nil),
		// with type Distance
		common.NewPropDef("defaultValueDistance", typeDefs[227], "defaultValue", false, nil),
		// with type Dosage
		common.NewPropDef("defaultValueDosage", typeDefs[234], "defaultValue", false, nil),
		// with type Duration
		common.NewPropDef("defaultValueDuration", typeDefs[236], "defaultValue", false, nil),
		// with type Expression
		common.NewPropDef("defaultValueExpression", typeDefs[15], "defaultValue", false, nil),
		// with type HumanName
		common.NewPropDef("defaultValueHumanName", typeDefs[17], "defaultValue", false, nil),
		// with type id
		common.NewPropDef("defaultValueId", typeDefs[675], "defaultValue", false, nil),
		// with type Identifier
		common.NewPropDef("defaultValueIdentifier", typeDefs[18], "defaultValue", false, nil),
		// with type instant
		common.NewPropDef("defaultValueInstant", typeDefs[44], "defaultValue", false, nil),
		// with type integer
		common.NewPropDef("defaultValueInteger", typeDefs[45], "defaultValue", false, nil),
		// with type markdown
		common.NewPropDef("defaultValueMarkdown", typeDefs[676], "defaultValue", false, nil),
		// with type Meta
		common.NewPropDef("defaultValueMeta", typeDefs[20], "defaultValue", false, nil),
		// with type Money
		common.NewPropDef("defaultValueMoney", typeDefs[21], "defaultValue", false, nil),
		// with type oid
		common.NewPropDef("defaultValueOid", typeDefs[677], "defaultValue", false, nil),
		// with type ParameterDefinition
		common.NewPropDef("defaultValueParameterDefinition", typeDefs[23], "defaultValue", false, nil),
		// with type Period
		common.NewPropDef("defaultValuePeriod", typeDefs[25], "defaultValue", false, nil),
		// with type positiveInt
		common.NewPropDef("defaultValuePositiveInt", typeDefs[678], "defaultValue", false, nil),
		// with type Quantity
		common.NewPropDef("defaultValueQuantity", typeDefs[29], "defaultValue", false, nil),
		// with type Range
		common.NewPropDef("defaultValueRange", typeDefs[30], "defaultValue", false, nil),
		// with type Ratio
		common.NewPropDef("defaultValueRatio", typeDefs[31], "defaultValue", false, nil),
		// with type Reference
		common.NewPropDef("defaultValueReference", typeDefs[32], "defaultValue", false, nil),
		// with type RelatedArtifact
		common.NewPropDef("defaultValueRelatedArtifact", typeDefs[33], "defaultValue", false, nil),
		// with type SampledData
		common.NewPropDef("defaultValueSampledData", typeDefs[34], "defaultValue", false, nil),
		// with type Signature
		common.NewPropDef("defaultValueSignature", typeDefs[35], "defaultValue", false, nil),
		// with type string
		common.NewPropDef("defaultValueString", typeDefs[46], "defaultValue", false, nil),
		// with type time
		common.NewPropDef("defaultValueTime", typeDefs[47], "defaultValue", false, nil),
		// with type Timing
		common.NewPropDef("defaultValueTiming", typeDefs[655], "defaultValue", false, nil),
		// with type TriggerDefinition
		common.NewPropDef("defaultValueTriggerDefinition", typeDefs[37], "defaultValue", false, nil),
		// with type unsignedInt
		common.NewPropDef("defaultValueUnsignedInt", typeDefs[679], "defaultValue", false, nil),
		// with type uri
		common.NewPropDef("defaultValueUri", typeDefs[48], "defaultValue", false, nil),
		// with type url
		common.NewPropDef("defaultValueUrl", typeDefs[680], "defaultValue", false, nil),
		// with type UsageContext
		common.NewPropDef("defaultValueUsageContext", typeDefs[38], "defaultValue", false, nil),
		// with type uuid
		common.NewPropDef("defaultValueUuid", typeDefs[681], "defaultValue", false, nil),
		// with type string
		common.NewPropDef("element", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("listMode", typeDefs[46], "", false, []string{
			"first",
			"not_first",
			"last",
			"not_last",
			"only_one",
		}),
		// with type string
		common.NewPropDef("logMessage", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("max", typeDefs[46], "", false, nil),
		// with type integer
		common.NewPropDef("min", typeDefs[45], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, nil),
		// with type id
		common.NewPropDef("variable", typeDefs[675], "", false, nil),
	})
	// properties of StructureMap_Structure
	typeDefs[565].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("alias", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("documentation", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("mode", typeDefs[46], "", false, []string{
			"source",
			"queried",
			"target",
			"produced",
		}),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type canonical
		common.NewPropDef("url", typeDefs[673], "", false, nil),
	})
	// properties of StructureMap_Target
	typeDefs[566].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type id
		common.NewPropDef("context", typeDefs[675], "", false, nil),
		// with type string
		common.NewPropDef("contextType", typeDefs[46], "", false, []string{
			"type",
			"variable",
		}),
		// with type string
		common.NewPropDef("element", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("listMode", typeDefs[46], "", true, []string{
			"first",
			"share",
			"last",
			"collate",
		}),
		// with type id
		common.NewPropDef("listRuleId", typeDefs[675], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type StructureMap_Parameter
		common.NewPropDef("parameter", typeDefs[562], "", true, nil),
		// with type string
		common.NewPropDef("transform", typeDefs[46], "", false, []string{
			"create",
			"copy",
			"truncate",
			"escape",
			"cast",
			"append",
			"translate",
			"reference",
			"dateOp",
			"uuid",
			"pointer",
			"evaluate",
			"cc",
			"c",
			"qty",
			"id",
			"cp",
		}),
		// with type id
		common.NewPropDef("variable", typeDefs[675], "", false, nil),
	})
	// properties of Subscription
	typeDefs[567].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Subscription_Channel
		common.NewPropDef("channel", typeDefs[568], "", false, nil),
		// with type ContactPoint
		common.NewPropDef("contact", typeDefs[11], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type string
		common.NewPropDef("criteria", typeDefs[46], "", false, nil),
		// with type instant
		common.NewPropDef("end", typeDefs[44], "", false, nil),
		// with type string
		common.NewPropDef("error", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("reason", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"requested",
			"active",
			"error",
			"off",
		}),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of Subscription_Channel
	typeDefs[568].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type url
		common.NewPropDef("endpoint", typeDefs[680], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("header", typeDefs[46], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type code
		common.NewPropDef("payload", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"rest-hook",
			"websocket",
			"email",
			"sms",
			"message",
		}),
	})
	// properties of Substance
	typeDefs[569].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type Substance_Ingredient
		common.NewPropDef("ingredient", typeDefs[607], "", true, nil),
		// with type Substance_Instance
		common.NewPropDef("instance", typeDefs[608], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"active",
			"inactive",
			"entered-in-error",
		}),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of SubstanceAmount
	typeDefs[36].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Quantity
		common.NewPropDef("amountQuantity", typeDefs[29], "", false, nil),
		// with type Range
		common.NewPropDef("amountRange", typeDefs[30], "", false, nil),
		// with type string
		common.NewPropDef("amountString", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("amountText", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("amountType", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type SubstanceAmount_ReferenceRange
		common.NewPropDef("referenceRange", typeDefs[570], "", false, nil),
	})
	// properties of SubstanceAmount_ReferenceRange
	typeDefs[570].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Quantity
		common.NewPropDef("highLimit", typeDefs[29], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Quantity
		common.NewPropDef("lowLimit", typeDefs[29], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of SubstanceNucleicAcid
	typeDefs[571].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("areaOfHybridisation", typeDefs[46], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type integer
		common.NewPropDef("numberOfSubunits", typeDefs[45], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("oligoNucleotideType", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("sequenceType", typeDefs[8], "", false, nil),
		// with type SubstanceNucleicAcid_Subunit
		common.NewPropDef("subunit", typeDefs[573], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of SubstanceNucleicAcid_Linkage
	typeDefs[572].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("connectivity", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("residueSite", typeDefs[46], "", false, nil),
	})
	// properties of SubstanceNucleicAcid_Subunit
	typeDefs[573].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("fivePrime", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type integer
		common.NewPropDef("length", typeDefs[45], "", false, nil),
		// with type SubstanceNucleicAcid_Linkage
		common.NewPropDef("linkage", typeDefs[572], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("sequence", typeDefs[46], "", false, nil),
		// with type Attachment
		common.NewPropDef("sequenceAttachment", typeDefs[4], "", false, nil),
		// with type integer
		common.NewPropDef("subunit", typeDefs[45], "", false, nil),
		// with type SubstanceNucleicAcid_Sugar
		common.NewPropDef("sugar", typeDefs[574], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("threePrime", typeDefs[8], "", false, nil),
	})
	// properties of SubstanceNucleicAcid_Sugar
	typeDefs[574].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("residueSite", typeDefs[46], "", false, nil),
	})
	// properties of SubstancePolymer
	typeDefs[575].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("class", typeDefs[8], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("copolymerConnectivity", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("geometry", typeDefs[8], "", false, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type string
		common.NewPropDef("modification", typeDefs[46], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type SubstancePolymer_MonomerSet
		common.NewPropDef("monomerSet", typeDefs[577], "", true, nil),
		// with type SubstancePolymer_Repeat
		common.NewPropDef("repeat", typeDefs[578], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of SubstancePolymer_DegreeOfPolymerisation
	typeDefs[576].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type SubstanceAmount
		common.NewPropDef("amount", typeDefs[36], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("degree", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of SubstancePolymer_MonomerSet
	typeDefs[577].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("ratioType", typeDefs[8], "", false, nil),
		// with type SubstancePolymer_StartingMaterial
		common.NewPropDef("startingMaterial", typeDefs[580], "", true, nil),
	})
	// properties of SubstancePolymer_Repeat
	typeDefs[578].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("averageMolecularFormula", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type integer
		common.NewPropDef("numberOfUnits", typeDefs[45], "", false, nil),
		// with type SubstancePolymer_RepeatUnit
		common.NewPropDef("repeatUnit", typeDefs[579], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("repeatUnitAmountType", typeDefs[8], "", false, nil),
	})
	// properties of SubstancePolymer_RepeatUnit
	typeDefs[579].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type SubstanceAmount
		common.NewPropDef("amount", typeDefs[36], "", false, nil),
		// with type SubstancePolymer_DegreeOfPolymerisation
		common.NewPropDef("degreeOfPolymerisation", typeDefs[576], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("orientationOfPolymerisation", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("repeatUnit", typeDefs[46], "", false, nil),
		// with type SubstancePolymer_StructuralRepresentation
		common.NewPropDef("structuralRepresentation", typeDefs[581], "", true, nil),
	})
	// properties of SubstancePolymer_StartingMaterial
	typeDefs[580].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type SubstanceAmount
		common.NewPropDef("amount", typeDefs[36], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("isDefining", typeDefs[40], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("material", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of SubstancePolymer_StructuralRepresentation
	typeDefs[581].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Attachment
		common.NewPropDef("attachment", typeDefs[4], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("representation", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of SubstanceProtein
	typeDefs[582].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type string
		common.NewPropDef("disulfideLinkage", typeDefs[46], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type integer
		common.NewPropDef("numberOfSubunits", typeDefs[45], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("sequenceType", typeDefs[8], "", false, nil),
		// with type SubstanceProtein_Subunit
		common.NewPropDef("subunit", typeDefs[583], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of SubstanceProtein_Subunit
	typeDefs[583].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("cTerminalModification", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("cTerminalModificationId", typeDefs[18], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type integer
		common.NewPropDef("length", typeDefs[45], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("nTerminalModification", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("nTerminalModificationId", typeDefs[18], "", false, nil),
		// with type string
		common.NewPropDef("sequence", typeDefs[46], "", false, nil),
		// with type Attachment
		common.NewPropDef("sequenceAttachment", typeDefs[4], "", false, nil),
		// with type integer
		common.NewPropDef("subunit", typeDefs[45], "", false, nil),
	})
	// properties of SubstanceReferenceInformation
	typeDefs[584].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type SubstanceReferenceInformation_Classification
		common.NewPropDef("classification", typeDefs[585], "", true, nil),
		// with type string
		common.NewPropDef("comment", typeDefs[46], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type SubstanceReferenceInformation_Gene
		common.NewPropDef("gene", typeDefs[586], "", true, nil),
		// with type SubstanceReferenceInformation_GeneElement
		common.NewPropDef("geneElement", typeDefs[587], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type SubstanceReferenceInformation_Target
		common.NewPropDef("target", typeDefs[588], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of SubstanceReferenceInformation_Classification
	typeDefs[585].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("classification", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("domain", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("source", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("subtype", typeDefs[8], "", true, nil),
	})
	// properties of SubstanceReferenceInformation_Gene
	typeDefs[586].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("gene", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("geneSequenceOrigin", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("source", typeDefs[32], "", true, nil),
	})
	// properties of SubstanceReferenceInformation_GeneElement
	typeDefs[587].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Identifier
		common.NewPropDef("element", typeDefs[18], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("source", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of SubstanceReferenceInformation_Target
	typeDefs[588].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Quantity
		common.NewPropDef("amountQuantity", typeDefs[29], "", false, nil),
		// with type Range
		common.NewPropDef("amountRange", typeDefs[30], "", false, nil),
		// with type string
		common.NewPropDef("amountString", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("amountType", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("interaction", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("organism", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("organismType", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("source", typeDefs[32], "", true, nil),
		// with type Identifier
		common.NewPropDef("target", typeDefs[18], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of SubstanceSourceMaterial
	typeDefs[589].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("countryOfOrigin", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("developmentStage", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type SubstanceSourceMaterial_FractionDescription
		common.NewPropDef("fractionDescription", typeDefs[591], "", true, nil),
		// with type string
		common.NewPropDef("geographicalLocation", typeDefs[46], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type SubstanceSourceMaterial_Organism
		common.NewPropDef("organism", typeDefs[593], "", false, nil),
		// with type Identifier
		common.NewPropDef("organismId", typeDefs[18], "", false, nil),
		// with type string
		common.NewPropDef("organismName", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("parentSubstanceId", typeDefs[18], "", true, nil),
		// with type string
		common.NewPropDef("parentSubstanceName", typeDefs[46], "", true, nil),
		// with type SubstanceSourceMaterial_PartDescription
		common.NewPropDef("partDescription", typeDefs[595], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("sourceMaterialClass", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("sourceMaterialState", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("sourceMaterialType", typeDefs[8], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of SubstanceSourceMaterial_Author
	typeDefs[590].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("authorDescription", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("authorType", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of SubstanceSourceMaterial_FractionDescription
	typeDefs[591].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("fraction", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("materialType", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of SubstanceSourceMaterial_Hybrid
	typeDefs[592].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("hybridType", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("maternalOrganismId", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("maternalOrganismName", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("paternalOrganismId", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("paternalOrganismName", typeDefs[46], "", false, nil),
	})
	// properties of SubstanceSourceMaterial_Organism
	typeDefs[593].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type SubstanceSourceMaterial_Author
		common.NewPropDef("author", typeDefs[590], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("family", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("genus", typeDefs[8], "", false, nil),
		// with type SubstanceSourceMaterial_Hybrid
		common.NewPropDef("hybrid", typeDefs[592], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("intraspecificDescription", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("intraspecificType", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type SubstanceSourceMaterial_OrganismGeneral
		common.NewPropDef("organismGeneral", typeDefs[594], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("species", typeDefs[8], "", false, nil),
	})
	// properties of SubstanceSourceMaterial_OrganismGeneral
	typeDefs[594].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("class", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("kingdom", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("order", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("phylum", typeDefs[8], "", false, nil),
	})
	// properties of SubstanceSourceMaterial_PartDescription
	typeDefs[595].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("part", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("partLocation", typeDefs[8], "", false, nil),
	})
	// properties of SubstanceSpecification
	typeDefs[596].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type SubstanceSpecification_Code
		common.NewPropDef("code", typeDefs[597], "", true, nil),
		// with type string
		common.NewPropDef("comment", typeDefs[46], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("domain", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type SubstanceSpecification_Moiety
		common.NewPropDef("moiety", typeDefs[599], "", true, nil),
		// with type SubstanceSpecification_MolecularWeight
		common.NewPropDef("molecularWeight", typeDefs[600], "", true, nil),
		// with type SubstanceSpecification_Name
		common.NewPropDef("name", typeDefs[601], "", true, nil),
		// with type Reference
		common.NewPropDef("nucleicAcid", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("polymer", typeDefs[32], "", false, nil),
		// with type SubstanceSpecification_Property
		common.NewPropDef("property", typeDefs[603], "", true, nil),
		// with type Reference
		common.NewPropDef("protein", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("referenceInformation", typeDefs[32], "", false, nil),
		// with type SubstanceSpecification_Relationship
		common.NewPropDef("relationship", typeDefs[604], "", true, nil),
		// with type Reference
		common.NewPropDef("source", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("sourceMaterial", typeDefs[32], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("status", typeDefs[8], "", false, nil),
		// with type SubstanceSpecification_Structure
		common.NewPropDef("structure", typeDefs[606], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of SubstanceSpecification_Code
	typeDefs[597].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type string
		common.NewPropDef("comment", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("source", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("status", typeDefs[8], "", false, nil),
		// with type dateTime
		common.NewPropDef("statusDate", typeDefs[42], "", false, nil),
	})
	// properties of SubstanceSpecification_Isotope
	typeDefs[598].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Quantity
		common.NewPropDef("halfLife", typeDefs[29], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type SubstanceSpecification_MolecularWeight
		common.NewPropDef("molecularWeight", typeDefs[600], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("name", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("substitution", typeDefs[8], "", false, nil),
	})
	// properties of SubstanceSpecification_Moiety
	typeDefs[599].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Quantity
		common.NewPropDef("amountQuantity", typeDefs[29], "", false, nil),
		// with type string
		common.NewPropDef("amountString", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("molecularFormula", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("opticalActivity", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("role", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("stereochemistry", typeDefs[8], "", false, nil),
	})
	// properties of SubstanceSpecification_MolecularWeight
	typeDefs[600].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Quantity
		common.NewPropDef("amount", typeDefs[29], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("method", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of SubstanceSpecification_Name
	typeDefs[601].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("domain", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("language", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type SubstanceSpecification_Official
		common.NewPropDef("official", typeDefs[602], "", true, nil),
		// with type boolean
		common.NewPropDef("preferred", typeDefs[40], "", false, nil),
		// with type Reference
		common.NewPropDef("source", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("status", typeDefs[8], "", false, nil),
		// with type SubstanceSpecification_Name
		common.NewPropDef("synonym", typeDefs[601], "", true, nil),
		// with type SubstanceSpecification_Name
		common.NewPropDef("translation", typeDefs[601], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of SubstanceSpecification_Official
	typeDefs[602].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("authority", typeDefs[8], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("status", typeDefs[8], "", false, nil),
	})
	// properties of SubstanceSpecification_Property
	typeDefs[603].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Quantity
		common.NewPropDef("amountQuantity", typeDefs[29], "", false, nil),
		// with type string
		common.NewPropDef("amountString", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("definingSubstanceCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("definingSubstanceReference", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("parameters", typeDefs[46], "", false, nil),
	})
	// properties of SubstanceSpecification_Relationship
	typeDefs[604].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Quantity
		common.NewPropDef("amountQuantity", typeDefs[29], "", false, nil),
		// with type Range
		common.NewPropDef("amountRange", typeDefs[30], "", false, nil),
		// with type Ratio
		common.NewPropDef("amountRatio", typeDefs[31], "", false, nil),
		// with type Ratio
		common.NewPropDef("amountRatioLowLimit", typeDefs[31], "", false, nil),
		// with type string
		common.NewPropDef("amountString", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("amountType", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("isDefining", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("relationship", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("source", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("substanceCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("substanceReference", typeDefs[32], "", false, nil),
	})
	// properties of SubstanceSpecification_Representation
	typeDefs[605].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Attachment
		common.NewPropDef("attachment", typeDefs[4], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("representation", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of SubstanceSpecification_Structure
	typeDefs[606].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type SubstanceSpecification_Isotope
		common.NewPropDef("isotope", typeDefs[598], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("molecularFormula", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("molecularFormulaByMoiety", typeDefs[46], "", false, nil),
		// with type SubstanceSpecification_MolecularWeight
		common.NewPropDef("molecularWeight", typeDefs[600], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("opticalActivity", typeDefs[8], "", false, nil),
		// with type SubstanceSpecification_Representation
		common.NewPropDef("representation", typeDefs[605], "", true, nil),
		// with type Reference
		common.NewPropDef("source", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("stereochemistry", typeDefs[8], "", false, nil),
	})
	// properties of Substance_Ingredient
	typeDefs[607].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Ratio
		common.NewPropDef("quantity", typeDefs[31], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("substanceCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("substanceReference", typeDefs[32], "", false, nil),
	})
	// properties of Substance_Instance
	typeDefs[608].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type dateTime
		common.NewPropDef("expiry", typeDefs[42], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
	})
	// properties of SupplyDelivery
	typeDefs[609].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Reference
		common.NewPropDef("basedOn", typeDefs[32], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("destination", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type dateTime
		common.NewPropDef("occurrenceDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("occurrencePeriod", typeDefs[25], "", false, nil),
		// with type Timing
		common.NewPropDef("occurrenceTiming", typeDefs[655], "", false, nil),
		// with type Reference
		common.NewPropDef("partOf", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("patient", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("receiver", typeDefs[32], "", true, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"in-progress",
			"completed",
			"abandoned",
			"entered-in-error",
		}),
		// with type SupplyDelivery_SuppliedItem
		common.NewPropDef("suppliedItem", typeDefs[610], "", false, nil),
		// with type Reference
		common.NewPropDef("supplier", typeDefs[32], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
	})
	// properties of SupplyDelivery_SuppliedItem
	typeDefs[610].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("itemCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("itemReference", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
	})
	// properties of SupplyRequest
	typeDefs[611].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type dateTime
		common.NewPropDef("authoredOn", typeDefs[42], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("category", typeDefs[8], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Reference
		common.NewPropDef("deliverFrom", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("deliverTo", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("itemCodeableConcept", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("itemReference", typeDefs[32], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type dateTime
		common.NewPropDef("occurrenceDateTime", typeDefs[42], "", false, nil),
		// with type Period
		common.NewPropDef("occurrencePeriod", typeDefs[25], "", false, nil),
		// with type Timing
		common.NewPropDef("occurrenceTiming", typeDefs[655], "", false, nil),
		// with type SupplyRequest_Parameter
		common.NewPropDef("parameter", typeDefs[612], "", true, nil),
		// with type code
		common.NewPropDef("priority", typeDefs[674], "", false, nil),
		// with type Quantity
		common.NewPropDef("quantity", typeDefs[29], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("reasonCode", typeDefs[8], "", true, nil),
		// with type Reference
		common.NewPropDef("reasonReference", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("requester", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"suspended",
			"cancelled",
			"completed",
			"entered-in-error",
			"unknown",
		}),
		// with type Reference
		common.NewPropDef("supplier", typeDefs[32], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of SupplyRequest_Parameter
	typeDefs[612].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("valueBoolean", typeDefs[40], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("valueCodeableConcept", typeDefs[8], "", false, nil),
		// with type Quantity
		common.NewPropDef("valueQuantity", typeDefs[29], "", false, nil),
		// with type Range
		common.NewPropDef("valueRange", typeDefs[30], "", false, nil),
	})
	// properties of Task
	typeDefs[613].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type dateTime
		common.NewPropDef("authoredOn", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("basedOn", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("businessStatus", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type Period
		common.NewPropDef("executionPeriod", typeDefs[25], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("focus", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("for", typeDefs[32], "", false, nil),
		// with type Identifier
		common.NewPropDef("groupIdentifier", typeDefs[18], "", false, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type Task_Input
		common.NewPropDef("input", typeDefs[614], "", true, nil),
		// with type canonical
		common.NewPropDef("instantiatesCanonical", typeDefs[673], "", false, nil),
		// with type uri
		common.NewPropDef("instantiatesUri", typeDefs[48], "", false, nil),
		// with type Reference
		common.NewPropDef("insurance", typeDefs[32], "", true, nil),
		// with type string
		common.NewPropDef("intent", typeDefs[46], "", false, []string{
			"unknown",
			"proposal",
			"plan",
			"order",
			"original-order",
			"reflex-order",
			"filler-order",
			"instance-order",
			"option",
		}),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type dateTime
		common.NewPropDef("lastModified", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("location", typeDefs[32], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type Task_Output
		common.NewPropDef("output", typeDefs[615], "", true, nil),
		// with type Reference
		common.NewPropDef("owner", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("partOf", typeDefs[32], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("performerType", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("priority", typeDefs[674], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("reasonCode", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("reasonReference", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("relevantHistory", typeDefs[32], "", true, nil),
		// with type Reference
		common.NewPropDef("requester", typeDefs[32], "", false, nil),
		// with type Task_Restriction
		common.NewPropDef("restriction", typeDefs[616], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"requested",
			"received",
			"accepted",
			"rejected",
			"ready",
			"cancelled",
			"in-progress",
			"on-hold",
			"failed",
			"completed",
			"entered-in-error",
		}),
		// with type CodeableConcept
		common.NewPropDef("statusReason", typeDefs[8], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of Task_Input
	typeDefs[614].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type Address
		common.NewPropDef("valueAddress", typeDefs[2], "value", false, nil),
		// with type Age
		common.NewPropDef("valueAge", typeDefs[59], "value", false, nil),
		// with type Annotation
		common.NewPropDef("valueAnnotation", typeDefs[3], "value", false, nil),
		// with type Attachment
		common.NewPropDef("valueAttachment", typeDefs[4], "value", false, nil),
		// with type base64Binary
		common.NewPropDef("valueBase64Binary", typeDefs[39], "value", false, nil),
		// with type boolean
		common.NewPropDef("valueBoolean", typeDefs[40], "value", false, nil),
		// with type canonical
		common.NewPropDef("valueCanonical", typeDefs[673], "value", false, nil),
		// with type code
		common.NewPropDef("valueCode", typeDefs[674], "value", false, nil),
		// with type CodeableConcept
		common.NewPropDef("valueCodeableConcept", typeDefs[8], "value", false, nil),
		// with type Coding
		common.NewPropDef("valueCoding", typeDefs[9], "value", false, nil),
		// with type ContactDetail
		common.NewPropDef("valueContactDetail", typeDefs[10], "value", false, nil),
		// with type ContactPoint
		common.NewPropDef("valueContactPoint", typeDefs[11], "value", false, nil),
		// with type Contributor
		common.NewPropDef("valueContributor", typeDefs[12], "value", false, nil),
		// with type Count
		common.NewPropDef("valueCount", typeDefs[186], "value", false, nil),
		// with type DataRequirement
		common.NewPropDef("valueDataRequirement", typeDefs[13], "value", false, nil),
		// with type date
		common.NewPropDef("valueDate", typeDefs[41], "value", false, nil),
		// with type dateTime
		common.NewPropDef("valueDateTime", typeDefs[42], "value", false, nil),
		// with type decimal
		common.NewPropDef("valueDecimal", typeDefs[43], "value", false, nil),
		// with type Distance
		common.NewPropDef("valueDistance", typeDefs[227], "value", false, nil),
		// with type Dosage
		common.NewPropDef("valueDosage", typeDefs[234], "value", false, nil),
		// with type Duration
		common.NewPropDef("valueDuration", typeDefs[236], "value", false, nil),
		// with type Expression
		common.NewPropDef("valueExpression", typeDefs[15], "value", false, nil),
		// with type HumanName
		common.NewPropDef("valueHumanName", typeDefs[17], "value", false, nil),
		// with type id
		common.NewPropDef("valueId", typeDefs[675], "value", false, nil),
		// with type Identifier
		common.NewPropDef("valueIdentifier", typeDefs[18], "value", false, nil),
		// with type instant
		common.NewPropDef("valueInstant", typeDefs[44], "value", false, nil),
		// with type integer
		common.NewPropDef("valueInteger", typeDefs[45], "value", false, nil),
		// with type markdown
		common.NewPropDef("valueMarkdown", typeDefs[676], "value", false, nil),
		// with type Meta
		common.NewPropDef("valueMeta", typeDefs[20], "value", false, nil),
		// with type Money
		common.NewPropDef("valueMoney", typeDefs[21], "value", false, nil),
		// with type oid
		common.NewPropDef("valueOid", typeDefs[677], "value", false, nil),
		// with type ParameterDefinition
		common.NewPropDef("valueParameterDefinition", typeDefs[23], "value", false, nil),
		// with type Period
		common.NewPropDef("valuePeriod", typeDefs[25], "value", false, nil),
		// with type positiveInt
		common.NewPropDef("valuePositiveInt", typeDefs[678], "value", false, nil),
		// with type Quantity
		common.NewPropDef("valueQuantity", typeDefs[29], "value", false, nil),
		// with type Range
		common.NewPropDef("valueRange", typeDefs[30], "value", false, nil),
		// with type Ratio
		common.NewPropDef("valueRatio", typeDefs[31], "value", false, nil),
		// with type Reference
		common.NewPropDef("valueReference", typeDefs[32], "value", false, nil),
		// with type RelatedArtifact
		common.NewPropDef("valueRelatedArtifact", typeDefs[33], "value", false, nil),
		// with type SampledData
		common.NewPropDef("valueSampledData", typeDefs[34], "value", false, nil),
		// with type Signature
		common.NewPropDef("valueSignature", typeDefs[35], "value", false, nil),
		// with type string
		common.NewPropDef("valueString", typeDefs[46], "value", false, nil),
		// with type time
		common.NewPropDef("valueTime", typeDefs[47], "value", false, nil),
		// with type Timing
		common.NewPropDef("valueTiming", typeDefs[655], "value", false, nil),
		// with type TriggerDefinition
		common.NewPropDef("valueTriggerDefinition", typeDefs[37], "value", false, nil),
		// with type unsignedInt
		common.NewPropDef("valueUnsignedInt", typeDefs[679], "value", false, nil),
		// with type uri
		common.NewPropDef("valueUri", typeDefs[48], "value", false, nil),
		// with type url
		common.NewPropDef("valueUrl", typeDefs[680], "value", false, nil),
		// with type UsageContext
		common.NewPropDef("valueUsageContext", typeDefs[38], "value", false, nil),
		// with type uuid
		common.NewPropDef("valueUuid", typeDefs[681], "value", false, nil),
	})
	// properties of Task_Output
	typeDefs[615].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", false, nil),
		// with type Address
		common.NewPropDef("valueAddress", typeDefs[2], "value", false, nil),
		// with type Age
		common.NewPropDef("valueAge", typeDefs[59], "value", false, nil),
		// with type Annotation
		common.NewPropDef("valueAnnotation", typeDefs[3], "value", false, nil),
		// with type Attachment
		common.NewPropDef("valueAttachment", typeDefs[4], "value", false, nil),
		// with type base64Binary
		common.NewPropDef("valueBase64Binary", typeDefs[39], "value", false, nil),
		// with type boolean
		common.NewPropDef("valueBoolean", typeDefs[40], "value", false, nil),
		// with type canonical
		common.NewPropDef("valueCanonical", typeDefs[673], "value", false, nil),
		// with type code
		common.NewPropDef("valueCode", typeDefs[674], "value", false, nil),
		// with type CodeableConcept
		common.NewPropDef("valueCodeableConcept", typeDefs[8], "value", false, nil),
		// with type Coding
		common.NewPropDef("valueCoding", typeDefs[9], "value", false, nil),
		// with type ContactDetail
		common.NewPropDef("valueContactDetail", typeDefs[10], "value", false, nil),
		// with type ContactPoint
		common.NewPropDef("valueContactPoint", typeDefs[11], "value", false, nil),
		// with type Contributor
		common.NewPropDef("valueContributor", typeDefs[12], "value", false, nil),
		// with type Count
		common.NewPropDef("valueCount", typeDefs[186], "value", false, nil),
		// with type DataRequirement
		common.NewPropDef("valueDataRequirement", typeDefs[13], "value", false, nil),
		// with type date
		common.NewPropDef("valueDate", typeDefs[41], "value", false, nil),
		// with type dateTime
		common.NewPropDef("valueDateTime", typeDefs[42], "value", false, nil),
		// with type decimal
		common.NewPropDef("valueDecimal", typeDefs[43], "value", false, nil),
		// with type Distance
		common.NewPropDef("valueDistance", typeDefs[227], "value", false, nil),
		// with type Dosage
		common.NewPropDef("valueDosage", typeDefs[234], "value", false, nil),
		// with type Duration
		common.NewPropDef("valueDuration", typeDefs[236], "value", false, nil),
		// with type Expression
		common.NewPropDef("valueExpression", typeDefs[15], "value", false, nil),
		// with type HumanName
		common.NewPropDef("valueHumanName", typeDefs[17], "value", false, nil),
		// with type id
		common.NewPropDef("valueId", typeDefs[675], "value", false, nil),
		// with type Identifier
		common.NewPropDef("valueIdentifier", typeDefs[18], "value", false, nil),
		// with type instant
		common.NewPropDef("valueInstant", typeDefs[44], "value", false, nil),
		// with type integer
		common.NewPropDef("valueInteger", typeDefs[45], "value", false, nil),
		// with type markdown
		common.NewPropDef("valueMarkdown", typeDefs[676], "value", false, nil),
		// with type Meta
		common.NewPropDef("valueMeta", typeDefs[20], "value", false, nil),
		// with type Money
		common.NewPropDef("valueMoney", typeDefs[21], "value", false, nil),
		// with type oid
		common.NewPropDef("valueOid", typeDefs[677], "value", false, nil),
		// with type ParameterDefinition
		common.NewPropDef("valueParameterDefinition", typeDefs[23], "value", false, nil),
		// with type Period
		common.NewPropDef("valuePeriod", typeDefs[25], "value", false, nil),
		// with type positiveInt
		common.NewPropDef("valuePositiveInt", typeDefs[678], "value", false, nil),
		// with type Quantity
		common.NewPropDef("valueQuantity", typeDefs[29], "value", false, nil),
		// with type Range
		common.NewPropDef("valueRange", typeDefs[30], "value", false, nil),
		// with type Ratio
		common.NewPropDef("valueRatio", typeDefs[31], "value", false, nil),
		// with type Reference
		common.NewPropDef("valueReference", typeDefs[32], "value", false, nil),
		// with type RelatedArtifact
		common.NewPropDef("valueRelatedArtifact", typeDefs[33], "value", false, nil),
		// with type SampledData
		common.NewPropDef("valueSampledData", typeDefs[34], "value", false, nil),
		// with type Signature
		common.NewPropDef("valueSignature", typeDefs[35], "value", false, nil),
		// with type string
		common.NewPropDef("valueString", typeDefs[46], "value", false, nil),
		// with type time
		common.NewPropDef("valueTime", typeDefs[47], "value", false, nil),
		// with type Timing
		common.NewPropDef("valueTiming", typeDefs[655], "value", false, nil),
		// with type TriggerDefinition
		common.NewPropDef("valueTriggerDefinition", typeDefs[37], "value", false, nil),
		// with type unsignedInt
		common.NewPropDef("valueUnsignedInt", typeDefs[679], "value", false, nil),
		// with type uri
		common.NewPropDef("valueUri", typeDefs[48], "value", false, nil),
		// with type url
		common.NewPropDef("valueUrl", typeDefs[680], "value", false, nil),
		// with type UsageContext
		common.NewPropDef("valueUsageContext", typeDefs[38], "value", false, nil),
		// with type uuid
		common.NewPropDef("valueUuid", typeDefs[681], "value", false, nil),
	})
	// properties of Task_Restriction
	typeDefs[616].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Period
		common.NewPropDef("period", typeDefs[25], "", false, nil),
		// with type Reference
		common.NewPropDef("recipient", typeDefs[32], "", true, nil),
		// with type positiveInt
		common.NewPropDef("repetitions", typeDefs[678], "", false, nil),
	})
	// properties of TerminologyCapabilities
	typeDefs[617].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type TerminologyCapabilities_Closure
		common.NewPropDef("closure", typeDefs[618], "", false, nil),
		// with type string
		common.NewPropDef("codeSearch", typeDefs[46], "", false, []string{
			"explicit",
			"all",
		}),
		// with type TerminologyCapabilities_CodeSystem
		common.NewPropDef("codeSystem", typeDefs[619], "", true, nil),
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type TerminologyCapabilities_Expansion
		common.NewPropDef("expansion", typeDefs[620], "", false, nil),
		// with type boolean
		common.NewPropDef("experimental", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type TerminologyCapabilities_Implementation
		common.NewPropDef("implementation", typeDefs[622], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("kind", typeDefs[674], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type boolean
		common.NewPropDef("lockedDate", typeDefs[40], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("purpose", typeDefs[676], "", false, nil),
		// with type TerminologyCapabilities_Software
		common.NewPropDef("software", typeDefs[624], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type TerminologyCapabilities_Translation
		common.NewPropDef("translation", typeDefs[625], "", false, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type TerminologyCapabilities_ValidateCode
		common.NewPropDef("validateCode", typeDefs[626], "", false, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of TerminologyCapabilities_Closure
	typeDefs[618].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("translation", typeDefs[40], "", false, nil),
	})
	// properties of TerminologyCapabilities_CodeSystem
	typeDefs[619].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("subsumption", typeDefs[40], "", false, nil),
		// with type canonical
		common.NewPropDef("uri", typeDefs[673], "", false, nil),
		// with type TerminologyCapabilities_Version
		common.NewPropDef("version", typeDefs[627], "", true, nil),
	})
	// properties of TerminologyCapabilities_Expansion
	typeDefs[620].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("hierarchical", typeDefs[40], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("incomplete", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("paging", typeDefs[40], "", false, nil),
		// with type TerminologyCapabilities_Parameter
		common.NewPropDef("parameter", typeDefs[623], "", true, nil),
		// with type markdown
		common.NewPropDef("textFilter", typeDefs[676], "", false, nil),
	})
	// properties of TerminologyCapabilities_Filter
	typeDefs[621].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("code", typeDefs[674], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type code
		common.NewPropDef("op", typeDefs[674], "", true, nil),
	})
	// properties of TerminologyCapabilities_Implementation
	typeDefs[622].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type url
		common.NewPropDef("url", typeDefs[680], "", false, nil),
	})
	// properties of TerminologyCapabilities_Parameter
	typeDefs[623].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("documentation", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type code
		common.NewPropDef("name", typeDefs[674], "", false, nil),
	})
	// properties of TerminologyCapabilities_Software
	typeDefs[624].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of TerminologyCapabilities_Translation
	typeDefs[625].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("needsMap", typeDefs[40], "", false, nil),
	})
	// properties of TerminologyCapabilities_ValidateCode
	typeDefs[626].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("translations", typeDefs[40], "", false, nil),
	})
	// properties of TerminologyCapabilities_Version
	typeDefs[627].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("code", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("compositional", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type TerminologyCapabilities_Filter
		common.NewPropDef("filter", typeDefs[621], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("isDefault", typeDefs[40], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type code
		common.NewPropDef("property", typeDefs[674], "", true, nil),
	})
	// properties of TestReport
	typeDefs[628].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type dateTime
		common.NewPropDef("issued", typeDefs[42], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type TestReport_Participant
		common.NewPropDef("participant", typeDefs[634], "", true, nil),
		// with type string
		common.NewPropDef("result", typeDefs[46], "", false, []string{
			"pass",
			"fail",
			"pending",
		}),
		// with type decimal
		common.NewPropDef("score", typeDefs[43], "", false, nil),
		// with type TestReport_Setup
		common.NewPropDef("setup", typeDefs[635], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"completed",
			"in-progress",
			"waiting",
			"stopped",
			"entered-in-error",
		}),
		// with type TestReport_Teardown
		common.NewPropDef("teardown", typeDefs[636], "", false, nil),
		// with type TestReport_Test
		common.NewPropDef("test", typeDefs[637], "", true, nil),
		// with type Reference
		common.NewPropDef("testScript", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("tester", typeDefs[46], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of TestReport_Action
	typeDefs[629].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type TestReport_Assert
		common.NewPropDef("assert", typeDefs[632], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type TestReport_Operation
		common.NewPropDef("operation", typeDefs[633], "", false, nil),
	})
	// properties of TestReport_Action1
	typeDefs[630].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type TestReport_Assert
		common.NewPropDef("assert", typeDefs[632], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type TestReport_Operation
		common.NewPropDef("operation", typeDefs[633], "", false, nil),
	})
	// properties of TestReport_Action2
	typeDefs[631].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type TestReport_Operation
		common.NewPropDef("operation", typeDefs[633], "", false, nil),
	})
	// properties of TestReport_Assert
	typeDefs[632].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("detail", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("message", typeDefs[676], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("result", typeDefs[46], "", false, []string{
			"pass",
			"skip",
			"fail",
			"warning",
			"error",
		}),
	})
	// properties of TestReport_Operation
	typeDefs[633].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type uri
		common.NewPropDef("detail", typeDefs[48], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("message", typeDefs[676], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("result", typeDefs[46], "", false, []string{
			"pass",
			"skip",
			"fail",
			"warning",
			"error",
		}),
	})
	// properties of TestReport_Participant
	typeDefs[634].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("display", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"test-engine",
			"client",
			"server",
		}),
		// with type uri
		common.NewPropDef("uri", typeDefs[48], "", false, nil),
	})
	// properties of TestReport_Setup
	typeDefs[635].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type TestReport_Action
		common.NewPropDef("action", typeDefs[629], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of TestReport_Teardown
	typeDefs[636].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type TestReport_Action2
		common.NewPropDef("action", typeDefs[631], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of TestReport_Test
	typeDefs[637].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type TestReport_Action1
		common.NewPropDef("action", typeDefs[630], "", true, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
	})
	// properties of TestScript
	typeDefs[638].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type TestScript_Destination
		common.NewPropDef("destination", typeDefs[644], "", true, nil),
		// with type boolean
		common.NewPropDef("experimental", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type TestScript_Fixture
		common.NewPropDef("fixture", typeDefs[645], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type TestScript_Metadata
		common.NewPropDef("metadata", typeDefs[647], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type TestScript_Origin
		common.NewPropDef("origin", typeDefs[649], "", true, nil),
		// with type Reference
		common.NewPropDef("profile", typeDefs[32], "", true, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("purpose", typeDefs[676], "", false, nil),
		// with type TestScript_Setup
		common.NewPropDef("setup", typeDefs[651], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type TestScript_Teardown
		common.NewPropDef("teardown", typeDefs[652], "", false, nil),
		// with type TestScript_Test
		common.NewPropDef("test", typeDefs[653], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type TestScript_Variable
		common.NewPropDef("variable", typeDefs[654], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of TestScript_Action
	typeDefs[639].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type TestScript_Assert
		common.NewPropDef("assert", typeDefs[642], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type TestScript_Operation
		common.NewPropDef("operation", typeDefs[648], "", false, nil),
	})
	// properties of TestScript_Action1
	typeDefs[640].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type TestScript_Assert
		common.NewPropDef("assert", typeDefs[642], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type TestScript_Operation
		common.NewPropDef("operation", typeDefs[648], "", false, nil),
	})
	// properties of TestScript_Action2
	typeDefs[641].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type TestScript_Operation
		common.NewPropDef("operation", typeDefs[648], "", false, nil),
	})
	// properties of TestScript_Assert
	typeDefs[642].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("compareToSourceExpression", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("compareToSourceId", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("compareToSourcePath", typeDefs[46], "", false, nil),
		// with type code
		common.NewPropDef("contentType", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("direction", typeDefs[46], "", false, []string{
			"response",
			"request",
		}),
		// with type string
		common.NewPropDef("expression", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("headerField", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("label", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("minimumId", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type boolean
		common.NewPropDef("navigationLinks", typeDefs[40], "", false, nil),
		// with type string
		common.NewPropDef("operator", typeDefs[46], "", false, []string{
			"equals",
			"notEquals",
			"in",
			"notIn",
			"greaterThan",
			"lessThan",
			"empty",
			"notEmpty",
			"contains",
			"notContains",
			"eval",
		}),
		// with type string
		common.NewPropDef("path", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("requestMethod", typeDefs[46], "", false, []string{
			"delete",
			"get",
			"options",
			"patch",
			"post",
			"put",
			"head",
		}),
		// with type string
		common.NewPropDef("requestURL", typeDefs[46], "", false, nil),
		// with type code
		common.NewPropDef("resource", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("response", typeDefs[46], "", false, []string{
			"okay",
			"created",
			"noContent",
			"notModified",
			"bad",
			"forbidden",
			"notFound",
			"methodNotAllowed",
			"conflict",
			"gone",
			"preconditionFailed",
			"unprocessable",
		}),
		// with type string
		common.NewPropDef("responseCode", typeDefs[46], "", false, nil),
		// with type id
		common.NewPropDef("sourceId", typeDefs[675], "", false, nil),
		// with type id
		common.NewPropDef("validateProfileId", typeDefs[675], "", false, nil),
		// with type string
		common.NewPropDef("value", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("warningOnly", typeDefs[40], "", false, nil),
	})
	// properties of TestScript_Capability
	typeDefs[643].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type canonical
		common.NewPropDef("capabilities", typeDefs[673], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type integer
		common.NewPropDef("destination", typeDefs[45], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("link", typeDefs[48], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type integer
		common.NewPropDef("origin", typeDefs[45], "", true, nil),
		// with type boolean
		common.NewPropDef("required", typeDefs[40], "", false, nil),
		// with type boolean
		common.NewPropDef("validated", typeDefs[40], "", false, nil),
	})
	// properties of TestScript_Destination
	typeDefs[644].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type integer
		common.NewPropDef("index", typeDefs[45], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Coding
		common.NewPropDef("profile", typeDefs[9], "", false, nil),
	})
	// properties of TestScript_Fixture
	typeDefs[645].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("autocreate", typeDefs[40], "", false, nil),
		// with type boolean
		common.NewPropDef("autodelete", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("resource", typeDefs[32], "", false, nil),
	})
	// properties of TestScript_Link
	typeDefs[646].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
	})
	// properties of TestScript_Metadata
	typeDefs[647].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type TestScript_Capability
		common.NewPropDef("capability", typeDefs[643], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type TestScript_Link
		common.NewPropDef("link", typeDefs[646], "", true, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of TestScript_Operation
	typeDefs[648].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("accept", typeDefs[674], "", false, nil),
		// with type code
		common.NewPropDef("contentType", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type integer
		common.NewPropDef("destination", typeDefs[45], "", false, nil),
		// with type boolean
		common.NewPropDef("encodeRequestUrl", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("label", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("method", typeDefs[46], "", false, []string{
			"delete",
			"get",
			"options",
			"patch",
			"post",
			"put",
			"head",
		}),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type integer
		common.NewPropDef("origin", typeDefs[45], "", false, nil),
		// with type string
		common.NewPropDef("params", typeDefs[46], "", false, nil),
		// with type TestScript_RequestHeader
		common.NewPropDef("requestHeader", typeDefs[650], "", true, nil),
		// with type id
		common.NewPropDef("requestId", typeDefs[675], "", false, nil),
		// with type code
		common.NewPropDef("resource", typeDefs[674], "", false, nil),
		// with type id
		common.NewPropDef("responseId", typeDefs[675], "", false, nil),
		// with type id
		common.NewPropDef("sourceId", typeDefs[675], "", false, nil),
		// with type id
		common.NewPropDef("targetId", typeDefs[675], "", false, nil),
		// with type Coding
		common.NewPropDef("type", typeDefs[9], "", false, nil),
		// with type string
		common.NewPropDef("url", typeDefs[46], "", false, nil),
	})
	// properties of TestScript_Origin
	typeDefs[649].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type integer
		common.NewPropDef("index", typeDefs[45], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Coding
		common.NewPropDef("profile", typeDefs[9], "", false, nil),
	})
	// properties of TestScript_RequestHeader
	typeDefs[650].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("field", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("value", typeDefs[46], "", false, nil),
	})
	// properties of TestScript_Setup
	typeDefs[651].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type TestScript_Action
		common.NewPropDef("action", typeDefs[639], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of TestScript_Teardown
	typeDefs[652].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type TestScript_Action2
		common.NewPropDef("action", typeDefs[641], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of TestScript_Test
	typeDefs[653].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type TestScript_Action1
		common.NewPropDef("action", typeDefs[640], "", true, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
	})
	// properties of TestScript_Variable
	typeDefs[654].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type string
		common.NewPropDef("defaultValue", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("description", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("expression", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("headerField", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("hint", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("path", typeDefs[46], "", false, nil),
		// with type id
		common.NewPropDef("sourceId", typeDefs[675], "", false, nil),
	})
	// properties of Timing
	typeDefs[655].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("code", typeDefs[8], "", false, nil),
		// with type dateTime
		common.NewPropDef("event", typeDefs[42], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Timing_Repeat
		common.NewPropDef("repeat", typeDefs[656], "", false, nil),
	})
	// properties of Timing_Repeat
	typeDefs[656].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Duration
		common.NewPropDef("boundsDuration", typeDefs[236], "", false, nil),
		// with type Period
		common.NewPropDef("boundsPeriod", typeDefs[25], "", false, nil),
		// with type Range
		common.NewPropDef("boundsRange", typeDefs[30], "", false, nil),
		// with type positiveInt
		common.NewPropDef("count", typeDefs[678], "", false, nil),
		// with type positiveInt
		common.NewPropDef("countMax", typeDefs[678], "", false, nil),
		// with type code
		common.NewPropDef("dayOfWeek", typeDefs[674], "", true, nil),
		// with type decimal
		common.NewPropDef("duration", typeDefs[43], "", false, nil),
		// with type decimal
		common.NewPropDef("durationMax", typeDefs[43], "", false, nil),
		// with type string
		common.NewPropDef("durationUnit", typeDefs[46], "", false, []string{
			"s",
			"min",
			"h",
			"d",
			"wk",
			"mo",
			"a",
		}),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type positiveInt
		common.NewPropDef("frequency", typeDefs[678], "", false, nil),
		// with type positiveInt
		common.NewPropDef("frequencyMax", typeDefs[678], "", false, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type unsignedInt
		common.NewPropDef("offset", typeDefs[679], "", false, nil),
		// with type decimal
		common.NewPropDef("period", typeDefs[43], "", false, nil),
		// with type decimal
		common.NewPropDef("periodMax", typeDefs[43], "", false, nil),
		// with type string
		common.NewPropDef("periodUnit", typeDefs[46], "", false, []string{
			"s",
			"min",
			"h",
			"d",
			"wk",
			"mo",
			"a",
		}),
		// with type time
		common.NewPropDef("timeOfDay", typeDefs[47], "", true, nil),
		// with type string
		common.NewPropDef("when", typeDefs[46], "", true, []string{
			"MORN",
			"MORN.early",
			"MORN.late",
			"NOON",
			"AFT",
			"AFT.early",
			"AFT.late",
			"EVE",
			"EVE.early",
			"EVE.late",
			"NIGHT",
			"PHS",
			"HS",
			"WAKE",
			"C",
			"CM",
			"CD",
			"CV",
			"AC",
			"ACM",
			"ACD",
			"ACV",
			"PC",
			"PCM",
			"PCD",
			"PCV",
		}),
	})
	// properties of TriggerDefinition
	typeDefs[37].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Expression
		common.NewPropDef("condition", typeDefs[15], "", false, nil),
		// with type DataRequirement
		common.NewPropDef("data", typeDefs[13], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type date
		common.NewPropDef("timingDate", typeDefs[41], "", false, nil),
		// with type dateTime
		common.NewPropDef("timingDateTime", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("timingReference", typeDefs[32], "", false, nil),
		// with type Timing
		common.NewPropDef("timingTiming", typeDefs[655], "", false, nil),
		// with type string
		common.NewPropDef("type", typeDefs[46], "", false, []string{
			"named-event",
			"periodic",
			"data-changed",
			"data-added",
			"data-modified",
			"data-removed",
			"data-accessed",
			"data-access-ended",
		}),
	})
	// properties of UsageContext
	typeDefs[38].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Coding
		common.NewPropDef("code", typeDefs[9], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("valueCodeableConcept", typeDefs[8], "", false, nil),
		// with type Quantity
		common.NewPropDef("valueQuantity", typeDefs[29], "", false, nil),
		// with type Range
		common.NewPropDef("valueRange", typeDefs[30], "", false, nil),
		// with type Reference
		common.NewPropDef("valueReference", typeDefs[32], "", false, nil),
	})
	// properties of ValueSet
	typeDefs[657].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ValueSet_Compose
		common.NewPropDef("compose", typeDefs[658], "", false, nil),
		// with type ContactDetail
		common.NewPropDef("contact", typeDefs[10], "", true, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type markdown
		common.NewPropDef("copyright", typeDefs[676], "", false, nil),
		// with type dateTime
		common.NewPropDef("date", typeDefs[42], "", false, nil),
		// with type markdown
		common.NewPropDef("description", typeDefs[676], "", false, nil),
		// with type ValueSet_Expansion
		common.NewPropDef("expansion", typeDefs[662], "", false, nil),
		// with type boolean
		common.NewPropDef("experimental", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type boolean
		common.NewPropDef("immutable", typeDefs[40], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("jurisdiction", typeDefs[8], "", true, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("publisher", typeDefs[46], "", false, nil),
		// with type markdown
		common.NewPropDef("purpose", typeDefs[676], "", false, nil),
		// with type string
		common.NewPropDef("status", typeDefs[46], "", false, []string{
			"draft",
			"active",
			"retired",
			"unknown",
		}),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type string
		common.NewPropDef("title", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("url", typeDefs[48], "", false, nil),
		// with type UsageContext
		common.NewPropDef("useContext", typeDefs[38], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of ValueSet_Compose
	typeDefs[658].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ValueSet_Include
		common.NewPropDef("exclude", typeDefs[664], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("inactive", typeDefs[40], "", false, nil),
		// with type ValueSet_Include
		common.NewPropDef("include", typeDefs[664], "", true, nil),
		// with type date
		common.NewPropDef("lockedDate", typeDefs[41], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of ValueSet_Concept
	typeDefs[659].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type code
		common.NewPropDef("code", typeDefs[674], "", false, nil),
		// with type ValueSet_Designation
		common.NewPropDef("designation", typeDefs[661], "", true, nil),
		// with type string
		common.NewPropDef("display", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})
	// properties of ValueSet_Contains
	typeDefs[660].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type boolean
		common.NewPropDef("abstract", typeDefs[40], "", false, nil),
		// with type code
		common.NewPropDef("code", typeDefs[674], "", false, nil),
		// with type ValueSet_Contains
		common.NewPropDef("contains", typeDefs[660], "", true, nil),
		// with type ValueSet_Designation
		common.NewPropDef("designation", typeDefs[661], "", true, nil),
		// with type string
		common.NewPropDef("display", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("inactive", typeDefs[40], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type uri
		common.NewPropDef("system", typeDefs[48], "", false, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of ValueSet_Designation
	typeDefs[661].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Coding
		common.NewPropDef("use", typeDefs[9], "", false, nil),
		// with type string
		common.NewPropDef("value", typeDefs[46], "", false, nil),
	})
	// properties of ValueSet_Expansion
	typeDefs[662].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ValueSet_Contains
		common.NewPropDef("contains", typeDefs[660], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type uri
		common.NewPropDef("identifier", typeDefs[48], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type integer
		common.NewPropDef("offset", typeDefs[45], "", false, nil),
		// with type ValueSet_Parameter
		common.NewPropDef("parameter", typeDefs[665], "", true, nil),
		// with type dateTime
		common.NewPropDef("timestamp", typeDefs[42], "", false, nil),
		// with type integer
		common.NewPropDef("total", typeDefs[45], "", false, nil),
	})
	// properties of ValueSet_Filter
	typeDefs[663].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("op", typeDefs[46], "", false, []string{
			"=",
			"is-a",
			"descendent-of",
			"is-not-a",
			"regex",
			"in",
			"not-in",
			"generalizes",
			"exists",
		}),
		// with type code
		common.NewPropDef("property", typeDefs[674], "", false, nil),
		// with type string
		common.NewPropDef("value", typeDefs[46], "", false, nil),
	})
	// properties of ValueSet_Include
	typeDefs[664].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type ValueSet_Concept
		common.NewPropDef("concept", typeDefs[659], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type ValueSet_Filter
		common.NewPropDef("filter", typeDefs[663], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type uri
		common.NewPropDef("system", typeDefs[48], "", false, nil),
		// with type canonical
		common.NewPropDef("valueSet", typeDefs[673], "", true, nil),
		// with type string
		common.NewPropDef("version", typeDefs[46], "", false, nil),
	})
	// properties of ValueSet_Parameter
	typeDefs[665].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("name", typeDefs[46], "", false, nil),
		// with type boolean
		common.NewPropDef("valueBoolean", typeDefs[40], "value", false, nil),
		// with type code
		common.NewPropDef("valueCode", typeDefs[674], "value", false, nil),
		// with type dateTime
		common.NewPropDef("valueDateTime", typeDefs[42], "value", false, nil),
		// with type decimal
		common.NewPropDef("valueDecimal", typeDefs[43], "value", false, nil),
		// with type integer
		common.NewPropDef("valueInteger", typeDefs[45], "value", false, nil),
		// with type string
		common.NewPropDef("valueString", typeDefs[46], "value", false, nil),
		// with type uri
		common.NewPropDef("valueUri", typeDefs[48], "value", false, nil),
	})
	// properties of VerificationResult
	typeDefs[666].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type VerificationResult_Attestation
		common.NewPropDef("attestation", typeDefs[667], "", false, nil),
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("failureAction", typeDefs[8], "", false, nil),
		// with type Timing
		common.NewPropDef("frequency", typeDefs[655], "", false, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type dateTime
		common.NewPropDef("lastPerformed", typeDefs[42], "", false, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("need", typeDefs[8], "", false, nil),
		// with type date
		common.NewPropDef("nextScheduled", typeDefs[41], "", false, nil),
		// with type VerificationResult_PrimarySource
		common.NewPropDef("primarySource", typeDefs[668], "", true, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type dateTime
		common.NewPropDef("statusDate", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("target", typeDefs[32], "", true, nil),
		// with type string
		common.NewPropDef("targetLocation", typeDefs[46], "", true, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("validationProcess", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("validationType", typeDefs[8], "", false, nil),
		// with type VerificationResult_Validator
		common.NewPropDef("validator", typeDefs[669], "", true, nil),
	})
	// properties of VerificationResult_Attestation
	typeDefs[667].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("communicationMethod", typeDefs[8], "", false, nil),
		// with type date
		common.NewPropDef("date", typeDefs[41], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("onBehalfOf", typeDefs[32], "", false, nil),
		// with type string
		common.NewPropDef("proxyIdentityCertificate", typeDefs[46], "", false, nil),
		// with type Signature
		common.NewPropDef("proxySignature", typeDefs[35], "", false, nil),
		// with type string
		common.NewPropDef("sourceIdentityCertificate", typeDefs[46], "", false, nil),
		// with type Signature
		common.NewPropDef("sourceSignature", typeDefs[35], "", false, nil),
		// with type Reference
		common.NewPropDef("who", typeDefs[32], "", false, nil),
	})
	// properties of VerificationResult_PrimarySource
	typeDefs[668].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type CodeableConcept
		common.NewPropDef("canPushUpdates", typeDefs[8], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("communicationMethod", typeDefs[8], "", true, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("pushTypeAvailable", typeDefs[8], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("type", typeDefs[8], "", true, nil),
		// with type dateTime
		common.NewPropDef("validationDate", typeDefs[42], "", false, nil),
		// with type CodeableConcept
		common.NewPropDef("validationStatus", typeDefs[8], "", false, nil),
		// with type Reference
		common.NewPropDef("who", typeDefs[32], "", false, nil),
	})
	// properties of VerificationResult_Validator
	typeDefs[669].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Signature
		common.NewPropDef("attestationSignature", typeDefs[35], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("identityCertificate", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("organization", typeDefs[32], "", false, nil),
	})
	// properties of VisionPrescription
	typeDefs[670].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type Resource
		common.NewPropDef("contained", typeDefs[1], "", true, nil),
		// with type dateTime
		common.NewPropDef("created", typeDefs[42], "", false, nil),
		// with type dateTime
		common.NewPropDef("dateWritten", typeDefs[42], "", false, nil),
		// with type Reference
		common.NewPropDef("encounter", typeDefs[32], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type id
		common.NewPropDef("id", typeDefs[675], "", false, nil),
		// with type Identifier
		common.NewPropDef("identifier", typeDefs[18], "", true, nil),
		// with type uri
		common.NewPropDef("implicitRules", typeDefs[48], "", false, nil),
		// with type code
		common.NewPropDef("language", typeDefs[674], "", false, nil),
		// with type VisionPrescription_LensSpecification
		common.NewPropDef("lensSpecification", typeDefs[671], "", true, nil),
		// with type Meta
		common.NewPropDef("meta", typeDefs[20], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Reference
		common.NewPropDef("patient", typeDefs[32], "", false, nil),
		// with type Reference
		common.NewPropDef("prescriber", typeDefs[32], "", false, nil),
		// with type code
		common.NewPropDef("status", typeDefs[674], "", false, nil),
		// with type Narrative
		common.NewPropDef("text", typeDefs[22], "", false, nil),
	})
	// properties of VisionPrescription_LensSpecification
	typeDefs[671].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type decimal
		common.NewPropDef("add", typeDefs[43], "", false, nil),
		// with type integer
		common.NewPropDef("axis", typeDefs[45], "", false, nil),
		// with type decimal
		common.NewPropDef("backCurve", typeDefs[43], "", false, nil),
		// with type string
		common.NewPropDef("brand", typeDefs[46], "", false, nil),
		// with type string
		common.NewPropDef("color", typeDefs[46], "", false, nil),
		// with type decimal
		common.NewPropDef("cylinder", typeDefs[43], "", false, nil),
		// with type decimal
		common.NewPropDef("diameter", typeDefs[43], "", false, nil),
		// with type Quantity
		common.NewPropDef("duration", typeDefs[29], "", false, nil),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("eye", typeDefs[46], "", false, []string{
			"right",
			"left",
		}),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
		// with type Annotation
		common.NewPropDef("note", typeDefs[3], "", true, nil),
		// with type decimal
		common.NewPropDef("power", typeDefs[43], "", false, nil),
		// with type VisionPrescription_Prism
		common.NewPropDef("prism", typeDefs[672], "", true, nil),
		// with type CodeableConcept
		common.NewPropDef("product", typeDefs[8], "", false, nil),
		// with type decimal
		common.NewPropDef("sphere", typeDefs[43], "", false, nil),
	})
	// properties of VisionPrescription_Prism
	typeDefs[672].(common.InitializableStructTypeDefAccessor).InitProps([]*common.PropDef{
		// with type decimal
		common.NewPropDef("amount", typeDefs[43], "", false, nil),
		// with type string
		common.NewPropDef("base", typeDefs[46], "", false, []string{
			"up",
			"down",
			"in",
			"out",
		}),
		// with type Extension
		common.NewPropDef("extension", typeDefs[16], "", true, nil),
		// with type string
		common.NewPropDef("id", typeDefs[46], "", false, nil),
		// with type Extension
		common.NewPropDef("modifierExtension", typeDefs[16], "", true, nil),
	})

	typeDefsByName := make(map[string]common.TypeDefAccessor)
	for _, ts := range typeDefs {
		typeDefsByName[ts.InternalName()] = ts
	}
	return common.NewTypeDefContainer("R4", "4.0.1", typeDefsByName)
}
