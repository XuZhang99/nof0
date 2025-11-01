package executor

import (
	"fmt"
	"strings"

	"nof0-api/pkg/prompt"
)

// PromptInputs contains dynamic data injected into the executor prompt template.
type PromptInputs struct {
	CurrentTime     string
	RuntimeMinutes  int
	SharpeRatio     float64
	AccountOverview string
	OpenPositions   string
	RiskBudget      string
	PerformanceView string
	CandidateCoins  string
	MarketSnapshots string
}

// PromptRenderer renders the executor system prompt from a template file.
type PromptRenderer struct {
	cfg *Config
	tpl *prompt.Template
}

// NewPromptRenderer constructs a renderer using the supplied template path.
func NewPromptRenderer(cfg *Config, templatePath string) (*PromptRenderer, error) {
	if cfg == nil {
		return nil, fmt.Errorf("executor prompt renderer requires config")
	}
	path := strings.TrimSpace(templatePath)
	if path == "" {
		return nil, fmt.Errorf("executor prompt renderer requires template path")
	}
	tpl, err := prompt.NewTemplate(path, nil)
	if err != nil {
		return nil, err
	}
	return &PromptRenderer{
		cfg: cfg,
		tpl: tpl,
	}, nil
}

// Render generates the final prompt string populated with inputs.
func (r *PromptRenderer) Render(inputs PromptInputs) (string, error) {
	if r == nil || r.tpl == nil {
		return "", fmt.Errorf("executor prompt renderer not initialised")
	}

	payload := struct {
		Config *Config
		PromptInputs
	}{
		Config:       r.cfg,
		PromptInputs: inputs,
	}

	return r.tpl.Render(payload)
}

// Digest returns the underlying template digest for observability.
func (r *PromptRenderer) Digest() string {
	if r == nil || r.tpl == nil {
		return ""
	}
	return r.tpl.Digest()
}
