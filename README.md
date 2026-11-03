# ASCII Art Web

## Description

A web application written in Go that converts text into ASCII Art using different banners:

- standard
- shadow
- thinkertoy

## Authors

Tony Maher

## Usage

Run:

```bash
go run .
```

Open:

http://localhost:8080

## Implementation Details

- HTTP server using net/http
- HTML templates using html/template
- POST requests send text and banner selection
- ASCII Art is generated from banner files
- Result displayed on webpage