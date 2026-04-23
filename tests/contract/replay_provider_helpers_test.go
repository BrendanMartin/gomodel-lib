//go:build contract

package contract

import (
	"testing"

	"github.com/BrendanMartin/gomodel-lib/core"
	"github.com/BrendanMartin/gomodel-lib/llmclient"
	"github.com/BrendanMartin/gomodel-lib/providers/gemini"
)

func newGeminiReplayProvider(t *testing.T, routes map[string]replayRoute) core.Provider {
	t.Helper()

	client := newReplayHTTPClient(t, routes)
	provider := gemini.NewWithHTTPClient("test-api-key", client, llmclient.Hooks{})
	provider.SetBaseURL("https://replay.local")
	provider.SetModelsURL("https://replay.local")
	return provider
}
