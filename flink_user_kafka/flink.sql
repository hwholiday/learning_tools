CREATE TABLE msg_event
(
    biz_tag     STRING,
    uid         INT,
    event       STRING,
    tag         STRING,
    create_time BIGINT,
    ts AS TO_TIMESTAMP(FROM_UNIXTIME(create_time / 1000, 'yyyy-MM-dd HH:mm:ss')),
    WATERMARK FOR ts AS ts - INTERVAL '5' SECOND
) WITH (
      'connector' = 'kafka',
      'topic' = 'msg_event',
      'properties.bootstrap.servers' = '172.12.17.161:9092',
      'format' = 'json'
      );

CREATE TABLE msg_event_sk
(
    biz_tag     STRING,
    event       STRING,
    tag         STRING,
    create_time BIGINT,
    uid         BIGINT,
    num         BIGINT
) WITH (
      'connector' = 'clickhouse',
      'url' = 'clickhouse://172.12.17.161:8123',
      'username' = 'default',
      'password' = '123456',
      'database-name' = 'msg',
      'table-name' = 'msg_event_sk',
      'sink.batch-size' = '1',
      'sink.flush-interval' = '1',
      'sink.max-retries' = '3',
      'sink.partition-strategy' = 'hash',
      'sink.partition-key' = 'biz_tag',
      'sink.ignore-delete' = 'true'
      );


create table msg.msg_event_sk
(
    biz_tag     String,
    event       String,
    tag         String,
    create_time Int64,
    uid         Int32,
    num         Int64
)ENGINE = MergeTree()
order by biz_tag;


INSERT INTO msg_event_sk
SELECT FIRST_VALUE(biz_tag)     as biz_tag,
       FIRST_VALUE(event)       as event,
       FIRST_VALUE(tag)         as tag,
       FIRST_VALUE(create_time) as create_time,
       FIRST_VALUE(uid)         as uid,
       COUNT(event)             as num
FROM msg_event
GROUP BY TUMBLE(ts, INTERVAL '1' MINUTE);


-- 测试数据
INSERT INTO msg_event_sk (biz_tag, event, tag, create_time, uid, num)
VALUES ('22', '1', '1', 1, 1, 1);

-- 客户端程序
clickhouse -client -h 172.12.17.161 --port 9090 -u default  --password 123456



-- 学习 ck DateTime
create database test;
-- ReplicatedMergeTree()
create table  test.dt (
    ds DateTime('Asia/Shanghai'),
    uid UInt32,
    event String
)ENGINE = MergeTree()
PARTITION BY toYYYYMM(ds)
order by (uid,event)
SETTINGS index_granularity = 8192;

INSERT INTO test.dt VALUES (1546300800, 1, 'login'), ('2019-01-01 00:00:00', 2,'game');
SELECT * FROM test.dt;
SELECT toYYYYMM(ds),uid FROM test.dt;
SELECT toYYYYMMDD(ds),uid FROM test.dt;