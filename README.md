# Discord Drive

## Build

```bash
go build discord-drive.go
```

## Usage

1. create .env file

    ```bash
    TOKEN=YOUR-DISCORD-BOT-TOKEN
    CHANNELID=STORAGE-CHANNEL-ID
    PORT=5000
    CHUNKSIZE=10000000
    ```

2. execute

- WebUI mode
    `./discord-drive`
- Command line mode
    `./discord-drive -c -u path/to/file`
    `./discord-drive -c -d filename`
