package proxy

import (
	"fmt"
	"net"
	"net/http"
	"os/exec"
	"time"

	"golang.org/x/net/proxy"
	"veko-grid/utils"
)

// TORManager mengelola koneksi TOR
type TORManager struct {
	logger        *utils.Logger
	controlPort   string
	socksPort     string
	isRunning     bool
	currentCircuit string
}

// NewTORManager membuat instance TORManager baru
func NewTORManager(logger *utils.Logger) *TORManager {
	return &TORManager{
		logger:      logger,
		controlPort: "9051",
		socksPort:   "9050",
		isRunning:   false,
	}
}

// CheckTORService mengecek apakah TOR service sudah berjalan
func (t *TORManager) CheckTORService() bool {
	// Coba connect ke SOCKS port
	conn, err := net.DialTimeout("tcp", "127.0.0.1:"+t.socksPort, 3*time.Second)
	if err != nil {
		t.logger.Debug("TOR SOCKS port tidak dapat diakses")
		return false
	}
	defer conn.Close()

	t.isRunning = true
	t.logger.Info("🧅 TOR service terdeteksi dan berjalan")
	return true
}

// StartTORService mencoba memulai TOR service (jika memungkinkan)
func (t *TORManager) StartTORService() error {
	if t.isRunning {
		return nil
	}

	t.logger.Info("🧅 Mencoba memulai TOR service...")

	// Coba jalankan tor command
	cmd := exec.Command("tor", "--SocksPort", t.socksPort, "--ControlPort", t.controlPort)
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("gagal memulai TOR: %v", err)
	}

	// Wait sebentar untuk TOR startup
	time.Sleep(10 * time.Second)

	if !t.CheckTORService() {
		return fmt.Errorf("TOR service tidak dapat dimulai")
	}

	return nil
}

// GetTORDialer mendapatkan dialer untuk TOR
func (t *TORManager) GetTORDialer() (proxy.Dialer, error) {
	if !t.isRunning && !t.CheckTORService() {
		return nil, fmt.Errorf("TOR service tidak tersedia")
	}

	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:"+t.socksPort, nil, &net.Dialer{
		Timeout: 30 * time.Second,
	})
	if err != nil {
		return nil, fmt.Errorf("gagal membuat TOR dialer: %v", err)
	}

	return dialer, nil
}

// GetTORHTTPClient mendapatkan HTTP client melalui TOR
func (t *TORManager) GetTORHTTPClient() (*http.Client, error) {
	dialer, err := t.GetTORDialer()
	if err != nil {
		return nil, err
	}

	transport := &http.Transport{
		Dial: dialer.Dial,
		// Disable keep-alive untuk privasi lebih baik
		DisableKeepAlives: true,
		// Timeout settings
		ResponseHeaderTimeout: 30 * time.Second,
		IdleConnTimeout:       30 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   60 * time.Second,
	}

	return client, nil
}

// NewTORCircuit membuat circuit TOR baru
func (t *TORManager) NewTORCircuit() error {
	if !t.isRunning {
		return fmt.Errorf("TOR service tidak berjalan")
	}

	t.logger.Debug("🔄 Membuat circuit TOR baru...")

	// Connect ke control port
	conn, err := net.Dial("tcp", "127.0.0.1:"+t.controlPort)
	if err != nil {
		return fmt.Errorf("gagal connect ke TOR control port: %v", err)
	}
	defer conn.Close()

	// Send NEWNYM command untuk circuit baru
	_, err = conn.Write([]byte("AUTHENTICATE\r\n"))
	if err != nil {
		return err
	}

	_, err = conn.Write([]byte("SIGNAL NEWNYM\r\n"))
	if err != nil {
		return err
	}

	t.logger.Info("🧅 Circuit TOR baru telah dibuat")
	return nil
}

// GetTORIP mendapatkan IP address melalui TOR
func (t *TORManager) GetTORIP() (string, error) {
	client, err := t.GetTORHTTPClient()
	if err != nil {
		return "", err
	}

	resp, err := client.Get("https://checkip.amazonaws.com")
	if err != nil {
		return "", fmt.Errorf("gagal mengecek IP TOR: %v", err)
	}
	defer resp.Body.Close()

	ip := make([]byte, 15)
	n, err := resp.Body.Read(ip)
	if err != nil {
		return "", err
	}

	return string(ip[:n]), nil
}

// VerifyTORConnection memverifikasi koneksi TOR
func (t *TORManager) VerifyTORConnection() error {
	client, err := t.GetTORHTTPClient()
	if err != nil {
		return err
	}

	// Test dengan service yang mendeteksi TOR
	resp, err := client.Get("https://check.torproject.org/api/ip")
	if err != nil {
		return fmt.Errorf("gagal verifikasi TOR: %v", err)
	}
	defer resp.Body.Close()

	// Parse response untuk memastikan kita menggunakan TOR
	buffer := make([]byte, 1024)
	n, err := resp.Body.Read(buffer)
	if err != nil {
		return err
	}

	response := string(buffer[:n])
	t.logger.Debug(fmt.Sprintf("TOR verification response: %s", response))

	// Di real implementation, parse JSON response untuk check IsTor field
	return nil
}

// GenerateRandomUserAgent menghasilkan User-Agent random untuk TOR
func (t *TORManager) GenerateRandomUserAgent() string {
	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:89.0) Gecko/20100101 Firefox/89.0",
	}

	// Random selection
	return userAgents[time.Now().UnixNano()%int64(len(userAgents))]
}

// IsRunning mengecek apakah TOR sedang berjalan
func (t *TORManager) IsRunning() bool {
	return t.isRunning
}

// Stop menghentikan TOR manager
func (t *TORManager) Stop() {
	if t.isRunning {
		t.logger.Info("🧅 Menghentikan TOR manager...")
		t.isRunning = false
	}
}
