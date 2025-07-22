package core

import (
	"fmt"
	"math"
	"sort"
	"time"

	"veko-grid/config"
	"veko-grid/utils"
)

// Grid mengelola grid-style scanning dan visualisasi
type Grid struct {
	config *config.Config
	logger *utils.Logger
}

// GridCell merepresentasikan satu cell dalam grid
type GridCell struct {
	Target     string
	Status     string // "scanning", "success", "failed", "pending"
	Result     *ScanResult
	Position   GridPosition
}

// GridPosition menyimpan posisi dalam grid
type GridPosition struct {
	Row    int
	Column int
}

// NewGrid membuat instance Grid baru
func NewGrid(cfg *config.Config, logger *utils.Logger) *Grid {
	return &Grid{
		config: cfg,
		logger: logger,
	}
}

// DisplayGridProgress menampilkan progress scanning dalam bentuk grid
func (g *Grid) DisplayGridProgress(targets []string, results []*ScanResult) {
	if g.config.Silent {
		return
	}

	// Hitung dimensi grid yang optimal
	totalTargets := len(targets)
	gridSize := g.calculateOptimalGridSize(totalTargets)
	
	fmt.Printf("\n📊 Grid Scanning Progress (%dx%d):\n", gridSize.Row, gridSize.Column)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	// Buat mapping hasil
	resultMap := make(map[string]*ScanResult)
	for _, result := range results {
		resultMap[result.Target] = result
	}

	// Display grid
	for row := 0; row < gridSize.Row; row++ {
		fmt.Print("║ ")
		for col := 0; col < gridSize.Column; col++ {
			index := row*gridSize.Column + col
			if index < len(targets) {
				target := targets[index]
				symbol := g.getStatusSymbol(target, resultMap)
				fmt.Printf("%s ", symbol)
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println(" ║")
	}

	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	g.displayLegend()
	g.displayStats(results)
}

// calculateOptimalGridSize menghitung ukuran grid yang optimal
func (g *Grid) calculateOptimalGridSize(totalTargets int) GridPosition {
	if totalTargets <= 0 {
		return GridPosition{Row: 1, Column: 1}
	}

	// Cari rasio yang mendekati square
	sqrt := math.Sqrt(float64(totalTargets))
	rows := int(sqrt)
	cols := int(math.Ceil(float64(totalTargets) / float64(rows)))

	// Adjust untuk tampilan yang lebih baik
	if cols > 50 { // Batasi lebar maksimal
		cols = 50
		rows = int(math.Ceil(float64(totalTargets) / float64(cols)))
	}

	return GridPosition{Row: rows, Column: cols}
}

// getStatusSymbol mendapatkan symbol untuk status target
func (g *Grid) getStatusSymbol(target string, resultMap map[string]*ScanResult) string {
	result, exists := resultMap[target]
	if !exists {
		return "⏳" // Pending
	}

	if result.Error != "" {
		return "❌" // Failed
	}

	// Success dengan gradasi berdasarkan hasil
	if len(result.OpenPorts) > 5 {
		return "🔴" // Banyak port terbuka
	} else if len(result.OpenPorts) > 0 {
		return "🟡" // Beberapa port terbuka
	} else {
		return "🟢" // Host aktif tapi port tertutup
	}
}

// displayLegend menampilkan legend untuk grid
func (g *Grid) displayLegend() {
	fmt.Println("\n📋 Legend:")
	fmt.Println("  ⏳ Pending   🟢 Host Active   🟡 Some Ports   🔴 Many Ports   ❌ Failed")
}

// displayStats menampilkan statistik scanning
func (g *Grid) displayStats(results []*ScanResult) {
	if len(results) == 0 {
		return
	}

	var successful, failed int
	var totalPorts int
	var totalScanTime time.Duration
	var avgScanTime time.Duration

	portStats := make(map[int]int)

	for _, result := range results {
		if result.Error != "" {
			failed++
		} else {
			successful++
			totalPorts += len(result.OpenPorts)
			
			// Count port occurrences
			for _, port := range result.OpenPorts {
				portStats[port]++
			}
		}
		totalScanTime += result.ScanTime
	}

	if len(results) > 0 {
		avgScanTime = totalScanTime / time.Duration(len(results))
	}

	fmt.Printf("\n📈 Scanning Statistics:\n")
	fmt.Printf("  🎯 Total Targets: %d\n", len(results))
	fmt.Printf("  ✅ Successful: %d (%.1f%%)\n", successful, float64(successful)/float64(len(results))*100)
	fmt.Printf("  ❌ Failed: %d (%.1f%%)\n", failed, float64(failed)/float64(len(results))*100)
	fmt.Printf("  🔓 Total Open Ports: %d\n", totalPorts)
	fmt.Printf("  ⏱️  Average Scan Time: %v\n", avgScanTime.Round(time.Millisecond))

	// Top ports
	g.displayTopPorts(portStats)
}

// displayTopPorts menampilkan port yang paling sering ditemukan
func (g *Grid) displayTopPorts(portStats map[int]int) {
	if len(portStats) == 0 {
		return
	}

	// Sort ports by frequency
	type portCount struct {
		Port  int
		Count int
	}

	var sortedPorts []portCount
	for port, count := range portStats {
		sortedPorts = append(sortedPorts, portCount{Port: port, Count: count})
	}

	sort.Slice(sortedPorts, func(i, j int) bool {
		return sortedPorts[i].Count > sortedPorts[j].Count
	})

	fmt.Printf("\n🔝 Top Open Ports:\n")
	limit := 5
	if len(sortedPorts) < limit {
		limit = len(sortedPorts)
	}

	for i := 0; i < limit; i++ {
		pc := sortedPorts[i]
		percentage := float64(pc.Count) / float64(len(portStats)) * 100
		fmt.Printf("  %d: %d targets (%.1f%%)\n", pc.Port, pc.Count, percentage)
	}
}

// DisplayRealTimeGrid menampilkan grid yang update secara real-time
func (g *Grid) DisplayRealTimeGrid(targets []string, resultChan <-chan *ScanResult) {
	if g.config.Silent {
		return
	}

	results := make([]*ScanResult, 0)
	
	// Ticker untuk update display
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case result, ok := <-resultChan:
			if !ok {
				// Channel closed, final display
				g.DisplayGridProgress(targets, results)
				return
			}
			results = append(results, result)

		case <-ticker.C:
			// Update display
			fmt.Print("\033[2J\033[H") // Clear screen
			g.DisplayGridProgress(targets, results)
		}
	}
}
