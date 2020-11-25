docker run --name es -d  -p 9200:9200 -p 9300:9300 -e  "discovery.type=single-node" elasticsearch:7.2.0


unzip elasticsearch-analysis-ik-7.2.0.zip -d ik


docker cp ik es:/usr/share/elasticsearch/plugins/ik/

docker restart es