# Ortszeit Android App

## Description
This is an Android app built using Go and Fyne that displays the "Ortszeit" or true solar time for the user's current location by means of the sun, instead of the time zone the user is in.

## Dependencies
- Fyne: `fyne.io/fyne/v2`
- Time: `time`
- Math: `math`

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
4. Build the app:
   ```sh
   cd cmd/ortszeit
   go build
   ```

## Running the App
To run the app, execute the built binary from the cmd/ortszeit directory:
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

## Building an APK
For building an installable Android APK:

1. Make sure you have Android SDK and NDK installed and properly configured
2. From the cmd/ortszeit directory, run the following command to build the APK:
   ```sh
   fyne package -os android -appID com.ortszeitapp.ortszeit -icon ../../assets/icons/sun.png
   ```
3. The APK will be generated in your current directory
4. To install on a connected Android device:
   ```sh
   adb install ortszeit.apk
   ```

For more information about mobile packaging with Fyne: https://docs.fyne.io/started/mobile.html

## Notes
- The app currently uses Munich (48.1351°N, 11.5820°E) as the default location
- Ensure you have the necessary environment set up to build and run Fyne applications on Android