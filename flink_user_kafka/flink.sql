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
clickhouse
-client -h 172.12.17.161 --port 9090 -u default  --password 123456



-- 学习 ck DateTime
-- 留存留存
create
database test;
-- ReplicatedMergeTree()
drop table test.t_event;
create table test.t_event
(
    uid   UInt32,
    event String,
    ds    DateTime('Asia/Shanghai')
)ENGINE = MergeTree()
PARTITION BY toYYYYMM(ds)
order by (uid,event)
SETTINGS index_granularity = 8192;

INSERT INTO test.t_event
VALUES (10001, 'login', '2020-08-01'),
       (10001, 'login', '2020-08-08'),
       (10001, 'login', '2020-08-09'),
       (10001, 'login', '2020-08-10'),
       (10001, 'login', '2020-08-12'),
       (10001, 'login', '2020-08-13'),
       (10001, 'login', '2020-08-14'),
       (10001, 'login', '2020-08-15'),
       (10001, 'login', '2020-08-16'),
       (10001, 'login', '2020-08-17'),
       (10001, 'login', '2020-08-18'),
       (10001, 'login', '2020-08-20'),
       (10001, 'login', '2020-08-22'),
       (10001, 'login', '2020-08-23'),
       (10001, 'login', '2020-08-24'),
       (10002, 'login', '2020-08-20'),
       (10002, 'login', '2020-08-22'),
       (10002, 'login', '2020-08-23'),
       (10002, 'login', '2020-08-01'),
       (10002, 'login', '2020-08-11'),
       (10002, 'login', '2020-08-12'),
       (10002, 'login', '2020-08-13'),
       (10002, 'login', '2020-08-20'),
       (10002, 'login', '2020-08-15'),
       (10002, 'login', '2020-08-30'),
       (10002, 'login', '2020-08-20'),
       (10002, 'login', '2020-08-01'),
       (10002, 'login', '2020-08-06'),
       (10002, 'login', '2020-08-24'),
       (10003, 'login', '2020-08-05'),
       (10003, 'login', '2020-08-08'),
       (10003, 'login', '2020-08-09'),
       (10003, 'login', '2020-08-10'),
       (10003, 'login', '2020-08-11'),
       (10003, 'login', '2020-08-13'),
       (10003, 'login', '2020-08-15'),
       (10003, 'login', '2020-08-16'),
       (10003, 'login', '2020-08-18'),
       (10003, 'login', '2020-08-20'),
       (10003, 'login', '2020-08-01'),
       (10003, 'login', '2020-08-21'),
       (10003, 'login', '2020-08-22'),
       (10003, 'login', '2020-08-24'),
       (10003, 'login', '2020-08-26'),
       (10003, 'login', '2020-08-25'),
       (10003, 'login', '2020-08-27'),
       (10003, 'login', '2020-08-28'),
       (10003, 'login', '2020-08-29'),
       (10003, 'login', '2020-08-30'),
       (10004, 'login', '2020-08-01'),
       (10004, 'login', '2020-08-02'),
       (10004, 'login', '2020-08-03'),
       (10004, 'login', '2020-08-04'),
       (10004, 'login', '2020-08-05'),
       (10004, 'login', '2020-08-08'),
       (10004, 'login', '2020-08-09'),
       (10004, 'login', '2020-08-10'),
       (10004, 'login', '2020-08-11'),
       (10004, 'login', '2020-08-14'),
       (10004, 'login', '2020-08-15'),
       (10004, 'login', '2020-08-16'),
       (10004, 'login', '2020-08-17'),
       (10004, 'login', '2020-08-19'),
       (10004, 'login', '2020-08-20'),
       (10004, 'login', '2020-08-21'),
       (10004, 'login', '2020-08-22'),
       (10004, 'login', '2020-08-23'),
       (10004, 'login', '2020-08-24'),
       (10004, 'login', '2020-08-23'),
       (10004, 'login', '2020-08-23'),
       (10004, 'login', '2020-08-25'),
       (10004, 'login', '2020-08-27'),
       (10004, 'login', '2020-08-30');

SELECT *
FROM test.t_event;

SELECT uid, toDate(ds), toYYYYMM(ds), toYYYYMMDD(ds)
FROM test.t_event;

select uid, toDate(ds) as t
from test.t_event
where t >= '2020-08-01'
  and t <= addDays(toDate('2020-08-01'), 7)
  and uid = 10001


-- 留存
SELECT toDate('2020-08-01')  AS ds,
       SUM(r[1])             AS activeAccountNum,
       SUM(r[2]) / SUM(r[1]) AS `次留`,
       SUM(r[3]) / SUM(r[1]) AS `3留`,
       SUM(r[4]) / SUM(r[1]) AS `7留`,
       SUM(r[5]) / SUM(r[1]) AS `14留`,
       SUM(r[6]) / SUM(r[1]) AS `30留`
FROM (
         --找到2020-08-01活跃用户在第2、3、6、13、29日的登录情况，1/0 => 登录/未登录
         WITH toDate('2020-08-01') AS tt
         SELECT uid,
             retention(
             toDate(ds) = tt and event = 'login',
             toDate(subtractDays(ds, 1)) = tt and event = 'login',
             toDate(subtractDays(ds, 2)) = tt and event = 'login',
             toDate(subtractDays(ds, 6)) = tt and event = 'login',
             toDate(subtractDays(ds, 13)) = tt and event = 'login',
             toDate(subtractDays(ds, 29)) = tt and event = 'login') AS r
         FROM test.t_event
         WHERE (toDate(ds) >= '2020-08-01') AND (toDate(ds) <= addDays(toDate('2020-08-01'), 29))
         GROUP BY uid
     )
--┌─────────ds─┬─activeAccountNum─┬─次留─┬──3留─┬─7留─┬─14留─┬─30留─┐
--│ 2020-08-01 │                4 │ 0.25 │ 0.25 │   0 │  0.5 │ 0.75 │
--└────────────┴──────────────────┴──────┴──────┴─────┴──────┴──────┘



