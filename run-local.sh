if [ -e .env ]
    go build
    ./khoros-community-cli debug
else
    echo "Env file not supplied"
fi