#!/bin/bash

# BTree v1.0
#
# $Copyright: $
# Copyright (c) 2007 by Steve Baker (ice@mama.indstate.edu)
# All Rights reserved
#
# This program is free software; you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation; either version 2 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program; if not, write to the Free Software
# Foundation, Inc., 59 Temple Place, Suite 330, Boston, MA  02111-1307  USA

colorize=0
pats=""
ccodes=""
npats=0

allfiles=0
dirsonly=0
printsize=0
printprot=0
printuid=0
printgid=0

ndirs=0
nfiles=0

listpats()
{
  local IFS code

  IFS=":"
  for code in $LS_COLORS
  do
    echo "$code"
  done
}

initcolor()
{
  local IFS

  if [ "$LS_COLORS" ]; then
    colorize=1
  else
    return
  fi

  IFS="="

  n=0

  while read pat colcode
  do
    case "$pat" in
      "no") no="$colcode" ;;
      "fi") fi="$colcode" ;;
      "di") di="$colcode" ;;
      "ln") ln="$colcode" ;;
      "pi") pi="$colcode" ;;
      "so") so="$colcode" ;;
      "bd") bd="$colcode" ;;
      "cd") cd="$colcode" ;;
      "su") su="$colcode" ;;
      "sg") sg="$colcode" ;;
      "ex") ex="$colcode" ;;
      *) pats[$n]="$pat";
	 ccodes[$n]="$colcode";
	 let n++;
	 ;;
    esac
  done < <(listpats)
  npats=$n
}

colorize()
{
  local n x

  if [ $colorize -eq 0 ]; then return; fi

  if [ -d "$1" ]; then echo -en "\033[${di}m"; return;
  elif [ -L "$1" ]; then echo -en "\033[${li}m"; return;
  elif [ -b "$1" ]; then echo -en "\033[${bd}m"; return;
  elif [ -c "$1" ]; then echo -en "\033[${cd}m"; return;
  elif [ -u "$1" ]; then echo -en "\033[${su}m"; return;
  elif [ -g "$1" ]; then echo -en "\033[${sg}m"; return;
  elif [ -p "$1" ]; then echo -en "\033[${pi}m"; return;
  elif [ -S "$1" ]; then echo -en "\033[${so}m"; return;
  fi

  if [ -f "$1" ]; then
    if [ -x "$1" ]; then echo -en "\033[${ex}m"; return; fi
    n=0
    for x in "${pats[@]}"
    do
      if [[ "$1" == $x ]]; then echo -en "\033[${ccodes[$n]}m"; return; fi
      let n++
    done
  fi
}

tree()
{
  local n cwd file files

  pushd -n "`pwd`" > /dev/null
  cd "$1" > /dev/null
  if [ $? -eq 1 ]; then return; fi

  n=0
  for f in $(if [ $allfiles -eq 1 ]; then echo ".*"; fi) *
  do
    if [[ "$f" = "." || "$f" = ".." ]]; then continue; fi
    if [ $dirsonly -eq 1 -a ! -d "$f" ]; then continue; fi
    if [ ! -e "$f" ]; then continue; fi
    file[$n]="$f"
    let n++
  done
  files=$n

  for ((n=0; n < files; n++ ))
  do
    if (( $n < $files-1 )); then
      echo -n "${2}|-- "
    else
      echo -n "${2}\`-- "
    fi
    if [ $printsize -eq 1 ]; then printf "[%11d]  " `stat -c %s "${file[$n]}"`; fi
    colorize "${file[$n]}"
    echo -n "${file[$n]}"
    if [ $colorize -eq 1 ]; then echo -en "\033[${no}m"; fi
    if [ -L "${file[$n]}" ]; then echo -n "-> `readlink \"${file[$n]}\"`"; fi
    echo ""
    if [ -d "${file[$n]}" -a ! -L "${file[$n]}" ]; then
      if (( $n < $files-1 )); then
        tree "${file[$n]}" "${2}|   "
      else
        tree "${file[$n]}" "${2}    "
      fi
      let ndirs++
    else
      let nfiles++
    fi
  done

  popd > /dev/null
}

initcolor

eval set -- $(getopt -- ads "$@")

while [ "$1" ]
do
  case "$1" in
   "-a") allfiles=1; shift ;;
   "-d") dirsonly=1; shift ;;
   "-s") printsize=1; shift ;;
   "--") shift; break ;;
  esac
done

for dir in "$@"
do
  cwd=`pwd`
  colorize "$dir"
  echo "$dir"
  if [ $colorize -eq 1 ]; then echo -en "\033[${no}m"; fi
  tree "$dir" ""
  echo ""
  cd $cwd
done

echo "$ndirs directories, $nfiles files"
