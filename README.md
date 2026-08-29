[![Go Version](https://img.shields.io/badge/Go-1.25-blue.svg)](https://go.dev/)
[![License](https://shields.io/badge/license-Apache%202-blue)](https://opensource.org)
# Gin & Tonic
A clean, production-ready Model-View-Controller boilerplate written in Go, for Gin Framework, and integrated with GORM, and a fully structured static assets pipeline.

## 🚀 Getting started
Follow these instructions to get a new **Gin & Tonic** project up and running on your local machine.

### Prerequisites
You need both **Git** and **Go** installed on your computer. You can check if you have them by running these commands in your terminal:
```bash
# Check Git version
git --version
```
```bash
# Check Go version
go version
```

### Git & Go Setup

1. **Install Git:** Download it from [git-scm.com](https://git-scm.com) if you don't have it.
2. **Install Go:** Download the installer from [go.dev/doc/install](https://go.dev).
3. **Configure Git:** Set up your identity in your terminal:

    ```bash
    # Set your name
    git config --global user.name "[Your Name]"

    # Set your email address
    git config --global user.email "[your.email@example.com]"
    ```

### New project intialization and installation
1. **Clone this repository** to your local machine:
    ```bash
    git clone https://github.com/adu013/GinAndTonic.git <!-- INSERT YOUR PROJECT NAME --> && rm -rf <!-- INSERT YOUR PROJECT NAME -->/.git
    ```
    ```bash
    cd <!-- INSERT YOUR PROJECT NAME -->

    ```

2. **Run the automation setup script** to dynamically rename the project module and update all internal imports:
    ```bash
    ./setup.sh
    ```

3. **Start your development server by using Makefile or by using standard Go run command**:
    ```bash
    make run
    ```
    OR
    ```bash
    go run main.go
    ```

Now open `http://localhost:8085` in your web browser!

## 🏗️ Project Architecture

```
├── cli/             # Custom command-line tools
├── config/          # Global system settings
├── controllers/     # Handles incoming HTTP requests and extracts data
├── database/        # Initializes your local SQLite connection and GORM setup
├── helpers/         # Context-aware helpers (e.g., password hashing, JWT generation)
├── middlewares/     # All middlewares
├── models/          # Defines your database schemas (GORM models)
├── routes/          # Connects your URLs to the proper controllers and middleware (CSRF, etc.)
├── utils/           # Independent utility functions (e.g., string formatting, date parsers)
├── views/           # Manages JSON API responses or HTML templates
├── .env             # Your local environment settings
├── .gitignore       # Tells Git to ignore the matched folder/file patterns
├── go.mod           # Go module dependency manifest (Crucial)
├── go.sum           # Go dependency checksum safety log (Crucial)
├── LICENSE          # The Apache license, Version 2 file (Mandatory)
├── local.db         # (Git-ignored) Automatically created on your very first run!
├── main.go          # The clean, minimal entry point of your app
├── Makefile         # Shortcuts for common developer commands (run, test, build)
├── README.md        # This beautiful documentation file (Mandatory)
└── setup.sh         # One-click automation script to configure the project for first-time use
```

## 💎 Key Features

*   **Gin Web Framework:** Ultra-fast routing and middleware handling.
*   **GORM Integration:** Powering the data layer with Go's excellent, developer-friendly ORM.
*   **Zero-Setup SQLite:** Start coding immediately without managing external database servers.
*   **Built-in Security & CSRF Protection:** Secure against Cross-Site Request Forgery attacks right from the very first boot.
*   **Decoupled MVC Layout:** Zero circular dependencies, designed carefully to respect Go's compiler rules.
*   **Production Ready:** Built-in environment file configuration, crash recovery, and structured logging middleware.


## 🤝 Contributing & Feedback
This project is 100% open-source and built for the community.

Whether you want to fix a bug, suggest an architectural improvement, or just share your thoughts, all feedback is welcome! Please feel free to open an issue or submit a pull request.

If this boilerplate helps you or your team, please consider giving it a ⭐ Star to help other developers find it!

## 🤝 How to Contribute

1. Fork the Project
2. Create your Feature Branch:   `git checkout -b feature/YourFeatureBranch`
3. Commit your Changes:          `git commit -m 'Add your amazing features'`
4. Push to the Branch:           `git push origin feature/YourFeatureBranch`
5. Open a Pull Request

## 📄 License

Distributed under the Apache License, Version 2.0. See `LICENSE` for more information.

## PS.
1. Deliberately kept "middlewares" as folder name. It intuitionally tells it contains a few middleware files.
