package config

var (
	OpenRouterAPIKey         = GetRequired("OPENROUTER_API_KEY")
	OpenRouterBaseURL        = Get("OPENROUTER_BASE_URL", "https://openrouter.ai/api/v1")
	OpenRouterAppURL         = Get("OPENROUTER_APP_URL", "")
	OpenRouterAppTitle       = Get("OPENROUTER_APP_TITLE", "go-service")
	OpenRouterEmbeddingModel = Get("OPENROUTER_EMBEDDING_MODEL", "openai/text-embedding-3-small")
	OpenRouterChatModel      = Get("OPENROUTER_CHAT_MODEL", "openai/gpt-4o-mini")
)
