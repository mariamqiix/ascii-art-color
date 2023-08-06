#!/bin/bash

echo "go run . --color red "banana""
go run . --color red "banana" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . --color=red \"hello world\""
go run . --color=red "hello world" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . --color=green \"1 + 1 = 2\""
go run . --color=green "1 + 1 = 2" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . --color=yellow \"(%&) ??\""
go run . --color=yellow "(%&) ??" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . --color=yellow  ello \"Hello\""
go run . --color=yellow  ello "Hello" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . --color=yellow  ello \"Hello\""
go run .  --color=yellow  e \"Hello\" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . --color=yellow  ello \"Hello\""
go run . --color=yellow  el \"Hello\" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . -color=orange GuYs "HeY GuYs""
go run . --color=orange GuYs "HeY GuYs" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . --color=blue B 'RGB()'"
go run . --color=blue B 'RGB()' | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . --color=cyan B \"ReBoOt\""
go run . --color=cyan B "ReBoOt" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . --color=purple ! \"Reboot 01!\""
go run . --color=purple ! "Reboot 01!" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . --color=purple R \"#Reboot 01!!!\""
go run . --color=purple R "#Reboot 01!!!" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . --color=purple R \"#R eb oot &&&&01!!!\""
go run . --color=purple R "#R eb oot &&&&01!!!" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . --color=purple R --color=cyan ! \"#R eb oot &&&&01!!!\""
go run . --color=purple R --color=cyan ! "#R eb oot &&&&01!!!" | cat -e 
echo
echo "---------------------------------------------------------"