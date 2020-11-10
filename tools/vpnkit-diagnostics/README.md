A simple Windows program to download vpnkit diagnostics

Usage:
```
go run main.go
```
The output is `vpnkit.tar` which contains the files (sorry about the inconvenient path separators):

- `connections`: a list of active socket connections
- `flows`: active TCP flows
- `endpoints`: active IP endpoints
- `capture\\http_proxy.pcap`: a capture of the recent traffic to/from the HTTP(S) proxy
- `capture\\bad.pcap`: a capture of any traffic which failed to parse
- `capture\\icmp.pcap`: a capture of the recent ICMP traffic
- `capture\\ntp.pcap`: a capture of the recent NTP traffic
- `capture\\all.pcap`: a capture of all recent traffic
- `capture\\dns.pcap`: a capture of the recent DNS traffic
