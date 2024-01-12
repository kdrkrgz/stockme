# StockMe CLI Tool

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)



StockMe is a simple CLI tool designed to provide stock market alerts and updates at regular intervals.

## Features

- Monitor stock prices for specified symbols.
- Display stock history with timestamps.
- Sound alerts for price increases and decreases.

## Installation

To run or build StockMe, you need to have Go installed on your machine, or you download direct bin file and run.

```bash
> git clone https://github.com/kdrkrgz/stockme.git
> make build
```
## Usage

```bash
stockme [flags]
```

## Flags
- -h, --help: Show help messages.
- -l, --list-symbols: List of stock symbols to monitor.
- -m, --mute: Mute alerts.
- -p, --period: Run period in minutes.
- --list-size: Stock list size.
- -s, --symbol: Stock symbol to monitor.

## Examples
Monitor stock prices for specific symbols:

```
stockme -l "ALTINS1" -p 10 --list-size 30
```

## Screenshots
![help](https://img001.prntscr.com/file/img001/NbyWVQnsShSTsrmYijPvnA.png)

![stocks](https://img001.prntscr.com/file/img001/H-mwWz9DRcOoz4SE-KiWzg.png)

License
This project is licensed under the MIT License - see the LICENSE file for details.