/**
 * @Author: lzw5399
 * @Date: 2021/3/18 22:00
 * @Desc: 审批/处理流程实例的接口的请求体
 */
package request

// 审批/处理流程实例的接口的请求体
type HandleInstancesRequest struct {
	EdgeID            string `json:"edgeId" form:"edgeId"`                       // 走的流程的id
	ProcessInstanceId uint   `json:"processInstanceId" form:"processInstanceId"` // 流程实例的id
	Remarks           string `json:"remarks" form:"remarks"`                     // 备注
}