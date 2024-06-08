#!/bin/bash

REPO_URL="https://github.com/silveiralexf/goflat"
printf "# CHANGELOG\n\n" >CHANGELOG.md
git --no-pager log \
	--no-merges \
	--format="%cd" \
	--date=short |
	sort -u -r |
	while read -r DATE; do
		echo
		echo "### [${DATE}]"
		GIT_PAGER=cat git log --no-merges --format=" * [[%h]]($REPO_URL/%H) %s (%aE)%n%n" --since="$DATE 00:00" --until="$DATE 23:59"
	done >>CHANGELOG.md
