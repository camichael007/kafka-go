package kafka

import (
	"common"
	"github.com/astaxie/beego"
	"errors"
	"strings"
	"fmt"
	"os/exec"
	"kafka/library/remote"
)

/**
*
*author : daozai
*time   : 2017/9/15 8:06
*file   : kafkaTopic.go
*
*/


type KafkaCommand struct {
	Zookeeper         string
	Topic             string
	Broker            string
	ConsumerGroup     string
	ReplicationFactor int
	Partition         int
	OperationType     string

}

func ExecuteKafkaCommand(command KafkaCommand) (string, error) {
	switch command.OperationType {
	case common.KAFKA_OPERATIONTYPE_TOPIC_LIST:
		return execShowTopicList(command.Zookeeper)
	case common.KAFKA_OPERATIONTYPE_TOPIC_INFO:
		return execShowTopicInfo(command.Zookeeper, command.Topic)
	case common.KAFKA_OPERATIONTYPE_SUBSCRIBED_GROUP:
		return execShowSubscribedGroup(command.Broker, command.ConsumerGroup)
	case common.KAFKA_OPERATIONTYPE_CREATE_TOPIC:
		return execCreateTopic(command.Zookeeper, command.ReplicationFactor, command.Partition, command.Topic)
	case common.KAFKA_OPERATIONTYPE_ALTER_TOPIC_PARTITION:
		return execAlterTopicPartition(command.Zookeeper, command.Topic, command.Partition)
	default:
		return "Not Supported command", errors.New("not supported")
	}

}

func execShowTopicList(Zookeeper string) (string, error) {
	command := common.ConcatBySpace("/data/svr/kafka/bin/kafka-topics.sh --list --zookeeper", Zookeeper)
	beego.Info(command)
	return execKafkaCommand(command)
}

func execShowTopicInfo(Zookeeper string, Topic string) (string, error) {
	command := common.ConcatBySpace("/data/svr/kafka/bin/kafka-topics.sh --describe --zookeeper", Zookeeper, "--topic", Topic)
	beego.Info(command)
	result, err := execKafkaCommand(command)
	if len(strings.TrimSpace(result)) == 0 {
		result = "Topic is not existed"
	}
	return result, err
}

func execShowSubscribedGroup(Broker string, ConsumerGroup string) (string, error) {
	command := common.ConcatBySpace("/data/svr/kafka/bin/kafka-consumer-groups.sh --describe --new-consumer  --bootstrap-server", Broker, "--group", ConsumerGroup)
	beego.Info(command)
	return execKafkaCommand(command)
}

func execCreateTopic(Zookeeper string, ReplicationFactor int, Partition int, Topic string) (string, error)  {
	command := common.ConcatBySpace("/data/svr/kafka/bin/kafka-topics.sh --create --zookeeper", Zookeeper, "--replication-factor", fmt.Sprint(ReplicationFactor), "--partitions", fmt.Sprint(Partition), "--topic", Topic)
	beego.Info(command)
	return execKafkaCommand(command)
}

func execAlterTopicPartition(Zookeeper string, Topic string, Partition int) (string, error) {
	command := common.ConcatBySpace("/data/svr/kafka/bin/kafka-topics.sh --alter --zookeeper", Zookeeper, "--topic", Topic, "--partition", Partition)
	beego.Info(command)
	return execKafkaCommand(command)

}

func execKafkaCommand(command string) (string, error) {
	mode := beego.AppConfig.String(common.KAFKA_RUNMODE)
	var (
		b []byte
		error error
	)
	switch mode {
	case common.RUNMODE_DEV:
		session, err := remote.SshConnect("root", "vagrant", "192.168.33.10", 22)
		if err != nil {
			beego.Error(err.Error(), err)
			return err.Error(), err
		}
		b, error = session.CombinedOutput(command)
	default:
		b, error = exec.Command(command).CombinedOutput()
	}
	if error != nil {
		beego.Info(common.Byte2String(b), error)
		return common.Byte2String(b), error
	}
	return  common.Byte2String(b), nil
}
