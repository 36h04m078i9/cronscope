# cronscope

Parse and visualize cron expressions in your terminal, showing next N execution times with timezone awareness.

## Installation

```bash
go install github.com/yourusername/cronscope@latest
```

Or build from source:

```bash
git clone https://github.com/yourusername/cronscope.git && cd cronscope && go build ./...
```

## Usage

```bash
# Show next 5 execution times for a cron expression
cronscope "0 9 * * 1-5"

# Show next 10 executions in a specific timezone
cronscope --next 10 --tz "America/New_York" "30 14 * * *"

# Use a named schedule
cronscope "@daily"
```

**Example output:**

```
Expression : 0 9 * * 1-5
Timezone   : America/New_York
Next 5 executions:

  1  →  Mon, 14 Jul 2025  09:00:00  (in 6h 23m)
  2  →  Tue, 15 Jul 2025  09:00:00  (in 1d 6h 23m)
  3  →  Wed, 16 Jul 2025  09:00:00  (in 2d 6h 23m)
  4  →  Thu, 17 Jul 2025  09:00:00  (in 3d 6h 23m)
  5  →  Fri, 18 Jul 2025  09:00:00  (in 4d 6h 23m)
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `--next` | `5` | Number of upcoming executions to display |
| `--tz` | system timezone | IANA timezone name |
| `--format` | `human` | Output format: `human`, `json`, `unix` |

## License

[MIT](LICENSE)