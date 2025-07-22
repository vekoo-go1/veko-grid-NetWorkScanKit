package proxy

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/proxy"
	"veko-grid/utils"
)

// Manager mengelola proxy connections dan rotations
type Manager struct {
	proxies     []ProxyConfig
	currentIdx  int
	logger      *utils.Logger
	useTor      bool
}

// ProxyConfig menyimpan konfigurasi proxy
type ProxyConfig struct {
	Address  string
	Type     string // "socks5", "http", "https"
	Username string
	Password string
	Active   bool
}

// NewManager membuat instance Manager baru
func NewManager(proxyAddr string, useTor bool, logger *utils.Logger) (*Manager, error) {
	manager := &Manager{
		proxies:    make([]ProxyConfig, 0),
		currentIdx: 0,
		logger:     logger,
		useTor:     useTor,
	}

	// Add proxy jika disediakan
	if proxyAddr != "" {
		proxyConfig, err := parseProxyAddress(proxyAddr)
		if err != nil {
			return nil, fmt.Errorf("invalid proxy address: %v", err)
		}
		manager.proxies = append(manager.proxies, *proxyConfig)
	}

	// Setup TOR jika diminta
	if useTor {
		torConfig := &ProxyConfig{
			Address: "127.0.0.1:9050",
			Type:    "socks5",
			Active:  true,
		}
		manager.proxies = append(manager.proxies, *torConfig)
		logger.Info("🧅 TOR proxy configured")
	}

	// Default proxy list untuk rotasi
	manager.addDefaultProxies()

	if len(manager.proxies) > 0 {
		logger.Info(fmt.Sprintf("🔄 Proxy manager initialized with %d proxies", len(manager.proxies)))
	}

	return manager, nil
}

// parseProxyAddress memparse alamat proxy
func parseProxyAddress(proxyAddr string) (*ProxyConfig, error) {
	u, err := url.Parse(proxyAddr)
	if err != nil {
		return nil, err
	}

	config := &ProxyConfig{
		Address: u.Host,
		Type:    u.Scheme,
		Active:  true,
	}

	if u.User != nil {
		config.Username = u.User.Username()
		config.Password, _ = u.User.Password()
	}

	return config, nil
}

// addDefaultProxies menambahkan proxy default untuk rotasi
func (m *Manager) addDefaultProxies() {
	// Proxy publik untuk testing (gunakan dengan hati-hati)
	defaultProxies := []ProxyConfig{
		{Address: "127.0.0.1:8080", Type: "http", Active: false},
		{Address: "127.0.0.1:3128", Type: "http", Active: false},
	}

	m.proxies = append(m.proxies, defaultProxies...)
}

// GetDialer mendapatkan dialer dengan proxy
func (m *Manager) GetDialer() (proxy.Dialer, error) {
	if len(m.proxies) == 0 {
		// Return direct dialer jika tidak ada proxy
		return &net.Dialer{
			Timeout: 10 * time.Second,
		}, nil
	}

	// Rotate proxy
	proxyConfig := m.getNextProxy()
	
	switch proxyConfig.Type {
	case "socks5":
		return m.createSOCKS5Dialer(proxyConfig)
	case "http", "https":
		return m.createHTTPDialer(proxyConfig)
	default:
		return nil, fmt.Errorf("unsupported proxy type: %s", proxyConfig.Type)
	}
}

// getNextProxy mendapatkan proxy berikutnya untuk rotasi
func (m *Manager) getNextProxy() ProxyConfig {
	if len(m.proxies) == 0 {
		return ProxyConfig{}
	}

	// Random selection untuk rotasi yang lebih baik
	if len(m.proxies) > 1 {
		m.currentIdx = rand.Intn(len(m.proxies))
	}

	proxy := m.proxies[m.currentIdx]
	m.currentIdx = (m.currentIdx + 1) % len(m.proxies)

	m.logger.Debug(fmt.Sprintf("Using proxy: %s (%s)", proxy.Address, proxy.Type))
	return proxy
}

// createSOCKS5Dialer membuat SOCKS5 dialer
func (m *Manager) createSOCKS5Dialer(config ProxyConfig) (proxy.Dialer, error) {
	var auth *proxy.Auth
	if config.Username != "" {
		auth = &proxy.Auth{
			User:     config.Username,
			Password: config.Password,
		}
	}

	return proxy.SOCKS5("tcp", config.Address, auth, &net.Dialer{
		Timeout: 10 * time.Second,
	})
}

// createHTTPDialer membuat HTTP proxy dialer
func (m *Manager) createHTTPDialer(config ProxyConfig) (proxy.Dialer, error) {
	// HTTP proxy implementation (simplified)
	return &HTTPProxyDialer{
		ProxyAddress: config.Address,
		Username:     config.Username,
		Password:     config.Password,
		Timeout:      10 * time.Second,
	}, nil
}

// HTTPProxyDialer implementasi sederhana HTTP proxy dialer
type HTTPProxyDialer struct {
	ProxyAddress string
	Username     string
	Password     string
	Timeout      time.Duration
}

// Dial implementasi proxy.Dialer interface
func (d *HTTPProxyDialer) Dial(network, addr string) (net.Conn, error) {
	// Connect ke proxy
	conn, err := net.DialTimeout("tcp", d.ProxyAddress, d.Timeout)
	if err != nil {
		return nil, err
	}

	// Send CONNECT request
	connectReq := fmt.Sprintf("CONNECT %s HTTP/1.1\r\nHost: %s\r\n", addr, addr)
	
	if d.Username != "" {
		// Basic auth
		auth := fmt.Sprintf("%s:%s", d.Username, d.Password)
		connectReq += fmt.Sprintf("Proxy-Authorization: Basic %s\r\n", auth)
	}
	
	connectReq += "\r\n"

	_, err = conn.Write([]byte(connectReq))
	if err != nil {
		conn.Close()
		return nil, err
	}

	// Read response (simplified)
	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		conn.Close()
		return nil, err
	}

	return conn, nil
}

// GetHTTPClient mendapatkan HTTP client dengan proxy
func (m *Manager) GetHTTPClient() (*http.Client, error) {
	if len(m.proxies) == 0 {
		return &http.Client{
			Timeout: 30 * time.Second,
		}, nil
	}

	proxyConfig := m.getNextProxy()
	
	var transport *http.Transport

	switch proxyConfig.Type {
	case "http", "https":
		proxyURL, err := url.Parse(fmt.Sprintf("%s://%s", proxyConfig.Type, proxyConfig.Address))
		if err != nil {
			return nil, err
		}

		if proxyConfig.Username != "" {
			proxyURL.User = url.UserPassword(proxyConfig.Username, proxyConfig.Password)
		}

		transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
			DialContext: (&net.Dialer{
				Timeout: 10 * time.Second,
			}).DialContext,
		}

	case "socks5":
		dialer, err := m.createSOCKS5Dialer(proxyConfig)
		if err != nil {
			return nil, err
		}

		transport = &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return dialer.Dial(network, addr)
			},
		}
	}

	return &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}, nil
}

// TestProxy mengetes konektivitas proxy
func (m *Manager) TestProxy(config ProxyConfig) error {
	m.logger.Debug(fmt.Sprintf("Testing proxy: %s", config.Address))

	dialer, err := m.createSOCKS5Dialer(config)
	if err != nil {
		return err
	}

	// Test connection ke Google DNS
	conn, err := dialer.Dial("tcp", "8.8.8.8:53")
	if err != nil {
		return err
	}
	defer conn.Close()

	m.logger.Debug(fmt.Sprintf("Proxy %s is working", config.Address))
	return nil
}

// RotateProxy secara manual merotasi ke proxy berikutnya
func (m *Manager) RotateProxy() {
	if len(m.proxies) > 1 {
		m.currentIdx = (m.currentIdx + 1) % len(m.proxies)
		m.logger.Debug("Rotated to next proxy")
	}
}

// GetActiveProxyCount mendapatkan jumlah proxy aktif
func (m *Manager) GetActiveProxyCount() int {
	count := 0
	for _, proxy := range m.proxies {
		if proxy.Active {
			count++
		}
	}
	return count
}
