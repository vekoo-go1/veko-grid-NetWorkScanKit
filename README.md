# 🛰️ Veko Grid - Network Exploration & Stealth Tool

**Veko Grid** adalah tool CLI berbasis Go untuk eksplorasi jaringan anonim dan stealth scanning dengan support TOR, proxy rotation, dan grid-style network mapping.

## 🎯 Tujuan Utama

Tool ini dibuat khusus untuk:
- 📚 **Penelitian akademik**
- 🔒 **Audit keamanan jaringan sendiri**
- 🌐 **Eksplorasi infrastruktur domain publik secara legal dan anonim**

> ⚠️ **PENTING**: Tool ini BUKAN untuk hacking atau serangan. Gunakan hanya untuk tujuan legal dan etis.

## 🧩 Fitur Utama

### 🕸️ Network Mapping
- Pemetaan jaringan eksternal (port, ping, traceroute)
- Scan domain atau IP publik dengan batas aman
- DNS Resolver lengkap (A/AAAA/MX/NS/CNAME/TXT)
- CDN detection dan analysis

### 🧊 Anonimitas Tinggi
- Support koneksi melalui TOR
- Proxy HTTP/SOCKS5 dengan rotasi otomatis
- Header spoofing (User-Agent, Accept, dll)
- TLS Fingerprint randomization

### 🧭 Intelligent Grid Scan
- Grid-style scanning dengan visualisasi ASCII
- Input dari file `.txt` berisi domain/IP
- Real-time progress monitoring
- Concurrent scanning dengan rate limiting

### 🔐 Stealth & Penyamaran
- Random delay antar permintaan (anti-radar)
- DNS-over-HTTPS (DoH) optional
- Multiple fingerprint modes
- Proxy rotation otomatis

### 📦 Output & Logging
- Export hasil ke JSON atau CSV
- Mode `--silent`, `--json`, dan `--debug`
- HTML report generation
- Comprehensive logging system

## 🚀 Instalasi dan Build

### Prerequisites
- Go 1.19 atau lebih baru
- Git (opsional)

### Build dari Source
```bash
# Download dependencies
go mod tidy

# Build binary untuk platform saat ini
go build -o veko-grid

# Build untuk Windows (.exe)
go build -o veko-grid.exe

# Build untuk berbagai platform dengan optimasi
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o veko-grid.exe
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o veko-grid-linux
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o veko-grid-mac
```

### Build dengan Script Otomatis

**Untuk Windows:**
```bash
# Jalankan build script
build.bat
```

**Untuk Linux/Mac:**
```bash
# Berikan permission dan jalankan
chmod +x build.sh
./build.sh
```

## 📖 Penggunaan

### Command Dasar

```bash
# Tampilkan help dan informasi tool
./veko-grid --help

# Lihat opsi untuk command scan
./veko-grid scan --help
```

### Contoh Penggunaan

#### 1. Scanning Dasar
```bash
# Scan dengan file target standar
./veko-grid scan --input targets.txt --output results.json

# Scan dengan mode debug untuk melihat detail proses
./veko-grid scan --input targets.txt --output results.json --debug
```

#### 2. Scanning dengan Anonimitas
```bash
# Menggunakan TOR untuk anonimitas (memerlukan TOR daemon)
./veko-grid scan --input targets.txt --tor --output results.json

# Menggunakan SOCKS5 proxy
./veko-grid scan --input targets.txt --proxy socks5://127.0.0.1:9050 --output results.json

# Kombinasi TOR dan proxy dengan delay custom
./veko-grid scan --input targets.txt --tor --proxy socks5://127.0.0.1:9050 --delay 200-1000 --output results.json
```

#### 3. Konfigurasi Advanced
```bash
# Scan dengan DNS over HTTPS dan timeout custom
./veko-grid scan --input targets.txt --dns doh --timeout 10 --output results.json

# Output dalam format CSV
./veko-grid scan --input targets.txt --output results.csv

# Mode silent untuk scripting
./veko-grid scan --input targets.txt --silent --output results.json

# Mode JSON output ke stdout untuk pipeline
./veko-grid scan --input targets.txt --json > results.json
```

#### 4. Performance Tuning
```bash
# Gunakan lebih banyak thread untuk scanning cepat
./veko-grid scan --input targets.txt --threads 20 --output results.json

# Delay minimal untuk scanning cepat (gunakan hati-hati)
./veko-grid scan --input targets.txt --delay 50-100 --threads 20 --output results.json
```

## 📁 Format File Input

### File Target (`targets.txt`)
```
# Format: satu domain atau IP per baris
# Gunakan # untuk komentar

# Domain examples
example.com
google.com
github.com

# IP address examples  
8.8.8.8
1.1.1.1

# Subdomain examples
api.example.com
www.example.com
```

## 📊 Format Output

### JSON Output
```json
{
  "metadata": {
    "tool": "Veko Grid",
    "version": "1.0.0",
    "start_time": "2024-01-01T10:00:00Z",
    "total_hosts": 10,
    "successful": 8,
    "failed": 2
  },
  "results": [
    {
      "target": "example.com",
      "ip": "93.184.216.34",
      "timestamp": "2024-01-01T10:00:01Z",
      "dns_records": {
        "A": ["93.184.216.34"],
        "AAAA": ["2606:2800:220:1:248:1893:25c8:1946"]
      },
      "open_ports": [80, 443],
      "services": {
        "80": "http",
        "443": "https"
      },
      "scan_time": "2.5s"
    }
  ]
}
```

### CSV Output
```csv
Target,IP,Timestamp,Open Ports,Services,CDN Provider,TLS Version,Scan Time,Error
example.com,93.184.216.34,2024-01-01T10:00:01Z,80;443,80:http;443:https,,TLS 1.3,2.5s,
```

## ⚙️ Konfigurasi

### Environment Variables (Opsional)
```bash
# TOR proxy configuration
export TOR_PROXY_HOST=127.0.0.1
export TOR_PROXY_PORT=9050

# HTTP proxy configuration  
export HTTP_PROXY=http://proxy.example.com:8080
export HTTPS_PROXY=https://proxy.example.com:8080
```

### TOR Setup (Opsional)
```bash
# Ubuntu/Debian
sudo apt install tor
sudo systemctl start tor

# Windows (menggunakan Tor Browser atau standalone)
# Download dari: https://www.torproject.org/download/

# macOS  
brew install tor
tor
```

## 🔒 Legal dan Etika

### ⚠️ PENTING - PENGGUNAAN YANG BERTANGGUNG JAWAB

Tool ini dibuat untuk:
- ✅ **Penelitian akademik dan edukasi**
- ✅ **Audit keamanan infrastruktur sendiri**
- ✅ **Testing penetrasi dengan izin tertulis**
- ✅ **Bug bounty hunting sesuai scope**

**TIDAK UNTUK:**
- ❌ Scanning sistem tanpa izin
- ❌ Aktivitas ilegal atau merusak
- ❌ Bypassing security measures
- ❌ Data harvesting tanpa izin

### Panduan Legal
1. **Selalu dapatkan izin tertulis** sebelum melakukan scanning
2. **Patuhi rate limiting** untuk menghindari DoS
3. **Gunakan proxy/TOR** hanya untuk perlindungan privasi legal
4. **Laporkan vulnerability** melalui jalur yang tepat
5. **Hormati robots.txt** dan terms of service

## 🐛 Troubleshooting

### Build Issues
```bash
# Jika dependency error
rm go.sum
go mod tidy
go mod download

# Jika Go version error
go version  # pastikan >= 1.19
```

### Runtime Issues
```bash
# Permission denied
chmod +x veko-grid

# TOR connection failed
# Pastikan TOR daemon berjalan di port 9050

# DNS resolution failed  
# Coba gunakan --dns doh atau ubah DNS server sistem
```

## 🤝 Kontribusi

Kontribusi untuk perbaikan dan pengembangan tool ini sangat diterima. Pastikan untuk:

1. Fork repository
2. Buat feature branch
3. Test perubahan secara menyeluruh
4. Submit pull request dengan deskripsi jelas

## 📝 License

Tool ini dibuat untuk tujuan edukasi dan penelitian. Pengguna bertanggung jawab penuh atas penggunaan tool sesuai hukum yang berlaku.

## 🆘 Support

Jika menemukan bug atau memerlukan fitur tambahan:
1. Buat issue dengan deskripsi detail
2. Sertakan log error dan konfigurasi
3. Jelaskan step untuk reproduce issue

---

**Dibuat dengan ❤️ untuk komunitas security research dan academic studies**
