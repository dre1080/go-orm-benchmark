package benchs

import (
	"fmt"

	pg "github.com/go-pg/pg/v10"
)

var pgdb *pg.DB

func init() {
	st := NewSuite("pg")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, PgInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, PgInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, PgUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, PgRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, PgReadSlice)

		opts, _ := pg.ParseURL(ORM_SOURCE)
		pgdb = pg.Connect(opts)
	}
}

func PgInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.ID = 0
		if err := pgdb.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PgInsertMulti(b *B) {
	var ms []*Model
	wrapExecute(b, func() {
		initDB()
	})

	for i := 0; i < b.N; i++ {
		ms = make([]*Model, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewModel())
		}
		if err := pgdb.Insert(&ms); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PgUpdate(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		if err := pgdb.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		m.Age = 20
		if err := pgdb.Update(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PgRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		if err := pgdb.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if err := pgdb.Select(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PgReadSlice(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.ID = 0
			if err := pgdb.Insert(m); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model
		if err := pgdb.Model(&models).Where("id > ?", 0).Limit(100).Select(); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
