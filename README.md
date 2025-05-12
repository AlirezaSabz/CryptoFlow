# CryptoFlow

Crypto Data Collector

Crypto Data Collector is a tool that collects cryptocurrency-related data from various sources and stores the collected data in a database for further analysis. 
This project uses **Temporal** for workflow orchestration, **Binance API** for real-time market data, and **SQLite** for persistent storage.

## Tools & Technologies Used
- **Temporal**: Workflow orchestration for managing data collection processes.
- **Binance API**: Fetching real-time market data for cryptocurrencies.
- **SQLite**: Database for storing collected data.
- **Go (Golang)**: Programming language used for development.

## Features
- Collects real-time cryptocurrency data from Binance.
- Stores collected data in an SQLite database.
- Uses Temporal for managing long-running workflows and retries.
- Scalable for future expansion to include additional data sources.



