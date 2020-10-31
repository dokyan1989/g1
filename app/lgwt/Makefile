name=new_app
gen:
	mkdir ./$(name) ./$(name)/cmd ./$(name)/internal
	touch ./$(name)/cmd/main.go
	touch ./$(name)/internal/.keep

run:
	go run ./$(name)/cmd

.PHONY:
	gen run