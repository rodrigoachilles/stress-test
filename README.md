Translations:

* [French](README_fr.md)
* [Portuguese (Brazil)](README_pt_br.md)

---

# 🌐 Stress Test CLI (stress-test)

![Project Logo](assets/stress_test-logo.png)

Welcome to the Stress Test CLI developed in Go! This project allows you to perform stress tests on a web service by specifying the URL, the total number of requests, and the number of simultaneous calls.

## 📑&nbsp;Table of Contents

- [📖 Introduction](#introduction)
- [🛠 Prerequisites](#prerequisites)
- [⚙️ Installation](#installation)
- [🚀 Usage](#usage)
- [🔍 Examples](#examples)
- [🤝 Contribution](#contribution)
- [📜 License](#license)

## 📖&nbsp;Introduction

This CLI tool allows you to perform stress tests on a web service by specifying parameters such as the URL, the total number of requests, and the number of simultaneous calls. It generates a detailed report with information about the requests made.

## 🛠&nbsp;Prerequisites

Make sure you have the following items installed before continuing:

- [Go](https://golang.org/doc/install)
- [Docker](https://www.docker.com/get-started)

## ⚙️&nbsp;Installation

1. Clone this repository:

    ```sh
    git clone git@github.com:rodrigoachilles/stress-test.git
    cd stress-test
    ```

2. Build the Docker image:

    ```sh
    docker build -t stress-test-app .
    ```

## 🚀&nbsp;Usage

After building the Docker image, you can run the stress test using the following command:

```sh
docker run stress-test-app --url=http://example.com --requests=1000 --concurrency=10
```

### 📄&nbsp;Command Line Parameters

- `--url`: URL of the web service to be tested.
- `--requests`: Total number of requests to be made.
- `--concurrency`: Number of simultaneous calls.

### 📋&nbsp;Report

After the test, a report will be generated with the following information:
- Total time taken
- Total number of requests made
- Number of requests with HTTP status 200
- Distribution of other HTTP status codes (e.g., 404, 500, etc.)

## 🔍&nbsp;Examples

Here are some usage examples of the stress test CLI:

- Test a web service with 1000 requests and 10 concurrent calls:
    ```sh
    docker run stress-test-app --url=http://example.com --requests=1000 --concurrency=10
    ```

## 🤝&nbsp;Contribution

Feel free to open issues or submit pull requests for improvements and bug fixes.

## 📜&nbsp;License

This project is licensed under the MIT License.
