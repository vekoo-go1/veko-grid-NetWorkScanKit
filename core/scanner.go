package core

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"

	"veko-grid/config"
	"veko-grid/proxy"
	"veko-grid/utils"
)

// Scanner adalah struct utama untuk melakukan network scanning
type Scanner struct {
	config       *config.Config
	logger       *utils.Logger
	proxyManager *proxy.Manager
	dnsResolver  *utils.DNSResolver
	fingerprint  *utils.FingerprintSpoofer
	grid         *Grid
}

// ScanResult menyimpan hasil scanning untuk satu target
type ScanResult struct {
	Target      string                 `json:"target"`
	IP          string                 `json:"ip,omitempty"`
	Timestamp   time.Time              `json:"timestamp"`
	DNSRecords  map[string][]string    `json:"dns_records,omitempty"`
	OpenPorts   []int                  `json:"open_ports,omitempty"`
	Services    map[int]string         `json:"services,omitempty"`
	Traceroute  []string               `json:"traceroute,omitempty"`
	CDNInfo     map[string]interface{} `json:"cdn_info,omitempty"`
	TLSInfo     map[string]interface{} `json:"tls_info,omitempty"`
	Error       string                 `json:"error,omitempty"`
	ScanTime    time.Duration          `json:"scan_time"`
}

// NewScanner membuat instance Scanner baru
func NewScanner(cfg *config.Config, logger *utils.Logger) (*Scanner, error) {
	scanner := &Scanner{
		config: cfg,
		logger: logger,
	}

	// Initialize proxy manager
	proxyMgr, err := proxy.NewManager(cfg.ProxyAddr, cfg.UseTor, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize proxy manager: %v", err)
	}
	scanner.proxyManager = proxyMgr

	// Initialize DNS resolver
	dnsResolver, err := utils.NewDNSResolver(cfg.IsDoHEnabled(), logger)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize DNS resolver: %v", err)
	}
	scanner.dnsResolver = dnsResolver

	// Initialize fingerprint spoofer
	scanner.fingerprint = utils.NewFingerprintSpoofer(logger)

	// Initialize grid scanner
	scanner.grid = NewGrid(cfg, logger)

	return scanner, nil
}

// ScanTargets melakukan scanning terhadap list target
func (s *Scanner) ScanTargets(targets []string) ([]*ScanResult, error) {
	s.logger.Info(fmt.Sprintf("🎯 Memulai scanning %d targets", len(targets)))

	var results []*ScanResult
	var mutex sync.Mutex
	var wg sync.WaitGroup

	// Channel untuk limit concurrent scanning
	semaphore := make(chan struct{}, s.config.MaxThreads)

	for i, target := range targets {
		wg.Add(1)
		go func(idx int, tgt string) {
			defer wg.Done()

			// Acquire semaphore
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			// Random delay untuk stealth
			s.randomDelay()

			// Scan single target
			result := s.scanSingleTarget(tgt, idx+1, len(targets))

			// Add to results
			mutex.Lock()
			results = append(results, result)
			mutex.Unlock()

		}(i, target)
	}

	wg.Wait()
	s.logger.Info("✅ Semua target selesai di-scan")

	return results, nil
}

// scanSingleTarget melakukan scanning untuk satu target
func (s *Scanner) scanSingleTarget(target string, current, total int) *ScanResult {
	startTime := time.Now()
	
	result := &ScanResult{
		Target:    target,
		Timestamp: startTime,
	}

	if !s.config.Silent {
		progress := fmt.Sprintf("[%d/%d]", current, total)
		s.logger.Info(fmt.Sprintf("🔍 %s Scanning: %s", progress, target))
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.config.GetTimeout())
	defer cancel()

	// DNS Resolution
	if dnsRecords, err := s.dnsResolver.ResolveAll(target); err == nil {
		result.DNSRecords = dnsRecords
		if ips, ok := dnsRecords["A"]; ok && len(ips) > 0 {
			result.IP = ips[0]
		}
	} else {
		s.logger.Debug(fmt.Sprintf("DNS resolution failed for %s: %v", target, err))
	}

	// Port Scanning
	if result.IP != "" {
		openPorts, services := s.scanPorts(ctx, result.IP)
		result.OpenPorts = openPorts
		result.Services = services
	}

	// Traceroute (simplified)
	if result.IP != "" {
		if traceroute, err := s.performTraceroute(ctx, result.IP); err == nil {
			result.Traceroute = traceroute
		}
	}

	// CDN Detection
	if cdnInfo := s.detectCDN(target); cdnInfo != nil {
		result.CDNInfo = cdnInfo
	}

	// TLS Fingerprinting
	if tlsInfo := s.performTLSFingerprinting(target); tlsInfo != nil {
		result.TLSInfo = tlsInfo
	}

	result.ScanTime = time.Since(startTime)

	if !s.config.Silent {
		s.displayScanResult(result)
	}

	return result
}

// randomDelay menerapkan delay random untuk stealth
func (s *Scanner) randomDelay() {
	minDelay, maxDelay, _ := s.config.GetDelayRange()
	
	// Generate random delay between min and max
	delayRange := maxDelay - minDelay
	randomDelay := minDelay + time.Duration(rand.Int63n(int64(delayRange)))
	
	time.Sleep(randomDelay)
}

// scanPorts melakukan port scanning
func (s *Scanner) scanPorts(ctx context.Context, ip string) ([]int, map[int]string) {
	commonPorts := []int{21, 22, 23, 25, 53, 80, 110, 443, 993, 995, 8080, 8443}
	var openPorts []int
	services := make(map[int]string)

	for _, port := range commonPorts {
		if s.isPortOpen(ctx, ip, port) {
			openPorts = append(openPorts, port)
			services[port] = s.identifyService(port)
		}
	}

	return openPorts, services
}

// isPortOpen mengecek apakah port terbuka
func (s *Scanner) isPortOpen(ctx context.Context, ip string, port int) bool {
	address := fmt.Sprintf("%s:%d", ip, port)
	
	conn, err := net.DialTimeout("tcp", address, 3*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	
	return true
}

// identifyService mengidentifikasi service berdasarkan port
func (s *Scanner) identifyService(port int) string {
	services := map[int]string{
		21:   "ftp",
		22:   "ssh", 
		23:   "telnet",
		25:   "smtp",
		53:   "dns",
		80:   "http",
		110:  "pop3",
		443:  "https",
		993:  "imaps",
		995:  "pop3s",
		8080: "http-alt",
		8443: "https-alt",
	}
	
	if service, exists := services[port]; exists {
		return service
	}
	return "unknown"
}

// performTraceroute melakukan traceroute sederhana
func (s *Scanner) performTraceroute(ctx context.Context, ip string) ([]string, error) {
	// Simplified traceroute implementation
	var hops []string
	
	// Untuk demo, kita buat traceroute sederhana
	// Di implementasi nyata, bisa menggunakan raw sockets
	hops = append(hops, "192.168.1.1")  // Gateway lokal
	hops = append(hops, ip)             // Target IP
	
	return hops, nil
}

// detectCDN mendeteksi penggunaan CDN
func (s *Scanner) detectCDN(target string) map[string]interface{} {
	cdnInfo := make(map[string]interface{})
	
	// DNS-based CDN detection
	if cnames, err := s.dnsResolver.LookupCNAME(target); err == nil {
		for _, cname := range cnames {
			if s.isCDNDomain(cname) {
				cdnInfo["provider"] = s.identifyCDNProvider(cname)
				cdnInfo["cname"] = cname
				break
			}
		}
	}
	
	if len(cdnInfo) == 0 {
		return nil
	}
	
	return cdnInfo
}

// isCDNDomain mengecek apakah domain adalah CDN
func (s *Scanner) isCDNDomain(domain string) bool {
	cdnPatterns := []string{
		"cloudflare.com",
		"fastly.com", 
		"amazonaws.com",
		"azureedge.net",
		"cdn77.com",
		"maxcdn.com",
	}
	
	for _, pattern := range cdnPatterns {
		if contains(domain, pattern) {
			return true
		}
	}
	
	return false
}

// identifyCDNProvider mengidentifikasi provider CDN
func (s *Scanner) identifyCDNProvider(domain string) string {
	if contains(domain, "cloudflare") {
		return "Cloudflare"
	} else if contains(domain, "fastly") {
		return "Fastly"
	} else if contains(domain, "amazonaws") {
		return "Amazon CloudFront"
	} else if contains(domain, "azureedge") {
		return "Azure CDN"
	}
	
	return "Unknown CDN"
}

// performTLSFingerprinting melakukan TLS fingerprinting
func (s *Scanner) performTLSFingerprinting(target string) map[string]interface{} {
	return s.fingerprint.AnalyzeTLS(target)
}

// displayScanResult menampilkan hasil scan ke terminal
func (s *Scanner) displayScanResult(result *ScanResult) {
	fmt.Printf("  📍 Target: %s", result.Target)
	if result.IP != "" {
		fmt.Printf(" (%s)", result.IP)
	}
	fmt.Println()

	if len(result.OpenPorts) > 0 {
		fmt.Printf("    🔓 Open Ports: %v\n", result.OpenPorts)
	}

	if result.CDNInfo != nil {
		if provider, ok := result.CDNInfo["provider"]; ok {
			fmt.Printf("    🌐 CDN: %s\n", provider)
		}
	}

	fmt.Printf("    ⏱️  Scan Time: %v\n", result.ScanTime.Round(time.Millisecond))
	fmt.Println()
}

// Helper function
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || 
		(len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr)))
}
