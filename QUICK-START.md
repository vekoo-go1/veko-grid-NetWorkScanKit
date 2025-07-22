# 🚀 QUICK START - VEKO GRID

## Langkah Cepat Mulai Menggunakan Veko Grid

---

## 1. ⚡ SETUP CEPAT (5 Menit)

### Download dan Siapkan
```cmd
# Buat folder kerja
mkdir C:\VekoGrid
cd C:\VekoGrid

# Copy file veko-grid.exe ke folder ini
# Ukuran file: 9.4MB
```

### Test Tool
```cmd
# Test apakah tool berjalan
veko-grid.exe --help

# Harusnya muncul:
# ╔═══════════════════════════════════════════════════════════════╗
# ║  🛰️  VEKO GRID v1.0.0 - Network Exploration & Stealth Tool   ║
# ╚═══════════════════════════════════════════════════════════════╝
```

---

## 2. 📝 BUAT FILE TARGET

### Contoh Sederhana (targets.txt)
```
# targets.txt - Daftar yang ingin di-scan
google.com
github.com
8.8.8.8
1.1.1.1
```

### Contoh untuk Website Sendiri
```
# my-sites.txt
mywebsite.com
www.mywebsite.com
mail.mywebsite.com
```

---

## 3. 🎯 COMMAND DASAR

### Scan Sederhana
```cmd
veko-grid.exe scan --input targets.txt --output results.json
```

### Scan dengan Detail Log
```cmd
veko-grid.exe scan --input targets.txt --output results.json --debug
```

### Scan Mode Silent
```cmd
veko-grid.exe scan --input targets.txt --output results.json --silent
```

---

## 4. 📊 LIHAT HASIL

### Hasil dalam JSON (results.json)
```json
{
  "metadata": {
    "tool": "Veko Grid",
    "total_targets": 4,
    "successful": 4,
    "failed": 0
  },
  "results": [
    {
      "target": "google.com",
      "ip": "142.250.125.101",
      "open_ports": [80, 443],
      "scan_time": "2.1s"
    }
  ]
}
```

### Hasil dalam CSV
```cmd
# Untuk output CSV
veko-grid.exe scan --input targets.txt --output results.csv
```

---

## 5. 🔧 COMMAND ADVANCED

### Dengan TOR (Anonymous)
```cmd
veko-grid.exe scan --input targets.txt --tor --output results.json
```
*Catatan: Butuh TOR Browser yang sedang berjalan*

### Dengan Custom Delay (Stealth)
```cmd
veko-grid.exe scan --input targets.txt --delay 1000-3000 --output results.json
```

### Multiple Threads (Cepat)
```cmd
veko-grid.exe scan --input targets.txt --threads 10 --output results.json
```

---

## 6. 📁 CONTOH FILE SIAP PAKAI

Tool sudah menyediakan contoh file:

- **contoh-targets-dns.txt** - DNS servers populer
- **contoh-targets-website.txt** - Website populer untuk research
- **targets.txt** - Contoh umum

### Cara Pakai:
```cmd
# Scan DNS servers
veko-grid.exe scan --input contoh-targets-dns.txt --output dns-results.json

# Scan websites (gunakan delay)
veko-grid.exe scan --input contoh-targets-website.txt --delay 2000-5000 --output web-results.json
```

---

## 7. ⚠️ TIPS PENTING

### Legal & Etika
- ✅ Scan domain/IP milik sendiri
- ✅ Scan dengan izin tertulis
- ❌ Jangan scan sistem orang lain tanpa izin

### Performance
```cmd
# Target sedikit (< 20)
--threads 5 --delay 500-1000

# Target banyak (50+)
--threads 10 --delay 200-800

# Maximum stealth
--tor --delay 2000-5000
```

### Troubleshooting
```cmd
# Jika ada error, gunakan debug
veko-grid.exe scan --input targets.txt --debug --output results.json

# Jika lambat, kurangi threads
--threads 3

# Jika timeout, tingkatkan timeout
--timeout 15
```

---

## 8. 🔄 WORKFLOW HARIAN

### Morning Routine
```cmd
# 1. Buat target file untuk hari ini
echo mysite.com > today-targets.txt
echo api.mysite.com >> today-targets.txt

# 2. Scan dengan stealth mode
veko-grid.exe scan --input today-targets.txt --delay 1000-2000 --output morning-scan.json --silent

# 3. Check results
type morning-scan.json | findstr "open_ports"
```

---

## 9. 📞 BANTUAN CEPAT

### Error Common
| Error | Solusi |
|-------|--------|
| `not recognized` | Pastikan di folder yang benar |
| `permission denied` | Run as Administrator |
| `DNS failed` | Gunakan `--dns doh` |
| `TOR failed` | Install TOR Browser dulu |

### Dokumentasi Lengkap
Baca **PANDUAN-LENGKAP.md** untuk tutorial detail dan advanced usage.

---

## 🎉 SELAMAT!

Anda sudah siap menggunakan Veko Grid! 

**Next Steps:**
1. Mulai dengan target kecil (2-3 domain)
2. Experiment dengan delay dan threads
3. Baca panduan lengkap untuk fitur advanced
4. Selalu patuhi etika dan aturan legal

**Happy Scanning!** 🛰️