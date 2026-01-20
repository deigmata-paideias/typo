package type

type CommandRecord struct {
	ID            int64     `db:"id" json:"id"`                           // 主键
	WrongCommand  string    `db:"wrong_command" json:"wrong_command"`     // 错误命令
	CorrectedCmd  string    `db:"corrected_cmd" json:"corrected_cmd"`     // 修正命令
	CorrectCount  int       `db:"correct_count" json:"correct_count"`     // 该纠错被使用次数
	FirstOccurAt  time.Time `db:"first_occur_at" json:"first_occur_at"`   // 首次出现时间
	LastOccurAt   time.Time `db:"last_occur_at" json:"last_occur_at"`     // 最后出现时间
}
