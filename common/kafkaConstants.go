package common

/**
*
*author : daozai
*time   : 2017/9/15 21:40
*file   : kafkaConstants.go
*
*/

//KAFKA

const (
	KAFKA_OPERATIONTYPE_TOPIC_LIST = "KAFKA_OPERATIONTYPE_TOPIC_LIST"
	KAFKA_OPERATIONTYPE_TOPIC_INFO = "KAFKA_OPERATIONTYPE_TOPIC_INFO"
	KAFKA_OPERATIONTYPE_SUBSCRIBED_GROUP = "KAFKA_OPERATIONTYPE_SUBSCRIBED_GROUP"
	KAFKA_OPERATIONTYPE_GROUP_LIST = "KAFKA_OPERATIONTYPE_GROUP_LIST"
	KAFKA_OPERATIONTYPE_CREATE_TOPIC = "KAFKA_OPERATIONTYPE_CREATE_TOPIC"
	KAFKA_OPERATIONTYPE_ALTER_TOPIC_PARTITION = "KAFKA_OPERATIONTYPE_ALTER_TOPIC_PARTITION"
	
)

const (
	KAFKA_RUNMODE = "runmode"
)

//RUNMODE

const (
	RUNMODE_DEV            = "dev"
	RUNMODE_TEST           = "test"
	RUNMODE_REGRESS        = "regress"
	RUNMODE_PRE_PRODUCTION = "pre_production"
	RUNMODE_PRODUCTION     = "production"

)