# Countdown Timer
This project was part of a challenge to create a countdown timer.

## Requirements

### Phase 1
A cli tool that takes optional arguments for duration, output file and format.

```bash
$ countdown -h
Use this tool to create a count down timer, it will create a file and
update it as the timer counts down to 0.

Usage:
  countdown [flags]

Flags:
  -d, --duration duration    Duration of the count down (default 5m0s)
  -f, --format string        The format to output the count down timer in (default "mm:ss")
  -h, --help                 help for countdown
  -o, --output-file string   Where the output will be written (default "~/.countdown")
```

The tool should write every second.

### Phase 2
The following additional features should form phase 2 of the project.
* The tool should run timers in the background.
* Multiple timers should be supported.
* There should be the ability to stop timers.

## Usage
```bash
# The default will set a 5 minute timer and write to the file .countdown in your home directory
$ countdown
# This will set a 10 minute timer and write to the file .cd in your home directory
$ countdown -d 10m -o ~/.cd
```
