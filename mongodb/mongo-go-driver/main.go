package main

//导入
import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"os"
	"time"
)

type Howie struct {
	//struct里面获取ObjectID
	HowieId     primitive.ObjectID `bson:"_id"`
	Name        string
	Pwd         string
	Age         int64
	CreateTime  int64
	ExpiredTime time.Time
}

func main() {
	TestMongo("mongodb://192.168.2.28:27017")
}

func TestMongo(url string) {
	var (
		err             error
		client          *mongo.Client
		collection      *mongo.Collection
		insertOneRes    *mongo.InsertOneResult
		insertManyRes   *mongo.InsertManyResult
		delRes          *mongo.DeleteResult
		updateRes       *mongo.UpdateResult
		cursor          *mongo.Cursor
		howieArray      = GetHowieArray()
		howie           Howie
		howieArrayEmpty []Howie
		size            int64
	)
	want, err := readpref.New(readpref.SecondaryMode) //表示只使用辅助节点
	if err != nil {
		checkErr(err)
	}
	wc := writeconcern.New(writeconcern.WMajority())
	readconcern.Majority()
	//链接mongo服务
	opt := options.Client().ApplyURI(url)
	opt.SetLocalThreshold(3 * time.Second)     //只使用与mongo操作耗时小于3秒的
	opt.SetMaxConnIdleTime(5 * time.Second)    //指定连接可以保持空闲的最大毫秒数
	opt.SetMaxPoolSize(200)                    //使用最大的连接数
	opt.SetReadPreference(want)                //表示只使用辅助节点
	opt.SetReadConcern(readconcern.Majority()) //指定查询应返回实例的最新数据确认为，已写入副本集中的大多数成员
	opt.SetWriteConcern(wc)                    //请求确认写操作传播到大多数mongod实例
	if client, err = mongo.Connect(getContext(), opt); err != nil {
		checkErr(err)
	}
	//StudySession(client)
	//判断服务是否可用
	if err = client.Ping(getContext(), readpref.Primary()); err != nil {
		checkErr(err)
	}
	//选择数据库和集合
	collection = client.Database("testing_base").Collection("howie")
	//集合数据自动过期
	/*k := mongo.IndexModel{
		Keys:    bsonx.Doc{{"expiredtime", bsonx.Int32(1)}},//创建索引
		Options: options.Index().SetExpireAfterSeconds(1 * 60),//创建数据过期
	}
	_, err = collection.Indexes().CreateOne(getContext(), k)
	checkErr(err)*/

	//删除这个集合
	collection.Drop(getContext())

	//插入一条数据
	if insertOneRes, err = collection.InsertOne(getContext(), howieArray[0]); err != nil {
		checkErr(err)
	}

	fmt.Printf("InsertOne插入的消息ID:%v\n", insertOneRes.InsertedID)

	//批量插入数据
	if insertManyRes, err = collection.InsertMany(getContext(), howieArray[1:]); err != nil {
		checkErr(err)
	}
	fmt.Printf("InsertMany插入的消息ID:%v\n", insertManyRes.InsertedIDs)
	var Dinfo = make(map[string]interface{})
	err = collection.FindOne(getContext(), bson.D{{"name", "howie_2"}, {"age", 11}}).Decode(&Dinfo)
	fmt.Println(Dinfo)
	fmt.Println("_id", Dinfo["_id"])

	//查询单条数据
	if err = collection.FindOne(getContext(), bson.D{{"name", "howie_2"}, {"age", 11}}).Decode(&howie); err != nil {
		checkErr(err)
	}
	fmt.Printf("FindOne查询到的数据:%v\n", howie)

	//查询单条数据后删除该数据
	if err = collection.FindOneAndDelete(getContext(), bson.D{{"name", "howie_3"}}).Decode(&howie); err != nil {
		checkErr(err)
	}
	fmt.Printf("FindOneAndDelete查询到的数据:%v\n", howie)

	//查询单条数据后修改该数据
	if err = collection.FindOneAndUpdate(getContext(), bson.D{{"name", "howie_4"}}, bson.M{"$set": bson.M{"name": "这条数据我需要修改了"}}).Decode(&howie); err != nil {
		checkErr(err)
	}
	fmt.Printf("FindOneAndUpdate查询到的数据:%v\n", howie)

	//查询单条数据后替换该数据(以前的数据全部清空)
	if err = collection.FindOneAndReplace(getContext(), bson.D{{"name", "howie_5"}}, bson.M{"hero": "这条数据我替换了"}).Decode(&howie); err != nil {
		checkErr(err)
	}

	fmt.Printf("FindOneAndReplace查询到的数据:%v\n", howie)
	//一次查询多条数据
	// 查询createtime>=3
	// 限制取2条
	// createtime从大到小排序的数据
	if cursor, err = collection.Find(getContext(), bson.M{"createtime": bson.M{"$gte": 2}}, options.Find().SetLimit(2), options.Find().SetSort(bson.M{"createtime": -1})); err != nil {
		checkErr(err)
	}
	if err = cursor.Err(); err != nil {
		checkErr(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		if err = cursor.Decode(&howie); err != nil {
			checkErr(err)
		}
		howieArrayEmpty = append(howieArrayEmpty, howie)
	}
	for _, v := range howieArrayEmpty {
		fmt.Printf("Find查询到的数据ObejectId值%s 值:%v\n", v.HowieId.Hex(), v)
	}
	//查询集合里面有多少数据
	if size, err = collection.CountDocuments(getContext(), bson.D{}); err != nil {
		checkErr(err)
	}
	fmt.Printf("Count里面有多少条数据:%d\n", size)

	//查询集合里面有多少数据(查询createtime>=3的数据)
	if size, err = collection.CountDocuments(getContext(), bson.M{"createtime": bson.M{"$gte": 3}}); err != nil {
		checkErr(err)
	}
	fmt.Printf("Count里面有多少条数据:%d\n", size)

	//修改一条数据
	if updateRes, err = collection.UpdateOne(getContext(), bson.M{"name": "howie_2"}, bson.M{"$set": bson.M{"name": "我要改了他的名字"}}); err != nil {
		checkErr(err)
	}
	fmt.Printf("UpdateOne的数据:%d\n", updateRes)

	//修改多条数据
	if updateRes, err = collection.UpdateMany(getContext(), bson.M{"createtime": bson.M{"$gte": 3}}, bson.M{"$set": bson.M{"name": "我要批量改了他的名字"}}); err != nil {
		checkErr(err)
	}

	fmt.Printf("UpdateMany的数据:%d\n", updateRes)
	//删除一条数据
	if delRes, err = collection.DeleteOne(getContext(), bson.M{"name": "howie_1"}); err != nil {
		checkErr(err)
	}
	fmt.Printf("DeleteOne删除了多少条数据:%d\n", delRes.DeletedCount)

	//删除多条数据
	if delRes, err = collection.DeleteMany(getContext(), bson.M{"createtime": bson.M{"$gte": 7}}); err != nil {
		checkErr(err)
	}
	fmt.Printf("DeleteMany删除了多少条数据:%d\n", delRes.DeletedCount)

}

func StudySession(client *mongo.Client) {
	client.UseSession(getContext(), func(sctx mongo.SessionContext) error {
		err := sctx.StartTransaction(options.Transaction().
			SetReadConcern(readconcern.Snapshot()).
			SetWriteConcern(writeconcern.New(writeconcern.WMajority())),
		)
		if err != nil {
			return err
		}
		_, err = client.Database("aa").Collection("bb").UpdateOne(sctx, bson.D{{"aa", "1"}}, bson.D{{"$set", bson.D{{"status", "123123"}}}})
		if err != nil {
			_ = sctx.AbortTransaction(sctx)
			return err
		}
		_, err = client.Database("cc").Collection("dd").InsertOne(sctx, bson.D{{"employee", 3}})
		if err != nil {
			_ = sctx.AbortTransaction(sctx)
			return err
		}
		for {
			err = sctx.CommitTransaction(sctx)
			switch e := err.(type) {
			case nil:
				return nil
			case mongo.CommandError:
				if e.HasErrorLabel("UnknownTransactionCommitResult") {
					continue
				}
				return e
			default:
				return e
			}
		}
	})
}

func checkErr(err error) {
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("没有查到数据")
			os.Exit(0)
		} else {
			fmt.Println(err)
			os.Exit(0)
		}

	}
}

func getContext() (ctx context.Context) {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	return
}

func GetHowieArray() (data []interface{}) {
	var i int64
	for i = 0; i <= 10; i++ {
		data = append(data, Howie{
			HowieId:     primitive.NewObjectID(),
			Name:        fmt.Sprintf("howie_%d", i+1),
			Pwd:         fmt.Sprintf("pwd_%d", i+1),
			Age:         i + 10,
			CreateTime:  time.Now().Unix(),
			ExpiredTime: time.Now(),
		})
	}
	return
}
