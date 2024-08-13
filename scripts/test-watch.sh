#!/usr/bin/env sh

run_test() {
		go test -v -race -buildvcs ./... | \
			sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | \
			sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/'' | \
			sed ''/RUN/s//$(printf "\033[33mRUN\033[0m")/''
}

echo "[$(date +%Y-%m-%d_%H-%m_%S)]"
echo "========================================"

run_test

echo ''

while true; do
	inotifywait -qq -r -e create,close_write,modify,move,delete ./ &&
		echo "[$(date +%Y-%m-%d_%H-%m_%S)]" &&
		echo "========================================" &&
		run_test

	echo ''
done
