package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type User struct {
	Id       bson.ObjectId `bson:"_id"`
	Name     string        `bson:"name"`
	PassWord string        `bson:"pass_word"`
	Age      int           `bson:"age"`
}

var db *mgo.Session

func main() {

	db, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	db.SetMode(mgo.Eventual, true)
	db.SetPoolLimit(2000)
	db.SetSocketTimeout(3 * time.Second)
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	c := db.DB("howie").C("person")

	//插入
	/*c.Insert(&User{
		Id:       bson.NewObjectId(),
		Name:     "JK_CHENG",
		PassWord: "123132",
		Age: 2,
	}, &User{
		Id:       bson.NewObjectId(),
		Name:     "JK_WEI",
		PassWord: "qwer",
		Age: 5,
	}, &User{
		Id:       bson.NewObjectId(),
		Name:     "JK_HE",
		PassWord: "6666",
		Age: 7,
	})*/
	var users []User
	c.Find(nil).All(&users) //查询全部数据
	log.Println(users)

	c.FindId(users[0].Id).All(&users) //通过ID查询
	log.Println(users)

	c.Find(bson.M{"name": "JK_WEI"}).All(&users) //单条件查询(=)
	log.Println(users)

	c.Find(bson.M{"name": bson.M{"$ne": "JK_WEI"}}).All(&users) //单条件查询(!=)
	log.Println(users)

	c.Find(bson.M{"age": bson.M{"$gt": 5}}).All(&users) //单条件查询(>)
	log.Println(users)

	c.Find(bson.M{"age": bson.M{"$gte": 5}}).All(&users) //单条件查询(>=)
	log.Println(users)

	c.Find(bson.M{"age": bson.M{"$lt": 5}}).All(&users) //单条件查询(<)
	log.Println(users)

	c.Find(bson.M{"age": bson.M{"$lte": 5}}).All(&users) //单条件查询(<=)
	log.Println(users)

	/*c.Find(bson.M{"name": bson.M{"$in": []string{"JK_WEI", "JK_HE"}}}).All(&users) //单条件查询(in)
	log.Println(users)

	c.Find(bson.M{"$or": []bson.M{bson.M{"name": "JK_WEI"}, bson.M{"age": 7}}}).All(&users) //多条件查询(or)
	log.Println(users)

	c.Update(bson.M{"_id": users[0].Id}, bson.M{"$set": bson.M{"name": "JK_HOWIE", "age": 61}}) //修改字段的值($set)

	c.FindId(users[0].Id).All(&users)
	log.Println(users)

	c.Find(bson.M{"name": "JK_CHENG", "age": 66}).All(&users) //多条件查询(and)
	log.Println(users)

	c.Update(bson.M{"_id": users[0].Id}, bson.M{"$inc": bson.M{"age": -6,}}) //字段增加值($inc)

	c.FindId(users[0].Id).All(&users)
	log.Println(users)*/

	//c.Update(bson.M{"_id": users[0].Id}, bson.M{"$push": bson.M{"interests": "PHP"}}) //从数组中增加一个元素($push)

	c.Update(bson.M{"_id": users[0].Id}, bson.M{"$pull": bson.M{"interests": "PHP"}}) //从数组中删除一个元素($pull)

	c.FindId(users[0].Id).All(&users)
	log.Println(users)

	c.Remove(bson.M{"name": "JK_CHENG"})//删除


}


type SessionStore struct {
	session *mgo.Session
}

//获取数据库的collection
func (d *SessionStore) C(name string) *mgo.Collection {
	return d.session.DB("howie").C(name)
}

func (d *SessionStore) Db() *mgo.Database {
	return d.session.DB("howie")
}

func NewMgoSession() *SessionStore {
	ds := &SessionStore{
		session: db.Copy(),
	}
	return ds
}

func (d *SessionStore) Close() {
	d.session.Close()
}

func (d *SessionStore) ErrNotFound() error {
	return mgo.ErrNotFound
}

func CloseMgoRedisConnection() {
	db.Close()
}