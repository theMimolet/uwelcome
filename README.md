# uWelcome

*uWelcome is a configurable CLI Banner made for your favorite Linux systems!*

**WIP**: Some features still need testing.

**Contributions are welcome!** If you want to contribute, you're welcome to submit a pull request or open an issue - it's very much appreciated ❤️

[Access the documentation](https://github.com/theMimolet/uwelcome/tree/main/docs)

## Roadmap

Here are features that are planned for the future:

- Shows a warning if the installed system image is over a certain amount of time (e.g. one month - configurable)
- Shows a warning if the installed system image is unverified (could be disabled via configuration)

> Both are already done in Bazzite, so it's only a matter of porting them over to uWelcome.

## How to try

### Install it with Homebrew

If you have Homebrew installed on your system, you can install uWelcome with the following command:
```
brew install themimolet/tap/uwelcome
```

It's the recommended way to install uWelcome on your system, as it will automatically update via Homebrew.

### Download it from the releases page

You can download the latest release from the [releases page](https://github.com/theMimolet/uwelcome/releases).

You can then rename it to `uwelcome` and place it in your usual `/bin` folder.

> Note: You won't receive automatic updates, you will have to download uWelcome at each new release.

### Compile from source

You'll need to have [`go`](https://repology.org/project/go/versions) installed on your system to compile uWelcome from source.

Then you'll have to simply clone the repository and then build the binary:
```
git clone https://github.com/theMimolet/uwelcome
cd uwelcome
go build
./uwelcome
```

You'll then have the `uwelcome` binary in the current directory, which you can just drop into your usual `/bin` folder and it will work without any further setup (except for the configuration file if you want to customize it).

## Commands

uWelcome supports the following commands:

- `toggle`: toggles the MOTD on or off for the current user
- `enable`: always enables the MOTD for the current user
- `disable`: always disables the MOTD for the current user
- `version`, `--version`, `-v`: displays the version of uWelcome you're currently using
