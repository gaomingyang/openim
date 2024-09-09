package define

const SUCCESSCODE = 200

// 用户状态
const USER_STATUS_NORMAL = 1 // 只对比这个是否正常

// const USER_STATUS_FORBIDDEN = 0

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	TraceID string      `json:"trace_id"`
	Data    interface{} `json:"data,omitempty"`
}
