package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ActiveContextTimeToLive represents the ActiveContextTimeToLive schema from the OpenAPI specification
type ActiveContextTimeToLive struct {
	Timetoliveinseconds interface{} `json:"timeToLiveInSeconds"`
	Turnstolive interface{} `json:"turnsToLive"`
}

// PutSessionRequest represents the PutSessionRequest schema from the OpenAPI specification
type PutSessionRequest struct {
	Messages interface{} `json:"messages,omitempty"`
	Requestattributes interface{} `json:"requestAttributes,omitempty"`
	Sessionstate interface{} `json:"sessionState"`
}

// Slot represents the Slot schema from the OpenAPI specification
type Slot struct {
	Values interface{} `json:"values,omitempty"`
	Subslots interface{} `json:"subSlots,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// DeleteSessionRequest represents the DeleteSessionRequest schema from the OpenAPI specification
type DeleteSessionRequest struct {
}

// PutSessionResponse represents the PutSessionResponse schema from the OpenAPI specification
type PutSessionResponse struct {
	Audiostream interface{} `json:"audioStream,omitempty"`
}

// SentimentScore represents the SentimentScore schema from the OpenAPI specification
type SentimentScore struct {
	Mixed interface{} `json:"mixed,omitempty"`
	Negative interface{} `json:"negative,omitempty"`
	Neutral interface{} `json:"neutral,omitempty"`
	Positive interface{} `json:"positive,omitempty"`
}

// SlotHintsIntentMap represents the SlotHintsIntentMap schema from the OpenAPI specification
type SlotHintsIntentMap struct {
}

// Value represents the Value schema from the OpenAPI specification
type Value struct {
	Interpretedvalue interface{} `json:"interpretedValue"`
	Originalvalue interface{} `json:"originalValue,omitempty"`
	Resolvedvalues interface{} `json:"resolvedValues,omitempty"`
}

// DeleteSessionResponse represents the DeleteSessionResponse schema from the OpenAPI specification
type DeleteSessionResponse struct {
	Localeid interface{} `json:"localeId,omitempty"`
	Sessionid interface{} `json:"sessionId,omitempty"`
	Botaliasid interface{} `json:"botAliasId,omitempty"`
	Botid interface{} `json:"botId,omitempty"`
}

// ElicitSubSlot represents the ElicitSubSlot schema from the OpenAPI specification
type ElicitSubSlot struct {
	Subslottoelicit interface{} `json:"subSlotToElicit,omitempty"`
	Name interface{} `json:"name"`
}

// RuntimeHints represents the RuntimeHints schema from the OpenAPI specification
type RuntimeHints struct {
	Slothints interface{} `json:"slotHints,omitempty"`
}

// DialogAction represents the DialogAction schema from the OpenAPI specification
type DialogAction struct {
	Slotelicitationstyle interface{} `json:"slotElicitationStyle,omitempty"`
	Slottoelicit interface{} `json:"slotToElicit,omitempty"`
	Subslottoelicit interface{} `json:"subSlotToElicit,omitempty"`
	TypeField interface{} `json:"type"`
}

// RuntimeHintValue represents the RuntimeHintValue schema from the OpenAPI specification
type RuntimeHintValue struct {
	Phrase interface{} `json:"phrase"`
}

// Interpretation represents the Interpretation schema from the OpenAPI specification
type Interpretation struct {
	Nluconfidence interface{} `json:"nluConfidence,omitempty"`
	Sentimentresponse interface{} `json:"sentimentResponse,omitempty"`
	Intent interface{} `json:"intent,omitempty"`
}

// SessionState represents the SessionState schema from the OpenAPI specification
type SessionState struct {
	Activecontexts interface{} `json:"activeContexts,omitempty"`
	Dialogaction interface{} `json:"dialogAction,omitempty"`
	Intent interface{} `json:"intent,omitempty"`
	Originatingrequestid interface{} `json:"originatingRequestId,omitempty"`
	Runtimehints interface{} `json:"runtimeHints,omitempty"`
	Sessionattributes interface{} `json:"sessionAttributes,omitempty"`
}

// ActiveContext represents the ActiveContext schema from the OpenAPI specification
type ActiveContext struct {
	Contextattributes interface{} `json:"contextAttributes"`
	Name interface{} `json:"name"`
	Timetolive interface{} `json:"timeToLive"`
}

// Slots represents the Slots schema from the OpenAPI specification
type Slots struct {
}

// ImageResponseCard represents the ImageResponseCard schema from the OpenAPI specification
type ImageResponseCard struct {
	Buttons interface{} `json:"buttons,omitempty"`
	Imageurl interface{} `json:"imageUrl,omitempty"`
	Subtitle interface{} `json:"subtitle,omitempty"`
	Title interface{} `json:"title"`
}

// RecognizeTextRequest represents the RecognizeTextRequest schema from the OpenAPI specification
type RecognizeTextRequest struct {
	Text interface{} `json:"text"`
	Requestattributes interface{} `json:"requestAttributes,omitempty"`
	Sessionstate interface{} `json:"sessionState,omitempty"`
}

// RecognizeUtteranceRequest represents the RecognizeUtteranceRequest schema from the OpenAPI specification
type RecognizeUtteranceRequest struct {
	Inputstream interface{} `json:"inputStream,omitempty"`
}

// SentimentResponse represents the SentimentResponse schema from the OpenAPI specification
type SentimentResponse struct {
	Sentiment interface{} `json:"sentiment,omitempty"`
	Sentimentscore SentimentScore `json:"sentimentScore,omitempty"` // The individual sentiment responses for the utterance.
}

// GetSessionRequest represents the GetSessionRequest schema from the OpenAPI specification
type GetSessionRequest struct {
}

// StringMap represents the StringMap schema from the OpenAPI specification
type StringMap struct {
}

// Button represents the Button schema from the OpenAPI specification
type Button struct {
	Text interface{} `json:"text"`
	Value interface{} `json:"value"`
}

// SlotHintsSlotMap represents the SlotHintsSlotMap schema from the OpenAPI specification
type SlotHintsSlotMap struct {
}

// ConfidenceScore represents the ConfidenceScore schema from the OpenAPI specification
type ConfidenceScore struct {
	Score interface{} `json:"score,omitempty"`
}

// RuntimeHintDetails represents the RuntimeHintDetails schema from the OpenAPI specification
type RuntimeHintDetails struct {
	Runtimehintvalues interface{} `json:"runtimeHintValues,omitempty"`
	Subslothints interface{} `json:"subSlotHints,omitempty"`
}

// RecognizeUtteranceResponse represents the RecognizeUtteranceResponse schema from the OpenAPI specification
type RecognizeUtteranceResponse struct {
	Audiostream interface{} `json:"audioStream,omitempty"`
}

// ActiveContextParametersMap represents the ActiveContextParametersMap schema from the OpenAPI specification
type ActiveContextParametersMap struct {
}

// Intent represents the Intent schema from the OpenAPI specification
type Intent struct {
	State interface{} `json:"state,omitempty"`
	Confirmationstate interface{} `json:"confirmationState,omitempty"`
	Name interface{} `json:"name"`
	Slots interface{} `json:"slots,omitempty"`
}

// GetSessionResponse represents the GetSessionResponse schema from the OpenAPI specification
type GetSessionResponse struct {
	Interpretations interface{} `json:"interpretations,omitempty"`
	Messages interface{} `json:"messages,omitempty"`
	Sessionid interface{} `json:"sessionId,omitempty"`
	Sessionstate interface{} `json:"sessionState,omitempty"`
}

// RecognizedBotMember represents the RecognizedBotMember schema from the OpenAPI specification
type RecognizedBotMember struct {
	Botid interface{} `json:"botId"`
	Botname interface{} `json:"botName,omitempty"`
}

// RecognizeTextResponse represents the RecognizeTextResponse schema from the OpenAPI specification
type RecognizeTextResponse struct {
	Messages interface{} `json:"messages,omitempty"`
	Recognizedbotmember interface{} `json:"recognizedBotMember,omitempty"`
	Requestattributes interface{} `json:"requestAttributes,omitempty"`
	Sessionid interface{} `json:"sessionId,omitempty"`
	Sessionstate interface{} `json:"sessionState,omitempty"`
	Interpretations interface{} `json:"interpretations,omitempty"`
}

// Message represents the Message schema from the OpenAPI specification
type Message struct {
	Content interface{} `json:"content,omitempty"`
	Contenttype interface{} `json:"contentType"`
	Imageresponsecard ImageResponseCard `json:"imageResponseCard,omitempty"` // <p>A card that is shown to the user by a messaging platform. You define the contents of the card, the card is displayed by the platform. </p> <p>When you use a response card, the response from the user is constrained to the text associated with a button on the card.</p>
}
