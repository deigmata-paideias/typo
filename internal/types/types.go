package type

type CommandRecord struct {
	ID            int64     `db:"id" json:"id"`                           // 主键
	WrongCommand  string    `db:"wrong_command" json:"wrong_command"`     // 错误命令
	CorrectedCmd  string    `db:"corrected_cmd" json:"corrected_cmd"`     // 修正命令
	CorrectCount  int       `db:"correct_count" json:"correct_count"`     // 该纠错被使用次数
	FirstOccurAt  time.Time `db:"first_occur_at" json:"first_occur_at"`   // 首次出现时间
	LastOccurAt   time.Time `db:"last_occur_at" json:"last_occur_at"`     // 最后出现时间
}

// CorrectionRule 纠错规则
type CorrectionRule struct {
	ID          int64  `db:"id"`
	Pattern     string `db:"pattern"`     // 匹配模式 (正则)
	Replacement string `db:"replacement"` // 替换为
	Category    string `db:"category"`    // 分类 (git.branch, docker.run等)
	Priority    int    `db:"priority"`    // 优先级
}

// CommandAlias 命令别名
type CommandAlias struct {
	ID      int64  `db:"id"`
	Alias   string `db:"alias"`       // 别名 (br, co等)
	Command string `db:"command"`     // 对应的完整命令 (branch, checkout等)
	Category string `db:"category"`   // 分类 (git.branch等)
	UserID  string `db:"user_id"`     // 用户ID (空为全局)
}

// CommandCategory 命令分类
type CommandCategory struct {
	ID         int64  `db:"id"`
	Category   string `db:"category"`   // 分类名 (git.branch, docker.run等)
	Parent     string `db:"parent"`     // 父命令 (git, docker等)
	SubCommand string `db:"sub_command"` // 子命令 (branch, run等)
}
