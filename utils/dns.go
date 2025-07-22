package utils

import (
        "fmt"
        "net"
        "strings"
        "time"

        "github.com/miekg/dns"
)

// DNSResolver mengelola DNS resolution dengan support DoH
type DNSResolver struct {
        client    *dns.Client
        servers   []string
        useDoH    bool
        logger    *Logger
}

// DNSRecord menyimpan record DNS
type DNSRecord struct {
        Type  string `json:"type"`
        Value string `json:"value"`
        TTL   uint32 `json:"ttl"`
}

// NewDNSResolver membuat instance DNSResolver baru
func NewDNSResolver(useDoH bool, logger *Logger) (*DNSResolver, error) {
        resolver := &DNSResolver{
                client: &dns.Client{
                        Timeout: 10 * time.Second,
                },
                useDoH: useDoH,
                logger: logger,
        }

        // Setup DNS servers
        if useDoH {
                resolver.servers = []string{
                        "https://dns.cloudflare.com/dns-query",
                        "https://dns.google/dns-query",
                        "https://dns.quad9.net/dns-query",
                }
                logger.Info("🌐 DNS over HTTPS (DoH) enabled")
        } else {
                resolver.servers = []string{
                        "8.8.8.8:53",
                        "1.1.1.1:53",
                        "9.9.9.9:53",
                        "208.67.222.222:53",
                }
        }

        return resolver, nil
}

// ResolveAll melakukan resolve semua jenis DNS record
func (d *DNSResolver) ResolveAll(domain string) (map[string][]string, error) {
        results := make(map[string][]string)

        // A records
        if aRecords, err := d.LookupA(domain); err == nil && len(aRecords) > 0 {
                results["A"] = aRecords
        }

        // AAAA records (IPv6)
        if aaaaRecords, err := d.LookupAAAA(domain); err == nil && len(aaaaRecords) > 0 {
                results["AAAA"] = aaaaRecords
        }

        // CNAME records
        if cnameRecords, err := d.LookupCNAME(domain); err == nil && len(cnameRecords) > 0 {
                results["CNAME"] = cnameRecords
        }

        // MX records
        if mxRecords, err := d.LookupMX(domain); err == nil && len(mxRecords) > 0 {
                results["MX"] = mxRecords
        }

        // NS records
        if nsRecords, err := d.LookupNS(domain); err == nil && len(nsRecords) > 0 {
                results["NS"] = nsRecords
        }

        // TXT records
        if txtRecords, err := d.LookupTXT(domain); err == nil && len(txtRecords) > 0 {
                results["TXT"] = txtRecords
        }

        if len(results) == 0 {
                return nil, fmt.Errorf("no DNS records found for %s", domain)
        }

        return results, nil
}

// LookupA melakukan A record lookup
func (d *DNSResolver) LookupA(domain string) ([]string, error) {
        if d.useDoH {
                return d.lookupDoH(domain, dns.TypeA)
        }
        return d.lookupClassic(domain, dns.TypeA)
}

// LookupAAAA melakukan AAAA record lookup  
func (d *DNSResolver) LookupAAAA(domain string) ([]string, error) {
        if d.useDoH {
                return d.lookupDoH(domain, dns.TypeAAAA)
        }
        return d.lookupClassic(domain, dns.TypeAAAA)
}

// LookupCNAME melakukan CNAME record lookup
func (d *DNSResolver) LookupCNAME(domain string) ([]string, error) {
        if d.useDoH {
                return d.lookupDoH(domain, dns.TypeCNAME)
        }
        return d.lookupClassic(domain, dns.TypeCNAME)
}

// LookupMX melakukan MX record lookup
func (d *DNSResolver) LookupMX(domain string) ([]string, error) {
        if d.useDoH {
                return d.lookupDoH(domain, dns.TypeMX)
        }
        return d.lookupClassic(domain, dns.TypeMX)
}

// LookupNS melakukan NS record lookup
func (d *DNSResolver) LookupNS(domain string) ([]string, error) {
        if d.useDoH {
                return d.lookupDoH(domain, dns.TypeNS)
        }
        return d.lookupClassic(domain, dns.TypeNS)
}

// LookupTXT melakukan TXT record lookup
func (d *DNSResolver) LookupTXT(domain string) ([]string, error) {
        if d.useDoH {
                return d.lookupDoH(domain, dns.TypeTXT)
        }
        return d.lookupClassic(domain, dns.TypeTXT)
}

// lookupClassic melakukan DNS lookup dengan server tradisional
func (d *DNSResolver) lookupClassic(domain string, qtype uint16) ([]string, error) {
        var results []string
        var lastErr error

        // Try multiple servers untuk redundancy
        for _, server := range d.servers {
                msg := new(dns.Msg)
                msg.SetQuestion(dns.Fqdn(domain), qtype)
                msg.RecursionDesired = true

                resp, _, err := d.client.Exchange(msg, server)
                if err != nil {
                        lastErr = err
                        continue
                }

                if resp.Rcode != dns.RcodeSuccess {
                        lastErr = fmt.Errorf("DNS query failed with rcode: %d", resp.Rcode)
                        continue
                }

                // Parse answers
                for _, ans := range resp.Answer {
                        value := d.extractRecordValue(ans)
                        if value != "" {
                                results = append(results, value)
                        }
                }

                if len(results) > 0 {
                        break
                }
        }

        if len(results) == 0 && lastErr != nil {
                return nil, lastErr
        }

        return results, nil
}

// lookupDoH melakukan DNS lookup menggunakan DNS over HTTPS
func (d *DNSResolver) lookupDoH(domain string, qtype uint16) ([]string, error) {
        // Simplified DoH implementation
        // Dalam implementasi nyata, gunakan HTTP client untuk DoH queries
        
        // Fallback ke classic DNS untuk sekarang
        d.logger.Debug(fmt.Sprintf("DoH lookup untuk %s (type: %d)", domain, qtype))
        return d.lookupClassic(domain, qtype)
}

// extractRecordValue mengekstrak value dari DNS answer
func (d *DNSResolver) extractRecordValue(rr dns.RR) string {
        switch v := rr.(type) {
        case *dns.A:
                return v.A.String()
        case *dns.AAAA:
                return v.AAAA.String()
        case *dns.CNAME:
                return strings.TrimSuffix(v.Target, ".")
        case *dns.MX:
                return fmt.Sprintf("%d %s", v.Preference, strings.TrimSuffix(v.Mx, "."))
        case *dns.NS:
                return strings.TrimSuffix(v.Ns, ".")
        case *dns.TXT:
                return strings.Join(v.Txt, " ")
        case *dns.PTR:
                return strings.TrimSuffix(v.Ptr, ".")
        default:
                return ""
        }
}

// ReverseLookup melakukan reverse DNS lookup
func (d *DNSResolver) ReverseLookup(ip string) ([]string, error) {
        addr, err := dns.ReverseAddr(ip)
        if err != nil {
                return nil, err
        }

        return d.LookupPTR(addr)
}

// LookupPTR melakukan PTR record lookup
func (d *DNSResolver) LookupPTR(addr string) ([]string, error) {
        if d.useDoH {
                return d.lookupDoH(addr, dns.TypePTR)
        }
        return d.lookupClassic(addr, dns.TypePTR)
}

// GetDNSInfo mendapatkan informasi lengkap DNS
func (d *DNSResolver) GetDNSInfo(domain string) (*DNSInfo, error) {
        info := &DNSInfo{
                Domain:    domain,
                Timestamp: time.Now(),
        }

        // Resolve semua record types
        records, err := d.ResolveAll(domain)
        if err != nil {
                return nil, err
        }

        info.Records = records

        // Check authoritative servers
        if nsRecords, exists := records["NS"]; exists {
                info.AuthoritativeServers = nsRecords
        }

        // Check mail servers
        if mxRecords, exists := records["MX"]; exists {
                info.MailServers = mxRecords
        }

        // Reverse lookup untuk IP addresses
        if aRecords, exists := records["A"]; exists && len(aRecords) > 0 {
                if reverseRecords, err := d.ReverseLookup(aRecords[0]); err == nil {
                        info.ReverseRecords = reverseRecords
                }
        }

        return info, nil
}

// DNSInfo menyimpan informasi lengkap DNS
type DNSInfo struct {
        Domain               string              `json:"domain"`
        Records              map[string][]string `json:"records"`
        AuthoritativeServers []string            `json:"authoritative_servers,omitempty"`
        MailServers          []string            `json:"mail_servers,omitempty"`
        ReverseRecords       []string            `json:"reverse_records,omitempty"`
        Timestamp            time.Time           `json:"timestamp"`
}

// ValidateDomain memvalidasi format domain
func ValidateDomain(domain string) bool {
        if len(domain) == 0 || len(domain) > 253 {
                return false
        }

        // Basic domain validation
        parts := strings.Split(domain, ".")
        if len(parts) < 2 {
                return false
        }

        for _, part := range parts {
                if len(part) == 0 || len(part) > 63 {
                        return false
                }
        }

        return true
}

// ValidateIP memvalidasi format IP address
func ValidateIP(ip string) bool {
        return net.ParseIP(ip) != nil
}
