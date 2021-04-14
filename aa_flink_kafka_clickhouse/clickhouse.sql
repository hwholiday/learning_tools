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


-- 留存分析
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

-- 路径分析
create table test.user_act
(
    uid UInt32,
    act String,
    ds  DateTime('Asia/Shanghai')
)ENGINE = MergeTree()
PARTITION BY toYYYYMM(ds)
order by (uid,act)
SETTINGS index_granularity = 8192;

INSERT INTO test.user_act
VALUES (222, 'login', '2021-04-13 10:01:10'),
       (222, 'blog', '2021-04-13 10:03:10'),
       (222, 'msg', '2021-04-13 10:10:10'),
       (333, 'login', '2021-04-13 10:01:10'),
       (333, 'msg', '2021-04-13 10:02:10');

-- 关键路径分析
-- 查询 2021-04-13 login 后大于 120 就开始 msg 的用户
select count(1) as userNum, sum(cn) as num
from (
         SELECT uid,
                sequenceCount('(?1)(?t>120)(?2)') (toDateTime(ds), act = 'login', act = 'msg') AS cn
         FROM test.user_act
         WHERE toDate(ds) = '2021-04-13'
         GROUP BY uid
     )
-- userNum 在2021-04-13 满足 即login也msg的用户数 /  num 是满足login 后大于 120 就开始 msg 的用户
--┌─userNum─┬─num─┐
--│       2 │   1 │
--└─────────┴─────┘

SELECT
    result_chain,
    uniqCombined(user_id) AS user_count
FROM (
         WITH
             toDateTime(maxIf(time, act = '会员支付成功')) AS end_event_maxt,  #以终点事件时间作为路径查找结束时间
             arrayCompact(arraySort(  #对事件按照时间维度排序后进行相邻去重
             x -> x.1,
             arrayFilter(  #根据end_event_maxt筛选出所有满足条件的事件 并按照<时间, <事件名, 页面名>>结构返回
             x -> x.1 <= end_event_maxt,
             groupArray((toDateTime(time), (act, page_name)))
             )
             )) AS sorted_events,
             arrayEnumerate(sorted_events) AS event_idxs,  #或取事件链的下标掩码序列，后面在对事件切割时会用到
             arrayFilter(  #将目标事件或当前事件与上一个事件间隔10分钟的数据为切割点
             (x, y, z) -> z.1 <= end_event_maxt AND (z.2.1 = '会员支付成功' OR y > 600),
             event_idxs,
             arrayDifference(sorted_events.1),
             sorted_events
             ) AS gap_idxs,
             arrayMap(x -> x + 1, gap_idxs) AS gap_idxs_,  #如果不加1的话上一个事件链的结尾事件会成为下个事件链的开始事件
             arrayMap(x -> if(has(gap_idxs_, x), 1, 0), event_idxs) AS gap_masks,  #标记切割点
             arraySplit((x, y) -> y, sorted_events, gap_masks) AS split_events  #把用户的访问数据切割成多个事件链
         SELECT
             user_id,
             arrayJoin(split_events) AS event_chain_,
             arrayCompact(event_chain_.2) AS event_chain,  #相邻去重
             hasAll(event_chain, [('pay_button_click', '会员购买页')]) AS has_midway_hit,
             arrayStringConcat(arrayMap(
             x -> concat(x.1, '#', x.2),
             event_chain
             ), ' -> ') AS result_chain  #用户访问路径字符串
         FROM (
             SELECT time,act,page_name,u_i as user_id
             FROM app.scene_tracker
             WHERE toDate(time) >= '2020-09-30' AND toDate(time) <= '2020-10-02'
             AND user_id IN (10266,10022,10339,10030)  #指定要分析的用户群
             )
         GROUP BY user_id
         HAVING length(event_chain) > 1
     )
WHERE event_chain[length(event_chain)].1 = '会员支付成功'  #事件链最后一个事件必须是目标事件
  AND has_midway_hit = 1   #必须包含途经点
GROUP BY result_chain
ORDER BY user_count DESC LIMIT 20;




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

SELECT level_index,
       count(1)
FROM (
         SELECT user_id,
                arrayWithConstant(level, 1)       AS levels,
                arrayJoin(arrayEnumerate(levels)) AS level_index
         FROM (
                  SELECT user_id,
                         windowFunnel(1800) (time, event_type = '浏览', event_type = '点击', event_type = '下单', event_type = '支付') AS level
                  FROM (
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