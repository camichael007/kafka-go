package kafka

import (
	"kafka/controllers"
	"kafka/library/kafka"
	"common"
	"github.com/astaxie/beego"
)

/**
*
*author : daozai
*time   : 2017/9/15 7:51
*file   : kafkaController.go
*
*/


type KafkaController struct {
	controllers.BaseController
}

func (c *KafkaController) ShowTopicList() {
	if c.IsGet() {
		zookeeper := c.GetString("zookeeper")
		beego.Info("zookeeper:" + zookeeper)
		command := kafka.KafkaCommand{Zookeeper:zookeeper, OperationType:common.KAFKA_OPERATIONTYPE_TOPIC_LIST}
		result, _ := kafka.ExecuteKafkaCommand(command)
		c.Data["result"] = &result
		c.TplName = "getKafkaTopicList.tpl"

	}
}

func (c *KafkaController) CreateTopic() {
	if c.IsGet() {
		c.TplName = "getKafkaTopicList.tpl"

	}

	if c.IsPost() {
		zookeeper := c.GetString("zookeeper", "192.168.33.10")
		topic := c.GetString("topic", "test")
		partition, _ := c.GetInt("partition", 1)
		replicationfactor, _ := c.GetInt("replicationfactor", 1)
		beego.Info("zookeeper:" , zookeeper , "topic" , topic , "partition", partition, "replicationfactor", replicationfactor)
		command := kafka.KafkaCommand{Zookeeper:zookeeper, Topic:topic, Partition:partition, ReplicationFactor:replicationfactor, OperationType:common.KAFKA_OPERATIONTYPE_CREATE_TOPIC}
		result, _ := kafka.ExecuteKafkaCommand(command)
		c.Data["result"] = &result
		c.TplName = "getKafkaTopicList.tpl"

	}
}