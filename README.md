# Ortszeit Android App

## Description

This is an Android app built using Go and Fyne that displays the "Ortszeit" or true solar time for the user's current location by means of the sun, instead of the time zone the user is in.

## Dependencies

- Fyne: `fyne.io/fyne/v2`
- Time: `time`
- Math: `math`
- Android SDK
- Android NDK
- Android Debug Bridge (ADB)

## Installation

1. Ensure you have Go installed on your system.
2. Ensure you have Fyne installed by running:
   ```sh
   go get fyne.io/fyne/v2
   ```
3. Clone this repository:
   ```sh
   git clone https://github.com/yourusername/OrtszeitApp.git
   cd OrtszeitApp
   ```
4. Install the Android SDK and NDK:
   - Download the Android SDK from the [official Android website](https://developer.android.com/studio#downloads).
   - Download the Android NDK from the [official Android website](https://developer.android.com/ndk/downloads).
   - Extract the SDK to a directory, for example, `C:\android`.
   - Extract the NDK to a directory, for example, `C:\android-ndk-r25c`.
5. Set the respective environment variables, e.g. in PowerShell:
   ```powershell
   $env:ANDROID_HOME = "C:\android"
   $env:ANDROID_NDK_HOME = "C:\android-ndk-r25c"
   $env:Path += ";C:\android\platform-tools"
   ```
6. Build the app:
   ```sh
   cd cmd/ortszeit
   go build
   ```

## Running the App

To run the app, execute the built binary from the `cmd/ortszeit` directory:

```sh
./ortszeit
```

## Command Line Interface (CLI)

This repository also includes a CLI version of the app, located in `cmd/cli.go`. To run the CLI, navigate to the `cmd` directory and run:

```sh
go run cli.go
```

You can also build the CLI binary using:

```sh
go build cli.go
```

And then run it using:

```sh
./cli
```

The CLI version of the app will print the current time, location, coordinates, and timezone to the console.

### Flags

- `--karlsruhe`: Use Karlsruhe, Germany as the location instead of the user's current location. This flag skips the API request and uses hardcoded values for Karlsruhe.
- `--debug`: Print JSON response for debugging purposes.

## Building an APK

For building an installable Android APK:

1. Make sure you have Android SDK and NDK are properly configured.
   - Define the environment variables, e.g. for PowerShell
       ```powershell
       $env:ANDROID_HOME = "C:\android"
       $env:ANDROID_NDK_HOME = "C:\android-ndk-r25c"
       $env:Path += ";C:\android\platform-tools"
       ```
2. Build the APK from within the `gui` directory
   ```powershell
   cd ./gui/
   $env:ANDROID_HOME="C:\android"
   $env:ANDROID_NDK_HOME="C:\android-ndk-r25c"
   fyne package -os android -appID com.example.ortszeitapp -icon ../assets/icons/sun.png -name Ortszeit
   ```
   This will generate `Ortszeit.apk` in the `gui` directory.
4. Install the APK on a connected Android device
   - Install the APK:
     ```powershell
     $env:Path += ";C:\android\platform-tools"
     adb install ortszeit.apk
     ```
   - If above fails, verify the device is recognized:
     ```sh
     adb devices
     ```

USB debugging needs to be enabled on your smartphone. If Developer Options is not visible, you can enable it by going to Settings > About phone and tapping the Build number several times until you see a message that Developer Options is enabled.

For more information about mobile packaging with Fyne: https://docs.fyne.io/started/mobile.html

**Note for Windows Users:** According to [Fyne Documentation](https://docs.fyne.io/started/), for development of the Android app on Windows, MSYS2 needs to be installed using the command:

```sh
winget install MSYS2.MSYS2
```


## Notes

- The app currently uses Munich (48.1351°N, 11.5820°E) as the default location.
- Ensure you have the necessary environment set up to build and run Fyne applications on Android.
