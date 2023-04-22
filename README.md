# iam

[In progress]

```
                            +---------+
                        +-->| CLient  |
                        |   +---------+
        +---------+     |      :3000             +---------+
        |  Proxy  +-----+                   +--->|   IAM   |
        +---------+     |                   |    +---------+
           :5000        |   +---------+     |       :7777
                        +-->| Gateway |-----+
                            +---------+     |
                               :7000        |    +---------+
                                            +--->|  Core   |
                                                 +---------+
                                                    :8000
```

```
docker-compose build
docker-compose up -d
docker-compose stop
docker stop $(docker ps -aq)
```