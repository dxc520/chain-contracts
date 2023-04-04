cd ../
docker build -t eth-security-toolbox:v0.1 . -f test/Dockerfile
docker run --rm eth-security-toolbox:v0.1
