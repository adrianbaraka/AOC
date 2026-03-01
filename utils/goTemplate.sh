#!/usr/bin/env bash

printUsage() {
	local code
	code=$1
	cat <<- EOF
	Usage $0 -d <day> directory

	Initialize a directory with the go template
	    -d (required) the day eg 1, 2
	    <directory> (required) the directory to initialize
	EOF
	exit "$code"

}

main(){
	while getopts d: opt; do
		case "$opt" in
			d) day="$OPTARG";;
			*) exit;;
		esac
	done

	shift $((OPTIND - 1))

	dir=$1
	[[ -z $dir ]] && printUsage 1



	if ! mkdir -pv "$dir"; then
		exit 1
	fi
	# copy the template
	cp -v  "starterCode/go.text" "${dir}/aoc${day}.go" || exit 1
	# done in a subshell
	(
		cd "$dir" || exit 1
		go mod init "aoc$day"
		go mod tidy

		# input files
		touch input.txt
		touch test.txt
	)
}

main "$@"