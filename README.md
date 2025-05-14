# PokeCLIpse-Client

PokeCLIpse is a command-line Pokémon game client where you can explore the world, encounter wild Pokémon, and battle with your party.

## Features

- **Player Account Management**: Register a new account or login to an existing one
- **Starter Pokémon Selection**: Choose your first Pokémon companion
- **World Exploration**: Navigate between different locations
- **Wild Encounters**: Find, battle and catch wild Pokémon
- **Progression System**: Level up your Pokémon, learn new moves, and evolve
- **Battle System**: Automated battle system

## Installation

### Prerequisites

- Go 1.23 or later
- Git

### Clone and Build

```bash
# Clone the repository
git clone https://github.com/moceviciusda/pokeCLIpse-client.git
cd pokeCLIpse-client

# Build the project
go build -o pokeclipse
```

## Configuration

Create a .env file in the root directory with the following content:

```
SERVER_API_BASE_URL=http://your-server-address:port/v1
```
or use the preview server from .env.example

## Usage

Run the client:

```bash
./pokeCLIpse-client
```

### Available Commands

- `help` - Display available commands and their descriptions
- `register <username> <password>` - Create a new account
- `login <username> <password>` - Log in to an existing account
- `location` - Get information about your current location
- `location next` - Move to the next location
- `location previous` - Move to the previous location
- `location search` - Look for wild Pokémon
- `exit` - Quit the game

## Battle System

When encountering a wild Pokémon using `location search`, you can choose to:
1. Enter battle mode with your current party
2. Run away to avoid the encounter

Battles is automated, meaning you won't have to manually select moves. The system will handle the battle mechanics for you.

If one of your Pokémon faints, you can choose which Pokémon to send out next. The battle will continue until either you or the wild Pokémon is defeated.

If wild Pokémon is defeated:
1. Your Pokémon will gain experience points, level up, and learn new moves
2. You will have the option to catch the wild Pokémon

## Dependencies

- [github.com/chzyer/readline](https://github.com/chzyer/readline) - Interactive command-line input
- [github.com/gorilla/websocket](https://github.com/gorilla/websocket) - WebSocket client for real-time battle interactions
- [github.com/joho/godotenv](https://github.com/joho/godotenv) - Load environment variables from .env file

## Project Structure

```
.
├── cli.go                 # CLI main interface
├── command_*.go           # Command implementations
├── commands.go            # Command definitions
├── main.go                # Application entry point
├── go.mod                 # Go module definition
├── go.sum                 # Go module checksums
├── internal/              # Internal packages
│   └── serverapi/         # Server API client
├── pkg/                   # Public packages
│   ├── ansiiutils/        # ANSI terminal utilities
│   └── pokeutils/         # Pokémon utilities
└── .env                   # Environment configuration
```
