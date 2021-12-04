# Ukaska

**Ukaska** *(eng. Pointer)* - is docker image with [mongodb](https://www.mongodb.com/), that has builtin backup to [Telegram](https://telegram.org/) 

## Getting Started

### Build with

* [Docker](https://hub.docker.com/)
* Mongo DB docker [image](https://hub.docker.com/_/mongo)
* [Golang](https://go.dev)
* [Jobber](https://github.com/dshearer/jobber) as cron alternative


### Building manualy
```sh
docker build -t ukaska .
```

### Running docker Image

1. Create bot with help of [@BotFather](https://t.me/BotFather) and set it to 
   ```sh
   export BOT_TOKEN="123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11"
   ```
2. Get channel ID with [this](https://gist.github.com/mraaroncruz/e76d19f7d61d59419002db54030ebe35) instruction
   ```sh
   export CHANNEL_ID="-100123456"
   ```
3. Set backup time for jobber *(in example below backup set on `00:00` daily)*
    > `❗ Keep in mind, that Jobber has another format, than cron, it starts with seconds `
    ```sh
    export JOBBER_STRING="0 0 0 * * *"
    ```

4. Set mongodb URL for backup *(e.g. for setting specific database)*
    ```sh
    export DB_URL="mongodb://localhost:27017/db"
    ```
5. Set mongodb collection names (use JSON serialization) 
    ```sh
    export COLLECTION_NAMES='["HamstersCollection"]'
    ```
6. Run docker container
    ```sh
    docker run -dp 27017:27017 --env DB_URL=${DB_URL} --env JOBBER_STRING='${JOBBER_STRING}' --env CHANNEL_ID=${CHANNEL_ID} --env BOT_TOKEN=${BOT_TOKEN} --env COLLECTION_NAMES='${COLLECTION_NAMES}' ukaska
    ```