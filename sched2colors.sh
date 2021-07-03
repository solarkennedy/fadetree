#!/bin/bash

function color2hex() {
  if [[ "$1" == "Blue" ]]; then
    echo '			"#0d0a7e",'
    echo '			"#000000",'
  elif [[ "$1" == "Purple" ]]; then
    echo '			"#6d418b",'
    echo '			"#000000",'
    echo '			"#ffffff",'
  elif [[ "$1" == "Gold" ]]; then
    echo '	  		"#d4af37",'
    echo '			"#000000",'
  elif [[ "$1" == "Red/white/blue" ]]; then
    echo '	  		"#ff0000",'
    echo '			"#ffffff",'
    echo '			"#0000ff",'
  elif [[ "$1" == "Blue/white" ]]; then
    echo '	  		"#005EB8",'
    echo '			"#ffffff",'
  else
    echo "// UNKNOWN COLORS FOR $1"
  fi
}

function print_code() {
  date=$(date "+%B %d" --date "$1")
  color="$2"
  occasion="$3"
  hex=$(color2hex "$color")

cat <<EOF
	} else if TodayIs("$date", day) {
		occasion = "$occasion"
		colors = []string{
$hex
		}
EOF
}

IFS="
"
for line in $(cat sched.txt); do
  date=$(echo $line | awk '{print $1}')
  color=$(echo $line | awk '{print $2}')
  occasion=$(echo $line | awk '{print $4 " " $5 " " $6 " " $7 " " $8}')
  print_code "$date" "$color" "$occasion"
done
