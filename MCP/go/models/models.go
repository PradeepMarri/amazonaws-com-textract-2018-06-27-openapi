package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ExpenseDetection represents the ExpenseDetection schema from the OpenAPI specification
type ExpenseDetection struct {
	Text interface{} `json:"Text,omitempty"`
	Confidence interface{} `json:"Confidence,omitempty"`
	Geometry Geometry `json:"Geometry,omitempty"` // Information about where the following items are located on a document page: detected page, text, key-value pairs, tables, table cells, and selection elements.
}

// GetLendingAnalysisSummaryRequest represents the GetLendingAnalysisSummaryRequest schema from the OpenAPI specification
type GetLendingAnalysisSummaryRequest struct {
	Jobid interface{} `json:"JobId"`
}

// Relationship represents the Relationship schema from the OpenAPI specification
type Relationship struct {
	Ids interface{} `json:"Ids,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// HumanLoopDataAttributes represents the HumanLoopDataAttributes schema from the OpenAPI specification
type HumanLoopDataAttributes struct {
	Contentclassifiers interface{} `json:"ContentClassifiers,omitempty"`
}

// StartExpenseAnalysisRequest represents the StartExpenseAnalysisRequest schema from the OpenAPI specification
type StartExpenseAnalysisRequest struct {
	Outputconfig interface{} `json:"OutputConfig,omitempty"`
	Clientrequesttoken interface{} `json:"ClientRequestToken,omitempty"`
	Documentlocation interface{} `json:"DocumentLocation"`
	Jobtag interface{} `json:"JobTag,omitempty"`
	Kmskeyid interface{} `json:"KMSKeyId,omitempty"`
	Notificationchannel interface{} `json:"NotificationChannel,omitempty"`
}

// StartDocumentAnalysisRequest represents the StartDocumentAnalysisRequest schema from the OpenAPI specification
type StartDocumentAnalysisRequest struct {
	Queriesconfig QueriesConfig `json:"QueriesConfig,omitempty"` // <p/>
	Clientrequesttoken interface{} `json:"ClientRequestToken,omitempty"`
	Documentlocation interface{} `json:"DocumentLocation"`
	Featuretypes interface{} `json:"FeatureTypes"`
	Jobtag interface{} `json:"JobTag,omitempty"`
	Kmskeyid interface{} `json:"KMSKeyId,omitempty"`
	Notificationchannel interface{} `json:"NotificationChannel,omitempty"`
	Outputconfig interface{} `json:"OutputConfig,omitempty"`
}

// GetDocumentTextDetectionResponse represents the GetDocumentTextDetectionResponse schema from the OpenAPI specification
type GetDocumentTextDetectionResponse struct {
	Documentmetadata interface{} `json:"DocumentMetadata,omitempty"`
	Jobstatus interface{} `json:"JobStatus,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Statusmessage interface{} `json:"StatusMessage,omitempty"`
	Warnings interface{} `json:"Warnings,omitempty"`
	Blocks interface{} `json:"Blocks,omitempty"`
	Detectdocumenttextmodelversion interface{} `json:"DetectDocumentTextModelVersion,omitempty"`
}

// SplitDocument represents the SplitDocument schema from the OpenAPI specification
type SplitDocument struct {
	Pages interface{} `json:"Pages,omitempty"`
	Index interface{} `json:"Index,omitempty"`
}

// ExpenseCurrency represents the ExpenseCurrency schema from the OpenAPI specification
type ExpenseCurrency struct {
	Confidence interface{} `json:"Confidence,omitempty"`
	Code interface{} `json:"Code,omitempty"`
}

// StartLendingAnalysisResponse represents the StartLendingAnalysisResponse schema from the OpenAPI specification
type StartLendingAnalysisResponse struct {
	Jobid interface{} `json:"JobId,omitempty"`
}

// Query represents the Query schema from the OpenAPI specification
type Query struct {
	Text interface{} `json:"Text"`
	Alias interface{} `json:"Alias,omitempty"`
	Pages interface{} `json:"Pages,omitempty"`
}

// ExpenseType represents the ExpenseType schema from the OpenAPI specification
type ExpenseType struct {
	Confidence interface{} `json:"Confidence,omitempty"`
	Text interface{} `json:"Text,omitempty"`
}

// DocumentLocation represents the DocumentLocation schema from the OpenAPI specification
type DocumentLocation struct {
	S3object interface{} `json:"S3Object,omitempty"`
}

// ExpenseDocument represents the ExpenseDocument schema from the OpenAPI specification
type ExpenseDocument struct {
	Lineitemgroups interface{} `json:"LineItemGroups,omitempty"`
	Summaryfields interface{} `json:"SummaryFields,omitempty"`
	Blocks interface{} `json:"Blocks,omitempty"`
	Expenseindex interface{} `json:"ExpenseIndex,omitempty"`
}

// DetectDocumentTextResponse represents the DetectDocumentTextResponse schema from the OpenAPI specification
type DetectDocumentTextResponse struct {
	Documentmetadata interface{} `json:"DocumentMetadata,omitempty"`
	Blocks interface{} `json:"Blocks,omitempty"`
	Detectdocumenttextmodelversion interface{} `json:"DetectDocumentTextModelVersion,omitempty"`
}

// ExpenseField represents the ExpenseField schema from the OpenAPI specification
type ExpenseField struct {
	Valuedetection interface{} `json:"ValueDetection,omitempty"`
	Currency interface{} `json:"Currency,omitempty"`
	Groupproperties interface{} `json:"GroupProperties,omitempty"`
	Labeldetection interface{} `json:"LabelDetection,omitempty"`
	Pagenumber interface{} `json:"PageNumber,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// StartDocumentTextDetectionRequest represents the StartDocumentTextDetectionRequest schema from the OpenAPI specification
type StartDocumentTextDetectionRequest struct {
	Clientrequesttoken interface{} `json:"ClientRequestToken,omitempty"`
	Documentlocation interface{} `json:"DocumentLocation"`
	Jobtag interface{} `json:"JobTag,omitempty"`
	Kmskeyid interface{} `json:"KMSKeyId,omitempty"`
	Notificationchannel interface{} `json:"NotificationChannel,omitempty"`
	Outputconfig interface{} `json:"OutputConfig,omitempty"`
}

// Warning represents the Warning schema from the OpenAPI specification
type Warning struct {
	Pages interface{} `json:"Pages,omitempty"`
	Errorcode interface{} `json:"ErrorCode,omitempty"`
}

// UndetectedSignature represents the UndetectedSignature schema from the OpenAPI specification
type UndetectedSignature struct {
	Page interface{} `json:"Page,omitempty"`
}

// GetDocumentTextDetectionRequest represents the GetDocumentTextDetectionRequest schema from the OpenAPI specification
type GetDocumentTextDetectionRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Jobid interface{} `json:"JobId"`
}

// GetLendingAnalysisResponse represents the GetLendingAnalysisResponse schema from the OpenAPI specification
type GetLendingAnalysisResponse struct {
	Jobstatus interface{} `json:"JobStatus,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Results interface{} `json:"Results,omitempty"`
	Statusmessage interface{} `json:"StatusMessage,omitempty"`
	Warnings interface{} `json:"Warnings,omitempty"`
	Analyzelendingmodelversion interface{} `json:"AnalyzeLendingModelVersion,omitempty"`
	Documentmetadata DocumentMetadata `json:"DocumentMetadata,omitempty"` // Information about the input document.
}

// GetDocumentAnalysisRequest represents the GetDocumentAnalysisRequest schema from the OpenAPI specification
type GetDocumentAnalysisRequest struct {
	Jobid interface{} `json:"JobId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// BoundingBox represents the BoundingBox schema from the OpenAPI specification
type BoundingBox struct {
	Height interface{} `json:"Height,omitempty"`
	Left interface{} `json:"Left,omitempty"`
	Top interface{} `json:"Top,omitempty"`
	Width interface{} `json:"Width,omitempty"`
}

// SignatureDetection represents the SignatureDetection schema from the OpenAPI specification
type SignatureDetection struct {
	Confidence interface{} `json:"Confidence,omitempty"`
	Geometry Geometry `json:"Geometry,omitempty"` // Information about where the following items are located on a document page: detected page, text, key-value pairs, tables, table cells, and selection elements.
}

// StartDocumentAnalysisResponse represents the StartDocumentAnalysisResponse schema from the OpenAPI specification
type StartDocumentAnalysisResponse struct {
	Jobid interface{} `json:"JobId,omitempty"`
}

// StartDocumentTextDetectionResponse represents the StartDocumentTextDetectionResponse schema from the OpenAPI specification
type StartDocumentTextDetectionResponse struct {
	Jobid interface{} `json:"JobId,omitempty"`
}

// Point represents the Point schema from the OpenAPI specification
type Point struct {
	X interface{} `json:"X,omitempty"`
	Y interface{} `json:"Y,omitempty"`
}

// LendingSummary represents the LendingSummary schema from the OpenAPI specification
type LendingSummary struct {
	Documentgroups interface{} `json:"DocumentGroups,omitempty"`
	Undetecteddocumenttypes interface{} `json:"UndetectedDocumentTypes,omitempty"`
}

// NormalizedValue represents the NormalizedValue schema from the OpenAPI specification
type NormalizedValue struct {
	Value interface{} `json:"Value,omitempty"`
	Valuetype interface{} `json:"ValueType,omitempty"`
}

// PageClassification represents the PageClassification schema from the OpenAPI specification
type PageClassification struct {
	Pagenumber interface{} `json:"PageNumber"`
	Pagetype interface{} `json:"PageType"`
}

// GetLendingAnalysisSummaryResponse represents the GetLendingAnalysisSummaryResponse schema from the OpenAPI specification
type GetLendingAnalysisSummaryResponse struct {
	Jobstatus interface{} `json:"JobStatus,omitempty"`
	Statusmessage interface{} `json:"StatusMessage,omitempty"`
	Summary interface{} `json:"Summary,omitempty"`
	Warnings interface{} `json:"Warnings,omitempty"`
	Analyzelendingmodelversion interface{} `json:"AnalyzeLendingModelVersion,omitempty"`
	Documentmetadata DocumentMetadata `json:"DocumentMetadata,omitempty"` // Information about the input document.
}

// GetExpenseAnalysisResponse represents the GetExpenseAnalysisResponse schema from the OpenAPI specification
type GetExpenseAnalysisResponse struct {
	Statusmessage interface{} `json:"StatusMessage,omitempty"`
	Warnings interface{} `json:"Warnings,omitempty"`
	Analyzeexpensemodelversion interface{} `json:"AnalyzeExpenseModelVersion,omitempty"`
	Documentmetadata interface{} `json:"DocumentMetadata,omitempty"`
	Expensedocuments interface{} `json:"ExpenseDocuments,omitempty"`
	Jobstatus interface{} `json:"JobStatus,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// NotificationChannel represents the NotificationChannel schema from the OpenAPI specification
type NotificationChannel struct {
	Rolearn interface{} `json:"RoleArn"`
	Snstopicarn interface{} `json:"SNSTopicArn"`
}

// AnalyzeDocumentRequest represents the AnalyzeDocumentRequest schema from the OpenAPI specification
type AnalyzeDocumentRequest struct {
	Queriesconfig interface{} `json:"QueriesConfig,omitempty"`
	Document interface{} `json:"Document"`
	Featuretypes interface{} `json:"FeatureTypes"`
	Humanloopconfig interface{} `json:"HumanLoopConfig,omitempty"`
}

// IdentityDocumentField represents the IdentityDocumentField schema from the OpenAPI specification
type IdentityDocumentField struct {
	TypeField AnalyzeIDDetections `json:"Type,omitempty"` // Used to contain the information detected by an AnalyzeID operation.
	Valuedetection AnalyzeIDDetections `json:"ValueDetection,omitempty"` // Used to contain the information detected by an AnalyzeID operation.
}

// LendingResult represents the LendingResult schema from the OpenAPI specification
type LendingResult struct {
	Extractions interface{} `json:"Extractions,omitempty"`
	Page interface{} `json:"Page,omitempty"`
	Pageclassification interface{} `json:"PageClassification,omitempty"`
}

// AnalyzeIDResponse represents the AnalyzeIDResponse schema from the OpenAPI specification
type AnalyzeIDResponse struct {
	Analyzeidmodelversion interface{} `json:"AnalyzeIDModelVersion,omitempty"`
	Documentmetadata DocumentMetadata `json:"DocumentMetadata,omitempty"` // Information about the input document.
	Identitydocuments interface{} `json:"IdentityDocuments,omitempty"`
}

// AnalyzeIDRequest represents the AnalyzeIDRequest schema from the OpenAPI specification
type AnalyzeIDRequest struct {
	Documentpages interface{} `json:"DocumentPages"`
}

// OutputConfig represents the OutputConfig schema from the OpenAPI specification
type OutputConfig struct {
	S3bucket interface{} `json:"S3Bucket"`
	S3prefix interface{} `json:"S3Prefix,omitempty"`
}

// AnalyzeDocumentResponse represents the AnalyzeDocumentResponse schema from the OpenAPI specification
type AnalyzeDocumentResponse struct {
	Analyzedocumentmodelversion interface{} `json:"AnalyzeDocumentModelVersion,omitempty"`
	Blocks interface{} `json:"Blocks,omitempty"`
	Documentmetadata interface{} `json:"DocumentMetadata,omitempty"`
	Humanloopactivationoutput interface{} `json:"HumanLoopActivationOutput,omitempty"`
}

// Document represents the Document schema from the OpenAPI specification
type Document struct {
	Bytes interface{} `json:"Bytes,omitempty"`
	S3object interface{} `json:"S3Object,omitempty"`
}

// DocumentGroup represents the DocumentGroup schema from the OpenAPI specification
type DocumentGroup struct {
	Splitdocuments interface{} `json:"SplitDocuments,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Undetectedsignatures interface{} `json:"UndetectedSignatures,omitempty"`
	Detectedsignatures interface{} `json:"DetectedSignatures,omitempty"`
}

// LendingField represents the LendingField schema from the OpenAPI specification
type LendingField struct {
	Keydetection LendingDetection `json:"KeyDetection,omitempty"` // The results extracted for a lending document.
	TypeField interface{} `json:"Type,omitempty"`
	Valuedetections interface{} `json:"ValueDetections,omitempty"`
}

// GetExpenseAnalysisRequest represents the GetExpenseAnalysisRequest schema from the OpenAPI specification
type GetExpenseAnalysisRequest struct {
	Jobid interface{} `json:"JobId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// LendingDetection represents the LendingDetection schema from the OpenAPI specification
type LendingDetection struct {
	Confidence interface{} `json:"Confidence,omitempty"`
	Geometry Geometry `json:"Geometry,omitempty"` // Information about where the following items are located on a document page: detected page, text, key-value pairs, tables, table cells, and selection elements.
	Selectionstatus interface{} `json:"SelectionStatus,omitempty"`
	Text interface{} `json:"Text,omitempty"`
}

// S3Object represents the S3Object schema from the OpenAPI specification
type S3Object struct {
	Bucket interface{} `json:"Bucket,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Version interface{} `json:"Version,omitempty"`
}

// ExpenseGroupProperty represents the ExpenseGroupProperty schema from the OpenAPI specification
type ExpenseGroupProperty struct {
	Id interface{} `json:"Id,omitempty"`
	Types interface{} `json:"Types,omitempty"`
}

// Block represents the Block schema from the OpenAPI specification
type Block struct {
	Entitytypes interface{} `json:"EntityTypes,omitempty"`
	Geometry interface{} `json:"Geometry,omitempty"`
	Confidence interface{} `json:"Confidence,omitempty"`
	Page interface{} `json:"Page,omitempty"`
	Text interface{} `json:"Text,omitempty"`
	Columnindex interface{} `json:"ColumnIndex,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Blocktype interface{} `json:"BlockType,omitempty"`
	Rowspan interface{} `json:"RowSpan,omitempty"`
	Relationships interface{} `json:"Relationships,omitempty"`
	Selectionstatus interface{} `json:"SelectionStatus,omitempty"`
	Columnspan interface{} `json:"ColumnSpan,omitempty"`
	Query interface{} `json:"Query,omitempty"`
	Rowindex interface{} `json:"RowIndex,omitempty"`
	Texttype interface{} `json:"TextType,omitempty"`
}

// QueriesConfig represents the QueriesConfig schema from the OpenAPI specification
type QueriesConfig struct {
	Queries interface{} `json:"Queries"`
}

// LineItemGroup represents the LineItemGroup schema from the OpenAPI specification
type LineItemGroup struct {
	Lineitemgroupindex interface{} `json:"LineItemGroupIndex,omitempty"`
	Lineitems interface{} `json:"LineItems,omitempty"`
}

// AnalyzeIDDetections represents the AnalyzeIDDetections schema from the OpenAPI specification
type AnalyzeIDDetections struct {
	Normalizedvalue interface{} `json:"NormalizedValue,omitempty"`
	Text interface{} `json:"Text"`
	Confidence interface{} `json:"Confidence,omitempty"`
}

// DocumentMetadata represents the DocumentMetadata schema from the OpenAPI specification
type DocumentMetadata struct {
	Pages interface{} `json:"Pages,omitempty"`
}

// Geometry represents the Geometry schema from the OpenAPI specification
type Geometry struct {
	Boundingbox interface{} `json:"BoundingBox,omitempty"`
	Polygon interface{} `json:"Polygon,omitempty"`
}

// DetectDocumentTextRequest represents the DetectDocumentTextRequest schema from the OpenAPI specification
type DetectDocumentTextRequest struct {
	Document interface{} `json:"Document"`
}

// LineItemFields represents the LineItemFields schema from the OpenAPI specification
type LineItemFields struct {
	Lineitemexpensefields interface{} `json:"LineItemExpenseFields,omitempty"`
}

// AnalyzeExpenseRequest represents the AnalyzeExpenseRequest schema from the OpenAPI specification
type AnalyzeExpenseRequest struct {
	Document Document `json:"Document"` // <p>The input document, either as bytes or as an S3 object.</p> <p>You pass image bytes to an Amazon Textract API operation by using the <code>Bytes</code> property. For example, you would use the <code>Bytes</code> property to pass a document loaded from a local file system. Image bytes passed by using the <code>Bytes</code> property must be base64 encoded. Your code might not need to encode document file bytes if you're using an AWS SDK to call Amazon Textract API operations. </p> <p>You pass images stored in an S3 bucket to an Amazon Textract API operation by using the <code>S3Object</code> property. Documents stored in an S3 bucket don't need to be base64 encoded.</p> <p>The AWS Region for the S3 bucket that contains the S3 object must match the AWS Region that you use for Amazon Textract operations.</p> <p>If you use the AWS CLI to call Amazon Textract operations, passing image bytes using the Bytes property isn't supported. You must first upload the document to an Amazon S3 bucket, and then call the operation using the S3Object property.</p> <p>For Amazon Textract to process an S3 object, the user must have permission to access the S3 object. </p>
}

// HumanLoopConfig represents the HumanLoopConfig schema from the OpenAPI specification
type HumanLoopConfig struct {
	Flowdefinitionarn interface{} `json:"FlowDefinitionArn"`
	Humanloopname interface{} `json:"HumanLoopName"`
	Dataattributes interface{} `json:"DataAttributes,omitempty"`
}

// GetLendingAnalysisRequest represents the GetLendingAnalysisRequest schema from the OpenAPI specification
type GetLendingAnalysisRequest struct {
	Jobid interface{} `json:"JobId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// LendingDocument represents the LendingDocument schema from the OpenAPI specification
type LendingDocument struct {
	Lendingfields interface{} `json:"LendingFields,omitempty"`
	Signaturedetections interface{} `json:"SignatureDetections,omitempty"`
}

// IdentityDocument represents the IdentityDocument schema from the OpenAPI specification
type IdentityDocument struct {
	Documentindex interface{} `json:"DocumentIndex,omitempty"`
	Identitydocumentfields interface{} `json:"IdentityDocumentFields,omitempty"`
	Blocks interface{} `json:"Blocks,omitempty"`
}

// Prediction represents the Prediction schema from the OpenAPI specification
type Prediction struct {
	Confidence interface{} `json:"Confidence,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// StartLendingAnalysisRequest represents the StartLendingAnalysisRequest schema from the OpenAPI specification
type StartLendingAnalysisRequest struct {
	Kmskeyid interface{} `json:"KMSKeyId,omitempty"`
	Notificationchannel NotificationChannel `json:"NotificationChannel,omitempty"` // The Amazon Simple Notification Service (Amazon SNS) topic to which Amazon Textract publishes the completion status of an asynchronous document operation.
	Outputconfig OutputConfig `json:"OutputConfig,omitempty"` // <p>Sets whether or not your output will go to a user created bucket. Used to set the name of the bucket, and the prefix on the output file.</p> <p> <code>OutputConfig</code> is an optional parameter which lets you adjust where your output will be placed. By default, Amazon Textract will store the results internally and can only be accessed by the Get API operations. With <code>OutputConfig</code> enabled, you can set the name of the bucket the output will be sent to the file prefix of the results where you can download your results. Additionally, you can set the <code>KMSKeyID</code> parameter to a customer master key (CMK) to encrypt your output. Without this parameter set Amazon Textract will encrypt server-side using the AWS managed CMK for Amazon S3.</p> <p>Decryption of Customer Content is necessary for processing of the documents by Amazon Textract. If your account is opted out under an AI services opt out policy then all unencrypted Customer Content is immediately and permanently deleted after the Customer Content has been processed by the service. No copy of of the output is retained by Amazon Textract. For information about how to opt out, see <a href="https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_ai-opt-out.html"> Managing AI services opt-out policy. </a> </p> <p>For more information on data privacy, see the <a href="https://aws.amazon.com/compliance/data-privacy-faq/">Data Privacy FAQ</a>.</p>
	Clientrequesttoken interface{} `json:"ClientRequestToken,omitempty"`
	Documentlocation DocumentLocation `json:"DocumentLocation"` // <p>The Amazon S3 bucket that contains the document to be processed. It's used by asynchronous operations.</p> <p>The input document can be an image file in JPEG or PNG format. It can also be a file in PDF format.</p>
	Jobtag interface{} `json:"JobTag,omitempty"`
}

// DetectedSignature represents the DetectedSignature schema from the OpenAPI specification
type DetectedSignature struct {
	Page interface{} `json:"Page,omitempty"`
}

// Extraction represents the Extraction schema from the OpenAPI specification
type Extraction struct {
	Identitydocument IdentityDocument `json:"IdentityDocument,omitempty"` // The structure that lists each document processed in an AnalyzeID operation.
	Lendingdocument interface{} `json:"LendingDocument,omitempty"`
	Expensedocument ExpenseDocument `json:"ExpenseDocument,omitempty"` // The structure holding all the information returned by AnalyzeExpense
}

// AnalyzeExpenseResponse represents the AnalyzeExpenseResponse schema from the OpenAPI specification
type AnalyzeExpenseResponse struct {
	Documentmetadata DocumentMetadata `json:"DocumentMetadata,omitempty"` // Information about the input document.
	Expensedocuments interface{} `json:"ExpenseDocuments,omitempty"`
}

// HumanLoopActivationOutput represents the HumanLoopActivationOutput schema from the OpenAPI specification
type HumanLoopActivationOutput struct {
	Humanloopactivationconditionsevaluationresults interface{} `json:"HumanLoopActivationConditionsEvaluationResults,omitempty"`
	Humanloopactivationreasons interface{} `json:"HumanLoopActivationReasons,omitempty"`
	Humanlooparn interface{} `json:"HumanLoopArn,omitempty"`
}

// GetDocumentAnalysisResponse represents the GetDocumentAnalysisResponse schema from the OpenAPI specification
type GetDocumentAnalysisResponse struct {
	Documentmetadata interface{} `json:"DocumentMetadata,omitempty"`
	Jobstatus interface{} `json:"JobStatus,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Statusmessage interface{} `json:"StatusMessage,omitempty"`
	Warnings interface{} `json:"Warnings,omitempty"`
	Analyzedocumentmodelversion interface{} `json:"AnalyzeDocumentModelVersion,omitempty"`
	Blocks interface{} `json:"Blocks,omitempty"`
}

// StartExpenseAnalysisResponse represents the StartExpenseAnalysisResponse schema from the OpenAPI specification
type StartExpenseAnalysisResponse struct {
	Jobid interface{} `json:"JobId,omitempty"`
}
