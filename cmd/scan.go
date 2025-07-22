package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"veko-grid/config"
	"veko-grid/core"
	"veko-grid/utils"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "🕸️ Melakukan grid-style network scanning",
	Long: `🕸️ Scan command melakukan eksplorasi jaringan dengan metode grid-style scanning.
	
Fitur utama:
• Port scanning dan ping detection
• DNS resolution (A/AAAA/MX/NS/CNAME)
• Traceroute dan CDN lookup
• Support TOR dan proxy rotation
• Random delay untuk stealth scanning
• TLS fingerprint randomization`,
	RunE: runScan,
}

var (
	inputFile   string
	outputFile  string
	proxyAddr   string
	useTor      bool
	delayRange  string
	timeout     int
	dnsMode     string
	silent      bool
	jsonOutput  bool
	debugMode   bool
	maxThreads  int
)

func init() {
	rootCmd.AddCommand(scanCmd)

	// Input/Output flags
	scanCmd.Flags().StringVarP(&inputFile, "input", "i", "", "File berisi daftar target (domain/IP)")
	scanCmd.Flags().StringVarP(&outputFile, "output", "o", "veko-results.json", "File output hasil scan")
	
	// Anonymity flags
	scanCmd.Flags().StringVarP(&proxyAddr, "proxy", "p", "", "Proxy address (socks5://127.0.0.1:9050)")
	scanCmd.Flags().BoolVar(&useTor, "tor", false, "Gunakan TOR untuk anonimitas")
	
	// Stealth flags
	scanCmd.Flags().StringVar(&delayRange, "delay", "100-500", "Random delay antar request (ms)")
	scanCmd.Flags().IntVar(&timeout, "timeout", 5, "Timeout koneksi (detik)")
	scanCmd.Flags().StringVar(&dnsMode, "dns", "default", "DNS mode: default/doh")
	
	// Output flags
	scanCmd.Flags().BoolVar(&silent, "silent", false, "Mode silent (minimal output)")
	scanCmd.Flags().BoolVar(&jsonOutput, "json", false, "Output dalam format JSON ke stdout")
	scanCmd.Flags().BoolVar(&debugMode, "debug", false, "Enable debug logging")
	
	// Performance flags
	scanCmd.Flags().IntVar(&maxThreads, "threads", 10, "Maksimum thread concurrent")

	// Required flags
	scanCmd.MarkFlagRequired("input")
}

func runScan(cmd *cobra.Command, args []string) error {
	// Initialize logger
	logger := utils.NewLogger(debugMode, silent)
	
	if !silent {
		fmt.Println("🚀 Memulai Veko Grid Scanning...")
	}

	// Baca konfigurasi
	cfg := &config.Config{
		InputFile:   inputFile,
		OutputFile:  outputFile,
		ProxyAddr:   proxyAddr,
		UseTor:      useTor,
		DelayRange:  delayRange,
		Timeout:     timeout,
		DNSMode:     dnsMode,
		Silent:      silent,
		JSONOutput:  jsonOutput,
		Debug:       debugMode,
		MaxThreads:  maxThreads,
	}

	// Validasi file input
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		return fmt.Errorf("❌ File input tidak ditemukan: %s", inputFile)
	}

	// Baca targets dari file
	targets, err := readTargetsFromFile(inputFile)
	if err != nil {
		return fmt.Errorf("❌ Error membaca file targets: %v", err)
	}

	if len(targets) == 0 {
		return fmt.Errorf("❌ Tidak ada target yang valid ditemukan")
	}

	logger.Info(fmt.Sprintf("📋 Loaded %d targets untuk scanning", len(targets)))

	// Initialize scanner
	scanner, err := core.NewScanner(cfg, logger)
	if err != nil {
		return fmt.Errorf("❌ Error inisialisasi scanner: %v", err)
	}

	// Mulai scanning
	results, err := scanner.ScanTargets(targets)
	if err != nil {
		return fmt.Errorf("❌ Error saat scanning: %v", err)
	}

	// Output hasil
	outputHandler := utils.NewOutputHandler(cfg, logger)
	if err := outputHandler.SaveResults(results); err != nil {
		return fmt.Errorf("❌ Error menyimpan hasil: %v", err)
	}

	if !silent {
		fmt.Printf("✅ Scanning selesai! Hasil disimpan di: %s\n", outputFile)
		fmt.Printf("📊 Total scanned: %d targets\n", len(results))
	}

	return nil
}

func readTargetsFromFile(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	var targets []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "#") {
			targets = append(targets, line)
		}
	}

	return targets, nil
}
