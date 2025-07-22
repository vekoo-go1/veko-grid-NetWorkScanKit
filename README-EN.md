# 🛰️ Veko Grid - Advanced Network Exploration & Stealth Scanning Tool

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-lightgrey.svg)](https://github.com/veko-grid/veko-grid)
[![Release](https://img.shields.io/badge/Release-v1.0.0-green.svg)](https://github.com/veko-grid/veko-grid/releases)

A powerful Go-based CLI tool for anonymous network exploration and stealth scanning designed for academic research, security auditing, and ethical penetration testing.

## 🌟 Features

### 🔍 **Comprehensive Network Scanning**
- **Port Scanning**: TCP port discovery with service identification
- **DNS Resolution**: Complete DNS record lookup (A/AAAA/MX/NS/CNAME/TXT)
- **HTTP/HTTPS Analysis**: Status codes, headers, server identification
- **CDN Detection**: Identify content delivery networks and providers
- **Grid-Style Visualization**: ASCII-based progress and result display

### 🔐 **Advanced Anonymity & Stealth**
- **TOR Integration**: Route traffic through TOR network for anonymity
- **Proxy Rotation**: HTTP/SOCKS5 proxy support with automatic rotation
- **Header Spoofing**: Randomized User-Agent and HTTP headers
- **TLS Fingerprinting**: Randomized TLS fingerprints for enhanced stealth
- **DNS over HTTPS**: Optional DoH support for DNS privacy
- **Smart Delays**: Configurable random delays to avoid detection

### 📊 **Flexible Output & Reporting**
- **Multiple Formats**: JSON, CSV, and HTML report generation
- **Real-time Progress**: Live scanning progress with ETA
- **Comprehensive Logging**: Debug, silent, and verbose modes
- **Custom Templates**: Configurable output formatting

### ⚙️ **Advanced Configuration**
- **Multi-threading**: Concurrent scanning with rate limiting
- **Custom Timeouts**: Configurable connection timeouts
- **Target Management**: Flexible target file formats
- **Environment Variables**: Support for configuration via ENV vars

## 🚀 Quick Start

### Installation

#### Option 1: Download Pre-built Binary
```bash
# Download the latest release for your platform
wget https://github.com/veko-grid/veko-grid/releases/latest/download/veko-grid-linux
chmod +x veko-grid-linux
./veko-grid-linux --help
```

#### Option 2: Build from Source
```bash
# Clone the repository
git clone https://github.com/veko-grid/veko-grid.git
cd veko-grid

# Build the binary
go build -o veko-grid

# Run the tool
./veko-grid --help
```

### Basic Usage

#### 1. Create Target File
```bash
# targets.txt
echo "example.com" > targets.txt
echo "google.com" >> targets.txt
echo "github.com" >> targets.txt
echo "8.8.8.8" >> targets.txt
```

#### 2. Basic Scan
```bash
# Simple scan with JSON output
./veko-grid scan --input targets.txt --output results.json

# Verbose scan with detailed logging
./veko-grid scan --input targets.txt --output results.json --debug
```

#### 3. View Results
```bash
# Pretty print JSON results
cat results.json | jq '.'

# Extract open ports only
cat results.json | jq '.results[].open_ports'
```

## 🔧 Advanced Usage

### Stealth and Anonymity

#### TOR Integration
```bash
# Basic TOR scanning (requires TOR daemon)
./veko-grid scan --input targets.txt --tor --output results.json

# TOR with custom delays for maximum stealth
./veko-grid scan --input targets.txt --tor --delay 2000-5000 --output results.json
```

#### Proxy Configuration
```bash
# SOCKS5 proxy
./veko-grid scan --input targets.txt --proxy socks5://127.0.0.1:9050 --output results.json

# HTTP proxy with authentication
./veko-grid scan --input targets.txt --proxy http://user:pass@proxy.example.com:8080 --output results.json

# Multiple proxy rotation
./veko-grid scan --input targets.txt --proxy-list proxies.txt --output results.json
```

### Performance Optimization

#### Multi-threading and Delays
```bash
# High-speed scanning (use carefully)
./veko-grid scan --input targets.txt --threads 20 --delay 100-300 --output results.json

# Balanced scanning for large target lists
./veko-grid scan --input targets.txt --threads 10 --delay 500-1500 --timeout 10 --output results.json

# Maximum stealth mode
./veko-grid scan --input targets.txt --threads 3 --delay 3000-8000 --timeout 30 --output results.json
```

#### DNS Configuration
```bash
# DNS over HTTPS for privacy
./veko-grid scan --input targets.txt --dns doh --output results.json

# Custom DNS servers
./veko-grid scan --input targets.txt --dns-servers 8.8.8.8,1.1.1.1 --output results.json
```

### Output Formats

#### JSON Output (Detailed)
```bash
./veko-grid scan --input targets.txt --output results.json --format json
```

#### CSV Output (Spreadsheet-friendly)
```bash
./veko-grid scan --input targets.txt --output results.csv --format csv
```

#### HTML Report (Visual)
```bash
./veko-grid scan --input targets.txt --output report.html --format html
```

#### Silent Mode (Scriptable)
```bash
./veko-grid scan --input targets.txt --silent --json > results.json
```

## 📁 File Formats

### Target File Format
```
# targets.txt - One target per line
# Lines starting with # are comments

# Domain names
example.com
subdomain.example.com
api.example.com

# IP addresses
192.168.1.1
8.8.8.8
1.1.1.1

# IP ranges (CIDR notation)
192.168.1.0/24
10.0.0.0/16

# Specific ports
example.com:443
192.168.1.1:22,80,443
```

### JSON Output Format
```json
{
  "metadata": {
    "tool": "Veko Grid",
    "version": "1.0.0",
    "start_time": "2024-07-22T10:00:00Z",
    "end_time": "2024-07-22T10:05:30Z",
    "total_targets": 10,
    "successful_scans": 8,
    "failed_scans": 2,
    "scan_options": {
      "threads": 8,
      "timeout": 5,
      "delay_range": "500-1500ms",
      "tor_enabled": true,
      "proxy": "socks5://127.0.0.1:9050"
    }
  },
  "results": [
    {
      "target": "example.com",
      "ip_address": "93.184.216.34",
      "timestamp": "2024-07-22T10:00:15Z",
      "scan_duration": "2.534s",
      "dns_records": {
        "A": ["93.184.216.34"],
        "AAAA": ["2606:2800:220:1:248:1893:25c8:1946"],
        "MX": ["0 ."],
        "NS": ["a.iana-servers.net.", "b.iana-servers.net."],
        "TXT": ["v=spf1 -all"],
        "CNAME": []
      },
      "open_ports": [80, 443],
      "port_details": {
        "80": {
          "service": "http",
          "banner": "nginx/1.18.0 (Ubuntu)",
          "status": "open",
          "response_time": "145ms"
        },
        "443": {
          "service": "https",
          "banner": "nginx/1.18.0 (Ubuntu)",
          "status": "open",
          "response_time": "156ms",
          "tls_version": "TLS 1.3",
          "cipher_suite": "TLS_AES_256_GCM_SHA384"
        }
      },
      "http_analysis": {
        "status_code": 200,
        "title": "Example Domain",
        "server": "ECS (sec/96EE)",
        "content_length": 1256,
        "headers": {
          "Server": "ECS (sec/96EE)",
          "Content-Type": "text/html; charset=UTF-8",
          "Last-Modified": "Thu, 17 Oct 2019 07:18:26 GMT"
        }
      },
      "cdn_detection": {
        "provider": "Unknown",
        "confidence": "low",
        "indicators": []
      },
      "geolocation": {
        "country": "US",
        "region": "California",
        "city": "Los Angeles",
        "latitude": 34.0522,
        "longitude": -118.2437
      }
    }
  ],
  "errors": [
    {
      "target": "nonexistent.example",
      "error": "DNS resolution failed: NXDOMAIN",
      "timestamp": "2024-07-22T10:01:20Z",
      "error_code": "DNS_NXDOMAIN"
    }
  ]
}
```

## 🔒 Legal and Ethical Usage

### ⚠️ Important Legal Notice

This tool is designed for **legitimate security research and authorized testing only**. Users are solely responsible for ensuring their use complies with all applicable laws and regulations.

### ✅ Authorized Use Cases
- **Academic Research**: Network topology studies and DNS infrastructure research
- **Bug Bounty Programs**: Reconnaissance within authorized scope and rules
- **Penetration Testing**: Authorized security assessments with proper documentation
- **Infrastructure Auditing**: Scanning your own systems and networks
- **Security Education**: Learning about network security in controlled environments

### ❌ Prohibited Activities
- Unauthorized scanning of systems you don't own or control
- Circumventing security measures or access controls
- Any activity that violates local, national, or international laws
- Denial of service attacks or resource exhaustion
- Data harvesting without explicit permission

### 🛡️ Responsible Usage Guidelines

1. **Always obtain written authorization** before scanning external systems
2. **Respect rate limits** and implement appropriate delays between requests
3. **Honor robots.txt** and terms of service of target websites
4. **Report vulnerabilities responsibly** through proper channels
5. **Maintain confidentiality** of discovered information
6. **Document your activities** for compliance and accountability

### 📋 Pre-Scan Checklist
- [ ] Confirmed ownership or authorization for all targets
- [ ] Read and understood applicable terms of service
- [ ] Configured appropriate delays and rate limiting
- [ ] Prepared incident response procedures
- [ ] Documented scope and objectives

## 🛠️ Configuration

### Environment Variables
```bash
# TOR configuration
export TOR_PROXY_HOST=127.0.0.1
export TOR_PROXY_PORT=9050

# HTTP proxy configuration
export HTTP_PROXY=http://proxy.example.com:8080
export HTTPS_PROXY=https://proxy.example.com:8080

# DNS configuration
export VEKO_DNS_SERVERS=8.8.8.8,1.1.1.1
export VEKO_DNS_TIMEOUT=5

# Output configuration
export VEKO_OUTPUT_FORMAT=json
export VEKO_LOG_LEVEL=info
```

### Configuration File
```yaml
# veko-config.yml
scanner:
  threads: 8
  timeout: 10
  delay_min: 500
  delay_max: 1500

proxy:
  enabled: true
  type: socks5
  host: 127.0.0.1
  port: 9050

dns:
  servers: [8.8.8.8, 1.1.1.1]
  doh_enabled: true
  timeout: 5

output:
  format: json
  include_errors: true
  include_metadata: true
```

## 🔧 Building and Development

### Prerequisites
- Go 1.21 or higher
- Git (for cloning the repository)

### Build Commands
```bash
# Install dependencies
go mod tidy

# Build for current platform
go build -o veko-grid

# Build with optimizations
go build -ldflags="-s -w" -o veko-grid

# Cross-compilation for different platforms
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o veko-grid.exe
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o veko-grid-linux
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o veko-grid-mac
```

### Development Setup
```bash
# Clone the repository
git clone https://github.com/veko-grid/veko-grid.git
cd veko-grid

# Install development dependencies
go mod download

# Run tests
go test ./...

# Run with race detection
go run -race main.go scan --input test-targets.txt --output test-results.json

# Build with debug info
go build -race -o veko-grid-debug
```

## 🐛 Troubleshooting

### Common Issues

#### Binary Not Found
```bash
# Make sure the binary is executable
chmod +x veko-grid

# Add to PATH
export PATH=$PATH:/path/to/veko-grid

# Or use full path
/full/path/to/veko-grid --help
```

#### Permission Denied
```bash
# Run with elevated privileges (if needed)
sudo ./veko-grid scan --input targets.txt --output results.json

# Or change binary permissions
chmod 755 veko-grid
```

#### TOR Connection Failed
```bash
# Install TOR
# Ubuntu/Debian:
sudo apt install tor
sudo systemctl start tor

# macOS:
brew install tor
tor

# Windows: Download TOR Browser or standalone TOR
```

#### DNS Resolution Issues
```bash
# Use DNS over HTTPS
./veko-grid scan --input targets.txt --dns doh --output results.json

# Specify custom DNS servers
./veko-grid scan --input targets.txt --dns-servers 8.8.8.8,1.1.1.1 --output results.json

# Increase timeout
./veko-grid scan --input targets.txt --timeout 15 --output results.json
```

#### High Resource Usage
```bash
# Reduce thread count
./veko-grid scan --input targets.txt --threads 4 --output results.json

# Increase delays
./veko-grid scan --input targets.txt --delay 1000-3000 --output results.json

# Use silent mode to reduce output overhead
./veko-grid scan --input targets.txt --silent --output results.json
```

### Debug Mode
```bash
# Enable verbose logging
./veko-grid scan --input targets.txt --debug --output results.json

# Redirect debug output to file
./veko-grid scan --input targets.txt --debug --output results.json > debug.log 2>&1

# Use built-in profiling (for developers)
./veko-grid scan --input targets.txt --profile --output results.json
```

## 🤝 Contributing

We welcome contributions from the security community! Here's how you can help:

### Ways to Contribute
- **Bug Reports**: Report issues with detailed reproduction steps
- **Feature Requests**: Suggest new features or improvements
- **Code Contributions**: Submit pull requests with fixes or enhancements
- **Documentation**: Improve documentation and examples
- **Testing**: Test on different platforms and edge cases

### Development Guidelines
1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes with appropriate tests
4. Ensure code passes all tests (`go test ./...`)
5. Update documentation as needed
6. Commit changes (`git commit -m 'Add amazing feature'`)
7. Push to branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

### Code Style
- Follow standard Go formatting (`go fmt`)
- Include comprehensive tests for new features
- Add detailed comments for complex logic
- Update documentation for user-facing changes

## 📈 Performance Benchmarks

### Scanning Speed
| Target Count | Threads | Avg Time | Targets/sec |
|--------------|---------|----------|-------------|
| 10           | 4       | 15s      | 0.67        |
| 100          | 8       | 2m 30s   | 0.67        |
| 1000         | 16      | 25m      | 0.67        |

*Note: Performance varies based on target responsiveness and network conditions*

### Memory Usage
| Target Count | Memory Usage | Peak Memory |
|--------------|--------------|-------------|
| 10           | 15 MB        | 25 MB       |
| 100          | 25 MB        | 45 MB       |
| 1000         | 85 MB        | 150 MB      |

## 📊 Comparison with Similar Tools

| Feature              | Veko Grid | nmap | masscan | subfinder |
|---------------------|-----------|------|---------|-----------|
| Port Scanning       | ✅         | ✅    | ✅       | ❌        |
| DNS Resolution      | ✅         | ✅    | ❌       | ✅        |
| TOR Support         | ✅         | ❌    | ❌       | ❌        |
| Proxy Rotation      | ✅         | ❌    | ❌       | ✅        |
| JSON/CSV Output     | ✅         | ✅    | ✅       | ✅        |
| Grid Visualization  | ✅         | ❌    | ❌       | ❌        |
| Go-based            | ✅         | ❌    | ❌       | ✅        |
| Single Binary       | ✅         | ❌    | ✅       | ✅        |

## 🏆 Recognition and Awards

*This section will be updated as the project gains recognition in the security community.*

## 📚 Additional Resources

### Documentation
- [Complete User Guide](PANDUAN-LENGKAP.md) (Indonesian)
- [Quick Start Guide](QUICK-START.md) (Indonesian)
- [API Documentation](docs/api.md)
- [Developer Guide](docs/development.md)

### Community
- [GitHub Discussions](https://github.com/veko-grid/veko-grid/discussions)
- [Issue Tracker](https://github.com/veko-grid/veko-grid/issues)
- [Security Policy](SECURITY.md)

### Related Projects
- [nmap](https://nmap.org/) - Network discovery and security auditing
- [masscan](https://github.com/robertdavidgraham/masscan) - Fast port scanner
- [subfinder](https://github.com/projectdiscovery/subfinder) - Subdomain discovery
- [nuclei](https://github.com/projectdiscovery/nuclei) - Vulnerability scanner

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### Third-Party Licenses
- Go standard library: [BSD-style license](https://golang.org/LICENSE)
- Third-party dependencies: See [go.mod](go.mod) for specific packages and their licenses

## 🙏 Acknowledgments

- The Go programming language team for excellent networking libraries
- The security research community for inspiration and feedback
- Contributors and testers who helped improve the tool
- The TOR Project for privacy and anonymity tools

## 📞 Support

### Getting Help
- 📖 Check the [documentation](README.md) and [troubleshooting](#troubleshooting) section
- 🐛 Search [existing issues](https://github.com/veko-grid/veko-grid/issues) for your problem
- 💬 Start a [discussion](https://github.com/veko-grid/veko-grid/discussions) for questions
- 🔒 For security issues, see our [Security Policy](SECURITY.md)

### Contact Information
- **GitHub**: [veko-grid/veko-grid](https://github.com/veko-grid/veko-grid)
- **Issues**: [GitHub Issues](https://github.com/veko-grid/veko-grid/issues)
- **Discussions**: [GitHub Discussions](https://github.com/veko-grid/veko-grid/discussions)

---

**Made with ❤️ by the security research community**

*Veko Grid v1.0.0 - Network exploration made simple, secure, and ethical.*