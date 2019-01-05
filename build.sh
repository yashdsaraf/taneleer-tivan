#!/bin/bash

ndk_standalone_toolchain=~/Toolchains/ndk-arm-27-oreo
out_bin=bin/tivan

export PATH=$PATH:$ndk_standalone_toolchain/bin

for os in android darwin linux windows
do
	export GOOS=$os
	should_strip=false
	cc=gcc
	strip=strip
	case $os in
		'android')
			export GOARCH=arm
			export CGO_ENABLED=1
			export GOARM=7
			cc=arm-linux-androideabi-gcc
			strip=arm-linux-androideabi-strip
			should_strip=true
		;;
		'windows')
			export GOARCH=386
		;;
		'darwin')
			export GOARCH=amd64
		;;
		'linux')
			export GOARCH=amd64
			should_strip=true
		;;
	esac

	mkdir -p bin
	CC=$cc go build -o $out_bin.$GOOS tivan.go
	chmod +x bin/*
	$should_strip && $strip -s $out_bin.$GOOS
done
