package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/amazon-textract/mcp-server/config"
	"github.com/amazon-textract/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetdocumenttextdetectionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.GetDocumentTextDetectionRequest
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/#X-Amz-Target=Textract.GetDocumentTextDetection", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["X-Amz-Target"]; ok {
			req.Header.Set("X-Amz-Target", fmt.Sprintf("%v", val))
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.GetDocumentTextDetectionResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGetdocumenttextdetectionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_#X-Amz-Target=Textract_GetDocumentTextDetection",
		mcp.WithDescription("<p>Gets the results for an Amazon Textract asynchronous operation that detects text in a document. Amazon Textract can detect lines of text and the words that make up a line of text.</p> <p>You start asynchronous text detection by calling <a>StartDocumentTextDetection</a>, which returns a job identifier (<code>JobId</code>). When the text detection operation finishes, Amazon Textract publishes a completion status to the Amazon Simple Notification Service (Amazon SNS) topic that's registered in the initial call to <code>StartDocumentTextDetection</code>. To get the results of the text-detection operation, first check that the status value published to the Amazon SNS topic is <code>SUCCEEDED</code>. If so, call <code>GetDocumentTextDetection</code>, and pass the job identifier (<code>JobId</code>) from the initial call to <code>StartDocumentTextDetection</code>.</p> <p> <code>GetDocumentTextDetection</code> returns an array of <a>Block</a> objects. </p> <p>Each document page has as an associated <code>Block</code> of type PAGE. Each PAGE <code>Block</code> object is the parent of LINE <code>Block</code> objects that represent the lines of detected text on a page. A LINE <code>Block</code> object is a parent for each word that makes up the line. Words are represented by <code>Block</code> objects of type WORD.</p> <p>Use the MaxResults parameter to limit the number of blocks that are returned. If there are more results than specified in <code>MaxResults</code>, the value of <code>NextToken</code> in the operation response contains a pagination token for getting the next set of results. To get the next page of results, call <code>GetDocumentTextDetection</code>, and populate the <code>NextToken</code> request parameter with the token value that's returned from the previous call to <code>GetDocumentTextDetection</code>.</p> <p>For more information, see <a href="https://docs.aws.amazon.com/textract/latest/dg/how-it-works-detecting.html">Document Text Detection</a>.</p>"),
		mcp.WithString("X-Amz-Target", mcp.Required(), mcp.Description("")),
		mcp.WithString("JobId", mcp.Required(), mcp.Description("")),
		mcp.WithString("MaxResults", mcp.Description("")),
		mcp.WithString("NextToken", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetdocumenttextdetectionHandler(cfg),
	}
}
