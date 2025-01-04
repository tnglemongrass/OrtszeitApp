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

## Notes
- The app currently uses a fixed location for demonstration purposes. You can modify the `main.go` file to fetch the user's location dynamically.
- Ensure you have the necessary environment set up to build and run Fyne applications on Android.
