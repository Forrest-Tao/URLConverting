package sequence

import (
	"database/sql"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const sqkReplaceIntoStub = `REPLACE INTO sequence (stub) VALUES ('a')`

type Mysql struct {
	conn sqlx.SqlConn
}

func NewMysql(DSN string) Sequence {
	return &Mysql{
		conn: sqlx.NewMysql(DSN),
	}
}

func (m *Mysql) Next() (seq uint64, err error) {
	var stmt sqlx.StmtSession
	stmt, err = m.conn.Prepare(sqkReplaceIntoStub)
	if err != nil {
		logx.Errorw("conn.Prepare failed", logx.Field("err", err.Error()))
		return
	}
	defer stmt.Close()

	// 执行
	var rest sql.Result
	rest, err = stmt.Exec()
	if err != nil {
		logx.Errorw("stmt.Exec failed", logx.Field("err", err.Error()))
		return
	}
	var lid int64
	lid, err = rest.LastInsertId()
	if err != nil {
		logx.Errorw("rest.LastInsertId failed", logx.Field("err", err.Error()))
		return
	}
	return uint64(lid), nil
}
