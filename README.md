# Search IPs and Server Names CLI Tool

This repository contains a Go (Golang) Command Line Interface (CLI) application that allows users to search for IP addresses and server names associated with a given host on the internet.

## Features

- Look up IP addresses for a specified host.
- Look up server names associated with a specified host.

## Installation

Ensure you have Go installed on your machine. Clone this repository and build the application using the following commands:

```bash
git clone https://github.com/otaviomart1ns/ipserverlookup.git
cd search_ips_go
go build -o ipserverlookup
```
## Usage

After installation, you can use the tool with the following commands:

### IP Lookup

To find IP addresses for a given host:

```bash
./ipserverlookup ip --host example.com
```
### Server Name Lookup

To find server names for a given host:

```bash
./ipserverlookup server --host example.com
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

Thanks to all the contributors who spend time to help improve this tool.  
Special thanks to [urfave/cli](https://github.com/urfave/cli) for the CLI library.

## Contact

Otavio Martins - [LinkedIn Profile](https://www.linkedin.com/in/otavio-mart1ns/)
