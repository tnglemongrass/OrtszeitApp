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
   go build
   ```

## Running the App
To run the app, execute the built binary:
```sh
./OrtszeitApp
```

## Building an APK
For building an installable Android APK:

1. Make sure you have Android SDK and NDK installed and properly configured
2. Run the following command to build the APK:
   ```sh
   fyne package -os android -appID com.example.ortszeit -icon assets/icons/sun.png
   ```
3. The APK will be generated in your current directory
4. To install on a connected Android device:
   ```sh
   adb install ortszeit.apk
   ```

For more information about mobile packaging with Fyne: https://docs.fyne.io/started/mobile.html

## Notes
- The app currently uses a fixed location for demonstration purposes. You can modify the `main.go` file to fetch the user's location dynamically.
- Ensure you have the necessary environment set up to build and run Fyne applications on Android.
