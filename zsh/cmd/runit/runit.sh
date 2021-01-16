#!/usr/bin/env zsh
# -*- coding: utf-8 -*-
    # shellcheck shell=bash
    # shellcheck source=/dev/null
    # shellcheck disable=2178,2128,2206,2034

# Do Over - run a command with args ... over and over and over ...
# it's just quicker and easier than debugging for short scripts
# add compile steps, CI pipeline stuff, logging ... whatever

# similar to https://linux.die.net/man/1/watch
# once I found 'watch' I renamed this script from 'do_over' to 'watch' to be
# consistent, so different versions may be available with either name

. $(which ansi_colors)

COUNTER=0                       #* not sure why we care how many  ... but i like to see it
SLEEP_INTERVAL=3                #* Interval between iterations
SCRIPT_NAME="${0##*/}"	        #* name of this script
SCRIPT_PATH="${0%/*}"		    #* path of this script
PID=$$						    #* PID of this script - we need more data to look important

ms () { printf '%i\n' "$(( $(gdate +%s%N) * 0.001 ))"; }

t() { t0=$(ms); }
deltat () {  echo $(( t1 - t0 )); }

# get timer so far with no t0 reset
lap () {
	t1=$(ms)
	dt=$(( t1 - t0 ))
	printf '%i\n' "$dt"
}

if [[ $1 == "--sleep" ]]; then
    SLEEP_INTERVAL=$2
    shift; shift
fi

while true; do
    echo "${LIME}<===============> CTRL-C to STOP the MADNESS <===============>${RESET:-}"
    #! literally just running whatever is passed to it  ------>    #! CAREFUL!!!    <------
    t0=$(ms)
    "$@"
    RETVAL=$?				# not very useful most of the time, but eh ...
    t1=$(ms)
    #! literally just running whatever is passed to it  ------>    #! CAREFUL!!!    <------
    echo "${ATTN:-}execution time: $(lap) ms${RESET:-}"
    echo "${BLUE:-}script text: $@${RESET:-}"
    echo "${CHERRY:-}script result: $RETVAL${RESET:-}"
    echo "${MAIN:-}${SCRIPT_NAME}${CANARY:-} (PID $PID)"
    echo "${GO:-}${SCRIPT_PATH}"
    echo "options: SLEEP = ${SLEEP_INTERVAL}  COUNT = ${COUNTER}${ATTN:-}  "
    ((COUNTER+=1))
    sleep $SLEEP_INTERVAL
done
