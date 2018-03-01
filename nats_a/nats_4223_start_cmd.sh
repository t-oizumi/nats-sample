#!/bin/bash

nats-streaming-server -store file -dir datastore -ft_group "ft" -cluster  nats://nats_a:6223 -routes nats://nats_a:6222 -p 4223

