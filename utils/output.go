package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"veko-grid/config"
)

// OutputHandler mengelola output hasil scanning
type OutputHandler struct {
	config *config.Config
	logger *Logger
}

// ScanOutput menyimpan format output untuk semua hasil
type ScanOutput struct {
	Metadata *ScanMetadata `json:"metadata"`
	Results  interface{}   `json:"results"`
}

// ScanMetadata menyimpan metadata scanning
type ScanMetadata struct {
	Tool        string    `json:"tool"`
	Version     string    `json:"version"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Duration    string    `json:"duration"`
	TotalHosts  int       `json:"total_hosts"`
	Successful  int       `json:"successful"`
	Failed      int       `json:"failed"`
	Config      *ScanConfigSummary `json:"config"`
}

// ScanConfigSummary menyimpan ringkasan konfigurasi scanning
type ScanConfigSummary struct {
	UseTor      bool   `json:"use_tor"`
	ProxyAddr   string `json:"proxy_addr,omitempty"`
	DNSMode     string `json:"dns_mode"`
	DelayRange  string `json:"delay_range"`
	Timeout     int    `json:"timeout"`
	MaxThreads  int    `json:"max_threads"`
}

// NewOutputHandler membuat instance OutputHandler baru
func NewOutputHandler(cfg *config.Config, logger *Logger) *OutputHandler {
	return &OutputHandler{
		config: cfg,
		logger: logger,
	}
}

// SaveResults menyimpan hasil scanning ke file
func (o *OutputHandler) SaveResults(results interface{}) error {
	// Tentukan format output berdasarkan ekstensi file
	ext := strings.ToLower(filepath.Ext(o.config.OutputFile))
	
	switch ext {
	case ".json":
		return o.saveAsJSON(results)
	case ".csv":
		return o.saveAsCSV(results)
	default:
		// Default ke JSON
		return o.saveAsJSON(results)
	}
}

// saveAsJSON menyimpan hasil dalam format JSON
func (o *OutputHandler) saveAsJSON(results interface{}) error {
	output := &ScanOutput{
		Metadata: o.generateMetadata(results),
		Results:  results,
	}

	// Create output directory if not exists
	dir := filepath.Dir(o.config.OutputFile)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	file, err := os.Create(o.config.OutputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty print

	if err := encoder.Encode(output); err != nil {
		return fmt.Errorf("failed to encode JSON: %v", err)
	}

	o.logger.Info(fmt.Sprintf("📄 Hasil disimpan dalam format JSON: %s", o.config.OutputFile))
	return nil
}

// saveAsCSV menyimpan hasil dalam format CSV
func (o *OutputHandler) saveAsCSV(results interface{}) error {
	// Create output directory if not exists
	dir := filepath.Dir(o.config.OutputFile)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	file, err := os.Create(o.config.OutputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	header := []string{
		"Target", "IP", "Timestamp", "Open Ports", "Services", 
		"CDN Provider", "TLS Version", "Scan Time", "Error",
	}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("failed to write CSV header: %v", err)
	}

	// Convert results to CSV format
	if err := o.writeCSVData(writer, results); err != nil {
		return fmt.Errorf("failed to write CSV data: %v", err)
	}

	o.logger.Info(fmt.Sprintf("📊 Hasil disimpan dalam format CSV: %s", o.config.OutputFile))
	return nil
}

// writeCSVData menulis data hasil ke CSV
func (o *OutputHandler) writeCSVData(writer *csv.Writer, results interface{}) error {
	// Type assertion untuk mendapatkan slice of ScanResult
	scanResults, ok := results.([]*ScanResult)
	if !ok {
		return fmt.Errorf("invalid results type for CSV output")
	}

	for _, result := range scanResults {
		// Convert open ports to string
		openPorts := ""
		if len(result.OpenPorts) > 0 {
			ports := make([]string, len(result.OpenPorts))
			for i, port := range result.OpenPorts {
				ports[i] = fmt.Sprintf("%d", port)
			}
			openPorts = strings.Join(ports, ";")
		}

		// Convert services to string
		services := ""
		if len(result.Services) > 0 {
			var serviceList []string
			for port, service := range result.Services {
				serviceList = append(serviceList, fmt.Sprintf("%d:%s", port, service))
			}
			services = strings.Join(serviceList, ";")
		}

		// Extract CDN provider
		cdnProvider := ""
		if result.CDNInfo != nil {
			if provider, ok := result.CDNInfo["provider"]; ok {
				cdnProvider = fmt.Sprintf("%v", provider)
			}
		}

		// Extract TLS version
		tlsVersion := ""
		if result.TLSInfo != nil {
			if fp, ok := result.TLSInfo["fingerprint"]; ok {
				if fpMap, ok := fp.(map[string]interface{}); ok {
					if ver, ok := fpMap["tls_version"]; ok {
						tlsVersion = fmt.Sprintf("%v", ver)
					}
				}
			}
		}

		// Write row
		row := []string{
			result.Target,
			result.IP,
			result.Timestamp.Format(time.RFC3339),
			openPorts,
			services,
			cdnProvider,
			tlsVersion,
			result.ScanTime.String(),
			result.Error,
		}

		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}

// generateMetadata menghasilkan metadata untuk output
func (o *OutputHandler) generateMetadata(results interface{}) *ScanMetadata {
	metadata := &ScanMetadata{
		Tool:      "Veko Grid",
		Version:   "1.0.0",
		StartTime: time.Now(), // Seharusnya dari waktu mulai scanning
		EndTime:   time.Now(),
		Config: &ScanConfigSummary{
			UseTor:     o.config.UseTor,
			ProxyAddr:  o.config.ProxyAddr,
			DNSMode:    o.config.DNSMode,
			DelayRange: o.config.DelayRange,
			Timeout:    o.config.Timeout,
			MaxThreads: o.config.MaxThreads,
		},
	}

	// Calculate statistics
	if scanResults, ok := results.([]*ScanResult); ok {
		metadata.TotalHosts = len(scanResults)
		
		for _, result := range scanResults {
			if result.Error != "" {
				metadata.Failed++
			} else {
				metadata.Successful++
			}
		}
	}

	metadata.Duration = metadata.EndTime.Sub(metadata.StartTime).String()
	return metadata
}

// PrintSummary mencetak ringkasan hasil ke terminal
func (o *OutputHandler) PrintSummary(results interface{}) {
	if o.config.Silent {
		return
	}

	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("📊 VEKO GRID SCAN SUMMARY")
	fmt.Println(strings.Repeat("=", 70))

	scanResults, ok := results.([]*ScanResult)
	if !ok {
		fmt.Println("❌ Invalid results format")
		return
	}

	// Statistics
	var successful, failed, totalPorts int
	var hosts []string

	for _, result := range scanResults {
		if result.Error != "" {
			failed++
		} else {
			successful++
			totalPorts += len(result.OpenPorts)
			if result.IP != "" {
				hosts = append(hosts, result.IP)
			}
		}
	}

	fmt.Printf("🎯 Total Targets Scanned: %d\n", len(scanResults))
	fmt.Printf("✅ Successful: %d (%.1f%%)\n", successful, float64(successful)/float64(len(scanResults))*100)
	fmt.Printf("❌ Failed: %d (%.1f%%)\n", failed, float64(failed)/float64(len(scanResults))*100)
	fmt.Printf("🔓 Total Open Ports: %d\n", totalPorts)
	fmt.Printf("🌐 Unique IPs: %d\n", len(o.uniqueStrings(hosts)))

	// Top results
	o.printTopResults(scanResults)

	fmt.Println(strings.Repeat("=", 70))
	fmt.Printf("💾 Hasil lengkap disimpan di: %s\n", o.config.OutputFile)
	fmt.Println(strings.Repeat("=", 70))
}

// printTopResults mencetak top results
func (o *OutputHandler) printTopResults(results []*ScanResult) {
	fmt.Println("\n🔝 TOP FINDINGS:")

	// Hosts dengan port terbanyak
	type hostPorts struct {
		Target string
		Ports  int
	}

	var topHosts []hostPorts
	for _, result := range results {
		if len(result.OpenPorts) > 0 {
			topHosts = append(topHosts, hostPorts{
				Target: result.Target,
				Ports:  len(result.OpenPorts),
			})
		}
	}

	// Sort by port count
	for i := 0; i < len(topHosts)-1; i++ {
		for j := i + 1; j < len(topHosts); j++ {
			if topHosts[i].Ports < topHosts[j].Ports {
				topHosts[i], topHosts[j] = topHosts[j], topHosts[i]
			}
		}
	}

	// Print top 5
	limit := 5
	if len(topHosts) < limit {
		limit = len(topHosts)
	}

	for i := 0; i < limit; i++ {
		fmt.Printf("  %d. %s - %d ports\n", i+1, topHosts[i].Target, topHosts[i].Ports)
	}
}

// uniqueStrings menghilangkan duplikat dari slice string
func (o *OutputHandler) uniqueStrings(input []string) []string {
	keys := make(map[string]bool)
	var result []string

	for _, item := range input {
		if !keys[item] {
			keys[item] = true
			result = append(result, item)
		}
	}

	return result
}

// OutputToStdout mencetak hasil ke stdout dalam format JSON
func (o *OutputHandler) OutputToStdout(results interface{}) error {
	if !o.config.JSONOutput {
		return nil
	}

	output := &ScanOutput{
		Metadata: o.generateMetadata(results),
		Results:  results,
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")

	return encoder.Encode(output)
}

// CreateReport membuat laporan scan yang lebih detail
func (o *OutputHandler) CreateReport(results interface{}) error {
	reportFile := strings.TrimSuffix(o.config.OutputFile, filepath.Ext(o.config.OutputFile)) + "_report.html"
	
	html := o.generateHTMLReport(results)
	
	if err := os.WriteFile(reportFile, []byte(html), 0644); err != nil {
		return fmt.Errorf("failed to create HTML report: %v", err)
	}

	o.logger.Info(fmt.Sprintf("📋 HTML report created: %s", reportFile))
	return nil
}

// generateHTMLReport menghasilkan laporan HTML
func (o *OutputHandler) generateHTMLReport(results interface{}) string {
	// Simplified HTML report template
	html := `<!DOCTYPE html>
<html>
<head>
    <title>Veko Grid Scan Report</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 20px; }
        .header { background: #f0f0f0; padding: 20px; border-radius: 5px; }
        .result { margin: 10px 0; padding: 10px; border: 1px solid #ddd; }
        .success { border-left: 5px solid #4CAF50; }
        .failed { border-left: 5px solid #f44336; }
    </style>
</head>
<body>
    <div class="header">
        <h1>🛰️ Veko Grid Scan Report</h1>
        <p>Generated: ` + time.Now().Format("2006-01-02 15:04:05") + `</p>
    </div>
    <div class="content">
        <h2>Scan Results</h2>
        <!-- Results will be added here -->
    </div>
</body>
</html>`

	return html
}

// ScanResult struct untuk type assertion (jika belum didefinisikan di tempat lain)
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
