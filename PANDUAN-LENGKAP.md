# 🛰️ PANDUAN LENGKAP VEKO GRID
## Tool CLI Network Exploration & Stealth Scanning

---

## 📋 DAFTAR ISI

1. [Pengenalan Tool](#pengenalan-tool)
2. [Persiapan dan Instalasi](#persiapan-dan-instalasi)
3. [Cara Penggunaan Dasar](#cara-penggunaan-dasar)
4. [Fitur-Fitur Advanced](#fitur-fitur-advanced)
5. [Format File Input dan Output](#format-file-input-dan-output)
6. [Contoh Kasus Penggunaan](#contoh-kasus-penggunaan)
7. [Troubleshooting](#troubleshooting)
8. [Tips dan Best Practices](#tips-dan-best-practices)
9. [Legal dan Etika](#legal-dan-etika)

---

## 🎯 PENGENALAN TOOL

### Apa itu Veko Grid?

Veko Grid adalah tool CLI (Command Line Interface) yang dibuat dengan bahasa Go untuk melakukan network exploration dan stealth scanning. Tool ini dirancang khusus untuk:

- **Penelitian akademik** tentang keamanan jaringan
- **Audit infrastruktur** milik sendiri atau dengan izin
- **Bug bounty hunting** sesuai scope yang diizinkan
- **Testing penetrasi** dengan otorisasi yang jelas

### Fitur Utama

✅ **Network Scanning**: Port scanning, ping, traceroute
✅ **DNS Resolution**: A/AAAA/MX/NS/CNAME/TXT records
✅ **Anonymity**: Support TOR dan proxy rotation
✅ **Stealth Mode**: Random delays, header spoofing
✅ **Multi-format Output**: JSON, CSV, HTML reports
✅ **Grid Visualization**: ASCII-based progress display

---

## 💻 PERSIAPAN DAN INSTALASI

### Sistem Requirements

- **Windows**: Windows 10/11 (64-bit)
- **RAM**: Minimal 1GB available
- **Storage**: 50MB untuk tool + space untuk results
- **Network**: Koneksi internet untuk scanning

### File yang Diperlukan

1. **veko-grid.exe** (9.4MB) - File utama executable
2. **targets.txt** - File berisi daftar target yang akan di-scan
3. **Folder output** - Tempat menyimpan hasil scanning

### Setup Awal

#### 1. Persiapan Folder
```cmd
# Buat folder kerja
mkdir C:\VekoGrid
cd C:\VekoGrid

# Copy file executable
copy veko-grid.exe C:\VekoGrid\
```

#### 2. Test Installation
```cmd
# Test apakah tool berjalan
veko-grid.exe --help

# Test version
veko-grid.exe --version
```

Jika berhasil, Anda akan melihat:
```
╔═══════════════════════════════════════════════════════════════╗
║  🛰️  VEKO GRID v1.0.0 - Network Exploration & Stealth Tool   ║
║  📡 Anonymous Grid Scanning • TOR/Proxy Support • DNS/DoH    ║
║  🔐 TLS Fingerprint Spoofing • Academic Research Tool        ║
╚═══════════════════════════════════════════════════════════════╝
```

---

## 🚀 CARA PENGGUNAAN DASAR

### Command Structure

```cmd
veko-grid.exe [COMMAND] [FLAGS] [OPTIONS]
```

### Commands Utama

#### 1. Help dan Informasi
```cmd
# Tampilkan help utama
veko-grid.exe --help

# Help untuk command scan
veko-grid.exe scan --help

# Tampilkan versi
veko-grid.exe --version
```

#### 2. Basic Scanning
```cmd
# Scan sederhana
veko-grid.exe scan --input targets.txt --output results.json

# Scan dengan mode debug (verbose)
veko-grid.exe scan --input targets.txt --output results.json --debug

# Scan mode silent (minimal output)
veko-grid.exe scan --input targets.txt --output results.json --silent
```

### Flag dan Parameter Penting

| Flag | Deskripsi | Contoh |
|------|-----------|---------|
| `--input` | File berisi target yang akan di-scan | `--input targets.txt` |
| `--output` | File output untuk hasil scanning | `--output results.json` |
| `--debug` | Mode verbose dengan detail logging | `--debug` |
| `--silent` | Mode tanpa output ke console | `--silent` |
| `--json` | Output langsung ke stdout dalam JSON | `--json` |

---

## 🔧 FITUR-FITUR ADVANCED

### 1. Anonymity dan Proxy

#### TOR Integration
```cmd
# Menggunakan TOR untuk anonimitas
veko-grid.exe scan --input targets.txt --tor --output results.json

# TOR dengan delay custom
veko-grid.exe scan --input targets.txt --tor --delay 500-2000 --output results.json
```

**Catatan**: Memerlukan TOR daemon yang berjalan di port 9050

#### Proxy Support
```cmd
# SOCKS5 proxy
veko-grid.exe scan --input targets.txt --proxy socks5://127.0.0.1:9050 --output results.json

# HTTP proxy
veko-grid.exe scan --input targets.txt --proxy http://proxy.example.com:8080 --output results.json

# Kombinasi TOR + Proxy
veko-grid.exe scan --input targets.txt --tor --proxy socks5://127.0.0.1:9050 --output results.json
```

### 2. DNS Configuration

#### DNS over HTTPS (DoH)
```cmd
# Menggunakan DoH untuk privasi DNS
veko-grid.exe scan --input targets.txt --dns doh --output results.json

# DoH dengan timeout custom
veko-grid.exe scan --input targets.txt --dns doh --timeout 10 --output results.json
```

### 3. Performance Tuning

#### Threading dan Delays
```cmd
# Lebih banyak thread untuk scanning cepat
veko-grid.exe scan --input targets.txt --threads 20 --output results.json

# Custom delay range (milliseconds)
veko-grid.exe scan --input targets.txt --delay 100-500 --output results.json

# Timeout custom (seconds)
veko-grid.exe scan --input targets.txt --timeout 15 --output results.json
```

#### Kombinasi Optimal
```cmd
# Setup untuk scanning cepat tapi stealth
veko-grid.exe scan --input targets.txt --threads 10 --delay 200-800 --timeout 10 --output results.json

# Setup untuk maximum stealth
veko-grid.exe scan --input targets.txt --tor --delay 1000-5000 --timeout 30 --output results.json
```

---

## 📁 FORMAT FILE INPUT DAN OUTPUT

### File Input (targets.txt)

#### Format Dasar
```
# Veko Grid Targets File
# Format: satu domain/IP per baris
# Gunakan # untuk komentar

# Domain examples
example.com
google.com
github.com
subdomain.example.com

# IP address examples
192.168.1.1
8.8.8.8
1.1.1.1

# Port specific (opsional)
example.com:80
192.168.1.1:22
```

#### Contoh File Target Lengkap
```
# === PERSONAL RESEARCH TARGETS ===
# Website pribadi untuk audit
mywebsite.com
blog.mywebsite.com
api.mywebsite.com

# === PUBLIC DNS SERVERS ===
# Google DNS
8.8.8.8
8.8.4.4
dns.google

# Cloudflare DNS
1.1.1.1
1.0.0.1
one.one.one.one

# === MAJOR WEBSITES FOR STUDY ===
# Popular sites for educational research
github.com
stackoverflow.com
reddit.com

# === CORPORATE INFRASTRUCTURE ===
# (Hanya jika memiliki izin tertulis)
# corp.example.com
# mail.example.com
```

### File Output

#### JSON Format (Recommended)
```json
{
  "metadata": {
    "tool": "Veko Grid",
    "version": "1.0.0",
    "scan_start": "2024-07-22T10:00:00Z",
    "scan_end": "2024-07-22T10:05:30Z",
    "total_targets": 10,
    "successful_scans": 8,
    "failed_scans": 2,
    "scan_options": {
      "threads": 8,
      "timeout": 5,
      "delay_range": "100-500ms",
      "tor_enabled": false,
      "proxy": null
    }
  },
  "results": [
    {
      "target": "example.com",
      "ip_address": "93.184.216.34",
      "timestamp": "2024-07-22T10:00:15Z",
      "scan_duration": "2.5s",
      "dns_records": {
        "A": ["93.184.216.34"],
        "AAAA": ["2606:2800:220:1:248:1893:25c8:1946"],
        "MX": ["0 ."],
        "NS": ["a.iana-servers.net.", "b.iana-servers.net."],
        "TXT": ["v=spf1 -all"]
      },
      "open_ports": [80, 443],
      "port_details": {
        "80": {
          "service": "http",
          "banner": "nginx/1.18.0",
          "status": "open"
        },
        "443": {
          "service": "https",
          "banner": "nginx/1.18.0",
          "status": "open",
          "tls_version": "TLS 1.3"
        }
      },
      "http_info": {
        "status_code": 200,
        "title": "Example Domain",
        "server": "ECS (sec/96EE)",
        "content_length": 1256
      },
      "cdn_detection": {
        "provider": "CloudFlare",
        "confidence": "high"
      }
    }
  ],
  "errors": [
    {
      "target": "invalid-domain.xyz",
      "error": "DNS resolution failed: NXDOMAIN",
      "timestamp": "2024-07-22T10:01:20Z"
    }
  ]
}
```

#### CSV Format
```csv
Target,IP,Timestamp,Open Ports,Services,HTTP Status,Server,CDN,TLS Version,Scan Time,Error
example.com,93.184.216.34,2024-07-22T10:00:15Z,80;443,80:http;443:https,200,nginx/1.18.0,CloudFlare,TLS 1.3,2.5s,
google.com,142.250.125.101,2024-07-22T10:00:45Z,80;443,80:http;443:https,200,gws,Google,TLS 1.3,3.1s,
invalid-domain.xyz,,,,,,,,,,"DNS resolution failed: NXDOMAIN"
```

---

## 💡 CONTOH KASUS PENGGUNAAN

### Case 1: Audit Website Pribadi

#### Scenario
Anda memiliki website pribadi dan ingin mengecek port terbuka dan konfigurasi DNS.

#### Setup Target File
```
# my-audit.txt
mywebsite.com
www.mywebsite.com
api.mywebsite.com
mail.mywebsite.com
```

#### Command
```cmd
veko-grid.exe scan --input my-audit.txt --output website-audit.json --debug
```

#### Analysis
- Cek port 80, 443 terbuka
- Verifikasi DNS records
- Pastikan tidak ada port tidak diinginkan terbuka

### Case 2: Research DNS Infrastructure

#### Scenario
Penelitian akademik tentang DNS servers publik yang populer.

#### Setup Target File
```
# dns-research.txt
# Major DNS Providers Study
8.8.8.8          # Google Primary
8.8.4.4          # Google Secondary
1.1.1.1          # Cloudflare Primary
1.0.0.1          # Cloudflare Secondary
208.67.222.222   # OpenDNS
9.9.9.9          # Quad9
```

#### Command
```cmd
veko-grid.exe scan --input dns-research.txt --output dns-study.json --delay 1000-2000
```

### Case 3: Bug Bounty Reconnaissance

#### Scenario
Legal reconnaissance untuk program bug bounty dengan scope yang jelas.

#### Setup Target File
```
# bounty-targets.txt
# Targets sesuai scope bug bounty program
target.com
api.target.com
mobile.target.com
admin.target.com
```

#### Command (Stealth Mode)
```cmd
veko-grid.exe scan --input bounty-targets.txt --tor --delay 2000-5000 --timeout 15 --output bounty-recon.json --silent
```

### Case 4: Network Infrastructure Mapping

#### Scenario
Mapping infrastruktur jaringan perusahaan (dengan izin IT).

#### Setup Target File
```
# corporate-audit.txt
# Internal infrastructure audit (authorized)
192.168.1.1      # Gateway
192.168.1.10     # DNS Server
192.168.1.20     # Mail Server
192.168.1.100    # Web Server
mail.company.com
www.company.com
```

#### Command
```cmd
veko-grid.exe scan --input corporate-audit.txt --threads 5 --timeout 10 --output infrastructure-map.json
```

---

## 🔧 TROUBLESHOOTING

### Problem 1: Tool Tidak Bisa Dijalankan

#### Symptoms
```
'veko-grid.exe' is not recognized as an internal or external command
```

#### Solutions
```cmd
# Pastikan di folder yang benar
dir veko-grid.exe

# Jalankan dengan full path
C:\VekoGrid\veko-grid.exe --help

# Atau add ke PATH
set PATH=%PATH%;C:\VekoGrid
```

### Problem 2: Permission Denied

#### Symptoms
```
Access denied or permission error
```

#### Solutions
```cmd
# Run as Administrator
# Klik kanan Command Prompt -> Run as Administrator

# Atau ubah folder permission
# Properties -> Security -> Edit
```

### Problem 3: TOR Connection Failed

#### Symptoms
```
[ERROR] TOR proxy connection failed
```

#### Solutions
1. **Install TOR Browser**
   - Download dari https://www.torproject.org/download/
   - Jalankan TOR Browser
   - TOR daemon otomatis berjalan di port 9050

2. **Standalone TOR Daemon**
   ```cmd
   # Download Tor Expert Bundle
   # Extract dan jalankan tor.exe
   ```

3. **Test TOR Connection**
   ```cmd
   # Test dengan curl (jika ada)
   curl --socks5 127.0.0.1:9050 https://check.torproject.org
   ```

### Problem 4: DNS Resolution Failed

#### Symptoms
```
DNS resolution failed: timeout
```

#### Solutions
```cmd
# Gunakan DNS over HTTPS
veko-grid.exe scan --input targets.txt --dns doh --output results.json

# Increase timeout
veko-grid.exe scan --input targets.txt --timeout 15 --output results.json

# Test DNS manual
nslookup example.com 8.8.8.8
```

### Problem 5: High CPU Usage

#### Symptoms
- Tool menggunakan CPU tinggi
- Sistem menjadi lambat

#### Solutions
```cmd
# Reduce threads
veko-grid.exe scan --input targets.txt --threads 4 --output results.json

# Increase delays
veko-grid.exe scan --input targets.txt --delay 500-2000 --output results.json

# Monitor usage
# Task Manager -> Performance -> CPU
```

### Problem 6: Output File Error

#### Symptoms
```
[ERROR] Cannot write to output file
```

#### Solutions
```cmd
# Check folder permission
# Pastikan folder bisa ditulis

# Use different location
veko-grid.exe scan --input targets.txt --output C:\Temp\results.json

# Check disk space
dir C:\
```

---

## 💎 TIPS DAN BEST PRACTICES

### Performance Optimization

#### 1. Threading Strategy
```cmd
# Small target list (< 50)
--threads 5

# Medium target list (50-200)
--threads 8

# Large target list (200+)
--threads 15 --delay 200-800
```

#### 2. Delay Configuration
```cmd
# Fast scanning (own servers)
--delay 50-200

# Normal scanning (authorized targets)
--delay 200-800

# Stealth scanning (public research)
--delay 1000-3000
```

#### 3. Timeout Settings
```cmd
# Local network
--timeout 5

# Internet targets
--timeout 10

# Slow/filtered targets
--timeout 20
```

### Security and Stealth

#### 1. Anonymity Layers
```cmd
# Level 1: Basic delay
--delay 1000-3000

# Level 2: TOR
--tor --delay 2000-5000

# Level 3: TOR + Proxy + Long delays
--tor --proxy socks5://127.0.0.1:9050 --delay 3000-8000
```

#### 2. Traffic Distribution
```cmd
# Don't scan everything at once
# Split large target lists into smaller files

# targets-1.txt (50 hosts)
# targets-2.txt (50 hosts)
# targets-3.txt (50 hosts)
```

#### 3. Time Distribution
```cmd
# Schedule scans at different times
# Use Windows Task Scheduler

# Morning batch
schtasks /create /tn "VekoGrid-Morning" /tr "C:\VekoGrid\veko-grid.exe scan --input morning-targets.txt --output morning-results.json"

# Evening batch  
schtasks /create /tn "VekoGrid-Evening" /tr "C:\VekoGrid\veko-grid.exe scan --input evening-targets.txt --output evening-results.json"
```

### Data Management

#### 1. Organize Output Files
```
C:\VekoGrid\
├── veko-grid.exe
├── targets\
│   ├── personal-sites.txt
│   ├── dns-servers.txt
│   └── research-domains.txt
├── results\
│   ├── 2024-07-22-personal-sites.json
│   ├── 2024-07-22-dns-servers.json
│   └── 2024-07-22-research-domains.json
└── logs\
    ├── debug-2024-07-22.log
    └── errors-2024-07-22.log
```

#### 2. Backup Strategy
```cmd
# Daily backup
xcopy C:\VekoGrid\results C:\Backup\VekoGrid\%DATE% /s /i

# Compress old results
# Use 7-Zip or WinRAR untuk compress hasil lama
```

#### 3. Log Management
```cmd
# Redirect logs to file
veko-grid.exe scan --input targets.txt --debug --output results.json > scan-log.txt 2>&1

# Silent mode with JSON output
veko-grid.exe scan --input targets.txt --silent --json > results.json
```

### Legal and Ethical

#### 1. Documentation
```
# Selalu dokumentasikan:
- Target yang di-scan dan alasan
- Otorisasi yang dimiliki
- Hasil yang ditemukan
- Tindakan yang diambil
```

#### 2. Rate Limiting
```cmd
# Jangan pernah scan terlalu agresif
# Minimum delay 100ms antara requests
--delay 100-500

# Untuk targets publik, gunakan delay lebih besar
--delay 1000-3000
```

#### 3. Scope Limitation
```
# Hanya scan yang diizinkan:
- Domain/IP milik sendiri
- Target dalam scope bug bounty
- Infrastruktur dengan otorisasi tertulis
- DNS servers publik untuk research
```

---

## ⚖️ LEGAL DAN ETIKA

### ATURAN PENGGUNAAN YANG BERTANGGUNG JAWAB

#### ✅ YANG DIIZINKAN:

1. **Audit Infrastruktur Sendiri**
   - Website, server, dan domain milik pribadi
   - Infrastruktur perusahaan dengan otorisasi IT
   - Testing setelah deployment aplikasi

2. **Penelitian Akademik**
   - Study tentang DNS infrastructure
   - Research tentang CDN distribution
   - Analysis port scanning patterns

3. **Bug Bounty Program**
   - Target yang listed dalam scope program
   - Sesuai dengan rules of engagement
   - Tidak melanggar batas yang ditetapkan

4. **Authorized Penetration Testing**
   - Dengan kontrak dan scope jelas
   - Letter of authorization dari client
   - Following industry standards

#### ❌ YANG DILARANG:

1. **Scanning Tanpa Izin**
   - Government websites tanpa otorisasi
   - Corporate networks tanpa permission
   - Personal websites orang lain

2. **Aktivitas Berbahaya**
   - DoS attacks dengan traffic tinggi
   - Bypassing security measures
   - Data exfiltration

3. **Commercial Misuse**
   - Competitive intelligence tanpa etika
   - Harvesting data untuk dijual
   - Spamming atau bulk operations

### PANDUAN LEGAL COMPLIANCE

#### 1. Dokumentasi Wajib
```
Sebelum scanning, dokumentasikan:
□ Target yang akan di-scan
□ Tujuan dan alasan scanning
□ Otorisasi yang dimiliki
□ Scope dan batasan
□ Timeline pelaksanaan
```

#### 2. Written Authorization
```
Untuk corporate/client scanning:
□ Signed contract atau work order
□ Explicit scope definition
□ Contact person untuk escalation
□ Incident response procedure
□ Data handling agreement
```

#### 3. Responsible Disclosure
```
Jika menemukan vulnerability:
□ Don't exploit atau access data
□ Contact security team/owner
□ Provide clear reproduction steps
□ Allow reasonable time untuk fix
□ Follow coordinated disclosure timeline
```

### COMPLIANCE CHECKLIST

#### Sebelum Memulai Scan
- [ ] Pastikan memiliki otorisasi legal
- [ ] Baca dan pahami ToS target
- [ ] Setup proper logging
- [ ] Configure appropriate delays
- [ ] Prepare incident response plan

#### Selama Scanning
- [ ] Monitor traffic untuk tidak overload
- [ ] Stop jika ada indikasi blocking
- [ ] Respect robots.txt dan security.txt
- [ ] Don't attempt to bypass protections
- [ ] Log semua aktivitas

#### Setelah Scanning
- [ ] Secure hasil scanning data
- [ ] Report findings sesuai procedure
- [ ] Delete sensitive data jika tidak diperlukan
- [ ] Update documentation
- [ ] Follow up dengan stakeholders

---

## 📞 SUPPORT DAN BANTUAN

### Jika Mengalami Masalah

1. **Check Troubleshooting Section**
   - Baca bagian troubleshooting di atas
   - Cari error message yang sama
   - Ikuti langkah-langkah solusi

2. **Gathering Information**
   ```cmd
   # Jalankan dengan debug mode
   veko-grid.exe scan --input targets.txt --debug --output results.json > debug.log 2>&1
   
   # Check system info
   systeminfo > system-info.txt
   ```

3. **Common Log Locations**
   ```
   C:\VekoGrid\debug.log        # Debug output
   C:\VekoGrid\error.log        # Error messages
   C:\Windows\System32\drivers\etc\hosts  # DNS resolution issues
   ```

### Best Practices untuk Support

- Sertakan versi tool dan sistem operasi
- Attach log file dengan error message
- Jelaskan step yang dilakukan sebelum error
- Sertakan contoh target file yang bermasalah
- Screenshot error message jika perlu

---

## 📝 KESIMPULAN

Veko Grid adalah tool yang powerful dan fleksibel untuk network exploration dan security research. Dengan mengikuti panduan ini, Anda dapat:

✅ Menggunakan tool dengan aman dan legal
✅ Melakukan scanning yang efektif dan stealth
✅ Menghasilkan laporan yang comprehensive
✅ Menghindari masalah teknis dan legal
✅ Berkontribusi pada security research community

**Ingat selalu: With great power comes great responsibility**

Tool ini dibuat untuk membantu security professionals, researchers, dan ethical hackers dalam melakukan pekerjaan mereka dengan cara yang legal dan bertanggung jawab.

---

**Veko Grid v1.0.0** - Dibuat dengan ❤️ untuk komunitas security research

*Last updated: 22 Juli 2024*