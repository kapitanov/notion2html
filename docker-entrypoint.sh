#!/bin/sh
set -e

cd /opt/notion2html

OUTPUT_DIR=${OUTPUT_DIR:-/out}
TIMER_PERIOD=${TIMER_PERIOD:-10m}

case "$1" in
generate)
	shift
	/opt/notion2html/notion2html $*
	;;
watch)
	shift
	/opt/notion2html/notion2html $*
	;;
*)
	/opt/notion2html/notion2html watch --output "$OUTPUT_DIR" --token "$NOTION_API_TOKEN" --period "$TIMER_PERIOD"
	;;
esac
