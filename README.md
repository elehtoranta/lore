# LORE - Ludicrously Offensive Retaliation Engine
## Description
This is the solution the space squirrels have been looking for. The mediocre gambling bot.

## Usage
1. Ensure you have Go installed, minimum of version 1.21.x. Check with `go version`.
2. Clone the repo.
3. In the project root, create a `.env` file for the API key, and add an environment variable to the file of form `LORE_API_KEY=[YOUR-KEY-HERE]`. Alternatively, you can run the binary with `env` to pass the API key from the command line.
4. In project root, run:
   ```bash
   go get .
   go install
   go run . [number-of-games]
   ```
5. Watch it go BRRRR. With a _single_ thread. In Go, the promised language of concurrency.
