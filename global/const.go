package global

const CONFIG_DEFAULT = "config.yaml"
const MINI_WECHAT = "micromessenger"
const RedisURL_CACHE = 7200
const DEFAULT_BOOK_ID = 0
const DEFAULT_NUM = 1
const DEFAULT_PAGE = 1
const DEFAULT_LEVEL = 1
const DEFAULT_PAGE_SIZE = 10
const QUEUE_LEN = 3
const DEFAULT_QUEUE = "queue_list_question"
const QUEUE = "queue_list_question_%d"
const AUTH_TOKEN = "wechat"
const AppId = "wx2afb8412b255e4fe"
const AppSecret = "82c58b3027d4e09230a06f8132b04577"

const (
	SUCCESS                = 10000
	FAIL                   = 10001
	FORBID                 = 403
	ERR_RES_PARAMS_ILLEGAL = 10002
)

const (
	SUCCESS_MSG                = "success"
	FAIL_MSG                   = "fail"
	FORBID_MSG                 = "非法请求，禁止访问"
	ERR_RES_PARAMS_ILLEGAL_MSG = "参数传递不合法"
)