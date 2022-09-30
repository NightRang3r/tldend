# tldend
Pipe subdomains via stdin and print to stdout only subdomains that are matching a specified tld domain

## Install

```go install -v github.com/NightRang3r/tldend@main```

## Build

```go build tldend.go```

## Usage

```cat starbucks_subdomains.txt | tldend -d starbucks.com```
