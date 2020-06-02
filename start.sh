lsof -ntP -i:7002|xargs kill -9
nohup ./httpscript &