# Hide File in PNG - Go Application

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Hiding a File](#hiding-a-file)
- [Retrieving a Hidden File](#retrieving-a-hidden-file)
- [License](#license)

## Introduction

This Go application allows you to hide any file within a PNG image, effectively concealing the file within the image's data. This can be a fun and creative way to protect sensitive information or share files discreetly.

## Features
  
- Hide any file within a PNG image.
- Retrieve hidden files from PNG images.

## Getting Started

### Prerequisites
  
Before using this application, you need to have Go 1.21.0 installed on your system.

### Installation

Clone the repository:

```bash
git clone https://github.com/DanteLorenzo/hpng.git

cd hpng
```

Build the application for Windows (64-bit):

```cmd
GOOS=windows GOARCH=amd64 go build -o myapp_windows_amd64.exe
```

Build the application for Linux (64-bit):
  
```bash
GOOS=linux GOARCH=amd64 go build -o myapp_linux_amd64
```

## Usage

To hide a file within a PNG image, follow these steps:

```bash
hpng -s PNG_IMAGE -f FILE_TO_HIDE
```

To retrieve hidden file:

```bash
hpng -r -s PNG_WITH_HIDDEN_FILE -f OUTPUT_NAME
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---
