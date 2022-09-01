# pushover-cli

A simple CLI tool to send Pushover messages.

I made this tool as an exercise to learn Go, so keep your expectations low on code quality. :)

## Usage

To use this tool, you need to have a Pushover API token and user key, which you can get from https://pushover.net.

In your shell, set these variables prior to running the binary: 

    export PUSHOVER_API_TOKEN=abcdefghijklomnpqrstuvxyz1234567890
    export PUSHOVER_USER_KEY=abcdefghijklomnpqrstuvxyz1234567890

Build or download the binary and put it in a directory in your `$PATH`. Then, run like this:

    pushover-cli --title "A Title" --message "This is the message" --url "https://google.com"    


## Flags


        --callbackurl string   A callback URL for your message. (optional)
        --device string        Device to send your message to (optional)(if left blank all devices will be messaged)
        --expire string        Time to expire (optional)
    -h, --help                 help for pushover-cli
        --html string          Parse HTML in message (True/False)(optional)
        --message string       Your message (required)
        --monospace string     Monospaced (True/False)(optional)
        --prio string          Your message's priority (-2 - 2) (optional)
        --retry string         Time to retry (optional)
        --sound string         A sound for your message (optional)
        --timestamp string     Include a timestamp in your message (optional)
        --title string         A title for your message (optional)
        --url string           An URL to send with your message (optional)
        --urltitle string      A title for your URL (optional)