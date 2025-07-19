# iplookup

<p align="center">
  <a href="https://github.com/gofika/iplookup/releases"><img src="https://img.shields.io/github/release/gofika/iplookup.svg?style=flat-square" alt="Release"></a>
  <a href="https://github.com/gofika/iplookup/blob/main/LICENSE"><img src="https://img.shields.io/github/license/gofika/iplookup?style=flat-square" alt="License"></a>
  <a href="https://github.com/gofika/iplookup/actions/workflows/ci.yml"><img src="https://img.shields.io/github/actions/workflow/status/gofika/iplookup/ci.yml?branch=main&style=flat-square" alt="CI"></a>
  <a href="https://goreportcard.com/report/github.com/gofika/iplookup"><img src="https://goreportcard.com/badge/github.com/gofika/iplookup?style=flat-square" alt="Go Report Card"></a>
</p>

<p align="center">
  <b>iplookup</b> is a fast and simple command-line tool for querying IP address geolocation and related information.
</p>

## ✨ Features

- 🚀 **Fast Response** - Uses high-performance ipinfo.io API with average response time < 200ms
- 🌍 **Detailed Information** - Get complete geolocation, ISP, timezone and more
- 🎯 **Accurate Detection** - Supports Anycast IP detection (e.g., 8.8.8.8)
- 🌐 **Domain Support** - Automatically resolves domain names to IP addresses
- 💻 **Cross-Platform** - Supports Windows, macOS, Linux
- 🔧 **Zero Configuration** - No API key required, works out of the box
- 📦 **Lightweight** - Single executable file with no dependencies

## 📥 Installation

### Using Package Managers (Recommended)

#### macOS (Homebrew)
```bash
brew tap gofika/tap
brew install iplookup
```

#### Linux (Script Install)
```bash
curl -sSL https://raw.githubusercontent.com/gofika/iplookup/main/install.sh | bash
```

#### Windows (PowerShell)
```powershell
irm https://raw.githubusercontent.com/gofika/iplookup/main/install.ps1 | iex
```

### Download Pre-built Binaries

Visit the [Releases](https://github.com/gofika/iplookup/releases/latest) page to download the appropriate version for your system:

- **Windows**: `iplookup_windows_amd64.exe`
- **macOS (Intel)**: `iplookup_darwin_amd64`
- **macOS (Apple Silicon)**: `iplookup_darwin_arm64`
- **Linux**: `iplookup_linux_amd64`

After downloading, rename the file to `iplookup` (or `iplookup.exe` on Windows) and add it to your system PATH.

### Install from Source

Requires Go 1.18 or higher:

```bash
go install github.com/gofika/iplookup@latest
```

## 📖 Usage

### Basic Usage

Query IP address information:

```bash
iplookup 8.8.8.8
```

Query domain name (automatically resolved to IP):

```bash
iplookup google.com
```

Sample output:
```json
{
    "ip": "8.8.8.8",
    "hostname": "dns.google",
    "city": "Mountain View",
    "region": "California",
    "country": "US",
    "loc": "37.4056,-122.0775",
    "org": "AS15169 Google LLC",
    "postal": "94043",
    "timezone": "America/Los_Angeles",
    "anycast": true
}
```

### Advanced Options

Get raw JSON output (unformatted):

```bash
iplookup -n 8.8.8.8
```

Show help information:

```bash
iplookup -h
```

## 🔍 Output Fields

| Field | Description | Example |
|------|------|------|
| `ip` | IP address queried | `8.8.8.8` |
| `hostname` | Reverse DNS lookup result | `dns.google` |
| `city` | City | `Mountain View` |
| `region` | Region/State | `California` |
| `country` | Country code | `US` |
| `loc` | Latitude,Longitude | `37.4056,-122.0775` |
| `org` | Organization/ISP | `AS15169 Google LLC` |
| `postal` | Postal code | `94043` |
| `timezone` | Timezone | `America/Los_Angeles` |
| `anycast` | Whether it's an Anycast IP | `true` |

## 🛠️ Building

Clone the repository and build:

```bash
git clone https://github.com/gofika/iplookup.git
cd iplookup
go build -o iplookup
```

Cross-platform builds:

```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o iplookup.exe

# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o iplookup

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o iplookup

# Linux
GOOS=linux GOARCH=amd64 go build -o iplookup
```

## 🤝 Contributing

Pull requests are welcome! Please:

1. Fork this repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Thanks to [ipinfo.io](https://ipinfo.io) for providing the free IP lookup API
- Built with [GoReleaser](https://goreleaser.com) for automated releases

---

<p align="center">
  Made with ❤️ by <a href="https://github.com/gofika">Gofika</a>
</p>