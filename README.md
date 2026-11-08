# Optinium

Optinium is a Windows payload builder written in Go, featuring a graphical user interface built with WebView.

It allows users to customize and compile Windows executables by selecting different modules and configuring a webhook URL.

### Features

- **System Information:** Collects system-related information.
- **Discord Data:** Retrieves Discord-related data.
- **Network Information:** Collects IP and Wi-Fi information.
- **Browser Data:** Collects visited URLs and cookie-related information.
- **Login Data:** Handles stored login-related information.
- **Card Data:** Includes a module for payment-card-related data.
- **Screenshots:** Supports screenshot functionality.
- **File Operations:** Includes file collection functionality.
- **Persistence:** Includes a persistence module.
- **Additional Options:** Fake error messages, execution delays, and self-deletion.
- **Executable Builder:** Compiles the selected modules into a standalone Windows executable.

The application uses embedded Go source files for its payload modules and invokes the Go compiler to generate the final executable.

**Note:** Originally created approximately one year ago and uploaded to GitHub today.
