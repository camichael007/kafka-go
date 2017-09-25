package routers

import (
	"github.com/astaxie/beego"
	"kafka/controllers/kafka"
)

func init() {
    beego.Router("/", &kafka.KafkaController{}, "*:CreateTopic")
    beego.Router("/kafka/topiclist", &kafka.KafkaController{}, "get:ShowTopicList")
    beego.Router("/kafka/createtopic", &kafka.KafkaController{}, "get,post:CreateTopic")
}
