package cmd

import (
        "fmt"

        "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
        Use:   "veko-grid",
        Short: "🛰️ Veko Grid - Tool eksplorasi jaringan anonim dan stealth",
        Long: `🛰️ Veko Grid adalah tool CLI berbasis Go untuk eksplorasi jaringan anonim
dan stealth scanning dengan support TOR, proxy rotation, dan grid-style network mapping.

Tool ini dibuat untuk:
• Penelitian akademik
• Audit keamanan jaringan sendiri  
• Eksplorasi infrastruktur domain publik secara legal dan anonim

Contoh penggunaan:
  veko-grid scan --input targets.txt --tor --proxy socks5://127.0.0.1:9050 --output results.json`,
        Version: "1.0.0",
}

func Execute() error {
        return rootCmd.Execute()
}

func init() {
        rootCmd.CompletionOptions.DisableDefaultCmd = true
        
        // Banner ASCII
        fmt.Println(`
╔═══════════════════════════════════════════════════════════════╗
║  🛰️  VEKO GRID v1.0.0 - Network Exploration & Stealth Tool   ║
║  📡 Anonymous Grid Scanning • TOR/Proxy Support • DNS/DoH    ║
║  🔐 TLS Fingerprint Spoofing • Academic Research Tool        ║
╚═══════════════════════════════════════════════════════════════╝
`)
}
