package rds_test

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/keitax/airlog/internal/domain"
	"github.com/keitax/airlog/internal/infrastructure/rds"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("PostRepository", func() {
	var (
		postrepo *rds.PostRepository
		db       *sql.DB
	)

	BeforeEach(func() {
		db_, err := sql.Open("mysql", "root@tcp(localhost:3306)/blog")
		if err != nil {
			panic(err)
		}
		db = db_
		postrepo = &rds.PostRepository{
			DB: db_,
		}
		if _, err := db.Exec(`delete from post`); err != nil {
			panic(err)
		}
	})

	Describe("Filename()", func() {
		Context("when some posts are saved", func() {
			BeforeEach(func() {
				for _, rec := range [][]interface{}{
					{"20190101-post.md", "2019-01-01 00:00:00", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "Title1", "Body1"},
					{"20190102-post.md", "2019-01-02 00:00:00", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "Title2", "Body2"},
					{"20190103-post.md", "2019-01-03 00:00:00", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "Title3", "Body3"},
				} {
					if _, err := db.Exec(`insert into post (filename, timestamp, hash, title, body) values (?, ?, ?, ?, ?)`, rec...); err != nil {
						panic(err)
					}
				}

			})

			It("selects the specify post by a filename", func() {
				got, err := postrepo.Filename("20190102-post.md")
				Expect(err).NotTo(HaveOccurred())
				Expect(got).To(Equal(&domain.Post{
					Filename:  "20190102-post.md",
					Timestamp: Time("2019-01-02 00:00:00"),
					Hash:      "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
					Title:     "Title2",
					Body:      "Body2",
				}))
			})
		})

		Context("when no posts are saved", func() {
			It("occurs a not found error", func() {
				got, err := postrepo.Filename("20190101-post.md")
				Expect(err).To(Equal(domain.ErrNotFound{}))
				Expect(got).To(BeNil())
			})
		})
	})

	Describe("All()", func() {
		Context("when three posts are saved", func() {
			BeforeEach(func() {
				for _, rec := range [][]interface{}{
					{"20190101-post.md", "2019-01-01 00:00:00", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "Title1", "Body1"},
					{"20190102-post.md", "2019-01-02 00:00:00", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "Title2", "Body2"},
					{"20190103-post.md", "2019-01-03 00:00:00", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "Title3", "Body3"},
				} {
					if _, err := db.Exec(`insert into post (filename, timestamp, hash, title, body) values (?, ?, ?, ?, ?)`, rec...); err != nil {
						panic(err)
					}
				}

			})

			It("selects all the posts ordered by timestamp", func() {
				got, err := postrepo.All()
				Expect(err).NotTo(HaveOccurred())
				Expect(got).To(HaveLen(3))
				Expect(got[0]).To(Equal(&domain.Post{
					Filename:  "20190103-post.md",
					Timestamp: Time("2019-01-03 00:00:00"),
					Hash:      "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
					Title:     "Title3",
					Body:      "Body3",
				}))
				Expect(got[len(got)-1]).To(Equal(&domain.Post{
					Filename:  "20190101-post.md",
					Timestamp: Time("2019-01-01 00:00:00"),
					Hash:      "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
					Title:     "Title1",
					Body:      "Body1",
				}))
			})
		})
	})
})

func Time(text string) time.Time {
	t, err := time.Parse("2006-01-02 15:04:05", text)
	if err != nil {
		panic(err)
	}
	return t
}
