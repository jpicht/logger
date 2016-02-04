package logger

import (
	"log"
	"strings"
)

type SeverityFn func(s string, fallback Severity) Severity

type Hijack struct {
	logger           Logger
	fallbackSeverity Severity
	severityFn       SeverityFn
}

func DefaultSeverityFn(s string, fallback Severity) Severity {
	if strings.Contains(s, "WARN") {
		return WARNING
	}
	if strings.Contains(s, "ERR") {
		return ERROR
	}
	if strings.Contains(s, "NOTICE") {
		return NOTICE
	}
	if strings.Contains(s, "DEBUG") {
		return DEBUG
	}
	return fallback
}

func NewHijack(l Logger, s Severity) *Hijack {
	return &Hijack{
		logger:           l,
		fallbackSeverity: s,
		severityFn:       DefaultSeverityFn,
	}
}

func (h *Hijack) Write(data []byte) (n int, err error) {
	parts := strings.Split(string(data), "\n")
	for _, s := range parts {
		if len(s) <= 20 {
			continue
		}
		s = strings.TrimSpace(s[20:])
		sev := h.severityFn(s, h.fallbackSeverity)
		h.logger.Log(sev, s)
	}
	return len(data), nil
}

func (h *Hijack) Do() {
	log.SetOutput(h)
	log.SetFlags(log.LstdFlags)
}
