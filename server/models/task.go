package models

import (
	//"github.com/go-xorm/xorm"
	"time"
)

type TaskLevel int8
type Status int8
type CommonMap map[string]interface{}
type TaskProtocol int8

const (
	Disabled Status = 0 // 禁用
	Failure  Status = 0 // 失败
	Enabled  Status = 1 // 启用
	Running  Status = 1 // 运行中
	Finish   Status = 2 // 完成
	Cancel   Status = 3 // 取消
)

const (
	TaskHTTP TaskProtocol = iota + 1 // HTTP协议
	TaskGRPC                         // GRPC方式执行命令
)

const (
	TaskLevelParent TaskLevel = 1 // 父任务
	TaskLevelChild  TaskLevel = 2 // 子任务
)

type Task struct {
	Id               int       `json:"id" xorm:"int pk autoincr"`
	Name             string    `json:"name" xorm:"varchar(32) notnull"`
	Level            TaskLevel `json:"level" xorm:"tinyint notnull index default 1"`
	DependencyTaskId []string  `json:"dependency_task_id" xorm:"varchar(64) notnull default ''"`
	Remark           string    // 备注
	Status           Status
	Protocol         TaskProtocol
	Parameter        interface{} // 参数
	Created          time.Time   `json:"created" xorm:"datetime notnull created"` // 创建时间
	Deleted          time.Time   `json:"deleted" xorm:"datetime deleted"`         // 删除时间
	Crontab          string      // crontab 字符串
}

// 更新
func (t *Task) Create() (int, error) {
	_, err := Db.Insert(t)
	if err == nil {
		return t.Id, nil
	}
	return 0, err
}

// 删除
func (t *Task) Delete(id int) (int64, error) {
	return Db.Id(id).Delete(t)
}

// 更新
func (t *Task) Update(id int, data CommonMap) (int64, error) {
	return Db.Table(t).ID(id).Update(data)
}

// 激活
func (t *Task) Enable(id int) (int64, error) {
	return t.Update(id, CommonMap{"status": Enabled})
}

// 禁用
func (t *Task) Disable(id int) (int64, error) {
	return t.Update(id, CommonMap{"status": Disabled})
}
