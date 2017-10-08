# emuromcleanup

Cleans up a directory of rom files and deletes broken [b*] and overdump [o*] files that you will probably never play. It ignores files that contain ! (verified working). 

So you just set up your raspberry pi emulator and you unzipped all of your torrented roms but realized that scrolling through 13,000 files in retropie is annoying enough as it is. But then you spent valuable minutes of your life setting favorites and after you updated/restarted your emulationstation most/all of your favorites disappear! No? Too specific? Well maybe you just have a directory full of files that match [o##] or [b##] and you realize you just don't want them anymore. This will help you thin the herd.

More info on ROM suffixes: https://64bitorless.wordpress.com/rom-suffix-explanations/

This _WILL DELETE THE FILES_ so do not run this on your only copy of the files.

## Usage

Build this and run it.

    emuromcleanup.exe <directory>

Optional parameters

    -d Dry run only, don't perform actual deletes
    -t Remove translations (files that have T+ or T- in square brackets)
    -h Remove hacks (files that have h in square brackets)
    -p Remove pirates (files that have p in square brackets)

