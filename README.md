# GoAlive

GoAlive is a tool written in Go for checking the availability of subdomains. It scans the provided list of subdomains and sends HTTP HEAD requests to each. It reports the results to the user and optionally displays the HTTP status code.

## Features

- Supports HTTP and HTTPS protocols.
- Timeout duration can be configured by the user.
- Verbose mode shows active and inactive subdomains with detailed status codes.
- Utilizes concurrent goroutines for parallel scanning, ensuring fast results.
- Simple and user-friendly command-line interface.

## Usage

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/goalive.git
   cd goalive

2. Build the executable:

   ```bash
   go build -o goalive main.go

3. Running the Program:

   ```bash
   ./goalive -i subdomains.txt -t 5000 -v


-i: File containing the list of subdomains.
-t: Request timeout in milliseconds.
-v: Verbose mode, prints active and inactive subdomains with status codes.

### Contributig
Found a bug? Report it using GitHub issues.
Have an idea for improvement? Open a pull request.

### License
This project is licensed under the MIT License - see the LICENSE file for details.
