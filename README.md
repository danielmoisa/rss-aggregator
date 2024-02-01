# RSS Aggregator API

## Overview

This project is a simple RSS aggregator API built with Go, which allows users to subscribe to RSS feeds and retrieve aggregated content. It also supports user authentication using API keys.

## Features

- Subscribe to multiple RSS feeds
- Fetch aggregated content from subscribed feeds
- User authentication using API keys

## Prerequisites

Before you begin, ensure you have the following dependencies installed:

- [Go](https://golang.org/dl/)
- [Docker](https://docs.docker.com/get-docker/)

## Getting Started

```bash
git clone https://github.com/danielmoisa/rss-aggregator.git
cd rss-aggregator

# Add the .env file
PORT=
DB_URL=

# Set Up the Database
docker-compose up -d
make migrate-up

# Generate sqlc  go code
make gen

# Build and Run the API
make run
