# khoros-community-cli
Get and manage Khoros communities via a CLI tool

This is a quick project to help me learn Go and also learn more about CLI Tooling

# Usage (WIP)

## Commands
| Command | Action |
| --- | --- |
| get | Gets a collection |
| create | Creates the following collection |
| update | Updates the following collection |

## Collections
| Collection | Required args |
| posts | `size-modifier` `params` |
| ideas | `size-modifier` `params` |
| users | `size-modifier` `params` |

## Examples
Get list of messages
```shell
get message latest
```
