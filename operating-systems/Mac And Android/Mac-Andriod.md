 macOS
(macOS commands are largely similar to UNIX/Linux, but here are a few macOS-specific commands)


brew

Scenario: Install software packages using Homebrew.
Usage: brew install wget

open

Scenario: Open a file or application.
Usage: open /Applications/Safari.app
diskutil

Scenario: Manage disks and volumes.
Usage: diskutil list
defaults

Scenario: Modify macOS system and application preferences.
Usage: defaults write com.apple.finder AppleShowAllFiles YES

say

Scenario: Make your Mac speak the specified text.
Usage: say "Hello, world!"



 Android (via ADB - Android Debug Bridge)

adb devices

Scenario: List connected Android devices.
Usage: adb devices

adb shell

Scenario: Start a remote shell on a connected Android device.
Usage: adb shell

adb install

Scenario: Install an APK on a connected device.
Usage: adb install app.apk

adb push

Scenario: Copy files from the host computer to the Android device.
Usage: adb push file.txt /sdcard/

adb logcat

Scenario: View device logs, useful for debugging.
Usage: adb logcat

adb reboot

Scenario: Reboot the connected Android device.
Usage: adb reboot