package Movie

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/abhinav0299/moviesApi/internal/models"
	"reflect"
	"regexp"
	"testing"
)

func TestAdd(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("failed")
	}
	h := New(db)
	testcase := []struct {
		input     models.Movie
		expOut    models.MovieModel
		expErr    error
		mockQuery interface{}
	}{
		{
			input: models.Movie{
				Name:     "abc",
				Genre:    "sus",
				Rating:   4.0,
				Plot:     "good",
				Released: false,
			},
			expOut: models.MovieModel{
				Code:   200,
				Status: "SUCCESS",
				Data: &models.Data{
					Movie: &models.Movie{
						ID:       1,
						Name:     "abc",
						Genre:    "sus",
						Rating:   4.0,
						Plot:     "good",
						Released: false,
					},
				},
				Error: nil,
			},
			mockQuery: mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO moviesTable(name,genre,rating,plot,released) values (?,?,?,?,?)`)).WithArgs("abc", "sus",
				4.0, "good", false).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			input: models.Movie{
				Name:     "abc",
				Genre:    "sus",
				Rating:   4.0,
				Plot:     "good",
				Released: false,
			},
			expErr: errors.New("error"),
			mockQuery: mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO moviesTable(name,genre,rating,plot,released) values (?,?,?,?,?)`)).WithArgs("abc", "sus",
				4.0, "good", false).WillReturnError(errors.New("error")),
		},
	}

	for _, tt := range testcase {
		res, er := h.AddOneMovie(tt.input)
		if er != nil && !reflect.DeepEqual(tt.expErr, er) {
			t.Errorf("failed, %v", er)
		}
		if er == nil && !reflect.DeepEqual(tt.expOut, res) {
			t.Errorf("failed, %v", res.Error.Error())
		}
	}

}
func TestGetBookById(t *testing.T) {
	db, mock, err := sqlmock.New()
	h := New(db)
	if err != nil {
		t.Errorf("failed")
	}
	tcs := []struct {
		expInp    int
		expOut    models.MovieModel
		mockQuery interface{}
		expError  error
	}{
		{
			expInp: 1,
			expOut: models.MovieModel{
				Code:   200,
				Status: "SUCCESS",
				Data: &models.Data{Movie: &models.Movie{
					ID:       1,
					Name:     "abc",
					Genre:    "sus",
					Rating:   4.8,
					Plot:     "Plot",
					Released: false,
				}},
				Error: nil,
			},
			expError: nil,
			mockQuery: mock.ExpectQuery(regexp.QuoteMeta(`select * from moviesTable where id = ?`)).WithArgs(1).
				WillReturnRows(sqlmock.NewRows([]string{"id", "name", "genre", "rating", "plot", "released"}).
					AddRow(1, "abc", "sus", 4.8, "Plot", false),
				),
		},
		{
			expInp: 2,
			expOut: models.MovieModel{
				Error: errors.New("db error"),
			},
			expError: errors.New("db error"),
			mockQuery: mock.ExpectQuery(regexp.QuoteMeta(`select * from moviesTable where id = ?`)).WithArgs(2).
				WillReturnRows(sqlmock.NewRows([]string{"id", "name", "genre", "rating", "plot", "released"}).
					AddRow(2, "abc", "sus", 4.8, "Plot", false),
				).WillReturnError(errors.New("db error")),
		},
	}
	for _, tt := range tcs {
		res, err := h.GetBookById(tt.expInp)
		if err != nil && !reflect.DeepEqual(tt.expError, err) {
			t.Errorf("Expected %v Got %v", err, tt.expError)
		} else {
			if !reflect.DeepEqual(tt.expOut, res) {
				t.Errorf("Expected %v Got %v", *res.Data, *tt.expOut.Data)
			}
		}
	}
}
