#!/usr/bin/env bash

#################################################################
# Generate tests with "gotests -all" for all Go source files.
# Generated files will contain dummy test funcs. 
#################################################################

# Var for source files
SOURCES=*.go

# Print usage message end exit
usage()
{
   echo "Usage : test_gen.sh -f|--file [FILE=""all""|FILENAME]" 
}

# Generate all
generate()
{
	gotests -all -w "$1"
}

# Messages
file_will()
{
	echo "Go tests for "$1" will be generated..."
}

file_done()
{
	echo "File "$1" done"
}

# ClI args processing

# if there is no args provided, print usage and exit
if [ "$1" = "" ]
then
    usage
    exit
fi

# args parsing
while [ "$1" != "" ]; do
    case $1 in
        -f | --file )
        FILE="$2"
        shift
        shift
        file_will "$FILE"
        ;;
        * )
        usage
        exit 1
    esac
done

# Move to package folder
cd ../pwgator

# Generating tests
 if [ "$FILE" = "all" ]
then
	for s in $SOURCES
	do
		generate "$s"
		file_done "$s"
	done
fi

generate "$FILE"
file_done "$FILE"