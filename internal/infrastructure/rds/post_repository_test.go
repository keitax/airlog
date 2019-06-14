package rds_test

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/keitam913/airlog/internal/domain"
	"github.com/keitam913/airlog/internal/infrastructure/rds"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PostRepository", func() {
	var (
		postrepo *rds.PostRepository
		db       *sql.DB
	)

	BeforeEach(func() {
		db_, err := sql.Open("mysql", "root:root@tcp(db:3306)/blog")
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
		if _, err := db.Exec(`delete from post_label`); err != nil {
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

				for _, rec := range [][]interface{}{
					{"20190101-post.md", "label-1"},
					{"20190102-post.md", "label-2"},
					{"20190103-post.md", "label-3"},
				} {
					if _, err := db.Exec(`insert into post_label (filename, label) values (?, ?)`, rec...); err != nil {
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
					Labels:    []string{"label-2"},
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

				for _, rec := range [][]interface{}{
					{"20190101-post.md", "label-1"},
					{"20190102-post.md", "label-2"},
					{"20190103-post.md", "label-3"},
				} {
					if _, err := db.Exec(`insert into post_label (filename, label) values (?, ?)`, rec...); err != nil {
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
					Labels:    []string{"label-3"},
				}))
				Expect(got[len(got)-1]).To(Equal(&domain.Post{
					Filename:  "20190101-post.md",
					Timestamp: Time("2019-01-01 00:00:00"),
					Hash:      "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
					Title:     "Title1",
					Body:      "Body1",
					Labels:    []string{"label-1"},
				}))
			})
		})
	})

	Describe("Put()", func() {
		var (
			post *domain.Post
			err  error
		)

		JustBeforeEach(func() {
			err = postrepo.Put(post)
		})

		Context("when a post is given", func() {
			BeforeEach(func() {
				post = &domain.Post{
					Filename:  "20190101-post.md",
					Timestamp: Time("2019-01-01 00:00:00"),
					Title:     "Title",
					Body:      "hello world",
					Labels:    []string{"label-0", "label-1"},
				}
			})

			It("inserts a post record", func() {
				Expect(err).NotTo(HaveOccurred())
				var rs *sql.Rows
				rs, err = db.Query("select filename, timestamp, title, body from post")
				if err != nil {
					panic(err)
				}
				defer rs.Close()
				Expect(rs.Next()).To(BeTrue())
				var filename, timestamp, title, body string
				if err := rs.Scan(&filename, &timestamp, &title, &body); err != nil {
					panic(err)
				}
				Expect(filename).To(Equal("20190101-post.md"))
				Expect(timestamp).To(Equal("2019-01-01 00:00:00"))
				Expect(title).To(Equal("Title"))
				Expect(body).To(Equal("hello world"))
			})

			It("inserts label records", func() {
				Expect(err).NotTo(HaveOccurred())
				rs, err := db.Query("select label from post_label where filename = ?", "20190101-post.md")
				if err != nil {
					panic(err)
				}
				defer rs.Close()
				var labels []string
				for rs.Next() {
					var label string
					if err := rs.Scan(&label); err != nil {
						panic(err)
					}
					labels = append(labels, label)
				}
				Expect(labels).To(ContainElement("label-0"))
				Expect(labels).To(ContainElement("label-1"))
			})
		})

		Context("when a post is inserted", func() {
			BeforeEach(func() {
				for _, rec := range [][]interface{}{
					{"20190101-post.md", "2019-01-01 00:00:00", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "Original Title", "orignal body"},
				} {
					if _, err := db.Exec(`insert into post (filename, timestamp, hash, title, body) values (?, ?, ?, ?, ?)`, rec...); err != nil {
						panic(err)
					}
				}
			})

			Context("given a post file whose filename is same", func() {
				BeforeEach(func() {
					post = &domain.Post{
						Filename:  "20190101-post.md",
						Timestamp: Time("2019-01-01 00:00:00"),
						Title:     "Changed Title",
						Body:      "changed body",
					}
				})

				It("updates the post record", func() {
					Expect(err).NotTo(HaveOccurred())
					var rs *sql.Rows
					rs, err = db.Query("select filename, timestamp, title, body from post")
					if err != nil {
						panic(err)
					}
					defer rs.Close()
					Expect(rs.Next()).To(BeTrue())
					var filename, timestamp, title, body string
					if err := rs.Scan(&filename, &timestamp, &title, &body); err != nil {
						panic(err)
					}
					Expect(filename).To(Equal("20190101-post.md"))
					Expect(timestamp).To(Equal("2019-01-01 00:00:00"))
					Expect(title).To(Equal("Changed Title"))
					Expect(body).To(Equal("changed body"))
				})
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
