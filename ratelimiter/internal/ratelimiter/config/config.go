package config

type Rule struct {
	Type      string `json:"type,omitempty"`
	PathMatch string `json:"path_match,omitempty"`
	Limit     int64  `json:"limit,omitempty"`

	FixedWindow   FixedWindowConfig   `json:"fixed_window,omitempty"`
	SlidingWindow SlidingWindowConfig `json:"sliding_window,omitempty"`
}

type FixedWindowConfig struct {
	Window int64 `json:"window_ms,omitempty"`
}

type SlidingWindowConfig struct {
	Window int64 `json:"window_ms,omitempty"`
}
