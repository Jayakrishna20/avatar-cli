# 🤝 Contributing to Avatar CLI

First off, thank you for considering contributing to **Avatar (`avtr`)**! It's people like you that make the open-source community such a fantastic place to learn, inspire, and create.

## 🛠️ How to Contribute

1. **Fork the repository** to your own GitHub account.
2. **Clone the project** to your local machine:
   ```bash
   git clone https://github.com/<your-username>/file-converter.git
   ```
3. **Create a new branch** for your feature, bug fix, or improvement:
   ```bash
   git checkout -b feature/amazing-new-feature
   ```
4. **Make your changes** and test them locally. (Ensure your code is well-formatted and documented).
5. **Commit your changes** clearly describing what you did:
   ```bash
   git commit -m "feat: added an amazing new feature"
   ```
6. **Push the branch** to your fork:
   ```bash
   git push origin feature/amazing-new-feature
   ```
7. **Open a Pull Request** against the `master` branch of the original repository.

## 🏗️ Development Setup

Ensure you have **Go 1.22+** installed on your system.

```bash
# Verify Go version
go version

# Download dependencies
go mod tidy

# Build the executable locally
go build -o avtr main.go

# Run it
./avtr --help
```

### Pull Request Guidelines
- Keep your Pull Requests focused on a single issue or feature natively.
- Provide a clear and comprehensive description of your changes.
- Ensure that the GitHub Actions CI pipeline stays green (passing) for your PR.

Thank you for contributing! 🚀
