# LocalAIBuddy

A simple terminal app in go to use AI in your terminal now currently only with gemini

## Installation

Make sure you have `go` installed and enter the following command

```bash
git clone <this-repo> # After forking this repo
go mod tidy

go run main.go

# or build and run sperately

go build main.go
./main # respective file extension .exe etc
```

### Important

Make sure you get an API Key for Gemini from [google](https://aistudio.google.com)

After that make an `.env` file and then add this variable

```env
GEMINI_API_KEY=<your-api-key>
```

