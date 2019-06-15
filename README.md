## Simple Web Application Firewall

## Features
 - XSS and SQL injection detection and blocking
 - Whitelist/blacklist of IP Addresses
 - Block file uploads by file extension

## Configuration
Configuration is in `config.toml`

```toml
# Forward requests to this URL
Upstream = "http://127.0.0.1:8000"

# Listen for requests on this address
ListenAddress = ":8080"

# Block or allow only these IP addresses, allowed modes are "whitelist" and "blacklist"
IpFilterMode = "whitelist"

# Array of IP Addresses that are whitelisted/blacklisted
IpAddresses = [
    "127.0.0.1"
]

# Block file uploads by extension
DenyExtensions = [
    "php",
    "aspx",
    "sh",
    "html",
    "jsp"
]
```

### To run

```bash
$ make run
```

### Execute test scripts with
```bash
$ make run-tests
```
