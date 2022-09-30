# tldend
Pipe subdomains via stdin,  check and print to stdout only subdomains that are  matching to a specified tld domain

## Install

```go install -v github.com/NightRang3r/tldend@main```

## Build

```go build tldend.go```

## Usage

```cat starbucks_subdomains.txt | tldend -d starbucks.com```
