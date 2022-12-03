#!/bin/sh
set -e

cd /opt/notion2html

case "$1" in
generate)
	shift
	/opt/notion2html/notion2html generate $*
	;;
watch)
	shift
	/opt/notion2html/notion2html watch $*
	;;
*)
	OUTPUT_DIR=${OUTPUT_DIR:-/out}
	TIMER_PERIOD=${TIMER_PERIOD:-10m}
	/opt/notion2html/notion2html watch --output "$OUTPUT_DIR" --token "$NOTION_API_TOKEN" --period "$TIMER_PERIOD"
	;;
esac
