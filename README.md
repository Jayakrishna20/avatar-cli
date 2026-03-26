<p align="center">
  <h1 align="center">🚀 Avatar (avtr) CLI</h1>
</p>

<p align="center">
  <img src="https://img.shields.io/github/actions/workflow/status/Jayakrishna20/file-converter/release.yml?style=for-the-badge&label=BUILD" />
  <img src="https://img.shields.io/github/go-mod/go-version/Jayakrishna20/file-converter?style=for-the-badge&label=GO+VERSION" />
  <img src="https://img.shields.io/github/actions/workflow/status/Jayakrishna20/file-converter/pages.yml?style=for-the-badge&label=PAGES&color=blue" />
</p>

Avatar (`avtr`) is a lightning-fast ⚡, highly capable command-line interface (CLI) tool designed for manipulating PDFs and images right from your terminal. It offers an intuitive, richly colored console experience 🎨 with detailed progress updates while gracefully handling your files!

📄 **About page:** [CLick Here](https://jayakrishna20.github.io/avatar=cli/)

<br>

<details>
<summary><h2>🧳 Installation</h2></summary>

<br>

Avatar is distributed as a standalone binary. You can install it in multiple ways:

### 1️⃣ Using Go Install (Recommended)
If you have Go installed, you can simply run:
```bash
go install github.com/Jayakrishna20/file-converter@latest
```

### 2️⃣ Package Managers (macOS, Linux, Windows)

**macOS (Homebrew):**
```bash
brew tap Jayakrishna20/file-converter
brew install avtr
```

**Linux (Debian/Ubuntu via APT):**
```bash
curl -1sLf 'https://dl.cloudsmith.io/public/jayakrishna20/file-converter/setup.deb.sh'  | sudo -E bash
sudo apt update
sudo apt install avtr
```
*(For other distributions like Fedora, use `sudo dnf install avtr`)*

**Windows (Winget / Chocolatey):**
```powershell
# Using Winget
winget install avtr

# Using Chocolatey
choco install avtr
```

### 3️⃣ Fast Installation Script (Linux & macOS)
Quickly grab the latest version and drop it right into your `/usr/local/bin` using curl:
```bash
curl -sSL https://raw.githubusercontent.com/Jayakrishna20/file-converter/main/install.sh | bash
```

### 4️⃣ Prebuilt Binaries (Windows, Linux, macOS)
You can head over to the [GitHub Releases](https://github.com/Jayakrishna20/file-converter/releases/latest) page and download the executable matching your OS directly!
</details>

<details>
<summary><h2>💺 Usage</h2></summary>

<br>

Check your active version anytime using:
```bash
avtr --version
```

### 🖼️ Convert Images to PDF (`img2pdf`)
Convert a folder filled with images into a combined PDF file.

```bash
avtr img2pdf --indir images/ --outdir output/ --output combined.pdf --pagesize A4
```
**Shorthand:**
```bash
avtr i -d images/ -o output/ -f combined.pdf -p A4
```

### 🔗 Merge PDFs (`merge`)
Combine two or more PDF files into a single, merged PDF.

```bash
avtr merge --input part1.pdf --input part2.pdf --outdir output/ --output merged.pdf
```
**Shorthand:**
```bash
avtr m -i part1.pdf -i part2.pdf -d output/ -o merged.pdf
```

### ✂️ Split PDF (`split`)
Split a PDF file into multiple parts based on your page span needs!

```bash
avtr split --input document.pdf --outdir output/ --span 1
```
**Shorthand:**
```bash
avtr s -i doc.pdf -o out/ -s 2
```

### 🔄 Update Avatar (`update`)
To grab the latest version instructions directly from your terminal, run:
```bash
avtr update
```
</details>

<details>
<summary><h2>🧩 Integrations</h2></summary>

<br>

Avatar relies on these rock-solid open-source libraries under the hood:
- **[Cobra](https://github.com/spf13/cobra)** - For the powerful CLI structure.
- **[pdfcpu](https://github.com/pdfcpu/pdfcpu)** - The amazing PDF processing engine.
- **[gopdf](https://github.com/signintech/gopdf)** - For creating PDFs from images.
- **[color](https://github.com/fatih/color)** & **[progressbar](https://github.com/schollz/progressbar)** - For the vibrant and helpful terminal UI.
</details>

<details>
<summary><h2>🏗️ Building</h2></summary>

<br>

To compile the executable locally from source:

1. Clone this repository.
2. Make sure you have Go 1.22+ installed.
3. Run the following:

```bash
go mod tidy
go build -o avtr main.go
```
The newly compiled `avtr` executable will be placed within your current directly.
</details>

<details>
<summary><h2>🤝 Contributing</h2></summary>

<br>

We love contributions! Please read our [Contributing Guide](CONTRIBUTING.md) to learn how you can report bugs, request features, or submit Pull Requests.
</details>

<details>
<summary><h2>📄 License</h2></summary>

<br>

This project is licensed under the **MIT License**. Check out the [LICENSE](LICENSE) file for more information!
</details>
