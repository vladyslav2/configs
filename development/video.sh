#!/bin/bash

slop=$(slop -f "%x %y %w %h %g %i") || exit 1
read -r X Y W H G ID <<< $slop
yes | ffmpeg -f x11grab -s "$W"x"$H" -i $DISPLAY.0+$X,$Y -f alsa -i pulse /tmp/video.webm
