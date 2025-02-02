# Ortszeit Android App <img src="assets/icons/sun.png" width="48" height="48" alt="Sun Icon" style="vertical-align: middle">

## Description

An Android app that displays the "Ortszeit" or true solar time for the user's current location by means of the sun, instead of the time zone the user is in. Built using Golang and Fyne.

## Requirements

For running and building the app:
- Go 1.22 or later

For building the Android APK:
- Android SDK
- Android NDK
- Android Debug Bridge (ADB)

## Development Setup

1. Clone this repository:
   ```sh
   git clone https://github.com/tnglemongrass/OrtszeitApp.git
   cd OrtszeitApp
   ```
2. Ensure you have Fyne installed by running:
   ```sh
   go get fyne.io/fyne/v2@latest
   go install fyne.io/fyne/v2/cmd/fyne@latest
   fyne version
   ```
3. Install the Android SDK and NDK:
   - Download the _Android SDK_ from the [official Android website](https://developer.android.com/studio#downloads).
   - Download the _Android NDK_ from the [official Android website](https://developer.android.com/ndk/downloads).
   - Extract the SDK to a directory, for example, `C:\android`.
   - Extract the NDK to a directory, for example, `C:\android-ndk-r25c`.
4. Set the respective environment variables, e.g. in PowerShell:
   ```powershell
   $env:ANDROID_HOME = "C:\android"
   $env:ANDROID_NDK_HOME = "C:\android-ndk-r25c"
   $env:Path += ";C:\android\platform-tools"
   ```
5. Build and run the CLI app:
   ```powershell
   cd cmd/cli
   go build main.go
   ./cli.exe
   ```
6. Build and install the Android app:
   ```powershell
   cd cmd/gui
   fyne package -os android -appID com.example.ortszeitapp -icon ../../assets/icons/sun.png -name Ortszeit
   adb install Ortszeit.apk
   ```

## Command Line Interface (CLI)

This repository includes a CLI version of the app, located in `cmd/cli/main.go`. To run the CLI:

```sh
go run cmd/cli/main.go
```

The CLI version prints the location, coordinates, timezone, and various calculated local and solar times based on the location. It supports options for using predefined locations (Karlsruhe and Munich) or fetching the user's current IP information (default). It also includes a debug option to print the JSON response and the IPInfo struct.

### CLI Flags

- `--karlsruhe`: Use Karlsruhe, Germany as the location instead of the user's current location
- `--munich`: Use Munich, Germany as the location instead of the user's current location
- `--debug`: Print JSON response for debugging purposes

## Building an APK

To build an installable Android APK:

1. Set up Android SDK and NDK environment variables, e.g. for PowerShell
       ```powershell
       $env:ANDROID_HOME = "C:\android"
       $env:ANDROID_NDK_HOME = "C:\android-ndk-r25c"
       $env:Path += ";C:\android\platform-tools"
       ```
2. Build the APK from within the `cmd/gui` directory
   ```powershell
   cd ./cmd/gui/
   $env:ANDROID_HOME="C:\android"
   $env:ANDROID_NDK_HOME="C:\android-ndk-r25c"
   fyne package -os android -appID com.example.ortszeitapp -icon ../assets/icons/sun.png -name Ortszeit
   ```
   This generates `Ortszeit.apk` in the `cmd/gui` directory.
3. Install the APK on a connected Android device
     ```powershell
     $env:Path += ";C:\android\platform-tools"
     adb install ortszeit.apk
     ```
   If above fails, verify the device is recognized:
   ```sh
   adb devices
   ```

Note: USB debugging must be enabled on your smartphone. Enable Developer Options in Settings > About phone by tapping Build number several times.

For more information about mobile packaging with Fyne: https://docs.fyne.io/started/mobile.html

## Building a Windows Executable

You need an installed c compiler (e.g. via `scoop install gcc`).

From within the `cmd/gui` directory:
   ```powershell
   cd .\cmd\gui
   fyne package -os windows -appID com.example.ortszeitapp -icon ..\..\assets\icons\sun.png -name Ortszeit
   ```

The executable is available as `Ortszeit.exe` in that directory.

## Notes

- The app supports switching between Karlsruhe and Munich locations. Cli also supports IP based geolocation.
- The project development started in a GitHub Codespace.
- Most of this app was developed using the _Cline (prev. Claude Dev)_ plugin in VSCode (on [github](https://github.com/cline/cline), on [marketplace](https://marketplace.visualstudio.com/items?itemName=saoudrizwan.claude-dev)). The very first lines of code were developed with the commercial model `Anthropic Claude-3-5-Sonnet-20241022`, and iterated further with `Qwen/Qwen2.5-Coder-32B-Instruct`. The initial Cline log is available [here](cline-logs/cline_task_jan-4-2025_2-29-03-pm.md).
- The icon was created by Cline with a single instruction and the sonnet model. The exact instructions were _create a simple square sun svg icon and also convert it to png and ico_. Full conversation is available [here](cline-logs/cline_task_jan-4-2025_3-23-06-pm.md).
- The readme was created and updated using Cline automatically or explicitly as needed.
- The initially generated version (Android + Fyne) did not compile and was way too complex to debug even the development setup. I started over with a simple cli version that just displayed the current time and iterated from there. It was much easier for cline to debug the cli on its own. Cline could run the script, read its output, iterate over server error messages, add temporary logging statements for the response jsons, and so on.
- The CLI was then converted to a Fyne app using Cline and the `Qwen/Qwen2.5-Coder-32B-Instruct` model. Conversation log is available [here](cline-logs/cline_task_jan-5-2025_11-21-30-am.md).

## Screenshots

![Android Screenshot](screenshots/android_screenshot.jpg "Android Screenshot")