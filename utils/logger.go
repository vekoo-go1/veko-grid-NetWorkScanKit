package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

// LogLevel menentukan level logging
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

// Logger mengelola logging untuk Veko Grid
type Logger struct {
	debugMode bool
	silent    bool
	logger    *log.Logger
}

// NewLogger membuat instance Logger baru
func NewLogger(debugMode, silent bool) *Logger {
	return &Logger{
		debugMode: debugMode,
		silent:    silent,
		logger:    log.New(os.Stdout, "", 0),
	}
}

// Debug mencetak debug message
func (l *Logger) Debug(message string) {
	if l.debugMode && !l.silent {
		timestamp := time.Now().Format("15:04:05")
		l.logger.Printf("🔍 [DEBUG] %s | %s", timestamp, message)
	}
}

// Info mencetak info message
func (l *Logger) Info(message string) {
	if !l.silent {
		timestamp := time.Now().Format("15:04:05")
		l.logger.Printf("ℹ️  [INFO]  %s | %s", timestamp, message)
	}
}

// Warn mencetak warning message
func (l *Logger) Warn(message string) {
	if !l.silent {
		timestamp := time.Now().Format("15:04:05")
		l.logger.Printf("⚠️  [WARN]  %s | %s", timestamp, message)
	}
}

// Error mencetak error message
func (l *Logger) Error(message string) {
	timestamp := time.Now().Format("15:04:05")
	l.logger.Printf("❌ [ERROR] %s | %s", timestamp, message)
}

// Fatal mencetak fatal error dan keluar dari program
func (l *Logger) Fatal(message string) {
	timestamp := time.Now().Format("15:04:05")
	l.logger.Printf("💀 [FATAL] %s | %s", timestamp, message)
	os.Exit(1)
}

// Progress mencetak progress message
func (l *Logger) Progress(current, total int, message string) {
	if !l.silent {
		percentage := float64(current) / float64(total) * 100
		progressBar := l.generateProgressBar(int(percentage))
		fmt.Printf("\r🔄 [%d/%d] %s %s %.1f%%", current, total, progressBar, message, percentage)
		
		if current == total {
			fmt.Println() // New line setelah selesai
		}
	}
}

// generateProgressBar menghasilkan progress bar ASCII
func (l *Logger) generateProgressBar(percentage int) string {
	const width = 20
	filled := percentage * width / 100
	bar := ""

	for i := 0; i < width; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}

	return fmt.Sprintf("[%s]", bar)
}

// LogRequest mencatat request yang dibuat
func (l *Logger) LogRequest(method, url string) {
	if l.debugMode {
		l.Debug(fmt.Sprintf("Request: %s %s", method, url))
	}
}

// LogResponse mencatat response yang diterima
func (l *Logger) LogResponse(statusCode int, size int64, duration time.Duration) {
	if l.debugMode {
		l.Debug(fmt.Sprintf("Response: %d | %d bytes | %v", statusCode, size, duration))
	}
}

// LogScanStart mencatat mulai scanning target
func (l *Logger) LogScanStart(target string, index, total int) {
	if !l.silent {
		timestamp := time.Now().Format("15:04:05")
		fmt.Printf("🎯 [%d/%d] %s | Scanning: %s\n", index, total, timestamp, target)
	}
}

// LogScanComplete mencatat selesai scanning target
func (l *Logger) LogScanComplete(target string, duration time.Duration, success bool) {
	if l.debugMode {
		status := "✅ SUCCESS"
		if !success {
			status = "❌ FAILED"
		}
		l.Debug(fmt.Sprintf("Scan %s: %s in %v", target, status, duration))
	}
}

// LogProxyRotation mencatat rotasi proxy
func (l *Logger) LogProxyRotation(oldProxy, newProxy string) {
	if l.debugMode {
		l.Debug(fmt.Sprintf("Proxy rotated: %s → %s", oldProxy, newProxy))
	}
}

// LogTORCircuit mencatat pembuatan circuit TOR baru
func (l *Logger) LogTORCircuit(circuitID string) {
	l.Info(fmt.Sprintf("🧅 New TOR circuit: %s", circuitID))
}

// LogDNSQuery mencatat DNS query
func (l *Logger) LogDNSQuery(domain, recordType string) {
	if l.debugMode {
		l.Debug(fmt.Sprintf("DNS Query: %s (%s)", domain, recordType))
	}
}

// LogTLSFingerprint mencatat TLS fingerprinting
func (l *Logger) LogTLSFingerprint(target, fingerprint string) {
	if l.debugMode {
		l.Debug(fmt.Sprintf("TLS Fingerprint %s: %s", target, fingerprint))
	}
}

// LogStatistics mencetak statistik scanning
func (l *Logger) LogStatistics(stats map[string]interface{}) {
	if !l.silent {
		l.Info("📊 Scanning Statistics:")
		for key, value := range stats {
			fmt.Printf("   %s: %v\n", key, value)
		}
	}
}

// SetSilent mengatur mode silent
func (l *Logger) SetSilent(silent bool) {
	l.silent = silent
}

// SetDebug mengatur mode debug
func (l *Logger) SetDebug(debug bool) {
	l.debugMode = debug
}

// IsSilent mengecek apakah mode silent aktif
func (l *Logger) IsSilent() bool {
	return l.silent
}

// IsDebug mengecek apakah mode debug aktif
func (l *Logger) IsDebug() bool {
	return l.debugMode
}

// LogWithLevel mencetak log dengan level tertentu
func (l *Logger) LogWithLevel(level LogLevel, message string) {
	switch level {
	case DEBUG:
		l.Debug(message)
	case INFO:
		l.Info(message)
	case WARN:
		l.Warn(message)
	case ERROR:
		l.Error(message)
	}
}

// CreateLogFile membuat file log untuk penyimpanan
func (l *Logger) CreateLogFile(filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to create log file: %v", err)
	}

	// Update logger untuk menulis ke file juga
	l.logger = log.New(file, "", log.LstdFlags)
	l.Info(fmt.Sprintf("Log file created: %s", filename))
	
	return nil
}

// Banner mencetak banner Veko Grid
func (l *Logger) Banner() {
	if !l.silent {
		banner := `
╔═══════════════════════════════════════════════════════════════╗
║  🛰️  VEKO GRID v1.0.0 - Network Exploration & Stealth Tool   ║
║  📡 Anonymous Grid Scanning • TOR/Proxy Support • DNS/DoH    ║
║  🔐 TLS Fingerprint Spoofing • Academic Research Tool        ║
╚═══════════════════════════════════════════════════════════════╝
`
		fmt.Print(banner)
	}
}

// Separator mencetak separator line
func (l *Logger) Separator() {
	if !l.silent {
		fmt.Println(fmt.Sprintf("%s", "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"))
	}
}
