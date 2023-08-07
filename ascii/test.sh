#!/bin/bash

echo "go run . --color red "banana""
go run . --color red "banana" 
echo
echo "---------------------------------------------------------"

echo "go run . --color=red \"hello world\""
go run . --color=red "hello world" 
echo
echo "---------------------------------------------------------"

echo "go run . --color=green \"1 + 1 = 2\""
go run . --color=green "1 + 1 = 2"  
echo
echo "---------------------------------------------------------"

echo "go run . --color=yellow \"(%&) ??\""
go run . --color=yellow "(%&) ??" 
echo
echo "---------------------------------------------------------"

echo "go run . --color=yellow  ello \"Hello\""
go run . --color=yellow  ello "Hello" 
echo
echo "---------------------------------------------------------"

echo "go run . --color=yellow  e \"Hello\""
go run .  --color=yellow  e \"Hello\" 
echo
echo "---------------------------------------------------------"

echo "go run . --color=yellow  el \"Hello\""
go run . --color=yellow  el \"Hello\" 
echo
echo "---------------------------------------------------------"

echo "go run . --color=orange GuYs \"HeY GuYs\""
go run . --color=orange GuYs "HeY GuYs"  
echo
echo "---------------------------------------------------------"

echo "go run . --color=blue B 'RGB()'"
go run . --color=blue B 'RGB()'  
echo
echo "---------------------------------------------------------"

echo "go run . --color=cyan B \"ReBoOt\""
go run . --color=cyan B "ReBoOt"  
echo
echo "---------------------------------------------------------"

echo "go run . --color=purple ! \"Reboot 01!\""
go run . --color=purple ! "Reboot 01!" 
echo
echo "---------------------------------------------------------"

echo "go run . --color=purple R \"#Reboot 01!!!\""
go run . --color=purple R "#Reboot 01!!!" 
echo
echo "---------------------------------------------------------"

echo "go run . --color=purple R \"#R eb oot &&&&01!!!\""
go run . --color=purple R "#R eb oot &&&&01!!!"  
echo
echo "---------------------------------------------------------"

echo "go run . --color=purple R --color=cyan ! \"#R eb oot &&&&01!!!\""
go run . --color=purple R --color=cyan ! "#R eb oot &&&&01!!!"  
echo
echo "---------------------------------------------------------"