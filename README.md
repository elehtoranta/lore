# LORE - Ludicrously Offensive Retaliation Engine
## Description
This is the solution the space squirrels have been looking for. The mediocre gambling bot.

## Installation
1. Ensure you have (https://go.dev/doc/install)[Go] installed, minimum of version 1.21.x. Check with `go version`.
2. Clone the repo.
3. You **must** have your API key in an environment variable called `LORE_API_KEY`. This can be achieved (for example) by either:
    1. in the project root, create a `.env` file, and add a line `LORE_API_KEY=<YOUR_KEY_HERE>`,
    2. run the program through `env` (`env LORE_API_KEY=<YOUR_KEY_HERE> ...`)
4. In project root, run:
   ```bash
   go get . # Get dependencies
   go build # Compile
   ```

## Usage
Takes the number of games to play as a parameter. Accepts a number from 1 to 100.
```bash
./lore <number-of-games> # OR go run . <number-of-games>
```

## Notes
The shell output is a text wall with pretty much all the information about the execution, which can naturally be grepped by a nifty developer.

## Author
Erkka Lehtoranta
[GitHub](https://github.com/elehtoranta)
[LinkedIn](https://linkedin.com/in/lehtoranta)
