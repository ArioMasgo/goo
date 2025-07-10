package models

type GPTRequest struct {
	Model    string        `json:"model"`
	Store    bool          `json:"store"`
	Messages []GPTMessages `json:"messages"`
}

type GPTResponse struct {
	Id                string              `json:"id"`
	Object            string              `json:"object"`
	Created           int                 `json:"created"`
	Model             string              `json:"model"`
	Choices           []GPTResponseChoice `json:"choices"`
	Usage             GPTResponseUsage    `json:"usage"`
	ServiceTier       string              `json:"service_tier"`
	SystemFingerprint string              `json:"system_fingerprint"`
}

type GPTResponseChoice struct {
	Index        int         `json:"index"`
	Message      GPTMessages `json:"message"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
}

type GPTMessages struct {
	Role    string      `json:"role"`
	Content string      `json:"content"`
	Refusal interface{} `json:"refusal,omitempty"`
}

type GPTResponseUsage struct {
	PromptTokens            int                               `json:"prompt_tokens"`
	CompletionTokens        int                               `json:"completion_tokens"`
	TotalTokens             int                               `json:"total_tokens"`
	PromptTokensDetails     GPTResponsePromptTokensDetail     `json:"prompt_tokens_details"`
	CompletionTokensDetails GPTResponseCompletionTokensDetail `json:"completion_tokens_details"`
}

type GPTResponsePromptTokensDetail struct {
	CachedTokens int `json:"cached_tokens"`
	AudioTokens  int `json:"audio_tokens"`
}

type GPTResponseCompletionTokensDetail struct {
	ReasoningTokens          int `json:"reasoning_tokens"`
	AudioTokens              int `json:"audio_tokens"`
	AcceptedPredictionTokens int `json:"accepted_prediction_tokens"`
	RejectedPredictionTokens int `json:"rejected_prediction_tokens"`
}
