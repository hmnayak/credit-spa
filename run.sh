trap "kill 0" EXIT

go run main.go &
npm --prefix web run dev &

wait
