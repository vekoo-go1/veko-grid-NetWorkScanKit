package utils

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"math/big"
	"net"
	"strings"
	"time"
)

// FingerprintSpoofer mengelola TLS fingerprint spoofing
type FingerprintSpoofer struct {
	logger *Logger
}

// TLSFingerprint menyimpan informasi TLS fingerprint
type TLSFingerprint struct {
	JA3         string            `json:"ja3,omitempty"`
	TLSVersion  string            `json:"tls_version"`
	CipherSuite string            `json:"cipher_suite"`
	ServerName  string            `json:"server_name"`
	Certificate *CertificateInfo  `json:"certificate,omitempty"`
	Extensions  map[string]string `json:"extensions,omitempty"`
}

// CertificateInfo menyimpan informasi sertifikat
type CertificateInfo struct {
	Subject     string    `json:"subject"`
	Issuer      string    `json:"issuer"`
	NotBefore   time.Time `json:"not_before"`
	NotAfter    time.Time `json:"not_after"`
	SerialNumber string   `json:"serial_number"`
	Fingerprint string    `json:"fingerprint"`
}

// NewFingerprintSpoofer membuat instance FingerprintSpoofer baru
func NewFingerprintSpoofer(logger *Logger) *FingerprintSpoofer {
	return &FingerprintSpoofer{
		logger: logger,
	}
}

// AnalyzeTLS menganalisis TLS connection dan fingerprint
func (f *FingerprintSpoofer) AnalyzeTLS(target string) map[string]interface{} {
	result := make(map[string]interface{})

	// Parse target untuk mendapatkan host dan port
	host, port := f.parseTarget(target)
	if port == "" {
		port = "443" // Default HTTPS port
	}

	address := net.JoinHostPort(host, port)

	// Custom TLS config untuk fingerprinting
	tlsConfig := f.createRandomTLSConfig(host)

	// Connect dengan timeout
	dialer := &net.Dialer{
		Timeout: 10 * time.Second,
	}

	conn, err := tls.DialWithDialer(dialer, "tcp", address, tlsConfig)
	if err != nil {
		f.logger.Debug(fmt.Sprintf("TLS connection failed for %s: %v", target, err))
		result["error"] = err.Error()
		return result
	}
	defer conn.Close()

	// Analyze connection state
	state := conn.ConnectionState()
	
	fingerprint := &TLSFingerprint{
		TLSVersion:  f.getTLSVersionString(state.Version),
		CipherSuite: f.getCipherSuiteString(state.CipherSuite),
		ServerName:  state.ServerName,
		Extensions:  make(map[string]string),
	}

	// Analyze certificates
	if len(state.PeerCertificates) > 0 {
		cert := state.PeerCertificates[0]
		fingerprint.Certificate = &CertificateInfo{
			Subject:      cert.Subject.String(),
			Issuer:       cert.Issuer.String(),
			NotBefore:    cert.NotBefore,
			NotAfter:     cert.NotAfter,
			SerialNumber: cert.SerialNumber.String(),
		}
	}

	// Generate JA3 fingerprint (simplified)
	fingerprint.JA3 = f.generateJA3(state)

	result["fingerprint"] = fingerprint
	result["handshake_complete"] = state.HandshakeComplete
	result["peer_certificates_count"] = len(state.PeerCertificates)
	
	return result
}

// parseTarget memparse target menjadi host dan port
func (f *FingerprintSpoofer) parseTarget(target string) (string, string) {
	// Remove protocol if present
	target = strings.TrimPrefix(target, "https://")
	target = strings.TrimPrefix(target, "http://")

	// Split host and port
	host, port, err := net.SplitHostPort(target)
	if err != nil {
		// No port specified
		return target, ""
	}

	return host, port
}

// createRandomTLSConfig membuat konfigurasi TLS dengan fingerprint random
func (f *FingerprintSpoofer) createRandomTLSConfig(serverName string) *tls.Config {
	config := &tls.Config{
		ServerName:         serverName,
		InsecureSkipVerify: false,
		MinVersion:         f.getRandomTLSVersion(),
		MaxVersion:         tls.VersionTLS13,
		CipherSuites:       f.getRandomCipherSuites(),
	}

	return config
}

// getRandomTLSVersion mendapatkan versi TLS random
func (f *FingerprintSpoofer) getRandomTLSVersion() uint16 {
	versions := []uint16{
		tls.VersionTLS12,
		tls.VersionTLS13,
	}

	idx := f.randomInt(len(versions))
	return versions[idx]
}

// getRandomCipherSuites mendapatkan cipher suites random
func (f *FingerprintSpoofer) getRandomCipherSuites() []uint16 {
	allSuites := []uint16{
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
	}

	// Select 3-5 random cipher suites
	count := 3 + f.randomInt(3)
	var selected []uint16

	for i := 0; i < count && i < len(allSuites); i++ {
		idx := f.randomInt(len(allSuites))
		selected = append(selected, allSuites[idx])
	}

	return selected
}

// getTLSVersionString mengkonversi TLS version ke string
func (f *FingerprintSpoofer) getTLSVersionString(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return fmt.Sprintf("Unknown (0x%x)", version)
	}
}

// getCipherSuiteString mengkonversi cipher suite ke string
func (f *FingerprintSpoofer) getCipherSuiteString(suite uint16) string {
	switch suite {
	case tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256:
		return "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256"
	case tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384:
		return "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"
	case tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256:
		return "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"
	case tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384:
		return "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"
	default:
		return fmt.Sprintf("Unknown (0x%x)", suite)
	}
}

// generateJA3 menghasilkan JA3 fingerprint (simplified)
func (f *FingerprintSpoofer) generateJA3(state tls.ConnectionState) string {
	// JA3 format: TLSVersion,Ciphers,Extensions,EllipticCurves,EllipticCurvePointFormats
	// Ini adalah implementasi sederhana
	
	version := fmt.Sprintf("%d", state.Version)
	cipher := fmt.Sprintf("%d", state.CipherSuite)
	
	// Generate pseudo-random extensions untuk demo
	extensions := "23-65281-10-11-35-16"
	curves := "29-23-24"
	pointFormats := "0"
	
	ja3String := fmt.Sprintf("%s,%s,%s,%s,%s", version, cipher, extensions, curves, pointFormats)
	
	return f.md5Hash(ja3String)[:16] // Simplified hash
}

// md5Hash menghasilkan hash MD5 sederhana (untuk demo)
func (f *FingerprintSpoofer) md5Hash(input string) string {
	// Simplified hash implementation
	hash := 0
	for _, char := range input {
		hash = hash*31 + int(char)
	}
	return fmt.Sprintf("%08x", hash)
}

// randomInt menghasilkan integer random
func (f *FingerprintSpoofer) randomInt(max int) int {
	if max <= 0 {
		return 0
	}
	
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0
	}
	
	return int(n.Int64())
}

// SpoofUserAgent menghasilkan User-Agent random
func (f *FingerprintSpoofer) SpoofUserAgent() string {
	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/121.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/121.0",
		"Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/121.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.1 Safari/605.1.15",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.0.0.0",
	}

	idx := f.randomInt(len(userAgents))
	return userAgents[idx]
}

// GenerateRandomHeaders menghasilkan HTTP headers random
func (f *FingerprintSpoofer) GenerateRandomHeaders() map[string]string {
	headers := map[string]string{
		"User-Agent":      f.SpoofUserAgent(),
		"Accept":          f.getRandomAcceptHeader(),
		"Accept-Language": f.getRandomLanguageHeader(),
		"Accept-Encoding": "gzip, deflate, br",
		"DNT":             "1",
		"Connection":      "keep-alive",
		"Upgrade-Insecure-Requests": "1",
	}

	return headers
}

// getRandomAcceptHeader menghasilkan Accept header random
func (f *FingerprintSpoofer) getRandomAcceptHeader() string {
	accepts := []string{
		"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
		"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
	}

	idx := f.randomInt(len(accepts))
	return accepts[idx]
}

// getRandomLanguageHeader menghasilkan Accept-Language header random
func (f *FingerprintSpoofer) getRandomLanguageHeader() string {
	languages := []string{
		"en-US,en;q=0.9",
		"en-GB,en;q=0.9",
		"en-US,en;q=0.8",
		"en,en-US;q=0.9",
		"en-US,en;q=0.5",
	}

	idx := f.randomInt(len(languages))
	return languages[idx]
}
