docker build -t testimg .
docker run --rm -p 9200:9200 testimg
docker run --rm -p 9200:9200 -e HOGE=hogge testimg
docker run --rm -p 9200:9200 -e HOGE=hogge -v $(pwd)/logs:/var/logs testimg
