#!/bin/sh
# apt install imagemagic xclip

rm /tmp/screenshot.png; import /tmp/screenshot.png && sleep 1 && cat /tmp/screenshot.png | xclip -selection clipboard -t image/png
