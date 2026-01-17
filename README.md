# Dev Multi Tool (DMT)

DMT (Dev Multi Tool) is a small, fast, and extensible command-line toolkit for common developer tasks.  
Written in Go, DMT aims to collect simple utilities and workflows (build helpers, checks, quick automation) into a single, consistent CLI so teams and individuals can save time and reduce repetition.

## Goals
- Provide a lightweight, zero-dependency CLI for common developer tasks.
- Be easy to install, use, and extend.
- Offer clear, scriptable commands that can be used in local development and CI.

## How it works
- The project is implemented in Go for fast startup and single-binary distribution.
- DMT exposes a CLI (binary typically named `dmt`) with subcommands and flags.
- Each subcommand performs a focused task (examples: format, lint, test, scaffold). Commands are implemented as Go packages and wired into the CLI entrypoint.
- The tool is designed to be modular: new commands can be added as separate packages and registered in the main CLI.

## Acknowledgements
- Built with Go — thanks to the Go community and ecosystem for tooling and libraries.
- Inspired by other small, modular CLI tools that favor composability and simplicity.
