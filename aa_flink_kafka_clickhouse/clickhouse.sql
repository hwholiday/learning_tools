-- 客户端程序
clickhouse-client -h 172.12.17.161 --port 9090 -u default  --password 123456



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

-- 漏斗分析
-- 分析"2020-01-02"这天 路径为“浏览->点击->下单->支付”的转化情况。

CREATE TABLE test.action
(
    `uid`        Int32,
    `event_type` String,
    `time`       DateTime
)ENGINE = MergeTree()
PARTITION BY uid
ORDER BY xxHash32(uid)
SAMPLE BY xxHash32(uid)
SETTINGS index_granularity = 8192

INSERT INTO action
VALUES (1,'浏览','2020-01-02 11:00:00'),
       (1,'点击','2020-01-02 11:10:00'),
       (1,'下单','2020-01-02 11:20:00'),
       (1,'支付','2020-01-02 11:30:00'),
       (2,'下单','2020-01-02 11:00:00'),
       (2,'支付','2020-01-02 11:10:00'),
       (1,'浏览','2020-01-02 11:00:00'),
       (3,'浏览','2020-01-02 11:20:00'),
       (3,'点击','2020-01-02 12:00:00'),
       (4,'浏览','2020-01-02 11:50:00'),
       (4,'点击','2020-01-02 12:00:00'),
       (5,'浏览','2020-01-02 11:50:00'),
       (5,'点击','2020-01-02 12:00:00'),
       (5,'下单','2020-01-02 11:10:00'),
       (6,'浏览','2020-01-02 11:50:00'),
       (6,'点击','2020-01-02 12:00:00'),
       (6,'下单','2020-01-02 12:10:00');

SELECT
    level_index,
    count(1)
FROM
    (
        SELECT
            user_id,
            arrayWithConstant(level, 1) AS levels,
            arrayJoin(arrayEnumerate(levels)) AS level_index
        FROM
            (
                SELECT
                    user_id,
                    windowFunnel(1800)(time, event_type = '浏览', event_type = '点击', event_type = '下单', event_type = '支付') AS level
                FROM
                    (
                        SELECT
                            time,
                            event_type,
                            uid AS user_id
                        FROM test.action
                        WHERE toDate(time) = '2020-01-02'
                    )
                GROUP BY user_id
            )
    )
GROUP BY level_index
ORDER BY level_index ASC

--┌─level_index─┬─count(1)─┐
--│           1 │        5 │
--│           2 │        4 │
--│           3 │        2 │
--│           4 │        1 │
--└─────────────┴──────────┘