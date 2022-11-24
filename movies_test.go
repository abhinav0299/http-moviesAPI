package main

//func TestPost(t *testing.T) {
//	var tcs = []struct {
//		desc   string
//		body   Movie
//		ExpOut MovieModel
//	}{
//		{desc: "hollywood", body: models.Movie{Name: "Titanic", Genre: "Thriller", Rating: 4.2, Plot: "Good", Released: true}, ExpOut: MovieModel{Code: 200, Status: "SUCCESS", Data: &Data{Movie: &Movie{
//			ID:       1,
//			Name:     "Titanic",
//			Genre:    "Thriller",
//			Rating:   4.2,
//			Plot:     "Good",
//			Released: true,
//		}}}},
//		{desc: "bollywood", body: Movie{Name: "Gangs of Wasseypur", Genre: "Thriller", Rating: 4.6, Plot: "Good", Released: true}, ExpOut: MovieModel{Code: 200, Status: "SUCCESS", Data: &Data{Movie: &Movie{
//			ID:       1,
//			Name:     "Gangs of Wasseypur",
//			Genre:    "Thriller",
//			Rating:   4.6,
//			Plot:     "Good",
//			Released: true,
//		}}}},
//	}
//	for _, tt := range tcs {
//		finalJson, _ := json.Marshal(tt.body)
//		r := httptest.NewRequest("Post", "/movies", bytes.NewReader(finalJson))
//		w := httptest.NewRecorder()
//		AddOneMovie(w, r)
//		res := w.Result()
//		data, err := io.ReadAll(res.Body)
//		if err != nil {
//			t.Errorf("expected error is nil got this %v", err)
//		}
//		parsed, err := json.Marshal(tt.ExpOut)
//
//		if err != nil {
//			t.Errorf("expected error is nil got this %v", err)
//		}
//		if string(parsed) != string(data[:len(data)-1]) {
//			t.Errorf("expected %v but got %v", string(parsed), string(data[:len(data)-1]))
//		} else {
//			fmt.Println("Done")
//		}
//
//	}
//}

//func TestGetById(t *testing.T) {
//	var tcs = []struct {
//		target         string
//		expectedError  error
//		expectedOutput MovieModel
//	}{
//		{
//			target: "/movies",
//			expectedOutput: MovieModel{
//				Error: errors.New("No movie Found with given id"),
//			},
//		},
//		{
//			//{ID: 1, Name: "Abhinav", Genre: "Comedy", Rating: 5.0, Plot: "Secret", Released: true},
//			target: "/movies",
//			expectedOutput: MovieModel{
//				Code:   200,
//				Status: "SUCCESS",
//				Data: &Data{&Movie{
//					ID:       1,
//					Name:     "Abhinav",
//					Genre:    "Comedy",
//					Rating:   5.0,
//					Plot:     "Secret",
//					Released: true,
//				}},
//				Error: nil,
//			},
//		},
//	}
//	for index, tt := range tcs {
//		r := httptest.NewRequest("GET", tt.target, nil)
//		w := httptest.NewRecorder()
//		params := map[string]string{
//			"id": "1",
//		}
//		if index == 0 {
//			params["id"] = "$"
//		}
//		r = mux.SetURLVars(r, params)
//		getMovieById(w, r)
//		var response MovieModel
//		res := w.Result()
//		data, err := io.ReadAll(res.Body)
//		if err != nil {
//			t.Errorf("Wrong output")
//		}
//		err = json.Unmarshal(data, &response)
//		if err != nil {
//			t.Errorf("Wrong output")
//		}
//		if !reflect.DeepEqual(response, tt.expectedOutput) {
//			t.Errorf("test Case Failed")
//		}
//	}
//}

//func TestUpdateById(t *testing.T) {
//
//	var tcs = []struct {
//		desc   string
//		body   Movie
//		ExpOut MovieModel
//	}{
//		{desc: "hollywood", body: Movie{Rating: 2, Plot: "Good", Released: true}, ExpOut: MovieModel{Code: 200, Status: "SUCCESS", Data: &Data{Movie: &Movie{
//			ID:       1,
//			Name:     "Titanic",
//			Genre:    "Thriller",
//			Rating:   2,
//			Plot:     "Good",
//			Released: true,
//		}}}},
//		{desc: "bollywood", body: Movie{Rating: 4.6, Plot: "Good", Released: true}, ExpOut: MovieModel{Code: 200, Status: "SUCCESS", Data: &Data{Movie: &Movie{
//			ID:       1,
//			Name:     "Titanic",
//			Genre:    "Thriller",
//			Rating:   4.6,
//			Plot:     "Good",
//			Released: true,
//		}}}},
//	}
//	for _, tt := range tcs {
//		finalJson, _ := json.Marshal(tt.body)
//		r := httptest.NewRequest("PUT", "/movies/1", bytes.NewReader(finalJson))
//		w := httptest.NewRecorder()
//		params := map[string]string{
//			"id": "1",
//		}
//
//		r = mux.SetURLVars(r, params)
//		updateOneMovie(w, r)
//		res := w.Result()
//		data, err := io.ReadAll(res.Body)
//		if err != nil {
//			t.Errorf("expected error is nil got this %v", err)
//		}
//		parsed, err := json.Marshal(tt.ExpOut)
//
//		if err != nil {
//			t.Errorf("expected error is nil got this %v", err)
//		}
//		if string(parsed) != string(data[:len(data)-1]) {
//			t.Errorf("expected %v but got %v", string(parsed), string(data[:len(data)-1]))
//		} else {
//			fmt.Println("Done")
//		}
//
//	}
//}
