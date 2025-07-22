# Veko Grid - Network Exploration & Stealth Tool

## Overview

Veko Grid is a Go-based CLI tool designed for anonymous network exploration and stealth scanning. The tool focuses on legal network reconnaissance, academic research, and security auditing with high anonymity features including TOR support, proxy rotation, and grid-style network mapping.

## User Preferences

Preferred communication style: Simple, everyday language.
Build target: Windows executable (.exe) for deployment

## System Architecture

### Architecture Pattern
- **Modular CLI Application**: Built using a clean separation of concerns with dedicated modules for different functionalities
- **Command-Line Interface**: Pure terminal-based tool using CLI frameworks (cobra or urfave/cli)
- **Concurrent Processing**: Multi-threaded scanning with rate limiting and intelligent delay mechanisms
- **Plugin-based Anonymity**: Modular proxy and anonymization layers

### Core Design Principles
1. **Anonymity First**: Every network request goes through anonymization layers
2. **Stealth Operations**: Random delays, header spoofing, and fingerprint randomization
3. **Legal Compliance**: Built-in safeguards for ethical usage only
4. **Modularity**: Clean separation between scanning, anonymization, and output modules

## Key Components

### 1. CLI Interface (`/cmd`)
- **Entry Point**: Main command processor and argument parsing
- **Command Structure**: Supports multiple subcommands (scan, etc.)
- **Flag Management**: Handles --tor, --proxy, --delay, --timeout, --dns, --output flags
- **Input Validation**: Validates target files and command arguments

### 2. Core Scanner (`/core`)
- **Grid Scanner**: Distributes scanning across multiple targets with visualization
- **Network Mapper**: Port scanning, ping, traceroute functionality
- **DNS Resolver**: Comprehensive DNS record lookup (A/AAAA/MX/NS/CNAME/TXT)
- **CDN Detection**: Identifies and analyzes CDN infrastructure
- **Rate Limiting**: Intelligent request throttling with random delays

### 3. Anonymization Layer (`/proxy`)
- **TOR Integration**: Routes traffic through TOR network when available
- **Proxy Support**: HTTP/SOCKS5 proxy with automatic rotation
- **Header Spoofing**: User-Agent and HTTP header randomization
- **TLS Fingerprinting**: Randomizes TLS fingerprints for stealth
- **DNS over HTTPS**: Optional DoH support for DNS privacy

### 4. Utilities (`/utils`)
- **Output Handlers**: JSON, CSV, and HTML report generation
- **Logging System**: Comprehensive logging with debug, silent, and verbose modes
- **File I/O**: Target file parsing and result file management
- **Progress Visualization**: ASCII-based progress display

## Data Flow

### 1. Input Processing
```
Target File (.txt) → Input Parser → Target Validation → Queue Generation
```

### 2. Anonymization Pipeline
```
Request → Proxy/TOR Selection → Header Spoofing → TLS Fingerprint → Network Request
```

### 3. Scanning Flow
```
Target Queue → Concurrent Workers → Rate Limiter → Anonymized Requests → Result Aggregation
```

### 4. Output Generation
```
Raw Results → Data Processing → Format Selection (JSON/CSV/HTML) → File Export
```

## External Dependencies

### Core Dependencies
- **CLI Framework**: cobra or urfave/cli for command-line interface
- **Network Libraries**: Standard Go net package with custom extensions
- **TOR Integration**: go-socks5 or similar for SOCKS5 proxy support
- **TLS Libraries**: crypto/tls with potential uTLS integration for fingerprint randomization
- **DNS Libraries**: miekg/dns for advanced DNS operations

### Optional Dependencies
- **DoH Support**: DNS over HTTPS client libraries
- **Proxy Libraries**: HTTP/SOCKS5 client implementations
- **Output Formatting**: JSON/CSV processing libraries

### System Requirements
- Go 1.21 or newer
- Optional: TOR daemon for anonymization
- Network access for target scanning

## Deployment Strategy

### Local Development
- **Direct Build**: Standard go build for immediate local execution
- **No Web Dependencies**: Pure CLI tool with no web server components
- **Cross-Platform**: Builds for Windows, Linux, and macOS

### Distribution Methods
- **Binary Releases**: Pre-compiled binaries for major platforms
- **Source Distribution**: Go source code for custom builds
- **Package Managers**: Potential future support for brew, apt, etc.

### Configuration Management
- **File-based Targets**: Plain text target files
- **Command-line Flags**: All configuration via CLI arguments
- **Environment Variables**: Optional support for proxy/TOR configuration

### Security Considerations
- **No Persistent Storage**: Results are exported but not cached
- **Anonymization Verification**: Built-in checks for proxy/TOR functionality
- **Rate Limiting**: Prevents accidental DoS-style scanning
- **Legal Safeguards**: Warnings and usage guidelines built into the tool

### Usage Examples
```bash
# Basic scanning with TOR
veko-grid scan --input targets.txt --tor --output results.json

# Advanced scanning with proxy rotation and custom delays
veko-grid scan --input targets.txt --proxy socks5://127.0.0.1:9050 \
--delay 100-500 --timeout 5 --dns doh --output results.json

# Silent mode with CSV output
veko-grid scan --input targets.txt --silent --output results.csv
```